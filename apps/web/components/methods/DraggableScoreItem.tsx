'use client'

import { useDraggable } from '@dnd-kit/core'
import type { ScoreItem } from '@plenya/types'
import { Card, CardContent } from '@/components/ui/card'
import { Badge } from '@/components/ui/badge'
import { Button } from '@/components/ui/button'
import { GripVertical, X, Loader2 } from 'lucide-react'
import { cn } from '@/lib/utils'
import { useUnassignItemFromPillar } from '@/lib/api/method-api'
import { toast } from 'sonner'

interface DraggableScoreItemProps {
  item: ScoreItem
  pillarId?: string
  isDragging?: boolean
}

export function DraggableScoreItem({ item, pillarId, isDragging = false }: DraggableScoreItemProps) {
  const { attributes, listeners, setNodeRef, transform } = useDraggable({
    id: item.id,
    data: { item },
  })

  const unassignMutation = useUnassignItemFromPillar()

  const style = transform
    ? {
        transform: `translate3d(${transform.x}px, ${transform.y}px, 0)`,
      }
    : undefined

  const handleUnassign = async (e: React.MouseEvent) => {
    e.stopPropagation()
    if (!pillarId) return

    try {
      await unassignMutation.mutateAsync({ itemId: item.id, pillarId })
      toast.success('Item removido do pilar')
    } catch (error) {
      toast.error('Erro ao remover item')
      console.error('Unassign error:', error)
    }
  }

  return (
    <Card
      ref={setNodeRef}
      style={style}
      className={cn(
        'transition-shadow hover:shadow-md',
        isDragging && 'shadow-xl opacity-80 cursor-grabbing',
        !isDragging && 'cursor-grab'
      )}
    >
      <CardContent className="p-3">
        <div className="flex items-start gap-2">
          {/* Drag Handle */}
          <div {...listeners} {...attributes} className="mt-1 cursor-grab active:cursor-grabbing">
            <GripVertical className="h-4 w-4 text-muted-foreground" />
          </div>

          {/* Content */}
          <div className="flex-1 space-y-2">
            <div className="flex items-start justify-between gap-2">
              <div>
                <p className="text-sm font-medium leading-tight">{item.name}</p>
                {item.unit && (
                  <p className="text-xs text-muted-foreground mt-1">Unidade: {item.unit}</p>
                )}
              </div>
              {pillarId && (
                <Button
                  variant="ghost"
                  size="sm"
                  className="h-6 w-6 p-0 hover:bg-destructive hover:text-destructive-foreground"
                  onClick={handleUnassign}
                  disabled={unassignMutation.isPending}
                >
                  {unassignMutation.isPending ? (
                    <Loader2 className="h-3 w-3 animate-spin" />
                  ) : (
                    <X className="h-3 w-3" />
                  )}
                </Button>
              )}
            </div>

            {/* Levels */}
            {item.levels && item.levels.length > 0 && (
              <div className="flex flex-wrap gap-1">
                {item.levels.slice(0, 3).map((level) => (
                  <Badge key={level.id} variant="outline" className="text-xs">
                    Nível {level.level}
                  </Badge>
                ))}
                {item.levels.length > 3 && (
                  <Badge variant="outline" className="text-xs">
                    +{item.levels.length - 3}
                  </Badge>
                )}
              </div>
            )}

            {/* Subgroup info */}
            {item.subgroup && (
              <p className="text-xs text-muted-foreground">
                {item.subgroup.group?.name} → {item.subgroup.name}
              </p>
            )}

            {/* Method Pillars badges (when viewing all items) */}
            {item.methodPillars && item.methodPillars.length > 0 && !pillarId && (
              <div className="flex flex-wrap gap-1">
                <span className="text-xs text-muted-foreground">Métodos:</span>
                {item.methodPillars.slice(0, 2).map((p) => (
                  <Badge key={p.id} variant="secondary" className="text-xs">
                    {p.name}
                  </Badge>
                ))}
                {item.methodPillars.length > 2 && (
                  <Badge variant="secondary" className="text-xs">
                    +{item.methodPillars.length - 2}
                  </Badge>
                )}
              </div>
            )}
          </div>
        </div>
      </CardContent>
    </Card>
  )
}
