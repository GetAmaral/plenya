'use client'

import { useEffect } from 'react'
import { apiClient } from '@/lib/api-client'

/**
 * Mantém a sessão deste aparelho renovada.
 *
 * A validade do refresh é deslizante no servidor (7 dias a cada renovação), mas ela só desliza
 * quando alguém chama /auth/refresh. Este componente faz isso ao abrir o app e toda vez que ele
 * volta do segundo plano — que é o momento em que o usuário "entra". Resultado prático: quem usa
 * o EMR com alguma frequência nunca mais vê a tela de login.
 *
 * Sem rede não acontece nada de ruim: `ensureFreshSession` mantém a sessão e tenta de novo
 * depois. Só uma recusa explícita do servidor (401/403) encerra a sessão.
 */
export function SessionKeepAlive() {
  useEffect(() => {
    void apiClient.ensureFreshSession()

    const onVisible = () => {
      if (document.visibilityState === 'visible') void apiClient.ensureFreshSession()
    }
    document.addEventListener('visibilitychange', onVisible)
    window.addEventListener('focus', onVisible)
    return () => {
      document.removeEventListener('visibilitychange', onVisible)
      window.removeEventListener('focus', onVisible)
    }
  }, [])

  return null
}
