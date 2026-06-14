"use client";

import { useState } from "react";
import { useQuery } from "@tanstack/react-query";
import { UseFormReturn } from "react-hook-form";
import { Trash2, TestTube, Sparkles, Dna } from "lucide-react";
import { labTestDefinitionApi } from "@/lib/api/lab-test-definition-api";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";
import {
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
  FormDescription,
} from "@/components/ui/form";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Separator } from "@/components/ui/separator";
import { LabTestDefinitionSelect } from "./LabTestDefinitionSelect";
import type { LabResultBatchFormValues } from "@/lib/validations/lab-result-batch";
import type { LabTestDefinition } from "@/lib/api/lab-test-definition-api";

interface LabResultInBatchFormProps {
  form: UseFormReturn<LabResultBatchFormValues>;
  index: number;
  onRemove: () => void;
  canRemove: boolean;
}

// Cor do badge por nível (0=pior … 5=melhor), reaproveitada do seletor de nível.
const levelColor = (level: number): string =>
  ["bg-red-100 border-red-500",
   "bg-orange-100 border-orange-500",
   "bg-yellow-100 border-yellow-500",
   "bg-blue-100 border-blue-500",
   "bg-green-100 border-green-500",
   "bg-emerald-100 border-emerald-500"][level] ?? "bg-muted border-border";

export function LabResultInBatchForm({
  form,
  index,
  onRemove,
  canRemove,
}: LabResultInBatchFormProps) {
  // Exame selecionado do catálogo (guarda resultType/code p/ decidir o tipo de entrada).
  const [selectedTest, setSelectedTest] = useState<LabTestDefinition | null>(null);

  // Handler para quando seleciona um exame do catálogo
  const handleTestSelect = (test: LabTestDefinition | null) => {
    setSelectedTest(test);
    if (test) {
      // Auto-preencher campos com dados do catálogo
      form.setValue(`labResults.${index}.labTestDefinitionId`, test.id);
      form.setValue(`labResults.${index}.testName`, test.name);
      form.setValue(`labResults.${index}.testType`, test.category);

      if (test.unit) {
        form.setValue(`labResults.${index}.unit`, test.unit);
      }
    } else {
      // Limpar seleção
      form.setValue(`labResults.${index}.labTestDefinitionId`, undefined);
    }
  };

  // Exame qualitativo (ex.: genótipo): busca as opções (níveis) e deixa selecionar uma,
  // gravando result_text + level de uma vez. Vale p/ qualquer exame categórico com opções.
  const isCategorical = selectedTest?.resultType === "categorical";
  const isGenetic = selectedTest?.category === "genetics";
  const { data: options } = useQuery({
    queryKey: ["qualitative-options", selectedTest?.code],
    queryFn: () => labTestDefinitionApi.getQualitativeOptions(selectedTest!.code),
    enabled: !!selectedTest?.code && isCategorical,
  });
  const hasOptions = isCategorical && (options?.length ?? 0) > 0;
  const currentLevel = form.watch(`labResults.${index}.level`);

  const handleOptionSelect = (levelStr: string) => {
    const opt = options?.find((o) => o.level === parseInt(levelStr));
    if (!opt) return;
    form.setValue(`labResults.${index}.level`, opt.level);
    form.setValue(`labResults.${index}.resultText`, opt.name);
    form.setValue(`labResults.${index}.resultNumeric`, undefined);
  };

  return (
    <Card>
      <CardHeader className="pb-3">
        <div className="flex items-center justify-between">
          <CardTitle className="text-base flex items-center gap-2">
            <TestTube className="h-4 w-4" />
            Resultado {index + 1}
          </CardTitle>
          {canRemove && (
            <Button
              type="button"
              variant="ghost"
              size="sm"
              onClick={onRemove}
              className="h-8 w-8 p-0 text-destructive hover:text-destructive"
            >
              <Trash2 className="h-4 w-4" />
            </Button>
          )}
        </div>
      </CardHeader>

      <CardContent className="pt-4 space-y-4">
        {/* Seletor de Exame do Catálogo */}
        <FormField
          control={form.control}
          name={`labResults.${index}.labTestDefinitionId`}
          render={({ field }) => (
            <FormItem>
              <FormLabel className="flex items-center gap-2">
                <Sparkles className="h-4 w-4" />
                Buscar no Catálogo de Exames
              </FormLabel>
              <FormControl>
                <LabTestDefinitionSelect
                  value={field.value}
                  onSelect={handleTestSelect}
                  showAll={true}
                />
              </FormControl>
              <FormDescription>
                Selecione do catálogo para auto-preencher ou preencha manualmente abaixo
              </FormDescription>
              <FormMessage />
            </FormItem>
          )}
        />

        <Separator className="my-2" />

        {/* Nome e Tipo do Exame */}
        <div className="grid grid-cols-2 gap-4">
          <FormField
            control={form.control}
            name={`labResults.${index}.testName`}
            render={({ field }) => (
              <FormItem>
                <FormLabel>Nome do Exame *</FormLabel>
                <FormControl>
                  <Input
                    {...field}
                    placeholder="Ex: Hemoglobina"
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name={`labResults.${index}.testType`}
            render={({ field }) => (
              <FormItem>
                <FormLabel>Tipo *</FormLabel>
                <FormControl>
                  <Input
                    {...field}
                    placeholder="Ex: Hematologia"
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
        </div>

        {/* Status é do lote (LabResultBatch), não por-resultado — campo removido (era órfão:
            não existia no schema/DTO nem era persistido). */}

        <Separator />

        {hasOptions ? (
          /* Exame qualitativo (genótipo/categórico): uma única seleção define result_text + level */
          <FormField
            control={form.control}
            name={`labResults.${index}.level`}
            render={() => (
              <FormItem>
                <FormLabel className="flex items-center gap-2">
                  <Dna className="h-4 w-4" />
                  {isGenetic ? "Genótipo" : "Resultado"} *
                </FormLabel>
                <Select
                  onValueChange={handleOptionSelect}
                  value={currentLevel != null ? String(currentLevel) : ""}
                >
                  <FormControl>
                    <SelectTrigger>
                      <SelectValue
                        placeholder={isGenetic ? "Selecione o genótipo" : "Selecione o resultado"}
                      />
                    </SelectTrigger>
                  </FormControl>
                  <SelectContent>
                    {options!.map((opt) => (
                      <SelectItem key={opt.id} value={String(opt.level)}>
                        <div className="flex items-center gap-2">
                          <div className={`w-3 h-3 rounded border ${levelColor(opt.level)}`} />
                          <span>{opt.name}</span>
                        </div>
                      </SelectItem>
                    ))}
                  </SelectContent>
                </Select>
                <FormDescription>
                  A opção define o resultado e o nível do escore (maior = melhor).
                </FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />
        ) : (
        <>
        {/* Resultado - Texto OU Numérico */}
        <div className="space-y-4">
          <div className="flex items-center justify-between">
            <FormLabel>Resultado</FormLabel>
            <span className="text-xs text-muted-foreground">
              Preencha valor numérico OU textual
            </span>
          </div>

          <div className="grid grid-cols-2 gap-4">
            <FormField
              control={form.control}
              name={`labResults.${index}.resultNumeric`}
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Valor Numérico</FormLabel>
                  <FormControl>
                    <Input
                      type="number"
                      step="any"
                      placeholder="0.00"
                      {...field}
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

            <FormField
              control={form.control}
              name={`labResults.${index}.unit`}
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Unidade</FormLabel>
                  <FormControl>
                    <Input
                      {...field}
                      placeholder="Ex: mg/dL"
                      value={field.value || ""}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
          </div>

          <FormField
            control={form.control}
            name={`labResults.${index}.resultText`}
            render={({ field }) => (
              <FormItem>
                <FormLabel>Resultado Textual</FormLabel>
                <FormControl>
                  <Input
                    {...field}
                    placeholder="Ex: Negativo, Positivo"
                    value={field.value || ""}
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
        </div>

        {/* Nível */}
        <FormField
          control={form.control}
          name={`labResults.${index}.level`}
          render={({ field }) => (
            <FormItem>
              <FormLabel>Nível</FormLabel>
              <Select
                onValueChange={(value) =>
                  field.onChange(value ? parseInt(value) : undefined)
                }
                value={field.value?.toString() || ""}
              >
                <FormControl>
                  <SelectTrigger>
                    <SelectValue placeholder="Selecione o nível" />
                  </SelectTrigger>
                </FormControl>
                <SelectContent>
                  <SelectItem value="0">
                    <div className="flex items-center gap-2">
                      <div className="w-3 h-3 rounded bg-red-100 border border-red-500" />
                      <span className="font-semibold">N0</span>
                    </div>
                  </SelectItem>
                  <SelectItem value="1">
                    <div className="flex items-center gap-2">
                      <div className="w-3 h-3 rounded bg-orange-100 border border-orange-500" />
                      <span className="font-semibold">N1</span>
                    </div>
                  </SelectItem>
                  <SelectItem value="2">
                    <div className="flex items-center gap-2">
                      <div className="w-3 h-3 rounded bg-yellow-100 border border-yellow-500" />
                      <span className="font-semibold">N2</span>
                    </div>
                  </SelectItem>
                  <SelectItem value="3">
                    <div className="flex items-center gap-2">
                      <div className="w-3 h-3 rounded bg-blue-100 border border-blue-500" />
                      <span className="font-semibold">N3</span>
                    </div>
                  </SelectItem>
                  <SelectItem value="4">
                    <div className="flex items-center gap-2">
                      <div className="w-3 h-3 rounded bg-green-100 border border-green-500" />
                      <span className="font-semibold">N4</span>
                    </div>
                  </SelectItem>
                  <SelectItem value="5">
                    <div className="flex items-center gap-2">
                      <div className="w-3 h-3 rounded bg-emerald-100 border border-emerald-500" />
                      <span className="font-semibold">N5</span>
                    </div>
                  </SelectItem>
                </SelectContent>
              </Select>
              <FormMessage />
            </FormItem>
          )}
        />
        </>
        )}

        {/* Interpretação */}
        <FormField
          control={form.control}
          name={`labResults.${index}.interpretation`}
          render={({ field }) => (
            <FormItem>
              <FormLabel>Interpretação</FormLabel>
              <FormControl>
                <Textarea
                  {...field}
                  placeholder="Observações sobre o resultado..."
                  rows={2}
                  value={field.value || ""}
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
      </CardContent>
    </Card>
  );
}
