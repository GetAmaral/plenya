"use client";

/**
 * CheckInsCard — leitura dos check-ins de bem-estar do paciente (app paciente).
 *
 * Mostra timeline dos últimos 30 e mini-gráfico SVG de energia/dor/humor.
 * Read-only — paciente preenche pelo app, profissional só consome.
 */
import { useQuery } from "@tanstack/react-query";
import { formatDate } from "@/lib/format-date";
import { Activity, Heart, Moon, Brain, Bed } from "lucide-react";

import { apiClient } from "@/lib/api-client";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Skeleton } from "@/components/ui/skeleton";
import { Badge } from "@/components/ui/badge";

interface CheckIn {
  id: string;
  energy: number;
  pain: number;
  painLocation?: string | null;
  mood: number;
  sleepHours: number;
  sleepQuality: number;
  stress: number;
  notes?: string | null;
  createdAt: string;
}

const W = 320;
const H = 80;
const PAD = 4;

function MiniChart({
  data,
  color,
  max = 5,
}: {
  data: number[];
  color: string;
  max?: number;
}) {
  if (data.length === 0) return null;
  const points = data
    .map((v, i) => {
      const x = PAD + (i / Math.max(data.length - 1, 1)) * (W - PAD * 2);
      const y = PAD + (1 - v / max) * (H - PAD * 2);
      return `${x.toFixed(1)},${y.toFixed(1)}`;
    })
    .join(" ");
  return (
    <svg
      viewBox={`0 0 ${W} ${H}`}
      className="h-20 w-full"
      preserveAspectRatio="none"
    >
      <polyline
        points={points}
        fill="none"
        stroke={color}
        strokeWidth="2"
        strokeLinecap="round"
        strokeLinejoin="round"
      />
    </svg>
  );
}

function Score({ value, max = 5 }: { value: number; max?: number }) {
  return (
    <span className="font-semibold">
      {value}
      <span className="text-muted-foreground">/{max}</span>
    </span>
  );
}

export function CheckInsCard({ patientId }: { patientId: string }) {
  const { data, isLoading } = useQuery({
    queryKey: ["patient", patientId, "check-ins"],
    queryFn: async () =>
      await apiClient.get<CheckIn[]>(
        `/api/v1/patients/${patientId}/check-ins?limit=30`,
      ),
  });

  if (isLoading) {
    return (
      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2 text-base">
            <Activity className="h-4 w-4" />
            Check-ins de bem-estar
          </CardTitle>
        </CardHeader>
        <CardContent>
          <Skeleton className="h-32 w-full" />
        </CardContent>
      </Card>
    );
  }

  const checkIns = data ?? [];

  if (checkIns.length === 0) {
    return (
      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2 text-base">
            <Activity className="h-4 w-4" />
            Check-ins de bem-estar
          </CardTitle>
          <CardDescription>
            Paciente ainda não fez check-in pelo app.
          </CardDescription>
        </CardHeader>
      </Card>
    );
  }

  // Pega ordem cronológica (mais antigo → mais recente) pro chart
  const ordered = [...checkIns].reverse();
  const energy = ordered.map((c) => c.energy);
  const pain = ordered.map((c) => c.pain);
  const mood = ordered.map((c) => c.mood);

  const latest = checkIns[0];

  return (
    <Card>
      <CardHeader>
        <CardTitle className="flex items-center gap-2 text-base">
          <Activity className="h-4 w-4" />
          Check-ins de bem-estar
        </CardTitle>
        <CardDescription>
          Últimos {checkIns.length} registros do paciente.
        </CardDescription>
      </CardHeader>
      <CardContent className="space-y-4">
        <div className="grid grid-cols-2 gap-3 sm:grid-cols-4">
          <div className="rounded-lg border p-3">
            <div className="flex items-center gap-1 text-xs text-muted-foreground">
              <Activity className="h-3 w-3" /> Energia
            </div>
            <Score value={latest.energy} />
          </div>
          <div className="rounded-lg border p-3">
            <div className="flex items-center gap-1 text-xs text-muted-foreground">
              <Heart className="h-3 w-3" /> Humor
            </div>
            <Score value={latest.mood} />
          </div>
          <div className="rounded-lg border p-3">
            <div className="flex items-center gap-1 text-xs text-muted-foreground">
              <Bed className="h-3 w-3" /> Sono
            </div>
            <span className="font-semibold">
              {latest.sleepHours}h{" "}
              <span className="text-muted-foreground text-xs">
                ({latest.sleepQuality}/5)
              </span>
            </span>
          </div>
          <div className="rounded-lg border p-3">
            <div className="flex items-center gap-1 text-xs text-muted-foreground">
              <Brain className="h-3 w-3" /> Estresse
            </div>
            <Score value={latest.stress} />
          </div>
        </div>

        <div className="space-y-2">
          <div>
            <div className="mb-1 flex items-center justify-between text-xs text-muted-foreground">
              <span>Energia</span>
              <span>1–5</span>
            </div>
            <MiniChart data={energy} color="#0e7c86" />
          </div>
          <div>
            <div className="mb-1 flex items-center justify-between text-xs text-muted-foreground">
              <span>Dor</span>
              <span>0–5</span>
            </div>
            <MiniChart data={pain} color="#dc2626" />
          </div>
          <div>
            <div className="mb-1 flex items-center justify-between text-xs text-muted-foreground">
              <span>Humor</span>
              <span>1–5</span>
            </div>
            <MiniChart data={mood} color="#7c3aed" />
          </div>
        </div>

        <div className="space-y-2 border-t pt-3">
          <p className="text-xs font-semibold uppercase text-muted-foreground">
            Últimos registros
          </p>
          <div className="max-h-64 space-y-2 overflow-y-auto">
            {checkIns.slice(0, 10).map((c) => (
              <div
                key={c.id}
                className="rounded-md border bg-muted/30 p-2 text-sm"
              >
                <div className="flex items-center justify-between">
                  <span className="text-xs text-muted-foreground">
                    {formatDate(c.createdAt, "dd/MM HH:mm")}
                  </span>
                  {c.pain > 0 && (
                    <Badge variant="destructive" className="text-xs">
                      Dor {c.pain}/5
                      {c.painLocation ? ` · ${c.painLocation}` : ""}
                    </Badge>
                  )}
                </div>
                <div className="mt-1 flex flex-wrap gap-x-4 gap-y-1 text-xs">
                  <span>Energia {c.energy}</span>
                  <span>Humor {c.mood}</span>
                  <span>Sono {c.sleepHours}h ({c.sleepQuality})</span>
                  <span>Estresse {c.stress}</span>
                </div>
                {c.notes && (
                  <p className="mt-1 text-xs italic text-muted-foreground">
                    “{c.notes}”
                  </p>
                )}
              </div>
            ))}
          </div>
        </div>
      </CardContent>
    </Card>
  );
}
