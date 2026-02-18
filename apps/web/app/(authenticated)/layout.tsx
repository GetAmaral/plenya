"use client";

import { useRequireAuth } from "@/lib/use-auth";
import { CollapsibleSidebar, useSidebarWidth } from "@/components/layout/collapsible-sidebar";
import { GlobalProcessingMonitor } from "@/components/processing/GlobalProcessingMonitor";
import dynamic from "next/dynamic";

const NotificationBell = dynamic(
  () => import("@/components/notifications/NotificationBell").then((m) => m.NotificationBell),
  { ssr: false }
);

export default function AuthenticatedLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  // Protect route - redirect to login if not authenticated
  useRequireAuth();

  const sidebarWidth = useSidebarWidth();

  return (
    <div className="min-h-screen bg-background">
      <CollapsibleSidebar />
      <GlobalProcessingMonitor />
      <main
        className="min-h-screen transition-all duration-300 ease-in-out"
        style={{
          marginLeft: `${sidebarWidth}px`,
        }}
      >
        {/* Header with notifications */}
        <div
          className="sticky top-0 z-30 flex h-16 items-center justify-end border-b bg-background/95 px-4 backdrop-blur supports-[backdrop-filter]:bg-background/60 sm:px-6 lg:px-8"
          style={{
            marginLeft: sidebarWidth === 0 ? '0' : `-${sidebarWidth}px`,
            paddingLeft: sidebarWidth === 0 ? '0' : `${sidebarWidth + 16}px`,
          }}
        >
          <NotificationBell />
        </div>

        {/* Extra padding-top on mobile to avoid menu button overlap */}
        <div className="p-4 pt-6 sm:p-6 lg:p-8">
          {children}
        </div>
      </main>
    </div>
  );
}
