import { ScrollView, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { useLocalSearchParams } from 'expo-router';
import { useQuery } from '@tanstack/react-query';
import { options, type LabResultValue } from '@plenya/api-client';
import { Card, CardHeader, ErrorState, Spinner, Text } from '@plenya/ui-mobile';
import { formatDate } from '@plenya/domain';
import { useScreenCaptureProtection } from '../../../../../lib/security/screenCapture';

const FLAG_COLOR: Record<NonNullable<LabResultValue['flag']>, string> = {
  low: 'text-amber-500',
  normal: 'text-emerald-600',
  high: 'text-amber-500',
  critical: 'text-destructive',
};

const FLAG_LABEL: Record<NonNullable<LabResultValue['flag']>, string> = {
  low: '↓ baixo',
  normal: 'normal',
  high: '↑ alto',
  critical: 'crítico',
};

export default function LabResultDetailScreen() {
  useScreenCaptureProtection();
  const { labResultId } = useLocalSearchParams<{ labResultId: string }>();
  const query = useQuery(options.labResultOptions(labResultId ?? ''));

  if (query.isLoading) return <Spinner centered />;
  if (query.isError) {
    return (
      <SafeAreaView className="flex-1 bg-background">
        <ErrorState onRetry={() => query.refetch()} />
      </SafeAreaView>
    );
  }
  if (!query.data) return null;

  const r = query.data;

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <ScrollView contentContainerClassName="gap-3 p-4">
        <View>
          <Text variant="heading">{r.labName ?? 'Exame laboratorial'}</Text>
          <Text variant="caption">Coletado em {formatDate(r.collectedAt)}</Text>
        </View>

        {r.notes && (
          <Card>
            <Text variant="body">{r.notes}</Text>
          </Card>
        )}

        <Card>
          <CardHeader>
            <Text variant="title">Valores</Text>
          </CardHeader>
          {r.values.length === 0 ? (
            <Text variant="caption">Sem valores registrados.</Text>
          ) : (
            <View className="gap-2">
              {r.values.map((v, idx) => (
                <View
                  key={`${v.testCode}-${idx}`}
                  className="flex-row items-center justify-between rounded-md bg-muted px-3 py-2"
                >
                  <View className="flex-1 pr-3">
                    <Text variant="body">{v.testName}</Text>
                    {(v.referenceMin !== undefined || v.referenceMax !== undefined) && (
                      <Text variant="caption">
                        Ref:{' '}
                        {v.referenceMin !== undefined ? v.referenceMin : '—'}
                        {' a '}
                        {v.referenceMax !== undefined ? v.referenceMax : '—'} {v.unit ?? ''}
                      </Text>
                    )}
                  </View>
                  <View className="items-end">
                    <Text variant="body">
                      {String(v.value)} {v.unit ?? ''}
                    </Text>
                    {v.flag && (
                      <Text className={`text-xs font-semibold ${FLAG_COLOR[v.flag]}`}>
                        {FLAG_LABEL[v.flag]}
                      </Text>
                    )}
                  </View>
                </View>
              ))}
            </View>
          )}
        </Card>

        {r.attachments && r.attachments.length > 0 && (
          <Card>
            <CardHeader>
              <Text variant="title">Anexos</Text>
            </CardHeader>
            <View className="gap-1">
              {r.attachments.map((a) => (
                <Text key={a.id} variant="caption">
                  • {a.type}
                </Text>
              ))}
            </View>
          </Card>
        )}
      </ScrollView>
    </SafeAreaView>
  );
}
