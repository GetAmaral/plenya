"use client";

import { useParams } from "next/navigation";
import { format } from "date-fns";
import { ptBR } from "date-fns/locale";
import { Heart, Scale, Ruler, Droplets } from "lucide-react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import { useRequireAuth } from "@/lib/use-auth";
import { useRequireSelectedPatient } from "@/lib/use-require-selected-patient";
import { SelectedPatientHeader } from "@/components/patients/SelectedPatientHeader";
import { PageHeader } from "@/components/layout/page-header";
import { usePhysicalAssessment } from "@/lib/api/physical-assessment-api";

const riskColors: Record<string, string> = {
  low: "bg-green-100 text-green-800 border-green-200",
  moderate: "bg-yellow-100 text-yellow-800 border-yellow-200",
  high: "bg-red-100 text-red-800 border-red-200",
};

const riskLabels: Record<string, string> = {
  low: "Baixo Risco",
  moderate: "Risco Moderado",
  high: "Alto Risco",
};

function MetricCard({ label, value, unit, icon: Icon }: { label: string; value?: number | null; unit?: string; icon?: any }) {
  if (value == null) return null;
  return (
    <Card>
      <CardContent className="p-4 flex items-center gap-3">
        {Icon && <Icon className="h-5 w-5 text-muted-foreground" />}
        <div>
          <p className="text-xs text-muted-foreground">{label}</p>
          <p className="text-lg font-semibold">
            {typeof value === "number" ? value.toFixed(1) : value}
            {unit && <span className="text-sm font-normal ml-1">{unit}</span>}
          </p>
        </div>
      </CardContent>
    </Card>
  );
}

const smokingLabels: Record<string, string> = {
  never: "Nunca fumou",
  former: "Ex-fumante",
  current: "Fumante ativo",
};

const activityLabels: Record<string, string> = {
  sedentary: "Sedentario",
  insufficient: "Insuficiente",
  moderate: "Moderado",
  active: "Ativo",
};

export default function PhysicalAssessmentDetailPage() {
  useRequireAuth();
  useRequireSelectedPatient();
  const params = useParams();
  const id = params.id as string;

  const { data: a, isLoading } = usePhysicalAssessment(id);

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
        <p className="text-muted-foreground">Avaliacao nao encontrada.</p>
      </div>
    );
  }

  return (
    <div className="space-y-6">
      <SelectedPatientHeader />
      <PageHeader
        title="Avaliacao Fisica"
        description={format(new Date(a.assessmentDate), "dd 'de' MMMM 'de' yyyy", { locale: ptBR })}
      />

      {/* ACSM Risk Badge */}
      {a.acsmRiskLevel && (
        <Card className={`border-2 ${riskColors[a.acsmRiskLevel] || ""}`}>
          <CardContent className="p-6 text-center">
            <h2 className="text-2xl font-bold">
              {riskLabels[a.acsmRiskLevel] || a.acsmRiskLevel}
            </h2>
            <p className="text-sm mt-1">
              {a.acsmRiskFactorsCount} fator(es) de risco cardiovascular
            </p>
          </CardContent>
        </Card>
      )}

      {/* Risk Factors */}
      {a.acsmRiskFactors?.length > 0 && (
        <Card>
          <CardHeader>
            <CardTitle className="text-base">Fatores de Risco</CardTitle>
          </CardHeader>
          <CardContent>
            <div className="flex flex-wrap gap-2">
              {a.acsmRiskFactors.map((f) => (
                <Badge key={f} variant="destructive">{f}</Badge>
              ))}
            </div>
            {a.acsmProtectiveFactors?.length > 0 && (
              <div className="mt-3">
                <p className="text-sm text-muted-foreground mb-1">Fatores Protetores:</p>
                <div className="flex flex-wrap gap-2">
                  {a.acsmProtectiveFactors.map((f) => (
                    <Badge key={f} variant="secondary" className="bg-green-100 text-green-800">{f}</Badge>
                  ))}
                </div>
              </div>
            )}
          </CardContent>
        </Card>
      )}

      {/* Metrics Grid */}
      <div className="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
        <MetricCard label="Peso" value={a.weight} unit="kg" icon={Scale} />
        <MetricCard label="Altura" value={a.height} unit="cm" icon={Ruler} />
        <MetricCard label="IMC" value={a.bmi} icon={Scale} />
        <MetricCard label="BRI" value={a.bri} />
        <MetricCard label="Cintura" value={a.waistCircumference} unit="cm" />
        <MetricCard label="% Gordura" value={a.bodyFatPercent} unit="%" />
        <MetricCard label="Massa Magra" value={a.leanMass} unit="kg" />
        <MetricCard label="PA" value={a.systolicBp != null && a.diastolicBp != null ? undefined : undefined} />
        <MetricCard label="LDL" value={a.ldl} unit="mg/dL" icon={Droplets} />
        <MetricCard label="HDL" value={a.hdl} unit="mg/dL" icon={Droplets} />
        <MetricCard label="Colesterol Total" value={a.totalCholesterol} unit="mg/dL" />
        <MetricCard label="Triglicerideos" value={a.triglycerides} unit="mg/dL" />
        <MetricCard label="Glicemia Jejum" value={a.fastingGlucose} unit="mg/dL" />
        <MetricCard label="HbA1c" value={a.hbA1c} unit="%" />
        <MetricCard label="FC Maxima" value={a.maxHr} unit="bpm" icon={Heart} />
        <MetricCard label="FC Repouso" value={a.restingHeartRate} unit="bpm" icon={Heart} />
      </div>

      {/* Blood Pressure (special display) */}
      {a.systolicBp != null && a.diastolicBp != null && (
        <Card>
          <CardContent className="p-4 flex items-center gap-3">
            <Heart className="h-5 w-5 text-muted-foreground" />
            <div>
              <p className="text-xs text-muted-foreground">Pressao Arterial</p>
              <p className="text-lg font-semibold">{a.systolicBp}/{a.diastolicBp} <span className="text-sm font-normal">mmHg</span></p>
            </div>
          </CardContent>
        </Card>
      )}

      {/* History Info */}
      {(a.smokingStatus || a.physicalActivityLevel || a.familyHistory) && (
        <Card>
          <CardHeader>
            <CardTitle className="text-base">Historico</CardTitle>
          </CardHeader>
          <CardContent>
            <div className="flex flex-wrap gap-2">
              {a.smokingStatus && (
                <Badge variant="outline">Tabagismo: {smokingLabels[a.smokingStatus] || a.smokingStatus}</Badge>
              )}
              {a.physicalActivityLevel && (
                <Badge variant="outline">Atividade: {activityLabels[a.physicalActivityLevel] || a.physicalActivityLevel}</Badge>
              )}
              {a.familyHistory && (
                <Badge variant="destructive">Hist. Familiar CV</Badge>
              )}
              {a.cardiovascularDisease && (
                <Badge variant="destructive">Doenca Cardiovascular</Badge>
              )}
            </div>
          </CardContent>
        </Card>
      )}

      {/* HR Zones */}
      {a.hrZones && a.hrZones.length > 0 && (
        <Card>
          <CardHeader>
            <CardTitle className="text-base">Zonas de Frequencia Cardiaca (Karvonen)</CardTitle>
          </CardHeader>
          <CardContent>
            <div className="space-y-2">
              {a.hrZones.map((zone) => (
                <div key={zone.zone} className="flex items-center gap-3 p-2 rounded-lg bg-muted/50">
                  <Badge variant="outline" className="w-16 justify-center">Z{zone.zone}</Badge>
                  <span className="text-sm font-medium flex-1">{zone.name}</span>
                  <span className="text-sm text-muted-foreground">{zone.percent}</span>
                  <span className="text-sm font-mono">
                    {zone.minBpm}-{zone.maxBpm} bpm
                  </span>
                </div>
              ))}
            </div>
          </CardContent>
        </Card>
      )}

      {/* ACSM Tags */}
      {a.acsmTags?.length > 0 && (
        <Card>
          <CardHeader>
            <CardTitle className="text-base">Tags ACSM</CardTitle>
          </CardHeader>
          <CardContent>
            <div className="flex flex-wrap gap-2">
              {a.acsmTags.map((tag) => (
                <Badge key={tag} variant="outline">{tag}</Badge>
              ))}
            </div>
          </CardContent>
        </Card>
      )}

      {/* Recommendation */}
      {a.acsmRecommendation && (
        <Card>
          <CardHeader>
            <CardTitle className="text-base">Recomendacao ACSM</CardTitle>
          </CardHeader>
          <CardContent>
            <p className="text-sm">{a.acsmRecommendation}</p>
          </CardContent>
        </Card>
      )}

      {/* AI Recommendation */}
      {a.aiRecommendation && (
        <Card>
          <CardHeader>
            <CardTitle className="text-base">Recomendacao IA</CardTitle>
          </CardHeader>
          <CardContent>
            <p className="text-sm whitespace-pre-wrap">{a.aiRecommendation}</p>
          </CardContent>
        </Card>
      )}
    </div>
  );
}
