"use client";

import { useRequireAuth } from "@/lib/use-auth";
import { useRequireSelectedPatient } from "@/lib/use-require-selected-patient";
import { SelectedPatientHeader } from "@/components/patients/SelectedPatientHeader";
import { PageHeader } from "@/components/layout/page-header";
import { Card, CardContent } from "@/components/ui/card";
import dynamic from "next/dynamic";

// Dynamic import para evitar SSR issues
const LabResultBatchForm = dynamic(
  () => import("@/components/lab-results/LabResultBatchForm").then(mod => ({ default: mod.LabResultBatchForm })),
  {
    ssr: false,
    loading: () => <Card><CardContent className="p-8 text-center">Carregando formulário...</CardContent></Card>,
  }
);

export default function NewLabResultBatchPage() {
  useRequireAuth();
  const { selectedPatient, isLoading } = useRequireSelectedPatient();

  if (isLoading) {
    return null;
  }

  if (!selectedPatient) {
    return null;
  }

  return (
    <div className="container mx-auto py-8 space-y-8">
      <SelectedPatientHeader />

      <PageHeader
        breadcrumbs={[
          { label: "Exames", href: "/lab-results" },
          { label: "Novo Lote" },
        ]}
        title="Novo Lote de Resultados"
        description="Adicione múltiplos resultados de exames de uma mesma coleta"
      />

      <LabResultBatchForm />
    </div>
  );
}
