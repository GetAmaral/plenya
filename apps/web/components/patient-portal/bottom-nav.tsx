"use client";

/**
 * BottomNav — barra inferior fixa em mobile (lg:hidden).
 * 5 ícones principais. Mais módulos seguem acessíveis via sidebar drawer.
 */
import Link from "next/link";
import { usePathname } from "next/navigation";
import { Home, Workflow, Calendar, MessageSquare, UserIcon } from "lucide-react";
import { cn } from "@/lib/utils";

const items = [
  { name: "Início", href: "/", icon: Home },
  { name: "Continuum", href: "/continuum", icon: Workflow },
  { name: "Consultas", href: "/consultas", icon: Calendar },
  { name: "Mensagens", href: "/mensagens", icon: MessageSquare },
  { name: "Eu", href: "/perfil", icon: UserIcon },
];

export function PatientBottomNav() {
  const pathname = usePathname();
  return (
    <nav className="fixed bottom-0 left-0 right-0 z-30 border-t border-border bg-card lg:hidden">
      <ul className="grid grid-cols-5">
        {items.map((it) => {
          const Icon = it.icon;
          const active =
            pathname === it.href ||
            (it.href !== "/" && pathname?.startsWith(it.href));
          return (
            <li key={it.name}>
              <Link
                href={it.href}
                className={cn(
                  "flex flex-col items-center gap-0.5 py-2 text-[10px] font-medium",
                  active ? "text-primary" : "text-muted-foreground",
                )}
              >
                <Icon className="h-5 w-5" />
                {it.name}
              </Link>
            </li>
          );
        })}
      </ul>
    </nav>
  );
}
