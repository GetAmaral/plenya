import { apiClient } from '../api-client'

export interface Patient {
  id: string
  name: string
  cpf: string
  email?: string
  phone?: string
  birthDate: string
  gender: 'male' | 'female' | 'other'
  address?: string
  city?: string
  state?: string
  zipCode?: string
  municipality?: string
  createdAt: string
  updatedAt: string
}

// Get all patients
export async function getPatients(): Promise<Patient[]> {
  const response = await apiClient.get('/patients')
  return response.data
}

// Get patient by ID
export async function getPatientById(id: string): Promise<Patient> {
  const response = await apiClient.get(`/patients/${id}`)
  return response.data
}

// Create patient
export async function createPatient(data: Partial<Patient>): Promise<Patient> {
  const response = await apiClient.post('/patients', data)
  return response.data
}

// Update patient
export async function updatePatient(id: string, data: Partial<Patient>): Promise<Patient> {
  const response = await apiClient.put(`/patients/${id}`, data)
  return response.data
}

// Delete patient
export async function deletePatient(id: string): Promise<void> {
  await apiClient.delete(`/patients/${id}`)
}
