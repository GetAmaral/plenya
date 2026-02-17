import { Badge } from '@/components/ui/badge'
import { cn } from '@/lib/utils'

interface SubscriptionStatusBadgeProps {
  status: 'active' | 'inactive' | 'cancelled' | 'expired' | 'suspended' | 'trial'
  className?: string
}

const statusConfig = {
  active: {
    label: 'Ativo',
    className: 'bg-green-500/10 text-green-700 border-green-500/20 hover:bg-green-500/20',
  },
  trial: {
    label: 'Trial',
    className: 'bg-blue-500/10 text-blue-700 border-blue-500/20 hover:bg-blue-500/20',
  },
  inactive: {
    label: 'Inativo',
    className: 'bg-gray-500/10 text-gray-700 border-gray-500/20 hover:bg-gray-500/20',
  },
  cancelled: {
    label: 'Cancelado',
    className: 'bg-red-500/10 text-red-700 border-red-500/20 hover:bg-red-500/20',
  },
  expired: {
    label: 'Expirado',
    className: 'bg-orange-500/10 text-orange-700 border-orange-500/20 hover:bg-orange-500/20',
  },
  suspended: {
    label: 'Suspenso',
    className: 'bg-yellow-500/10 text-yellow-700 border-yellow-500/20 hover:bg-yellow-500/20',
  },
}

export function SubscriptionStatusBadge({ status, className }: SubscriptionStatusBadgeProps) {
  const config = statusConfig[status]

  return (
    <Badge variant="outline" className={cn(config.className, className)}>
      {config.label}
    </Badge>
  )
}
