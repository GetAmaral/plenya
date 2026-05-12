"""HTTP client for Meta Graph API with auto-refresh of long-lived IG tokens."""

import sys
from datetime import datetime, timezone
from typing import Any

import httpx

from . import config, storage


class MetaError(RuntimeError):
    def __init__(self, status: int, body: Any):
        super().__init__(f"Meta API {status}: {body}")
        self.status = status
        self.body = body


def _now() -> datetime:
    return datetime.now(timezone.utc)


def _seconds_until_expiry(record: dict) -> int:
    expires_at = datetime.fromisoformat(record["expires_at"])
    return int((expires_at - _now()).total_seconds())


def _refresh_if_needed(record: dict) -> dict:
    seconds_left = _seconds_until_expiry(record)
    if seconds_left > config.REFRESH_WINDOW_SECONDS:
        return record
    print(f"[plenya-social-mcp] token expires in {seconds_left}s, refreshing...", file=sys.stderr)
    r = httpx.get(
        config.IG_REFRESH_URL,
        params={"grant_type": "ig_refresh_token", "access_token": record["access_token"]},
        timeout=30,
    )
    if r.status_code >= 400:
        print(f"[plenya-social-mcp] refresh failed: {r.status_code} {r.text}", file=sys.stderr)
        return record
    payload = r.json()
    new_expires_in = int(payload.get("expires_in", record["expires_in"]))
    from datetime import timedelta
    record = {
        **record,
        "access_token": payload["access_token"],
        "expires_in": new_expires_in,
        "expires_at": (_now() + timedelta(seconds=new_expires_in)).isoformat(),
        "obtained_at": _now().isoformat(),
    }
    storage.save("instagram", record)
    return record


def _ig_token() -> tuple[str, str]:
    """Returns (access_token, ig_user_id)."""
    record = storage.load("instagram")
    if not record:
        raise RuntimeError("No Instagram token. Run `python -m plenya_social.oauth_setup` first.")
    record = _refresh_if_needed(record)
    return record["access_token"], record["user_id"]


def ig_request(
    method: str,
    path: str,
    *,
    params: dict | None = None,
    data: dict | None = None,
    json_body: dict | None = None,
) -> Any:
    """Make an authenticated Instagram Graph API request.

    Path may start with '/' (treated as relative to v21.0 base) or be a full URL.
    Adds access_token to params/data automatically.
    """
    token, _ = _ig_token()
    url = path if path.startswith("http") else f"{config.IG_GRAPH_BASE}{path}"

    merged_params = dict(params or {})
    if method.upper() == "GET":
        merged_params["access_token"] = token

    if method.upper() in ("POST", "DELETE") and data is None and json_body is None:
        data = {"access_token": token}
    elif data is not None:
        data = {**data, "access_token": token}

    r = httpx.request(
        method,
        url,
        params=merged_params if method.upper() == "GET" else (params or None),
        data=data,
        json=json_body,
        timeout=60,
    )
    if r.status_code >= 400:
        try:
            body = r.json()
        except Exception:
            body = r.text
        raise MetaError(r.status_code, body)
    return r.json() if r.content else {}


def ig_user_id() -> str:
    _, uid = _ig_token()
    return uid
