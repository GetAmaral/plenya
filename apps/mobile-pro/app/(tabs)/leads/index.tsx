import { FlatList, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { Link } from 'expo-router';
import { useQuery } from '@tanstack/react-query';
import { options } from '@plenya/api-client';
import { Card, Spinner, Text } from '@plenya/ui-mobile';
import { formatRelative } from '@plenya/domain';

export default function LeadsListScreen() {
  const query = useQuery(options.leadsListOptions());

  if (query.isLoading) return <Spinner centered />;

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <FlatList
        data={query.data ?? []}
        keyExtractor={(item) => item.id}
        contentContainerClassName="gap-2 p-4"
        ListEmptyComponent={
          <View className="items-center py-12">
            <Text variant="caption">Nenhum lead no momento</Text>
          </View>
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
