import { z } from 'zod'

export const medicationDefinitionSchema = z.object({
  commonName: z
    .string()
    .min(3, 'Nome deve ter no mínimo 3 caracteres')
    .max(500, 'Nome deve ter no máximo 500 caracteres'),
  activeIngredient: z
    .string()
    .min(3, 'Princípio ativo deve ter no mínimo 3 caracteres')
    .max(500, 'Princípio ativo deve ter no máximo 500 caracteres'),
  category: z.enum(['simple', 'c1', 'c5', 'antibiotic', 'glp1'], {
    required_error: 'Selecione uma categoria',
  }),
  validityDays: z.coerce
    .number()
    .int('Deve ser um número inteiro')
    .min(1, 'Validade deve ser no mínimo 1 dia')
    .max(365, 'Validade deve ser no máximo 365 dias'),
  maxPerPrescription: z.coerce
    .number()
    .int('Deve ser um número inteiro')
    .min(1, 'Quantidade mínima é 1')
    .max(100, 'Quantidade máxima é 100'),
  maxTreatmentDays: z.coerce
    .number()
    .int('Deve ser um número inteiro')
    .min(1, 'Duração mínima é 1 dia')
    .max(365, 'Duração máxima é 365 dias'),
  requiresDigitalSignature: z.boolean(),
  requiresSNCR: z.boolean(),
  anvisaCode: z.string().optional(),
})

export type MedicationDefinitionInput = z.infer<typeof medicationDefinitionSchema>
