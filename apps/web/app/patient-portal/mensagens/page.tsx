"use client";

/**
 * /mensagens — caixa única do paciente com a equipe Plenya.
 * Sem múltiplos canais visíveis pro paciente: ele fala "com a clínica" e ponto.
 * Polling 30s + marca como lido ao montar.
 */
import { useEffect, useRef, useState } from "react";
import { format, formatRelative } from "date-fns";
import { ptBR } from "date-fns/locale";
import { Loader2, Send, MessageSquare } from "lucide-react";
import { toast } from "sonner";

import { useRequirePatientAuth } from "@/lib/use-patient-auth";
import {
  useMyMessages,
  useSendMessage,
  useMarkMessagesRead,
  type PatientMessage,
} from "@/lib/api/patient-portal-api";
import { Card, CardContent } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Textarea } from "@/components/ui/textarea";
import { cn } from "@/lib/utils";

export default function MyMessagesPage() {
  const { ready } = useRequirePatientAuth();
  const { data, isLoading } = useMyMessages();
  const send = useSendMessage();
  const markRead = useMarkMessagesRead();
  const [draft, setDraft] = useState("");
  const scrollRef = useRef<HTMLDivElement>(null);
  const markedRef = useRef(false);

  // Marca lido uma vez ao montar (depois de carregar)
  useEffect(() => {
    if (!markedRef.current && data) {
      markedRef.current = true;
      markRead.mutate();
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [data]);

  // Auto-scroll para o final quando lista atualiza
  useEffect(() => {
    if (scrollRef.current) {
      scrollRef.current.scrollTop = scrollRef.current.scrollHeight;
    }
  }, [data?.length]);

  const handleSend = async () => {
    if (!draft.trim()) return;
    try {
      await send.mutateAsync(draft.trim());
      setDraft("");
    } catch (e: any) {
      toast.error(e?.message ?? "Falha ao enviar");
    }
  };

  const handleKeyDown = (e: React.KeyboardEvent<HTMLTextAreaElement>) => {
    if ((e.metaKey || e.ctrlKey) && e.key === "Enter") {
      e.preventDefault();
      handleSend();
    }
  };

  if (!ready) {
    return (
      <div className="flex h-[60vh] items-center justify-center">
        <Loader2 className="h-6 w-6 animate-spin text-muted-foreground" />
      </div>
    );
  }

  // Mensagens vêm DESC do backend, invertemos pra mostrar antiga em cima
  const messages = (data ?? []).slice().reverse();

  return (
    <div className="mx-auto max-w-2xl space-y-4">
      <header className="space-y-2">
        <p className="text-sm uppercase tracking-wider text-muted-foreground">
          Mensagens
        </p>
        <h1 className="text-3xl font-light">Sua conversa com a Plenya</h1>
        <p className="text-sm text-muted-foreground">
          A equipe responde em horário comercial. Para urgências, ligue.
        </p>
      </header>

      <Card className="flex h-[65vh] flex-col">
        <div ref={scrollRef} className="flex-1 space-y-3 overflow-y-auto p-4">
          {isLoading ? (
            <div className="flex h-full items-center justify-center">
              <Loader2 className="h-5 w-5 animate-spin text-muted-foreground" />
            </div>
          ) : messages.length === 0 ? (
            <div className="flex h-full flex-col items-center justify-center gap-3 text-center text-sm text-muted-foreground">
              <MessageSquare className="h-8 w-8" />
              <p>Nenhuma mensagem ainda. Escreva pra equipe abaixo.</p>
            </div>
          ) : (
            messages.map((m) => <MessageBubble key={m.id} m={m} />)
          )}
        </div>
        <div className="border-t p-3">
          <div className="flex gap-2">
            <Textarea
              rows={2}
              value={draft}
              onChange={(e) => setDraft(e.target.value)}
              onKeyDown={handleKeyDown}
              placeholder="Escreva uma mensagem… (Ctrl/⌘+Enter envia)"
              className="resize-none"
            />
            <Button onClick={handleSend} disabled={!draft.trim() || send.isPending}>
              {send.isPending ? <Loader2 className="h-4 w-4 animate-spin" /> : <Send className="h-4 w-4" />}
            </Button>
          </div>
        </div>
      </Card>
    </div>
  );
}

function MessageBubble({ m }: { m: PatientMessage }) {
  const fromMe = m.from === "patient";
  return (
    <div className={cn("flex", fromMe ? "justify-end" : "justify-start")}>
      <div
        className={cn(
          "max-w-[80%] rounded-2xl px-4 py-2.5 text-sm",
          fromMe
            ? "bg-primary text-primary-foreground rounded-br-sm"
            : "bg-muted rounded-bl-sm",
        )}
      >
        {!fromMe && m.authorName && (
          <p className="mb-0.5 text-xs font-medium opacity-70">{m.authorName}</p>
        )}
        <p className="whitespace-pre-wrap">{m.content}</p>
        <p
          className={cn(
            "mt-1 text-[10px] opacity-60",
            fromMe ? "text-right" : "text-left",
          )}
          title={format(new Date(m.createdAt), "dd/MM/yyyy HH:mm")}
        >
          {formatRelative(new Date(m.createdAt), new Date(), { locale: ptBR })}
        </p>
      </div>
    </div>
  );
}
