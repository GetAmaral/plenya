// Aliases ergonômicos sobre os tipos GERADOS (openapi-typescript a partir dos Go models).
// Regra: a base é sempre o schema gerado — campos novos do model propagam automaticamente.
// Só refinamos a obrigatoriedade dos campos que a RESPOSTA da API sempre traz (o swag os marca
// opcionais por causa do `omitempty`, mas em resposta eles existem). Nunca redeclarar campos à mão.
import type { components } from './generated/api-types'

type Schemas = components['schemas']

/** Marca as chaves K de T como obrigatórias, preservando o resto do tipo gerado. */
type WithRequired<T, K extends keyof T> = Omit<T, K> & Required<Pick<T, K>>

/** Campo JSON do backend (datatypes.JSONMap) — o swag emite `unknown`; aqui é objeto. */
type Json = Record<string, unknown>

/** Sobrescreve em T as propriedades declaradas em R (Omit + re-add — evita interseção de
 *  arrays, que faz `.map` resolver pelo elemento cru). Use p/ refinar relações aninhadas. */
type Refine<T, R> = Omit<T, keyof R> & R

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
  'id' | 'isActive' | 'isRequestable' | 'displayOrder' | 'createdAt' | 'updatedAt'
>

/** Anamnese (resposta da API). */
export type Anamnesis = Refine<
  WithRequired<
    Schemas['models.Anamnesis'],
    'id' | 'patientId' | 'authorId' | 'consultationDate' | 'visibility' | 'createdAt' | 'updatedAt'
  >,
  { items?: AnamnesisItem[] }
>
export type AnamnesisItem = WithRequired<
  Schemas['models.AnamnesisItem'],
  'id' | 'scoreItemId' | 'order' | 'createdAt' | 'updatedAt'
>
export type AnamnesisTemplate = Refine<
  WithRequired<Schemas['models.AnamnesisTemplate'], 'id' | 'name' | 'area' | 'createdAt' | 'updatedAt'>,
  { items?: AnamnesisTemplateItem[] }
>
export type AnamnesisTemplateItem = WithRequired<
  Schemas['models.AnamnesisTemplateItem'],
  'id' | 'anamnesisTemplateId' | 'scoreItemId' | 'order' | 'createdAt' | 'updatedAt'
>

/** Snapshot de escore do paciente (resposta da API). */
export type PatientScoreSnapshot = Refine<
  WithRequired<
    Schemas['models.PatientScoreSnapshot'],
    | 'id' | 'patientId' | 'calculatedByUserId' | 'calculatedAt' | 'totalActualPoints'
    | 'totalPossiblePoints' | 'totalScorePercentage' | 'itemsEvaluatedCount'
    | 'itemsNotEvaluatedCount' | 'createdAt' | 'updatedAt'
  >,
  { groupResults?: PatientScoreGroupResult[]; itemResults?: PatientScoreItemResult[] }
>
export type PatientScoreGroupResult = WithRequired<
  Schemas['models.PatientScoreGroupResult'],
  | 'id' | 'snapshotId' | 'groupId' | 'actualPoints' | 'possiblePoints'
  | 'scorePercentage' | 'itemsEvaluatedCount' | 'createdAt' | 'updatedAt'
>
export type PatientScoreItemResult = WithRequired<
  Schemas['models.PatientScoreItemResult'],
  'id' | 'snapshotId' | 'itemId' | 'groupId' | 'status' | 'maxPoints' | 'actualPoints' | 'createdAt' | 'updatedAt'
>

/** Pedido de exames (resposta da API). */
export type LabRequest = WithRequired<
  Schemas['models.LabRequest'],
  'id' | 'patientId' | 'date' | 'exams' | 'createdAt' | 'updatedAt'
>

/** View de resultados (resposta da API). */
export type LabResultView = Refine<
  WithRequired<Schemas['models.LabResultView'], 'id' | 'name' | 'isActive' | 'displayOrder'>,
  { items?: LabResultViewItem[] }
>
export type LabResultViewItem = WithRequired<
  Schemas['models.LabResultViewItem'],
  'id' | 'labResultViewId' | 'labTestDefinitionId' | 'order'
>
export type LabResultValue = WithRequired<
  Schemas['models.LabResultValue'],
  'id' | 'labResultId' | 'labTestDefinitionId' | 'createdAt' | 'updatedAt'
>

/** Lead / atividade de lead (resposta da API). */
export type LeadActivity = Refine<
  WithRequired<Schemas['models.LeadActivity'], 'id' | 'leadId' | 'type' | 'channel' | 'createdAt'>,
  { metadata?: Json }
>
export type Lead = Refine<
  WithRequired<
    Schemas['models.Lead'],
    'id' | 'source' | 'status' | 'emailOptIn' | 'whatsAppOptIn' | 'newsletterOptIn' | 'createdAt' | 'updatedAt'
  >,
  // refina relações aninhadas (activities) e o JSON
  { activities?: LeadActivity[]; metadata?: Json }
>


/** Template de pedido de exames (resposta da API). */
export type LabRequestTemplate = Refine<
  WithRequired<Schemas['models.LabRequestTemplate'], 'id' | 'displayOrder' | 'isActive' | 'createdAt' | 'updatedAt'>,
  // o schema gerado aponta labTests pro tipo cru (id opcional); refina pro mesmo alias acima
  { labTests?: LabTestDefinition[] }
>

