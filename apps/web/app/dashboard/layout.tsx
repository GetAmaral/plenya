"use client";

import { useRequireAuth } from "@/lib/use-auth";
import { Sidebar } from "@/components/dashboard/sidebar";

export default function DashboardLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  // Proteger rota - redireciona para login se não autenticado
  useRequireAuth();

  return (
    <div className="min-h-screen bg-background">
      <Sidebar />
      <main className="pl-64">
        {children}
      </main>
    </div>
  );
}
