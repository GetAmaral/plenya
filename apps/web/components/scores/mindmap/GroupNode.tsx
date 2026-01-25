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

export const GroupNode = memo(({ data }: NodeProps) => {
  const hasChildren = (data.subgroupCount || 0) > 0
  const isExpanded = data.isExpanded || false

  return (
    <ContextMenu>
      <ContextMenuTrigger asChild>
        <div className="w-[336px] px-4 py-3 rounded-lg bg-primary text-primary-foreground shadow-lg border-2 border-primary">
          <Handle type="source" position={Position.Right} className="w-3 h-3" />

          <div className="space-y-2">
            <div className="flex items-start gap-2">
              {hasChildren && (
                <button
                  onClick={(e) => {
                    e.stopPropagation()
                    data.onToggle?.()
                  }}
                  className="hover:bg-primary-foreground/20 rounded p-1 transition-colors shrink-0 mt-0.5"
                  title={isExpanded ? 'Recolher' : 'Expandir'}
                >
                  {isExpanded ? (
                    <ChevronDown className="h-4 w-4" />
                  ) : (
                    <ChevronRight className="h-4 w-4" />
                  )}
                </button>
              )}
              <h3 className="font-bold text-base leading-tight break-words flex-1">{data.name}</h3>
            </div>

            <div className="flex items-center justify-between gap-2 text-xs opacity-90">
              <span>{data.subgroupCount || 0} subgrupo{data.subgroupCount !== 1 ? 's' : ''}</span>
              <Badge variant="secondary" className="bg-primary-foreground/20 text-primary-foreground shrink-0">
                #{data.order}
              </Badge>
            </div>
          </div>
        </div>
      </ContextMenuTrigger>
      <ContextMenuContent>
        <ContextMenuItem onClick={() => data.onEdit?.()}>
          <Edit className="mr-2 h-4 w-4" />
          Editar Grupo
        </ContextMenuItem>
        <ContextMenuItem onClick={() => data.onAdd?.()}>
          <Plus className="mr-2 h-4 w-4" />
          Adicionar Subgrupo
        </ContextMenuItem>
        <ContextMenuSeparator />
        <ContextMenuItem onClick={() => data.onDelete?.()} className="text-destructive focus:text-destructive">
          <Trash2 className="mr-2 h-4 w-4" />
          Excluir Grupo
        </ContextMenuItem>
      </ContextMenuContent>
    </ContextMenu>
  )
})

GroupNode.displayName = 'GroupNode'
