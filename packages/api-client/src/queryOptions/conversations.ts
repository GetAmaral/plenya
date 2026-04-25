import { queryOptions } from '@tanstack/react-query';
import { api } from '../fetcher';
import { queryKeys } from '../queryKeys';

export type ConversationOwnerType = 'lead' | 'patient';
export type ConversationChannel = 'email' | 'whatsapp' | 'internal';
export type ConversationDirection = 'in' | 'out';

export interface ConversationItem {
  ownerType: ConversationOwnerType;
  ownerId: string;
  name: string;
  email?: string;
  phone?: string;
  lastChannel: ConversationChannel;
  lastDirection: ConversationDirection;
  lastSnippet: string;
  lastAt: string;
  unreadCount: number;
  assignedToUserId?: string;
  channels: string[];
  leadStatus?: string;
  emailOptIn: boolean;
  whatsAppOptIn: boolean;
  lastInboundAt?: string;
}

export interface ListConversationsResult {
  items: ConversationItem[];
  nextCursor?: string;
}

export interface ConversationMessage {
  id: string;
  type: string;
  direction?: ConversationDirection;
  channel?: ConversationChannel;
  content?: string;
  metadata?: Record<string, unknown>;
  attachments?: Array<{
    url: string;
    filename: string;
    contentType?: string;
    size?: number;
  }>;
  createdAt: string;
  actorName?: string;
}

export interface ListConversationsParams {
  assignedToMe?: boolean;
  unreadOnly?: boolean;
  channel?: 'email' | 'whatsapp';
  search?: string;
  limit?: number;
  cursor?: string;
}

const conversationKeys = {
  all: () => [...queryKeys.all, 'conversations'] as const,
  list: (params: ListConversationsParams) =>
    [...conversationKeys.all(), 'list', params] as const,
  messages: (ownerType: ConversationOwnerType, ownerId: string) =>
    [...conversationKeys.all(), 'messages', ownerType, ownerId] as const,
};

export const conversationKeysFor = conversationKeys;

export const conversationsListOptions = (params: ListConversationsParams = {}) =>
  queryOptions({
    queryKey: conversationKeys.list(params),
    queryFn: ({ signal }) => {
      const qs = new URLSearchParams();
      if (params.assignedToMe) qs.set('assigned_to_me', 'true');
      if (params.unreadOnly) qs.set('unread_only', 'true');
      if (params.channel) qs.set('channel', params.channel);
      if (params.search) qs.set('search', params.search);
      if (params.limit) qs.set('limit', String(params.limit));
      if (params.cursor) qs.set('cursor', params.cursor);
      const suffix = qs.toString() ? `?${qs.toString()}` : '';
      return api.get<ListConversationsResult>(`/api/v1/conversations${suffix}`, { signal });
    },
    refetchInterval: 30_000,
  });

export const conversationMessagesOptions = (
  ownerType: ConversationOwnerType,
  ownerId: string,
) =>
  queryOptions({
    queryKey: conversationKeys.messages(ownerType, ownerId),
    queryFn: ({ signal }) =>
      api
        .get<{ items: ConversationMessage[] }>(
          `/api/v1/conversations/${ownerType}/${ownerId}/messages`,
          { signal },
        )
        .then((r) => r.items ?? []),
    enabled: Boolean(ownerType && ownerId),
    refetchInterval: 15_000, // polling real-time
  });

export const conversationMutations = {
  markRead: (ownerType: ConversationOwnerType, ownerId: string) =>
    api.post<void>(`/api/v1/conversations/${ownerType}/${ownerId}/read`),
  sendEmail: (
    ownerType: ConversationOwnerType,
    ownerId: string,
    body: { content: string; attachments?: string[] },
  ) =>
    api.post<ConversationMessage>(
      `/api/v1/conversations/${ownerType}/${ownerId}/email`,
      body,
    ),
  sendWhatsApp: (
    ownerType: ConversationOwnerType,
    ownerId: string,
    body: { content: string },
  ) =>
    api.post<ConversationMessage>(
      `/api/v1/conversations/${ownerType}/${ownerId}/whatsapp`,
      body,
    ),
  aiSummary: (ownerType: ConversationOwnerType, ownerId: string) =>
    api.post<{ summary: string }>(
      `/api/v1/conversations/${ownerType}/${ownerId}/ai/summary`,
    ),
  aiSuggestReply: (
    ownerType: ConversationOwnerType,
    ownerId: string,
    body: { instruction?: string } = {},
  ) =>
    api.post<{ reply: string }>(
      `/api/v1/conversations/${ownerType}/${ownerId}/ai/suggest-reply`,
      body,
    ),
};

/** Path da rota do anexo conversa (compat com upload genérico /uploads). */
export const conversationAttachmentUploadUrl = '/api/v1/conversations/attachments/upload';
