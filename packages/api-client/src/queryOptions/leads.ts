import { queryOptions } from '@tanstack/react-query';
import { api } from '../fetcher';
import { queryKeys } from '../queryKeys';

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

export const leadStatusLabels: Record<LeadStatus, string> = {
  new: 'Novo',
  contacted: 'Contatado',
  qualified: 'Qualificado',
  converted: 'Convertido',
  lost: 'Perdido',
  unsubscribed: 'Descadastrado',
};

export const leadSourceLabels: Record<LeadSource, string> = {
  light_claim: 'Escore Light',
  contact_form: 'Form contato',
  whatsapp_inbound: 'WhatsApp',
  email_inbound: 'Email',
  newsletter: 'Newsletter',
  manual: 'Manual',
};

export interface LeadAssignedUser {
  id: string;
  name: string;
}

export interface LeadSummary {
  id: string;
  name?: string;
  phone?: string;
  email?: string;
  source: LeadSource;
  status: LeadStatus;
  /** @deprecated alias antigo — use status */
  stage?: string;
  assignedToUserId?: string;
  assignedTo?: LeadAssignedUser;
  lastInboundAt?: string;
  createdAt: string;
  updatedAt: string;
}

export interface LeadActivity {
  id: string;
  kind: 'whatsapp_inbound' | 'whatsapp_outbound' | 'email' | 'note' | 'stage_change';
  content: string;
  createdAt: string;
  actorName?: string;
}

export interface LeadDetail extends LeadSummary {
  activities: LeadActivity[];
  notes?: string;
  convertedPatientId?: string;
}

export interface LeadDashboard {
  byStage: Record<string, number>;
  bySource: Record<LeadSource, number>;
  newInLast24h: number;
}

export interface LeadListParams {
  search?: string;
  status?: LeadStatus;
  source?: LeadSource;
  assignedToUserId?: string;
  page?: number;
  pageSize?: number;
}

export interface LeadListResult {
  items: LeadSummary[];
  total: number;
  pageIndex: number;
  pageSize: number;
  totalPages: number;
}

export const leadsListOptions = (params: LeadListParams = {}) =>
  queryOptions({
    queryKey: queryKeys.leads.list(params as Record<string, unknown>),
    queryFn: ({ signal }) => {
      const qs = new URLSearchParams();
      if (params.search) qs.set('search', params.search);
      if (params.status) qs.set('status', params.status);
      if (params.source) qs.set('source', params.source);
      if (params.assignedToUserId) qs.set('assignedToUserId', params.assignedToUserId);
      qs.set('page', String(params.page ?? 0));
      qs.set('pageSize', String(params.pageSize ?? 50));
      return api.get<LeadListResult>(`/api/v1/leads?${qs.toString()}`, { signal });
    },
  });

export const leadOptions = (id: string) =>
  queryOptions({
    queryKey: queryKeys.leads.detail(id),
    queryFn: ({ signal }) => api.get<LeadDetail>(`/api/v1/leads/${id}`, { signal }),
    enabled: Boolean(id),
    refetchInterval: 15_000,
  });

export const leadsDashboardOptions = () =>
  queryOptions({
    queryKey: queryKeys.leads.dashboard(),
    queryFn: ({ signal }) => api.get<LeadDashboard>('/api/v1/leads/dashboard', { signal }),
    staleTime: 60_000,
  });

export interface LeadUpdateInput {
  status?: LeadStatus;
  assignedToUserId?: string | null;
  name?: string;
}

export const leadMutations = {
  convert: (id: string) =>
    api.post<{ patientId: string }>(`/api/v1/leads/${id}/convert`, {}),
  addNote: (id: string, content: string) =>
    api.post<LeadActivity>(`/api/v1/leads/${id}/notes`, { content }),
  update: (id: string, body: LeadUpdateInput) =>
    api.patch<LeadDetail>(`/api/v1/leads/${id}`, body),
  /** @deprecated use conversationMutations.sendEmail / sendWhatsApp */
  sendReply: (id: string, body: { channel: 'whatsapp' | 'email'; content: string }) =>
    api.post<LeadActivity>(`/api/v1/leads/${id}/reply`, body),
};
