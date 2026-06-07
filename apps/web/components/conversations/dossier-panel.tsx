'use client';

/**
 * DossierPanel — visão 360 SOCIAL de uma pessoa (memória de relacionamento da Lívia).
 * Resumo rolante + fatos atômicos (key/value) agrupados por categoria, com proveniência
 * (IA vs equipe) e edição manual. Só dado social — nada clínico (LGPD/CFM). O conteúdo é
 * alimentado pela IA (job) e pela equipe (aqui). Plano: docs/emr/plano-livia-memoria-dossie-360.md.
 */

import { useMemo, useState } from 'react';
import { Bot, Cake, Check, Gift, Pencil, Plus, Trash2, User, X } from 'lucide-react';
import { toast } from 'sonner';

import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetHeader,
  SheetTitle,
} from '@/components/ui/sheet';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Badge } from '@/components/ui/badge';
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';
import {
  DOSSIER_CATEGORY_LABELS,
  DOSSIER_EVENT_TYPE_LABELS,
  useAddDossierEvent,
  useAddDossierFact,
  useAddDossierPerson,
  useDeleteDossierEvent,
  useDeleteDossierFact,
  useDeleteDossierPerson,
  useDossier,
  useUpdateDossierEvent,
  useUpdateDossierFact,
  type ConversationOwnerType,
  type DossierEvent,
  type DossierFact,
  type DossierPerson,
} from '@/lib/api/conversations-api';

const CATEGORY_ORDER = [
  'identidade_social',
  'familia_rede',
  'preferencias_atendimento',
  'contexto_chegada',
  'relacionamento',
];

interface Props {
  ownerType: ConversationOwnerType;
  ownerId: string;
  name: string;
  open: boolean;
  onOpenChange: (open: boolean) => void;
}

export function DossierPanel({ ownerType, ownerId, name, open, onOpenChange }: Props) {
  const { data, isLoading } = useDossier(ownerType, ownerId);
  const addFact = useAddDossierFact(ownerType, ownerId);
  const updateFact = useUpdateDossierFact(ownerType, ownerId);
  const deleteFact = useDeleteDossierFact(ownerType, ownerId);
  const addPerson = useAddDossierPerson(ownerType, ownerId);
  const deletePerson = useDeleteDossierPerson(ownerType, ownerId);
  const addEvent = useAddDossierEvent(ownerType, ownerId);
  const updateEvent = useUpdateDossierEvent(ownerType, ownerId);
  const deleteEvent = useDeleteDossierEvent(ownerType, ownerId);

  const grouped = useMemo(() => {
    const g: Record<string, DossierFact[]> = {};
    for (const f of data?.facts ?? []) {
      (g[f.category] ??= []).push(f);
    }
    return g;
  }, [data?.facts]);

  const categories = useMemo(
    () =>
      CATEGORY_ORDER.filter((c) => grouped[c]?.length).concat(
        Object.keys(grouped).filter((c) => !CATEGORY_ORDER.includes(c))
      ),
    [grouped]
  );

  return (
    <Sheet open={open} onOpenChange={onOpenChange}>
      <SheetContent side="right" className="flex w-full flex-col gap-0 p-0 sm:max-w-md">
        <SheetHeader className="border-b border-border px-5 py-4">
          <SheetTitle className="flex flex-wrap items-center gap-2">
            Dossiê
            <Badge variant="outline" className="text-[10px] uppercase tracking-wide">
              {ownerType === 'patient' ? 'Paciente' : 'Lead'}
            </Badge>
            {data?.continuumActive && (
              <Badge className="bg-amber-100 text-amber-900 hover:bg-amber-100 text-[10px] uppercase tracking-wide">
                Continuum
              </Badge>
            )}
            {data?.frequent && (
              <Badge className="bg-emerald-100 text-emerald-900 hover:bg-emerald-100 text-[10px] uppercase tracking-wide">
                Frequente
              </Badge>
            )}
          </SheetTitle>
          <SheetDescription className="text-xs">
            O que a equipe e a Lívia sabem sobre {name}. Apenas informações sociais, nada clínico.
            {data?.appointmentsCompleted ? (
              <>
                {' '}
                {data.appointmentsCompleted} consulta{data.appointmentsCompleted > 1 ? 's' : ''} realizada
                {data.appointmentsCompleted > 1 ? 's' : ''}
                {data.lastConsultAt ? ` · última em ${formatDateBR(data.lastConsultAt)}` : ''}.
              </>
            ) : null}
          </SheetDescription>
        </SheetHeader>

        <div className="min-h-0 flex-1 overflow-y-auto px-5 py-4">
          {isLoading ? (
            <p className="text-sm text-muted-foreground">Carregando…</p>
          ) : (
            <div className="space-y-6">
              {/* Resumo da relação */}
              {data?.rollingSummary ? (
                <section>
                  <h3 className="mb-1.5 text-xs font-semibold uppercase tracking-wide text-muted-foreground">
                    Resumo da relação
                  </h3>
                  <p className="whitespace-pre-wrap rounded-md bg-muted/40 p-3 text-sm leading-relaxed">
                    {data.rollingSummary}
                  </p>
                </section>
              ) : null}

              {/* Fatos por categoria */}
              {categories.length > 0 ? (
                categories.map((cat) => (
                  <section key={cat}>
                    <h3 className="mb-1.5 text-xs font-semibold uppercase tracking-wide text-muted-foreground">
                      {DOSSIER_CATEGORY_LABELS[cat] ?? cat}
                    </h3>
                    <ul className="space-y-1.5">
                      {grouped[cat].map((f) => (
                        <FactRow
                          key={f.id}
                          fact={f}
                          onSave={(value) =>
                            updateFact.mutate(
                              { factId: f.id, value },
                              { onError: (e) => toast.error(e instanceof Error ? e.message : 'Falha ao salvar') }
                            )
                          }
                          onDelete={() =>
                            deleteFact.mutate(f.id, {
                              onError: (e) => toast.error(e instanceof Error ? e.message : 'Falha ao remover'),
                            })
                          }
                        />
                      ))}
                    </ul>
                  </section>
                ))
              ) : !data?.rollingSummary ? (
                <p className="text-sm text-muted-foreground">
                  Ainda não há nada registrado. A Lívia vai aprendendo na conversa, e você pode
                  anotar algo abaixo.
                </p>
              ) : null}

              {/* Adicionar fato */}
              <AddFactForm
                pending={addFact.isPending}
                onAdd={(category, value) =>
                  addFact.mutate(
                    { category, value },
                    { onError: (e) => toast.error(e instanceof Error ? e.message : 'Falha ao adicionar') }
                  )
                }
              />

              {/* Pessoas importantes */}
              <PeopleSection
                people={data?.people ?? []}
                pending={addPerson.isPending}
                onAdd={(name, relation, birthday) =>
                  addPerson.mutate(
                    { name, relation, birthday },
                    { onError: (e) => toast.error(e instanceof Error ? e.message : 'Falha ao adicionar pessoa') }
                  )
                }
                onDelete={(id) =>
                  deletePerson.mutate(id, {
                    onError: (e) => toast.error(e instanceof Error ? e.message : 'Falha ao remover'),
                  })
                }
              />

              {/* Eventos / datas */}
              <EventsSection
                events={data?.events ?? []}
                pending={addEvent.isPending}
                onAdd={(type, title, date, recurring) =>
                  addEvent.mutate(
                    { type, title, date, recurring },
                    { onError: (e) => toast.error(e instanceof Error ? e.message : 'Falha ao adicionar evento') }
                  )
                }
                onStatus={(eventId, status) =>
                  updateEvent.mutate(
                    { eventId, status },
                    { onError: (e) => toast.error(e instanceof Error ? e.message : 'Falha ao atualizar') }
                  )
                }
                onDelete={(id) =>
                  deleteEvent.mutate(id, {
                    onError: (e) => toast.error(e instanceof Error ? e.message : 'Falha ao remover'),
                  })
                }
              />
            </div>
          )}
        </div>
      </SheetContent>
    </Sheet>
  );
}

function sourceIcon(source: string) {
  return source === 'ai' ? (
    <Bot className="h-3.5 w-3.5" />
  ) : source === 'system' ? (
    <Cake className="h-3.5 w-3.5" />
  ) : (
    <User className="h-3.5 w-3.5" />
  );
}

function formatDateBR(iso?: string | null): string {
  if (!iso) return '';
  const d = new Date(iso);
  if (Number.isNaN(d.getTime())) return '';
  return d.toLocaleDateString('pt-BR', { day: '2-digit', month: '2-digit' });
}

function PeopleSection({
  people,
  pending,
  onAdd,
  onDelete,
}: {
  people: DossierPerson[];
  pending: boolean;
  onAdd: (name: string, relation: string, birthday?: string) => void;
  onDelete: (id: string) => void;
}) {
  const [name, setName] = useState('');
  const [relation, setRelation] = useState('');
  const [birthday, setBirthday] = useState('');

  const submit = () => {
    if (!name.trim()) return;
    onAdd(name.trim(), relation.trim(), birthday || undefined);
    setName('');
    setRelation('');
    setBirthday('');
  };

  return (
    <section className="border-t border-border pt-4">
      <h3 className="mb-2 text-xs font-semibold uppercase tracking-wide text-muted-foreground">
        Pessoas importantes
      </h3>
      {people.length > 0 && (
        <ul className="mb-3 space-y-1.5">
          {people.map((p) => (
            <li key={p.id} className="group flex items-start gap-2 text-sm">
              <span className="mt-0.5 shrink-0 text-muted-foreground" title={p.source === 'ai' ? 'Registrado pela Lívia' : 'Registrado pela equipe'}>
                {sourceIcon(p.source)}
              </span>
              <span className="min-w-0 flex-1">
                <span className="font-medium">{p.name}</span>
                {p.relation ? <span className="text-muted-foreground"> · {p.relation}</span> : null}
                {p.birthday ? (
                  <span className="text-muted-foreground"> · 🎂 {formatDateBR(p.birthday)}</span>
                ) : null}
              </span>
              <Button
                type="button"
                size="icon"
                variant="ghost"
                className="h-6 w-6 shrink-0 opacity-0 transition-opacity group-hover:opacity-100"
                onClick={() => onDelete(p.id)}
              >
                <Trash2 className="h-3 w-3" />
              </Button>
            </li>
          ))}
        </ul>
      )}
      <div className="space-y-2">
        <div className="flex items-center gap-1.5">
          <Input value={name} onChange={(e) => setName(e.target.value)} placeholder="Nome" className="h-8 text-sm" />
          <Input value={relation} onChange={(e) => setRelation(e.target.value)} placeholder="Relação (filha, esposo…)" className="h-8 text-sm" />
        </div>
        <div className="flex items-center gap-1.5">
          <Input
            type="date"
            value={birthday}
            onChange={(e) => setBirthday(e.target.value)}
            className="h-8 text-sm"
            title="Aniversário (opcional)"
          />
          <Button type="button" size="icon" className="h-8 w-8 shrink-0" disabled={pending} onClick={submit}>
            <Plus className="h-4 w-4" />
          </Button>
        </div>
      </div>
    </section>
  );
}

function EventsSection({
  events,
  pending,
  onAdd,
  onStatus,
  onDelete,
}: {
  events: DossierEvent[];
  pending: boolean;
  onAdd: (type: string, title: string, date: string, recurring: boolean) => void;
  onStatus: (eventId: string, status: 'done' | 'dismissed') => void;
  onDelete: (id: string) => void;
}) {
  const [type, setType] = useState('birthday');
  const [title, setTitle] = useState('');
  const [date, setDate] = useState('');

  const submit = () => {
    if (!title.trim() || !date) return;
    onAdd(type, title.trim(), date, type === 'birthday');
    setTitle('');
    setDate('');
  };

  return (
    <section className="border-t border-border pt-4">
      <h3 className="mb-2 text-xs font-semibold uppercase tracking-wide text-muted-foreground">
        Datas e eventos
      </h3>
      {events.length > 0 && (
        <ul className="mb-3 space-y-1.5">
          {events.map((e) => (
            <li key={e.id} className="group flex items-start gap-2 text-sm">
              <span className="mt-0.5 shrink-0 text-muted-foreground">
                <Gift className="h-3.5 w-3.5" />
              </span>
              <span className="min-w-0 flex-1">
                <span className="font-medium">{e.title}</span>
                <span className="text-muted-foreground">
                  {' '}
                  · {DOSSIER_EVENT_TYPE_LABELS[e.type] ?? e.type} · {formatDateBR(e.eventDate)}
                  {e.recurring ? ' (todo ano)' : ''}
                  {e.status === 'done' ? ' · ✓' : ''}
                </span>
              </span>
              <span className="flex shrink-0 gap-0.5 opacity-0 transition-opacity group-hover:opacity-100">
                {e.status !== 'done' && (
                  <Button type="button" size="icon" variant="ghost" className="h-6 w-6" title="Marcar como feito" onClick={() => onStatus(e.id, 'done')}>
                    <Check className="h-3 w-3" />
                  </Button>
                )}
                <Button type="button" size="icon" variant="ghost" className="h-6 w-6" onClick={() => onDelete(e.id)}>
                  <Trash2 className="h-3 w-3" />
                </Button>
              </span>
            </li>
          ))}
        </ul>
      )}
      <div className="space-y-2">
        <div className="flex items-center gap-1.5">
          <Select value={type} onValueChange={setType}>
            <SelectTrigger className="h-8 w-40 text-sm">
              <SelectValue />
            </SelectTrigger>
            <SelectContent>
              {Object.entries(DOSSIER_EVENT_TYPE_LABELS).map(([v, label]) => (
                <SelectItem key={v} value={v} className="text-sm">
                  {label}
                </SelectItem>
              ))}
            </SelectContent>
          </Select>
          <Input type="date" value={date} onChange={(e) => setDate(e.target.value)} className="h-8 text-sm" />
        </div>
        <div className="flex items-center gap-1.5">
          <Input
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            placeholder="Ex.: Formatura da filha; aniversário"
            className="h-8 text-sm"
            onKeyDown={(e) => {
              if (e.key === 'Enter') submit();
            }}
          />
          <Button type="button" size="icon" className="h-8 w-8 shrink-0" disabled={pending} onClick={submit}>
            <Plus className="h-4 w-4" />
          </Button>
        </div>
      </div>
    </section>
  );
}

function FactRow({
  fact,
  onSave,
  onDelete,
}: {
  fact: DossierFact;
  onSave: (value: string) => void;
  onDelete: () => void;
}) {
  const [editing, setEditing] = useState(false);
  const [value, setValue] = useState(fact.value);

  if (editing) {
    return (
      <li className="flex items-center gap-1.5">
        <Input
          value={value}
          onChange={(e) => setValue(e.target.value)}
          className="h-8 text-sm"
          autoFocus
          onKeyDown={(e) => {
            if (e.key === 'Enter') {
              setEditing(false);
              if (value.trim() && value !== fact.value) onSave(value.trim());
            }
            if (e.key === 'Escape') {
              setValue(fact.value);
              setEditing(false);
            }
          }}
        />
        <Button
          type="button"
          size="icon"
          variant="ghost"
          className="h-7 w-7 shrink-0"
          onClick={() => {
            setEditing(false);
            if (value.trim() && value !== fact.value) onSave(value.trim());
          }}
        >
          <Check className="h-3.5 w-3.5" />
        </Button>
        <Button
          type="button"
          size="icon"
          variant="ghost"
          className="h-7 w-7 shrink-0"
          onClick={() => {
            setValue(fact.value);
            setEditing(false);
          }}
        >
          <X className="h-3.5 w-3.5" />
        </Button>
      </li>
    );
  }

  return (
    <li className="group flex items-start gap-2 text-sm">
      <span
        title={fact.source === 'ai' ? 'Registrado pela Lívia' : 'Registrado pela equipe'}
        className="mt-0.5 shrink-0 text-muted-foreground"
      >
        {fact.source === 'ai' ? <Bot className="h-3.5 w-3.5" /> : <User className="h-3.5 w-3.5" />}
      </span>
      <span className="min-w-0 flex-1">
        <span className="text-muted-foreground">{fact.label}: </span>
        <span className="font-medium">{fact.value}</span>
      </span>
      <span className="flex shrink-0 gap-0.5 opacity-0 transition-opacity group-hover:opacity-100">
        <Button
          type="button"
          size="icon"
          variant="ghost"
          className="h-6 w-6"
          onClick={() => {
            setValue(fact.value);
            setEditing(true);
          }}
        >
          <Pencil className="h-3 w-3" />
        </Button>
        <Button type="button" size="icon" variant="ghost" className="h-6 w-6" onClick={onDelete}>
          <Trash2 className="h-3 w-3" />
        </Button>
      </span>
    </li>
  );
}

function AddFactForm({
  pending,
  onAdd,
}: {
  pending: boolean;
  onAdd: (category: string, value: string) => void;
}) {
  const [category, setCategory] = useState('relacionamento');
  const [value, setValue] = useState('');

  const submit = () => {
    const v = value.trim();
    if (!v) return;
    onAdd(category, v);
    setValue('');
  };

  return (
    <section className="border-t border-border pt-4">
      <h3 className="mb-2 text-xs font-semibold uppercase tracking-wide text-muted-foreground">
        Anotar algo
      </h3>
      <div className="space-y-2">
        <Select value={category} onValueChange={setCategory}>
          <SelectTrigger className="h-8 text-sm">
            <SelectValue />
          </SelectTrigger>
          <SelectContent>
            {CATEGORY_ORDER.map((c) => (
              <SelectItem key={c} value={c} className="text-sm">
                {DOSSIER_CATEGORY_LABELS[c]}
              </SelectItem>
            ))}
          </SelectContent>
        </Select>
        <div className="flex items-center gap-1.5">
          <Input
            value={value}
            onChange={(e) => setValue(e.target.value)}
            placeholder="ex.: prefere ser chamado de Dr. João; veio por indicação da filha"
            className="h-8 text-sm"
            onKeyDown={(e) => {
              if (e.key === 'Enter') submit();
            }}
          />
          <Button type="button" size="icon" className="h-8 w-8 shrink-0" disabled={pending} onClick={submit}>
            <Plus className="h-4 w-4" />
          </Button>
        </div>
        <p className="text-[11px] text-muted-foreground">
          Só informações sociais (preferências, família, como chegou). Nada clínico vai aqui.
        </p>
      </div>
    </section>
  );
}
