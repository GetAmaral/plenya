"use client";

import { useParams } from "next/navigation";

import { ArrowLeft, Copy, FileText, Pencil } from "lucide-react";
import { toast } from "sonner";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";

import { Skeleton } from "@/components/ui/skeleton";
import { useRequireAuth } from "@/lib/use-auth";
import { useRequireSelectedPatient } from "@/lib/use-require-selected-patient";
import { SelectedPatientHeader } from "@/components/patients/SelectedPatientHeader";
import { PageHeader } from "@/components/layout/page-header";
import { useWorkoutPlan, useGenerateHtml } from "@/lib/api/workout-plan-api";

const API_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:3001";

const phaseLabels: Record<string, string> = {
  warmup: "Aquecimento",
  main: "Principal",
  cooldown: "Finalização",
};

const cadenceLabels: Record<string, string> = {
  normal: "Normal",
  slow: "Lento",
  paused: "Pausado",
  explosive: "Explosivo",
  free: "Livre",
};

export default function WorkoutPlanDetailPage() {
  useRequireAuth();
  useRequireSelectedPatient();
  const params = useParams();
  const id = params.id as string;

  const { data: plan, isLoading, refetch } = useWorkoutPlan(id);
  const generateHtml = useGenerateHtml();

  const handleGenerateHtml = async () => {
    try {
      await generateHtml.mutateAsync(id);
      toast.success("HTML gerado com sucesso!");
      refetch();
    } catch (error: any) {
      toast.error(error?.message || "Erro ao gerar HTML");
    }
  };

  if (isLoading) {
    return <Skeleton className="h-96" />;
  }

  if (!plan) {
    return <p className="text-center text-muted-foreground py-12">Plano não encontrado.</p>;
  }

  const copyLink = () => {
    navigator.clipboard.writeText(`${window.location.origin}/workout?code=${plan.publicCode}`);
    toast.success("Link copiado!");
  };

  return (
    <div className="space-y-6">
      <SelectedPatientHeader />
      <PageHeader
        title={plan.name}
        description={plan.displayTitle}
        actions={[
          {
            label: "Editar",
            icon: <Pencil className="h-4 w-4" />,
            onClick: () => window.location.assign(`/training/workout-plans/${id}/edit`),
            variant: "outline",
          },
          {
            label: generateHtml.isPending ? "Gerando..." : "Gerar HTML",
            icon: <FileText className="h-4 w-4" />,
            onClick: handleGenerateHtml,
            variant: "outline",
            disabled: generateHtml.isPending,
            loading: generateHtml.isPending,
          },
          {
            label: "Copiar Link",
            icon: <Copy className="h-4 w-4" />,
            onClick: copyLink,
            variant: "outline",
          },
          {
            label: "Voltar",
            icon: <ArrowLeft className="h-4 w-4" />,
            onClick: () => window.location.assign("/training/workout-plans"),
            variant: "outline",
          },
        ]}
      />

      {plan.sessions?.map((session) => (
        <Card key={session.id}>
          <CardHeader>
            <CardTitle className="text-base">{session.name}</CardTitle>
          </CardHeader>
          <CardContent>
            {(["warmup", "main", "cooldown"] as const).map((phase) => {
              const exercises = session.exercises?.filter((e) => e.phase === phase);
              if (!exercises?.length) return null;
              return (
                <div key={phase} className="mb-4">
                  <h4 className="text-sm font-semibold text-muted-foreground mb-2">
                    {phaseLabels[phase]}
                  </h4>
                  <div className="space-y-2">
                    {exercises.map((ex) => (
                      <div key={ex.id} className="flex items-center gap-3 p-3 rounded-lg bg-muted/50">
                        {ex.exercise?.gifUrl && (
                          <img
                            src={`${API_URL}${ex.exercise.gifUrl}`}
                            alt={ex.exercise.namePt || ex.exercise.name}
                            className="w-16 h-16 rounded object-cover"
                            loading="lazy"
                          />
                        )}
                        <div className="flex-1 min-w-0">
                          <p className="font-medium text-sm truncate">
                            {ex.exercise?.namePt || ex.exercise?.name || "Exercício"}
                          </p>
                          <div className="flex flex-wrap gap-1 mt-1">
                            <Badge variant="outline" className="text-xs">{ex.sets}x{ex.reps}</Badge>
                            <Badge variant="secondary" className="text-xs">{cadenceLabels[ex.cadence]}</Badge>
                            <Badge variant="secondary" className="text-xs">Desc: {ex.restBetweenSetsSec}s</Badge>
                          </div>
                        </div>
                      </div>
                    ))}
                  </div>
                </div>
              );
            })}
          </CardContent>
        </Card>
      ))}

      {plan.htmlContent && (
        <Card>
          <CardHeader>
            <CardTitle className="text-base">Visualização HTML</CardTitle>
          </CardHeader>
          <CardContent>
            <div
              className="prose prose-sm max-w-none"
              dangerouslySetInnerHTML={{ __html: plan.htmlContent }}
            />
          </CardContent>
        </Card>
      )}
    </div>
  );
}
