"use client";

import { motion } from "framer-motion";
import Link from "next/link";
import { format } from "date-fns";
import { ptBR } from "date-fns/locale";
import { Plus, Activity } from "lucide-react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import { useRequireAuth } from "@/lib/use-auth";
import { useRequireSelectedPatient } from "@/lib/use-require-selected-patient";
import { SelectedPatientHeader } from "@/components/patients/SelectedPatientHeader";
import { PageHeader } from "@/components/layout/page-header";
import { usePhysicalAssessments } from "@/lib/api/physical-assessment-api";

const riskColors: Record<string, string> = {
  low: "bg-green-100 text-green-800",
  moderate: "bg-yellow-100 text-yellow-800",
  high: "bg-red-100 text-red-800",
};

const riskLabels: Record<string, string> = {
  low: "Baixo Risco",
  moderate: "Risco Moderado",
  high: "Alto Risco",
};

export default function PhysicalAssessmentsPage() {
  useRequireAuth();
  useRequireSelectedPatient();

  const { data: assessments, isLoading } = usePhysicalAssessments();

  return (
    <div className="space-y-6">
      <SelectedPatientHeader />
      <PageHeader
        title="Avaliações Físicas"
        subtitle="Estratificação de risco ACSM"
        icon={Activity}
        actions={
          <Link href="/training/physical-assessments/new">
            <Button>
              <Plus className="h-4 w-4 mr-2" />
              Nova Avaliação
            </Button>
          </Link>
        }
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
            <Activity className="h-12 w-12 mb-4" />
            <p>Nenhuma avaliação física encontrada.</p>
            <Link href="/training/physical-assessments/new" className="mt-4">
              <Button variant="outline">Criar primeira avaliação</Button>
            </Link>
          </CardContent>
        </Card>
      ) : (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
          {assessments.map((assessment) => (
            <motion.div
              key={assessment.id}
              initial={{ opacity: 0, y: 10 }}
              animate={{ opacity: 1, y: 0 }}
            >
              <Link href={`/training/physical-assessments/${assessment.id}`}>
                <Card className="hover:shadow-lg transition-shadow cursor-pointer">
                  <CardHeader className="pb-2">
                    <div className="flex items-center justify-between">
                      <CardTitle className="text-sm font-medium">
                        {format(new Date(assessment.assessmentDate), "dd/MM/yyyy", { locale: ptBR })}
                      </CardTitle>
                      {assessment.acsmRiskLevel && (
                        <Badge className={riskColors[assessment.acsmRiskLevel] || ""}>
                          {riskLabels[assessment.acsmRiskLevel] || assessment.acsmRiskLevel}
                        </Badge>
                      )}
                    </div>
                  </CardHeader>
                  <CardContent>
                    <div className="space-y-2">
                      <p className="text-sm text-muted-foreground">
                        {assessment.acsmRiskFactorsCount} fator(es) de risco
                      </p>
                      <div className="flex flex-wrap gap-1">
                        {assessment.acsmTags?.slice(0, 4).map((tag) => (
                          <Badge key={tag} variant="outline" className="text-xs">
                            {tag}
                          </Badge>
                        ))}
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
