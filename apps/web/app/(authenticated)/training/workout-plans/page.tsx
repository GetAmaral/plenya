"use client";

import { motion } from "framer-motion";
import Link from "next/link";
import { format } from "date-fns";
import { ptBR } from "date-fns/locale";
import { Plus, Dumbbell, Copy, ExternalLink } from "lucide-react";
import { toast } from "sonner";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import { useRequireAuth } from "@/lib/use-auth";
import { useRequireSelectedPatient } from "@/lib/use-require-selected-patient";
import { SelectedPatientHeader } from "@/components/patients/SelectedPatientHeader";
import { PageHeader } from "@/components/layout/page-header";
import { useWorkoutPlans } from "@/lib/api/workout-plan-api";

const objectiveLabels: Record<string, string> = {
  hypertrophy: "Hipertrofia",
  strength: "Força",
  endurance: "Resistência",
  weight_loss: "Emagrecimento",
  conditioning: "Condicionamento",
  rehabilitation: "Reabilitação",
};

const intensityLabels: Record<string, string> = {
  very_light: "Muito Leve",
  light: "Leve",
  moderate: "Moderado",
  high: "Alto",
  very_high: "Muito Alto",
};

export default function WorkoutPlansPage() {
  useRequireAuth();
  useRequireSelectedPatient();

  const { data: plans, isLoading } = useWorkoutPlans();

  const copyPublicLink = (code: string) => {
    const url = `${window.location.origin}/workout?code=${code}`;
    navigator.clipboard.writeText(url);
    toast.success("Link copiado!");
  };

  return (
    <div className="space-y-6">
      <SelectedPatientHeader />
      <PageHeader
        title="Planos de Treino"
        description="Gerencie planos de treino do paciente"
        actions={[
          {
            label: "Novo Plano",
            icon: <Plus className="h-4 w-4" />,
            onClick: () => window.location.assign("/training/workout-plans/new"),
          },
        ]}
      />

      {isLoading ? (
        <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
          {Array.from({ length: 4 }).map((_, i) => (
            <Skeleton key={i} className="h-48" />
          ))}
        </div>
      ) : !plans?.length ? (
        <Card>
          <CardContent className="flex flex-col items-center justify-center py-12 text-muted-foreground">
            <Dumbbell className="h-12 w-12 mb-4" />
            <p>Nenhum plano de treino encontrado.</p>
            <Link href="/training/workout-plans/new" className="mt-4">
              <Button variant="outline">Criar primeiro plano</Button>
            </Link>
          </CardContent>
        </Card>
      ) : (
        <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
          {plans.map((plan) => (
            <motion.div
              key={plan.id}
              initial={{ opacity: 0, y: 10 }}
              animate={{ opacity: 1, y: 0 }}
            >
              <Card className="hover:shadow-lg transition-shadow">
                <CardHeader className="pb-2">
                  <div className="flex items-center justify-between">
                    <CardTitle className="text-base">{plan.name}</CardTitle>
                    {!plan.isActive && <Badge variant="secondary">Inativo</Badge>}
                  </div>
                </CardHeader>
                <CardContent className="space-y-3">
                  <div className="flex flex-wrap gap-2">
                    <Badge>{objectiveLabels[plan.objective] || plan.objective}</Badge>
                    <Badge variant="outline">{intensityLabels[plan.intensity] || plan.intensity}</Badge>
                    <Badge variant="secondary">{plan.weeklyFrequency}x/sem</Badge>
                    <Badge variant="secondary">{plan.durationMinutes} min</Badge>
                  </div>
                  <p className="text-sm text-muted-foreground">
                    {plan.sessions?.length || 0} sessão(ões) •{" "}
                    {format(new Date(plan.createdAt), "dd/MM/yyyy", { locale: ptBR })}
                  </p>
                  <div className="flex gap-2">
                    <Link href={`/training/workout-plans/${plan.id}`} className="flex-1">
                      <Button variant="outline" size="sm" className="w-full">Ver Plano</Button>
                    </Link>
                    <Button
                      variant="ghost"
                      size="sm"
                      onClick={() => copyPublicLink(plan.publicCode)}
                      title="Copiar link público"
                    >
                      <Copy className="h-4 w-4" />
                    </Button>
                    <a
                      href={`/workout?code=${plan.publicCode}`}
                      target="_blank"
                      rel="noopener noreferrer"
                    >
                      <Button variant="ghost" size="sm" title="Abrir link público">
                        <ExternalLink className="h-4 w-4" />
                      </Button>
                    </a>
                  </div>
                </CardContent>
              </Card>
            </motion.div>
          ))}
        </div>
      )}
    </div>
  );
}
