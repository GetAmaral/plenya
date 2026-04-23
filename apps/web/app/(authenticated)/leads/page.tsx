'use client';

import { useState } from 'react';
import { useRouter } from 'next/navigation';
import { format } from 'date-fns';
import { ptBR } from 'date-fns/locale';
import { ChevronLeft, ChevronRight, Search } from 'lucide-react';

import { useQueryClient } from '@tanstack/react-query';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Input } from '@/components/ui/input';
import { Button } from '@/components/ui/button';
import { Badge } from '@/components/ui/badge';
import { Skeleton } from '@/components/ui/skeleton';
import { PageHeader } from '@/components/layout/page-header';
import { useRequireAuth, useAuth } from '@/lib/use-auth';
import { usePatientGuard } from '@/lib/use-patient-guard';
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

const SOURCE_OPTIONS: Array<{ value: LeadSource | ''; label: string }> = [
  { value: '', label: 'Todas as origens' },
  { value: 'light_claim', label: 'Escore Light' },
  { value: 'contact_form', label: 'Formulário /contato' },
  { value: 'whatsapp_inbound', label: 'WhatsApp inbound' },
  { value: 'newsletter', label: 'Newsletter' },
  { value: 'manual', label: 'Manual' },
];

const STATUS_OPTIONS: Array<{ value: LeadStatus | ''; label: string }> = [
  { value: '', label: 'Todos os status' },
  { value: 'new', label: 'Novos' },
  { value: 'contacted', label: 'Contatados' },
  { value: 'qualified', label: 'Qualificados' },
  { value: 'converted', label: 'Convertidos' },
  { value: 'lost', label: 'Perdidos' },
  { value: 'unsubscribed', label: 'Descadastraram' },
];

export default function LeadsPage() {
  useRequireAuth();
  usePatientGuard();
  const router = useRouter();
  const qc = useQueryClient();
  const { user } = useAuth();

  const [search, setSearch] = useState('');
  const [source, setSource] = useState<LeadSource | ''>('');
  const [status, setStatus] = useState<LeadStatus | ''>('');
  const [emailOptIn, setEmailOptIn] = useState<'' | 'yes' | 'no'>('');
  const [whatsAppOptIn, setWhatsAppOptIn] = useState<'' | 'yes' | 'no'>('');
  const [assignedFilter, setAssignedFilter] = useState<'' | 'me' | 'unassigned' | string>('');
  const [page, setPage] = useState(0);
  const [actionPending, setActionPending] = useState<string | null>(null);
  const pageSize = 25;

  // Resolve assigned filter — backend aceita user UUID; "unassigned" precisa client-side
  const assignedToUserId =
    assignedFilter === 'me' ? user?.id : (assignedFilter && assignedFilter !== 'unassigned' ? assignedFilter : undefined);

  const filter: LeadFilter = {
    search: search || undefined,
    source: source || undefined,
    status: status || undefined,
    hasEmailOptIn: emailOptIn === '' ? undefined : emailOptIn === 'yes',
    hasWhatsAppOptIn: whatsAppOptIn === '' ? undefined : whatsAppOptIn === 'yes',
    assignedToUserId,
  };

  const { data, isLoading } = useLeads(filter, page, pageSize);
  const { data: staffUsers } = useStaffUsers();

  // Contador rápido pra "meus leads" (só dispara quando user.id resolveu —
  // evita fetch sem filtro retornando total geral durante o flash de carregamento)
  const myLeadsQuery = useLeads(
    { assignedToUserId: user?.id },
    0,
    1,
    { enabled: !!user?.id }
  );
  const myLeadsCount = myLeadsQuery.data?.total ?? 0;

  async function quickAction(lead: Lead, patch: { status?: LeadStatus; assignedToUserId?: string }) {
    setActionPending(lead.id);
    try {
      await apiClient.patch(`/api/v1/leads/${lead.id}`, patch);
      await qc.invalidateQueries({ queryKey: leadKeys.all });
    } finally {
      setActionPending(null);
    }
  }

  return (
    <div className="space-y-6 p-6">
      <PageHeader
        title="Leads"
        description="Captura de contatos via Escore Light, formulário /contato, WhatsApp e newsletter."
      />

      <Card>
        <CardHeader>
          <CardTitle>Filtros</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="grid gap-3 md:grid-cols-[1fr_220px_220px]">
            <div className="relative">
              <Search className="absolute left-3 top-2.5 h-4 w-4 text-muted-foreground" />
              <Input
                placeholder="Buscar por nome, email ou telefone…"
                value={search}
                onChange={(e) => {
                  setSearch(e.target.value);
                  setPage(0);
                }}
                className="pl-9"
              />
            </div>
            <select
              value={source}
              onChange={(e) => {
                setSource(e.target.value as LeadSource | '');
                setPage(0);
              }}
              className="h-9 w-full rounded-md border border-input bg-transparent px-3 py-1 text-sm shadow-sm focus:outline-none focus:ring-1 focus:ring-ring"
            >
              {SOURCE_OPTIONS.map((o) => (
                <option key={o.value} value={o.value}>
                  {o.label}
                </option>
              ))}
            </select>
            <select
              value={status}
              onChange={(e) => {
                setStatus(e.target.value as LeadStatus | '');
                setPage(0);
              }}
              className="h-9 w-full rounded-md border border-input bg-transparent px-3 py-1 text-sm shadow-sm focus:outline-none focus:ring-1 focus:ring-ring"
            >
              {STATUS_OPTIONS.map((o) => (
                <option key={o.value} value={o.value}>
                  {o.label}
                </option>
              ))}
            </select>
          </div>

          <div className="mt-3 grid gap-3 md:grid-cols-2">
            <select
              value={emailOptIn}
              onChange={(e) => {
                setEmailOptIn(e.target.value as '' | 'yes' | 'no');
                setPage(0);
              }}
              className="h-9 w-full rounded-md border border-input bg-transparent px-3 py-1 text-sm shadow-sm focus:outline-none focus:ring-1 focus:ring-ring"
            >
              <option value="">Email opt-in: qualquer</option>
              <option value="yes">Com opt-in de email</option>
              <option value="no">Sem opt-in de email</option>
            </select>
            <select
              value={whatsAppOptIn}
              onChange={(e) => {
                setWhatsAppOptIn(e.target.value as '' | 'yes' | 'no');
                setPage(0);
              }}
              className="h-9 w-full rounded-md border border-input bg-transparent px-3 py-1 text-sm shadow-sm focus:outline-none focus:ring-1 focus:ring-ring"
            >
              <option value="">WhatsApp opt-in: qualquer</option>
              <option value="yes">Com opt-in de WhatsApp</option>
              <option value="no">Sem opt-in de WhatsApp</option>
            </select>
          </div>
          <div className="mt-3">
            <select
              value={assignedFilter}
              onChange={(e) => {
                setAssignedFilter(e.target.value);
                setPage(0);
              }}
              className="h-9 w-full rounded-md border border-input bg-transparent px-3 py-1 text-sm shadow-sm focus:outline-none focus:ring-1 focus:ring-ring md:max-w-xs"
            >
              <option value="">Atribuído a: qualquer</option>
              <option value="me">Atribuídos a mim</option>
              {staffUsers
                ?.filter((u) => u.id !== user?.id)
                .map((u) => (
                  <option key={u.id} value={u.id}>
                    Atribuídos a {u.name}
                  </option>
                ))}
            </select>
          </div>
        </CardContent>
      </Card>

      {/* Contador rápido */}
      {user?.id && (
        <p className="text-sm text-muted-foreground">
          <strong>{myLeadsCount}</strong> lead{myLeadsCount === 1 ? '' : 's'} atribuído{myLeadsCount === 1 ? '' : 's'} a mim
        </p>
      )}

      <Card>
        <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-3">
          <CardTitle className="text-base">
            {data ? `${data.total} lead${data.total === 1 ? '' : 's'}` : 'Carregando…'}
          </CardTitle>
          {data && data.totalPages > 1 && (
            <div className="flex items-center gap-2">
              <Button
                variant="outline"
                size="sm"
                disabled={page === 0}
                onClick={() => setPage((p) => Math.max(0, p - 1))}
              >
                <ChevronLeft className="h-4 w-4" />
              </Button>
              <span className="text-sm text-muted-foreground">
                {page + 1} / {data.totalPages}
              </span>
              <Button
                variant="outline"
                size="sm"
                disabled={page >= data.totalPages - 1}
                onClick={() => setPage((p) => p + 1)}
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
          ) : !data || data.items.length === 0 ? (
            <div className="py-16 text-center text-muted-foreground">
              Nenhum lead encontrado para os filtros atuais.
            </div>
          ) : (
            <div className="overflow-x-auto">
              <table className="w-full text-sm">
                <thead>
                  <tr className="border-b text-left text-muted-foreground">
                    <th className="px-2 py-3 font-medium">Quando</th>
                    <th className="px-2 py-3 font-medium">Origem</th>
                    <th className="px-2 py-3 font-medium">Contato</th>
                    <th className="px-2 py-3 font-medium">Status</th>
                    <th className="px-2 py-3 font-medium">Canais</th>
                    <th className="px-2 py-3 font-medium">Atribuído</th>
                    <th className="px-2 py-3 font-medium text-right">Ações</th>
                  </tr>
                </thead>
                <tbody>
                  {data.items.map((lead) => (
                    <tr key={lead.id} className="border-b hover:bg-muted/40">
                      <td className="px-2 py-3 cursor-pointer" onClick={() => router.push(`/leads/${lead.id}`)}>
                        <div>{format(new Date(lead.createdAt), 'dd/MM HH:mm', { locale: ptBR })}</div>
                      </td>
                      <td className="px-2 py-3 cursor-pointer" onClick={() => router.push(`/leads/${lead.id}`)}>
                        <Badge variant="outline">{SOURCE_LABELS[lead.source]}</Badge>
                      </td>
                      <td className="px-2 py-3 cursor-pointer" onClick={() => router.push(`/leads/${lead.id}`)}>
                        <div className="font-medium">{lead.name ?? '—'}</div>
                        <div className="text-xs text-muted-foreground space-x-2">
                          {lead.email && <span>{lead.email}</span>}
                          {lead.phone && <span>{lead.phone}</span>}
                        </div>
                      </td>
                      <td className="px-2 py-3 cursor-pointer" onClick={() => router.push(`/leads/${lead.id}`)}>
                        <span
                          className={`inline-flex items-center rounded-md border px-2 py-0.5 text-xs ${STATUS_COLORS[lead.status]}`}
                        >
                          {STATUS_LABELS[lead.status]}
                        </span>
                      </td>
                      <td className="px-2 py-3 cursor-pointer" onClick={() => router.push(`/leads/${lead.id}`)}>
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
                      <td className="px-2 py-3 cursor-pointer text-xs" onClick={() => router.push(`/leads/${lead.id}`)}>
                        {lead.assignedTo?.name ?? <span className="text-muted-foreground">—</span>}
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
