"use client";

/**
 * /patients/[id] — CAPA do prontuário (P1 do plano UX do médico).
 *
 * Abre na SÍNTESE proprietária (Escore Plenya + radar AGIR derivado do snapshot
 * + biomarcadores em Atenção/Positivos), não em demografia. Os dados cadastrais
 * descem para uma seção secundária abaixo. Quando não há Escore, mostra estado
 * vazio com CTA. Frontend-only: reusa o snapshot mais recente (useLatestHealthScore),
 * o RadarAgir e a timeline agregada (useMedicalRecord) — sem backend novo.
 */
import { useMemo } from "react";
import { useParams, useRouter } from "next/navigation";
import { useQuery } from "@tanstack/react-query";
import { format } from "date-fns";
import { ptBR } from "date-fns/locale";
import {
  Activity,
  AlertTriangle,
  CheckCircle2,
  Edit,
  FileText,
  Workflow,
} from "lucide-react";

import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import { apiClient } from "@/lib/api-client";
import { cn } from "@/lib/utils";
import { useRequireAuth } from "@/lib/use-auth";
import { PageHeader } from "@/components/layout/page-header";
import { PortalAccessCard } from "@/components/patients/portal-access-card";
import { CheckInsCard } from "@/components/patients/check-ins-card";
import { PatientFinanceCard } from "@/components/patients/patient-finance-card";
import { RadarAgir, type RadarLetter, type RadarPillar } from "@/components/health-scores/RadarAgir";
import {
  useLatestHealthScore,
  type PatientScoreSnapshot,
  type PatientScoreItemResult,
} from "@/lib/api/health-score-api";
import { useMedicalRecord } from "@/lib/api/medical-record-api";

interface Patient {
  id: string;
  userId: string;
  name: string;
  cpf?: string;
  birthDate: string;
  gender: "male" | "female" | "other";
  menopause?: boolean;
  phone?: string;
  address?: string;
  municipality?: string;
  state?: string;
  motherName?: string;
  fatherName?: string;
  height?: number;
  weight?: number;
  createdAt: string;
  updatedAt: string;
}

// Semântica canônica dos níveis (espelha ScoreLevelBadge): 0=Crítico … 5=Ótimo.
const LEVEL_META: Record<number, { label: string; chip: string }> = {
  0: { label: "Crítico", chip: "border-red-300 bg-red-50 text-red-800" },
  1: { label: "Muito Baixo/Alto", chip: "border-orange-300 bg-orange-50 text-orange-800" },
  2: { label: "Subótimo", chip: "border-yellow-300 bg-yellow-50 text-yellow-800" },
  3: { label: "Limítrofe", chip: "border-blue-300 bg-blue-50 text-blue-800" },
  4: { label: "Bom", chip: "border-green-300 bg-green-50 text-green-800" },
  5: { label: "Ótimo", chip: "border-emerald-300 bg-emerald-50 text-emerald-800" },
  6: { label: "Reservado", chip: "border-gray-300 bg-gray-50 text-gray-800" },
};

// Deriva letras/pilares AGIR do próprio snapshot (itemResults carregam
// item.methodPillars). Retorna null se o snapshot não tiver mapeamento de método.
function buildAgir(snapshot: PatientScoreSnapshot): { letters: RadarLetter[]; pillars: RadarPillar[] } | null {
  const items = (snapshot.itemResults ?? []).filter(
    (ir) => ir.status === "evaluated" && (ir.item?.methodPillars?.length ?? 0) > 0,
  );
  if (items.length === 0) return null;

  const pillarMap = new Map<string, { name: string; letter: string; actual: number; max: number }>();
  const letterMap = new Map<string, { name: string; color: string; order: number; actual: number; max: number }>();

  for (const ir of items) {
    for (const mp of ir.item?.methodPillars ?? []) {
      const letter = mp.letter;
      if (!letter) continue;
      const p = pillarMap.get(mp.id) ?? { name: mp.name, letter: letter.code, actual: 0, max: 0 };
      p.actual += ir.actualPoints;
      p.max += ir.maxPoints;
      pillarMap.set(mp.id, p);

      const l = letterMap.get(letter.code) ?? {
        name: letter.name,
        color: letter.color || "#94a3b8",
        order: letter.order ?? 0,
        actual: 0,
        max: 0,
      };
      l.actual += ir.actualPoints;
      l.max += ir.maxPoints;
      letterMap.set(letter.code, l);
    }
  }
  if (letterMap.size === 0) return null;

  const letters: RadarLetter[] = [...letterMap.entries()]
    .sort((a, b) => a[1].order - b[1].order)
    .map(([code, l]) => ({ code, name: l.name, color: l.color, score: l.max > 0 ? (l.actual / l.max) * 100 : 0 }));
  const pillars: RadarPillar[] = [...pillarMap.values()].map((p) => ({
    letter: p.letter,
    name: p.name,
    score: p.max > 0 ? (p.actual / p.max) * 100 : 0,
  }));
  return { letters, pillars };
}

function fmtValue(ir: PatientScoreItemResult): string {
  if (ir.valueUsed === undefined || ir.valueUsed === null) return "";
  const unit = ir.item?.unit ? ` ${ir.item.unit}` : "";
  return `${ir.valueUsed}${unit}`;
}

export default function PatientDetailPage() {
  useRequireAuth();
  const params = useParams();
  const router = useRouter();
  const patientId = params.id as string;

  const { data: patient, isLoading } = useQuery({
    queryKey: ["patient", patientId],
    queryFn: () => apiClient.get<Patient>(`/api/v1/patients/${patientId}`),
  });

  const { data: score, isLoading: scoreLoading } = useLatestHealthScore(patientId);
  const mr = useMedicalRecord(patientId, { types: ["appointment"] });
  const lastAppointment = useMemo(() => (mr.data?.pages ?? []).flat()[0], [mr.data]);

  const agir = useMemo(() => (score ? buildAgir(score) : null), [score]);

  const { atencao, positivos } = useMemo(() => {
    const evaluated = (score?.itemResults ?? []).filter(
      (ir) => ir.status === "evaluated" && ir.levelNumber !== undefined && ir.levelNumber !== null,
    );
    const atencao = evaluated
      .filter((ir) => (ir.levelNumber ?? 9) <= 2)
      .sort((a, b) => (a.levelNumber ?? 0) - (b.levelNumber ?? 0))
      .slice(0, 8);
    const positivos = evaluated
      .filter((ir) => (ir.levelNumber ?? -1) >= 4 && (ir.levelNumber ?? 9) <= 5)
      .sort((a, b) => (b.levelNumber ?? 0) - (a.levelNumber ?? 0))
      .slice(0, 8);
    return { atencao, positivos };
  }, [score]);

  if (isLoading) {
    return (
      <div className="container mx-auto space-y-4 py-8">
        <Skeleton className="h-10 w-64" />
        <Skeleton className="h-96 w-full" />
      </div>
    );
  }
  if (!patient) {
    return (
      <div className="container mx-auto py-8">
        <p>Paciente não encontrado</p>
      </div>
    );
  }

  const age = patient.birthDate
    ? Math.floor((Date.now() - new Date(patient.birthDate).getTime()) / (1000 * 60 * 60 * 24 * 365.25))
    : null;
  const genderLabel = { male: "Masculino", female: "Feminino", other: "Outro" }[patient.gender];
  const formatCPF = (cpf: string) => {
    const c = cpf.replace(/\D/g, "");
    return c.length === 11 ? `${c.slice(0, 3)}.${c.slice(3, 6)}.${c.slice(6, 9)}-${c.slice(9)}` : cpf;
  };

  return (
    <div className="container mx-auto space-y-6 py-8">
      <PageHeader
        title={patient.name}
        description={[genderLabel, age !== null ? `${age} anos` : null].filter(Boolean).join(" · ")}
        breadcrumbs={[{ label: "Pacientes", href: "/patients" }, { label: patient.name }]}
        actions={[
          {
            label: "Prontuário",
            icon: <FileText className="h-4 w-4" />,
            variant: "outline",
            onClick: () => router.push(`/patients/${patientId}/prontuario`),
          },
          {
            label: "Continuum",
            icon: <Workflow className="h-4 w-4" />,
            variant: "outline",
            onClick: () => router.push(`/patients/${patientId}/continuum`),
          },
          {
            label: "Editar",
            icon: <Edit className="h-4 w-4" />,
            variant: "default",
            onClick: () => router.push(`/patients/${patientId}/edit`),
          },
        ]}
      />

      {/* ===== SÍNTESE ===== */}
      {scoreLoading ? (
        <Skeleton className="h-80 w-full" />
      ) : score ? (
        <>
          <div className="grid gap-6 lg:grid-cols-3">
            {/* Radar / Escore */}
            <Card className="lg:col-span-2">
              <CardHeader className="pb-0">
                <CardTitle className="flex items-center gap-2 text-base">
                  <Activity className="h-5 w-5 text-primary" />
                  Escore Plenya
                  <Badge variant="outline" className="ml-1">
                    {score.totalScorePercentage.toFixed(0)}%
                  </Badge>
                </CardTitle>
              </CardHeader>
              <CardContent className="pt-4">
                {agir ? (
                  <RadarAgir letters={agir.letters} pillars={agir.pillars} globalScore={score.totalScorePercentage} />
                ) : (
                  // Sem mapeamento de método: barras por grupo clínico.
                  <div className="space-y-2">
                    {(score.groupResults ?? [])
                      .slice()
                      .sort((a, b) => (a.group?.order ?? 0) - (b.group?.order ?? 0))
                      .map((g) => (
                        <div key={g.id} className="space-y-1">
                          <div className="flex items-center justify-between text-sm">
                            <span>{g.group?.name ?? "Grupo"}</span>
                            <span className="text-muted-foreground">{g.scorePercentage.toFixed(0)}%</span>
                          </div>
                          <div className="h-2 w-full overflow-hidden rounded-full bg-muted">
                            <div className="h-full rounded-full bg-primary" style={{ width: `${Math.min(100, Math.max(0, g.scorePercentage))}%` }} />
                          </div>
                        </div>
                      ))}
                  </div>
                )}
              </CardContent>
            </Card>

            {/* Resumo */}
            <Card>
              <CardHeader className="pb-2">
                <CardTitle className="text-base">Resumo</CardTitle>
              </CardHeader>
              <CardContent className="space-y-3 text-sm">
                <div className="flex items-center justify-between">
                  <span className="text-muted-foreground">Itens avaliados</span>
                  <span className="font-medium">{score.itemsEvaluatedCount}</span>
                </div>
                <div className="flex items-center justify-between">
                  <span className="text-muted-foreground">Escore calculado em</span>
                  <span className="font-medium">{format(new Date(score.calculatedAt), "dd/MM/yyyy")}</span>
                </div>
                <div className="border-t pt-3">
                  <p className="text-xs uppercase text-muted-foreground">Última consulta</p>
                  {lastAppointment ? (
                    <button
                      type="button"
                      onClick={() => router.push(`/appointments/${lastAppointment.id}`)}
                      className="mt-1 flex w-full items-center justify-between gap-2 text-left hover:underline"
                    >
                      <span>{format(new Date(lastAppointment.eventDate), "dd/MM/yyyy")} — {lastAppointment.title}</span>
                      <Badge
                        variant="outline"
                        className={cn(
                          "shrink-0 text-[10px]",
                          lastAppointment.hasNote
                            ? "border-emerald-200 bg-emerald-50 text-emerald-700"
                            : "border-amber-200 bg-amber-50 text-amber-700",
                        )}
                      >
                        {lastAppointment.hasNote ? "Nota registrada" : "Sem nota"}
                      </Badge>
                    </button>
                  ) : (
                    <p className="mt-1 text-muted-foreground">Sem consultas registradas.</p>
                  )}
                </div>
                <Button variant="link" className="h-auto p-0 text-xs" onClick={() => router.push(`/patients/${patientId}/prontuario`)}>
                  Ver linha do tempo completa →
                </Button>
              </CardContent>
            </Card>
          </div>

          {/* Atenção / Positivos */}
          <div className="grid gap-6 md:grid-cols-2">
            <BiomarkerCard
              title="Atenção"
              icon={<AlertTriangle className="h-5 w-5 text-amber-600" />}
              items={atencao}
              emptyMsg="Nenhum biomarcador em faixa de atenção."
            />
            <BiomarkerCard
              title="No ótimo"
              icon={<CheckCircle2 className="h-5 w-5 text-emerald-600" />}
              items={positivos}
              emptyMsg="Nenhum biomarcador em faixa ótima ainda."
            />
          </div>
        </>
      ) : (
        <Card>
          <CardContent className="flex flex-col items-center gap-3 py-12 text-center">
            <Activity className="h-10 w-10 text-muted-foreground" />
            <div>
              <p className="font-medium">Este paciente ainda não tem Escore calculado.</p>
              <p className="text-sm text-muted-foreground">
                Calcule o Escore Plenya para abrir o prontuário na síntese de saúde.
              </p>
            </div>
            <Button onClick={() => router.push("/health-scores")}>
              <Activity className="mr-2 h-4 w-4" />
              Calcular Escore
            </Button>
          </CardContent>
        </Card>
      )}

      {/* ===== DADOS CADASTRAIS (secundário) ===== */}
      <div className="pt-2">
        <h2 className="mb-3 text-sm font-medium uppercase tracking-wide text-muted-foreground">
          Dados cadastrais
        </h2>
        <div className="grid gap-6 md:grid-cols-2">
          <Card>
            <CardHeader>
              <CardTitle className="text-base">Informações Pessoais</CardTitle>
            </CardHeader>
            <CardContent className="space-y-3 text-sm">
              {patient.cpf && <Row label="CPF" value={formatCPF(patient.cpf)} mono />}
              {patient.birthDate && (
                <Row label="Nascimento" value={format(new Date(patient.birthDate), "dd/MM/yyyy", { locale: ptBR })} />
              )}
              <Row label="Gênero" value={genderLabel} />
              {patient.phone && <Row label="Telefone" value={patient.phone} />}
              {patient.address && <Row label="Endereço" value={patient.address} />}
              {(patient.municipality || patient.state) && (
                <Row label="Localização" value={[patient.municipality, patient.state].filter(Boolean).join(" - ")} />
              )}
            </CardContent>
          </Card>

          <Card>
            <CardHeader>
              <CardTitle className="text-base">Medidas</CardTitle>
            </CardHeader>
            <CardContent className="space-y-3 text-sm">
              {patient.height ? <Row label="Altura" value={`${patient.height} cm`} /> : null}
              {patient.weight ? <Row label="Peso" value={`${patient.weight} kg`} /> : null}
              {patient.height && patient.weight && (
                <Row label="IMC" value={(patient.weight / Math.pow(patient.height / 100, 2)).toFixed(1)} />
              )}
              {!patient.height && !patient.weight && <p className="text-muted-foreground">Sem medidas registradas.</p>}
            </CardContent>
          </Card>

          <CheckInsCard patientId={patientId} />
          <PortalAccessCard patientId={patientId} />
          <PatientFinanceCard patientId={patientId} />
        </div>
      </div>
    </div>
  );
}

function Row({ label, value, mono }: { label: string; value: string; mono?: boolean }) {
  return (
    <div className="flex items-center justify-between gap-3">
      <span className="text-muted-foreground">{label}</span>
      <span className={cn("font-medium", mono && "font-mono")}>{value}</span>
    </div>
  );
}

function BiomarkerCard({
  title,
  icon,
  items,
  emptyMsg,
}: {
  title: string;
  icon: React.ReactNode;
  items: PatientScoreItemResult[];
  emptyMsg: string;
}) {
  return (
    <Card>
      <CardHeader className="pb-2">
        <CardTitle className="flex items-center gap-2 text-base">
          {icon}
          {title}
        </CardTitle>
      </CardHeader>
      <CardContent>
        {items.length === 0 ? (
          <p className="py-4 text-sm text-muted-foreground">{emptyMsg}</p>
        ) : (
          <div className="space-y-2">
            {items.map((ir) => {
              const meta = LEVEL_META[ir.levelNumber ?? 6] ?? LEVEL_META[6];
              return (
                <div key={ir.id} className="flex items-center justify-between gap-2 text-sm">
                  <span className="min-w-0 flex-1 truncate">{ir.item?.name ?? "Item"}</span>
                  {fmtValue(ir) && <span className="shrink-0 font-mono text-xs text-muted-foreground">{fmtValue(ir)}</span>}
                  <Badge variant="outline" className={cn("shrink-0 text-[10px]", meta.chip)}>
                    N{ir.levelNumber}: {ir.levelMatched?.name ?? meta.label}
                  </Badge>
                </div>
              );
            })}
          </div>
        )}
      </CardContent>
    </Card>
  );
}
