"use client";

import { useRequireAuth } from "@/lib/use-auth";
import { CollapsibleSidebar } from "@/components/layout/collapsible-sidebar";
import { useEffect, useState } from "react";

export default function AuthenticatedLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  // Protect route - redirect to login if not authenticated
  useRequireAuth();

  const [isCollapsed, setIsCollapsed] = useState(false);
  const [isMounted, setIsMounted] = useState(false);

  // Prevent hydration mismatch
  useEffect(() => {
    setIsMounted(true);
  }, []);

  // Sync with sidebar's collapsed state
  useEffect(() => {
    const updateCollapsedState = () => {
      const stored = localStorage.getItem("sidebar-collapsed");
      setIsCollapsed(stored === "true");
    };

    updateCollapsedState();

    // Listen for storage changes (e.g., from other tabs or sidebar toggle)
    window.addEventListener("storage", updateCollapsedState);

    // Custom event for same-tab updates
    const handleSidebarToggle = () => updateCollapsedState();
    window.addEventListener("sidebar-toggle", handleSidebarToggle);

    return () => {
      window.removeEventListener("storage", updateCollapsedState);
      window.removeEventListener("sidebar-toggle", handleSidebarToggle);
    };
  }, []);

  return (
    <div className="min-h-screen bg-background">
      <CollapsibleSidebar />
      <main
        className="min-h-screen transition-all duration-300 ease-in-out"
        style={{
          marginLeft: isMounted && window.innerWidth >= 1024
            ? (isCollapsed ? "80px" : "256px")
            : "0px",
        }}
      >
        <div className="p-4 pt-16 lg:pt-6 lg:p-8">
          {children}
        </div>
      </main>
    </div>
  );
}
