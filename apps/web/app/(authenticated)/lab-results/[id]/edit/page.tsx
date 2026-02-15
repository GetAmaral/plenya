"use client";

import { useParams } from "next/navigation";
import { useQuery } from "@tanstack/react-query";
import { useEffect, useState } from "react";
import dynamic from "next/dynamic";
import { Card, CardContent } from "@/components/ui/card";
import { Skeleton } from "@/components/ui/skeleton";
import { useRequireAuth } from "@/lib/use-auth";
import { useRequireSelectedPatient } from "@/lib/use-require-selected-patient";
import { SelectedPatientHeader } from "@/components/patients/SelectedPatientHeader";
import { PageHeader } from "@/components/layout/page-header";
import { labResultBatchApi } from "@/lib/api/lab-result-batch-api";
import { apiToFormValues } from "@/lib/validations/lab-result-batch";
import { ArrowLeft } from "lucide-react";

const LabResultBatchForm = dynamic(
  () =>
    import("@/components/lab-results/LabResultBatchForm").then((mod) => ({
      default: mod.LabResultBatchForm,
    })),
  {
    ssr: false,
    loading: () => (
      <Card>
        <CardContent className="py-12">
          <Skeleton className="h-96 w-full" />
        </CardContent>
      </Card>
    ),
  }
);

export default function EditLabResultBatchPage() {
  useRequireAuth();
  useRequireSelectedPatient();
  const params = useParams();
  const batchId = params.id as string;
  const [focusLabResultId, setFocusLabResultId] = useState<string | null>(null);

  const { data: batch, isLoading } = useQuery({
    queryKey: ["lab-result-batch", batchId],
    queryFn: () => labResultBatchApi.getById(batchId),
  });

  // Detect lab result ID from URL hash (for auto-scroll and focus)
  useEffect(() => {
    const hash = window.location.hash.slice(1); // Remove '#' prefix
    if (hash) {
      setFocusLabResultId(hash);
    }
  }, []);

  if (isLoading) {
    return (
      <div className="container mx-auto py-8 space-y-8">
        <SelectedPatientHeader />
        <PageHeader
          breadcrumbs={[
            { label: "Exames", href: "/lab-results" },
            { label: "Editar Lote" },
          ]}
          title="Editar Lote de Resultados"
          description="Atualize as informações do lote de exames"
        />
        <Card>
          <CardContent className="py-12">
            <Skeleton className="h-96 w-full" />
          </CardContent>
        </Card>
      </div>
    );
  }

  if (!batch) {
    return (
      <div className="container mx-auto py-8 space-y-8">
        <SelectedPatientHeader />
        <Card>
          <CardContent className="py-12 text-center">
            <p className="text-muted-foreground">Lote não encontrado</p>
          </CardContent>
        </Card>
      </div>
    );
  }

  const initialValues = apiToFormValues(batch);

  return (
    <div className="container mx-auto py-8 space-y-8">
      <SelectedPatientHeader />

      <PageHeader
        breadcrumbs={[
          { label: "Exames", href: "/lab-results" },
          { label: batch.laboratoryName, href: `/lab-results/${batchId}` },
          { label: "Editar" },
        ]}
        title="Editar Lote de Resultados"
        description={`Editando lote: ${batch.laboratoryName}`}
      />

      <LabResultBatchForm
        batchId={batchId}
        initialValues={initialValues}
        focusLabResultId={focusLabResultId}
      />
    </div>
  );
}
