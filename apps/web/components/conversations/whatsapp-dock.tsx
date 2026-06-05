'use client';

import { useEffect } from 'react';
import { MessageSquare, X } from 'lucide-react';

import { cn } from '@/lib/utils';
import { useAuth } from '@/lib/use-auth';
import { isGranted } from '@/lib/auth-store';
import { useConversationsUnreadCount } from '@/lib/api/conversations-api';
import { useWhatsAppDock } from '@/lib/whatsapp-dock-store';
import { WhatsAppChat } from './whatsapp-chat';

/**
 * Dock global de WhatsApp — painel deslizante NÃO-modal, montado no layout autenticado
 * (persiste entre navegações). Permite à secretária responder WhatsApp de qualquer tela
 * sem trocar de rota: abre por botão flutuante ou atalho (Ctrl/Cmd+J), responde e fecha,
 * voltando exatamente pro que estava fazendo. Como é não-modal, a tela atrás continua
 * clicável (dá pra navegar com o dock aberto, mantendo o rascunho).
 */
export function WhatsAppDock() {
  const { user } = useAuth();
  const { isOpen, selected, toggle, close, setSelected } = useWhatsAppDock();
  const unread = useConversationsUnreadCount('whatsapp');

  // Atalho global: Ctrl/Cmd+J abre/fecha; Esc fecha quando aberto.
  useEffect(() => {
    const onKey = (e: KeyboardEvent) => {
      if ((e.metaKey || e.ctrlKey) && (e.key === 'j' || e.key === 'J')) {
        e.preventDefault();
        toggle();
      } else if (e.key === 'Escape' && useWhatsAppDock.getState().isOpen) {
        close();
      }
    };
    window.addEventListener('keydown', onKey);
    return () => window.removeEventListener('keydown', onKey);
  }, [toggle, close]);

  const allowed =
    isGranted(user, 'admin') || isGranted(user, 'secretary') || isGranted(user, 'manager');
  if (!allowed) return null;

  const count = unread.data ?? 0;

  return (
    <>
      {/* Botão flutuante (some quando o dock está aberto) */}
      {!isOpen && (
        <button
          type="button"
          onClick={toggle}
          aria-label="Abrir WhatsApp (Ctrl+J)"
          title="WhatsApp (Ctrl+J)"
          className="fixed bottom-5 right-5 z-40 flex h-14 w-14 items-center justify-center rounded-full bg-emerald-600 text-white shadow-lg transition-transform hover:scale-105 hover:bg-emerald-700 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-emerald-400 print:hidden"
        >
          <MessageSquare className="h-6 w-6" />
          {count > 0 && (
            <span className="absolute -right-1 -top-1 flex h-5 min-w-5 items-center justify-center rounded-full bg-rose-500 px-1 text-[11px] font-semibold text-white ring-2 ring-background">
              {count > 99 ? '99+' : count}
            </span>
          )}
        </button>
      )}

      {/* Painel não-modal ancorado à direita */}
      <div
        role="dialog"
        aria-label="WhatsApp"
        aria-hidden={!isOpen}
        className={cn(
          'fixed right-0 top-0 z-40 flex h-screen w-full max-w-md flex-col border-l border-border bg-background shadow-2xl transition-transform duration-300 ease-in-out print:hidden',
          isOpen ? 'translate-x-0' : 'pointer-events-none translate-x-full'
        )}
      >
        <header className="flex shrink-0 items-center justify-between border-b border-border px-3 py-2">
          <div className="flex items-center gap-2 text-sm font-semibold">
            <MessageSquare className="h-4 w-4 text-emerald-600" /> WhatsApp
          </div>
          <button
            type="button"
            onClick={close}
            aria-label="Fechar WhatsApp"
            className="flex h-8 w-8 items-center justify-center rounded-md hover:bg-muted"
          >
            <X className="h-4 w-4" />
          </button>
        </header>
        <div className="min-h-0 flex-1">
          {isOpen && (
            <WhatsAppChat variant="dock" selected={selected} onSelect={setSelected} />
          )}
        </div>
      </div>
    </>
  );
}
