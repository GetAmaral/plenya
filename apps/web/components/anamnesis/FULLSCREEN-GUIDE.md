# Guia Rápido - Modo Fullscreen Anamnese

## 🎯 Para que serve?

Formulário otimizado para profissionais que fazem anamnese **durante a consulta** em **tablets**. Design limpo, touch-friendly e sem distrações.

---

## 📱 Como Acessar

### Opção 1: Da Lista de Anamneses
```
/anamnesis → Botão "Modo Tablet" (ao lado de "Nova Anamnese")
```

### Opção 2: URL Direta
```
/anamnesis/fullscreen
```

---

## 🖥️ Interface

### Header Minimalista
```
┌────────────────────────────────────────────────────────┐
│ ← Voltar  │  Nova Anamnese                             │
│           │  Paciente: João Silva              [🖵] Tela│
└────────────────────────────────────────────────────────┘
```

### Quick Info Bar (sempre visível)
```
┌─────────────────────────────────────────────────────────┐
│ 📅 08/02/2026 às 14:30   📄 Template: Primeira Consulta │
│                          👁️ Visibilidade: Todos          │
└─────────────────────────────────────────────────────────┘
```

### Seções Expansíveis (Accordion)

#### 1️⃣ Data e Template
```
▼ 📅 Data e Template
  ┌─────────────────────────────┐
  │ Data da Consulta *          │
  │ [08/02/2026 14:30]          │
  ├─────────────────────────────┤
  │ Template (opcional)         │
  │ [Selecionar template...  📋]│
  └─────────────────────────────┘
```

#### 2️⃣ Resumo
```
▼ 📄 Resumo
  ┌─────────────────────────────┐
  │ [Texto formatado ou plain]  │
  │                             │
  │ Suporta rich text editor    │
  └─────────────────────────────┘
```

#### 3️⃣ Conteúdo Detalhado
```
▶ 📄 Conteúdo Detalhado
  (clique para expandir)
```

#### 4️⃣ Detalhes Adicionais
```
▶ 📝 Detalhes Adicionais
  - Visibilidade
  - Observações
```

#### 5️⃣ Items do Template
```
▼ 📋 Items do Template  [3 preenchidos]

  ┌─ Grupo: Sintomas Físicos ──────────┐
  │ Subgrupo: Dor                      │
  │                                    │
  │ ┌─ Item: Intensidade da Dor ✓ ──┐ │
  │ │                                │ │
  │ │ Selecione o nível:             │ │
  │ │ ┌────┐ ┌────┐ ┌────┐           │ │
  │ │ │ N0 │ │ N1 │ │ N2 │ ...      │ │
  │ │ │Sem │ │Leve│ │Mod │           │ │
  │ │ └────┘ └────┘ └────┘           │ │
  │ │                                │ │
  │ │ Observações:                   │ │
  │ │ [Paciente relata dor...]       │ │
  │ │                                │ │
  │ │ ✓ Nível: 2  ✓ Com observações │ │
  │ └────────────────────────────────┘ │
  └────────────────────────────────────┘
```

### Ações (sempre no final)
```
┌─────────────────────────────────────────────────┐
│ [✨ Formatação Ativada]    [Cancelar] [💾 Salvar│
└─────────────────────────────────────────────────┘
```

---

## 🎨 Principais Diferenças

| Elemento | Normal | Fullscreen |
|----------|--------|------------|
| **Layout** | 3 colunas | 1 coluna wide |
| **Templates** | Sidebar fixa | Modal popup |
| **Seções** | Todas abertas | Accordion |
| **Botões Level** | Small (36px) | Large (80px) |
| **Fonte** | 14px | 16px (text-base) |
| **Espaçamento** | Compacto | Generoso |
| **Touch** | Desktop-first | Touch-first |

---

## 💡 Dicas de Uso

### Para Profissionais

1. **Use tela cheia** - Elimina distrações durante consulta
2. **Expanda apenas o necessário** - Foco em uma seção por vez
3. **Templates são opcionais** - Use apenas se relevante
4. **Enter navega campos** - Não precisa do mouse
5. **Níveis com cores** - Identificação visual rápida
   - 🔴 N0: Vermelho (grave)
   - 🟠 N1: Laranja
   - 🟡 N2: Amarelo
   - 🔵 N3: Azul
   - 🟢 N4-5: Verde (normal)

### Workflow Recomendado

```
1. Abrir formulário fullscreen
2. Confirmar data/hora
3. Selecionar template (se aplicável)
4. Preencher resumo rápido
5. Durante consulta:
   - Expandir seção necessária
   - Preencher items relevantes
   - Colapsar quando terminar
6. Adicionar observações finais
7. Salvar
```

---

## ⌨️ Atalhos de Teclado (futuros)

- `Ctrl+1` - Seção Data/Template
- `Ctrl+2` - Seção Resumo
- `Ctrl+3` - Seção Conteúdo
- `Ctrl+4` - Seção Detalhes
- `Ctrl+5` - Seção Items
- `Ctrl+S` - Salvar
- `Esc` - Voltar/Cancelar
- `F11` - Toggle fullscreen

---

## 🐛 Resolução de Problemas

### Tela não fica fullscreen
- Navegador pode bloquear fullscreen automático
- Clique manualmente no botão "Tela Cheia"
- Use F11 como alternativa

### Template não carrega items
- Aguarde alguns segundos (lazy loading)
- Recarregue a página se persistir
- Verifique se template tem items configurados

### Dados não salvam
- Verifique conexão com internet
- Confirme que paciente está selecionado
- Data da consulta é obrigatória

---

## 📱 Dispositivos Recomendados

### Tablets (Ideal)
- **iPad Pro 11"/12.9"** - Landscape mode
- **Samsung Tab S** - 10"+ tablets
- **Surface Pro** - Windows tablets

### Desktop (Alternativa)
- Funciona em qualquer resolução
- Ideal: 1280px+ width
- Use zoom do navegador se necessário

### Mobile
- ❌ **Não recomendado** para smartphones
- Tela muito pequena para workflow eficiente
- Use versão normal do formulário

---

## 🔗 Links Úteis

- [README Técnico](./README-FULLSCREEN.md) - Documentação para devs
- [Formulário Normal](../AnamnesisForm.tsx) - Versão padrão
- [Templates](../AnamnesisTemplateItemsForm.tsx) - Items originais

---

**Dúvidas?** Consulte a documentação técnica ou contate o desenvolvedor.
