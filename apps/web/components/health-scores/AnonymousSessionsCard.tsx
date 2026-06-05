'use client'

import { toast } from 'sonner'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { Badge } from '@/components/ui/badge'
import { Download, Check, Sparkles } from 'lucide-react'
import {
  usePatientAnonymousSessions,
  useImportAnonymousSession,
} from '@/lib/api/health-score-api'

function fmtDate(iso: string): string {
  try {
    return new Date(iso).toLocaleDateString('pt-BR', {
      day: '2-digit',
      month: '2-digit',
      year: 'numeric',
    })
  } catch {
    return iso
  }
}

/**
 * Lista as sessões do Escore Plenya (Light/Triagem) que o paciente fez e vinculou à conta,
 * com botão "Importar pro prontuário" que materializa um patient_score_snapshot parcial.
 * Não renderiza nada quando o paciente não tem sessões (mantém a página limpa). Fase 4.
 */
export function AnonymousSessionsCard({ patientId }: { patientId: string }) {
  const { data: sessions, isLoading } = usePatientAnonymousSessions(patientId)
  const importMut = useImportAnonymousSession(patientId)

  if (!patientId || isLoading || !sessions || sessions.length === 0) return null

  const handleImport = async (code: string) => {
    try {
      await importMut.mutateAsync(code)
      toast.success('Sessão importada pro prontuário')
    } catch (e: any) {
      toast.error(e?.message || 'Erro ao importar sessão')
    }
  }

  return (
    <Card>
      <CardHeader>
        <CardTitle className="flex items-center gap-2">
          <Sparkles className="h-5 w-5 text-muted-foreground" />
          Escore Plenya (Light/Triagem)
        </CardTitle>
        <CardDescription>
          Avaliações públicas que o paciente vinculou à conta. Importe pro prontuário para gerar
          um snapshot oficial (parcial) no histórico de escores.
        </CardDescription>
      </CardHeader>
      <CardContent>
        <div className="divide-y">
          {sessions.map((s) => {
            const pct = s.snapshot ? Math.round(s.snapshot.totalScorePercentage) : null
            const imported = !!s.importedSnapshotId
            return (
              <div
                key={s.publicCode}
                className="flex flex-wrap items-center justify-between gap-3 py-3"
              >
                <div className="min-w-0">
                  <div className="flex items-center gap-2">
                    <span className="font-medium">{fmtDate(s.createdAt)}</span>
                    {pct !== null && <Badge variant="outline">{pct}%</Badge>}
                    {imported && (
                      <Badge variant="secondary" className="gap-1">
                        <Check className="h-3 w-3" /> Importada
                      </Badge>
                    )}
                  </div>
                  <p className="text-xs text-muted-foreground">
                    {s.snapshot
                      ? `${s.snapshot.itemsEvaluatedCount} itens avaliados`
                      : 'Sem resultado calculado'}
                  </p>
                </div>
                <Button
                  size="sm"
                  variant={imported ? 'outline' : 'default'}
                  disabled={imported || !s.snapshot || importMut.isPending}
                  onClick={() => handleImport(s.publicCode)}
                  className="gap-2"
                >
                  <Download className="h-4 w-4" />
                  {imported ? 'Importada' : 'Importar pro prontuário'}
                </Button>
              </div>
            )
          })}
        </div>
      </CardContent>
    </Card>
  )
}
