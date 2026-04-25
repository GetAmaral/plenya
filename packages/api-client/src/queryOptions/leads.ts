import { queryOptions } from '@tanstack/react-query';
import { api } from '../fetcher';
import { queryKeys } from '../queryKeys';

export type LeadSource =
  | 'light_claim'
  | 'contact_form'
  | 'whatsapp_inbound'
  | 'newsletter'
  | 'direct';

export interface LeadSummary {
  id: string;
  name: string;
  phone?: string;
  email?: string;
  source: LeadSource;
  stage: string;
  lastActivityAt: string;
  unreadCount: number;
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

export const leadsListOptions = (params: { stage?: string; search?: string } = {}) =>
  queryOptions({
    queryKey: queryKeys.leads.list(params),
    queryFn: ({ signal }) => {
      const qs = new URLSearchParams();
      if (params.stage) qs.set('stage', params.stage);
      if (params.search) qs.set('search', params.search);
      const suffix = qs.toString() ? `?${qs.toString()}` : '';
      return api.get<LeadSummary[]>(`/api/v1/leads${suffix}`, { signal });
    },
  });

export const leadOptions = (id: string) =>
  queryOptions({
    queryKey: queryKeys.leads.detail(id),
    queryFn: ({ signal }) => api.get<LeadDetail>(`/api/v1/leads/${id}`, { signal }),
    enabled: Boolean(id),
  });

export const leadsDashboardOptions = () =>
  queryOptions({
    queryKey: queryKeys.leads.dashboard(),
    queryFn: ({ signal }) => api.get<LeadDashboard>('/api/v1/leads/dashboard', { signal }),
    staleTime: 60_000,
  });

export const leadMutations = {
  convert: (id: string) =>
    api.post<{ patientId: string }>(`/api/v1/leads/${id}/convert`, {}),
  addNote: (id: string, content: string) =>
    api.post<LeadActivity>(`/api/v1/leads/${id}/notes`, { content }),
  sendReply: (id: string, body: { channel: 'whatsapp' | 'email'; content: string }) =>
    api.post<LeadActivity>(`/api/v1/leads/${id}/reply`, body),
};
