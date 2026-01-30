"use client";

import { useRequireAuth } from "@/lib/use-auth";
import { CollapsibleSidebar, useSidebarWidth } from "@/components/layout/collapsible-sidebar";

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
      <main
        className="min-h-screen transition-all duration-300 ease-in-out"
        style={{
          marginLeft: `${sidebarWidth}px`,
        }}
      >
        <div className="p-6 lg:p-8">
          {children}
        </div>
      </main>
    </div>
  );
}
