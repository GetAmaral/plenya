"use client";

import { useMemo } from "react";
import { useRouter } from "next/navigation";
import { useQuery } from "@tanstack/react-query";
import {
  ArrowDown,
  ArrowUp,
  ChevronLeft,
  ChevronRight,
  Plus,
  Search,
} from "lucide-react";
import type { DateRange } from "react-day-picker";

import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import { FiltersMobileSheet } from "@/components/filters/filters-mobile-sheet";
import { PatientsFilterBar } from "@/components/patients/patients-filter-bar";
import { PageHeader } from "@/components/layout/page-header";
import { apiClient } from "@/lib/api-client";
import { useRequireAuth } from "@/lib/use-auth";
import { usePatientGuard } from "@/lib/use-patient-guard";
import { useDebouncedValue } from "@/lib/use-debounced-value";
import { useUrlFilters } from "@/lib/use-url-filters";
import { formatDate } from "@/lib/format-date";

interface Patient {
  id: string;
  userId: string;
  name: string;
  cpf?: string;
  email?: string;
  birthDate: string;
  age: number;
  ageText: string;
  gender: "male" | "female" | "other";
  phone: string;
  createdAt: string;
  updatedAt: string;
  /** Pode vir do backend em alguns endpoints. Se ausente, switch "plano ativo"
   *  é desligado de fato (caveat documentado). */
  hasActiveSubscription?: boolean;
}

interface PatientsResponse {
  data: Patient[];
  total: number;
  page: number;
  limit: number;
}

type SortKey = "name" | "createdAt" | "age";
type SortDir = "asc" | "desc";

const PAGE_SIZE = 25;

/**
 * Busca é server-side: /api/v1/patients?search= casa nome/telefone (ILIKE) e
 * CPF (blind index HMAC, já que o CPF é criptografado). Sem termo, baixamos um
 * lote (MAX_BUFFER) pra navegação paginada client-side. Os filtros de período/
 * plano e a ordenação seguem client-side sobre o conjunto retornado.
 */
const MAX_BUFFER = 500;

export default function PatientsPage() {
  useRequireAuth();
  usePatientGuard();
  const router = useRouter();

  const { params, setParam, setParams, clearAll } = useUrlFilters();

  // ---- estado UI vinda da URL ------------------------------------------------
  const search = params.get("q") ?? "";
  const fromIso = params.get("from");
  const toIso = params.get("to");
  const dateRange: DateRange | undefined = useMemo(() => {
    if (!fromIso && !toIso) return undefined;
    return {
      from: fromIso ? new Date(fromIso) : undefined,
      to: toIso ? new Date(toIso) : undefined,
    };
  }, [fromIso, toIso]);
  const onlyActivePlan = params.get("activePlan") === "1";
  // Default alfabético por nome (antes era por data de cadastro desc).
  const sortKey = (params.get("sort") as SortKey | null) ?? "name";
  const sortDir = (params.get("dir") as SortDir | null) ?? "asc";
  const page = Math.max(0, parseInt(params.get("page") ?? "0", 10));

  const debouncedSearch = useDebouncedValue(search, 300);

  // ---- fetch (buffer maior pra suportar filtros client-side) ----------------
  const { data, isLoading } = useQuery({
    queryKey: ["patients", "list", debouncedSearch],
    queryFn: async () => {
      // Com termo: busca server-side (sem teto de buffer, CPF funciona via blind
      // index). Sem termo: lote pra navegação paginada client-side.
      const qs = debouncedSearch
        ? `limit=100&search=${encodeURIComponent(debouncedSearch)}`
        : `limit=${MAX_BUFFER}&offset=0`;
      const result = await apiClient.get<Patient[] | PatientsResponse>(
        `/api/v1/patients?${qs}`,
      );

      if (Array.isArray(result)) {
        return {
          data: result,
          total: result.length,
          page: 1,
          limit: MAX_BUFFER,
        } satisfies PatientsResponse;
      }
      return result;
    },
    staleTime: 30 * 1000,
  });

  // ---- filtragem + ordenação client-side ------------------------------------
  const filtered = useMemo(() => {
    let items = data?.data ?? [];

    // Busca já foi aplicada server-side (ver useQuery). Aqui só os filtros
    // locais de período/plano e a ordenação sobre o conjunto retornado.
    if (dateRange?.from) {
      const fromMs = dateRange.from.getTime();
      items = items.filter((p) => new Date(p.createdAt).getTime() >= fromMs);
    }
    if (dateRange?.to) {
      const toMs = new Date(dateRange.to).setHours(23, 59, 59, 999);
      items = items.filter((p) => new Date(p.createdAt).getTime() <= toMs);
    }

    if (onlyActivePlan) {
      items = items.filter((p) => p.hasActiveSubscription === true);
    }

    const dir = sortDir === "asc" ? 1 : -1;
    items = [...items].sort((a, b) => {
      switch (sortKey) {
        case "name":
          return a.name.localeCompare(b.name, "pt", { sensitivity: "base" }) * dir;
        case "age":
          return (a.age - b.age) * dir;
        case "createdAt":
        default:
          return (
            (new Date(a.createdAt).getTime() - new Date(b.createdAt).getTime()) * dir
          );
      }
    });

    return items;
  }, [
    data?.data,
    dateRange,
    onlyActivePlan,
    sortKey,
    sortDir,
  ]);

  // Paginação client-side
  const total = filtered.length;
  const totalPages = Math.max(1, Math.ceil(total / PAGE_SIZE));
  const safePage = Math.min(page, totalPages - 1);
  const paged = filtered.slice(
    safePage * PAGE_SIZE,
    safePage * PAGE_SIZE + PAGE_SIZE,
  );

  // ---- handlers --------------------------------------------------------------
  function handleSearchChange(value: string) {
    setParams({ q: value || null, page: null });
  }

  function handleDateRangeChange(range: DateRange | undefined) {
    setParams({
      from: range?.from ? range.from.toISOString().slice(0, 10) : null,
      to: range?.to ? range.to.toISOString().slice(0, 10) : null,
      page: null,
    });
  }

  function handleActivePlanChange(next: boolean) {
    setParams({ activePlan: next ? "1" : null, page: null });
  }

  function toggleSort(key: SortKey) {
    if (sortKey === key) {
      setParam("dir", sortDir === "asc" ? "desc" : "asc");
    } else {
      setParams({ sort: key, dir: "asc" });
    }
  }

  const activeFilterCount =
    (search ? 1 : 0) + (dateRange ? 1 : 0) + (onlyActivePlan ? 1 : 0);
  const hasActiveFilters = activeFilterCount > 0;

  function emptyMessage() {
    if (debouncedSearch) {
      return `Nenhum paciente encontrado para "${debouncedSearch}".`;
    }
    if (hasActiveFilters) {
      return "Nenhum paciente encontrado com os filtros aplicados.";
    }
    return "Nenhum paciente cadastrado ainda.";
  }

  const filterBar = (
    <PatientsFilterBar
      dateRange={dateRange}
      onDateRangeChange={handleDateRangeChange}
      onlyActivePlan={onlyActivePlan}
      onOnlyActivePlanChange={handleActivePlanChange}
      hasActiveFilters={hasActiveFilters}
      onClear={clearAll}
    />
  );

  return (
    <div className="container mx-auto py-8 space-y-8">
      <PageHeader
        breadcrumbs={[{ label: "Pacientes" }]}
        title="Pacientes"
        description={`${data?.total ?? 0} pacientes cadastrados no sistema`}
        actions={[
          {
            label: "Novo",
            icon: <Plus className="h-4 w-4" />,
            onClick: () => router.push("/patients/new"),
            variant: "default",
          },
        ]}
      />

      {/* Filtros */}
      <div className="space-y-3">
        <div className="flex flex-col gap-2 sm:flex-row sm:items-center">
          <div className="relative flex-1">
            <Search className="absolute left-3 top-2.5 h-4 w-4 text-muted-foreground" />
            <Input
              placeholder="Buscar por nome, email, telefone ou CPF…"
              value={search}
              onChange={(e) => handleSearchChange(e.target.value)}
              className="pl-9"
              aria-label="Buscar pacientes"
            />
          </div>
          <div className="md:hidden">
            <FiltersMobileSheet activeCount={activeFilterCount}>
              {filterBar}
            </FiltersMobileSheet>
          </div>
        </div>
        <div className="hidden md:block">{filterBar}</div>
      </div>

      <Card className="border-0 shadow-lg">
        <CardHeader>
          <CardTitle className="flex items-center justify-between">
            <span>
              {isLoading
                ? "Carregando…"
                : `${total}${
                    total !== (data?.total ?? 0) ? ` de ${data?.total ?? 0}` : ""
                  } paciente${total === 1 ? "" : "s"}`}
            </span>
            {totalPages > 1 && (
              <div className="flex items-center gap-2">
                <Button
                  variant="outline"
                  size="sm"
                  disabled={safePage === 0}
                  onClick={() => setParam("page", String(Math.max(0, safePage - 1)))}
                  aria-label="Página anterior"
                >
                  <ChevronLeft className="h-4 w-4" />
                </Button>
                <span className="text-sm text-muted-foreground">
                  {safePage + 1} / {totalPages}
                </span>
                <Button
                  variant="outline"
                  size="sm"
                  disabled={safePage >= totalPages - 1}
                  onClick={() => setParam("page", String(safePage + 1))}
                  aria-label="Próxima página"
                >
                  <ChevronRight className="h-4 w-4" />
                </Button>
              </div>
            )}
          </CardTitle>
        </CardHeader>
        <CardContent>
          <div className="rounded-md border">
            <table className="w-full">
              <thead className="bg-muted/50">
                <tr>
                  <SortableHeader
                    label="Nome"
                    active={sortKey === "name"}
                    dir={sortDir}
                    onClick={() => toggleSort("name")}
                  />
                  <SortableHeader
                    label="Idade / Nascimento"
                    active={sortKey === "age"}
                    dir={sortDir}
                    onClick={() => toggleSort("age")}
                  />
                  <th className="px-4 py-3 text-left text-sm font-semibold">Gênero</th>
                  <th className="px-4 py-3 text-left text-sm font-semibold">Telefone</th>
                  <SortableHeader
                    label="Cadastrado em"
                    active={sortKey === "createdAt"}
                    dir={sortDir}
                    onClick={() => toggleSort("createdAt")}
                  />
                  <th className="px-4 py-3 text-left text-sm font-semibold">Ações</th>
                </tr>
              </thead>
              <tbody className="divide-y">
                {isLoading ? (
                  Array.from({ length: 5 }).map((_, i) => (
                    <tr key={i}>
                      {Array.from({ length: 6 }).map((_, j) => (
                        <td key={j} className="px-4 py-3">
                          <Skeleton className="h-6 w-full" />
                        </td>
                      ))}
                    </tr>
                  ))
                ) : paged.length === 0 ? (
                  <tr>
                    <td
                      colSpan={6}
                      className="px-4 py-8 text-center text-muted-foreground"
                    >
                      {emptyMessage()}
                    </td>
                  </tr>
                ) : (
                  paged.map((patient) => (
                    <tr
                      key={patient.id}
                      className="transition-colors hover:bg-muted/50"
                    >
                      <td className="px-4 py-3">
                        <div className="font-medium">{patient.name}</div>
                        <div className="text-xs text-muted-foreground">
                          <span className="font-mono">#{patient.id.slice(-6).toUpperCase()}</span>
                          {patient.email && <span> · {patient.email}</span>}
                        </div>
                      </td>
                      <td className="px-4 py-3">
                        <div className="text-sm">
                          <span className="font-medium">{patient.age}</span>
                          <span className="text-muted-foreground">
                            {" "}({formatDate(patient.birthDate)})
                          </span>
                        </div>
                      </td>
                      <td className="px-4 py-3">
                        <Badge variant="outline">
                          {patient.gender === "male"
                            ? "Masculino"
                            : patient.gender === "female"
                              ? "Feminino"
                              : "Outro"}
                        </Badge>
                      </td>
                      <td className="px-4 py-3">
                        <div className="text-sm">{patient.phone}</div>
                      </td>
                      <td className="px-4 py-3">
                        <div className="text-sm text-muted-foreground">
                          {formatDate(patient.createdAt)}
                        </div>
                      </td>
                      <td className="px-4 py-3">
                        <div className="flex gap-2">
                          <Button
                            variant="ghost"
                            size="sm"
                            onClick={() => router.push(`/patients/${patient.id}`)}
                          >
                            Ver
                          </Button>
                          <Button
                            variant="ghost"
                            size="sm"
                            onClick={() => router.push(`/patients/${patient.id}/edit`)}
                          >
                            Editar
                          </Button>
                        </div>
                      </td>
                    </tr>
                  ))
                )}
              </tbody>
            </table>
          </div>

          <div className="mt-4 flex items-center justify-between text-sm text-muted-foreground">
            <div>
              {total > 0
                ? `Mostrando ${safePage * PAGE_SIZE + 1} até ${Math.min(
                    (safePage + 1) * PAGE_SIZE,
                    total,
                  )} de ${total}`
                : "Nenhum resultado"}
            </div>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}

interface SortableHeaderProps {
  label: string;
  active: boolean;
  dir: SortDir;
  onClick: () => void;
}

function SortableHeader({ label, active, dir, onClick }: SortableHeaderProps) {
  return (
    <th className="px-4 py-3 text-left text-sm font-semibold">
      <button
        type="button"
        onClick={onClick}
        className="inline-flex items-center gap-1 hover:text-foreground"
      >
        {label}
        {active &&
          (dir === "asc" ? (
            <ArrowUp className="h-3 w-3" />
          ) : (
            <ArrowDown className="h-3 w-3" />
          ))}
      </button>
    </th>
  );
}
