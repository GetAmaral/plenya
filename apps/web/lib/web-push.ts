'use client'

import { useCallback, useEffect, useState } from 'react'
import { apiClient } from './api-client'

// Estados possíveis do canal de avisos neste aparelho.
export type WebPushState =
  | 'loading' // ainda verificando
  | 'unsupported' // navegador sem suporte (ex: iOS Safari fora da tela de início)
  | 'ios-needs-install' // iPhone: precisa "Adicionar à Tela de Início" antes
  | 'denied' // usuário bloqueou notificações no navegador
  | 'default' // suportado, ainda não ativado
  | 'subscribed' // ativo neste aparelho

function urlBase64ToUint8Array(base64String: string): Uint8Array {
  const padding = '='.repeat((4 - (base64String.length % 4)) % 4)
  const base64 = (base64String + padding).replace(/-/g, '+').replace(/_/g, '/')
  const raw = atob(base64)
  const out = new Uint8Array(raw.length)
  for (let i = 0; i < raw.length; i++) out[i] = raw.charCodeAt(i)
  return out
}

function isIOS(): boolean {
  if (typeof navigator === 'undefined') return false
  return (
    /iPad|iPhone|iPod/.test(navigator.userAgent) ||
    // iPadOS se apresenta como Mac com touch
    (navigator.platform === 'MacIntel' && (navigator as any).maxTouchPoints > 1)
  )
}

function isStandalone(): boolean {
  if (typeof window === 'undefined') return false
  return (
    window.matchMedia?.('(display-mode: standalone)').matches ||
    (navigator as any).standalone === true
  )
}

function supportsPush(): boolean {
  return (
    typeof window !== 'undefined' &&
    'serviceWorker' in navigator &&
    'PushManager' in window &&
    'Notification' in window
  )
}

function deviceLabel(): string {
  if (typeof navigator === 'undefined') return ''
  const ua = navigator.userAgent
  const browser = /Edg/.test(ua)
    ? 'Edge'
    : /Chrome/.test(ua)
    ? 'Chrome'
    : /Firefox/.test(ua)
    ? 'Firefox'
    : /Safari/.test(ua)
    ? 'Safari'
    : 'Navegador'
  const os = /Windows/.test(ua)
    ? 'Windows'
    : /Mac/.test(ua)
    ? 'macOS'
    : /iPhone|iPad/.test(ua)
    ? 'iOS'
    : /Android/.test(ua)
    ? 'Android'
    : /Linux/.test(ua)
    ? 'Linux'
    : ''
  return os ? `${browser} · ${os}` : browser
}

async function getRegistration(): Promise<ServiceWorkerRegistration> {
  const existing = await navigator.serviceWorker.getRegistration('/sw.js')
  if (existing) return existing
  return navigator.serviceWorker.register('/sw.js')
}

export function useWebPush() {
  const [state, setState] = useState<WebPushState>('loading')
  const [busy, setBusy] = useState(false)

  const refresh = useCallback(async () => {
    if (!supportsPush()) {
      // iPhone fora da tela de início não expõe PushManager: orienta a instalar.
      setState(isIOS() && !isStandalone() ? 'ios-needs-install' : 'unsupported')
      return
    }
    if (Notification.permission === 'denied') {
      setState('denied')
      return
    }
    try {
      const reg = await getRegistration()
      const sub = await reg.pushManager.getSubscription()
      setState(sub ? 'subscribed' : 'default')
    } catch {
      setState('default')
    }
  }, [])

  useEffect(() => {
    void refresh()
  }, [refresh])

  const enable = useCallback(async () => {
    setBusy(true)
    try {
      const permission = await Notification.requestPermission()
      if (permission !== 'granted') {
        setState(permission === 'denied' ? 'denied' : 'default')
        return false
      }

      const { publicKey, enabled } = await apiClient.get<{
        publicKey: string
        enabled: boolean
      }>('/api/v1/web-push/vapid-public-key')
      if (!enabled || !publicKey) {
        setState('unsupported')
        return false
      }

      const reg = await getRegistration()
      await navigator.serviceWorker.ready
      const sub = await reg.pushManager.subscribe({
        userVisibleOnly: true,
        applicationServerKey: urlBase64ToUint8Array(publicKey) as BufferSource,
      })

      await apiClient.post('/api/v1/web-push/subscribe', {
        ...sub.toJSON(),
        deviceLabel: deviceLabel(),
      })

      setState('subscribed')
      return true
    } catch (e) {
      console.error('[web-push] enable falhou', e)
      await refresh()
      return false
    } finally {
      setBusy(false)
    }
  }, [refresh])

  const disable = useCallback(async () => {
    setBusy(true)
    try {
      const reg = await getRegistration()
      const sub = await reg.pushManager.getSubscription()
      if (sub) {
        await apiClient
          .post('/api/v1/web-push/unsubscribe', { endpoint: sub.endpoint })
          .catch(() => {})
        await sub.unsubscribe().catch(() => {})
      }
      setState('default')
      return true
    } finally {
      setBusy(false)
    }
  }, [])

  const sendTest = useCallback(async () => {
    await apiClient.post('/api/v1/web-push/test', {})
  }, [])

  return { state, busy, enable, disable, sendTest, refresh }
}
