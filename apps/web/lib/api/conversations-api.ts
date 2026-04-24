import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { apiClient } from '../api-client';

// =====================================================
// Types — espelham response do backend (Bloco B)
// =====================================================

export type ConversationOwnerType = 'lead' | 'patient';
export type ConversationChannel = 'email' | 'whatsapp' | 'internal';
export type ConversationDirection = 'in' | 'out';

export interface ConversationItem {
  ownerType: ConversationOwnerType;
  ownerId: string;
  name: string;
  email?: string;
  phone?: string;
  /** Lead-only: opt-in flags (vem do backend). */
  emailOptIn?: boolean;
  whatsAppOptIn?: boolean;
  /** Lead-only: timestamp da última inbound — pra calcular janela 24h. */
  lastInboundAt?: string;
  lastChannel: ConversationChannel;
  lastDirection: ConversationDirection;
  lastSnippet: string;
  lastAt: string;
  unreadCount: number;
  assignedToUserId?: string;
  channels: string[];
}

export interface ConversationListResult {
  items: ConversationItem[];
  nextCursor?: string;
}

export interface ConversationMessageActor {
  id: string;
  name: string;
  email?: string;
}

/** Anexo persistido na metadata da activity (mesmo schema inbound + outbound). */
export interface ConversationMessageAttachment {
  filename: string;
  /** Path relativo a /app/uploads (ex: "conversation-outbound/<uid>/<file>.pdf"). */
  path: string;
  content_type?: string;
  size_bytes?: number;
}

export interface ConversationMessageMetadata {
  subject?: string;
  from?: string;
  messageId?: string;
  recipient?: string;
  status?: string;
  wa_message_id?: string;
  attachments?: ConversationMessageAttachment[];
  [k: string]: unknown;
}

export interface ConversationMessage {
  id: string;
  leadId?: string;
  patientId?: string;
  type: string; // 'message_received' | 'message_sent' | 'message_status_changed' | etc.
  channel: ConversationChannel;
  content?: string;
  metadata?: ConversationMessageMetadata;
  actorUserId?: string;
  actor?: ConversationMessageActor;
  createdAt: string;
}

export interface ConversationListFilters {
  assignedToMe?: boolean;
  unreadOnly?: boolean;
  channel?: 'email' | 'whatsapp';
  search?: string;
  cursor?: string;
  limit?: number;
}

// =====================================================
// Query keys
// =====================================================

/**
 * Conta total de conversas não-lidas (Lead + Patient unificado).
 * Usado pelo badge do sidebar — polling 20s pra não bater muito.
 */
export function useConversationsUnreadCount() {
  return useQuery({
    queryKey: ['conversations', 'unread-count'],
    queryFn: async () => {
      const data = await apiClient.get<ConversationListResult>(
        '/api/v1/conversations?unread_only=true&limit=200'
      );
      return data.items.reduce((sum, it) => sum + it.unreadCount, 0);
    },
    staleTime: 10_000,
    refetchInterval: 20_000,
    refetchOnWindowFocus: true,
  });
}

export const conversationKeys = {
  all: ['conversations'] as const,
  list: (filters: ConversationListFilters) =>
    [...conversationKeys.all, 'list', filters] as const,
  messages: (type: ConversationOwnerType, id: string) =>
    [...conversationKeys.all, 'messages', type, id] as const,
};

// =====================================================
// Helpers
// =====================================================

function buildListQuery(filters: ConversationListFilters): string {
  const params = new URLSearchParams();
  if (filters.assignedToMe) params.set('assigned_to_me', 'true');
  if (filters.unreadOnly) params.set('unread_only', 'true');
  if (filters.channel) params.set('channel', filters.channel);
  if (filters.search) params.set('search', filters.search);
  if (filters.cursor) params.set('cursor', filters.cursor);
  if (filters.limit) params.set('limit', String(filters.limit));
  const qs = params.toString();
  return qs ? `?${qs}` : '';
}

// =====================================================
// Queries
// =====================================================

/**
 * Lista conversas unificadas (leads + patients).
 *
 * Polling 15s + refetchOnWindowFocus mantém a lista quase-realtime sem WebSocket.
 * staleTime curto (10s) garante que filtros mudam imediatamente sem cache stale.
 */
export function useConversations(filters: ConversationListFilters) {
  return useQuery({
    queryKey: conversationKeys.list(filters),
    queryFn: () =>
      apiClient.get<ConversationListResult>(
        `/api/v1/conversations${buildListQuery(filters)}`
      ),
    staleTime: 10_000,
    refetchInterval: 15_000,
    refetchOnWindowFocus: true,
  });
}

/**
 * Timeline da conversa (ASC por created_at).
 *
 * Side-effect documentado: o backend marca a conversa como lida ao chamar este endpoint.
 * Por isso é seguro chamar sempre que a viewer abre — invalidamos a list pra zerar o badge.
 */
export function useConversationMessages(
  type: ConversationOwnerType | null,
  id: string | null
) {
  const qc = useQueryClient();
  return useQuery({
    queryKey: type && id ? conversationKeys.messages(type, id) : ['conversations', 'messages', 'noop'],
    queryFn: async () => {
      if (!type || !id) return [] as ConversationMessage[];
      const raw = await apiClient.get<ConversationMessage[] | { items: ConversationMessage[] }>(
        `/api/v1/conversations/${type}/${id}/messages`
      );
      const data = Array.isArray(raw) ? raw : raw?.items ?? [];
      // Side-effect: backend marcou como lido. Invalidamos só a list (não a si mesmo
      // — invalidar conversationKeys.all causava loop refetch que travava o Query Client).
      qc.invalidateQueries({ queryKey: [...conversationKeys.all, 'list'] });
      return data;
    },
    enabled: !!type && !!id,
    refetchInterval: 5_000, // 5s — responsivo pra inbound real-time
    staleTime: 1_000,
  });
}

// =====================================================
// Mutations
// =====================================================

/** Marca conversa como lida sem precisar abrir/fetch das mensagens. */
export function useMarkConversationRead() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ type, id }: { type: ConversationOwnerType; id: string }) =>
      apiClient.post<void>(`/api/v1/conversations/${type}/${id}/read`),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: conversationKeys.all });
    },
  });
}

export interface SendConversationEmailAttachment {
  /** Path relativo retornado pelo upload endpoint. */
  path: string;
  /** Nome original mostrado ao destinatário. */
  filename: string;
}

export interface SendConversationEmailPayload {
  subject?: string;
  bodyText: string;
  inReplyTo?: string;
  references?: string[];
  attachments?: SendConversationEmailAttachment[];
}

/** Resposta do POST /conversations/attachments/upload. */
export interface UploadedAttachment {
  path: string;
  filename: string;
  contentType: string;
  size: number;
}

/**
 * Upload de um arquivo individual. Cada Drop dispara um POST imediato pra ter feedback
 * por-arquivo (status: uploading|done|error) — composer só habilita Send quando todos
 * estão done.
 */
export function useUploadConversationAttachment() {
  return useMutation({
    mutationFn: async (file: File): Promise<UploadedAttachment> => {
      const fd = new FormData();
      fd.append('file', file);
      return apiClient.post<UploadedAttachment>(
        '/api/v1/conversations/attachments/upload',
        fd
      );
    },
  });
}

/**
 * URL pública do arquivo no static server. Usada nos chips clicáveis
 * (download direto pelo browser).
 *
 * Em dev: http://localhost:3001/uploads/<path>
 * Em prod: https://api.plenyasaude.com.br/uploads/<path>
 */
export function attachmentDownloadUrl(path: string): string {
  const base = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:3001';
  const clean = path.replace(/^\/+/, '').replace(/^uploads\//, '');
  return `${base}/uploads/${clean}`;
}

export function useSendConversationEmail(type: ConversationOwnerType, id: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (payload: SendConversationEmailPayload) =>
      apiClient.post<{ message: string; to?: string }>(
        `/api/v1/conversations/${type}/${id}/email`,
        payload
      ),
    onSuccess: async () => {
      // refetchQueries força o re-fetch imediato — invalidate sozinho não dispara
      // refetch se a query estiver marcada inactive ou se tiver outra refetch em curso.
      await qc.refetchQueries({ queryKey: conversationKeys.messages(type, id) });
      qc.invalidateQueries({ queryKey: [...conversationKeys.all, 'list'] });
    },
  });
}

// =====================================================
// Compose (novo email global) — Bloco 7
// =====================================================

/**
 * Payload do POST /conversations/compose — vendedor inicia email pra qualquer
 * endereço. Backend resolve owner: Patient (email match) → Lead ativo → cria Lead.
 */
export interface ComposeConversationPayload {
  to: string;
  /** Opcional — se Lead novo for criado, vira Lead.Name. */
  name?: string;
  subject: string;
  bodyText: string;
  bodyHTML?: string;
  attachments?: SendConversationEmailAttachment[];
}

/**
 * Resposta do compose. `leadCreated=true` distingue Lead novo de Lead/Patient
 * pré-existente — UI mostra toast diferenciado e navega pra conversa criada.
 */
export interface ComposeConversationResult {
  ownerType: ConversationOwnerType;
  ownerId: string;
  leadCreated: boolean;
  url?: string;
}

/**
 * Regex de validação client-side — espelha a do backend (RFC 5322 simplificada).
 * Não substitui validação do backend; serve só pra UX (desabilitar Send antes).
 */
export const COMPOSE_EMAIL_REGEX = /^[^@\s]+@[^@\s]+\.[^@\s]+$/;

export function useConversationCompose() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (payload: ComposeConversationPayload) =>
      apiClient.post<ComposeConversationResult>(
        '/api/v1/conversations/compose',
        payload
      ),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: conversationKeys.all });
    },
  });
}

export interface SendConversationWhatsAppPayload {
  bodyText: string;
}

export function useSendConversationWhatsApp(
  type: ConversationOwnerType,
  id: string
) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (payload: SendConversationWhatsAppPayload) =>
      apiClient.post<{ message: string }>(
        `/api/v1/conversations/${type}/${id}/whatsapp`,
        payload
      ),
    onSuccess: async () => {
      await qc.refetchQueries({ queryKey: conversationKeys.messages(type, id) });
      qc.invalidateQueries({ queryKey: [...conversationKeys.all, 'list'] });
    },
  });
}

// =====================================================
// Helpers visuais
// =====================================================

/** Cor consistente baseada no hash do nome — usado nos avatares. */
export function avatarColorClass(name: string): string {
  const palette = [
    'bg-rose-200 text-rose-900',
    'bg-amber-200 text-amber-900',
    'bg-emerald-200 text-emerald-900',
    'bg-sky-200 text-sky-900',
    'bg-violet-200 text-violet-900',
    'bg-fuchsia-200 text-fuchsia-900',
    'bg-teal-200 text-teal-900',
    'bg-orange-200 text-orange-900',
  ];
  let h = 0;
  for (let i = 0; i < name.length; i++) {
    h = (h * 31 + name.charCodeAt(i)) >>> 0;
  }
  return palette[h % palette.length];
}

export function initials(name: string): string {
  const parts = name.trim().split(/\s+/).filter(Boolean);
  if (parts.length === 0) return '?';
  if (parts.length === 1) return parts[0].slice(0, 2).toUpperCase();
  return (parts[0][0] + parts[parts.length - 1][0]).toUpperCase();
}
