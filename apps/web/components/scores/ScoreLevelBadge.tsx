'use client'

import { cn } from '@/lib/utils'
import { Badge } from '@/components/ui/badge'
import { Edit, Trash2, Calendar } from 'lucide-react'
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from '@/components/ui/tooltip'
import {
  ContextMenu,
  ContextMenuContent,
  ContextMenuItem,
  ContextMenuTrigger,
  ContextMenuSeparator,
} from '@/components/ui/context-menu'
import { ScoreLevel } from '@/lib/api/score-api'
import { format } from 'date-fns'
import { ptBR } from 'date-fns/locale'

interface ScoreLevelBadgeProps {
  level: ScoreLevel
  size?: 'sm' | 'md' | 'lg'
  showTooltip?: boolean
  onEdit?: () => void
  onDelete?: () => void
}

const LEVEL_STYLES = {
  0: {
    bg: 'bg-red-100 dark:bg-red-950',
    text: 'text-red-900 dark:text-red-100',
    border: 'border-red-500',
    label: 'Crítico',
  },
  1: {
    bg: 'bg-orange-100 dark:bg-orange-950',
    text: 'text-orange-900 dark:text-orange-100',
    border: 'border-orange-500',
    label: 'Muito Baixo/Alto',
  },
  2: {
    bg: 'bg-yellow-100 dark:bg-yellow-950',
    text: 'text-yellow-900 dark:text-yellow-100',
    border: 'border-yellow-500',
    label: 'Subótimo',
  },
  3: {
    bg: 'bg-blue-100 dark:bg-blue-950',
    text: 'text-blue-900 dark:text-blue-100',
    border: 'border-blue-500',
    label: 'Limítrofe',
  },
  4: {
    bg: 'bg-green-100 dark:bg-green-950',
    text: 'text-green-900 dark:text-green-100',
    border: 'border-green-500',
    label: 'Bom',
  },
  5: {
    bg: 'bg-emerald-100 dark:bg-emerald-950',
    text: 'text-emerald-900 dark:text-emerald-100',
    border: 'border-emerald-500',
    label: 'Ótimo',
  },
  6: {
    bg: 'bg-gray-100 dark:bg-gray-950',
    text: 'text-gray-900 dark:text-gray-100',
    border: 'border-gray-500',
    label: 'Reservado',
  },
}

const SIZE_STYLES = {
  sm: 'text-xs px-2 py-0.5',
  md: 'text-sm px-2.5 py-1',
  lg: 'text-base px-3 py-1.5',
}

export function ScoreLevelBadge({
  level,
  size = 'md',
  showTooltip = true,
  onEdit,
  onDelete,
}: ScoreLevelBadgeProps) {
  const style = LEVEL_STYLES[level.level as keyof typeof LEVEL_STYLES] || LEVEL_STYLES[6]
  const hasActions = onEdit || onDelete

  const badgeContent = (
    <Badge
      variant="outline"
      className={cn(
        'font-medium border',
        style.bg,
        style.text,
        style.border,
        SIZE_STYLES[size],
        hasActions && 'cursor-context-menu'
      )}
    >
      <span className="font-bold mr-1">N{level.level}:</span>
      {level.name}
    </Badge>
  )

  const badge = hasActions ? (
    <ContextMenu>
      <ContextMenuTrigger asChild>
        {badgeContent}
      </ContextMenuTrigger>
      <ContextMenuContent>
        {onEdit && (
          <ContextMenuItem onClick={onEdit}>
            <Edit className="mr-2 h-4 w-4" />
            Editar Nível
          </ContextMenuItem>
        )}
        {onEdit && onDelete && <ContextMenuSeparator />}
        {onDelete && (
          <ContextMenuItem onClick={onDelete} className="text-destructive focus:text-destructive">
            <Trash2 className="mr-2 h-4 w-4" />
            Excluir Nível
          </ContextMenuItem>
        )}
      </ContextMenuContent>
    </ContextMenu>
  ) : (
    badgeContent
  )

  if (!showTooltip) {
    return badge
  }

  return (
    <TooltipProvider>
      <Tooltip>
        <TooltipTrigger asChild>
          <div>{badge}</div>
        </TooltipTrigger>
        <TooltipContent className="max-w-md">
          <div className="space-y-2">
            <p className="font-semibold">
              Nível {level.level}: {style.label}
            </p>
            {(level.lowerLimit || level.upperLimit) && (
              <p className="text-xs font-mono">
                {level.operator === 'between' && level.lowerLimit && level.upperLimit
                  ? `${level.lowerLimit} - ${level.upperLimit}`
                  : level.operator === '='
                  ? `= ${level.lowerLimit || level.upperLimit}`
                  : level.operator === '>'
                  ? `> ${level.lowerLimit || level.upperLimit}`
                  : level.operator === '>='
                  ? `≥ ${level.lowerLimit || level.upperLimit}`
                  : level.operator === '<'
                  ? `< ${level.upperLimit || level.lowerLimit}`
                  : level.operator === '<='
                  ? `≤ ${level.upperLimit || level.lowerLimit}`
                  : `${level.lowerLimit || ''} ${level.upperLimit || ''}`}
              </p>
            )}
            {level.lastReview && (
              <div className="flex items-center gap-1 text-xs text-muted-foreground">
                <Calendar className="h-3 w-3" />
                <span>Revisado em {format(new Date(level.lastReview), "dd/MM/yyyy", { locale: ptBR })}</span>
              </div>
            )}
            {level.patientExplanation && (
              <div className="pt-1 border-t">
                <p className="text-xs text-muted-foreground">
                  {level.patientExplanation.length > 150
                    ? `${level.patientExplanation.substring(0, 150)}...`
                    : level.patientExplanation}
                </p>
              </div>
            )}
            {level.clinicalRelevance && (
              <div className="pt-1 border-t">
                <p className="text-xs font-medium mb-0.5">Relevância Clínica:</p>
                <p className="text-xs text-muted-foreground">
                  {level.clinicalRelevance.length > 150
                    ? `${level.clinicalRelevance.substring(0, 150)}...`
                    : level.clinicalRelevance}
                </p>
              </div>
            )}
          </div>
        </TooltipContent>
      </Tooltip>
    </TooltipProvider>
  )
}
