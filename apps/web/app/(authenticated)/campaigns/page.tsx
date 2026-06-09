'use client';

import Link from 'next/link';
import { useState } from 'react';
import { Plus, Archive, Trash2, ExternalLink, QrCode, Copy, Check } from 'lucide-react';
import { format } from 'date-fns';
import { ptBR } from 'date-fns/locale';

import { useRequireAuth } from '@/lib/use-auth';
import { safeUrl } from '@/lib/security';
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Badge } from '@/components/ui/badge';
import { Skeleton } from '@/components/ui/skeleton';
import { Switch } from '@/components/ui/switch';
import {
  useCampaigns,
  useArchiveCampaign,
  useDeleteCampaign,
  type Campaign,
} from '@/lib/api/campaigns-api';
import { toast } from 'sonner';

export default function CampaignsPage() {
  useRequireAuth();
  const [includeArchived, setIncludeArchived] = useState(false);
  const { data: campaigns, isLoading } = useCampaigns(includeArchived);
  const archive = useArchiveCampaign();
  const del = useDeleteCampaign();
  const [copiedId, setCopiedId] = useState<string | null>(null);

  const copyURL = async (c: Campaign) => {
    try {
      await navigator.clipboard.writeText(c.url);
      setCopiedId(c.id);
      toast.success('URL copiada');
      setTimeout(() => setCopiedId(null), 2000);
    } catch {
      toast.error('Falha ao copiar');
    }
  };

  return (
    <div className="space-y-6">
      <div className="flex items-start justify-between gap-3 flex-wrap">
        <div>
          <h1 className="text-2xl font-semibold">Campanhas</h1>
          <p className="text-sm text-muted-foreground mt-1">
            Atribui ações de marketing (UTM + QR) e mede sessões e leads gerados.
          </p>
        </div>
        <Link href="/campaigns/new">
          <Button>
            <Plus className="mr-2 h-4 w-4" /> Nova campanha
          </Button>
        </Link>
      </div>

      <div className="flex items-center gap-3 text-sm">
        <Switch
          id="incl-archived"
          checked={includeArchived}
          onCheckedChange={setIncludeArchived}
        />
        <label htmlFor="incl-archived" className="text-muted-foreground cursor-pointer">
          Mostrar arquivadas
        </label>
      </div>

      {isLoading && (
        <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
          <Skeleton className="h-44" />
          <Skeleton className="h-44" />
          <Skeleton className="h-44" />
        </div>
      )}

      {!isLoading && campaigns && campaigns.length === 0 && (
        <Card>
          <CardContent className="py-12 text-center text-muted-foreground">
            Nenhuma campanha ainda. Crie a primeira pra começar a rastrear leads por origem.
          </CardContent>
        </Card>
      )}

      {!isLoading && campaigns && campaigns.length > 0 && (
        <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
          {campaigns.map((c) => (
            <Card key={c.id} className={c.status === 'archived' ? 'opacity-60' : ''}>
              <CardHeader className="pb-3">
                <div className="flex items-start justify-between gap-2">
                  <CardTitle className="text-base leading-tight">
                    <Link href={`/campaigns/${c.id}`} className="hover:underline">
                      {c.name}
                    </Link>
                  </CardTitle>
                  <Badge variant={c.status === 'active' ? 'default' : 'secondary'}>
                    {c.status === 'active' ? 'ativa' : 'arquivada'}
                  </Badge>
                </div>
                <p className="text-xs text-muted-foreground font-mono">{c.slug}</p>
              </CardHeader>
              <CardContent className="space-y-3 text-sm">
                <div className="flex flex-wrap gap-1.5 text-xs">
                  <Badge variant="outline">src: {c.utmSource}</Badge>
                  <Badge variant="outline">med: {c.utmMedium}</Badge>
                  <Badge variant="outline">cmp: {c.utmCampaign}</Badge>
                </div>
                <p className="text-xs text-muted-foreground truncate" title={c.url}>
                  {c.url}
                </p>
                <div className="flex items-center gap-3 text-xs text-muted-foreground pt-1 border-t">
                  <span>📊 {c.stats?.sessionsCount ?? 0} sessões</span>
                  <span>👤 {c.stats?.leadsCount ?? 0} leads</span>
                  <span className="ml-auto">
                    {format(new Date(c.createdAt), "dd MMM ''yy", { locale: ptBR })}
                  </span>
                </div>
                <div className="flex gap-1.5 pt-2">
                  <Button
                    size="sm"
                    variant="ghost"
                    onClick={() => copyURL(c)}
                    title="Copiar URL"
                  >
                    {copiedId === c.id ? (
                      <Check className="h-4 w-4 text-green-600" />
                    ) : (
                      <Copy className="h-4 w-4" />
                    )}
                  </Button>
                  <a href={safeUrl(c.url)} target="_blank" rel="noreferrer">
                    <Button size="sm" variant="ghost" title="Abrir landing">
                      <ExternalLink className="h-4 w-4" />
                    </Button>
                  </a>
                  <Link href={`/campaigns/${c.id}`}>
                    <Button size="sm" variant="ghost" title="QR + detalhes">
                      <QrCode className="h-4 w-4" />
                    </Button>
                  </Link>
                  {c.status === 'active' && (
                    <Button
                      size="sm"
                      variant="ghost"
                      onClick={() => {
                        if (confirm(`Arquivar "${c.name}"?`)) {
                          archive.mutate(c.id, {
                            onSuccess: () => toast.success('Arquivada'),
                          });
                        }
                      }}
                      title="Arquivar"
                    >
                      <Archive className="h-4 w-4" />
                    </Button>
                  )}
                  <Button
                    size="sm"
                    variant="ghost"
                    onClick={() => {
                      if (
                        confirm(
                          `Excluir "${c.name}" definitivamente? Leads atribuídos preservam o utm_campaign.`,
                        )
                      ) {
                        del.mutate(c.id, { onSuccess: () => toast.success('Excluída') });
                      }
                    }}
                    title="Excluir"
                  >
                    <Trash2 className="h-4 w-4 text-destructive" />
                  </Button>
                </div>
              </CardContent>
            </Card>
          ))}
        </div>
      )}
    </div>
  );
}
