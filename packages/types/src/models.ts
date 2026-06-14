// Aliases ergonômicos sobre os tipos GERADOS (openapi-typescript a partir dos Go models).
// Regra: a base é sempre o schema gerado — campos novos do model propagam automaticamente.
// Só refinamos a obrigatoriedade dos campos que a RESPOSTA da API sempre traz (o swag os marca
// opcionais por causa do `omitempty`, mas em resposta eles existem). Nunca redeclarar campos à mão.
import type { components } from './generated/api-types'

type Schemas = components['schemas']

/** Marca as chaves K de T como obrigatórias, preservando o resto do tipo gerado. */
type WithRequired<T, K extends keyof T> = Omit<T, K> & Required<Pick<T, K>>

/** Paciente (resposta da API). */
export type Patient = Schemas['models.Patient']

/** Definição de medicamento (resposta da API). */
export type MedicationDefinition = WithRequired<
  Schemas['models.MedicationDefinition'],
  | 'id' | 'commonName' | 'activeIngredient' | 'category' | 'validityDays'
  | 'maxPerPrescription' | 'maxTreatmentDays' | 'requiresDigitalSignature'
  | 'requiresSNCR' | 'createdAt' | 'updatedAt'
>

/** Definição de exame laboratorial (resposta da API). */
export type LabTestDefinition = WithRequired<
  Schemas['models.LabTestDefinition'],
  'id' | 'isActive' | 'isRequestable'
>

/** Template de pedido de exames (resposta da API). */
export type LabRequestTemplate = WithRequired<
  Schemas['models.LabRequestTemplate'],
  'id' | 'displayOrder' | 'isActive' | 'createdAt' | 'updatedAt'
> & {
  // o schema gerado aponta labTests pro tipo cru (id opcional); refina pro mesmo alias acima
  labTests?: LabTestDefinition[]
}
