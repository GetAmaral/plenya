'use client'

import type { AnamnesisTemplate } from '@/lib/api/anamnesis-templates'
import { AnamnesisTemplateItemsRenderer } from './AnamnesisTemplateItemsRenderer'

// Type for the form values
export interface AnamnesisItemFormValue {
  scoreItemId: string
  numericValue?: number
  textValue?: string
  order: number
}

interface AnamnesisTemplateItemsFormProps {
  template: AnamnesisTemplate
  initialValues?: AnamnesisItemFormValue[]
  onChange: (values: AnamnesisItemFormValue[]) => void
  focusScoreItemId?: string | null
}

export function AnamnesisTemplateItemsForm({
  template,
  initialValues = [],
  onChange,
  focusScoreItemId,
}: AnamnesisTemplateItemsFormProps) {
  return (
    <AnamnesisTemplateItemsRenderer
      template={template}
      initialValues={initialValues}
      onChange={onChange}
      compact={true}
      focusScoreItemId={focusScoreItemId}
    />
  )
}
