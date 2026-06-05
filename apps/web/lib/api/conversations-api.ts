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
  /** "phone_app" quando a mensagem foi enviada pela equipe pelo app do celular (coexistence). */
  origin?: string;
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
  // Mídia (WhatsApp Fase 2). mediaType: image|audio|voice|video|document|sticker.
  mediaType?: string;
  mediaMime?: string;
  mediaFilename?: string;
  mediaSizeBytes?: number;
  /** Quando setado, o arquivo foi salvo nas mídias do prontuário do paciente. */
  patientDocumentId?: string;
  /** Transcrição de áudio (Fase 2.2). */
  transcription?: string;
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
 * Com `channel`, conta só aquele canal — usado pelos badges separados de WhatsApp e E-mail
 * na sidebar/dock. Sem `channel`, conta tudo. Polling 20s pra não bater muito.
 */
export function useConversationsUnreadCount(channel?: 'whatsapp' | 'email') {
  return useQuery({
    queryKey: ['conversations', 'unread-count', channel ?? 'all'],
    queryFn: async () => {
      const ch = channel ? `&channel=${channel}` : '';
      const data = await apiClient.get<ConversationListResult>(
        `/api/v1/conversations?unread_only=true&limit=200${ch}`
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
  messages: (type: ConversationOwnerType, id: string, channel?: string) =>
    [...conversationKeys.all, 'messages', type, id, channel ?? 'all'] as const,
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
  id: string | null,
  channel?: 'whatsapp' | 'email'
) {
  const qc = useQueryClient();
  return useQuery({
    queryKey: type && id ? conversationKeys.messages(type, id, channel) : ['conversations', 'messages', 'noop'],
    queryFn: async () => {
      if (!type || !id) return [] as ConversationMessage[];
      const ch = channel ? `?channel=${channel}` : '';
      const raw = await apiClient.get<ConversationMessage[] | { items: ConversationMessage[] }>(
        `/api/v1/conversations/${type}/${id}/messages${ch}`
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

/** Endpoint autenticado que serve a mídia (WhatsApp Fase 2) de uma mensagem. */
export function conversationMediaEndpoint(
  type: ConversationOwnerType,
  id: string,
  activityId: string,
): string {
  return `/api/v1/conversations/${type}/${id}/media/${activityId}`;
}

/** Busca a mídia de uma mensagem como Blob (autenticado, via apiClient). */
export function fetchConversationMedia(
  type: ConversationOwnerType,
  id: string,
  activityId: string,
): Promise<Blob> {
  return apiClient.getBlob(conversationMediaEndpoint(type, id, activityId));
}

export interface InterpretExamResult {
  documentId: string;
  batchId: string;
  jobId: string;
}

/** Ação sobre a mídia de uma mensagem. attachmentIndex aponta um anexo de e-mail;
 * ausente = mídia WhatsApp da própria atividade. */
export interface MediaActionArgs {
  activityId: string;
  attachmentIndex?: number;
}

/** Busca um anexo de e-mail como Blob (autenticado). */
export function fetchConversationAttachment(
  type: ConversationOwnerType,
  id: string,
  activityId: string,
  idx: number,
): Promise<Blob> {
  return apiClient.getBlob(
    `/api/v1/conversations/${type}/${id}/messages/${activityId}/attachments/${idx}`,
  );
}

/** Salva a mídia/anexo de uma mensagem nas mídias do prontuário do paciente. */
export function useSaveToProntuario(type: ConversationOwnerType, id: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ activityId, attachmentIndex }: MediaActionArgs) =>
      apiClient.post<{ documentId: string }>(
        `/api/v1/conversations/${type}/${id}/messages/${activityId}/save-to-prontuario`,
        attachmentIndex != null ? { attachmentIndex } : {},
      ),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: conversationKeys.messages(type, id) });
      qc.invalidateQueries({ queryKey: ['patient-documents'] });
    },
  });
}

/** Submete a mídia de uma mensagem ao interpretador de exames (salva no prontuário se preciso). */
export function useInterpretExam(type: ConversationOwnerType, id: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ activityId, attachmentIndex }: MediaActionArgs) =>
      apiClient.post<InterpretExamResult>(
        `/api/v1/conversations/${type}/${id}/media/${activityId}/interpret-exam`,
        attachmentIndex != null ? { attachmentIndex } : {},
      ),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: ['lab-result-batches'] });
      qc.invalidateQueries({ queryKey: conversationKeys.messages(type, id) });
    },
  });
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
  bodyText?: string;
  /** Mídia (uma por mensagem) já enviada via /conversations/attachments/upload. */
  attachments?: SendConversationEmailAttachment[];
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

export interface SendConversationWhatsAppTemplatePayload {
  name: string;
  language?: string;
  params?: string[];
}

/** Envia template aprovado (reabre conversa fora da janela de 24h). */
export function useSendConversationWhatsAppTemplate(
  type: ConversationOwnerType,
  id: string
) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (payload: SendConversationWhatsAppTemplatePayload) =>
      apiClient.post<{ message: string }>(
        `/api/v1/conversations/${type}/${id}/whatsapp/template`,
        payload
      ),
    onSuccess: async () => {
      await qc.refetchQueries({ queryKey: conversationKeys.messages(type, id) });
      qc.invalidateQueries({ queryKey: [...conversationKeys.all, 'list'] });
    },
  });
}

// =====================================================
// IA — resumir conversa + sugerir resposta
// =====================================================

export interface AISummaryResult {
  summary: string;
  generatedAt: string;
  cached: boolean;
}

export interface AISuggestionResult {
  suggestion: string;
  model: string;
}

export interface AISuggestionPayload {
  intent?: string;
}

/**
 * Resumir conversa via Claude Haiku. Backend cacheia 1h por hash do transcript;
 * passe `force: true` pra regenerar. 422 = sem mensagens, 502 = Claude falhou,
 * 504 = timeout.
 */
export function useConversationAISummary(
  type: ConversationOwnerType,
  id: string
) {
  return useMutation({
    mutationFn: ({ force }: { force?: boolean } = {}) =>
      apiClient.post<AISummaryResult>(
        `/api/v1/conversations/${type}/${id}/ai/summary${force ? '?force=true' : ''}`,
        {}
      ),
  });
}

/**
 * Sugerir resposta no tom Plenya via Claude Sonnet. `intent` opcional guia o modelo
 * (ex: "agendar consulta", "recusar educadamente"). Sem cache — sempre gera nova.
 */
export function useConversationAISuggestion(
  type: ConversationOwnerType,
  id: string
) {
  return useMutation({
    mutationFn: (payload: AISuggestionPayload = {}) =>
      apiClient.post<AISuggestionResult>(
        `/api/v1/conversations/${type}/${id}/ai/suggest-reply`,
        payload
      ),
  });
}

export interface ReceptionReplyResult {
  reply: string;
  action: 'ask' | 'answer' | 'handle_objection' | 'propose_schedule' | 'handoff';
  handoffReason?: string;
  discloseAI: boolean;
  model: string;
}

export type AutomationMode = 'off' | 'copilot' | 'auto';

export interface ConversationAutomationView {
  ownerType: string;
  ownerId: string;
  mode: AutomationMode;
  fallbackMinutes: number;
  pausedUntil?: string;
  globallyEnabled: boolean;
}

/** Modo do recepcionista virtual desta conversa (off|copilot|auto) + fallback. */
export function useConversationAutomation(
  type: ConversationOwnerType,
  id: string
) {
  return useQuery({
    queryKey: ['conversation-automation', type, id],
    queryFn: () =>
      apiClient.get<ConversationAutomationView>(
        `/api/v1/conversations/${type}/${id}/automation`
      ),
    staleTime: 30_000,
  });
}

export function useSetConversationAutomation(
  type: ConversationOwnerType,
  id: string
) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (body: { mode: AutomationMode; fallbackMinutes?: number }) =>
      apiClient.put<ConversationAutomationView>(
        `/api/v1/conversations/${type}/${id}/automation`,
        body
      ),
    onSuccess: () =>
      qc.invalidateQueries({ queryKey: ['conversation-automation', type, id] }),
  });
}

export interface ReceptionMetrics {
  periodDays: number;
  autoReplies: number;
  handoffs: number;
  conversationsWithBot: number;
  convertedAfterBot: number;
  byMode: Record<string, number>;
  globallyEnabled: boolean;
}

/** Métricas de atuação do recepcionista virtual (últimos N dias). */
export function useReceptionMetrics(days = 30) {
  return useQuery({
    queryKey: ['reception-metrics', days],
    queryFn: () =>
      apiClient.get<ReceptionMetrics>(`/api/v1/conversations/ai/metrics?days=${days}`),
    staleTime: 60_000,
  });
}

/**
 * Recepcionista virtual: gera a próxima mensagem ancorada no script da recepção +
 * banco de objeções + guardrails (CFM/LGPD/marca). No Copiloto, o atendente revisa o
 * `reply` e envia. `action=handoff` indica que o caso deve ir para um humano.
 */
export function useConversationReceptionReply(
  type: ConversationOwnerType,
  id: string
) {
  return useMutation({
    mutationFn: () =>
      apiClient.post<ReceptionReplyResult>(
        `/api/v1/conversations/${type}/${id}/ai/reception-reply`,
        {}
      ),
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
