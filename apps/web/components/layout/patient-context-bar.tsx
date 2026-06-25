"use client";

import { useRouter, usePathname } from "next/navigation";
import { calcAge, formatDateOnly } from "@/lib/format-date";
import { useSelectedPatient } from "@/lib/use-selected-patient";

function getInitials(name: string): string {
  return name
    .split(" ")
    .slice(0, 2)
    .map((n) => n[0])
    .join("")
    .toUpperCase();
}

function getGenderLabel(gender: string): string {
  const labels: Record<string, string> = {
    male: "M",
    female: "F",
    other: "Outro",
  };
  return labels[gender] ?? gender;
}

export function PatientContextBar() {
  const router = useRouter();
  const pathname = usePathname();
  const { selectedPatient } = useSelectedPatient();

  if (!selectedPatient) return null;
  // Inbox de CRM (conversas) não é tela clínica — esconde o contexto de paciente
  // pra não confundir com o contato da conversa (lead/paciente da própria thread).
  if (pathname?.includes('/conversas')) return null;

  // Ao trocar de paciente, volta pra tela de origem (não pro dashboard). A tela de seleção
  // lê esta chave; mesmo padrão do redirect automático em use-require-selected-patient.
  const handleTrocar = () => {
    if (typeof window !== "undefined" && pathname && !pathname.startsWith("/patients/select")) {
      sessionStorage.setItem("plenya-redirect-after-patient-select", pathname);
    }
    router.push("/patients/select");
  };

  // birthDate pode vir vazio (paciente cadastrado sem data). O util é resiliente
  // a data inválida (sem ele, date-fns lança e derruba a barra em toda tela).
  const age = calcAge(selectedPatient.birthDate);
  const dob = selectedPatient.birthDate ? formatDateOnly(selectedPatient.birthDate, "dd/MM/yyyy", { fallback: "" }) : "";
  const gender = getGenderLabel(selectedPatient.gender);
  const shortId = selectedPatient.id.slice(-6).toUpperCase();
  const initials = getInitials(selectedPatient.name);

  return (
    <div className="flex h-10 items-center justify-between border-b border-ocean-200/70 bg-ocean-50/70 px-4 dark:border-petrol-700 dark:bg-petrol/30 sm:px-6">
      {/* Patient info */}
      <div className="flex items-center gap-2 text-sm overflow-hidden">
        {/* Avatar */}
        <span className="flex h-6 w-6 shrink-0 items-center justify-center rounded-full bg-gold text-[10px] font-semibold text-petrol">
          {initials}
        </span>

        {/* Name */}
        <span className="font-semibold text-petrol dark:text-cream truncate max-w-[200px]">
          {selectedPatient.name}
        </span>

        {dob && (
          <>
            <span className="hidden text-ocean-300 sm:inline">·</span>

            {/* DOB */}
            <span className="hidden text-muted-foreground sm:inline">{dob}</span>
          </>
        )}

        {age !== null && (
          <>
            <span className="hidden text-ocean-300 sm:inline">·</span>

            {/* Age */}
            <span className="hidden text-muted-foreground sm:inline">{age} anos</span>
          </>
        )}

        <span className="hidden text-ocean-300 md:inline">·</span>

        {/* Gender */}
        <span className="hidden text-muted-foreground md:inline">{gender}</span>

        <span className="hidden text-ocean-300 lg:inline">·</span>

        {/* Short ID */}
        <span className="hidden font-mono text-xs text-muted-foreground lg:inline">
          #{shortId}
        </span>
      </div>

      {/* Action */}
      <button
        onClick={handleTrocar}
        className="shrink-0 rounded px-2 py-0.5 text-xs font-medium text-ocean-700 hover:bg-ocean-100 dark:text-ocean-200 dark:hover:bg-petrol/50 transition-colors"
      >
        Trocar ▾
      </button>
    </div>
  );
}
