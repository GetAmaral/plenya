import { Linking, ScrollView, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { useLocalSearchParams } from 'expo-router';
import { useQuery } from '@tanstack/react-query';
import { options } from '@plenya/api-client';
import { Button, Card, CardHeader, ErrorState, Spinner, Text } from '@plenya/ui-mobile';
import { formatDate } from '@plenya/domain';
import { useScreenCaptureProtection } from '../../../../../lib/security/screenCapture';

const STATUS_LABEL: Record<string, string> = {
  draft: 'Rascunho',
  signed: 'Assinada',
  cancelled: 'Cancelada',
};

const STATUS_COLOR: Record<string, string> = {
  draft: 'text-muted-foreground',
  signed: 'text-emerald-600',
  cancelled: 'text-destructive',
};

export default function PrescriptionDetailScreen() {
  useScreenCaptureProtection();
  const { prescriptionId } = useLocalSearchParams<{ prescriptionId: string }>();
  const query = useQuery(options.prescriptionOptions(prescriptionId ?? ''));

  if (query.isLoading) return <Spinner centered />;
  if (query.isError) {
    return (
      <SafeAreaView className="flex-1 bg-background">
        <ErrorState onRetry={() => query.refetch()} />
      </SafeAreaView>
    );
  }
  if (!query.data) return null;

  const p = query.data;

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <ScrollView contentContainerClassName="gap-3 p-4">
        <View>
          <Text variant="heading">Prescrição</Text>
          <View className="flex-row items-center gap-2">
            <Text variant="caption">Emitida em {formatDate(p.issuedAt)}</Text>
            <Text className={`text-xs font-semibold ${STATUS_COLOR[p.status] ?? ''}`}>
              · {STATUS_LABEL[p.status] ?? p.status}
            </Text>
          </View>
          {p.signedAt && (
            <Text variant="caption">Assinada em {formatDate(p.signedAt)}</Text>
          )}
        </View>

        {p.notes && (
          <Card>
            <Text variant="body">{p.notes}</Text>
          </Card>
        )}

        <Card>
          <CardHeader>
            <Text variant="title">Medicamentos</Text>
          </CardHeader>
          {p.items.length === 0 ? (
            <Text variant="caption">Sem itens.</Text>
          ) : (
            <View className="gap-3">
              {p.items.map((item, idx) => (
                <View key={idx} className="rounded-md bg-muted px-3 py-2">
                  <Text variant="body">{item.name}</Text>
                  <Text variant="caption">
                    {item.dosage} · {item.frequency}
                    {item.duration ? ` · ${item.duration}` : ''}
                  </Text>
                  {item.notes && (
                    <Text variant="caption" className="mt-1 italic">
                      {item.notes}
                    </Text>
                  )}
                </View>
              ))}
            </View>
          )}
        </Card>

        {p.pdfUrl && (
          <Button
            onPress={() => {
              if (!p.pdfUrl) return;
              Linking.openURL(p.pdfUrl).catch(() => {});
            }}
            fullWidth
          >
            Abrir PDF assinado
          </Button>
        )}
      </ScrollView>
    </SafeAreaView>
  );
}
