/**
 * Medical Record (prontuário agregado) API hooks.
 *
 * Query UNION sobre 8 modelos de evento do paciente. Filtros: tipo, autor,
 * janela de datas. Paginação obrigatória.
 */
import { useInfiniteQuery, useQuery } from '@tanstack/react-query';
import { apiClient } from '../api-client';

export type MedicalRecordType =
  | 'anamnesis'
  | 'lab_batch'
  | 'score_snapshot'
  | 'physical_assessment'
  | 'fitness_test'
  | 'postural_assessment'
  | 'prescription'
  | 'appointment';

export interface MedicalRecordEntry {
  id: string;
  type: MedicalRecordType;
  eventDate: string; // ISO
  title: string;
  subtitle?: string;
  status?: string;
  hasNote?: boolean; // só p/ type=appointment: existe nota de evolução vinculada
  authorId: string;
  authorName?: string;
  authorRoles?: string; // CSV
  patientId: string;
  createdAt: string;
}

export const MR_TYPE_LABELS: Record<MedicalRecordType, string> = {
  anamnesis: 'Anamnese',
  lab_batch: 'Exames laboratoriais',
  score_snapshot: 'Escore Plenya',
  physical_assessment: 'Avaliação Física',
  fitness_test: 'Teste de Fitness',
  postural_assessment: 'Avaliação Postural',
  prescription: 'Prescrição',
  appointment: 'Consulta',
};

export const MR_TYPE_COLORS: Record<MedicalRecordType, string> = {
  anamnesis: 'bg-violet-50 text-violet-700 border-violet-200',
  lab_batch: 'bg-orange-50 text-orange-700 border-orange-200',
  score_snapshot: 'bg-amber-50 text-amber-700 border-amber-200',
  physical_assessment: 'bg-cyan-50 text-cyan-700 border-cyan-200',
  fitness_test: 'bg-teal-50 text-teal-700 border-teal-200',
  postural_assessment: 'bg-sky-50 text-sky-700 border-sky-200',
  prescription: 'bg-pink-50 text-pink-700 border-pink-200',
  appointment: 'bg-blue-50 text-blue-700 border-blue-200',
};

export interface AggregateFilter {
  types?: MedicalRecordType[];
  authors?: string[];
  from?: string;
  to?: string;
}

const PAGE_SIZE = 30;

export function useMedicalRecord(patientId: string | null | undefined, filter: AggregateFilter = {}) {
  return useInfiniteQuery({
    queryKey: ['medical-record', patientId, filter],
    queryFn: ({ pageParam = 0 }) => {
      const params = new URLSearchParams();
      params.set('limit', String(PAGE_SIZE));
      params.set('offset', String(pageParam));
      if (filter.types?.length) params.set('types', filter.types.join(','));
      if (filter.authors?.length) params.set('authors', filter.authors.join(','));
      if (filter.from) params.set('from', filter.from);
      if (filter.to) params.set('to', filter.to);
      return apiClient.get<MedicalRecordEntry[]>(
        `/api/v1/patients/${patientId}/prontuario?${params.toString()}`,
      );
    },
    initialPageParam: 0,
    getNextPageParam: (lastPage, allPages) => {
      if (!lastPage || lastPage.length < PAGE_SIZE) return undefined;
      return allPages.length * PAGE_SIZE;
    },
    enabled: !!patientId,
  });
}
