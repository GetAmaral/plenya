'use client';

/**
 * /configuracoes/continuum-templates — lista de templates de programa Continuum.
 * Admin/manager cria, edita, clona ou arquiva. Cada template é o "esqueleto"
 * de um programa (Semestral, Anual, Trimestral...) com N marcos por semana.
 */
import { useState } from 'react';
import Link from 'next/link';
import { toast } from 'sonner';
import { Plus, Copy, Pencil, Trash2, Loader2 } from 'lucide-react';

import { useRequireAuth } from '@/lib/use-auth';
import { useAuthStore, isGranted } from '@/lib/auth-store';
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Badge } from '@/components/ui/badge';
import { Input } from '@/components/ui/input';
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog';
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from '@/components/ui/alert-dialog';
import { PageHeader } from '@/components/layout/page-header';
import {
  useContinuumTemplates,
  useCloneContinuumTemplate,
  useDeleteContinuumTemplate,
  useCreateContinuumTemplate,
  type ContinuumTemplate,
} from '@/lib/api/continuum-api';

export default function ContinuumTemplatesPage() {
  useRequireAuth();
  const { user } = useAuthStore();
  const canManage = isGranted(user, 'admin') || isGranted(user, 'manager');

  const { data: templates = [], isLoading } = useContinuumTemplates(true);
  const clone = useCloneContinuumTemplate();
  const remove = useDeleteContinuumTemplate();
  const create = useCreateContinuumTemplate();

  const [cloneSource, setCloneSource] = useState<ContinuumTemplate | null>(null);
  const [cloneName, setCloneName] = useState('');
  const [confirmDelete, setConfirmDelete] = useState<ContinuumTemplate | null>(null);
  const [createOpen, setCreateOpen] = useState(false);
  const [newName, setNewName] = useState('');
  const [newWeeks, setNewWeeks] = useState(26);

  const handleClone = async () => {
    if (!cloneSource || !cloneName.trim()) return;
    try {
      await clone.mutateAsync({ id: cloneSource.id, name: cloneName.trim() });
      toast.success(`Template "${cloneName}" criado a partir de "${cloneSource.name}"`);
      setCloneSource(null);
      setCloneName('');
    } catch (e: any) {
      toast.error(e?.response?.data?.error ?? 'Falha ao clonar');
    }
  };

  const handleDelete = async () => {
    if (!confirmDelete) return;
    try {
      await remove.mutateAsync(confirmDelete.id);
      toast.success('Template removido');
      setConfirmDelete(null);
    } catch (e: any) {
      toast.error(e?.response?.data?.error ?? 'Falha ao remover');
    }
  };

  const handleCreateBlank = async () => {
    if (!newName.trim() || newWeeks < 1) return;
    try {
      const created = await create.mutateAsync({
        name: newName.trim(),
        durationWeeks: newWeeks,
        items: [],
      });
      toast.success('Template criado. Edite para adicionar items.');
      setCreateOpen(false);
      setNewName('');
      setNewWeeks(26);
      // Redireciona pra editor
      window.location.href = `/configuracoes/continuum-templates/${created.id}`;
    } catch (e: any) {
      toast.error(e?.response?.data?.error ?? 'Falha ao criar');
    }
  };

  return (
    <div className="container mx-auto space-y-4 py-6">
      <PageHeader
        breadcrumbs={[{ label: 'Configurações' }, { label: 'Templates Continuum' }]}
        title="Templates Continuum"
        description="Cada template define a jornada — duração, marcos semanais, reavaliações trimestrais e Boxes. Inscrição de paciente faz snapshot e segue trajetória própria."
        actions={
          canManage
            ? [
                {
                  label: 'Novo template',
                  icon: <Plus className="h-4 w-4" />,
                  onClick: () => setCreateOpen(true),
                  variant: 'default',
                },
              ]
            : []
        }
      />

      {isLoading ? (
        <Card className="flex h-40 items-center justify-center">
          <Loader2 className="h-5 w-5 animate-spin text-muted-foreground" />
        </Card>
      ) : templates.length === 0 ? (
        <Card>
          <CardContent className="py-10 text-center text-sm text-muted-foreground">
            Nenhum template cadastrado.
          </CardContent>
        </Card>
      ) : (
        <div className="grid gap-3">
          {templates.map((t: ContinuumTemplate) => (
            <Card key={t.id}>
              <CardHeader className="flex flex-row items-start justify-between gap-4 space-y-0">
                <div className="space-y-1">
                  <div className="flex items-center gap-2">
                    <CardTitle className="text-lg">{t.name}</CardTitle>
                    {t.status === 'archived' && <Badge variant="outline">Arquivado</Badge>}
                    <Badge variant="secondary">{t.durationWeeks} sem</Badge>
                  </div>
                  {t.description && <CardDescription>{t.description}</CardDescription>}
                </div>
                <div className="flex shrink-0 gap-1">
                  <Link href={`/configuracoes/continuum-templates/${t.id}`}>
                    <Button variant="outline" size="sm">
                      <Pencil className="mr-1 h-3.5 w-3.5" />
                      Editar
                    </Button>
                  </Link>
                  {canManage && (
                    <>
                      <Button
                        variant="outline"
                        size="sm"
                        onClick={() => {
                          setCloneSource(t);
                          setCloneName(`${t.name} (cópia)`);
                        }}
                      >
                        <Copy className="mr-1 h-3.5 w-3.5" />
                        Clonar
                      </Button>
                      <Button
                        variant="outline"
                        size="sm"
                        onClick={() => setConfirmDelete(t)}
                      >
                        <Trash2 className="h-3.5 w-3.5 text-destructive" />
                      </Button>
                    </>
                  )}
                </div>
              </CardHeader>
            </Card>
          ))}
        </div>
      )}

      {/* Clone dialog */}
      <Dialog open={!!cloneSource} onOpenChange={(v) => !v && setCloneSource(null)}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Clonar template</DialogTitle>
            <DialogDescription>
              Cria um novo template copiando todos os items de "{cloneSource?.name}".
            </DialogDescription>
          </DialogHeader>
          <Input
            value={cloneName}
            onChange={(e) => setCloneName(e.target.value)}
            placeholder="Nome do novo template"
          />
          <DialogFooter>
            <Button variant="outline" onClick={() => setCloneSource(null)}>
              Cancelar
            </Button>
            <Button onClick={handleClone} disabled={clone.isPending || !cloneName.trim()}>
              {clone.isPending ? <Loader2 className="h-4 w-4 animate-spin" /> : 'Clonar'}
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>

      {/* Create blank dialog */}
      <Dialog open={createOpen} onOpenChange={setCreateOpen}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Novo template</DialogTitle>
            <DialogDescription>
              Cria um template em branco. Você adiciona os items em seguida no editor.
            </DialogDescription>
          </DialogHeader>
          <div className="space-y-3">
            <div>
              <label className="mb-1 block text-xs font-medium">Nome</label>
              <Input
                value={newName}
                onChange={(e) => setNewName(e.target.value)}
                placeholder="Ex: Continuum Trimestral"
              />
            </div>
            <div>
              <label className="mb-1 block text-xs font-medium">Duração (semanas)</label>
              <Input
                type="number"
                min={1}
                value={newWeeks}
                onChange={(e) => setNewWeeks(Number(e.target.value) || 0)}
              />
            </div>
          </div>
          <DialogFooter>
            <Button variant="outline" onClick={() => setCreateOpen(false)}>
              Cancelar
            </Button>
            <Button onClick={handleCreateBlank} disabled={create.isPending || !newName.trim()}>
              {create.isPending ? <Loader2 className="h-4 w-4 animate-spin" /> : 'Criar'}
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>

      {/* Delete confirm */}
      <AlertDialog open={!!confirmDelete} onOpenChange={(v) => !v && setConfirmDelete(null)}>
        <AlertDialogContent>
          <AlertDialogHeader>
            <AlertDialogTitle>Remover template?</AlertDialogTitle>
            <AlertDialogDescription>
              "{confirmDelete?.name}" será arquivado (soft delete). Inscrições já criadas a
              partir dele continuam funcionando — usam snapshot.
            </AlertDialogDescription>
          </AlertDialogHeader>
          <AlertDialogFooter>
            <AlertDialogCancel>Cancelar</AlertDialogCancel>
            <AlertDialogAction onClick={handleDelete}>Remover</AlertDialogAction>
          </AlertDialogFooter>
        </AlertDialogContent>
      </AlertDialog>
    </div>
  );
}
