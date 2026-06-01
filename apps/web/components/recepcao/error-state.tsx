"use client";

import { AlertTriangle, RefreshCw, type LucideIcon } from "lucide-react";
import { Button } from "@/components/ui/button";

/**
 * Estado de erro — distinto do EmptyState (caloroso). Quando uma busca falha,
 * a secretaria precisa saber que houve falha de rede, e nao acreditar que a
 * agenda/lista esta vazia. Mostra um aviso sobrio com acao de tentar de novo.
 */
export function ErrorState({
  title = "Nao foi possivel carregar",
  hint = "Verifique a conexao e tente de novo.",
  onRetry,
  icon: Icon = AlertTriangle,
}: {
  title?: string;
  hint?: string;
  onRetry?: () => void;
  icon?: LucideIcon;
}) {
  return (
    <div className="flex flex-col items-center justify-center py-10 text-center">
      <div className="flex h-12 w-12 items-center justify-center rounded-full bg-destructive/10 text-destructive">
        <Icon className="h-6 w-6" />
      </div>
      <p className="mt-3 text-sm font-medium text-foreground">{title}</p>
      {hint && <p className="mt-1 text-sm text-muted-foreground">{hint}</p>}
      {onRetry && (
        <Button
          variant="outline"
          size="sm"
          className="mt-3 gap-2"
          onClick={onRetry}
        >
          <RefreshCw className="h-4 w-4" />
          Tentar de novo
        </Button>
      )}
    </div>
  );
}
