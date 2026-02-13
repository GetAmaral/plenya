"use client"

import { PatientScoreSnapshot } from "@/lib/api/health-score-api"
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table"
import { Badge } from "@/components/ui/badge"
import { Button } from "@/components/ui/button"
import { format } from "date-fns"
import { ptBR } from "date-fns/locale"
import { Eye, Trash2 } from "lucide-react"
import { useDeleteHealthScore } from "@/lib/api/health-score-api"
import { useRouter } from "next/navigation"

interface ScoreHistoryTableProps {
  snapshots: PatientScoreSnapshot[]
}

export function ScoreHistoryTable({ snapshots }: ScoreHistoryTableProps) {
  const router = useRouter()
  const deleteSnapshotMutation = useDeleteHealthScore()

  const getScoreColor = (percentage: number) => {
    if (percentage >= 86) return "bg-green-500"
    if (percentage >= 71) return "bg-blue-500"
    if (percentage >= 51) return "bg-yellow-500"
    if (percentage >= 31) return "bg-orange-500"
    return "bg-red-500"
  }

  const handleDelete = (id: string) => {
    if (confirm("Tem certeza que deseja deletar este snapshot?")) {
      deleteSnapshotMutation.mutate(id)
    }
  }

  return (
    <Table>
      <TableHeader>
        <TableRow>
          <TableHead>Data</TableHead>
          <TableHead>Score Total</TableHead>
          <TableHead>Itens Avaliados</TableHead>
          <TableHead>Calculado Por</TableHead>
          <TableHead className="text-right">Ações</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {snapshots.map((snapshot) => (
          <TableRow key={snapshot.id}>
            <TableCell>
              {format(new Date(snapshot.calculatedAt), "dd/MM/yyyy 'às' HH:mm", {
                locale: ptBR,
              })}
            </TableCell>
            <TableCell>
              <Badge className={getScoreColor(snapshot.totalScorePercentage)}>
                {snapshot.totalScorePercentage.toFixed(1)}%
              </Badge>
            </TableCell>
            <TableCell>
              {snapshot.itemsEvaluatedCount} avaliados
              {snapshot.itemsNotEvaluatedCount > 0 &&
                ` • ${snapshot.itemsNotEvaluatedCount} sem dados`}
            </TableCell>
            <TableCell>{snapshot.calculatedByUser?.name || "—"}</TableCell>
            <TableCell className="text-right">
              <div className="flex items-center justify-end gap-2">
                <Button
                  variant="ghost"
                  size="sm"
                  onClick={() => router.push(`/health-scores/${snapshot.id}`)}
                  title="Ver detalhes"
                >
                  <Eye className="h-4 w-4" />
                </Button>
                <Button
                  variant="ghost"
                  size="sm"
                  onClick={() => handleDelete(snapshot.id)}
                  disabled={deleteSnapshotMutation.isPending}
                  title="Deletar snapshot"
                >
                  <Trash2 className="h-4 w-4" />
                </Button>
              </div>
            </TableCell>
          </TableRow>
        ))}
      </TableBody>
    </Table>
  )
}
