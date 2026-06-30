'use client';

/**
 * "Enviar por WhatsApp" de documentos clínicos ao paciente (pedido de exames, documento emitido,
 * receita, doc de prontuário) sem o round-trip de baixar+reanexar. Backend:
 * /api/v1/patients/:id/clinical-documents/* (ver clinical_document_handler.go).
 */

import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { apiClient } from '../api-client';

export type ClinicalDocType =
  | 'lab_request'
  | 'issued_document'
  | 'prescription'
  | 'patient_document';

export type SendMode = 'file' | 'link';

export interface WhatsAppWindowState {
  hasPhone: boolean;
  windowOpen: boolean;
  lastInboundAt?: string;
}

export interface ClinicalDocItem {
  docType: ClinicalDocType;
  docId: string;
  title: string;
  filename: string;
  createdAt: string;
  signed: boolean;
}

/** Estado da janela de 24h do paciente — os cartões decidem quais botões mostrar. */
export function useWhatsAppWindow(patientId?: string) {
  return useQuery({
    queryKey: ['wa-window', patientId],
    enabled: !!patientId,
    staleTime: 60_000,
    queryFn: () =>
      apiClient.get<WhatsAppWindowState>(
        `/api/v1/patients/${patientId}/clinical-documents/whatsapp-window`,
      ),
  });
}

/** Documentos enviáveis do paciente (com PDF pronto) — picker "Anexar arquivo do EMR". */
export function useClinicalDocuments(patientId?: string, enabled = true) {
  return useQuery({
    queryKey: ['clinical-documents', patientId],
    enabled: !!patientId && enabled,
    queryFn: () =>
      apiClient
        .get<{ data: ClinicalDocItem[] }>(`/api/v1/patients/${patientId}/clinical-documents`)
        .then((r) => r.data ?? []),
  });
}

export function useSendClinicalDocWhatsApp(patientId: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (vars: { docType: ClinicalDocType; docId: string; mode: SendMode }) =>
      apiClient.post(
        `/api/v1/patients/${patientId}/clinical-documents/send-whatsapp`,
        vars,
      ),
    onSuccess: () => {
      // A conversa do paciente ganhou uma mensagem — invalida pra refletir na timeline.
      qc.invalidateQueries({ queryKey: ['conversation'] });
      qc.invalidateQueries({ queryKey: ['conversations'] });
    },
  });
}
