"use client";

import { useParams } from "next/navigation";
import { format } from "date-fns";
import { ptBR } from "date-fns/locale";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import { useRequireAuth } from "@/lib/use-auth";
import { useRequireSelectedPatient } from "@/lib/use-require-selected-patient";
import { SelectedPatientHeader } from "@/components/patients/SelectedPatientHeader";
import { PageHeader } from "@/components/layout/page-header";
import { usePosturalAssessment } from "@/lib/api/posture-api";

// --- Severity classification ---

type Severity = "Normal" | "Leve" | "Moderado" | "Severo";

function classifyStandard(value: number): Severity {
  if (value <= 2) return "Normal";
  if (value <= 5) return "Leve";
  if (value <= 10) return "Moderado";
  return "Severo";
}

function classifyFhp(value: number): Severity {
  if (value <= 3) return "Normal";
  if (value <= 8) return "Leve";
  if (value <= 15) return "Moderado";
  return "Severo";
}

function classifySpinal(value: number): Severity {
  if (value <= 5) return "Normal";
  if (value <= 15) return "Leve";
  if (value <= 30) return "Moderado";
  return "Severo";
}

const severityColors: Record<Severity, string> = {
  Normal: "text-green-600",
  Leve: "text-yellow-600",
  Moderado: "text-orange-500",
  Severo: "text-red-600",
};

const severityBadgeColors: Record<Severity, string> = {
  Normal: "bg-green-100 text-green-800 border-green-200",
  Leve: "bg-yellow-100 text-yellow-800 border-yellow-200",
  Moderado: "bg-orange-100 text-orange-800 border-orange-200",
  Severo: "bg-red-100 text-red-800 border-red-200",
};

// --- Overall score classification ---

type OverallClassification = "Excelente" | "Boa" | "Regular" | "Necessita Atencao";

const overallColors: Record<OverallClassification, string> = {
  "Excelente": "bg-green-100 text-green-800 border-green-200",
  "Boa": "bg-blue-100 text-blue-800 border-blue-200",
  "Regular": "bg-yellow-100 text-yellow-800 border-yellow-200",
  "Necessita Atencao": "bg-red-100 text-red-800 border-red-200",
};

function classifyOverall(score: number): OverallClassification {
  if (score >= 90) return "Excelente";
  if (score >= 70) return "Boa";
  if (score >= 50) return "Regular";
  return "Necessita Atencao";
}

// --- Measurement config ---

interface MeasurementDef {
  key: string;
  label: string;
  classify: (v: number) => Severity;
}

const measurements: MeasurementDef[] = [
  { key: "shoulderDeviation", label: "Desvio Ombro", classify: classifyStandard },
  { key: "hipDeviation", label: "Desvio Quadril", classify: classifyStandard },
  { key: "headLateralDeviation", label: "Desvio Lateral Cabeca", classify: classifyStandard },
  { key: "fhp", label: "FHP - Anterioridade Cabeca", classify: classifyFhp },
  { key: "thoracicKyphosis", label: "Cifose Toracica", classify: classifySpinal },
  { key: "lumbarLordosis", label: "Lordose Lombar", classify: classifySpinal },
  { key: "kneeAngle", label: "Angulo Joelho", classify: classifyStandard },
];

// --- Metric card ---

function PostureMetricCard({ label, value, classify }: {
  label: string; value?: number | null; classify: (v: number) => Severity;
}) {
  if (value == null) return null;
  const severity = classify(Math.abs(value));
  return (
    <Card>
      <CardContent className="p-4">
        <p className="text-xs text-muted-foreground mb-1">{label}</p>
        <p className={`text-2xl font-bold ${severityColors[severity]}`}>
          {value.toFixed(1)}<span className="text-sm font-normal ml-1">graus</span>
        </p>
        <Badge variant="outline" className={`mt-2 text-xs ${severityBadgeColors[severity]}`}>
          {severity}
        </Badge>
      </CardContent>
    </Card>
  );
}

// --- View type labels ---

const viewTypeLabels: Record<string, string> = {
  front: "Anterior (Frente)",
  side_left: "Lateral Esquerda",
  side_right: "Lateral Direita",
  back: "Posterior (Costas)",
};

export default function PosturalAssessmentDetailPage() {
  useRequireAuth();
  useRequireSelectedPatient();
  const params = useParams();
  const id = params.id as string;

  const { data: a, isLoading } = usePosturalAssessment(id);

  if (isLoading) {
    return (
      <div className="space-y-6">
        <Skeleton className="h-20" />
        <Skeleton className="h-64" />
      </div>
    );
  }

  if (!a) {
    return (
      <div className="text-center py-12">
        <p className="text-muted-foreground">Avaliacao postural nao encontrada.</p>
      </div>
    );
  }

  // Compute overall score and severe count
  const severities: Severity[] = [];
  for (const m of measurements) {
    const val = a[m.key as keyof typeof a] as number | null | undefined;
    if (val != null) {
      severities.push(m.classify(Math.abs(val)));
    }
  }

  const severeCount = severities.filter((s) => s === "Severo").length;

  // Score: each measurement contributes up to ~14.3 points (100/7)
  // Normal=100%, Leve=75%, Moderado=50%, Severo=25%
  const severityScoreMap: Record<Severity, number> = {
    Normal: 100,
    Leve: 75,
    Moderado: 50,
    Severo: 25,
  };

  const overallScore = severities.length > 0
    ? Math.round(severities.reduce((acc, s) => acc + severityScoreMap[s], 0) / severities.length)
    : 0;

  const classification = classifyOverall(overallScore);

  return (
    <div className="space-y-6">
      <SelectedPatientHeader />
      <PageHeader
        title="Avaliacao Postural"
        description={format(new Date(a.assessmentDate), "dd 'de' MMMM 'de' yyyy", { locale: ptBR })}
      />

      {/* View Type Badge */}
      {a.viewType && (
        <Badge variant="outline" className="text-sm">
          Vista: {viewTypeLabels[a.viewType] || a.viewType}
        </Badge>
      )}

      {/* Overall Score Card */}
      <Card className={`border-2 ${overallColors[classification]}`}>
        <CardContent className="p-6 text-center">
          <p className="text-5xl font-bold">{overallScore}</p>
          <h2 className="text-xl font-semibold mt-2">{classification}</h2>
          <p className="text-sm mt-1 text-muted-foreground">
            Pontuacao postural geral ({severities.length} medicoes avaliadas)
          </p>
        </CardContent>
      </Card>

      {/* Severe Deviations Count */}
      {severeCount > 0 && (
        <div className="flex justify-center">
          <Badge variant="destructive" className="text-sm px-4 py-1">
            {severeCount} desvio(s) severo(s) detectado(s)
          </Badge>
        </div>
      )}

      {/* Individual Measurements Grid */}
      <div>
        <h3 className="text-base font-semibold mb-3">Medicoes Individuais</h3>
        <div className="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
          {measurements.map((m) => (
            <PostureMetricCard
              key={m.key}
              label={m.label}
              value={a[m.key as keyof typeof a] as number | null | undefined}
              classify={m.classify}
            />
          ))}
        </div>
      </div>

      {/* Notes */}
      {a.notes && (
        <Card>
          <CardHeader>
            <CardTitle className="text-base">Observacoes</CardTitle>
          </CardHeader>
          <CardContent>
            <p className="text-sm whitespace-pre-wrap">{a.notes}</p>
          </CardContent>
        </Card>
      )}
    </div>
  );
}
