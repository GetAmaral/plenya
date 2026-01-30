"use client";

import { useState, useEffect, useRef, useMemo } from "react";
import { useForm, useFieldArray, useWatch } from "react-hook-form";
import { useFormNavigation } from "@/lib/use-form-navigation";
import { zodResolver } from "@hookform/resolvers/zod";
import { useQuery } from "@tanstack/react-query";
import { CalendarIcon, Check, ChevronsUpDown, Loader2, User2 } from "lucide-react";
import { format, differenceInYears } from "date-fns";
import { ptBR } from "date-fns/locale";
import { cn } from "@/lib/utils";
import { useSelectedPatient } from "@/lib/use-selected-patient";
import {
  labResultValuesBatchSchema,
  type LabResultValuesBatchInput,
  getDefaultValueForType,
} from "@/lib/validations/lab-result-value";
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
import { Calendar } from "@/components/ui/calendar";
import { Switch } from "@/components/ui/switch";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Separator } from "@/components/ui/separator";
import {
  labTestDefinitionsApi,
  type LabTestDefinition,
} from "@/lib/api/lab-test-api";
import { apiClient } from "@/lib/api-client";
import { Badge } from "@/components/ui/badge";

interface LabResultEntryFormProps {
  currentUserId: string;
  onSubmit: (data: LabResultValuesBatchInput) => Promise<void>;
  onCancel: () => void;
  isSubmitting?: boolean;
}

export function LabResultEntryForm({
  currentUserId,
  onSubmit,
  onCancel,
  isSubmitting = false,
}: LabResultEntryFormProps) {
  const [testOpen, setTestOpen] = useState(false);
  const { selectedPatient } = useSelectedPatient();

  const formRef = useRef<HTMLFormElement>(null);
  useFormNavigation({ formRef });

  // Fetch requestable tests
  const { data: requestableTests = [] } = useQuery({
    queryKey: ["lab-tests-requestable"],
    queryFn: () => labTestDefinitionsApi.getRequestable(),
  });

  const form = useForm<LabResultValuesBatchInput>({
    resolver: zodResolver(labResultValuesBatchSchema),
    defaultValues: {
      patientId: undefined, // Will be auto-filled by backend from selectedPatient
      doctorId: currentUserId,
      testDate: new Date(),
      selectedTestId: "",
      notes: "",
      values: [],
    },
  });

  const { fields, replace } = useFieldArray({
    control: form.control,
    name: "values",
  });

  // Watch selected test ID using useWatch (stable subscription)
  const selectedTestId = useWatch({
    control: form.control,
    name: "selectedTestId",
  });

  // Derive selectedTest from selectedTestId (no state needed)
  const selectedTest = useMemo(() => {
    if (!selectedTestId || !requestableTests.length) return null;
    return requestableTests.find((t) => t.id === selectedTestId) || null;
  }, [selectedTestId, requestableTests]);

  // Derive testParameters from selectedTest (no state needed)
  const testParameters = useMemo(() => {
    if (!selectedTest) return [];

    if (selectedTest.subTests && selectedTest.subTests.length > 0) {
      return [...selectedTest.subTests].sort((a, b) => a.displayOrder - b.displayOrder);
    }

    return [selectedTest];
  }, [selectedTest]);

  // Track last processed test to avoid re-initializing form unnecessarily
  const lastProcessedTestId = useRef<string | undefined>(undefined);
  const isInitializing = useRef(false);

  // Initialize form values when test parameters change
  useEffect(() => {
    // Avoid processing if already in progress
    if (isInitializing.current) return;

    // Check if we need to update
    const currentTestId = selectedTestId || undefined;
    if (lastProcessedTestId.current === currentTestId) return;

    isInitializing.current = true;
    lastProcessedTestId.current = currentTestId;

    if (!selectedTestId || testParameters.length === 0) {
      replace([]);
    } else {
      const initialValues = testParameters.map((param) => ({
        labTestDefinitionId: param.id,
        ...getDefaultValueForType(param.resultType),
        _testName: param.name,
        _resultType: param.resultType,
        _unit: param.unit,
      }));
      replace(initialValues as any);
    }

    // Reset flag after a microtask to allow render to complete
    Promise.resolve().then(() => {
      isInitializing.current = false;
    });
  }, [selectedTestId, testParameters, replace]);

  const handleSubmit = async (data: LabResultValuesBatchInput) => {
    await onSubmit(data);
  };

  return (
    <Form {...form}>
      <form ref={formRef} onSubmit={form.handleSubmit(handleSubmit)} className="space-y-6">
        {/* Basic Information */}
        <Card>
          <CardHeader>
            <CardTitle>Informações Básicas</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            {/* Selected Patient Info (Read-only) */}
            {selectedPatient && (
              <div className="rounded-lg border bg-muted/50 p-4">
                <div className="flex items-center gap-3">
                  <div className="flex h-10 w-10 items-center justify-center rounded-full bg-blue-600 text-white">
                    <User2 className="h-5 w-5" />
                  </div>
                  <div>
                    <div className="flex items-center gap-2">
                      <p className="font-medium">{selectedPatient.name}</p>
                      <Badge variant="outline" className="text-xs">
                        {differenceInYears(new Date(), new Date(selectedPatient.birthDate))} anos
                      </Badge>
                    </div>
                    <p className="text-sm text-muted-foreground">
                      Paciente selecionado para este exame
                    </p>
                  </div>
                </div>
              </div>
            )}

            <div className="grid grid-cols-2 gap-4">

              {/* Test Date */}
              <FormField
                control={form.control}
                name="testDate"
                render={({ field }) => (
                  <FormItem className="flex flex-col">
                    <FormLabel>Data do Exame *</FormLabel>
                    <Popover>
                      <PopoverTrigger asChild>
                        <FormControl>
                          <Button
                            variant="outline"
                            className={cn(
                              "justify-start text-left font-normal",
                              !field.value && "text-muted-foreground"
                            )}
                          >
                            <CalendarIcon className="mr-2 h-4 w-4" />
                            {field.value ? (
                              format(field.value, "PPP", { locale: ptBR })
                            ) : (
                              <span>Selecione a data</span>
                            )}
                          </Button>
                        </FormControl>
                      </PopoverTrigger>
                      <PopoverContent className="w-auto p-0" align="start">
                        <Calendar
                          mode="single"
                          selected={field.value}
                          onSelect={field.onChange}
                          disabled={(date) =>
                            date > new Date() || date < new Date("1900-01-01")
                          }
                          initialFocus
                        />
                      </PopoverContent>
                    </Popover>
                    <FormMessage />
                  </FormItem>
                )}
              />
            </div>

            {/* Test Selection */}
            <FormField
              control={form.control}
              name="selectedTestId"
              render={({ field }) => (
                <FormItem className="flex flex-col">
                  <FormLabel>Exame Realizado *</FormLabel>
                  <Popover open={testOpen} onOpenChange={setTestOpen}>
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
                            : "Selecione o exame"}
                          <ChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
                        </Button>
                      </FormControl>
                    </PopoverTrigger>
                    <PopoverContent className="w-[500px] p-0">
                      <Command>
                        <CommandInput placeholder="Buscar exame..." />
                        <CommandEmpty>Nenhum exame encontrado.</CommandEmpty>
                        <CommandGroup className="max-h-64 overflow-auto">
                          {requestableTests.map((test) => (
                            <CommandItem
                              key={test.id}
                              value={test.name}
                              onSelect={() => {
                                form.setValue("selectedTestId", test.id);
                                setTestOpen(false);
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
                                  {test.subTests && test.subTests.length > 0 && (
                                    <> • {test.subTests.length} parâmetros</>
                                  )}
                                </div>
                              </div>
                            </CommandItem>
                          ))}
                        </CommandGroup>
                      </Command>
                    </PopoverContent>
                  </Popover>
                  <FormDescription>
                    Selecione qual exame foi realizado
                  </FormDescription>
                  <FormMessage />
                </FormItem>
              )}
            />

            {/* General Notes */}
            <FormField
              control={form.control}
              name="notes"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Observações Gerais</FormLabel>
                  <FormControl>
                    <Textarea
                      placeholder="Observações sobre a coleta, condições do paciente, etc."
                      className="resize-none"
                      rows={2}
                      {...field}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
          </CardContent>
        </Card>

        {/* Test Parameters/Values */}
        {selectedTest && testParameters.length > 0 && (
          <Card>
            <CardHeader>
              <CardTitle>
                Valores - {selectedTest.name}
                {testParameters.length > 1 && ` (${testParameters.length} parâmetros)`}
              </CardTitle>
            </CardHeader>
            <CardContent>
              <div className="space-y-4">
                {fields.map((field, index) => {
                  const param = testParameters[index];
                  if (!param) return null;

                  return (
                    <div key={field.id}>
                      {index > 0 && <Separator className="my-4" />}

                      <div className="space-y-3">
                        <div className="flex items-center justify-between">
                          <div>
                            <h4 className="font-medium">{param.name}</h4>
                            {param.shortName && (
                              <p className="text-sm text-muted-foreground">
                                {param.shortName}
                              </p>
                            )}
                          </div>
                          {param.unit && (
                            <div className="text-sm text-muted-foreground">
                              Unidade: {param.unit}
                            </div>
                          )}
                        </div>

                        {/* Numeric Input */}
                        {param.resultType === "numeric" && (
                          <div className="grid grid-cols-2 gap-4">
                            <FormField
                              control={form.control}
                              name={`values.${index}.value` as any}
                              render={({ field }) => (
                                <FormItem>
                                  <FormLabel>Valor *</FormLabel>
                                  <FormControl>
                                    <Input
                                      type="number"
                                      step="any"
                                      placeholder="0.00"
                                      {...field}
                                      onChange={(e) => {
                                        const value = e.target.value;
                                        field.onChange(value === "" ? null : parseFloat(value));
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
                              name={`values.${index}.unit` as any}
                              render={({ field }) => (
                                <FormItem>
                                  <FormLabel>Unidade</FormLabel>
                                  <FormControl>
                                    <Input
                                      placeholder={param.unit || "mg/dL"}
                                      {...field}
                                    />
                                  </FormControl>
                                  <FormMessage />
                                </FormItem>
                              )}
                            />
                          </div>
                        )}

                        {/* Text Input */}
                        {(param.resultType === "text" || param.resultType === "categorical") && (
                          <FormField
                            control={form.control}
                            name={`values.${index}.value` as any}
                            render={({ field }) => (
                              <FormItem>
                                <FormLabel>Valor *</FormLabel>
                                <FormControl>
                                  <Input
                                    placeholder="Ex: Negativo, Positivo, Normal"
                                    {...field}
                                  />
                                </FormControl>
                                <FormMessage />
                              </FormItem>
                            )}
                          />
                        )}

                        {/* Boolean Input */}
                        {param.resultType === "boolean" && (
                          <FormField
                            control={form.control}
                            name={`values.${index}.value` as any}
                            render={({ field }) => (
                              <FormItem className="flex flex-row items-center justify-between rounded-lg border p-4">
                                <div className="space-y-0.5">
                                  <FormLabel className="text-base">Resultado</FormLabel>
                                  <FormDescription>
                                    {field.value === true && "Positivo / Presente"}
                                    {field.value === false && "Negativo / Ausente"}
                                    {field.value === null && "Não informado"}
                                  </FormDescription>
                                </div>
                                <FormControl>
                                  <Switch
                                    checked={field.value === true}
                                    onCheckedChange={field.onChange}
                                  />
                                </FormControl>
                              </FormItem>
                            )}
                          />
                        )}

                        {/* Notes for this parameter */}
                        <FormField
                          control={form.control}
                          name={`values.${index}.notes` as any}
                          render={({ field }) => (
                            <FormItem>
                              <FormLabel>Observações</FormLabel>
                              <FormControl>
                                <Textarea
                                  placeholder="Observações específicas para este parâmetro"
                                  className="resize-none"
                                  rows={2}
                                  {...field}
                                />
                              </FormControl>
                              <FormMessage />
                            </FormItem>
                          )}
                        />
                      </div>
                    </div>
                  );
                })}
              </div>
            </CardContent>
          </Card>
        )}

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
          <Button type="submit" disabled={isSubmitting || !selectedTest}>
            {isSubmitting && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
            Salvar Resultado
          </Button>
        </div>
      </form>
    </Form>
  );
}
