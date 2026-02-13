'use client'

import type { AnamnesisTemplate } from '@/lib/api/anamnesis-templates'
import type { AnamnesisItemFormValue } from './AnamnesisTemplateItemsForm'
import { AnamnesisTemplateItemsRenderer } from './AnamnesisTemplateItemsRenderer'

interface AnamnesisFullscreenTemplateItemsProps {
  template: AnamnesisTemplate
  initialValues?: AnamnesisItemFormValue[]
  onChange: (values: AnamnesisItemFormValue[]) => void
}

export function AnamnesisFullscreenTemplateItems({
  template,
  initialValues = [],
  onChange,
}: AnamnesisFullscreenTemplateItemsProps) {
  return (
    <AnamnesisTemplateItemsRenderer
      template={template}
      initialValues={initialValues}
      onChange={onChange}
      compact={false}
    />
  )
}
