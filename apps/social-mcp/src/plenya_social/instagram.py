"""Instagram tool implementations. Each function returns JSON-serializable data."""

from typing import Any

from .meta_client import ig_request, ig_user_id


def get_user_info() -> dict:
    """Get authenticated Instagram Business/Creator profile info."""
    return ig_request(
        "GET",
        "/me",
        params={"fields": "id,username,account_type,biography,followers_count,follows_count,media_count,profile_picture_url,website"},
    )


def list_media(limit: int = 25, after: str | None = None) -> dict:
    """List recent media (posts/reels) from the authenticated user."""
    params: dict[str, Any] = {
        "fields": "id,caption,media_type,media_product_type,permalink,timestamp,comments_count,like_count,thumbnail_url",
        "limit": min(max(limit, 1), 100),
    }
    if after:
        params["after"] = after
    return ig_request("GET", f"/{ig_user_id()}/media", params=params)


def get_media_comments(media_id: str, limit: int = 50, after: str | None = None) -> dict:
    """List comments on a media post. Includes replies."""
    params: dict[str, Any] = {
        "fields": "id,text,username,timestamp,like_count,replies{id,text,username,timestamp,from},parent_id",
        "limit": min(max(limit, 1), 100),
    }
    if after:
        params["after"] = after
    return ig_request("GET", f"/{media_id}/comments", params=params)


def post_comment_reply(comment_id: str, message: str) -> dict:
    """Reply to an existing comment."""
    return ig_request(
        "POST",
        f"/{comment_id}/replies",
        data={"message": message},
    )


def like_comment(comment_id: str) -> dict:
    """Like a comment. Tests whether `instagram_business_manage_comments` covers likes."""
    return ig_request("POST", f"/{comment_id}/likes")


def unlike_comment(comment_id: str) -> dict:
    """Remove a like from a comment."""
    return ig_request("DELETE", f"/{comment_id}/likes")


def delete_comment(comment_id: str) -> dict:
    """Delete a comment your account created or moderates."""
    return ig_request("DELETE", f"/{comment_id}")


def hide_comment(comment_id: str, hidden: bool = True) -> dict:
    """Hide or unhide a comment."""
    return ig_request("POST", f"/{comment_id}", data={"hide": "true" if hidden else "false"})


def list_conversations(limit: int = 25, after: str | None = None) -> dict:
    """List DM conversations for the authenticated user."""
    params: dict[str, Any] = {
        "platform": "instagram",
        "fields": "id,updated_time,participants",
        "limit": min(max(limit, 1), 200),
    }
    if after:
        params["after"] = after
    return ig_request("GET", f"/{ig_user_id()}/conversations", params=params)


def list_messages(conversation_id: str, limit: int = 25, after: str | None = None) -> dict:
    """List messages in a conversation (newest first)."""
    params: dict[str, Any] = {
        "fields": "id,from,to,message,created_time,attachments",
        "limit": min(max(limit, 1), 200),
    }
    if after:
        params["after"] = after
    return ig_request("GET", f"/{conversation_id}/messages", params=params)


def send_message(recipient_user_id: str, text: str) -> dict:
    """Send a text DM to a user (must be inside the 24h messaging window or be a HUMAN_AGENT reply).

    `recipient_user_id` is the IGSID (Instagram-scoped user ID) of the target user — usually
    obtained from the `from.id` field of a message in `list_messages`.
    """
    return ig_request(
        "POST",
        f"/{ig_user_id()}/messages",
        json_body={
            "recipient": {"id": recipient_user_id},
            "message": {"text": text},
        },
    )


def mark_seen(recipient_user_id: str) -> dict:
    """Mark messages from a user as seen."""
    return ig_request(
        "POST",
        f"/{ig_user_id()}/messages",
        json_body={
            "recipient": {"id": recipient_user_id},
            "sender_action": "mark_seen",
        },
    )


def get_user_tags(limit: int = 25, after: str | None = None) -> dict:
    """List media where the authenticated user has been tagged/mentioned."""
    params: dict[str, Any] = {
        "fields": "id,caption,media_type,permalink,timestamp,username,comments_count,like_count",
        "limit": min(max(limit, 1), 100),
    }
    if after:
        params["after"] = after
    return ig_request("GET", f"/{ig_user_id()}/tags", params=params)
