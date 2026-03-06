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
import { Switch } from "@/components/ui/switch";
import { useRequireAuth } from "@/lib/use-auth";
import { useRequireSelectedPatient } from "@/lib/use-require-selected-patient";
import { SelectedPatientHeader } from "@/components/patients/SelectedPatientHeader";
import { PageHeader } from "@/components/layout/page-header";
import { useFormNavigation } from "@/lib/use-form-navigation";
import { useCreatePhysicalAssessment, type CreatePhysicalAssessmentDTO } from "@/lib/api/physical-assessment-api";

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

export default function NewPhysicalAssessmentPage() {
  useRequireAuth();
  useRequireSelectedPatient();
  const router = useRouter();
  const formRef = useRef<HTMLFormElement>(null);
  useFormNavigation({ formRef });

  const [assessmentDate, setAssessmentDate] = useState(format(new Date(), "yyyy-MM-dd"));

  // Antropometria
  const [weight, setWeight] = useState("");
  const [height, setHeight] = useState("");
  const [waistCircumference, setWaistCircumference] = useState("");

  // Composição corporal
  const [bodyFatPercent, setBodyFatPercent] = useState("");
  const [leanMass, setLeanMass] = useState("");

  // Cardiovascular
  const [systolicBp, setSystolicBp] = useState("");
  const [diastolicBp, setDiastolicBp] = useState("");
  const [restingHeartRate, setRestingHeartRate] = useState("");

  // Laboratorial
  const [ldl, setLdl] = useState("");
  const [hdl, setHdl] = useState("");
  const [totalCholesterol, setTotalCholesterol] = useState("");
  const [triglycerides, setTriglycerides] = useState("");
  const [fastingGlucose, setFastingGlucose] = useState("");
  const [hbA1c, setHbA1c] = useState("");

  // Histórico
  const [familyHistory, setFamilyHistory] = useState(false);
  const [smokingStatus, setSmokingStatus] = useState("");
  const [physicalActivityLevel, setPhysicalActivityLevel] = useState("");

  // Condições
  const [cardiovascularDisease, setCardiovascularDisease] = useState(false);
  const [diabetesType, setDiabetesType] = useState("");
  const [symptoms, setSymptoms] = useState("");
  const [clinicalAlert, setClinicalAlert] = useState(false);

  const createMutation = useCreatePhysicalAssessment();

  const toNum = (v: string) => v ? Number(v) : undefined;
  const toInt = (v: string) => v ? Math.round(Number(v)) : undefined;

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    const data: CreatePhysicalAssessmentDTO = {
      assessmentDate,
      weight: toNum(weight),
      height: toNum(height),
      waistCircumference: toNum(waistCircumference),
      bodyFatPercent: toNum(bodyFatPercent),
      leanMass: toNum(leanMass),
      systolicBp: toInt(systolicBp),
      diastolicBp: toInt(diastolicBp),
      restingHeartRate: toInt(restingHeartRate),
      ldl: toNum(ldl),
      hdl: toNum(hdl),
      totalCholesterol: toNum(totalCholesterol),
      triglycerides: toNum(triglycerides),
      fastingGlucose: toNum(fastingGlucose),
      hbA1c: toNum(hbA1c),
      familyHistory: familyHistory || undefined,
      smokingStatus: smokingStatus as CreatePhysicalAssessmentDTO['smokingStatus'] || undefined,
      physicalActivityLevel: physicalActivityLevel as CreatePhysicalAssessmentDTO['physicalActivityLevel'] || undefined,
      cardiovascularDisease: cardiovascularDisease || undefined,
      diabetesType: diabetesType || undefined,
      symptoms: symptoms || undefined,
      clinicalAlert: clinicalAlert || undefined,
    };

    try {
      const result = await createMutation.mutateAsync(data);
      toast.success("Avaliacao fisica criada com sucesso!");
      router.push(`/training/physical-assessments/${result.id}`);
    } catch (error: any) {
      toast.error(error?.message || "Erro ao criar avaliacao");
    }
  };

  return (
    <div className="space-y-6">
      <SelectedPatientHeader />
      <PageHeader
        title="Nova Avaliacao Fisica"
        description="Preencha os dados para calcular a estratificacao ACSM. Dados laboratoriais sao preenchidos automaticamente se houver exames recentes."
      />

      <form ref={formRef} onSubmit={handleSubmit} className="space-y-6 max-w-3xl">
        {/* Data */}
        <Card>
          <CardHeader><CardTitle className="text-base">Data da Avaliacao</CardTitle></CardHeader>
          <CardContent>
            <Input type="date" value={assessmentDate} onChange={(e) => setAssessmentDate(e.target.value)} required />
          </CardContent>
        </Card>

        {/* Antropometria */}
        <Card>
          <CardHeader><CardTitle className="text-base">Antropometria</CardTitle></CardHeader>
          <CardContent>
            <div className="grid grid-cols-2 md:grid-cols-3 gap-4">
              <NumberInput label="Peso" value={weight} onChange={setWeight} unit="kg" min="20" max="300" />
              <NumberInput label="Altura" value={height} onChange={setHeight} unit="cm" min="50" max="250" />
              <NumberInput label="Circunferencia Abdominal" value={waistCircumference} onChange={setWaistCircumference} unit="cm" min="40" max="200" />
            </div>
          </CardContent>
        </Card>

        {/* Composição Corporal */}
        <Card>
          <CardHeader><CardTitle className="text-base">Composicao Corporal</CardTitle></CardHeader>
          <CardContent>
            <div className="grid grid-cols-2 gap-4">
              <NumberInput label="% Gordura Corporal" value={bodyFatPercent} onChange={setBodyFatPercent} unit="%" min="2" max="60" />
              <NumberInput label="Massa Magra" value={leanMass} onChange={setLeanMass} unit="kg" min="10" max="150" />
            </div>
            <p className="text-xs text-muted-foreground mt-2">IMC e BRI sao calculados automaticamente.</p>
          </CardContent>
        </Card>

        {/* Cardiovascular */}
        <Card>
          <CardHeader><CardTitle className="text-base">Cardiovascular</CardTitle></CardHeader>
          <CardContent>
            <div className="grid grid-cols-2 md:grid-cols-3 gap-4">
              <NumberInput label="PA Sistolica" value={systolicBp} onChange={setSystolicBp} unit="mmHg" step="1" min="60" max="250" />
              <NumberInput label="PA Diastolica" value={diastolicBp} onChange={setDiastolicBp} unit="mmHg" step="1" min="30" max="150" />
              <NumberInput label="FC Repouso" value={restingHeartRate} onChange={setRestingHeartRate} unit="bpm" step="1" min="30" max="200" />
            </div>
          </CardContent>
        </Card>

        {/* Laboratorial */}
        <Card>
          <CardHeader>
            <CardTitle className="text-base">Laboratorial</CardTitle>
            <p className="text-xs text-muted-foreground">Deixe em branco para preencher automaticamente do ultimo exame do paciente.</p>
          </CardHeader>
          <CardContent>
            <div className="grid grid-cols-2 md:grid-cols-3 gap-4">
              <NumberInput label="LDL" value={ldl} onChange={setLdl} unit="mg/dL" />
              <NumberInput label="HDL" value={hdl} onChange={setHdl} unit="mg/dL" />
              <NumberInput label="Colesterol Total" value={totalCholesterol} onChange={setTotalCholesterol} unit="mg/dL" />
              <NumberInput label="Triglicerideos" value={triglycerides} onChange={setTriglycerides} unit="mg/dL" />
              <NumberInput label="Glicemia Jejum" value={fastingGlucose} onChange={setFastingGlucose} unit="mg/dL" />
              <NumberInput label="HbA1c" value={hbA1c} onChange={setHbA1c} unit="%" />
            </div>
          </CardContent>
        </Card>

        {/* Histórico / Fatores de Risco */}
        <Card>
          <CardHeader><CardTitle className="text-base">Historico e Fatores de Risco ACSM</CardTitle></CardHeader>
          <CardContent className="space-y-4">
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div className="space-y-1">
                <Label className="text-xs">Tabagismo</Label>
                <Select value={smokingStatus} onValueChange={setSmokingStatus}>
                  <SelectTrigger><SelectValue placeholder="Selecione" /></SelectTrigger>
                  <SelectContent>
                    <SelectItem value="never">Nunca fumou</SelectItem>
                    <SelectItem value="former">Ex-fumante</SelectItem>
                    <SelectItem value="current">Fumante ativo</SelectItem>
                  </SelectContent>
                </Select>
              </div>
              <div className="space-y-1">
                <Label className="text-xs">Nivel de Atividade Fisica</Label>
                <Select value={physicalActivityLevel} onValueChange={setPhysicalActivityLevel}>
                  <SelectTrigger><SelectValue placeholder="Selecione" /></SelectTrigger>
                  <SelectContent>
                    <SelectItem value="sedentary">Sedentario</SelectItem>
                    <SelectItem value="insufficient">Insuficiente</SelectItem>
                    <SelectItem value="moderate">Moderado</SelectItem>
                    <SelectItem value="active">Ativo</SelectItem>
                  </SelectContent>
                </Select>
              </div>
            </div>

            <div className="flex items-center gap-6">
              <div className="flex items-center gap-2">
                <Switch checked={familyHistory} onCheckedChange={setFamilyHistory} />
                <Label className="text-xs">Historico familiar de doenca cardiovascular</Label>
              </div>
            </div>
          </CardContent>
        </Card>

        {/* Condições */}
        <Card>
          <CardHeader><CardTitle className="text-base">Condicoes Clinicas</CardTitle></CardHeader>
          <CardContent className="space-y-4">
            <div className="flex items-center gap-6 flex-wrap">
              <div className="flex items-center gap-2">
                <Switch checked={cardiovascularDisease} onCheckedChange={setCardiovascularDisease} />
                <Label className="text-xs">Doenca cardiovascular</Label>
              </div>
              <div className="flex items-center gap-2">
                <Switch checked={clinicalAlert} onCheckedChange={setClinicalAlert} />
                <Label className="text-xs">Alerta clinico</Label>
              </div>
            </div>

            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div className="space-y-1">
                <Label className="text-xs">Tipo de Diabetes</Label>
                <Select value={diabetesType} onValueChange={setDiabetesType}>
                  <SelectTrigger><SelectValue placeholder="Nenhum" /></SelectTrigger>
                  <SelectContent>
                    <SelectItem value="none">Nenhum</SelectItem>
                    <SelectItem value="type1">Tipo 1</SelectItem>
                    <SelectItem value="type2">Tipo 2</SelectItem>
                    <SelectItem value="gestational">Gestacional</SelectItem>
                  </SelectContent>
                </Select>
              </div>
            </div>

            <div className="space-y-1">
              <Label className="text-xs">Sintomas</Label>
              <Textarea
                value={symptoms}
                onChange={(e) => setSymptoms(e.target.value)}
                placeholder="Descreva sintomas relevantes (dor toracica, dispneia, claudicacao, etc.)"
                rows={2}
              />
            </div>
          </CardContent>
        </Card>

        <Button type="submit" disabled={createMutation.isPending} className="w-full">
          {createMutation.isPending ? "Calculando ACSM..." : "Gerar Avaliacao ACSM"}
        </Button>
      </form>
    </div>
  );
}
