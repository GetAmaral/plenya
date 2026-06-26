'use client';

/**
 * /configuracoes/whatsapp-templates
 *
 * Catálogo dos templates WhatsApp: status sincronizado da Meta + metadados nossos (propósito,
 * variáveis, cabeamento, toggle). Admin/manager. Backend: /api/v1/admin/whatsapp-templates.
 */
import { useMemo, useState } from 'react';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { RefreshCw, Loader2, Pencil } from 'lucide-react';
import { toast } from 'sonner';

import { useRequireAuth } from '@/lib/use-auth';
import { apiClient } from '@/lib/api-client';
import { formatDateTime } from '@/lib/format-date';
import { PageHeader } from '@/components/layout/page-header';
import { Card, CardContent } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Badge } from '@/components/ui/badge';
import { Switch } from '@/components/ui/switch';
import { Skeleton } from '@/components/ui/skeleton';
import { Input } from '@/components/ui/input';
import { Textarea } from '@/components/ui/textarea';
import { Label } from '@/components/ui/label';
import {
  Dialog, DialogContent, DialogHeader, DialogTitle, DialogFooter,
} from '@/components/ui/dialog';

interface TemplateVar { index: number; label: string; example: string }
interface WaTemplate {
  id: string; name: string; language: string; category: string; status: string;
  bodyText: string; variableCount: number; variables: TemplateVar[];
  purpose: string; wiringNotes: string; enabled: boolean; lastSyncedAt?: string;
}

function statusVariant(s: string): 'default' | 'secondary' | 'destructive' | 'outline' {
  if (s === 'APPROVED') return 'default';
  if (s === 'PENDING' || s === 'IN_APPEAL' || s === 'PENDING_DELETION') return 'secondary';
  if (s === 'REJECTED' || s === 'DISABLED' || s === 'PAUSED') return 'destructive';
  return 'outline';
}

export default function WhatsAppTemplatesPage() {
  useRequireAuth();
  const qc = useQueryClient();
  const [editing, setEditing] = useState<WaTemplate | null>(null);

  const { data, isLoading } = useQuery({
    queryKey: ['wa-templates'],
    queryFn: () => apiClient.get<{ data: WaTemplate[] }>('/api/v1/admin/whatsapp-templates'),
  });
  const templates = useMemo(() => data?.data ?? [], [data]);

  const sync = useMutation({
    mutationFn: () => apiClient.post<{ synced: number }>('/api/v1/admin/whatsapp-templates/sync'),
    onSuccess: (r) => { toast.success(`${r.synced} templates sincronizados da Meta`); qc.invalidateQueries({ queryKey: ['wa-templates'] }); },
    onError: (e: unknown) => toast.error(e instanceof Error ? e.message : 'Falha ao sincronizar'),
  });

  const patch = useMutation({
    mutationFn: (p: { id: string; body: Partial<WaTemplate> }) =>
      apiClient.patch<WaTemplate>(`/api/v1/admin/whatsapp-templates/${p.id}`, p.body),
    onSuccess: () => { qc.invalidateQueries({ queryKey: ['wa-templates'] }); },
    onError: (e: unknown) => toast.error(e instanceof Error ? e.message : 'Falha ao salvar'),
  });

  return (
    <div className="container mx-auto py-8 space-y-6">
      <PageHeader
        breadcrumbs={[{ label: 'Configurações' }, { label: 'Templates WhatsApp' }]}
        title="Templates WhatsApp"
        description="Catálogo dos templates aprovados/pendentes na Meta + como cada um é usado."
        actions={[]}
      />

      <div className="flex justify-end">
        <Button onClick={() => sync.mutate()} disabled={sync.isPending} variant="outline">
          {sync.isPending ? <Loader2 className="h-4 w-4 animate-spin" /> : <RefreshCw className="h-4 w-4" />}
          Sincronizar Meta
        </Button>
      </div>

      {isLoading ? (
        <div className="space-y-3">{[0, 1, 2].map((i) => <Skeleton key={i} className="h-28 w-full" />)}</div>
      ) : templates.length === 0 ? (
        <Card><CardContent className="py-10 text-center text-muted-foreground">
          Nenhum template no catálogo. Clique em "Sincronizar Meta" para puxar da WhatsApp Business.
        </CardContent></Card>
      ) : (
        <div className="space-y-3">
          {templates.map((t) => (
            <Card key={t.id}>
              <CardContent className="py-4 space-y-3">
                <div className="flex items-start justify-between gap-3">
                  <div className="space-y-1 min-w-0">
                    <div className="flex items-center gap-2 flex-wrap">
                      <span className="font-mono font-semibold">{t.name}</span>
                      <Badge variant={statusVariant(t.status)}>{t.status}</Badge>
                      <Badge variant="outline">{t.category || '—'}</Badge>
                      <Badge variant="outline">{t.language}</Badge>
                      <Badge variant="outline">{t.variableCount} var</Badge>
                    </div>
                    {t.purpose && <p className="text-sm text-muted-foreground">{t.purpose}</p>}
                  </div>
                  <div className="flex items-center gap-3 shrink-0">
                    <div className="flex items-center gap-2" title="Habilitado para envio">
                      <Switch checked={t.enabled} onCheckedChange={(v) => patch.mutate({ id: t.id, body: { enabled: v } })} />
                      <span className="text-xs text-muted-foreground">{t.enabled ? 'on' : 'off'}</span>
                    </div>
                    <Button size="sm" variant="ghost" onClick={() => setEditing(t)}><Pencil className="h-4 w-4" /></Button>
                  </div>
                </div>

                <p className="text-sm whitespace-pre-wrap rounded bg-muted/50 p-3">{t.bodyText || '—'}</p>

                {t.variables?.length > 0 && (
                  <div className="text-xs text-muted-foreground space-y-0.5">
                    {t.variables.map((v) => (
                      <div key={v.index}>
                        <span className="font-mono">{`{{${v.index}}}`}</span>
                        {v.label ? ` = ${v.label}` : ' = (sem rótulo)'}
                        {v.example ? ` · ex.: ${v.example}` : ''}
                      </div>
                    ))}
                  </div>
                )}

                {t.wiringNotes && <p className="text-xs text-muted-foreground border-l-2 pl-2 italic">{t.wiringNotes}</p>}
                {t.lastSyncedAt && <p className="text-[11px] text-muted-foreground">Sincronizado: {formatDateTime(t.lastSyncedAt)}</p>}
              </CardContent>
            </Card>
          ))}
        </div>
      )}

      {editing && (
        <EditDialog
          template={editing}
          onClose={() => setEditing(null)}
          onSave={(body) => { patch.mutate({ id: editing.id, body }); setEditing(null); }}
          saving={patch.isPending}
        />
      )}
    </div>
  );
}

function EditDialog({ template, onClose, onSave, saving }: {
  template: WaTemplate; onClose: () => void; onSave: (b: Partial<WaTemplate>) => void; saving: boolean;
}) {
  const [purpose, setPurpose] = useState(template.purpose ?? '');
  const [wiringNotes, setWiringNotes] = useState(template.wiringNotes ?? '');
  const [vars, setVars] = useState<TemplateVar[]>(
    template.variables?.length
      ? template.variables
      : Array.from({ length: template.variableCount }, (_, i) => ({ index: i + 1, label: '', example: '' })),
  );

  return (
    <Dialog open onOpenChange={(o) => { if (!o) onClose(); }}>
      <DialogContent className="max-w-lg">
        <DialogHeader><DialogTitle className="font-mono">{template.name}</DialogTitle></DialogHeader>
        <div className="space-y-4">
          <div className="space-y-1">
            <Label>Propósito</Label>
            <Textarea value={purpose} onChange={(e) => setPurpose(e.target.value)} rows={2} placeholder="Para que serve este template" />
          </div>
          <div className="space-y-1">
            <Label>Cabeamento / como usar</Label>
            <Textarea value={wiringNotes} onChange={(e) => setWiringNotes(e.target.value)} rows={2} placeholder="Onde/como é disparado no sistema" />
          </div>
          {vars.length > 0 && (
            <div className="space-y-2">
              <Label>Variáveis</Label>
              {vars.map((v, i) => (
                <div key={v.index} className="flex items-center gap-2">
                  <span className="font-mono text-sm w-10">{`{{${v.index}}}`}</span>
                  <Input placeholder="rótulo (ex.: nome)" value={v.label}
                    onChange={(e) => setVars(vars.map((x, j) => j === i ? { ...x, label: e.target.value } : x))} />
                  <Input placeholder="exemplo" value={v.example}
                    onChange={(e) => setVars(vars.map((x, j) => j === i ? { ...x, example: e.target.value } : x))} />
                </div>
              ))}
            </div>
          )}
        </div>
        <DialogFooter>
          <Button variant="ghost" onClick={onClose}>Cancelar</Button>
          <Button onClick={() => onSave({ purpose, wiringNotes, variables: vars })} disabled={saving}>
            {saving && <Loader2 className="h-4 w-4 animate-spin" />} Salvar
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}
