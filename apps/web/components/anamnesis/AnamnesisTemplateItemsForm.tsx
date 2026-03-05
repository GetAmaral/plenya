'use client'

import type { AnamnesisTemplate } from '@/lib/api/anamnesis-templates'
import type { Patient } from '@/lib/auth-store'
import { AnamnesisTemplateItemsRenderer } from './AnamnesisTemplateItemsRenderer'

// Type for the form values
export interface AnamnesisItemFormValue {
  scoreItemId: string
  numericValue?: number    // real measured value (e.g. 82.3 kg)
  selectedLevel?: number  // classified level (0-6), auto-detected or manually selected
  textValue?: string
  order: number
}

interface AnamnesisTemplateItemsFormProps {
  template: AnamnesisTemplate
  initialValues?: AnamnesisItemFormValue[]
  onChange: (values: AnamnesisItemFormValue[]) => void
  focusScoreItemId?: string | null
  patient?: Patient | null
}

export function AnamnesisTemplateItemsForm({
  template,
  initialValues = [],
  onChange,
  focusScoreItemId,
  patient,
}: AnamnesisTemplateItemsFormProps) {
  return (
    <AnamnesisTemplateItemsRenderer
      template={template}
      initialValues={initialValues}
      onChange={onChange}
      compact={true}
      focusScoreItemId={focusScoreItemId}
      patient={patient}
    />
  )
}
