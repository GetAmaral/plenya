"use client";

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
import { useEffect, useMemo, useState } from "react";
import { useParams, useRouter } from "next/navigation";
import { formatDate } from "@/lib/format-date";
import { toast } from "sonner";
import {
  AlertTriangle,
  ArrowLeft,
  CheckCircle2,
  Eye,
  FileText,
  Loader2,
  MessageSquare,
  Plus,
  Ruler,
  Send,
  Trash2,
} from "lucide-react";

import { useRequireAuth } from "@/lib/use-auth";
import { useRequireSelectedPatient } from "@/lib/use-require-selected-patient";
import { usePatient } from "@/lib/api/patient-api";
import { apiClient } from "@/lib/api-client";
import {
  usePatientPlans,
  useCreatePatientPlan,
  useUpdatePatientPlan,
  useDeletePatientPlan,
  useCheckPlanOverflow,
  type PlanAssistantTurn,
  usePlanConversation,
  usePlanDossier,
  usePlanDossierStaleness,
  usePlanRevisions,
  useRestorePlanRevision,
  usePlanSuggestions,
  useResolveSuggestions,
  useSendPlanMessage,
  usePublishPatientPlan,
  useRefreshPlanDossier,
  patientPlansApi,
  type PatientPlan,
  type DeckOverflow,
  type DeckSlide,
} from "@/lib/api/patient-plans";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Textarea } from "@/components/ui/textarea";
import { DossierColumn } from "@/components/plan/dossier-column";
import { SlideList } from "@/components/plan/deck/slide-list";
import { PlanHistoryPanel } from "@/components/plan/history-panel";
import { PlanChatPanel } from "@/components/plan/chat/plan-chat-panel";
import {
  Sheet,
  SheetContent,
  SheetTitle,
  SheetTrigger,
} from "@/components/ui/sheet";
import { cn } from "@/lib/utils";

/** Um plano novo já nasce com a gramática dos decks que existem, para não começar do zero. */
const SLIDES_INICIAIS: DeckSlide[] = [
  {
    kind: "cover",
    variant: "deep",
    eyebrow: "Seus exames",
    title: "",
    lede: "",
  },
  { kind: "closing", variant: "deep", eyebrow: "Em uma página", title: "" },
];

export default function PatientPlanPage() {
  useRequireAuth();
  const router = useRouter();
  const params = useParams();
  const routeId = String(params?.id ?? "");

  // O paciente desta tela é o da ROTA, não o do seletor global. A distinção importa aqui mais do
  // que em qualquer outra tela: abrir /patients/X/plano com o paciente Y selecionado no topo
  // mostraria os planos de Y sob a URL de X — e "Publicar no portal" entregaria a devolutiva de Y.
  // O seletor continua sendo exigido como guarda de sessão, igual às telas irmãs.
  useRequireSelectedPatient();
  const patientId = routeId;
  const { data: patient } = usePatient(patientId);

  const { data: plans = [], isLoading } = usePatientPlans(
    patientId || undefined,
  );
  const createPlan = useCreatePatientPlan(patientId);
  const updatePlan = useUpdatePatientPlan(patientId);
  const deletePlan = useDeletePatientPlan(patientId);
  const checkOverflow = useCheckPlanOverflow(patientId);
  const publishPlan = usePublishPatientPlan(patientId);

  const [selectedId, setSelectedId] = useState<string | null>(null);
  // O dossiê CONGELADO deste plano, não o vivo do paciente. Depende de selectedId, então vem
  // depois dele. Ver components/plan/dossier-column.
  const { data: dossie, isLoading: carregandoDossie } = usePlanDossier(
    patientId,
    selectedId ?? undefined,
  );
  const { data: envelhecimento } = usePlanDossierStaleness(
    patientId,
    selectedId ?? undefined,
  );
  const refrescarDossie = useRefreshPlanDossier(patientId);
  const { data: conversa = [] } = usePlanConversation(
    patientId,
    selectedId ?? undefined,
  );
  const { data: sugestoes = [] } = usePlanSuggestions(
    patientId,
    selectedId ?? undefined,
  );
  const { data: revisoes = [], isLoading: carregandoRevisoes } =
    usePlanRevisions(patientId, selectedId ?? undefined);
  const restaurar = useRestorePlanRevision(patientId ?? "");
  const enviarMensagem = useSendPlanMessage(patientId);
  const resolverSugestao = useResolveSuggestions(patientId);
  const [ultimoTurno, setUltimoTurno] = useState<PlanAssistantTurn | null>(
    null,
  );
  const [title, setTitle] = useState("");
  // Os slides viram estado estruturado. O texto do JSON some daqui: a escotilha por slide vive
  // dentro do próprio cartão, onde ela é útil sem ser a única forma de editar.
  const [slides, setSlides] = useState<DeckSlide[]>([]);
  const [overflow, setOverflow] = useState<DeckOverflow[] | null>(null);
  const [previewHTML, setPreviewHTML] = useState<string | null>(null);

  const selected = useMemo(
    () => plans.find((p) => p.id === selectedId) ?? null,
    [plans, selectedId],
  );

  useEffect(() => {
    if (!selected) return;
    setTitle(selected.title);
    setSlides(selected.content ?? []);
    setOverflow(null);
    setPreviewHTML(null);
  }, [selected?.id]);

  // "Conferir" e "Publicar" agem sobre o que está SALVO no servidor. Sem esta marca, editar e
  // clicar em conferir media o conteúdo antigo e dizia "todos os slides cabem"; publicar entregava
  // ao paciente o deck de antes da edição, com o toast dizendo que deu certo.
  //
  // A comparação é POR SLIDE e não pela string do JSON inteiro: o cartão precisa saber se ELE está
  // sujo para mostrar o selo, e comparar texto bruto era frágil de qualquer forma — o Go remove
  // campo vazio no `omitempty`, então digitar num campo e apagar deixava o plano sujo para sempre.
  const sujosPorSlide = useMemo(() => {
    const out = new Set<string>();
    if (!selected) return out;
    const base = selected.content ?? [];
    slides.forEach((s, i) => {
      const anterior = base.find((b) => b.id && b.id === s.id) ?? base[i];
      if (!anterior || JSON.stringify(anterior) !== JSON.stringify(s)) {
        out.add(s.id || `pos-${i}`);
      }
    });
    return out;
  }, [selected, slides]);

  const sujo =
    !!selected &&
    (title !== selected.title ||
      slides.length !== (selected.content ?? []).length ||
      sujosPorSlide.size > 0);

  const exigeSalvar = () => {
    if (sujo) {
      toast.error(
        "Salve as alterações antes: conferir e publicar usam o que está gravado.",
      );
      return true;
    }
    return false;
  };

  const parsedContent = (): DeckSlide[] | null => slides;

  // O turno é síncrono e leva de dez a vinte segundos. Bloqueia enquanto há edição não salva:
  // a IA escreve no SERVIDOR, e mandar um turno com o rascunho local divergindo daria 409 ou
  // apagaria o que está na tela.
  const handleEnviarMensagem = async (texto: string) => {
    if (!selectedId || !selected) return;
    try {
      const r = await enviarMensagem.mutateAsync({
        planId: selectedId,
        body: texto,
        expectedRevision: selected.revisionSeq,
      });
      setUltimoTurno(r);
      if (r.plan?.content) setSlides(r.plan.content as DeckSlide[]);
      if (r.failed)
        toast.error(r.error ?? "A rodada falhou e nada foi alterado.");
    } catch (e) {
      toast.error(
        e instanceof Error ? e.message : "Falha no turno do assistente",
      );
    }
  };

  const resolver = (action: "accept" | "reject") => async (id: string) => {
    if (!selectedId || !selected) return;
    try {
      const r = await resolverSugestao.mutateAsync({
        planId: selectedId,
        action,
        suggestionIds: [id],
        expectedRevision: selected.revisionSeq,
      });
      if (r.plan?.content) setSlides(r.plan.content as DeckSlide[]);
      const pulou = r.skipped ?? [];
      if (pulou.length > 0)
        toast.warning(pulou[0].reason ?? "Sugestão não aplicada");
      else if (action === "accept") toast.success("Sugestão aplicada");
    } catch (e) {
      toast.error(
        e instanceof Error ? e.message : "Falha ao resolver a sugestão",
      );
    }
  };

  // Refrescar é ato explícito: automático trocaria número debaixo de quem está escrevendo. O que
  // volta é restrito ao que o DECK cita, então cabe num toast em vez de virar tela.
  const handleRefreshDossie = async () => {
    if (!selectedId) return;
    try {
      const r = await refrescarDossie.mutateAsync(selectedId);
      const mudou = r.changed ?? [];
      if (mudou.length === 0) {
        toast.success(
          "Dossiê atualizado. Nenhum exame citado no deck mudou de valor.",
        );
        return;
      }
      toast.warning(
        `${mudou.length} exame${mudou.length > 1 ? "s" : ""} citado${mudou.length > 1 ? "s" : ""} mudou de valor: ` +
          mudou.map((c) => `${c.name} ${c.was} → ${c.now}`).join("; "),
        { duration: 12000 },
      );
    } catch (e) {
      toast.error(
        e instanceof Error ? e.message : "Falha ao atualizar o dossiê",
      );
    }
  };

  const handleCreate = async () => {
    try {
      const plan = await createPlan.mutateAsync({
        title: "Seus exames",
        content: SLIDES_INICIAIS,
      });
      setSelectedId(plan.id);
      toast.success("Rascunho criado");
    } catch (e) {
      toast.error(e instanceof Error ? e.message : "Falha ao criar o rascunho");
    }
  };

  const handleSave = async () => {
    if (!selected) return;
    const content = parsedContent();
    if (!content) return;
    try {
      const salvo = await updatePlan.mutateAsync({
        id: selected.id,
        payload: { title, content },
      });
      // Reescreve a caixa com o que o SERVIDOR gravou. Sem isto, "sujo" ficava true para sempre
      // depois de salvar — o Go normaliza a ordem das chaves e derruba campo vazio (`omitempty`),
      // então o texto colado nunca voltava a bater com o salvo, e Prévia/Conferir/Publicar ficavam
      // travados até o usuário trocar de plano e voltar.
      setTitle(salvo.title);
      setSlides(salvo.content ?? []);
      setOverflow(null);
      toast.success("Plano salvo");
    } catch (e) {
      // Salvar que falha em silêncio é o pior caso: o médico segue achando que gravou.
      toast.error(e instanceof Error ? e.message : "Falha ao salvar o plano");
    }
  };

  const handlePreview = async () => {
    if (!selected || exigeSalvar()) return;
    try {
      // A prévia é HTML autenticado, então não dá para apontar o iframe direto para a URL: o token
      // vai no header. E não dá para usar apiClient.get, que sempre faz response.json(). O getBlob
      // é o caminho que já leva o token e devolve o corpo cru; daí entra por srcdoc, isolado.
      const blob = await apiClient.getBlob(
        patientPlansApi.previewURL(patientId, selected.id),
      );
      setPreviewHTML(await blob.text());
    } catch (e) {
      toast.error(e instanceof Error ? e.message : "Falha ao montar a prévia");
    }
  };

  const handleCheck = async () => {
    if (!selected || exigeSalvar()) return;
    try {
      const res = await checkOverflow.mutateAsync(selected.id);
      setOverflow(res.slides);
      if (res.slides.length === 0) toast.success("Todos os slides cabem");
    } catch (e) {
      toast.error(e instanceof Error ? e.message : "Falha ao medir os slides");
    }
  };

  const handlePublish = async () => {
    if (!selected || exigeSalvar()) return;
    try {
      const plan = await publishPlan.mutateAsync(selected.id);
      setOverflow([]);
      toast.success(`Publicado no portal (versão ${plan.version})`);
    } catch (e) {
      // 422 = conteúdo que não cabe. Vale mostrar QUAIS slides, não só que falhou.
      const slides = (e as { data?: { slides?: DeckOverflow[] } })?.data
        ?.slides;
      if (slides?.length) {
        setOverflow(slides);
        toast.error("Há slides que não cabem — corrija antes de publicar");
        return;
      }
      toast.error(e instanceof Error ? e.message : "Falha ao publicar");
    }
  };

  const handleRestore = async (revisionId: string) => {
    if (!selected) return;
    try {
      const plan = await restaurar.mutateAsync({
        planId: selected.id,
        revisionId,
      });
      // Reescreve o rascunho local com o que o SERVIDOR gravou, pelo mesmo motivo do salvar: sem
      // isto o editor continuaria mostrando o conteúdo antigo e o selo de "sujo" nunca sairia.
      // O título entra junto: restaurar troca os dois, e reescrever só os slides deixaria o selo
      // preso pela diferença no título.
      setTitle(plan.title);
      setSlides(plan.content ?? []);
      setOverflow(null);
      toast.success("Rascunho restaurado");
    } catch (e) {
      toast.error(e instanceof Error ? e.message : "Falha ao restaurar");
    }
  };

  const handleDelete = async (plan: PatientPlan) => {
    try {
      await deletePlan.mutateAsync(plan.id);
      if (selectedId === plan.id) setSelectedId(null);
      toast.success("Plano apagado");
    } catch (e) {
      toast.error(e instanceof Error ? e.message : "Falha ao apagar o plano");
    }
  };

  if (!patientId) {
    return (
      <div className="flex h-[60vh] items-center justify-center">
        <Loader2 className="h-6 w-6 animate-spin text-muted-foreground" />
      </div>
    );
  }

  return (
    <div className="space-y-6 p-6">
      <div className="flex items-center gap-3">
        <Button
          variant="ghost"
          size="sm"
          onClick={() => router.push(`/patients/${routeId}`)}
        >
          <ArrowLeft className="mr-1 h-4 w-4" />
          Voltar
        </Button>
        <div>
          <h1 className="text-xl font-semibold">Plano de devolutiva</h1>
          <p className="text-sm text-muted-foreground">{patient?.name ?? ""}</p>
        </div>
        {selected && (
          <Sheet>
            <SheetTrigger asChild>
              <Button variant="outline" className="ml-auto 2xl:hidden">
                <MessageSquare className="mr-2 h-4 w-4" />
                Conversa
                {sugestoes.length > 0 && (
                  <span className="ml-2 rounded bg-amber-100 px-1.5 text-[10px] text-amber-900">
                    {sugestoes.length}
                  </span>
                )}
              </Button>
            </SheetTrigger>
            <SheetContent
              side="right"
              className="flex w-full flex-col sm:max-w-md"
            >
              <SheetTitle className="text-sm">Conversa</SheetTitle>
              <div className="min-h-0 flex-1">
                <PlanChatPanel
                  mensagens={conversa}
                  sugestoes={sugestoes}
                  onEnviar={handleEnviarMensagem}
                  enviando={enviarMensagem.isPending}
                  ultimoTurno={ultimoTurno}
                  desabilitado={sujo}
                  motivoDesabilitado="Salve as alterações antes: o assistente escreve no servidor."
                />
              </div>
            </SheetContent>
          </Sheet>
        )}
        <Button
          className={selected ? "" : "ml-auto"}
          onClick={handleCreate}
          disabled={createPlan.isPending}
        >
          {createPlan.isPending ? (
            <Loader2 className="mr-2 h-4 w-4 animate-spin" />
          ) : (
            <Plus className="mr-2 h-4 w-4" />
          )}
          Novo plano
        </Button>
      </div>

      {/* Duas colunas, não três. Medido: com a sidebar de 256px e o padding do layout, num laptop
          de 1366 um arranjo de três deixa ~310px para o centro — e o centro é o produto. O
          prontuário compilado divide a coluna da esquerda com a lista de planos; a conversa entra
          na fase 5, como terceira coluna só a partir de 2xl. */}
      <div className="grid gap-6 lg:grid-cols-[340px_1fr] 2xl:grid-cols-[320px_1fr_360px]">
        <div className="space-y-4">
          <Card>
            <CardHeader>
              <CardTitle className="text-base">Planos</CardTitle>
            </CardHeader>
            <CardContent className="space-y-2">
              {isLoading && (
                <Loader2 className="h-4 w-4 animate-spin text-muted-foreground" />
              )}
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
                    "w-full rounded-md border p-3 text-left transition-colors",
                    selectedId === p.id
                      ? "border-primary bg-accent"
                      : "hover:bg-accent/50",
                  )}
                >
                  <div className="flex items-center justify-between gap-2">
                    <span className="truncate text-sm font-medium">
                      {p.title}
                    </span>
                    <Badge
                      variant={
                        p.status === "published" ? "default" : "secondary"
                      }
                    >
                      {p.status === "published" ? `v${p.version}` : "rascunho"}
                    </Badge>
                  </div>
                  <div className="mt-1 text-xs text-muted-foreground">
                    {p.content?.length ?? 0} slides ·{" "}
                    {p.status === "published" && p.publishedAt
                      ? `publicado em ${formatDate(p.publishedAt)}`
                      : `editado em ${formatDate(p.updatedAt)}`}
                    {p.status === "draft" && p.publishedAt
                      ? ` · v${p.version} no portal desde ${formatDate(p.publishedAt)}`
                      : ""}
                  </div>
                </button>
              ))}
            </CardContent>
          </Card>

          {selected && (
            <Card>
              <CardContent className="h-[calc(100vh-22rem)] min-h-[20rem] pt-6">
                <DossierColumn
                  dossier={dossie?.dossier}
                  frozenAt={dossie?.frozenAt}
                  carregando={carregandoDossie}
                  motivosDeEnvelhecimento={envelhecimento?.reasons}
                  onRefresh={handleRefreshDossie}
                  refrescando={refrescarDossie.isPending}
                />
              </CardContent>
            </Card>
          )}
        </div>

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
                  <div className="flex items-baseline justify-between">
                    <Label>Slides</Label>
                    <span className="text-[11px] text-muted-foreground">
                      {slides.length} slide{slides.length === 1 ? "" : "s"} · a
                      miniatura mostra como o paciente vê na tela, não a moldura
                      impressa
                    </span>
                  </div>
                  <SlideList
                    slides={slides}
                    onChange={setSlides}
                    dossier={dossie?.dossier}
                    overflow={overflow ?? []}
                    sujos={sujosPorSlide}
                    sugestoes={sugestoes}
                    onAceitarSugestao={resolver("accept")}
                    onDescartarSugestao={resolver("reject")}
                    resolvendo={resolverSugestao.isPending}
                  />
                </div>
                {sujo && (
                  <p className="rounded-md border border-amber-300 bg-amber-50 px-3 py-2 text-sm text-amber-900">
                    Há alterações não salvas. Conferir e publicar usam o que
                    está gravado no servidor.
                  </p>
                )}
                <div className="flex flex-wrap gap-2">
                  <Button onClick={handleSave} disabled={updatePlan.isPending}>
                    {updatePlan.isPending && (
                      <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                    )}
                    Salvar
                  </Button>
                  <Button
                    variant="outline"
                    onClick={handlePreview}
                    disabled={sujo}
                  >
                    <Eye className="mr-2 h-4 w-4" />
                    Prévia
                  </Button>
                  <Button
                    variant="outline"
                    onClick={handleCheck}
                    disabled={sujo || checkOverflow.isPending}
                  >
                    {checkOverflow.isPending ? (
                      <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                    ) : (
                      <Ruler className="mr-2 h-4 w-4" />
                    )}
                    Conferir se cabe
                  </Button>
                  <Button
                    onClick={handlePublish}
                    disabled={sujo || publishPlan.isPending}
                  >
                    {publishPlan.isPending ? (
                      <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                    ) : (
                      <Send className="mr-2 h-4 w-4" />
                    )}
                    Publicar no portal
                  </Button>
                  <PlanHistoryPanel
                    revisoes={revisoes}
                    carregando={carregandoRevisoes}
                    restaurando={restaurar.isPending}
                    sujo={sujo}
                    onRestaurar={handleRestore}
                  />
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
                      "rounded-md border p-3 text-sm",
                      overflow.length === 0
                        ? "border-emerald-200 bg-emerald-50 text-emerald-900"
                        : "border-amber-300 bg-amber-50 text-amber-900",
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
                          {overflow.map((o) => {
                            const embaixo = o.bottom ?? 0;
                            const direita = o.right ?? 0;
                            return (
                              <li key={o.slide}>
                                Slide {String(o.slide ?? 0).padStart(2, "0")}
                                {o.title ? ` (${o.title})` : ""} passa{" "}
                                {embaixo > 0 ? `${embaixo}px embaixo` : ""}
                                {embaixo > 0 && direita > 0 ? " e " : ""}
                                {direita > 0 ? `${direita}px à direita` : ""}
                              </li>
                            );
                          })}
                        </ul>
                      </>
                    )}
                  </div>
                )}

                {selected.status === "published" && (
                  <div className="flex items-center gap-2 text-sm text-muted-foreground">
                    <FileText className="h-4 w-4" />
                    Versão {selected.version} publicada no portal: PDF 16:9 e A4
                    para impressão.
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

        {/* A conversa é terceira coluna só a partir de 2xl. Abaixo disso vira Sheet à direita, que
            é o padrão da casa para painel auxiliar (ver dossier-panel.tsx, aberto assim em tela
            cheia) — não é degradação de responsivo. */}
        {selected && (
          <Card className="hidden 2xl:block">
            <CardContent className="h-[calc(100vh-13rem)] pt-6">
              <PlanChatPanel
                mensagens={conversa}
                sugestoes={sugestoes}
                onEnviar={handleEnviarMensagem}
                enviando={enviarMensagem.isPending}
                ultimoTurno={ultimoTurno}
                desabilitado={sujo}
                motivoDesabilitado="Salve as alterações antes: o assistente escreve no servidor."
              />
            </CardContent>
          </Card>
        )}
      </div>
    </div>
  );
}
