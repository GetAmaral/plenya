'use client'

import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'
import { apiClient } from '../api-client'
import {
  updateAnamnesis,
  type Anamnesis,
  type CreateAnamnesisRequest,
  type UpdateAnamnesisRequest,
} from './anamnesis'

// Anamnese vinculada a uma consulta (Appointment.AnamnesisID). 404 → null (estado normal).
export async function getAppointmentAnamnesis(appointmentId: string): Promise<Anamnesis | null> {
  try {
    return await apiClient.get<Anamnesis>(`/api/v1/appointments/${appointmentId}/anamnesis`)
  } catch (err) {
    if ((err as { status?: number }).status === 404) return null
    throw err
  }
}

export const appointmentAnamnesisKeys = {
  byAppointment: (id: string) => ['appointment-anamnesis', id] as const,
}

export function useAppointmentAnamnesis(appointmentId: string | undefined) {
  return useQuery({
    queryKey: appointmentAnamnesisKeys.byAppointment(appointmentId ?? ''),
    enabled: !!appointmentId,
    queryFn: () => getAppointmentAnamnesis(appointmentId!),
  })
}

export function useCreateAppointmentAnamnesis(appointmentId: string) {
  const qc = useQueryClient()
  return useMutation({
    mutationFn: (data: CreateAnamnesisRequest) =>
      apiClient.post<Anamnesis>(`/api/v1/appointments/${appointmentId}/anamnesis`, data),
    onSuccess: () =>
      qc.invalidateQueries({ queryKey: appointmentAnamnesisKeys.byAppointment(appointmentId) }),
  })
}

export function useUpdateAppointmentAnamnesis(appointmentId: string) {
  const qc = useQueryClient()
  return useMutation({
    mutationFn: ({ id, data }: { id: string; data: UpdateAnamnesisRequest }) =>
      updateAnamnesis(id, data),
    onSuccess: () =>
      qc.invalidateQueries({ queryKey: appointmentAnamnesisKeys.byAppointment(appointmentId) }),
  })
}
