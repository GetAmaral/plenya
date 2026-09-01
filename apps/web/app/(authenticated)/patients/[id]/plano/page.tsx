'use client';

/**
 * /patients/[id]/plano — a devolutiva de resultados do paciente (o "deck").
 *
 * Esta tela é de MONTAGEM E PUBLICAÇÃO, não de desenho slide a slide. O conteúdo é escrito a
 * partir do dossiê (que já deriva do prontuário as réguas e os achados ordenados por peso) e
 * chega aqui como a lista de slides; aqui se confere a prévia, mede-se o transbordo e publica-se.
 *
 * Antes de publicar, o servidor mede cada slide contra a moldura de 1920×1080 e RECUSA se algo não
 * couber: o slide tem altura fixa e overflow:hidden, então conteúdo demais não dá erro — ele
 * simplesmente some do PDF que o paciente recebe.
 */
import { useEffect, useMemo, useState } from 'react';
import { useParams, useRouter } from 'next/navigation';
import { formatDate } from '@/lib/format-date';
import { toast } from 'sonner';
import {
  ArrowLeft,
  Loader2,
  Plus,
  Eye,
  Ruler,
  Send,
  Trash2,
  FileText,
  AlertTriangle,
  CheckCircle2,
} from 'lucide-react';

import { useRequireAuth } from '@/lib/use-auth';
import { useRequireSelectedPatient } from '@/lib/use-require-selected-patient';
import { apiClient } from '@/lib/api-client';
import {
  usePatientPlans,
  useCreatePatientPlan,
  useUpdatePatientPlan,
  useDeletePatientPlan,
  useCheckPlanOverflow,
  usePublishPatientPlan,
  patientPlansApi,
  type PatientPlan,
  type DeckOverflow,
  type DeckSlide,
} from '@/lib/api/patient-plans';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Badge } from '@/components/ui/badge';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { Textarea } from '@/components/ui/textarea';
import { cn } from '@/lib/utils';

/** Um plano novo já nasce com a gramática dos decks que existem, para não começar do zero. */
const SLIDES_INICIAIS: DeckSlide[] = [
  { kind: 'cover', variant: 'deep', eyebrow: 'Seus exames', title: '', lede: '' },
  { kind: 'closing', variant: 'deep', eyebrow: 'Em uma página', title: '' },
];

export default function PatientPlanPage() {
  useRequireAuth();
  const router = useRouter();
  const params = useParams();
  const routeId = String(params?.id ?? '');

  // O prontuário é SEMPRE isolado pelo paciente selecionado: toda consulta abaixo é escopada por
  // ele, e a tela não abre sem ele.
  const { selectedPatient } = useRequireSelectedPatient();
  const patientId = selectedPatient?.id ?? '';

  const { data: plans = [], isLoading } = usePatientPlans(patientId || undefined);
  const createPlan = useCreatePatientPlan(patientId);
  const updatePlan = useUpdatePatientPlan(patientId);
  const deletePlan = useDeletePatientPlan(patientId);
  const checkOverflow = useCheckPlanOverflow(patientId);
  const publishPlan = usePublishPatientPlan(patientId);

  const [selectedId, setSelectedId] = useState<string | null>(null);
  const [title, setTitle] = useState('');
  const [contentText, setContentText] = useState('');
  const [contentError, setContentError] = useState<string | null>(null);
  const [overflow, setOverflow] = useState<DeckOverflow[] | null>(null);
  const [previewHTML, setPreviewHTML] = useState<string | null>(null);

  const selected = useMemo(
    () => plans.find((p) => p.id === selectedId) ?? null,
    [plans, selectedId],
  );

  useEffect(() => {
    if (!selected) return;
    setTitle(selected.title);
    setContentText(JSON.stringify(selected.content ?? [], null, 2));
    setContentError(null);
    setOverflow(null);
    setPreviewHTML(null);
  }, [selected?.id]);

  const parsedContent = (): DeckSlide[] | null => {
    try {
      const parsed = JSON.parse(contentText);
      if (!Array.isArray(parsed)) {
        setContentError('O conteúdo tem que ser uma lista de slides.');
        return null;
      }
      setContentError(null);
      return parsed as DeckSlide[];
    } catch (e) {
      setContentError(e instanceof Error ? e.message : 'JSON inválido');
      return null;
    }
  };

  const handleCreate = async () => {
    try {
      const plan = await createPlan.mutateAsync({
        title: 'Seus exames',
        content: SLIDES_INICIAIS,
      });
      setSelectedId(plan.id);
      toast.success('Rascunho criado');
    } catch (e) {
      toast.error(e instanceof Error ? e.message : 'Falha ao criar o rascunho');
    }
  };

  const handleSave = async () => {
    if (!selected) return;
    const content = parsedContent();
    if (!content) return;
    try {
      await updatePlan.mutateAsync({ id: selected.id, payload: { title, content } });
      setOverflow(null);
      toast.success('Plano salvo');
    } catch (e) {
      // Salvar que falha em silêncio é o pior caso: o médico segue achando que gravou.
      toast.error(e instanceof Error ? e.message : 'Falha ao salvar o plano');
    }
  };

  const handlePreview = async () => {
    if (!selected) return;
    try {
      // A prévia é HTML autenticado, então não dá para apontar o iframe direto para a URL: o token
      // vai no header. E não dá para usar apiClient.get, que sempre faz response.json(). O getBlob
      // é o caminho que já leva o token e devolve o corpo cru; daí entra por srcdoc, isolado.
      const blob = await apiClient.getBlob(patientPlansApi.previewURL(patientId, selected.id));
      setPreviewHTML(await blob.text());
    } catch (e) {
      toast.error(e instanceof Error ? e.message : 'Falha ao montar a prévia');
    }
  };

  const handleCheck = async () => {
    if (!selected) return;
    try {
      const res = await checkOverflow.mutateAsync(selected.id);
      setOverflow(res.slides);
      if (res.slides.length === 0) toast.success('Todos os slides cabem');
    } catch (e) {
      toast.error(e instanceof Error ? e.message : 'Falha ao medir os slides');
    }
  };

  const handlePublish = async () => {
    if (!selected) return;
    try {
      const plan = await publishPlan.mutateAsync(selected.id);
      setOverflow([]);
      toast.success(`Publicado no portal (versão ${plan.version})`);
    } catch (e) {
      // 422 = conteúdo que não cabe. Vale mostrar QUAIS slides, não só que falhou.
      const slides = (e as { data?: { slides?: DeckOverflow[] } })?.data?.slides;
      if (slides?.length) {
        setOverflow(slides);
        toast.error('Há slides que não cabem — corrija antes de publicar');
        return;
      }
      toast.error(e instanceof Error ? e.message : 'Falha ao publicar');
    }
  };

  const handleDelete = async (plan: PatientPlan) => {
    try {
      await deletePlan.mutateAsync(plan.id);
      if (selectedId === plan.id) setSelectedId(null);
      toast.success('Plano apagado');
    } catch (e) {
      toast.error(e instanceof Error ? e.message : 'Falha ao apagar o plano');
    }
  };

  if (!selectedPatient) {
    return (
      <div className="flex h-[60vh] items-center justify-center">
        <Loader2 className="h-6 w-6 animate-spin text-muted-foreground" />
      </div>
    );
  }

  return (
    <div className="space-y-6 p-6">
      <div className="flex items-center gap-3">
        <Button variant="ghost" size="sm" onClick={() => router.push(`/patients/${routeId}`)}>
          <ArrowLeft className="mr-1 h-4 w-4" />
          Voltar
        </Button>
        <div>
          <h1 className="text-xl font-semibold">Plano de devolutiva</h1>
          <p className="text-sm text-muted-foreground">{selectedPatient.name}</p>
        </div>
        <Button className="ml-auto" onClick={handleCreate} disabled={createPlan.isPending}>
          {createPlan.isPending ? (
            <Loader2 className="mr-2 h-4 w-4 animate-spin" />
          ) : (
            <Plus className="mr-2 h-4 w-4" />
          )}
          Novo plano
        </Button>
      </div>

      <div className="grid gap-6 lg:grid-cols-[320px_1fr]">
        <Card>
          <CardHeader>
            <CardTitle className="text-base">Planos</CardTitle>
          </CardHeader>
          <CardContent className="space-y-2">
            {isLoading && <Loader2 className="h-4 w-4 animate-spin text-muted-foreground" />}
            {!isLoading && plans.length === 0 && (
              <p className="text-sm text-muted-foreground">
                Nenhum plano ainda. Crie um rascunho para começar.
              </p>
            )}
            {plans.map((p) => (
              <button
                key={p.id}
                onClick={() => setSelectedId(p.id)}
                className={cn(
                  'w-full rounded-md border p-3 text-left transition-colors',
                  selectedId === p.id ? 'border-primary bg-accent' : 'hover:bg-accent/50',
                )}
              >
                <div className="flex items-center justify-between gap-2">
                  <span className="truncate text-sm font-medium">{p.title}</span>
                  <Badge variant={p.status === 'published' ? 'default' : 'secondary'}>
                    {p.status === 'published' ? `v${p.version}` : 'rascunho'}
                  </Badge>
                </div>
                <div className="mt-1 text-xs text-muted-foreground">
                  {p.content?.length ?? 0} slides ·{' '}
                  {p.status === 'published' && p.publishedAt
                    ? `publicado em ${formatDate(p.publishedAt)}`
                    : `editado em ${formatDate(p.updatedAt)}`}
                  {p.status === 'draft' && p.publishedAt
                    ? ` · v${p.version} no portal desde ${formatDate(p.publishedAt)}`
                    : ''}
                </div>
              </button>
            ))}
          </CardContent>
        </Card>

        {selected ? (
          <div className="space-y-4">
            <Card>
              <CardContent className="space-y-4 pt-6">
                <div className="space-y-2">
                  <Label htmlFor="plan-title">Título</Label>
                  <Input
                    id="plan-title"
                    value={title}
                    onChange={(e) => setTitle(e.target.value)}
                  />
                </div>
                <div className="space-y-2">
                  <Label htmlFor="plan-content">Slides</Label>
                  <Textarea
                    id="plan-content"
                    value={contentText}
                    onChange={(e) => setContentText(e.target.value)}
                    rows={16}
                    className="font-mono text-xs"
                  />
                  {contentError && (
                    <p className="text-sm text-destructive">{contentError}</p>
                  )}
                </div>
                <div className="flex flex-wrap gap-2">
                  <Button onClick={handleSave} disabled={updatePlan.isPending}>
                    {updatePlan.isPending && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
                    Salvar
                  </Button>
                  <Button variant="outline" onClick={handlePreview}>
                    <Eye className="mr-2 h-4 w-4" />
                    Prévia
                  </Button>
                  <Button variant="outline" onClick={handleCheck} disabled={checkOverflow.isPending}>
                    {checkOverflow.isPending ? (
                      <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                    ) : (
                      <Ruler className="mr-2 h-4 w-4" />
                    )}
                    Conferir se cabe
                  </Button>
                  <Button onClick={handlePublish} disabled={publishPlan.isPending}>
                    {publishPlan.isPending ? (
                      <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                    ) : (
                      <Send className="mr-2 h-4 w-4" />
                    )}
                    Publicar no portal
                  </Button>
                  <Button
                    variant="ghost"
                    className="text-destructive"
                    onClick={() => handleDelete(selected)}
                  >
                    <Trash2 className="h-4 w-4" />
                  </Button>
                </div>

                {overflow !== null && (
                  <div
                    className={cn(
                      'rounded-md border p-3 text-sm',
                      overflow.length === 0
                        ? 'border-emerald-200 bg-emerald-50 text-emerald-900'
                        : 'border-amber-300 bg-amber-50 text-amber-900',
                    )}
                  >
                    {overflow.length === 0 ? (
                      <span className="flex items-center gap-2">
                        <CheckCircle2 className="h-4 w-4" />
                        Todos os slides cabem na moldura.
                      </span>
                    ) : (
                      <>
                        <span className="flex items-center gap-2 font-medium">
                          <AlertTriangle className="h-4 w-4" />
                          Conteúdo que não cabe seria cortado do PDF sem aviso:
                        </span>
                        <ul className="mt-2 space-y-1">
                          {overflow.map((o) => (
                            <li key={o.slide}>
                              Slide {String(o.slide).padStart(2, '0')}
                              {o.title ? ` (${o.title})` : ''} passa{' '}
                              {o.bottom > 0 ? `${o.bottom}px embaixo` : ''}
                              {o.bottom > 0 && o.right > 0 ? ' e ' : ''}
                              {o.right > 0 ? `${o.right}px à direita` : ''}
                            </li>
                          ))}
                        </ul>
                      </>
                    )}
                  </div>
                )}

                {selected.status === 'published' && (
                  <div className="flex items-center gap-2 text-sm text-muted-foreground">
                    <FileText className="h-4 w-4" />
                    Versão {selected.version} publicada no portal: PDF 16:9 e A4 para impressão.
                  </div>
                )}
              </CardContent>
            </Card>

            {previewHTML && (
              <Card>
                <CardHeader>
                  <CardTitle className="text-base">Prévia</CardTitle>
                </CardHeader>
                <CardContent>
                  {/* sandbox sem allow-scripts: a prévia é só desenho, não roda nada. */}
                  <iframe
                    title="Prévia do plano"
                    srcDoc={previewHTML}
                    sandbox=""
                    className="h-[70vh] w-full rounded-md border bg-white"
                  />
                </CardContent>
              </Card>
            )}
          </div>
        ) : (
          <Card>
            <CardContent className="py-16 text-center text-sm text-muted-foreground">
              Escolha um plano à esquerda, ou crie um novo.
            </CardContent>
          </Card>
        )}
      </div>
    </div>
  );
}
