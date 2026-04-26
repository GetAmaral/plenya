import { api } from '../fetcher';

export interface TrainingAIChatMessage {
  role: 'user' | 'assistant';
  content: string;
}

export interface TrainingAIPatientContext {
  name: string;
  age: number;
  gender: string;
  weight?: number;
  height?: number;
  bmi?: number;
  objective?: string;
  riskLevel?: string;
  riskFactors?: string[];
}

export interface TrainingAIChatRequest {
  messages: TrainingAIChatMessage[];
  patientContext?: TrainingAIPatientContext;
}

export interface TrainingAIChatResponse {
  message: string;
  sources?: string[];
}

export const trainingAiMutations = {
  chat: (body: TrainingAIChatRequest) =>
    api.post<TrainingAIChatResponse>('/api/v1/training/ai/chat', body),
};
