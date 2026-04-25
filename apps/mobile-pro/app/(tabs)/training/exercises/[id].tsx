import { ScrollView, View } from 'react-native';
import { Image } from 'expo-image';
import { SafeAreaView } from 'react-native-safe-area-context';
import { useLocalSearchParams } from 'expo-router';
import { useQuery } from '@tanstack/react-query';
import { options } from '@plenya/api-client';
import { Card, CardHeader, ErrorState, Spinner, Text } from '@plenya/ui-mobile';

export default function ExerciseDetailScreen() {
  const { id } = useLocalSearchParams<{ id: string }>();
  const exerciseId = id ?? '';
  const query = useQuery(options.exerciseDetailOptions(exerciseId));

  if (query.isLoading) return <Spinner centered />;
  if (query.isError) {
    return (
      <SafeAreaView className="flex-1 bg-background">
        <ErrorState onRetry={() => query.refetch()} />
      </SafeAreaView>
    );
  }

  const e = query.data;
  if (!e) return null;

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <ScrollView contentContainerClassName="gap-3 p-4">
        {(e.gifUrl ?? e.thumbnailUrl) && (
          <View className="overflow-hidden rounded-xl bg-muted">
            <Image
              source={{ uri: e.gifUrl ?? e.thumbnailUrl }}
              style={{ width: '100%', height: 220 }}
              contentFit="cover"
            />
          </View>
        )}

        <Text variant="heading">{e.name}</Text>
        {e.category && <Text variant="caption">{e.category}</Text>}

        {e.description && (
          <Card>
            <CardHeader>
              <Text variant="title">Descrição</Text>
            </CardHeader>
            <Text variant="body">{e.description}</Text>
          </Card>
        )}

        {e.muscleGroups && e.muscleGroups.length > 0 && (
          <Card>
            <CardHeader>
              <Text variant="title">Grupos musculares</Text>
            </CardHeader>
            <View className="flex-row flex-wrap gap-2">
              {e.muscleGroups.map((g) => (
                <View key={g} className="rounded-full bg-secondary px-3 py-1">
                  <Text className="text-xs text-secondary-foreground">{g}</Text>
                </View>
              ))}
            </View>
          </Card>
        )}

        {e.equipment && e.equipment.length > 0 && (
          <Card>
            <CardHeader>
              <Text variant="title">Equipamentos</Text>
            </CardHeader>
            <View className="flex-row flex-wrap gap-2">
              {e.equipment.map((q) => (
                <View key={q} className="rounded-full bg-secondary px-3 py-1">
                  <Text className="text-xs text-secondary-foreground">{q}</Text>
                </View>
              ))}
            </View>
          </Card>
        )}

        {e.nscaReferences && e.nscaReferences.length > 0 && (
          <Card>
            <CardHeader>
              <Text variant="title">Referências NSCA</Text>
            </CardHeader>
            <View className="gap-2">
              {e.nscaReferences.map((r, idx) => (
                <Text key={idx} variant="caption">
                  • {r.citation}
                </Text>
              ))}
            </View>
          </Card>
        )}
      </ScrollView>
    </SafeAreaView>
  );
}
