---
name: plenya-deck
description: Gerador de decks comerciais Plenya em PPTX nativo, on-brand. Wrapper sobre a skill `pptx` que injeta paleta oficial (gold/petrol/ocean/sage/cream), tipografia, frases-âncora dos vídeos institucionais, estrutura narrativa Plenya (origem 40 anos → fragmentação → Normal vs Ótimo → Método AGIR → Escore 800+ → equipe → fechamento) e regras editoriais invariantes (sem preços, sem marcas comerciais, sem casos clínicos identificáveis, sem "medicina preditiva"). Invocar quando o usuário pedir "deck", "apresentação", "slides" ou "pitch" com qualquer relação à Plenya — corporativo, parceiros, paciente PJ, fornecedor, investidor, equipe. Para deck científico de Dr. Getúlio (voz pessoal, não institucional Plenya), usar contexto específico no prompt.
---

# Skill: `/plenya-deck` — Decks comerciais Plenya

## Quando usar

- Deck comercial para plano corporativo PJ
- Apresentação para parceiros (academias, clínicas, operadoras)
- Pitch institucional Plenya (investidores, imprensa, eventos)
- Material de equipe/onboarding interno
- Qualquer slide deck que represente a **voz institucional Plenya**

**NÃO usar para:** voz pessoal Dr. Getúlio (palestras médicas, congressos) — esses usam tom observacional clínico, não institucional. Para casos assim, instruir o usuário a usar `pptx` skill direta + memória `getulio_credenciais_medicas`.

## 🚨 REGRA ZERO — Ler o site ANTES de escrever uma linha

**Esta regra é inviolável.** Nasceu de uma falha grave: deck do Continuum gerado sem nem citar o nome do programa, com posicionamento corporativo B2B inventado quando o programa é B2C individual.

Antes de gerar QUALQUER copy de deck, ler obrigatoriamente:

| Tema do deck | Página oficial (source-of-truth) | Translation keys |
|---|---|---|
| Continuum (programa principal) | `apps/site/app/[locale]/continuum/page.tsx` | `messages/pt.json` → `continuum.*` |
| Escore Plenya (instrumento) | `apps/site/app/[locale]/escore-plenya/page.tsx` | `messages/pt.json` → `escorePlenya.*` |
| Método AGIR | `apps/site/app/[locale]/metodo-agir/page.tsx` | `messages/pt.json` → `metodoAgir.*` |
| Equipe | `apps/site/app/[locale]/equipe/page.tsx` | `messages/pt.json` → `equipe.*` |
| Consultas (avulsas) | `apps/site/app/[locale]/consultas/page.tsx` | `messages/pt.json` → `consultas.*` |
| Home/Plenya geral | `apps/site/app/[locale]/page.tsx` | `messages/pt.json` → `home.*` |
| Checkup Longevidade | `apps/site/app/[locale]/checkup-longevidade/page.tsx` | `messages/pt.json` → `checkupLongevidade.*` |

**Como ler:** o `.tsx` mostra a estrutura/ordem das seções; o `pt.json` traz a copy oficial verbatim. Sempre dump das chaves antes de redigir:

```bash
node -e "const c=require('./apps/site/messages/pt.json').<namespace>; for (const [k,v] of Object.entries(c)) console.log(k+' :: '+v);"
```

**Princípios de uso da copy oficial:**
- Frases-âncora (hero, claims, manifestos) — usar **verbatim**. Não reescrever, não "melhorar".
- Listas (retratos "para quem é", razões "não é", fases da jornada, itens de escopo) — preservar a estrutura e o número exatos do site.
- Conceitos próprios (nomes de programas, entregas, instrumentos) — citar pelo nome oficial sempre. Exemplos canônicos: **Continuum Plenya**, **Box Plenya**, **Escore Plenya**, **Método AGIR**, **médico-gestor**, **janela silenciosa**, **Consultas Plenya**.
- Modalidades comerciais (Semestral/Anual, etc) — usar nomes do site.
- Preços — nunca aparecem em deck. Site fala "sob consulta" — manter assim.

**Se a copy oficial não cobrir alguma seção do deck**, registrar o gap explicitamente e pedir input do usuário, em vez de inventar.

## Como funciona

Esta skill **não executa sozinha** — ela compõe um briefing rico em cima do qual a skill `pptx` (oficial Anthropic) gera o `.pptx` final via `pptxgenjs`. Workflow:

1. **Lê a fonte de verdade no site** (Regra Zero acima) — sem isso, parar.
2. Lê o pedido do usuário (objetivo, público, número de slides)
3. Carrega brand assets desta skill (cores, fontes, frases-âncora)
4. Estrutura a narrativa **espelhando a página oficial**, não inventando
5. Aplica regras editoriais invariantes
6. Delega para skill `pptx` gerar arquivo final
7. Salva em `docs/decks/<slug>-YYYYMMDD.pptx` (criar diretório se não existir)

## Brand assets — fonte única

### Paleta oficial (HEX)

```javascript
const PLENYA = {
  gold:   "B38645",  // primário — CTA, números grandes, ícones acento
  petrol: "063B4F",  // dominante — fundos escuros, headers
  ocean:  "417E8E",  // secundário — gráficos, divisores
  sage:   "92B8B4",  // suporte — backgrounds claros, ilustrações
  cream:  "EAE7DA",  // base clara — fundo content slides
  white:  "FFFFFF",
  ink:    "0A1F26",  // texto sobre claro (petrol mais escuro)
};
```

**Regra de dominância:** petrol domina 60-70%, ocean+sage compõem 20-30%, gold é o acento (5-10%). Cream é base alternativa pra content slides.

**Sandwich pattern obrigatório:**
- Capa + fechamento → fundo **petrol**, texto cream/gold
- Slides de conteúdo → fundo **cream**, texto petrol/ink
- Slides de transição/destaque → fundo **ocean**, texto cream

### Tipografia

Fonte da verdade: `packages/brand/src/tokens/typography.ts`.

| Camada | Fonte primária | Fallbacks |
|---|---|---|
| Heading | **Nalieta** (proprietária) | Cormorant Garamond → Georgia → serif |
| Body | **Inter** | Geist → system-ui → sans-serif |

**Em slides PPTX:**
- Nalieta NÃO é usada como texto (proprietária, não distribuível). Aparece apenas no **wordmark logo embedado como PNG** em capa/fechamento.
- Heading dos slides → **Cormorant Garamond** (segundo nível do fallback oficial, Google Fonts gratuita, ampla disponibilidade).
- Body dos slides → **Inter** (Google Fonts gratuita).
- Se o viewer não tem Cormorant/Inter instaladas, PowerPoint faz fallback automático para Georgia/Calibri — mantém a hierarquia.

| Elemento | Fonte | Tamanho | Peso |
|---|---|---|---|
| Wordmark capa | (logo PNG embedado) | — | — |
| Slide title | Cormorant Garamond | 36pt | bold |
| Kicker uppercase | Inter | 11pt | bold, charSpacing 5 |
| Pillar letter (A/G/I/R) | Cormorant Garamond | 130pt | bold |
| Big stat (800+) | Cormorant Garamond | 280pt | bold |
| Body | Inter | 14pt | regular |
| Caption | Inter | 10-11pt | regular, muted |
| Pull-quote text | Inter | 13pt | italic |

**Nunca usar:** Arial, Helvetica, Impact, Times New Roman, Comic Sans.

### Frases-âncora (verbatim — NÃO reescrever)

Sempre que possível, citar frases oficiais dos vídeos institucionais em vez de inventar copy nova:

| Frase | Slot recomendado |
|---|---|
| "Saúde, Performance & Longevidade" | Tagline (capa) |
| "Viva bem, viva mais." | Claim (fechamento) |
| "Normal não é o mesmo que ótimo." | Diferenciador central |
| "Saúde não é sorte. É estrutura." | Hero / abertura |
| "Saúde não melhora com intenção. Melhora com direção." | Slide AGIR |
| "Saúde não é sobre reagir. É sobre antecipar." | Posicionamento |
| "O corpo não funciona em partes." | Slide fragmentação |
| "Plenitude não é um ponto de chegada. É uma linha contínua." | Manifesto |

## Template comercial padrão (12 slides)

Estrutura narrativa default para deck comercial corporativo. Ajustar contagem conforme briefing.

| # | Slide | Conteúdo |
|---|---|---|
| 1 | **Capa** | Logo Plenya + tagline "Saúde, Performance & Longevidade" + claim. Fundo petrol. |
| 2 | **Origem** | 40 anos de história — Nefroclínica → Plenya. Não é startup. |
| 3 | **O problema** | Fragmentação do cuidado. "O corpo não funciona em partes." Visual: paciente cercado de 5-6 especialistas isolados. |
| 4 | **Insight central** | "Normal não é o mesmo que ótimo." Slide-impacto, fundo petrol, frase grande gold. |
| 5 | **Quem atendemos** | 35-55 anos, alta demanda profissional, sinais de desgaste, exames "normais" mas mal. |
| 6 | **Método AGIR** | 4 pilares: A (Alimentação+Atividade) / G (Gestão metabólica) / I (Integração corpo-mente) / R (Ritmo circadiano). Layout 2x2 grid com ícone+texto cada pilar. |
| 7 | **Escore Plenya** | Big stat 800+ itens avaliados. Visual: dial/gauge. Texto: "pontuação única e evolutiva". |
| 8 | **A equipe integrada** | Médico + nutricionista + psicólogo + educador físico. **Sempre os 4.** Plano único. |
| 9 | **Diferencial** | "Saúde não é sobre reagir. É sobre antecipar." Tabela comparativa: modelo tradicional vs Plenya. |
| 10 | **Jornada do paciente** | Timeline: avaliação inicial → escore → plano integrado → reavaliação contínua. |
| 11 | **Casos / Resultados** | **NUNCA** identificáveis. Perfil composto recorrente ("perfil que atendemos recorrentemente"). Métricas agregadas (qualidade do sono, composição corporal, marcadores metabólicos), sem nomes nem datas específicas. |
| 12 | **Fechamento + contato** | Claim "Viva bem, viva mais." + Plenya logo + canais (site, WhatsApp). Fundo petrol. |

## Regras editoriais invariantes

Aplicar a TODOS os decks gerados por esta skill, sem exceção:

1. **Sem preços.** Nunca citar valores, mensalidades, tabelas comerciais. Negociação comercial é offline.
2. **Sem marcas comerciais.** Nunca citar wearables (Oura, Whoop, Garmin), suplementos por marca, varejistas (Amazon, ML, Vitaminer), labs por nome. Usar categoria genérica + especificação técnica.
3. **Sem "medicina preditiva".** Não está na voz de marca. Usar "antecipatório", "estruturado", "personalizado".
4. **Sem casos identificáveis.** Casos sempre como perfil composto recorrente. LGPD + ética médica.
5. **Sem promessas absolutas.** Sem "garantimos", sem "sempre", sem números de resultado individuais ("perdeu 12kg em 3 meses"). Falar em direção, processo, continuidade.
6. **Sem hashtags em slides** (resíduo de social).
7. **Equipe sempre completa.** Quando mencionar profissionais, citar os 4 perfis. Nunca "médico e nutricionista" sozinho.
8. **800+ itens** no Escore — sempre que falar do Escore, mencionar o número concreto.
9. **40 anos** sempre que falar de origem.
10. **Sem travessões em-dash em copy de slide.** Trocar por vírgula, dois-pontos ou reescrever. (Travessão é AI-tell em 2026.)

## Princípios visuais (consolidados após v2)

1. **TODO slide em fundo petrol** (`#063B4F`). Nunca alternar pra cream — a marca é dark-mode.
2. **Wordmark Plenya real** em capa e fechamento — `apps/site/public/brand/wordmark/cream.png`, embedado como imagem.
3. **Infinity símbolo** (`apps/site/public/brand/symbol/gold.png` ou ocean) como watermark — sutil em cantos de slides de conteúdo, hero centralizado em capa/fechamento.
4. **Wave pattern** em ocean tom-sobre-tom — `docs/decks/_assets/wave-bottom.svg` ou `wave-side.svg` — usar como motif em slides de timeline, equipe, ecosistema.
5. **Cards** sempre `roundRect` com `rectRadius: 0.08–0.1` + borda gold thin (0.5–1pt). Fundo `petrolDeep` (#041F2A) para destacar do background.
6. **Pull-quotes** em cards cream com aspa gold serif gigante (70-180pt) e texto italic Inter.
7. **Letras AGIR gigantes** em Cormorant Garamond bold gold, 130pt+, alinhadas à esquerda dos cards.
8. **Kickers** uppercase gold Inter bold, charSpacing 5, sempre acima do title.
9. **Footer canônico**: linha divisória gold sutil + "PLENYA" mini esquerda + "SAÚDE · PERFORMANCE · LONGEVIDADE" centro + "NN / TT" direita.
10. **Hierarquia clara**: kicker (11pt gold) → title (32-36pt cream serif) → subtitle uppercase sage → body cream Inter.

## Defaults técnicos

- Layout: `LAYOUT_WIDE` (13.33" × 7.5") — formato moderno, mais espaço pra layouts ricos
- Margens: 0.6" laterais, 0.5" topo, 0.45" base (com footer)
- Filename: `<slug-kebab>-<YYYYMMDD>.pptx` em `docs/decks/`
- Speaker notes: incluir sempre — frase de transição + dado de apoio + objeção comum + resposta
- Assets reutilizáveis: helpers em `scripts/deck-builder/plenya-base.js`. Para deck novo, criar `scripts/deck-builder/build-<slug>.js` que importa o base.

## Workflow de execução

Quando invocada:

### 1. Coleta de briefing

Se o pedido não trouxer, perguntar (1 mensagem só, todas as perguntas):
- **Objetivo do deck?** (vender plano PJ, atrair parceiro, apresentar a investidor, onboarding equipe)
- **Público específico?** (ex: RH de empresa 500+ funcionários, gestor de academia, fundo VC)
- **Duração da apresentação?** (15min → 8 slides, 30min → 12 slides, 45min → 16-18)
- **Tem algum ângulo específico a priorizar?** (ex: ROI corporativo, ciência por trás do Método, diferencial da equipe)
- **Onde salvar?** (default: `docs/decks/<slug>-YYYYMMDD.pptx`)

### 2. Estruturação

Com base no briefing, ajustar o template padrão:
- Reordenar slides se o ângulo pede (ex: deck pra RH começa pelo problema corporativo, não pela origem)
- Cortar/adicionar slides — manter os 12 como spinal column
- Definir cada slide com: title, subtitle, body bullets, visual hint, speaker notes

### 3. Geração via pptx skill

Invocar a skill `pptx` com payload completo:
- Paleta, fontes, layout, regras de dominância (passar como spec, não deixar a `pptx` improvisar)
- Estrutura slide-a-slide já pronta
- Pedir formato `pptxgenjs` (criar do zero — não temos template `.pptx` referência por enquanto)

### 4. Output e validação

- Confirmar arquivo gerado em `docs/decks/`
- Reportar ao usuário: caminho do arquivo, nº de slides, tamanho aprox, próximos ajustes possíveis
- **Não abrir o arquivo automaticamente** — usuário decide

## Onde está documentado o que

- **Marca Plenya:** memória `plenya_brand_essence` (40 anos, AGIR, Escore 800+, equipe, frases)
- **Voz/tom:** memórias `plenya_brand_voice_no_preditiva`, `plenya_regras_editoriais`, `plenya_agir_acts_canonical`
- **Casos clínicos:** memórias `linkedin_casos_clinicos_genericos`, `no_chutar_dados_verificaveis`
- **Travessões:** memória `linkedin_no_em_dash`
- **PPTX técnico:** skill `pptx` (irmã desta) — `editing.md` e `pptxgenjs.md`

## Roadmap (futuro, NÃO implementar agora)

- [ ] Template `.pptx` de referência em `docs/decks/_template/plenya-base.pptx` (slide master com logo+cores oficiais embedded)
- [ ] Variante `plenya-deck-cientifico` (voz Getúlio, citações, livro Antes)
- [ ] Variante `plenya-deck-rh` (deck corporativo com ângulo afastamento + ROI)
- [ ] Export Canva opcional via API
