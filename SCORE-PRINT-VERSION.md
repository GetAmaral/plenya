# Versão de Impressão do Escore - Implementação Completa

## Visão Geral

Implementada página dedicada para impressão e exportação em PDF do Sistema de Escore Plenya, com todos os grupos, subgrupos, itens e níveis expandidos e otimizados para impressão.

## Arquivos Criados/Modificados

### 1. `/apps/web/app/scores/print/page.tsx`
- **Página principal de impressão**
- Mostra todos os dados expandidos
- Sem elementos interativos (botões de edição, accordions)
- Botão "Imprimir / Salvar PDF" que chama `window.print()`
- Layout otimizado para papel A4

### 2. `/apps/web/app/globals.css`
- Adicionados estilos `@media print` para otimizar impressão
- Configurações de página A4 com margens adequadas
- Controle de quebras de página
- Garantia de impressão de cores de fundo
- Ajustes de tamanhos de fonte para impressão

### 3. `/apps/web/app/scores/page.tsx`
- Adicionado botão "Versão Impressão" no header do dashboard
- Navegação para `/scores/print`

## Funcionalidades

### Layout de Impressão

**Estrutura Hierárquica Completa:**
```
┌─ Grupo (numerado: 1, 2, 3...)
│  ├─ Subgrupo (numerado: 1.1, 1.2, 1.3...)
│  │  ├─ Item (com pontos)
│  │  │  ├─ Níveis (0-5 com cores)
│  │  │  │  └─ Range de valores
│  │  │  └─ Itens Filhos (hierárquicos)
│  │  │     └─ Níveis
```

### Cores dos Níveis (preservadas na impressão)

| Nível | Cor | Representação |
|-------|-----|---------------|
| 0 | Vermelho | Crítico/Muito Alto Risco |
| 1 | Laranja | Alto Risco |
| 2 | Amarelo | Risco Moderado |
| 3 | Azul | Risco Baixo-Moderado |
| 4 | Verde | Baixo Risco |
| 5 | Verde Esmeralda | Ótimo/Ideal |
| 6+ | Cinza | Outros |

### Elementos Visíveis

**Na Tela:**
- Cabeçalho com botão "Imprimir / Salvar PDF"
- Título e descrição
- Todos os grupos, subgrupos e itens expandidos
- Cores e bordas completas

**Na Impressão:**
- Cabeçalho formal com:
  - Título "Escore Plenya"
  - Descrição do sistema
  - Data e hora de geração
- Todo conteúdo expandido
- Rodapé com:
  - Nome do sistema
  - Aviso LGPD
- Sem botões ou elementos de navegação

### Informações Exibidas

**Por Item:**
- Nome do item
- Unidade de medida (se houver)
- Conversão de unidade (se houver)
- Pontuação

**Por Nível:**
- Número do nível (0-5)
- Nome/descrição do nível
- Range de valores (formatado: "<10", "10-19", ">100", etc.)
- Definição (se houver)

**Itens Hierárquicos:**
- Items filhos são exibidos indentados
- Mesma estrutura de níveis
- Visualmente diferenciados com borda lateral

## Otimizações de Impressão

### Configurações de Página
```css
@page {
  size: A4;
  margin: 1.5cm;
}
```

### Controle de Quebras
- `break-inside-avoid` em grupos, subgrupos e itens
- Evita quebra de títulos no meio
- Mantém itens e seus níveis juntos

### Cores
- `-webkit-print-color-adjust: exact` garante impressão de cores
- Cores otimizadas para impressão (tons mais claros)
- Bordas definidas para clareza

### Tipografia
- Tamanhos de fonte ajustados para legibilidade em papel
- Texto: 9pt-20pt conforme hierarquia
- Font rendering otimizado

## Como Usar

### Acessar
1. Dashboard → "Versão Impressão" (botão no header)
2. Ou navegar direto para `/scores/print`

### Imprimir
1. Clicar no botão "Imprimir / Salvar PDF"
2. Na janela de impressão:
   - **Destino**: "Salvar como PDF" (para PDF) ou impressora física
   - **Layout**: Retrato
   - **Margens**: Padrão
   - **Opções**:
     - ✅ Gráficos de fundo (para preservar cores)
     - ✅ Cabeçalhos e rodapés (opcional)
3. Salvar ou Imprimir

### Atalhos de Teclado
- `Ctrl/Cmd + P` - Abrir janela de impressão
- `Esc` - Cancelar

## Estatísticas de Importação Atual

**Dados Importados:**
- **Grupos**: 13
- **Subgrupos**: 30
- **Itens**: 772 (603 com pontos)
- **Níveis**: 3.028
- **Ranges "between"**: Centenas (corrigidos)

## Exemplo de Saída

### Grupo - Exames Laboratoriais
```
┌─────────────────────────────────────────────────────┐
│ 7. Exames Laboratoriais                             │
└─────────────────────────────────────────────────────┘

  ┌───────────────────────────────────────────────────┐
  │ 7.1 Hematologia                                   │
  └───────────────────────────────────────────────────┘

    ┌─────────────────────────────────────────────────┐
    │ 25-hidroxivitamina D              20 pts        │
    │ Unidade: ng/mL | ng/mL × 2,5 = nmol/L           │
    ├─────────────────────────────────────────────────┤
    │ Níveis:                                         │
    │ [0] <10           (vermelho)  < 10              │
    │ [1] 10–19         (laranja)   10 - 19           │
    │ [2] 20–39         (amarelo)   20 - 39           │
    │ [3] >100          (azul)      > 100             │
    │ [4] 40–59         (verde)     40 - 59           │
    │ [5] 60–100        (esmeralda) 60 - 100          │
    └─────────────────────────────────────────────────┘
```

## Próximos Passos (Opcional)

### Melhorias Futuras
1. **Exportação automática para PDF server-side** (usando Puppeteer/Playwright)
2. **Customização de layout** (escolher quais grupos incluir)
3. **Logotipo da clínica** no cabeçalho
4. **Assinatura digital** do médico
5. **Geração de relatórios** com dados de pacientes preenchidos
6. **Múltiplos formatos** (PDF, Excel, Word)

## Observações Técnicas

- **Performance**: Renderiza ~3000 níveis sem problemas
- **Compatibilidade**: Chrome, Firefox, Safari, Edge
- **Responsivo**: Otimizado para A4, mas funciona em outros tamanhos
- **Acessibilidade**: Estrutura semântica com headings corretos
- **SEO**: Título dinâmico para identificação

## Troubleshooting

### Cores não aparecem na impressão
- Verificar "Gráficos de fundo" está habilitado
- Usar Chrome/Edge (melhor suporte)

### Quebras de página indesejadas
- Já configurado `break-inside-avoid`
- Se persistir, pode ser limitação do navegador

### Fonte muito pequena/grande
- Ajustar zoom antes de imprimir
- Tamanhos são otimizados para A4 em 100%

---

**Implementado em**: 2026-01-25
**Status**: ✅ Completo e funcional
**Testado**: Importação de 3.028 níveis com sucesso
