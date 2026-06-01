"use client";

/**
 * PatientSidebar — sidebar colapsável da área do paciente.
 * Mesmo idioma visual do CollapsibleSidebar do EMR profissional, items próprios.
 *
 * Em mobile, vira drawer overlay; bottom-nav cobre a navegação principal.
 */
import { useEffect, useState } from "react";
import Link from "next/link";
import Image from "next/image";
import { usePathname, useRouter } from "next/navigation";
import {
  Home,
  Workflow,
  Calendar,
  Microscope,
  Activity,
  MessageSquare,
  UserIcon,
  LogOut,
  Menu,
  X,
  ChevronLeft,
  ChevronRight,
  Package,
  FileText,
} from "lucide-react";

import { cn } from "@/lib/utils";
import { useAuthStore } from "@/lib/auth-store";
import { Avatar, AvatarFallback } from "@/components/ui/avatar";

type Item = { name: string; href: string; icon: any };

const items: Item[] = [
  { name: "Início", href: "/", icon: Home },
  { name: "Meu Continuum", href: "/continuum", icon: Workflow },
  { name: "Consultas", href: "/consultas", icon: Calendar },
  { name: "Exames", href: "/exames", icon: Microscope },
  { name: "Escores", href: "/escores", icon: Activity },
  { name: "Mensagens", href: "/mensagens", icon: MessageSquare },
  { name: "Boxes", href: "/boxes", icon: Package },
  { name: "Documentos", href: "/documentos", icon: FileText },
  { name: "Perfil", href: "/perfil", icon: UserIcon },
];

export function PatientSidebar() {
  const pathname = usePathname();
  const router = useRouter();
  const { user, clearAuth } = useAuthStore();
  const [collapsed, setCollapsed] = useState(false);
  const [mobileOpen, setMobileOpen] = useState(false);
  const [isMobile, setIsMobile] = useState(false);

  useEffect(() => {
    const check = () => setIsMobile(window.innerWidth < 1024);
    check();
    window.addEventListener("resize", check);
    return () => window.removeEventListener("resize", check);
  }, []);

  useEffect(() => {
    const stored = localStorage.getItem("patient-sidebar-collapsed");
    if (stored !== null) setCollapsed(stored === "true");
  }, []);

  useEffect(() => {
    setMobileOpen(false);
  }, [pathname]);

  const toggle = () => {
    const next = !collapsed;
    setCollapsed(next);
    localStorage.setItem("patient-sidebar-collapsed", String(next));
  };

  const handleLogout = () => {
    clearAuth();
    router.replace("/login");
  };

  const initials = (user?.name || user?.email || "P").substring(0, 2).toUpperCase();
  const width = collapsed ? 80 : 256;

  // Mobile: drawer overlay (mesma estética do EMR)
  if (isMobile) {
    return (
      <>
        {!mobileOpen && (
          <button
            onClick={() => setMobileOpen(true)}
            className="fixed top-4 left-4 z-50 flex h-12 w-12 items-center justify-center rounded-lg bg-primary text-primary-foreground shadow-lg lg:hidden"
            aria-label="Abrir menu"
          >
            <Menu className="h-6 w-6" />
          </button>
        )}
        {mobileOpen && (
          <div
            className="fixed inset-0 z-40 bg-black/50 lg:hidden"
            onClick={() => setMobileOpen(false)}
          />
        )}
        <aside
          className={cn(
            "fixed left-0 top-0 z-50 h-screen w-64 border-r border-border bg-card shadow-lg transition-transform duration-300 lg:hidden",
            mobileOpen ? "translate-x-0" : "-translate-x-full",
          )}
        >
          <SidebarBody
            items={items}
            pathname={pathname}
            collapsed={false}
            initials={initials}
            user={user}
            onClose={() => setMobileOpen(false)}
            onLogout={handleLogout}
          />
        </aside>
      </>
    );
  }

  // Desktop
  return (
    <aside
      style={{ width }}
      className="fixed left-0 top-0 z-30 hidden h-screen border-r border-border bg-card transition-[width] duration-300 lg:block"
    >
      <SidebarBody
        items={items}
        pathname={pathname}
        collapsed={collapsed}
        initials={initials}
        user={user}
        onLogout={handleLogout}
      />
      <button
        onClick={toggle}
        className="absolute -right-3 top-20 z-40 hidden h-6 w-6 items-center justify-center rounded-full border border-border bg-card shadow-sm lg:flex"
        aria-label={collapsed ? "Expandir" : "Recolher"}
      >
        {collapsed ? <ChevronRight className="h-3 w-3" /> : <ChevronLeft className="h-3 w-3" />}
      </button>
    </aside>
  );
}

function SidebarBody({
  items,
  pathname,
  collapsed,
  initials,
  user,
  onClose,
  onLogout,
}: {
  items: Item[];
  pathname: string | null;
  collapsed: boolean;
  initials: string;
  user: any;
  onClose?: () => void;
  onLogout: () => void;
}) {
  return (
    <div className="flex h-full flex-col">
      <div className="flex h-16 items-center gap-3 border-b border-border px-4">
        <div className="relative flex h-10 w-10 shrink-0 items-center justify-center">
          <Image src="/logo_infinity.svg" alt="Plenya" fill className="object-contain" />
        </div>
        {!collapsed && (
          <div className="flex-1 min-w-0">
            <h1 className="text-lg font-bold">Plenya</h1>
            <p className="text-xs text-muted-foreground">Área do paciente</p>
          </div>
        )}
        {onClose && (
          <button
            onClick={onClose}
            className="flex h-8 w-8 items-center justify-center rounded-lg hover:bg-accent"
            aria-label="Fechar menu"
          >
            <X className="h-5 w-5" />
          </button>
        )}
      </div>

      <nav className="flex-1 space-y-1 overflow-y-auto p-3">
        {items.map((item) => {
          const Icon = item.icon;
          const active =
            pathname === item.href ||
            (item.href !== "/" && pathname?.startsWith(item.href));
          return (
            <Link
              key={item.name}
              href={item.href}
              className={cn(
                "flex items-center gap-3 rounded-lg px-3 py-2 text-sm font-medium transition-colors",
                active
                  ? "bg-primary/10 text-primary"
                  : "text-muted-foreground hover:bg-accent hover:text-foreground",
                collapsed && "justify-center px-0",
              )}
              title={collapsed ? item.name : undefined}
            >
              <Icon className="h-4 w-4 shrink-0" />
              {!collapsed && <span className="truncate">{item.name}</span>}
            </Link>
          );
        })}
      </nav>

      <div className="border-t border-border p-3">
        <Link href="/perfil" className={cn("flex items-center gap-3 rounded-lg p-2 hover:bg-accent", collapsed && "justify-center")}>
          <Avatar className="h-9 w-9 shrink-0">
            <AvatarFallback className="bg-linear-to-br from-emerald-500 to-emerald-700 text-sm font-semibold text-white">
              {initials}
            </AvatarFallback>
          </Avatar>
          {!collapsed && (
            <div className="min-w-0 flex-1">
              <p className="truncate text-sm font-medium">{user?.name ?? "Paciente"}</p>
              <p className="truncate text-xs text-muted-foreground">{user?.email}</p>
            </div>
          )}
        </Link>

        <button
          onClick={onLogout}
          className={cn(
            "mt-2 flex w-full items-center gap-2 rounded-lg px-3 py-2 text-sm font-medium text-muted-foreground transition-colors hover:bg-destructive/10 hover:text-destructive",
            collapsed && "justify-center px-0",
          )}
          title={collapsed ? "Sair" : undefined}
        >
          <LogOut className="h-4 w-4" />
          {!collapsed && "Sair"}
        </button>
      </div>
    </div>
  );
}
