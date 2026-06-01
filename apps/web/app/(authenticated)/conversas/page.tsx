'use client';

import { useEffect, useMemo, useState } from 'react';
import { Inbox, Mail, MessageSquare, Search } from 'lucide-react';

import { cn } from '@/lib/utils';
import { useRequireAuth } from '@/lib/use-auth';
import { usePatientGuard } from '@/lib/use-patient-guard';
import { Input } from '@/components/ui/input';
import { Skeleton } from '@/components/ui/skeleton';
import { Badge } from '@/components/ui/badge';
import { Tabs, TabsList, TabsTrigger } from '@/components/ui/tabs';
import { Switch } from '@/components/ui/switch';
import { Sheet, SheetContent } from '@/components/ui/sheet';
import { PageHeader } from '@/components/layout/page-header';
import {
  type ConversationItem,
  type ConversationListFilters,
  type ConversationOwnerType,
  useConversations,
} from '@/lib/api/conversations-api';
import { ConversationListRow } from '@/components/conversations/conversation-list-item';
import { ConversationViewer } from '@/components/conversations/conversation-viewer';
import { NewEmailDialog } from '@/components/conversations/new-email-dialog';

type AssignedFilter = 'all' | 'mine';
type ChannelFilter = 'all' | 'email' | 'whatsapp';

type Selected = { type: ConversationOwnerType; id: string } | null;

/**
 * Debounce simples para busca — evita spam de requests enquanto o usuário digita.
 */
function useDebouncedValue<T>(value: T, ms = 300): T {
  const [debounced, setDebounced] = useState(value);
  useEffect(() => {
    const t = setTimeout(() => setDebounced(value), ms);
    return () => clearTimeout(t);
  }, [value, ms]);
  return debounced;
}

export default function ConversasPage() {
  useRequireAuth();
  usePatientGuard();

  const [assigned, setAssigned] = useState<AssignedFilter>('all');
  const [channel, setChannel] = useState<ChannelFilter>('all');
  const [unreadOnly, setUnreadOnly] = useState(false);
  const [search, setSearch] = useState('');
  const debouncedSearch = useDebouncedValue(search, 300);

  const [selected, setSelected] = useState<Selected>(null);
  const [mobileViewerOpen, setMobileViewerOpen] = useState(false);
  const [newEmailOpen, setNewEmailOpen] = useState(false);

  const filters: ConversationListFilters = useMemo(
    () => ({
      assignedToMe: assigned === 'mine',
      unreadOnly,
      channel: channel === 'all' ? undefined : channel,
      search: debouncedSearch.trim() || undefined,
      limit: 50,
    }),
    [assigned, unreadOnly, channel, debouncedSearch]
  );

  const { data, isLoading, isError } = useConversations(filters);

  const items = data?.items ?? [];
  // SLA: conversas aguardando resposta (última msg inbound) sobem ao topo da
  // página carregada, da mais antiga aguardando pra mais nova; o resto segue por
  // recência. Ordenação dentro da página carregada (cross-page server-side = follow-up).
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
  const totalUnread = useMemo(
    () => items.reduce((sum, it) => sum + it.unreadCount, 0),
    [items]
  );

  const selectedItem: ConversationItem | undefined = useMemo(() => {
    if (!selected) return undefined;
    return items.find((i) => i.ownerType === selected.type && i.ownerId === selected.id);
  }, [selected, items]);

  // Auto-seleciona primeira conversa em desktop, mas NUNCA em mobile (UX: fica na lista).
  useEffect(() => {
    if (selected) return;
    if (typeof window !== 'undefined' && window.innerWidth < 768) return;
    if (sortedItems.length > 0) {
      setSelected({ type: sortedItems[0].ownerType, id: sortedItems[0].ownerId });
    }
  }, [sortedItems, selected]);

  // Se a conversa selecionada sumir da lista (ex: filtro mudou), limpa a seleção
  useEffect(() => {
    if (!selected) return;
    if (data && !selectedItem) {
      setSelected(null);
      setMobileViewerOpen(false);
    }
  }, [selected, selectedItem, data]);

  const handleSelect = (item: ConversationItem) => {
    setSelected({ type: item.ownerType, id: item.ownerId });
    setMobileViewerOpen(true);
  };

  return (
    <div className="flex h-[calc(100vh-4rem)] flex-col gap-3 p-3 sm:p-4">
      <PageHeader
        title="Conversas"
        description="Caixa unificada de email e WhatsApp — leads e pacientes em um só lugar."
        actions={[
          {
            label: 'Novo email',
            icon: <Mail className="h-4 w-4" />,
            onClick: () => setNewEmailOpen(true),
          },
        ]}
      />

      <NewEmailDialog
        open={newEmailOpen}
        onOpenChange={setNewEmailOpen}
        onSent={({ ownerType, ownerId }) => {
          // Seleciona a conversa criada/reusada — espera próximo refetch da lista
          // (mutation já invalida `conversations` keys em onSuccess).
          setSelected({ type: ownerType, id: ownerId });
          if (typeof window !== 'undefined' && window.innerWidth < 768) {
            setMobileViewerOpen(true);
          }
        }}
      />

      {/* Filtros */}
      <div className="flex flex-wrap items-center gap-2">
        <Tabs value={assigned} onValueChange={(v) => setAssigned(v as AssignedFilter)}>
          <TabsList>
            <TabsTrigger value="all">Todas</TabsTrigger>
            <TabsTrigger value="mine">Atribuídas a mim</TabsTrigger>
          </TabsList>
        </Tabs>

        <Tabs value={channel} onValueChange={(v) => setChannel(v as ChannelFilter)}>
          <TabsList>
            <TabsTrigger value="all">Todos</TabsTrigger>
            <TabsTrigger value="email">Email</TabsTrigger>
            <TabsTrigger value="whatsapp">WhatsApp</TabsTrigger>
          </TabsList>
        </Tabs>

        <label className="inline-flex shrink-0 items-center gap-2 rounded-md border border-border px-3 py-1.5 text-sm">
          <Switch
            checked={unreadOnly}
            onCheckedChange={(v) => setUnreadOnly(!!v)}
            aria-label="Mostrar apenas não lidas"
          />
          <span className="text-xs text-muted-foreground">Não lidas</span>
        </label>

        {totalUnread > 0 && (
          <Badge variant="secondary" className="ml-auto bg-rose-100 text-rose-900 border-rose-200">
            {totalUnread} não {totalUnread === 1 ? 'lida' : 'lidas'}
          </Badge>
        )}
      </div>

      <div className="relative">
        <Search className="absolute left-3 top-2.5 h-4 w-4 text-muted-foreground" aria-hidden />
        <Input
          placeholder="Buscar por nome, email ou telefone…"
          value={search}
          onChange={(e) => setSearch(e.target.value)}
          className="pl-9"
          aria-label="Buscar conversa"
        />
      </div>

      {/* Layout: split desktop / só lista mobile */}
      <div className="flex min-h-0 flex-1 overflow-hidden rounded-lg border border-border bg-card">
        {/* Lista esquerda */}
        <div
          className={cn(
            'flex min-w-0 min-h-0 flex-1 flex-col border-r border-border md:max-w-[380px] md:flex-none md:basis-[380px]'
          )}
        >
          <div className="flex-1 overflow-y-auto" role="list" aria-label="Lista de conversas">
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
                Falha ao carregar conversas. Atualize a página em alguns segundos.
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
        </div>

        {/* Viewer direita (desktop) */}
        <div className="hidden min-w-0 flex-1 md:flex md:flex-col">
          {selectedItem ? (
            <ConversationViewer item={selectedItem} />
          ) : (
            <div className="flex h-full flex-col items-center justify-center gap-3 p-6 text-center text-sm text-muted-foreground">
              <MessageSquare className="h-12 w-12 text-muted-foreground/50" aria-hidden />
              <p>Selecione uma conversa para ler e responder.</p>
            </div>
          )}
        </div>
      </div>

      {/* Drawer viewer mobile */}
      <Sheet
        open={mobileViewerOpen && !!selectedItem}
        onOpenChange={(open) => {
          if (!open) setMobileViewerOpen(false);
        }}
      >
        <SheetContent
          side="right"
          className="w-full max-w-full p-0 sm:max-w-full md:hidden"
        >
          {selectedItem && (
            <div className="flex h-full min-h-0 flex-col">
              <ConversationViewer
                item={selectedItem}
                onBack={() => setMobileViewerOpen(false)}
              />
            </div>
          )}
        </SheetContent>
      </Sheet>
    </div>
  );
}

function EmptyList({ unreadOnly, hasSearch }: { unreadOnly: boolean; hasSearch: boolean }) {
  let msg = 'Nenhuma conversa por enquanto. Quando chegar email ou WhatsApp, aparece aqui.';
  if (hasSearch) msg = 'Nenhuma conversa bate com a busca.';
  else if (unreadOnly) msg = 'Sem mensagens não lidas. Tudo em dia.';
  return (
    <div className="flex h-full flex-col items-center justify-center gap-3 p-6 text-center text-sm text-muted-foreground">
      <Inbox className="h-10 w-10 text-muted-foreground/50" aria-hidden />
      <p>{msg}</p>
    </div>
  );
}
