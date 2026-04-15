"use client";

import { useRef, useState } from "react";
import { useRouter } from "next/navigation";
import { format } from "date-fns";
import { toast } from "sonner";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Textarea } from "@/components/ui/textarea";
import { useRequireAuth } from "@/lib/use-auth";
import { useRequireSelectedPatient } from "@/lib/use-require-selected-patient";
import { SelectedPatientHeader } from "@/components/patients/SelectedPatientHeader";
import { PageHeader } from "@/components/layout/page-header";
import { useFormNavigation } from "@/lib/use-form-navigation";
import { useCreateFitnessTest, type CreateFitnessTestDTO } from "@/lib/api/fitness-test-api";

function NumberInput({ label, value, onChange, unit, step = "1", min, max }: {
  label: string; value: string; onChange: (v: string) => void; unit?: string;
  step?: string; min?: string; max?: string;
}) {
  return (
    <div className="space-y-1">
      <Label className="text-xs">{label}{unit && <span className="text-muted-foreground ml-1">({unit})</span>}</Label>
      <Input type="number" step={step} min={min} max={max} value={value} onChange={(e) => onChange(e.target.value)} />
    </div>
  );
}

export default function NewFitnessTestPage() {
  useRequireAuth();
  useRequireSelectedPatient();
  const router = useRouter();
  const formRef = useRef<HTMLFormElement>(null);
  useFormNavigation({ formRef });

  const [assessmentDate, setAssessmentDate] = useState(format(new Date(), "yyyy-MM-dd"));

  // 5 fitness tests
  const [abdominal, setAbdominal] = useState("");
  const [flexao, setFlexao] = useState("");
  const [prancha, setPrancha] = useState("");
  const [burpee, setBurpee] = useState("");
  const [frt, setFrt] = useState("");

  // Notes
  const [notes, setNotes] = useState("");

  const createMutation = useCreateFitnessTest();

  const toInt = (v: string) => v ? Math.round(Number(v)) : undefined;

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    const data: CreateFitnessTestDTO = {
      assessmentDate,
      abdominal: toInt(abdominal),
      flexao: toInt(flexao),
      prancha: toInt(prancha),
      burpee: toInt(burpee),
      frt: toInt(frt),
      notes: notes || undefined,
    };

    try {
      const result = await createMutation.mutateAsync(data);
      toast.success("Teste de fitness criado com sucesso!");
      router.push(`/training/fitness-tests/${result.id}`);
    } catch (error: any) {
      toast.error(error?.message || "Erro ao criar teste de fitness");
    }
  };

  return (
    <div className="space-y-6">
      <SelectedPatientHeader />
      <PageHeader
        title="Novo Teste de Fitness"
        description="Avalie a aptidao fisica do paciente com 5 testes: abdominal, flexao, prancha, burpee e FRT. Os resultados sao classificados automaticamente."
      />

      <form ref={formRef} onSubmit={handleSubmit} className="space-y-6 max-w-3xl">
        {/* Data */}
        <Card>
          <CardHeader><CardTitle className="text-base">Data da Avaliacao</CardTitle></CardHeader>
          <CardContent>
            <Input type="date" value={assessmentDate} onChange={(e) => setAssessmentDate(e.target.value)} required />
          </CardContent>
        </Card>

        {/* 5 Testes */}
        <Card>
          <CardHeader>
            <CardTitle className="text-base">Testes de Fitness</CardTitle>
            <p className="text-xs text-muted-foreground">Preencha os resultados de cada teste. Deixe em branco os testes nao realizados.</p>
          </CardHeader>
          <CardContent>
            <div className="grid grid-cols-2 md:grid-cols-3 gap-4">
              <NumberInput label="Abdominal" value={abdominal} onChange={setAbdominal} unit="rep/min" min="0" max="200" />
              <NumberInput label="Flexao" value={flexao} onChange={setFlexao} unit="rep/min" min="0" max="200" />
              <NumberInput label="Prancha" value={prancha} onChange={setPrancha} unit="segundos" min="0" max="600" />
              <NumberInput label="Burpee" value={burpee} onChange={setBurpee} unit="ciclos/3min" min="0" max="100" />
              <NumberInput label="FRT" value={frt} onChange={setFrt} unit="rep/90s" min="0" max="200" />
            </div>
          </CardContent>
        </Card>

        {/* Observacoes */}
        <Card>
          <CardHeader><CardTitle className="text-base">Observacoes</CardTitle></CardHeader>
          <CardContent>
            <Textarea
              value={notes}
              onChange={(e) => setNotes(e.target.value)}
              placeholder="Observacoes sobre a execucao dos testes, limitacoes, etc."
              rows={3}
            />
          </CardContent>
        </Card>

        <Button type="submit" disabled={createMutation.isPending} className="w-full">
          {createMutation.isPending ? "Avaliando..." : "Avaliar Fitness"}
        </Button>
      </form>
    </div>
  );
}
