'use client'

import { useAuthStore } from '@/lib/auth-store'
import { useRequireAuth } from '@/lib/use-auth'
import { PageHeader } from '@/components/layout/page-header'
import { ProfileForm } from '@/components/profile/ProfileForm'
import { User as UserIcon } from 'lucide-react'

export default function ProfilePage() {
  const { user } = useRequireAuth()
  const updateUser = useAuthStore((state) => state.updateUser)

  if (!user) {
    return null
  }

  const handleSuccess = (updatedUser: any) => {
    // Atualiza o usuário no Zustand store
    updateUser(updatedUser)
  }

  return (
    <div className="container mx-auto py-8 space-y-8">
      <PageHeader
        breadcrumbs={[{ label: 'Meu Perfil' }]}
        icon={UserIcon}
      />

      <div className="max-w-3xl">
        <ProfileForm user={user as any} onSuccess={handleSuccess} />
      </div>
    </div>
  )
}
