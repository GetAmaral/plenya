import { ApiError, NetworkError } from './errors';

export interface FetcherContext {
  baseUrl: string;
  getAccessToken?: () => string | null | Promise<string | null>;
  onUnauthorized?: () => void | Promise<void>;
  extraHeaders?: () => Record<string, string> | Promise<Record<string, string>>;
  fetchImpl?: typeof fetch;
}

export interface FetcherConfig {
  signal?: AbortSignal;
  headers?: Record<string, string>;
  skipAuth?: boolean;
}

export type RequestBody = object | FormData | undefined;

let globalContext: FetcherContext | null = null;

export function configureFetcher(ctx: FetcherContext): void {
  globalContext = ctx;
}

export function getFetcherContext(): FetcherContext {
  if (!globalContext) {
    throw new Error(
      '@plenya/api-client: configureFetcher() must be called before using the fetcher.',
    );
  }
  return globalContext;
}

async function request<T>(
  method: string,
  path: string,
  body?: RequestBody,
  config: FetcherConfig = {},
): Promise<T> {
  const ctx = getFetcherContext();
  const fetchImpl = ctx.fetchImpl ?? fetch;

  const headers: Record<string, string> = {
    Accept: 'application/json',
    ...(config.headers ?? {}),
  };

  if (!(body instanceof FormData) && body !== undefined) {
    headers['Content-Type'] = 'application/json';
  }

  if (!config.skipAuth && ctx.getAccessToken) {
    const token = await ctx.getAccessToken();
    if (token) headers.Authorization = `Bearer ${token}`;
  }

  if (ctx.extraHeaders) {
    Object.assign(headers, await ctx.extraHeaders());
  }

  let response: Response;
  try {
    response = await fetchImpl(joinUrl(ctx.baseUrl, path), {
      method,
      headers,
      body: body instanceof FormData ? body : body !== undefined ? JSON.stringify(body) : undefined,
      signal: config.signal,
    });
  } catch (err) {
    throw new NetworkError('Network request failed', err);
  }

  if (response.status === 401 && ctx.onUnauthorized) {
    await ctx.onUnauthorized();
  }

  if (!response.ok) {
    let payload: { code?: string; message?: string; details?: unknown } = {};
    try {
      payload = (await response.json()) as typeof payload;
    } catch {
      /* body not JSON */
    }
    throw new ApiError(
      response.status,
      payload.code ?? 'http_error',
      payload.message ?? response.statusText,
      payload.details,
    );
  }

  if (response.status === 204) return undefined as T;

  const text = await response.text();
  if (!text) return undefined as T;

  try {
    return JSON.parse(text) as T;
  } catch {
    return text as unknown as T;
  }
}

function joinUrl(base: string, path: string): string {
  if (/^https?:\/\//i.test(path)) return path;
  const b = base.endsWith('/') ? base.slice(0, -1) : base;
  const p = path.startsWith('/') ? path : `/${path}`;
  return `${b}${p}`;
}

export const api = {
  get: <T>(path: string, config?: FetcherConfig): Promise<T> =>
    request<T>('GET', path, undefined, config),
  post: <T>(path: string, body?: RequestBody, config?: FetcherConfig): Promise<T> =>
    request<T>('POST', path, body, config),
  put: <T>(path: string, body?: RequestBody, config?: FetcherConfig): Promise<T> =>
    request<T>('PUT', path, body, config),
  patch: <T>(path: string, body?: RequestBody, config?: FetcherConfig): Promise<T> =>
    request<T>('PATCH', path, body, config),
  delete: <T>(path: string, config?: FetcherConfig): Promise<T> =>
    request<T>('DELETE', path, undefined, config),
};
