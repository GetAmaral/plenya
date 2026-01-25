import { Node, Edge } from 'reactflow'
import { ScoreGroup } from '@/lib/api/score-api'

interface LayoutResult {
  nodes: Node[]
  edges: Edge[]
}

// Espaçamento entre nodes - UNIFORME para todos os níveis
const SPACING = {
  horizontal: 384, // Espaçamento uniforme entre colunas (grupo → subgrupo → item → nível) - 320px + 20%
  vertical: 20,    // Espaçamento vertical mínimo entre cards (compacto)
  group: 20,       // Espaçamento mínimo entre grupos
  subgroup: 20,    // Espaçamento mínimo entre subgrupos
  item: 20,        // Espaçamento mínimo entre items
  level: 20,       // Espaçamento mínimo entre níveis
}

// Largura fixa para todos os cards (definida aqui para consistência)
export const NODE_WIDTH = 336 // Todos os cards terão exatamente 336px de largura (280px + 20%)

// Função para estimar altura de um card baseado no texto
function estimateCardHeight(text: string, cardType: 'group' | 'subgroup' | 'item' | 'level'): number {
  const baseHeights = {
    group: 80,      // Altura base do GroupNode
    subgroup: 70,   // Altura base do SubgroupNode
    item: 70,       // Altura base do ItemNode
    level: 60,      // Altura base do LevelNode
  }

  const charsPerLine = cardType === 'level' ? 35 : 40 // Aproximadamente quantos caracteres cabem por linha
  const lineHeight = 20 // Altura aproximada de cada linha de texto

  const textLength = text?.length || 0
  const estimatedLines = Math.ceil(textLength / charsPerLine)

  // Se o texto cabe em 1 linha, retorna altura base
  if (estimatedLines <= 1) {
    return baseHeights[cardType]
  }

  // Adiciona altura extra para linhas adicionais
  const extraHeight = (estimatedLines - 1) * lineHeight
  return baseHeights[cardType] + extraHeight
}

export function buildMindmapLayout(
  scoreGroups: ScoreGroup[],
  expandedNodes: Record<string, boolean> = {},
  zoomLevel: number = 1
): LayoutResult {
  const nodes: Node[] = []
  const edges: Edge[] = []

  let currentY = 0

  // Level of Detail baseado no zoom (ajustado para thresholds mais baixos)
  const showSubgroups = zoomLevel >= 0.25
  const showItems = zoomLevel >= 0.6
  const showLevels = zoomLevel >= 1.0

  scoreGroups.forEach((group, groupIndex) => {
    const groupId = `group-${group.id}`
    const groupExpanded = expandedNodes[groupId]

    // Estimar altura do card do grupo
    const groupCardHeight = estimateCardHeight(group.name, 'group')

    // Sempre mostrar grupos
    nodes.push({
      id: groupId,
      type: 'groupNode',
      position: { x: 0, y: currentY },
      data: {
        ...group,
        subgroupCount: group.subgroups?.length || 0,
        isExpanded: groupExpanded,
      },
    })

    let groupHeight = groupCardHeight

    // Mostrar subgrupos se:
    // 1. Zoom permite (>= 0.25) OU
    // 2. Grupo está expandido manualmente (usuário clicou no chevron)
    // Lógica: Se o usuário expandiu manualmente, mostrar independente do zoom
    const shouldShowSubgroups = (showSubgroups || groupExpanded) && groupExpanded && group.subgroups && group.subgroups.length > 0

    if (shouldShowSubgroups) {
      let subgroupY = currentY

      group.subgroups.forEach((subgroup) => {
        const subgroupId = `subgroup-${subgroup.id}`
        const subgroupExpanded = expandedNodes[subgroupId]

        // Estimar altura do card do subgrupo
        const subgroupCardHeight = estimateCardHeight(subgroup.name, 'subgroup')

        nodes.push({
          id: subgroupId,
          type: 'subgroupNode',
          position: { x: SPACING.horizontal, y: subgroupY },
          data: {
            ...subgroup,
            groupId: group.id,
            itemCount: subgroup.items?.length || 0,
            isExpanded: subgroupExpanded,
          },
        })

        edges.push({
          id: `${groupId}-${subgroupId}`,
          source: groupId,
          target: subgroupId,
          type: 'smoothstep',
          animated: true,
        })

        let subgroupHeight = subgroupCardHeight

        // Mostrar items se:
        // 1. Zoom permite (>= 0.6) OU
        // 2. Subgrupo está expandido manualmente (usuário clicou no chevron)
        // Lógica: Se o usuário expandiu manualmente, mostrar independente do zoom
        const shouldShowItems = (showItems || subgroupExpanded) && subgroupExpanded && subgroup.items && subgroup.items.length > 0

        if (shouldShowItems) {
          let itemY = subgroupY

          subgroup.items.forEach((item) => {
            const itemId = `item-${item.id}`
            const itemExpanded = expandedNodes[itemId]

            // Estimar altura do card do item
            const itemCardHeight = estimateCardHeight(item.name, 'item')

            nodes.push({
              id: itemId,
              type: 'itemNode',
              position: { x: SPACING.horizontal * 2, y: itemY },
              data: {
                ...item,
                subgroupId: subgroup.id,
                levelCount: item.levels?.length || 0,
                isExpanded: itemExpanded,
              },
            })

            edges.push({
              id: `${subgroupId}-${itemId}`,
              source: subgroupId,
              target: itemId,
              type: 'smoothstep',
              animated: true,
            })

            let itemHeight = itemCardHeight

            // Mostrar níveis se:
            // 1. Zoom permite (>= 1.0) OU
            // 2. Item está expandido manualmente (usuário clicou no chevron)
            // Lógica: Se o usuário expandiu manualmente, mostrar independente do zoom
            const shouldShowLevels = (showLevels || itemExpanded) && itemExpanded && item.levels && item.levels.length > 0

            if (shouldShowLevels) {
              const sortedLevels = [...item.levels].sort((a, b) => a.level - b.level)
              let levelY = itemY
              let totalLevelsHeight = 0

              sortedLevels.forEach((level, levelIndex) => {
                const levelId = `level-${level.id}`

                // Estimar altura do card do nível
                const levelCardHeight = estimateCardHeight(level.name, 'level')

                nodes.push({
                  id: levelId,
                  type: 'levelNode',
                  position: { x: SPACING.horizontal * 3, y: levelY },
                  data: {
                    ...level,
                    itemId: item.id,
                  },
                })

                edges.push({
                  id: `${itemId}-${levelId}`,
                  source: itemId,
                  target: levelId,
                  type: 'smoothstep',
                  animated: true,
                  style: {
                    stroke: getLevelColor(level.level),
                    strokeWidth: 2,
                  },
                })

                // Próximo nível posicionado abaixo: altura do card atual + espaçamento mínimo
                const nextLevelY = levelCardHeight + SPACING.level
                levelY += nextLevelY
                totalLevelsHeight += nextLevelY
              })

              itemHeight = Math.max(itemHeight, totalLevelsHeight)
            }

            // Próximo item: altura atual do item + espaçamento mínimo
            itemY += itemHeight + SPACING.item
          })

          subgroupHeight = Math.max(subgroupHeight, itemY - subgroupY)
        }

        // Próximo subgrupo: altura atual do subgrupo + espaçamento mínimo
        subgroupY += subgroupHeight + SPACING.subgroup
      })

      groupHeight = Math.max(groupHeight, subgroupY - currentY)
    }

    // Próximo grupo: altura atual do grupo + espaçamento mínimo
    currentY += groupHeight + SPACING.group
  })

  return { nodes, edges }
}

function getLevelColor(level: number): string {
  const colors: Record<number, string> = {
    0: '#ef4444', // red
    1: '#f97316', // orange
    2: '#eab308', // yellow
    3: '#3b82f6', // blue
    4: '#22c55e', // green
    5: '#10b981', // emerald
  }
  return colors[level] || '#6b7280' // gray default
}
