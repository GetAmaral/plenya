'use client';

/**
 * DossierPanel — visão 360 SOCIAL de uma pessoa (memória de relacionamento da Lívia).
 * Resumo rolante + fatos atômicos (key/value) agrupados por categoria, com proveniência
 * (IA vs equipe) e edição manual. Só dado social — nada clínico (LGPD/CFM). O conteúdo é
 * alimentado pela IA (job) e pela equipe (aqui). Plano: docs/emr/plano-livia-memoria-dossie-360.md.
 */

import { useMemo, useState } from 'react';
import { Bot, Check, Pencil, Plus, Trash2, User, X } from 'lucide-react';
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
  useAddDossierFact,
  useDeleteDossierFact,
  useDossier,
  useUpdateDossierFact,
  type ConversationOwnerType,
  type DossierFact,
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
          <SheetTitle className="flex items-center gap-2">
            Dossiê
            <Badge variant="outline" className="text-[10px] uppercase tracking-wide">
              {ownerType === 'patient' ? 'Paciente' : 'Lead'}
            </Badge>
          </SheetTitle>
          <SheetDescription className="text-xs">
            O que a equipe e a Lívia sabem sobre {name}. Apenas informações sociais, nada clínico.
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
            </div>
          )}
        </div>
      </SheetContent>
    </Sheet>
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
        <Button type="button" size="icon" variant="ghost" className="h-6 w-6" onClick={() => setEditing(true)}>
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
