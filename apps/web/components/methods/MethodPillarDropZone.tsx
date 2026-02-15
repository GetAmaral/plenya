import { useDroppable } from '@dnd-kit/core'
import { cn } from '@/lib/utils'

interface MethodPillarDropZoneProps {
  pillarId: string
  children: React.ReactNode
}

export function MethodPillarDropZone({ pillarId, children }: MethodPillarDropZoneProps) {
  const { setNodeRef, isOver } = useDroppable({ id: pillarId })

  return (
    <div
      ref={setNodeRef}
      className={cn(
        'min-h-[80px] p-3 rounded-lg transition-all',
        isOver && 'bg-primary/10 ring-2 ring-primary ring-offset-2'
      )}
    >
      {children}
    </div>
  )
}
