import { useState } from 'react';
import { FlatList, Pressable, RefreshControl, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { Link } from 'expo-router';
import { useQuery } from '@tanstack/react-query';
import { options, type ConversationItem } from '@plenya/api-client';
import { Card, EmptyState, ErrorState, Spinner, Text } from '@plenya/ui-mobile';
import { formatRelative } from '@plenya/domain';

const CHANNEL_ICON: Record<string, string> = {
  email: '✉',
  whatsapp: '💬',
  internal: '·',
};

const FILTER_TABS: Array<{ key: string; label: string }> = [
  { key: 'all', label: 'Tudo' },
  { key: 'unread', label: 'Não lidas' },
  { key: 'mine', label: 'Minhas' },
];

export default function ConversationsScreen() {
  const [filter, setFilter] = useState<'all' | 'unread' | 'mine'>('all');
  const query = useQuery(
    options.conversationsListOptions({
      unreadOnly: filter === 'unread',
      assignedToMe: filter === 'mine',
      limit: 100,
    }),
  );

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <View className="flex-row gap-2 p-4">
        {FILTER_TABS.map((t) => (
          <Pressable
            key={t.key}
            onPress={() => setFilter(t.key as typeof filter)}
            className={`flex-1 items-center rounded-full py-2 ${
              filter === t.key ? 'bg-primary' : 'bg-muted'
            }`}
          >
            <Text
              className={`text-sm font-semibold ${
                filter === t.key ? 'text-primary-foreground' : 'text-foreground'
              }`}
            >
              {t.label}
            </Text>
          </Pressable>
        ))}
      </View>

      {query.isLoading ? (
        <Spinner centered />
      ) : query.isError ? (
        <ErrorState onRetry={() => query.refetch()} />
      ) : (
        <FlatList
          data={query.data?.items ?? []}
          keyExtractor={(item) => `${item.ownerType}:${item.ownerId}`}
          contentContainerClassName="gap-2 px-4 pb-8"
          refreshControl={
            <RefreshControl refreshing={query.isRefetching} onRefresh={() => query.refetch()} />
          }
          ListEmptyComponent={
            <EmptyState
              title="Sem conversas"
              description={
                filter === 'unread'
                  ? 'Nada não lido por aqui.'
                  : 'Nenhuma conversa ativa.'
              }
            />
          }
          renderItem={({ item }) => <ConversationRow item={item} />}
        />
      )}
    </SafeAreaView>
  );
}

function ConversationRow({ item }: { item: ConversationItem }) {
  return (
    <Link href={`/(tabs)/conversations/${item.ownerType}/${item.ownerId}`} asChild>
      <Card className={item.unreadCount > 0 ? 'border-primary' : ''}>
        <View className="flex-row items-start gap-2">
          <Text className="text-base">{CHANNEL_ICON[item.lastChannel] ?? '·'}</Text>
          <View className="flex-1">
            <View className="flex-row items-center justify-between">
              <Text variant="title" numberOfLines={1} className="flex-1 pr-2">
                {item.name}
              </Text>
              <Text variant="caption">{formatRelative(item.lastAt)}</Text>
            </View>
            <Text variant="caption" numberOfLines={2} className="mt-0.5">
              {item.lastDirection === 'in' ? '↓ ' : '↑ '}
              {item.lastSnippet}
            </Text>
            <View className="mt-1 flex-row items-center gap-2">
              <View
                className={`rounded-full px-2 py-0.5 ${
                  item.ownerType === 'patient' ? 'bg-emerald-100' : 'bg-amber-100'
                }`}
              >
                <Text
                  className={`text-[10px] font-semibold uppercase ${
                    item.ownerType === 'patient' ? 'text-emerald-800' : 'text-amber-800'
                  }`}
                >
                  {item.ownerType}
                </Text>
              </View>
              {item.unreadCount > 0 && (
                <View className="rounded-full bg-primary px-2 py-0.5">
                  <Text className="text-[10px] font-semibold text-primary-foreground">
                    {item.unreadCount}
                  </Text>
                </View>
              )}
            </View>
          </View>
        </View>
      </Card>
    </Link>
  );
}
