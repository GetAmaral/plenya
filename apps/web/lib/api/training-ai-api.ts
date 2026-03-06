import { useMutation } from '@tanstack/react-query'
import { apiClient } from '../api-client'

export interface ChatMessage {
  role: 'user' | 'assistant'
  content: string
}

export interface PatientContext {
  name: string
  age: number
  gender: string
  weight?: number
  height?: number
  bmi?: number
  objective?: string
  riskLevel?: string
  riskFactors?: string[]
}

export interface ChatRequest {
  messages: ChatMessage[]
  patientContext?: PatientContext
}

export interface ChatResponse {
  message: string
  sources?: string[]
}

export function useTrainingAIChat() {
  return useMutation({
    mutationFn: (data: ChatRequest) =>
      apiClient.post<ChatResponse>('/api/v1/training/ai/chat', data),
  })
}

export function useTrainingAIRecommendation() {
  return useMutation({
    mutationFn: (data: { patientContext?: PatientContext; objective: string }) =>
      apiClient.post<{ recommendation: string }>('/api/v1/training/ai/recommendations', data),
  })
}
