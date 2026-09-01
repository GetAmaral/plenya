'use client';

import { useEffect, useRef, useState } from 'react';
import { Bot, Loader2, Send, ShieldAlert, User } from 'lucide-react';
import type { PlanAssistantTurn, PlanMessage, PlanSuggestion } from '@plenya/types';

import { Badge } from '@/components/ui/badge';
import { Button } from '@/components/ui/button';
import { ScrollArea } from '@/components/ui/scroll-area';
import { Textarea } from '@/components/ui/textarea';
import { formatDate } from '@/lib/format-date';
import { cn } from '@/lib/utils';

/**
 * A conversa que edita o rascunho.
 *
 * Sem streaming, porque não existe streaming em lugar nenhum deste app e um turno leva de dez a
 * vinte segundos: um POST com estado de espera honesto, e a conversa persistida no servidor, para
 * fechar o notebook não perder o turno.
 *
 * O resultado do turno aparece como resumo do que ENTROU e do que ESPERA aceite. As sugestões em si
 * moram no cartão do slide que elas alteram, e não aqui: decidir sobre um número olhando a lista de
 * sugestões, longe do slide, é decidir no escuro.
 */
export function PlanChatPanel({
  mensagens,
  sugestoes,
  onEnviar,
  enviando,
  ultimoTurno,
  desabilitado,
  motivoDesabilitado,
}: {
  mensagens: PlanMessage[];
  sugestoes: PlanSuggestion[];
  onEnviar: (texto: string) => void;
  enviando?: boolean;
  ultimoTurno?: PlanAssistantTurn | null;
  desabilitado?: boolean;
  motivoDesabilitado?: string;
}) {
  const [texto, setTexto] = useState('');
  const fim = useRef<HTMLDivElement>(null);

  useEffect(() => {
    fim.current?.scrollIntoView({ behavior: 'smooth' });
  }, [mensagens.length, enviando]);

  const enviar = () => {
    const t = texto.trim();
    if (!t || enviando || desabilitado) return;
    onEnviar(t);
    setTexto('');
  };

  return (
    <div className="flex h-full flex-col">
      <div className="flex items-baseline justify-between gap-2 pb-2">
        <h2 className="text-sm font-semibold">Conversa</h2>
        {sugestoes.length > 0 && (
          <Badge variant="outline" className="h-5 gap-1 border-amber-500 px-1.5 text-[10px] text-amber-700">
            <ShieldAlert className="h-2.5 w-2.5" />
            {sugestoes.length} esperando aceite
          </Badge>
        )}
      </div>

      <ScrollArea className="min-h-0 flex-1 pr-2">
        <div className="space-y-3">
          {mensagens.length === 0 && (
            <p className="px-1 py-6 text-center text-xs text-muted-foreground">
              Peça uma alteração em português. Reescrita de texto entra direto e fica no histórico;
              qualquer coisa que toque número vira sugestão para você conferir, com a origem do valor
              ao lado.
            </p>
          )}

          {mensagens.map((m) => (
            <div key={m.id} className={cn('flex gap-2', m.role === 'user' && 'flex-row-reverse')}>
              <span
                className={cn(
                  'mt-0.5 flex h-6 w-6 shrink-0 items-center justify-center rounded-full',
                  m.role === 'user' ? 'bg-primary/10' : 'bg-muted',
                )}
              >
                {m.role === 'user' ? <User className="h-3 w-3" /> : <Bot className="h-3 w-3" />}
              </span>
              <div
                className={cn(
                  'min-w-0 max-w-[85%] rounded-lg px-2.5 py-1.5 text-xs leading-relaxed',
                  m.role === 'user' ? 'bg-primary/5' : 'bg-muted/60',
                  m.status === 'failed' && 'border border-destructive/40 text-destructive',
                )}
              >
                <p className="whitespace-pre-wrap">{m.body}</p>
                <p className="mt-0.5 text-[10px] text-muted-foreground">{formatDate(m.createdAt, 'HH:mm')}</p>
              </div>
            </div>
          ))}

          {enviando && (
            <div className="flex gap-2">
              <span className="mt-0.5 flex h-6 w-6 items-center justify-center rounded-full bg-muted">
                <Bot className="h-3 w-3" />
              </span>
              <div className="rounded-lg bg-muted/60 px-2.5 py-1.5 text-xs text-muted-foreground">
                <Loader2 className="mr-1 inline h-3 w-3 animate-spin" />
                lendo o prontuário e montando as alterações
              </div>
            </div>
          )}

          {ultimoTurno && !enviando && (
            <div className="rounded-md border bg-card p-2 text-[11px]">
              <p className="font-medium">Nesta rodada</p>
              <ul className="mt-1 space-y-0.5 text-muted-foreground">
                <li>{ultimoTurno.applied?.length ?? 0} alteração(ões) de texto entraram</li>
                <li>{ultimoTurno.suggestions?.length ?? 0} esperando seu aceite, no cartão do slide</li>
                {(ultimoTurno.rejected?.length ?? 0) > 0 && (
                  <li className="text-destructive">
                    {ultimoTurno.rejected!.length} recusada(s): campo que vem do prontuário
                  </li>
                )}
              </ul>
            </div>
          )}

          <div ref={fim} />
        </div>
      </ScrollArea>

      <div className="space-y-1 pt-2">
        <Textarea
          value={texto}
          onChange={(e) => setTexto(e.target.value)}
          onKeyDown={(e) => {
            if (e.key === 'Enter' && (e.metaKey || e.ctrlKey)) {
              e.preventDefault();
              enviar();
            }
          }}
          rows={3}
          disabled={desabilitado || enviando}
          placeholder={desabilitado ? motivoDesabilitado : 'Encurte o título do slide 3…'}
          className="resize-none text-xs"
        />
        <div className="flex items-center justify-between gap-2">
          <span className="text-[10px] text-muted-foreground">Ctrl+Enter envia</span>
          <Button
            type="button"
            size="sm"
            className="h-7 text-xs"
            onClick={enviar}
            disabled={!texto.trim() || enviando || desabilitado}
          >
            {enviando ? <Loader2 className="mr-1 h-3 w-3 animate-spin" /> : <Send className="mr-1 h-3 w-3" />}
            Enviar
          </Button>
        </div>
      </div>
    </div>
  );
}
