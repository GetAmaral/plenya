'use client';

import { useMemo, useState } from 'react';
import { useRouter } from 'next/navigation';
import { format } from 'date-fns';
import { ptBR } from 'date-fns/locale';
import { ArrowDown, ArrowUp, ChevronLeft, ChevronRight, Search } from 'lucide-react';
import type { DateRange } from 'react-day-picker';

import { useQueryClient } from '@tanstack/react-query';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Input } from '@/components/ui/input';
import { Button } from '@/components/ui/button';
import { Badge } from '@/components/ui/badge';
import { Skeleton } from '@/components/ui/skeleton';
import { PageHeader } from '@/components/layout/page-header';
import { FiltersMobileSheet } from '@/components/filters/filters-mobile-sheet';
import { LeadsFilterBar } from '@/components/leads/leads-filter-bar';
import { useRequireAuth, useAuth } from '@/lib/use-auth';
import { usePatientGuard } from '@/lib/use-patient-guard';
import { useDebouncedValue } from '@/lib/use-debounced-value';
import { parseList, serializeList, useUrlFilters } from '@/lib/use-url-filters';
import { apiClient } from '@/lib/api-client';
import {
  useLeads,
  useStaffUsers,
  leadKeys,
  SOURCE_LABELS,
  STATUS_LABELS,
  STATUS_COLORS,
  type Lead,
  type LeadFilter,
  type LeadSource,
  type LeadStatus,
} from '@/lib/api/leads-api';

const ALL_STATUSES: ReadonlyArray<LeadStatus> = [
  'new',
  'contacted',
  'qualified',
  'converted',
  'lost',
  'unsubscribed',
];

const ALL_SOURCES: ReadonlyArray<LeadSource> = [
  'light_claim',
  'contact_form',
  'whatsapp_inbound',
  'email_inbound',
  'newsletter',
  'manual',
];

type SortKey = 'createdAt' | 'name' | 'status';
type SortDir = 'asc' | 'desc';

export default function LeadsPage() {
  useRequireAuth();
  usePatientGuard();
  const router = useRouter();
  const qc = useQueryClient();
  const { user } = useAuth();

  const { params, setParam, setParams, clearAll } = useUrlFilters();

  // ---- estado UI vinda da URL ------------------------------------------------
  const search = params.get('q') ?? '';
  const statuses = parseList(params.get('status')).filter((s): s is LeadStatus =>
    ALL_STATUSES.includes(s as LeadStatus),
  );
  const sources = parseList(params.get('source')).filter((s): s is LeadSource =>
    ALL_SOURCES.includes(s as LeadSource),
  );
  const assignee = params.get('assignee'); // "me" | "unassigned" | uuid | null
  const fromIso = params.get('from');
  const toIso = params.get('to');
  const dateRange: DateRange | undefined = useMemo(() => {
    if (!fromIso && !toIso) return undefined;
    return {
      from: fromIso ? new Date(fromIso) : undefined,
      to: toIso ? new Date(toIso) : undefined,
    };
  }, [fromIso, toIso]);
  const sortKey = (params.get('sort') as SortKey | null) ?? 'createdAt';
  const sortDir = (params.get('dir') as SortDir | null) ?? 'desc';
  const page = Math.max(0, parseInt(params.get('page') ?? '0', 10));

  const debouncedSearch = useDebouncedValue(search, 300);
  const pageSize = 25;
  const [actionPending, setActionPending] = useState<string | null>(null);

  // ---- staff e filtros derivados --------------------------------------------
  const { data: staffUsers } = useStaffUsers();

  const assignedToUserId =
    assignee === 'me'
      ? user?.id
      : assignee && assignee !== 'unassigned'
        ? assignee
        : undefined;
  const wantsUnassigned = assignee === 'unassigned';

  // Backend só aceita 1 status/source. Mando filtro só quando há exatamente 1;
  // quando há 2+ selecionados, faço filter client-side na página atual.
  const filter: LeadFilter = {
    search: debouncedSearch || undefined,
    source: sources.length === 1 ? sources[0] : undefined,
    status: statuses.length === 1 ? statuses[0] : undefined,
    assignedToUserId,
  };

  const { data, isLoading } = useLeads(filter, page, pageSize);

  // Contador "meus leads" — só dispara quando user.id resolveu
  const myLeadsQuery = useLeads(
    { assignedToUserId: user?.id },
    0,
    1,
    { enabled: !!user?.id },
  );
  const myLeadsCount = myLeadsQuery.data?.total ?? 0;

  // ---- filtros client-side (multi status/source + range + unassigned) -------
  const filteredItems = useMemo(() => {
    let items = data?.items ?? [];

    if (statuses.length >= 2) {
      items = items.filter((l) => statuses.includes(l.status));
    }
    if (sources.length >= 2) {
      items = items.filter((l) => sources.includes(l.source));
    }
    if (wantsUnassigned) {
      items = items.filter((l) => !l.assignedToUserId);
    }
    if (dateRange?.from) {
      const fromMs = dateRange.from.getTime();
      items = items.filter((l) => new Date(l.createdAt).getTime() >= fromMs);
    }
    if (dateRange?.to) {
      // inclui o dia inteiro
      const toMs = new Date(dateRange.to).setHours(23, 59, 59, 999);
      items = items.filter((l) => new Date(l.createdAt).getTime() <= toMs);
    }

    // Ordenação
    const dir = sortDir === 'asc' ? 1 : -1;
    items = [...items].sort((a, b) => {
      switch (sortKey) {
        case 'name':
          return ((a.name ?? '').localeCompare(b.name ?? '')) * dir;
        case 'status':
          return a.status.localeCompare(b.status) * dir;
        case 'createdAt':
        default:
          return (
            (new Date(a.createdAt).getTime() - new Date(b.createdAt).getTime()) * dir
          );
      }
    });

    return items;
  }, [data?.items, statuses, sources, wantsUnassigned, dateRange, sortKey, sortDir]);

  const totalPages = data?.totalPages ?? 0;
  const total = data?.total ?? 0;

  // ---- handlers --------------------------------------------------------------
  function handleSearchChange(value: string) {
    setParams({ q: value || null, page: null });
  }

  function handleStatusesChange(next: LeadStatus[]) {
    setParams({ status: serializeList(next), page: null });
  }

  function handleSourcesChange(next: LeadSource[]) {
    setParams({ source: serializeList(next), page: null });
  }

  function handleAssigneeChange(next: string | null) {
    setParams({ assignee: next, page: null });
  }

  function handleDateRangeChange(range: DateRange | undefined) {
    setParams({
      from: range?.from ? range.from.toISOString().slice(0, 10) : null,
      to: range?.to ? range.to.toISOString().slice(0, 10) : null,
      page: null,
    });
  }

  function toggleSort(key: SortKey) {
    if (sortKey === key) {
      setParam('dir', sortDir === 'asc' ? 'desc' : 'asc');
    } else {
      setParams({ sort: key, dir: 'desc' });
    }
  }

  async function quickAction(
    lead: Lead,
    patch: { status?: LeadStatus; assignedToUserId?: string },
  ) {
    setActionPending(lead.id);
    try {
      await apiClient.patch(`/api/v1/leads/${lead.id}`, patch);
      await qc.invalidateQueries({ queryKey: leadKeys.all });
    } finally {
      setActionPending(null);
    }
  }

  const activeFilterCount =
    (statuses.length > 0 ? 1 : 0) +
    (sources.length > 0 ? 1 : 0) +
    (assignee ? 1 : 0) +
    (dateRange ? 1 : 0) +
    (search ? 1 : 0);
  const hasActiveFilters = activeFilterCount > 0;

  const filterBar = (
    <LeadsFilterBar
      statuses={statuses}
      onStatusesChange={handleStatusesChange}
      sources={sources}
      onSourcesChange={handleSourcesChange}
      assignee={assignee}
      onAssigneeChange={handleAssigneeChange}
      staffUsers={staffUsers ?? []}
      currentUserId={user?.id}
      dateRange={dateRange}
      onDateRangeChange={handleDateRangeChange}
      hasActiveFilters={hasActiveFilters}
      onClear={clearAll}
    />
  );

  function emptyMessage() {
    if (debouncedSearch) {
      return `Nenhum lead encontrado para "${debouncedSearch}".`;
    }
    if (hasActiveFilters) {
      return 'Nenhum lead encontrado com os filtros aplicados. Tente limpar algum.';
    }
    return 'Nenhum lead cadastrado ainda.';
  }

  return (
    <div className="space-y-6 p-6">
      <PageHeader
        title="Leads"
        description="Captura de contatos via Escore Light, formulário /contato, WhatsApp, email e newsletter."
      />

      {/* Quick stats */}
      {user?.id && (
        <div className="grid gap-3 sm:grid-cols-3">
          <Card>
            <CardContent className="pt-6">
              <p className="text-xs text-muted-foreground">Total</p>
              <p className="text-2xl font-semibold">{total}</p>
            </CardContent>
          </Card>
          <Card>
            <CardContent className="pt-6">
              <p className="text-xs text-muted-foreground">Atribuídos a mim</p>
              <p className="text-2xl font-semibold">{myLeadsCount}</p>
            </CardContent>
          </Card>
          <Card>
            <CardContent className="pt-6">
              <p className="text-xs text-muted-foreground">Filtros ativos</p>
              <p className="text-2xl font-semibold">{activeFilterCount}</p>
            </CardContent>
          </Card>
        </div>
      )}

      {/* Filtros */}
      <div className="space-y-3">
        <div className="flex flex-col gap-2 sm:flex-row sm:items-center">
          <div className="relative flex-1">
            <Search className="absolute left-3 top-2.5 h-4 w-4 text-muted-foreground" />
            <Input
              placeholder="Buscar por nome, email ou telefone…"
              value={search}
              onChange={(e) => handleSearchChange(e.target.value)}
              className="pl-9"
              aria-label="Buscar leads"
            />
          </div>
          {/* Mobile: chips dentro do sheet. >=md: chips inline ao lado da busca */}
          <div className="md:hidden">
            <FiltersMobileSheet activeCount={activeFilterCount}>
              {filterBar}
            </FiltersMobileSheet>
          </div>
        </div>
        <div className="hidden md:block">{filterBar}</div>
      </div>

      <Card>
        <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-3">
          <CardTitle className="text-base">
            {isLoading
              ? 'Carregando…'
              : `${filteredItems.length}${
                  filteredItems.length !== total ? ` de ${total}` : ''
                } lead${total === 1 ? '' : 's'}`}
          </CardTitle>
          {totalPages > 1 && (
            <div className="flex items-center gap-2">
              <Button
                variant="outline"
                size="sm"
                disabled={page === 0}
                onClick={() => setParam('page', String(Math.max(0, page - 1)))}
                aria-label="Página anterior"
              >
                <ChevronLeft className="h-4 w-4" />
              </Button>
              <span className="text-sm text-muted-foreground">
                {page + 1} / {totalPages}
              </span>
              <Button
                variant="outline"
                size="sm"
                disabled={page >= totalPages - 1}
                onClick={() => setParam('page', String(page + 1))}
                aria-label="Próxima página"
              >
                <ChevronRight className="h-4 w-4" />
              </Button>
            </div>
          )}
        </CardHeader>
        <CardContent>
          {isLoading ? (
            <div className="space-y-3">
              {Array.from({ length: 5 }).map((_, i) => (
                <Skeleton key={i} className="h-16 w-full" />
              ))}
            </div>
          ) : filteredItems.length === 0 ? (
            <div className="py-16 text-center text-muted-foreground">{emptyMessage()}</div>
          ) : (
            <div className="overflow-x-auto">
              <table className="w-full text-sm">
                <thead>
                  <tr className="border-b text-left text-muted-foreground">
                    <SortableHeader
                      label="Quando"
                      active={sortKey === 'createdAt'}
                      dir={sortDir}
                      onClick={() => toggleSort('createdAt')}
                    />
                    <th className="px-2 py-3 font-medium">Origem</th>
                    <SortableHeader
                      label="Contato"
                      active={sortKey === 'name'}
                      dir={sortDir}
                      onClick={() => toggleSort('name')}
                    />
                    <SortableHeader
                      label="Status"
                      active={sortKey === 'status'}
                      dir={sortDir}
                      onClick={() => toggleSort('status')}
                    />
                    <th className="px-2 py-3 font-medium">Canais</th>
                    <th className="px-2 py-3 font-medium">Atribuído</th>
                    <th className="px-2 py-3 font-medium text-right">Ações</th>
                  </tr>
                </thead>
                <tbody>
                  {filteredItems.map((lead) => (
                    <tr key={lead.id} className="border-b hover:bg-muted/40">
                      <td
                        className="px-2 py-3 cursor-pointer"
                        onClick={() => router.push(`/leads/${lead.id}`)}
                      >
                        <div>
                          {format(new Date(lead.createdAt), 'dd/MM HH:mm', { locale: ptBR })}
                        </div>
                      </td>
                      <td
                        className="px-2 py-3 cursor-pointer"
                        onClick={() => router.push(`/leads/${lead.id}`)}
                      >
                        <Badge variant="outline">{SOURCE_LABELS[lead.source]}</Badge>
                        {lead.utmCampaign && (
                          <div
                            className="mt-1 text-[10px] font-mono text-amber-700 truncate max-w-[160px]"
                            title={`utm_source: ${lead.utmSource ?? '—'}\nutm_medium: ${lead.utmMedium ?? '—'}\nutm_campaign: ${lead.utmCampaign}`}
                          >
                            📣 {lead.utmCampaign}
                          </div>
                        )}
                      </td>
                      <td
                        className="px-2 py-3 cursor-pointer"
                        onClick={() => router.push(`/leads/${lead.id}`)}
                      >
                        <div className="font-medium">{lead.name ?? '—'}</div>
                        <div className="text-xs text-muted-foreground space-x-2">
                          {lead.email && <span>{lead.email}</span>}
                          {lead.phone && <span>{lead.phone}</span>}
                        </div>
                      </td>
                      <td
                        className="px-2 py-3 cursor-pointer"
                        onClick={() => router.push(`/leads/${lead.id}`)}
                      >
                        <span
                          className={`inline-flex items-center rounded-md border px-2 py-0.5 text-xs ${STATUS_COLORS[lead.status]}`}
                        >
                          {STATUS_LABELS[lead.status]}
                        </span>
                      </td>
                      <td
                        className="px-2 py-3 cursor-pointer"
                        onClick={() => router.push(`/leads/${lead.id}`)}
                      >
                        <div className="flex gap-1">
                          {lead.emailOptIn && (
                            <Badge variant="secondary" className="text-xs">
                              Email
                            </Badge>
                          )}
                          {lead.whatsAppOptIn && (
                            <Badge variant="secondary" className="text-xs">
                              WhatsApp
                            </Badge>
                          )}
                          {lead.newsletterOptIn && (
                            <Badge variant="secondary" className="text-xs">
                              Newsletter
                            </Badge>
                          )}
                        </div>
                      </td>
                      <td
                        className="px-2 py-3 cursor-pointer text-xs"
                        onClick={() => router.push(`/leads/${lead.id}`)}
                      >
                        {lead.assignedTo?.name ?? (
                          <span className="text-muted-foreground">—</span>
                        )}
                      </td>
                      <td className="px-2 py-3 text-right whitespace-nowrap">
                        <div className="inline-flex gap-1">
                          {lead.status === 'new' && (
                            <Button
                              size="sm"
                              variant="outline"
                              disabled={actionPending === lead.id}
                              onClick={() => quickAction(lead, { status: 'contacted' })}
                              title="Marcar como contatado"
                            >
                              Contatado
                            </Button>
                          )}
                          {user?.id && lead.assignedToUserId !== user.id && (
                            <Button
                              size="sm"
                              variant="outline"
                              disabled={actionPending === lead.id}
                              onClick={() => quickAction(lead, { assignedToUserId: user.id })}
                              title="Atribuir a mim"
                            >
                              Pra mim
                            </Button>
                          )}
                        </div>
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          )}
        </CardContent>
      </Card>
    </div>
  );
}

interface SortableHeaderProps {
  label: string;
  active: boolean;
  dir: SortDir;
  onClick: () => void;
}

function SortableHeader({ label, active, dir, onClick }: SortableHeaderProps) {
  return (
    <th className="px-2 py-3 font-medium">
      <button
        type="button"
        onClick={onClick}
        className="inline-flex items-center gap-1 hover:text-foreground"
      >
        {label}
        {active &&
          (dir === 'asc' ? (
            <ArrowUp className="h-3 w-3" />
          ) : (
            <ArrowDown className="h-3 w-3" />
          ))}
      </button>
    </th>
  );
}
