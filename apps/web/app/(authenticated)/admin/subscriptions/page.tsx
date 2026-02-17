'use client'

import { useState, useMemo } from 'react'
import { useAllSubscriptions, useSuspendSubscription, useActivateSubscription } from '@/lib/api/subscription-api'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Skeleton } from '@/components/ui/skeleton'
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs'
import { AlertCircle, DollarSign, Users, TrendingUp, Calendar, MoreVertical, Eye, Edit, Ban, CheckCircle } from 'lucide-react'
import { SubscriptionStatusBadge } from '@/components/subscriptions/SubscriptionStatusBadge'
import { SubscriptionDialog } from '@/components/subscriptions/SubscriptionDialog'
import { Badge } from '@/components/ui/badge'
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu'
import { Button } from '@/components/ui/button'
import Link from 'next/link'
import { useRouter } from 'next/navigation'
import { toast } from 'sonner'
import type { PatientSubscription } from '@plenya/types'

export default function AdminSubscriptionsPage() {
  const { data: allSubscriptions, isLoading, error } = useAllSubscriptions()
  const router = useRouter()
  const suspendMutation = useSuspendSubscription()
  const activateMutation = useActivateSubscription()

  const [selectedTab, setSelectedTab] = useState('all')
  const [selectedPlan, setSelectedPlan] = useState<string>('all')
  const [editDialogOpen, setEditDialogOpen] = useState(false)
  const [editingSubscription, setEditingSubscription] = useState<PatientSubscription | null>(null)

  // Extract unique plans
  const uniquePlans = useMemo(() => {
    if (!allSubscriptions) return []
    const planSet = new Set<string>()
    allSubscriptions.forEach(sub => {
      try {
        const snapshot = JSON.parse(sub.planSnapshot)
        if (snapshot.name) planSet.add(snapshot.name)
      } catch {}
    })
    return Array.from(planSet).sort()
  }, [allSubscriptions])

  // Group by status and filter by plan (must be before early returns to maintain hook order)
  const subscriptionsByStatus = useMemo(() => {
    const filterByPlan = (subs: PatientSubscription[]) => {
      if (selectedPlan === 'all') return subs
      return subs.filter(s => {
        try {
          const snapshot = JSON.parse(s.planSnapshot)
          return snapshot.name === selectedPlan
        } catch {
          return false
        }
      })
    }

    return {
      all: filterByPlan(allSubscriptions || []),
      active: filterByPlan(allSubscriptions?.filter((s) => s.status === 'active') || []),
      trial: filterByPlan(allSubscriptions?.filter((s) => s.status === 'trial') || []),
      inactive: filterByPlan(allSubscriptions?.filter((s) => s.status === 'inactive') || []),
      cancelled: filterByPlan(allSubscriptions?.filter((s) => s.status === 'cancelled') || []),
      expired: filterByPlan(allSubscriptions?.filter((s) => s.status === 'expired') || []),
      suspended: filterByPlan(allSubscriptions?.filter((s) => s.status === 'suspended') || []),
    }
  }, [allSubscriptions, selectedPlan])

  if (isLoading) {
    return (
      <div className="space-y-6">
        <div className="space-y-2">
          <Skeleton className="h-10 w-[400px]" />
          <Skeleton className="h-5 w-[600px]" />
        </div>
        <div className="grid gap-4 md:grid-cols-4">
          <Skeleton className="h-[120px]" />
          <Skeleton className="h-[120px]" />
          <Skeleton className="h-[120px]" />
          <Skeleton className="h-[120px]" />
        </div>
        <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
          <Skeleton className="h-[300px]" />
          <Skeleton className="h-[300px]" />
          <Skeleton className="h-[300px]" />
        </div>
      </div>
    )
  }

  if (error) {
    return (
      <div className="space-y-6">
        <div>
          <h1 className="text-3xl font-bold tracking-tight">Dashboard de Subscriptions</h1>
          <p className="text-muted-foreground">
            Visão geral de todos os pacotes de acompanhamento
          </p>
        </div>
        <Card className="border-destructive">
          <CardContent className="flex items-center gap-3 p-6">
            <AlertCircle className="h-5 w-5 text-destructive" />
            <div>
              <p className="font-medium">Erro ao carregar subscriptions</p>
              <p className="text-sm text-muted-foreground">
                Ocorreu um erro ao carregar os dados. Tente novamente mais tarde.
              </p>
            </div>
          </CardContent>
        </Card>
      </div>
    )
  }

  // Calculate metrics
  const totalSubscriptions = allSubscriptions?.length || 0
  const activeCount = allSubscriptions?.filter((s) => s.status === 'active').length || 0
  const trialCount = allSubscriptions?.filter((s) => s.status === 'trial').length || 0
  const cancelledCount = allSubscriptions?.filter((s) => s.status === 'cancelled').length || 0

  // Calculate revenue considering customPrice and discountPercent
  const totalRevenue = allSubscriptions
    ?.filter((s) => s.status === 'active' || s.status === 'trial')
    .reduce((sum, s) => {
      // Parse plan snapshot to get original price
      let planPrice = 0
      try {
        const snapshot = JSON.parse(s.planSnapshot)
        planPrice = snapshot.price || 0
      } catch {
        planPrice = 0
      }

      // Custom price takes precedence, otherwise apply discount to plan price
      const finalPrice = s.customPrice !== undefined
        ? s.customPrice
        : planPrice * (1 - s.discountPercent / 100)

      return sum + finalPrice
    }, 0) || 0

  // Calculate average ticket
  const avgTicket = totalSubscriptions > 0
    ? allSubscriptions.reduce((sum, s) => {
        let planPrice = 0
        try {
          const snapshot = JSON.parse(s.planSnapshot)
          planPrice = snapshot.price || 0
        } catch {
          planPrice = 0
        }

        const finalPrice = s.customPrice !== undefined
          ? s.customPrice
          : planPrice * (1 - s.discountPercent / 100)

        return sum + finalPrice
      }, 0) / totalSubscriptions
    : 0

  // Calculate churn rate (cancelled / total)
  const churnRate = totalSubscriptions > 0
    ? (cancelledCount / totalSubscriptions) * 100
    : 0

  const currentSubscriptions = subscriptionsByStatus[selectedTab as keyof typeof subscriptionsByStatus]

  const formatDate = (date: string | undefined) => {
    if (!date) return '-'
    return new Date(date).toLocaleDateString('pt-BR')
  }

  const formatCurrency = (value: number, currency = 'BRL') => {
    return new Intl.NumberFormat('pt-BR', {
      style: 'currency',
      currency,
    }).format(value)
  }

  const parsePlanSnapshot = (snapshot: string) => {
    try {
      return JSON.parse(snapshot)
    } catch {
      return null
    }
  }

  const handleSuspend = async (subscriptionId: string) => {
    try {
      await suspendMutation.mutateAsync(subscriptionId)
      toast.success('Subscription suspenso com sucesso')
    } catch (error) {
      toast.error('Erro ao suspender subscription')
    }
  }

  const handleActivate = async (subscriptionId: string) => {
    try {
      await activateMutation.mutateAsync(subscriptionId)
      toast.success('Subscription ativado com sucesso')
    } catch (error) {
      toast.error('Erro ao ativar subscription')
    }
  }

  const handleEdit = (subscription: PatientSubscription) => {
    setEditingSubscription(subscription)
    setEditDialogOpen(true)
  }

  return (
    <div className="space-y-6">
      {/* Header */}
      <div>
        <h1 className="text-3xl font-bold tracking-tight">Dashboard de Subscriptions</h1>
        <p className="text-muted-foreground">
          Visão geral e gestão de todos os pacotes de acompanhamento
        </p>
      </div>

      {/* Metrics Cards */}
      <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">Total de Subscriptions</CardTitle>
            <Users className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">{totalSubscriptions}</div>
            <p className="text-xs text-muted-foreground">
              {activeCount} ativos, {trialCount} em trial
            </p>
          </CardContent>
        </Card>

        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">Receita Mensal (Estimada)</CardTitle>
            <DollarSign className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">
              {new Intl.NumberFormat('pt-BR', {
                style: 'currency',
                currency: 'BRL',
              }).format(totalRevenue)}
            </div>
            <p className="text-xs text-muted-foreground">
              Baseado em {activeCount + trialCount} subscriptions ativos
            </p>
          </CardContent>
        </Card>

        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">Ticket Médio</CardTitle>
            <TrendingUp className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">
              {new Intl.NumberFormat('pt-BR', {
                style: 'currency',
                currency: 'BRL',
              }).format(avgTicket)}
            </div>
            <p className="text-xs text-muted-foreground">
              Valor médio por subscription
            </p>
          </CardContent>
        </Card>

        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">Taxa de Churn</CardTitle>
            <Calendar className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">
              {churnRate.toFixed(1)}%
            </div>
            <p className="text-xs text-muted-foreground">
              {cancelledCount} cancelados de {totalSubscriptions} total
            </p>
          </CardContent>
        </Card>
      </div>

      {/* Filters */}
      <div className="flex items-center gap-4">
        <div className="flex-1">
          <Tabs value={selectedTab} onValueChange={setSelectedTab}>
            <TabsList>
              <TabsTrigger value="all">
                Todos ({subscriptionsByStatus.all.length})
              </TabsTrigger>
              <TabsTrigger value="active">
                Ativos ({subscriptionsByStatus.active.length})
              </TabsTrigger>
              <TabsTrigger value="trial">
                Trial ({subscriptionsByStatus.trial.length})
              </TabsTrigger>
              <TabsTrigger value="inactive">
                Inativos ({subscriptionsByStatus.inactive.length})
              </TabsTrigger>
              <TabsTrigger value="cancelled">
                Cancelados ({subscriptionsByStatus.cancelled.length})
              </TabsTrigger>
              <TabsTrigger value="expired">
                Expirados ({subscriptionsByStatus.expired.length})
              </TabsTrigger>
              <TabsTrigger value="suspended">
                Suspensos ({subscriptionsByStatus.suspended.length})
              </TabsTrigger>
            </TabsList>
          </Tabs>
        </div>

        <Select value={selectedPlan} onValueChange={setSelectedPlan}>
          <SelectTrigger className="w-[250px]">
            <SelectValue placeholder="Filtrar por plano" />
          </SelectTrigger>
          <SelectContent>
            <SelectItem value="all">Todos os planos</SelectItem>
            {uniquePlans.map(plan => (
              <SelectItem key={plan} value={plan}>{plan}</SelectItem>
            ))}
          </SelectContent>
        </Select>
      </div>

      {/* Table */}
      <Card>
        <CardContent className="p-0">
          {currentSubscriptions.length > 0 ? (
            <Table>
              <TableHeader>
                <TableRow>
                  <TableHead>Paciente</TableHead>
                  <TableHead>Plano</TableHead>
                  <TableHead>Metodologia</TableHead>
                  <TableHead>Status</TableHead>
                  <TableHead>Valor</TableHead>
                  <TableHead>Início</TableHead>
                  <TableHead>Término</TableHead>
                  <TableHead className="text-right">Ações</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {currentSubscriptions.map((subscription) => {
                  const planData = parsePlanSnapshot(subscription.planSnapshot)
                  const snapshotPrice = planData?.price || 0
                  const finalPrice = subscription.customPrice
                    ? subscription.customPrice
                    : snapshotPrice * (1 - subscription.discountPercent / 100)
                  const hasDiscount = subscription.customPrice !== undefined || subscription.discountPercent > 0

                  return (
                    <TableRow key={subscription.id}>
                      <TableCell className="font-medium">
                        {subscription.patient?.name || '-'}
                      </TableCell>
                      <TableCell>
                        <div>
                          <div className="font-medium">{planData?.name || 'N/A'}</div>
                          {hasDiscount && (
                            <div className="text-xs text-green-600">
                              {subscription.customPrice !== undefined
                                ? 'Preço customizado'
                                : `${subscription.discountPercent}% desconto`}
                            </div>
                          )}
                        </div>
                      </TableCell>
                      <TableCell>
                        {subscription.subscriptionPlan?.method ? (
                          <Badge
                            variant="outline"
                            style={{
                              borderColor: subscription.subscriptionPlan.method.color || undefined
                            }}
                          >
                            {subscription.subscriptionPlan.method.shortName}
                          </Badge>
                        ) : (
                          <span className="text-muted-foreground text-sm">-</span>
                        )}
                      </TableCell>
                      <TableCell>
                        <SubscriptionStatusBadge status={subscription.status} />
                      </TableCell>
                      <TableCell>
                        <div>
                          <div className="font-medium">
                            {formatCurrency(finalPrice, planData?.currency)}
                          </div>
                          {hasDiscount && !subscription.customPrice && (
                            <div className="text-xs text-muted-foreground line-through">
                              {formatCurrency(snapshotPrice, planData?.currency)}
                            </div>
                          )}
                        </div>
                      </TableCell>
                      <TableCell>{formatDate(subscription.startDate)}</TableCell>
                      <TableCell>
                        {subscription.trialEndDate ? (
                          <div className="text-blue-600 font-medium">
                            Trial: {formatDate(subscription.trialEndDate)}
                          </div>
                        ) : subscription.endDate ? (
                          formatDate(subscription.endDate)
                        ) : (
                          '-'
                        )}
                      </TableCell>
                      <TableCell className="text-right">
                        <DropdownMenu>
                          <DropdownMenuTrigger asChild>
                            <Button variant="ghost" size="sm" className="h-8 w-8 p-0">
                              <span className="sr-only">Abrir menu</span>
                              <MoreVertical className="h-4 w-4" />
                            </Button>
                          </DropdownMenuTrigger>
                          <DropdownMenuContent align="end" className="w-[160px]">
                            <DropdownMenuItem
                              onClick={() => router.push(`/patients/${subscription.patientId}/subscriptions`)}
                              className="cursor-pointer"
                            >
                              <Eye className="mr-2 h-4 w-4" />
                              Ver
                            </DropdownMenuItem>
                            <DropdownMenuItem
                              onClick={() => handleEdit(subscription)}
                              className="cursor-pointer"
                            >
                              <Edit className="mr-2 h-4 w-4" />
                              Editar
                            </DropdownMenuItem>
                            <DropdownMenuSeparator />
                            {subscription.status === 'active' || subscription.status === 'trial' ? (
                              <DropdownMenuItem
                                onClick={() => handleSuspend(subscription.id)}
                                className="cursor-pointer text-destructive focus:text-destructive"
                              >
                                <Ban className="mr-2 h-4 w-4" />
                                Suspender
                              </DropdownMenuItem>
                            ) : subscription.status === 'suspended' ? (
                              <DropdownMenuItem
                                onClick={() => handleActivate(subscription.id)}
                                className="cursor-pointer text-green-600 focus:text-green-600"
                              >
                                <CheckCircle className="mr-2 h-4 w-4" />
                                Ativar
                              </DropdownMenuItem>
                            ) : null}
                          </DropdownMenuContent>
                        </DropdownMenu>
                      </TableCell>
                    </TableRow>
                  )
                })}
              </TableBody>
            </Table>
          ) : (
            <div className="flex flex-col items-center justify-center py-12">
              <p className="text-muted-foreground text-center">
                Nenhum subscription encontrado com os filtros selecionados.
              </p>
            </div>
          )}
        </CardContent>
      </Card>

      {/* Edit Dialog */}
      {editingSubscription && (
        <SubscriptionDialog
          open={editDialogOpen}
          onOpenChange={(open) => {
            setEditDialogOpen(open)
            if (!open) setEditingSubscription(null)
          }}
          subscription={editingSubscription}
          patientId={editingSubscription.patientId}
        />
      )}
    </div>
  )
}
