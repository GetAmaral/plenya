'use client'

import { memo } from 'react'
import { Handle, Position, NodeProps } from 'reactflow'
import { Badge } from '@/components/ui/badge'
import { ChevronRight, ChevronDown, Edit, Plus, Trash2 } from 'lucide-react'
import {
  ContextMenu,
  ContextMenuContent,
  ContextMenuItem,
  ContextMenuTrigger,
  ContextMenuSeparator,
} from '@/components/ui/context-menu'

export const ItemNode = memo(({ data }: NodeProps) => {
  const hasChildren = (data.levelCount || 0) > 0
  const isExpanded = data.isExpanded || false

  return (
    <ContextMenu>
      <ContextMenuTrigger asChild>
        <div className="w-[336px] px-4 py-3 rounded-md bg-accent border shadow">
          {/* Handles: left (entrada), right (saída para levels), bottom (saída para filhos) */}
          <Handle type="target" position={Position.Left} id="left" className="w-2 h-2" />
          <Handle type="source" position={Position.Right} id="right" className="w-2 h-2" />
          <Handle type="source" position={Position.Bottom} id="bottom" className="w-2 h-2" />

          <div className="space-y-1.5">
            <div className="flex items-start gap-2">
              {hasChildren && (
                <button
                  onClick={(e) => {
                    e.stopPropagation()
                    data.onToggle?.()
                  }}
                  className="hover:bg-muted rounded p-0.5 transition-colors shrink-0 mt-0.5"
                  title={isExpanded ? 'Recolher' : 'Expandir'}
                >
                  {isExpanded ? (
                    <ChevronDown className="h-3 w-3" />
                  ) : (
                    <ChevronRight className="h-3 w-3" />
                  )}
                </button>
              )}
              <div className="font-medium text-sm leading-tight break-words flex-1">{data.name}</div>
            </div>

            <div className="flex items-center flex-wrap gap-2 text-xs">
              {data.unit && (
                <Badge variant="secondary" className="text-xs px-1.5 py-0">
                  {data.unit}
                </Badge>
              )}
              <span className="text-muted-foreground">{data.points} pts</span>
              {hasChildren && (
                <span className="text-muted-foreground">
                  · {data.levelCount} nível{data.levelCount !== 1 ? 'is' : ''}
                </span>
              )}
            </div>
          </div>
        </div>
      </ContextMenuTrigger>
      <ContextMenuContent>
        <ContextMenuItem onClick={() => data.onEdit?.(data.subgroupId)}>
          <Edit className="mr-2 h-4 w-4" />
          Editar Item
        </ContextMenuItem>
        <ContextMenuItem onClick={() => data.onAdd?.()}>
          <Plus className="mr-2 h-4 w-4" />
          Adicionar Nível
        </ContextMenuItem>
        <ContextMenuSeparator />
        <ContextMenuItem onClick={() => data.onDelete?.()} className="text-destructive focus:text-destructive">
          <Trash2 className="mr-2 h-4 w-4" />
          Excluir Item
        </ContextMenuItem>
      </ContextMenuContent>
    </ContextMenu>
  )
})

ItemNode.displayName = 'ItemNode'
