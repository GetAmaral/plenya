'use client';

/**
 * /configuracoes/integracoes
 *
 * Lista das integrações disponíveis. V1 só tem Google Calendar.
 * Quando OAuth callback retorna ?google_connected=true, mostra toast.
 * Se ?google_error=X, mostra toast de erro com mensagem amigável.
 */
import { useEffect, useState } from 'react';
import { useSearchParams, useRouter } from 'next/navigation';
import { formatDate } from '@/lib/format-date';
import {
  Calendar as CalendarIcon,
  CheckCircle2,
  XCircle,
  AlertTriangle,
  Loader2,
  ExternalLink,
} from 'lucide-react';
import { toast } from 'sonner';

import { useRequireAuth } from '@/lib/use-auth';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Badge } from '@/components/ui/badge';
import { Skeleton } from '@/components/ui/skeleton';
import { PageHeader } from '@/components/layout/page-header';
import {
  useFetchGoogleAuthURL,
  useGoogleDisconnect,
  useGoogleStatus,
} from '@/lib/api/calendar-api';

const GOOGLE_ERROR_MESSAGES: Record<string, string> = {
  not_configured: 'Google Calendar não configurado pelo administrador.',
  consent_denied: 'Você negou o acesso. Tente novamente para conectar.',
  missing_params: 'Parâmetros ausentes no callback do Google.',
  invalid_state: 'Sessão expirou. Tente conectar novamente.',
  invalid_state_claims: 'Estado inválido. Tente novamente.',
  invalid_user_id: 'Identificação de usuário inválida.',
  exchange_failed: 'Falha ao trocar o código por tokens. Tente novamente.',
  userinfo_failed: 'Falha ao obter dados da conta Google.',
  create_calendar_failed: 'Falha ao criar calendário dedicado no Google.',
  encrypt_failed: 'Erro ao salvar credenciais (criptografia).',
  db_lookup_failed: 'Erro ao acessar credenciais existentes.',
  db_create_failed: 'Erro ao salvar credenciais.',
  db_save_failed: 'Erro ao atualizar credenciais.',
};

export default function IntegracoesPage() {
  useRequireAuth();
  const params = useSearchParams();
  const router = useRouter();
  const [calledOnce, setCalledOnce] = useState(false);

  const { data: status, isLoading, refetch } = useGoogleStatus();
  const fetchAuthURL = useFetchGoogleAuthURL();
  const disconnectMutation = useGoogleDisconnect();

  // Processa callback do Google (success/error) — só uma vez.
  useEffect(() => {
    if (calledOnce) return;
    const connected = params.get('google_connected');
    const error = params.get('google_error');
    if (connected === 'true') {
      toast.success('Google Calendar conectado!', {
        description: 'Eventos serão sincronizados automaticamente.',
      });
      void refetch();
      router.replace('/configuracoes/integracoes');
      setCalledOnce(true);
    } else if (error) {
      const msg = GOOGLE_ERROR_MESSAGES[error] ?? `Erro: ${error}`;
      toast.error('Falha ao conectar Google', { description: msg });
      router.replace('/configuracoes/integracoes');
      setCalledOnce(true);
    }
  }, [params, refetch, router, calledOnce]);

  const handleConnect = async () => {
    try {
      const result = await fetchAuthURL.mutateAsync();
      if (!result.url) {
        toast.error('URL de OAuth não retornada pelo backend.');
        return;
      }
      window.location.href = result.url;
    } catch (err) {
      const msg = err instanceof Error ? err.message : 'Erro desconhecido';
      // Backend retorna 503 se não configured.
      if (msg.includes('google_not_configured') || msg.includes('not_configured')) {
        toast.error('Google Calendar não configurado', {
          description: 'O administrador precisa configurar GOOGLE_CLIENT_ID/SECRET.',
        });
      } else {
        toast.error('Erro ao gerar URL de OAuth', { description: msg });
      }
    }
  };

  const handleDisconnect = async () => {
    if (!confirm('Desconectar Google Calendar? Eventos antigos permanecem, mas novos não serão sincronizados.')) {
      return;
    }
    try {
      await disconnectMutation.mutateAsync();
      toast.success('Google Calendar desconectado');
    } catch (err) {
      toast.error('Erro ao desconectar', {
        description: err instanceof Error ? err.message : undefined,
      });
    }
  };

  return (
    <div className="container mx-auto space-y-6 py-8">
      <PageHeader
        breadcrumbs={[{ label: 'Configurações' }, { label: 'Integrações' }]}
        title="Integrações"
        description="Conecte serviços externos ao seu fluxo de trabalho."
      />

      {/* Google Calendar */}
      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-3">
            <div className="flex h-10 w-10 items-center justify-center rounded-lg bg-blue-100 text-blue-700">
              <CalendarIcon className="h-5 w-5" />
            </div>
            <div className="flex-1">
              <div className="text-base">Google Calendar</div>
              <div className="text-xs font-normal text-muted-foreground">
                Sincronize consultas com seu calendário Google.
              </div>
            </div>
          </CardTitle>
        </CardHeader>
        <CardContent>
          {isLoading ? (
            <div className="space-y-3">
              <Skeleton className="h-6 w-32" />
              <Skeleton className="h-10 w-40" />
            </div>
          ) : !status?.configured ? (
            <div className="flex items-start gap-3 rounded border border-amber-300 bg-amber-50 p-4">
              <AlertTriangle className="mt-0.5 h-5 w-5 shrink-0 text-amber-600" />
              <div className="flex-1 text-sm">
                <p className="font-semibold text-amber-900">Não configurado</p>
                <p className="text-amber-800">
                  O administrador ainda não definiu GOOGLE_CLIENT_ID/SECRET. Entre em
                  contato para habilitar.
                </p>
              </div>
            </div>
          ) : status.connected ? (
            <div className="space-y-4">
              <div className="flex items-center gap-2">
                <Badge variant="outline" className="border-emerald-300 bg-emerald-100 text-emerald-900">
                  <CheckCircle2 className="mr-1 h-3 w-3" />
                  Conectado
                </Badge>
                {status.googleEmail && (
                  <span className="text-sm text-muted-foreground">{status.googleEmail}</span>
                )}
              </div>
              {status.lastSyncAt && (
                <p className="text-xs text-muted-foreground">
                  Última sincronização:{' '}
                  {formatDate(status.lastSyncAt, "dd 'de' MMM yyyy 'às' HH:mm")}
                </p>
              )}
              <div className="flex gap-2">
                <Button
                  variant="outline"
                  size="sm"
                  onClick={handleConnect}
                  disabled={fetchAuthURL.isPending}
                >
                  {fetchAuthURL.isPending && <Loader2 className="mr-2 h-3 w-3 animate-spin" />}
                  Reconectar
                </Button>
                <Button
                  variant="ghost"
                  size="sm"
                  onClick={handleDisconnect}
                  disabled={disconnectMutation.isPending}
                  className="text-rose-600 hover:bg-rose-50 hover:text-rose-700"
                >
                  {disconnectMutation.isPending && (
                    <Loader2 className="mr-2 h-3 w-3 animate-spin" />
                  )}
                  Desconectar
                </Button>
              </div>
            </div>
          ) : (
            <div className="space-y-3">
              {status.revokedAt ? (
                <Badge variant="outline" className="border-rose-300 bg-rose-100 text-rose-900">
                  <XCircle className="mr-1 h-3 w-3" />
                  Revogado
                </Badge>
              ) : (
                <Badge variant="outline" className="border-stone-300 bg-stone-100 text-stone-900">
                  Desconectado
                </Badge>
              )}
              <Button onClick={handleConnect} disabled={fetchAuthURL.isPending}>
                {fetchAuthURL.isPending ? (
                  <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                ) : (
                  <ExternalLink className="mr-2 h-4 w-4" />
                )}
                Conectar com Google
              </Button>
            </div>
          )}
        </CardContent>
      </Card>
    </div>
  );
}
