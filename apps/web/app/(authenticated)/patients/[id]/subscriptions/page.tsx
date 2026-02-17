'use client'

import { useState } from 'react'
import { useParams } from 'next/navigation'
import { usePatientSubscriptions } from '@/lib/api/subscription-api'
import { usePatient } from '@/lib/api/patient-api'
import { Card, CardContent } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { Skeleton } from '@/components/ui/skeleton'
import { AlertCircle, Plus } from 'lucide-react'
import { SubscriptionCard } from '@/components/subscriptions/SubscriptionCard'
import { SubscriptionDialog } from '@/components/subscriptions/SubscriptionDialog'
import { useAuth } from '@/lib/use-auth'
import { useRequireSelectedPatient } from '@/lib/use-require-selected-patient'
import type { PatientSubscription } from '@plenya/types'

export default function PatientSubscriptionsPage() {
  useRequireSelectedPatient()

  const params = useParams()
  const patientId = params.id as string

  const { user } = useAuth()
  const { data: patient, isLoading: isLoadingPatient } = usePatient(patientId)
  const { data: subscriptions, isLoading, error } = usePatientSubscriptions(patientId)

  const [createDialogOpen, setCreateDialogOpen] = useState(false)
  const [editDialogOpen, setEditDialogOpen] = useState(false)
  const [editingSubscription, setEditingSubscription] = useState<PatientSubscription | null>(null)

  const isMedicalStaff = user?.roles?.some((role) =>
    ['admin', 'doctor', 'nurse', 'psychologist', 'nutritionist'].includes(role)
  )

  const handleEdit = (subscription: PatientSubscription) => {
    setEditingSubscription(subscription)
    setEditDialogOpen(true)
  }

  if (isLoadingPatient) {
    return (
      <div className="space-y-6">
        <Skeleton className="h-10 w-[300px]" />
        <div className="grid gap-4 md:grid-cols-2">
          <Skeleton className="h-[300px]" />
          <Skeleton className="h-[300px]" />
        </div>
      </div>
    )
  }

  if (isLoading) {
    return (
      <div className="space-y-6">
        <div className="space-y-2">
          <h1 className="text-3xl font-bold tracking-tight">
            Subscriptions - {patient?.name}
          </h1>
          <p className="text-muted-foreground">
            Pacotes de acompanhamento assinados pelo paciente
          </p>
        </div>
        <div className="grid gap-4 md:grid-cols-2">
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
          <h1 className="text-3xl font-bold tracking-tight">
            Subscriptions - {patient?.name}
          </h1>
          <p className="text-muted-foreground">
            Pacotes de acompanhamento assinados pelo paciente
          </p>
        </div>
        <Card className="border-destructive">
          <CardContent className="flex items-center gap-3 p-6">
            <AlertCircle className="h-5 w-5 text-destructive" />
            <div>
              <p className="font-medium">Erro ao carregar subscriptions</p>
              <p className="text-sm text-muted-foreground">
                Ocorreu um erro ao carregar os subscriptions. Tente novamente mais tarde.
              </p>
            </div>
          </CardContent>
        </Card>
      </div>
    )
  }

  const activeSubscriptions = subscriptions?.filter((s) => s.status === 'active' || s.status === 'trial') || []
  const inactiveSubscriptions = subscriptions?.filter((s) => s.status !== 'active' && s.status !== 'trial') || []

  return (
    <div className="space-y-6">
      {/* Header */}
      <div className="flex items-start justify-between">
        <div>
          <h1 className="text-3xl font-bold tracking-tight">
            Subscriptions - {patient?.name}
          </h1>
          <p className="text-muted-foreground">
            Pacotes de acompanhamento assinados pelo paciente
          </p>
        </div>
        {isMedicalStaff && (
          <Button onClick={() => setCreateDialogOpen(true)}>
            <Plus className="h-4 w-4 mr-1" />
            Novo Subscription
          </Button>
        )}
      </div>

      {/* Active Subscriptions */}
      {activeSubscriptions.length > 0 && (
        <div className="space-y-4">
          <h2 className="text-xl font-semibold">Ativos</h2>
          <div className="grid gap-4 md:grid-cols-2">
            {activeSubscriptions.map((subscription) => (
              <SubscriptionCard
                key={subscription.id}
                subscription={subscription}
                onEdit={isMedicalStaff ? () => handleEdit(subscription) : undefined}
                showActions
              />
            ))}
          </div>
        </div>
      )}

      {/* Inactive Subscriptions */}
      {inactiveSubscriptions.length > 0 && (
        <div className="space-y-4">
          <h2 className="text-xl font-semibold">Histórico</h2>
          <div className="grid gap-4 md:grid-cols-2">
            {inactiveSubscriptions.map((subscription) => (
              <SubscriptionCard
                key={subscription.id}
                subscription={subscription}
                onEdit={isMedicalStaff ? () => handleEdit(subscription) : undefined}
                showActions
              />
            ))}
          </div>
        </div>
      )}

      {/* Empty State */}
      {subscriptions && subscriptions.length === 0 && (
        <Card>
          <CardContent className="flex flex-col items-center justify-center py-12">
            <p className="text-muted-foreground text-center">
              Nenhum subscription cadastrado ainda.
            </p>
            {isMedicalStaff && (
              <p className="text-sm text-muted-foreground text-center mt-2">
                Clique em "Novo Subscription" para criar o primeiro pacote de acompanhamento.
              </p>
            )}
          </CardContent>
        </Card>
      )}

      {/* Create Dialog */}
      <SubscriptionDialog
        open={createDialogOpen}
        onOpenChange={setCreateDialogOpen}
        patientId={patientId}
      />

      {/* Edit Dialog */}
      {editingSubscription && (
        <SubscriptionDialog
          open={editDialogOpen}
          onOpenChange={(open) => {
            setEditDialogOpen(open)
            if (!open) setEditingSubscription(null)
          }}
          subscription={editingSubscription}
          patientId={patientId}
        />
      )}
    </div>
  )
}
