import { redirect } from 'next/navigation';

/**
 * `/conversas` foi separada em duas superfícies: WhatsApp (chat) e E-mails (webmail).
 * Mantemos a rota antiga como redirect pro WhatsApp (canal síncrono, padrão da recepção)
 * pra não quebrar deep links / atalhos existentes.
 */
export default function ConversasIndexPage() {
  redirect('/conversas/whatsapp');
}
