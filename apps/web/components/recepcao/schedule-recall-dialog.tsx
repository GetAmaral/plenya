"use client";

import { useState } from "react";
import { addMonths, format } from "date-fns";
import { toast } from "sonner";
import { CalendarClock, Loader2 } from "lucide-react";

import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogDescription,
  DialogFooter,
} from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useCreateRecall } from "@/lib/api/recalls";

function isoDate(d: Date): string {
  return format(d, "yyyy-MM-dd");
}

/**
 * Dialog de "Agendar retorno": registra um retorno previsto (recall) que entra
 * na fila interna da recepção. Sem aviso automático ao paciente — a secretária
 * contata e agenda manualmente. Reutilizável (detalhe da consulta, ficha).
 */
export function ScheduleRecallDialog({
  open,
  onOpenChange,
  patientId,
  doctorId,
  sourceAppointmentId,
  patientName,
}: {
  open: boolean;
  onOpenChange: (o: boolean) => void;
  patientId: string;
  doctorId?: string;
  sourceAppointmentId?: string;
  patientName?: string;
}) {
  const create = useCreateRecall();
  const [dueDate, setDueDate] = useState("");
  const [reason, setReason] = useState("");

  function quick(months: number) {
    setDueDate(isoDate(addMonths(new Date(), months)));
  }

  function handleSave() {
    if (!dueDate) {
      toast.error("Escolha a data do retorno");
      return;
    }
    create.mutate(
      {
        patientId,
        doctorId,
        sourceAppointmentId,
        dueDate,
        reason: reason.trim() || undefined,
      },
      {
        onSuccess: () => {
          toast.success("Retorno registrado", {
            description: patientName
              ? `${patientName} em ${dueDate.split("-").reverse().join("/")}.`
              : undefined,
          });
          setDueDate("");
          setReason("");
          onOpenChange(false);
        },
        onError: (e: unknown) =>
          toast.error(
            e instanceof Error ? e.message : "Falha ao registrar retorno",
          ),
      },
    );
  }

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Agendar retorno</DialogTitle>
          <DialogDescription>
            Registra um retorno previsto
            {patientName ? ` de ${patientName}` : ""}. Entra na fila de retornos
            da recepção para contato e agendamento.
          </DialogDescription>
        </DialogHeader>
        <div className="space-y-4">
          <div className="space-y-2">
            <Label>Quando</Label>
            <div className="flex flex-wrap gap-2">
              <Button type="button" size="sm" variant="outline" onClick={() => quick(3)}>
                3 meses
              </Button>
              <Button type="button" size="sm" variant="outline" onClick={() => quick(6)}>
                6 meses
              </Button>
              <Button type="button" size="sm" variant="outline" onClick={() => quick(12)}>
                1 ano
              </Button>
            </div>
            <Input
              type="date"
              value={dueDate}
              onChange={(e) => setDueDate(e.target.value)}
            />
          </div>
          <div className="space-y-2">
            <Label htmlFor="recall-reason">Motivo (opcional)</Label>
            <Input
              id="recall-reason"
              value={reason}
              onChange={(e) => setReason(e.target.value)}
              placeholder="Ex: reavaliação anual de biomarcadores"
            />
          </div>
        </div>
        <DialogFooter>
          <Button variant="outline" onClick={() => onOpenChange(false)}>
            Cancelar
          </Button>
          <Button onClick={handleSave} disabled={create.isPending} className="gap-2">
            {create.isPending ? (
              <Loader2 className="h-4 w-4 animate-spin" />
            ) : (
              <CalendarClock className="h-4 w-4" />
            )}
            Registrar retorno
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}

export default ScheduleRecallDialog;
