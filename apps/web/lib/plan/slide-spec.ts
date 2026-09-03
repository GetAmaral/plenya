import type { DeckSlideKind } from "@plenya/types";

/**
 * Que blocos cada tipo de slide aceita, e os tetos de cada um.
 *
 * A ideia que evita nove editores quase iguais: `kind` não tem editor próprio, ele apenas
 * SELECIONA quais blocos aparecem. Quem tem editor é o bloco. São seis editores para nove tipos,
 * e `plan-step` não é um décimo — é a combinação de três que já existem.
 */
export type BlockId =
  | "header"
  | "cards"
  | "table"
  | "rulers"
  | "summary"
  | "takeaway"
  | "steps";

export interface SlideSpec {
  label: string;
  blocks: BlockId[];
  /**
   * Tetos medidos, não estimados. Vêm da contagem nos decks reais e do teste de estouro do
   * backend: oito réguas num slide comprovadamente transbordam, quatro cabem. São teto DURO na
   * interface, não aviso — é dado empírico.
   */
  ceilings?: Partial<
    Record<"rulers" | "rows" | "cols" | "cards" | "lines" | "groups", number>
  >;
  /** Só leitura + escotilha de JSON. Ver `sequence` abaixo. */
  readOnly?: boolean;
  hint?: string;
}

export const SLIDE_SPEC: Record<DeckSlideKind, SlideSpec> = {
  cover: { label: "Capa", blocks: ["header"] },
  summary: {
    label: "Resumo",
    blocks: ["header", "summary"],
    ceilings: { cards: 2, lines: 4 },
    hint: "Dois cartões, até quatro linhas cada. É o slide que o paciente mais relê.",
  },
  rulers: {
    label: "Réguas",
    blocks: ["header", "rulers"],
    ceilings: { rulers: 4 },
    hint: "Até quatro réguas. Nenhuma régua entra sem um rótulo avaliativo visível no mesmo slide.",
  },
  "rulers-cards": {
    label: "Réguas e cartões",
    blocks: ["header", "rulers", "cards"],
    ceilings: { rulers: 2, cards: 2 },
    hint:
      "Duas réguas em cima, dois cartões embaixo. É o slide 08 dos dois decks aprovados, e nos " +
      "dois ele contrasta o exame que temos com o exame que falta.",
  },
  "two-cards": {
    label: "Dois caminhos",
    blocks: ["header", "cards", "table"],
    ceilings: { cards: 4, rows: 8, cols: 3 },
    hint:
      "Dois cartões, ou quatro na grade da decisão. É o contraste entre o caminho descartado e o " +
      "que vale que ensina; o veredicto do cartão é o que faz o slide decidir.",
  },
  "plan-step": {
    label: "Uma conduta",
    blocks: ["header", "cards", "table", "takeaway"],
    ceilings: { cards: 3, rows: 8, cols: 3, groups: 3 },
  },
  table: {
    label: "Tabela",
    blocks: ["header", "table"],
    ceilings: { rows: 8, cols: 3 },
    hint: "Coluna de dose não quebra linha: nunca use prosa nela.",
  },
  takeaway: {
    label: "Para levar",
    blocks: ["header", "takeaway"],
    ceilings: { groups: 3 },
    hint: "Três grupos lado a lado é o limite do layout, não um conselho.",
  },
  closing: { label: "Fecho", blocks: ["header"] },
  sequence: {
    label: "Sequência",
    blocks: [],
    readOnly: true,
    hint:
      "Os dois decks reais usaram tabela para a sequência, e este tipo nunca foi usado. Editável " +
      "pelo JSON até que alguém precise dele.",
  },
};

/** Os campos que aparecem em quase todo slide, e por isso ficam sempre visíveis no cartão. */
export const CAMPOS_UNIVERSAIS = ["eyebrow", "title", "punch"] as const;

/** Os que aparecem em menos de um terço dos slides e vivem atrás de "mais opções". */
export const CAMPOS_RAROS = ["lede", "kicker", "source", "legend"] as const;
