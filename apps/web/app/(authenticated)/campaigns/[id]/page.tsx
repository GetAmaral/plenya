'use client';

import Link from 'next/link';
import { useParams, useRouter } from 'next/navigation';
import { useEffect, useState } from 'react';
import { ArrowLeft, Copy, Check, Download, Archive, ExternalLink } from 'lucide-react';
import { format } from 'date-fns';
import { ptBR } from 'date-fns/locale';
import { toast } from 'sonner';

import { useRequireAuth } from '@/lib/use-auth';
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Badge } from '@/components/ui/badge';
import { Skeleton } from '@/components/ui/skeleton';
import { CampaignForm } from '@/components/campaigns/campaign-form';
import {
  useCampaign,
  useUpdateCampaign,
  useArchiveCampaign,
  useCampaignQRCode,
  type CreateCampaignInput,
} from '@/lib/api/campaigns-api';

export default function CampaignDetailPage() {
  useRequireAuth();
  const params = useParams<{ id: string }>();
  const router = useRouter();
  const id = params?.id ?? '';

  const { data: campaign, isLoading } = useCampaign(id);
  const update = useUpdateCampaign(id);
  const archive = useArchiveCampaign();
  const { data: qrURL } = useCampaignQRCode(id, 512);

  const [editing, setEditing] = useState(false);
  const [copied, setCopied] = useState(false);

  // Limpa object URL quando trocar de campanha
  useEffect(() => () => {
    if (qrURL) URL.revokeObjectURL(qrURL);
  }, [qrURL]);

  if (isLoading || !campaign) {
    return (
      <div className="space-y-4">
        <Skeleton className="h-8 w-32" />
        <Skeleton className="h-44" />
      </div>
    );
  }

  const copyURL = async () => {
    try {
      await navigator.clipboard.writeText(campaign.url);
      setCopied(true);
      toast.success('URL copiada');
      setTimeout(() => setCopied(false), 2000);
    } catch {
      toast.error('Falha ao copiar');
    }
  };

  const onSubmit = (input: CreateCampaignInput) => {
    update.mutate(input, {
      onSuccess: () => {
        toast.success('Campanha atualizada');
        setEditing(false);
      },
      onError: (err) =>
        toast.error('Erro', {
          description: err instanceof Error ? err.message : 'Tente novamente',
        }),
    });
  };

  return (
    <div className="space-y-6 max-w-4xl">
      <Link href="/campaigns">
        <Button variant="ghost" size="sm">
          <ArrowLeft className="mr-2 h-4 w-4" /> Todas as campanhas
        </Button>
      </Link>

      <div className="flex items-start justify-between gap-3 flex-wrap">
        <div>
          <h1 className="text-2xl font-semibold">{campaign.name}</h1>
          <p className="font-mono text-xs text-muted-foreground mt-1">
            {campaign.slug}{' '}
            <Badge variant={campaign.status === 'active' ? 'default' : 'secondary'} className="ml-2">
              {campaign.status === 'active' ? 'ativa' : 'arquivada'}
            </Badge>
          </p>
        </div>
        {!editing && (
          <div className="flex gap-2">
            <Button variant="outline" onClick={() => setEditing(true)}>
              Editar
            </Button>
            {campaign.status === 'active' && (
              <Button
                variant="outline"
                onClick={() => {
                  if (confirm(`Arquivar "${campaign.name}"?`)) {
                    archive.mutate(campaign.id, {
                      onSuccess: () => toast.success('Arquivada'),
                    });
                  }
                }}
              >
                <Archive className="mr-2 h-4 w-4" /> Arquivar
              </Button>
            )}
          </div>
        )}
      </div>

      {editing ? (
        <>
          <CampaignForm
            initial={campaign}
            onSubmit={onSubmit}
            submitting={update.isPending}
            submitLabel="Salvar alterações"
          />
          <Button variant="ghost" onClick={() => setEditing(false)}>
            Cancelar
          </Button>
        </>
      ) : (
        <div className="grid gap-6 md:grid-cols-[1fr_auto]">
          {/* Detalhes */}
          <div className="space-y-4">
            <Card>
              <CardHeader>
                <CardTitle className="text-base">URL final</CardTitle>
              </CardHeader>
              <CardContent className="space-y-3">
                <p className="font-mono text-xs break-all bg-muted p-3 rounded-md">
                  {campaign.url}
                </p>
                <div className="flex gap-2">
                  <Button size="sm" variant="outline" onClick={copyURL}>
                    {copied ? (
                      <>
                        <Check className="mr-2 h-4 w-4" /> Copiada
                      </>
                    ) : (
                      <>
                        <Copy className="mr-2 h-4 w-4" /> Copiar
                      </>
                    )}
                  </Button>
                  <a href={campaign.url} target="_blank" rel="noreferrer">
                    <Button size="sm" variant="outline">
                      <ExternalLink className="mr-2 h-4 w-4" /> Abrir
                    </Button>
                  </a>
                </div>
              </CardContent>
            </Card>

            <Card>
              <CardHeader>
                <CardTitle className="text-base">Atribuição UTM</CardTitle>
              </CardHeader>
              <CardContent className="space-y-2 text-sm">
                <Row label="utm_source" value={campaign.utmSource} />
                <Row label="utm_medium" value={campaign.utmMedium} />
                <Row label="utm_campaign" value={campaign.utmCampaign} />
                {campaign.utmTerm && <Row label="utm_term" value={campaign.utmTerm} />}
                <Row label="landing" value={campaign.landingPath} />
              </CardContent>
            </Card>

            <Card>
              <CardHeader>
                <CardTitle className="text-base">Métricas</CardTitle>
              </CardHeader>
              <CardContent>
                <div className="grid grid-cols-2 gap-4">
                  <Stat label="Sessões do Escore" value={campaign.stats?.sessionsCount ?? 0} />
                  <Stat label="Leads gerados" value={campaign.stats?.leadsCount ?? 0} />
                </div>
                <p className="text-xs text-muted-foreground mt-3">
                  Sessões: respostas registradas via Painel/Triagem com este utm_campaign.
                  Leads: visitantes que fizeram claim e viraram lead atribuído.
                </p>
              </CardContent>
            </Card>

            {campaign.description && (
              <Card>
                <CardHeader>
                  <CardTitle className="text-base">Descrição</CardTitle>
                </CardHeader>
                <CardContent>
                  <p className="text-sm whitespace-pre-wrap">{campaign.description}</p>
                </CardContent>
              </Card>
            )}

            <p className="text-xs text-muted-foreground">
              Criada em{' '}
              {format(new Date(campaign.createdAt), "dd 'de' MMMM 'de' yyyy 'às' HH:mm", {
                locale: ptBR,
              })}
            </p>
          </div>

          {/* QR code */}
          <Card className="md:w-[280px]">
            <CardHeader>
              <CardTitle className="text-base">QR Code</CardTitle>
            </CardHeader>
            <CardContent className="space-y-3">
              {qrURL ? (
                <>
                  {/* eslint-disable-next-line @next/next/no-img-element */}
                  <img
                    src={qrURL}
                    alt={`QR code da campanha ${campaign.name}`}
                    className="w-full h-auto rounded-md border"
                  />
                  <a href={qrURL} download={`plenya-campanha-${campaign.slug}.png`}>
                    <Button size="sm" variant="outline" className="w-full">
                      <Download className="mr-2 h-4 w-4" /> Baixar PNG
                    </Button>
                  </a>
                </>
              ) : (
                <Skeleton className="aspect-square w-full" />
              )}
            </CardContent>
          </Card>
        </div>
      )}
    </div>
  );
}

function Row({ label, value }: { label: string; value: string }) {
  return (
    <div className="flex items-baseline gap-3">
      <span className="text-xs uppercase tracking-wider text-muted-foreground w-32 shrink-0">
        {label}
      </span>
      <code className="font-mono text-sm">{value}</code>
    </div>
  );
}

function Stat({ label, value }: { label: string; value: number }) {
  return (
    <div className="border rounded-md p-3">
      <p className="text-2xl font-semibold tabular-nums">{value}</p>
      <p className="text-xs text-muted-foreground mt-1">{label}</p>
    </div>
  );
}
