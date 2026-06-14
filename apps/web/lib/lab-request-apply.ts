import type { LabTestDefinition } from './api/lab-request-templates'

// normalizeName — espelha o normalizeTestName do backend (minúsculas, sem acento, sem pontuação).
// Usado p/ remover da textarea exames já cobertos por um pedido externo (dedup best-effort).
export function normalizeName(s: string): string {
  return s
    .toLowerCase()
    .normalize('NFD')
    .replace(/[̀-ͯ]/g, '')
    .replace(/[-,/():]/g, ' ')
    .replace(/[^a-z0-9 ]/g, '')
    .replace(/\s+/g, ' ')
    .trim()
}

// applyTemplate — monta o texto da textarea a partir dos exames de um template:
//  - filtra por sexo do paciente (sexApplicability 'all' ou == gender; ex.: PSA só ♂);
//  - exclui exames já cobertos por um pedido externo (coveredCodes);
//  - ordena por nome; 1 exame por linha;
//  - exame com justificativa clínica (requestJustification) emite linhas "# ..." logo abaixo
//    do nome (mesmo padrão que o render do PDF reconhece — pdfdoc.parseExamBlocks).
export function applyTemplate(
  labTests: LabTestDefinition[],
  gender: string | undefined,
  coveredCodes: Set<string>,
): string {
  return [...labTests]
    .filter((t) => {
      const sex = t.sexApplicability || 'all'
      if (sex !== 'all' && gender && sex !== gender) return false
      if (coveredCodes.has(t.code)) return false
      return true
    })
    .sort((a, b) => a.name.localeCompare(b.name))
    .map((t) => {
      const lines = [t.name]
      const just = (t.requestJustification || '').trim()
      if (just) {
        for (const ln of just.split('\n')) {
          const s = ln.trim()
          lines.push(s ? `# ${s}` : '#')
        }
      }
      return lines.join('\n')
    })
    .join('\n')
}
