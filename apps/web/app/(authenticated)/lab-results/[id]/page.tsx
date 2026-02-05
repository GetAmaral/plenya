"use client";

import { useQuery } from "@tanstack/react-query";
import { useParams, useRouter } from "next/navigation";
import { format } from "date-fns";
import { ptBR } from "date-fns/locale";
import {
  ArrowLeft,
  Calendar,
  FileText,
  TestTube,
  Building2,
  Edit,
  Trash2,
  TrendingUp,
  Minus,
  Package,
} from "lucide-react";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import { Separator } from "@/components/ui/separator";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { useRequireAuth } from "@/lib/use-auth";
import { useRequireSelectedPatient } from "@/lib/use-require-selected-patient";
import { SelectedPatientHeader } from "@/components/patients/SelectedPatientHeader";
import { PageHeader } from "@/components/layout/page-header";
import { labResultBatchApi } from "@/lib/api/lab-result-batch-api";

const statusConfig = {
  pending: { label: "Pendente", variant: "outline" as const, icon: Minus },
  partial: { label: "Parcial", variant: "secondary" as const, icon: Package },
  completed: { label: "Concluído", variant: "default" as const, icon: TrendingUp },
};

const resultStatusConfig = {
  pending: { label: "Pendente", variant: "outline" as const },
  completed: { label: "Concluído", variant: "default" as const },
  cancelled: { label: "Cancelado", variant: "destructive" as const },
};

export default function LabResultBatchDetailPage() {
  useRequireAuth();
  useRequireSelectedPatient();
  const params = useParams();
  const router = useRouter();
  const batchId = params.id as string;

  const { data: batch, isLoading } = useQuery({
    queryKey: ["lab-result-batch", batchId],
    queryFn: () => labResultBatchApi.getById(batchId),
  });

  if (isLoading) {
    return (
      <div className="container mx-auto py-8 space-y-8">
        <SelectedPatientHeader />
        <div className="space-y-4">
          <Skeleton className="h-12 w-full" />
          <Skeleton className="h-64 w-full" />
          <Skeleton className="h-96 w-full" />
        </div>
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
            <Button
              variant="outline"
              className="mt-4"
              onClick={() => router.push("/lab-results")}
            >
              Voltar para Exames
            </Button>
          </CardContent>
        </Card>
      </div>
    );
  }

  const statusBadge = statusConfig[batch.status];
  const StatusIcon = statusBadge.icon;

  return (
    <div className="container mx-auto py-8 space-y-8">
      <SelectedPatientHeader />

      <PageHeader
        breadcrumbs={[
          { label: "Exames", href: "/lab-results" },
          { label: "Detalhes do Lote" },
        ]}
        title={`Lote: ${batch.laboratoryName}`}
        description={`${batch.resultCount} resultado(s) • Coletado em ${format(
          new Date(batch.collectionDate),
          "dd/MM/yyyy",
          { locale: ptBR }
        )}`}
        actions={[
          {
            label: "Voltar",
            icon: <ArrowLeft className="h-4 w-4" />,
            onClick: () => router.push("/lab-results"),
            variant: "outline" as const,
          },
          {
            label: "Editar Lote",
            icon: <Edit className="h-4 w-4" />,
            onClick: () => router.push(`/lab-results/${batchId}/edit`),
            variant: "default" as const,
          },
        ]}
      />

      {/* Informações do Lote */}
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
        <Card>
          <CardContent className="p-4">
            <div className="flex items-center gap-3">
              <Building2 className="h-8 w-8 text-muted-foreground" />
              <div>
                <p className="text-sm text-muted-foreground">Laboratório</p>
                <p className="font-semibold">{batch.laboratoryName}</p>
              </div>
            </div>
          </CardContent>
        </Card>

        <Card>
          <CardContent className="p-4">
            <div className="flex items-center gap-3">
              <Calendar className="h-8 w-8 text-muted-foreground" />
              <div>
                <p className="text-sm text-muted-foreground">Data da Coleta</p>
                <p className="font-semibold">
                  {format(new Date(batch.collectionDate), "dd/MM/yyyy", {
                    locale: ptBR,
                  })}
                </p>
              </div>
            </div>
          </CardContent>
        </Card>

        <Card>
          <CardContent className="p-4">
            <div className="flex items-center gap-3">
              <TestTube className="h-8 w-8 text-muted-foreground" />
              <div>
                <p className="text-sm text-muted-foreground">Resultados</p>
                <p className="font-semibold">{batch.resultCount}</p>
              </div>
            </div>
          </CardContent>
        </Card>

        <Card>
          <CardContent className="p-4">
            <div className="flex items-center gap-3">
              <StatusIcon className="h-8 w-8 text-muted-foreground" />
              <div>
                <p className="text-sm text-muted-foreground">Status</p>
                <Badge variant={statusBadge.variant} className="mt-1">
                  {statusBadge.label}
                </Badge>
              </div>
            </div>
          </CardContent>
        </Card>
      </div>

      {/* Observações */}
      {batch.observations && (
        <Card>
          <CardHeader>
            <CardTitle className="flex items-center gap-2 text-base">
              <FileText className="h-4 w-4" />
              Observações
            </CardTitle>
          </CardHeader>
          <CardContent>
            <p className="text-sm text-muted-foreground whitespace-pre-wrap">
              {batch.observations}
            </p>
          </CardContent>
        </Card>
      )}

      {/* Resultados */}
      <Card>
        <CardHeader>
          <CardTitle>Resultados do Lote</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="rounded-md border">
            <Table>
              <TableHeader>
                <TableRow>
                  <TableHead>Exame</TableHead>
                  <TableHead>Tipo</TableHead>
                  <TableHead>Resultado</TableHead>
                  <TableHead>Valores de Referência</TableHead>
                  <TableHead>Status</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {batch.labResults && batch.labResults.length > 0 ? (
                  batch.labResults.map((result) => {
                    const resultValue =
                      result.resultNumeric !== undefined && result.resultNumeric !== null
                        ? `${result.resultNumeric}${result.unit ? ` ${result.unit}` : ""}`
                        : result.resultText || "-";

                    const resultStatusBadge =
                      resultStatusConfig[result.status as keyof typeof resultStatusConfig];

                    return (
                      <TableRow key={result.id}>
                        <TableCell className="font-medium">
                          {result.testName}
                        </TableCell>
                        <TableCell className="text-sm text-muted-foreground">
                          {result.testType}
                        </TableCell>
                        <TableCell>
                          <div>
                            <div className="font-medium">{resultValue}</div>
                            {result.interpretation && (
                              <div className="text-xs text-muted-foreground mt-1">
                                {result.interpretation}
                              </div>
                            )}
                          </div>
                        </TableCell>
                        <TableCell className="text-sm text-muted-foreground">
                          {result.referenceRange || "-"}
                        </TableCell>
                        <TableCell>
                          <Badge variant={resultStatusBadge?.variant || "outline"}>
                            {resultStatusBadge?.label || result.status}
                          </Badge>
                        </TableCell>
                      </TableRow>
                    );
                  })
                ) : (
                  <TableRow>
                    <TableCell colSpan={5} className="text-center py-8 text-muted-foreground">
                      Nenhum resultado encontrado
                    </TableCell>
                  </TableRow>
                )}
              </TableBody>
            </Table>
          </div>
        </CardContent>
      </Card>

      {/* Metadados */}
      <Card>
        <CardHeader>
          <CardTitle className="text-base">Informações Adicionais</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="grid grid-cols-2 gap-4 text-sm">
            <div>
              <span className="text-muted-foreground">Criado em:</span>
              <p className="font-medium">
                {format(new Date(batch.createdAt), "dd/MM/yyyy 'às' HH:mm", {
                  locale: ptBR,
                })}
              </p>
            </div>
            <div>
              <span className="text-muted-foreground">Última atualização:</span>
              <p className="font-medium">
                {format(new Date(batch.updatedAt), "dd/MM/yyyy 'às' HH:mm", {
                  locale: ptBR,
                })}
              </p>
            </div>
            {batch.resultDate && (
              <div>
                <span className="text-muted-foreground">Data do resultado:</span>
                <p className="font-medium">
                  {format(new Date(batch.resultDate), "dd/MM/yyyy", {
                    locale: ptBR,
                  })}
                </p>
              </div>
            )}
          </div>
        </CardContent>
      </Card>
    </div>
  );
}
