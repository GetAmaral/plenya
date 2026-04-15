"use client";

import { motion } from "framer-motion";
import Link from "next/link";
import { format } from "date-fns";
import { ptBR } from "date-fns/locale";
import { Plus, User } from "lucide-react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import { useRequireAuth } from "@/lib/use-auth";
import { useRequireSelectedPatient } from "@/lib/use-require-selected-patient";
import { SelectedPatientHeader } from "@/components/patients/SelectedPatientHeader";
import { PageHeader } from "@/components/layout/page-header";
import { usePosturalAssessments } from "@/lib/api/posture-api";

const classColors: Record<string, string> = {
  Excelente: "bg-green-100 text-green-800",
  Boa: "bg-blue-100 text-blue-800",
  Regular: "bg-yellow-100 text-yellow-800",
  "Necessita Atencao": "bg-red-100 text-red-800",
};

const viewLabels: Record<string, string> = {
  front: "Frontal",
  side_left: "Lateral Esq.",
  side_right: "Lateral Dir.",
  back: "Posterior",
};

export default function PosturalAssessmentsPage() {
  useRequireAuth();
  useRequireSelectedPatient();

  const { data: assessments, isLoading } = usePosturalAssessments();

  return (
    <div className="space-y-6">
      <SelectedPatientHeader />
      <PageHeader
        title="Avaliacoes Posturais"
        description="Analise postural com medicoes goniometricas"
        actions={[
          {
            label: "Nova Avaliacao",
            icon: <Plus className="h-4 w-4" />,
            onClick: () => window.location.assign("/training/posture/new"),
          },
        ]}
      />

      {isLoading ? (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
          {Array.from({ length: 3 }).map((_, i) => (
            <Skeleton key={i} className="h-48" />
          ))}
        </div>
      ) : !assessments?.length ? (
        <Card>
          <CardContent className="flex flex-col items-center justify-center py-12 text-muted-foreground">
            <User className="h-12 w-12 mb-4" />
            <p>Nenhuma avaliacao postural encontrada.</p>
            <Link href="/training/posture/new" className="mt-4">
              <Button variant="outline">Criar primeira avaliacao</Button>
            </Link>
          </CardContent>
        </Card>
      ) : (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
          {assessments.map((a) => (
            <motion.div
              key={a.id}
              initial={{ opacity: 0, y: 10 }}
              animate={{ opacity: 1, y: 0 }}
            >
              <Link href={`/training/posture/${a.id}`}>
                <Card className="hover:shadow-lg transition-shadow cursor-pointer">
                  <CardHeader className="pb-2">
                    <div className="flex items-center justify-between">
                      <CardTitle className="text-sm font-medium">
                        {format(new Date(a.assessmentDate), "dd/MM/yyyy", { locale: ptBR })}
                      </CardTitle>
                      <Badge className={classColors[a.posturalClassification] || ""}>
                        {a.posturalClassification}
                      </Badge>
                    </div>
                  </CardHeader>
                  <CardContent>
                    <div className="flex items-center justify-between">
                      <div className="text-center flex-1">
                        <p className="text-3xl font-bold">{a.posturalScore}</p>
                        <p className="text-xs text-muted-foreground">Score</p>
                      </div>
                      <div className="text-right">
                        <Badge variant="outline">{viewLabels[a.viewType] || a.viewType}</Badge>
                        {a.severeDeviations > 0 && (
                          <p className="text-xs text-red-600 mt-1">{a.severeDeviations} desvio(s) severo(s)</p>
                        )}
                      </div>
                    </div>
                  </CardContent>
                </Card>
              </Link>
            </motion.div>
          ))}
        </div>
      )}
    </div>
  );
}
