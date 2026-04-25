import { useEffect, useRef } from 'react';
import { options } from '@plenya/api-client';

/**
 * Garante que o paciente passado seja o `User.SelectedPatientID` no backend.
 * Endpoints de prontuário (anamnese, labs, prescrições, treino, avaliação física)
 * filtram pelo selected patient — sem essa chamada, todas as listas vêm vazias.
 *
 * Idempotente: dispara só uma vez por mount + patientId. Não bloqueia render.
 */
export function useEnsureSelectedPatient(patientId: string | null | undefined): void {
  const lastSentRef = useRef<string | null>(null);

  useEffect(() => {
    if (!patientId) return;
    if (lastSentRef.current === patientId) return;
    lastSentRef.current = patientId;
    options.patientMutations.setSelected(patientId).catch(() => {
      lastSentRef.current = null; // permite retry no próximo mount
    });
  }, [patientId]);
}
