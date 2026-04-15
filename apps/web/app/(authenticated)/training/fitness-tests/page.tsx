"use client";

import { motion } from "framer-motion";
import Link from "next/link";
import { format } from "date-fns";
import { ptBR } from "date-fns/locale";
import { Plus, Target } from "lucide-react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import { useRequireAuth } from "@/lib/use-auth";
import { useRequireSelectedPatient } from "@/lib/use-require-selected-patient";
import { SelectedPatientHeader } from "@/components/patients/SelectedPatientHeader";
import { PageHeader } from "@/components/layout/page-header";
import { useFitnessTests } from "@/lib/api/fitness-test-api";

const classColors: Record<string, string> = {
  Excelente: "bg-green-100 text-green-800",
  Bom: "bg-blue-100 text-blue-800",
  Medio: "bg-yellow-100 text-yellow-800",
  "Abaixo da Media": "bg-orange-100 text-orange-800",
  Fraco: "bg-red-100 text-red-800",
};

export default function FitnessTestsPage() {
  useRequireAuth();
  useRequireSelectedPatient();

  const { data: tests, isLoading } = useFitnessTests();

  return (
    <div className="space-y-6">
      <SelectedPatientHeader />
      <PageHeader
        title="Testes de Fitness"
        description="Avaliacao de resistencia muscular, forca e condicionamento"
        actions={[
          {
            label: "Novo Teste",
            icon: <Plus className="h-4 w-4" />,
            onClick: () => window.location.assign("/training/fitness-tests/new"),
          },
        ]}
      />

      {isLoading ? (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
          {Array.from({ length: 3 }).map((_, i) => (
            <Skeleton key={i} className="h-48" />
          ))}
        </div>
      ) : !tests?.length ? (
        <Card>
          <CardContent className="flex flex-col items-center justify-center py-12 text-muted-foreground">
            <Target className="h-12 w-12 mb-4" />
            <p>Nenhum teste de fitness encontrado.</p>
            <Link href="/training/fitness-tests/new" className="mt-4">
              <Button variant="outline">Criar primeiro teste</Button>
            </Link>
          </CardContent>
        </Card>
      ) : (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
          {tests.map((test) => (
            <motion.div
              key={test.id}
              initial={{ opacity: 0, y: 10 }}
              animate={{ opacity: 1, y: 0 }}
            >
              <Link href={`/training/fitness-tests/${test.id}`}>
                <Card className="hover:shadow-lg transition-shadow cursor-pointer">
                  <CardHeader className="pb-2">
                    <div className="flex items-center justify-between">
                      <CardTitle className="text-sm font-medium">
                        {format(new Date(test.assessmentDate), "dd/MM/yyyy", { locale: ptBR })}
                      </CardTitle>
                      <Badge className={classColors[test.overallClassification] || ""}>
                        {test.overallClassification}
                      </Badge>
                    </div>
                  </CardHeader>
                  <CardContent>
                    <div className="text-center">
                      <p className="text-3xl font-bold">{test.overallScore}</p>
                      <p className="text-xs text-muted-foreground">Score Geral</p>
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
