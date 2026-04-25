import { ScrollView, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { useLocalSearchParams } from 'expo-router';
import { useQuery } from '@tanstack/react-query';
import { options } from '@plenya/api-client';
import { Card, CardHeader, Spinner, Text } from '@plenya/ui-mobile';
import { formatRelative } from '@plenya/domain';

export default function LeadDetailScreen() {
  const { id } = useLocalSearchParams<{ id: string }>();
  const lead = useQuery(options.leadOptions(id ?? ''));

  if (lead.isLoading) return <Spinner centered />;
  if (!lead.data) {
    return (
      <SafeAreaView className="flex-1 bg-background">
        <View className="p-6">
          <Text variant="caption">Lead não encontrado</Text>
        </View>
      </SafeAreaView>
    );
  }

  const l = lead.data;

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <ScrollView contentContainerClassName="gap-3 p-4">
        <Text variant="heading">{l.name}</Text>
        <Text variant="caption">
          {l.source} · estágio {l.stage}
        </Text>

        <Card>
          <CardHeader>
            <Text variant="title">Contato</Text>
          </CardHeader>
          <Text variant="body">{l.phone ?? l.email ?? '—'}</Text>
        </Card>

        <Card>
          <CardHeader>
            <Text variant="title">Timeline</Text>
          </CardHeader>
          <View className="gap-3">
            {l.activities.map((a) => (
              <View key={a.id}>
                <Text variant="body">{a.content}</Text>
                <Text variant="caption">
                  {a.kind} · {formatRelative(a.createdAt)}
                  {a.actorName ? ` · ${a.actorName}` : ''}
                </Text>
              </View>
            ))}
            {l.activities.length === 0 && <Text variant="caption">Sem atividades</Text>}
          </View>
        </Card>
      </ScrollView>
    </SafeAreaView>
  );
}
