import 'server-only';

type RdLead = {
  name: string;
  email: string;
  phone: string;
  reason?: string;
  window?: string;
  source?: string;
};

const RD_CONVERSION_URL = 'https://api.rd.services/platform/conversions';

export async function sendToRdStation(lead: RdLead) {
  const token = process.env.RD_STATION_TOKEN;
  if (!token) {
    return { skipped: true };
  }
  try {
    const res = await fetch(RD_CONVERSION_URL, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify({
        event_type: 'CONVERSION',
        event_family: 'CDP',
        payload: {
          conversion_identifier: lead.source ?? 'site-contato',
          name: lead.name,
          email: lead.email,
          mobile_phone: lead.phone,
          cf_motivo: lead.reason,
          cf_janela_preferida: lead.window,
          tags: ['plenya-site', lead.source ?? 'contato'],
        },
      }),
    });
    if (!res.ok) {
      const body = await res.text();
      console.error('[rdstation] failed', res.status, body);
      return { sent: false, status: res.status };
    }
    return { sent: true };
  } catch (error) {
    console.error('[rdstation] error', error);
    return { sent: false, error: 'network' };
  }
}
