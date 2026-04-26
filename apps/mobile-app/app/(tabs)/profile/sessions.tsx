import { Pressable, ScrollView, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { options, queryKeys } from '@plenya/api-client';
import { Card, EmptyState, Spinner, Text, useToast } from '@plenya/ui-mobile';
import { formatDateTime } from '@plenya/domain';

export default function SessionsScreen() {
  const qc = useQueryClient();
  const toast = useToast();
  const sessions = useQuery(options.meSessionsOptions());

  const revoke = useMutation({
    mutationFn: (id: string) => options.meMutations.revokeSession(id),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: queryKeys.meSessions() });
      toast.show('Sessão revogada', 'success');
    },
    onError: (err) =>
      toast.show(err instanceof Error ? err.message : 'Erro', 'error'),
  });

  if (sessions.isLoading) return <Spinner centered />;

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <ScrollView contentContainerClassName="gap-3 p-4">
        <Text variant="caption">
          Cada dispositivo conectado aparece aqui. Você pode revogar acesso de
          qualquer um (exceto o atual).
        </Text>

        {(sessions.data ?? []).length === 0 ? (
          <EmptyState title="Sem sessões ativas" />
        ) : (
          (sessions.data ?? []).map((s) => (
            <Card key={s.id}>
              <View className="flex-row items-start justify-between gap-3">
                <View className="flex-1">
                  <Text variant="body" className="font-semibold">
                    {s.device || 'Dispositivo'}
                  </Text>
                  <Text variant="caption">
                    {s.platform} · {s.appVariant}
                    {s.appVersion ? ` · v${s.appVersion}` : ''}
                  </Text>
                  <Text variant="caption">
                    Última atividade: {formatDateTime(s.lastSeenAt)}
                  </Text>
                  {s.current && (
                    <Text variant="caption" className="font-semibold text-primary">
                      Sessão atual
                    </Text>
                  )}
                </View>
                {!s.current && (
                  <Pressable
                    onPress={() => revoke.mutate(s.id)}
                    disabled={revoke.isPending}
                    className="rounded-md bg-red-500 px-3 py-2"
                    accessibilityRole="button"
                    accessibilityLabel="Revogar sessão"
                  >
                    <Text className="text-xs font-semibold text-primary-foreground">
                      Revogar
                    </Text>
                  </Pressable>
                )}
              </View>
            </Card>
          ))
        )}
      </ScrollView>
    </SafeAreaView>
  );
}
