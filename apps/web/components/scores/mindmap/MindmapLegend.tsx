'use client'

import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Badge } from '@/components/ui/badge'

export function MindmapLegend() {
  return (
    <Card className="absolute bottom-4 left-4 w-64 shadow-lg z-10">
      <CardHeader className="pb-3">
        <CardTitle className="text-sm">Legenda de Níveis</CardTitle>
      </CardHeader>
      <CardContent className="space-y-2">
        <div className="flex items-center gap-2">
          <div className="w-4 h-4 rounded bg-red-500" />
          <span className="text-xs">Nível 0 - Crítico</span>
        </div>
        <div className="flex items-center gap-2">
          <div className="w-4 h-4 rounded bg-orange-500" />
          <span className="text-xs">Nível 1 - Muito Baixo/Alto</span>
        </div>
        <div className="flex items-center gap-2">
          <div className="w-4 h-4 rounded bg-yellow-500" />
          <span className="text-xs">Nível 2 - Subótimo</span>
        </div>
        <div className="flex items-center gap-2">
          <div className="w-4 h-4 rounded bg-blue-500" />
          <span className="text-xs">Nível 3 - Limítrofe</span>
        </div>
        <div className="flex items-center gap-2">
          <div className="w-4 h-4 rounded bg-green-500" />
          <span className="text-xs">Nível 4 - Bom</span>
        </div>
        <div className="flex items-center gap-2">
          <div className="w-4 h-4 rounded bg-emerald-500" />
          <span className="text-xs">Nível 5 - Ótimo</span>
        </div>

        <div className="pt-2 mt-2 border-t text-xs text-muted-foreground space-y-1">
          <div className="flex items-center gap-2">
            <Badge variant="default" className="h-4 text-xs">Grupo</Badge>
            <span>→</span>
            <Badge variant="outline" className="h-4 text-xs">Subgrupo</Badge>
          </div>
          <div className="flex items-center gap-2">
            <Badge variant="secondary" className="h-4 text-xs">Item</Badge>
            <span>→</span>
            <Badge variant="outline" className="h-4 text-xs">Níveis</Badge>
          </div>
        </div>
      </CardContent>
    </Card>
  )
}
