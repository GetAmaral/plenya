"use client";

import { useParams } from "next/navigation";
import { ArrowLeft, Calendar } from "lucide-react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import { useRequireAuth } from "@/lib/use-auth";
import { useRequireSelectedPatient } from "@/lib/use-require-selected-patient";
import { SelectedPatientHeader } from "@/components/patients/SelectedPatientHeader";
import { PageHeader } from "@/components/layout/page-header";
import { usePeriodization } from "@/lib/api/periodization-api";

const phaseLabels: Record<string, string> = {
  accumulation: "Acumulacao",
  transformation: "Transformacao",
  realization: "Realizacao",
  hypertrophy: "Hipertrofia",
  strength: "Forca",
  endurance: "Resistencia",
  power: "Potencia",
};

const phaseColors: Record<string, string> = {
  accumulation: "bg-blue-500",
  transformation: "bg-yellow-500",
  realization: "bg-red-500",
  hypertrophy: "bg-purple-500",
  strength: "bg-orange-500",
  endurance: "bg-green-500",
  power: "bg-pink-500",
};

const frameworkLabels: Record<string, string> = {
  bompa: "Bompa",
  linear: "Linear",
  undulating: "Ondulatoria",
  block: "Bloco",
};

export default function PeriodizationDetailPage() {
  useRequireAuth();
  useRequireSelectedPatient();
  const params = useParams();
  const id = params.id as string;

  const { data: periodization, isLoading } = usePeriodization(id);

  if (isLoading) return <Skeleton className="h-96" />;
  if (!periodization) return <p className="text-center text-muted-foreground py-12">Periodizacao nao encontrada.</p>;

  return (
    <div className="space-y-6">
      <SelectedPatientHeader />
      <PageHeader
        title={periodization.objective}
        description={`${frameworkLabels[periodization.framework] || periodization.framework} - ${periodization.totalWeeks} semanas`}
        actions={[
          {
            label: "Voltar",
            icon: <ArrowLeft className="h-4 w-4" />,
            onClick: () => window.location.assign("/training/periodization"),
            variant: "outline",
          },
        ]}
      />

      {/* Timeline visual */}
      <Card>
        <CardHeader>
          <CardTitle className="text-base flex items-center gap-2">
            <Calendar className="h-4 w-4" />
            Macrociclo
          </CardTitle>
        </CardHeader>
        <CardContent>
          <div className="flex gap-1 h-8 rounded-lg overflow-hidden">
            {periodization.mesocycles?.map((m, i) => {
              const widthPercent = (m.durationWeeks / periodization.totalWeeks) * 100;
              return (
                <div
                  key={m.id || i}
                  className={`${phaseColors[m.phase] || "bg-gray-400"} flex items-center justify-center text-white text-[10px] font-medium`}
                  style={{ width: `${widthPercent}%` }}
                  title={`${phaseLabels[m.phase] || m.phase} - ${m.durationWeeks} sem`}
                >
                  {m.durationWeeks >= 3 ? `${m.durationWeeks}s` : ""}
                </div>
              );
            })}
          </div>
          <div className="flex gap-3 mt-3 flex-wrap">
            {periodization.mesocycles?.map((m, i) => (
              <div key={m.id || i} className="flex items-center gap-1 text-xs">
                <div className={`w-3 h-3 rounded ${phaseColors[m.phase] || "bg-gray-400"}`} />
                {phaseLabels[m.phase] || m.phase}
              </div>
            ))}
          </div>
        </CardContent>
      </Card>

      {/* Mesocycle cards */}
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
        {periodization.mesocycles?.map((m, i) => (
          <Card key={m.id || i}>
            <CardHeader className="pb-2">
              <div className="flex items-center justify-between">
                <CardTitle className="text-sm">
                  Mesociclo {i + 1}
                </CardTitle>
                <Badge className={`${phaseColors[m.phase] || "bg-gray-400"} text-white text-[10px]`}>
                  {phaseLabels[m.phase] || m.phase}
                </Badge>
              </div>
            </CardHeader>
            <CardContent className="space-y-2">
              <div className="grid grid-cols-2 gap-2 text-sm">
                <div>
                  <p className="text-muted-foreground text-xs">Duracao</p>
                  <p className="font-medium">{m.durationWeeks} semanas</p>
                </div>
                <div>
                  <p className="text-muted-foreground text-xs">Foco</p>
                  <p className="font-medium text-xs">{m.physiologicalFocus}</p>
                </div>
              </div>
              <div className="space-y-1">
                <div className="flex items-center justify-between text-xs">
                  <span className="text-muted-foreground">Volume</span>
                  <span>{m.volumePercent}%</span>
                </div>
                <div className="w-full bg-muted rounded-full h-2">
                  <div className="bg-blue-500 h-2 rounded-full" style={{ width: `${m.volumePercent}%` }} />
                </div>
              </div>
              <div className="space-y-1">
                <div className="flex items-center justify-between text-xs">
                  <span className="text-muted-foreground">Intensidade</span>
                  <span>{m.intensityPercent}%</span>
                </div>
                <div className="w-full bg-muted rounded-full h-2">
                  <div className="bg-red-500 h-2 rounded-full" style={{ width: `${m.intensityPercent}%` }} />
                </div>
              </div>
              {m.startDate && (
                <p className="text-[10px] text-muted-foreground">
                  {new Date(m.startDate).toLocaleDateString("pt-BR")} - {new Date(m.endDate).toLocaleDateString("pt-BR")}
                </p>
              )}
            </CardContent>
          </Card>
        ))}
      </div>

      {/* Scientific justification */}
      {periodization.scientificJustification && (
        <Card>
          <CardHeader>
            <CardTitle className="text-base">Justificativa Cientifica</CardTitle>
          </CardHeader>
          <CardContent>
            <p className="text-sm text-muted-foreground whitespace-pre-wrap">
              {periodization.scientificJustification}
            </p>
          </CardContent>
        </Card>
      )}
    </div>
  );
}
