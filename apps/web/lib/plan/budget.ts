import type { DeckSlide } from '@plenya/types';

/**
 * Orçamento de altura do slide, estimado no cliente.
 *
 * Existe porque a medição de verdade é cara e disputada: ela roda no Chromium, mede o deck INTEIRO
 * e passa por um mutex global compartilhado com a impressão de receita e de pedido de exames. Medir
 * a cada tecla poria a impressão da clínica atrás do editor de deck.
 *
 * Então são duas camadas. Esta é a barata e imediata, feita das constantes reais de
 * `pdfdoc/ruler.go` e dos tetos observados; ela alimenta o selo do cartão enquanto se digita. A
 * outra é `CheckDeckOverflow`, que roda ao salvar e é a única verdade geométrica.
 *
 * É ESTIMATIVA, e a interface precisa dizer isso. Um selo que promete o que não pode cumprir é
 * pior do que selo nenhum.
 */

/** Altura útil: 1080 menos os 84px de padding de cima e de baixo (`--pad` em deck.go). */
const ALTURA_UTIL = 1080 - 84 * 2;

/** Constantes de `internal/pdfdoc/ruler.go`. */
const REGUA_LINHA = 132;
const REGUA_NOTA = 34;

/** Aproximações do cabeçalho, medidas nos decks reais. */
const ALTURA_EYEBROW = 46;
const ALTURA_TITULO_LINHA = 92;
const ALTURA_LEDE_LINHA = 44;
const ALTURA_PUNCH = 86;
const ALTURA_LINHA_TABELA = 64;
const ALTURA_CABECALHO_TABELA = 58;
const ALTURA_CARTAO = 210;
const ALTURA_GRUPO_ITEM = 46;

/** Caracteres que cabem numa linha, por tipografia. Grosso, e serve: o erro é de uma linha. */
const CHARS_TITULO = 46;
const CHARS_LEDE = 96;

export type NivelDeOcupacao = 'ok' | 'apertado' | 'provavel-estouro';

export interface Orcamento {
  usado: number;
  disponivel: number;
  nivel: NivelDeOcupacao;
  /** O que mais pesa neste slide, para o aviso apontar onde cortar. */
  maiorBloco?: string;
}

function linhas(texto: string | undefined, porLinha: number): number {
  if (!texto) return 0;
  return Math.max(1, Math.ceil(texto.length / porLinha));
}

/** Estima quanto do slide já está ocupado. */
export function orcamentoDoSlide(slide: DeckSlide): Orcamento {
  const partes: { nome: string; altura: number }[] = [];

  if (slide.eyebrow) partes.push({ nome: 'eyebrow', altura: ALTURA_EYEBROW });
  if (slide.title) partes.push({ nome: 'título', altura: linhas(slide.title, CHARS_TITULO) * ALTURA_TITULO_LINHA });
  if (slide.lede) partes.push({ nome: 'lede', altura: linhas(slide.lede, CHARS_LEDE) * ALTURA_LEDE_LINHA });
  if (slide.kicker) partes.push({ nome: 'kicker', altura: linhas(slide.kicker, CHARS_LEDE) * ALTURA_LEDE_LINHA });
  if (slide.punch) partes.push({ nome: 'punch', altura: ALTURA_PUNCH });
  if (slide.source) partes.push({ nome: 'fonte', altura: ALTURA_LEDE_LINHA });

  const rulers = slide.rulers ?? [];
  if (rulers.length > 0) {
    const alturaReguas = rulers.reduce(
      (soma, r) => soma + REGUA_LINHA + (r.note ? REGUA_NOTA : 0),
      0,
    );
    partes.push({ nome: `${rulers.length} régua${rulers.length > 1 ? 's' : ''}`, altura: alturaReguas });
  }

  const rows = slide.table?.rows?.length ?? 0;
  if (rows > 0) {
    partes.push({
      nome: `tabela com ${rows} linha${rows > 1 ? 's' : ''}`,
      altura: ALTURA_CABECALHO_TABELA + rows * ALTURA_LINHA_TABELA * (slide.table?.dense ? 0.8 : 1),
    });
  }

  const cards = slide.cards?.length ?? 0;
  if (cards > 0) partes.push({ nome: `${cards} cartão(ões)`, altura: ALTURA_CARTAO });

  const grupos = slide.takeaway?.groups ?? [];
  if (grupos.length > 0) {
    const maiorGrupo = Math.max(...grupos.map((g) => g.items?.length ?? 0), 0);
    partes.push({
      nome: 'para levar',
      altura: ALTURA_CARTAO + maiorGrupo * ALTURA_GRUPO_ITEM,
    });
  }

  const linhasResumo = (slide.summary?.cards ?? []).reduce(
    (m, c) => Math.max(m, c.lines?.length ?? 0),
    0,
  );
  if (linhasResumo > 0 || (slide.summary?.steps?.length ?? 0) > 0) {
    partes.push({
      nome: 'resumo',
      altura: ALTURA_CARTAO + linhasResumo * 78 + (slide.summary?.steps?.length ?? 0) * 40,
    });
  }

  const usado = Math.round(partes.reduce((s, p) => s + p.altura, 0));
  const maior = partes.sort((a, b) => b.altura - a.altura)[0];

  let nivel: NivelDeOcupacao = 'ok';
  if (usado > ALTURA_UTIL) nivel = 'provavel-estouro';
  else if (usado > ALTURA_UTIL * 0.85) nivel = 'apertado';

  return { usado, disponivel: ALTURA_UTIL, nivel, maiorBloco: maior?.nome };
}
