"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { ArrowLeft, Dumbbell, Plus, Trash2 } from "lucide-react";
import { toast } from "sonner";
import { Card, CardContent, CardHeader, CardTitle, CardDescription } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import { useRequireAuth } from "@/lib/use-auth";
import { useRequireSelectedPatient } from "@/lib/use-require-selected-patient";
import { SelectedPatientHeader } from "@/components/patients/SelectedPatientHeader";
import { PageHeader } from "@/components/layout/page-header";
import {
  useCreateWorkoutPlan,
  type WorkoutObjective,
  type WorkoutIntensity,
  type CreateSessionDTO,
} from "@/lib/api/workout-plan-api";
import Link from "next/link";

export default function NewWorkoutPlanPage() {
  useRequireAuth();
  useRequireSelectedPatient();
  const router = useRouter();

  const [name, setName] = useState("");
  const [objective, setObjective] = useState<WorkoutObjective>("hypertrophy");
  const [intensity, setIntensity] = useState<WorkoutIntensity>("moderate");
  const [durationMinutes, setDurationMinutes] = useState(60);
  const [weeklyFrequency, setWeeklyFrequency] = useState(3);
  const [sessions, setSessions] = useState<{ name: string }[]>([{ name: "Treino A" }]);

  const createMutation = useCreateWorkoutPlan();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!name) {
      toast.error("Nome do plano é obrigatório");
      return;
    }

    try {
      const sessionDtos: CreateSessionDTO[] = sessions.map((s, i) => ({
        name: s.name,
        order: i,
        exercises: [],
      }));

      const result = await createMutation.mutateAsync({
        name,
        objective,
        intensity,
        durationMinutes,
        weeklyFrequency,
        sessions: sessionDtos,
      });
      toast.success("Plano criado com sucesso!");
      router.push(`/training/workout-plans/${result.id}`);
    } catch (error: any) {
      toast.error(error?.message || "Erro ao criar plano");
    }
  };

  return (
    <div className="space-y-6">
      <SelectedPatientHeader />
      <PageHeader
        title="Novo Plano de Treino"
        subtitle="Configure o plano e adicione exercícios depois"
        icon={Dumbbell}
        actions={
          <Link href="/training/workout-plans">
            <Button variant="outline">
              <ArrowLeft className="h-4 w-4 mr-2" />
              Voltar
            </Button>
          </Link>
        }
      />

      <Card className="max-w-xl">
        <CardHeader>
          <CardTitle>Dados do Plano</CardTitle>
        </CardHeader>
        <CardContent>
          <form onSubmit={handleSubmit} className="space-y-4">
            <div className="space-y-2">
              <Label>Nome do Plano</Label>
              <Input value={name} onChange={(e) => setName(e.target.value)} placeholder="Ex: Plano Hipertrofia Março" required />
            </div>

            <div className="grid grid-cols-2 gap-4">
              <div className="space-y-2">
                <Label>Objetivo</Label>
                <Select value={objective} onValueChange={(v) => setObjective(v as WorkoutObjective)}>
                  <SelectTrigger><SelectValue /></SelectTrigger>
                  <SelectContent>
                    <SelectItem value="hypertrophy">Hipertrofia</SelectItem>
                    <SelectItem value="strength">Força</SelectItem>
                    <SelectItem value="endurance">Resistência</SelectItem>
                    <SelectItem value="weight_loss">Emagrecimento</SelectItem>
                    <SelectItem value="conditioning">Condicionamento</SelectItem>
                    <SelectItem value="rehabilitation">Reabilitação</SelectItem>
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
                <Label>Duração (min)</Label>
                <Input type="number" min={10} max={300} value={durationMinutes} onChange={(e) => setDurationMinutes(Number(e.target.value))} />
              </div>
              <div className="space-y-2">
                <Label>Frequência Semanal</Label>
                <Input type="number" min={1} max={7} value={weeklyFrequency} onChange={(e) => setWeeklyFrequency(Number(e.target.value))} />
              </div>
            </div>

            <div className="space-y-2">
              <Label>Sessões</Label>
              {sessions.map((s, i) => (
                <div key={i} className="flex gap-2">
                  <Input
                    value={s.name}
                    onChange={(e) => {
                      const updated = [...sessions];
                      updated[i] = { name: e.target.value };
                      setSessions(updated);
                    }}
                    placeholder={`Treino ${String.fromCharCode(65 + i)}`}
                  />
                  {sessions.length > 1 && (
                    <Button type="button" variant="ghost" size="icon" onClick={() => setSessions(sessions.filter((_, j) => j !== i))}>
                      <Trash2 className="h-4 w-4" />
                    </Button>
                  )}
                </div>
              ))}
              <Button type="button" variant="outline" size="sm" onClick={() => setSessions([...sessions, { name: `Treino ${String.fromCharCode(65 + sessions.length)}` }])}>
                <Plus className="h-4 w-4 mr-1" /> Adicionar Sessão
              </Button>
            </div>

            <Button type="submit" disabled={createMutation.isPending} className="w-full">
              {createMutation.isPending ? "Criando..." : "Criar Plano"}
            </Button>
          </form>
        </CardContent>
      </Card>
    </div>
  );
}
