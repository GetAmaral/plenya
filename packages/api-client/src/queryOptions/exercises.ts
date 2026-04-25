import { queryOptions } from '@tanstack/react-query';
import { api } from '../fetcher';
import { queryKeys } from '../queryKeys';

export interface ExerciseSummary {
  id: string;
  name: string;
  category?: string;
  muscleGroups?: string[];
  equipment?: string[];
  thumbnailUrl?: string;
  gifUrl?: string;
}

export interface NSCAReference {
  citation: string;
  url?: string;
}

export interface ExerciseDetail extends ExerciseSummary {
  description?: string;
  instructionsHtml?: string;
  biomechanicsData?: Record<string, unknown>;
  programDesign?: Record<string, unknown>;
  nscaReferences?: NSCAReference[];
  videoUrl?: string;
  difficulty?: 'beginner' | 'intermediate' | 'advanced';
}

export const exercisesListOptions = (params: { search?: string } = {}) =>
  queryOptions({
    queryKey: [...queryKeys.exercises(), params] as const,
    queryFn: ({ signal }) => {
      const qs = new URLSearchParams();
      if (params.search) qs.set('search', params.search);
      const suffix = qs.toString() ? `?${qs.toString()}` : '';
      return api.get<ExerciseSummary[]>(`/api/v1/exercises${suffix}`, { signal });
    },
  });

export const exerciseDetailOptions = (id: string) =>
  queryOptions({
    queryKey: [...queryKeys.exercises(), 'detail', id] as const,
    queryFn: ({ signal }) => api.get<ExerciseDetail>(`/api/v1/exercises/${id}`, { signal }),
    enabled: Boolean(id),
  });
