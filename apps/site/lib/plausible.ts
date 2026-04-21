import Plausible from 'plausible-tracker';

const domain = process.env.NEXT_PUBLIC_PLAUSIBLE_DOMAIN || 'plenyasaude.com.br';

export const plausible =
  typeof window !== 'undefined'
    ? Plausible({ domain, trackLocalhost: false, apiHost: 'https://plausible.io' })
    : null;

export type PlenyaEvent =
  | 'cta_agendar_clicou'
  | 'cta_planos_clicou'
  | 'whatsapp_clicou'
  | 'form_contato_enviado'
  | 'form_agir_enviado'
  | 'newsletter_inscreveu'
  | 'blog_post_lido'
  | 'diagnostico_started'
  | 'diagnostico_completed';

export function track(event: PlenyaEvent, props?: Record<string, string | number | boolean>) {
  plausible?.trackEvent(event, { props });
}
