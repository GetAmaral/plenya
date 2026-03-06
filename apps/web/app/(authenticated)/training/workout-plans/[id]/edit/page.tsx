"use client";

import { useEffect, useState } from "react";
import { useParams, useRouter } from "next/navigation";
import { toast } from "sonner";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import { Switch } from "@/components/ui/switch";
import { Skeleton } from "@/components/ui/skeleton";
import { useRequireAuth } from "@/lib/use-auth";
import { useRequireSelectedPatient } from "@/lib/use-require-selected-patient";
import { SelectedPatientHeader } from "@/components/patients/SelectedPatientHeader";
import { PageHeader } from "@/components/layout/page-header";
import {
  useWorkoutPlan,
  useUpdateWorkoutPlan,
  type WorkoutObjective,
  type WorkoutIntensity,
} from "@/lib/api/workout-plan-api";

export default function EditWorkoutPlanPage() {
  useRequireAuth();
  useRequireSelectedPatient();
  const params = useParams();
  const router = useRouter();
  const id = params.id as string;

  const { data: plan, isLoading } = useWorkoutPlan(id);
  const updateMutation = useUpdateWorkoutPlan();

  const [name, setName] = useState("");
  const [objective, setObjective] = useState<WorkoutObjective>("hypertrophy");
  const [intensity, setIntensity] = useState<WorkoutIntensity>("moderate");
  const [durationMinutes, setDurationMinutes] = useState(60);
  const [weeklyFrequency, setWeeklyFrequency] = useState(3);
  const [isActive, setIsActive] = useState(true);

  useEffect(() => {
    if (plan) {
      setName(plan.name);
      setObjective(plan.objective);
      setIntensity(plan.intensity);
      setDurationMinutes(plan.durationMinutes);
      setWeeklyFrequency(plan.weeklyFrequency);
      setIsActive(plan.isActive);
    }
  }, [plan]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!name) {
      toast.error("Nome do plano e obrigatorio");
      return;
    }

    try {
      await updateMutation.mutateAsync({
        id,
        data: { name, objective, intensity, durationMinutes, weeklyFrequency, isActive },
      });
      toast.success("Plano atualizado!");
      router.push(`/training/workout-plans/${id}`);
    } catch (error: any) {
      toast.error(error?.message || "Erro ao atualizar plano");
    }
  };

  if (isLoading) return <Skeleton className="h-96" />;
  if (!plan) return <p className="text-center text-muted-foreground py-12">Plano nao encontrado.</p>;

  return (
    <div className="space-y-6">
      <SelectedPatientHeader />
      <PageHeader
        title="Editar Plano de Treino"
        description={plan.displayTitle}
      />

      <form onSubmit={handleSubmit} className="space-y-6 max-w-xl">
        <Card>
          <CardHeader>
            <CardTitle className="text-base">Dados do Plano</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <div className="space-y-2">
              <Label>Nome do Plano</Label>
              <Input value={name} onChange={(e) => setName(e.target.value)} required />
            </div>

            <div className="grid grid-cols-2 gap-4">
              <div className="space-y-2">
                <Label>Objetivo</Label>
                <Select value={objective} onValueChange={(v) => setObjective(v as WorkoutObjective)}>
                  <SelectTrigger><SelectValue /></SelectTrigger>
                  <SelectContent>
                    <SelectItem value="hypertrophy">Hipertrofia</SelectItem>
                    <SelectItem value="strength">Forca</SelectItem>
                    <SelectItem value="endurance">Resistencia</SelectItem>
                    <SelectItem value="weight_loss">Emagrecimento</SelectItem>
                    <SelectItem value="conditioning">Condicionamento</SelectItem>
                    <SelectItem value="rehabilitation">Reabilitacao</SelectItem>
                  </SelectContent>
                </Select>
              </div>

              <div className="space-y-2">
                <Label>Intensidade</Label>
                <Select value={intensity} onValueChange={(v) => setIntensity(v as WorkoutIntensity)}>
                  <SelectTrigger><SelectValue /></SelectTrigger>
                  <SelectContent>
                    <SelectItem value="very_light">Muito Leve</SelectItem>
                    <SelectItem value="light">Leve</SelectItem>
                    <SelectItem value="moderate">Moderado</SelectItem>
                    <SelectItem value="high">Alto</SelectItem>
                    <SelectItem value="very_high">Muito Alto</SelectItem>
                  </SelectContent>
                </Select>
              </div>
            </div>

            <div className="grid grid-cols-2 gap-4">
              <div className="space-y-2">
                <Label>Duracao (min)</Label>
                <Input type="number" min={10} max={300} value={durationMinutes} onChange={(e) => setDurationMinutes(Number(e.target.value))} />
              </div>
              <div className="space-y-2">
                <Label>Frequencia Semanal</Label>
                <Input type="number" min={1} max={7} value={weeklyFrequency} onChange={(e) => setWeeklyFrequency(Number(e.target.value))} />
              </div>
            </div>

            <div className="flex items-center gap-2">
              <Switch checked={isActive} onCheckedChange={setIsActive} />
              <Label>Plano ativo</Label>
            </div>
          </CardContent>
        </Card>

        <div className="flex gap-2">
          <Button type="submit" disabled={updateMutation.isPending} className="flex-1">
            {updateMutation.isPending ? "Salvando..." : "Salvar Alteracoes"}
          </Button>
          <Button type="button" variant="outline" onClick={() => router.push(`/training/workout-plans/${id}`)}>
            Cancelar
          </Button>
        </div>
      </form>
    </div>
  );
}
