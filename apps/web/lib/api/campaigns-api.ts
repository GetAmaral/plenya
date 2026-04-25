import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { apiClient } from '../api-client';

export type CampaignStatus = 'active' | 'archived';

export interface CampaignStats {
  sessionsCount: number;
  leadsCount: number;
}

export interface Campaign {
  id: string;
  name: string;
  slug: string;
  description?: string | null;
  landingPath: string;
  utmSource: string;
  utmMedium: string;
  utmCampaign: string;
  utmTerm?: string | null;
  status: CampaignStatus;
  createdByUserId?: string | null;
  createdAt: string;
  updatedAt: string;
  url: string;
  stats?: CampaignStats;
}

export interface CreateCampaignInput {
  name: string;
  slug?: string;
  description?: string;
  landingPath: string;
  utmSource: string;
  utmMedium: string;
  utmCampaign?: string;
  utmTerm?: string;
}

export type UpdateCampaignInput = Partial<CreateCampaignInput> & {
  status?: CampaignStatus;
};

export const campaignKeys = {
  all: ['campaigns'] as const,
  list: (includeArchived: boolean) => [...campaignKeys.all, 'list', { includeArchived }] as const,
  detail: (id: string) => [...campaignKeys.all, 'detail', id] as const,
};

export function useCampaigns(includeArchived = false) {
  return useQuery({
    queryKey: campaignKeys.list(includeArchived),
    queryFn: () =>
      apiClient.get<Campaign[]>(
        `/api/v1/campaigns${includeArchived ? '?includeArchived=true' : ''}`,
      ),
  });
}

export function useCampaign(id: string | undefined) {
  return useQuery({
    queryKey: id ? campaignKeys.detail(id) : ['campaigns', 'detail', 'noop'],
    queryFn: () => apiClient.get<Campaign>(`/api/v1/campaigns/${id}`),
    enabled: !!id,
  });
}

export function useCreateCampaign() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (input: CreateCampaignInput) =>
      apiClient.post<Campaign>('/api/v1/campaigns', input),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: campaignKeys.all });
    },
  });
}

export function useUpdateCampaign(id: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (input: UpdateCampaignInput) =>
      apiClient.patch<Campaign>(`/api/v1/campaigns/${id}`, input),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: campaignKeys.all });
      qc.invalidateQueries({ queryKey: campaignKeys.detail(id) });
    },
  });
}

export function useArchiveCampaign() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: string) =>
      apiClient.post<void>(`/api/v1/campaigns/${id}/archive`, {}),
    onSuccess: () => qc.invalidateQueries({ queryKey: campaignKeys.all }),
  });
}

export function useDeleteCampaign() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: string) =>
      apiClient.delete<void>(`/api/v1/campaigns/${id}`),
    onSuccess: () => qc.invalidateQueries({ queryKey: campaignKeys.all }),
  });
}

/** Hook que baixa o QR code (PNG) autenticado e devolve como object URL pra usar em <img>. */
export function useCampaignQRCode(id: string | undefined, sizePx = 512) {
  return useQuery({
    queryKey: id ? [...campaignKeys.detail(id), 'qrcode', sizePx] : ['campaigns', 'qr', 'noop'],
    queryFn: async () => {
      const blob = await apiClient.getBlob(`/api/v1/campaigns/${id}/qrcode?size=${sizePx}`);
      return URL.createObjectURL(blob);
    },
    enabled: !!id,
    staleTime: Infinity, // QR de uma campanha existente é estável
  });
}

/** URL direta (sem auth) — só útil em contextos server-side ou pra debug; navegador NÃO consegue usar em <img> pq exige Bearer token. */
export function campaignQRCodeRawURL(id: string, sizePx = 512): string {
  const base = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:3001';
  return `${base}/api/v1/campaigns/${id}/qrcode?size=${sizePx}`;
}
