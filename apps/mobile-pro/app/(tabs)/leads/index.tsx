import { FlatList, RefreshControl } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { Link } from 'expo-router';
import { useQuery } from '@tanstack/react-query';
import { options } from '@plenya/api-client';
import { Card, EmptyState, ErrorState, Spinner, Text } from '@plenya/ui-mobile';
import { formatRelative } from '@plenya/domain';

export default function LeadsListScreen() {
  const query = useQuery(options.leadsListOptions());

  if (query.isLoading) return <Spinner centered />;
  if (query.isError) {
    return (
      <SafeAreaView className="flex-1 bg-background">
        <ErrorState onRetry={() => query.refetch()} />
      </SafeAreaView>
    );
  }

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <FlatList
        data={query.data ?? []}
        keyExtractor={(item) => item.id}
        contentContainerClassName="gap-2 p-4"
        refreshControl={
          <RefreshControl refreshing={query.isRefetching} onRefresh={() => query.refetch()} />
        }
        ListEmptyComponent={
          <EmptyState title="Nenhum lead" description="Você ainda não tem leads no funil." />
        }
        renderItem={({ item }) => (
          <Link href={`/(tabs)/leads/${item.id}`} asChild>
            <Card>
              <Text variant="title">{item.name}</Text>
              <Text variant="caption">
                {item.source} · {item.stage} · {formatRelative(item.lastActivityAt)}
              </Text>
            </Card>
          </Link>
        )}
      />
    </SafeAreaView>
  );
}
