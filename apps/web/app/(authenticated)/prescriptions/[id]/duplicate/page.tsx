'use client'

import { useEffect } from 'react'
import { useParams, useRouter } from 'next/navigation'
import { Loader2 } from 'lucide-react'

/**
 * Duplicar receita.
 *
 * A listagem já linkava esta rota, que não existia (404). Duplicar não precisa de tela própria:
 * é a tela de nova prescrição partindo de uma receita anterior, onde o médico revisa antes de
 * assinar. Criar a cópia direto no banco produziria rascunhos que ninguém pediu.
 */
export default function DuplicatePrescriptionPage() {
  const router = useRouter()
  const params = useParams()
  const prescriptionId = params.id as string

  useEffect(() => {
    router.replace(`/prescriptions/new?from=${prescriptionId}`)
  }, [router, prescriptionId])

  return (
    <div className="flex min-h-screen items-center justify-center">
      <Loader2 className="h-8 w-8 animate-spin text-muted-foreground" />
    </div>
  )
}
