'use client';

/**
 * Documentos clínicos emitidos/assináveis (P3 frente 2): atestado, declaração, laudo.
 * Cria rascunho → assina (ICP-Brasil quando há cert; degrada p/ assinatura manual) →
 * publicado no portal do paciente. Download autenticado via getBlob.
 */

import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { openServerPdf } from "@/lib/open-pdf";
import { apiClient } from '../api-client';

export type IssuedDocumentType = 'certificate' | 'declaration' | 'report' | 'orientation';
export type IssuedDocumentStatus = 'draft' | 'signed';

export interface IssuedDocument {
  id: string;
  patientId: string;
  appointmentId?: string;
  doctorId: string;
  doctorName?: string;
  type: IssuedDocumentType;
  title: string;
  body: string;
  bodyHtml?: string;
  purpose?: string;
  daysOff?: number;
  includesCid: boolean;
  cidCode?: string;
  cidConsent: boolean;
  status: IssuedDocumentStatus;
  hasDigitalSignature: boolean;
  signedAt?: string;
  signedPdfHash?: string;
  certificateSerial?: string;
  qrCodeData?: string;
  patientDocumentId?: string;
  issuedByUserId: string;
  issuedAt: string;
  createdAt: string;
}

export interface CreateIssuedDocumentPayload {
  appointmentId?: string;
  type: IssuedDocumentType;
  title: string;
  body?: string;
  bodyHtml?: string;
  purpose?: string;
  daysOff?: number;
  includesCid?: boolean;
  cidCode?: string;
  cidConsent?: boolean;
}

export const issuedDocKeys = {
  byPatient: (patientId: string) => ['issued-documents', patientId] as const,
};

export function useIssuedDocuments(patientId: string | undefined) {
  return useQuery({
    queryKey: issuedDocKeys.byPatient(patientId ?? ''),
    enabled: !!patientId,
    queryFn: () => apiClient.get<IssuedDocument[]>(`/api/v1/patients/${patientId}/issued-documents`),
  });
}

export function useCreateIssuedDocument(patientId: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (payload: CreateIssuedDocumentPayload) =>
      apiClient.post<IssuedDocument>(`/api/v1/patients/${patientId}/issued-documents`, payload),
    onSuccess: () => qc.invalidateQueries({ queryKey: issuedDocKeys.byPatient(patientId) }),
  });
}

export function useUpdateIssuedDocument(patientId: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ docId, payload }: { docId: string; payload: CreateIssuedDocumentPayload }) =>
      apiClient.put<IssuedDocument>(`/api/v1/issued-documents/${docId}`, payload),
    onSuccess: () => qc.invalidateQueries({ queryKey: issuedDocKeys.byPatient(patientId) }),
  });
}

export function useSignIssuedDocument(patientId: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (docId: string) => apiClient.post<IssuedDocument>(`/api/v1/issued-documents/${docId}/sign`),
    onSuccess: () => qc.invalidateQueries({ queryKey: issuedDocKeys.byPatient(patientId) }),
  });
}

export function useDeleteIssuedDocument(patientId: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (docId: string) => apiClient.delete<void>(`/api/v1/issued-documents/${docId}`),
    onSuccess: () => qc.invalidateQueries({ queryKey: issuedDocKeys.byPatient(patientId) }),
  });
}

/** Baixa o PDF assinado (autenticado) e entrega com o nome do servidor. Ver lib/open-pdf.ts. */
export async function openIssuedDocumentPDF(docId: string) {
  await openServerPdf(`/api/v1/issued-documents/${docId}/pdf`, `Documento_${docId.slice(0, 8)}.pdf`);
}
