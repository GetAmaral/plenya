'use client';

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
  useContinuumBoxTemplates,
  useCloneContinuumBoxTemplate,
  useDeleteContinuumBoxTemplate,
  useCreateContinuumBoxTemplate,
  type ContinuumBoxTemplate,
} from '@/lib/api/continuum-api';

export default function ContinuumBoxTemplatesPage() {
  useRequireAuth();
  const { user } = useAuthStore();
  const canManage = isGranted(user, 'admin') || isGranted(user, 'manager');

  const { data: boxes = [], isLoading } = useContinuumBoxTemplates(true);
  const clone = useCloneContinuumBoxTemplate();
  const remove = useDeleteContinuumBoxTemplate();
  const create = useCreateContinuumBoxTemplate();

  const [cloneSource, setCloneSource] = useState<ContinuumBoxTemplate | null>(null);
  const [cloneName, setCloneName] = useState('');
  const [confirmDelete, setConfirmDelete] = useState<ContinuumBoxTemplate | null>(null);
  const [createOpen, setCreateOpen] = useState(false);
  const [newName, setNewName] = useState('');

  const handleClone = async () => {
    if (!cloneSource || !cloneName.trim()) return;
    try {
      await clone.mutateAsync({ id: cloneSource.id, name: cloneName.trim() });
      toast.success(`Box "${cloneName}" criado`);
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
      toast.success('Box removido');
      setConfirmDelete(null);
    } catch (e: any) {
      toast.error(e?.response?.data?.error ?? 'Falha ao remover');
    }
  };

  const handleCreateBlank = async () => {
    if (!newName.trim()) return;
    try {
      const created = await create.mutateAsync({ name: newName.trim() });
      toast.success('Box criado. Edite o conteúdo.');
      setCreateOpen(false);
      setNewName('');
      window.location.href = `/continuum/box-templates/${created.id}`;
    } catch (e: any) {
      toast.error(e?.response?.data?.error ?? 'Falha ao criar');
    }
  };

  return (
    <div className="container mx-auto space-y-4 py-6">
      <PageHeader
        breadcrumbs={[
          { label: 'Continuum', href: '/continuum' },
          { label: 'Templates de Box' },
        ]}
        title="Templates de Box"
        description="Caixas Plenya — boas-vindas, mensal, reavaliação. Reusados pelos templates de programa."
        actions={
          canManage
            ? [
                {
                  label: 'Novo Box',
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
      ) : boxes.length === 0 ? (
        <Card>
          <CardContent className="py-10 text-center text-sm text-muted-foreground">
            Nenhum box cadastrado.
          </CardContent>
        </Card>
      ) : (
        <div className="grid gap-3 md:grid-cols-2">
          {boxes.map((b: ContinuumBoxTemplate) => (
            <Card key={b.id}>
              <CardHeader>
                <div className="flex items-start justify-between gap-3">
                  <div className="space-y-1">
                    <div className="flex items-center gap-2">
                      <CardTitle className="text-base">{b.name}</CardTitle>
                      {b.status === 'archived' && <Badge variant="outline">Arquivado</Badge>}
                    </div>
                    {b.description && <CardDescription>{b.description}</CardDescription>}
                  </div>
                  <div className="flex shrink-0 gap-1">
                    <Link href={`/continuum/box-templates/${b.id}`}>
                      <Button variant="outline" size="sm">
                        <Pencil className="h-3.5 w-3.5" />
                      </Button>
                    </Link>
                    {canManage && (
                      <>
                        <Button
                          variant="outline"
                          size="sm"
                          onClick={() => {
                            setCloneSource(b);
                            setCloneName(`${b.name} (cópia)`);
                          }}
                        >
                          <Copy className="h-3.5 w-3.5" />
                        </Button>
                        <Button
                          variant="outline"
                          size="sm"
                          onClick={() => setConfirmDelete(b)}
                        >
                          <Trash2 className="h-3.5 w-3.5 text-destructive" />
                        </Button>
                      </>
                    )}
                  </div>
                </div>
              </CardHeader>
            </Card>
          ))}
        </div>
      )}

      <Dialog open={!!cloneSource} onOpenChange={(v) => !v && setCloneSource(null)}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Clonar box</DialogTitle>
            <DialogDescription>Cria nova cópia de "{cloneSource?.name}".</DialogDescription>
          </DialogHeader>
          <Input value={cloneName} onChange={(e) => setCloneName(e.target.value)} />
          <DialogFooter>
            <Button variant="outline" onClick={() => setCloneSource(null)}>
              Cancelar
            </Button>
            <Button onClick={handleClone} disabled={clone.isPending}>
              {clone.isPending ? <Loader2 className="h-4 w-4 animate-spin" /> : 'Clonar'}
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>

      <Dialog open={createOpen} onOpenChange={setCreateOpen}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Novo box</DialogTitle>
          </DialogHeader>
          <Input
            value={newName}
            onChange={(e) => setNewName(e.target.value)}
            placeholder="Ex: Box Reavaliação Anual"
          />
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

      <AlertDialog open={!!confirmDelete} onOpenChange={(v) => !v && setConfirmDelete(null)}>
        <AlertDialogContent>
          <AlertDialogHeader>
            <AlertDialogTitle>Remover box?</AlertDialogTitle>
            <AlertDialogDescription>
              "{confirmDelete?.name}" será arquivado.
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
