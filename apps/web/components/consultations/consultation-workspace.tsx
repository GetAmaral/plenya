'use client';

/**
 * Workspace de Consulta — tela única do atendimento (P0 do plano UX do médico).
 *
 * 3 zonas:
 *  - Esquerda: SÍNTESE do paciente (Escore Plenya + pilares AGIR + motivo).
 *  - Centro: NOTA DE EVOLUÇÃO SOAP/APSO (rich-text, reusa RichTextEditor).
 *  - Direita: AÇÕES clínicas (prescrever, pedir exame, escore, retorno) +
 *    Finalizar consulta.
 *
 * Ao montar, seleciona o paciente da consulta (selectedPatient) — assim os
 * fluxos de prescrição/exame resolvem o patientId pelo backend sem mudança.
 * Nota assinada é read-only (correção via adendo, fora do P0).
 */
import { useEffect, useRef, useState } from 'react';
import { useRouter } from 'next/navigation';
import { toast } from 'sonner';
import DOMPurify from 'dompurify';
import {
  Activity,
  CalendarClock,
  CheckCircle2,
  FlaskConical,
  Loader2,
  Lock,
  Pill,
  Save,
  Stethoscope,
} from 'lucide-react';

import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Badge } from '@/components/ui/badge';
import { RichTextEditor } from '@/components/ui/rich-text-editor';
import { ScheduleRecallDialog } from '@/components/recepcao/schedule-recall-dialog';
import { useSelectedPatient } from '@/lib/use-selected-patient';
import { useLatestHealthScore } from '@/lib/api/health-score-api';
import {
  type Appointment,
  useStartAppointment,
  useUpdateAppointment,
} from '@/lib/api/calendar-api';
import {
  type ClinicalNoteLayout,
  useClinicalNoteByAppointment,
  useCreateClinicalNote,
  useSignClinicalNote,
  useUpdateClinicalNote,
} from '@/lib/api/clinical-notes';

const SECTION_LABELS: Record<string, { label: string; placeholder: string }> = {
  subjective: { label: 'Subjetivo', placeholder: 'Queixa, história da doença atual, relato do paciente...' },
  objective: { label: 'Objetivo', placeholder: 'Exame físico, sinais vitais, achados objetivos...' },
  assessment: { label: 'Avaliação', placeholder: 'Impressão clínica, hipóteses, raciocínio...' },
  plan: { label: 'Plano', placeholder: 'Conduta, prescrição, exames, orientações, retorno...' },
};

const LAYOUT_ORDER: Record<ClinicalNoteLayout, Array<keyof typeof SECTION_LABELS>> = {
  soap: ['subjective', 'objective', 'assessment', 'plan'],
  apso: ['assessment', 'plan', 'subjective', 'objective'],
};

export function ConsultationWorkspace({ appt }: { appt: Appointment }) {
  const router = useRouter();
  const { selectedPatientId, setSelectedPatient } = useSelectedPatient();

  // Seleciona o paciente da consulta ao abrir (uma vez), pra dar contexto aos
  // fluxos embutidos (prescrição/exame). Evita re-disparo via ref.
  const autoSelected = useRef(false);
  useEffect(() => {
    if (!autoSelected.current && appt.patientId && selectedPatientId !== appt.patientId) {
      autoSelected.current = true;
      setSelectedPatient(appt.patientId);
    }
  }, [appt.patientId, selectedPatientId, setSelectedPatient]);

  const { data: score } = useLatestHealthScore(appt.patientId);

  const noteQuery = useClinicalNoteByAppointment(appt.id);
  const note = noteQuery.data ?? null;
  const createNote = useCreateClinicalNote();
  const updateNote = useUpdateClinicalNote();
  const signNote = useSignClinicalNote();
  const startAppt = useStartAppointment(appt.id);
  const completeAppt = useUpdateAppointment(appt.id);

  const [layout, setLayout] = useState<ClinicalNoteLayout>('soap');
  const [subjective, setSubjective] = useState('');
  const [objective, setObjective] = useState('');
  const [assessment, setAssessment] = useState('');
  const [plan, setPlan] = useState('');
  const [showRecall, setShowRecall] = useState(false);

  // Hidrata os campos a partir da nota existente (uma vez).
  const hydrated = useRef(false);
  useEffect(() => {
    if (note && !hydrated.current) {
      hydrated.current = true;
      setLayout(note.layout ?? 'soap');
      setSubjective(note.subjectiveHtml ?? '');
      setObjective(note.objectiveHtml ?? '');
      setAssessment(note.assessmentHtml ?? '');
      setPlan(note.planHtml ?? '');
    }
  }, [note]);

  const isSigned = note?.status === 'signed';
  const readOnly = isSigned;
  const saving = createNote.isPending || updateNote.isPending;
  const finalizing = signNote.isPending || completeAppt.isPending;

  const htmlBySection: Record<string, string> = {
    subjective,
    objective,
    assessment,
    plan,
  };
  const setterBySection: Record<string, (v: string) => void> = {
    subjective: setSubjective,
    objective: setObjective,
    assessment: setAssessment,
    plan: setPlan,
  };

  async function saveDraft(): Promise<string | null> {
    const payload = {
      layout,
      subjectiveHtml: subjective,
      objectiveHtml: objective,
      assessmentHtml: assessment,
      planHtml: plan,
    };
    try {
      if (note) {
        const updated = await updateNote.mutateAsync({ id: note.id, payload });
        return updated.id;
      }
      const created = await createNote.mutateAsync({ appointmentId: appt.id, ...payload });
      return created.id;
    } catch (err) {
      toast.error('Erro ao salvar a nota', {
        description: err instanceof Error ? err.message : undefined,
      });
      return null;
    }
  }

  async function handleSaveDraft() {
    const id = await saveDraft();
    if (id) toast.success('Rascunho salvo');
  }

  // Salva o rascunho antes de navegar pra prescrição/exame (não perde o que foi
  // digitado), depois abre o fluxo existente já com o paciente selecionado.
  async function navigateAfterSave(href: string) {
    if (!readOnly) await saveDraft();
    router.push(href);
  }

  async function handleFinalize() {
    const id = await saveDraft();
    if (!id) return;
    try {
      await signNote.mutateAsync(id);
    } catch (err) {
      toast.error('Erro ao assinar a nota', {
        description: err instanceof Error ? err.message : undefined,
      });
      return;
    }
    try {
      await completeAppt.mutateAsync({ status: 'completed' });
    } catch (err) {
      toast.error('Nota assinada, mas falha ao concluir a consulta', {
        description: err instanceof Error ? err.message : 'Tente concluir novamente.',
      });
      return;
    }
    toast.success('Consulta finalizada e documentada');
  }

  async function handleSignOnly() {
    const id = await saveDraft();
    if (!id) return;
    try {
      await signNote.mutateAsync(id);
      toast.success('Nota assinada');
    } catch (err) {
      toast.error('Erro ao assinar a nota', {
        description: err instanceof Error ? err.message : undefined,
      });
    }
  }

  async function handleStart() {
    try {
      await startAppt.mutateAsync();
      toast.success('Atendimento iniciado');
    } catch (err) {
      toast.error('Erro ao iniciar o atendimento', {
        description: err instanceof Error ? err.message : undefined,
      });
    }
  }

  const totalScore = score?.totalScorePercentage;

  return (
    <div className="grid gap-4 lg:grid-cols-[18rem_minmax(0,1fr)_15rem]">
      {/* ESQUERDA — Síntese */}
      <div className="space-y-4">
        <Card>
          <CardHeader className="pb-2">
            <CardTitle className="flex items-center gap-2 text-sm">
              <Activity className="h-4 w-4 text-primary" />
              Síntese clínica
            </CardTitle>
          </CardHeader>
          <CardContent className="space-y-3">
            {typeof totalScore === 'number' ? (
              <div>
                <p className="text-xs uppercase text-muted-foreground">Escore Plenya</p>
                <p className="text-2xl font-semibold">{totalScore.toFixed(0)}%</p>
              </div>
            ) : (
              <p className="text-xs text-muted-foreground">
                Sem escore calculado ainda.
              </p>
            )}

            {score?.groupResults && score.groupResults.length > 0 && (
              <div className="space-y-1.5">
                <p className="text-xs uppercase text-muted-foreground">Pilares AGIR</p>
                {score.groupResults
                  .slice()
                  .sort((a, b) => (a.group?.order ?? 0) - (b.group?.order ?? 0))
                  .map((g) => (
                    <div key={g.id} className="space-y-0.5">
                      <div className="flex items-center justify-between text-xs">
                        <span className="truncate">{g.group?.name ?? 'Grupo'}</span>
                        <span className="text-muted-foreground">
                          {g.scorePercentage.toFixed(0)}%
                        </span>
                      </div>
                      <div className="h-1.5 w-full overflow-hidden rounded-full bg-muted">
                        <div
                          className="h-full rounded-full bg-primary"
                          style={{ width: `${Math.min(100, Math.max(0, g.scorePercentage))}%` }}
                        />
                      </div>
                    </div>
                  ))}
              </div>
            )}

            <div>
              <p className="text-xs uppercase text-muted-foreground">Motivo</p>
              <p className="text-sm">{appt.reason}</p>
            </div>
            {appt.patientNotes && (
              <div>
                <p className="text-xs uppercase text-muted-foreground">Obs. do paciente</p>
                <p className="text-sm text-muted-foreground">{appt.patientNotes}</p>
              </div>
            )}

            <Button
              variant="link"
              className="h-auto p-0 text-xs"
              onClick={() => router.push(`/patients/${appt.patientId}/prontuario`)}
            >
              Ver prontuário completo →
            </Button>
          </CardContent>
        </Card>
      </div>

      {/* CENTRO — Nota de evolução */}
      <Card>
        <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
          <CardTitle className="flex items-center gap-2 text-sm">
            <Stethoscope className="h-4 w-4 text-primary" />
            Evolução
            {isSigned ? (
              <Badge variant="outline" className="border-emerald-200 bg-emerald-50 text-emerald-900">
                <Lock className="mr-1 h-3 w-3" />
                Assinada
              </Badge>
            ) : (
              <Badge variant="outline">Rascunho</Badge>
            )}
          </CardTitle>
          {!readOnly && (
            <div className="flex rounded-md border p-0.5 text-xs">
              {(['soap', 'apso'] as ClinicalNoteLayout[]).map((l) => (
                <button
                  key={l}
                  type="button"
                  onClick={() => setLayout(l)}
                  className={`rounded px-2 py-0.5 uppercase ${
                    layout === l ? 'bg-primary text-primary-foreground' : 'text-muted-foreground'
                  }`}
                >
                  {l}
                </button>
              ))}
            </div>
          )}
        </CardHeader>
        <CardContent className="space-y-4">
          {LAYOUT_ORDER[layout].map((key) => (
            <div key={key} className="space-y-1">
              <p className="text-xs font-medium uppercase text-muted-foreground">
                {SECTION_LABELS[key].label}
              </p>
              {readOnly ? (
                <div
                  className="prose prose-sm max-w-none rounded-md border bg-muted/30 p-3 text-sm"
                  dangerouslySetInnerHTML={{
                    __html: DOMPurify.sanitize(htmlBySection[key] || '<p class="text-muted-foreground">—</p>'),
                  }}
                />
              ) : (
                <RichTextEditor
                  editorId={`note-${key}`}
                  value={htmlBySection[key]}
                  onChange={setterBySection[key]}
                  placeholder={SECTION_LABELS[key].placeholder}
                  minHeight="120px"
                />
              )}
            </div>
          ))}

          {!readOnly && (
            <div className="flex justify-end">
              <Button variant="outline" size="sm" onClick={handleSaveDraft} disabled={saving}>
                {saving ? <Loader2 className="mr-2 h-4 w-4 animate-spin" /> : <Save className="mr-2 h-4 w-4" />}
                Salvar rascunho
              </Button>
            </div>
          )}
        </CardContent>
      </Card>

      {/* DIREITA — Ações clínicas */}
      <div className="space-y-4">
        <Card>
          <CardHeader className="pb-2">
            <CardTitle className="text-sm">Ações</CardTitle>
          </CardHeader>
          <CardContent className="space-y-2">
            {appt.status === 'checked_in' && (
              <Button className="w-full" onClick={handleStart} disabled={startAppt.isPending}>
                {startAppt.isPending && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
                Iniciar atendimento
              </Button>
            )}

            <Button variant="outline" className="w-full justify-start" onClick={() => navigateAfterSave('/prescriptions/new')}>
              <Pill className="mr-2 h-4 w-4" />
              Prescrever
            </Button>
            <Button variant="outline" className="w-full justify-start" onClick={() => navigateAfterSave('/lab-requests')}>
              <FlaskConical className="mr-2 h-4 w-4" />
              Pedir exame
            </Button>
            <Button variant="outline" className="w-full justify-start" onClick={() => navigateAfterSave('/health-scores')}>
              <Activity className="mr-2 h-4 w-4" />
              Escore
            </Button>
            <Button variant="outline" className="w-full justify-start" onClick={() => setShowRecall(true)}>
              <CalendarClock className="mr-2 h-4 w-4" />
              Agendar retorno
            </Button>

            {!isSigned && appt.status !== 'completed' && (
              <Button className="w-full" onClick={handleFinalize} disabled={finalizing || saving}>
                {finalizing ? <Loader2 className="mr-2 h-4 w-4 animate-spin" /> : <CheckCircle2 className="mr-2 h-4 w-4" />}
                Finalizar consulta
              </Button>
            )}
            {!isSigned && appt.status === 'completed' && (
              <Button className="w-full" onClick={handleSignOnly} disabled={finalizing || saving}>
                {finalizing ? <Loader2 className="mr-2 h-4 w-4 animate-spin" /> : <CheckCircle2 className="mr-2 h-4 w-4" />}
                Assinar nota
              </Button>
            )}
            {isSigned && (
              <p className="text-center text-xs text-muted-foreground">
                Nota assinada{note?.signedAt ? '' : ''}. Correções via adendo.
              </p>
            )}
          </CardContent>
        </Card>
      </div>

      <ScheduleRecallDialog
        open={showRecall}
        onOpenChange={setShowRecall}
        patientId={appt.patientId}
        doctorId={appt.doctorId}
        sourceAppointmentId={appt.id}
        patientName={appt.patient?.name ?? ''}
      />
    </div>
  );
}
