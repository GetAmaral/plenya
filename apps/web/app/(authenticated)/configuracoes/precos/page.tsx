"use client";

import { useState } from "react";
import { toast } from "sonner";
import { DollarSign, Loader2, Save } from "lucide-react";

import { PageHeader } from "@/components/layout/page-header";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Skeleton } from "@/components/ui/skeleton";
import { useRequireAuth } from "@/lib/use-auth";
import { useAuthStore, isGranted } from "@/lib/auth-store";
import {
  APPOINTMENT_TYPE_LABELS,
  type AppointmentType,
} from "@/lib/api/calendar-api";
import {
  useConsultationPrices,
  useUpsertConsultationPrice,
  formatBRL,
  type ConsultationPrice,
} from "@/lib/api/payments";

const TYPES = Object.keys(APPOINTMENT_TYPE_LABELS) as AppointmentType[];

// "350,00" / "350.00" / "1.350,50" -> centavos. Null se inválido.
function parseReaisToCents(input: string): number | null {
  const cleaned = input
    .replace(/[^\d,.-]/g, "")
    .replace(/\./g, "")
    .replace(",", ".");
  const n = parseFloat(cleaned);
  if (isNaN(n) || n < 0) return null;
  return Math.round(n * 100);
}

function PriceRow({
  type,
  price,
}: {
  type: AppointmentType;
  price?: ConsultationPrice;
}) {
  const upsert = useUpsertConsultationPrice();
  const [value, setValue] = useState(
    price ? (price.amountCents / 100).toFixed(2).replace(".", ",") : "",
  );

  function handleSave() {
    const cents = parseReaisToCents(value);
    if (cents === null) {
      toast.error("Valor inválido");
      return;
    }
    upsert.mutate(
      { type, amountCents: cents, active: true },
      {
        onSuccess: () =>
          toast.success(`Preço de ${APPOINTMENT_TYPE_LABELS[type]} salvo`),
        onError: (e: unknown) =>
          toast.error(e instanceof Error ? e.message : "Falha ao salvar"),
      },
    );
  }

  return (
    <div className="flex flex-wrap items-end justify-between gap-3 rounded-lg border p-3">
      <div className="min-w-0">
        <p className="font-medium">{APPOINTMENT_TYPE_LABELS[type]}</p>
        <p className="text-sm text-muted-foreground">
          {price ? `Atual: ${formatBRL(price.amountCents)}` : "Sem preço definido"}
        </p>
      </div>
      <div className="flex items-end gap-2">
        <div className="space-y-1">
          <Label htmlFor={`price-${type}`} className="text-xs">
            Valor (R$)
          </Label>
          <Input
            id={`price-${type}`}
            value={value}
            onChange={(e) => setValue(e.target.value)}
            placeholder="350,00"
            className="w-32"
            inputMode="decimal"
          />
        </div>
        <Button onClick={handleSave} disabled={upsert.isPending} className="gap-2">
          {upsert.isPending ? (
            <Loader2 className="h-4 w-4 animate-spin" />
          ) : (
            <Save className="h-4 w-4" />
          )}
          Salvar
        </Button>
      </div>
    </div>
  );
}

export default function PrecosPage() {
  useRequireAuth();
  const { user } = useAuthStore();
  const canEdit = isGranted(user, "admin") || isGranted(user, "manager");
  const { data: prices = [], isLoading } = useConsultationPrices();

  return (
    <div className="container mx-auto space-y-6 py-8">
      <PageHeader
        breadcrumbs={[{ label: "Configurações" }, { label: "Preços" }]}
        title="Preços de Consulta"
        description="Valores por tipo de consulta. Pré-preenchem o registro de pagamento na recepção."
      />

      {!canEdit ? (
        <Card>
          <CardContent className="py-8 text-center text-sm text-muted-foreground">
            Apenas admin e gestor podem editar a tabela de preços.
          </CardContent>
        </Card>
      ) : (
        <Card>
          <CardHeader>
            <CardTitle className="flex items-center gap-2 text-base">
              <DollarSign className="h-4 w-4" /> Tabela de preços
            </CardTitle>
          </CardHeader>
          <CardContent className="space-y-3">
            {isLoading
              ? Array.from({ length: 5 }).map((_, i) => (
                  <Skeleton key={i} className="h-16 w-full" />
                ))
              : TYPES.map((t) => (
                  <PriceRow
                    key={t}
                    type={t}
                    price={prices.find((p) => p.type === t)}
                  />
                ))}
          </CardContent>
        </Card>
      )}
    </div>
  );
}
