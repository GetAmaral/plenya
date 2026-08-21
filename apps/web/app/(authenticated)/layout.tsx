"use client";

import { useEffect } from "react";
import { useRequireAuth } from "@/lib/use-auth";
import { CollapsibleSidebar, useSidebarWidth } from "@/components/layout/collapsible-sidebar";
import { GlobalProcessingMonitor } from "@/components/processing/GlobalProcessingMonitor";
import { GlobalSearch } from "@/components/global-search";
import { TopBar } from "@/components/layout/top-bar";
import { PatientContextBar } from "@/components/layout/patient-context-bar";
import { WhatsAppDock } from "@/components/conversations/whatsapp-dock";
import { InactivityLock } from "@/components/auth/inactivity-lock";
import { SessionKeepAlive } from "@/components/auth/session-keep-alive";
import { WebPushSync } from "@/components/notifications/WebPushSync";
import { PageHeaderProvider } from "@/lib/page-context";

export default function AuthenticatedLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  useRequireAuth();
  const sidebarWidth = useSidebarWidth();

  // Pede storage persistente ao navegador (reduz evicção do token no PWA iOS, que é
  // script-writable e some após ~7 dias sem isso). Best-effort, uma vez.
  useEffect(() => {
    navigator.storage?.persist?.().catch(() => {});
  }, []);

  return (
    <PageHeaderProvider>
      <SessionKeepAlive />
      <WebPushSync />
      <InactivityLock>
      <div className="min-h-screen bg-background">
        <div className="print:hidden">
          <CollapsibleSidebar />
        </div>
        <div className="print:hidden">
          <GlobalProcessingMonitor />
        </div>
        {/* Busca global de pacientes (Cmd/Ctrl+K). Renderizada uma vez,
            disponível em todas as telas autenticadas do staff. */}
        <GlobalSearch />
        {/* Dock global de WhatsApp — persiste entre navegações (responder de qualquer tela). */}
        <WhatsAppDock />
        <main
          className="min-h-screen transition-all duration-300 ease-in-out print:ml-0"
          style={{ marginLeft: `${sidebarWidth}px` }}
        >
          {/* Sticky header: TopBar + PatientContextBar */}
          <div
            className="print:hidden sticky top-0 z-30 border-b bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60"
            style={{
              marginLeft: sidebarWidth === 0 ? "0" : `-${sidebarWidth}px`,
              paddingLeft: sidebarWidth === 0 ? "0" : `${sidebarWidth}px`,
            }}
          >
            <TopBar />
            <PatientContextBar />
          </div>

          {/* Page content */}
          <div className="p-4 pt-6 sm:p-6 lg:p-8 print:p-0">
            {children}
          </div>
        </main>
      </div>
      </InactivityLock>
    </PageHeaderProvider>
  );
}
