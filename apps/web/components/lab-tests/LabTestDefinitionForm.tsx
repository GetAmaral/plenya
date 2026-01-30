"use client";

import { useEffect, useState, useRef } from "react";
import { useForm } from "react-hook-form";
import { useFormNavigation } from "@/lib/use-form-navigation";
import { zodResolver } from "@hookform/resolvers/zod";
import { Check, ChevronsUpDown, Loader2 } from "lucide-react";
import { cn } from "@/lib/utils";
import {
  labTestDefinitionSchema,
  type LabTestDefinitionFormValues,
  getDefaultValues,
  apiToFormValues,
  formToApiValues,
  categoryOptions,
  resultTypeOptions,
} from "@/lib/validations/lab-test-definition";
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Textarea } from "@/components/ui/textarea";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
} from "@/components/ui/command";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { Switch } from "@/components/ui/switch";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import type { LabTestDefinition } from "@/lib/api/lab-test-api";

interface LabTestDefinitionFormProps {
  mode: "create" | "edit";
  initialData?: LabTestDefinition;
  availableTests: LabTestDefinition[];
  onSubmit: (values: LabTestDefinitionFormValues) => Promise<void>;
  onCancel: () => void;
  isSubmitting?: boolean;
}

export function LabTestDefinitionForm({
  mode,
  initialData,
  availableTests,
  onSubmit,
  onCancel,
  isSubmitting = false,
}: LabTestDefinitionFormProps) {
  const [parentOpen, setParentOpen] = useState(false);

  const formRef = useRef<HTMLFormElement>(null);
  useFormNavigation({ formRef });

  const form = useForm<LabTestDefinitionFormValues>({
    resolver: zodResolver(labTestDefinitionSchema),
    defaultValues: initialData
      ? apiToFormValues(initialData)
      : getDefaultValues(),
  });

  // Filter requestable tests for parent selection
  const requestableTests = availableTests.filter((t) => t.isRequestable);

  // Watch isRequestable to show/hide parent field
  const isRequestable = form.watch("isRequestable");

  const handleSubmit = async (values: LabTestDefinitionFormValues) => {
    const apiValues = formToApiValues(values);
    await onSubmit(apiValues as LabTestDefinitionFormValues);
  };

  return (
    <Form {...form}>
      <form ref={formRef} onSubmit={form.handleSubmit(handleSubmit)} className="space-y-6">
        <Tabs defaultValue="basic" className="w-full">
          <TabsList className="grid w-full grid-cols-4">
            <TabsTrigger value="basic">Básico</TabsTrigger>
            <TabsTrigger value="technical">Técnico</TabsTrigger>
            <TabsTrigger value="collection">Coleta</TabsTrigger>
            <TabsTrigger value="documentation">Documentação</TabsTrigger>
          </TabsList>

          {/* Basic Information Tab */}
          <TabsContent value="basic" className="space-y-4 mt-4">
            <div className="grid grid-cols-2 gap-4">
              <FormField
                control={form.control}
                name="code"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Código *</FormLabel>
                    <FormControl>
                      <Input
                        placeholder="HEMOGRAMA_COMPLETO"
                        className="font-mono"
                        {...field}
                      />
                    </FormControl>
                    <FormDescription>
                      Código único (A-Z, 0-9, _)
                    </FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />

              <FormField
                control={form.control}
                name="name"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Nome do Exame *</FormLabel>
                    <FormControl>
                      <Input placeholder="Hemograma Completo" {...field} />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
            </div>

            <div className="grid grid-cols-2 gap-4">
              <FormField
                control={form.control}
                name="shortName"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Nome Curto</FormLabel>
                    <FormControl>
                      <Input placeholder="Hemograma" {...field} />
                    </FormControl>
                    <FormDescription>Abreviação (opcional)</FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />

              <FormField
                control={form.control}
                name="category"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Categoria *</FormLabel>
                    <Select
                      onValueChange={field.onChange}
                      defaultValue={field.value}
                    >
                      <FormControl>
                        <SelectTrigger>
                          <SelectValue placeholder="Selecione uma categoria" />
                        </SelectTrigger>
                      </FormControl>
                      <SelectContent>
                        {categoryOptions.map((option) => (
                          <SelectItem key={option.value} value={option.value}>
                            {option.label}
                          </SelectItem>
                        ))}
                      </SelectContent>
                    </Select>
                    <FormMessage />
                  </FormItem>
                )}
              />
            </div>

            <div className="grid grid-cols-2 gap-4">
              <FormField
                control={form.control}
                name="isRequestable"
                render={({ field }) => (
                  <FormItem className="flex flex-row items-center justify-between rounded-lg border p-4">
                    <div className="space-y-0.5">
                      <FormLabel className="text-base">Solicitável</FormLabel>
                      <FormDescription>
                        Pode ser solicitado individualmente
                      </FormDescription>
                    </div>
                    <FormControl>
                      <Switch
                        checked={field.value}
                        onCheckedChange={field.onChange}
                      />
                    </FormControl>
                  </FormItem>
                )}
              />

              <FormField
                control={form.control}
                name="isActive"
                render={({ field }) => (
                  <FormItem className="flex flex-row items-center justify-between rounded-lg border p-4">
                    <div className="space-y-0.5">
                      <FormLabel className="text-base">Ativo</FormLabel>
                      <FormDescription>
                        Exame disponível no sistema
                      </FormDescription>
                    </div>
                    <FormControl>
                      <Switch
                        checked={field.value}
                        onCheckedChange={field.onChange}
                      />
                    </FormControl>
                  </FormItem>
                )}
              />
            </div>

            {!isRequestable && (
              <FormField
                control={form.control}
                name="parentTestId"
                render={({ field }) => (
                  <FormItem className="flex flex-col">
                    <FormLabel>Exame Pai</FormLabel>
                    <Popover open={parentOpen} onOpenChange={setParentOpen}>
                      <PopoverTrigger asChild>
                        <FormControl>
                          <Button
                            variant="outline"
                            role="combobox"
                            className={cn(
                              "justify-between",
                              !field.value && "text-muted-foreground"
                            )}
                          >
                            {field.value
                              ? requestableTests.find((t) => t.id === field.value)
                                  ?.name
                              : "Selecione o exame pai"}
                            <ChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
                          </Button>
                        </FormControl>
                      </PopoverTrigger>
                      <PopoverContent className="w-[400px] p-0">
                        <Command>
                          <CommandInput placeholder="Buscar exame..." />
                          <CommandEmpty>Nenhum exame encontrado.</CommandEmpty>
                          <CommandGroup className="max-h-64 overflow-auto">
                            {requestableTests.map((test) => (
                              <CommandItem
                                key={test.id}
                                value={test.name}
                                onSelect={() => {
                                  form.setValue("parentTestId", test.id);
                                  setParentOpen(false);
                                }}
                              >
                                <Check
                                  className={cn(
                                    "mr-2 h-4 w-4",
                                    test.id === field.value
                                      ? "opacity-100"
                                      : "opacity-0"
                                  )}
                                />
                                <div>
                                  <div className="font-medium">{test.name}</div>
                                  <div className="text-xs text-muted-foreground">
                                    {test.code}
                                  </div>
                                </div>
                              </CommandItem>
                            ))}
                          </CommandGroup>
                        </Command>
                      </PopoverContent>
                    </Popover>
                    <FormDescription>
                      Parâmetros não solicitáveis devem ter um exame pai
                    </FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />
            )}
          </TabsContent>

          {/* Technical Information Tab */}
          <TabsContent value="technical" className="space-y-4 mt-4">
            <div className="grid grid-cols-2 gap-4">
              <FormField
                control={form.control}
                name="tussCode"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Código TUSS</FormLabel>
                    <FormControl>
                      <Input placeholder="40304485" {...field} />
                    </FormControl>
                    <FormDescription>
                      Código para faturamento no Brasil
                    </FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />

              <FormField
                control={form.control}
                name="loincCode"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Código LOINC</FormLabel>
                    <FormControl>
                      <Input placeholder="718-7" {...field} />
                    </FormControl>
                    <FormDescription>Código internacional</FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />
            </div>

            <FormField
              control={form.control}
              name="resultType"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Tipo de Resultado *</FormLabel>
                  <Select
                    onValueChange={field.onChange}
                    defaultValue={field.value}
                  >
                    <FormControl>
                      <SelectTrigger>
                        <SelectValue placeholder="Selecione o tipo" />
                      </SelectTrigger>
                    </FormControl>
                    <SelectContent>
                      {resultTypeOptions.map((option) => (
                        <SelectItem key={option.value} value={option.value}>
                          <div>
                            <div className="font-medium">{option.label}</div>
                            <div className="text-xs text-muted-foreground">
                              {option.description}
                            </div>
                          </div>
                        </SelectItem>
                      ))}
                    </SelectContent>
                  </Select>
                  <FormMessage />
                </FormItem>
              )}
            />

            <div className="grid grid-cols-2 gap-4">
              <FormField
                control={form.control}
                name="unit"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Unidade de Medida</FormLabel>
                    <FormControl>
                      <Input placeholder="g/dL, mg/dL, mU/L" {...field} />
                    </FormControl>
                    <FormDescription>Para resultados numéricos</FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />

              <FormField
                control={form.control}
                name="displayOrder"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Ordem de Exibição</FormLabel>
                    <FormControl>
                      <Input
                        type="number"
                        placeholder="0"
                        {...field}
                        onChange={(e) =>
                          field.onChange(parseInt(e.target.value) || 0)
                        }
                      />
                    </FormControl>
                    <FormDescription>
                      Ordem para listar parâmetros
                    </FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />
            </div>

            <FormField
              control={form.control}
              name="unitConversion"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Conversão de Unidades</FormLabel>
                  <FormControl>
                    <Input placeholder="1 g/dL = 10 g/L" {...field} />
                  </FormControl>
                  <FormDescription>
                    Fórmula de conversão (opcional)
                  </FormDescription>
                  <FormMessage />
                </FormItem>
              )}
            />
          </TabsContent>

          {/* Collection Information Tab */}
          <TabsContent value="collection" className="space-y-4 mt-4">
            <FormField
              control={form.control}
              name="specimenType"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Tipo de Amostra</FormLabel>
                  <FormControl>
                    <Input
                      placeholder="Sangue total, Soro, Urina, Fezes"
                      {...field}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name="fastingHours"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Horas de Jejum</FormLabel>
                  <FormControl>
                    <Input
                      type="number"
                      placeholder="8"
                      {...field}
                      value={field.value ?? ""}
                      onChange={(e) => {
                        const value = e.target.value;
                        field.onChange(value === "" ? null : parseInt(value));
                      }}
                    />
                  </FormControl>
                  <FormDescription>
                    Tempo de jejum necessário (em horas)
                  </FormDescription>
                  <FormMessage />
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name="collectionMethod"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Método de Coleta</FormLabel>
                  <FormControl>
                    <Textarea
                      placeholder="Sangue venoso com jejum de 8-12h"
                      className="resize-none"
                      rows={3}
                      {...field}
                    />
                  </FormControl>
                  <FormDescription>
                    Instruções para coleta da amostra
                  </FormDescription>
                  <FormMessage />
                </FormItem>
              )}
            />
          </TabsContent>

          {/* Documentation Tab */}
          <TabsContent value="documentation" className="space-y-4 mt-4">
            <FormField
              control={form.control}
              name="description"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Descrição</FormLabel>
                  <FormControl>
                    <Textarea
                      placeholder="Descrição técnica do exame..."
                      className="resize-none"
                      rows={4}
                      {...field}
                    />
                  </FormControl>
                  <FormDescription>
                    Informações técnicas sobre o exame
                  </FormDescription>
                  <FormMessage />
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name="clinicalIndications"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Indicações Clínicas</FormLabel>
                  <FormControl>
                    <Textarea
                      placeholder="Indicado para diagnóstico de..."
                      className="resize-none"
                      rows={4}
                      {...field}
                    />
                  </FormControl>
                  <FormDescription>
                    Quando e por que solicitar este exame
                  </FormDescription>
                  <FormMessage />
                </FormItem>
              )}
            />
          </TabsContent>
        </Tabs>

        {/* Form Actions */}
        <div className="flex justify-end gap-2 pt-4 border-t">
          <Button
            type="button"
            variant="outline"
            onClick={onCancel}
            disabled={isSubmitting}
          >
            Cancelar
          </Button>
          <Button type="submit" disabled={isSubmitting}>
            {isSubmitting && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
            {mode === "create" ? "Criar Definição" : "Salvar Alterações"}
          </Button>
        </div>
      </form>
    </Form>
  );
}
