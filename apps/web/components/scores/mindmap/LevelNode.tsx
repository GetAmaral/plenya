'use client'

import { memo } from 'react'
import { Handle, Position, NodeProps } from 'reactflow'
import { Edit, Trash2 } from 'lucide-react'
import { cn } from '@/lib/utils'
import { getScoreLevelColorFull } from '@/lib/utils/score-level-colors'
import {
  ContextMenu,
  ContextMenuContent,
  ContextMenuItem,
  ContextMenuTrigger,
  ContextMenuSeparator,
} from '@/components/ui/context-menu'

export const LevelNode = memo(({ data }: NodeProps) => {
  const colorClass = getScoreLevelColorFull(data.level)

  return (
    <ContextMenu>
      <ContextMenuTrigger asChild>
        <div
          className={cn(
            'w-[336px] px-3 py-2 rounded border-2 shadow-sm',
            colorClass
          )}
        >
          <Handle type="target" position={Position.Left} id="left" className="w-2 h-2" />

          <div className="space-y-1">
            <div className="flex items-center justify-between gap-2">
              <div className="font-bold text-xs">Nível {data.level}</div>
              {(data.lowerLimit || data.upperLimit) && (
                <div className="text-xs opacity-80 font-mono shrink-0">
                  {data.operator === 'between' && data.lowerLimit && data.upperLimit
                    ? `${data.lowerLimit} - ${data.upperLimit}`
                    : data.operator === '='
                    ? `= ${data.lowerLimit || data.upperLimit}`
                    : data.operator === '>'
                    ? `> ${data.lowerLimit || data.upperLimit}`
                    : data.operator === '>='
                    ? `≥ ${data.lowerLimit || data.upperLimit}`
                    : data.operator === '<'
                    ? `< ${data.upperLimit || data.lowerLimit}`
                    : data.operator === '<='
                    ? `≤ ${data.upperLimit || data.lowerLimit}`
                    : `${data.lowerLimit || ''} ${data.upperLimit || ''}`}
                </div>
              )}
            </div>
            <div className="text-xs font-medium leading-tight break-words">{data.name}</div>
          </div>
        </div>
      </ContextMenuTrigger>
      <ContextMenuContent>
        <ContextMenuItem onClick={() => data.onEdit?.(data.itemId)}>
          <Edit className="mr-2 h-4 w-4" />
          Editar Nível
        </ContextMenuItem>
        <ContextMenuSeparator />
        <ContextMenuItem onClick={() => data.onDelete?.()} className="text-destructive focus:text-destructive">
          <Trash2 className="mr-2 h-4 w-4" />
          Excluir Nível
        </ContextMenuItem>
      </ContextMenuContent>
    </ContextMenu>
  )
})

LevelNode.displayName = 'LevelNode'
