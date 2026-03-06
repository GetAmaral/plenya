"use client";

import { motion } from "framer-motion";
import Link from "next/link";
import { format } from "date-fns";
import { ptBR } from "date-fns/locale";
import { Plus, Timer } from "lucide-react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import { useRequireAuth } from "@/lib/use-auth";
import { useRequireSelectedPatient } from "@/lib/use-require-selected-patient";
import { SelectedPatientHeader } from "@/components/patients/SelectedPatientHeader";
import { PageHeader } from "@/components/layout/page-header";
import { usePeriodizations } from "@/lib/api/periodization-api";

const frameworkLabels: Record<string, string> = {
  bompa: "Bompa",
  linear: "Linear",
  undulating: "Ondulatória",
  block: "Bloco",
};

const phaseColors: Record<string, string> = {
  accumulation: "bg-blue-100 text-blue-800",
  transformation: "bg-purple-100 text-purple-800",
  realization: "bg-green-100 text-green-800",
  hypertrophy: "bg-orange-100 text-orange-800",
  strength: "bg-red-100 text-red-800",
  endurance: "bg-teal-100 text-teal-800",
  power: "bg-yellow-100 text-yellow-800",
};

export default function PeriodizationPage() {
  useRequireAuth();
  useRequireSelectedPatient();

  const { data: periodizations, isLoading } = usePeriodizations();

  return (
    <div className="space-y-6">
      <SelectedPatientHeader />
      <PageHeader
        title="Periodização"
        description="Macrociclos e mesociclos de treinamento"
        actions={[
          {
            label: "Nova Periodização",
            icon: <Plus className="h-4 w-4" />,
            onClick: () => window.location.assign("/training/periodization/new"),
          },
        ]}
      />

      {isLoading ? (
        <Skeleton className="h-48" />
      ) : !periodizations?.length ? (
        <Card>
          <CardContent className="flex flex-col items-center justify-center py-12 text-muted-foreground">
            <Timer className="h-12 w-12 mb-4" />
            <p>Nenhuma periodização encontrada.</p>
            <Link href="/training/periodization/new" className="mt-4">
              <Button variant="outline">Criar periodização</Button>
            </Link>
          </CardContent>
        </Card>
      ) : (
        <div className="space-y-4">
          {periodizations.map((p) => (
            <motion.div
              key={p.id}
              initial={{ opacity: 0, y: 10 }}
              animate={{ opacity: 1, y: 0 }}
            >
              <Card>
                <CardHeader className="pb-2">
                  <div className="flex items-center justify-between">
                    <CardTitle className="text-base">{p.objective}</CardTitle>
                    <Badge>{frameworkLabels[p.framework] || p.framework}</Badge>
                  </div>
                  <p className="text-sm text-muted-foreground">
                    {format(new Date(p.startDate), "dd/MM/yyyy", { locale: ptBR })} •{" "}
                    {p.totalWeeks} semanas
                  </p>
                </CardHeader>
                <CardContent>
                  {p.mesocycles?.length > 0 && (
                    <div className="flex gap-1 flex-wrap">
                      {p.mesocycles.map((m) => (
                        <Badge key={m.id} className={phaseColors[m.phase] || ""}>
                          {m.phase} ({m.durationWeeks}sem)
                        </Badge>
                      ))}
                    </div>
                  )}
                </CardContent>
              </Card>
            </motion.div>
          ))}
        </div>
      )}
    </div>
  );
}
