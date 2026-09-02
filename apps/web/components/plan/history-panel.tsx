"use client";

import { useState } from "react";
import { History, RotateCcw, Sparkles, Upload, User } from "lucide-react";
import type { PlanRevision } from "@/lib/api/patient-plans";

import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { ScrollArea } from "@/components/ui/scroll-area";
import {
  Sheet,
  SheetContent,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
} from "@/components/ui/sheet";
import { cn } from "@/lib/utils";

/**
 * O histórico do rascunho, e o desfazer.
 *
 * Toda gravação já virava linha no banco desde o começo, mas ninguém no EMR conseguia ver: sem esta
 * tela, "quem escreveu esta frase" e "volta como estava antes da ferramenta mexer" continuavam
 * sendo consulta SQL.
 *
 * Fica num Sheet, e não numa quarta coluna, porque num laptop de 1366 as três colunas já disputam
 * espaço com a sidebar de 256px.
 */
const MOTIVO: Record<PlanRevision["reason"], string> = {
  edit: "edição",
  ai_apply: "aplicado pela ferramenta",
  suggestion_accept: "sugestão aceita",
  restore: "restauração",
  publish: "publicação",
};

/**
 * Op estrutural não tem caminho de campo, e o servidor grava o VERBO no lugar. Renderizado como se
 * fosse campo, saía "alterou 1 campo: add", que não diz nada a quem lê.
 */
const VERBO: Record<string, string> = {
  add: "acrescentou um slide",
  remove: "removeu um slide",
  reorder: "mudou a ordem dos slides",
};

function descreveMudanca(caminhos: string[]) {
  const campos: string[] = [];
  const estruturais: string[] = [];
  for (const c of caminhos) {
    const ultimo = c.split(":").pop() ?? "";
    if (VERBO[ultimo]) estruturais.push(VERBO[ultimo]);
    else campos.push(ultimo);
  }
  const partes: string[] = [];
  if (estruturais.length > 0) partes.push([...new Set(estruturais)].join(", "));
  if (campos.length > 0) {
    partes.push(
      `alterou ${campos.length} campo${campos.length > 1 ? "s" : ""}: ` +
        campos.slice(0, 3).join(", ") +
        (campos.length > 3 ? "…" : ""),
    );
  }
  return partes.join(" · ");
}

function horaCurta(iso: string) {
  const d = new Date(iso);
  if (Number.isNaN(d.getTime())) return "";
  return d.toLocaleString("pt-BR", {
    day: "2-digit",
    month: "2-digit",
    hour: "2-digit",
    minute: "2-digit",
  });
}

export interface PlanHistoryPanelProps {
  revisoes: PlanRevision[];
  carregando?: boolean;
  restaurando?: boolean;
  onRestaurar: (revisionId: string) => void;
  /** Bloqueia restaurar enquanto há alteração não salva, que seria perdida em silêncio. */
  sujo?: boolean;
}

export function PlanHistoryPanel({
  revisoes,
  carregando,
  restaurando,
  onRestaurar,
  sujo,
}: PlanHistoryPanelProps) {
  const [aberto, setAberto] = useState(false);
  const [confirmando, setConfirmando] = useState<string | null>(null);

  return (
    <Sheet open={aberto} onOpenChange={setAberto}>
      <SheetTrigger asChild>
        <Button
          type="button"
          variant="outline"
          size="sm"
          className="h-8 text-xs"
        >
          <History className="mr-1.5 h-3.5 w-3.5" />
          Histórico
          {revisoes.length > 0 && (
            <span className="ml-1.5 text-muted-foreground">
              {revisoes.length}
            </span>
          )}
        </Button>
      </SheetTrigger>
      <SheetContent side="right" className="w-full sm:max-w-md">
        <SheetHeader>
          <SheetTitle>Histórico do rascunho</SheetTitle>
        </SheetHeader>
        <p className="mt-1 text-[11px] text-muted-foreground">
          Restaurar não apaga nada: grava uma versão nova com o conteúdo antigo,
          então restaurar por engano também se desfaz.
        </p>
        <ScrollArea className="mt-3 h-[calc(100vh-9rem)] pr-3">
          {carregando && (
            <p className="text-xs text-muted-foreground">Carregando…</p>
          )}
          {!carregando && revisoes.length === 0 && (
            <p className="text-xs text-muted-foreground">
              Nenhuma gravação ainda.
            </p>
          )}
          <ol className="space-y-2">
            {revisoes.map((r) => {
              const daFerramenta = r.authorKind === "assistant";
              return (
                <li
                  key={r.id}
                  className={cn(
                    "rounded-md border p-2.5",
                    r.isPublication && "border-l-2 border-l-primary",
                  )}
                >
                  <div className="flex items-start justify-between gap-2">
                    <div className="min-w-0">
                      <p className="flex items-center gap-1.5 text-xs font-medium">
                        {r.isPublication ? (
                          <Upload className="h-3 w-3 shrink-0 text-primary" />
                        ) : daFerramenta ? (
                          <Sparkles className="h-3 w-3 shrink-0 text-amber-600" />
                        ) : (
                          <User className="h-3 w-3 shrink-0 text-muted-foreground" />
                        )}
                        <span className="tabular-nums text-muted-foreground">
                          #{r.seq}
                        </span>
                        {MOTIVO[r.reason] ?? r.reason}
                      </p>
                      <p className="mt-0.5 truncate text-[11px] text-muted-foreground">
                        {horaCurta(r.createdAt)}
                        {r.authorName ? ` · ${r.authorName}` : ""}
                        {daFerramenta && r.aiModel ? ` · ${r.aiModel}` : ""}
                      </p>
                      {(r.changedPaths?.length ?? 0) > 0 && (
                        <p className="mt-1 text-[11px] text-muted-foreground">
                          {descreveMudanca(r.changedPaths!)}
                        </p>
                      )}
                      {/*
                        Só aparece na publicação, e é a pergunta que a coluna existe para responder:
                        o que a ferramenta escreveu e ninguém reescreveu antes de o paciente ler.
                        Super-reporta de propósito, e o texto diz isso.
                      */}
                      {r.isPublication &&
                        (r.aiTouchedPaths?.length ?? 0) > 0 && (
                          <p
                            className="mt-1 text-[11px] text-amber-700"
                            title={r
                              .aiTouchedPaths!.map((c) => c.split(":").pop())
                              .join("\n")}
                          >
                            {r.aiTouchedPaths!.length} campo
                            {r.aiTouchedPaths!.length > 1 ? "s" : ""} sem
                            reescrita registrada depois da ferramenta
                          </p>
                        )}
                    </div>
                    <div className="shrink-0">
                      {confirmando === r.id ? (
                        <div className="flex flex-col gap-1">
                          <Button
                            type="button"
                            size="sm"
                            className="h-6 px-2 text-[11px]"
                            disabled={restaurando}
                            onClick={() => {
                              onRestaurar(r.id);
                              setConfirmando(null);
                            }}
                          >
                            Confirmar
                          </Button>
                          <Button
                            type="button"
                            size="sm"
                            variant="ghost"
                            className="h-6 px-2 text-[11px]"
                            onClick={() => setConfirmando(null)}
                          >
                            Cancelar
                          </Button>
                        </div>
                      ) : (
                        <Button
                          type="button"
                          size="sm"
                          variant="ghost"
                          className="h-6 px-2 text-[11px]"
                          disabled={restaurando || sujo}
                          title={
                            sujo
                              ? "Salve ou descarte a alteração aberta antes de restaurar"
                              : undefined
                          }
                          onClick={() => setConfirmando(r.id)}
                        >
                          <RotateCcw className="mr-1 h-3 w-3" />
                          Restaurar
                        </Button>
                      )}
                    </div>
                  </div>
                  {r.isPublication && (
                    <Badge
                      variant="outline"
                      className="mt-1.5 h-5 px-1.5 text-[10px] font-normal"
                    >
                      v{r.planVersion} no portal
                    </Badge>
                  )}
                </li>
              );
            })}
          </ol>
        </ScrollArea>
      </SheetContent>
    </Sheet>
  );
}
