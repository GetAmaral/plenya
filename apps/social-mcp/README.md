# Plenya Social MCP

MCP server for managing Instagram (and eventually Facebook Page) interactions via Meta Graph API. Replaces Composio for IG-related actions so we can use endpoints Composio doesn't expose (like `instagram_business_manage_comments` like-comment).

## Setup

```bash
cd apps/social-mcp
python3 -m venv .venv
source .venv/bin/activate
pip install -e .

cp .env.example .env
# Edit .env: set PLENYA_SOCIAL_IG_CLIENT_SECRET and generate PLENYA_SOCIAL_ENCRYPTION_KEY

# Generate encryption key
python -c "from cryptography.fernet import Fernet; print(Fernet.generate_key().decode())"
# Paste output into .env as PLENYA_SOCIAL_ENCRYPTION_KEY
```

## One-time OAuth

```bash
python -m plenya_social.oauth_setup
```

Opens browser, captures token, exchanges for long-lived (60d), saves encrypted to `tokens/instagram.json`.

Requires `http://localhost:8765/callback` in the Meta App's Instagram Business Login redirect URIs.

## Register with Claude Code

```bash
claude mcp add plenya-social python -m plenya_social.server \
  --env PLENYA_SOCIAL_TOKENS_DIR=/home/user/plenya/apps/social-mcp/tokens \
  --env PLENYA_SOCIAL_ENCRYPTION_KEY=...
```

Restart Claude Code session to load tools.

## Token Refresh

Long-lived tokens last 60 days. Auto-refreshed on use when expiry is within 7 days. Manual refresh:

```bash
python -m plenya_social.oauth_setup --refresh
```
