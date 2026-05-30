"use client";

import { useEffect, useState } from "react";
import { toast } from "sonner";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogFooter,
  DialogDescription,
} from "@/components/ui/dialog";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import {
  useCreatePayment,
  useConsultationPrices,
  openReceipt,
  PAYMENT_METHOD_LABELS,
  type PaymentMethod,
} from "@/lib/api/payments";
import type { AppointmentType } from "@/lib/api/calendar-api";

interface RegisterPaymentDialogProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  patientId: string;
  patientName?: string;
  appointmentId?: string;
  appointmentType?: AppointmentType;
}

export function RegisterPaymentDialog({
  open,
  onOpenChange,
  patientId,
  patientName,
  appointmentId,
  appointmentType,
}: RegisterPaymentDialogProps) {
  const { data: prices } = useConsultationPrices();
  const createMutation = useCreatePayment();

  const [amount, setAmount] = useState(""); // em reais, string
  const [method, setMethod] = useState<PaymentMethod>("pix");
  const [notes, setNotes] = useState("");

  // Pré-preenche o valor pelo tipo de consulta (preço cadastrado).
  useEffect(() => {
    if (!open || !appointmentType || !prices) return;
    const p = prices.find((x) => x.type === appointmentType && x.active);
    if (p) setAmount((p.amountCents / 100).toFixed(2).replace(".", ","));
  }, [open, appointmentType, prices]);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    const cents = Math.round(
      parseFloat(amount.replace(/\./g, "").replace(",", ".")) * 100,
    );
    if (!cents || cents <= 0) {
      toast.error("Informe um valor válido");
      return;
    }
    createMutation.mutate(
      {
        patientId,
        appointmentId,
        amountCents: cents,
        method,
        notes: notes.trim() || undefined,
      },
      {
        onSuccess: (payment) => {
          toast.success(`Pagamento registrado (recibo ${payment.receiptNumber})`);
          setAmount("");
          setNotes("");
          onOpenChange(false);
          // Abre o recibo automaticamente.
          openReceipt(payment.id).catch(() => {
            toast.error("Falha ao abrir o recibo");
          });
        },
        onError: (err: unknown) => {
          toast.error(err instanceof Error ? err.message : "Falha ao registrar");
        },
      },
    );
  };

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Registrar pagamento</DialogTitle>
          <DialogDescription>
            {patientName ? `Paciente: ${patientName}` : "Pagamento de consulta"}
          </DialogDescription>
        </DialogHeader>
        <form onSubmit={handleSubmit} className="space-y-4">
          <div className="space-y-2">
            <Label htmlFor="pay-amount">Valor (R$)</Label>
            <Input
              id="pay-amount"
              value={amount}
              onChange={(e) => setAmount(e.target.value)}
              placeholder="0,00"
              inputMode="decimal"
              autoFocus
            />
          </div>
          <div className="space-y-2">
            <Label>Forma de pagamento</Label>
            <Select
              value={method}
              onValueChange={(v) => setMethod(v as PaymentMethod)}
            >
              <SelectTrigger>
                <SelectValue />
              </SelectTrigger>
              <SelectContent>
                {Object.entries(PAYMENT_METHOD_LABELS).map(([k, label]) => (
                  <SelectItem key={k} value={k}>
                    {label}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
          </div>
          <div className="space-y-2">
            <Label htmlFor="pay-notes">Observações</Label>
            <Input
              id="pay-notes"
              value={notes}
              onChange={(e) => setNotes(e.target.value)}
              placeholder="Opcional"
            />
          </div>
          <DialogFooter>
            <Button
              type="button"
              variant="outline"
              onClick={() => onOpenChange(false)}
            >
              Cancelar
            </Button>
            <Button type="submit" disabled={createMutation.isPending}>
              {createMutation.isPending
                ? "Registrando..."
                : "Registrar e emitir recibo"}
            </Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  );
}

export default RegisterPaymentDialog;
