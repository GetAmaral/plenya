"use client";

import Link from "next/link";
import { Activity, Dumbbell, Timer, Library } from "lucide-react";
import { Card, CardContent, CardHeader, CardTitle, CardDescription } from "@/components/ui/card";
import { useRequireAuth } from "@/lib/use-auth";
import { useRequireSelectedPatient } from "@/lib/use-require-selected-patient";
import { SelectedPatientHeader } from "@/components/patients/SelectedPatientHeader";
import { PageHeader } from "@/components/layout/page-header";

const modules = [
  {
    title: "Avaliação Física",
    description: "Estratificação de risco ACSM, composição corporal, zonas de FC",
    href: "/training/physical-assessments",
    icon: Activity,
  },
  {
    title: "Planos de Treino",
    description: "Criar e gerenciar planos de treino com exercícios e séries",
    href: "/training/workout-plans",
    icon: Dumbbell,
  },
  {
    title: "Periodização",
    description: "Macrociclos e mesociclos de treinamento (Bompa, Linear, Ondulatória)",
    href: "/training/periodization",
    icon: Timer,
  },
  {
    title: "Exercícios",
    description: "Catálogo com ~1500 exercícios, GIFs e instruções",
    href: "/training/exercises",
    icon: Library,
  },
];

export default function TrainingDashboardPage() {
  useRequireAuth();
  useRequireSelectedPatient();

  return (
    <div className="space-y-6">
      <SelectedPatientHeader />
      <PageHeader
        title="Treinamento"
        subtitle="Módulo de avaliação física e prescrição de treinos"
        icon={Dumbbell}
      />

      <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
        {modules.map((mod) => (
          <Link key={mod.href} href={mod.href}>
            <Card className="hover:shadow-lg transition-shadow cursor-pointer h-full">
              <CardHeader>
                <div className="flex items-center gap-3">
                  <mod.icon className="h-8 w-8 text-primary" />
                  <div>
                    <CardTitle className="text-base">{mod.title}</CardTitle>
                    <CardDescription>{mod.description}</CardDescription>
                  </div>
                </div>
              </CardHeader>
            </Card>
          </Link>
        ))}
      </div>
    </div>
  );
}
