import * as z from "zod";
import type { LabTestCategory, ResultType } from "@/lib/api/lab-test-api";

/**
 * Zod schema for Lab Test Definition form validation
 */
export const labTestDefinitionSchema = z.object({
  code: z
    .string()
    .min(2, "Código deve ter no mínimo 2 caracteres")
    .max(100, "Código deve ter no máximo 100 caracteres")
    .regex(
      /^[A-Z0-9_]+$/,
      "Código deve conter apenas letras maiúsculas, números e underscore"
    ),

  name: z
    .string()
    .min(3, "Nome deve ter no mínimo 3 caracteres")
    .max(300, "Nome deve ter no máximo 300 caracteres"),

  shortName: z
    .string()
    .max(50, "Nome curto deve ter no máximo 50 caracteres")
    .optional()
    .or(z.literal("")),

  tussCode: z
    .string()
    .max(20, "Código TUSS deve ter no máximo 20 caracteres")
    .optional()
    .or(z.literal("")),

  loincCode: z
    .string()
    .max(20, "Código LOINC deve ter no máximo 20 caracteres")
    .optional()
    .or(z.literal("")),

  category: z.enum([
    "hematology",
    "biochemistry",
    "hormones",
    "immunology",
    "microbiology",
    "urine",
    "imaging",
    "functional",
    "genetics",
    "other",
  ] as const, {
    required_error: "Categoria é obrigatória",
  }),

  isRequestable: z.boolean().default(true),

  parentTestId: z.string().uuid().optional().or(z.literal("")),

  unit: z
    .string()
    .max(50, "Unidade deve ter no máximo 50 caracteres")
    .optional()
    .or(z.literal("")),

  unitConversion: z.string().optional().or(z.literal("")),

  resultType: z.enum(["numeric", "text", "boolean", "categorical"] as const, {
    required_error: "Tipo de resultado é obrigatório",
  }),

  collectionMethod: z.string().optional().or(z.literal("")),

  fastingHours: z
    .number()
    .int("Horas de jejum deve ser um número inteiro")
    .min(0, "Horas de jejum não pode ser negativo")
    .max(48, "Horas de jejum não pode exceder 48 horas")
    .optional()
    .nullable(),

  specimenType: z
    .string()
    .max(100, "Tipo de amostra deve ter no máximo 100 caracteres")
    .optional()
    .or(z.literal("")),

  description: z.string().optional().or(z.literal("")),

  clinicalIndications: z.string().optional().or(z.literal("")),

  displayOrder: z
    .number()
    .int("Ordem de exibição deve ser um número inteiro")
    .min(0, "Ordem de exibição não pode ser negativo")
    .default(0),

  isActive: z.boolean().default(true),
});

export type LabTestDefinitionFormValues = z.infer<
  typeof labTestDefinitionSchema
>;

/**
 * Get default values for the form
 */
export const getDefaultValues = (): LabTestDefinitionFormValues => ({
  code: "",
  name: "",
  shortName: "",
  tussCode: "",
  loincCode: "",
  category: "biochemistry",
  isRequestable: true,
  parentTestId: "",
  unit: "",
  unitConversion: "",
  resultType: "numeric",
  collectionMethod: "",
  fastingHours: null,
  specimenType: "",
  description: "",
  clinicalIndications: "",
  displayOrder: 0,
  isActive: true,
});

/**
 * Convert API response to form values
 */
export const apiToFormValues = (
  data: any
): LabTestDefinitionFormValues => ({
  code: data.code || "",
  name: data.name || "",
  shortName: data.shortName || "",
  tussCode: data.tussCode || "",
  loincCode: data.loincCode || "",
  category: data.category,
  isRequestable: data.isRequestable ?? true,
  parentTestId: data.parentTestId || "",
  unit: data.unit || "",
  unitConversion: data.unitConversion || "",
  resultType: data.resultType,
  collectionMethod: data.collectionMethod || "",
  fastingHours: data.fastingHours ?? null,
  specimenType: data.specimenType || "",
  description: data.description || "",
  clinicalIndications: data.clinicalIndications || "",
  displayOrder: data.displayOrder ?? 0,
  isActive: data.isActive ?? true,
});

/**
 * Convert form values to API payload
 */
export const formToApiValues = (values: LabTestDefinitionFormValues) => {
  return {
    code: values.code,
    name: values.name,
    shortName: values.shortName || undefined,
    tussCode: values.tussCode || undefined,
    loincCode: values.loincCode || undefined,
    category: values.category,
    isRequestable: values.isRequestable,
    parentTestId: values.parentTestId || undefined,
    unit: values.unit || undefined,
    unitConversion: values.unitConversion || undefined,
    resultType: values.resultType,
    collectionMethod: values.collectionMethod || undefined,
    fastingHours: values.fastingHours ?? undefined,
    specimenType: values.specimenType || undefined,
    description: values.description || undefined,
    clinicalIndications: values.clinicalIndications || undefined,
    displayOrder: values.displayOrder,
    isActive: values.isActive,
  };
};

/**
 * Category options for select
 */
export const categoryOptions: Array<{
  value: LabTestCategory;
  label: string;
}> = [
  { value: "hematology", label: "Hematologia" },
  { value: "biochemistry", label: "Bioquímica" },
  { value: "hormones", label: "Hormônios" },
  { value: "immunology", label: "Imunologia" },
  { value: "microbiology", label: "Microbiologia" },
  { value: "urine", label: "Urinálise" },
  { value: "imaging", label: "Imagem" },
  { value: "functional", label: "Medicina Funcional" },
  { value: "genetics", label: "Genética" },
  { value: "other", label: "Outros" },
];

/**
 * Result type options for select
 */
export const resultTypeOptions: Array<{
  value: ResultType;
  label: string;
  description: string;
}> = [
  {
    value: "numeric",
    label: "Numérico",
    description: "Valores quantitativos (ex: 12.5 g/dL)",
  },
  {
    value: "text",
    label: "Texto",
    description: "Valores textuais (ex: Negativo, Positivo)",
  },
  {
    value: "boolean",
    label: "Sim/Não",
    description: "Valores booleanos (ex: Presente, Ausente)",
  },
  {
    value: "categorical",
    label: "Categórico",
    description: "Valores de categoria (ex: Leve, Moderado, Grave)",
  },
];
