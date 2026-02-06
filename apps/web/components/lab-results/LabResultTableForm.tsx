"use client";

import { useState, useEffect, useRef } from "react";
import { UseFormReturn } from "react-hook-form";
import { Trash2, Plus, MoreHorizontal } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import {
  FormControl,
  FormField,
  FormItem,
  FormMessage,
} from "@/components/ui/form";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { LabTestDefinitionSelect } from "./LabTestDefinitionSelect";
import { UnmatchedBadge } from "./UnmatchedBadge";
import type { LabResultBatchFormValues } from "@/lib/validations/lab-result-batch";
import type { LabTestDefinition } from "@/lib/api/lab-test-definition-api";

interface LabResultTableFormProps {
  form: UseFormReturn<LabResultBatchFormValues>;
  fields: any[];
  append: (value: any) => void;
  remove: (index: number) => void;
  getDefaultResultValues: () => any;
}

const levelColors = {
  0: "bg-red-100 border-red-500",
  1: "bg-orange-100 border-orange-500",
  2: "bg-yellow-100 border-yellow-500",
  3: "bg-blue-100 border-blue-500",
  4: "bg-green-100 border-green-500",
  5: "bg-emerald-100 border-emerald-500",
};

export function LabResultTableForm({
  form,
  fields,
  append,
  remove,
  getDefaultResultValues,
}: LabResultTableFormProps) {
  // Estado para controlar se cada linha está em modo manual (true) ou seleção (false)
  const [manualModeRows, setManualModeRows] = useState<Record<number, boolean>>({});
  // Estado para rastrear qual linha acabou de ser adicionada (para auto-focus)
  const [lastAddedIndex, setLastAddedIndex] = useState<number | null>(null);
  // Refs para os inputs de resultado (para focar após seleção)
  const resultInputRefs = useRef<(HTMLInputElement | null)[]>([]);

  const handleAddResult = () => {
    append(getDefaultResultValues());
    // Define o índice da nova linha (será o último)
    setLastAddedIndex(fields.length);
    // Limpa após um pequeno delay
    setTimeout(() => setLastAddedIndex(null), 100);
  };

  const toggleManualMode = (index: number) => {
    setManualModeRows(prev => ({
      ...prev,
      [index]: !prev[index]
    }));
  };

  // Garantir que sempre há pelo menos uma linha vazia
  useEffect(() => {
    if (fields.length === 0) {
      append(getDefaultResultValues());
    }
  }, [fields.length, append, getDefaultResultValues]);

  const handleTestSelect = (index: number, test: LabTestDefinition | null) => {
    if (test) {
      form.setValue(`labResults.${index}.labTestDefinitionId`, test.id);
      form.setValue(`labResults.${index}.testName`, test.name);
      form.setValue(`labResults.${index}.testType`, test.category);
      if (test.unit) {
        form.setValue(`labResults.${index}.unit`, test.unit);
      }

      // Focar no campo de resultado após seleção
      setTimeout(() => {
        resultInputRefs.current[index]?.focus();
      }, 100);
    } else {
      form.setValue(`labResults.${index}.labTestDefinitionId`, undefined);
    }
  };

  return (
    <div className="space-y-3">
      <div className="flex justify-between items-center">
        <h3 className="text-base font-semibold">
          Resultados ({fields.length})
        </h3>
        <Button
          type="button"
          variant="outline"
          size="sm"
          onClick={handleAddResult}
          className="gap-1.5 h-8 text-sm"
        >
          <Plus className="h-3.5 w-3.5" />
          Adicionar Linha
        </Button>
      </div>

      <div className="rounded-md border overflow-x-auto max-h-[600px] overflow-y-auto">
          <Table>
            <TableHeader className="sticky top-0 bg-background z-10 shadow-sm">
              <TableRow className="hover:bg-transparent border-b-2">
                <TableHead className="w-[35px] px-2 py-2 text-xs">#</TableHead>
                <TableHead className="min-w-[280px] px-2 py-2 text-xs">Exame / Nome / Tipo</TableHead>
                <TableHead className="min-w-[80px] px-2 py-2 text-xs">Valor</TableHead>
                <TableHead className="min-w-[70px] px-2 py-2 text-xs">Unid.</TableHead>
                <TableHead className="min-w-[200px] px-2 py-2 text-xs">Texto</TableHead>
                <TableHead className="min-w-[85px] px-2 py-2 text-xs">Nível</TableHead>
                <TableHead className="min-w-[160px] px-2 py-2 text-xs">Interpretação</TableHead>
                <TableHead className="w-[50px] px-2 py-2 text-xs">Ações</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {fields.map((field, index) => (
                <TableRow key={field.id} className="even:bg-muted/30">
                  {/* Número */}
                  <TableCell className="px-2 py-1.5 text-center text-muted-foreground font-mono text-xs">
                    {index + 1}
                  </TableCell>

                  {/* Exame / Nome / Tipo (com toggle) */}
                  <TableCell className="px-2 py-1.5">
                    <div className="flex items-start gap-1.5">
                      <div className="flex-1 space-y-1.5">
                        {!manualModeRows[index] ? (
                          // Modo Seleção: LabTestDefinitionSelect
                          <FormField
                            control={form.control}
                            name={`labResults.${index}.labTestDefinitionId`}
                            render={({ field }) => (
                              <FormItem>
                                <FormControl>
                                  <LabTestDefinitionSelect
                                    value={field.value}
                                    onSelect={(test) => handleTestSelect(index, test)}
                                    autoFocus={index === lastAddedIndex}
                                  />
                                </FormControl>
                                <FormMessage />
                              </FormItem>
                            )}
                          />
                        ) : (
                          // Modo Manual: Nome + Tipo
                          <>
                            <div className="flex items-center gap-2">
                              <FormField
                                control={form.control}
                                name={`labResults.${index}.testName`}
                                render={({ field }) => (
                                  <FormItem className="flex-1">
                                    <FormControl>
                                      <Input
                                        {...field}
                                        placeholder="Nome do exame"
                                        className="h-8 text-sm"
                                      />
                                    </FormControl>
                                    <FormMessage />
                                  </FormItem>
                                )}
                              />
                              {form.watch(`labResults.${index}.matched`) === false && (
                                <UnmatchedBadge />
                              )}
                            </div>
                            <FormField
                              control={form.control}
                              name={`labResults.${index}.testType`}
                              render={({ field }) => (
                                <FormItem>
                                  <FormControl>
                                    <Input
                                      {...field}
                                      placeholder="Tipo/Categoria"
                                      className="h-8 text-sm"
                                    />
                                  </FormControl>
                                  <FormMessage />
                                </FormItem>
                              )}
                            />
                          </>
                        )}
                      </div>
                      <Button
                        type="button"
                        variant="ghost"
                        size="sm"
                        onClick={() => toggleManualMode(index)}
                        className="h-8 w-8 p-0 mt-0 flex-shrink-0"
                        title={manualModeRows[index] ? "Voltar para busca no catálogo" : "Preencher manualmente"}
                      >
                        <MoreHorizontal className="h-4 w-4" />
                      </Button>
                    </div>
                  </TableCell>

                  {/* Resultado Numérico */}
                  <TableCell className="px-2 py-1.5">
                    <FormField
                      control={form.control}
                      name={`labResults.${index}.resultNumeric`}
                      render={({ field }) => (
                        <FormItem>
                          <FormControl>
                            <Input
                              type="number"
                              step="any"
                              placeholder="0.00"
                              className="h-8 text-sm"
                              {...field}
                              ref={(el) => {
                                resultInputRefs.current[index] = el;
                              }}
                              onChange={(e) => {
                                const value = e.target.value;
                                field.onChange(value === "" ? undefined : parseFloat(value));
                              }}
                              value={field.value ?? ""}
                            />
                          </FormControl>
                          <FormMessage />
                        </FormItem>
                      )}
                    />
                  </TableCell>

                  {/* Unidade */}
                  <TableCell className="px-2 py-1.5">
                    <FormField
                      control={form.control}
                      name={`labResults.${index}.unit`}
                      render={({ field }) => (
                        <FormItem>
                          <FormControl>
                            <Input
                              {...field}
                              placeholder="mg/dL"
                              className="h-8 text-sm"
                              value={field.value || ""}
                            />
                          </FormControl>
                          <FormMessage />
                        </FormItem>
                      )}
                    />
                  </TableCell>

                  {/* Resultado Textual */}
                  <TableCell className="px-2 py-1.5">
                    <FormField
                      control={form.control}
                      name={`labResults.${index}.resultText`}
                      render={({ field }) => (
                        <FormItem>
                          <FormControl>
                            <Input
                              {...field}
                              placeholder="Negativo"
                              className="h-8 text-sm"
                              value={field.value || ""}
                            />
                          </FormControl>
                          <FormMessage />
                        </FormItem>
                      )}
                    />
                  </TableCell>

                  {/* Nível */}
                  <TableCell className="px-2 py-1.5">
                    <FormField
                      control={form.control}
                      name={`labResults.${index}.level`}
                      render={({ field }) => (
                        <FormItem>
                          <Select
                            onValueChange={(value) =>
                              field.onChange(value ? parseInt(value) : undefined)
                            }
                            value={field.value?.toString() || ""}
                          >
                            <FormControl>
                              <SelectTrigger className="h-8 text-sm">
                                <SelectValue placeholder="Nível" />
                              </SelectTrigger>
                            </FormControl>
                            <SelectContent>
                              {[0, 1, 2, 3, 4, 5].map((level) => (
                                <SelectItem key={level} value={level.toString()}>
                                  <div className="flex items-center gap-1.5">
                                    <div className={`w-2.5 h-2.5 rounded border ${levelColors[level as keyof typeof levelColors]}`} />
                                    <span className="text-xs">N{level}</span>
                                  </div>
                                </SelectItem>
                              ))}
                            </SelectContent>
                          </Select>
                          <FormMessage />
                        </FormItem>
                      )}
                    />
                  </TableCell>

                  {/* Interpretação */}
                  <TableCell className="px-2 py-1.5">
                    <FormField
                      control={form.control}
                      name={`labResults.${index}.interpretation`}
                      render={({ field }) => (
                        <FormItem>
                          <FormControl>
                            <Input
                              {...field}
                              placeholder="Observações..."
                              className="h-8 text-sm"
                              value={field.value || ""}
                            />
                          </FormControl>
                          <FormMessage />
                        </FormItem>
                      )}
                    />
                  </TableCell>

                  {/* Ações */}
                  <TableCell className="px-2 py-1.5">
                    <Button
                      type="button"
                      variant="ghost"
                      size="sm"
                      onClick={() => remove(index)}
                      className="h-7 w-7 p-0 text-destructive hover:text-destructive hover:bg-destructive/10"
                    >
                      <Trash2 className="h-3.5 w-3.5" />
                    </Button>
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </div>
    </div>
  );
}
