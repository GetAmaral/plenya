'use client'

import type { AnamnesisTemplate } from '@/lib/api/anamnesis-templates'
import type { AnamnesisItemFormValue } from './AnamnesisTemplateItemsForm'
import type { Patient } from '@/lib/auth-store'
import { AnamnesisTemplateItemsRenderer } from './AnamnesisTemplateItemsRenderer'

interface AnamnesisFullscreenTemplateItemsProps {
  template: AnamnesisTemplate
  initialValues?: AnamnesisItemFormValue[]
  onChange: (values: AnamnesisItemFormValue[]) => void
  patient?: Patient | null
}

export function AnamnesisFullscreenTemplateItems({
  template,
  initialValues = [],
  onChange,
  patient,
}: AnamnesisFullscreenTemplateItemsProps) {
  return (
    <AnamnesisTemplateItemsRenderer
      template={template}
      initialValues={initialValues}
      onChange={onChange}
      compact={false}
      patient={patient}
    />
  )
}
