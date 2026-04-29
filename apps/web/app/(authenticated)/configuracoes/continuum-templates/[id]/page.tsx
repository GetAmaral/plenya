'use client';

/**
 * /configuracoes/continuum-templates/[id] — editor de template Continuum.
 *
 * Tabela de items ordenada por (weekOffset, position). Cada linha define
 * um marco do programa: tipo, especialidade (se appointment), título,
 * semana, dia da semana, dias até virar atrasado, e box template (se box).
 *
 * Salvar substitui todos os items por inteiro (semântica replace-all).
 */
import { useEffect, useMemo, useState } from 'react';
import { useParams, useRouter } from 'next/navigation';
import { toast } from 'sonner';
import { Loader2, Plus, Trash2, ArrowDown, ArrowUp, Save } from 'lucide-react';

import { useRequireAuth } from '@/lib/use-auth';
import { useAuthStore, isGranted } from '@/lib/auth-store';
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Input } from '@/components/ui/input';
import { Textarea } from '@/components/ui/textarea';
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';
import { PageHeader } from '@/components/layout/page-header';
import {
  useContinuumTemplate,
  useContinuumBoxTemplates,
  useUpdateContinuumTemplate,
  type ContinuumTemplateItem,
  type ContinuumItemType,
  type ContinuumItemSpecialty,
  ITEM_TYPE_LABELS,
  SPECIALTY_LABELS,
} from '@/lib/api/continuum-api';

const WEEKDAYS = ['Dom', 'Seg', 'Ter', 'Qua', 'Qui', 'Sex', 'Sáb'];

export default function ContinuumTemplateEditorPage() {
  const params = useParams<{ id: string }>();
  useRequireAuth();
  const router = useRouter();
  const { user } = useAuthStore();
  const canManage = isGranted(user, 'admin') || isGranted(user, 'manager');

  const { data: template, isLoading } = useContinuumTemplate(params.id);
  const { data: boxTemplates = [] } = useContinuumBoxTemplates();
  const update = useUpdateContinuumTemplate(params.id);

  const [name, setName] = useState('');
  const [description, setDescription] = useState('');
  const [durationWeeks, setDurationWeeks] = useState(26);
  const [items, setItems] = useState<ContinuumTemplateItem[]>([]);

  // Hidrata estado quando query carrega.
  useEffect(() => {
    if (!template) return;
    setName(template.name);
    setDescription(template.description ?? '');
    setDurationWeeks(template.durationWeeks);
    setItems(
      [...(template.items ?? [])].sort(
        (a, b) => a.weekOffset - b.weekOffset || a.position - b.position,
      ),
    );
  }, [template]);

  const sortedItems = useMemo(
    () =>
      [...items].sort((a, b) => a.weekOffset - b.weekOffset || a.position - b.position),
    [items],
  );

  const addItem = () => {
    setItems((prev) => [
      ...prev,
      {
        type: 'appointment',
        specialty: 'doctor',
        title: 'Novo marco',
        description: '',
        weekOffset: 0,
        expectedOffsetDays: 0,
        lateAfterDays: 7,
        position: prev.length,
      },
    ]);
  };

  const removeItem = (index: number) => {
    setItems((prev) => prev.filter((_, i) => i !== index));
  };

  const updateItem = (index: number, patch: Partial<ContinuumTemplateItem>) => {
    setItems((prev) => prev.map((it, i) => (i === index ? { ...it, ...patch } : it)));
  };

  const moveItem = (index: number, dir: -1 | 1) => {
    setItems((prev) => {
      const next = [...prev];
      const target = index + dir;
      if (target < 0 || target >= next.length) return prev;
      [next[index], next[target]] = [next[target], next[index]];
      return next;
    });
  };

  const handleSave = async () => {
    if (!name.trim()) {
      toast.error('Nome é obrigatório');
      return;
    }
    try {
      await update.mutateAsync({
        name: name.trim(),
        description,
        durationWeeks,
        items: items.map((it, i) => ({ ...it, position: i })),
      });
      toast.success('Template salvo');
    } catch (e: any) {
      toast.error(e?.response?.data?.error ?? 'Falha ao salvar');
    }
  };

  if (isLoading) {
    return (
      <div className="flex h-96 items-center justify-center">
        <Loader2 className="h-6 w-6 animate-spin text-muted-foreground" />
      </div>
    );
  }

  if (!template) {
    return <div className="container mx-auto py-6">Template não encontrado.</div>;
  }

  return (
    <div className="container mx-auto space-y-4 py-6">
      <PageHeader
        breadcrumbs={[
          { label: 'Configurações' },
          { label: 'Templates Continuum', href: '/configuracoes/continuum-templates' },
          { label: template.name },
        ]}
        title={template.name}
        description={`${template.durationWeeks} semanas · ${items.length} marcos`}
        actions={
          canManage
            ? [
                {
                  label: update.isPending ? 'Salvando...' : 'Salvar',
                  icon: <Save className="h-4 w-4" />,
                  onClick: handleSave,
                  variant: 'default',
                  disabled: update.isPending,
                },
              ]
            : []
        }
      />

      {/* Header form */}
      <Card>
        <CardHeader>
          <CardTitle className="text-base">Identificação</CardTitle>
        </CardHeader>
        <CardContent className="grid gap-4 md:grid-cols-3">
          <div className="md:col-span-2">
            <label className="mb-1 block text-xs font-medium">Nome</label>
            <Input
              value={name}
              onChange={(e) => setName(e.target.value)}
              disabled={!canManage}
            />
          </div>
          <div>
            <label className="mb-1 block text-xs font-medium">Duração (semanas)</label>
            <Input
              type="number"
              min={1}
              value={durationWeeks}
              onChange={(e) => setDurationWeeks(Number(e.target.value) || 0)}
              disabled={!canManage}
            />
          </div>
          <div className="md:col-span-3">
            <label className="mb-1 block text-xs font-medium">Descrição</label>
            <Textarea
              rows={2}
              value={description}
              onChange={(e) => setDescription(e.target.value)}
              disabled={!canManage}
            />
          </div>
        </CardContent>
      </Card>

      {/* Items */}
      <Card>
        <CardHeader className="flex flex-row items-center justify-between space-y-0">
          <CardTitle className="text-base">Marcos do programa ({items.length})</CardTitle>
          {canManage && (
            <Button size="sm" onClick={addItem}>
              <Plus className="mr-1 h-4 w-4" />
              Adicionar marco
            </Button>
          )}
        </CardHeader>
        <CardContent>
          {sortedItems.length === 0 ? (
            <p className="py-8 text-center text-sm text-muted-foreground">
              Nenhum marco. Adicione o primeiro.
            </p>
          ) : (
            <div className="overflow-x-auto">
              <table className="w-full text-sm">
                <thead className="text-xs uppercase text-muted-foreground">
                  <tr className="border-b">
                    <th className="py-2 text-left">Sem</th>
                    <th className="py-2 text-left">Dia</th>
                    <th className="py-2 text-left">Tipo</th>
                    <th className="py-2 text-left">Especialidade</th>
                    <th className="py-2 text-left">Título</th>
                    <th className="py-2 text-left">Box</th>
                    <th className="py-2 text-left">Atraso (d)</th>
                    {canManage && <th className="py-2 text-right">Ações</th>}
                  </tr>
                </thead>
                <tbody>
                  {items.map((it, idx) => (
                    <tr key={idx} className="border-b last:border-0">
                      <td className="py-2 pr-2">
                        <Input
                          type="number"
                          min={0}
                          max={durationWeeks - 1}
                          className="h-8 w-16"
                          value={it.weekOffset}
                          onChange={(e) =>
                            updateItem(idx, { weekOffset: Number(e.target.value) || 0 })
                          }
                          disabled={!canManage}
                        />
                      </td>
                      <td className="py-2 pr-2">
                        <Select
                          value={String(it.expectedOffsetDays)}
                          onValueChange={(v) =>
                            updateItem(idx, { expectedOffsetDays: Number(v) })
                          }
                          disabled={!canManage}
                        >
                          <SelectTrigger className="h-8 w-20">
                            <SelectValue />
                          </SelectTrigger>
                          <SelectContent>
                            {WEEKDAYS.map((d, i) => (
                              <SelectItem key={i} value={String(i)}>
                                {d}
                              </SelectItem>
                            ))}
                          </SelectContent>
                        </Select>
                      </td>
                      <td className="py-2 pr-2">
                        <Select
                          value={it.type}
                          onValueChange={(v) =>
                            updateItem(idx, { type: v as ContinuumItemType })
                          }
                          disabled={!canManage}
                        >
                          <SelectTrigger className="h-8 w-32">
                            <SelectValue />
                          </SelectTrigger>
                          <SelectContent>
                            {Object.entries(ITEM_TYPE_LABELS).map(([k, v]) => (
                              <SelectItem key={k} value={k}>
                                {v}
                              </SelectItem>
                            ))}
                          </SelectContent>
                        </Select>
                      </td>
                      <td className="py-2 pr-2">
                        {it.type === 'appointment' ? (
                          <Select
                            value={it.specialty ?? 'doctor'}
                            onValueChange={(v) =>
                              updateItem(idx, { specialty: v as ContinuumItemSpecialty })
                            }
                            disabled={!canManage}
                          >
                            <SelectTrigger className="h-8 w-36">
                              <SelectValue />
                            </SelectTrigger>
                            <SelectContent>
                              {Object.entries(SPECIALTY_LABELS).map(([k, v]) => (
                                <SelectItem key={k} value={k}>
                                  {v}
                                </SelectItem>
                              ))}
                            </SelectContent>
                          </Select>
                        ) : (
                          <span className="text-xs text-muted-foreground">—</span>
                        )}
                      </td>
                      <td className="py-2 pr-2">
                        <Input
                          className="h-8 min-w-[180px]"
                          value={it.title}
                          onChange={(e) => updateItem(idx, { title: e.target.value })}
                          disabled={!canManage}
                        />
                      </td>
                      <td className="py-2 pr-2">
                        {it.type === 'box' ? (
                          <Select
                            value={it.boxTemplateId ?? ''}
                            onValueChange={(v) =>
                              updateItem(idx, { boxTemplateId: v || null })
                            }
                            disabled={!canManage}
                          >
                            <SelectTrigger className="h-8 w-44">
                              <SelectValue placeholder="Selecionar box" />
                            </SelectTrigger>
                            <SelectContent>
                              {boxTemplates.map((b: { id: string; name: string }) => (
                                <SelectItem key={b.id} value={b.id}>
                                  {b.name}
                                </SelectItem>
                              ))}
                            </SelectContent>
                          </Select>
                        ) : (
                          <span className="text-xs text-muted-foreground">—</span>
                        )}
                      </td>
                      <td className="py-2 pr-2">
                        <Input
                          type="number"
                          min={0}
                          className="h-8 w-16"
                          value={it.lateAfterDays}
                          onChange={(e) =>
                            updateItem(idx, { lateAfterDays: Number(e.target.value) || 0 })
                          }
                          disabled={!canManage}
                        />
                      </td>
                      {canManage && (
                        <td className="py-2">
                          <div className="flex justify-end gap-1">
                            <Button
                              variant="ghost"
                              size="icon"
                              className="h-7 w-7"
                              onClick={() => moveItem(idx, -1)}
                            >
                              <ArrowUp className="h-3.5 w-3.5" />
                            </Button>
                            <Button
                              variant="ghost"
                              size="icon"
                              className="h-7 w-7"
                              onClick={() => moveItem(idx, 1)}
                            >
                              <ArrowDown className="h-3.5 w-3.5" />
                            </Button>
                            <Button
                              variant="ghost"
                              size="icon"
                              className="h-7 w-7"
                              onClick={() => removeItem(idx)}
                            >
                              <Trash2 className="h-3.5 w-3.5 text-destructive" />
                            </Button>
                          </div>
                        </td>
                      )}
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          )}
        </CardContent>
      </Card>

      {canManage && (
        <div className="flex justify-end">
          <Button onClick={handleSave} disabled={update.isPending}>
            {update.isPending && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
            Salvar template
          </Button>
        </div>
      )}
    </div>
  );
}
