"use client";

import { useRequireAuth } from "@/lib/use-auth";
import { Sidebar } from "@/components/dashboard/sidebar";

export default function AppointmentsLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  useRequireAuth();

  return (
    <div className="min-h-screen bg-background">
      <Sidebar />
      <main className="pl-64">{children}</main>
    </div>
  );
}
