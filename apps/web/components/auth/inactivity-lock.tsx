'use client'

import { useCallback, useEffect, useRef, useState } from 'react'
import { useRouter } from 'next/navigation'
import { Lock, Loader2 } from 'lucide-react'
import { toast } from 'sonner'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { apiClient } from '@/lib/api-client'
import { useAuthStore } from '@/lib/auth-store'

// Bloqueio de tela por inatividade (CFM/SBIS): após 4h sem atividade, cobre o app e exige
// reautenticação por senha (checagem de presença; a sessão continua a mesma). Ver
// docs/emr/estudo-sessao-login-persistente.md.
const INACTIVITY_MS = 4 * 60 * 60 * 1000 // 4 horas
const ACTIVITY_KEY = 'plenya-last-activity'
const THROTTLE_MS = 15 * 1000

export function touchActivity() {
  try {
    localStorage.setItem(ACTIVITY_KEY, String(Date.now()))
  } catch {
    /* storage indisponível */
  }
}

export function InactivityLock({ children }: { children: React.ReactNode }) {
  const [locked, setLocked] = useState(false)
  const lastRef = useRef<number>(Date.now())

  const touch = useCallback(() => {
    const now = Date.now()
    lastRef.current = now
    try {
      localStorage.setItem(ACTIVITY_KEY, String(now))
    } catch {
      /* noop */
    }
  }, [])

  // Inicializa a partir do storage; se já passou do limite (ex.: app reaberto depois de horas),
  // trava de cara.
  useEffect(() => {
    const stored = Number(localStorage.getItem(ACTIVITY_KEY) || 0)
    const last = stored || Date.now()
    lastRef.current = last
    if (Date.now() - last >= INACTIVITY_MS) {
      setLocked(true)
    } else if (!stored) {
      touch()
    }
  }, [touch])

  // Registra atividade (throttle) enquanto desbloqueado.
  useEffect(() => {
    if (locked) return
    let lastFire = 0
    const onActivity = () => {
      const now = Date.now()
      if (now - lastFire < THROTTLE_MS) return
      lastFire = now
      touch()
    }
    const events = ['mousemove', 'mousedown', 'keydown', 'touchstart', 'scroll', 'click']
    events.forEach((e) => window.addEventListener(e, onActivity, { passive: true }))
    return () => events.forEach((e) => window.removeEventListener(e, onActivity))
  }, [locked, touch])

  // Verifica inatividade periodicamente e ao reganhar foco/visibilidade.
  useEffect(() => {
    const check = () => {
      if (Date.now() - lastRef.current >= INACTIVITY_MS) setLocked(true)
    }
    const id = window.setInterval(check, 30 * 1000)
    const onVis = () => {
      if (document.visibilityState === 'visible') {
        // relê o storage (outra aba pode ter atualizado) antes de checar
        const stored = Number(localStorage.getItem(ACTIVITY_KEY) || 0)
        if (stored) lastRef.current = stored
        check()
      }
    }
    window.addEventListener('focus', check)
    document.addEventListener('visibilitychange', onVis)
    return () => {
      window.clearInterval(id)
      window.removeEventListener('focus', check)
      document.removeEventListener('visibilitychange', onVis)
    }
  }, [])

  const handleUnlock = useCallback(() => {
    touch()
    setLocked(false)
  }, [touch])

  return (
    <>
      {children}
      {locked && <LockScreen onUnlock={handleUnlock} />}
    </>
  )
}

function LockScreen({ onUnlock }: { onUnlock: () => void }) {
  const router = useRouter()
  const user = useAuthStore((s) => s.user)
  const clearAuth = useAuthStore((s) => s.clearAuth)
  const [password, setPassword] = useState('')
  const [loading, setLoading] = useState(false)

  const submit = async (e: React.FormEvent) => {
    e.preventDefault()
    if (!password) return
    setLoading(true)
    try {
      await apiClient.post('/api/v1/auth/verify-password', { password })
      setPassword('')
      onUnlock()
    } catch {
      toast.error('Senha incorreta')
    } finally {
      setLoading(false)
    }
  }

  const handleLogout = () => {
    clearAuth()
    router.push('/login')
  }

  return (
    <div className="fixed inset-0 z-[100] flex items-center justify-center bg-background/80 backdrop-blur-xl">
      <div className="w-full max-w-sm rounded-2xl border bg-card p-8 shadow-2xl">
        <div className="mb-6 flex flex-col items-center text-center">
          <div className="mb-3 flex h-14 w-14 items-center justify-center rounded-full bg-primary/10">
            <Lock className="h-7 w-7 text-primary" />
          </div>
          <h2 className="text-lg font-semibold">Tela bloqueada</h2>
          <p className="mt-1 text-sm text-muted-foreground">
            Sua sessão foi bloqueada por inatividade. Digite sua senha para continuar.
          </p>
          {user?.email && (
            <p className="mt-2 text-xs font-medium text-muted-foreground">{user.email}</p>
          )}
        </div>

        <form onSubmit={submit} className="space-y-4">
          <div className="space-y-2">
            <Label htmlFor="unlock-password">Senha</Label>
            <Input
              id="unlock-password"
              type="password"
              autoFocus
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              placeholder="••••••••"
            />
          </div>
          <Button type="submit" className="w-full" disabled={loading || !password}>
            {loading ? <Loader2 className="mr-2 h-4 w-4 animate-spin" /> : null}
            Desbloquear
          </Button>
        </form>

        <button
          type="button"
          onClick={handleLogout}
          className="mt-4 w-full text-center text-xs text-muted-foreground hover:text-foreground"
        >
          Sair e entrar com outra conta
        </button>
      </div>
    </div>
  )
}
