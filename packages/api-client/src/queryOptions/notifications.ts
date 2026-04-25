import { queryOptions } from '@tanstack/react-query';
import { api } from '../fetcher';
import { queryKeys } from '../queryKeys';

export interface NotificationItem {
  id: string;
  type: string;
  title: string;
  message: string;
  read: boolean;
  actionUrl?: string;
  createdAt: string;
}

export const notificationsOptions = () =>
  queryOptions({
    queryKey: queryKeys.notifications.all(),
    queryFn: ({ signal }) => api.get<NotificationItem[]>('/api/v1/notifications', { signal }),
  });

export const unreadCountOptions = () =>
  queryOptions({
    queryKey: queryKeys.notifications.unread(),
    queryFn: ({ signal }) =>
      api.get<{ count: number }>('/api/v1/notifications/unread-count', { signal }),
    refetchInterval: 30_000,
  });

export const notificationMutations = {
  markRead: (id: string) => api.post<void>(`/api/v1/notifications/${id}/read`, {}),
  markAllRead: () => api.post<void>('/api/v1/notifications/read-all', {}),
};
