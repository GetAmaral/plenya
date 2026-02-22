"use client";

import { useRequireAuth } from "@/lib/use-auth";
import { CollapsibleSidebar, useSidebarWidth } from "@/components/layout/collapsible-sidebar";
import { GlobalProcessingMonitor } from "@/components/processing/GlobalProcessingMonitor";
import { TopBar } from "@/components/layout/top-bar";
import { PatientContextBar } from "@/components/layout/patient-context-bar";
import { PageHeaderProvider } from "@/lib/page-context";

export default function AuthenticatedLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  useRequireAuth();
  const sidebarWidth = useSidebarWidth();

  return (
    <PageHeaderProvider>
      <div className="min-h-screen bg-background">
        <div className="print:hidden">
          <CollapsibleSidebar />
        </div>
        <div className="print:hidden">
          <GlobalProcessingMonitor />
        </div>
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
    </PageHeaderProvider>
  );
}
