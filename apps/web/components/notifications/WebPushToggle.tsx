'use client'

import { BellRing, BellOff, Loader2, Share, Check } from 'lucide-react'
import { Button } from '@/components/ui/button'
import { useWebPush } from '@/lib/web-push'
import { toast } from 'sonner'

// Faixa "Ativar avisos neste aparelho" no topo do painel de notificações.
// Web Push: o aviso chega mesmo com o EMR fechado (no iPhone, exige
// "Adicionar à Tela de Início" antes — ver estado ios-needs-install).
export function WebPushToggle() {
  const { state, busy, enable, disable, sendTest } = useWebPush()

  if (state === 'loading' || state === 'unsupported') return null

  if (state === 'ios-needs-install') {
    return (
      <div className="px-3 py-2 border-b bg-muted/40 text-xs text-muted-foreground flex items-start gap-2">
        <Share className="h-3.5 w-3.5 mt-0.5 shrink-0" />
        <span>
          Para receber avisos no iPhone, toque em{' '}
          <span className="font-medium">Compartilhar</span> e em{' '}
          <span className="font-medium">Adicionar à Tela de Início</span>. Depois abra o app pelo
          ícone e ative aqui.
        </span>
      </div>
    )
  }

  if (state === 'denied') {
    return (
      <div className="px-3 py-2 border-b bg-muted/40 text-xs text-muted-foreground flex items-start gap-2">
        <BellOff className="h-3.5 w-3.5 mt-0.5 shrink-0" />
        <span>
          Avisos bloqueados neste navegador. Libere as notificações nas configurações do site para
          ativar.
        </span>
      </div>
    )
  }

  const handleEnable = async () => {
    const ok = await enable()
    if (ok) {
      toast.success('Avisos ativados neste aparelho')
      sendTest().catch(() => {})
    } else {
      toast.error('Não foi possível ativar os avisos')
    }
  }

  const handleDisable = async () => {
    await disable()
    toast.success('Avisos desativados neste aparelho')
  }

  return (
    <div className="px-3 py-2 border-b bg-muted/30 flex items-center justify-between gap-2">
      <div className="flex items-center gap-2 min-w-0">
        {state === 'subscribed' ? (
          <Check className="h-3.5 w-3.5 text-emerald-600 shrink-0" />
        ) : (
          <BellRing className="h-3.5 w-3.5 text-muted-foreground shrink-0" />
        )}
        <span className="text-xs text-muted-foreground truncate">
          {state === 'subscribed'
            ? 'Avisos ativos neste aparelho'
            : 'Receber avisos com o EMR fechado'}
        </span>
      </div>
      {state === 'subscribed' ? (
        <Button variant="ghost" size="sm" className="h-7 text-xs shrink-0" onClick={handleDisable} disabled={busy}>
          {busy ? <Loader2 className="h-3 w-3 animate-spin" /> : 'Desativar'}
        </Button>
      ) : (
        <Button variant="outline" size="sm" className="h-7 text-xs shrink-0" onClick={handleEnable} disabled={busy}>
          {busy ? <Loader2 className="h-3 w-3 animate-spin" /> : 'Ativar'}
        </Button>
      )}
    </div>
  )
}
