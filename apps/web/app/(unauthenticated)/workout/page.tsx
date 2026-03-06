"use client";

import { useSearchParams } from "next/navigation";
import { Suspense } from "react";
import { Dumbbell } from "lucide-react";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import { usePublicWorkoutPlan } from "@/lib/api/workout-plan-api";

const API_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:3001";

const phaseLabels: Record<string, string> = {
  warmup: "Aquecimento",
  main: "Principal",
  cooldown: "Finalização",
};

function WorkoutContent() {
  const searchParams = useSearchParams();
  const code = searchParams.get("code") || "";

  const { data: plan, isLoading, error } = usePublicWorkoutPlan(code);

  if (!code) {
    return (
      <div className="flex flex-col items-center justify-center min-h-screen p-4">
        <Dumbbell className="h-12 w-12 text-muted-foreground mb-4" />
        <p className="text-muted-foreground">Código do treino não fornecido.</p>
      </div>
    );
  }

  if (isLoading) {
    return (
      <div className="p-4 space-y-4">
        <Skeleton className="h-12 w-full" />
        <Skeleton className="h-96 w-full" />
      </div>
    );
  }

  if (error || !plan) {
    return (
      <div className="flex flex-col items-center justify-center min-h-screen p-4">
        <Dumbbell className="h-12 w-12 text-muted-foreground mb-4" />
        <p className="text-lg font-semibold">Treino não encontrado</p>
        <p className="text-sm text-muted-foreground">Verifique o código e tente novamente.</p>
      </div>
    );
  }

  // If HTML content is available, render it directly
  if (plan.htmlContent) {
    return (
      <div
        className="min-h-screen"
        dangerouslySetInnerHTML={{ __html: plan.htmlContent }}
      />
    );
  }

  // Fallback: render structured data
  return (
    <div className="max-w-2xl mx-auto p-4 space-y-6">
      <div className="text-center py-6">
        <h1 className="text-2xl font-bold">{plan.name}</h1>
        <p className="text-muted-foreground">{plan.displayTitle}</p>
      </div>

      {plan.sessions?.map((session) => (
        <div key={session.id} className="space-y-4">
          <h2 className="text-lg font-semibold border-b pb-2">{session.name}</h2>

          {(["warmup", "main", "cooldown"] as const).map((phase) => {
            const exercises = session.exercises?.filter((e) => e.phase === phase);
            if (!exercises?.length) return null;
            return (
              <div key={phase}>
                <h3 className="text-sm font-semibold text-muted-foreground mb-2 uppercase">
                  {phaseLabels[phase]}
                </h3>
                <div className="space-y-3">
                  {exercises.map((ex) => (
                    <div key={ex.id} className="flex items-start gap-3 p-3 rounded-lg border">
                      {ex.exercise?.gifUrl && (
                        <img
                          src={`${API_URL}${ex.exercise.gifUrl}`}
                          alt={ex.exercise.namePt || ex.exercise.name}
                          className="w-20 h-20 rounded object-cover flex-shrink-0"
                          loading="lazy"
                        />
                      )}
                      <div className="flex-1">
                        <p className="font-medium">
                          {ex.exercise?.namePt || ex.exercise?.name || "Exercício"}
                        </p>
                        <div className="flex flex-wrap gap-1 mt-1">
                          <Badge variant="outline">{ex.sets} x {ex.reps}</Badge>
                          <Badge variant="secondary">Desc: {ex.restBetweenSetsSec}s</Badge>
                        </div>
                        {ex.notes && (
                          <p className="text-xs text-muted-foreground mt-1">{ex.notes}</p>
                        )}
                      </div>
                    </div>
                  ))}
                </div>
              </div>
            );
          })}
        </div>
      ))}

      <div className="text-center py-6 text-xs text-muted-foreground border-t">
        Gerado por Plenya
      </div>
    </div>
  );
}

export default function WorkoutPublicPage() {
  return (
    <Suspense fallback={<Skeleton className="h-screen" />}>
      <WorkoutContent />
    </Suspense>
  );
}
