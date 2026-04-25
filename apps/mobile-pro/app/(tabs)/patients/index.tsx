import { useState } from 'react';
import { FlatList, Pressable, RefreshControl, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { Link, router } from 'expo-router';
import { useInfiniteQuery } from '@tanstack/react-query';
import { options } from '@plenya/api-client';
import { Button, Card, EmptyState, ErrorState, Input, Spinner, Text } from '@plenya/ui-mobile';
import { formatRelative } from '@plenya/domain';
import { useScreenCaptureProtection } from '../../../lib/security/screenCapture';

export default function PatientsListScreen() {
  useScreenCaptureProtection();
  const [search, setSearch] = useState('');

  const query = useInfiniteQuery(options.patientsListOptions({ search, limit: 30 }));

  const items = query.data?.pages.flatMap((p) => p.items) ?? [];

  function handleSelect(patientId: string) {
    // Fire-and-forget: setSelected garante que endpoints scope-by-selected
    // (lab-results, prescriptions etc) achem o paciente certo já no detail.
    options.patientMutations.setSelected(patientId).catch(() => {});
    router.push(`/(tabs)/patients/${patientId}`);
  }

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <View className="gap-3 p-4">
        <Input
          placeholder="Buscar paciente"
          value={search}
          onChangeText={setSearch}
          autoCapitalize="words"
        />
        <Link href="/(tabs)/patients/new" asChild>
          <Button variant="outline" size="sm" fullWidth>
            + Novo paciente
          </Button>
        </Link>
      </View>

      {query.isLoading ? (
        <Spinner centered />
      ) : query.isError ? (
        <ErrorState onRetry={() => query.refetch()} />
      ) : (
        <FlatList
          data={items}
          keyExtractor={(item) => item.id}
          contentContainerClassName="gap-2 px-4 pb-8"
          onEndReached={() => query.hasNextPage && query.fetchNextPage()}
          onEndReachedThreshold={0.4}
          refreshControl={
            <RefreshControl
              refreshing={query.isRefetching && !query.isFetchingNextPage}
              onRefresh={() => query.refetch()}
            />
          }
          ListEmptyComponent={
            <EmptyState
              title="Nenhum paciente"
              description={search ? 'Sem resultados para a busca.' : 'Você ainda não tem pacientes cadastrados.'}
            />
          }
          renderItem={({ item }) => (
            <Pressable onPress={() => handleSelect(item.id)}>
              <Card>
                <Text variant="title">{item.name}</Text>
                <Text variant="caption">
                  {item.phone ?? item.email ?? '—'} ·{' '}
                  {item.lastVisitAt ? `Última visita ${formatRelative(item.lastVisitAt)}` : 'Sem visitas'}
                </Text>
              </Card>
            </Pressable>
          )}
        />
      )}
    </SafeAreaView>
  );
}
