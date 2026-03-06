"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { format } from "date-fns";
import { ArrowLeft, Timer } from "lucide-react";
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
import { useCreatePeriodization, type PeriodizationFramework } from "@/lib/api/periodization-api";
import Link from "next/link";

export default function NewPeriodizationPage() {
  useRequireAuth();
  useRequireSelectedPatient();
  const router = useRouter();

  const [framework, setFramework] = useState<PeriodizationFramework>("bompa");
  const [startDate, setStartDate] = useState(format(new Date(), "yyyy-MM-dd"));
  const [totalWeeks, setTotalWeeks] = useState(24);
  const [objective, setObjective] = useState("");

  const createMutation = useCreatePeriodization();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!objective) {
      toast.error("Objetivo é obrigatório");
      return;
    }

    try {
      const result = await createMutation.mutateAsync({
        framework,
        startDate,
        totalWeeks,
        objective,
      });
      toast.success("Periodização criada com sucesso!");
      router.push("/training/periodization");
    } catch (error: any) {
      toast.error(error?.message || "Erro ao criar periodização");
    }
  };

  return (
    <div className="space-y-6">
      <SelectedPatientHeader />
      <PageHeader
        title="Nova Periodização"
        subtitle="Configure o macrociclo de treinamento"
        icon={Timer}
        actions={
          <Link href="/training/periodization">
            <Button variant="outline">
              <ArrowLeft className="h-4 w-4 mr-2" />
              Voltar
            </Button>
          </Link>
        }
      />

      <Card className="max-w-xl">
        <CardHeader>
          <CardTitle>Dados da Periodização</CardTitle>
        </CardHeader>
        <CardContent>
          <form onSubmit={handleSubmit} className="space-y-4">
            <div className="space-y-2">
              <Label>Objetivo</Label>
              <Input value={objective} onChange={(e) => setObjective(e.target.value)} placeholder="Ex: Hipertrofia para competição" required />
            </div>

            <div className="space-y-2">
              <Label>Framework</Label>
              <Select value={framework} onValueChange={(v) => setFramework(v as PeriodizationFramework)}>
                <SelectTrigger><SelectValue /></SelectTrigger>
                <SelectContent>
                  <SelectItem value="bompa">Bompa</SelectItem>
                  <SelectItem value="linear">Linear</SelectItem>
                  <SelectItem value="undulating">Ondulatória</SelectItem>
                  <SelectItem value="block">Bloco</SelectItem>
                </SelectContent>
              </Select>
            </div>

            <div className="grid grid-cols-2 gap-4">
              <div className="space-y-2">
                <Label>Data de Início</Label>
                <Input type="date" value={startDate} onChange={(e) => setStartDate(e.target.value)} required />
              </div>
              <div className="space-y-2">
                <Label>Total de Semanas</Label>
                <Input type="number" min={1} max={104} value={totalWeeks} onChange={(e) => setTotalWeeks(Number(e.target.value))} />
              </div>
            </div>

            <Button type="submit" disabled={createMutation.isPending} className="w-full">
              {createMutation.isPending ? "Criando..." : "Criar Periodização"}
            </Button>
          </form>
        </CardContent>
      </Card>
    </div>
  );
}
