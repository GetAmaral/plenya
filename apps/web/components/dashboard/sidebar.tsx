"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";
import { motion } from "framer-motion";
import {
  Activity,
  Calendar,
  FileText,
  Home,
  Users,
  Microscope,
  LogOut,
  Stethoscope,
  Network,
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
  { name: "Escores", href: "/scores", icon: Network },
];

export function Sidebar() {
  const pathname = usePathname();
  const { user, logout } = useAuth();

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

  return (
    <motion.aside
      initial={{ x: -300 }}
      animate={{ x: 0 }}
      className="fixed left-0 top-0 z-40 h-screen w-64 border-r border-border bg-card"
    >
      <div className="flex h-full flex-col">
        {/* Logo */}
        <div className="flex h-16 items-center gap-3 border-b border-border px-6">
          <div className="flex h-10 w-10 items-center justify-center rounded-xl bg-gradient-to-br from-blue-500 to-blue-700 shadow-md">
            <Activity className="h-5 w-5 text-white" />
          </div>
          <div>
            <h1 className="text-lg font-bold">Plenya EMR</h1>
            <p className="text-xs text-muted-foreground">Medical System</p>
          </div>
        </div>

        {/* Navigation */}
        <nav className="flex-1 space-y-1 overflow-y-auto p-4">
          {navigation.map((item) => {
            const Icon = item.icon;
            const isActive = pathname === item.href;

            return (
              <Link key={item.name} href={item.href}>
                <motion.div
                  whileHover={{ x: 4 }}
                  whileTap={{ scale: 0.98 }}
                  className={cn(
                    "group flex items-center gap-3 rounded-lg px-3 py-2.5 text-sm font-medium transition-colors",
                    isActive
                      ? "bg-primary text-primary-foreground shadow-sm"
                      : "text-muted-foreground hover:bg-accent hover:text-accent-foreground"
                  )}
                >
                  <Icon className="h-5 w-5 shrink-0" />
                  <span>{item.name}</span>
                </motion.div>
              </Link>
            );
          })}
        </nav>

        {/* User section */}
        <div className="border-t border-border p-4">
          <div className="flex items-center gap-3 rounded-lg bg-accent/50 p-3">
            <Avatar className="h-10 w-10">
              <AvatarFallback className="bg-gradient-to-br from-blue-500 to-blue-700 text-white text-sm font-semibold">
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

          <motion.button
            whileHover={{ scale: 1.02 }}
            whileTap={{ scale: 0.98 }}
            onClick={logout}
            className="mt-2 flex w-full items-center gap-2 rounded-lg px-3 py-2 text-sm font-medium text-muted-foreground transition-colors hover:bg-destructive/10 hover:text-destructive"
          >
            <LogOut className="h-4 w-4" />
            Sair
          </motion.button>
        </div>
      </div>
    </motion.aside>
  );
}
