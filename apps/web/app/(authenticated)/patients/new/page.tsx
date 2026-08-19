"use client";

import { useRef } from "react";
import { useRouter } from "next/navigation";
import { useForm } from "react-hook-form";
import { motion } from "framer-motion";
import { UserPlus } from "lucide-react";
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
import { useRequireAuth } from "@/lib/use-auth";
import { useCreatePatient } from "@/lib/api/patient-api";
import { useSelectedPatient } from "@/lib/use-selected-patient";
import { toast } from "sonner";
import { PageHeader } from "@/components/layout/page-header";

interface NewPatientForm {
  name: string;
  cpf?: string;
  rg?: string;
  birthDate?: string;
  gender?: "male" | "female" | "other";
  socialGender?:
    | "male"
    | "female"
    | "non_binary"
    | "trans_male"
    | "trans_female"
    | "other"
    | "prefer_not_to_say";
  email?: string;
  phone?: string;
  address?: string;
  municipality?: string;
  state?: string;
  maritalStatus?: "single" | "married" | "divorced" | "widowed" | "other";
  occupation?: string;
}

export default function NewPatientPage() {
  useRequireAuth();
  const router = useRouter();
  const formRef = useRef<HTMLFormElement>(null);

  // Navegação por Enter nos campos do formulário (padrão obrigatório).
  useFormNavigation({ formRef });

  const createMutation = useCreatePatient();
  const { setSelectedPatient } = useSelectedPatient();

  const {
    register,
    handleSubmit,
    setValue,
    watch,
    formState: { errors },
  } = useForm<NewPatientForm>({
    defaultValues: {
      name: "",
      cpf: "",
      rg: "",
      birthDate: "",
      gender: undefined,
      socialGender: undefined,
      email: "",
      phone: "",
      address: "",
      municipality: "",
      state: "",
      maritalStatus: undefined,
      occupation: "",
    },
  });

  const onSubmit = (data: NewPatientForm) => {
    // CPF: aceitar vazio; quando preenchido, exigir 11 dígitos.
    let cleanCPF: string | undefined = undefined;
    if (data.cpf && data.cpf.trim() !== "") {
      cleanCPF = data.cpf.replace(/\D/g, "");
      if (cleanCPF.length !== 11) {
        toast.error("CPF deve ter exatamente 11 dígitos");
        return;
      }
    }

    let cleanRG: string | undefined = undefined;
    if (data.rg && data.rg.trim() !== "") {
      cleanRG = data.rg.replace(/[^\dA-Za-z]/g, "");
    }

    const blank = (v?: string) => (v && v.trim() !== "" ? v.trim() : undefined);

    // Envia apenas o que foi preenchido. Nome é o único obrigatório no balcão;
    // o restante da ficha pode ser completado depois (cadastro tolerante).
    const payload = {
      name: data.name.trim(),
      cpf: cleanCPF,
      rg: cleanRG,
      birthDate: blank(data.birthDate),
      gender: data.gender,
      socialGender: data.socialGender,
      email: blank(data.email),
      phone: blank(data.phone),
      address: blank(data.address),
      municipality: blank(data.municipality),
      state: data.state && data.state.trim() !== "" ? data.state.trim().toUpperCase() : undefined,
      maritalStatus: data.maritalStatus,
      occupation: blank(data.occupation),
    };

    createMutation.mutate(payload, {
      onSuccess: (patient) => {
        toast.success("Paciente cadastrado");
        // Seleciona automaticamente o paciente recém-criado (vira o selectedPatient do user).
        setSelectedPatient(patient.id);
        router.push(`/patients/${patient.id}`);
      },
      onError: (error: unknown) => {
        toast.error(
          error instanceof Error
            ? error.message
            : "Não foi possível cadastrar o paciente",
        );
      },
    });
  };

  const selectedGender = watch("gender");
  const selectedSocialGender = watch("socialGender");
  const selectedMaritalStatus = watch("maritalStatus");

  return (
    <div className="min-h-screen p-6">
      <div className="mx-auto max-w-4xl">
        <PageHeader
          title="Novo Paciente"
          description="Cadastro completo. Apenas o nome é obrigatório; o restante pode ser preenchido depois."
          breadcrumbs={[
            { label: "Pacientes", href: "/patients" },
            { label: "Novo" },
          ]}
        />

        <motion.div
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ delay: 0.1 }}
        >
          <form
            ref={formRef}
            onSubmit={handleSubmit(onSubmit)}
            className="space-y-6"
          >
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
                    autoFocus
                  />
                  {errors.name && (
                    <p className="text-sm text-destructive">
                      {errors.name.message}
                    </p>
                  )}
                </div>

                <div className="grid gap-4 md:grid-cols-2">
                  <div className="space-y-2">
                    <Label htmlFor="birthDate">Data de Nascimento</Label>
                    <Input id="birthDate" type="date" {...register("birthDate")} />
                  </div>

                  <div className="space-y-2">
                    <Label htmlFor="gender">Gênero</Label>
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
                      setValue(
                        "socialGender",
                        value as NewPatientForm["socialGender"],
                      )
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
                      <SelectItem value="prefer_not_to_say">
                        Prefiro não dizer
                      </SelectItem>
                    </SelectContent>
                  </Select>
                </div>
              </CardContent>
            </Card>

            {/* Contato e Documentos */}
            <Card className="border-0 shadow-lg">
              <CardHeader>
                <CardTitle>Contato e Documentos</CardTitle>
              </CardHeader>
              <CardContent className="space-y-4">
                <div className="grid gap-4 md:grid-cols-2">
                  <div className="space-y-2">
                    <Label htmlFor="phone">Telefone</Label>
                    <Input
                      id="phone"
                      {...register("phone", {
                        validate: (value) => {
                          if (!value || value.trim() === "") return true;
                          if (value.length < 10)
                            return "Telefone deve ter pelo menos 10 caracteres";
                          if (value.length > 20)
                            return "Telefone não pode ter mais de 20 caracteres";
                          return true;
                        },
                      })}
                      placeholder="(43) 99999-9999"
                    />
                    {errors.phone && (
                      <p className="text-sm text-destructive">
                        {errors.phone.message}
                      </p>
                    )}
                  </div>

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
                </div>

                <div className="grid gap-4 md:grid-cols-2">
                  <div className="space-y-2">
                    <Label htmlFor="cpf">CPF</Label>
                    <Input
                      id="cpf"
                      {...register("cpf", {
                        validate: (value) => {
                          if (!value || value.trim() === "") return true;
                          const cleanCPF = value.replace(/\D/g, "");
                          if (cleanCPF.length !== 11)
                            return "CPF deve ter 11 dígitos";
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
                          if (value.length > 20)
                            return "RG não pode ter mais de 20 caracteres";
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
                    <Label htmlFor="maritalStatus">Estado Civil</Label>
                    <Select
                      value={selectedMaritalStatus}
                      onValueChange={(value) =>
                        setValue(
                          "maritalStatus",
                          value as NewPatientForm["maritalStatus"],
                        )
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

                  <div className="space-y-2">
                    <Label htmlFor="occupation">Ocupação</Label>
                    <Input
                      id="occupation"
                      {...register("occupation", {
                        validate: (value) => {
                          if (!value || value.trim() === "") return true;
                          if (value.length > 100)
                            return "Ocupação não pode ter mais de 100 caracteres";
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
                </div>
              </CardContent>
            </Card>

            {/* Endereço */}
            <Card className="border-0 shadow-lg">
              <CardHeader>
                <CardTitle>Endereço</CardTitle>
              </CardHeader>
              <CardContent className="space-y-4">
                <div className="space-y-2">
                  <Label htmlFor="address">Endereço</Label>
                  <Input
                    id="address"
                    {...register("address", {
                      validate: (value) => {
                        if (!value || value.trim() === "") return true;
                        if (value.length > 500)
                          return "Endereço não pode ter mais de 500 caracteres";
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
                          if (value.length > 100)
                            return "Município não pode ter mais de 100 caracteres";
                          return true;
                        },
                      })}
                      placeholder="Londrina"
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
                          if (value.length !== 2)
                            return "UF deve ter 2 caracteres (ex: PR)";
                          return true;
                        },
                      })}
                      placeholder="PR"
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

            {/* Actions — no celular empilha (o par lado a lado ocupava a largura toda e
                espremia o formulário); ação principal em cima, na altura do polegar. */}
            <div className="flex flex-col-reverse gap-3 sm:flex-row sm:justify-end sm:gap-4">
              <Button
                type="button"
                variant="outline"
                onClick={() => router.push("/patients")}
                className="w-full sm:w-auto"
              >
                Cancelar
              </Button>
              <Button
                type="submit"
                disabled={createMutation.isPending}
                className="w-full gap-2 sm:w-auto"
              >
                {createMutation.isPending ? (
                  <>Salvando...</>
                ) : (
                  <>
                    <UserPlus className="h-4 w-4" />
                    Cadastrar Paciente
                  </>
                )}
              </Button>
            </div>
          </form>
        </motion.div>
      </div>
    </div>
  );
}
