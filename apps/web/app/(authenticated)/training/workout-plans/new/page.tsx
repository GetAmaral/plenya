"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { Plus, Trash2, GripVertical } from "lucide-react";
import { toast } from "sonner";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Badge } from "@/components/ui/badge";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { useRequireAuth } from "@/lib/use-auth";
import { useRequireSelectedPatient } from "@/lib/use-require-selected-patient";
import { SelectedPatientHeader } from "@/components/patients/SelectedPatientHeader";
import { PageHeader } from "@/components/layout/page-header";
import {
  useCreateWorkoutPlan,
  type WorkoutObjective,
  type WorkoutIntensity,
  type ExercisePhase,
  type ExerciseCadence,
  type CreateSessionDTO,
  type CreateSessionExerciseDTO,
} from "@/lib/api/workout-plan-api";
import type { Exercise } from "@/lib/api/exercise-api";
import { ExercisePicker } from "@/components/training/ExercisePicker";


interface SessionExerciseLocal extends CreateSessionExerciseDTO {
  _exercise?: Exercise; // local display only
}

interface SessionLocal {
  name: string;
  exercises: SessionExerciseLocal[];
}

const phaseLabels: Record<ExercisePhase, string> = {
  warmup: "Aquecimento",
  main: "Principal",
  cooldown: "Finalizacao",
};

const cadenceOptions: { value: ExerciseCadence; label: string }[] = [
  { value: "normal", label: "Normal" },
  { value: "slow", label: "Lento" },
  { value: "paused", label: "Pausado" },
  { value: "explosive", label: "Explosivo" },
  { value: "free", label: "Livre" },
];

export default function NewWorkoutPlanPage() {
  useRequireAuth();
  useRequireSelectedPatient();
  const router = useRouter();

  const [name, setName] = useState("");
  const [objective, setObjective] = useState<WorkoutObjective>("hypertrophy");
  const [intensity, setIntensity] = useState<WorkoutIntensity>("moderate");
  const [durationMinutes, setDurationMinutes] = useState(60);
  const [weeklyFrequency, setWeeklyFrequency] = useState(3);
  const [sessions, setSessions] = useState<SessionLocal[]>([{ name: "Treino A", exercises: [] }]);
  const [activeSession, setActiveSession] = useState(0);
  const [activePhase, setActivePhase] = useState<ExercisePhase>("main");

  const createMutation = useCreateWorkoutPlan();

  const addExercise = (exercise: Exercise) => {
    const updated = [...sessions];
    const session = updated[activeSession];
    const phaseExercises = session.exercises.filter((e) => e.phase === activePhase);
    session.exercises.push({
      exerciseId: exercise.id,
      phase: activePhase,
      order: phaseExercises.length,
      sets: 3,
      reps: "10",
      cadence: "normal",
      restBetweenSetsSec: 60,
      restBetweenExercisesSec: 90,
      _exercise: exercise,
    });
    setSessions(updated);
  };

  const removeExercise = (sessionIdx: number, exIdx: number) => {
    const updated = [...sessions];
    updated[sessionIdx].exercises.splice(exIdx, 1);
    setSessions(updated);
  };

  const updateExercise = (sessionIdx: number, exIdx: number, field: string, value: string | number) => {
    const updated = [...sessions];
    (updated[sessionIdx].exercises[exIdx] as any)[field] = value;
    setSessions(updated);
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!name) {
      toast.error("Nome do plano e obrigatorio");
      return;
    }

    try {
      const sessionDtos: CreateSessionDTO[] = sessions.map((s, i) => ({
        name: s.name,
        order: i,
        exercises: s.exercises.map((ex, j) => ({
          exerciseId: ex.exerciseId,
          phase: ex.phase,
          order: j,
          sets: ex.sets,
          reps: ex.reps,
          cadence: ex.cadence,
          restBetweenSetsSec: ex.restBetweenSetsSec,
          restBetweenExercisesSec: ex.restBetweenExercisesSec,
          notes: ex.notes,
        })),
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

  const currentSession = sessions[activeSession];
  const currentPhaseExercises = currentSession?.exercises.filter((e) => e.phase === activePhase) || [];

  return (
    <div className="space-y-6">
      <SelectedPatientHeader />
      <PageHeader
        title="Novo Plano de Treino"
        description="Configure o plano e adicione exercicios"
      />

      <form onSubmit={handleSubmit} className="space-y-6">
        {/* Plan Config */}
        <Card>
          <CardHeader>
            <CardTitle className="text-base">Dados do Plano</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <div className="space-y-2">
              <Label>Nome do Plano</Label>
              <Input value={name} onChange={(e) => setName(e.target.value)} placeholder="Ex: Plano Hipertrofia Marco" required />
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
          </CardContent>
        </Card>

        {/* Sessions */}
        <Card>
          <CardHeader>
            <div className="flex items-center justify-between">
              <CardTitle className="text-base">Sessoes</CardTitle>
              <Button
                type="button"
                variant="outline"
                size="sm"
                onClick={() => {
                  setSessions([...sessions, { name: `Treino ${String.fromCharCode(65 + sessions.length)}`, exercises: [] }]);
                }}
              >
                <Plus className="h-4 w-4 mr-1" /> Sessao
              </Button>
            </div>
          </CardHeader>
          <CardContent>
            {/* Session tabs */}
            <div className="flex gap-2 mb-4 flex-wrap">
              {sessions.map((s, i) => (
                <div key={i} className="flex items-center gap-1">
                  <Button
                    type="button"
                    variant={activeSession === i ? "default" : "outline"}
                    size="sm"
                    onClick={() => setActiveSession(i)}
                  >
                    {s.name}
                    <Badge variant="secondary" className="ml-1 text-[10px]">{s.exercises.length}</Badge>
                  </Button>
                  {sessions.length > 1 && (
                    <Button
                      type="button"
                      variant="ghost"
                      size="icon"
                      className="h-7 w-7"
                      onClick={() => {
                        setSessions(sessions.filter((_, j) => j !== i));
                        if (activeSession >= sessions.length - 1) setActiveSession(Math.max(0, sessions.length - 2));
                      }}
                    >
                      <Trash2 className="h-3 w-3" />
                    </Button>
                  )}
                </div>
              ))}
            </div>

            {/* Session name */}
            <div className="mb-4">
              <Input
                value={currentSession?.name || ""}
                onChange={(e) => {
                  const updated = [...sessions];
                  updated[activeSession].name = e.target.value;
                  setSessions(updated);
                }}
                placeholder="Nome da sessao"
              />
            </div>

            {/* Phase tabs */}
            <Tabs value={activePhase} onValueChange={(v) => setActivePhase(v as ExercisePhase)}>
              <div className="flex items-center justify-between mb-3">
                <TabsList>
                  <TabsTrigger value="warmup">Aquecimento</TabsTrigger>
                  <TabsTrigger value="main">Principal</TabsTrigger>
                  <TabsTrigger value="cooldown">Finalizacao</TabsTrigger>
                </TabsList>
                <ExercisePicker onSelect={addExercise} />
              </div>

              {(["warmup", "main", "cooldown"] as const).map((phase) => (
                <TabsContent key={phase} value={phase} className="space-y-2">
                  {currentSession?.exercises
                    .map((ex, originalIdx) => ({ ex, originalIdx }))
                    .filter(({ ex }) => ex.phase === phase)
                    .map(({ ex, originalIdx }) => (
                      <div key={originalIdx} className="flex items-start gap-2 p-3 rounded-lg border bg-card">
                        <div className="flex-1 space-y-2">
                          <div className="flex items-center gap-2">
                            <GripVertical className="h-4 w-4 text-muted-foreground flex-shrink-0" />
                            <span className="font-medium text-sm truncate">
                              {ex._exercise?.namePt || ex._exercise?.name || "Exercicio"}
                            </span>
                            <Button
                              type="button"
                              variant="ghost"
                              size="icon"
                              className="h-6 w-6 ml-auto"
                              onClick={() => removeExercise(activeSession, originalIdx)}
                            >
                              <Trash2 className="h-3 w-3" />
                            </Button>
                          </div>
                          <div className="grid grid-cols-5 gap-2">
                            <div>
                              <Label className="text-[10px]">Series</Label>
                              <Input
                                type="number"
                                min={1}
                                value={ex.sets}
                                onChange={(e) => updateExercise(activeSession, originalIdx, "sets", Number(e.target.value))}
                                className="h-8 text-xs"
                              />
                            </div>
                            <div>
                              <Label className="text-[10px]">Reps</Label>
                              <Input
                                value={ex.reps}
                                onChange={(e) => updateExercise(activeSession, originalIdx, "reps", e.target.value)}
                                className="h-8 text-xs"
                                placeholder="10"
                              />
                            </div>
                            <div>
                              <Label className="text-[10px]">Cadencia</Label>
                              <Select
                                value={ex.cadence}
                                onValueChange={(v) => updateExercise(activeSession, originalIdx, "cadence", v)}
                              >
                                <SelectTrigger className="h-8 text-xs"><SelectValue /></SelectTrigger>
                                <SelectContent>
                                  {cadenceOptions.map((c) => (
                                    <SelectItem key={c.value} value={c.value}>{c.label}</SelectItem>
                                  ))}
                                </SelectContent>
                              </Select>
                            </div>
                            <div>
                              <Label className="text-[10px]">Desc.(s)</Label>
                              <Input
                                type="number"
                                min={0}
                                value={ex.restBetweenSetsSec}
                                onChange={(e) => updateExercise(activeSession, originalIdx, "restBetweenSetsSec", Number(e.target.value))}
                                className="h-8 text-xs"
                              />
                            </div>
                            <div>
                              <Label className="text-[10px]">Int.(s)</Label>
                              <Input
                                type="number"
                                min={0}
                                value={ex.restBetweenExercisesSec}
                                onChange={(e) => updateExercise(activeSession, originalIdx, "restBetweenExercisesSec", Number(e.target.value))}
                                className="h-8 text-xs"
                              />
                            </div>
                          </div>
                        </div>
                      </div>
                    ))}
                  {currentSession?.exercises.filter((e) => e.phase === phase).length === 0 && (
                    <p className="text-center text-muted-foreground text-sm py-6">
                      Nenhum exercicio no {phaseLabels[phase].toLowerCase()}. Use o botao acima para adicionar.
                    </p>
                  )}
                </TabsContent>
              ))}
            </Tabs>
          </CardContent>
        </Card>

        <Button type="submit" disabled={createMutation.isPending} className="w-full" size="lg">
          {createMutation.isPending ? "Criando..." : "Criar Plano de Treino"}
        </Button>
      </form>
    </div>
  );
}
