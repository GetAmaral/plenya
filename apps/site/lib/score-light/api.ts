/**
 * Cliente do Escore Light — chama as rotas públicas do EMR.
 * Usado pelos client components do fluxo /escore-plenya/avaliar.
 */

import type {
  ConfirmClaimResult,
  CreateSessionRequest,
  PublicSession,
  RequestClaimPayload,
} from './types';

const apiBase =
  process.env.NEXT_PUBLIC_API_URL?.replace(/\/$/, '') ?? 'http://localhost:3001';

async function request<T>(path: string, init?: RequestInit): Promise<T> {
  const res = await fetch(`${apiBase}/api/v1${path}`, {
    ...init,
    headers: { 'Content-Type': 'application/json', ...(init?.headers ?? {}) },
  });
  if (!res.ok) {
    const text = await res.text().catch(() => '');
    throw new Error(`API ${res.status}: ${text || res.statusText}`);
  }
  return res.json() as Promise<T>;
}

export function createSession(payload: CreateSessionRequest) {
  return request<PublicSession>('/score-light/sessions', {
    method: 'POST',
    body: JSON.stringify(payload),
  });
}

export function getSession(code: string) {
  return request<PublicSession>(`/score-light/sessions/${encodeURIComponent(code)}`);
}

export function requestClaim(code: string, payload: RequestClaimPayload) {
  return request<{ message: string }>(
    `/score-light/sessions/${encodeURIComponent(code)}/claim`,
    {
      method: 'POST',
      body: JSON.stringify(payload),
    }
  );
}

export function confirmClaim(token: string) {
  return request<ConfirmClaimResult>('/score-light/claim/confirm', {
    method: 'POST',
    body: JSON.stringify({ token }),
  });
}

export type LightLabMatch = {
  scoreItemId: string;
  itemName: string;
  numericValue: number;
  originalRaw: string;
  unit?: string;
};

export type LightLabUnmatched = {
  nomeExame: string;
  resultado: string;
  unidade?: string;
};

export type LightLabExtractionResult = {
  matched: LightLabMatch[];
  unmatched: LightLabUnmatched[];
  warnings?: string[];
};

/** URL de download da sessão em JSON (LGPD portabilidade — art. 18 V). */
export function exportSessionURL(code: string): string {
  return `${apiBase}/api/v1/score-light/sessions/${encodeURIComponent(code)}/export`;
}

export async function deleteSession(code: string): Promise<void> {
  const res = await fetch(`${apiBase}/api/v1/score-light/sessions/${encodeURIComponent(code)}`, {
    method: 'DELETE',
  });
  if (!res.ok && res.status !== 204) {
    const text = await res.text().catch(() => '');
    throw new Error(`Falha ao excluir (${res.status}): ${text || res.statusText}`);
  }
}

export async function extractLabsFromPDF(pdf: File): Promise<LightLabExtractionResult> {
  const fd = new FormData();
  fd.append('pdf', pdf);
  const res = await fetch(`${apiBase}/api/v1/score-light/extract-labs`, {
    method: 'POST',
    body: fd,
  });
  if (!res.ok) {
    const text = await res.text().catch(() => '');
    throw new Error(`Extração falhou (${res.status}): ${text || res.statusText}`);
  }
  return res.json() as Promise<LightLabExtractionResult>;
}
