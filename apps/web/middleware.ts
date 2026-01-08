import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";

// Middleware desabilitado - usando proteção client-side com useRequireAuth()
// O middleware não pode acessar localStorage, apenas cookies
// Como usamos localStorage para auth, a proteção é feita no client

export function middleware(request: NextRequest) {
  // Apenas permitir todas as requisições
  // A proteção de rotas é feita client-side com useRequireAuth() hook
  return NextResponse.next();
}

// Configurar matcher para aplicar middleware apenas em rotas específicas
export const config = {
  matcher: [
    /*
     * Match all request paths except:
     * - _next/static (static files)
     * - _next/image (image optimization files)
     * - favicon.ico (favicon file)
     * - public folder
     */
    "/((?!_next/static|_next/image|favicon.ico|.*\\.png$).*)",
  ],
};
