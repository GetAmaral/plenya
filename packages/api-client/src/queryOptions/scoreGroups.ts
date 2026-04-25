import { queryOptions } from '@tanstack/react-query';
import { api } from '../fetcher';
import { queryKeys } from '../queryKeys';

export interface ScoreLevelNode {
  id: string;
  level: string;
  color: string;
  operatorType: string;
  minValue?: number;
  maxValue?: number;
  textValue?: string;
  description?: string;
}

export interface ScoreItemNode {
  id: string;
  name: string;
  description?: string;
  clinicalRelevance?: string;
  patientExplanation?: string;
  conduct?: string;
  levels: ScoreLevelNode[];
}

export interface ScoreSubgroupNode {
  id: string;
  name: string;
  items: ScoreItemNode[];
}

export interface ScoreGroupNode {
  id: string;
  name: string;
  color?: string;
  subgroups: ScoreSubgroupNode[];
}

export interface PatientScoreResult {
  itemId: string;
  levelId?: string;
  value?: number | string;
  takenAt: string;
}

export const scoreTreeOptions = () =>
  queryOptions({
    queryKey: queryKeys.scoreGroups.tree(),
    queryFn: ({ signal }) =>
      api.get<ScoreGroupNode[]>('/api/v1/score-groups/tree', { signal }),
    staleTime: 10 * 60_000,
  });

export const patientScoresOptions = (patientId: string) =>
  queryOptions({
    queryKey: queryKeys.patients.scores(patientId),
    queryFn: ({ signal }) =>
      api.get<PatientScoreResult[]>(`/api/v1/patients/${patientId}/scores`, { signal }),
    enabled: Boolean(patientId),
  });
