"use client";

import { useRef, useState, useEffect } from "react";
import { useRouter } from "next/navigation";
import { useForm, useFieldArray } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { format } from "date-fns";
import { CalendarIcon, Loader2, Package } from "lucide-react";
import { toast } from "sonner";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";
import { Calendar } from "@/components/ui/calendar";
import {
  Form,
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
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { cn } from "@/lib/utils";
import { useFormNavigation } from "@/lib/use-form-navigation";
import {
  labResultBatchSchema,
  getDefaultBatchValues,
  getDefaultResultValues,
  formToApiValues,
  type LabResultBatchFormValues,
} from "@/lib/validations/lab-result-batch";
import { labResultBatchApi } from "@/lib/api/lab-result-batch-api";
import { useProcessingJobStore } from "@/lib/processing-job-store";
import { LabResultTableForm } from "./LabResultTableForm";
import { PDFUploadZone } from "./PDFUploadZone";
import { ProcessingStatus } from "./ProcessingStatus";

interface LabResultBatchFormProps {
  batchId?: string;
  initialValues?: LabResultBatchFormValues;
}

export function LabResultBatchForm({ batchId, initialValues }: LabResultBatchFormProps = {}) {
  const router = useRouter();
  const queryClient = useQueryClient();
  const formRef = useRef<HTMLFormElement>(null);
  const isEditMode = !!batchId;
  const { addJob } = useProcessingJobStore();

  // PDF processing states
  const [processingJobId, setProcessingJobId] = useState<string | null>(null);
  const [processingStatus, setProcessingStatus] = useState<"idle" | "processing" | "completed" | "failed">("idle");
  const [tempBatchId, setTempBatchId] = useState<string | null>(null); // Batch temporário criado no upload
  const [isManualSubmitting, setIsManualSubmitting] = useState(false); // Loading state for manual save

  useFormNavigation({ formRef });

  const form = useForm<LabResultBatchFormValues>({
    resolver: zodResolver(labResultBatchSchema),
    defaultValues: initialValues || getDefaultBatchValues(),
    mode: "onBlur",
  });


  const { fields, append, remove } = useFieldArray({
    control: form.control,
    name: "labResults",
  });

  const createMutation = useMutation({
    mutationFn: async (values: LabResultBatchFormValues) => {
      const apiData = formToApiValues(values);
      return labResultBatchApi.create(apiData);
    },
    onSuccess: (data) => {
      toast.success("Lote de resultados criado com sucesso!");
      queryClient.invalidateQueries({ queryKey: ["lab-result-batches"] });
      router.push("/lab-results");
    },
    onError: (error: any) => {
      toast.error(
        error?.response?.data?.message || "Erro ao criar lote de resultados"
      );
    },
  });

  // Mutation para criar batch temporário (usado no upload de PDF em modo de criação)
  const createTempBatchMutation = useMutation({
    mutationFn: async () => {
      // Criar batch temporário com dados mínimos do formulário
      const currentValues = form.getValues();
      const apiData = {
        laboratoryName: currentValues.laboratoryName || "Importado via PDF",
        collectionDate: currentValues.collectionDate?.toISOString() || new Date().toISOString(),
        status: "pending" as const,
        labResults: [
          {
            testName: "Processando...",
            testType: "pending",
          }
        ], // Resultado placeholder que será substituído pelo processamento
      };
      return labResultBatchApi.create(apiData);
    },
    onSuccess: (data) => {
      setTempBatchId(data.id);
      toast.success("Batch temporário criado. Fazendo upload do PDF...");
    },
    onError: (error: any) => {
      toast.error(
        error?.response?.data?.message || "Erro ao criar batch temporário"
      );
    },
  });

  const updateMutation = useMutation({
    mutationFn: async (values: LabResultBatchFormValues) => {
      if (!batchId) throw new Error("Batch ID is required for update");
      // Envia batch metadata + resultados (sync automático no backend)
      const apiData = formToApiValues(values);
      return labResultBatchApi.update(batchId, apiData);
    },
    onSuccess: (data) => {
      toast.success("Lote de resultados atualizado com sucesso!");
      queryClient.invalidateQueries({ queryKey: ["lab-result-batches"] });
      queryClient.invalidateQueries({ queryKey: ["lab-result-batch", batchId] });
      router.push("/lab-results");
    },
    onError: (error: any) => {
      toast.error(
        error?.response?.data?.message || "Erro ao atualizar lote de resultados"
      );
    },
  });

  const onSubmit = async (values: LabResultBatchFormValues) => {
    console.log("🚀 onSubmit called", { tempBatchId, batchId, isEditMode, values });

    // Se já criou batch temporário via upload, atualizar em vez de criar
    if (tempBatchId) {
      console.log("📝 Updating temp batch", tempBatchId);
      setIsManualSubmitting(true);
      try {
        const apiData = formToApiValues(values);
        console.log("📤 Sending data to API", apiData);
        await labResultBatchApi.update(tempBatchId, apiData as any);
        toast.success("Lote de resultados salvo com sucesso!");
        queryClient.invalidateQueries({ queryKey: ["lab-result-batches"] });
        router.push("/lab-results");
      } catch (error: any) {
        console.error("❌ Error saving batch", error);
        toast.error(error?.response?.data?.message || "Erro ao salvar lote");
      } finally {
        setIsManualSubmitting(false);
      }
    } else if (isEditMode) {
      console.log("✏️ Edit mode - updating batch", batchId);
      updateMutation.mutate(values);
    } else {
      console.log("➕ Create mode - creating new batch");
      createMutation.mutate(values);
    }
  };

  const handleUploadSuccess = (jobId: string) => {
    setProcessingJobId(jobId);
    setProcessingStatus("processing");

    // Add to global processing store for background monitoring
    const effectiveBatchId = batchId || tempBatchId;
    if (effectiveBatchId) {
      addJob(jobId, effectiveBatchId);
    }
  };

  const handleProcessingCompleted = async () => {
    setProcessingStatus("completed");

    // Refresh batch data to get AI-extracted results
    const effectiveBatchId = batchId || tempBatchId;
    if (effectiveBatchId) {
      try {
        const batch = await labResultBatchApi.getById(effectiveBatchId);

        // Verificar se há resultados extraídos
        const hasResults = batch.labResults && batch.labResults.length > 0;

        if (hasResults) {
          toast.success("Interpretação concluída!", {
            description: "Os resultados foram preenchidos automaticamente. Revise e confirme.",
          });
          // Reset form with new data
          form.reset(batch as any);
        } else {
          toast.warning("Nenhum resultado extraído", {
            description: "A IA não conseguiu extrair resultados do PDF. Adicione manualmente ou tente outro arquivo.",
          });
        }

        queryClient.invalidateQueries({ queryKey: ["lab-result-batch", effectiveBatchId] });
        queryClient.invalidateQueries({ queryKey: ["lab-result-batches"] });

        // Após sucesso no processamento, redirecionar para a lista
        router.push("/lab-results");
      } catch (error) {
        console.error("Failed to refresh batch data:", error);
        toast.error("Erro ao carregar dados do batch");
      }
    }
  };

  const handleProcessingFailed = (error: string) => {
    setProcessingStatus("failed");
    toast.error("Falha na interpretação", { description: error });
  };

  const isSubmitting = createMutation.isPending || updateMutation.isPending || isManualSubmitting;

  return (
    <Form {...form}>
      <form
        ref={formRef}
        onSubmit={form.handleSubmit(
          (values) => {
            console.log("✅ Form valid, calling onSubmit");
            onSubmit(values);
          },
          (errors) => {
            console.error("❌ Form validation failed:", errors);
            toast.error("Preencha todos os campos obrigatórios corretamente");
          }
        )}
        className="space-y-6"
      >
        <div className="space-y-6">
            {/* Informações Gerais */}
            <Card>
              <CardHeader>
                <CardTitle className="flex items-center gap-2">
                  <Package className="h-5 w-5" />
                  Informações do Lote
                </CardTitle>
              </CardHeader>
              <CardContent className="space-y-4">
                {/* Laboratório */}
                <FormField
                  control={form.control}
                  name="laboratoryName"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Laboratório *</FormLabel>
                      <FormControl>
                        <Input
                          {...field}
                          placeholder="Nome do laboratório"
                        />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />

                {/* Datas */}
                <div className="grid grid-cols-2 gap-4">
                  <FormField
                    control={form.control}
                    name="collectionDate"
                    render={({ field }) => (
                      <FormItem className="flex flex-col">
                        <FormLabel>Data da Coleta *</FormLabel>
                        <Popover>
                          <PopoverTrigger asChild>
                            <FormControl>
                              <Button
                                variant="outline"
                                className={cn(
                                  "pl-3 text-left font-normal",
                                  !field.value && "text-muted-foreground"
                                )}
                              >
                                {field.value ? (
                                  format(field.value, "dd/MM/yyyy")
                                ) : (
                                  <span>Selecione a data</span>
                                )}
                                <CalendarIcon className="ml-auto h-4 w-4 opacity-50" />
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

                  <FormField
                    control={form.control}
                    name="resultDate"
                    render={({ field }) => (
                      <FormItem className="flex flex-col">
                        <FormLabel>Data do Resultado</FormLabel>
                        <Popover>
                          <PopoverTrigger asChild>
                            <FormControl>
                              <Button
                                variant="outline"
                                className={cn(
                                  "pl-3 text-left font-normal",
                                  !field.value && "text-muted-foreground"
                                )}
                              >
                                {field.value ? (
                                  format(field.value, "dd/MM/yyyy")
                                ) : (
                                  <span>Selecione a data</span>
                                )}
                                <CalendarIcon className="ml-auto h-4 w-4 opacity-50" />
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

                {/* Status */}
                <FormField
                  control={form.control}
                  name="status"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Status do Lote *</FormLabel>
                      <Select onValueChange={field.onChange} value={field.value}>
                        <FormControl>
                          <SelectTrigger>
                            <SelectValue placeholder="Selecione o status" />
                          </SelectTrigger>
                        </FormControl>
                        <SelectContent>
                          <SelectItem value="pending">Pendente</SelectItem>
                          <SelectItem value="partial">Parcial</SelectItem>
                          <SelectItem value="completed">Concluído</SelectItem>
                        </SelectContent>
                      </Select>
                      <FormMessage />
                    </FormItem>
                  )}
                />

                {/* Observações */}
                <FormField
                  control={form.control}
                  name="observations"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Observações Gerais</FormLabel>
                      <FormControl>
                        <Textarea
                          {...field}
                          placeholder="Observações sobre o lote de exames..."
                          rows={3}
                          value={field.value || ""}
                        />
                      </FormControl>
                      <FormDescription>
                        Informações adicionais sobre o lote
                      </FormDescription>
                      <FormMessage />
                    </FormItem>
                  )}
                />
              </CardContent>
            </Card>

            {/* PDF Upload - Disponível tanto para criar quanto editar */}
            <Card>
              <CardHeader>
                <CardTitle>
                  {isEditMode ? "Importar Laudo (PDF)" : "Importar Laudo via PDF (Opcional)"}
                </CardTitle>
                {!isEditMode && (
                  <p className="text-sm text-muted-foreground">
                    Faça upload de um PDF de laudo médico e a IA preencherá automaticamente os resultados
                  </p>
                )}
              </CardHeader>
              <CardContent className="space-y-4">
                <PDFUploadZone
                  batchId={batchId || tempBatchId || undefined}
                  onUploadSuccess={handleUploadSuccess}
                  onBeforeUpload={async () => {
                    // Criar batch temporário antes do upload (modo de criação)
                    const result = await createTempBatchMutation.mutateAsync()
                    return result.id
                  }}
                  disabled={processingStatus === "processing" || createTempBatchMutation.isPending}
                />

                {processingJobId && processingStatus !== "completed" && (
                  <ProcessingStatus
                    jobId={processingJobId}
                    onCompleted={handleProcessingCompleted}
                    onFailed={handleProcessingFailed}
                  />
                )}

                {!isEditMode && processingStatus === "completed" && (
                  <div className="text-sm text-muted-foreground p-4 border rounded-lg bg-blue-50">
                    ℹ️ Resultados preenchidos automaticamente. Revise os dados abaixo e clique em "Criar Lote" para salvar.
                  </div>
                )}
              </CardContent>
            </Card>

            {/* Resultados */}
            <Card>
              <CardContent className="pt-6">
                <LabResultTableForm
                  form={form}
                  fields={fields}
                  append={append}
                  remove={remove}
                  getDefaultResultValues={getDefaultResultValues}
                />
              </CardContent>
            </Card>

            {/* Actions */}
            <div className="flex justify-end gap-2 pt-4 border-t">
              <Button
                type="button"
                variant="outline"
                onClick={() => router.back()}
                disabled={isSubmitting}
              >
                Cancelar
              </Button>
              <Button type="submit" disabled={isSubmitting}>
                {isSubmitting && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
                {isSubmitting ? "Salvando..." : (isEditMode ? "Atualizar Lote" : "Salvar Lote")}
              </Button>
            </div>
        </div>
      </form>
    </Form>
  );
}
