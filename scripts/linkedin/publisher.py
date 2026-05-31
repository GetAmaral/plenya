#!/usr/bin/env python3
"""
LinkedIn cron publisher — Dr. Getúlio pipeline.

Runs every 15 min via WSL cron. For each post in queue.yaml with
status=approved AND scheduled_at within the firing window, publishes
via LinkedIn API and updates queue state.

Idempotent: posts already published are skipped. Late-firing window of
60 min protects against cron misses.

Logs to scripts/linkedin/publisher.log (rotated by user's logrotate or manually).
"""
from __future__ import annotations

import json
import logging
import os
import sys
import time
import urllib.parse
import urllib.request
import urllib.error
from datetime import datetime, timezone, timedelta
from pathlib import Path

import yaml

BASE = Path(__file__).parent
REPO = BASE.parent.parent  # /home/user/plenya
QUEUE_PATH = BASE / "queue.yaml"
ENV_PATH = BASE / ".env"
LOG_PATH = BASE / "publisher.log"
LINKEDIN_VERSION = "202604"
ONTIME_WINDOW = timedelta(minutes=60)   # considered "on time"
CATCHUP_WINDOW = timedelta(hours=12)    # catch-up: recover posts missed by cron gaps (machine off)
TOKEN_WARN_DAYS = 7

logging.basicConfig(
    filename=LOG_PATH,
    level=logging.INFO,
    format="%(asctime)s [%(levelname)s] %(message)s",
)
log = logging.getLogger("linkedin-publisher")


# ─── ENV ──────────────────────────────────────────────────────────────────────

def load_env() -> dict[str, str]:
    env = {}
    for line in ENV_PATH.read_text().splitlines():
        line = line.strip()
        if not line or line.startswith("#") or "=" not in line:
            continue
        k, v = line.split("=", 1)
        env[k] = v
    return env


# ─── HTTP ─────────────────────────────────────────────────────────────────────

def li_post(url: str, body: dict, token: str, extra_headers: dict | None = None) -> dict:
    headers = {
        "Authorization": f"Bearer {token}",
        "Content-Type": "application/json",
        "LinkedIn-Version": LINKEDIN_VERSION,
        "X-Restli-Protocol-Version": "2.0.0",
    }
    if extra_headers:
        headers.update(extra_headers)
    req = urllib.request.Request(url, data=json.dumps(body).encode(), method="POST", headers=headers)
    try:
        with urllib.request.urlopen(req, timeout=30) as r:
            return {"ok": True, "status": r.status, "headers": dict(r.headers), "body": r.read().decode()}
    except urllib.error.HTTPError as e:
        return {"ok": False, "status": e.code, "body": e.read().decode(errors="ignore")}


def li_put_binary(url: str, data: bytes, token: str, content_type: str) -> dict:
    req = urllib.request.Request(url, data=data, method="PUT",
                                   headers={"Authorization": f"Bearer {token}", "Content-Type": content_type})
    try:
        with urllib.request.urlopen(req, timeout=120) as r:
            return {"ok": True, "status": r.status}
    except urllib.error.HTTPError as e:
        return {"ok": False, "status": e.code, "body": e.read().decode(errors="ignore")}


# ─── IMAGE UPLOAD ─────────────────────────────────────────────────────────────

def upload_image(image_path: Path, author: str, token: str) -> str | None:
    """Returns image URN on success, None on failure (caller may choose to post text-only)."""
    if not image_path or not image_path.exists():
        return None
    # LinkedIn accepts JPG/PNG. Convert WebP if needed.
    if image_path.suffix.lower() == ".webp":
        jpg = Path(f"/tmp/li-img-{int(time.time())}.jpg")
        os.system(f"convert {image_path.as_posix()!r} -quality 92 {jpg.as_posix()!r}")
        if not jpg.exists():
            log.error("WebP→JPG conversion failed for %s", image_path)
            return None
        image_path = jpg
        content_type = "image/jpeg"
    elif image_path.suffix.lower() in (".jpg", ".jpeg"):
        content_type = "image/jpeg"
    elif image_path.suffix.lower() == ".png":
        content_type = "image/png"
    else:
        log.error("Unsupported image format: %s", image_path)
        return None

    init = li_post("https://api.linkedin.com/rest/images?action=initializeUpload",
                   {"initializeUploadRequest": {"owner": author}}, token)
    if not init["ok"]:
        log.error("image initializeUpload failed (%s): %s", init["status"], init.get("body"))
        return None
    val = json.loads(init["body"])["value"]
    put = li_put_binary(val["uploadUrl"], image_path.read_bytes(), token, content_type)
    if not put["ok"]:
        log.error("image PUT failed (%s): %s", put["status"], put.get("body"))
        return None
    return val["image"]


# ─── POST CREATION ────────────────────────────────────────────────────────────

def publish_post(entry: dict, author: str, token: str) -> tuple[bool, str | None, str | None]:
    """Returns (ok, post_urn, error)."""
    body = {
        "author": author,
        "commentary": entry["commentary"].rstrip(),
        "visibility": entry.get("visibility", "PUBLIC"),
        "distribution": {
            "feedDistribution": "MAIN_FEED",
            "targetEntities": [],
            "thirdPartyDistributionChannels": [],
        },
        "lifecycleState": "PUBLISHED",
        "isReshareDisabledByAuthor": False,
    }

    image_urn = None
    if entry.get("image_path"):
        image_full = REPO / entry["image_path"]
        image_urn = upload_image(image_full, author, token)
        if image_urn:
            body["content"] = {
                "media": {
                    "title": entry.get("image_title", ""),
                    "id": image_urn,
                }
            }
        else:
            log.warning("Posting text-only (image upload failed) for %s", entry["slug"])

    r = li_post("https://api.linkedin.com/rest/posts", body, token)
    if not r["ok"]:
        return False, None, f"{r['status']}: {r.get('body','')[:500]}"
    post_urn = r["headers"].get("x-restli-id") or r["headers"].get("X-RestLi-Id")
    return True, post_urn, None


# ─── QUEUE STATE ──────────────────────────────────────────────────────────────

def parse_iso(s) -> datetime:
    # accept datetime (YAML auto-parses unquoted timestamps),
    # "YYYY-MM-DD HH:MM:SS-03:00", or "YYYY-MM-DDTHH:MM:SS-03:00"
    if isinstance(s, datetime):
        return s
    s = s.replace(" ", "T", 1) if "T" not in s else s
    return datetime.fromisoformat(s)


def in_fire_window(scheduled_at: datetime, now: datetime) -> bool:
    # Catch-up: fire any post whose scheduled time has passed, up to CATCHUP_WINDOW later.
    # Tolerates cron gaps (machine off/asleep) within the same workday without firing
    # a morning post in the middle of the night.
    return scheduled_at <= now <= scheduled_at + CATCHUP_WINDOW


def main() -> int:
    if not ENV_PATH.exists():
        log.error("missing .env at %s", ENV_PATH); return 1
    if not QUEUE_PATH.exists():
        log.error("missing queue.yaml"); return 1

    env = load_env()
    token = env.get("LINKEDIN_ACCESS_TOKEN")
    author = env.get("LINKEDIN_AUTHOR_URN")
    if not token or not author:
        log.error("missing LINKEDIN_ACCESS_TOKEN or LINKEDIN_AUTHOR_URN in .env"); return 1

    # Token expiry check
    exp = int(env.get("LINKEDIN_TOKEN_EXPIRES_AT", "0"))
    days_left = (exp - int(time.time())) // 86400 if exp else None
    if days_left is not None and days_left <= TOKEN_WARN_DAYS:
        log.warning("⚠ token expires in %d days — re-run OAuth flow soon", days_left)

    queue = yaml.safe_load(QUEUE_PATH.read_text())
    posts = queue.get("posts", [])
    now = datetime.now(timezone(timedelta(hours=-3)))  # BRT

    fired = 0
    for entry in posts:
        if entry.get("status") != "approved":
            continue
        sched_raw = entry.get("scheduled_at")
        if not sched_raw:
            continue
        sched = parse_iso(sched_raw)
        if not in_fire_window(sched, now):
            continue

        slug = entry["slug"]
        late = now > sched + ONTIME_WINDOW
        if late:
            mins = int((now - sched).total_seconds() // 60)
            log.warning("CATCH-UP publishing %s (scheduled %s, %d min late — cron gap?)", slug, sched_raw, mins)
        else:
            log.info("publishing %s (scheduled %s)", slug, sched_raw)
        ok, post_urn, err = publish_post(entry, author, token)
        if ok:
            entry["status"] = "published"
            entry["published_at"] = now.isoformat(timespec="seconds")
            entry["post_urn"] = post_urn
            log.info("  ✓ published %s → %s", slug, post_urn)
        else:
            entry["status"] = "failed"
            entry["failed_at"] = now.isoformat(timespec="seconds")
            entry["error"] = err
            log.error("  ✗ failed %s: %s", slug, err)
        fired += 1

    if fired:
        # Write back queue (preserve YAML structure as best as possible)
        QUEUE_PATH.write_text(yaml.safe_dump(queue, allow_unicode=True, sort_keys=False, default_flow_style=False, width=120))
        log.info("queue updated (%d entries fired)", fired)
    else:
        log.info("no posts in fire window — nothing to do")

    return 0


if __name__ == "__main__":
    sys.exit(main())
