"use client";

import { useRouter } from "next/navigation";
import { differenceInYears, format } from "date-fns";
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
  const { selectedPatient } = useSelectedPatient();

  if (!selectedPatient) return null;

  const age = differenceInYears(new Date(), new Date(selectedPatient.birthDate));
  const dob = format(new Date(selectedPatient.birthDate), "dd/MM/yyyy");
  const gender = getGenderLabel(selectedPatient.gender);
  const shortId = selectedPatient.id.slice(-6).toUpperCase();
  const initials = getInitials(selectedPatient.name);

  return (
    <div className="flex h-10 items-center justify-between border-b border-blue-200 bg-blue-50/80 px-4 dark:border-blue-800 dark:bg-blue-950/30 sm:px-6">
      {/* Patient info */}
      <div className="flex items-center gap-2 text-sm overflow-hidden">
        {/* Avatar */}
        <span className="flex h-6 w-6 shrink-0 items-center justify-center rounded-full bg-blue-600 text-[10px] font-semibold text-white">
          {initials}
        </span>

        {/* Name */}
        <span className="font-semibold text-blue-900 dark:text-blue-100 truncate max-w-[200px]">
          {selectedPatient.name}
        </span>

        <span className="hidden text-blue-400 sm:inline">·</span>

        {/* DOB */}
        <span className="hidden text-muted-foreground sm:inline">{dob}</span>

        <span className="hidden text-blue-400 sm:inline">·</span>

        {/* Age */}
        <span className="hidden text-muted-foreground sm:inline">{age} anos</span>

        <span className="hidden text-blue-400 md:inline">·</span>

        {/* Gender */}
        <span className="hidden text-muted-foreground md:inline">{gender}</span>

        <span className="hidden text-blue-400 lg:inline">·</span>

        {/* Short ID */}
        <span className="hidden font-mono text-xs text-muted-foreground lg:inline">
          #{shortId}
        </span>
      </div>

      {/* Action */}
      <button
        onClick={() => router.push("/patients/select")}
        className="shrink-0 rounded px-2 py-0.5 text-xs font-medium text-blue-700 hover:bg-blue-100 dark:text-blue-300 dark:hover:bg-blue-900/50 transition-colors"
      >
        Trocar ▾
      </button>
    </div>
  );
}
