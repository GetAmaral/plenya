"use client"

import { HelpCircle } from "lucide-react"
import { Badge } from "@/components/ui/badge"
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "@/components/ui/tooltip"

/**
 * Sinaliza um exame que está NO catálogo mas não recebeu nível de classificação,
 * mostrando o motivo (sem item de score, não se aplica ao paciente, valor fora das
 * faixas, etc.). Diferente do "Não catalogado" (UnmatchedBadge).
 */
export function UnclassifiedBadge({ reason }: { reason?: string | null }) {
  return (
    <TooltipProvider>
      <Tooltip>
        <TooltipTrigger asChild>
          <Badge
            variant="outline"
            className="bg-muted border-border text-muted-foreground hover:bg-muted/70"
          >
            <HelpCircle className="h-3 w-3 mr-1" />
            Sem classificação
          </Badge>
        </TooltipTrigger>
        <TooltipContent className="max-w-xs">
          <p>{reason || "Este exame não recebeu nível de classificação."}</p>
        </TooltipContent>
      </Tooltip>
    </TooltipProvider>
  )
}
