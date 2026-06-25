'use client';

/**
 * /issued-documents — lista de documentos clínicos do paciente selecionado
 * (atestado, declaração, laudo, orientações). Igual às telas de Prescrições e
 * Pedidos de Exames: isolada por selectedPatient. Mostra o CONTEÚDO inline
 * (showContent), além de emitir/editar/assinar/baixar via IssuedDocumentsCard.
 */

import { usePatientGuard } from '@/lib/use-patient-guard';
import { useRequireSelectedPatient } from '@/lib/use-require-selected-patient';
import { SelectedPatientHeader } from '@/components/patients/SelectedPatientHeader';
import { PageHeader } from '@/components/layout/page-header';
import { IssuedDocumentsCard } from '@/components/clinical/issued-documents-card';

export default function IssuedDocumentsPage() {
  usePatientGuard(); // staff only
  const { selectedPatient, isLoading } = useRequireSelectedPatient();

  if (isLoading) {
    return <div className="container mx-auto py-8">Carregando...</div>;
  }

  return (
    <div className="container mx-auto py-8 space-y-8">
      <SelectedPatientHeader />

      <PageHeader
        breadcrumbs={[{ label: 'Documentos' }]}
        title="Documentos"
        description="Atestados, declarações, laudos e orientações do paciente"
      />

      {selectedPatient && (
        <IssuedDocumentsCard patientId={selectedPatient.id} showContent />
      )}
    </div>
  );
}
