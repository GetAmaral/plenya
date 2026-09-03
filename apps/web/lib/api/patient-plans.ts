"use client";

/**
 * Plano de devolutiva do paciente (o "deck") — o mesmo conteúdo com três saídas: a tela do portal,
 * o PDF 16:9 (apresentar e mandar) e o PDF A4 paisagem (imprimir).
 *
 * O dossiê é o insumo DERIVADO do prontuário (réguas por exame, achados classificados e ordenados
 * por peso); o plano é o que alguém escreve em cima dele. Path-scoped por paciente, sempre.
 */

import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { apiClient } from "../api-client";
import type {
  DeckOverflow,
  DeckSlide,
  PlanAssistantTurn,
  PlanDossier,
  PlanDossierRefresh,
  PlanDossierStaleness,
  PlanMessage,
  PlanResolveResult,
  PlanSuggestion,
} from "@plenya/types";

export type {
  DeckOverflow,
  DeckSlide,
  PlanAssistantTurn,
  PlanDossier,
  PlanDossierStaleness,
  PlanMessage,
  PlanSuggestion,
};

export type PatientPlanStatus = "draft" | "published";

export interface PatientPlan {
  id: string;
  patientId: string;
  title: string;
  status: PatientPlanStatus;
  version: number;
  /** Conta EDIÇÕES do rascunho, e é o token de concorrência. `version` conta publicações. */
  revisionSeq: number;
  content: DeckSlide[];
  sourceSnapshotId?: string;
  authorUserId: string;
  publishedAt?: string;
  document16x9Id?: string;
  documentA4Id?: string;
  createdAt: string;
  updatedAt: string;
}

export interface SavePatientPlanPayload {
  title?: string;
  content?: DeckSlide[];
  sourceSnapshotId?: string;
  /**
   * A revisão que o cliente acha ser a corrente. Quando vem e não bate, o servidor responde 409 em
   * vez de sobrescrever escrita que este cliente não viu. É o que protege a edição do médico de um
   * salvamento atrasado, e vice-versa.
   */
  expectedRevision?: number;
}

export const patientPlanKeys = {
  list: (patientId: string) => ["patient-plans", patientId] as const,
  one: (patientId: string, planId: string) =>
    ["patient-plans", patientId, planId] as const,
  dossier: (patientId: string, planId: string) =>
    ["patient-plans", patientId, planId, "dossier"] as const,
  staleness: (patientId: string, planId: string) =>
    ["patient-plans", patientId, planId, "staleness"] as const,
  conversa: (patientId: string, planId: string) =>
    ["patient-plans", patientId, planId, "assistant"] as const,
  revisoes: (patientId: string, planId: string) =>
    ["patient-plans", patientId, planId, "revisoes"] as const,
  sugestoes: (patientId: string, planId: string) =>
    ["patient-plans", patientId, planId, "suggestions"] as const,
};

const base = (patientId: string) => `/api/v1/patients/${patientId}/plans`;

export function usePatientPlans(patientId: string | undefined) {
  return useQuery({
    queryKey: patientPlanKeys.list(patientId ?? ""),
    enabled: !!patientId,
    queryFn: () => apiClient.get<PatientPlan[]>(base(patientId!)),
  });
}

export function usePatientPlan(
  patientId: string | undefined,
  planId: string | undefined,
) {
  return useQuery({
    queryKey: patientPlanKeys.one(patientId ?? "", planId ?? ""),
    enabled: !!patientId && !!planId,
    queryFn: () => apiClient.get<PatientPlan>(`${base(patientId!)}/${planId}`),
  });
}

export function useCreatePatientPlan(patientId: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (payload: SavePatientPlanPayload) =>
      apiClient.post<PatientPlan>(base(patientId), payload),
    onSuccess: () =>
      qc.invalidateQueries({ queryKey: patientPlanKeys.list(patientId) }),
  });
}

export function useUpdatePatientPlan(patientId: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({
      id,
      payload,
    }: {
      id: string;
      payload: SavePatientPlanPayload;
    }) => apiClient.put<PatientPlan>(`${base(patientId)}/${id}`, payload),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: patientPlanKeys.list(patientId) });
    },
  });
}

export function useDeletePatientPlan(patientId: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: string) =>
      apiClient.delete<void>(`${base(patientId)}/${id}`),
    onSuccess: () =>
      qc.invalidateQueries({ queryKey: patientPlanKeys.list(patientId) }),
  });
}

/** Mede quais slides transbordam. Lista vazia = pode publicar. */
export function useCheckPlanOverflow(patientId: string) {
  return useMutation({
    mutationFn: (id: string) =>
      apiClient.get<{ slides: DeckOverflow[] }>(
        `${base(patientId)}/${id}/overflow`,
      ),
  });
}

/**
 * Publica no portal. Se algum slide transbordar o servidor responde 422 com a lista — conteúdo que
 * não cabe some do PDF em silêncio, então é bloqueio, não aviso.
 */
export function usePublishPatientPlan(patientId: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: string) =>
      apiClient.post<PatientPlan>(`${base(patientId)}/${id}/publish`, {}),
    onSuccess: () =>
      qc.invalidateQueries({ queryKey: patientPlanKeys.list(patientId) }),
  });
}

/**
 * O prontuário compilado CONGELADO deste plano.
 *
 * Não é o dossiê vivo do paciente: é o que estava valendo quando o plano nasceu. A distinção é o
 * ponto — números que mudam debaixo do texto que está sendo escrito foi o que motivou congelar.
 * `staleTime` alto porque, por construção, ele não muda sozinho.
 */
export function usePlanDossier(
  patientId: string | undefined,
  planId: string | undefined,
) {
  return useQuery({
    queryKey: patientPlanKeys.dossier(patientId ?? "", planId ?? ""),
    enabled: !!patientId && !!planId,
    staleTime: 30 * 60 * 1000,
    queryFn: () =>
      apiClient.get<{
        dossierId: string;
        seq: number;
        frozenAt: string;
        dossier: PlanDossier;
      }>(`${base(patientId!)}/${planId}/dossier`),
  });
}

/** Se o prontuário andou desde o congelamento, e em quê. Uma consulta de marcas d'água. */
export function usePlanDossierStaleness(
  patientId: string | undefined,
  planId: string | undefined,
) {
  return useQuery({
    queryKey: patientPlanKeys.staleness(patientId ?? "", planId ?? ""),
    enabled: !!patientId && !!planId,
    queryFn: () =>
      apiClient.get<PlanDossierStaleness>(
        `${base(patientId!)}/${planId}/dossier/staleness`,
      ),
  });
}

/**
 * Congela de novo e devolve o que mudou NO QUE O DECK CITA. Ato explícito: refrescar sozinho
 * trocaria número debaixo de quem está escrevendo.
 */
export function useRefreshPlanDossier(patientId: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (planId: string) =>
      apiClient.post<PlanDossierRefresh>(
        `${base(patientId)}/${planId}/dossier/refresh`,
        {},
      ),
    onSuccess: (_data, planId) => {
      qc.invalidateQueries({
        queryKey: patientPlanKeys.dossier(patientId, planId),
      });
      qc.invalidateQueries({
        queryKey: patientPlanKeys.staleness(patientId, planId),
      });
    },
  });
}

/** A conversa do plano, persistida: fechar o notebook não perde o turno. */
export function usePlanConversation(
  patientId: string | undefined,
  planId: string | undefined,
) {
  return useQuery({
    queryKey: patientPlanKeys.conversa(patientId ?? "", planId ?? ""),
    enabled: !!patientId && !!planId,
    queryFn: () =>
      apiClient.get<PlanMessage[]>(
        `${base(patientId!)}/${planId}/assistant/messages`,
      ),
  });
}

/** As sugestões esperando aceite. */
export function usePlanSuggestions(
  patientId: string | undefined,
  planId: string | undefined,
) {
  return useQuery({
    queryKey: patientPlanKeys.sugestoes(patientId ?? "", planId ?? ""),
    enabled: !!patientId && !!planId,
    queryFn: () =>
      apiClient.get<PlanSuggestion[]>(
        `${base(patientId!)}/${planId}/suggestions`,
      ),
  });
}

/**
 * Um turno da conversa.
 *
 * A chamada é síncrona e leva de dez a vinte segundos. `clientMessageId` é o que impede um turno
 * duplicado quando o médico reenvia depois de fechar a aba: o servidor devolve 409 em vez de rodar
 * de novo.
 */
export function useSendPlanMessage(patientId: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({
      planId,
      body,
      expectedRevision,
      clientMessageId,
    }: {
      planId: string;
      body: string;
      expectedRevision?: number;
      /**
       * Chave de idempotência. Tem que ser criada UMA vez por mensagem composta e repetida no
       * reenvio: gerada aqui dentro, cada tentativa levava um id novo e o `client_message_id`
       * único do servidor nunca casava, deixando toda a proteção contra turno duplicado (e o 409
       * que a acompanha) como código morto.
       */
      clientMessageId: string;
    }) =>
      apiClient.post<PlanAssistantTurn>(
        `${base(patientId)}/${planId}/assistant/messages`,
        {
          body,
          clientMessageId,
          expectedRevision,
        },
      ),
    onSuccess: (_d, { planId }) => {
      qc.invalidateQueries({
        queryKey: patientPlanKeys.conversa(patientId, planId),
      });
      qc.invalidateQueries({
        queryKey: patientPlanKeys.sugestoes(patientId, planId),
      });
      qc.invalidateQueries({ queryKey: patientPlanKeys.list(patientId) });
    },
  });
}

/** Aceita ou recusa sugestões. Resultado parcial é resposta legítima: ver `skipped`. */
export function useResolveSuggestions(patientId: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({
      planId,
      action,
      suggestionIds,
      slideId,
      expectedRevision,
    }: {
      planId: string;
      action: "accept" | "reject";
      suggestionIds?: string[];
      slideId?: string;
      expectedRevision?: number;
    }) =>
      apiClient.post<PlanResolveResult>(
        `${base(patientId)}/${planId}/suggestions/resolve`,
        {
          action,
          suggestionIds,
          slideId,
          expectedRevision,
        },
      ),
    onSuccess: (_d, { planId }) => {
      qc.invalidateQueries({
        queryKey: patientPlanKeys.sugestoes(patientId, planId),
      });
      qc.invalidateQueries({ queryKey: patientPlanKeys.list(patientId) });
    },
  });
}

/**
 * Uma linha do histórico do rascunho. Sem `content`: a lista carregaria dezenas de decks inteiros
 * para desenhar dezenas de linhas.
 */
export interface PlanRevision {
  id: string;
  seq: number;
  planVersion: number;
  title: string;
  authorKind: "human" | "assistant" | "system";
  authorName: string;
  reason: "edit" | "ai_apply" | "suggestion_accept" | "restore" | "publish";
  isPublication: boolean;
  slides: number;
  changedPaths?: string[];
  /** Só na publicação: o que a ferramenta escreveu e ninguém reescreveu depois. */
  aiTouchedPaths?: string[];
  aiModel?: string;
  createdAt: string;
}

export function usePlanRevisions(
  patientId: string | undefined,
  planId: string | undefined,
) {
  return useQuery({
    queryKey: patientPlanKeys.revisoes(patientId ?? "", planId ?? ""),
    enabled: !!patientId && !!planId,
    queryFn: () =>
      apiClient.get<PlanRevision[]>(`${base(patientId!)}/${planId}/revisions`),
  });
}

export function useRestorePlanRevision(patientId: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({
      planId,
      revisionId,
    }: {
      planId: string;
      revisionId: string;
    }) =>
      apiClient.post<PatientPlan>(
        `${base(patientId)}/${planId}/revisions/${revisionId}/restore`,
        {},
      ),
    onSuccess: (_d, { planId }) => {
      // Restaurar grava uma revisão nova, então a própria lista de histórico muda também.
      qc.invalidateQueries({
        queryKey: patientPlanKeys.one(patientId, planId),
      });
      qc.invalidateQueries({
        queryKey: patientPlanKeys.revisoes(patientId, planId),
      });
      qc.invalidateQueries({ queryKey: patientPlanKeys.list(patientId) });
    },
  });
}

/** Um número que o modelo escreveu e o servidor não achou no prontuário compilado. */
export interface PlanGenWarning {
  slideIndex: number;
  slideId?: string;
  title?: string;
  numeral: string;
  reason: string;
}

export interface GenerateDraftResult {
  plan: PatientPlan;
  /** O arco que o modelo escolheu e o que deixou de fora. Para o médico ler, nunca vai pro deck. */
  reply: string;
  warnings?: PlanGenWarning[];
  overflow?: DeckOverflow[];
  model?: string;
}

/**
 * Gera o rascunho INTEIRO a partir do prontuário compilado.
 *
 * É o passo que faltava: até aqui "novo plano" criava dois slides vazios e a conversa editava um
 * documento que ninguém tinha escrito. Leva de 40 a 90 segundos, porque o modelo redige o deck todo.
 */
export function useGeneratePlanDraft(patientId: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (body: { title?: string; instruction?: string } = {}) =>
      apiClient.post<GenerateDraftResult>(`${base(patientId)}/generate`, body),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: patientPlanKeys.list(patientId) });
    },
  });
}

export const patientPlansApi = {
  previewURL: (patientId: string, planId: string) =>
    `${base(patientId)}/${planId}/preview`,
};
