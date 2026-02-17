'use client'

import { useState, useEffect } from 'react'
import { useAuth } from '@/lib/use-auth'
import { useSubscriptionPlans } from '@/lib/api/subscription-plan-api'
import { Card, CardContent } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { Skeleton } from '@/components/ui/skeleton'
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs'
import { AlertCircle, Plus } from 'lucide-react'
import { SubscriptionPlanCard } from '@/components/subscriptions/SubscriptionPlanCard'
import { SubscriptionPlanDialog } from '@/components/subscriptions/SubscriptionPlanDialog'
import type { SubscriptionPlan } from '@plenya/types'
import { useRouter } from 'next/navigation'

export default function SubscriptionPlansPage() {
  const router = useRouter()
  const { user, isLoading: isAuthLoading } = useAuth()
  const { data: plans, isLoading, error, refetch } = useSubscriptionPlans()

  const [selectedTab, setSelectedTab] = useState('all')
  const [createDialogOpen, setCreateDialogOpen] = useState(false)
  const [editDialogOpen, setEditDialogOpen] = useState(false)
  const [duplicateDialogOpen, setDuplicateDialogOpen] = useState(false)
  const [editingPlan, setEditingPlan] = useState<SubscriptionPlan | null>(null)
  const [duplicateData, setDuplicateData] = useState<any>(null)

  // Admin check
  const isAdmin = user?.roles?.includes('admin')

  // Redirect if not admin (use useEffect to avoid render-time side effects)
  useEffect(() => {
    if (!isAuthLoading && !isAdmin) {
      router.push('/dashboard')
    }
  }, [isAuthLoading, isAdmin, router])

  // Don't render anything if not admin
  if (!isAuthLoading && !isAdmin) {
    return null
  }

  const handleEdit = (plan: SubscriptionPlan) => {
    setEditingPlan(plan)
    setEditDialogOpen(true)
  }

  const handleDuplicate = (plan: SubscriptionPlan) => {
    // Create initial data for duplication
    setDuplicateData({
      name: `${plan.name} - Cópia`,
      description: plan.description,
      features: plan.features,
      price: plan.price,
      currency: plan.currency,
      billingCycle: plan.billingCycle,
      methodId: plan.methodId,
      trialPeriodDays: plan.trialPeriodDays,
      order: plan.order + 1,
      isActive: false, // Start as inactive
    })
    setDuplicateDialogOpen(true)
  }

  if (isLoading || isAuthLoading) {
    return (
      <div className="space-y-6">
        <div className="space-y-2">
          <Skeleton className="h-10 w-[400px]" />
          <Skeleton className="h-5 w-[600px]" />
        </div>
        <Skeleton className="h-10 w-full max-w-md" />
        <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
          <Skeleton className="h-[400px]" />
          <Skeleton className="h-[400px]" />
          <Skeleton className="h-[400px]" />
        </div>
      </div>
    )
  }

  if (error) {
    return (
      <div className="space-y-6">
        <div>
          <h1 className="text-3xl font-bold tracking-tight">Catálogo de Planos de Assinatura</h1>
          <p className="text-muted-foreground">
            Gerencie os planos disponíveis para subscriptions
          </p>
        </div>
        <Card className="border-destructive">
          <CardContent className="flex items-center gap-3 p-6">
            <AlertCircle className="h-5 w-5 text-destructive" />
            <div className="flex-1">
              <p className="font-medium">Erro ao carregar planos</p>
              <p className="text-sm text-muted-foreground">
                Ocorreu um erro ao carregar os planos. Tente novamente.
              </p>
            </div>
            <Button onClick={() => refetch()} variant="outline">
              Tentar Novamente
            </Button>
          </CardContent>
        </Card>
      </div>
    )
  }

  // Filter plans by status
  const allPlans = plans || []
  const activePlans = allPlans.filter((p) => p.isActive)
  const inactivePlans = allPlans.filter((p) => !p.isActive)

  const plansByTab = {
    all: allPlans,
    active: activePlans,
    inactive: inactivePlans,
  }

  const currentPlans = plansByTab[selectedTab as keyof typeof plansByTab]

  return (
    <div className="space-y-6">
      {/* Header */}
      <div className="flex items-start justify-between">
        <div>
          <h1 className="text-3xl font-bold tracking-tight">Catálogo de Planos de Assinatura</h1>
          <p className="text-muted-foreground">
            Gerencie os planos disponíveis para subscriptions de pacientes
          </p>
        </div>
        <Button onClick={() => setCreateDialogOpen(true)}>
          <Plus className="h-4 w-4 mr-1" />
          Novo Plano
        </Button>
      </div>

      {/* Tabs */}
      <Tabs value={selectedTab} onValueChange={setSelectedTab}>
        <TabsList>
          <TabsTrigger value="all">
            Todos ({allPlans.length})
          </TabsTrigger>
          <TabsTrigger value="active">
            Ativos ({activePlans.length})
          </TabsTrigger>
          <TabsTrigger value="inactive">
            Inativos ({inactivePlans.length})
          </TabsTrigger>
        </TabsList>

        <TabsContent value={selectedTab} className="space-y-4">
          {currentPlans.length > 0 ? (
            <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
              {currentPlans.map((plan) => (
                <SubscriptionPlanCard
                  key={plan.id}
                  plan={plan}
                  onEdit={handleEdit}
                  onDuplicate={handleDuplicate}
                  showActions
                />
              ))}
            </div>
          ) : (
            <Card>
              <CardContent className="flex flex-col items-center justify-center py-12">
                <p className="text-muted-foreground text-center">
                  Nenhum plano nesta categoria.
                </p>
                {selectedTab === 'all' && (
                  <p className="text-sm text-muted-foreground text-center mt-2">
                    Clique em "Novo Plano" para criar o primeiro plano de subscription.
                  </p>
                )}
              </CardContent>
            </Card>
          )}
        </TabsContent>
      </Tabs>

      {/* Create Dialog */}
      <SubscriptionPlanDialog
        open={createDialogOpen}
        onOpenChange={setCreateDialogOpen}
      />

      {/* Edit Dialog */}
      {editingPlan && (
        <SubscriptionPlanDialog
          open={editDialogOpen}
          onOpenChange={(open) => {
            setEditDialogOpen(open)
            if (!open) setEditingPlan(null)
          }}
          plan={editingPlan}
        />
      )}

      {/* Duplicate Dialog */}
      {duplicateData && (
        <SubscriptionPlanDialog
          open={duplicateDialogOpen}
          onOpenChange={(open) => {
            setDuplicateDialogOpen(open)
            if (!open) setDuplicateData(null)
          }}
          initialData={duplicateData}
        />
      )}
    </div>
  )
}
