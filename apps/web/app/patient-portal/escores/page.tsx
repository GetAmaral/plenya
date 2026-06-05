"use client";

/**
 * /escores — histórico unificado Light + Completo + radar do mais recente.
 */
import { useMemo } from "react";
import { format } from "date-fns";
import { ptBR } from "date-fns/locale";
import { Loader2, Activity, ArrowUp, ArrowDown } from "lucide-react";

import { useRequirePatientAuth } from "@/lib/use-patient-auth";
import { useMyScores, useMyLatestSnapshot, type ScoreEntryView } from "@/lib/api/patient-portal-api";
import { RadarAgir } from "@/components/health-scores/RadarAgir";
import { buildAgir } from "@/components/health-scores/build-agir";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";

const SITE_URL = process.env.NEXT_PUBLIC_SITE_URL || "https://plenyasaude.com.br";

export default function MyScoresPage() {
  const { ready } = useRequirePatientAuth();
  const { data, isLoading } = useMyScores();
  const { data: latestSnapshot } = useMyLatestSnapshot();

  // Radar AGIR do escore completo mais recente — MESMA fonte (buildAgir) do EMR.
  const agir = useMemo(
    () => (latestSnapshot ? buildAgir(latestSnapshot) : null),
    [latestSnapshot],
  );

  if (!ready || isLoading) {
    return (
      <div className="flex h-[60vh] items-center justify-center">
        <Loader2 className="h-6 w-6 animate-spin text-muted-foreground" />
      </div>
    );
  }

  return (
    <div className="mx-auto max-w-3xl space-y-6">
      <header className="space-y-2">
        <p className="text-sm uppercase tracking-wider text-muted-foreground">Saúde</p>
        <h1 className="text-3xl font-light">Meus escores</h1>
        <p className="text-muted-foreground">
          Acompanhe sua evolução. Refazer a avaliação a cada 3 meses ajuda a ver mudanças.
        </p>
      </header>

      {!data?.length ? (
        <Card>
          <CardContent className="space-y-4 py-10 text-center text-sm text-muted-foreground">
            <Activity className="mx-auto h-8 w-8" />
            <p>Você ainda não tem nenhum escore registrado.</p>
            <a
              href={`${SITE_URL}/pt/escore-plenya/avaliar`}
              className="inline-block rounded-md bg-primary px-5 py-2.5 text-sm font-medium text-primary-foreground hover:opacity-90"
            >
              Fazer minha primeira avaliação
            </a>
          </CardContent>
        </Card>
      ) : (
        <>
          {/* Radar AGIR do escore completo mais recente — mesmo radar do EMR */}
          {agir && latestSnapshot && (
            <Card>
              <CardHeader>
                <CardTitle className="text-base">
                  Radar AGIR · {format(new Date(latestSnapshot.calculatedAt), "d 'de' MMM 'de' yyyy", { locale: ptBR })}
                </CardTitle>
              </CardHeader>
              <CardContent className="flex justify-center py-2">
                <RadarAgir
                  letters={agir.letters}
                  pillars={agir.pillars}
                  globalScore={latestSnapshot.totalScorePercentage}
                />
              </CardContent>
            </Card>
          )}

          {/* Timeline */}
          <section className="space-y-3">
            <h2 className="text-lg font-medium">Histórico</h2>
            {data.map((entry, i) => (
              <ScoreEntryRow
                key={entry.id}
                entry={entry}
                previous={i + 1 < data.length ? data[i + 1] : undefined}
              />
            ))}
          </section>
        </>
      )}
    </div>
  );
}

function ScoreEntryRow({ entry, previous }: { entry: ScoreEntryView; previous?: ScoreEntryView }) {
  const total = Math.round(entry.totalScorePercentage);
  const delta = previous ? entry.totalScorePercentage - previous.totalScorePercentage : null;
  const isLight = entry.source === "light";

  const lightHref = entry.publicCode
    ? `${SITE_URL}/pt/escore-plenya/resultado/${entry.publicCode}`
    : undefined;

  const Wrapper = ({ children }: { children: React.ReactNode }) =>
    lightHref ? (
      <a href={lightHref} target="_blank" rel="noreferrer" className="block">
        {children}
      </a>
    ) : (
      <div>{children}</div>
    );

  return (
    <Wrapper>
      <Card className="transition-colors hover:border-primary/40">
        <CardContent className="flex items-center justify-between gap-3 py-4">
          <div className="space-y-1">
            <p className="text-xs text-muted-foreground">
              {format(new Date(entry.createdAt), "EEEE, d 'de' MMMM 'de' yyyy", { locale: ptBR })}
            </p>
            <div className="flex items-center gap-2">
              <Badge variant={isLight ? "secondary" : "default"}>
                {isLight ? "Autoavaliação Light" : "Avaliação completa"}
              </Badge>
              <span className="text-xs text-muted-foreground">
                {entry.itemsEvaluatedCount} itens
              </span>
            </div>
          </div>
          <div className="flex items-baseline gap-2 text-right">
            <span className="text-3xl font-light tabular-nums">{total}</span>
            <span className="text-muted-foreground">/100</span>
            {delta !== null && (
              <span
                className={`flex items-center gap-0.5 text-xs ${
                  delta >= 0 ? "text-emerald-600" : "text-amber-600"
                }`}
              >
                {delta >= 0 ? (
                  <ArrowUp className="h-3 w-3" />
                ) : (
                  <ArrowDown className="h-3 w-3" />
                )}
                {Math.abs(Math.round(delta))}
              </span>
            )}
          </div>
        </CardContent>
      </Card>
    </Wrapper>
  );
}
