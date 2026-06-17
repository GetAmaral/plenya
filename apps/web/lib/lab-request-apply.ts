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
//  - LABORATÓRIO primeiro (tudo que não é imagem), em ordem alfabética, todos juntos (fluem na
//    mesma página, sem linha em branco entre eles);
//  - IMAGEM depois (category === 'imaging'), em ordem alfabética, UM POR PÁGINA — cada um isolado
//    por linha em branco (no render do PDF, linha em branco = quebra de página).
export function applyTemplate(
  labTests: LabTestDefinition[],
  gender: string | undefined,
  coveredCodes: Set<string>,
): string {
  const visible = [...labTests].filter((t) => {
    const sex = t.sexApplicability || 'all'
    if (sex !== 'all' && gender && sex !== gender) return false
    if (coveredCodes.has(t.code)) return false
    return true
  })

  const byName = (a: LabTestDefinition, b: LabTestDefinition) => a.name.localeCompare(b.name)
  const lab = visible.filter((t) => t.category !== 'imaging').sort(byName).map(examBlock)
  const imaging = visible.filter((t) => t.category === 'imaging').sort(byName).map(examBlock)

  // Laboratório vira um único bloco (juntado por \n, sem página entre exames). Cada exame de
  // imagem vira seu próprio bloco. join('\n\n') insere a linha em branco que separa as páginas.
  const parts: string[] = []
  if (lab.length) parts.push(lab.join('\n'))
  parts.push(...imaging)
  return parts.join('\n\n')
}
