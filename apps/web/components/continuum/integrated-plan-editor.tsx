'use client';

/**
 * IntegratedPlanEditor — markdown colaborativo do plano integrado do Continuum.
 *
 * Mostra editor com last-saved indicator, botão Salvar (cria revisão imutável)
 * e botão "Histórico" que abre dialog com todas as revisões cronológicas.
 *
 * Sem locking concorrente — UI alerta "última edição há X" pra mitigar.
 */
import { useEffect, useState } from 'react';
import { History, Loader2, Save } from 'lucide-react';
import { toast } from 'sonner';

import { Button } from '@/components/ui/button';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Textarea } from '@/components/ui/textarea';
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog';
import {
  useIntegratedPlanRevisions,
  useUpdateIntegratedPlan,
  type PatientContinuum,
} from '@/lib/api/continuum-api';
import { formatDate, formatRelativeToNow } from '@/lib/format-date';

interface Props {
  enrollment: PatientContinuum;
  canEdit: boolean;
}

export function IntegratedPlanEditor({ enrollment, canEdit }: Props) {
  const [draft, setDraft] = useState(enrollment.integratedPlanMarkdown ?? '');
  const [historyOpen, setHistoryOpen] = useState(false);
  const update = useUpdateIntegratedPlan(enrollment.id);
  const revisions = useIntegratedPlanRevisions(historyOpen ? enrollment.id : null);

  // Hidrata draft quando o servidor traz update (ex: outro usuário salvou).
  useEffect(() => {
    setDraft(enrollment.integratedPlanMarkdown ?? '');
  }, [enrollment.integratedPlanMarkdown, enrollment.id]);

  const dirty = draft !== (enrollment.integratedPlanMarkdown ?? '');

  const handleSave = async () => {
    try {
      await update.mutateAsync(draft);
      toast.success('Plano atualizado. Revisão criada.');
    } catch (e: any) {
      toast.error(e?.response?.data?.error ?? 'Falha ao salvar');
    }
  };

  return (
    <>
      <Card>
        <CardHeader className="flex flex-row items-start justify-between gap-3 space-y-0">
          <div className="space-y-1">
            <CardTitle className="text-base">Plano integrado</CardTitle>
            <CardDescription>
              Documento único compartilhado por toda a equipe. Cada salvamento gera revisão.
            </CardDescription>
          </div>
          <div className="flex shrink-0 gap-2">
            <Button
              variant="outline"
              size="sm"
              onClick={() => setHistoryOpen(true)}
            >
              <History className="mr-1 h-3.5 w-3.5" />
              Histórico
            </Button>
            {canEdit && (
              <Button size="sm" onClick={handleSave} disabled={!dirty || update.isPending}>
                {update.isPending ? (
                  <Loader2 className="mr-1 h-3.5 w-3.5 animate-spin" />
                ) : (
                  <Save className="mr-1 h-3.5 w-3.5" />
                )}
                Salvar
              </Button>
            )}
          </div>
        </CardHeader>
        <CardContent className="space-y-2">
          {enrollment.integratedPlanUpdatedAt && (
            <p className="text-xs text-muted-foreground">
              Última edição{' '}
              {formatRelativeToNow(enrollment.integratedPlanUpdatedAt, {
                addSuffix: true,
              })}
              {dirty && (
                <span className="ml-2 text-amber-600">
                  · você tem alterações não salvas
                </span>
              )}
            </p>
          )}
          <Textarea
            rows={14}
            value={draft}
            onChange={(e) => setDraft(e.target.value)}
            disabled={!canEdit}
            placeholder="Leitura clínica integrada da equipe — direção, prioridades e aposta terapêutica do período. O paciente consulta um plano único, não silos por especialidade. Markdown é suportado."
            className="font-mono text-sm"
          />
        </CardContent>
      </Card>

      <Dialog open={historyOpen} onOpenChange={setHistoryOpen}>
        <DialogContent className="max-w-2xl">
          <DialogHeader>
            <DialogTitle>Histórico do plano integrado</DialogTitle>
            <DialogDescription>
              Todas as revisões em ordem cronológica (mais recente primeiro).
            </DialogDescription>
          </DialogHeader>
          {revisions.isLoading ? (
            <div className="flex h-32 items-center justify-center">
              <Loader2 className="h-5 w-5 animate-spin text-muted-foreground" />
            </div>
          ) : (revisions.data ?? []).length === 0 ? (
            <p className="py-8 text-center text-sm text-muted-foreground">
              Nenhuma revisão ainda.
            </p>
          ) : (
            <div className="max-h-[60vh] space-y-3 overflow-y-auto pr-2">
              {(revisions.data ?? []).map((rev) => (
                <Card key={rev.id} className="border bg-muted/40">
                  <CardHeader className="pb-2">
                    <div className="flex items-center justify-between text-xs text-muted-foreground">
                      <span className="font-medium text-foreground">
                        {rev.updatedBy?.name ?? '—'}
                      </span>
                      <span>
                        {formatDate(rev.createdAt, "dd 'de' MMM yyyy 'às' HH:mm")}
                      </span>
                    </div>
                  </CardHeader>
                  <CardContent>
                    <pre className="whitespace-pre-wrap font-sans text-xs text-muted-foreground">
                      {rev.content || <em>(vazio)</em>}
                    </pre>
                  </CardContent>
                </Card>
              ))}
            </div>
          )}
        </DialogContent>
      </Dialog>
    </>
  );
}
