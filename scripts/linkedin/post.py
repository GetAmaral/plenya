#!/usr/bin/env python3
"""
Post text/article to LinkedIn personal profile via Posts API.

Usage:
  python3 post.py "Texto do post" [--visibility PUBLIC|CONNECTIONS]
  python3 post.py --stdin                       # read text from stdin
  python3 post.py --file post.md                # read text from file

Uses LinkedIn-Version: 202604 (current as of May 2026) to avoid the
deprecated-version bug that breaks Composio's wrapper.
"""
import argparse
import json
import os
import sys
import time
import urllib.request
import urllib.error
from pathlib import Path

ENV_PATH = Path(__file__).parent / ".env"
LINKEDIN_VERSION = "202604"
POSTS_API = "https://api.linkedin.com/rest/posts"


def load_env() -> dict:
    env = {}
    for line in ENV_PATH.read_text().splitlines():
        line = line.strip()
        if not line or line.startswith("#") or "=" not in line:
            continue
        k, v = line.split("=", 1)
        env[k] = v
    return env


def check_token_expiry(env: dict) -> None:
    exp = int(env.get("LINKEDIN_TOKEN_EXPIRES_AT", "0"))
    if exp and exp - int(time.time()) < 86400 * 7:
        days = (exp - int(time.time())) // 86400
        print(f"⚠️  Token expira em {days} dias — renovar logo.", file=sys.stderr)


def create_post(env: dict, commentary: str, visibility: str = "PUBLIC") -> dict:
    payload = {
        "author": env["LINKEDIN_AUTHOR_URN"],
        "commentary": commentary,
        "visibility": visibility,
        "distribution": {
            "feedDistribution": "MAIN_FEED",
            "targetEntities": [],
            "thirdPartyDistributionChannels": [],
        },
        "lifecycleState": "PUBLISHED",
        "isReshareDisabledByAuthor": False,
    }
    req = urllib.request.Request(
        POSTS_API,
        data=json.dumps(payload).encode(),
        method="POST",
        headers={
            "Authorization": f"Bearer {env['LINKEDIN_ACCESS_TOKEN']}",
            "Content-Type": "application/json",
            "LinkedIn-Version": LINKEDIN_VERSION,
            "X-Restli-Protocol-Version": "2.0.0",
        },
    )
    try:
        with urllib.request.urlopen(req, timeout=20) as r:
            post_id = r.headers.get("x-restli-id")
            return {"ok": True, "post_id": post_id, "status": r.status}
    except urllib.error.HTTPError as e:
        body = e.read().decode(errors="ignore")
        return {"ok": False, "status": e.code, "error": body}


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("text", nargs="?", help="Post commentary text")
    parser.add_argument("--stdin", action="store_true", help="Read text from stdin")
    parser.add_argument("--file", help="Read text from file")
    parser.add_argument(
        "--visibility",
        default="PUBLIC",
        choices=["PUBLIC", "CONNECTIONS", "LOGGED_IN"],
    )
    args = parser.parse_args()

    if args.stdin:
        commentary = sys.stdin.read()
    elif args.file:
        commentary = Path(args.file).read_text()
    elif args.text:
        commentary = args.text
    else:
        parser.error("provide text, --stdin, or --file")

    commentary = commentary.strip()
    if not commentary:
        parser.error("empty text")

    env = load_env()
    check_token_expiry(env)

    result = create_post(env, commentary, args.visibility)
    if result["ok"]:
        post_id = result["post_id"]
        urn_part = post_id.replace("urn:li:share:", "").replace("urn:li:ugcPost:", "")
        url = f"https://www.linkedin.com/feed/update/{post_id}/"
        print(f"✓ Posted: {post_id}")
        print(f"  URL:    {url}")
        sys.exit(0)
    else:
        print(f"✗ FAILED ({result['status']}):", file=sys.stderr)
        print(result["error"], file=sys.stderr)
        sys.exit(1)


if __name__ == "__main__":
    main()
