# Frontend - TanStack Query Patterns

## Query Keys Hierárquicos

```typescript
export const scoreKeys = {
  all: ['scores'] as const,
  groups: () => [...scoreKeys.all, 'groups'] as const,
  group: (id: string) => [...scoreKeys.groups(), id] as const,
  groupTree: (id: string) => [...scoreKeys.group(id), 'tree'] as const,
  allGroupTrees: () => [...scoreKeys.groups(), 'trees'] as const,
  items: () => [...scoreKeys.all, 'items'] as const,
  item: (id: string) => [...scoreKeys.items(), id] as const,
  itemsBySubgroup: (subgroupId: string) => [...scoreKeys.items(), 'subgroup', subgroupId] as const,
}
```

## useQuery Pattern

```typescript
export function useScoreGroups() {
  return useQuery({
    queryKey: scoreKeys.groups(),
    queryFn: () => apiClient.get<ScoreGroup[]>('/api/v1/score-groups'),
    staleTime: 5 * 60 * 1000, // 5 minutes
  })
}

export function useScoreGroup(id: string) {
  return useQuery({
    queryKey: scoreKeys.group(id),
    queryFn: () => apiClient.get<ScoreGroup>(`/api/v1/score-groups/${id}`),
    enabled: !!id, // Só executa se id presente
  })
}
```

## useMutation Pattern

```typescript
export function useCreateScoreItem() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (data: CreateScoreItemDTO) =>
      apiClient.post<ScoreItem>('/api/v1/score-items', data),
    onSuccess: (newItem) => {
      // Invalidar queries relacionadas
      queryClient.invalidateQueries({ queryKey: scoreKeys.items() })
      queryClient.invalidateQueries({ queryKey: scoreKeys.itemsBySubgroup(newItem.subgroupId) })
      queryClient.invalidateQueries({ queryKey: scoreKeys.allGroupTrees() })
    },
  })
}

export function useUpdateScoreItem() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: ({ id, data }: { id: string; data: UpdateScoreItemDTO }) =>
      apiClient.put<ScoreItem>(`/api/v1/score-items/${id}`, data),
    onSuccess: (updatedItem, { id }) => {
      queryClient.invalidateQueries({ queryKey: scoreKeys.item(id) })
      queryClient.invalidateQueries({ queryKey: scoreKeys.items() })
      queryClient.invalidateQueries({ queryKey: scoreKeys.allGroupTrees() })
    },
  })
}
```

## Optimistic Updates (Delete Example)

```typescript
export function useDeleteScoreLevel() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (id: string) => apiClient.delete(`/api/v1/score-levels/${id}`),
    onMutate: async (id) => {
      // Cancel outgoing refetches
      await queryClient.cancelQueries({ queryKey: scoreKeys.allGroupTrees() })

      // Snapshot previous value
      const previousData = queryClient.getQueryData<ScoreGroup[]>(scoreKeys.allGroupTrees())

      // Optimistically update (remove level da UI imediatamente)
      queryClient.setQueryData<ScoreGroup[]>(
        scoreKeys.allGroupTrees(),
        (oldData) => {
          if (!oldData) return oldData

          return oldData.map(group => ({
            ...group,
            subgroups: group.subgroups?.map(subgroup => ({
              ...subgroup,
              items: subgroup.items?.map(item => ({
                ...item,
                levels: item.levels?.filter(level => level.id !== id)
              }))
            }))
          }))
        }
      )

      return { previousData }
    },
    onError: (_err, _id, context) => {
      // Rollback on error
      if (context?.previousData) {
        queryClient.setQueryData(scoreKeys.allGroupTrees(), context.previousData)
      }
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: scoreKeys.levels() })
    },
  })
}
```

## Uso no Componente

```tsx
function ScoreItemList() {
  const { data: items, isLoading, error } = useScoreItems()
  const createMutation = useCreateScoreItem()

  if (isLoading) return <LoadingSpinner />
  if (error) return <ErrorMessage error={error} />

  const handleCreate = (data: CreateScoreItemDTO) => {
    createMutation.mutate(data, {
      onSuccess: () => {
        toast.success('Item criado')
      },
      onError: (err) => {
        toast.error(err.message)
      },
    })
  }

  return (
    <>
      {items?.map(item => <ScoreItemCard key={item.id} item={item} />)}
      <CreateButton onClick={handleCreate} isPending={createMutation.isPending} />
    </>
  )
}
```

## Regras

- ✅ Query keys hierárquicos
- ✅ Invalidar queries relacionadas após mutations
- ✅ `enabled: !!id` para queries condicionais
- ✅ `staleTime` para caching (5 min para listas)
- ✅ Optimistic updates para UX (delete, toggle)
- ❌ NÃO invalidar excessivamente
- ❌ NÃO usar `refetchOnWindowFocus` em tudo

Ver: `apps/web/lib/api/*.ts`
