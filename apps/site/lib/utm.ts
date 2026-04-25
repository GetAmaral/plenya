/**
 * Captura e persistência de parâmetros UTM.
 *
 * Quando o usuário aterra no site com `?utm_source=...&utm_medium=...&utm_campaign=...`
 * (vindo de QR code, link de Insta-bio, anúncio etc), persiste o trio em sessionStorage
 * para que o submit de qualquer formulário (Triagem, Painel, Contato) possa atribuir
 * corretamente — mesmo que o usuário tenha navegado por várias páginas antes.
 *
 * Sessão (não localStorage) é deliberado: a atribuição morre quando a aba fecha.
 * Se o usuário voltar amanhã sem o link UTM, é um lead orgânico.
 */

const STORAGE_KEY = 'plenya:utm';

export type UTMParams = {
  utm_source?: string;
  utm_medium?: string;
  utm_campaign?: string;
  utm_term?: string;
};

const KEYS: (keyof UTMParams)[] = [
  'utm_source',
  'utm_medium',
  'utm_campaign',
  'utm_term',
];

/**
 * Lê os UTMs da URL atual e, se houver pelo menos um, persiste em sessionStorage.
 * Idempotente: roda em todo mount do hook sem custo se nada mudou.
 *
 * Retorna o que ficou armazenado (URL > sessionStorage > nada).
 */
export function captureUTMFromURL(): UTMParams {
  if (typeof window === 'undefined') return {};
  const params = new URLSearchParams(window.location.search);
  const fromURL: UTMParams = {};
  for (const k of KEYS) {
    const v = params.get(k);
    if (v && v.trim().length > 0) fromURL[k] = v.trim().slice(0, 120);
  }
  if (Object.keys(fromURL).length > 0) {
    try {
      sessionStorage.setItem(STORAGE_KEY, JSON.stringify(fromURL));
    } catch {
      // sessionStorage indisponível (modo privado iOS antigo) — segue sem persistir
    }
    return fromURL;
  }
  return readStoredUTM();
}

/**
 * Lê UTMs persistidos em sessionStorage.
 */
export function readStoredUTM(): UTMParams {
  if (typeof window === 'undefined') return {};
  try {
    const raw = sessionStorage.getItem(STORAGE_KEY);
    if (!raw) return {};
    const parsed = JSON.parse(raw) as UTMParams;
    const out: UTMParams = {};
    for (const k of KEYS) {
      if (typeof parsed[k] === 'string' && (parsed[k] as string).length > 0) {
        out[k] = parsed[k];
      }
    }
    return out;
  } catch {
    return {};
  }
}

/**
 * Converte UTMs do formato URL (`utm_source`) pro formato camelCase do payload da API.
 */
export function utmToPayload(utm: UTMParams): {
  utmSource?: string;
  utmMedium?: string;
  utmCampaign?: string;
  utmTerm?: string;
} {
  return {
    utmSource: utm.utm_source,
    utmMedium: utm.utm_medium,
    utmCampaign: utm.utm_campaign,
    utmTerm: utm.utm_term,
  };
}
