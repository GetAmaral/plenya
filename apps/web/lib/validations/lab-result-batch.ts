import { z } from "zod";

// Schema para um resultado individual dentro do batch
export const labResultInBatchSchema = z.object({
  id: z.string().uuid().optional(), // ID para update de resultados existentes
  labTestDefinitionId: z.string().uuid().optional(),
  testName: z.string().max(200, "Nome do exame deve ter no máximo 200 caracteres").optional(),
  testType: z.string().max(100, "Tipo do exame deve ter no máximo 100 caracteres").optional(),
  resultText: z.string().optional(),
  resultNumeric: z.number().optional(),
  unit: z.string().max(50).optional(),
  interpretation: z.string().optional(),
  level: z.number().int().optional(),
  matched: z.boolean().optional(), // true se matched com LabTestDefinition, false se extraído mas não catalogado
});

// Schema para o batch completo com preprocessamento
export const labResultBatchSchema = z.preprocess(
  (data: any) => {
    // Filtrar linhas vazias ANTES da validação
    if (data?.labResults) {
      data.labResults = data.labResults.filter((r: any) => {
        const hasDefinition = !!r.labTestDefinitionId;
        const hasName = r.testName && r.testName.trim().length > 0;
        const hasValue = r.resultText || r.resultNumeric !== undefined;

        // Mantém a linha se tiver: definição OU nome OU valor
        return hasDefinition || hasName || hasValue;
      });
    }
    return data;
  },
  z
    .object({
      labRequestId: z.string().uuid().optional(),
      requestingDoctorId: z.string().uuid().optional(),
      laboratoryName: z
        .string()
        .min(2, "Nome do laboratório deve ter no mínimo 2 caracteres")
        .max(200, "Nome do laboratório deve ter no máximo 200 caracteres"),
      collectionDate: z.date({ required_error: "Data da coleta é obrigatória" }),
      resultDate: z.date().optional(),
      status: z.enum(["pending", "partial", "completed"] as const, {
        required_error: "Status do lote é obrigatório",
      }),
      observations: z.string().optional(),
      attachments: z.string().optional(),
      labResults: z.array(labResultInBatchSchema),
    })
    .refine(
      (data) => data.labResults.length > 0,
      {
        message: "Adicione pelo menos um resultado ao lote",
        path: ["labResults"],
      }
    )
);

export type LabResultInBatchFormValues = z.infer<typeof labResultInBatchSchema>;
export type LabResultBatchFormValues = z.infer<typeof labResultBatchSchema>;

// Valores padrão
export const getDefaultBatchValues = (): LabResultBatchFormValues => ({
  laboratoryName: "",
  collectionDate: new Date(),
  status: "pending",
  labResults: [getDefaultResultValues()], // Sempre inicia com uma linha vazia
});

export const getDefaultResultValues = (): LabResultInBatchFormValues => ({
  testName: "",
  testType: "",
});

// Helpers de conversão
export const formToApiValues = (values: LabResultBatchFormValues) => {
  // Filtrar linhas vazias (sem testName)
  const filledResults = values.labResults.filter(r => r.testName && r.testName.trim());

  return {
    laboratoryName: values.laboratoryName,
    collectionDate: values.collectionDate.toISOString(),
    resultDate: values.resultDate?.toISOString(),
    status: values.status,
    observations: values.observations,
    attachments: values.attachments,
    labRequestId: values.labRequestId,
    requestingDoctorId: values.requestingDoctorId,
    labResults: filledResults.map((result) => ({
      id: result.id, // Inclui ID para update (se existir)
      labTestDefinitionId: result.labTestDefinitionId,
      testName: result.testName,
      testType: result.testType,
      resultText: result.resultText,
      resultNumeric: result.resultNumeric,
      unit: result.unit,
      interpretation: result.interpretation,
      level: result.level,
    })),
  };
};

export const apiToFormValues = (batch: any): LabResultBatchFormValues => {
  const existingResults = (batch.labResults || []).map((result: any) => ({
    id: result.id, // Preserva ID do resultado existente
    labTestDefinitionId: result.labTestDefinitionId,
    testName: result.testName || "",
    testType: result.testType || "",
    resultText: result.resultText,
    resultNumeric: result.resultNumeric,
    unit: result.unit,
    interpretation: result.interpretation,
    level: result.level,
  }));

  return {
    laboratoryName: batch.laboratoryName || "",
    collectionDate: batch.collectionDate ? new Date(batch.collectionDate) : new Date(),
    resultDate: batch.resultDate ? new Date(batch.resultDate) : undefined,
    status: batch.status || "pending",
    observations: batch.observations,
    attachments: batch.attachments,
    labRequestId: batch.labRequestId,
    requestingDoctorId: batch.requestingDoctorId,
    // Sempre adiciona uma linha vazia no final ao editar
    labResults: [...existingResults, getDefaultResultValues()],
  };
};
