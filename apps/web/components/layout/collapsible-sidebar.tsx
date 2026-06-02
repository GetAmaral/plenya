"use client";

import { useState, useEffect } from "react";
import Link from "next/link";
import Image from "next/image";
import { usePathname } from "next/navigation";
import { motion } from "framer-motion";
import { Search, DollarSign } from "lucide-react";

const DEV_BYPASS_AUTH = process.env.NEXT_PUBLIC_DEV_BYPASS_AUTH === 'true';
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
  LayoutList,
  ChevronLeft,
  ChevronRight,
  Menu,
  X,
  FileCheck,
  Shield,
  User as UserIcon,
  Pill,
  ShieldCheck,
  Activity,
  Target,
  CreditCard,
  Dumbbell,
  Library,
  Timer,
  Bot,
  UserPlus,
  BarChart3,
  MessageSquare,
  Plug,
  Settings,
  Workflow,
  Megaphone,
  Package,
} from "lucide-react";
import { cn } from "@/lib/utils";
import { useAuth } from "@/lib/use-auth";
import { isGranted, type UserRole } from "@/lib/auth-store";
import { Avatar, AvatarFallback } from "@/components/ui/avatar";
import { Badge } from "@/components/ui/badge";
import { ROLES, ROLE_BADGE_COLORS, getRoleColor, getRoleLabel } from "@/lib/roles";
import { useConversationsUnreadCount } from "@/lib/api/conversations-api";
import { useLeads } from "@/lib/api/leads-api";
import { useLabReviewCount } from "@/lib/api/lab-inbox";

type NavigationItem = {
  name: string;
  href: string;
  icon: any;
  adminOnly?: boolean;
  staffOnly?: boolean; // Restritas a profissionais (não patients)
  // requiredRoles: visível apenas se usuário tem PELO MENOS UM dos roles listados.
  // Útil quando staff genérico não basta — ex: apenas admin/secretary/manager.
  requiredRoles?: UserRole[];
  // badgeKey: chave que identifica qual contagem exibir no badge.
  // 'conversations' (mensagens não lidas), 'leads' (leads novos) ou 'lab-review' (exames a revisar).
  badgeKey?: 'conversations' | 'leads' | 'lab-review';
};

type NavGroup = { title: string; items: NavigationItem[] };

// Navegação agrupada por IA: as ferramentas clínicas diárias do médico no topo,
// depois operação (balcão/CRM), treino, biblioteca/config e admin.
const navGroups: NavGroup[] = [
  {
    title: "Clínico",
    items: [
      { name: "Início", href: "/dashboard", icon: Home },
      { name: "Pacientes", href: "/patients", icon: Users, staffOnly: true },
      { name: "Consultas", href: "/appointments", icon: Calendar },
      { name: "Anamneses", href: "/anamnesis", icon: Stethoscope },
      { name: "Escores de Saúde", href: "/health-scores", icon: Activity },
      { name: "Prescrições", href: "/prescriptions", icon: FileText },
      { name: "Exames", href: "/lab-results", icon: Microscope },
      { name: "Exames a revisar", href: "/lab-results/revisar", icon: Microscope, staffOnly: true, requiredRoles: ['admin', 'doctor', 'nurse', 'psychologist', 'nutritionist', 'physicalEducator'], badgeKey: 'lab-review' },
      { name: "Pedidos de Exames", href: "/lab-requests", icon: ClipboardList, staffOnly: true },
      { name: "Views de Resultados", href: "/lab-result-views", icon: LayoutList, staffOnly: true },
      { name: "Continuum", href: "/continuum", icon: Workflow, staffOnly: true },
      { name: "Boxes", href: "/continuum/boxes", icon: Package, staffOnly: true },
    ],
  },
  {
    title: "Operação",
    items: [
      { name: "Recepção", href: "/recepcao", icon: ClipboardList, staffOnly: true, requiredRoles: ['admin', 'secretary', 'manager'] },
      { name: "Calendário", href: "/calendario", icon: Calendar, staffOnly: true },
      { name: "Conversas", href: "/conversas", icon: MessageSquare, staffOnly: true, requiredRoles: ['admin', 'secretary', 'manager'], badgeKey: 'conversations' },
      { name: "Leads", href: "/leads", icon: UserPlus, staffOnly: true, requiredRoles: ['admin', 'secretary', 'manager'], badgeKey: 'leads' },
      { name: "Dashboard de Leads", href: "/leads/dashboard", icon: BarChart3, staffOnly: true, requiredRoles: ['admin', 'secretary', 'manager'] },
      { name: "Campanhas", href: "/campaigns", icon: Megaphone, staffOnly: true, requiredRoles: ['admin', 'secretary', 'manager'] },
    ],
  },
  {
    title: "Treino",
    items: [
      { name: "Avaliação Física", href: "/training/physical-assessments", icon: Activity, staffOnly: true },
      { name: "Testes de Fitness", href: "/training/fitness-tests", icon: Target, staffOnly: true },
      { name: "Avaliação Postural", href: "/training/posture", icon: UserIcon, staffOnly: true },
      { name: "Planos de Treino", href: "/training/workout-plans", icon: Dumbbell, staffOnly: true },
      { name: "Periodização", href: "/training/periodization", icon: Timer, staffOnly: true },
      { name: "Exercícios", href: "/training/exercises", icon: Library, staffOnly: true },
      { name: "Agente IA Treino", href: "/training/ai-agent", icon: Bot, staffOnly: true },
    ],
  },
  {
    title: "Biblioteca & Config",
    items: [
      { name: "Artigos", href: "/articles", icon: BookOpen },
      { name: "Escores", href: "/scores", icon: Network, staffOnly: true },
      { name: "Metodologias", href: "/methods", icon: Target, staffOnly: true },
      { name: "Templates de Anamnese", href: "/anamnesis-templates", icon: FileCheck, staffOnly: true },
      { name: "Template de Pedido de Exame", href: "/lab-request-templates", icon: LayoutTemplate, staffOnly: true },
      { name: "Configurar Agenda", href: "/configuracoes/agenda", icon: Settings, staffOnly: true },
      { name: "Preços de Consulta", href: "/configuracoes/precos", icon: DollarSign, staffOnly: true, requiredRoles: ['admin', 'manager'] },
      { name: "Templates Continuum", href: "/configuracoes/continuum-templates", icon: Workflow, staffOnly: true, requiredRoles: ['admin', 'manager'] },
      { name: "Templates de Box", href: "/configuracoes/box-templates", icon: Package, staffOnly: true, requiredRoles: ['admin', 'manager'] },
      { name: "Integrações", href: "/configuracoes/integracoes", icon: Plug, staffOnly: true },
    ],
  },
  {
    title: "Admin",
    items: [
      { name: "Planos de Assinatura", href: "/admin/subscription-plans", icon: LayoutList, adminOnly: true },
      { name: "Assinaturas", href: "/admin/subscriptions", icon: CreditCard, adminOnly: true },
      { name: "Usuários", href: "/admin/users", icon: Shield, adminOnly: true },
      { name: "Definições de Medicamentos", href: "/admin/medication-definitions", icon: Pill, adminOnly: true },
      { name: "Certificados Digitais", href: "/admin/certificates", icon: ShieldCheck, adminOnly: true },
    ],
  },
];

// Helper: check if user is patient-only (only has patient role, no other roles)
const isPatientOnly = (user: any) => {
  if (!user?.roles || !Array.isArray(user.roles)) return false;
  return user.roles.includes("patient") && user.roles.length === 1;
};

export function CollapsibleSidebar() {
  const pathname = usePathname();
  const { user, logout } = useAuth();
  const [isCollapsed, setIsCollapsed] = useState(false);
  const [isMobileOpen, setIsMobileOpen] = useState(false);
  const [isMobile, setIsMobile] = useState(false);
  const [isHydrated, setIsHydrated] = useState(false);
  const [searchQuery, setSearchQuery] = useState("");

  // Wait for Zustand hydration before showing admin-only items
  useEffect(() => {
    setIsHydrated(true);
  }, []);

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

  // Close mobile menu and clear search when route changes
  useEffect(() => {
    setIsMobileOpen(false);
    setSearchQuery("");
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
              <div className="relative flex h-10 w-10 shrink-0 items-center justify-center">
                <Image
                  src="/logo_infinity.svg"
                  alt="Plenya Logo"
                  fill
                  className="object-contain"
                />
              </div>
              <div className="flex-1 min-w-0">
                <h1 className="text-lg font-bold">Plenya EMR</h1>
                <p className="text-xs text-muted-foreground">Sistema Médico</p>
              </div>
              <button
                onClick={() => setIsMobileOpen(false)}
                className="flex h-8 w-8 shrink-0 items-center justify-center rounded-lg hover:bg-accent transition-colors"
                aria-label="Fechar menu"
              >
                <X className="h-5 w-5" />
              </button>
            </div>

            {/* Search */}
            <div className="px-4 pb-2 pt-3">
              <div className="relative">
                <Search className="absolute left-2.5 top-1/2 h-3.5 w-3.5 -translate-y-1/2 text-muted-foreground" />
                <input
                  type="text"
                  placeholder="Filtrar menu..."
                  value={searchQuery}
                  onChange={(e) => setSearchQuery(e.target.value)}
                  className="w-full rounded-md border border-input bg-background py-1.5 pl-8 pr-3 text-sm outline-none placeholder:text-muted-foreground focus:border-primary focus:ring-1 focus:ring-primary"
                />
              </div>
            </div>

            {/* Navigation */}
            <nav className="flex-1 space-y-1 overflow-y-auto p-4 pt-1">
              <NavGroupsView
                groups={navGroups}
                user={user}
                isHydrated={isHydrated}
                searchQuery={searchQuery}
                pathname={pathname}
                collapsed={false}
              />
            </nav>

            {/* User section */}
            <div className="border-t border-border p-4">
              <Link href="/profile">
                <div className="flex items-center gap-3 rounded-lg bg-accent/50 p-3 cursor-pointer hover:bg-accent transition-colors">
                  <Avatar className="h-10 w-10 shrink-0">
                    <AvatarFallback className="bg-linear-to-br from-blue-500 to-blue-700 text-sm font-semibold text-white">
                      {user?.name?.substring(0, 2).toUpperCase() || user?.email.substring(0, 2).toUpperCase() || "US"}
                    </AvatarFallback>
                  </Avatar>
                  <div className="flex-1 overflow-hidden">
                    <p className="truncate text-sm font-medium">
                      {user?.name || user?.email || "Usuário"}
                    </p>
                    {user?.roles && user.roles.length > 0 && (
                      <div className="mt-1 flex flex-wrap gap-1">
                        {(() => {
                          const primaryRole = user.roles[0];
                          const colorScheme = getRoleColor(primaryRole as UserRole);
                          const badgeClasses = ROLE_BADGE_COLORS[colorScheme as keyof typeof ROLE_BADGE_COLORS];
                          return (
                            <Badge
                              variant="outline"
                              className={cn("text-xs border", badgeClasses.active, badgeClasses.border)}
                            >
                              {getRoleLabel(primaryRole as UserRole)}
                            </Badge>
                          );
                        })()}
                        {user.roles.length > 1 && (
                          <Badge variant="outline" className="text-xs">
                            +{user.roles.length - 1}
                          </Badge>
                        )}
                      </div>
                    )}
                  </div>
                </div>
              </Link>

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
          <div className="relative flex h-10 w-10 shrink-0 items-center justify-center">
            <Image
              src="/logo_infinity.svg"
              alt="Plenya Logo"
              fill
              className="object-contain"
            />
          </div>
          {!isCollapsed && (
            <div className="flex-1 min-w-0">
              <h1 className="whitespace-nowrap text-lg font-bold">Plenya EMR</h1>
              <p className="whitespace-nowrap text-xs text-muted-foreground">
                Sistema Médico
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

        {/* Search (desktop, only when expanded) */}
        {!isCollapsed && (
          <div className="px-4 pb-2 pt-3">
            <div className="relative">
              <Search className="absolute left-2.5 top-1/2 h-3.5 w-3.5 -translate-y-1/2 text-muted-foreground" />
              <input
                type="text"
                placeholder="Filtrar menu..."
                value={searchQuery}
                onChange={(e) => setSearchQuery(e.target.value)}
                className="w-full rounded-md border border-input bg-background py-1.5 pl-8 pr-3 text-sm outline-none placeholder:text-muted-foreground focus:border-primary focus:ring-1 focus:ring-primary"
              />
            </div>
          </div>
        )}

        {/* Navigation */}
        <nav className="flex-1 space-y-1 overflow-y-auto p-4 pt-1">
          <NavGroupsView
            groups={navGroups}
            user={user}
            isHydrated={isHydrated}
            searchQuery={searchQuery}
            pathname={pathname}
            collapsed={isCollapsed}
          />
        </nav>

        {/* User Section */}
        <div className="border-t border-border p-4">
          {!isCollapsed ? (
            <>
              <Link href="/profile">
                <div className="flex items-center gap-3 rounded-lg bg-accent/50 p-3 cursor-pointer hover:bg-accent transition-colors">
                  <Avatar className="h-10 w-10 shrink-0">
                    <AvatarFallback className="bg-linear-to-br from-blue-500 to-blue-700 text-sm font-semibold text-white">
                      {user?.name?.substring(0, 2).toUpperCase() || user?.email.substring(0, 2).toUpperCase() || "US"}
                    </AvatarFallback>
                  </Avatar>
                  <div className="flex-1 overflow-hidden">
                    <p className="truncate text-sm font-medium">
                      {user?.name || user?.email || "Usuário"}
                    </p>
                    {user?.roles && user.roles.length > 0 && (
                      <div className="mt-1 flex flex-wrap gap-1">
                        {(() => {
                          const primaryRole = user.roles[0];
                          const colorScheme = getRoleColor(primaryRole as UserRole);
                          const badgeClasses = ROLE_BADGE_COLORS[colorScheme as keyof typeof ROLE_BADGE_COLORS];
                          return (
                            <Badge
                              variant="outline"
                              className={cn("text-xs border", badgeClasses.active, badgeClasses.border)}
                            >
                              {getRoleLabel(primaryRole as UserRole)}
                            </Badge>
                          );
                        })()}
                        {user.roles.length > 1 && (
                          <Badge variant="outline" className="text-xs">
                            +{user.roles.length - 1}
                          </Badge>
                        )}
                      </div>
                    )}
                  </div>
                </div>
              </Link>

              {DEV_BYPASS_AUTH && (
                <div className="mt-2 rounded-lg bg-amber-500/10 border border-amber-500/20 p-2">
                  <p className="text-xs font-semibold text-amber-600 dark:text-amber-400 text-center">
                    ⚠️ DEV BYPASS MODE
                  </p>
                </div>
              )}

              <button
                onClick={logout}
                className="mt-2 flex w-full items-center gap-2 rounded-lg px-3 py-2 text-sm font-medium text-muted-foreground transition-colors hover:bg-destructive/10 hover:text-destructive"
                title={DEV_BYPASS_AUTH ? "Recarregar página (DEV MODE)" : "Sair do sistema"}
              >
                <LogOut className="h-4 w-4" />
                {DEV_BYPASS_AUTH ? "Reload" : "Sair"}
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
              {DEV_BYPASS_AUTH && (
                <div className="w-10 h-10 rounded-lg bg-amber-500/10 border border-amber-500/20 flex items-center justify-center" title="DEV BYPASS MODE">
                  <span className="text-amber-600 dark:text-amber-400 text-xs font-bold">⚠️</span>
                </div>
              )}

              <Link href="/profile" title="Meu Perfil">
                <Avatar className="h-10 w-10 cursor-pointer hover:ring-2 hover:ring-primary transition-all">
                  <AvatarFallback className="bg-linear-to-br from-blue-500 to-blue-700 text-sm font-semibold text-white">
                    {user?.name?.substring(0, 2).toUpperCase() || user?.email.substring(0, 2).toUpperCase() || "US"}
                  </AvatarFallback>
                </Avatar>
              </Link>

              <button
                onClick={logout}
                className="flex h-10 w-10 items-center justify-center rounded-lg text-muted-foreground transition-colors hover:bg-destructive/10 hover:text-destructive"
                title={DEV_BYPASS_AUTH ? "Reload (DEV MODE)" : "Sair"}
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

// NavGroupsView: renderiza a navegação agrupada (Clínico/Operação/…), aplicando
// o gating por role + filtro de busca. Reusado no sidebar mobile e desktop.
// Grupos sem itens visíveis somem; o cabeçalho da seção só aparece expandido.
function NavGroupsView({
  groups,
  user,
  isHydrated,
  searchQuery,
  pathname,
  collapsed,
}: {
  groups: NavGroup[];
  user: any;
  isHydrated: boolean;
  searchQuery: string;
  pathname: string;
  collapsed: boolean;
}) {
  return (
    <>
      {groups.map((group) => {
        const items = group.items.filter((item) => {
          if (item.adminOnly && (!isHydrated || !isGranted(user, 'admin'))) return false;
          if (item.staffOnly && isHydrated && isPatientOnly(user)) return false;
          if (item.requiredRoles && isHydrated && !item.requiredRoles.some((r) => isGranted(user, r))) return false;
          if (searchQuery && !item.name.toLowerCase().includes(searchQuery.toLowerCase())) return false;
          return true;
        });
        if (items.length === 0) return null;
        return (
          <div key={group.title} className="space-y-1 pb-1">
            {!collapsed && (
              <p className="px-3 pb-1 pt-3 text-[10px] font-semibold uppercase tracking-wider text-muted-foreground/60">
                {group.title}
              </p>
            )}
            {items.map((item) => (
              <NavItemRow key={item.name} item={item} pathname={pathname} collapsed={collapsed} />
            ))}
          </div>
        );
      })}
    </>
  );
}

// NavItemRow: linha do menu que opcionalmente exibe badge de não-lidas.
// Renderizado tanto no sidebar desktop quanto na versão collapsed.
function NavItemRow({
  item,
  pathname,
  collapsed,
}: {
  item: NavigationItem;
  pathname: string;
  collapsed: boolean;
}) {
  const Icon = item.icon;
  const isActive = pathname === item.href || pathname.startsWith(item.href + '/');
  const conversationsUnread = useConversationsUnreadCount();
  // Contagem de leads novos: gateada por badgeKey pra rodar só na linha de Leads
  // (que só aparece pra admin/secretary/manager), evitando 403 nos demais perfis.
  const leadsNew = useLeads({ status: 'new' }, 0, 1, {
    enabled: item.badgeKey === 'leads',
  });
  // Exames a revisar: gateado por badgeKey (a rota é RequireClinician — evita 403 fora da linha).
  const labReview = useLabReviewCount({ enabled: item.badgeKey === 'lab-review' });
  const badgeValue =
    item.badgeKey === 'conversations'
      ? conversationsUnread.data ?? 0
      : item.badgeKey === 'leads'
        ? leadsNew.data?.total ?? 0
        : item.badgeKey === 'lab-review'
          ? labReview.data?.total ?? 0
          : 0;
  const badgeNoun =
    item.badgeKey === 'leads' ? 'novos' : item.badgeKey === 'lab-review' ? 'a revisar' : 'não lidas';
  const showBadge = badgeValue > 0;

  return (
    <Link href={item.href}>
      <div
        className={cn(
          'flex items-center gap-3 rounded-lg px-3 py-2.5 text-sm font-medium transition-colors',
          isActive
            ? 'bg-primary text-primary-foreground shadow-sm'
            : 'text-muted-foreground hover:bg-accent hover:text-accent-foreground',
          collapsed && 'justify-center'
        )}
        title={collapsed ? `${item.name}${showBadge ? ` (${badgeValue} ${badgeNoun})` : ''}` : undefined}
      >
        <div className="relative shrink-0">
          <Icon className="h-5 w-5" />
          {showBadge && collapsed && (
            <span className="absolute -right-1.5 -top-1.5 h-2 w-2 rounded-full bg-rose-500" />
          )}
        </div>
        {!collapsed && (
          <>
            <span className="flex-1 truncate">{item.name}</span>
            {showBadge && (
              <Badge
                variant="secondary"
                className="ml-auto bg-rose-100 text-rose-900 border-rose-200 text-[10px] px-1.5 py-0 h-5"
              >
                {badgeValue > 99 ? '99+' : badgeValue}
              </Badge>
            )}
          </>
        )}
      </div>
    </Link>
  );
}
