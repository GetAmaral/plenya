'use client'

import { useState } from 'react'
import { useRouter } from 'next/navigation'
import { ArrowLeft, Maximize2, Minimize2 } from 'lucide-react'
import { Button } from '@/components/ui/button'
import { useRequireAuth } from '@/lib/use-auth'
import { useRequireSelectedPatient } from '@/lib/use-require-selected-patient'
import { AnamnesisFullscreenForm } from '@/components/anamnesis/AnamnesisFullscreenForm'

export default function AnamnesisFullscreenPage() {
  useRequireAuth()
  const { selectedPatient } = useRequireSelectedPatient()
  const router = useRouter()
  const [isFullscreen, setIsFullscreen] = useState(false)

  const toggleFullscreen = async () => {
    if (!document.fullscreenElement) {
      await document.documentElement.requestFullscreen()
      setIsFullscreen(true)
    } else {
      await document.exitFullscreen()
      setIsFullscreen(false)
    }
  }

  const handleSuccess = () => {
    router.push('/anamnesis')
  }

  const handleCancel = () => {
    router.push('/anamnesis')
  }

  return (
    <div className="h-screen flex flex-col bg-background overflow-hidden">
      {/* Minimal header */}
      <header className="flex-none border-b bg-card/50 backdrop-blur-sm px-6 py-3">
        <div className="flex items-center justify-between">
          <div className="flex items-center gap-4">
            <Button
              variant="ghost"
              size="sm"
              onClick={handleCancel}
              className="gap-2"
            >
              <ArrowLeft className="h-4 w-4" />
              Voltar
            </Button>
            <div className="h-6 w-px bg-border" />
            <div>
              <h1 className="text-lg font-semibold">Nova Anamnese</h1>
              {selectedPatient && (
                <p className="text-xs text-muted-foreground">
                  Paciente: {selectedPatient.name}
                </p>
              )}
            </div>
          </div>
          <Button
            variant="outline"
            size="sm"
            onClick={toggleFullscreen}
            className="gap-2"
          >
            {isFullscreen ? (
              <>
                <Minimize2 className="h-4 w-4" />
                Sair Tela Cheia
              </>
            ) : (
              <>
                <Maximize2 className="h-4 w-4" />
                Tela Cheia
              </>
            )}
          </Button>
        </div>
      </header>

      {/* Form content */}
      <div className="flex-1 overflow-hidden">
        <AnamnesisFullscreenForm
          onSuccess={handleSuccess}
          onCancel={handleCancel}
        />
      </div>
    </div>
  )
}
