import 'server-only';

const BEEHIIV_API = 'https://api.beehiiv.com/v2';

type SubscribeInput = {
  email: string;
  source?: string;
  utmSource?: string;
  utmMedium?: string;
  utmCampaign?: string;
};

export async function subscribeToNewsletter(input: SubscribeInput) {
  const apiKey = process.env.BEEHIIV_API_KEY;
  const publicationId = process.env.BEEHIIV_PUBLICATION_ID;

  if (!apiKey || !publicationId) {
    console.log('[beehiiv] not configured. Would subscribe:', input.email);
    return { skipped: true };
  }

  try {
    const res = await fetch(`${BEEHIIV_API}/publications/${publicationId}/subscriptions`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${apiKey}`,
      },
      body: JSON.stringify({
        email: input.email,
        send_welcome_email: true,
        reactivate_existing: true,
        utm_source: input.utmSource ?? input.source ?? 'plenya-site',
        utm_medium: input.utmMedium ?? 'inline-form',
        utm_campaign: input.utmCampaign ?? 'boletim-plenya',
      }),
    });

    if (!res.ok) {
      const body = await res.text();
      console.error('[beehiiv] failed', res.status, body);
      return { sent: false, status: res.status };
    }
    return { sent: true };
  } catch (error) {
    console.error('[beehiiv] error', error);
    return { sent: false, error: 'network' };
  }
}
