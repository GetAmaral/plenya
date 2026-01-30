"use client";

import { useParams, useRouter } from "next/navigation";
import { useQuery } from "@tanstack/react-query";
import { motion } from "framer-motion";
import { format } from "date-fns";
import { ptBR } from "date-fns/locale";
import {
  ArrowLeft,
  User,
  Calendar,
  Phone,
  Users,
  Ruler,
  Weight as WeightIcon,
  Edit,
} from "lucide-react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import { apiClient } from "@/lib/api-client";
import { useRequireAuth } from "@/lib/use-auth";
import { PageHeader } from "@/components/layout/page-header";

interface Patient {
  id: string;
  userId: string;
  name: string;
  birthDate: string;
  gender: "male" | "female" | "other";
  phone?: string;
  motherName?: string;
  fatherName?: string;
  height?: number;
  weight?: number;
  createdAt: string;
  updatedAt: string;
}

export default function PatientDetailPage() {
  useRequireAuth();
  const params = useParams();
  const router = useRouter();
  const patientId = params.id as string;

  const { data: patient, isLoading } = useQuery({
    queryKey: ["patient", patientId],
    queryFn: async () => {
      return await apiClient.get<Patient>(`/api/v1/patients/${patientId}`);
    },
  });

  if (isLoading) {
    return (
      <div className="min-h-screen p-6">
        <div className="mx-auto max-w-4xl space-y-6">
          <Skeleton className="h-8 w-64" />
          <Skeleton className="h-96 w-full" />
        </div>
      </div>
    );
  }

  if (!patient) {
    return (
      <div className="min-h-screen p-6">
        <div className="mx-auto max-w-4xl">
          <p>Paciente não encontrado</p>
        </div>
      </div>
    );
  }

  const genderLabel = {
    male: "Masculino",
    female: "Feminino",
    other: "Outro",
  }[patient.gender];

  const age = Math.floor(
    (new Date().getTime() - new Date(patient.birthDate).getTime()) /
      (1000 * 60 * 60 * 24 * 365.25)
  );

  return (
    <div className="min-h-screen p-6">
      <div className="mx-auto max-w-4xl">
        {/* Header */}
        <PageHeader
          title={patient.name}
          description="Detalhes do paciente"
          breadcrumbs={[
            { label: "Pacientes", href: "/patients" },
            { label: patient.name }
          ]}
          actions={
            <>
              <Button
                onClick={() => router.push(`/patients/${patientId}/edit`)}
                className="gap-2"
              >
                <Edit className="h-4 w-4" />
                Editar
              </Button>
            </>
          }
        />

        {/* Patient Info */}
        <motion.div
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ delay: 0.1 }}
          className="grid gap-6 md:grid-cols-2"
        >
          {/* Informações Pessoais */}
          <Card className="border-0 shadow-lg">
            <CardHeader>
              <CardTitle>Informações Pessoais</CardTitle>
            </CardHeader>
            <CardContent className="space-y-4">
              <div>
                <p className="text-sm text-muted-foreground">Nome Completo</p>
                <p className="font-medium">{patient.name}</p>
              </div>

              <div>
                <p className="text-sm text-muted-foreground">
                  Data de Nascimento
                </p>
                <div className="flex items-center gap-2">
                  <Calendar className="h-4 w-4 text-muted-foreground" />
                  <p className="font-medium">
                    {format(new Date(patient.birthDate), "dd/MM/yyyy", {
                      locale: ptBR,
                    })}{" "}
                    <span className="text-sm text-muted-foreground">
                      ({age} anos)
                    </span>
                  </p>
                </div>
              </div>

              <div>
                <p className="text-sm text-muted-foreground">Gênero</p>
                <Badge variant="outline">{genderLabel}</Badge>
              </div>

              {patient.phone && (
                <div>
                  <p className="text-sm text-muted-foreground">Telefone</p>
                  <div className="flex items-center gap-2">
                    <Phone className="h-4 w-4 text-muted-foreground" />
                    <p className="font-medium">{patient.phone}</p>
                  </div>
                </div>
              )}
            </CardContent>
          </Card>

          {/* Informações Familiares */}
          <Card className="border-0 shadow-lg">
            <CardHeader>
              <CardTitle>Informações Familiares</CardTitle>
            </CardHeader>
            <CardContent className="space-y-4">
              {patient.motherName ? (
                <div>
                  <p className="text-sm text-muted-foreground">Nome da Mãe</p>
                  <div className="flex items-center gap-2">
                    <Users className="h-4 w-4 text-muted-foreground" />
                    <p className="font-medium">{patient.motherName}</p>
                  </div>
                </div>
              ) : (
                <p className="text-sm text-muted-foreground">
                  Nome da mãe não informado
                </p>
              )}

              {patient.fatherName ? (
                <div>
                  <p className="text-sm text-muted-foreground">Nome do Pai</p>
                  <div className="flex items-center gap-2">
                    <Users className="h-4 w-4 text-muted-foreground" />
                    <p className="font-medium">{patient.fatherName}</p>
                  </div>
                </div>
              ) : (
                <p className="text-sm text-muted-foreground">
                  Nome do pai não informado
                </p>
              )}
            </CardContent>
          </Card>

          {/* Medidas */}
          <Card className="border-0 shadow-lg">
            <CardHeader>
              <CardTitle>Medidas</CardTitle>
            </CardHeader>
            <CardContent className="space-y-4">
              {patient.height ? (
                <div>
                  <p className="text-sm text-muted-foreground">Altura</p>
                  <div className="flex items-center gap-2">
                    <Ruler className="h-4 w-4 text-muted-foreground" />
                    <p className="font-medium">{patient.height} cm</p>
                  </div>
                </div>
              ) : (
                <p className="text-sm text-muted-foreground">
                  Altura não informada
                </p>
              )}

              {patient.weight ? (
                <div>
                  <p className="text-sm text-muted-foreground">Peso</p>
                  <div className="flex items-center gap-2">
                    <WeightIcon className="h-4 w-4 text-muted-foreground" />
                    <p className="font-medium">{patient.weight} kg</p>
                  </div>
                </div>
              ) : (
                <p className="text-sm text-muted-foreground">
                  Peso não informado
                </p>
              )}

              {patient.height && patient.weight && (
                <div>
                  <p className="text-sm text-muted-foreground">IMC</p>
                  <p className="font-medium">
                    {(
                      patient.weight /
                      Math.pow(patient.height / 100, 2)
                    ).toFixed(1)}
                  </p>
                </div>
              )}
            </CardContent>
          </Card>

          {/* Metadados */}
          <Card className="border-0 shadow-lg">
            <CardHeader>
              <CardTitle>Informações do Sistema</CardTitle>
            </CardHeader>
            <CardContent className="space-y-4">
              <div>
                <p className="text-sm text-muted-foreground">ID</p>
                <p className="font-mono text-xs">{patient.id}</p>
              </div>

              <div>
                <p className="text-sm text-muted-foreground">
                  Cadastrado em
                </p>
                <p className="text-sm">
                  {format(new Date(patient.createdAt), "dd/MM/yyyy 'às' HH:mm", {
                    locale: ptBR,
                  })}
                </p>
              </div>

              <div>
                <p className="text-sm text-muted-foreground">
                  Última atualização
                </p>
                <p className="text-sm">
                  {format(new Date(patient.updatedAt), "dd/MM/yyyy 'às' HH:mm", {
                    locale: ptBR,
                  })}
                </p>
              </div>
            </CardContent>
          </Card>
        </motion.div>
      </div>
    </div>
  );
}
