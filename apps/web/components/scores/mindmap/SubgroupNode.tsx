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

export const SubgroupNode = memo(({ data }: NodeProps) => {
  const hasChildren = (data.itemCount || 0) > 0
  const isExpanded = data.isExpanded || false

  return (
    <ContextMenu>
      <ContextMenuTrigger asChild>
        <div className="w-[336px] px-4 py-3 rounded-md bg-card border-2 border-primary/30 shadow-md">
          <Handle type="target" position={Position.Left} id="left" className="w-3 h-3" />
          <Handle type="source" position={Position.Right} id="right" className="w-3 h-3" />

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
              <div className="font-semibold text-sm leading-tight break-words flex-1">{data.name}</div>
            </div>

            <div className="flex items-center gap-2 text-xs text-muted-foreground">
              <Badge variant="outline" className="text-xs">
                {data.itemCount || 0} item{data.itemCount !== 1 ? 'ns' : ''}
              </Badge>
              <span>#{data.order}</span>
            </div>
          </div>
        </div>
      </ContextMenuTrigger>
      <ContextMenuContent>
        <ContextMenuItem onClick={() => data.onEdit?.(data.groupId)}>
          <Edit className="mr-2 h-4 w-4" />
          Editar Subgrupo
        </ContextMenuItem>
        <ContextMenuItem onClick={() => data.onAdd?.()}>
          <Plus className="mr-2 h-4 w-4" />
          Adicionar Item
        </ContextMenuItem>
        <ContextMenuSeparator />
        <ContextMenuItem onClick={() => data.onDelete?.()} className="text-destructive focus:text-destructive">
          <Trash2 className="mr-2 h-4 w-4" />
          Excluir Subgrupo
        </ContextMenuItem>
      </ContextMenuContent>
    </ContextMenu>
  )
})

SubgroupNode.displayName = 'SubgroupNode'
