"use client";

import { useRef, useState } from "react";
import { useRouter } from "next/navigation";
import { format } from "date-fns";
import { toast } from "sonner";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import { Textarea } from "@/components/ui/textarea";
import { useRequireAuth } from "@/lib/use-auth";
import { useRequireSelectedPatient } from "@/lib/use-require-selected-patient";
import { SelectedPatientHeader } from "@/components/patients/SelectedPatientHeader";
import { PageHeader } from "@/components/layout/page-header";
import { useFormNavigation } from "@/lib/use-form-navigation";
import { useCreatePosturalAssessment, type CreatePosturalAssessmentDTO } from "@/lib/api/posture-api";

function NumberInput({ label, value, onChange, unit, step = "0.1", min, max }: {
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

export default function NewPosturalAssessmentPage() {
  useRequireAuth();
  useRequireSelectedPatient();
  const router = useRouter();
  const formRef = useRef<HTMLFormElement>(null);
  useFormNavigation({ formRef });

  const [assessmentDate, setAssessmentDate] = useState(format(new Date(), "yyyy-MM-dd"));
  const [viewType, setViewType] = useState("");

  // Measurements (degrees)
  const [shoulderDeviation, setShoulderDeviation] = useState("");
  const [hipDeviation, setHipDeviation] = useState("");
  const [headLateralDeviation, setHeadLateralDeviation] = useState("");
  const [fhp, setFhp] = useState("");
  const [thoracicKyphosis, setThoracicKyphosis] = useState("");
  const [lumbarLordosis, setLumbarLordosis] = useState("");
  const [kneeAngle, setKneeAngle] = useState("");

  const [notes, setNotes] = useState("");

  const createMutation = useCreatePosturalAssessment();

  const toNum = (v: string) => v ? Number(v) : undefined;

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    const data: CreatePosturalAssessmentDTO = {
      assessmentDate,
      viewType: viewType as CreatePosturalAssessmentDTO["viewType"],
      shoulderDeviation: toNum(shoulderDeviation),
      hipDeviation: toNum(hipDeviation),
      headLateralDeviation: toNum(headLateralDeviation),
      fhp: toNum(fhp),
      thoracicKyphosis: toNum(thoracicKyphosis),
      lumbarLordosis: toNum(lumbarLordosis),
      kneeAngle: toNum(kneeAngle),
      notes: notes || undefined,
    };

    try {
      const result = await createMutation.mutateAsync(data);
      toast.success("Avaliacao postural criada com sucesso!");
      router.push(`/training/posture/${result.id}`);
    } catch (error: any) {
      toast.error(error?.message || "Erro ao criar avaliacao postural");
    }
  };

  return (
    <div className="space-y-6">
      <SelectedPatientHeader />
      <PageHeader
        title="Nova Avaliacao Postural"
        description="Preencha os desvios posturais em graus para gerar a avaliacao automatica."
      />

      <form ref={formRef} onSubmit={handleSubmit} className="space-y-6 max-w-3xl">
        {/* Data */}
        <Card>
          <CardHeader><CardTitle className="text-base">Data da Avaliacao</CardTitle></CardHeader>
          <CardContent>
            <Input type="date" value={assessmentDate} onChange={(e) => setAssessmentDate(e.target.value)} required />
          </CardContent>
        </Card>

        {/* Tipo de Vista */}
        <Card>
          <CardHeader><CardTitle className="text-base">Tipo de Vista</CardTitle></CardHeader>
          <CardContent>
            <div className="space-y-1">
              <Label className="text-xs">Vista</Label>
              <Select value={viewType} onValueChange={setViewType}>
                <SelectTrigger><SelectValue placeholder="Selecione a vista" /></SelectTrigger>
                <SelectContent>
                  <SelectItem value="front">Anterior (Frente)</SelectItem>
                  <SelectItem value="side_left">Lateral Esquerda</SelectItem>
                  <SelectItem value="side_right">Lateral Direita</SelectItem>
                  <SelectItem value="back">Posterior (Costas)</SelectItem>
                </SelectContent>
              </Select>
            </div>
          </CardContent>
        </Card>

        {/* Desvios Posturais */}
        <Card>
          <CardHeader><CardTitle className="text-base">Desvios Posturais</CardTitle></CardHeader>
          <CardContent>
            <div className="grid grid-cols-2 md:grid-cols-3 gap-4">
              <NumberInput label="Desvio Ombro" value={shoulderDeviation} onChange={setShoulderDeviation} unit="graus" min="0" max="90" />
              <NumberInput label="Desvio Quadril" value={hipDeviation} onChange={setHipDeviation} unit="graus" min="0" max="90" />
              <NumberInput label="Desvio Lateral Cabeca" value={headLateralDeviation} onChange={setHeadLateralDeviation} unit="graus" min="0" max="90" />
              <NumberInput label="FHP - Anterioridade Cabeca" value={fhp} onChange={setFhp} unit="graus" min="0" max="90" />
              <NumberInput label="Cifose Toracica" value={thoracicKyphosis} onChange={setThoracicKyphosis} unit="graus" min="0" max="90" />
              <NumberInput label="Lordose Lombar" value={lumbarLordosis} onChange={setLumbarLordosis} unit="graus" min="0" max="90" />
              <NumberInput label="Angulo Joelho" value={kneeAngle} onChange={setKneeAngle} unit="graus" min="0" max="90" />
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
              placeholder="Observacoes adicionais sobre a postura do paciente..."
              rows={3}
            />
          </CardContent>
        </Card>

        <Button type="submit" disabled={createMutation.isPending} className="w-full">
          {createMutation.isPending ? "Avaliando postura..." : "Avaliar Postura"}
        </Button>
      </form>
    </div>
  );
}
