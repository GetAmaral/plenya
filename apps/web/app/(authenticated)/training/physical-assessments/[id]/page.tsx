"use client";

import { useParams } from "next/navigation";
import Link from "next/link";
import { format } from "date-fns";
import { ptBR } from "date-fns/locale";
import { Activity, ArrowLeft, Heart, Scale, Ruler, Droplets } from "lucide-react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
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

export default function PhysicalAssessmentDetailPage() {
  useRequireAuth();
  useRequireSelectedPatient();
  const params = useParams();
  const id = params.id as string;

  const { data: assessment, isLoading } = usePhysicalAssessment(id);

  if (isLoading) {
    return (
      <div className="space-y-6">
        <Skeleton className="h-20" />
        <Skeleton className="h-64" />
      </div>
    );
  }

  if (!assessment) {
    return (
      <div className="text-center py-12">
        <p className="text-muted-foreground">Avaliação não encontrada.</p>
      </div>
    );
  }

  const resolved = assessment.resolvedData;

  return (
    <div className="space-y-6">
      <SelectedPatientHeader />
      <PageHeader
        title="Avaliação Física"
        subtitle={format(new Date(assessment.assessmentDate), "dd 'de' MMMM 'de' yyyy", { locale: ptBR })}
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

      {/* ACSM Risk Badge */}
      {assessment.acsmRiskLevel && (
        <Card className={`border-2 ${riskColors[assessment.acsmRiskLevel] || ""}`}>
          <CardContent className="p-6 text-center">
            <h2 className="text-2xl font-bold">
              {riskLabels[assessment.acsmRiskLevel] || assessment.acsmRiskLevel}
            </h2>
            <p className="text-sm mt-1">
              {assessment.acsmRiskFactorsCount} fator(es) de risco cardiovascular
            </p>
          </CardContent>
        </Card>
      )}

      {/* Risk Factors */}
      {assessment.acsmRiskFactors?.length > 0 && (
        <Card>
          <CardHeader>
            <CardTitle className="text-base">Fatores de Risco</CardTitle>
          </CardHeader>
          <CardContent>
            <div className="flex flex-wrap gap-2">
              {assessment.acsmRiskFactors.map((f) => (
                <Badge key={f} variant="destructive">{f}</Badge>
              ))}
            </div>
            {assessment.acsmProtectiveFactors?.length > 0 && (
              <div className="mt-3">
                <p className="text-sm text-muted-foreground mb-1">Fatores Protetores:</p>
                <div className="flex flex-wrap gap-2">
                  {assessment.acsmProtectiveFactors.map((f) => (
                    <Badge key={f} variant="secondary" className="bg-green-100 text-green-800">{f}</Badge>
                  ))}
                </div>
              </div>
            )}
          </CardContent>
        </Card>
      )}

      {/* Metrics Grid */}
      {resolved && (
        <div className="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
          <MetricCard label="Peso" value={resolved.weight} unit="kg" icon={Scale} />
          <MetricCard label="Altura" value={resolved.height} unit="cm" icon={Ruler} />
          <MetricCard label="IMC" value={resolved.bmi} icon={Scale} />
          <MetricCard label="BRI" value={resolved.bri} />
          <MetricCard label="Cintura" value={resolved.waistCm} unit="cm" />
          <MetricCard label="% Gordura" value={resolved.bodyFatPercent} unit="%" />
          <MetricCard label="LDL" value={resolved.ldl} unit="mg/dL" icon={Droplets} />
          <MetricCard label="HDL" value={resolved.hdl} unit="mg/dL" icon={Droplets} />
          <MetricCard label="Colesterol Total" value={resolved.totalChol} unit="mg/dL" />
          <MetricCard label="Triglicerídeos" value={resolved.triglycerides} unit="mg/dL" />
          <MetricCard label="Glicemia Jejum" value={resolved.fastingGlucose} unit="mg/dL" />
          <MetricCard label="HbA1c" value={resolved.hbA1c} unit="%" />
          <MetricCard label="FC Máxima" value={resolved.maxHr} unit="bpm" icon={Heart} />
        </div>
      )}

      {/* HR Zones */}
      {resolved?.hrZones && resolved.hrZones.length > 0 && (
        <Card>
          <CardHeader>
            <CardTitle className="text-base">Zonas de Frequência Cardíaca (Karvonen)</CardTitle>
          </CardHeader>
          <CardContent>
            <div className="space-y-2">
              {resolved.hrZones.map((zone) => (
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
      {assessment.acsmTags?.length > 0 && (
        <Card>
          <CardHeader>
            <CardTitle className="text-base">Tags ACSM</CardTitle>
          </CardHeader>
          <CardContent>
            <div className="flex flex-wrap gap-2">
              {assessment.acsmTags.map((tag) => (
                <Badge key={tag} variant="outline">{tag}</Badge>
              ))}
            </div>
          </CardContent>
        </Card>
      )}

      {/* Recommendation */}
      {assessment.acsmRecommendation && (
        <Card>
          <CardHeader>
            <CardTitle className="text-base">Recomendação ACSM</CardTitle>
          </CardHeader>
          <CardContent>
            <p className="text-sm">{assessment.acsmRecommendation}</p>
          </CardContent>
        </Card>
      )}

      {/* AI Recommendation */}
      {assessment.aiRecommendation && (
        <Card>
          <CardHeader>
            <CardTitle className="text-base">Recomendação IA</CardTitle>
          </CardHeader>
          <CardContent>
            <p className="text-sm whitespace-pre-wrap">{assessment.aiRecommendation}</p>
          </CardContent>
        </Card>
      )}
    </div>
  );
}
