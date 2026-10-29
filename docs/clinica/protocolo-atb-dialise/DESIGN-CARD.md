# Design do card em forma de algoritmo

Conteúdo: `BRIEFING-CARD.md` (fonte única). Este arquivo define a forma. Estudo de 2026-10-29.

## Fontes de boas práticas

- **SMDM 1992**, Committee on Standardization of Clinical Algorithms. Proposal for clinical algorithm
  standards. Med Decis Making 1992;12:149-54. Formas: estado clínico = retângulo arredondado; decisão =
  hexágono; ação = retângulo; conexão = oval pequeno. (BSI 4058 usa losango para decisão.)
- **Moi D, Maulder J, Zeloof D, Iyer D.** Principles of visual design in cognitive aid development
  (ANZCA, Australasian Anaesthesia): uma única fonte sem serifa; Title Case nos títulos e **nunca caixa
  alta** (piora a legibilidade); alinhamento à esquerda (leitura em F); paleta de ~6 cores; contraste
  >= 4,5:1; espaço em branco; hierarquia por posição e agrupamento; ícones para criar seções; testar
  com quem vai usar.
- **RapidRead (UCSD, PervasiveHealth 2014):** foco no passo atual; linguagem objeto-ação estereotipada
  (verbo + objeto, curto); "information patches", blocos graficamente delimitados para achar rápido.
- **ISMP, lista de abreviações e designações de dose propensas a erro:** sem zero à direita (5 mg, não
  5,0 mg); zero à esquerda (0,5 ml); unidade sempre escrita; não usar "cc" nem "µg".
- Cor nunca sozinha: sempre com forma ou ícone (acessibilidade a daltonismo).

## Regras do card

1. **Algoritmo, não tabela de texto.** Cada infecção é um fluxo de cima para baixo:
   gatilho → ação → decisão → ação → fim (duração).
2. **Gramática visual fixa (4 formas, sempre a mesma cor):**
   - Estado/gatilho: retângulo arredondado, contorno petróleo.
   - Decisão: hexágono (ou losango), fundo âmbar claro, pergunta curta com "?".
   - Ação/antibiótico: retângulo sólido petróleo, texto branco, verbo + droga.
   - Fim/duração: pílula com relógio, fundo verde-água, só o número ("14 dias").
   - Alerta (retirar cateter, cirurgia): borda coral + ícone de alerta.
3. **Linguagem objeto-ação:** "Colher 2 hemoculturas", "Iniciar vanco + cefta", "Trocar por fio-guia".
   Máximo ~6 palavras por caixa. Nada de frase com vírgula.
4. **Doses fora do fluxo.** O fluxo traz o nome da droga em "chip" colorido; a dose fica na ficha de
   drogas (mesma cor do chip), no molde Dose → Diluição → Via → Tempo → Momento.
5. **Tipografia:** uma família sem serifa (Atkinson Hyperlegible), Title Case, sem caixa alta; corpo
   >= 9 pt; números em peso bold.
6. **Paleta (6):** petróleo (ação), âmbar (decisão), verde-água (duração), coral (alerta), cinza-azulado
   (texto secundário), branco. Contraste >= 4,5:1 conferido.
7. **Setas:** uma direção principal (de cima para baixo); ramos "sim/não" rotulados na seta.
8. **Espaço em branco > conteúdo.** Se não couber, divide em mais páginas, não encolhe fonte.
9. **Desenho em SVG/HTML posicionado à mão** (não layout automático): alinhamento em grade, mesma largura
   de caixa por coluna.
10. **Testar:** mostrar a quem prescreve e a quem administra (enfermagem) antes de fechar.

## Estrutura proposta (A4 vertical desde 2026-10-29; antes paisagem)

- **Folha 1, Bacteremia por cateter:** gatilho (febre/calafrio na sessão) → colher HC D0 → iniciar
  vanco + cefta + lock → decisão "Cultura?" com ramos MSSA / MRSA-SCN / gram-negativo / *Candida* →
  decisão "72 h: febre ou HC positiva?" → manter com lock ou trocar por fio-guia → pílulas de duração.
  Faixa lateral: hemoculturas (linha do tempo D0 / 72 h / fim) e regras de troca e novo acesso.
- **Folha 2, Outras infecções:** cinco mini-algoritmos lado a lado, mesmo molde (óstio/túnel,
  FAV/prótese, urina, pneumonia, celulite): gatilho → 1 decisão → droga → duração.
- **Folha 3, Ficha de drogas + lock:** tabelas no molde fixo, chips nas mesmas cores do fluxo.
