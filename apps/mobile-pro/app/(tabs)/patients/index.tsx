import { useState } from 'react';
import { FlatList, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { Link } from 'expo-router';
import { useInfiniteQuery } from '@tanstack/react-query';
import { options } from '@plenya/api-client';
import { Card, Input, Spinner, Text } from '@plenya/ui-mobile';
import { formatRelative } from '@plenya/domain';
import { useScreenCaptureProtection } from '../../../lib/security/screenCapture';

export default function PatientsListScreen() {
  useScreenCaptureProtection();
  const [search, setSearch] = useState('');

  const query = useInfiniteQuery(options.patientsListOptions({ search, limit: 30 }));

  const items = query.data?.pages.flatMap((p) => p.items) ?? [];

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <View className="p-4">
        <Input
          placeholder="Buscar paciente"
          value={search}
          onChangeText={setSearch}
          autoCapitalize="words"
        />
      </View>

      {query.isLoading ? (
        <Spinner centered />
      ) : (
        <FlatList
          data={items}
          keyExtractor={(item) => item.id}
          contentContainerClassName="gap-2 px-4 pb-8"
          onEndReached={() => query.hasNextPage && query.fetchNextPage()}
          onEndReachedThreshold={0.4}
          ListEmptyComponent={
            <View className="items-center py-12">
              <Text variant="caption">Nenhum paciente encontrado</Text>
            </View>
          }
          renderItem={({ item }) => (
            <Link href={`/(tabs)/patients/${item.id}`} asChild>
              <Card>
                <Text variant="title">{item.name}</Text>
                <Text variant="caption">
                  {item.phone ?? item.email ?? '—'} ·{' '}
                  {item.lastVisitAt ? `Última visita ${formatRelative(item.lastVisitAt)}` : 'Sem visitas'}
                </Text>
              </Card>
            </Link>
          )}
        />
      )}
    </SafeAreaView>
  );
}
