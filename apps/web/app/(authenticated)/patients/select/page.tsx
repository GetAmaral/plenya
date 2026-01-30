"use client";

import { useState } from "react";
import { useQuery } from "@tanstack/react-query";
import { useRouter } from "next/navigation";
import { motion } from "framer-motion";
import { Users, Search, MapPin, Calendar, Phone, Loader2 } from "lucide-react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import { apiClient } from "@/lib/api-client";
import { useRequireAuth } from "@/lib/use-auth";
import { useSelectedPatient } from "@/lib/use-selected-patient";
import { Patient } from "@/lib/auth-store";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import { format, differenceInYears } from "date-fns";
import { ptBR } from "date-fns/locale";

export default function PatientSelectPage() {
  useRequireAuth();
  const router = useRouter();
  const { setSelectedPatient, isLoading: isUpdating } = useSelectedPatient();

  const [searchTerm, setSearchTerm] = useState("");
  const [municipalityFilter, setMunicipalityFilter] = useState<string>("all");
  const [stateFilter, setStateFilter] = useState<string>("all");

  // Fetch all patients
  const { data: patients = [], isLoading } = useQuery<Patient[]>({
    queryKey: ["patients"],
    queryFn: async () => {
      return apiClient.get<Patient[]>("/api/v1/patients?limit=1000");
    },
  });

  // Get unique municipalities and states for filters
  const municipalities = Array.from(
    new Set(patients.map((p) => p.municipality).filter(Boolean))
  ).sort() as string[];

  const states = Array.from(
    new Set(patients.map((p) => p.state).filter(Boolean))
  ).sort() as string[];

  // Filter patients
  const filteredPatients = patients.filter((patient) => {
    const matchesSearch =
      !searchTerm ||
      patient.name.toLowerCase().includes(searchTerm.toLowerCase());

    const matchesMunicipality =
      !municipalityFilter || municipalityFilter === "all" || patient.municipality === municipalityFilter;

    const matchesState = !stateFilter || stateFilter === "all" || patient.state === stateFilter;

    return matchesSearch && matchesMunicipality && matchesState;
  });

  const handleSelectPatient = async (patientId: string) => {
    await setSelectedPatient(patientId);

    // Check if there's a redirect path saved
    const redirectPath =
      typeof window !== "undefined"
        ? sessionStorage.getItem("plenya-redirect-after-patient-select")
        : null;

    // Clear the saved path
    if (typeof window !== "undefined") {
      sessionStorage.removeItem("plenya-redirect-after-patient-select");
    }

    // Redirect to saved path or dashboard
    router.push(redirectPath || "/dashboard");
  };

  const calculateAge = (birthDate: string) => {
    return differenceInYears(new Date(), new Date(birthDate));
  };

  const getGenderLabel = (gender: string) => {
    const labels = {
      male: "Masculino",
      female: "Feminino",
      other: "Outro",
    };
    return labels[gender as keyof typeof labels] || gender;
  };

  return (
    <div className="min-h-screen p-6 bg-gradient-to-br from-blue-50 via-white to-purple-50">
      <div className="mx-auto max-w-7xl">
        {/* Header */}
        <motion.div
          initial={{ opacity: 0, y: -20 }}
          animate={{ opacity: 1, y: 0 }}
          className="mb-8"
        >
          <div className="flex items-center gap-3 mb-2">
            <Users className="h-10 w-10 text-blue-600" />
            <h1 className="text-4xl font-bold tracking-tight">
              Selecionar Paciente
            </h1>
          </div>
          <p className="text-muted-foreground text-lg">
            Escolha um paciente para visualizar e gerenciar seus registros médicos
          </p>
        </motion.div>

        {/* Filters */}
        <motion.div
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ delay: 0.1 }}
        >
          <Card className="mb-6 border-0 shadow-lg">
            <CardHeader>
              <CardTitle>Filtros de Busca</CardTitle>
            </CardHeader>
            <CardContent>
              <div className="grid gap-4 md:grid-cols-3">
                {/* Search by name */}
                <div className="relative">
                  <Search className="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground" />
                  <Input
                    placeholder="Buscar por nome..."
                    value={searchTerm}
                    onChange={(e) => setSearchTerm(e.target.value)}
                    className="pl-9"
                  />
                </div>

                {/* Filter by municipality */}
                <Select value={municipalityFilter} onValueChange={setMunicipalityFilter}>
                  <SelectTrigger>
                    <SelectValue placeholder="Município" />
                  </SelectTrigger>
                  <SelectContent>
                    <SelectItem value="all">Todos os municípios</SelectItem>
                    {municipalities.map((municipality) => (
                      <SelectItem key={municipality} value={municipality}>
                        {municipality}
                      </SelectItem>
                    ))}
                  </SelectContent>
                </Select>

                {/* Filter by state */}
                <Select value={stateFilter} onValueChange={setStateFilter}>
                  <SelectTrigger>
                    <SelectValue placeholder="Estado" />
                  </SelectTrigger>
                  <SelectContent>
                    <SelectItem value="all">Todos os estados</SelectItem>
                    {states.map((state) => (
                      <SelectItem key={state} value={state}>
                        {state}
                      </SelectItem>
                    ))}
                  </SelectContent>
                </Select>
              </div>

              {/* Active filters summary */}
              {(searchTerm || (municipalityFilter && municipalityFilter !== "all") || (stateFilter && stateFilter !== "all")) && (
                <div className="mt-4 flex gap-2 flex-wrap">
                  {searchTerm && (
                    <Badge variant="secondary" className="gap-1">
                      Busca: {searchTerm}
                    </Badge>
                  )}
                  {municipalityFilter && municipalityFilter !== "all" && (
                    <Badge variant="secondary" className="gap-1">
                      {municipalityFilter}
                    </Badge>
                  )}
                  {stateFilter && stateFilter !== "all" && (
                    <Badge variant="secondary" className="gap-1">
                      {stateFilter}
                    </Badge>
                  )}
                  <Button
                    variant="ghost"
                    size="sm"
                    onClick={() => {
                      setSearchTerm("");
                      setMunicipalityFilter("all");
                      setStateFilter("all");
                    }}
                  >
                    Limpar filtros
                  </Button>
                </div>
              )}
            </CardContent>
          </Card>
        </motion.div>

        {/* Results count */}
        <div className="mb-4 text-sm text-muted-foreground">
          {isLoading ? (
            <Skeleton className="h-4 w-48" />
          ) : (
            <>
              Exibindo {filteredPatients.length} de {patients.length} pacientes
            </>
          )}
        </div>

        {/* Patient Cards */}
        <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
          {isLoading ? (
            Array.from({ length: 6 }).map((_, i) => (
              <Card key={i} className="border-0 shadow-md">
                <CardContent className="p-6">
                  <Skeleton className="h-6 w-3/4 mb-4" />
                  <Skeleton className="h-4 w-1/2 mb-2" />
                  <Skeleton className="h-4 w-2/3 mb-2" />
                  <Skeleton className="h-10 w-full mt-4" />
                </CardContent>
              </Card>
            ))
          ) : filteredPatients.length === 0 ? (
            <div className="col-span-full">
              <Card className="border-0 shadow-md">
                <CardContent className="p-12 text-center">
                  <Users className="h-16 w-16 mx-auto text-muted-foreground mb-4" />
                  <h3 className="text-xl font-semibold mb-2">
                    Nenhum paciente encontrado
                  </h3>
                  <p className="text-muted-foreground">
                    Tente ajustar os filtros ou limpar a busca
                  </p>
                </CardContent>
              </Card>
            </div>
          ) : (
            filteredPatients.map((patient) => (
              <motion.div
                key={patient.id}
                initial={{ opacity: 0, scale: 0.95 }}
                animate={{ opacity: 1, scale: 1 }}
                transition={{ duration: 0.2 }}
              >
                <Card className="border-0 shadow-md hover:shadow-xl transition-shadow h-full">
                  <CardContent className="p-6">
                    <div className="mb-4">
                      <h3 className="text-xl font-bold mb-1">{patient.name}</h3>
                      <div className="flex items-center gap-2 text-sm text-muted-foreground">
                        <Calendar className="h-3 w-3" />
                        <span>
                          {calculateAge(patient.birthDate)} anos •{" "}
                          {getGenderLabel(patient.gender)}
                        </span>
                      </div>
                    </div>

                    <div className="space-y-2 mb-4">
                      {patient.phone && (
                        <div className="flex items-center gap-2 text-sm">
                          <Phone className="h-3 w-3 text-muted-foreground" />
                          <span>{patient.phone}</span>
                        </div>
                      )}

                      {(patient.municipality || patient.state) && (
                        <div className="flex items-center gap-2 text-sm">
                          <MapPin className="h-3 w-3 text-muted-foreground" />
                          <span>
                            {patient.municipality}
                            {patient.municipality && patient.state && " - "}
                            {patient.state}
                          </span>
                        </div>
                      )}
                    </div>

                    <Button
                      className="w-full"
                      onClick={() => handleSelectPatient(patient.id)}
                      disabled={isUpdating}
                    >
                      {isUpdating ? (
                        <>
                          <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                          Selecionando...
                        </>
                      ) : (
                        "Selecionar Paciente"
                      )}
                    </Button>
                  </CardContent>
                </Card>
              </motion.div>
            ))
          )}
        </div>
      </div>
    </div>
  );
}
