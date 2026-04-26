import { useEffect, useState } from 'react';
import { Pressable, ScrollView, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { router } from 'expo-router';
import { useQueries } from '@tanstack/react-query';
import { api, options } from '@plenya/api-client';
import { Card, EmptyState, Input, Spinner, Text } from '@plenya/ui-mobile';

/**
 * Busca global em 3 buckets paralelos: pacientes, leads, exercícios.
 * Debounce simples via setTimeout (sem dep extra). Mínimo 2 chars.
 */
export default function SearchScreen() {
  const [query, setQuery] = useState('');
  const [debounced, setDebounced] = useState('');

  useEffect(() => {
    const id = setTimeout(() => setDebounced(query.trim()), 300);
    return () => clearTimeout(id);
  }, [query]);

  const enabled = debounced.length >= 2;

  const [patients, leads, exercises] = useQueries({
    queries: [
      {
        queryKey: ['plenya', 'search', 'patients', debounced] as const,
        queryFn: ({ signal }: { signal: AbortSignal }) =>
          api
            .get<{ items: Array<{ id: string; name: string; phone?: string }> }>(
              `/api/v1/patients?search=${encodeURIComponent(debounced)}&limit=10`,
              { signal },
            )
            .then((r) => r.items ?? []),
        enabled,
      },
      {
        ...options.leadsListOptions({ search: debounced, pageSize: 10 }),
        select: (data: { items: Array<{ id: string; name?: string; email?: string; phone?: string }> }) =>
          data.items ?? [],
        enabled,
      },
      {
        ...options.exercisesListOptions({ search: debounced }),
        enabled,
      },
    ],
  });

  const isLoading =
    enabled && (patients.isLoading || leads.isLoading || exercises.isLoading);

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <View className="p-4">
        <Input
          value={query}
          onChangeText={setQuery}
          placeholder="Buscar paciente, exercício, lead..."
          autoFocus
          autoCapitalize="none"
          accessibilityLabel="Campo de busca global"
        />
      </View>

      <ScrollView contentContainerClassName="gap-4 px-4 pb-8">
        {!enabled ? (
          <EmptyState
            title="Digite ao menos 2 caracteres"
            description="Buscamos em pacientes, exercícios e (em breve) leads."
          />
        ) : isLoading ? (
          <Spinner centered />
        ) : (
          <>
            <Section title="Pacientes" count={patients.data?.length ?? 0}>
              {(patients.data ?? []).map((p) => (
                <Pressable
                  key={p.id}
                  onPress={() => {
                    options.patientMutations.setSelected(p.id).catch(() => {});
                    router.push(`/(tabs)/patients/${p.id}` as never);
                  }}
                  accessibilityRole="button"
                  accessibilityLabel={`Abrir paciente ${p.name}`}
                >
                  <Card>
                    <Text variant="body">{p.name}</Text>
                    {p.phone && <Text variant="caption">{p.phone}</Text>}
                  </Card>
                </Pressable>
              ))}
            </Section>

            <Section title="Leads" count={leads.data?.length ?? 0}>
              {(leads.data ?? []).map((l) => (
                <Pressable
                  key={l.id}
                  onPress={() => router.push(`/(tabs)/leads/${l.id}` as never)}
                  accessibilityRole="button"
                  accessibilityLabel={`Abrir lead ${l.name ?? l.email ?? l.phone}`}
                >
                  <Card>
                    <Text variant="body">{l.name ?? l.email ?? l.phone ?? 'Lead'}</Text>
                    {(l.phone || l.email) && (
                      <Text variant="caption">{l.phone ?? l.email}</Text>
                    )}
                  </Card>
                </Pressable>
              ))}
            </Section>

            <Section title="Exercícios" count={exercises.data?.length ?? 0}>
              {(exercises.data ?? []).map((ex) => (
                <Pressable
                  key={ex.id}
                  onPress={() => router.push(`/(tabs)/training/exercises/${ex.id}` as never)}
                  accessibilityRole="button"
                  accessibilityLabel={`Abrir exercício ${ex.name}`}
                >
                  <Card>
                    <Text variant="body">{ex.name}</Text>
                    <Text variant="caption">{ex.category ?? '—'}</Text>
                  </Card>
                </Pressable>
              ))}
            </Section>
          </>
        )}
      </ScrollView>
    </SafeAreaView>
  );
}

function Section({
  title,
  count,
  children,
}: {
  title: string;
  count: number;
  children: React.ReactNode;
}) {
  return (
    <View className="gap-2">
      <Text variant="caption" className="font-semibold uppercase">
        {title} {count > 0 ? `(${count})` : ''}
      </Text>
      {count === 0 ? (
        <Text variant="caption">Sem resultados nesta categoria.</Text>
      ) : (
        children
      )}
    </View>
  );
}
