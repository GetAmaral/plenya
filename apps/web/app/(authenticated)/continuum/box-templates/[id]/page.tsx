'use client';

import { useEffect, useState } from 'react';
import { useParams } from 'next/navigation';
import { toast } from 'sonner';
import { Loader2, Save } from 'lucide-react';

import { useRequireAuth } from '@/lib/use-auth';
import { useAuthStore, isGranted } from '@/lib/auth-store';
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Input } from '@/components/ui/input';
import { Textarea } from '@/components/ui/textarea';
import { PageHeader } from '@/components/layout/page-header';
import {
  useContinuumBoxTemplate,
  useUpdateContinuumBoxTemplate,
} from '@/lib/api/continuum-api';

export default function ContinuumBoxTemplateEditorPage() {
  const params = useParams<{ id: string }>();
  useRequireAuth();
  const { user } = useAuthStore();
  const canManage = isGranted(user, 'admin') || isGranted(user, 'manager');

  const { data: box, isLoading } = useContinuumBoxTemplate(params.id);
  const update = useUpdateContinuumBoxTemplate(params.id);

  const [name, setName] = useState('');
  const [description, setDescription] = useState('');
  const [contents, setContents] = useState('');
  const [notes, setNotes] = useState('');

  useEffect(() => {
    if (!box) return;
    setName(box.name);
    setDescription(box.description ?? '');
    setContents(box.contents ?? '');
    setNotes(box.notes ?? '');
  }, [box]);

  const handleSave = async () => {
    if (!name.trim()) {
      toast.error('Nome é obrigatório');
      return;
    }
    try {
      await update.mutateAsync({ name: name.trim(), description, contents, notes });
      toast.success('Box salvo');
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
  if (!box) return <div className="container mx-auto py-6">Box não encontrado.</div>;

  return (
    <div className="container mx-auto space-y-4 py-6">
      <PageHeader
        breadcrumbs={[
          { label: 'Continuum', href: '/continuum' },
          { label: 'Templates de Box', href: '/continuum/box-templates' },
          { label: box.name },
        ]}
        title={box.name}
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

      <Card>
        <CardHeader>
          <CardTitle className="text-base">Identificação</CardTitle>
        </CardHeader>
        <CardContent className="grid gap-4">
          <div>
            <label className="mb-1 block text-xs font-medium">Nome</label>
            <Input
              value={name}
              onChange={(e) => setName(e.target.value)}
              disabled={!canManage}
            />
          </div>
          <div>
            <label className="mb-1 block text-xs font-medium">Descrição curta</label>
            <Input
              value={description}
              onChange={(e) => setDescription(e.target.value)}
              disabled={!canManage}
            />
          </div>
        </CardContent>
      </Card>

      <Card>
        <CardHeader>
          <CardTitle className="text-base">Conteúdo (Markdown)</CardTitle>
        </CardHeader>
        <CardContent>
          <Textarea
            rows={10}
            value={contents}
            onChange={(e) => setContents(e.target.value)}
            placeholder="Lista de itens que compõem o box. Pode usar Markdown."
            disabled={!canManage}
          />
        </CardContent>
      </Card>

      <Card>
        <CardHeader>
          <CardTitle className="text-base">Instruções pra logística</CardTitle>
        </CardHeader>
        <CardContent>
          <Textarea
            rows={4}
            value={notes}
            onChange={(e) => setNotes(e.target.value)}
            placeholder="Anotações internas — embalagem, alertas, prazos."
            disabled={!canManage}
          />
        </CardContent>
      </Card>

      {canManage && (
        <div className="flex justify-end">
          <Button onClick={handleSave} disabled={update.isPending}>
            {update.isPending && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
            Salvar
          </Button>
        </div>
      )}
    </div>
  );
}
