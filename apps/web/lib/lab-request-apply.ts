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

// examBlock — nome do exame + linhas de justificativa ("# ...") logo abaixo, no padrão que o
// render do PDF reconhece (pdfdoc.parseExamBlocks: "#" adere ao exame de cima, não conta página).
function examBlock(t: LabTestDefinition): string {
  const lines = [t.name]
  const just = (t.requestJustification || '').trim()
  if (just) {
    for (const ln of just.split('\n')) {
      const s = ln.trim()
      lines.push(s ? `# ${s}` : '#')
    }
  }
  return lines.join('\n')
}

// applyTemplate — monta o texto da textarea a partir dos exames de um template:
//  - filtra por sexo do paciente (sexApplicability 'all' ou == gender; ex.: PSA só ♂);
//  - exclui exames já cobertos por um pedido externo (coveredCodes);
//  - A ORDEM e a PAGINAÇÃO vêm do TEMPLATE (dado), não do código: labTests já chega ordenado por
//    display_order e cada exame traz `pageBreakBefore`. Insere uma linha em branco (= nova página
//    no render) ANTES de cada exame marcado. Nenhuma regra de agrupamento aqui.
export function applyTemplate(
  labTests: LabTestDefinition[],
  gender: string | undefined,
  coveredCodes: Set<string>,
): string {
  const visible = labTests.filter((t) => {
    const sex = t.sexApplicability || 'all'
    if (sex !== 'all' && gender && sex !== gender) return false
    if (coveredCodes.has(t.code)) return false
    return true
  })

  const out: string[] = []
  visible.forEach((t, i) => {
    // linha em branco antes deste exame (nova página) — exceto se for o 1º visível
    if (i > 0 && t.pageBreakBefore) out.push('')
    out.push(examBlock(t))
  })
  return out.join('\n')
}
