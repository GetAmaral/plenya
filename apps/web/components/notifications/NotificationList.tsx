'use client'

import { useNotifications, useMarkAsRead, useMarkAllAsRead, useDeleteNotification } from '@/lib/api/notification-api'
import { Button } from '@/components/ui/button'
import { ScrollArea } from '@/components/ui/scroll-area'
import { Separator } from '@/components/ui/separator'
import { CheckCheck, X, Trash2, Bell } from 'lucide-react'
import { toast } from 'sonner'
import { formatDistanceToNow } from 'date-fns'
import { ptBR } from 'date-fns/locale'
import Link from 'next/link'
import { cn } from '@/lib/utils'
import type { Notification } from '@plenya/types'

interface NotificationListProps {
  onClose: () => void
}

const typeColors = {
  trial_expiring: 'bg-blue-500/10 border-blue-500/20',
  renewal_upcoming: 'bg-green-500/10 border-green-500/20',
  subscription_expired: 'bg-orange-500/10 border-orange-500/20',
  payment_pending: 'bg-yellow-500/10 border-yellow-500/20',
  general: 'bg-gray-500/10 border-gray-500/20',
}

export function NotificationList({ onClose }: NotificationListProps) {
  const { data: notifications, isLoading } = useNotifications(20)
  const markAsRead = useMarkAsRead()
  const markAllAsRead = useMarkAllAsRead()
  const deleteNotification = useDeleteNotification()

  const handleMarkAsRead = async (id: string) => {
    try {
      await markAsRead.mutateAsync(id)
    } catch (error) {
      toast.error('Erro ao marcar como lida')
    }
  }

  const handleMarkAllAsRead = async () => {
    try {
      await markAllAsRead.mutateAsync()
      toast.success('Todas as notificações marcadas como lidas')
    } catch (error) {
      toast.error('Erro ao marcar todas como lidas')
    }
  }

  const handleDelete = async (id: string, e: React.MouseEvent) => {
    e.stopPropagation()
    e.preventDefault()
    try {
      await deleteNotification.mutateAsync(id)
      toast.success('Notificação removida')
    } catch (error) {
      toast.error('Erro ao remover notificação')
    }
  }

  const handleNotificationClick = (notification: Notification) => {
    if (!notification.read) {
      handleMarkAsRead(notification.id)
    }
    if (notification.actionUrl) {
      onClose()
    }
  }

  const unreadCount = notifications?.filter((n) => !n.read).length || 0

  if (isLoading) {
    return (
      <div className="p-4">
        <div className="space-y-3">
          {[1, 2, 3].map((i) => (
            <div key={i} className="h-16 bg-muted animate-pulse rounded" />
          ))}
        </div>
      </div>
    )
  }

  return (
    <div className="flex flex-col h-full">
      {/* Header */}
      <div className="p-3 border-b flex items-center justify-between">
        <div>
          <h3 className="font-semibold">Notificações</h3>
          {unreadCount > 0 && (
            <p className="text-xs text-muted-foreground">{unreadCount} não lida{unreadCount !== 1 && 's'}</p>
          )}
        </div>
        {unreadCount > 0 && (
          <Button
            variant="ghost"
            size="sm"
            onClick={handleMarkAllAsRead}
            className="h-8 text-xs"
          >
            <CheckCheck className="h-3 w-3 mr-1" />
            Marcar todas
          </Button>
        )}
      </div>

      {/* List */}
      <ScrollArea className="h-[400px]">
        {notifications && notifications.length > 0 ? (
          <div className="divide-y">
            {notifications.map((notification) => {
              const content = (
                <div
                  className={cn(
                    'p-3 hover:bg-accent cursor-pointer transition-colors relative',
                    !notification.read && 'bg-primary/5'
                  )}
                  onClick={() => handleNotificationClick(notification)}
                >
                  {!notification.read && (
                    <div className="absolute top-3 right-3 w-2 h-2 bg-primary rounded-full" />
                  )}

                  <div className={cn(
                    'rounded-lg p-3 border',
                    typeColors[notification.type]
                  )}>
                    <div className="flex items-start justify-between gap-2 mb-1">
                      <h4 className="font-medium text-sm">{notification.title}</h4>
                      <Button
                        variant="ghost"
                        size="icon"
                        className="h-6 w-6 -mt-1 -mr-1"
                        onClick={(e) => handleDelete(notification.id, e)}
                      >
                        <X className="h-3 w-3" />
                      </Button>
                    </div>
                    <p className="text-sm text-muted-foreground mb-2">
                      {notification.message}
                    </p>
                    <div className="flex items-center justify-between">
                      <span className="text-xs text-muted-foreground">
                        {formatDistanceToNow(new Date(notification.createdAt), {
                          addSuffix: true,
                          locale: ptBR,
                        })}
                      </span>
                      {notification.actionText && notification.actionUrl && (
                        <span className="text-xs text-primary font-medium">
                          {notification.actionText} →
                        </span>
                      )}
                    </div>
                  </div>
                </div>
              )

              if (notification.actionUrl) {
                return (
                  <Link key={notification.id} href={notification.actionUrl}>
                    {content}
                  </Link>
                )
              }

              return <div key={notification.id}>{content}</div>
            })}
          </div>
        ) : (
          <div className="flex flex-col items-center justify-center py-12 px-4">
            <Bell className="h-12 w-12 text-muted-foreground/50 mb-3" />
            <p className="text-sm text-muted-foreground text-center">
              Nenhuma notificação
            </p>
          </div>
        )}
      </ScrollArea>
    </div>
  )
}
