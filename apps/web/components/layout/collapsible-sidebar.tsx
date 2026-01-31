"use client";

import { useState, useEffect } from "react";
import Link from "next/link";
import Image from "next/image";
import { usePathname } from "next/navigation";
import { motion } from "framer-motion";
import {
  Calendar,
  FileText,
  Home,
  Users,
  Microscope,
  LogOut,
  Stethoscope,
  Network,
  BookOpen,
  ClipboardList,
  LayoutTemplate,
  ChevronLeft,
  ChevronRight,
  Menu,
  X,
} from "lucide-react";
import { cn } from "@/lib/utils";
import { useAuth } from "@/lib/use-auth";
import { Avatar, AvatarFallback } from "@/components/ui/avatar";
import { Badge } from "@/components/ui/badge";

const navigation = [
  { name: "Dashboard", href: "/dashboard", icon: Home },
  { name: "Pacientes", href: "/patients", icon: Users },
  { name: "Consultas", href: "/appointments", icon: Calendar },
  { name: "Anamnesis", href: "/anamnesis", icon: Stethoscope },
  { name: "Prescrições", href: "/prescriptions", icon: FileText },
  { name: "Exames", href: "/lab-results", icon: Microscope },
  { name: "Pedidos de Exames", href: "/lab-requests", icon: ClipboardList },
  { name: "Templates de Exames", href: "/lab-request-templates", icon: LayoutTemplate },
  { name: "Escores", href: "/scores", icon: Network },
  { name: "Artigos", href: "/articles", icon: BookOpen },
];

export function CollapsibleSidebar() {
  const pathname = usePathname();
  const { user, logout } = useAuth();
  const [isCollapsed, setIsCollapsed] = useState(false);
  const [isMobileOpen, setIsMobileOpen] = useState(false);
  const [isMobile, setIsMobile] = useState(false);

  // Detect mobile
  useEffect(() => {
    const checkMobile = () => {
      setIsMobile(window.innerWidth < 1024);
    };

    checkMobile();
    window.addEventListener("resize", checkMobile);
    return () => window.removeEventListener("resize", checkMobile);
  }, []);

  // Load collapsed state from localStorage (desktop only)
  useEffect(() => {
    if (!isMobile) {
      const stored = localStorage.getItem("sidebar-collapsed");
      if (stored !== null) {
        setIsCollapsed(stored === "true");
      }
    }
  }, [isMobile]);

  // Save collapsed state to localStorage
  const toggleCollapse = () => {
    const newState = !isCollapsed;
    setIsCollapsed(newState);
    localStorage.setItem("sidebar-collapsed", newState.toString());
  };

  // Close mobile menu when route changes
  useEffect(() => {
    setIsMobileOpen(false);
  }, [pathname]);

  // Keyboard shortcut: Cmd/Ctrl + B to toggle (desktop only)
  useEffect(() => {
    if (isMobile) return;

    const handleKeyDown = (e: KeyboardEvent) => {
      if ((e.metaKey || e.ctrlKey) && e.key === "b") {
        e.preventDefault();
        toggleCollapse();
      }
    };

    window.addEventListener("keydown", handleKeyDown);
    return () => window.removeEventListener("keydown", handleKeyDown);
  }, [isCollapsed, isMobile]);

  const getRoleColor = (role: string) => {
    switch (role) {
      case "admin":
        return "destructive";
      case "doctor":
        return "default";
      case "nurse":
        return "secondary";
      case "patient":
        return "outline";
      default:
        return "outline";
    }
  };

  const getRoleLabel = (role: string) => {
    switch (role) {
      case "admin":
        return "Administrador";
      case "doctor":
        return "Médico";
      case "nurse":
        return "Enfermeiro";
      case "patient":
        return "Paciente";
      default:
        return role;
    }
  };

  const sidebarWidth = isCollapsed ? 80 : 256;

  // Mobile: Render as overlay
  if (isMobile) {
    return (
      <>
        {/* Mobile Menu Button - Only show when sidebar is closed */}
        {!isMobileOpen && (
          <button
            onClick={() => setIsMobileOpen(true)}
            className="fixed top-4 left-4 z-50 flex items-center justify-center h-12 w-12 rounded-lg bg-primary text-primary-foreground shadow-lg lg:hidden hover:bg-primary/90 transition-colors"
            aria-label="Abrir menu"
          >
            <Menu className="h-6 w-6" />
          </button>
        )}

        {/* Overlay */}
        {isMobileOpen && (
          <div
            className="fixed inset-0 z-40 bg-black/50 lg:hidden"
            onClick={() => setIsMobileOpen(false)}
          />
        )}

        {/* Mobile Sidebar */}
        <aside
          className={cn(
            "fixed left-0 top-0 z-50 h-screen w-64 border-r border-border bg-card shadow-lg transition-transform duration-300 lg:hidden",
            isMobileOpen ? "translate-x-0" : "-translate-x-full"
          )}
        >
          <div className="flex h-full flex-col">
            {/* Logo and Close Button */}
            <div className="flex h-16 items-center border-b border-border px-4 gap-3">
              <div className="flex h-10 w-10 shrink-0 items-center justify-center">
                <Image
                  src="/logo_infinity.svg"
                  alt="Plenya Logo"
                  width={40}
                  height={40}
                  className="object-contain"
                  priority
                />
              </div>
              <div className="flex-1 min-w-0">
                <h1 className="text-lg font-bold">Plenya EMR</h1>
                <p className="text-xs text-muted-foreground">Medical System</p>
              </div>
              <button
                onClick={() => setIsMobileOpen(false)}
                className="flex h-8 w-8 shrink-0 items-center justify-center rounded-lg hover:bg-accent transition-colors"
                aria-label="Fechar menu"
              >
                <X className="h-5 w-5" />
              </button>
            </div>

            {/* Navigation */}
            <nav className="flex-1 space-y-1 overflow-y-auto p-4">
              {navigation.map((item) => {
                const Icon = item.icon;
                const isActive = pathname.startsWith(item.href);

                return (
                  <Link key={item.name} href={item.href}>
                    <div
                      className={cn(
                        "flex items-center gap-3 rounded-lg px-3 py-2.5 text-sm font-medium transition-colors",
                        isActive
                          ? "bg-primary text-primary-foreground shadow-sm"
                          : "text-muted-foreground hover:bg-accent hover:text-accent-foreground"
                      )}
                    >
                      <Icon className="h-5 w-5 shrink-0" />
                      <span>{item.name}</span>
                    </div>
                  </Link>
                );
              })}
            </nav>

            {/* User section */}
            <div className="border-t border-border p-4">
              <div className="flex items-center gap-3 rounded-lg bg-accent/50 p-3">
                <Avatar className="h-10 w-10 shrink-0">
                  <AvatarFallback className="bg-gradient-to-br from-blue-500 to-blue-700 text-sm font-semibold text-white">
                    {user?.email.substring(0, 2).toUpperCase() || "US"}
                  </AvatarFallback>
                </Avatar>
                <div className="flex-1 overflow-hidden">
                  <p className="truncate text-sm font-medium">
                    {user?.email || "Usuário"}
                  </p>
                  {user?.role && (
                    <Badge
                      variant={
                        getRoleColor(user.role) as
                          | "default"
                          | "secondary"
                          | "destructive"
                          | "outline"
                      }
                      className="mt-1 text-xs"
                    >
                      {getRoleLabel(user.role)}
                    </Badge>
                  )}
                </div>
              </div>

              <button
                onClick={logout}
                className="mt-2 flex w-full items-center gap-2 rounded-lg px-3 py-2 text-sm font-medium text-muted-foreground transition-colors hover:bg-destructive/10 hover:text-destructive"
              >
                <LogOut className="h-4 w-4" />
                Sair
              </button>
            </div>
          </div>
        </aside>
      </>
    );
  }

  // Desktop: Render fixed sidebar
  return (
    <motion.aside
      animate={{ width: sidebarWidth }}
      transition={{ duration: 0.3, ease: "easeInOut" }}
      className="fixed left-0 top-0 z-40 h-screen border-r border-border bg-card shadow-sm"
    >
      <div className="flex h-full flex-col">
        {/* Logo and Toggle */}
        <div className="flex h-16 items-center border-b border-border px-4 gap-3">
          <div className="flex h-10 w-10 shrink-0 items-center justify-center">
            <Image
              src="/logo_infinity.svg"
              alt="Plenya Logo"
              width={40}
              height={40}
              className="object-contain"
              priority
            />
          </div>
          {!isCollapsed && (
            <div className="flex-1 min-w-0">
              <h1 className="whitespace-nowrap text-lg font-bold">Plenya EMR</h1>
              <p className="whitespace-nowrap text-xs text-muted-foreground">
                Medical System
              </p>
            </div>
          )}
          <button
            onClick={toggleCollapse}
            className="flex h-8 w-8 shrink-0 items-center justify-center rounded-lg hover:bg-accent transition-colors ml-auto"
            title={isCollapsed ? "Expandir sidebar" : "Recolher sidebar"}
          >
            {isCollapsed ? (
              <ChevronRight className="h-4 w-4" />
            ) : (
              <ChevronLeft className="h-4 w-4" />
            )}
          </button>
        </div>

        {/* Navigation */}
        <nav className="flex-1 space-y-1 overflow-y-auto p-4">
          {navigation.map((item) => {
            const Icon = item.icon;
            const isActive = pathname.startsWith(item.href);

            return (
              <Link key={item.name} href={item.href}>
                <div
                  className={cn(
                    "flex items-center gap-3 rounded-lg px-3 py-2.5 text-sm font-medium transition-colors",
                    isActive
                      ? "bg-primary text-primary-foreground shadow-sm"
                      : "text-muted-foreground hover:bg-accent hover:text-accent-foreground",
                    isCollapsed && "justify-center"
                  )}
                  title={isCollapsed ? item.name : undefined}
                >
                  <Icon className="h-5 w-5 shrink-0" />
                  {!isCollapsed && <span>{item.name}</span>}
                </div>
              </Link>
            );
          })}
        </nav>

        {/* User Section */}
        <div className="border-t border-border p-4">
          {!isCollapsed ? (
            <>
              <div className="flex items-center gap-3 rounded-lg bg-accent/50 p-3">
                <Avatar className="h-10 w-10 shrink-0">
                  <AvatarFallback className="bg-gradient-to-br from-blue-500 to-blue-700 text-sm font-semibold text-white">
                    {user?.email.substring(0, 2).toUpperCase() || "US"}
                  </AvatarFallback>
                </Avatar>
                <div className="flex-1 overflow-hidden">
                  <p className="truncate text-sm font-medium">
                    {user?.email || "Usuário"}
                  </p>
                  {user?.role && (
                    <Badge
                      variant={
                        getRoleColor(user.role) as
                          | "default"
                          | "secondary"
                          | "destructive"
                          | "outline"
                      }
                      className="mt-1 text-xs"
                    >
                      {getRoleLabel(user.role)}
                    </Badge>
                  )}
                </div>
              </div>

              <button
                onClick={logout}
                className="mt-2 flex w-full items-center gap-2 rounded-lg px-3 py-2 text-sm font-medium text-muted-foreground transition-colors hover:bg-destructive/10 hover:text-destructive"
              >
                <LogOut className="h-4 w-4" />
                Sair
              </button>

              <div className="mt-3 text-center">
                <p className="text-xs text-muted-foreground">
                  <kbd className="rounded bg-muted px-1.5 py-0.5 font-mono text-xs">⌘</kbd>
                  {" + "}
                  <kbd className="rounded bg-muted px-1.5 py-0.5 font-mono text-xs">B</kbd>
                </p>
              </div>
            </>
          ) : (
            <div className="flex flex-col items-center gap-2">
              <Avatar className="h-10 w-10">
                <AvatarFallback className="bg-gradient-to-br from-blue-500 to-blue-700 text-sm font-semibold text-white">
                  {user?.email.substring(0, 2).toUpperCase() || "US"}
                </AvatarFallback>
              </Avatar>

              <button
                onClick={logout}
                className="flex h-10 w-10 items-center justify-center rounded-lg text-muted-foreground transition-colors hover:bg-destructive/10 hover:text-destructive"
                title="Sair"
              >
                <LogOut className="h-4 w-4" />
              </button>
            </div>
          )}
        </div>
      </div>
    </motion.aside>
  );
}

// Export the width for the layout to use
export function useSidebarWidth() {
  const [width, setWidth] = useState(256);
  const [isMobile, setIsMobile] = useState(false);

  useEffect(() => {
    const checkMobile = () => {
      const mobile = window.innerWidth < 1024;
      setIsMobile(mobile);
      if (mobile) {
        setWidth(0);
      }
    };

    const updateWidth = () => {
      if (isMobile) {
        setWidth(0);
        return;
      }

      const stored = localStorage.getItem("sidebar-collapsed");
      setWidth(stored === "true" ? 80 : 256);
    };

    checkMobile();
    updateWidth();

    window.addEventListener("resize", checkMobile);
    window.addEventListener("storage", updateWidth);

    // Poll for changes (for same-tab updates)
    const interval = setInterval(updateWidth, 100);

    return () => {
      window.removeEventListener("resize", checkMobile);
      window.removeEventListener("storage", updateWidth);
      clearInterval(interval);
    };
  }, [isMobile]);

  return width;
}
