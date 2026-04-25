import { useState } from 'react';
import { FlatList, RefreshControl, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { Link } from 'expo-router';
import { useQuery } from '@tanstack/react-query';
import { options } from '@plenya/api-client';
import { Card, EmptyState, ErrorState, Input, Spinner, Text } from '@plenya/ui-mobile';

export default function ExerciseLibraryScreen() {
  const [search, setSearch] = useState('');
  const query = useQuery(options.exercisesListOptions({ search }));

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <View className="p-4">
        <Input
          placeholder="Buscar exercício"
          value={search}
          onChangeText={setSearch}
          autoCapitalize="none"
        />
      </View>

      {query.isLoading ? (
        <Spinner centered />
      ) : query.isError ? (
        <ErrorState onRetry={() => query.refetch()} />
      ) : (
        <FlatList
          data={query.data ?? []}
          keyExtractor={(item) => item.id}
          contentContainerClassName="gap-2 px-4 pb-8"
          refreshControl={
            <RefreshControl refreshing={query.isRefetching} onRefresh={() => query.refetch()} />
          }
          ListEmptyComponent={
            <EmptyState
              title="Nenhum exercício"
              description={search ? 'Tente outro termo.' : 'A biblioteca está vazia.'}
            />
          }
          renderItem={({ item }) => (
            <Link href={`/(tabs)/training/exercises/${item.id}`} asChild>
              <Card>
                <Text variant="title">{item.name}</Text>
                <Text variant="caption">
                  {item.category ?? '—'}
                  {item.muscleGroups && item.muscleGroups.length > 0
                    ? ` · ${item.muscleGroups.slice(0, 3).join(', ')}`
                    : ''}
                </Text>
              </Card>
            </Link>
          )}
        />
      )}
    </SafeAreaView>
  );
}
