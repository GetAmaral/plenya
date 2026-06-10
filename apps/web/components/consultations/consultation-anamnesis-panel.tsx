'use client'

/**
 * Painel de Anamnese DENTRO da consulta — itens (ScoreItem) linkados ao appointment.
 *
 * Sem anamnese vinculada: oferece escolher um template e iniciar (POST cria + linka via
 * Appointment.AnamnesisID). Com anamnese: renderiza o renderer denso (compact) com histórico
 * por item, e salva (autosave debounced) via PUT /anamnesis/:id. O médico não sai da consulta.
 */
import { useEffect, useRef, useState } from 'react'
import { useQuery } from '@tanstack/react-query'
import { toast } from 'sonner'
import { ClipboardList, ChevronDown, Loader2, CheckCircle2, Plus } from 'lucide-react'

import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { cn } from '@/lib/utils'
import type { Patient } from '@/lib/auth-store'
import {
  getAllAnamnesisTemplates,
  getAnamnesisTemplateById,
} from '@/lib/api/anamnesis-templates'
import {
  useAppointmentAnamnesis,
  useCreateAppointmentAnamnesis,
  useUpdateAppointmentAnamnesis,
} from '@/lib/api/appointment-anamnesis'
import { AnamnesisTemplateItemsForm } from '@/components/anamnesis/AnamnesisTemplateItemsForm'
import type { AnamnesisItemFormValue } from '@/components/anamnesis/AnamnesisTemplateItemsForm'

interface ConsultationAnamnesisPanelProps {
  appointmentId: string
  patient?: Patient | null
  // ID do paciente da consulta — usado no histórico por item (confiável mesmo se o
  // selectedPatient ainda não estiver hidratado).
  patientId: string
}

export function ConsultationAnamnesisPanel({ appointmentId, patient, patientId }: ConsultationAnamnesisPanelProps) {
  const [open, setOpen] = useState(false)
  const anamnesisQuery = useAppointmentAnamnesis(appointmentId)
  const anamnesis = anamnesisQuery.data ?? null

  const createAnamnesis = useCreateAppointmentAnamnesis(appointmentId)
  const updateAnamnesis = useUpdateAppointmentAnamnesis(appointmentId)

  // Templates (para iniciar quando ainda não há anamnese vinculada).
  const templatesQuery = useQuery({
    queryKey: ['anamnesis-templates', 'list'],
    queryFn: () => getAllAnamnesisTemplates(false),
    enabled: open && !anamnesis && !anamnesisQuery.isLoading,
  })

  // Template (com itens) da anamnese existente — necessário p/ renderizar o renderer.
  const templateId = anamnesis?.anamnesisTemplateId
  const templateQuery = useQuery({
    queryKey: ['anamnesis-template', templateId],
    queryFn: () => getAnamnesisTemplateById(templateId!, true),
    enabled: !!templateId,
  })

  const initialValues: AnamnesisItemFormValue[] = (anamnesis?.items ?? []).map((it) => ({
    scoreItemId: it.scoreItemId,
    numericValue: it.numericValue,
    selectedLevel: it.selectedLevel,
    textValue: it.textValue,
    order: it.order,
  }))

  // Autosave debounced dos itens.
  const [saveState, setSaveState] = useState<'idle' | 'saving' | 'saved'>('idle')
  const saveTimer = useRef<ReturnType<typeof setTimeout> | null>(null)
  const latestValues = useRef<AnamnesisItemFormValue[]>([])

  useEffect(() => () => { if (saveTimer.current) clearTimeout(saveTimer.current) }, [])

  function handleItemsChange(values: AnamnesisItemFormValue[]) {
    latestValues.current = values
    if (!anamnesis) return
    if (saveTimer.current) clearTimeout(saveTimer.current)
    setSaveState('saving')
    saveTimer.current = setTimeout(() => {
      updateAnamnesis.mutate(
        { id: anamnesis.id, data: { items: latestValues.current } },
        {
          onSuccess: () => setSaveState('saved'),
          onError: (e) => {
            setSaveState('idle')
            toast.error('Erro ao salvar a anamnese', {
              description: e instanceof Error ? e.message : undefined,
            })
          },
        }
      )
    }, 1200)
  }

  async function handleStart(templateIdToUse: string) {
    try {
      await createAnamnesis.mutateAsync({
        anamnesisTemplateId: templateIdToUse,
        consultationDate: new Date().toISOString(),
        visibility: 'all',
        items: [],
      })
      toast.success('Anamnese iniciada nesta consulta')
    } catch (e) {
      toast.error('Erro ao iniciar a anamnese', {
        description: e instanceof Error ? e.message : undefined,
      })
    }
  }

  const filledCount = anamnesis?.items?.length ?? 0

  return (
    <Card>
      <CardHeader
        className="flex cursor-pointer flex-row items-center justify-between space-y-0 pb-2"
        onClick={() => setOpen((v) => !v)}
      >
        <CardTitle className="flex items-center gap-2 text-sm">
          <ClipboardList className="h-4 w-4 text-primary" />
          Anamnese da consulta
          {anamnesis && (
            <span className="rounded-full bg-muted px-2 py-0.5 text-[11px] font-normal text-muted-foreground">
              {filledCount} {filledCount === 1 ? 'item' : 'itens'}
            </span>
          )}
          {saveState === 'saving' && (
            <span className="flex items-center gap-1 text-[11px] font-normal text-muted-foreground">
              <Loader2 className="h-3 w-3 animate-spin" /> salvando
            </span>
          )}
          {saveState === 'saved' && (
            <span className="flex items-center gap-1 text-[11px] font-normal text-emerald-700">
              <CheckCircle2 className="h-3 w-3" /> salvo
            </span>
          )}
        </CardTitle>
        <ChevronDown className={cn('h-4 w-4 text-muted-foreground transition-transform', open && 'rotate-180')} />
      </CardHeader>

      {open && (
        <CardContent className="space-y-3">
          {anamnesisQuery.isLoading ? (
            <div className="flex items-center gap-2 py-2 text-sm text-muted-foreground">
              <Loader2 className="h-4 w-4 animate-spin" /> Carregando…
            </div>
          ) : !anamnesis ? (
            // Sem anamnese: escolher template e iniciar.
            <div className="space-y-2">
              <p className="text-xs text-muted-foreground">
                Nenhuma anamnese nesta consulta. Escolha um modelo para iniciar:
              </p>
              {templatesQuery.isLoading ? (
                <div className="flex items-center gap-2 py-1 text-sm text-muted-foreground">
                  <Loader2 className="h-4 w-4 animate-spin" /> Carregando modelos…
                </div>
              ) : (templatesQuery.data ?? []).length === 0 ? (
                <p className="text-xs text-muted-foreground">Nenhum modelo de anamnese disponível.</p>
              ) : (
                <div className="flex flex-col gap-1.5">
                  {(templatesQuery.data ?? []).map((t) => (
                    <Button
                      key={t.id}
                      variant="outline"
                      size="sm"
                      className="justify-start"
                      disabled={createAnamnesis.isPending}
                      onClick={() => handleStart(t.id)}
                    >
                      {createAnamnesis.isPending ? (
                        <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                      ) : (
                        <Plus className="mr-2 h-4 w-4" />
                      )}
                      {t.name}
                    </Button>
                  ))}
                </div>
              )}
            </div>
          ) : templateQuery.isLoading ? (
            <div className="flex items-center gap-2 py-2 text-sm text-muted-foreground">
              <Loader2 className="h-4 w-4 animate-spin" /> Carregando itens…
            </div>
          ) : !templateQuery.data ? (
            <p className="text-xs text-muted-foreground">
              Não foi possível carregar o modelo desta anamnese.
            </p>
          ) : (
            <AnamnesisTemplateItemsForm
              template={templateQuery.data}
              initialValues={initialValues}
              onChange={handleItemsChange}
              patient={patient}
              patientId={patientId}
            />
          )}
        </CardContent>
      )}
    </Card>
  )
}
