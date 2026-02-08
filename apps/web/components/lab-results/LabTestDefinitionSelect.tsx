"use client";

import { useState, useMemo, useEffect } from "react";
import { useQuery } from "@tanstack/react-query";
import { Check, ChevronsUpDown, Search, Loader2, TestTube } from "lucide-react";
import { Button } from "@/components/ui/button";
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
} from "@/components/ui/command";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { Badge } from "@/components/ui/badge";
import { cn } from "@/lib/utils";
import { labTestDefinitionApi, type LabTestDefinition } from "@/lib/api/lab-test-definition-api";

// Remove accents from string (unaccent)
function removeAccents(str: string): string {
  return str.normalize('NFD').replace(/[\u0300-\u036f]/g, '')
}

// Normalize string for search (lowercase + remove accents)
function normalizeForSearch(str: string): string {
  return removeAccents(str.toLowerCase())
}

interface LabTestDefinitionSelectProps {
  value?: string;
  onSelect: (test: LabTestDefinition | null) => void;
  disabled?: boolean;
  autoFocus?: boolean;
  showAll?: boolean; // Se true, mostra todos os exames (não apenas requestable)
}

const categoryLabels: Record<string, string> = {
  hematology: "Hematologia",
  biochemistry: "Bioquímica",
  hormones: "Hormônios",
  immunology: "Imunologia",
  microbiology: "Microbiologia",
  urine: "Urinálise",
  imaging: "Imagem",
  functional: "Funcional",
  genetics: "Genética",
  other: "Outros",
};

const categoryColors: Record<string, string> = {
  hematology: "bg-red-100 text-red-700",
  biochemistry: "bg-blue-100 text-blue-700",
  hormones: "bg-purple-100 text-purple-700",
  immunology: "bg-green-100 text-green-700",
  microbiology: "bg-yellow-100 text-yellow-700",
  urine: "bg-orange-100 text-orange-700",
  imaging: "bg-cyan-100 text-cyan-700",
  functional: "bg-pink-100 text-pink-700",
  genetics: "bg-indigo-100 text-indigo-700",
  other: "bg-gray-100 text-gray-700",
};

export function LabTestDefinitionSelect({
  value,
  onSelect,
  disabled,
  autoFocus = false,
  showAll = false,
}: LabTestDefinitionSelectProps) {
  const [open, setOpen] = useState(false);
  const [search, setSearch] = useState("");

  // Auto-abrir o popover quando autoFocus é true
  useEffect(() => {
    if (autoFocus && !open) {
      // Pequeno delay para garantir que o componente está montado
      const timer = setTimeout(() => setOpen(true), 100);
      return () => clearTimeout(timer);
    }
  }, [autoFocus]);

  const { data: tests, isLoading } = useQuery({
    queryKey: showAll ? ["lab-test-definitions-all"] : ["lab-test-definitions-requestable"],
    queryFn: () => showAll ? labTestDefinitionApi.getAll() : labTestDefinitionApi.getRequestable(),
  });

  // Agrupar por categoria
  const groupedTests = useMemo(() => {
    if (!tests) return new Map<string, LabTestDefinition[]>();

    const grouped = new Map<string, LabTestDefinition[]>();

    tests.forEach((test) => {
      if (!grouped.has(test.category)) {
        grouped.set(test.category, []);
      }
      grouped.get(test.category)!.push(test);
    });

    // Ordenar cada grupo por nome (alfabético pt-BR)
    grouped.forEach((group) => {
      group.sort((a, b) => a.name.localeCompare(b.name, 'pt-BR'));
    });

    return grouped;
  }, [tests]);

  // Filtrar por busca (case-insensitive, unaccent, search in name, shortName, altNames)
  const filteredGroups = useMemo(() => {
    if (!search.trim()) return groupedTests;

    const filtered = new Map<string, LabTestDefinition[]>();
    const searchNormalized = normalizeForSearch(search);

    groupedTests.forEach((tests, category) => {
      const matchingTests = tests.filter((test) => {
        // Search in name
        if (normalizeForSearch(test.name).includes(searchNormalized)) return true;
        // Search in shortName
        if (test.shortName && normalizeForSearch(test.shortName).includes(searchNormalized)) return true;
        // Search in altNames
        if (test.altNames && test.altNames.some(alt => normalizeForSearch(alt).includes(searchNormalized))) return true;
        return false;
      });

      if (matchingTests.length > 0) {
        filtered.set(category, matchingTests);
      }
    });

    return filtered;
  }, [groupedTests, search]);

  const selectedTest = useMemo(() => {
    if (!value || !tests) return null;
    return tests.find((t) => t.id === value) || null;
  }, [value, tests]);

  const handleSelect = (test: LabTestDefinition) => {
    onSelect(test);
    setOpen(false);
  };

  const handleClear = () => {
    onSelect(null);
    setOpen(false);
  };

  return (
    <Popover open={open} onOpenChange={setOpen}>
      <PopoverTrigger asChild>
        <Button
          variant="outline"
          role="combobox"
          aria-expanded={open}
          className={cn(
            "w-full justify-between",
            !selectedTest && "text-muted-foreground"
          )}
          disabled={disabled}
        >
          <div className="flex items-center gap-2 truncate">
            {selectedTest ? (
              <>
                <TestTube className="h-4 w-4 shrink-0" />
                <span className="truncate">{selectedTest.name}</span>
                {selectedTest.shortName && (
                  <span className="text-xs text-muted-foreground">
                    ({selectedTest.shortName})
                  </span>
                )}
              </>
            ) : (
              <>
                <Search className="h-4 w-4 shrink-0" />
                <span>Selecionar exame do catálogo...</span>
              </>
            )}
          </div>
          <ChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
        </Button>
      </PopoverTrigger>
      <PopoverContent className="w-[500px] p-0" align="start">
        <Command shouldFilter={false}>
          <CommandInput
            placeholder="Buscar por nome, sigla ou sinônimo..."
            value={search}
            onValueChange={setSearch}
          />
          <CommandList>
            {isLoading ? (
              <div className="flex items-center justify-center py-6">
                <Loader2 className="h-6 w-6 animate-spin text-muted-foreground" />
              </div>
            ) : filteredGroups.size === 0 ? (
              <CommandEmpty>
                {search ? "Nenhum exame encontrado." : "Nenhum exame disponível."}
              </CommandEmpty>
            ) : (
              <>
                {/* Opção para limpar seleção */}
                {selectedTest && (
                  <CommandGroup heading="Ação">
                    <CommandItem
                      onSelect={handleClear}
                      className="text-destructive"
                    >
                      Limpar seleção e preencher manualmente
                    </CommandItem>
                  </CommandGroup>
                )}

                {/* Grupos por categoria */}
                {Array.from(filteredGroups.entries()).map(([category, tests]) => (
                  <CommandGroup
                    key={category}
                    heading={
                      <div className="flex items-center gap-2">
                        <Badge
                          variant="secondary"
                          className={cn(
                            "text-xs",
                            categoryColors[category] || "bg-gray-100"
                          )}
                        >
                          {categoryLabels[category] || category}
                        </Badge>
                        <span className="text-xs text-muted-foreground">
                          {tests.length} {tests.length === 1 ? "exame" : "exames"}
                        </span>
                      </div>
                    }
                  >
                    {tests.map((test) => (
                      <CommandItem
                        key={test.id}
                        value={test.id}
                        onSelect={() => handleSelect(test)}
                        className="flex items-center justify-between"
                      >
                        <div className="flex items-center gap-2 flex-1 min-w-0">
                          <Check
                            className={cn(
                              "h-4 w-4 shrink-0",
                              selectedTest?.id === test.id
                                ? "opacity-100"
                                : "opacity-0"
                            )}
                          />
                          <div className="flex items-center gap-1 min-w-0">
                            <span className="font-medium truncate">
                              {test.name}
                            </span>
                            {(test.shortName || test.unit) && (
                              <span className="text-xs text-muted-foreground whitespace-nowrap">
                                ({[test.shortName, test.unit].filter(Boolean).join(' | ')})
                              </span>
                            )}
                          </div>
                        </div>
                      </CommandItem>
                    ))}
                  </CommandGroup>
                ))}
              </>
            )}
          </CommandList>
        </Command>
      </PopoverContent>
    </Popover>
  );
}
