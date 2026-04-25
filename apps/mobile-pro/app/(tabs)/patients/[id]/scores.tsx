import { useMemo } from 'react';
import { Pressable, RefreshControl, ScrollView, Share, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { useLocalSearchParams } from 'expo-router';
import { useQuery } from '@tanstack/react-query';
import {
  options,
  queryKeys,
  type ScoreGroupNode,
  type ScoreItemNode,
  type ScoreLevelNode,
  type PatientScoreResult,
} from '@plenya/api-client';
import { Card, EmptyState, ErrorState, Spinner, Text } from '@plenya/ui-mobile';
import { matchLevel } from '@plenya/domain';
import { useScreenCaptureProtection } from '../../../../lib/security/screenCapture';
import { useRefresh } from '../../../../features/patients/usePatientRefresh';
import { useEnsureSelectedPatient } from '../../../../features/patients/useEnsureSelectedPatient';

export default function PatientScoresScreen() {
  useScreenCaptureProtection();
  const { id } = useLocalSearchParams<{ id: string }>();
  const patientId = id ?? '';
  useEnsureSelectedPatient(patientId);

  const tree = useQuery(options.scoreTreeOptions());
  const results = useQuery(options.patientScoresOptions(patientId));

  const { refreshing, onRefresh } = useRefresh([
    queryKeys.scoreGroups.tree(),
    queryKeys.patients.scores(patientId),
  ]);

  const resultsByItem = useMemo(() => {
    const map = new Map<string, PatientScoreResult>();
    for (const r of results.data ?? []) map.set(r.itemId, r);
    return map;
  }, [results.data]);

  if (tree.isLoading || results.isLoading) return <Spinner centered />;
  if (tree.isError || results.isError) {
    return (
      <SafeAreaView className="flex-1 bg-background">
        <ErrorState
          onRetry={() => {
            tree.refetch();
            results.refetch();
          }}
        />
      </SafeAreaView>
    );
  }

  const groups = tree.data ?? [];

  async function handleShare() {
    const lines: string[] = ['Escore de saúde — resumo'];
    for (const g of groups) {
      const groupItems = collectScored(g, resultsByItem);
      if (groupItems.length === 0) continue;
      lines.push(`\n${g.name}`);
      for (const { item, level, value } of groupItems) {
        lines.push(`  • ${item.name}: ${formatValue(value)} → ${level?.level ?? 'sem nível'}`);
      }
    }
    await Share.share({ message: lines.join('\n') });
  }

  if (groups.length === 0) {
    return (
      <SafeAreaView className="flex-1 bg-background">
        <EmptyState
          title="Sem escores configurados"
          description="O sistema de escores ainda não foi populado."
        />
      </SafeAreaView>
    );
  }

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <ScrollView
        contentContainerClassName="gap-3 p-4"
        refreshControl={<RefreshControl refreshing={refreshing} onRefresh={onRefresh} />}
      >
        <Pressable
          onPress={handleShare}
          className="self-end rounded-full bg-primary px-4 py-2 active:opacity-90"
        >
          <Text className="text-sm font-semibold text-primary-foreground">Compartilhar resumo</Text>
        </Pressable>

        {groups.map((group) => (
          <ScoreGroupCard key={group.id} group={group} resultsByItem={resultsByItem} />
        ))}
      </ScrollView>
    </SafeAreaView>
  );
}

interface ScoreGroupCardProps {
  group: ScoreGroupNode;
  resultsByItem: Map<string, PatientScoreResult>;
}

function ScoreGroupCard({ group, resultsByItem }: ScoreGroupCardProps) {
  const scoredItems = collectScored(group, resultsByItem);
  if (scoredItems.length === 0) return null;

  return (
    <Card>
      <Text variant="title" className="mb-2">
        {group.name}
      </Text>
      <View className="gap-3">
        {group.subgroups.map((sub) => {
          const items = sub.items.filter((it) => resultsByItem.has(it.id));
          if (items.length === 0) return null;
          return (
            <View key={sub.id} className="gap-2">
              <Text variant="caption" className="font-semibold">
                {sub.name}
              </Text>
              {items.map((item) => {
                const result = resultsByItem.get(item.id);
                const level = pickLevel(item, result);
                return (
                  <View
                    key={item.id}
                    className="flex-row items-center justify-between rounded-md bg-muted px-3 py-2"
                  >
                    <View className="flex-1 pr-3">
                      <Text variant="body">{item.name}</Text>
                      <Text variant="caption">{formatValue(result?.value)}</Text>
                    </View>
                    {level && (
                      <View
                        className="rounded-full px-2.5 py-1"
                        style={{ backgroundColor: level.color || '#94a3b8' }}
                      >
                        <Text className="text-xs font-semibold text-white">{level.level}</Text>
                      </View>
                    )}
                  </View>
                );
              })}
            </View>
          );
        })}
      </View>
    </Card>
  );
}

interface ScoredEntry {
  item: ScoreItemNode;
  level: ScoreLevelNode | undefined;
  value: PatientScoreResult['value'] | undefined;
}

function collectScored(
  group: ScoreGroupNode,
  resultsByItem: Map<string, PatientScoreResult>,
): ScoredEntry[] {
  const out: ScoredEntry[] = [];
  for (const sub of group.subgroups) {
    for (const item of sub.items) {
      const result = resultsByItem.get(item.id);
      if (!result) continue;
      out.push({ item, level: pickLevel(item, result), value: result.value });
    }
  }
  return out;
}

function pickLevel(
  item: ScoreItemNode,
  result?: PatientScoreResult,
): ScoreLevelNode | undefined {
  if (!result) return undefined;
  if (result.levelId) return item.levels.find((l) => l.id === result.levelId);
  if (result.value === undefined || result.value === null) return undefined;
  return matchLevel(
    item.levels.map((l) => ({
      id: l.id,
      level: l.level,
      color: l.color,
      operatorType: l.operatorType as never,
      minValue: l.minValue,
      maxValue: l.maxValue,
      textValue: l.textValue,
    })),
    result.value,
  ) as ScoreLevelNode | undefined;
}

function formatValue(value: PatientScoreResult['value'] | undefined): string {
  if (value === undefined || value === null) return '—';
  if (typeof value === 'number') return value.toString();
  return String(value);
}
