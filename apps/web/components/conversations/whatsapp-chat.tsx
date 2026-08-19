'use client';

import { useEffect, useMemo, useState } from 'react';
import { Inbox, MessageSquare, Search } from 'lucide-react';

import { cn } from '@/lib/utils';
import { Input } from '@/components/ui/input';
import { Skeleton } from '@/components/ui/skeleton';
import { Switch } from '@/components/ui/switch';
import { Sheet, SheetContent, SheetTitle } from '@/components/ui/sheet';
import {
  type ConversationItem,
  type ConversationListFilters,
  type ConversationOwnerType,
  useConversations,
} from '@/lib/api/conversations-api';
import { ConversationListRow } from './conversation-list-item';
import { ConversationViewer } from './conversation-viewer';

export type ChatSelection = { type: ConversationOwnerType; id: string } | null;

type Props = {
  /** 'page' = split-pane cheio (rota dedicada); 'dock' = coluna única (Sheet lateral). */
  variant?: 'page' | 'dock';
  selected: ChatSelection;
  onSelect: (sel: ChatSelection) => void;
};

function useDebouncedValue<T>(value: T, ms = 300): T {
  const [debounced, setDebounced] = useState(value);
  useEffect(() => {
    const t = setTimeout(() => setDebounced(value), ms);
    return () => clearTimeout(t);
  }, [value, ms]);
  return debounced;
}

/**
 * Superfície de chat WhatsApp reutilizável — lista + thread em balões, canal travado em
 * 'whatsapp'. Usada pela página `/conversas/whatsapp` (variant page) e pelo dock global
 * (variant dock). Estado de seleção é controlado pelo pai (no dock vive no Zustand, pra
 * sobreviver à navegação).
 */
export function WhatsAppChat({ variant = 'page', selected, onSelect }: Props) {
  const isDock = variant === 'dock';

  const [assigned, setAssigned] = useState<'all' | 'mine'>('all');
  const [unreadOnly, setUnreadOnly] = useState(false);
  const [search, setSearch] = useState('');
  const debouncedSearch = useDebouncedValue(search, 300);
  const [mobileViewerOpen, setMobileViewerOpen] = useState(false);

  // O viewer mobile é um Sheet cujo overlay (bg-black/80, fixed inset-0) NÃO tem
  // hiding responsivo — só o conteúdo é md:hidden. Se abrir no desktop, o backdrop
  // cobre a tela inteira sem conteúdo visível = tela preta. Por isso o Sheet só pode
  // abrir abaixo de md (768px). Reativo a resize pra não travar preto ao redimensionar.
  const [isMobile, setIsMobile] = useState(false);
  useEffect(() => {
    const mql = window.matchMedia('(max-width: 767px)');
    const update = () => setIsMobile(mql.matches);
    update();
    mql.addEventListener('change', update);
    return () => mql.removeEventListener('change', update);
  }, []);

  const filters: ConversationListFilters = useMemo(
    () => ({
      channel: 'whatsapp',
      assignedToMe: assigned === 'mine',
      unreadOnly,
      search: debouncedSearch.trim() || undefined,
      limit: 50,
    }),
    [assigned, unreadOnly, debouncedSearch]
  );

  const { data, isLoading, isError } = useConversations(filters);
  const items = data?.items ?? [];

  // Uma ordem só: mensagem mais recente no topo, seja ela minha ou do contato — igual a
  // qualquer app de mensagem. Havia aqui uma ordenação por SLA (não respondidas primeiro,
  // MAIS ANTIGA no topo, respondidas empurradas pra baixo) que, misturada com a ordem do
  // servidor, fazia a lista parecer aleatória.
  const sortedItems = useMemo(
    () =>
      [...items].sort(
        (a, b) => new Date(b.lastAt).getTime() - new Date(a.lastAt).getTime()
      ),
    [items]
  );

  const selectedItem = useMemo(() => {
    if (!selected) return undefined;
    return items.find((i) => i.ownerType === selected.type && i.ownerId === selected.id);
  }, [selected, items]);

  // Auto-seleciona a primeira conversa só na página em desktop (dock começa na lista).
  useEffect(() => {
    if (isDock || selected) return;
    if (typeof window !== 'undefined' && window.innerWidth < 768) return;
    if (sortedItems.length > 0) {
      onSelect({ type: sortedItems[0].ownerType, id: sortedItems[0].ownerId });
    }
  }, [isDock, sortedItems, selected, onSelect]);

  // Limpa seleção que sumiu da lista (filtro mudou).
  useEffect(() => {
    if (!selected) return;
    if (data && !selectedItem) {
      onSelect(null);
      setMobileViewerOpen(false);
    }
  }, [selected, selectedItem, data, onSelect]);

  const handleSelect = (item: ConversationItem) => {
    onSelect({ type: item.ownerType, id: item.ownerId });
    setMobileViewerOpen(true);
  };

  const list = (
    <div className="flex-1 overflow-y-auto" role="list" aria-label="Conversas de WhatsApp">
      {isLoading ? (
        <div className="space-y-0">
          {Array.from({ length: 5 }).map((_, i) => (
            <div key={i} className="border-b border-border p-3">
              <div className="flex gap-3">
                <Skeleton className="h-10 w-10 rounded-full" />
                <div className="flex-1 space-y-2">
                  <Skeleton className="h-4 w-2/3" />
                  <Skeleton className="h-3 w-full" />
                </div>
              </div>
            </div>
          ))}
        </div>
      ) : isError ? (
        <div className="flex h-full items-center justify-center p-6 text-center text-sm text-rose-700">
          Falha ao carregar conversas. Atualize em alguns segundos.
        </div>
      ) : items.length === 0 ? (
        <EmptyList unreadOnly={unreadOnly} hasSearch={!!debouncedSearch.trim()} />
      ) : (
        sortedItems.map((item) => (
          <div role="listitem" key={`${item.ownerType}:${item.ownerId}`}>
            <ConversationListRow
              item={item}
              active={
                !!selected &&
                selected.type === item.ownerType &&
                selected.id === item.ownerId
              }
              onSelect={() => handleSelect(item)}
            />
          </div>
        ))
      )}
    </div>
  );

  const filterBar = (
    <div className="shrink-0 space-y-2 border-b border-border p-2">
      <div className="relative">
        <Search className="absolute left-3 top-2.5 h-4 w-4 text-muted-foreground" aria-hidden />
        <Input
          placeholder="Buscar por nome ou telefone…"
          value={search}
          onChange={(e) => setSearch(e.target.value)}
          className="pl-9"
          aria-label="Buscar conversa"
        />
      </div>
      <div className="flex items-center justify-between gap-2 px-1">
        <button
          type="button"
          onClick={() => setAssigned((a) => (a === 'mine' ? 'all' : 'mine'))}
          className={cn(
            'rounded-full border px-2.5 py-1 text-xs font-medium transition-colors',
            assigned === 'mine'
              ? 'border-emerald-300 bg-emerald-100 text-emerald-900'
              : 'border-border bg-background text-muted-foreground hover:bg-muted'
          )}
        >
          Atribuídas a mim
        </button>
        <label className="inline-flex shrink-0 items-center gap-2 text-xs text-muted-foreground">
          <Switch
            checked={unreadOnly}
            onCheckedChange={(v) => setUnreadOnly(!!v)}
            aria-label="Mostrar apenas não lidas"
          />
          Não lidas
        </label>
      </div>
    </div>
  );

  // ===== DOCK: coluna única (lista ↔ thread) =====
  if (isDock) {
    return (
      <div className="flex h-full min-h-0 flex-col">
        {selectedItem ? (
          <ConversationViewer
            item={selectedItem}
            channel="whatsapp"
            menuControls
            compact
            onBack={() => onSelect(null)}
          />
        ) : (
          <>
            {filterBar}
            {list}
          </>
        )}
      </div>
    );
  }

  // ===== PAGE: split-pane + drawer mobile =====
  return (
    <>
      <div className="flex min-h-0 flex-1 overflow-hidden rounded-lg border border-border bg-card">
        <div className="flex min-w-0 min-h-0 flex-1 flex-col border-r border-border md:max-w-[380px] md:flex-none md:basis-[380px]">
          {filterBar}
          {list}
        </div>
        <div className="hidden min-w-0 flex-1 md:flex md:flex-col">
          {selectedItem ? (
            <ConversationViewer item={selectedItem} channel="whatsapp" menuControls />
          ) : (
            <div className="flex h-full flex-col items-center justify-center gap-3 p-6 text-center text-sm text-muted-foreground">
              <MessageSquare className="h-12 w-12 text-muted-foreground/50" aria-hidden />
              <p>Selecione uma conversa para ler e responder.</p>
            </div>
          )}
        </div>
      </div>

      <Sheet
        open={mobileViewerOpen && !!selectedItem && isMobile}
        onOpenChange={(open) => {
          if (!open) setMobileViewerOpen(false);
        }}
      >
        <SheetContent side="right" className="w-full max-w-full p-0 sm:max-w-full md:hidden">
          <SheetTitle className="sr-only">Conversa de WhatsApp</SheetTitle>
          {selectedItem && (
            <div className="flex h-full min-h-0 flex-col">
              <ConversationViewer
                item={selectedItem}
                channel="whatsapp"
                menuControls
                onBack={() => setMobileViewerOpen(false)}
              />
            </div>
          )}
        </SheetContent>
      </Sheet>
    </>
  );
}

function EmptyList({ unreadOnly, hasSearch }: { unreadOnly: boolean; hasSearch: boolean }) {
  let msg = 'Nenhuma conversa de WhatsApp por enquanto.';
  if (hasSearch) msg = 'Nenhuma conversa bate com a busca.';
  else if (unreadOnly) msg = 'Sem mensagens não lidas. Tudo em dia.';
  return (
    <div className="flex h-full flex-col items-center justify-center gap-3 p-6 text-center text-sm text-muted-foreground">
      <Inbox className="h-10 w-10 text-muted-foreground/50" aria-hidden />
      <p>{msg}</p>
    </div>
  );
}
