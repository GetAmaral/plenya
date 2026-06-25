"use client";

import { useMemo, useState } from "react";
import { useQuery } from "@tanstack/react-query";
import { useRouter } from "next/navigation";
import { Users, Search, MapPin, Phone, Loader2, Cake, ChevronRight, X } from "lucide-react";
import { PageHeader } from "@/components/layout/page-header";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Avatar, AvatarFallback } from "@/components/ui/avatar";
import { Skeleton } from "@/components/ui/skeleton";
import { apiClient } from "@/lib/api-client";
import { useRequireAuth } from "@/lib/use-auth";
import { useSelectedPatient } from "@/lib/use-selected-patient";
import { Patient } from "@/lib/auth-store";
import { cn } from "@/lib/utils";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import { formatDateOnly, calcAge } from "@/lib/format-date";

type SortBy = "name" | "recent";

function initials(name: string): string {
  const parts = name.trim().split(/\s+/).filter(Boolean);
  if (parts.length === 0) return "?";
  if (parts.length === 1) return parts[0].slice(0, 2).toUpperCase();
  return (parts[0][0] + parts[parts.length - 1][0]).toUpperCase();
}

function genderLabel(gender: string): string {
  return { male: "Masc.", female: "Fem.", other: "Outro" }[gender] || gender;
}

/** Nº de registro curto (MRN-like) — mesmo padrão da barra de contexto do paciente. */
function recordId(id: string): string {
  return id.slice(-6).toUpperCase();
}

export default function PatientSelectPage() {
  useRequireAuth();
  const router = useRouter();
  const { setSelectedPatient } = useSelectedPatient();

  const [searchTerm, setSearchTerm] = useState("");
  const [stateFilter, setStateFilter] = useState<string>("all");
  const [sortBy, setSortBy] = useState<SortBy>("name");
  const [selectingId, setSelectingId] = useState<string | null>(null);

  const { data: patients = [], isLoading } = useQuery<Patient[]>({
    queryKey: ["patients"],
    queryFn: async () => apiClient.get<Patient[]>("/api/v1/patients?limit=1000"),
  });

  const states = useMemo(
    () => Array.from(new Set(patients.map((p) => p.state).filter(Boolean))).sort() as string[],
    [patients]
  );

  const filteredPatients = useMemo(() => {
    const term = searchTerm.trim().toLowerCase();
    const list = patients.filter((p) => {
      const matchesSearch =
        !term ||
        p.name.toLowerCase().includes(term) ||
        (p.phone || "").toLowerCase().includes(term) ||
        recordId(p.id).toLowerCase().includes(term);
      const matchesState = stateFilter === "all" || p.state === stateFilter;
      return matchesSearch && matchesState;
    });
    list.sort((a, b) =>
      sortBy === "name"
        ? a.name.localeCompare(b.name, "pt", { sensitivity: "base" })
        : new Date(b.createdAt).getTime() - new Date(a.createdAt).getTime()
    );
    return list;
  }, [patients, searchTerm, stateFilter, sortBy]);

  const handleSelectPatient = async (patientId: string) => {
    setSelectingId(patientId);
    try {
      await setSelectedPatient(patientId);
      const redirectPath =
        typeof window !== "undefined"
          ? sessionStorage.getItem("plenya-redirect-after-patient-select")
          : null;
      if (typeof window !== "undefined") {
        sessionStorage.removeItem("plenya-redirect-after-patient-select");
      }
      router.push(redirectPath || "/dashboard");
    } catch {
      setSelectingId(null);
    }
  };

  const hasFilters = !!searchTerm || stateFilter !== "all";

  return (
    <div className="mx-auto max-w-4xl p-4 sm:p-6">
      <PageHeader
        title="Selecionar Paciente"
        description="Escolha um paciente para abrir o prontuário."
      />

      {/* Busca + filtros (compactos) */}
      <div className="mt-4 flex flex-col gap-2 sm:flex-row sm:items-center">
        <div className="relative flex-1">
          <Search className="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground" />
          <Input
            autoFocus
            placeholder="Buscar por nome, telefone ou nº de registro…"
            value={searchTerm}
            onChange={(e) => setSearchTerm(e.target.value)}
            className="pl-9"
          />
        </div>
        <Select value={sortBy} onValueChange={(v) => setSortBy(v as SortBy)}>
          <SelectTrigger className="w-full sm:w-[170px]">
            <SelectValue placeholder="Ordenar" />
          </SelectTrigger>
          <SelectContent>
            <SelectItem value="name">Nome (A–Z)</SelectItem>
            <SelectItem value="recent">Mais recentes</SelectItem>
          </SelectContent>
        </Select>
        {states.length > 0 && (
          <Select value={stateFilter} onValueChange={setStateFilter}>
            <SelectTrigger className="w-full sm:w-[130px]">
              <SelectValue placeholder="Estado" />
            </SelectTrigger>
            <SelectContent>
              <SelectItem value="all">Todos UF</SelectItem>
              {states.map((s) => (
                <SelectItem key={s} value={s}>
                  {s}
                </SelectItem>
              ))}
            </SelectContent>
          </Select>
        )}
      </div>

      {/* Contagem */}
      <div className="mt-3 flex items-center gap-2 text-xs text-muted-foreground">
        {isLoading ? (
          <Skeleton className="h-4 w-40" />
        ) : (
          <span>
            {filteredPatients.length} de {patients.length} paciente
            {patients.length === 1 ? "" : "s"}
          </span>
        )}
        {hasFilters && (
          <Button
            variant="ghost"
            size="sm"
            className="h-6 gap-1 px-2 text-xs"
            onClick={() => {
              setSearchTerm("");
              setStateFilter("all");
            }}
          >
            <X className="h-3 w-3" /> Limpar
          </Button>
        )}
      </div>

      {/* Lista densa */}
      <div className="mt-2 overflow-hidden rounded-lg border border-border bg-card">
        {isLoading ? (
          <div className="divide-y divide-border">
            {Array.from({ length: 8 }).map((_, i) => (
              <div key={i} className="flex items-center gap-3 p-3">
                <Skeleton className="h-10 w-10 rounded-full" />
                <div className="flex-1 space-y-2">
                  <Skeleton className="h-4 w-1/3" />
                  <Skeleton className="h-3 w-2/3" />
                </div>
              </div>
            ))}
          </div>
        ) : filteredPatients.length === 0 ? (
          <div className="flex flex-col items-center justify-center gap-3 p-12 text-center">
            <Users className="h-12 w-12 text-muted-foreground/50" />
            <div>
              <p className="font-medium">Nenhum paciente encontrado</p>
              <p className="text-sm text-muted-foreground">Ajuste a busca ou os filtros.</p>
            </div>
          </div>
        ) : (
          <ul className="divide-y divide-border" role="list">
            {filteredPatients.map((p) => {
              const busy = selectingId === p.id;
              const dob = p.birthDate
                ? formatDateOnly(p.birthDate, "dd/MM/yyyy")
                : null;
              const age = calcAge(p.birthDate);
              const city = [p.municipality, p.state].filter(Boolean).join(" - ");
              return (
                <li key={p.id}>
                  <button
                    type="button"
                    onClick={() => handleSelectPatient(p.id)}
                    disabled={!!selectingId}
                    className={cn(
                      "flex w-full items-center gap-3 p-3 text-left transition-colors hover:bg-muted/60 focus-visible:bg-muted focus-visible:outline-none disabled:opacity-60",
                      busy && "bg-muted"
                    )}
                  >
                    <Avatar className="h-10 w-10 shrink-0">
                      <AvatarFallback className="bg-primary/10 text-sm font-semibold text-primary">
                        {initials(p.name)}
                      </AvatarFallback>
                    </Avatar>
                    <div className="min-w-0 flex-1">
                      <div className="flex items-center gap-2">
                        <span className="truncate font-semibold">{p.name}</span>
                        <span className="shrink-0 font-mono text-[11px] text-muted-foreground">
                          #{recordId(p.id)}
                        </span>
                      </div>
                      <div className="mt-0.5 flex flex-wrap items-center gap-x-3 gap-y-0.5 text-xs text-muted-foreground">
                        {dob && (
                          <span className="inline-flex items-center gap-1">
                            <Cake className="h-3 w-3" />
                            {dob}
                            {age !== null && ` (${age} anos)`}
                          </span>
                        )}
                        <span>{genderLabel(p.gender)}</span>
                        {city && (
                          <span className="inline-flex items-center gap-1">
                            <MapPin className="h-3 w-3" />
                            {city}
                          </span>
                        )}
                        {p.phone && (
                          <span className="inline-flex items-center gap-1">
                            <Phone className="h-3 w-3" />
                            {p.phone}
                          </span>
                        )}
                      </div>
                    </div>
                    {busy ? (
                      <Loader2 className="h-4 w-4 shrink-0 animate-spin text-muted-foreground" />
                    ) : (
                      <ChevronRight className="h-4 w-4 shrink-0 text-muted-foreground/50" />
                    )}
                  </button>
                </li>
              );
            })}
          </ul>
        )}
      </div>
    </div>
  );
}
