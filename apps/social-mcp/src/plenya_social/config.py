import os
from pathlib import Path
from dotenv import load_dotenv

_ROOT = Path(__file__).resolve().parents[2]
load_dotenv(_ROOT / ".env")

IG_CLIENT_ID = os.environ["PLENYA_SOCIAL_IG_CLIENT_ID"]
IG_CLIENT_SECRET = os.environ["PLENYA_SOCIAL_IG_CLIENT_SECRET"]
ENCRYPTION_KEY = os.environ["PLENYA_SOCIAL_ENCRYPTION_KEY"].encode()

TOKENS_DIR = Path(os.environ.get("PLENYA_SOCIAL_TOKENS_DIR") or (_ROOT / "tokens")).resolve()
TOKENS_DIR.mkdir(parents=True, exist_ok=True)

OAUTH_REDIRECT_URI = os.environ.get("PLENYA_SOCIAL_OAUTH_REDIRECT_URI", "http://localhost:8765/callback")
OAUTH_PORT = int(os.environ.get("PLENYA_SOCIAL_OAUTH_PORT", "8765"))

IG_SCOPES = [
    "instagram_business_basic",
    "instagram_business_manage_messages",
    "instagram_business_manage_comments",
    "instagram_business_content_publish",
    "instagram_business_manage_insights",
]

IG_AUTHORIZE_URL = "https://www.instagram.com/oauth/authorize"
IG_TOKEN_URL = "https://api.instagram.com/oauth/access_token"
IG_LONG_LIVED_URL = "https://graph.instagram.com/access_token"
IG_REFRESH_URL = "https://graph.instagram.com/refresh_access_token"
IG_GRAPH_BASE = "https://graph.instagram.com/v21.0"

REFRESH_WINDOW_SECONDS = 7 * 24 * 3600
