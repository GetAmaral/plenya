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

/** Artigo (resposta da API). */
export type Article = WithRequired<
  Schemas['models.Article'],
  'id' | 'title' | 'articleType' | 'publishDate' | 'createdAt' | 'updatedAt'
>

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
export type AnamnesisItem = Refine<
  WithRequired<
    Schemas['models.AnamnesisItem'],
    'id' | 'scoreItemId' | 'order' | 'createdAt' | 'updatedAt'
  >,
  { scoreItem?: ScoreItem }
>
export type AnamnesisTemplate = Refine<
  WithRequired<Schemas['models.AnamnesisTemplate'], 'id' | 'name' | 'area' | 'createdAt' | 'updatedAt'>,
  { items?: AnamnesisTemplateItem[] }
>
export type AnamnesisTemplateItem = Refine<
  WithRequired<
    Schemas['models.AnamnesisTemplateItem'],
    'id' | 'anamnesisTemplateId' | 'scoreItemId' | 'order' | 'createdAt' | 'updatedAt'
  >,
  { scoreItem?: ScoreItem }
>

/** Snapshot de escore do paciente (resposta da API). */
export type PatientScoreSnapshot = Refine<
  WithRequired<
    Schemas['models.PatientScoreSnapshot'],
    | 'id' | 'patientId' | 'calculatedByUserId' | 'calculatedAt' | 'totalActualPoints'
    | 'totalPossiblePoints' | 'totalScorePercentage' | 'itemsEvaluatedCount'
    | 'itemsNotEvaluatedCount' | 'createdAt' | 'updatedAt'
  >,
  // o detalhe do snapshot sempre traz os resultados materializados (os consumidores assumem presença).
  { groupResults: PatientScoreGroupResult[]; itemResults: PatientScoreItemResult[] }
>
export type PatientScoreGroupResult = Refine<
  WithRequired<
    Schemas['models.PatientScoreGroupResult'],
    | 'id' | 'snapshotId' | 'groupId' | 'actualPoints' | 'possiblePoints'
    | 'scorePercentage' | 'itemsEvaluatedCount' | 'createdAt' | 'updatedAt'
  >,
  { group?: ScoreGroup }
>
export type PatientScoreItemResult = Refine<
  WithRequired<
    Schemas['models.PatientScoreItemResult'],
    'id' | 'snapshotId' | 'itemId' | 'groupId' | 'status' | 'maxPoints' | 'actualPoints' | 'createdAt' | 'updatedAt'
  >,
  { group?: ScoreGroup; item?: ScoreItem }
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


/** Escore: Group → Subgroup → Item → Level + Method/Letter/Pillar (resposta da API).
 *  Refinam as relações aninhadas pros próprios aliases (o gerado aponta pro tipo cru). */
export type ScoreLevel = Refine<
  WithRequired<
    Schemas['models.ScoreLevel'],
    'id' | 'itemId' | 'level' | 'name' | 'operator' | 'createdAt' | 'updatedAt'
  >,
  // o swag emite `operator: string`, mas os valores Go são restritos a este conjunto.
  { operator: '=' | '>' | '>=' | '<' | '<=' | 'between' }
>
export type ScoreItem = Refine<
  WithRequired<
    Schemas['models.ScoreItem'],
    'id' | 'name' | 'order' | 'subgroupId' | 'createdAt' | 'updatedAt'
  >,
  {
    subgroup?: ScoreSubgroup
    levels?: ScoreLevel[]
    childItems?: ScoreItem[]
    parentItem?: ScoreItem
    methodPillars?: MethodPillar[]
  }
>
export type ScoreSubgroup = Refine<
  WithRequired<
    Schemas['models.ScoreSubgroup'],
    'id' | 'name' | 'order' | 'maxSelect' | 'groupId' | 'createdAt' | 'updatedAt'
  >,
  { group?: ScoreGroup; items?: ScoreItem[] }
>
export type ScoreGroup = Refine<
  WithRequired<Schemas['models.ScoreGroup'], 'id' | 'name' | 'order' | 'createdAt' | 'updatedAt'>,
  { subgroups?: ScoreSubgroup[] }
>
export type MethodPillar = Refine<
  WithRequired<
    Schemas['models.MethodPillar'],
    'id' | 'letterId' | 'name' | 'order' | 'createdAt' | 'updatedAt'
  >,
  { scoreItems?: ScoreItem[]; letter?: MethodLetter }
>
export type MethodLetter = Refine<
  WithRequired<
    Schemas['models.MethodLetter'],
    'id' | 'methodId' | 'code' | 'name' | 'order' | 'createdAt' | 'updatedAt'
  >,
  { pillars?: MethodPillar[]; method?: Method }
>
export type Method = Refine<
  WithRequired<
    Schemas['models.Method'],
    'id' | 'name' | 'shortName' | 'order' | 'isDefault' | 'createdAt' | 'updatedAt'
  >,
  { letters?: MethodLetter[] }
>

/** Prescrição (resposta da API = dto.PrescriptionResponse, não o model cru). */
export type MedicationResponse = WithRequired<
  Schemas['dto.MedicationResponse'],
  | 'id' | 'medicationName' | 'activeIngredient' | 'category' | 'concentration'
  | 'dosage' | 'frequency' | 'route' | 'duration' | 'quantity'
>
/** Fórmula magistral (receita de manipulado). */
export type FormulaComponentResponse = WithRequired<
  Schemas['dto.FormulaComponentResponse'],
  'id' | 'substance' | 'quantity' | 'unit' | 'category'
>
export type FormulaResponse = Refine<
  WithRequired<
    Schemas['dto.FormulaResponse'],
    | 'id' | 'pharmaceuticalForm' | 'usageType' | 'quantityToDispense'
    | 'quantityUnit' | 'posology' | 'highestCategory'
  >,
  { components: FormulaComponentResponse[] }
>
export type Prescription = Refine<
  WithRequired<
    Schemas['dto.PrescriptionResponse'],
    | 'id' | 'patientId' | 'doctorId' | 'status' | 'prescriptionDate'
    | 'validUntil' | 'isUsed' | 'createdAt' | 'updatedAt' | 'type'
  >,
  // o backend sempre devolve as duas listas (vazias quando não se aplicam ao tipo)
  { medications: MedicationResponse[]; formulas: FormulaResponse[] }
>

/** Template de pedido de exames (resposta da API). */
export type LabRequestTemplate = Refine<
  WithRequired<Schemas['models.LabRequestTemplate'], 'id' | 'displayOrder' | 'isActive' | 'createdAt' | 'updatedAt'>,
  // o schema gerado aponta labTests pro tipo cru (id opcional); refina pro mesmo alias acima
  { labTests?: LabTestDefinition[] }
>

