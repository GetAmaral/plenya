"use client";

import React, { useRef } from "react";
import { useParams, useRouter } from "next/navigation";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { useForm } from "react-hook-form";
import { motion } from "framer-motion";
import { ArrowLeft, Save, User } from "lucide-react";
import { useFormNavigation } from "@/lib/use-form-navigation";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Skeleton } from "@/components/ui/skeleton";
import { apiClient } from "@/lib/api-client";
import { useRequireAuth } from "@/lib/use-auth";
import { toast } from "sonner";
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

interface UpdatePatientForm {
  name: string;
  birthDate: string;
  gender: "male" | "female" | "other";
  phone?: string;
  motherName?: string;
  fatherName?: string;
  height?: number;
  weight?: number;
}

export default function EditPatientPage() {
  useRequireAuth();
  const params = useParams();
  const router = useRouter();
  const queryClient = useQueryClient();
  const patientId = params.id as string;
  const formRef = useRef<HTMLFormElement>(null);

  // Navegação por Enter nos campos do formulário
  useFormNavigation({ formRef });

  const { data: patient, isLoading } = useQuery({
    queryKey: ["patient", patientId],
    queryFn: async () => {
      return await apiClient.get<Patient>(`/api/v1/patients/${patientId}`);
    },
  });

  const {
    register,
    handleSubmit,
    setValue,
    watch,
    formState: { errors, isDirty },
  } = useForm<UpdatePatientForm>({
    defaultValues: {
      name: patient?.name || "",
      birthDate: patient?.birthDate || "",
      gender: patient?.gender || "male",
      phone: patient?.phone || "",
      motherName: patient?.motherName || "",
      fatherName: patient?.fatherName || "",
      height: patient?.height || undefined,
      weight: patient?.weight || undefined,
    },
  });

  // Atualizar valores do formulário quando patient carregar
  React.useEffect(() => {
    if (patient) {
      setValue("name", patient.name, { shouldDirty: false });
      setValue("birthDate", patient.birthDate, { shouldDirty: false });
      setValue("gender", patient.gender, { shouldDirty: false });
      setValue("phone", patient.phone || "", { shouldDirty: false });
      setValue("motherName", patient.motherName || "", { shouldDirty: false });
      setValue("fatherName", patient.fatherName || "", { shouldDirty: false });
      setValue("height", patient.height, { shouldDirty: false });
      setValue("weight", patient.weight, { shouldDirty: false });
    }
  }, [patient, setValue]);

  const updateMutation = useMutation({
    mutationFn: async (data: UpdatePatientForm) => {
      return await apiClient.put(`/api/v1/patients/${patientId}`, data);
    },
    onSuccess: () => {
      toast.success("Paciente atualizado com sucesso!");
      queryClient.invalidateQueries({ queryKey: ["patient", patientId] });
      queryClient.invalidateQueries({ queryKey: ["patients"] });
      router.push(`/patients/${patientId}`);
    },
    onError: (error: Error) => {
      toast.error("Erro ao atualizar paciente: " + error.message);
    },
  });

  const onSubmit = (data: UpdatePatientForm) => {
    // Limpar campos vazios (enviar undefined ao invés de string vazia)
    const cleanData = {
      name: data.name,
      birthDate: data.birthDate,
      gender: data.gender,
      phone: data.phone && data.phone.trim() !== "" ? data.phone : undefined,
      motherName: data.motherName && data.motherName.trim() !== "" ? data.motherName : undefined,
      fatherName: data.fatherName && data.fatherName.trim() !== "" ? data.fatherName : undefined,
      height: data.height && data.height > 0 ? data.height : undefined,
      weight: data.weight && data.weight > 0 ? data.weight : undefined,
    };

    updateMutation.mutate(cleanData);
  };

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

  const selectedGender = watch("gender");

  return (
    <div className="min-h-screen p-6">
      <div className="mx-auto max-w-4xl">
        {/* Header */}
        <PageHeader
          title="Editar Paciente"
          description={patient.name}
          breadcrumbs={[
            { label: "Pacientes", href: "/patients" },
            { label: patient.name, href: `/patients/${patientId}` },
            { label: "Editar" }
          ]}
        />

        {/* Form */}
        <motion.div
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ delay: 0.1 }}
        >
          <form ref={formRef} onSubmit={handleSubmit(onSubmit)} className="space-y-6">
            {/* Informações Pessoais */}
            <Card className="border-0 shadow-lg">
              <CardHeader>
                <CardTitle>Informações Pessoais</CardTitle>
              </CardHeader>
              <CardContent className="space-y-4">
                <div className="space-y-2">
                  <Label htmlFor="name">Nome Completo *</Label>
                  <Input
                    id="name"
                    {...register("name", {
                      required: "Nome é obrigatório",
                      minLength: {
                        value: 3,
                        message: "Nome deve ter pelo menos 3 caracteres",
                      },
                      maxLength: {
                        value: 200,
                        message: "Nome não pode ter mais de 200 caracteres",
                      },
                    })}
                    placeholder="João da Silva"
                  />
                  {errors.name && (
                    <p className="text-sm text-destructive">
                      {errors.name.message}
                    </p>
                  )}
                </div>

                <div className="grid gap-4 md:grid-cols-2">
                  <div className="space-y-2">
                    <Label htmlFor="birthDate">Data de Nascimento *</Label>
                    <Input
                      id="birthDate"
                      type="date"
                      {...register("birthDate", {
                        required: "Data de nascimento é obrigatória",
                      })}
                    />
                    {errors.birthDate && (
                      <p className="text-sm text-destructive">
                        {errors.birthDate.message}
                      </p>
                    )}
                  </div>

                  <div className="space-y-2">
                    <Label htmlFor="gender">Gênero *</Label>
                    <Select
                      value={selectedGender}
                      onValueChange={(value) =>
                        setValue("gender", value as "male" | "female" | "other")
                      }
                    >
                      <SelectTrigger>
                        <SelectValue placeholder="Selecione o gênero" />
                      </SelectTrigger>
                      <SelectContent>
                        <SelectItem value="male">Masculino</SelectItem>
                        <SelectItem value="female">Feminino</SelectItem>
                        <SelectItem value="other">Outro</SelectItem>
                      </SelectContent>
                    </Select>
                  </div>
                </div>

                <div className="space-y-2">
                  <Label htmlFor="phone">Telefone</Label>
                  <Input
                    id="phone"
                    {...register("phone", {
                      validate: (value) => {
                        if (!value || value.trim() === "") return true;
                        if (value.length < 10) return "Telefone deve ter pelo menos 10 caracteres";
                        if (value.length > 20) return "Telefone não pode ter mais de 20 caracteres";
                        return true;
                      },
                    })}
                    placeholder="(11) 98765-4321"
                  />
                  {errors.phone && (
                    <p className="text-sm text-destructive">
                      {errors.phone.message}
                    </p>
                  )}
                </div>
              </CardContent>
            </Card>

            {/* Informações Familiares */}
            <Card className="border-0 shadow-lg">
              <CardHeader>
                <CardTitle>Informações Familiares</CardTitle>
              </CardHeader>
              <CardContent className="space-y-4">
                <div className="space-y-2">
                  <Label htmlFor="motherName">Nome da Mãe</Label>
                  <Input
                    id="motherName"
                    {...register("motherName", {
                      validate: (value) => {
                        if (!value || value.trim() === "") return true;
                        if (value.length < 3) return "Nome deve ter pelo menos 3 caracteres";
                        if (value.length > 200) return "Nome não pode ter mais de 200 caracteres";
                        return true;
                      },
                    })}
                    placeholder="Maria da Silva"
                  />
                  {errors.motherName && (
                    <p className="text-sm text-destructive">
                      {errors.motherName.message}
                    </p>
                  )}
                </div>

                <div className="space-y-2">
                  <Label htmlFor="fatherName">Nome do Pai</Label>
                  <Input
                    id="fatherName"
                    {...register("fatherName", {
                      validate: (value) => {
                        if (!value || value.trim() === "") return true;
                        if (value.length < 3) return "Nome deve ter pelo menos 3 caracteres";
                        if (value.length > 200) return "Nome não pode ter mais de 200 caracteres";
                        return true;
                      },
                    })}
                    placeholder="José da Silva"
                  />
                  {errors.fatherName && (
                    <p className="text-sm text-destructive">
                      {errors.fatherName.message}
                    </p>
                  )}
                </div>
              </CardContent>
            </Card>

            {/* Medidas */}
            <Card className="border-0 shadow-lg">
              <CardHeader>
                <CardTitle>Medidas</CardTitle>
              </CardHeader>
              <CardContent className="grid gap-4 md:grid-cols-2">
                <div className="space-y-2">
                  <Label htmlFor="height">Altura (cm)</Label>
                  <Input
                    id="height"
                    type="number"
                    step="0.1"
                    {...register("height", {
                      valueAsNumber: true,
                      validate: (value) => {
                        if (!value || isNaN(value)) return true;
                        if (value <= 0) return "Altura deve ser maior que zero";
                        if (value > 300) return "Altura inválida";
                        return true;
                      },
                    })}
                    placeholder="180"
                  />
                  {errors.height && (
                    <p className="text-sm text-destructive">
                      {errors.height.message}
                    </p>
                  )}
                </div>

                <div className="space-y-2">
                  <Label htmlFor="weight">Peso (kg)</Label>
                  <Input
                    id="weight"
                    type="number"
                    step="0.1"
                    {...register("weight", {
                      valueAsNumber: true,
                      validate: (value) => {
                        if (!value || isNaN(value)) return true;
                        if (value <= 0) return "Peso deve ser maior que zero";
                        if (value > 500) return "Peso inválido";
                        return true;
                      },
                    })}
                    placeholder="68.6"
                  />
                  {errors.weight && (
                    <p className="text-sm text-destructive">
                      {errors.weight.message}
                    </p>
                  )}
                </div>
              </CardContent>
            </Card>

            {/* Actions */}
            <div className="flex gap-4 justify-end">
              <Button
                type="button"
                variant="outline"
                onClick={() => router.push(`/patients/${patientId}`)}
              >
                Cancelar
              </Button>
              <Button
                type="submit"
                disabled={!isDirty || updateMutation.isPending || Object.keys(errors).length > 0}
                className="gap-2"
              >
                <Save className="h-4 w-4" />
                {updateMutation.isPending ? "Salvando..." : "Salvar Alterações"}
              </Button>
            </div>
          </form>
        </motion.div>
      </div>
    </div>
  );
}
