"use client"

import { useRef, FormEvent } from "react"
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog"
import { Button } from "@/components/ui/button"
import { Label } from "@/components/ui/label"
import { Textarea } from "@/components/ui/textarea"
import { Alert, AlertDescription } from "@/components/ui/alert"
import { Loader2, Info } from "lucide-react"
import { useCalculateHealthScore } from "@/lib/api/health-score-api"
import { useFormNavigation } from "@/lib/use-form-navigation"

interface CalculateScoreDialogProps {
  patientId: string
  open: boolean
  onOpenChange: (open: boolean) => void
}

export function CalculateScoreDialog({
  patientId,
  open,
  onOpenChange,
}: CalculateScoreDialogProps) {
  const formRef = useRef<HTMLFormElement>(null)
  const calculateMutation = useCalculateHealthScore()

  useFormNavigation({ formRef, disabled: !open })

  const handleSubmit = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    const formData = new FormData(e.currentTarget)
    const notes = (formData.get("notes") as string) || undefined

    calculateMutation.mutate(
      { patientId, notes },
      {
        onSuccess: () => {
          onOpenChange(false)
          if (formRef.current) {
            formRef.current.reset()
          }
        },
      }
    )
  }

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Calcular Escore de Saúde</DialogTitle>
          <DialogDescription>
            O escore será calculado com base nos exames laboratoriais e anamnese mais recentes.
          </DialogDescription>
        </DialogHeader>

        <form ref={formRef} onSubmit={handleSubmit}>
          <div className="space-y-4">
            {/* Info sobre dados disponíveis */}
            <Alert>
              <Info className="h-4 w-4" />
              <AlertDescription>
                O cálculo buscará valores em TODO o histórico de exames e anamneses, usando sempre
                o dado mais recente de cada parâmetro.
              </AlertDescription>
            </Alert>

            {/* Campo de notas */}
            <div className="space-y-2">
              <Label htmlFor="notes">Observações (opcional)</Label>
              <Textarea
                id="notes"
                name="notes"
                placeholder="Ex: Primeira avaliação, pós-tratamento, etc."
                rows={3}
              />
            </div>
          </div>

          <DialogFooter className="mt-6">
            <Button
              type="button"
              variant="outline"
              onClick={() => onOpenChange(false)}
              disabled={calculateMutation.isPending}
            >
              Cancelar
            </Button>
            <Button type="submit" disabled={calculateMutation.isPending}>
              {calculateMutation.isPending && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
              Calcular Escore
            </Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  )
}
