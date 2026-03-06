"use client";

import { useParams } from "next/navigation";
import Link from "next/link";
import { ArrowLeft, Dumbbell, Copy } from "lucide-react";
import { toast } from "sonner";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { Skeleton } from "@/components/ui/skeleton";
import { useRequireAuth } from "@/lib/use-auth";
import { useRequireSelectedPatient } from "@/lib/use-require-selected-patient";
import { SelectedPatientHeader } from "@/components/patients/SelectedPatientHeader";
import { PageHeader } from "@/components/layout/page-header";
import { useWorkoutPlan } from "@/lib/api/workout-plan-api";

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

  const { data: plan, isLoading } = useWorkoutPlan(id);

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
        subtitle={plan.displayTitle}
        icon={Dumbbell}
        actions={
          <div className="flex gap-2">
            <Button variant="outline" size="sm" onClick={copyLink}>
              <Copy className="h-4 w-4 mr-2" />
              Copiar Link
            </Button>
            <Link href="/training/workout-plans">
              <Button variant="outline" size="sm">
                <ArrowLeft className="h-4 w-4 mr-2" />
                Voltar
              </Button>
            </Link>
          </div>
        }
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
