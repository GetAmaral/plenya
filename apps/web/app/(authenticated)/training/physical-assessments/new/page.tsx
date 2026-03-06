"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { format } from "date-fns";
import { Activity, ArrowLeft } from "lucide-react";
import { toast } from "sonner";
import { Card, CardContent, CardHeader, CardTitle, CardDescription } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import { Skeleton } from "@/components/ui/skeleton";
import { useRequireAuth } from "@/lib/use-auth";
import { useRequireSelectedPatient } from "@/lib/use-require-selected-patient";
import { SelectedPatientHeader } from "@/components/patients/SelectedPatientHeader";
import { PageHeader } from "@/components/layout/page-header";
import { useCreatePhysicalAssessment } from "@/lib/api/physical-assessment-api";
import { apiClient } from "@/lib/api-client";
import { useQuery } from "@tanstack/react-query";
import Link from "next/link";

interface AnamnesisOption {
  id: string;
  consultationDate: string;
  displayTitle: string;
}

export default function NewPhysicalAssessmentPage() {
  useRequireAuth();
  useRequireSelectedPatient();
  const router = useRouter();

  const [anamnesisId, setAnamnesisId] = useState("");
  const [assessmentDate, setAssessmentDate] = useState(format(new Date(), "yyyy-MM-dd"));

  const { data: anamneses, isLoading: loadingAnamneses } = useQuery({
    queryKey: ["anamnesis-for-assessment"],
    queryFn: () => apiClient.get<AnamnesisOption[]>("/api/v1/anamnesis?limit=50"),
  });

  const createMutation = useCreatePhysicalAssessment();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!anamnesisId) {
      toast.error("Selecione uma anamnese");
      return;
    }

    try {
      const result = await createMutation.mutateAsync({
        anamnesisId,
        assessmentDate,
      });
      toast.success("Avaliação física criada com sucesso!");
      router.push(`/training/physical-assessments/${result.id}`);
    } catch (error: any) {
      toast.error(error?.message || "Erro ao criar avaliação");
    }
  };

  return (
    <div className="space-y-6">
      <SelectedPatientHeader />
      <PageHeader
        title="Nova Avaliação Física"
        subtitle="Selecione a anamnese e gere a avaliação ACSM automaticamente"
        icon={Activity}
        actions={
          <Link href="/training/physical-assessments">
            <Button variant="outline">
              <ArrowLeft className="h-4 w-4 mr-2" />
              Voltar
            </Button>
          </Link>
        }
      />

      <Card className="max-w-xl">
        <CardHeader>
          <CardTitle>Dados da Avaliação</CardTitle>
          <CardDescription>
            O sistema buscará automaticamente peso, altura, IMC, BRI e resultados de exames
            da anamnese selecionada para calcular a estratificação ACSM.
          </CardDescription>
        </CardHeader>
        <CardContent>
          <form onSubmit={handleSubmit} className="space-y-4">
            <div className="space-y-2">
              <Label htmlFor="assessmentDate">Data da Avaliação</Label>
              <Input
                id="assessmentDate"
                type="date"
                value={assessmentDate}
                onChange={(e) => setAssessmentDate(e.target.value)}
                required
              />
            </div>

            <div className="space-y-2">
              <Label htmlFor="anamnesisId">Anamnese de Referência</Label>
              {loadingAnamneses ? (
                <Skeleton className="h-10 w-full" />
              ) : (
                <Select value={anamnesisId} onValueChange={setAnamnesisId}>
                  <SelectTrigger>
                    <SelectValue placeholder="Selecione a anamnese com medidas físicas" />
                  </SelectTrigger>
                  <SelectContent>
                    {anamneses?.map((a) => (
                      <SelectItem key={a.id} value={a.id}>
                        {a.displayTitle}
                      </SelectItem>
                    ))}
                  </SelectContent>
                </Select>
              )}
              <p className="text-xs text-muted-foreground">
                A anamnese deve conter medidas como peso, altura, cintura e dados de saúde.
              </p>
            </div>

            <Button type="submit" disabled={createMutation.isPending} className="w-full">
              {createMutation.isPending ? "Calculando ACSM..." : "Gerar Avaliação ACSM"}
            </Button>
          </form>
        </CardContent>
      </Card>
    </div>
  );
}
