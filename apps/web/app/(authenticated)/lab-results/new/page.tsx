"use client";

import { useRequireAuth } from "@/lib/use-auth";
import { useRequireSelectedPatient } from "@/lib/use-require-selected-patient";
import { SelectedPatientHeader } from "@/components/patients/SelectedPatientHeader";
import { PageHeader } from "@/components/layout/page-header";
import { Card, CardContent, CardHeader, CardTitle, CardDescription } from "@/components/ui/card";
import dynamic from "next/dynamic";

// Dynamic import para evitar SSR issues
const LabResultBatchForm = dynamic(
  () => import("@/components/lab-results/LabResultBatchForm").then(mod => ({ default: mod.LabResultBatchForm })),
  {
    ssr: false,
    loading: () => <Card><CardContent className="p-8 text-center">Carregando formulário...</CardContent></Card>,
  }
);

const MultiPDFImport = dynamic(
  () => import("@/components/lab-results/MultiPDFImport").then(mod => ({ default: mod.MultiPDFImport })),
  { ssr: false }
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

      <Card>
        <CardHeader>
          <CardTitle className="text-base">Importar laudos em PDF (vários de uma vez)</CardTitle>
          <CardDescription>
            Cada PDF vira um lote separado, com laboratório e data de coleta lidos do próprio laudo.
            A interpretação e a classificação rodam em segundo plano.
          </CardDescription>
        </CardHeader>
        <CardContent>
          <MultiPDFImport />
        </CardContent>
      </Card>

      <div className="relative">
        <div className="absolute inset-0 flex items-center">
          <span className="w-full border-t" />
        </div>
        <div className="relative flex justify-center text-xs uppercase">
          <span className="bg-background px-2 text-muted-foreground">ou lançar um lote manualmente</span>
        </div>
      </div>

      <LabResultBatchForm />
    </div>
  );
}
