/**
 * API client para sessões do Escore Plenya Light na área do paciente.
 */
import { apiClient } from "@/lib/api-client";

export type LightGroupResult = {
  groupId: string;
  groupName: string;
  actualPoints: number;
  possiblePoints: number;
  scorePercentage: number;
  itemsEvaluatedCount: number;
};

export type LightSnapshot = {
  totalActualPoints: number;
  totalPossiblePoints: number;
  totalScorePercentage: number;
  itemsEvaluatedCount: number;
  itemsNotEvaluatedCount: number;
  groupResults: LightGroupResult[];
};

export type LightSession = {
  publicCode: string;
  age: number;
  gender: string;
  postMenopause?: boolean;
  height?: number;
  weight?: number;
  createdAt: string;
  expiresAt?: string;
  isClaimed: boolean;
  snapshot?: LightSnapshot;
};

export const scoreLightApi = {
  mySessions: () => apiClient.get<LightSession[]>("/api/v1/score-light/my-sessions"),
};
