import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { apiClient } from '../api-client';

export type LeadSource =
  | 'light_claim'
  | 'contact_form'
  | 'whatsapp_inbound'
  | 'email_inbound'
  | 'newsletter'
  | 'manual';

export type LeadStatus =
  | 'new'
  | 'contacted'
  | 'qualified'
  | 'converted'
  | 'lost'
  | 'unsubscribed';

export type LeadActivityType =
  | 'created'
  | 'message_sent'
  | 'message_received'
  | 'message_status_changed'
  | 'status_changed'
  | 'note_added'
  | 'converted'
  | 'assigned'
  | 'unsubscribed';

export type LeadActivityChannel = 'email' | 'whatsapp' | 'internal';

export interface LeadActivity {
  id: string;
  leadId: string;
  type: LeadActivityType;
  channel: LeadActivityChannel;
  content?: string;
  metadata?: Record<string, unknown>;
  actorUserId?: string;
  actor?: { id: string; name: string; email: string };
  createdAt: string;
}

export interface Lead {
  id: string;
  source: LeadSource;
  status: LeadStatus;
  name?: string;
  email?: string;
  phone?: string;
  message?: string;
  metadata?: Record<string, unknown>;
  emailOptIn: boolean;
  whatsAppOptIn: boolean;
  newsletterOptIn: boolean;
  consentVersion?: string;
  consentTimestamp?: string;
  anonymousScoreSessionId?: string;
  convertedPatientId?: string;
  convertedAt?: string;
  convertedByUserId?: string;
  assignedToUserId?: string;
  /** Última inbound do cliente — usado pra calcular janela 24h de session messages. */
  lastInboundAt?: string;
  createdAt: string;
  updatedAt: string;
  activities?: LeadActivity[];
  assignedTo?: { id: string; name: string; email: string };
  convertedPatient?: { id: string; name: string };
}

export interface LeadFilter {
  source?: LeadSource;
  status?: LeadStatus;
  search?: string;
  hasEmailOptIn?: boolean;
  hasWhatsAppOptIn?: boolean;
  assignedToUserId?: string;
}

export interface LeadListResult {
  items: Lead[];
  total: number;
  pageIndex: number;
  pageSize: number;
  totalPages: number;
}

export const leadKeys = {
  all: ['leads'] as const,
  list: (filter: LeadFilter, page: number, pageSize: number) =>
    [...leadKeys.all, 'list', filter, page, pageSize] as const,
  detail: (id: string) => [...leadKeys.all, 'detail', id] as const,
};

function buildQuery(filter: LeadFilter, page: number, pageSize: number) {
  const params = new URLSearchParams();
  if (filter.source) params.set('source', filter.source);
  if (filter.status) params.set('status', filter.status);
  if (filter.search) params.set('search', filter.search);
  if (filter.hasEmailOptIn !== undefined) params.set('hasEmailOptIn', String(filter.hasEmailOptIn));
  if (filter.hasWhatsAppOptIn !== undefined)
    params.set('hasWhatsAppOptIn', String(filter.hasWhatsAppOptIn));
  if (filter.assignedToUserId) params.set('assignedToUserId', filter.assignedToUserId);
  params.set('page', String(page));
  params.set('pageSize', String(pageSize));
  return params.toString();
}

export function useLeads(
  filter: LeadFilter,
  page = 0,
  pageSize = 25,
  options?: { enabled?: boolean }
) {
  return useQuery({
    queryKey: leadKeys.list(filter, page, pageSize),
    queryFn: () => apiClient.get<LeadListResult>(`/api/v1/leads?${buildQuery(filter, page, pageSize)}`),
    enabled: options?.enabled ?? true,
  });
}

export function useLead(id: string) {
  return useQuery({
    queryKey: leadKeys.detail(id),
    queryFn: () => apiClient.get<Lead>(`/api/v1/leads/${id}`),
    enabled: !!id,
  });
}

export interface UpdateLeadPatch {
  status?: LeadStatus;
  assignedToUserId?: string;
  name?: string;
}

export function useUpdateLead(id: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (patch: UpdateLeadPatch) =>
      apiClient.patch<Lead>(`/api/v1/leads/${id}`, patch),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: leadKeys.detail(id) });
      qc.invalidateQueries({ queryKey: leadKeys.all });
    },
  });
}

export function useAddLeadNote(id: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (note: string) => apiClient.post<void>(`/api/v1/leads/${id}/notes`, { note }),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: leadKeys.detail(id) });
    },
  });
}

export interface ConvertLeadPayload {
  name?: string;
  email?: string;
  phone?: string;
  birthDate?: string; // ISO
  gender?: 'male' | 'female' | 'other';
}

export function useConvertLead(id: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (payload: ConvertLeadPayload) =>
      apiClient.post<{ id: string; name: string }>(`/api/v1/leads/${id}/convert`, payload),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: leadKeys.detail(id) });
      qc.invalidateQueries({ queryKey: leadKeys.all });
    },
  });
}

export function useDeleteLead(id: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: () => apiClient.delete<void>(`/api/v1/leads/${id}`),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: leadKeys.all });
    },
  });
}

export type StatsPeriodKey = '7d' | '30d' | '90d' | 'custom';

export interface LeadStats {
  period: { from: string; to: string; label: string };
  totals: Record<string, number>;
  bySource: Record<string, number>;
  conversionRate: number;
  avgTimeToContact?: number; // segundos
  avgTimeToConvert?: number; // segundos
  byDay: Array<{ date: string; count: number }>;
}

export function useLeadStats(period: StatsPeriodKey = '30d', from?: string, to?: string) {
  return useQuery({
    queryKey: ['lead-stats', period, from, to],
    queryFn: () => {
      const params = new URLSearchParams({ period });
      if (period === 'custom') {
        if (from) params.set('from', from);
        if (to) params.set('to', to);
      }
      return apiClient.get<LeadStats>(`/api/v1/leads/stats?${params.toString()}`);
    },
    staleTime: 60 * 1000,
  });
}

export interface StaffUser {
  id: string;
  name: string;
  email: string;
  roles: string[];
}

/**
 * Lista de usuários staff (admin/secretary/manager) elegíveis pra atribuição de leads.
 * Endpoint dedicado /api/v1/users/staff (acessível por qualquer staff, não exige admin).
 */
export function useStaffUsers() {
  return useQuery({
    queryKey: ['users', 'staff'],
    queryFn: () => apiClient.get<StaffUser[]>('/api/v1/users/staff'),
    staleTime: 5 * 60 * 1000, // 5min — staff list muda raramente
  });
}

export function useSendLeadMessage(id: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (content: string) =>
      apiClient.post<LeadActivity>(`/api/v1/leads/${id}/messages`, { content }),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: leadKeys.detail(id) });
    },
  });
}

// =====================================================
// Email reply (Bloco 3) — responder email do lead
// =====================================================

export type SendEmailReplyPayload = {
  /** Vazio → backend monta "Re: <subject anterior>" via histórico. */
  subject?: string;
  bodyText: string;
  /** Opcional: HTML pré-renderizado pela UI. Se vazio, backend gera a partir do bodyText. */
  bodyHTML?: string;
  /** Message-ID do email respondido — pra threading. */
  inReplyTo?: string;
  /** Chain de Message-IDs anteriores (RFC 5322 References). */
  references?: string[];
};

export type SendEmailReplyResponse = {
  message: string;
  to?: string;
};

/**
 * Envia resposta por email via Resend, espelha em Sent (IMAP) e cria activity message_sent.
 * Backend rejeita com 422 se Lead.emailOptIn=false ou Lead.email=null.
 */
export function useSendLeadEmailReply(id: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (payload: SendEmailReplyPayload) =>
      apiClient.post<SendEmailReplyResponse>(`/api/v1/leads/${id}/email-reply`, payload),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: leadKeys.detail(id) });
    },
  });
}

// Helpers visuais
export const SOURCE_LABELS: Record<LeadSource, string> = {
  light_claim: 'Escore Light',
  contact_form: 'Formulário /contato',
  whatsapp_inbound: 'WhatsApp inbound',
  email_inbound: 'Email inbound',
  newsletter: 'Newsletter',
  manual: 'Manual',
};

export const STATUS_LABELS: Record<LeadStatus, string> = {
  new: 'Novo',
  contacted: 'Contatado',
  qualified: 'Qualificado',
  converted: 'Convertido',
  lost: 'Perdido',
  unsubscribed: 'Descadastrou',
};

export const STATUS_COLORS: Record<LeadStatus, string> = {
  new: 'bg-blue-100 text-blue-900 border-blue-200',
  contacted: 'bg-amber-100 text-amber-900 border-amber-200',
  qualified: 'bg-purple-100 text-purple-900 border-purple-200',
  converted: 'bg-emerald-100 text-emerald-900 border-emerald-200',
  lost: 'bg-stone-100 text-stone-700 border-stone-200',
  unsubscribed: 'bg-rose-100 text-rose-900 border-rose-200',
};
