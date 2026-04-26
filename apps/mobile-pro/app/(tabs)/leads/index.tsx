import { useMemo, useState } from 'react';
import { Alert, FlatList, Pressable, RefreshControl, ScrollView, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { Link } from 'expo-router';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import {
  options,
  queryKeys,
  leadStatusLabels,
  leadSourceLabels,
  type LeadStatus,
  type LeadSource,
  type LeadSummary,
} from '@plenya/api-client';
import {
  Button,
  Card,
  EmptyState,
  ErrorState,
  Input,
  Sheet,
  Spinner,
  Text,
  useToast,
} from '@plenya/ui-mobile';
import { formatRelative } from '@plenya/domain';

const STATUS_FILTERS: Array<{ key: LeadStatus | 'all'; label: string }> = [
  { key: 'all', label: 'Todos' },
  { key: 'new', label: 'Novos' },
  { key: 'contacted', label: 'Contatados' },
  { key: 'qualified', label: 'Qualificados' },
  { key: 'converted', label: 'Convertidos' },
];

const SOURCES: LeadSource[] = [
  'light_claim',
  'contact_form',
  'whatsapp_inbound',
  'email_inbound',
  'newsletter',
  'manual',
];

const STATUS_COLOR: Record<LeadStatus, string> = {
  new: 'bg-primary',
  contacted: 'bg-amber-500',
  qualified: 'bg-emerald-600',
  converted: 'bg-secondary',
  lost: 'bg-muted',
  unsubscribed: 'bg-muted',
};

export default function LeadsListScreen() {
  const queryClient = useQueryClient();
  const toast = useToast();

  const [search, setSearch] = useState('');
  const [statusFilter, setStatusFilter] = useState<LeadStatus | 'all'>('all');
  const [sourceFilter, setSourceFilter] = useState<LeadSource | undefined>();
  const [assigneeFilter, setAssigneeFilter] = useState<'all' | 'mine' | string>('all');
  const [filtersOpen, setFiltersOpen] = useState(false);
  const [actionLead, setActionLead] = useState<LeadSummary | null>(null);

  const me = useQuery(options.meOptions());
  const staff = useQuery({ ...options.staffListOptions(), enabled: filtersOpen });

  const params = useMemo(
    () => ({
      search: search.trim() || undefined,
      status: statusFilter === 'all' ? undefined : statusFilter,
      source: sourceFilter,
      assignedToUserId:
        assigneeFilter === 'all'
          ? undefined
          : assigneeFilter === 'mine'
            ? me.data?.id
            : assigneeFilter,
      pageSize: 50,
    }),
    [search, statusFilter, sourceFilter, assigneeFilter, me.data?.id],
  );

  const list = useQuery(options.leadsListOptions(params));

  const update = useMutation({
    mutationFn: (vars: {
      id: string;
      patch: { status?: LeadStatus; assignedToUserId?: string | null };
    }) => options.leadMutations.update(vars.id, vars.patch),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.leads.all() });
      setActionLead(null);
      toast.show('Lead atualizado', 'success');
    },
    onError: (err) =>
      toast.show(err instanceof Error ? err.message : 'Falha', 'error'),
  });

  const activeFilterCount =
    (sourceFilter ? 1 : 0) + (assigneeFilter !== 'all' ? 1 : 0);

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <View className="gap-2 p-4">
        <Input
          placeholder="Buscar por nome, telefone, email..."
          value={search}
          onChangeText={setSearch}
          autoCapitalize="none"
        />

        <ScrollView
          horizontal
          showsHorizontalScrollIndicator={false}
          contentContainerClassName="gap-2"
        >
          {STATUS_FILTERS.map((f) => (
            <Pressable
              key={f.key}
              onPress={() => setStatusFilter(f.key)}
              className={`rounded-full px-3 py-1.5 ${
                statusFilter === f.key ? 'bg-primary' : 'bg-muted'
              }`}
            >
              <Text
                className={`text-xs font-semibold ${
                  statusFilter === f.key ? 'text-primary-foreground' : 'text-foreground'
                }`}
              >
                {f.label}
              </Text>
            </Pressable>
          ))}
          <Pressable
            onPress={() => setFiltersOpen(true)}
            className={`rounded-full px-3 py-1.5 ${
              activeFilterCount > 0 ? 'bg-secondary' : 'bg-muted'
            }`}
          >
            <Text className="text-xs font-semibold text-foreground">
              ⚙ Filtros{activeFilterCount > 0 ? ` (${activeFilterCount})` : ''}
            </Text>
          </Pressable>
        </ScrollView>
      </View>

      {list.isLoading ? (
        <Spinner centered />
      ) : list.isError ? (
        <ErrorState onRetry={() => list.refetch()} />
      ) : (
        <FlatList
          data={list.data?.items ?? []}
          keyExtractor={(item) => item.id}
          contentContainerClassName="gap-2 px-4 pb-8"
          refreshControl={
            <RefreshControl
              refreshing={list.isRefetching}
              onRefresh={() => list.refetch()}
            />
          }
          ListEmptyComponent={
            <EmptyState title="Nenhum lead" description="Sem leads pra esses filtros." />
          }
          renderItem={({ item }) => (
            <LeadRow item={item} onLongPress={() => setActionLead(item)} />
          )}
        />
      )}

      <Sheet open={filtersOpen} onClose={() => setFiltersOpen(false)}>
        <Text variant="title" className="mb-3">
          Filtros avançados
        </Text>

        <Text variant="caption" className="mb-2 font-semibold">
          Origem
        </Text>
        <View className="flex-row flex-wrap gap-2">
          <Pressable
            onPress={() => setSourceFilter(undefined)}
            className={`rounded-full px-3 py-1.5 ${
              !sourceFilter ? 'bg-primary' : 'bg-muted'
            }`}
          >
            <Text
              className={`text-xs font-semibold ${
                !sourceFilter ? 'text-primary-foreground' : 'text-foreground'
              }`}
            >
              Todas
            </Text>
          </Pressable>
          {SOURCES.map((s) => (
            <Pressable
              key={s}
              onPress={() => setSourceFilter(s)}
              className={`rounded-full px-3 py-1.5 ${
                sourceFilter === s ? 'bg-primary' : 'bg-muted'
              }`}
            >
              <Text
                className={`text-xs font-semibold ${
                  sourceFilter === s ? 'text-primary-foreground' : 'text-foreground'
                }`}
              >
                {leadSourceLabels[s]}
              </Text>
            </Pressable>
          ))}
        </View>

        <Text variant="caption" className="mb-2 mt-4 font-semibold">
          Responsável
        </Text>
        <View className="flex-row flex-wrap gap-2">
          <Pressable
            onPress={() => setAssigneeFilter('all')}
            className={`rounded-full px-3 py-1.5 ${
              assigneeFilter === 'all' ? 'bg-primary' : 'bg-muted'
            }`}
          >
            <Text
              className={`text-xs font-semibold ${
                assigneeFilter === 'all' ? 'text-primary-foreground' : 'text-foreground'
              }`}
            >
              Qualquer
            </Text>
          </Pressable>
          <Pressable
            onPress={() => setAssigneeFilter('mine')}
            className={`rounded-full px-3 py-1.5 ${
              assigneeFilter === 'mine' ? 'bg-primary' : 'bg-muted'
            }`}
          >
            <Text
              className={`text-xs font-semibold ${
                assigneeFilter === 'mine' ? 'text-primary-foreground' : 'text-foreground'
              }`}
            >
              Meus
            </Text>
          </Pressable>
          {(staff.data ?? []).map((u) => (
            <Pressable
              key={u.id}
              onPress={() => setAssigneeFilter(u.id)}
              className={`rounded-full px-3 py-1.5 ${
                assigneeFilter === u.id ? 'bg-primary' : 'bg-muted'
              }`}
            >
              <Text
                className={`text-xs font-semibold ${
                  assigneeFilter === u.id ? 'text-primary-foreground' : 'text-foreground'
                }`}
              >
                {u.name}
              </Text>
            </Pressable>
          ))}
        </View>

        <Button
          onPress={() => setFiltersOpen(false)}
          fullWidth
          className="mt-5"
        >
          Aplicar
        </Button>
        {activeFilterCount > 0 && (
          <Button
            variant="ghost"
            onPress={() => {
              setSourceFilter(undefined);
              setAssigneeFilter('all');
            }}
            fullWidth
            className="mt-2"
          >
            Limpar filtros
          </Button>
        )}
      </Sheet>

      <Sheet open={Boolean(actionLead)} onClose={() => setActionLead(null)}>
        {actionLead && (
          <View className="gap-2">
            <Text variant="title">{actionLead.name ?? 'Lead'}</Text>
            <Text variant="caption" className="mb-2">
              {leadSourceLabels[actionLead.source] ?? actionLead.source} ·{' '}
              {leadStatusLabels[actionLead.status] ?? actionLead.status}
            </Text>

            {actionLead.assignedToUserId !== me.data?.id && me.data?.id && (
              <Button
                onPress={() =>
                  update.mutate({
                    id: actionLead.id,
                    patch: { assignedToUserId: me.data!.id },
                  })
                }
                loading={update.isPending}
              >
                Atribuir a mim
              </Button>
            )}

            {actionLead.status === 'new' && (
              <Button
                variant="outline"
                onPress={() =>
                  update.mutate({
                    id: actionLead.id,
                    patch: { status: 'contacted' },
                  })
                }
                loading={update.isPending}
              >
                Marcar como contatado
              </Button>
            )}

            {actionLead.status !== 'qualified' && actionLead.status !== 'converted' && (
              <Button
                variant="outline"
                onPress={() =>
                  update.mutate({
                    id: actionLead.id,
                    patch: { status: 'qualified' },
                  })
                }
                loading={update.isPending}
              >
                Marcar como qualificado
              </Button>
            )}

            {actionLead.status !== 'lost' && actionLead.status !== 'converted' && (
              <Button
                variant="ghost"
                onPress={() =>
                  Alert.alert('Marcar como perdido?', actionLead.name ?? 'Lead', [
                    { text: 'Cancelar', style: 'cancel' },
                    {
                      text: 'Confirmar',
                      style: 'destructive',
                      onPress: () =>
                        update.mutate({
                          id: actionLead.id,
                          patch: { status: 'lost' },
                        }),
                    },
                  ])
                }
              >
                Marcar como perdido
              </Button>
            )}
          </View>
        )}
      </Sheet>
    </SafeAreaView>
  );
}

function LeadRow({
  item,
  onLongPress,
}: {
  item: LeadSummary;
  onLongPress: () => void;
}) {
  return (
    <Link href={`/(tabs)/leads/${item.id}`} asChild>
      <Pressable onLongPress={onLongPress} delayLongPress={400}>
        <Card>
          <View className="flex-row items-start justify-between">
            <View className="flex-1 pr-2">
              <Text variant="title">{item.name ?? item.email ?? item.phone ?? 'Sem nome'}</Text>
              <Text variant="caption">
                {leadSourceLabels[item.source] ?? item.source} ·{' '}
                {formatRelative(item.lastInboundAt ?? item.updatedAt)}
                {item.assignedTo?.name ? ` · ${item.assignedTo.name}` : ''}
              </Text>
            </View>
            <View
              className={`rounded-full px-2 py-0.5 ${
                STATUS_COLOR[item.status] ?? 'bg-muted'
              }`}
            >
              <Text className="text-[10px] font-semibold uppercase text-white">
                {leadStatusLabels[item.status] ?? item.status}
              </Text>
            </View>
          </View>
        </Card>
      </Pressable>
    </Link>
  );
}
