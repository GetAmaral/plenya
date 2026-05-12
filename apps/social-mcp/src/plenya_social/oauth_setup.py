"""One-time OAuth setup for Instagram Business Login.

Run once to capture a long-lived (60-day) access token. Re-run with --refresh
to manually refresh an existing long-lived token.
"""

import argparse
import http.server
import secrets
import sys
import threading
import time
import urllib.parse
import webbrowser
from datetime import datetime, timedelta, timezone

import httpx

from . import config, storage


def _build_authorize_url(state: str) -> str:
    qs = urllib.parse.urlencode({
        "client_id": config.IG_CLIENT_ID,
        "redirect_uri": config.OAUTH_REDIRECT_URI,
        "response_type": "code",
        "scope": ",".join(config.IG_SCOPES),
        "state": state,
    })
    return f"{config.IG_AUTHORIZE_URL}?{qs}"


class _CallbackHandler(http.server.BaseHTTPRequestHandler):
    captured: dict = {}

    def do_GET(self):
        parsed = urllib.parse.urlparse(self.path)
        if parsed.path != "/callback":
            self.send_response(404)
            self.end_headers()
            return
        params = urllib.parse.parse_qs(parsed.query)
        if "code" in params and params.get("state", [""])[0] == _CallbackHandler.captured.get("expected_state"):
            _CallbackHandler.captured["code"] = params["code"][0]
            body = b"<h1>OK</h1><p>Token captured. You can close this tab.</p>"
        else:
            _CallbackHandler.captured["error"] = parsed.query
            body = f"<h1>Erro</h1><pre>{parsed.query}</pre>".encode()
        self.send_response(200)
        self.send_header("Content-Type", "text/html; charset=utf-8")
        self.end_headers()
        self.wfile.write(body)

    def log_message(self, *args, **kwargs):
        pass


def _exchange_code_for_short_lived(code: str) -> dict:
    r = httpx.post(
        config.IG_TOKEN_URL,
        data={
            "client_id": config.IG_CLIENT_ID,
            "client_secret": config.IG_CLIENT_SECRET,
            "grant_type": "authorization_code",
            "redirect_uri": config.OAUTH_REDIRECT_URI,
            "code": code,
        },
        timeout=30,
    )
    r.raise_for_status()
    return r.json()


def _exchange_for_long_lived(short_token: str) -> dict:
    r = httpx.get(
        config.IG_LONG_LIVED_URL,
        params={
            "grant_type": "ig_exchange_token",
            "client_secret": config.IG_CLIENT_SECRET,
            "access_token": short_token,
        },
        timeout=30,
    )
    r.raise_for_status()
    return r.json()


def _refresh_long_lived(token: str) -> dict:
    r = httpx.get(
        config.IG_REFRESH_URL,
        params={"grant_type": "ig_refresh_token", "access_token": token},
        timeout=30,
    )
    r.raise_for_status()
    return r.json()


def _fetch_user_profile(token: str) -> dict:
    r = httpx.get(
        f"{config.IG_GRAPH_BASE}/me",
        params={"fields": "id,username,account_type", "access_token": token},
        timeout=30,
    )
    r.raise_for_status()
    return r.json()


def _save_token(payload: dict, profile: dict) -> None:
    expires_in = int(payload.get("expires_in", 0))
    record = {
        "access_token": payload["access_token"],
        "token_type": payload.get("token_type", "bearer"),
        "expires_in": expires_in,
        "expires_at": (datetime.now(timezone.utc) + timedelta(seconds=expires_in)).isoformat(),
        "obtained_at": datetime.now(timezone.utc).isoformat(),
        "scopes": config.IG_SCOPES,
        "user_id": profile["id"],
        "username": profile["username"],
        "account_type": profile.get("account_type"),
    }
    storage.save("instagram", record)


def run_initial() -> None:
    state = secrets.token_urlsafe(16)
    _CallbackHandler.captured = {"expected_state": state}

    server = http.server.HTTPServer(("127.0.0.1", config.OAUTH_PORT), _CallbackHandler)
    thread = threading.Thread(target=server.serve_forever, daemon=True)
    thread.start()

    url = _build_authorize_url(state)
    print(f"\nAbra este URL no browser (deve abrir automaticamente):\n{url}\n", file=sys.stderr)
    try:
        webbrowser.open(url)
    except Exception:
        pass

    timeout = time.time() + 300
    while "code" not in _CallbackHandler.captured and "error" not in _CallbackHandler.captured:
        if time.time() > timeout:
            server.shutdown()
            raise SystemExit("OAuth timeout (5 min). Tente novamente.")
        time.sleep(0.5)
    server.shutdown()

    if "error" in _CallbackHandler.captured:
        raise SystemExit(f"OAuth erro: {_CallbackHandler.captured['error']}")

    code = _CallbackHandler.captured["code"]
    print("Code capturado. Trocando por short-lived token...", file=sys.stderr)
    short = _exchange_code_for_short_lived(code)
    short_token = short["access_token"]
    print(f"Short-lived OK (user_id: {short.get('user_id')}). Trocando por long-lived...", file=sys.stderr)
    long_payload = _exchange_for_long_lived(short_token)
    print(f"Long-lived OK (expires_in: {long_payload.get('expires_in')} s ≈ {long_payload.get('expires_in', 0) // 86400}d)", file=sys.stderr)

    profile = _fetch_user_profile(long_payload["access_token"])
    _save_token(long_payload, profile)
    print(f"\nToken salvo em tokens/instagram.enc\nUsuário: @{profile['username']} (id {profile['id']}, {profile.get('account_type')})", file=sys.stderr)


def run_refresh() -> None:
    current = storage.load("instagram")
    if not current:
        raise SystemExit("Sem token salvo. Rode sem --refresh primeiro.")
    print(f"Refrescando token de @{current['username']}...", file=sys.stderr)
    refreshed = _refresh_long_lived(current["access_token"])
    profile = _fetch_user_profile(refreshed["access_token"])
    _save_token(refreshed, profile)
    print(f"Refresh OK. Novo expires_in: {refreshed.get('expires_in')} s", file=sys.stderr)


def main() -> None:
    parser = argparse.ArgumentParser(description="OAuth setup for Plenya Social MCP")
    parser.add_argument("--refresh", action="store_true", help="Refresh existing long-lived token")
    args = parser.parse_args()
    if args.refresh:
        run_refresh()
    else:
        run_initial()


if __name__ == "__main__":
    main()
