/**
 * Tipos do Escore Plenya Light — espelham os DTOs do backend Go
 * (apps/api/internal/services/anonymous_score_service.go)
 *
 * Mantenha sincronizado manualmente. Como o site é estático e o EMR é source
 * of truth, este arquivo evolui sob demanda.
 */

export type LightLevelConfig = {
  id: string;
  level: number;
  name: string;
  lowerLimit?: string;
  upperLimit?: string;
  operator: '=' | '>' | '>=' | '<' | '<=' | 'between';
  patientExplanation?: string;
};

export type LightItemConfig = {
  id: string;
  name: string;
  lightQuestion: string;
  unit?: string;
  gender?: 'male' | 'female' | 'not_applicable';
  ageRangeMin?: number;
  ageRangeMax?: number;
  postMenopause?: boolean;
  points?: number;
  lightOrder: number;
  labTestCode?: string;
  anamneseItemCode?: string;
  patientExplanation?: string;
  levels: LightLevelConfig[];
};

export type LightSubgroupConfig = {
  id: string;
  name: string;
  order: number;
  items: LightItemConfig[];
};

export type LightGroupConfig = {
  id: string;
  name: string;
  order: number;
  subgroups: LightSubgroupConfig[];
};

export type LightConfig = {
  version: string;
  generatedAt: string;
  itemCount: number;
  groups: LightGroupConfig[];
};

// === Sessão / respostas ===

export type SessionResponse = {
  scoreItemId: string;
  numericValue?: number;
  selectedLevel?: number;
  textValue?: string;
};

export type CreateSessionRequest = {
  age: number;
  gender: 'male' | 'female' | 'other';
  postMenopause?: boolean;
  height?: number;
  weight?: number;
  responses: SessionResponse[];
};

// === Resposta pública (sem dados técnicos) ===

export type PublicGroupResult = {
  groupId: string;
  groupName: string;
  actualPoints: number;
  possiblePoints: number;
  scorePercentage: number;
  itemsEvaluatedCount: number;
};

export type PublicSnapshot = {
  totalActualPoints: number;
  totalPossiblePoints: number;
  totalScorePercentage: number;
  itemsEvaluatedCount: number;
  itemsNotEvaluatedCount: number;
  groupResults: PublicGroupResult[];
};

export type PublicAnonymousScoreItem = {
  scoreItemId: string;
  question: string;
  numericValue?: number;
  selectedLevel?: number;
  textValue?: string;
};

export type PublicSession = {
  publicCode: string;
  age: number;
  gender: string;
  postMenopause?: boolean;
  height?: number;
  weight?: number;
  createdAt: string;
  expiresAt?: string;
  isClaimed: boolean;
  items: PublicAnonymousScoreItem[];
  snapshot?: PublicSnapshot;
};

export type ConfirmClaimResult = {
  sessionCode: string;
  patientId: string;
  userId: string;
  email: string;
  accessToken: string;
  refreshToken: string;
};
