'use client'

import { useState } from 'react'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { Badge } from '@/components/ui/badge'
import { BellRing, BellOff, Check, Loader2, RefreshCw, Send, Share } from 'lucide-react'
import { toast } from 'sonner'
import { useWebPush } from '@/lib/web-push'

/**
 * "Avisos neste aparelho" no perfil — controle de Web Push alcançável.
 *
 * O controle existia só dentro do painel do sino, e o sino foi escondido em junho/2026
 * (`top-bar.tsx`, SHOW_NOTIFICATION_BELL). Resultado: por meses não houve como ligar, testar
 * nem renovar os avisos — e quando a inscrição do iPhone ficou órfã (a Apple seguia aceitando
 * o envio com 201, mas nada aparecia na tela), não havia caminho para refazê-la.
 *
 * Por isso o botão "Renovar inscrição": ele descarta a inscrição atual e cria outra, amarrada
 * ao service worker que está rodando agora. É o remédio para push aceito pelo servidor que não
 * chega ao aparelho.
 */
export function WebPushCard() {
  const { state, busy, enable, disable, sendTest, refresh } = useWebPush()
  const [renewing, setRenewing] = useState(false)

  if (state === 'loading') return null

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

  const handleRenew = async () => {
    setRenewing(true)
    try {
      await disable()
      const ok = await enable()
      if (ok) {
        toast.success('Inscrição renovada', {
          description: 'Um aviso de teste foi enviado para este aparelho.',
        })
        sendTest().catch(() => {})
      } else {
        toast.error('Não foi possível renovar a inscrição')
      }
    } finally {
      setRenewing(false)
      await refresh()
    }
  }

  const handleTest = async () => {
    try {
      await sendTest()
      toast.success('Aviso de teste enviado', {
        description: 'Se ele não aparecer em alguns segundos, use "Renovar inscrição".',
      })
    } catch {
      toast.error('Não foi possível enviar o teste')
    }
  }

  const working = busy || renewing

  return (
    <Card>
      <CardHeader>
        <CardTitle className="flex items-center gap-2">
          <BellRing className="h-5 w-5" />
          Avisos neste aparelho
          {state === 'subscribed' && (
            <Badge variant="secondary" className="ml-1 gap-1">
              <Check className="h-3 w-3" />
              ativo
            </Badge>
          )}
        </CardTitle>
      </CardHeader>
      <CardContent className="space-y-4">
        {state === 'unsupported' && (
          <p className="text-sm text-muted-foreground">
            Este navegador não recebe avisos com o EMR fechado.
          </p>
        )}

        {state === 'ios-needs-install' && (
          <div className="flex items-start gap-2 rounded-lg border bg-muted/40 p-3 text-sm text-muted-foreground">
            <Share className="mt-0.5 h-4 w-4 shrink-0" />
            <span>
              No iPhone os avisos só funcionam com o app na tela de início. Toque em{' '}
              <span className="font-medium">Compartilhar</span>, depois em{' '}
              <span className="font-medium">Adicionar à Tela de Início</span>, abra pelo ícone e
              volte aqui.
            </span>
          </div>
        )}

        {state === 'denied' && (
          <div className="flex items-start gap-2 rounded-lg border bg-muted/40 p-3 text-sm text-muted-foreground">
            <BellOff className="mt-0.5 h-4 w-4 shrink-0" />
            <span>
              Os avisos estão bloqueados para este site. Libere as notificações nas configurações
              do aparelho e volte aqui.
            </span>
          </div>
        )}

        {(state === 'default' || state === 'subscribed') && (
          <>
            <p className="text-sm text-muted-foreground">
              {state === 'subscribed'
                ? 'Você recebe avisos de mensagens mesmo com o EMR fechado. Se pararem de aparecer, renove a inscrição.'
                : 'Receba avisos de mensagens mesmo com o EMR fechado.'}
            </p>

            <div className="flex flex-wrap gap-2">
              {state === 'subscribed' ? (
                <>
                  <Button variant="outline" onClick={handleRenew} disabled={working}>
                    {renewing ? (
                      <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                    ) : (
                      <RefreshCw className="mr-2 h-4 w-4" />
                    )}
                    Renovar inscrição
                  </Button>
                  <Button variant="outline" onClick={handleTest} disabled={working}>
                    <Send className="mr-2 h-4 w-4" />
                    Enviar teste
                  </Button>
                  <Button variant="ghost" onClick={handleDisable} disabled={working}>
                    Desativar
                  </Button>
                </>
              ) : (
                <Button onClick={handleEnable} disabled={working}>
                  {busy ? (
                    <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                  ) : (
                    <BellRing className="mr-2 h-4 w-4" />
                  )}
                  Ativar avisos
                </Button>
              )}
            </div>
          </>
        )}
      </CardContent>
    </Card>
  )
}
