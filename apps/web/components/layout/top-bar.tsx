"use client";

import { usePathname, useRouter } from "next/navigation";
import { Loader2, Home, ChevronRight, User as UserIcon, LogOut } from "lucide-react";
import { useProcessingJobStore } from "@/lib/processing-job-store";
import { useAuth } from "@/lib/use-auth";
import { usePageHeader, type PageBreadcrumb } from "@/lib/page-context";
import { cn } from "@/lib/utils";
import { Badge } from "@/components/ui/badge";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { getRoleColor, getRoleLabel, ROLE_BADGE_COLORS } from "@/lib/roles";
import type { UserRole } from "@/lib/auth-store";
import dynamic from "next/dynamic";

const DEV_BYPASS_AUTH = process.env.NEXT_PUBLIC_DEV_BYPASS_AUTH === "true";

const NotificationBell = dynamic(
  () =>
    import("@/components/notifications/NotificationBell").then(
      (m) => m.NotificationBell
    ),
  { ssr: false }
);

const ROUTE_LABELS: Record<string, string> = {
  "/dashboard": "Dashboard",
  "/patients": "Pacientes",
  "/appointments": "Consultas",
  "/anamnesis": "Anamnese",
  "/prescriptions": "Prescrições",
  "/lab-results": "Resultados de Exames",
  "/lab-requests": "Solicitações de Exames",
  "/lab-result-views": "Visões de Exames",
  "/health-scores": "Escores de Saúde",
  "/scores": "Escore Plenya",
  "/articles": "Artigos",
  "/methods": "Métodos",
  "/admin": "Administração",
  "/profile": "Perfil",
};

function getModuleInfo(pathname: string): { label: string; route: string } | null {
  const match = Object.keys(ROUTE_LABELS)
    .filter((key) => pathname.startsWith(key))
    .sort((a, b) => b.length - a.length)[0];
  return match ? { label: ROUTE_LABELS[match], route: match } : null;
}

function truncate(text: string, max = 40): string {
  return text.length > max ? text.slice(0, max) + "…" : text;
}

/** Build the full crumb trail from context data + pathname fallback */
function buildCrumbs(
  breadcrumbs: PageBreadcrumb[],
  title: string,
  pathname: string
): PageBreadcrumb[] {
  const module = getModuleInfo(pathname);

  if (breadcrumbs.length > 0) {
    // If the last provided crumb already IS the title, use as-is
    const lastLabel = breadcrumbs[breadcrumbs.length - 1]?.label;
    if (!title || lastLabel === title) return breadcrumbs;
    // Otherwise append the title as the final crumb
    return [...breadcrumbs, { label: title }];
  }

  if (title) {
    // No breadcrumbs from page: synthesise parent from module label
    if (module && module.label !== title) {
      return [{ label: module.label, href: module.route }, { label: title }];
    }
    return [{ label: title }];
  }

  // Pure fallback: just the module label
  if (module) return [{ label: module.label }];
  return [];
}

function ProcessingIndicator() {
  const { activeJobs } = useProcessingJobStore();
  if (activeJobs.length === 0) return null;
  return (
    <div className="flex items-center gap-1.5 text-sm text-muted-foreground">
      <Loader2 className="h-4 w-4 animate-spin" />
      <span className="hidden sm:inline">{activeJobs.length}</span>
    </div>
  );
}

function UserMenu() {
  const { user, logout } = useAuth();
  const router = useRouter();
  const initials = user?.name
    ? user.name.split(" ").slice(0, 2).map((n) => n[0]).join("").toUpperCase()
    : user?.email?.[0]?.toUpperCase() ?? "U";

  const primaryRole = user?.roles?.[0] as UserRole | undefined;
  const extraRoles = (user?.roles?.length ?? 0) - 1;
  const badge = primaryRole
    ? ROLE_BADGE_COLORS[getRoleColor(primaryRole) as keyof typeof ROLE_BADGE_COLORS]
    : null;

  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <button
          className="flex h-8 w-8 items-center justify-center rounded-full bg-primary text-primary-foreground text-xs font-semibold transition-opacity hover:opacity-80 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring"
          title="Conta"
          aria-label="Menu da conta"
        >
          {initials}
        </button>
      </DropdownMenuTrigger>
      <DropdownMenuContent align="end" className="w-64">
        <DropdownMenuLabel className="font-normal">
          <div className="flex flex-col gap-1">
            <span className="truncate text-sm font-semibold leading-tight">
              {user?.name || user?.email || "Usuário"}
            </span>
            {user?.email && (
              <span className="truncate text-xs font-normal text-muted-foreground">
                {user.email}
              </span>
            )}
            {primaryRole && (
              <div className="mt-1 flex flex-wrap gap-1">
                <Badge
                  variant="outline"
                  className={cn("text-[10px] border", badge?.active, badge?.border)}
                >
                  {getRoleLabel(primaryRole)}
                </Badge>
                {extraRoles > 0 && (
                  <Badge variant="outline" className="text-[10px]">
                    +{extraRoles}
                  </Badge>
                )}
              </div>
            )}
            {DEV_BYPASS_AUTH && (
              <span className="mt-1 text-[10px] font-semibold text-amber-600 dark:text-amber-400">
                ⚠️ DEV BYPASS MODE
              </span>
            )}
          </div>
        </DropdownMenuLabel>
        <DropdownMenuSeparator />
        <DropdownMenuItem onSelect={() => router.push("/profile")}>
          <UserIcon className="mr-2 h-4 w-4" /> Meu perfil
        </DropdownMenuItem>
        <DropdownMenuItem
          onSelect={logout}
          className="text-destructive focus:text-destructive"
        >
          <LogOut className="mr-2 h-4 w-4" /> {DEV_BYPASS_AUTH ? "Reload" : "Sair"}
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  );
}

function BreadcrumbTrail({ crumbs }: { crumbs: PageBreadcrumb[] }) {
  const router = useRouter();

  return (
    <>
      {crumbs.map((crumb, i) => {
        const isLast = i === crumbs.length - 1;
        const label = isLast ? truncate(crumb.label) : crumb.label;
        const isClickable = !!crumb.href && !isLast;

        return (
          <span key={i} className="flex items-center gap-1.5 min-w-0 shrink-0">
            <ChevronRight className="h-3.5 w-3.5 text-muted-foreground/50 shrink-0" />
            {isClickable ? (
              <button
                onClick={() => router.push(crumb.href!)}
                className="text-muted-foreground hover:text-foreground transition-colors whitespace-nowrap"
              >
                {label}
              </button>
            ) : (
              <span
                className={
                  isLast
                    ? "heading-section text-foreground text-[1.5em] truncate"
                    : "text-muted-foreground whitespace-nowrap"
                }
                title={isLast && crumb.label.length > 40 ? crumb.label : undefined}
              >
                {label}
              </span>
            )}
          </span>
        );
      })}
    </>
  );
}

export function TopBar() {
  const pathname = usePathname();
  const router = useRouter();
  const { breadcrumbs, title } = usePageHeader();

  const crumbs = buildCrumbs(breadcrumbs, title, pathname);

  return (
    <div className="flex h-16 items-center justify-between pl-20 pr-4 lg:px-6">
      {/* Left: 🏠 > parent > Current Page */}
      <div className="flex items-center gap-1.5 text-sm min-w-0 overflow-hidden">
        <button
          onClick={() => router.push("/dashboard")}
          className="text-muted-foreground hover:text-foreground transition-colors shrink-0"
          title="Dashboard"
        >
          <Home className="h-4 w-4" />
        </button>

        {crumbs.length > 0 && <BreadcrumbTrail crumbs={crumbs} />}
      </div>

      {/* Right: processing + notifications + avatar */}
      <div className="flex items-center gap-2 ml-auto">
        <ProcessingIndicator />
        <NotificationBell />
        <UserMenu />
      </div>
    </div>
  );
}
