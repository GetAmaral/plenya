"use client";

/**
 * Layout da área do paciente (minha.plenyasaude.com.br).
 *
 * - Páginas de auth (/login, /auth/*, /esqueci-senha) são detectadas e renderizadas
 *   com layout limpo (só logo + main centralizado).
 * - Páginas autenticadas usam PatientSidebar + PatientBottomNav.
 *
 * Não há proteção de rota aqui — cada page chama useRequirePatientAuth quando precisa.
 */
import { usePathname } from "next/navigation";
import { PatientSidebar } from "@/components/patient-portal/sidebar";
import { PatientBottomNav } from "@/components/patient-portal/bottom-nav";

const PUBLIC_PATHS = ["/login", "/auth", "/esqueci-senha", "/escore-light"];

export default function PatientPortalLayout({ children }: { children: React.ReactNode }) {
  const pathname = usePathname();
  const isPublic = PUBLIC_PATHS.some((p) => pathname?.startsWith(p));

  if (isPublic) {
    return (
      <div className="min-h-screen bg-background">
        <header className="border-b">
          <div className="container mx-auto flex items-center justify-between px-4 py-5">
            <a href="/" className="text-lg font-medium tracking-wide">
              Plenya
            </a>
            <span className="text-sm text-muted-foreground">Área do paciente</span>
          </div>
        </header>
        <main className="container mx-auto px-4 py-10">{children}</main>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-background">
      <PatientSidebar />
      <div className="lg:pl-64">
        <main className="px-4 pb-20 pt-6 lg:px-8 lg:pb-8">{children}</main>
      </div>
      <PatientBottomNav />
    </div>
  );
}
