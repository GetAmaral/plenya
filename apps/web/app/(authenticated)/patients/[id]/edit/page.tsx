"use client";

import React, { useRef } from "react";
import { useParams, useRouter } from "next/navigation";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { useForm } from "react-hook-form";
import { motion } from "framer-motion";
import { ArrowLeft, Save, User, Heart, AlertCircle } from "lucide-react";
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
  cpf?: string;
  birthDate: string;
  gender: "male" | "female" | "other";
  socialGender?: "male" | "female" | "non_binary" | "trans_male" | "trans_female" | "other" | "prefer_not_to_say";
  menopause?: boolean;
  age: number;
  ageText: string;
  phone?: string;
  address?: string;
  municipality?: string;
  state?: string;
  motherName?: string;
  fatherName?: string;
  height?: number;
  weight?: number;
  rg?: string;
  email?: string;
  bloodType?: "A+" | "A-" | "B+" | "B-" | "AB+" | "AB-" | "O+" | "O-";
  maritalStatus?: "single" | "married" | "divorced" | "widowed" | "other";
  occupation?: string;
  emergencyContact?: string;
  emergencyPhone?: string;
  createdAt: string;
  updatedAt: string;
}

interface UpdatePatientForm {
  name: string;
  cpf?: string;
  rg?: string;
  birthDate: string;
  gender: "male" | "female" | "other";
  socialGender?: "male" | "female" | "non_binary" | "trans_male" | "trans_female" | "other" | "prefer_not_to_say";
  menopause?: boolean;
  email?: string;
  phone?: string;
  address?: string;
  municipality?: string;
  state?: string;
  motherName?: string;
  fatherName?: string;
  height?: number;
  weight?: number;
  bloodType?: "A+" | "A-" | "B+" | "B-" | "AB+" | "AB-" | "O+" | "O-";
  maritalStatus?: "single" | "married" | "divorced" | "widowed" | "other";
  occupation?: string;
  emergencyContact?: string;
  emergencyPhone?: string;
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
      cpf: patient?.cpf || "",
      rg: patient?.rg || "",
      birthDate: patient?.birthDate || "",
      gender: patient?.gender || "male",
      socialGender: patient?.socialGender || undefined,
      menopause: patient?.menopause || undefined,
      email: patient?.email || "",
      phone: patient?.phone || "",
      address: patient?.address || "",
      municipality: patient?.municipality || "",
      state: patient?.state || "",
      motherName: patient?.motherName || "",
      fatherName: patient?.fatherName || "",
      height: patient?.height || undefined,
      weight: patient?.weight || undefined,
      bloodType: patient?.bloodType || undefined,
      maritalStatus: patient?.maritalStatus || undefined,
      occupation: patient?.occupation || "",
      emergencyContact: patient?.emergencyContact || "",
      emergencyPhone: patient?.emergencyPhone || "",
    },
  });

  // Atualizar valores do formulário quando patient carregar
  React.useEffect(() => {
    if (patient) {
      setValue("name", patient.name, { shouldDirty: false });
      setValue("cpf", patient.cpf || "", { shouldDirty: false });
      setValue("rg", patient.rg || "", { shouldDirty: false });
      setValue("birthDate", patient.birthDate, { shouldDirty: false });
      setValue("gender", patient.gender, { shouldDirty: false });
      setValue("socialGender", patient.socialGender, { shouldDirty: false });
      setValue("menopause", patient.menopause, { shouldDirty: false });
      setValue("email", patient.email || "", { shouldDirty: false });
      setValue("phone", patient.phone || "", { shouldDirty: false });
      setValue("address", patient.address || "", { shouldDirty: false });
      setValue("municipality", patient.municipality || "", { shouldDirty: false });
      setValue("state", patient.state || "", { shouldDirty: false });
      setValue("motherName", patient.motherName || "", { shouldDirty: false });
      setValue("fatherName", patient.fatherName || "", { shouldDirty: false });
      setValue("height", patient.height, { shouldDirty: false });
      setValue("weight", patient.weight, { shouldDirty: false });
      setValue("bloodType", patient.bloodType, { shouldDirty: false });
      setValue("maritalStatus", patient.maritalStatus, { shouldDirty: false });
      setValue("occupation", patient.occupation || "", { shouldDirty: false });
      setValue("emergencyContact", patient.emergencyContact || "", { shouldDirty: false });
      setValue("emergencyPhone", patient.emergencyPhone || "", { shouldDirty: false });
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
    // Limpar CPF (remover formatação: apenas dígitos)
    let cleanCPF: string | undefined = undefined;
    if (data.cpf && data.cpf.trim() !== "") {
      cleanCPF = data.cpf.replace(/\D/g, "");
      if (cleanCPF.length !== 11) {
        toast.error("CPF deve ter exatamente 11 dígitos");
        return;
      }
    }

    // Limpar RG (remover formatação: apenas dígitos e letras)
    let cleanRG: string | undefined = undefined;
    if (data.rg && data.rg.trim() !== "") {
      cleanRG = data.rg.replace(/[^\dA-Za-z]/g, "");
    }

    // Limpar campos vazios (enviar undefined ao invés de string vazia)
    const cleanData = {
      name: data.name,
      cpf: cleanCPF,
      rg: cleanRG,
      birthDate: data.birthDate,
      gender: data.gender,
      socialGender: data.socialGender,
      menopause: data.menopause,
      email: data.email && data.email.trim() !== "" ? data.email : undefined,
      phone: data.phone && data.phone.trim() !== "" ? data.phone : undefined,
      address: data.address && data.address.trim() !== "" ? data.address : undefined,
      municipality: data.municipality && data.municipality.trim() !== "" ? data.municipality : undefined,
      state: data.state && data.state.trim() !== "" ? data.state?.toUpperCase() : undefined,
      motherName: data.motherName && data.motherName.trim() !== "" ? data.motherName : undefined,
      fatherName: data.fatherName && data.fatherName.trim() !== "" ? data.fatherName : undefined,
      height: data.height && data.height > 0 ? data.height : undefined,
      weight: data.weight && data.weight > 0 ? data.weight : undefined,
      bloodType: data.bloodType,
      maritalStatus: data.maritalStatus,
      occupation: data.occupation && data.occupation.trim() !== "" ? data.occupation : undefined,
      emergencyContact: data.emergencyContact && data.emergencyContact.trim() !== "" ? data.emergencyContact : undefined,
      emergencyPhone: data.emergencyPhone && data.emergencyPhone.trim() !== "" ? data.emergencyPhone : undefined,
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
  const selectedSocialGender = watch("socialGender");
  const selectedBloodType = watch("bloodType");
  const selectedMaritalStatus = watch("maritalStatus");

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
                  <Label htmlFor="socialGender">Gênero Social</Label>
                  <Select
                    value={selectedSocialGender}
                    onValueChange={(value) =>
                      setValue("socialGender", value as "male" | "female" | "non_binary" | "trans_male" | "trans_female" | "other" | "prefer_not_to_say")
                    }
                  >
                    <SelectTrigger>
                      <SelectValue placeholder="Selecione o gênero social" />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectItem value="male">Masculino</SelectItem>
                      <SelectItem value="female">Feminino</SelectItem>
                      <SelectItem value="non_binary">Não-binário</SelectItem>
                      <SelectItem value="trans_male">Homem Trans</SelectItem>
                      <SelectItem value="trans_female">Mulher Trans</SelectItem>
                      <SelectItem value="other">Outro</SelectItem>
                      <SelectItem value="prefer_not_to_say">Prefiro não dizer</SelectItem>
                    </SelectContent>
                  </Select>
                </div>

                {selectedGender === "female" && (
                  <div className="space-y-2">
                    <Label htmlFor="menopause">Menopausa</Label>
                    <Select
                      value={watch("menopause") === undefined ? "undefined" : watch("menopause") ? "true" : "false"}
                      onValueChange={(value) =>
                        setValue("menopause", value === "undefined" ? undefined : value === "true")
                      }
                    >
                      <SelectTrigger>
                        <SelectValue placeholder="Selecione" />
                      </SelectTrigger>
                      <SelectContent>
                        <SelectItem value="undefined">Não informado</SelectItem>
                        <SelectItem value="true">Sim</SelectItem>
                        <SelectItem value="false">Não</SelectItem>
                      </SelectContent>
                    </Select>
                  </div>
                )}

                <div className="space-y-2">
                  <Label htmlFor="email">E-mail</Label>
                  <Input
                    id="email"
                    type="email"
                    {...register("email", {
                      validate: (value) => {
                        if (!value || value.trim() === "") return true;
                        if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(value))
                          return "E-mail inválido";
                        return true;
                      },
                    })}
                    placeholder="joao@example.com"
                  />
                  {errors.email && (
                    <p className="text-sm text-destructive">
                      {errors.email.message}
                    </p>
                  )}
                </div>

                <div className="grid gap-4 md:grid-cols-2">
                  <div className="space-y-2">
                    <Label htmlFor="cpf">CPF</Label>
                    <Input
                      id="cpf"
                      {...register("cpf", {
                        validate: (value) => {
                          if (!value || value.trim() === "") return true;
                          // Remove formatting
                          const cleanCPF = value.replace(/\D/g, "");
                          if (cleanCPF.length !== 11) return "CPF deve ter 11 dígitos";
                          return true;
                        },
                      })}
                      placeholder="000.000.000-00"
                      maxLength={14}
                    />
                    {errors.cpf && (
                      <p className="text-sm text-destructive">
                        {errors.cpf.message}
                      </p>
                    )}
                  </div>

                  <div className="space-y-2">
                    <Label htmlFor="rg">RG</Label>
                    <Input
                      id="rg"
                      {...register("rg", {
                        validate: (value) => {
                          if (!value || value.trim() === "") return true;
                          if (value.length > 20) return "RG não pode ter mais de 20 caracteres";
                          return true;
                        },
                      })}
                      placeholder="00.000.000-0"
                    />
                    {errors.rg && (
                      <p className="text-sm text-destructive">
                        {errors.rg.message}
                      </p>
                    )}
                  </div>
                </div>

                <div className="grid gap-4 md:grid-cols-2">
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

                  <div className="space-y-2">
                    <Label htmlFor="maritalStatus">Estado Civil</Label>
                    <Select
                      value={selectedMaritalStatus}
                      onValueChange={(value) =>
                        setValue("maritalStatus", value as "single" | "married" | "divorced" | "widowed" | "other")
                      }
                    >
                      <SelectTrigger>
                        <SelectValue placeholder="Selecione o estado civil" />
                      </SelectTrigger>
                      <SelectContent>
                        <SelectItem value="single">Solteiro(a)</SelectItem>
                        <SelectItem value="married">Casado(a)</SelectItem>
                        <SelectItem value="divorced">Divorciado(a)</SelectItem>
                        <SelectItem value="widowed">Viúvo(a)</SelectItem>
                        <SelectItem value="other">Outro</SelectItem>
                      </SelectContent>
                    </Select>
                  </div>
                </div>

                <div className="space-y-2">
                  <Label htmlFor="occupation">Ocupação</Label>
                  <Input
                    id="occupation"
                    {...register("occupation", {
                      validate: (value) => {
                        if (!value || value.trim() === "") return true;
                        if (value.length > 100) return "Ocupação não pode ter mais de 100 caracteres";
                        return true;
                      },
                    })}
                    placeholder="Engenheiro de Software"
                  />
                  {errors.occupation && (
                    <p className="text-sm text-destructive">
                      {errors.occupation.message}
                    </p>
                  )}
                </div>

                <div className="space-y-2">
                  <Label htmlFor="address">Endereço</Label>
                  <Input
                    id="address"
                    {...register("address", {
                      validate: (value) => {
                        if (!value || value.trim() === "") return true;
                        if (value.length > 500) return "Endereço não pode ter mais de 500 caracteres";
                        return true;
                      },
                    })}
                    placeholder="Rua Exemplo, 123 - Bairro"
                  />
                  {errors.address && (
                    <p className="text-sm text-destructive">
                      {errors.address.message}
                    </p>
                  )}
                </div>

                <div className="grid gap-4 md:grid-cols-2">
                  <div className="space-y-2">
                    <Label htmlFor="municipality">Município</Label>
                    <Input
                      id="municipality"
                      {...register("municipality", {
                        validate: (value) => {
                          if (!value || value.trim() === "") return true;
                          if (value.length > 100) return "Município não pode ter mais de 100 caracteres";
                          return true;
                        },
                      })}
                      placeholder="São Paulo"
                    />
                    {errors.municipality && (
                      <p className="text-sm text-destructive">
                        {errors.municipality.message}
                      </p>
                    )}
                  </div>

                  <div className="space-y-2">
                    <Label htmlFor="state">Estado (UF)</Label>
                    <Input
                      id="state"
                      {...register("state", {
                        validate: (value) => {
                          if (!value || value.trim() === "") return true;
                          if (value.length !== 2) return "UF deve ter 2 caracteres (ex: SP)";
                          return true;
                        },
                      })}
                      placeholder="SP"
                      maxLength={2}
                    />
                    {errors.state && (
                      <p className="text-sm text-destructive">
                        {errors.state.message}
                      </p>
                    )}
                  </div>
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

            {/* Informações de Saúde */}
            <Card className="border-0 shadow-lg">
              <CardHeader className="flex flex-row items-center gap-2">
                <Heart className="h-5 w-5 text-muted-foreground" />
                <CardTitle>Informações de Saúde</CardTitle>
              </CardHeader>
              <CardContent className="space-y-4">
                <div className="grid gap-4 md:grid-cols-2">
                  <div className="space-y-2">
                    <Label htmlFor="age">Idade</Label>
                    <Input
                      id="age"
                      type="number"
                      value={patient.age}
                      readOnly
                      className="bg-muted"
                    />
                  </div>

                  <div className="space-y-2">
                    <Label htmlFor="ageText">Idade Detalhada</Label>
                    <Input
                      id="ageText"
                      value={patient.ageText}
                      readOnly
                      className="bg-muted"
                    />
                  </div>
                </div>

                <div className="space-y-2">
                  <Label htmlFor="bloodType">Tipo Sanguíneo</Label>
                  <Select
                    value={selectedBloodType}
                    onValueChange={(value) =>
                      setValue("bloodType", value as "A+" | "A-" | "B+" | "B-" | "AB+" | "AB-" | "O+" | "O-")
                    }
                  >
                    <SelectTrigger>
                      <SelectValue placeholder="Selecione o tipo sanguíneo" />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectItem value="A+">A+</SelectItem>
                      <SelectItem value="A-">A-</SelectItem>
                      <SelectItem value="B+">B+</SelectItem>
                      <SelectItem value="B-">B-</SelectItem>
                      <SelectItem value="AB+">AB+</SelectItem>
                      <SelectItem value="AB-">AB-</SelectItem>
                      <SelectItem value="O+">O+</SelectItem>
                      <SelectItem value="O-">O-</SelectItem>
                    </SelectContent>
                  </Select>
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

            {/* Contato de Emergência */}
            <Card className="border-0 shadow-lg">
              <CardHeader className="flex flex-row items-center gap-2">
                <AlertCircle className="h-5 w-5 text-muted-foreground" />
                <CardTitle>Contato de Emergência</CardTitle>
              </CardHeader>
              <CardContent className="grid gap-4 md:grid-cols-2">
                <div className="space-y-2">
                  <Label htmlFor="emergencyContact">Nome do Contato</Label>
                  <Input
                    id="emergencyContact"
                    {...register("emergencyContact", {
                      validate: (value) => {
                        if (!value || value.trim() === "") return true;
                        if (value.length < 3) return "Nome deve ter pelo menos 3 caracteres";
                        if (value.length > 200) return "Nome não pode ter mais de 200 caracteres";
                        return true;
                      },
                    })}
                    placeholder="Maria da Silva"
                  />
                  {errors.emergencyContact && (
                    <p className="text-sm text-destructive">
                      {errors.emergencyContact.message}
                    </p>
                  )}
                </div>

                <div className="space-y-2">
                  <Label htmlFor="emergencyPhone">Telefone de Emergência</Label>
                  <Input
                    id="emergencyPhone"
                    type="tel"
                    {...register("emergencyPhone", {
                      validate: (value) => {
                        if (!value || value.trim() === "") return true;
                        if (value.length < 10) return "Telefone deve ter pelo menos 10 caracteres";
                        if (value.length > 20) return "Telefone não pode ter mais de 20 caracteres";
                        return true;
                      },
                    })}
                    placeholder="(11) 98765-4321"
                  />
                  {errors.emergencyPhone && (
                    <p className="text-sm text-destructive">
                      {errors.emergencyPhone.message}
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
