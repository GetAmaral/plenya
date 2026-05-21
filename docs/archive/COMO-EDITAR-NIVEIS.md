# Como Adicionar e Editar Níveis de Escores

## 📋 Visão Geral

Os **Níveis** (Levels) são as diferentes faixas de estratificação de risco dentro de cada **Item**. Cada item pode ter de 0 a 7 níveis (0-6), representando diferentes graus de risco ou condição clínica.

---

## 🎯 Hierarquia do Sistema

```
Grupo
 └─ Subgrupo
     └─ Item (ex: "Glicose em jejum", "Colesterol Total")
         └─ Níveis (ex: Crítico, Subótimo, Bom, Ótimo)
```

---

## ➕ Como Adicionar um Novo Nível

### Passo 1: Localizar o Item

1. Acesse **Gestão de Escores** (`/scores`)
2. Use **Ctrl+F** para buscar rapidamente ou navegue manualmente
3. Localize o **Item** onde deseja adicionar um nível
4. Expanda o subgrupo clicando na seta

### Passo 2: Clicar no Botão "+"

No card do item, você verá três botões no canto superior direito:

```
┌──────────────────────────────────────────┐
│ Glicose em jejum         [mg/dL]         │
│ 10 pontos                                │
│                                  [+] [✏️] [🗑️]
│ ─────────────────────────────────────    │
│ N0: Crítico    N2: Subótimo    N4: Bom   │
└──────────────────────────────────────────┘
```

- **[+]** = Adicionar novo nível
- **[✏️]** = Editar o item
- **[🗑️]** = Excluir o item

### Passo 3: Preencher o Formulário

Ao clicar em **[+]**, abre um diálogo com os seguintes campos:

#### 1. **Número do Nível** (obrigatório)
Escolha de 0 a 6:
- **Nível 0** - 🔴 Crítico (pior condição)
- **Nível 1** - 🟠 Muito Baixo/Alto
- **Nível 2** - 🟡 Subótimo
- **Nível 3** - 🔵 Limítrofe
- **Nível 4** - 🟢 Bom
- **Nível 5** - 🟩 Ótimo (melhor condição)
- **Nível 6** - ⚪ Reservado

#### 2. **Nome** (obrigatório)
Descrição curta do nível.

**Exemplos:**
- "Hipoglicemia severa"
- "Dentro da faixa ótima"
- "Elevado - risco cardiovascular"

#### 3. **Definição** (opcional)
Explicação detalhada sobre o que significa esse nível.

**Exemplo:**
```
Glicose muito baixa, indica risco de choque hipoglicêmico.
Requer intervenção imediata.
```

#### 4. **Operador** (obrigatório)
Como interpretar os limites:

| Operador | Símbolo | Uso |
|----------|---------|-----|
| Igual a | = | Valor exato (raro) |
| Maior que | > | Acima de um valor |
| Maior ou igual | ≥ | A partir de um valor |
| Menor que | < | Abaixo de um valor |
| Menor ou igual | ≤ | Até um valor |
| **Entre (intervalo)** | - | Faixa de valores (mais comum) |

#### 5. **Limites** (opcional, mas recomendado)

**Para operador "Entre (intervalo)":**
- **Limite Inferior**: Ex: 70
- **Limite Superior**: Ex: 100
- Resultado: 70 - 100 mg/dL

**Para outros operadores:**
- Preencha apenas um dos campos
- Ex: Operador ">=" + Limite Inferior "200" = ≥ 200 mg/dL

### Exemplo Completo: Glicose em Jejum

```
Nível 0 - Crítico
  Nome: Hipoglicemia severa
  Definição: Glicose muito baixa, risco de choque hipoglicêmico
  Operador: <
  Limite Inferior: 55
  Resultado: < 55 mg/dL

Nível 2 - Subótimo
  Nome: Pré-diabetes
  Definição: Glicemia de jejum alterada
  Operador: Entre
  Limite Inferior: 100
  Limite Superior: 125
  Resultado: 100 - 125 mg/dL

Nível 4 - Bom
  Nome: Glicemia normal
  Definição: Dentro da faixa de normalidade
  Operador: Entre
  Limite Inferior: 70
  Limite Superior: 99
  Resultado: 70 - 99 mg/dL
```

---

## ✏️ Como Editar um Nível Existente

### Método 1: Clicar no Ícone de Lápis (NOVO!)

Cada badge de nível agora tem botões de ação:

```
┌──────────────────────────────────────────┐
│ N0: Crítico [✏️] [🗑️]                     │
│ N2: Subótimo [✏️] [🗑️]                    │
│ N4: Bom [✏️] [🗑️]                         │
└──────────────────────────────────────────┘
```

**Passos:**
1. Passe o mouse sobre a badge do nível
2. Clique no ícone **[✏️]** (editar)
3. O diálogo abre com os dados preenchidos
4. Modifique os campos necessários
5. Clique em **"Salvar"**

### Método 2: Tooltip para Visualizar Detalhes

Passe o mouse sobre a badge para ver um tooltip com:
- Número e nome do nível
- Classificação (Crítico, Ótimo, etc.)
- Definição completa
- Faixa de valores

---

## 🗑️ Como Excluir um Nível

### Clicar no Ícone de Lixeira

```
┌──────────────────────────────────────────┐
│ N0: Crítico [✏️] [🗑️] ← Clique aqui      │
└──────────────────────────────────────────┘
```

**Passos:**
1. Localize a badge do nível
2. Clique no ícone **[🗑️]** (excluir)
3. Confirme a exclusão no diálogo
4. O nível é removido imediatamente

**⚠️ Atenção:** A exclusão de um nível é permanente e não pode ser desfeita!

---

## 🎨 Sistema de Cores

As badges têm cores que representam visualmente o nível de risco:

| Nível | Cor | Significado | Exemplo |
|-------|-----|-------------|---------|
| 0 | 🔴 Vermelho | Crítico - pior condição | Glicose < 55 mg/dL |
| 1 | 🟠 Laranja | Muito Baixo/Alto | Pressão < 90/60 mmHg |
| 2 | 🟡 Amarelo | Subótimo | Glicose 100-125 mg/dL |
| 3 | 🔵 Azul | Limítrofe | Colesterol 200-239 mg/dL |
| 4 | 🟢 Verde | Bom | Glicose 70-99 mg/dL |
| 5 | 🟩 Verde-Esmeralda | Ótimo - melhor condição | HDL > 60 mg/dL |
| 6 | ⚪ Cinza | Reservado | (uso futuro) |

---

## 🔍 Visualizar no Mindmap

Após adicionar/editar níveis, você pode visualizá-los no **Mindmap**:

1. Clique em **"Visualizar Mindmap"** no topo da página
2. Expanda o item desejado
3. Os níveis aparecem como cards coloridos à direita
4. Use **Ctrl+F** para buscar rapidamente

```
Grupo ──→ Subgrupo ──→ Item ──→ Níveis
                                 ├─ N0: Crítico (vermelho)
                                 ├─ N2: Subótimo (amarelo)
                                 └─ N4: Bom (verde)
```

---

## ✅ Boas Práticas

### 1. **Nomenclatura Clara**
- ✅ "Hipertensão estágio 2"
- ❌ "Nível alto 2"

### 2. **Definições Completas**
- ✅ "Pressão arterial elevada que requer medicação e mudanças de estilo de vida"
- ❌ "PA alta"

### 3. **Limites Precisos**
- ✅ Use intervalos quando possível: 100 - 125 mg/dL
- ❌ Evite valores vagos

### 4. **Cobertura Completa**
- Certifique-se que os níveis cobrem toda a faixa de valores possíveis
- Evite "buracos" entre níveis

**Exemplo de cobertura completa (Glicose):**
```
N0: < 55 mg/dL         (Crítico baixo)
N1: 55 - 69 mg/dL      (Baixo)
N4: 70 - 99 mg/dL      (Normal)
N2: 100 - 125 mg/dL    (Pré-diabetes)
N0: ≥ 126 mg/dL        (Diabetes)
```

### 5. **Consistência com Literatura Médica**
- Use referências científicas atualizadas
- Siga guidelines internacionais quando aplicável
- Documente a fonte na definição quando relevante

---

## 🐛 Troubleshooting

### Problema: Não consigo adicionar mais níveis

**Possíveis causas:**
- Já existem 7 níveis (máximo permitido: 0-6)
- O item foi bloqueado para edição

**Solução:** Verifique quantos níveis já existem. Se necessário, edite ou exclua um nível existente.

### Problema: Os limites não aparecem

**Causa:** Campos de limite não foram preenchidos ou operador não requer limites.

**Solução:**
- Preencha "Limite Inferior" e/ou "Limite Superior"
- Para "Entre", preencha ambos os campos

### Problema: Nível não aparece no mindmap

**Causa:** Zoom muito baixo ou item não expandido.

**Solução:**
- Use zoom ≥ 100% no mindmap
- Clique no chevron do item para expandir
- Ou use "Expandir Tudo"

---

## 📊 Campos do Formulário (Referência Completa)

| Campo | Tipo | Obrigatório | Validação | Exemplo |
|-------|------|-------------|-----------|---------|
| Número do Nível | Select | ✅ Sim | 0-6 | 4 (Bom) |
| Nome | Text | ✅ Sim | 1-200 caracteres | "Dentro da faixa normal" |
| Definição | Textarea | ❌ Não | Até 1000 caracteres | "Glicemia de jejum considerada saudável..." |
| Operador | Select | ✅ Sim | Opções fixas | "entre" |
| Limite Inferior | Number | ❌ Não* | Número decimal | 70 |
| Limite Superior | Number | ❌ Não* | Número decimal | 99 |

*Recomendado preencher para ter faixas de valores definidas.

---

## 🚀 Atalhos Úteis

| Ação | Atalho |
|------|--------|
| Buscar item | **Ctrl+F** |
| Fechar diálogo | **Esc** |
| Salvar formulário | **Ctrl+Enter** (se suportado) |

---

## 📸 Exemplo Visual

```
┌─────────────────────────────────────────────────────┐
│ ITEM: Colesterol Total                  [mg/dL]     │
│ 15 pontos                                 [+] [✏️] [🗑️]│
│ ─────────────────────────────────────────────────── │
│                                                      │
│ 🔴 N0: Muito Alto (≥ 240 mg/dL)          [✏️] [🗑️]  │
│    Risco cardiovascular significativo               │
│                                                      │
│ 🔵 N3: Limítrofe (200-239 mg/dL)         [✏️] [🗑️]  │
│    Requer mudanças no estilo de vida                │
│                                                      │
│ 🟢 N4: Desejável (< 200 mg/dL)           [✏️] [🗑️]  │
│    Dentro da faixa saudável                         │
│                                                      │
└─────────────────────────────────────────────────────┘
```

---

**Última atualização**: 2026-01-24
**Versão**: 2.0 - Com edição e exclusão direta nas badges
