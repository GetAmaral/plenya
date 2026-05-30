import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { apiClient } from "@/lib/api-client";
import type { AppointmentType } from "@/lib/api/calendar-api";

// ===== Types =====

export type PaymentMethod = "cash" | "pix" | "card" | "transfer" | "other";
export type PaymentStatus = "paid" | "refunded";

export interface Payment {
  id: string;
  patientId: string;
  appointmentId?: string;
  amountCents: number;
  method: PaymentMethod;
  status: PaymentStatus;
  paidAt: string;
  receiptNumber: string;
  notes?: string;
  refundedAt?: string;
  refundReason?: string;
  createdByUserId: string;
  createdAt: string;
  patient?: { id: string; name: string; email?: string; phone?: string };
}

export interface CreatePaymentInput {
  patientId: string;
  appointmentId?: string;
  amountCents: number;
  method: PaymentMethod;
  paidAt?: string;
  notes?: string;
}

export interface ConsultationPrice {
  id: string;
  type: AppointmentType;
  amountCents: number;
  active: boolean;
  createdAt: string;
  updatedAt: string;
}

export const PAYMENT_METHOD_LABELS: Record<PaymentMethod, string> = {
  cash: "Dinheiro",
  pix: "Pix",
  card: "Cartão",
  transfer: "Transferência",
  other: "Outro",
};

// ===== Query keys =====

export const paymentKeys = {
  all: ["payments"] as const,
  list: (patientId?: string) => [...paymentKeys.all, "list", patientId] as const,
  prices: ["consultation-prices"] as const,
};

// ===== Queries =====

export function usePayments(patientId?: string) {
  return useQuery({
    queryKey: paymentKeys.list(patientId),
    queryFn: () => {
      const qs = patientId ? `?patientId=${patientId}` : "";
      return apiClient.get<Payment[]>(`/api/v1/payments${qs}`);
    },
    staleTime: 30_000,
  });
}

export function useConsultationPrices() {
  return useQuery({
    queryKey: paymentKeys.prices,
    queryFn: () =>
      apiClient.get<ConsultationPrice[]>("/api/v1/consultation-prices"),
    staleTime: 60_000,
  });
}

// ===== Mutations =====

export function useCreatePayment() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (data: CreatePaymentInput) =>
      apiClient.post<Payment>("/api/v1/payments", data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: paymentKeys.all });
    },
  });
}

export function useRefundPayment(id: string) {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (reason: string) =>
      apiClient.post<Payment>(`/api/v1/payments/${id}/refund`, { reason }),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: paymentKeys.all });
    },
  });
}

export function useUpsertConsultationPrice() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (data: {
      type: AppointmentType;
      amountCents: number;
      active?: boolean;
    }) => apiClient.put<ConsultationPrice>("/api/v1/consultation-prices", data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: paymentKeys.prices });
    },
  });
}

// Busca o recibo (PDF, atrás de auth) como blob e abre em nova aba.
export async function openReceipt(paymentId: string): Promise<void> {
  const blob = await apiClient.getBlob(`/api/v1/payments/${paymentId}/receipt`);
  const url = URL.createObjectURL(blob);
  window.open(url, "_blank");
  // Revoga depois de um tempo para o navegador conseguir abrir.
  setTimeout(() => URL.revokeObjectURL(url), 60_000);
}

// Formata centavos em "R$ 1.234,56".
export function formatBRL(cents: number): string {
  return (cents / 100).toLocaleString("pt-BR", {
    style: "currency",
    currency: "BRL",
  });
}
