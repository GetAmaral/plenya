"use client";

import { forwardRef, useState } from "react";
import { useQuery } from "@tanstack/react-query";
import Link from "next/link";
import { Search, Loader2, UserPlus } from "lucide-react";
import { apiClient } from "@/lib/api-client";
import type { Patient } from "@/lib/api/patient-api";
import { useDebouncedValue } from "@/lib/use-debounced-value";
import { Input } from "@/components/ui/input";
import { Card, CardContent } from "@/components/ui/card";

/**
 * Busca rapida de paciente. Foca um input real na pagina (debounced) e bate
 * no endpoint de busca de pacientes.
 *
 * Derivacao client-side documentada: o hook usePatients() compartilhado nao
 * expoe o parametro ?search, entao chamamos apiClient.get direto no endpoint
 * server-side ja existente: GET /api/v1/patients?search=<q>&limit=8. Nenhum
 * endpoint novo foi inventado.
 */
export const PatientQuickSearch = forwardRef<HTMLInputElement>(
  function PatientQuickSearch(_props, ref) {
    const [term, setTerm] = useState("");
    const debounced = useDebouncedValue(term.trim(), 300);
    const enabled = debounced.length >= 2;

    const { data: results = [], isFetching } = useQuery({
      queryKey: ["recepcao", "patient-search", debounced],
      queryFn: () =>
        apiClient.get<Patient[]>(
          `/api/v1/patients?search=${encodeURIComponent(debounced)}&limit=8`,
        ),
      enabled,
      staleTime: 10_000,
    });

    return (
      <div className="relative">
        <div className="relative">
          <Search className="absolute left-3 top-1/2 -translate-y-1/2 h-4 w-4 text-muted-foreground" />
          <Input
            ref={ref}
            value={term}
            onChange={(e) => setTerm(e.target.value)}
            placeholder="Buscar paciente por nome, telefone ou CPF"
            className="pl-9"
            aria-label="Buscar paciente"
          />
          {isFetching && enabled && (
            <Loader2 className="absolute right-3 top-1/2 -translate-y-1/2 h-4 w-4 animate-spin text-muted-foreground" />
          )}
        </div>

        {enabled && (
          <Card className="absolute z-20 mt-2 w-full border-0 shadow-lg">
            <CardContent className="p-2">
              {results.length === 0 && !isFetching ? (
                <div className="px-3 py-4 text-sm text-muted-foreground">
                  <p>Nenhum paciente encontrado.</p>
                  <Link
                    href="/patients/new"
                    className="mt-2 inline-flex items-center gap-1.5 text-primary hover:underline"
                  >
                    <UserPlus className="h-4 w-4" />
                    Cadastrar novo paciente
                  </Link>
                </div>
              ) : (
                <ul className="max-h-72 overflow-auto">
                  {results.map((p) => (
                    <li key={p.id}>
                      <Link
                        href={`/patients/${p.id}`}
                        className="flex flex-col gap-0.5 rounded-md px-3 py-2 hover:bg-accent transition-colors"
                        onClick={() => setTerm("")}
                      >
                        <span className="font-medium">{p.name}</span>
                        {(p.phone || p.email) && (
                          <span className="text-xs text-muted-foreground">
                            {[p.phone, p.email].filter(Boolean).join(" · ")}
                          </span>
                        )}
                      </Link>
                    </li>
                  ))}
                </ul>
              )}
            </CardContent>
          </Card>
        )}
      </div>
    );
  },
);
