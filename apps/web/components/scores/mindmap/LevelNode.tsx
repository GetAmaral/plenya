'use client'

import { memo } from 'react'
import { Handle, Position, NodeProps } from 'reactflow'
import { Edit, Trash2 } from 'lucide-react'
import { cn } from '@/lib/utils'
import {
  ContextMenu,
  ContextMenuContent,
  ContextMenuItem,
  ContextMenuTrigger,
  ContextMenuSeparator,
} from '@/components/ui/context-menu'

const LEVEL_COLORS = {
  0: 'bg-red-100 dark:bg-red-950 border-red-500 text-red-900 dark:text-red-100',
  1: 'bg-orange-100 dark:bg-orange-950 border-orange-500 text-orange-900 dark:text-orange-100',
  2: 'bg-yellow-100 dark:bg-yellow-950 border-yellow-500 text-yellow-900 dark:text-yellow-100',
  3: 'bg-blue-100 dark:bg-blue-950 border-blue-500 text-blue-900 dark:text-blue-100',
  4: 'bg-green-100 dark:bg-green-950 border-green-500 text-green-900 dark:text-green-100',
  5: 'bg-emerald-100 dark:bg-emerald-950 border-emerald-500 text-emerald-900 dark:text-emerald-100',
  6: 'bg-gray-100 dark:bg-gray-950 border-gray-500 text-gray-900 dark:text-gray-100',
}

export const LevelNode = memo(({ data }: NodeProps) => {
  const colorClass = LEVEL_COLORS[data.level as keyof typeof LEVEL_COLORS] || LEVEL_COLORS[6]

  return (
    <ContextMenu>
      <ContextMenuTrigger asChild>
        <div
          className={cn(
            'w-[336px] px-3 py-2 rounded border-2 shadow-sm',
            colorClass
          )}
        >
          <Handle type="target" position={Position.Left} className="w-2 h-2" />

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
