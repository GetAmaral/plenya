"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { format } from "date-fns";
import { Sparkles } from "lucide-react";
import { toast } from "sonner";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import { useRequireAuth } from "@/lib/use-auth";
import { useRequireSelectedPatient } from "@/lib/use-require-selected-patient";
import { SelectedPatientHeader } from "@/components/patients/SelectedPatientHeader";
import { PageHeader } from "@/components/layout/page-header";
import { useCreatePeriodization, useGeneratePeriodization, type PeriodizationFramework } from "@/lib/api/periodization-api";

export default function NewPeriodizationPage() {
  useRequireAuth();
  useRequireSelectedPatient();
  const router = useRouter();

  const [framework, setFramework] = useState<PeriodizationFramework>("bompa");
  const [startDate, setStartDate] = useState(format(new Date(), "yyyy-MM-dd"));
  const [totalWeeks, setTotalWeeks] = useState(24);
  const [objective, setObjective] = useState("");

  const createMutation = useCreatePeriodization();
  const generateMutation = useGeneratePeriodization();

  const handleSubmit = async (e: React.FormEvent, useAI: boolean) => {
    e.preventDefault();
    if (!objective) {
      toast.error("Objetivo e obrigatorio");
      return;
    }

    const data = { framework, startDate, totalWeeks, objective };

    try {
      const mutation = useAI ? generateMutation : createMutation;
      const result = await mutation.mutateAsync(data);
      toast.success(useAI ? "Periodizacao gerada com IA!" : "Periodizacao criada!");
      router.push(`/training/periodization/${result.id}`);
    } catch (error: any) {
      toast.error(error?.message || "Erro ao criar periodizacao");
    }
  };

  const isPending = createMutation.isPending || generateMutation.isPending;

  return (
    <div className="space-y-6">
      <SelectedPatientHeader />
      <PageHeader
        title="Nova Periodizacao"
        description="Configure o macrociclo de treinamento"
      />

      <Card className="max-w-xl">
        <CardHeader>
          <CardTitle>Dados da Periodizacao</CardTitle>
        </CardHeader>
        <CardContent>
          <form onSubmit={(e) => handleSubmit(e, false)} className="space-y-4">
            <div className="space-y-2">
              <Label>Objetivo</Label>
              <Input value={objective} onChange={(e) => setObjective(e.target.value)} placeholder="Ex: Hipertrofia para competicao" required />
            </div>

            <div className="space-y-2">
              <Label>Framework</Label>
              <Select value={framework} onValueChange={(v) => setFramework(v as PeriodizationFramework)}>
                <SelectTrigger><SelectValue /></SelectTrigger>
                <SelectContent>
                  <SelectItem value="bompa">Bompa</SelectItem>
                  <SelectItem value="linear">Linear</SelectItem>
                  <SelectItem value="undulating">Ondulatoria</SelectItem>
                  <SelectItem value="block">Bloco</SelectItem>
                </SelectContent>
              </Select>
            </div>

            <div className="grid grid-cols-2 gap-4">
              <div className="space-y-2">
                <Label>Data de Inicio</Label>
                <Input type="date" value={startDate} onChange={(e) => setStartDate(e.target.value)} required />
              </div>
              <div className="space-y-2">
                <Label>Total de Semanas</Label>
                <Input type="number" min={1} max={104} value={totalWeeks} onChange={(e) => setTotalWeeks(Number(e.target.value))} />
              </div>
            </div>

            <div className="flex gap-2">
              <Button type="submit" disabled={isPending} className="flex-1">
                {createMutation.isPending ? "Criando..." : "Criar Manual"}
              </Button>
              <Button
                type="button"
                disabled={isPending}
                variant="default"
                className="flex-1"
                onClick={(e) => handleSubmit(e, true)}
              >
                <Sparkles className="h-4 w-4 mr-2" />
                {generateMutation.isPending ? "Gerando com IA..." : "Gerar com IA"}
              </Button>
            </div>
          </form>
        </CardContent>
      </Card>
    </div>
  );
}
