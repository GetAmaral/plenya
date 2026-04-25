import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";

// PORTAL_HOST detecta requisições vindo do subdomínio meu.plenyasaude.com.br
// (ou alias portal.). Em dev, a mesma codebase atende ambos via host header.
const PORTAL_HOST_PATTERNS = [/^meu\./i, /^portal\./i];

// Path "público" do portal — passa direto sem rewrite.
const PORTAL_PASSTHROUGH = ["/_next", "/api", "/favicon", "/static"];

/**
 * Quando a request vem do subdomínio do portal (meu.plenyasaude.com.br),
 * reescreve internamente pra /patient-portal/<path>. O usuário continua
 * vendo a URL limpa "meu.plenyasaude.com.br/perfil" mas o Next serve
 * o conteúdo de app/patient-portal/perfil/page.tsx.
 *
 * No domínio principal (app.plenyasaude.com.br), seguimos com EMR profissional.
 *
 * Exceção: /sala/[token] tem layout próprio (lobby telemedicina sem login),
 * então passa direto sem entrar em /patient-portal.
 */
export function middleware(request: NextRequest) {
  const host = request.headers.get("host") ?? "";
  const isPortalHost = PORTAL_HOST_PATTERNS.some((p) => p.test(host));

  if (!isPortalHost) {
    return NextResponse.next();
  }

  const url = request.nextUrl;
  const path = url.pathname;

  if (PORTAL_PASSTHROUGH.some((prefix) => path.startsWith(prefix))) {
    return NextResponse.next();
  }

  if (path.startsWith("/sala")) {
    return NextResponse.next();
  }

  if (path.startsWith("/patient-portal")) {
    return NextResponse.next();
  }

  const newURL = url.clone();
  newURL.pathname = `/patient-portal${path === "/" ? "" : path}`;
  return NextResponse.rewrite(newURL);
}

export const config = {
  matcher: ["/((?!_next/static|_next/image|favicon.ico|.*\\.png$).*)"],
};
