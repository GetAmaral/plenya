/**
 * Continuum API hooks (Fase 1).
 *
 * Cobre:
 *  - Templates de programa Continuum (CRUD + clone)
 *  - Templates de Box (CRUD + clone)
 *
 * Fases 2+ adicionam: inscrição, items, plano integrado, dashboard.
 */
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { apiClient } from '../api-client';

// =====================================================
// Types
// =====================================================

export type ContinuumItemType =
  | 'appointment'
  | 'box'
  | 'reassessment'
  | 'milestone'
  | 'custom';

export type ContinuumItemSpecialty =
  | 'doctor'
  | 'nutritionist'
  | 'psychologist'
  | 'physicalEducator';

export interface ContinuumTemplateItem {
  id?: string;
  type: ContinuumItemType;
  specialty?: ContinuumItemSpecialty | null;
  title: string;
  description?: string;
  weekOffset: number;
  expectedOffsetDays: number;
  lateAfterDays: number;
  boxTemplateId?: string | null;
  position: number;
}

export interface ContinuumTemplate {
  id: string;
  name: string;
  description: string;
  durationWeeks: number;
  status: 'active' | 'archived';
  createdByUserId: string;
  createdAt: string;
  updatedAt: string;
  items?: ContinuumTemplateItem[];
}

export interface ContinuumTemplatePayload {
  name: string;
  description?: string;
  durationWeeks: number;
  status?: 'active' | 'archived';
  items: ContinuumTemplateItem[];
}

export interface ContinuumBoxTemplate {
  id: string;
  name: string;
  description: string;
  contents: string;
  notes: string;
  status: 'active' | 'archived';
  createdByUserId: string;
  createdAt: string;
  updatedAt: string;
}

export interface ContinuumBoxTemplatePayload {
  name: string;
  description?: string;
  contents?: string;
  notes?: string;
  status?: 'active' | 'archived';
}

// =====================================================
// Query Keys
// =====================================================

export const continuumKeys = {
  templates: () => ['continuum', 'templates'] as const,
  template: (id: string) => ['continuum', 'templates', id] as const,
  boxTemplates: () => ['continuum', 'box-templates'] as const,
  boxTemplate: (id: string) => ['continuum', 'box-templates', id] as const,
};

// =====================================================
// Templates de Programa
// =====================================================

export function useContinuumTemplates(includeArchived = false) {
  return useQuery({
    queryKey: [...continuumKeys.templates(), { includeArchived }],
    queryFn: () =>
      apiClient.get<ContinuumTemplate[]>(
        `/api/v1/continuum/templates?includeArchived=${includeArchived}`,
      ),
  });
}

export function useContinuumTemplate(id: string | null | undefined) {
  return useQuery({
    queryKey: continuumKeys.template(id ?? ''),
    queryFn: () =>
      apiClient.get<ContinuumTemplate>(`/api/v1/continuum/templates/${id}`),
    enabled: !!id,
  });
}

export function useCreateContinuumTemplate() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (payload: ContinuumTemplatePayload) =>
      apiClient.post<ContinuumTemplate>('/api/v1/continuum/templates', payload),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: continuumKeys.templates() });
    },
  });
}

export function useUpdateContinuumTemplate(id: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (payload: ContinuumTemplatePayload) =>
      apiClient.put<ContinuumTemplate>(`/api/v1/continuum/templates/${id}`, payload),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: continuumKeys.templates() });
      qc.invalidateQueries({ queryKey: continuumKeys.template(id) });
    },
  });
}

export function useDeleteContinuumTemplate() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: string) => apiClient.delete(`/api/v1/continuum/templates/${id}`),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: continuumKeys.templates() });
    },
  });
}

export function useCloneContinuumTemplate() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ id, name }: { id: string; name: string }) =>
      apiClient.post<ContinuumTemplate>(`/api/v1/continuum/templates/${id}/clone`, { name }),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: continuumKeys.templates() });
    },
  });
}

// =====================================================
// Templates de Box
// =====================================================

export function useContinuumBoxTemplates(includeArchived = false) {
  return useQuery({
    queryKey: [...continuumKeys.boxTemplates(), { includeArchived }],
    queryFn: () =>
      apiClient.get<ContinuumBoxTemplate[]>(
        `/api/v1/continuum/box-templates?includeArchived=${includeArchived}`,
      ),
  });
}

export function useContinuumBoxTemplate(id: string | null | undefined) {
  return useQuery({
    queryKey: continuumKeys.boxTemplate(id ?? ''),
    queryFn: () =>
      apiClient.get<ContinuumBoxTemplate>(`/api/v1/continuum/box-templates/${id}`),
    enabled: !!id,
  });
}

export function useCreateContinuumBoxTemplate() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (payload: ContinuumBoxTemplatePayload) =>
      apiClient.post<ContinuumBoxTemplate>('/api/v1/continuum/box-templates', payload),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: continuumKeys.boxTemplates() });
    },
  });
}

export function useUpdateContinuumBoxTemplate(id: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (payload: ContinuumBoxTemplatePayload) =>
      apiClient.put<ContinuumBoxTemplate>(`/api/v1/continuum/box-templates/${id}`, payload),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: continuumKeys.boxTemplates() });
      qc.invalidateQueries({ queryKey: continuumKeys.boxTemplate(id) });
    },
  });
}

export function useDeleteContinuumBoxTemplate() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: string) => apiClient.delete(`/api/v1/continuum/box-templates/${id}`),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: continuumKeys.boxTemplates() });
    },
  });
}

export function useCloneContinuumBoxTemplate() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ id, name }: { id: string; name: string }) =>
      apiClient.post<ContinuumBoxTemplate>(`/api/v1/continuum/box-templates/${id}/clone`, {
        name,
      }),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: continuumKeys.boxTemplates() });
    },
  });
}

// =====================================================
// Inscrição de paciente (Fase 2)
// =====================================================

export type ContinuumStatus = 'active' | 'paused' | 'completed' | 'cancelled';

export type ContinuumItemStatus =
  | 'pending'
  | 'scheduled'
  | 'completed'
  | 'missed'
  | 'cancelled'
  | 'skipped';

export interface PatientContinuumItem {
  id: string;
  continuumId: string;
  type: ContinuumItemType;
  specialty?: ContinuumItemSpecialty | null;
  title: string;
  description?: string;
  weekOffset: number;
  expectedDate: string; // ISO
  lateAfterDate: string;
  status: ContinuumItemStatus;
  appointmentId?: string | null;
  boxId?: string | null;
  completedAt?: string | null;
  position: number;
}

export interface PatientContinuum {
  id: string;
  patientId: string;
  templateId: string;
  status: ContinuumStatus;
  startDate: string;
  endDate: string;
  coordinatorDoctorId?: string | null;
  integratedPlanMarkdown?: string;
  integratedPlanUpdatedAt?: string | null;
  integratedPlanUpdatedBy?: string | null;
  whatsappGroupName?: string | null;
  whatsappGroupInviteLink?: string | null;
  notes?: string;
  createdAt: string;
  updatedAt: string;
  items?: PatientContinuumItem[];
  patient?: { id: string; name: string };
  coordinatorDoctor?: { id: string; name: string };
}

export interface EnrollPayload {
  templateId: string;
  startDate: string; // YYYY-MM-DD
  coordinatorDoctorId?: string;
  notes?: string;
}

export interface UpdateItemPayload {
  status?: ContinuumItemStatus;
  appointmentId?: string;
  completedAt?: string;
}

export const patientContinuumKeys = {
  byPatient: (patientId: string) => ['continuum', 'patient', patientId] as const,
  enrollment: (id: string) => ['continuum', 'enrollment', id] as const,
};

export function usePatientContinuums(patientId: string | null | undefined) {
  return useQuery({
    queryKey: patientContinuumKeys.byPatient(patientId ?? ''),
    queryFn: () =>
      apiClient.get<PatientContinuum[]>(`/api/v1/patients/${patientId}/continuum`),
    enabled: !!patientId,
  });
}

export function useContinuumEnrollment(id: string | null | undefined) {
  return useQuery({
    queryKey: patientContinuumKeys.enrollment(id ?? ''),
    queryFn: () =>
      apiClient.get<PatientContinuum>(`/api/v1/continuum/enrollments/${id}`),
    enabled: !!id,
  });
}

export function useEnrollPatientContinuum(patientId: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (payload: EnrollPayload) =>
      apiClient.post<PatientContinuum>(
        `/api/v1/patients/${patientId}/continuum`,
        payload,
      ),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: patientContinuumKeys.byPatient(patientId) });
    },
  });
}

export function useUpdateContinuumItem(continuumId: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ itemId, payload }: { itemId: string; payload: UpdateItemPayload }) =>
      apiClient.put<PatientContinuumItem>(`/api/v1/continuum/items/${itemId}`, payload),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: patientContinuumKeys.enrollment(continuumId) });
    },
  });
}

// === Dashboard (Fase 7) ===

export interface PerPatientRow {
  continuumId: string;
  patientId: string;
  patientName: string;
  startDate: string;
  endDate: string;
  durationWeeks: number;
  currentWeek: number;
  totalItems: number;
  completedItems: number;
  scheduledItems: number;
  missedItems: number;
  pendingItems: number;
  nextItemTitle?: string;
  nextItemDate?: string;
  coordinatorDoctorId?: string;
  coordinatorName?: string;
}

export interface PerWeekItem {
  id: string;
  continuumId: string;
  type: ContinuumItemType;
  specialty?: ContinuumItemSpecialty;
  title: string;
  expectedDate: string;
  status: ContinuumItemStatus;
  patientId: string;
  patientName: string;
  appointmentId?: string;
  boxId?: string;
}

export interface AlertRow {
  id: string;
  continuumId: string;
  patientId: string;
  patientName: string;
  type: ContinuumItemType;
  specialty?: ContinuumItemSpecialty;
  title: string;
  status: ContinuumItemStatus;
  expectedDate: string;
  lateAfterDate: string;
  severity: 'missed' | 'due-soon';
  appointmentId?: string;
}

export function useContinuumDashboardPatients() {
  return useQuery({
    queryKey: ['continuum', 'dashboard', 'patients'],
    queryFn: () => apiClient.get<PerPatientRow[]>('/api/v1/continuum/dashboard/patients'),
  });
}

export function useContinuumDashboardWeek(weekStart?: string) {
  const qs = weekStart ? `?start=${weekStart}` : '';
  return useQuery({
    queryKey: ['continuum', 'dashboard', 'week', { weekStart }],
    queryFn: () =>
      apiClient.get<{ weekStart: string; items: PerWeekItem[] }>(
        `/api/v1/continuum/dashboard/week${qs}`,
      ),
  });
}

export function useContinuumDashboardAlerts(dueSoonDays = 7) {
  return useQuery({
    queryKey: ['continuum', 'dashboard', 'alerts', { dueSoonDays }],
    queryFn: () =>
      apiClient.get<AlertRow[]>(`/api/v1/continuum/dashboard/alerts?dueSoonDays=${dueSoonDays}`),
  });
}

// === Box logístico (Fase 5) ===

export type BoxStatus = 'planned' | 'preparing' | 'shipped' | 'delivered' | 'cancelled';

export interface ContinuumBox {
  id: string;
  continuumItemId: string;
  name: string;
  contents: string;
  status: BoxStatus;
  preparedAt?: string | null;
  shippedAt?: string | null;
  deliveredAt?: string | null;
  trackingCode?: string | null;
  carrier?: string | null;
  addressSnapshot?: string;
  notes?: string;
  createdAt: string;
  updatedAt: string;
  // Joins
  patientId: string;
  patientName: string;
  expectedDate: string;
  weekOffset: number;
  continuumId: string;
}

export interface UpdateBoxPayload {
  status?: BoxStatus;
  trackingCode?: string;
  carrier?: string;
  notes?: string;
  contents?: string;
  address?: string;
}

export const BOX_STATUS_LABELS: Record<BoxStatus, string> = {
  planned: 'Planejado',
  preparing: 'Preparando',
  shipped: 'Enviado',
  delivered: 'Entregue',
  cancelled: 'Cancelado',
};

export const BOX_STATUS_COLORS: Record<BoxStatus, string> = {
  planned: 'bg-slate-100 text-slate-700 border-slate-200',
  preparing: 'bg-amber-50 text-amber-700 border-amber-200',
  shipped: 'bg-blue-50 text-blue-700 border-blue-200',
  delivered: 'bg-emerald-50 text-emerald-700 border-emerald-200',
  cancelled: 'bg-zinc-100 text-zinc-500 border-zinc-200',
};

export function useContinuumBoxes(statuses?: BoxStatus[]) {
  const qs = statuses && statuses.length > 0 ? `?status=${statuses.join(',')}` : '';
  return useQuery({
    queryKey: ['continuum', 'boxes', { statuses }],
    queryFn: () => apiClient.get<ContinuumBox[]>(`/api/v1/continuum/boxes${qs}`),
  });
}

export function useContinuumBoxCounts() {
  return useQuery({
    queryKey: ['continuum', 'boxes', 'counts'],
    queryFn: () => apiClient.get<Record<BoxStatus, number>>('/api/v1/continuum/boxes/counts'),
  });
}

export function useUpdateContinuumBox() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ id, payload }: { id: string; payload: UpdateBoxPayload }) =>
      apiClient.put<ContinuumBox>(`/api/v1/continuum/boxes/${id}`, payload),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: ['continuum', 'boxes'] });
      qc.invalidateQueries({ queryKey: ['continuum', 'enrollment'] });
    },
  });
}

// === Plano integrado (Fase 4) ===

export interface IntegratedPlanRevision {
  id: string;
  continuumId: string;
  content: string;
  updatedById: string;
  createdAt: string;
  updatedBy?: { id: string; name: string };
}

export function useIntegratedPlanRevisions(continuumId: string | null | undefined) {
  return useQuery({
    queryKey: ['continuum', 'integrated-plan-revisions', continuumId],
    queryFn: () =>
      apiClient.get<IntegratedPlanRevision[]>(
        `/api/v1/continuum/enrollments/${continuumId}/integrated-plan/revisions`,
      ),
    enabled: !!continuumId,
  });
}

export function useUpdateIntegratedPlan(continuumId: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (content: string) =>
      apiClient.put<PatientContinuum>(
        `/api/v1/continuum/enrollments/${continuumId}/integrated-plan`,
        { content },
      ),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: patientContinuumKeys.enrollment(continuumId) });
      qc.invalidateQueries({ queryKey: ['continuum', 'integrated-plan-revisions', continuumId] });
    },
  });
}

export function useCancelContinuumEnrollment() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: string) =>
      apiClient.delete(`/api/v1/continuum/enrollments/${id}`),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: ['continuum', 'patient'] });
      qc.invalidateQueries({ queryKey: ['continuum', 'enrollment'] });
    },
  });
}

// =====================================================
// Helpers
// =====================================================

export const ITEM_STATUS_LABELS: Record<ContinuumItemStatus, string> = {
  pending: 'Pendente',
  scheduled: 'Agendado',
  completed: 'Concluído',
  missed: 'Atrasado',
  cancelled: 'Cancelado',
  skipped: 'Pulado',
};

export const ITEM_STATUS_COLORS: Record<ContinuumItemStatus, string> = {
  pending: 'bg-slate-100 text-slate-700 border-slate-200',
  scheduled: 'bg-blue-50 text-blue-700 border-blue-200',
  completed: 'bg-emerald-50 text-emerald-700 border-emerald-200',
  missed: 'bg-red-50 text-red-700 border-red-200',
  cancelled: 'bg-zinc-100 text-zinc-500 border-zinc-200',
  skipped: 'bg-amber-50 text-amber-700 border-amber-200',
};

export const ITEM_TYPE_LABELS: Record<ContinuumItemType, string> = {
  appointment: 'Consulta',
  box: 'Box Plenya',
  reassessment: 'Reavaliação',
  milestone: 'Marco',
  custom: 'Personalizado',
};

export const SPECIALTY_LABELS: Record<ContinuumItemSpecialty, string> = {
  doctor: 'Médico',
  nutritionist: 'Nutricionista',
  psychologist: 'Psicólogo',
  physicalEducator: 'Educador Físico',
};
