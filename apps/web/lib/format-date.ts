import {
  format as dfFormat,
  formatDistanceToNow as dfFormatDistanceToNow,
  differenceInYears,
  type FormatDistanceToNowOptions,
} from "date-fns";
import { ptBR } from "date-fns/locale";

/**
 * Formatação de datas resiliente a valores ausentes/ inválidos.
 *
 * Motivação: campos de data em branco são comuns no EMR (paciente cadastrado
 * sem data de nascimento, registros legados, etc.). O `format` do date-fns
 * lança `RangeError: Invalid time value` quando recebe uma data inválida, e
 * isso derruba a renderização da página inteira (não só do campo). Todo
 * componente que exibe data deve passar por aqui em vez de chamar
 * `format(new Date(x), …)` direto.
 *
 * Regra: entrada vazia/inválida → `fallback` (default "—"), nunca exceção.
 */

export type DateInput = string | number | Date | null | undefined;

/** Converte uma entrada qualquer em `Date` válido, ou `null`. */
export function toDate(value: DateInput): Date | null {
  if (value === null || value === undefined || value === "") return null;
  const d = value instanceof Date ? value : new Date(value);
  return Number.isNaN(d.getTime()) ? null : d;
}

/**
 * Substitui `format(new Date(x), pattern, { locale: ptBR })`.
 * Locale ptBR é o default; passe `options` para sobrescrever.
 */
export function formatDate(
  value: DateInput,
  pattern = "dd/MM/yyyy",
  options?: { fallback?: string } & Parameters<typeof dfFormat>[2],
): string {
  const { fallback = "—", ...fmtOptions } = options ?? {};
  const d = toDate(value);
  if (!d) return fallback;
  return dfFormat(d, pattern, { locale: ptBR, ...fmtOptions });
}

/**
 * Formata uma DATA PURA (sem hora) — data de nascimento, validade, etc.
 *
 * Datas puras chegam do backend como meia-noite UTC ("1990-01-01T00:00:00Z").
 * `formatDate`/`new Date(x)` formatam no fuso LOCAL: em BRT (UTC-3) isso vira
 * 31/12/1989 (o famoso "-1 dia"). Aqui lemos os componentes em UTC e os
 * reinterpretamos como data local, então o dia exibido é sempre o armazenado,
 * independente do fuso do navegador. Use SEMPRE isto para birthDate & cia.
 */
export function formatDateOnly(
  value: DateInput,
  pattern = "dd/MM/yyyy",
  options?: { fallback?: string } & Parameters<typeof dfFormat>[2],
): string {
  const { fallback = "—", ...fmtOptions } = options ?? {};
  const d = toDate(value);
  if (!d) return fallback;
  const cal = new Date(d.getUTCFullYear(), d.getUTCMonth(), d.getUTCDate());
  return dfFormat(cal, pattern, { locale: ptBR, ...fmtOptions });
}

/**
 * Diz se uma data PURA (sem hora) já passou, comparando pelo calendário.
 *
 * `new Date("2026-09-18") < new Date()` compara meia-noite UTC com o agora local: em BRT a data
 * "vence" às 21h do dia anterior. Era o que fazia a receita aparecer com selo "Expirado" ao lado
 * da coluna Validade mostrando o dia seguinte.
 */
export function isPastDateOnly(value: DateInput): boolean {
  const d = toDate(value);
  if (!d) return false;
  const alvo = new Date(d.getUTCFullYear(), d.getUTCMonth(), d.getUTCDate());
  const hoje = new Date();
  const hojeCal = new Date(hoje.getFullYear(), hoje.getMonth(), hoje.getDate());
  return alvo < hojeCal;
}

/** Atalho data + hora (dd/MM/yyyy HH:mm). */
export function formatDateTime(
  value: DateInput,
  options?: { fallback?: string } & Parameters<typeof dfFormat>[2],
): string {
  return formatDate(value, "dd/MM/yyyy HH:mm", options);
}

/** Substitui `formatDistanceToNow(new Date(x), …)`. */
export function formatRelativeToNow(
  value: DateInput,
  options?: { fallback?: string } & FormatDistanceToNowOptions,
): string {
  const { fallback = "—", ...distOptions } = options ?? {};
  const d = toDate(value);
  if (!d) return fallback;
  return dfFormatDistanceToNow(d, { locale: ptBR, ...distOptions });
}

/** Idade em anos a partir da data de nascimento, ou `null` se inválida. */
export function calcAge(value: DateInput): number | null {
  const d = toDate(value);
  if (!d) return null;
  // Mesma lógica de formatDateOnly: birthDate é data pura (meia-noite UTC).
  // Comparar pelo calendário em UTC evita idade -1 perto do aniversário em BRT.
  const cal = new Date(d.getUTCFullYear(), d.getUTCMonth(), d.getUTCDate());
  return differenceInYears(new Date(), cal);
}

/** Substitui `new Date(x).toISOString()` (que lança em data inválida). */
export function toIsoOrNull(value: DateInput): string | null {
  const d = toDate(value);
  return d ? d.toISOString() : null;
}
