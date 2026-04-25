'use client';

import { useState, use } from 'react';
import Link from 'next/link';
import { useRouter } from 'next/navigation';
import { format } from 'date-fns';
import { ptBR } from 'date-fns/locale';
import {
  ArrowLeft,
  ExternalLink,
  Mail,
  MessageSquare,
  StickyNote,
  Trash2,
  UserCheck,
  Activity,
  ShieldCheck,
} from 'lucide-react';

import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Badge } from '@/components/ui/badge';
import { Skeleton } from '@/components/ui/skeleton';
import { useRequireAuth, useAuth } from '@/lib/use-auth';
import { usePatientGuard } from '@/lib/use-patient-guard';
import { isGranted } from '@/lib/auth-store';
import {
  useLead,
  useUpdateLead,
  useAddLeadNote,
  useConvertLead,
  useDeleteLead,
  useSendLeadMessage,
  useStaffUsers,
  SOURCE_LABELS,
  STATUS_LABELS,
  STATUS_COLORS,
  type LeadStatus,
} from '@/lib/api/leads-api';
import { EmailReplyBlock } from '@/components/leads/EmailReplyBlock';

const SITE_PUBLIC_URL =
  process.env.NEXT_PUBLIC_SITE_PUBLIC_URL?.replace(/\/$/, '') ?? 'http://localhost:3002';

const STATUS_OPTIONS: LeadStatus[] = [
  'new',
  'contacted',
  'qualified',
  'converted',
  'lost',
  'unsubscribed',
];

function activityIcon(type: string, channel: string) {
  if (type === 'message_status_changed') return Activity;
  if (channel === 'email') return Mail;
  if (channel === 'whatsapp') return MessageSquare;
  if (type === 'note_added') return StickyNote;
  if (type === 'converted') return UserCheck;
  if (type === 'status_changed') return Activity;
  return Activity;
}

export default function LeadDetailPage({ params }: { params: Promise<{ id: string }> }) {
  useRequireAuth();
  usePatientGuard();
  const { user } = useAuth();
  const router = useRouter();

  const { id } = use(params);
  const { data: lead, isLoading } = useLead(id);
  const updateLead = useUpdateLead(id);
  const addNote = useAddLeadNote(id);
  const convertLead = useConvertLead(id);
  const deleteLead = useDeleteLead(id);
  const sendMessage = useSendLeadMessage(id);
  const { data: staffUsers } = useStaffUsers();

  const [note, setNote] = useState('');
  const [waMessage, setWaMessage] = useState('');
  const [convertOpen, setConvertOpen] = useState(false);
  const [convertGender, setConvertGender] = useState<'male' | 'female' | 'other'>('other');
  const [convertBirthDate, setConvertBirthDate] = useState('');
  const [confirmDelete, setConfirmDelete] = useState(false);

  const isAdmin = isGranted(user, 'admin');

  if (isLoading || !lead) {
    return (
      <div className="space-y-4 p-6">
        <Skeleton className="h-10 w-64" />
        <Skeleton className="h-48 w-full" />
        <Skeleton className="h-96 w-full" />
      </div>
    );
  }

  return (
    <div className="space-y-6 p-6">
      <Link href="/leads" className="inline-flex items-center text-sm text-muted-foreground hover:text-foreground">
        <ArrowLeft className="mr-1 h-4 w-4" /> Todos os leads
      </Link>

      {/* Identificação */}
      <Card>
        <CardHeader>
          <div className="flex flex-wrap items-start justify-between gap-3">
            <div className="min-w-0">
              <CardTitle className="text-2xl break-words">{lead.name ?? '(sem nome)'}</CardTitle>
              <p className="mt-1 text-sm text-muted-foreground">
                Capturado {format(new Date(lead.createdAt), "dd 'de' MMMM 'às' HH:mm", { locale: ptBR })} via{' '}
                <Badge variant="outline">{SOURCE_LABELS[lead.source]}</Badge>
                {lead.source === 'whatsapp_inbound' && lead.convertedPatientId && (
                  <Badge className="ml-2 bg-emerald-100 text-emerald-900 border-emerald-200">
                    Paciente cadastrado
                  </Badge>
                )}
              </p>
            </div>
            <span
              className={`inline-flex items-center rounded-md border px-3 py-1 text-sm ${STATUS_COLORS[lead.status]}`}
            >
              {STATUS_LABELS[lead.status]}
            </span>
          </div>
        </CardHeader>
        <CardContent>
          <div className="grid gap-4 sm:grid-cols-2">
            <div>
              <div className="text-xs font-medium uppercase tracking-wide text-muted-foreground">Email</div>
              <div className="mt-1">
                {lead.email ? (
                  <a href={`mailto:${lead.email}`} className="text-blue-600 hover:underline">
                    {lead.email}
                  </a>
                ) : (
                  <span className="text-muted-foreground">—</span>
                )}
                {lead.emailOptIn && <Badge variant="secondary" className="ml-2 text-xs">Opt-in</Badge>}
              </div>
            </div>
            <div>
              <div className="text-xs font-medium uppercase tracking-wide text-muted-foreground">Telefone</div>
              <div className="mt-1">
                {lead.phone ? (
                  <a
                    href={`https://wa.me/${lead.phone.replace(/\D/g, '')}`}
                    target="_blank"
                    rel="noopener noreferrer"
                    className="text-blue-600 hover:underline"
                  >
                    {lead.phone}
                  </a>
                ) : (
                  <span className="text-muted-foreground">—</span>
                )}
                {lead.whatsAppOptIn && <Badge variant="secondary" className="ml-2 text-xs">WhatsApp</Badge>}
              </div>
            </div>
            {lead.message && (
              <div className="sm:col-span-2">
                <div className="text-xs font-medium uppercase tracking-wide text-muted-foreground">Primeira mensagem</div>
                <p className="mt-1 whitespace-pre-wrap text-sm">{lead.message}</p>
              </div>
            )}
            {lead.metadata && Object.keys(lead.metadata).length > 0 && (
              <div className="sm:col-span-2">
                <div className="text-xs font-medium uppercase tracking-wide text-muted-foreground">Metadados</div>
                <pre className="mt-1 overflow-x-auto rounded bg-muted p-2 text-xs">{JSON.stringify(lead.metadata, null, 2)}</pre>
              </div>
            )}
            {(lead.utmSource || lead.utmMedium || lead.utmCampaign) && (
              <div className="sm:col-span-2 rounded-md border border-amber-200 bg-amber-50/50 p-3">
                <div className="text-xs font-medium uppercase tracking-wide text-amber-900">
                  Origem da campanha
                </div>
                <div className="mt-2 flex flex-wrap items-center gap-2">
                  {lead.utmSource && (
                    <Badge variant="outline" className="font-mono text-xs">
                      src: {lead.utmSource}
                    </Badge>
                  )}
                  {lead.utmMedium && (
                    <Badge variant="outline" className="font-mono text-xs">
                      med: {lead.utmMedium}
                    </Badge>
                  )}
                  {lead.utmCampaign && (
                    <Badge variant="outline" className="font-mono text-xs">
                      cmp: {lead.utmCampaign}
                    </Badge>
                  )}
                  {lead.utmTerm && (
                    <Badge variant="outline" className="font-mono text-xs">
                      term: {lead.utmTerm}
                    </Badge>
                  )}
                </div>
              </div>
            )}
          </div>

          {/* Ações */}
          <div className="mt-6 flex flex-wrap items-center gap-3 border-t pt-4">
            <span className="text-sm text-muted-foreground">Mudar status:</span>
            {STATUS_OPTIONS.map((s) => (
              <Button
                key={s}
                variant={s === lead.status ? 'default' : 'outline'}
                size="sm"
                disabled={updateLead.isPending || s === lead.status}
                onClick={() => updateLead.mutate({ status: s })}
              >
                {STATUS_LABELS[s]}
              </Button>
            ))}
          </div>

          {/* Atribuição */}
          <div className="mt-3 flex flex-wrap items-center gap-3 border-t pt-4">
            <span className="text-sm text-muted-foreground">Atribuído a:</span>
            <select
              value={lead.assignedToUserId ?? ''}
              onChange={(e) => {
                const value = e.target.value;
                updateLead.mutate({
                  assignedToUserId: value === '' ? undefined : value,
                });
              }}
              disabled={updateLead.isPending}
              className="h-9 rounded-md border border-input bg-transparent px-3 py-1 text-sm shadow-sm focus:outline-none focus:ring-1 focus:ring-ring"
            >
              <option value="">— Sem atribuição —</option>
              {user?.id && (
                <option value={user.id}>Eu ({user.name})</option>
              )}
              {staffUsers
                ?.filter((u) => u.id !== user?.id)
                .map((u) => (
                  <option key={u.id} value={u.id}>
                    {u.name}
                  </option>
                ))}
            </select>
          </div>

          {/* Atalhos: ver Light, converter, etc. */}
          <div className="mt-3 flex flex-wrap items-center gap-3 border-t pt-4">
            {lead.anonymousScoreSessionId && (
              <a
                href={`${SITE_PUBLIC_URL}/pt/escore-plenya/resultado/${lead.anonymousScoreSessionId}`}
                target="_blank"
                rel="noopener noreferrer"
                className="inline-flex items-center gap-1 text-sm text-blue-600 hover:underline"
              >
                <ExternalLink className="h-3.5 w-3.5" /> Ver resultado do Light
              </a>
            )}

            {lead.status !== 'converted' && (
              <Button size="sm" variant="default" onClick={() => setConvertOpen((v) => !v)}>
                <UserCheck className="mr-1 h-4 w-4" /> Converter em paciente
              </Button>
            )}
            {lead.status === 'converted' && lead.convertedPatientId && (
              <Button
                size="sm"
                variant="outline"
                onClick={() => router.push(`/patients/${lead.convertedPatientId}`)}
              >
                <UserCheck className="mr-1 h-4 w-4" /> Abrir paciente
              </Button>
            )}
          </div>

          {convertOpen && lead.status !== 'converted' && (
            <div className="mt-4 rounded-md border border-dashed border-primary/30 bg-primary/5 p-4 space-y-3">
              <p className="text-sm">
                Os campos abaixo viram o cadastro inicial do Patient — admin pode editar depois.
              </p>
              <div className="grid gap-3 sm:grid-cols-2">
                <label className="block">
                  <span className="text-xs uppercase tracking-wide text-muted-foreground">Data de nascimento</span>
                  <input
                    type="date"
                    value={convertBirthDate}
                    onChange={(e) => setConvertBirthDate(e.target.value)}
                    className="mt-1 w-full rounded-md border border-input bg-transparent px-3 py-1.5 text-sm"
                  />
                </label>
                <label className="block">
                  <span className="text-xs uppercase tracking-wide text-muted-foreground">Gênero</span>
                  <select
                    value={convertGender}
                    onChange={(e) => setConvertGender(e.target.value as 'male' | 'female' | 'other')}
                    className="mt-1 w-full rounded-md border border-input bg-transparent px-3 py-1.5 text-sm"
                  >
                    <option value="other">Outro / não informado</option>
                    <option value="female">Feminino</option>
                    <option value="male">Masculino</option>
                  </select>
                </label>
              </div>
              <div className="flex justify-end gap-2">
                <Button variant="outline" size="sm" onClick={() => setConvertOpen(false)}>
                  Cancelar
                </Button>
                <Button
                  size="sm"
                  disabled={convertLead.isPending || !convertBirthDate}
                  onClick={() => {
                    if (!convertBirthDate) return;
                    convertLead.mutate(
                      {
                        gender: convertGender,
                        birthDate: new Date(convertBirthDate).toISOString(),
                      },
                      {
                        onSuccess: () => setConvertOpen(false),
                      }
                    );
                  }}
                >
                  {convertLead.isPending ? 'Convertendo…' : 'Confirmar conversão'}
                </Button>
              </div>
            </div>
          )}
        </CardContent>
      </Card>

      {/* Adicionar nota */}
      <Card>
        <CardHeader>
          <CardTitle className="text-base">Adicionar nota interna</CardTitle>
        </CardHeader>
        <CardContent>
          <textarea
            value={note}
            onChange={(e) => setNote(e.target.value)}
            placeholder="Ex: ligação às 14h, ficou de retornar amanhã..."
            className="min-h-[80px] w-full rounded-md border border-input bg-transparent p-3 text-sm focus:outline-none focus:ring-1 focus:ring-ring"
          />
          <div className="mt-3 flex justify-end">
            <Button
              size="sm"
              disabled={!note.trim() || addNote.isPending}
              onClick={() => {
                addNote.mutate(note.trim(), {
                  onSuccess: () => setNote(''),
                });
              }}
            >
              {addNote.isPending ? 'Salvando…' : 'Salvar nota'}
            </Button>
          </div>
        </CardContent>
      </Card>

      {/* Email reply (Bloco 3) — lista threads e abre modal de resposta */}
      <EmailReplyBlock lead={lead} />

      {/* Chat WhatsApp inline (session messages, janela 24h) */}
      {lead.phone && lead.whatsAppOptIn && lead.status !== 'unsubscribed' && (
        <Card>
          <CardHeader>
            <CardTitle className="flex items-center gap-2 text-base">
              <MessageSquare className="h-4 w-4" /> Conversa WhatsApp
            </CardTitle>
            <p className="text-sm text-muted-foreground">
              {(() => {
                if (!lead.lastInboundAt) {
                  return 'Sem mensagem inbound recente — fora da janela de 24h. Use template aprovado para iniciar conversa.';
                }
                const inboundTs = new Date(lead.lastInboundAt).getTime();
                const elapsedH = (Date.now() - inboundTs) / 36e5;
                const remainingH = 24 - elapsedH;
                if (remainingH <= 0) {
                  return 'Janela de 24h expirada — para iniciar nova conversa, use template aprovado.';
                }
                const h = Math.floor(remainingH);
                const m = Math.floor((remainingH - h) * 60);
                return `Janela de 24h: ${h}h ${m}min restantes`;
              })()}
            </p>
          </CardHeader>
          <CardContent>
            {(() => {
              const inWindow =
                lead.lastInboundAt &&
                Date.now() - new Date(lead.lastInboundAt).getTime() < 24 * 36e5;
              return (
                <>
                  <textarea
                    value={waMessage}
                    onChange={(e) => setWaMessage(e.target.value)}
                    disabled={!inWindow}
                    placeholder={
                      inWindow
                        ? 'Mensagem para o cliente…'
                        : 'Janela 24h fechada. Mande um template ao invés disso.'
                    }
                    className="min-h-[100px] w-full rounded-md border border-input bg-transparent p-3 text-sm focus:outline-none focus:ring-1 focus:ring-ring disabled:opacity-50"
                  />
                  <div className="mt-3 flex justify-end">
                    <Button
                      size="sm"
                      disabled={!inWindow || !waMessage.trim() || sendMessage.isPending}
                      onClick={() => {
                        sendMessage.mutate(waMessage.trim(), {
                          onSuccess: () => setWaMessage(''),
                          onError: (err: unknown) => {
                            alert(
                              err instanceof Error ? err.message : 'Falha ao enviar mensagem'
                            );
                          },
                        });
                      }}
                    >
                      {sendMessage.isPending ? 'Enviando…' : 'Enviar WhatsApp'}
                    </Button>
                  </div>
                </>
              );
            })()}
          </CardContent>
        </Card>
      )}

      {/* Timeline */}
      <Card>
        <CardHeader>
          <CardTitle className="text-base">Histórico</CardTitle>
        </CardHeader>
        <CardContent>
          {!lead.activities || lead.activities.length === 0 ? (
            <p className="py-6 text-center text-sm text-muted-foreground">Nenhuma atividade registrada.</p>
          ) : (
            (() => {
              // Agrupa message_status_changed sob a message_sent original (via wa_message_id)
              const statusByMsgID = new Map<string, typeof lead.activities>();
              const main: typeof lead.activities = [];
              for (const a of lead.activities) {
                if (a.type === 'message_status_changed') {
                  const wamid = (a.metadata as { wa_message_id?: string } | undefined)?.wa_message_id;
                  if (wamid) {
                    const arr = statusByMsgID.get(wamid) ?? [];
                    arr.push(a);
                    statusByMsgID.set(wamid, arr);
                    continue;
                  }
                }
                main.push(a);
              }
              return (
                <ol className="space-y-3">
                  {main.map((a) => {
                    const Icon = activityIcon(a.type, a.channel);
                    const wamid = (a.metadata as { wa_message_id?: string } | undefined)?.wa_message_id;
                    const childStatuses = wamid ? statusByMsgID.get(wamid) ?? [] : [];
                    return (
                      <li key={a.id} className="flex gap-3 border-l-2 border-muted pl-4">
                        <Icon className="mt-0.5 h-4 w-4 shrink-0 text-muted-foreground" />
                        <div className="min-w-0 flex-1">
                          <div className="text-xs text-muted-foreground break-words">
                            {format(new Date(a.createdAt), "dd/MM 'às' HH:mm", { locale: ptBR })} ·{' '}
                            <span className="capitalize">{a.type.replace('_', ' ')}</span>
                            {a.channel !== 'internal' && <span> · {a.channel}</span>}
                            {a.actor && <span> · por {a.actor.name}</span>}
                          </div>
                          {a.content && (
                            <p className="mt-1 whitespace-pre-wrap break-words text-sm">{a.content}</p>
                          )}
                          {childStatuses.length > 0 && (
                            <div className="mt-1 flex flex-wrap gap-2 text-xs">
                              {childStatuses.map((s) => {
                                const meta = s.metadata as { status?: string } | undefined;
                                const status = meta?.status ?? '?';
                                const failed = status === 'failed';
                                return (
                                  <span
                                    key={s.id}
                                    className={
                                      failed
                                        ? 'text-red-700'
                                        : 'text-muted-foreground'
                                    }
                                  >
                                    {format(new Date(s.createdAt), 'HH:mm:ss')} {status}
                                  </span>
                                );
                              })}
                            </div>
                          )}
                        </div>
                      </li>
                    );
                  })}
                </ol>
              );
            })()
          )}
        </CardContent>
      </Card>

      {/* LGPD info */}
      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2 text-base">
            <ShieldCheck className="h-4 w-4" /> Consentimento LGPD
          </CardTitle>
        </CardHeader>
        <CardContent className="text-sm text-muted-foreground">
          <p>
            Versão: <strong>{lead.consentVersion ?? '—'}</strong>
          </p>
          {lead.consentTimestamp && (
            <p>
              Aceito em:{' '}
              <strong>
                {format(new Date(lead.consentTimestamp), "dd/MM/yyyy 'às' HH:mm", { locale: ptBR })}
              </strong>
            </p>
          )}
          <p className="mt-2 text-xs">
            IP do aceite armazenado como hash SHA-256 (não recuperável). Lead pode solicitar exclusão pelo DPO.
          </p>
        </CardContent>
      </Card>

      {/* Zona de perigo — admin only */}
      {isAdmin && (
        <Card className="border-destructive/30">
          <CardHeader>
            <CardTitle className="flex items-center gap-2 text-base text-destructive">
              <Trash2 className="h-4 w-4" /> Excluir lead
            </CardTitle>
          </CardHeader>
          <CardContent className="space-y-3">
            <p className="text-sm text-muted-foreground">
              Excluir é definitivo (soft delete + atividades em cascade). Use para atender pedido LGPD do
              titular ou limpar leads de teste.
            </p>
            {!confirmDelete ? (
              <Button variant="destructive" size="sm" onClick={() => setConfirmDelete(true)}>
                Excluir este lead
              </Button>
            ) : (
              <div className="flex flex-wrap items-center gap-3">
                <span className="text-sm">Tem certeza?</span>
                <Button
                  variant="destructive"
                  size="sm"
                  disabled={deleteLead.isPending}
                  onClick={() => {
                    deleteLead.mutate(undefined, {
                      onSuccess: () => router.push('/leads'),
                    });
                  }}
                >
                  {deleteLead.isPending ? 'Excluindo…' : 'Sim, excluir definitivamente'}
                </Button>
                <Button variant="outline" size="sm" onClick={() => setConfirmDelete(false)}>
                  Cancelar
                </Button>
              </div>
            )}
          </CardContent>
        </Card>
      )}
    </div>
  );
}
