'use client'

import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { Badge } from '@/components/ui/badge'
import { Loader2, LogOut, Monitor, Smartphone } from 'lucide-react'
import { formatRelativeToNow } from '@/lib/format-date'
import { toast } from 'sonner'
import { useSessions, useRevokeSession, type Session } from '@/lib/api/sessions'

// Resumo legível do device a partir do user-agent.
function deviceLabel(ua?: string): { label: string; mobile: boolean } {
  if (!ua) return { label: 'Aparelho desconhecido', mobile: false }
  const mobile = /iPhone|iPad|Android|Mobile/i.test(ua)
  const os = /iPhone|iPad/i.test(ua)
    ? 'iPhone/iPad'
    : /Android/i.test(ua)
    ? 'Android'
    : /Windows/i.test(ua)
    ? 'Windows'
    : /Mac/i.test(ua)
    ? 'Mac'
    : /Linux/i.test(ua)
    ? 'Linux'
    : 'Outro'
  const browser = /Edg/i.test(ua)
    ? 'Edge'
    : /Chrome/i.test(ua)
    ? 'Chrome'
    : /Firefox/i.test(ua)
    ? 'Firefox'
    : /Safari/i.test(ua)
    ? 'Safari'
    : 'Navegador'
  return { label: `${browser} · ${os}`, mobile }
}

function SessionRow({ session }: { session: Session }) {
  const revoke = useRevokeSession()
  const { label, mobile } = deviceLabel(session.userAgent)
  const Icon = mobile ? Smartphone : Monitor

  const handleRevoke = async () => {
    try {
      await revoke.mutateAsync(session.id)
      toast.success('Aparelho desconectado')
    } catch {
      toast.error('Erro ao desconectar aparelho')
    }
  }

  const last = session.lastUsedAt || session.createdAt

  return (
    <div className="flex items-center justify-between gap-3 rounded-lg border p-3">
      <div className="flex items-center gap-3 min-w-0">
        <Icon className="h-5 w-5 shrink-0 text-muted-foreground" />
        <div className="min-w-0">
          <div className="flex items-center gap-2">
            <span className="text-sm font-medium truncate">{label}</span>
            {session.remember && (
              <Badge variant="outline" className="text-[10px]">Mantido conectado</Badge>
            )}
          </div>
          <p className="text-xs text-muted-foreground">
            {session.ipAddress ? `${session.ipAddress} · ` : ''}
            ativo {formatRelativeToNow(last, { addSuffix: true })}
          </p>
        </div>
      </div>
      <Button
        variant="ghost"
        size="sm"
        className="shrink-0 text-destructive hover:text-destructive"
        onClick={handleRevoke}
        disabled={revoke.isPending}
      >
        {revoke.isPending ? <Loader2 className="h-4 w-4 animate-spin" /> : <LogOut className="h-4 w-4" />}
      </Button>
    </div>
  )
}

export function SessionsCard() {
  const { data: sessions, isLoading } = useSessions()

  return (
    <Card>
      <CardHeader>
        <CardTitle>Aparelhos conectados</CardTitle>
        <p className="text-sm text-muted-foreground">
          Sessões ativas da sua conta. Desconecte aparelhos que você não reconhece.
        </p>
      </CardHeader>
      <CardContent className="space-y-2">
        {isLoading ? (
          <div className="flex items-center gap-2 text-sm text-muted-foreground">
            <Loader2 className="h-4 w-4 animate-spin" /> Carregando…
          </div>
        ) : sessions && sessions.length > 0 ? (
          sessions.map((s) => <SessionRow key={s.id} session={s} />)
        ) : (
          <p className="text-sm text-muted-foreground">Nenhuma sessão ativa.</p>
        )}
      </CardContent>
    </Card>
  )
}
