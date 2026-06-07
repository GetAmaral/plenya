"use client";

import { useMemo, useRef, useState } from "react";
import { useRouter } from "next/navigation";
import { motion } from "framer-motion";
import { toast } from "sonner";
import { CalendarPlus, UserPlus, Search } from "lucide-react";
import { PageHeader } from "@/components/layout/page-header";
import { Button } from "@/components/ui/button";
import { useAppointments, useDoctors } from "@/lib/api/calendar-api";
import type { Appointment } from "@/lib/api/calendar-api";
import { useDayPayments, type Payment } from "@/lib/api/payments";
import { useSelectedPatient } from "@/lib/use-selected-patient";
import { QuickAddPatientDialog } from "@/components/patients/quick-add-patient-dialog";
import { AgendaHojeCard } from "@/components/recepcao/agenda-hoje-card";
import { AConfirmarCard } from "@/components/recepcao/a-confirmar-card";
import { MensagensCard } from "@/components/recepcao/mensagens-card";
import { LeadsNovosCard } from "@/components/recepcao/leads-novos-card";
import { ListaEsperaCard } from "@/components/recepcao/lista-espera-card";
import { RetornosCard } from "@/components/recepcao/retornos-card";
import { ProximosEventosCard } from "@/components/recepcao/proximos-eventos-card";
import { PatientQuickSearch } from "@/components/recepcao/patient-quick-search";
import {
  addDays,
  formatFullDatePtBR,
  startOfTodaySaoPaulo,
} from "@/components/recepcao/recepcao-helpers";

const container = {
  hidden: { opacity: 0 },
  show: { opacity: 1, transition: { staggerChildren: 0.08 } },
};
const item = {
  hidden: { opacity: 0, y: 16 },
  show: { opacity: 1, y: 0 },
};

/**
 * Recepcao · Hoje — cockpit do front-desk.
 *
 * Junta numa unica tela o que a secretaria buscava em Calendario, Conversas e
 * Leads: agenda do dia (todos os medicos), confirmacoes pendentes (hoje +
 * amanha), mensagens nao lidas e leads novos, mais acoes rapidas.
 *
 * Polling: as consultas vem de useAppointments (refetchInterval 15s, igual ao
 * Calendario). Mensagens, leads e lista de espera tambem fazem polling (~30s)
 * nos seus proprios hooks, entao o cockpit inteiro fica vivo sem refresh manual.
 */
export default function RecepcaoPage() {
  const router = useRouter();
  const searchRef = useRef<HTMLInputElement>(null);
  const [quickAddOpen, setQuickAddOpen] = useState(false);
  const { setSelectedPatient } = useSelectedPatient();

  const {
    data: doctors = [],
    isLoading: doctorsLoading,
    isError: doctorsError,
    refetch: refetchDoctors,
  } = useDoctors();
  const doctorIds = useMemo(() => doctors.map((d) => d.id), [doctors]);

  // Janela [hoje 00:00, depois-de-amanha 00:00): cobre hoje (Agenda) e
  // amanha (A confirmar) numa unica busca. Sem horario de verao no Brasil.
  const { dateFrom, dateTo, todayStartMs, tomorrowEndMs } = useMemo(() => {
    const start = startOfTodaySaoPaulo();
    const tomorrow = addDays(start, 1);
    const dayAfter = addDays(start, 2);
    return {
      dateFrom: start.toISOString(),
      dateTo: dayAfter.toISOString(),
      todayStartMs: start.getTime(),
      tomorrowEndMs: tomorrow.getTime() + 24 * 60 * 60 * 1000 - 1, // fim de amanha
    };
  }, []);

  const {
    data: appointments = [],
    isLoading: appointmentsLoading,
    isError: appointmentsError,
    refetch: refetchAppointments,
  } = useAppointments({
    doctorIds,
    dateFrom,
    dateTo,
    limit: 500,
  });

  // Enquanto a lista de medicos nao resolve, useAppointments fica desabilitado
  // (enabled=false => isLoading false). Tratamos como "carregando" pra mostrar
  // skeletons em vez de piscar estados vazios.
  const isLoading = doctorsLoading || appointmentsLoading;
  const isError = doctorsError || appointmentsError;
  const retryAgenda = () => {
    refetchDoctors();
    refetchAppointments();
  };

  // Pagamentos do dia → mapa por consulta, pra marcar "Pago" na agenda e evitar
  // cobrança em duplicidade. Mesma janela das consultas (hoje + amanhã).
  const { data: dayPayments = [] } = useDayPayments(dateFrom, dateTo);
  const paidByAppointment = useMemo(() => {
    const m = new Map<string, Payment>();
    for (const p of dayPayments) {
      if (p.status === "paid" && p.appointmentId) m.set(p.appointmentId, p);
    }
    return m;
  }, [dayPayments]);

  // Hoje vs hoje+amanha — derivados client-side da mesma janela.
  const todayStart = todayStartMs;
  const tomorrowStart = todayStartMs + 24 * 60 * 60 * 1000;

  const todayAppointments = useMemo(
    () =>
      (appointments as Appointment[]).filter((a) => {
        const t = new Date(a.scheduledAt).getTime();
        return t >= todayStart && t < tomorrowStart;
      }),
    [appointments, todayStart, tomorrowStart],
  );

  const todayAndTomorrow = useMemo(
    () =>
      (appointments as Appointment[]).filter((a) => {
        const t = new Date(a.scheduledAt).getTime();
        return t >= todayStart && t <= tomorrowEndMs;
      }),
    [appointments, todayStart, tomorrowEndMs],
  );

  const today = new Date();
  const fullDate = formatFullDatePtBR(today);

  return (
    <div className="container mx-auto py-8 space-y-8">
      <PageHeader title="Recepção" description={fullDate} />

      {/* Barra de acoes rapidas */}
      <motion.div
        initial={{ opacity: 0, y: -8 }}
        animate={{ opacity: 1, y: 0 }}
        className="flex flex-col gap-3 lg:flex-row lg:items-center"
      >
        <div className="flex flex-wrap gap-2">
          <Button onClick={() => router.push("/appointments/new")}>
            <CalendarPlus className="h-4 w-4" />
            Nova consulta
          </Button>
          <Button variant="outline" onClick={() => setQuickAddOpen(true)}>
            <UserPlus className="h-4 w-4" />
            Novo paciente
          </Button>
          <Button
            variant="ghost"
            onClick={() => searchRef.current?.focus()}
          >
            <Search className="h-4 w-4" />
            Buscar paciente
          </Button>
        </div>
        <div className="lg:ml-auto lg:w-96">
          <PatientQuickSearch ref={searchRef} />
        </div>
      </motion.div>

      {/* Grade principal */}
      <motion.div
        variants={container}
        initial="hidden"
        animate="show"
        className="grid gap-6 lg:grid-cols-3"
      >
        {/* Coluna larga: agenda do dia */}
        <motion.div variants={item} className="lg:col-span-2">
          <AgendaHojeCard
            appointments={todayAppointments}
            isLoading={isLoading}
            isError={isError}
            onRetry={retryAgenda}
            paidByAppointment={paidByAppointment}
          />
        </motion.div>

        {/* Coluna lateral: a confirmar */}
        <motion.div variants={item}>
          <AConfirmarCard
            appointments={todayAndTomorrow}
            isLoading={isLoading}
            isError={isError}
            onRetry={retryAgenda}
          />
        </motion.div>

        {/* Mensagens */}
        <motion.div variants={item}>
          <MensagensCard />
        </motion.div>

        {/* Leads novos */}
        <motion.div variants={item}>
          <LeadsNovosCard />
        </motion.div>

        {/* Lista de espera */}
        <motion.div variants={item}>
          <ListaEsperaCard />
        </motion.div>

        {/* Retornos previstos */}
        <motion.div variants={item}>
          <RetornosCard />
        </motion.div>

        <motion.div variants={item}>
          <ProximosEventosCard />
        </motion.div>
      </motion.div>

      <QuickAddPatientDialog
        open={quickAddOpen}
        onOpenChange={setQuickAddOpen}
        onCreated={(patient) => {
          toast.success("Paciente cadastrado", {
            description: `Vamos agendar a consulta de ${patient.name}.`,
          });
          // Coloca o novo paciente em contexto e encadeia para Nova consulta.
          setSelectedPatient(patient.id);
          router.push("/appointments/new");
        }}
      />
    </div>
  );
}
