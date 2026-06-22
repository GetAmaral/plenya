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
  return differenceInYears(new Date(), d);
}

/** Substitui `new Date(x).toISOString()` (que lança em data inválida). */
export function toIsoOrNull(value: DateInput): string | null {
  const d = toDate(value);
  return d ? d.toISOString() : null;
}
