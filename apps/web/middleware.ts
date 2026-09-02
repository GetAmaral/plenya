import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";

// PORTAL_HOST detecta requisições vindo do subdomínio minha.plenyasaude.com.br
// (ou alias portal.). Em dev, a mesma codebase atende ambos via host header.
const PORTAL_HOST_PATTERNS = [/^minha\./i, /^portal\./i];

// Path "público" do portal — passa direto sem rewrite.
const PORTAL_PASSTHROUGH = ["/_next", "/api", "/favicon", "/static"];

// buildCSP — Content-Security-Policy estrita com nonce por requisição.
//
// script-src usa nonce + 'strict-dynamic': navegadores modernos IGNORAM
// 'unsafe-inline'/'self'/host-list e só executam scripts com o nonce do request (ou criados por
// um script já confiável, ex.: o loader do Google Identity Services). Isso bloqueia <script>
// inline injetado (XSS). 'unsafe-inline' e https: ficam só como fallback p/ navegadores antigos
// que não entendem strict-dynamic (eles são ignorados pelos modernos). 'unsafe-eval' foi
// REMOVIDO (nada no app usa eval). Ver docs/emr/estudo-xss-hardening.md.
function buildCSP(nonce: string): string {
  const api = process.env.NEXT_PUBLIC_API_URL || "http://localhost:3001";
  // 'unsafe-eval' SÓ em desenvolvimento: o HMR/react-refresh do Next usa eval. O build de
  // produção não precisa, então em prod o script-src fica sem eval (mais estrito).
  const evalSrc =
    process.env.NODE_ENV === "development" ? " 'unsafe-eval'" : "";
  return [
    "default-src 'self'",
    `script-src 'self' 'nonce-${nonce}' 'strict-dynamic' 'unsafe-inline' https:${evalSrc}`,
    "style-src 'self' 'unsafe-inline' https://fonts.googleapis.com",
    "img-src 'self' data: blob: https:",
    // media-src é OBRIGATÓRIO p/ <audio>/<video>: sem ele cai no default-src 'self' e
    // bloqueia o streaming de áudio das conversas (que vem da api, cross-origin) — o
    // <audio> nem chega a requisitar. blob: cobre mídia montada client-side.
    `media-src 'self' blob: data: ${api}`,
    // ${api} é OBRIGATÓRIO: a prévia do deck passou a linkar as sete fontes em vez de embutir
    // 1,9 MB de base64, e elas são servidas por /api/v1/deck-fonts, que é OUTRA origem (na dev
    // também: :3001 contra :3000). O iframe `srcdoc` herda a CSP deste documento, então sem isto
    // as fontes falham em silêncio, a métrica do texto muda e a prévia mente sobre o que cabe —
    // que é exatamente o modo de falha que o teste de fontes existe para impedir.
    `font-src 'self' data: https://fonts.gstatic.com ${api}`,
    `connect-src 'self' https://api.openai.com https://api.anthropic.com https://*.daily.co wss://*.daily.co ${api}`,
    "frame-src 'self' https://*.daily.co https://accounts.google.com",
    "frame-ancestors 'none'",
    "object-src 'none'",
    "base-uri 'self'",
    "form-action 'self'",
  ].join("; ");
}

// withSecurity injeta o nonce no request (pro Next aplicar nos próprios <script>) e a CSP
// na resposta. Recebe a resposta-base (next ou rewrite) e devolve uma equivalente com o nonce.
function withSecurity(request: NextRequest, rewriteURL?: URL): NextResponse {
  const nonce = btoa(crypto.randomUUID());
  const csp = buildCSP(nonce);

  const requestHeaders = new Headers(request.headers);
  requestHeaders.set("x-nonce", nonce);
  // Next lê a CSP do request pra propagar o nonce aos scripts do framework.
  requestHeaders.set("content-security-policy", csp);

  const response = rewriteURL
    ? NextResponse.rewrite(rewriteURL, { request: { headers: requestHeaders } })
    : NextResponse.next({ request: { headers: requestHeaders } });

  response.headers.set("content-security-policy", csp);
  return response;
}

/**
 * Reescreve o subdomínio do portal (minha.plenyasaude.com.br) pra /patient-portal e aplica a
 * CSP com nonce em todas as respostas (EMR + portal). URL fica limpa pro usuário.
 * Exceção: /sala/[token] (lobby telemedicina) não é reescrito.
 */
export function middleware(request: NextRequest) {
  const host = request.headers.get("host") ?? "";
  const isPortalHost = PORTAL_HOST_PATTERNS.some((p) => p.test(host));

  if (!isPortalHost) {
    return withSecurity(request);
  }

  const url = request.nextUrl;
  const path = url.pathname;

  if (PORTAL_PASSTHROUGH.some((prefix) => path.startsWith(prefix))) {
    return withSecurity(request);
  }

  if (path.startsWith("/sala")) {
    return withSecurity(request);
  }

  if (path.startsWith("/patient-portal")) {
    return withSecurity(request);
  }

  const newURL = url.clone();
  newURL.pathname = `/patient-portal${path === "/" ? "" : path}`;
  return withSecurity(request, newURL);
}

export const config = {
  matcher: ["/((?!_next/static|_next/image|favicon.ico|.*\\.png$).*)"],
};
