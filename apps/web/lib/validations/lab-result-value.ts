import * as z from "zod";

/**
 * Schema for a single lab result value entry
 */
export const labResultValueEntrySchema = z.object({
  labTestDefinitionId: z.string().uuid(),
  numericValue: z.number().nullable().optional(),
  textValue: z.string().optional(),
  booleanValue: z.boolean().nullable().optional(),
  unit: z.string().max(50).optional(),
  notes: z.string().optional(),
});

/**
 * Schema for batch lab result values entry
 * Note: patientId is optional - if not provided, backend will auto-fill from selectedPatient
 */
export const labResultValuesBatchSchema = z.object({
  patientId: z.string().uuid().optional(),
  doctorId: z.string().uuid({
    error: "Médico é obrigatório",
  }),
  testDate: z.date({
    error: "Data do exame é obrigatória",
  }),
  selectedTestId: z.string().uuid({
    error: "Exame é obrigatório",
  }),
  notes: z.string().optional(),
  values: z.array(labResultValueEntrySchema).min(1, "Pelo menos um valor é obrigatório"),
});

export type LabResultValueEntry = z.infer<typeof labResultValueEntrySchema>;
export type LabResultValuesBatchInput = z.infer<typeof labResultValuesBatchSchema>;

/**
 * Dynamic field schema based on test definition
 */
export const createDynamicFieldSchema = (resultType: string) => {
  switch (resultType) {
    case "numeric":
      return z.object({
        value: z.number({ error: "Valor é obrigatório" }),
        unit: z.string().optional(),
        notes: z.string().optional(),
      });

    case "text":
      return z.object({
        value: z.string().min(1, "Valor é obrigatório"),
        notes: z.string().optional(),
      });

    case "boolean":
      return z.object({
        value: z.boolean({
          error: "Valor é obrigatório",
        }),
        notes: z.string().optional(),
      });

    case "categorical":
      return z.object({
        value: z.string().min(1, "Valor é obrigatório"),
        notes: z.string().optional(),
      });

    default:
      return z.object({
        value: z.string().min(1, "Valor é obrigatório"),
        notes: z.string().optional(),
      });
  }
};

/**
 * Get default value based on result type
 */
export const getDefaultValueForType = (resultType: string): any => {
  switch (resultType) {
    case "numeric":
      return { value: null, unit: "", notes: "" };
    case "text":
      return { value: "", notes: "" };
    case "boolean":
      return { value: null, notes: "" };
    case "categorical":
      return { value: "", notes: "" };
    default:
      return { value: "", notes: "" };
  }
};

/**
 * Convert form value to API payload based on result type
 */
export const valueToApiPayload = (
  labTestDefinitionId: string,
  formValue: any,
  resultType: string,
  defaultUnit?: string
): LabResultValueEntry => {
  const base = {
    labTestDefinitionId,
    notes: formValue.notes || undefined,
  };

  switch (resultType) {
    case "numeric":
      return {
        ...base,
        numericValue: formValue.value !== null && formValue.value !== "" ? Number(formValue.value) : null,
        unit: formValue.unit || defaultUnit || undefined,
      };

    case "text":
    case "categorical":
      return {
        ...base,
        textValue: formValue.value || undefined,
      };

    case "boolean":
      return {
        ...base,
        booleanValue: formValue.value,
      };

    default:
      return {
        ...base,
        textValue: String(formValue.value) || undefined,
      };
  }
};
