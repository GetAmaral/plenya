"use client";

import { useState } from "react";
import { formatDate } from "@/lib/format-date";
import { Loader2, Receipt, RotateCcw } from "lucide-react";
import { toast } from "sonner";

import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { Skeleton } from "@/components/ui/skeleton";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogDescription,
  DialogFooter,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useAuthStore, isGranted } from "@/lib/auth-store";
import {
  usePayments,
  useRefundPayment,
  openReceipt,
  formatBRL,
  PAYMENT_METHOD_LABELS,
  type Payment,
} from "@/lib/api/payments";

function fmtDate(iso: string): string {
  return formatDate(iso, "dd/MM/yyyy");
}

function PaymentRow({
  payment,
  canRefund,
}: {
  payment: Payment;
  canRefund: boolean;
}) {
  const refund = useRefundPayment(payment.id);
  const [open, setOpen] = useState(false);
  const [reason, setReason] = useState("");
  const [opening, setOpening] = useState(false);
  const refunded = payment.status === "refunded";

  async function handleReceipt() {
    setOpening(true);
    try {
      await openReceipt(payment.id);
    } catch {
      toast.error("Falha ao abrir o recibo");
    } finally {
      setOpening(false);
    }
  }

  function handleRefund() {
    if (reason.trim().length < 3) {
      toast.error("Informe o motivo do estorno");
      return;
    }
    refund.mutate(reason.trim(), {
      onSuccess: () => {
        toast.success("Pagamento estornado");
        setOpen(false);
        setReason("");
      },
      onError: (e: unknown) =>
        toast.error(e instanceof Error ? e.message : "Falha ao estornar"),
    });
  }

  return (
    <tr className="border-t">
      <td className="px-3 py-2 text-sm">{fmtDate(payment.paidAt)}</td>
      <td className="px-3 py-2 text-sm font-medium">
        {formatBRL(payment.amountCents)}
      </td>
      <td className="px-3 py-2 text-sm">
        {PAYMENT_METHOD_LABELS[payment.method]}
      </td>
      <td className="px-3 py-2 text-xs text-muted-foreground">
        {payment.receiptNumber}
      </td>
      <td className="px-3 py-2">
        <Badge
          variant="outline"
          className={
            refunded
              ? "border-stone-300 text-stone-600"
              : "border-emerald-200 bg-emerald-100 text-emerald-900"
          }
        >
          {refunded ? "Estornado" : "Pago"}
        </Badge>
      </td>
      <td className="px-3 py-2">
        <div className="flex justify-end gap-1">
          <Button
            size="sm"
            variant="ghost"
            onClick={handleReceipt}
            disabled={opening}
          >
            {opening ? (
              <Loader2 className="h-4 w-4 animate-spin" />
            ) : (
              <Receipt className="h-4 w-4" />
            )}
            <span className="hidden sm:inline">Recibo</span>
          </Button>
          {canRefund && !refunded && (
            <Button
              size="sm"
              variant="ghost"
              onClick={() => setOpen(true)}
              title="Estornar pagamento"
            >
              <RotateCcw className="h-4 w-4" />
            </Button>
          )}
        </div>

        <Dialog open={open} onOpenChange={setOpen}>
          <DialogContent>
            <DialogHeader>
              <DialogTitle>Estornar pagamento</DialogTitle>
              <DialogDescription>
                Recibo {payment.receiptNumber} · {formatBRL(payment.amountCents)}.
                Informe o motivo do estorno.
              </DialogDescription>
            </DialogHeader>
            <div className="space-y-2">
              <Label htmlFor={`refund-${payment.id}`}>Motivo</Label>
              <Input
                id={`refund-${payment.id}`}
                value={reason}
                onChange={(e) => setReason(e.target.value)}
                placeholder="Ex: cobrança em duplicidade"
                autoFocus
              />
            </div>
            <DialogFooter>
              <Button variant="outline" onClick={() => setOpen(false)}>
                Cancelar
              </Button>
              <Button
                variant="destructive"
                onClick={handleRefund}
                disabled={refund.isPending}
              >
                {refund.isPending ? "Estornando..." : "Confirmar estorno"}
              </Button>
            </DialogFooter>
          </DialogContent>
        </Dialog>
      </td>
    </tr>
  );
}

/**
 * Histórico financeiro do paciente: pagamentos com data, valor, método, status,
 * reimpressão de recibo (PDF) e estorno (admin/manager). Antes disso, um erro de
 * cobrança só se corrigia via banco direto.
 */
export function PatientFinanceCard({ patientId }: { patientId: string }) {
  const { user } = useAuthStore();
  const canRefund = isGranted(user, "admin") || isGranted(user, "manager");
  const { data: payments = [], isLoading, isError, refetch } =
    usePayments(patientId);

  const totalPaid = payments
    .filter((p) => p.status === "paid")
    .reduce((sum, p) => sum + p.amountCents, 0);

  return (
    <Card className="border-0 shadow-lg md:col-span-2">
      <CardHeader>
        <CardTitle className="flex items-center justify-between">
          <span>Financeiro</span>
          {!isLoading && !isError && payments.length > 0 && (
            <span className="text-sm font-normal text-muted-foreground">
              Total pago: {formatBRL(totalPaid)}
            </span>
          )}
        </CardTitle>
      </CardHeader>
      <CardContent>
        {isLoading ? (
          <div className="space-y-2">
            {Array.from({ length: 3 }).map((_, i) => (
              <Skeleton key={i} className="h-9 w-full" />
            ))}
          </div>
        ) : isError ? (
          <div className="py-6 text-center text-sm text-muted-foreground">
            Não foi possível carregar o financeiro.{" "}
            <Button
              variant="link"
              className="h-auto p-0"
              onClick={() => refetch()}
            >
              Tentar de novo
            </Button>
          </div>
        ) : payments.length === 0 ? (
          <p className="py-6 text-center text-sm text-muted-foreground">
            Nenhum pagamento registrado para este paciente.
          </p>
        ) : (
          <div className="overflow-x-auto rounded-md border">
            <table className="w-full">
              <thead className="bg-muted/50 text-left text-xs font-semibold uppercase tracking-wide text-muted-foreground">
                <tr>
                  <th className="px-3 py-2">Data</th>
                  <th className="px-3 py-2">Valor</th>
                  <th className="px-3 py-2">Método</th>
                  <th className="px-3 py-2">Recibo</th>
                  <th className="px-3 py-2">Status</th>
                  <th className="px-3 py-2 text-right">Ações</th>
                </tr>
              </thead>
              <tbody>
                {payments.map((p) => (
                  <PaymentRow key={p.id} payment={p} canRefund={canRefund} />
                ))}
              </tbody>
            </table>
          </div>
        )}
      </CardContent>
    </Card>
  );
}

export default PatientFinanceCard;
