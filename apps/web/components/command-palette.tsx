"use client";

import * as React from "react";
import { useRouter } from "next/navigation";
import {
  CommandDialog,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
  CommandSeparator,
} from "@/components/ui/command";
import {
  LayoutDashboard,
  Users,
  Calendar,
  ClipboardList,
  FileText,
  TestTube,
  LogOut,
} from "lucide-react";
import { useAuth } from "@/lib/use-auth";

export function CommandPalette() {
  const [open, setOpen] = React.useState(false);
  const router = useRouter();
  const { logout } = useAuth();

  React.useEffect(() => {
    const down = (e: KeyboardEvent) => {
      if (e.key === "k" && (e.metaKey || e.ctrlKey)) {
        e.preventDefault();
        setOpen((open) => !open);
      }
    };

    document.addEventListener("keydown", down);
    return () => document.removeEventListener("keydown", down);
  }, []);

  const runCommand = React.useCallback((command: () => void) => {
    setOpen(false);
    command();
  }, []);

  return (
    <CommandDialog open={open} onOpenChange={setOpen}>
      <CommandInput placeholder="Digite um comando ou pesquise..." />
      <CommandList>
        <CommandEmpty>Nenhum resultado encontrado.</CommandEmpty>
        <CommandGroup heading="Navegação">
          <CommandItem
            onSelect={() => runCommand(() => router.push("/dashboard"))}
          >
            <LayoutDashboard className="mr-2 h-4 w-4" />
            <span>Dashboard</span>
          </CommandItem>
          <CommandItem
            onSelect={() => runCommand(() => router.push("/patients"))}
          >
            <Users className="mr-2 h-4 w-4" />
            <span>Pacientes</span>
          </CommandItem>
          <CommandItem
            onSelect={() => runCommand(() => router.push("/appointments"))}
          >
            <Calendar className="mr-2 h-4 w-4" />
            <span>Consultas</span>
          </CommandItem>
          <CommandItem
            onSelect={() => runCommand(() => router.push("/anamnesis"))}
          >
            <ClipboardList className="mr-2 h-4 w-4" />
            <span>Anamnese</span>
          </CommandItem>
          <CommandItem
            onSelect={() => runCommand(() => router.push("/prescriptions"))}
          >
            <FileText className="mr-2 h-4 w-4" />
            <span>Prescrições</span>
          </CommandItem>
          <CommandItem
            onSelect={() => runCommand(() => router.push("/lab-results"))}
          >
            <TestTube className="mr-2 h-4 w-4" />
            <span>Exames Laboratoriais</span>
          </CommandItem>
        </CommandGroup>
        <CommandSeparator />
        <CommandGroup heading="Ações">
          <CommandItem onSelect={() => runCommand(() => logout())}>
            <LogOut className="mr-2 h-4 w-4" />
            <span>Sair</span>
          </CommandItem>
        </CommandGroup>
      </CommandList>
    </CommandDialog>
  );
}
