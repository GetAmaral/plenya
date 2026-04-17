import 'server-only';
import nodemailer from 'nodemailer';
import { brand } from '@plenya/brand';

let cached: nodemailer.Transporter | null = null;

export function getMailer(): nodemailer.Transporter | null {
  if (cached) return cached;
  const host = process.env.SMTP_HOST;
  const user = process.env.SMTP_USER;
  const pass = process.env.SMTP_PASSWORD;
  if (!host || !user || !pass) return null;
  cached = nodemailer.createTransport({
    host,
    port: Number(process.env.SMTP_PORT ?? 587),
    secure: process.env.SMTP_SECURE === 'true',
    auth: { user, pass },
  });
  return cached;
}

type LeadEmail = {
  name: string;
  email: string;
  phone: string;
  reason?: string;
  window?: string;
  source?: string;
  message?: string;
};

export async function sendLeadEmail(lead: LeadEmail) {
  const mailer = getMailer();
  if (!mailer) {
    console.log('[email] SMTP not configured. Lead would have been:', lead);
    return { skipped: true };
  }
  const to = process.env.LEADS_NOTIFY_EMAIL ?? brand.email;
  const subject = `Novo lead — ${lead.name} (${lead.source ?? 'site'})`;
  const text = [
    `Nome: ${lead.name}`,
    `Email: ${lead.email}`,
    `Telefone: ${lead.phone}`,
    lead.reason ? `Motivo: ${lead.reason}` : '',
    lead.window ? `Janela preferida: ${lead.window}` : '',
    lead.message ? `\nMensagem:\n${lead.message}` : '',
    '',
    `Origem: ${lead.source ?? 'site'}`,
    `Recebido em: ${new Date().toISOString()}`,
  ]
    .filter(Boolean)
    .join('\n');

  await mailer.sendMail({
    from: `"Plenya Site" <${process.env.SMTP_USER}>`,
    to,
    replyTo: lead.email,
    subject,
    text,
  });
  return { sent: true };
}
