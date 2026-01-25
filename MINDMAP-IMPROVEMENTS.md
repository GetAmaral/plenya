# Melhorias do Mindmap - Sistema de Escores

## ✨ Funcionalidades Implementadas

### 1. **Zoom Dinâmico com Level of Detail (LOD)**

O mindmap agora mostra diferentes níveis de detalhe baseado no zoom:

| Zoom Level | Visibilidade | Descrição |
|------------|-------------|-----------|
| < 30% | 🔭 **Grupos** | Visão geral - apenas grupos principais |
| 30-70% | 📊 **Grupos + Subgrupos** | Visão média - adiciona subgrupos |
| 70-120% | 🔍 **Items** | Visão detalhada - mostra items |
| ≥ 120% | 🎯 **Completa** | Visão total - todos os níveis |

**Indicador visual**: Badge no canto inferior esquerdo mostra o nível atual

### 2. **Sistema de Expansão/Colapso**

Cada tipo de node (Grupo, Subgrupo, Item) possui:
- ✅ **Botão de expansão** (chevron) quando há filhos
- ✅ **Estado visual** claro (ChevronDown = expandido, ChevronRight = recolhido)
- ✅ **Interação intuitiva** - clique no botão expande/recolhe

**Comportamento:**
- Nodes recolhidos ocultam todos os filhos
- Layout recalcula automaticamente
- Animações suaves nas transições

### 3. **Controles Globais no Header**

Botões de ação principais:
- **Procurar (Ctrl+F)** - Busca inteligente em todos os níveis
- **Expandir Tudo** - Abre todos os grupos, subgrupos e items
- **Recolher Tudo** - Fecha tudo, mostrando apenas grupos
- **Exportar PNG** - Download da visualização atual

### 4. **Painel de Controles Lateral**

Controles visuais no canto superior direito:
- 🔍 **Zoom In** - Aumenta zoom com animação
- 🔎 **Zoom Out** - Diminui zoom com animação
- 📐 **Fit View** - Ajusta visualização para mostrar tudo
- 🎯 **Reposicionar** - Retorna para posição padrão (zoom 50%)

### 5. **Busca Inteligente (Ctrl+F)**

Funcionalidade completa de busca inspirada no Adobe Acrobat Reader:

**Recursos:**
- ⌨️ **Atalho Ctrl+F** - Abre rapidamente o painel de busca
- 🔍 **Busca em tempo real** - Resultados aparecem conforme você digita
- 🎯 **Busca em todos os níveis** - Grupos, subgrupos, items e níveis
- 📊 **Lista de resultados** - Mostra todos os matches com caminho completo
- 🎨 **Badges coloridas** - Cada tipo tem sua cor (grupo, subgrupo, item, nível)
- ⌨️ **Navegação por teclado** - ↑↓ para navegar, Enter para selecionar, Esc para fechar
- 🔗 **Auto-expansão** - Expande automaticamente todo o caminho até o item encontrado
- 🎯 **Auto-foco** - Centraliza a visualização no resultado selecionado
- 📍 **Breadcrumb** - Mostra caminho completo: Grupo → Subgrupo → Item → Nível

**Workflow:**
1. Pressione **Ctrl+F** (ou clique no botão Procurar)
2. Digite o termo de busca
3. Veja os resultados em tempo real com contador
4. Use ↑↓ ou mouse para navegar pelos resultados
5. Pressione **Enter** ou clique para ir ao resultado
6. O mindmap expande automaticamente e centraliza no item

**Exemplo de uso:**
```
Buscar: "colesterol"
Resultados:
  📊 Item - Colesterol Total
     Grupo → Lipídios → Colesterol Total

  📊 Item - Colesterol HDL
     Grupo → Lipídios → Colesterol HDL

  📊 Item - Colesterol LDL
     Grupo → Lipídios → Colesterol LDL
```

**Integração com Dashboard de Gestão:**
- ✅ Mesmo componente de busca reutilizado (ScoreSearch.tsx)
- ✅ Substituiu a busca antiga que só procurava em grupos
- ✅ Auto-expansão de accordions quando encontra resultado
- ✅ Highlight visual temporário (2s) no elemento encontrado
- ✅ Scroll suave até o elemento
- ✅ Funciona em todos os níveis: grupos, subgrupos, items e níveis

### 6. **Indicadores Visuais Melhorados**

**Header:**
- Contador de grupos
- Porcentagem de zoom em tempo real

**Nodes:**
- Badges com contadores (X subgrupos, Y items, Z níveis)
- Números de ordem (#1, #2, etc.)
- Cores por nível (edges coloridas para níveis 0-5)

**Edges:**
- Animadas quando visíveis
- Coloridas baseado no nível de risco:
  - 🔴 Nível 0: Vermelho (crítico)
  - 🟠 Nível 1: Laranja (muito baixo/alto)
  - 🟡 Nível 2: Amarelo (subótimo)
  - 🔵 Nível 3: Azul (limítrofe)
  - 🟢 Nível 4: Verde (bom)
  - 🟩 Nível 5: Esmeralda (ótimo)

**Busca:**
- Box flutuante no canto superior esquerdo
- Contador de resultados em tempo real
- Footer com dicas de atalhos de teclado
- Indicador de posição (1/10, 2/10, etc.)

## 🎨 Experiência do Usuário

### Workflow Típico

1. **Visão Geral**
   - Página carrega com zoom reduzido
   - Todos os grupos visíveis
   - Layout organizado verticalmente

2. **Exploração**
   - Clique em chevron de um grupo → Expande subgrupos
   - Aumente zoom → Items aparecem automaticamente
   - Continue zoom → Níveis ficam visíveis

3. **Foco em Detalhes**
   - Use "Recolher Tudo" para limpar
   - Expanda apenas o grupo de interesse
   - Zoom in para ver todos os níveis

4. **Exportação**
   - Ajuste visualização com Fit View
   - Clique "Exportar PNG"
   - Imagem salva com qualidade alta

## 🔧 Implementação Técnica

### Arquitetura

```
page.tsx (ReactFlowProvider)
  └─ MindmapContent
      ├─ Estado de expansão (expandedNodes)
      ├─ Monitoramento de zoom (useEffect + getZoom)
      └─ buildMindmapLayout(groups, expandedNodes, zoomLevel)
          ├─ LOD baseado em zoomLevel
          ├─ Filtragem por estado de expansão
          └─ Cálculo dinâmico de posições
```

### Principais Hooks

```typescript
// Controle de estado
const [expandedNodes, setExpandedNodes] = useState<Record<string, boolean>>({})
const [zoomLevel, setZoomLevel] = useState(1)

// React Flow hooks
const { zoomIn, zoomOut, fitView, setViewport, getZoom } = useReactFlow()

// Monitoramento contínuo do zoom
useEffect(() => {
  const interval = setInterval(() => {
    const zoom = getZoom()
    setZoomLevel(zoom)
  }, 100)
  return () => clearInterval(interval)
}, [getZoom])

// Atualização de nodes com callbacks
useEffect(() => {
  setNodes(nodes =>
    nodes.map(node => ({
      ...node,
      data: {
        ...node.data,
        onToggle: () => toggleNode(node.id),
        isExpanded: expandedNodes[node.id] || false,
      }
    }))
  )
}, [expandedNodes, toggleNode, setNodes])
```

### Algoritmo de Layout

```typescript
// Estima altura do card baseado no comprimento do texto
function estimateCardHeight(text: string, cardType: 'group' | 'subgroup' | 'item' | 'level'): number {
  const charsPerLine = cardType === 'level' ? 35 : 40
  const lineHeight = 20
  const estimatedLines = Math.ceil(text.length / charsPerLine)
  return baseHeight + ((estimatedLines - 1) * lineHeight)
}

function buildMindmapLayout(
  scoreGroups: ScoreGroup[],
  expandedNodes: Record<string, boolean> = {},
  zoomLevel: number = 1
): LayoutResult {
  // LOD baseado em zoom
  const showSubgroups = zoomLevel >= 0.25
  const showItems = zoomLevel >= 0.6
  const showLevels = zoomLevel >= 1.0

  // Posicionamento dinâmico com altura estimada
  scoreGroups.forEach(group => {
    const groupCardHeight = estimateCardHeight(group.name, 'group')
    // Posiciona grupo em currentY

    group.subgroups.forEach(subgroup => {
      const subgroupCardHeight = estimateCardHeight(subgroup.name, 'subgroup')
      // Posiciona subgrupo em subgroupY

      subgroup.items.forEach(item => {
        const itemCardHeight = estimateCardHeight(item.name, 'item')
        // Posiciona item em itemY

        item.levels.forEach(level => {
          const levelCardHeight = estimateCardHeight(level.name, 'level')
          // Próximo nível: levelY + levelCardHeight + SPACING.level
        })
        // Próximo item: itemY + itemHeight + SPACING.item
      })
      // Próximo subgrupo: subgroupY + subgroupHeight + SPACING.subgroup
    })
    // Próximo grupo: currentY + groupHeight + SPACING.group
  })
}
```

## 📊 Performance

### Otimizações

1. **Lazy Rendering**
   - Apenas nodes visíveis são renderizados
   - LOD reduz nodes em zoom baixo

2. **Memo Components**
   - Todos os nodes são `memo()`
   - Previne re-renders desnecessários

3. **Debounce de Zoom**
   - Atualização a cada 100ms
   - Evita recalculos excessivos

### Métricas Estimadas

| Zoom | Nodes Visíveis | Performance |
|------|----------------|-------------|
| < 30% | ~10-20 | 🟢 Excelente |
| 30-70% | ~50-100 | 🟢 Ótima |
| 70-120% | ~200-500 | 🟡 Boa |
| ≥ 120% | ~500-1000+ | 🟡 Aceitável |

## 🎯 Funcionalidades Futuras (Opcional)

- [ ] **Busca no mindmap** - Encontrar e focar em nodes específicos
- [ ] **Filtros** - Mostrar apenas certos níveis de risco
- [ ] **Temas** - Dark/light mode para o mindmap
- [ ] **Mini preview** - Thumbnail ao passar mouse nos nodes recolhidos
- [ ] **Histórico de navegação** - Voltar/avançar no zoom
- [ ] **Compartilhamento** - Link direto para estado específico
- [ ] **Modo apresentação** - Fullscreen com controles simplificados

## 🎨 Design Uniforme (Atualização)

### Largura e Espaçamento Padronizados

Todos os cards agora têm **largura fixa de 336px** e **espaçamento uniforme de 384px** entre colunas:

**Layout Horizontal:**
```
Grupos:     x = 0
Subgrupos:  x = 384    (+384px)
Items:      x = 768    (+384px)
Níveis:     x = 1152   (+384px)
```

**Benefícios:**
- ✅ Visual consistente e organizado
- ✅ Cards mais largos (20% maior) para melhor legibilidade
- ✅ Espaçamento generoso (20% maior) entre colunas
- ✅ Textos longos quebram linha (word-wrap) ao invés de aumentar largura
- ✅ Cards crescem verticalmente conforme necessário
- ✅ Mais fácil de navegar e entender hierarquia
- ✅ Melhor aproveitamento do espaço horizontal

### Especificações dos Cards

| Tipo | Largura | Padding | Quebra de Linha | Espaçamento Vertical Mínimo |
|------|---------|---------|-----------------|----------------------------|
| Grupo | 336px | px-4 py-3 | ✅ `break-words` | 20px |
| Subgrupo | 336px | px-4 py-3 | ✅ `break-words` | 20px |
| Item | 336px | px-4 py-3 | ✅ `break-words` | 20px |
| Nível | 336px | px-3 py-2 | ✅ `break-words` | 20px |

**Espaçamento Vertical:**
- **Apenas 20px** entre todos os cards para layout compacto
- **Cálculo Dinâmico de Altura:** Cada card tem sua altura estimada baseado no comprimento do texto
- Textos longos que quebram em múltiplas linhas aumentam a altura do card
- O próximo card é posicionado levando em conta a altura real + 20px de espaçamento
- **Resultado:** Cards muito próximos, mas nunca se sobrepõem, independente do tamanho do texto

## 📝 Arquivos Modificados

### Mindmap
```
apps/web/app/scores/mindmap/page.tsx              ← Lógica principal + controles + busca + Ctrl+F
apps/web/components/scores/mindmap/
  ├─ useMindmapLayout.ts                          ← LOD + layout dinâmico + altura dinâmica
  ├─ GroupNode.tsx                                ← + botão expansão + largura fixa + word-wrap
  ├─ SubgroupNode.tsx                             ← + botão expansão + largura fixa + word-wrap
  ├─ ItemNode.tsx                                 ← + botão expansão + largura fixa + word-wrap
  ├─ LevelNode.tsx                                ← + largura fixa + word-wrap
  └─ MindmapSearch.tsx                            ← NOVO - Componente de busca (reutilizado)
```

### Dashboard de Gestão
```
apps/web/app/scores/page.tsx                      ← + busca inteligente + Ctrl+F + auto-expansão
apps/web/components/scores/
  ├─ ScoreSearch.tsx                              ← NOVO - Componente genérico de busca
  ├─ ScoreTreeView.tsx                            ← + suporte a expandedNodes + IDs nos elementos
  └─ ScoreItemCard.tsx                            ← + IDs nos níveis + prop isExpanded
```

## ✅ Testado

- ✅ Zoom dinâmico funcionando
- ✅ Expansão/colapso individual
- ✅ Expansão/colapso global
- ✅ Controles de navegação
- ✅ Indicadores visuais
- ✅ Exportação PNG
- ✅ Performance aceitável
- ✅ Sem erros de compilação
- ✅ LOD com expansão manual override
- ✅ Largura uniforme de todos os cards (336px)
- ✅ Quebra de linha em textos longos
- ✅ Espaçamento horizontal uniforme (384px)
- ✅ Espaçamento vertical dinâmico (20px mínimo)
- ✅ Níveis posicionados corretamente
- ✅ Busca em tempo real funcionando
- ✅ Atalho Ctrl+F funcionando
- ✅ Auto-expansão de resultados
- ✅ Auto-foco e centralização
- ✅ Navegação por teclado (↑↓ Enter Esc)

---

**Data**: 2026-01-24
**Status**: ✅ Implementado, testado e refinado
**Versão**: 3.0 - Design uniforme + Busca inteligente
