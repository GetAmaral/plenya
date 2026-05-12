"""Plenya Social MCP server (stdio transport).

Run via:  python -m plenya_social.server
Or register: claude mcp add plenya-social python -m plenya_social.server
"""

import asyncio
import json
import sys
from typing import Any

from mcp.server import Server
from mcp.server.stdio import stdio_server
from mcp.types import TextContent, Tool

from . import instagram
from .meta_client import MetaError

server: Server = Server("plenya-social")


_TOOLS: list[tuple[Tool, Any]] = [
    (
        Tool(
            name="ig_get_user_info",
            description="Get the authenticated Instagram Business/Creator account profile (username, followers, media count, etc.).",
            inputSchema={"type": "object", "properties": {}, "additionalProperties": False},
        ),
        lambda args: instagram.get_user_info(),
    ),
    (
        Tool(
            name="ig_list_media",
            description="List recent media (posts/reels) from the authenticated user, newest first.",
            inputSchema={
                "type": "object",
                "properties": {
                    "limit": {"type": "integer", "minimum": 1, "maximum": 100, "default": 25},
                    "after": {"type": "string", "description": "Cursor from previous response's paging.cursors.after"},
                },
                "additionalProperties": False,
            },
        ),
        lambda args: instagram.list_media(limit=args.get("limit", 25), after=args.get("after")),
    ),
    (
        Tool(
            name="ig_get_media_comments",
            description="List comments on a specific media post. Includes nested replies.",
            inputSchema={
                "type": "object",
                "properties": {
                    "media_id": {"type": "string", "description": "Instagram media (post/reel) ID"},
                    "limit": {"type": "integer", "minimum": 1, "maximum": 100, "default": 50},
                    "after": {"type": "string"},
                },
                "required": ["media_id"],
                "additionalProperties": False,
            },
        ),
        lambda args: instagram.get_media_comments(args["media_id"], args.get("limit", 50), args.get("after")),
    ),
    (
        Tool(
            name="ig_post_comment_reply",
            description="Reply to an existing Instagram comment. Returns the new reply's ID.",
            inputSchema={
                "type": "object",
                "properties": {
                    "comment_id": {"type": "string", "description": "ID of the comment to reply to"},
                    "message": {"type": "string", "description": "Reply text (no hard char limit on API)"},
                },
                "required": ["comment_id", "message"],
                "additionalProperties": False,
            },
        ),
        lambda args: instagram.post_comment_reply(args["comment_id"], args["message"]),
    ),
    (
        Tool(
            name="ig_like_comment",
            description="Like a comment (the coraçãozinho). Tests the API capability under instagram_business_manage_comments scope.",
            inputSchema={
                "type": "object",
                "properties": {"comment_id": {"type": "string"}},
                "required": ["comment_id"],
                "additionalProperties": False,
            },
        ),
        lambda args: instagram.like_comment(args["comment_id"]),
    ),
    (
        Tool(
            name="ig_unlike_comment",
            description="Remove a like from a comment.",
            inputSchema={
                "type": "object",
                "properties": {"comment_id": {"type": "string"}},
                "required": ["comment_id"],
                "additionalProperties": False,
            },
        ),
        lambda args: instagram.unlike_comment(args["comment_id"]),
    ),
    (
        Tool(
            name="ig_delete_comment",
            description="Delete a comment (use for spam/abusive content). Irreversible.",
            inputSchema={
                "type": "object",
                "properties": {"comment_id": {"type": "string"}},
                "required": ["comment_id"],
                "additionalProperties": False,
            },
        ),
        lambda args: instagram.delete_comment(args["comment_id"]),
    ),
    (
        Tool(
            name="ig_hide_comment",
            description="Hide or unhide a comment (moderation, not deletion).",
            inputSchema={
                "type": "object",
                "properties": {
                    "comment_id": {"type": "string"},
                    "hidden": {"type": "boolean", "default": True},
                },
                "required": ["comment_id"],
                "additionalProperties": False,
            },
        ),
        lambda args: instagram.hide_comment(args["comment_id"], args.get("hidden", True)),
    ),
    (
        Tool(
            name="ig_list_conversations",
            description="List DM conversations for the authenticated user, ordered by last update.",
            inputSchema={
                "type": "object",
                "properties": {
                    "limit": {"type": "integer", "minimum": 1, "maximum": 200, "default": 25},
                    "after": {"type": "string"},
                },
                "additionalProperties": False,
            },
        ),
        lambda args: instagram.list_conversations(args.get("limit", 25), args.get("after")),
    ),
    (
        Tool(
            name="ig_list_messages",
            description="List messages within a DM conversation (newest first).",
            inputSchema={
                "type": "object",
                "properties": {
                    "conversation_id": {"type": "string"},
                    "limit": {"type": "integer", "minimum": 1, "maximum": 200, "default": 25},
                    "after": {"type": "string"},
                },
                "required": ["conversation_id"],
                "additionalProperties": False,
            },
        ),
        lambda args: instagram.list_messages(args["conversation_id"], args.get("limit", 25), args.get("after")),
    ),
    (
        Tool(
            name="ig_send_message",
            description="Send a DM text message to a user (must be within Meta's 24h messaging window).",
            inputSchema={
                "type": "object",
                "properties": {
                    "recipient_user_id": {"type": "string", "description": "IGSID of recipient (from message.from.id)"},
                    "text": {"type": "string"},
                },
                "required": ["recipient_user_id", "text"],
                "additionalProperties": False,
            },
        ),
        lambda args: instagram.send_message(args["recipient_user_id"], args["text"]),
    ),
    (
        Tool(
            name="ig_mark_seen",
            description="Mark messages from a user as seen (clears unread indicator).",
            inputSchema={
                "type": "object",
                "properties": {"recipient_user_id": {"type": "string"}},
                "required": ["recipient_user_id"],
                "additionalProperties": False,
            },
        ),
        lambda args: instagram.mark_seen(args["recipient_user_id"]),
    ),
    (
        Tool(
            name="ig_get_user_tags",
            description="List media where the authenticated user was tagged or mentioned by others.",
            inputSchema={
                "type": "object",
                "properties": {
                    "limit": {"type": "integer", "minimum": 1, "maximum": 100, "default": 25},
                    "after": {"type": "string"},
                },
                "additionalProperties": False,
            },
        ),
        lambda args: instagram.get_user_tags(args.get("limit", 25), args.get("after")),
    ),
]

_TOOL_MAP = {tool.name: handler for tool, handler in _TOOLS}


@server.list_tools()
async def _list_tools() -> list[Tool]:
    return [tool for tool, _ in _TOOLS]


@server.call_tool()
async def _call_tool(name: str, arguments: dict[str, Any]) -> list[TextContent]:
    handler = _TOOL_MAP.get(name)
    if handler is None:
        return [TextContent(type="text", text=json.dumps({"error": f"unknown tool: {name}"}))]
    try:
        result = handler(arguments or {})
        return [TextContent(type="text", text=json.dumps(result, ensure_ascii=False, indent=2))]
    except MetaError as e:
        print(f"[plenya-social-mcp] MetaError on {name}: {e}", file=sys.stderr)
        return [TextContent(type="text", text=json.dumps({"error": str(e), "status": e.status, "body": e.body}, ensure_ascii=False))]
    except Exception as e:
        print(f"[plenya-social-mcp] Unhandled on {name}: {e!r}", file=sys.stderr)
        return [TextContent(type="text", text=json.dumps({"error": repr(e)}, ensure_ascii=False))]


async def _run() -> None:
    async with stdio_server() as (read, write):
        await server.run(read, write, server.create_initialization_options())


def main() -> None:
    asyncio.run(_run())


if __name__ == "__main__":
    main()
