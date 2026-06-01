"use client";

import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { useQuery } from "@tanstack/react-query";
import {
  Command,
  CommandInput,
  CommandList,
  CommandEmpty,
  CommandGroup,
  CommandItem,
} from "@/components/ui/command";
import { Dialog, DialogContent } from "@/components/ui/dialog";
import { apiClient } from "@/lib/api-client";
import type { Patient } from "@/lib/api/patient-api";

// Busca global de pacientes (nome / CPF / telefone). Atalho Cmd/Ctrl+K,
// disponível de qualquer tela do EMR.
export function GlobalSearch() {
  const router = useRouter();
  const [open, setOpen] = useState(false);
  const [term, setTerm] = useState("");
  const [debounced, setDebounced] = useState("");

  // Atalho global Cmd/Ctrl+K.
  useEffect(() => {
    const onKey = (e: KeyboardEvent) => {
      if ((e.metaKey || e.ctrlKey) && e.key.toLowerCase() === "k") {
        e.preventDefault();
        setOpen((v) => !v);
      }
    };
    document.addEventListener("keydown", onKey);
    return () => document.removeEventListener("keydown", onKey);
  }, []);

  // Debounce 300ms.
  useEffect(() => {
    const t = setTimeout(() => setDebounced(term.trim()), 300);
    return () => clearTimeout(t);
  }, [term]);

  const { data, isFetching } = useQuery({
    queryKey: ["global-search", debounced],
    queryFn: () =>
      apiClient.get<Patient[] | { patients: Patient[] }>(
        `/api/v1/patients?limit=20&search=${encodeURIComponent(debounced)}`,
      ),
    enabled: open && debounced.length >= 2,
    staleTime: 10_000,
  });

  const results: Patient[] = Array.isArray(data)
    ? data
    : (data?.patients ?? []);

  const goToPatient = (id: string) => {
    setOpen(false);
    setTerm("");
    router.push(`/patients/${id}`);
  };

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogContent className="overflow-hidden p-0">
        {/* shouldFilter=false: filtragem é server-side (cobre CPF, que não está
            no texto do item). */}
        <Command shouldFilter={false}>
          <CommandInput
            placeholder="Buscar paciente por nome, CPF ou telefone..."
            value={term}
            onValueChange={setTerm}
          />
          <CommandList>
        {debounced.length < 2 ? (
          <CommandEmpty>Digite ao menos 2 caracteres.</CommandEmpty>
        ) : isFetching ? (
          <CommandEmpty>Buscando...</CommandEmpty>
        ) : results.length === 0 ? (
          <CommandEmpty>Nenhum paciente encontrado.</CommandEmpty>
        ) : (
          <CommandGroup heading="Pacientes">
            {results.map((p) => (
              <CommandItem
                key={p.id}
                value={`${p.id} ${p.name} ${p.phone ?? ""}`}
                onSelect={() => goToPatient(p.id)}
              >
                <div className="flex flex-col">
                  <span>{p.name}</span>
                  {p.phone && (
                    <span className="text-xs text-muted-foreground">
                      {p.phone}
                    </span>
                  )}
                </div>
              </CommandItem>
            ))}
            </CommandGroup>
          )}
          </CommandList>
        </Command>
      </DialogContent>
    </Dialog>
  );
}

export default GlobalSearch;
