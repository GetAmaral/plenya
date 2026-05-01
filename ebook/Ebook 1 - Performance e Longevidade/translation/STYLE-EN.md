# BEFORE — Style Bible (PT-BR → EN tradução)

Documento de referência da tradução do livro **ANTES** para edição internacional **BEFORE**. Tudo abaixo é decisão fechada. Antes de traduzir qualquer trecho, ler este documento inteiro.

Versão: 1.0 · Data: 2026-05-01 · Manutenção: Plenya / Editorial.

---

## 0. Documentos relacionados

- `GLOSSARY.md` (mesma pasta) — tabela PT→EN de termos travados (~150 entradas).
- `PRE-AUDIT.md` (mesma pasta) — auditoria capítulo a capítulo dos itens BR-específicos a adaptar.
- `../md/pt-BR/00-indice.md` — fonte da verdade para AGIR↔ACTS mapping.
- `../md/pt-BR/briefing-branding.md` — fonte da verdade para voz, posicionamento, personagens.
- `apps/site/TRANSLATION-GUIDE.md` (raiz do monorepo) — guia gêmeo para o site Plenya.

---

## 1. Decisões macro (já travadas)

| Decisão | Valor |
|---|---|
| Variante | **US English** |
| Texto-fonte PT | **Congelado** — não há ajustes editoriais pendentes |
| Personagens | **Fictícios** (compostos clínicos) — sem issue de consentimento |
| Plenya | **Brand internacional** — não traduzir |
| Imprint | **Edição do Autor** (mesmo selo PT) |
| ISBN | **Novo** para EN — registrar no CBL antes do build |
| Referências (Cap 18) | **Manter idênticas** (estudos já estão em inglês) |
| Index | **Sem índice remissivo** |
| Agradecimentos | **Traduzir** — sem mencionar Goiânia (não está no original) |

---

## 2. Título e identidade da obra

**Título:** *BEFORE*
**Subtítulo:** *The Silent Window Between Normal and Optimal — where health is decided.*

- Capa **não cita o método ACTS** (paralelo à capa PT — método aparece dentro do livro).
- Autor: **Dr. Getulio Amaral Filho** (sem acento; assim mesmo na versão internacional).
- Bio: *"Brazilian nephrologist and clinical physician, founder of Plenya — a precision medicine clinic for healthspan and longevity."* (Plenya como diferencial; Brazilian como contexto, não como limitação.)

---

## 3. Posicionamento autoral

A voz autoral é **médico-narrador**: alguém que prescreve longevidade no consultório todos os dias, não pesquisador acadêmico, não influenciador de Instagram.

### Pares de contraste — o que SOMOS vs o que NÃO SOMOS

| Somos | Não somos |
|---|---|
| Medicina preventiva proativa | Medicina reativa convencional |
| Ciência que prescreve | Ciência que apenas contempla |
| Autoridade clínica brasileira global | Guru de Instagram / coach de longevidade |
| Premium acessível | Luxury elitista |
| Narrativa anedótica + dado | Manual técnico sem alma |
| Direto, claro, prático | Acadêmico, hedge, qualifying everything |

### Modelos literários-âncora (referência de tom para o tradutor)

Quando estiver inseguro sobre tom de um trecho, pergunte: *"como isso seria escrito em…"*

1. **Outlive — Peter Attia (2023)** — *âncora primária*. Frases médias, voz primeira pessoa do médico, anedotas clínicas, ciência embutida no caso. Esse é o tom que o leitor-alvo já conhece e espera.
2. **Why We Sleep — Matthew Walker (2017)** — para Cap 14 (sono/ritmo).
3. **Lifespan — David Sinclair (2019)** — para discussões sobre senescência e biomarcadores.
4. **Lost Connections — Johann Hari (2018)** — para Cap 13 (conexão/propósito).
5. **The Body Keeps the Score — Bessel van der Kolk (2014)** — para Cap 12 (mente/trauma).
6. **Mary Oliver, "Sometimes"** — referência para o uso de "tend" ("Pay attention. Be astonished. Tell about it.")

---

## 4. Voz, registro, ritmo

### Pessoa gramatical

- **"I"** — narrador médico. Casos clínicos, opiniões, decisões clínicas. Default no livro inteiro.
- **"You"** — leitor. Quando ensinando, prescrevendo, perguntando.
- **"We"** — Plenya / equipe / medicina como prática coletiva. Usar com parcimônia.
- **"They"** — pacientes apresentados em terceira pessoa nos casos.

### Ritmo

- **Frases curtas-médias.** Médias quando explicando ciência. Curtas quando concluindo, prescrevendo, marcando ponto.
- **Parágrafos médios** — 3 a 6 frases. Parágrafo de 1 frase é OK como ponto de impacto.
- **Variação rítmica é regra.** Três frases longas seguidas matam a página.
- **Itálicos para ênfase** (não bold, exceto em headers).

### Registro

- **Inglês conversacional educado.** Não acadêmico. Não casual demais.
- **Pode usar contrações.** "It's", "don't", "we're" — são OK e tornam o texto mais humano. Attia usa.
- **Termos técnicos são OK** quando contextualizados na primeira menção. Não inventar perífrase para evitar termo técnico.
- **Evitar hedging excessivo.** Não escrever "it might possibly be argued that perhaps". Escrever "the data suggests" ou "we know".

### Padrões a manter

- **Voz ativa > voz passiva** sempre que cabível.
- **Verbo presente** quando descrevendo ciência atemporal ou prática clínica.
- **Verbo passado** apenas em casos clínicos narrativos.

### Padrões a evitar

- *Translatedese* — frases que soam traduzidas. Sintoma: leitor americano hesita ao ler. Cura: ler em voz alta; se travar, reescrever.
- **Latinismos desnecessários** — *prior to* → *before*; *utilize* → *use*; *commence* → *start*; *terminate* → *end*. Inglês US prefere a palavra anglo-saxã quando há.
- **Falsos cognatos** (ver `GLOSSARY.md` seção Falsos Cognatos).
- **"Sublime", "magnificent", "ultimate"** e outros adjetivos hiperbólicos — soam Instagram health coach.
- **"Solution", "answer", "key", "secret"** — vocabulário de self-help. Substituir por "approach", "strategy", "mechanism", "lever".

---

## 5. ACTS Method — referência canônica

**Recapitulando** (a tabela completa está em `GLOSSARY.md` e em `briefing-branding.md`):

| Letra | Curta (capa/método) | Expandida (1ª menção) |
|---|---|---|
| **A** | Activity, Alimentation & Smart Adjuncts | Activity (Exercise & Movement), Alimentation (Eating Well) & Smart Adjuncts |
| **C** | Clinical Optimization | Clinical Optimization (Systems, Panels & Exposures) |
| **T** | Tending Mind, Body & Bonds | Tending Mind, Body & Bonds (Inner Work + Connection) |
| **S** | Sleep, Rhythm & Recovery | Sleep, Rhythm & Recovery (Sleep + Circadian Timing + Active Rest) |

### Regras de uso no livro

- **Forma expandida:** primeira menção de cada pilar **na introdução da Parte III** (Cap 6b). Em todos os outros lugares, forma curta.
- **Substantivo do bloco:** *Pillar A*, *Pillar C*, *Pillar T*, *Pillar S*. **Não** *Practice*, *Step*, *Phase*.
- **Substantivo do método:** *The ACTS Method* (sempre com artigo, sempre capitalizado). **Não** *the ACTS framework* nem *ACTS approach* como termo de marca.
- **Ordem inviolável:** A → C → T → S. Nunca reordenar por estética.
- **Tagline canônica do método:** *"AGIR means 'to act' in Portuguese. ACTS is its English counterpart — four pillars that, together, protect the decades of health that standard medicine gives up on."*

### O que dizer em vez de "AGIR"

Quando o PT diz "Método AGIR" ou "AGIR" como termo, o EN usa **"The ACTS Method"** ou **"ACTS"**. **Exceção única:** quando o autor explica a origem etimológica/conceitual em uma única passagem (provavelmente Cap 6b), aí pode-se dizer:

> *"In Portuguese, the method is called AGIR — which means 'to act'. The English counterpart, ACTS, preserves both the imperative spelling and the underlying call to action."*

Fora desse momento explicativo, **não usar "AGIR" em texto inglês.**

---

## 6. Personagens-âncora — nomes e localização

| Personagem PT | EN | Notas |
|---|---|---|
| Ricardo, 52 | **Ricardo, 52** | Manter nome. Ancora a brasilianidade do médico-narrador (paralelo a "Joe", "Sam" em Outlive). |
| Fernanda, 41 | **Fernanda, 41** | Manter. |
| André, 45 | **André, 45** | Manter o acento (US imprime acentos sem problema; ele é parte do nome). |
| Marcos, 57 | **Marcos, 57** | Manter. |
| Paulo, 48 | **Paulo, 48** | Manter. |
| Ana, 44 | **Ana, 44** | Manter. |

**Não anglicizar** (nada de "Richard", "Andrew", "Paul"). Os nomes brasileiros são parte do diferencial autoral — médico brasileiro com prática real, não tradução genérica.

**Profissões e contexto dos casos:** se o PT diz "executivo de São Paulo", o EN pode dizer "São Paulo executive" — manter referência geográfica brasileira, sem traduzir cidade. Se for ornamental e não-essencial, considere generalizar para "executive" sem cidade.

---

## 7. Unidades de medida

| Categoria | Estratégia |
|---|---|
| **Biomarcadores** (colesterol, glicose, etc.) | **Manter mg/dL** (US usa idêntico ao Brasil; UK usa mmol/L mas o livro é US). |
| **Hormônios** | Manter unidade do estudo original (ng/dL, pmol/L, etc.). |
| **Distância em copy descritivo** | Imperial primário com métrico em parênteses na 1ª menção: *"a 5-mile walk (8 km)"*. |
| **Distância em prescrição/treino** | Imperial primário ("run 3 miles, twice a week"). |
| **Peso corporal em copy** | Imperial primário: *"180 lbs (82 kg)"*. |
| **Peso em prescrição/dose** | **Manter kg** (científico): *"1.6 g/kg of protein per day"*. |
| **Altura** | Pés/polegadas: *"5'10" (178 cm)"*. |
| **Temperatura ambiente/clima** | Fahrenheit primário com Celsius em parênteses na 1ª menção. |
| **Temperatura corporal/clínica** | **Celsius** (mesma convenção médica internacional). |
| **Velocidade em treino** | Manter km/h ou mph conforme contexto do estudo citado; em prescrição genérica, usar mph. |

**Princípio:** no espírito do leitor americano (familiar com imperial), mas sem distorcer ciência (estudos publicados em métrico mantêm métrico).

---

## 8. Adaptação cultural e regulatória

Detalhes específicos por capítulo estão em `PRE-AUDIT.md`. Princípios gerais:

### Anvisa, CFM, SUS, ANS

Nunca traduzir como "FDA", "AMA", etc. — incorreto e legalmente confuso. Em vez disso:

- **Anvisa** → *"Brazilian regulators (Anvisa)"* + se necessário, nota de equivalência.
- **CFM** → *"Brazil's medical authority (Conselho Federal de Medicina)"*.
- **SUS** → *"Brazil's public health system (SUS)"*.
- **Plano de saúde / convênio** → *"private health insurance"* ou *"healthcare coverage"* (depende do contexto).

### Estatísticas brasileiras

Quando o PT cita dado IBGE/DATASUS:
- **Manter o dado BR** como anchor narrativo.
- **Adicionar dado equivalente CDC/WHO** entre parênteses ou em frase seguinte. Ex.: *"In Brazil, 70% of adults… (in the U.S., the CDC reports 73%)."*

### Anedotas culturais

- **Manter** quando essenciais à narrativa ou ao charme autoral (ex.: rotina de paciente brasileiro, particularidades do sistema BR).
- **Substituir** quando ornamentais e o equivalente cultural mudaria a metáfora (ex.: futebol → especificar ou universalizar para "soccer/football, depending on where you live"). Avaliar caso a caso.

### Estudos brasileiros

- **ELSA-Brasil** e outros cohorts brasileiros — manter como diferencial. Adicionar em parênteses *"(a Brazilian longitudinal study of adult health, comparable to the U.S. Atherosclerosis Risk in Communities cohort)"* na primeira menção.
- **Pesquisador BR citado** — manter; legitima a expertise brasileira do autor.

---

## 9. Convenções de pontuação e formatação

- **Vírgula de Oxford:** sempre usar.
- **Datas:** *April 28, 2026* (não *28 April 2026*).
- **Aspas:** duplas curvas *"…"* (não retas). Aspas simples *'…'* apenas para citação dentro de citação.
- **Travessão (em-dash):** *—* sem espaços. *"the silent window — between normal and optimal"* fica *"the silent window—between normal and optimal"* em US convention. **Atenção:** o autor usa em-dash com espaços em PT e em copy do site EN. Para o livro, **manter em-dash com espaços** (escolha de design tipográfico do livro PT) — isso é coerente com o que já foi gerado.
- **Itálico:** para ênfase suave, títulos de obras, palavras estrangeiras (incluindo "AGIR" quando referenciado).
- **Bold:** apenas para headers. Não usar bold em prosa para ênfase.
- **Capitalização de títulos:** *Title Case* para H1, H2, títulos de capítulo. *Sentence case* para subseções.

---

## 10. Workflow recomendado para o tradutor

1. **Leia o capítulo PT inteiro antes de traduzir** — entender arco narrativo evita travar em frase isolada.
2. **Identifique pré-audit items do capítulo** (`PRE-AUDIT.md`) — saber de antemão o que vai precisar adaptar.
3. **Traduza com glossário aberto** (`GLOSSARY.md`) — termos travados não negociam.
4. **Primeiro draft em prosa fluente** — não trave em palavra perfeita; marque com `[?]` e siga.
5. **Releia em voz alta** — translatedese só revela na boca.
6. **Self-checklist** (final do `apps/site/TRANSLATION-GUIDE.md` seção 7) — passar antes de entregar.
7. **Lista de dúvidas no final** — `[?]` items consolidados em comentário ao revisor.

---

## 11. Critérios de qualidade do produto final

Um capítulo está pronto quando:

- [ ] Termos do glossário aplicados consistentemente.
- [ ] Voz primeira pessoa do médico mantida.
- [ ] ACTS / Pillar / Plenya seguem convenções fechadas.
- [ ] Anvisa / SUS / dados BR adaptados conforme PRE-AUDIT.
- [ ] Unidades convertidas conforme tabela seção 7.
- [ ] Personagens com nomes brasileiros preservados.
- [ ] Vírgula de Oxford presente.
- [ ] Spelling 100% US (não UK).
- [ ] Lido em voz alta sem travas.
- [ ] Título do capítulo em Title Case.
- [ ] Casos clínicos com idade/biomarcadores conferidos.
- [ ] Referências literárias-âncora (Outlive, Walker) ressoam quando aplicável.

---

## 12. Mudanças neste guia

Qualquer alteração estrutural (variante, ACTS, posicionamento, voz) precisa de **aprovação editorial** antes de ser aplicada. Pequenos ajustes de glossário e adições ao PRE-AUDIT podem ser feitos pelo tradutor com nota de mudança no commit.
