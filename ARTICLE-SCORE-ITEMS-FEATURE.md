# 📊 Feature: Visualizador de Score Items em Articles

## 🎯 Visão Geral

Nova feature que permite visualizar na página de detalhes de um artigo científico quais score items estão vinculados a ele, organizados hierarquicamente por grupo e subgrupo.

## ✨ Funcionalidades

### 1. Visualização Hierárquica

O componente mostra os score items organizados em 3 níveis:

```
┌─────────────────────────────────────────────────────────────┐
│ 📊 Score Items Vinculados                    Ver todos →    │
│ Este artigo está vinculado a 5 items do sistema de escores │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│ • Exames Laboratoriais                                     │
│   │                                                         │
│   └─ Lipidograma                                   2 itens │
│      • Colesterol Total                    mg/dL     5pt   │
│      • Colesterol LDL                      mg/dL     3pt   │
│                                                             │
│   └─ Hormônios                                     3 itens │
│      • Testosterona Total                  ng/dL     4pt   │
│      • Estradiol                           pg/mL     3pt   │
│      • TSH                                 µIU/mL    2pt   │
│                                                             │
│ 🔗 Visualizar no Mindmap de Escores                        │
└─────────────────────────────────────────────────────────────┘
```

### 2. Informações Exibidas

Para cada score item vinculado, mostra:

- ✅ **Nome do item** (ex: "Colesterol Total")
- ✅ **Descrição** (se disponível, truncada)
- ✅ **Unidade de medida** (ex: "mg/dL")
- ✅ **Pontuação** (ex: "5pt")
- ✅ **Hierarquia completa**: Grupo > Subgrupo > Item

### 3. Indicadores Visuais

- **Contadores**: Mostra quantos items existem em cada subgrupo
- **Badges coloridos**:
  - Cinza (secondary) para contadores
  - Contornado (outline) para unidades
  - Primário (default) para pontos
- **Separadores visuais**: Bullets e bordas laterais para hierarquia
- **Hover states**: Destaque ao passar o mouse sobre items

### 4. Links de Navegação

- **"Ver todos os escores"** → `/scores`
- **"Visualizar no Mindmap de Escores"** → `/scores/mindmap`

## 🔧 Implementação Técnica

### Backend

**Arquivo:** `apps/api/internal/repository/article_repository.go`

```go
// FindByID busca um artigo por ID
func (r *ArticleRepository) FindByID(id uuid.UUID) (*models.Article, error) {
	var article models.Article
	if err := r.db.
		Preload("ScoreItems.ScoreSubgroup.ScoreGroup").
		Preload("ScoreItems.ScoreSubgroup").
		First(&article, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &article, nil
}
```

**O que faz:**
- Usa GORM Preload para carregar relações aninhadas
- Carrega ScoreItems → ScoreSubgroup → ScoreGroup em uma query
- Evita problema N+1 queries

### Frontend - Tipos

**Arquivo:** `apps/web/lib/api/article-api.ts`

```typescript
export interface ScoreGroup {
  id: string
  name: string
  description?: string
}

export interface ScoreSubgroup {
  id: string
  name: string
  description?: string
  scoreGroup?: ScoreGroup
}

export interface ScoreItem {
  id: string
  name: string
  description?: string
  unit?: string
  points: number
  scoreSubgroup?: ScoreSubgroup
}

export interface Article {
  // ... outros campos
  scoreItems?: ScoreItem[]
  // ...
}
```

### Frontend - Componente

**Arquivo:** `apps/web/components/articles/ArticleScoreItems.tsx`

**Principais funcionalidades:**

1. **Agrupamento inteligente**:
```typescript
const groupedItems = scoreItems.reduce((acc, item) => {
  // Agrupa por grupo > subgrupo > items
  // ...
}, {})
```

2. **Renderização hierárquica**:
```jsx
{groups.map(group => (
  <div>
    <h3>{group.name}</h3>
    {Object.values(group.subgroups).map(subgroup => (
      <div>
        <h4>{subgroup.name}</h4>
        {subgroup.items.map(item => (
          <ItemCard item={item} />
        ))}
      </div>
    ))}
  </div>
))}
```

3. **Conditional rendering**:
```jsx
export function ArticleScoreItems({ scoreItems }: Props) {
  if (!scoreItems || scoreItems.length === 0) {
    return null // Não renderiza nada se não houver items
  }
  // ...
}
```

### Integração na Página

**Arquivo:** `apps/web/app/articles/[id]/page.tsx`

```jsx
{/* Score Items Vinculados */}
{article.scoreItems && article.scoreItems.length > 0 && (
  <ArticleScoreItems scoreItems={article.scoreItems} />
)}
```

**Posicionamento:**
- Após Keywords & MeSH Terms
- Antes de Notes

## 📱 Exemplo Visual

### Artigo sem Score Items

Não renderiza nada (componente retorna `null`)

### Artigo com Score Items

```
┌────────────────────────────────────────────────────────────────┐
│ Artigo: "Lipid Management in Cardiovascular Disease"          │
├────────────────────────────────────────────────────────────────┤
│                                                                │
│ [Autores]                                                      │
│ [Resumo...]                                                    │
│                                                                │
│ ┌────────────────────────────────────────────────────────────┐ │
│ │ 🏷️  Palavras-chave e Termos                               │ │
│ │ Keywords: lipids, cholesterol, cardiovascular              │ │
│ └────────────────────────────────────────────────────────────┘ │
│                                                                │
│ ┌────────────────────────────────────────────────────────────┐ │
│ │ 📊 Score Items Vinculados        Ver todos os escores →   │ │
│ │ Este artigo está vinculado a 4 items                      │ │
│ │                                                            │ │
│ │ • Exames Laboratoriais                                    │ │
│ │   │                                                        │ │
│ │   └─ Lipidograma                              4 itens     │ │
│ │      ┌────────────────────────────────────────────────┐   │ │
│ │      │ • Colesterol Total              mg/dL    5pt   │   │ │
│ │      │   Níveis séricos de colesterol total           │   │ │
│ │      └────────────────────────────────────────────────┘   │ │
│ │      ┌────────────────────────────────────────────────┐   │ │
│ │      │ • Colesterol LDL                mg/dL    3pt   │   │ │
│ │      │   Low-density lipoprotein cholesterol          │   │ │
│ │      └────────────────────────────────────────────────┘   │ │
│ │      ┌────────────────────────────────────────────────┐   │ │
│ │      │ • Colesterol HDL                mg/dL    4pt   │   │ │
│ │      └────────────────────────────────────────────────┘   │ │
│ │      ┌────────────────────────────────────────────────┐   │ │
│ │      │ • Triglicerídeos                mg/dL    2pt   │   │ │
│ │      └────────────────────────────────────────────────┘   │ │
│ │                                                            │ │
│ │ 🔗 Visualizar no Mindmap de Escores                       │ │
│ └────────────────────────────────────────────────────────────┘ │
│                                                                │
│ [Notas...]                                                     │
└────────────────────────────────────────────────────────────────┘
```

## 🎨 Design System

### Cores e Estilos

**Hierarquia:**
```css
Grupo:
  - Bullet: bg-primary (dot)
  - Texto: font-semibold text-base

Subgrupo:
  - Borda esquerda: border-l-2 border-muted
  - Texto: font-medium text-sm text-muted-foreground
  - Badge contador: variant="secondary"

Item:
  - Bullet: bg-muted-foreground/50 (smaller dot)
  - Card hover: hover:bg-muted/50
  - Borda hover: hover:border-muted
  - Badge unidade: variant="outline" font-mono
  - Badge pontos: variant="default"
```

### Espaçamentos

```
Card padding: p-6
Espaçamento vertical entre grupos: space-y-6
Espaçamento vertical entre subgrupos: space-y-4
Espaçamento vertical entre items: space-y-1.5
Indentação grupo: base (0)
Indentação subgrupo: ml-3
Indentação items: ml-3 (total: ml-6 do grupo)
```

### Responsividade

- **Mobile**: Stack vertical, padding reduzido
- **Tablet**: Mantém hierarquia, mais espaçamento
- **Desktop**: Layout completo com hover states

## 🚀 Casos de Uso

### 1. Médico Visualizando Artigo

**Cenário:**
Dr. João está lendo um artigo sobre diabetes e quer saber se ele suporta os critérios de escore que ele usa.

**Fluxo:**
1. Acessa artigo `/articles/{id}`
2. Rola para baixo e vê "Score Items Vinculados"
3. Identifica que o artigo suporta:
   - Glicemia em Jejum
   - Hemoglobina Glicada
   - Peptídeo C
4. Clica em "Visualizar no Mindmap" para ver contexto completo

### 2. Curadoria de Conteúdo

**Cenário:**
Equipe médica está organizando biblioteca e quer garantir que artigos estejam vinculados aos itens corretos.

**Fluxo:**
1. Abre artigo
2. Verifica "Score Items Vinculados"
3. Se faltarem itens, clica "Editar" e adiciona vinculações
4. Se tiver itens errados, remove vinculações

### 3. Pesquisa de Evidências

**Cenário:**
Médico quer encontrar todos os artigos que suportam "Colesterol LDL".

**Fluxo:**
1. Vai para `/scores/mindmap`
2. Busca "Colesterol LDL"
3. Vê artigos vinculados
4. Clica em um artigo
5. Confirma que o artigo realmente trata do assunto

## 📊 Dados de Exemplo

### Artigo com Múltiplos Grupos

```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "title": "Comprehensive Health Assessment",
  "scoreItems": [
    {
      "id": "item-1",
      "name": "Colesterol Total",
      "unit": "mg/dL",
      "points": 5,
      "scoreSubgroup": {
        "id": "subgroup-1",
        "name": "Lipidograma",
        "scoreGroup": {
          "id": "group-1",
          "name": "Exames Laboratoriais"
        }
      }
    },
    {
      "id": "item-2",
      "name": "Pressão Arterial Sistólica",
      "unit": "mmHg",
      "points": 4,
      "scoreSubgroup": {
        "id": "subgroup-2",
        "name": "Sinais Vitais",
        "scoreGroup": {
          "id": "group-2",
          "name": "Exame Físico"
        }
      }
    }
  ]
}
```

## ✅ Checklist de Implementação

- [x] Backend: Adicionar Preload de ScoreItems no repository
- [x] Frontend: Adicionar tipos ScoreItem, ScoreSubgroup, ScoreGroup
- [x] Frontend: Criar componente ArticleScoreItems
- [x] Frontend: Integrar componente na página de detalhes
- [x] Design: Implementar hierarquia visual
- [x] Design: Adicionar badges e indicadores
- [x] UX: Links de navegação para scores e mindmap
- [x] Performance: Evitar N+1 queries com Preload
- [x] Responsive: Testar em mobile/tablet/desktop

## 🎯 Próximos Passos (Futuro)

### Melhorias Possíveis

1. **Filtros e Busca**
   - Buscar items dentro do card
   - Filtrar por grupo/subgrupo

2. **Ações Rápidas**
   - Adicionar/remover vinculações inline
   - Editar item sem sair da página

3. **Estatísticas**
   - Mostrar quantos artigos estão vinculados a cada item
   - Indicar se item tem artigos suficientes

4. **Visualização Alternativa**
   - Toggle entre lista e cards
   - Modo compacto vs expandido

5. **Export**
   - Exportar lista de items como CSV/PDF
   - Copiar lista formatada para clipboard

## 📝 Notas Técnicas

### Performance

**Query otimizada:**
```sql
SELECT * FROM articles WHERE id = ?
  -- Com Preload:
  + SELECT * FROM score_items WHERE article_id IN (?)
  + SELECT * FROM score_subgroups WHERE id IN (?)
  + SELECT * FROM score_groups WHERE id IN (?)
```

**Total:** 4 queries (em vez de N+1)

### Segurança

- ✅ Autenticação: Requer JWT válido
- ✅ Autorização: Usuário deve ter acesso ao artigo
- ✅ Validação: IDs validados no backend
- ✅ Sanitização: Dados escapados no frontend

### Acessibilidade

- ✅ Estrutura semântica (headings, lists)
- ✅ Contraste de cores adequado (WCAG AA)
- ✅ Navegação por teclado funcional
- ✅ Screen reader friendly

---

**Versão:** 1.0
**Data:** 29 de Janeiro de 2026
**Autor:** Sistema Plenya
**Commit:** `feat: Adicionar visualizador de score items vinculados em articles`
