"use client"

import { Info } from "lucide-react"
import { Badge } from "@/components/ui/badge"
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "@/components/ui/tooltip"

export function UnmatchedBadge() {
  return (
    <TooltipProvider>
      <Tooltip>
        <TooltipTrigger asChild>
          <Badge
            variant="outline"
            className="bg-yellow-50 border-yellow-300 text-yellow-700 hover:bg-yellow-100"
          >
            <Info className="h-3 w-3 mr-1" />
            Não catalogado
          </Badge>
        </TooltipTrigger>
        <TooltipContent className="max-w-xs">
          <p>
            Este exame não foi encontrado no catálogo de exames padronizados. Os dados foram
            extraídos do PDF, mas sem correspondência automática com as definições do sistema.
          </p>
        </TooltipContent>
      </Tooltip>
    </TooltipProvider>
  )
}
