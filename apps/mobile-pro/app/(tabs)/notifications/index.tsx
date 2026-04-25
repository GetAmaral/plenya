import { FlatList, Pressable, RefreshControl, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { router } from 'expo-router';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { options, queryKeys, type NotificationItem } from '@plenya/api-client';
import { Button, Card, EmptyState, ErrorState, Spinner, Text, useToast } from '@plenya/ui-mobile';
import { formatRelative } from '@plenya/domain';

export default function NotificationsScreen() {
  const queryClient = useQueryClient();
  const toast = useToast();
  const list = useQuery(options.notificationsOptions());

  const markRead = useMutation({
    mutationFn: (id: string) => options.notificationMutations.markRead(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.notifications.all() });
      queryClient.invalidateQueries({ queryKey: queryKeys.notifications.unread() });
    },
  });

  const markAllRead = useMutation({
    mutationFn: () => options.notificationMutations.markAllRead(),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.notifications.all() });
      queryClient.invalidateQueries({ queryKey: queryKeys.notifications.unread() });
      toast.show('Todas marcadas como lidas', 'success');
    },
  });

  function handleOpen(n: NotificationItem) {
    if (!n.read) markRead.mutate(n.id);
    if (n.actionUrl) {
      if (n.actionUrl.startsWith('plenyapro://') || n.actionUrl.startsWith('http')) {
        // deep link externo — usa router push direto pra paths internos
        router.push(n.actionUrl as never);
      } else {
        router.push(n.actionUrl as never);
      }
    }
  }

  if (list.isLoading) return <Spinner centered />;
  if (list.isError) {
    return (
      <SafeAreaView className="flex-1 bg-background">
        <ErrorState onRetry={() => list.refetch()} />
      </SafeAreaView>
    );
  }

  const items = list.data ?? [];
  const hasUnread = items.some((n) => !n.read);

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      {hasUnread && (
        <View className="px-4 pt-3">
          <Button
            variant="outline"
            size="sm"
            onPress={() => markAllRead.mutate()}
            loading={markAllRead.isPending}
            fullWidth
          >
            Marcar todas como lidas
          </Button>
        </View>
      )}

      <FlatList
        data={items}
        keyExtractor={(item) => item.id}
        contentContainerClassName="gap-2 p-4"
        refreshControl={
          <RefreshControl refreshing={list.isRefetching} onRefresh={() => list.refetch()} />
        }
        ListEmptyComponent={
          <EmptyState title="Sem notificações" description="Você está em dia." />
        }
        renderItem={({ item }) => (
          <Pressable onPress={() => handleOpen(item)}>
            <Card className={item.read ? 'opacity-70' : ''}>
              <View className="flex-row items-start gap-2">
                {!item.read && (
                  <View className="mt-1.5 h-2 w-2 rounded-full bg-primary" />
                )}
                <View className="flex-1">
                  <Text variant="title">{item.title}</Text>
                  <Text variant="body" className="mt-0.5">
                    {item.message}
                  </Text>
                  <Text variant="caption" className="mt-1">
                    {formatRelative(item.createdAt)}
                  </Text>
                </View>
              </View>
            </Card>
          </Pressable>
        )}
      />
    </SafeAreaView>
  );
}
