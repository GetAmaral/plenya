"use client";

import { useParams } from "next/navigation";
import { format } from "date-fns";
import { ptBR } from "date-fns/locale";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import { useRequireAuth } from "@/lib/use-auth";
import { useRequireSelectedPatient } from "@/lib/use-require-selected-patient";
import { SelectedPatientHeader } from "@/components/patients/SelectedPatientHeader";
import { PageHeader } from "@/components/layout/page-header";
import { useFitnessTest } from "@/lib/api/fitness-test-api";

const levelColors: Record<string, string> = {
  excelente: "#22c55e",
  bom: "#3b82f6",
  medio: "#f59e0b",
  abaixo: "#f97316",
  fraco: "#ef4444",
};

const levelWidths: Record<string, string> = {
  excelente: "100%",
  bom: "75%",
  medio: "50%",
  abaixo: "25%",
  fraco: "10%",
};

const levelLabels: Record<string, string> = {
  excelente: "Excelente",
  bom: "Bom",
  medio: "Medio",
  abaixo: "Abaixo",
  fraco: "Fraco",
};

const classificationColors: Record<string, string> = {
  excelente: "bg-green-100 text-green-800 border-green-200",
  bom: "bg-blue-100 text-blue-800 border-blue-200",
  medio: "bg-yellow-100 text-yellow-800 border-yellow-200",
  abaixo: "bg-orange-100 text-orange-800 border-orange-200",
  fraco: "bg-red-100 text-red-800 border-red-200",
};

interface TestResult {
  name: string;
  value?: number | null;
  unit: string;
  level?: string | null;
}

export default function FitnessTestDetailPage() {
  useRequireAuth();
  useRequireSelectedPatient();
  const params = useParams();
  const id = params.id as string;

  const { data: ft, isLoading } = useFitnessTest(id);

  if (isLoading) {
    return (
      <div className="space-y-6">
        <Skeleton className="h-20" />
        <Skeleton className="h-40" />
        <Skeleton className="h-64" />
      </div>
    );
  }

  if (!ft) {
    return (
      <div className="text-center py-12">
        <p className="text-muted-foreground">Teste de fitness nao encontrado.</p>
      </div>
    );
  }

  const tests: TestResult[] = [
    { name: "Abdominal", value: ft.abdominal, unit: "rep/min", level: ft.abdominalLevel },
    { name: "Flexao", value: ft.flexao, unit: "rep/min", level: ft.flexaoLevel },
    { name: "Prancha", value: ft.prancha, unit: "segundos", level: ft.pranchaLevel },
    { name: "Burpee", value: ft.burpee, unit: "ciclos/3min", level: ft.burpeeLevel },
    { name: "FRT", value: ft.frt, unit: "rep/90s", level: ft.frtLevel },
  ];

  const performedTests = tests.filter((t) => t.value != null);

  return (
    <div className="space-y-6">
      <SelectedPatientHeader />
      <PageHeader
        title="Resultado do Teste de Fitness"
        description={format(new Date(ft.assessmentDate), "dd 'de' MMMM 'de' yyyy", { locale: ptBR })}
      />

      {/* Overall Score */}
      {ft.overallScore != null && ft.classification && (
        <Card className={`border-2 ${classificationColors[ft.classification] || ""}`}>
          <CardContent className="p-6 text-center">
            <p className="text-5xl font-bold">{ft.overallScore.toFixed(1)}</p>
            <p className="text-xl font-semibold mt-2">
              {levelLabels[ft.classification] || ft.classification}
            </p>
            <p className="text-sm text-muted-foreground mt-1">
              Pontuacao geral baseada em {performedTests.length} teste(s)
            </p>
          </CardContent>
        </Card>
      )}

      {/* Individual Test Results */}
      <Card>
        <CardHeader>
          <CardTitle className="text-base">Resultados Individuais</CardTitle>
        </CardHeader>
        <CardContent className="space-y-4">
          {performedTests.map((test) => {
            const level = test.level || "fraco";
            const color = levelColors[level] || levelColors.fraco;
            const width = levelWidths[level] || levelWidths.fraco;

            return (
              <div key={test.name} className="space-y-1.5">
                <div className="flex items-center justify-between">
                  <div className="flex items-center gap-2">
                    <span className="text-sm font-medium">{test.name}</span>
                    <span className="text-sm text-muted-foreground">
                      {test.value} {test.unit}
                    </span>
                  </div>
                  <Badge
                    variant="outline"
                    style={{
                      borderColor: color,
                      color: color,
                      backgroundColor: `${color}15`,
                    }}
                  >
                    {levelLabels[level] || level}
                  </Badge>
                </div>
                <div className="h-3 w-full rounded-full bg-muted overflow-hidden">
                  <div
                    className="h-full rounded-full transition-all duration-500"
                    style={{
                      width,
                      backgroundColor: color,
                    }}
                  />
                </div>
              </div>
            );
          })}

          {performedTests.length === 0 && (
            <p className="text-sm text-muted-foreground text-center py-4">
              Nenhum teste foi realizado.
            </p>
          )}
        </CardContent>
      </Card>

      {/* Notes */}
      {ft.notes && (
        <Card>
          <CardHeader>
            <CardTitle className="text-base">Observacoes</CardTitle>
          </CardHeader>
          <CardContent>
            <p className="text-sm whitespace-pre-wrap">{ft.notes}</p>
          </CardContent>
        </Card>
      )}
    </div>
  );
}
