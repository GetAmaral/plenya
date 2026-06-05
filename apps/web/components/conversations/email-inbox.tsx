'use client';

import { useEffect, useMemo, useState } from 'react';
import { Inbox, Mail, Search } from 'lucide-react';

import { Input } from '@/components/ui/input';
import { Skeleton } from '@/components/ui/skeleton';
import { Switch } from '@/components/ui/switch';
import { cn } from '@/lib/utils';
import { Sheet, SheetContent, SheetTitle } from '@/components/ui/sheet';
import {
  type ConversationItem,
  type ConversationListFilters,
  type ConversationOwnerType,
  useConversations,
} from '@/lib/api/conversations-api';
import { ConversationListRow } from './conversation-list-item';
import { ConversationViewer } from './conversation-viewer';

export type EmailSelection = { type: ConversationOwnerType; id: string } | null;

type Props = {
  selected: EmailSelection;
  onSelect: (sel: EmailSelection) => void;
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
 * Caixa de e-mail estilo webmail — lista + thread com assunto/anexos, canal travado em
 * 'email'. Sem automação/recepção/janela 24h (isso é do WhatsApp). Seleção controlada
 * pelo pai (pra integrar com o "Novo email").
 */
export function EmailInbox({ selected, onSelect }: Props) {
  const [assigned, setAssigned] = useState<'all' | 'mine'>('all');
  const [unreadOnly, setUnreadOnly] = useState(false);
  const [search, setSearch] = useState('');
  const debouncedSearch = useDebouncedValue(search, 300);
  const [mobileViewerOpen, setMobileViewerOpen] = useState(false);

  const filters: ConversationListFilters = useMemo(
    () => ({
      channel: 'email',
      assignedToMe: assigned === 'mine',
      unreadOnly,
      search: debouncedSearch.trim() || undefined,
      limit: 50,
    }),
    [assigned, unreadOnly, debouncedSearch]
  );

  const { data, isLoading, isError } = useConversations(filters);
  const items = data?.items ?? [];

  const sortedItems = useMemo(() => {
    const waitMs = (it: ConversationItem) =>
      new Date(it.lastInboundAt ?? it.lastAt).getTime();
    return [...items].sort((a, b) => {
      const aw = a.lastDirection === 'in' ? 0 : 1;
      const bw = b.lastDirection === 'in' ? 0 : 1;
      if (aw !== bw) return aw - bw;
      if (aw === 0) return waitMs(a) - waitMs(b);
      return new Date(b.lastAt).getTime() - new Date(a.lastAt).getTime();
    });
  }, [items]);

  const selectedItem = useMemo(() => {
    if (!selected) return undefined;
    return items.find((i) => i.ownerType === selected.type && i.ownerId === selected.id);
  }, [selected, items]);

  useEffect(() => {
    if (selected) return;
    if (typeof window !== 'undefined' && window.innerWidth < 768) return;
    if (sortedItems.length > 0) {
      onSelect({ type: sortedItems[0].ownerType, id: sortedItems[0].ownerId });
    }
  }, [sortedItems, selected, onSelect]);

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

  const filterBar = (
    <div className="shrink-0 space-y-2 border-b border-border p-2">
      <div className="relative">
        <Search className="absolute left-3 top-2.5 h-4 w-4 text-muted-foreground" aria-hidden />
        <Input
          placeholder="Buscar por nome ou email…"
          value={search}
          onChange={(e) => setSearch(e.target.value)}
          className="pl-9"
          aria-label="Buscar e-mail"
        />
      </div>
      <div className="flex items-center justify-between gap-2 px-1">
        <button
          type="button"
          onClick={() => setAssigned((a) => (a === 'mine' ? 'all' : 'mine'))}
          className={cn(
            'rounded-full border px-2.5 py-1 text-xs font-medium transition-colors',
            assigned === 'mine'
              ? 'border-sky-300 bg-sky-100 text-sky-900'
              : 'border-border bg-background text-muted-foreground hover:bg-muted'
          )}
        >
          Atribuídos a mim
        </button>
        <label className="inline-flex shrink-0 items-center gap-2 text-xs text-muted-foreground">
          <Switch
            checked={unreadOnly}
            onCheckedChange={(v) => setUnreadOnly(!!v)}
            aria-label="Mostrar apenas não lidos"
          />
          Não lidos
        </label>
      </div>
    </div>
  );

  const list = (
    <div className="flex-1 overflow-y-auto" role="list" aria-label="E-mails">
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
          Falha ao carregar e-mails. Atualize em alguns segundos.
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

  return (
    <>
      <div className="flex min-h-0 flex-1 overflow-hidden rounded-lg border border-border bg-card">
        <div className="flex min-w-0 min-h-0 flex-1 flex-col border-r border-border md:max-w-[380px] md:flex-none md:basis-[380px]">
          {filterBar}
          {list}
        </div>
        <div className="hidden min-w-0 flex-1 md:flex md:flex-col">
          {selectedItem ? (
            <ConversationViewer item={selectedItem} channel="email" menuControls />
          ) : (
            <div className="flex h-full flex-col items-center justify-center gap-3 p-6 text-center text-sm text-muted-foreground">
              <Mail className="h-12 w-12 text-muted-foreground/50" aria-hidden />
              <p>Selecione um e-mail para ler e responder.</p>
            </div>
          )}
        </div>
      </div>

      <Sheet
        open={mobileViewerOpen && !!selectedItem}
        onOpenChange={(open) => {
          if (!open) setMobileViewerOpen(false);
        }}
      >
        <SheetContent side="right" className="w-full max-w-full p-0 sm:max-w-full md:hidden">
          <SheetTitle className="sr-only">E-mail</SheetTitle>
          {selectedItem && (
            <div className="flex h-full min-h-0 flex-col">
              <ConversationViewer
                item={selectedItem}
                channel="email"
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
  let msg = 'Nenhum e-mail por enquanto.';
  if (hasSearch) msg = 'Nenhum e-mail bate com a busca.';
  else if (unreadOnly) msg = 'Sem e-mails não lidos. Tudo em dia.';
  return (
    <div className="flex h-full flex-col items-center justify-center gap-3 p-6 text-center text-sm text-muted-foreground">
      <Inbox className="h-10 w-10 text-muted-foreground/50" aria-hidden />
      <p>{msg}</p>
    </div>
  );
}
