/**
 * Helpers da Recepção (cockpit do front-desk).
 *
 * Centraliza formatação de data/hora em America/Sao_Paulo e os mapas de cor
 * dos badges de status. Mantém a página fina e reaproveitável.
 *
 * Timezone: o backend trabalha em UTC; a UI sempre mostra em America/Sao_Paulo.
 */

import type {
  AppointmentStatus,
  AppointmentType,
} from "@/lib/api/calendar-api";

const SAO_PAULO_TZ = "America/Sao_Paulo";

/** "2026-05-30T13:00:00Z" -> "10:00" (horario de Brasilia). */
export function formatTimeBRT(iso: string): string {
  return new Date(iso).toLocaleTimeString("pt-BR", {
    timeZone: SAO_PAULO_TZ,
    hour: "2-digit",
    minute: "2-digit",
  });
}

/** Data por extenso em PT-BR, ex: "sexta-feira, 30 de maio de 2026". */
export function formatFullDatePtBR(date: Date): string {
  return date.toLocaleDateString("pt-BR", {
    timeZone: SAO_PAULO_TZ,
    weekday: "long",
    day: "numeric",
    month: "long",
    year: "numeric",
  });
}

/** Hora curta de um instante qualquer (mensagens, leads). Ex: "14:32". */
export function formatShortTime(iso: string): string {
  return new Date(iso).toLocaleTimeString("pt-BR", {
    timeZone: SAO_PAULO_TZ,
    hour: "2-digit",
    minute: "2-digit",
  });
}

/**
 * "Hoje as 14:32", "Ontem as 09:10" ou "28/05 as 16:00" dependendo de quao
 * recente foi. Usado nos cartoes de mensagens e leads novos.
 */
export function formatRelativeDayTime(iso: string): string {
  const d = new Date(iso);
  const now = new Date();
  const sameDay = (a: Date, b: Date) =>
    a.toLocaleDateString("pt-BR", { timeZone: SAO_PAULO_TZ }) ===
    b.toLocaleDateString("pt-BR", { timeZone: SAO_PAULO_TZ });

  const yesterday = new Date(now);
  yesterday.setDate(now.getDate() - 1);

  const time = formatShortTime(iso);
  if (sameDay(d, now)) return `Hoje as ${time}`;
  if (sameDay(d, yesterday)) return `Ontem as ${time}`;

  const date = d.toLocaleDateString("pt-BR", {
    timeZone: SAO_PAULO_TZ,
    day: "2-digit",
    month: "2-digit",
  });
  return `${date} as ${time}`;
}

/** Inicio do dia local (00:00 BRT) como instante. */
export function startOfTodaySaoPaulo(): Date {
  const now = new Date();
  // Formata o dia corrente em Sao Paulo, depois reconstroi o inicio do dia.
  const parts = new Intl.DateTimeFormat("en-CA", {
    timeZone: SAO_PAULO_TZ,
    year: "numeric",
    month: "2-digit",
    day: "2-digit",
  }).formatToParts(now);
  const get = (t: string) => parts.find((p) => p.type === t)?.value ?? "00";
  // -03:00 e o offset padrao do Brasil (sem horario de verao desde 2019).
  return new Date(`${get("year")}-${get("month")}-${get("day")}T00:00:00-03:00`);
}

export function addDays(d: Date, n: number): Date {
  const r = new Date(d);
  r.setDate(r.getDate() + n);
  return r;
}

/**
 * Variant do Badge (design system) por status de consulta.
 * Mapeia para as cores semanticas pedidas: agendada=neutra/ambar,
 * confirmada=azul/verde, concluida=verde, cancelada=vermelho/mudo,
 * faltou=vermelho.
 */
export const STATUS_BADGE_VARIANT: Record<
  AppointmentStatus,
  "secondary" | "stable" | "observation" | "critical" | "outline"
> = {
  scheduled: "observation", // ambar suave: aguardando confirmacao
  confirmed: "stable", // verde: confirmada
  checked_in: "observation", // ambar: paciente chegou, aguardando
  in_progress: "stable", // em atendimento
  completed: "secondary", // neutra: ja aconteceu
  cancelled: "outline", // mudo: cancelada
  no_show: "critical", // vermelho: nao compareceu
};

/** Classe de cor pro ponto/realce da hora por status. */
export const STATUS_DOT_CLASS: Record<AppointmentStatus, string> = {
  scheduled: "bg-amber-500",
  confirmed: "bg-emerald-500",
  checked_in: "bg-amber-500",
  in_progress: "bg-violet-500",
  completed: "bg-stone-400",
  cancelled: "bg-stone-300",
  no_show: "bg-rose-500",
};

export const CHANNEL_LABELS: Record<string, string> = {
  email: "E-mail",
  whatsapp: "WhatsApp",
  internal: "Interno",
};

export function appointmentTypeIsTelemed(type: AppointmentType): boolean {
  return type === "telemedicine";
}
