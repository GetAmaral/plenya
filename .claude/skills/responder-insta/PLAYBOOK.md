# PLAYBOOK — Categorias de resposta

> Tipologia de interações que chegam no IG do Dr Getúlio + política de tratamento por tipo.

## Como classificar

Lê o texto da pergunta/mensagem. Procure sinais:
- Verbo de ação ("posso", "devo", "como faço") → pergunta clínica
- Adjetivos de elogio ("incrível", "obrigada", "ótimo") → elogio/agradecimento
- Negações fortes ("isso não é verdade", "exagero") → objeção
- Promo/spam patterns → descartar

Se houver dúvida entre 2 categorias, escolha a **mais conservadora** (mais clínica > menos clínica).

---

## Tipo 1 — Pergunta clínica simples

**Sinais:** pergunta direta sobre tratamento, exame, medicamento, hábito (ex.: "Quando devo fazer creatinina?", "Posso treinar em jejum?", "Vitamina D faz mal?")

**Política:**
- Resposta educativa baseada em evidência (RAG + WebSearch se necessário)
- Estrutura: abertura → tese → mecanismo/dado → fechamento colaborativo
- 600-1.000 chars
- Fechar com convite suave pra consulta SE houver elemento individual (ex.: "depende do seu caso, dá pra avaliar em consulta")
- NÃO citar dose específica sem contexto. OK falar de classe ou intervalo de dose padrão.

**Exemplo de fechamento padrão:**
> "Se quiser ajustar pro seu contexto específico, dá pra avaliar em consulta."

---

## Tipo 2 — Pergunta clínica complexa (caso individual)

**Sinais:** usuário menciona condições próprias específicas (TFG, diagnóstico, lista de medicamentos, comorbidade), pede "posso usar X no meu caso?"

**Política:**
- Mais conservador. Resposta indica caminho mas NÃO prescreve.
- Mencionar variáveis que afetam decisão (idade, função renal, outras comorbidades, medicações concomitantes).
- Cita evidência principal (1-2 estudos por nome).
- Fechamento canônico **mais forte** pra consulta presencial:
  > "No seu caso, a princípio vale a pena considerar — conversa com seu nefrologista e decide junto com ele."
- 1.000-1.500 chars (caso justifica detalhe)
- Se a pergunta exige diagnóstico ("estou com X, é isso mesmo?"): **recusa explícita** + redirecionamento.
  > "Diagnóstico não dá pra fechar pelo Instagram. Procura um nefrologista pra investigar direito."

---

## Tipo 3 — Pedido de diagnóstico ou prescrição

**Sinais:** "Tenho dor X, é Y?", "Que remédio eu tomo pra...?", "Me passa receita de..."

**Política:**
- **Recusa firme mas gentil.**
- Não inventar diagnóstico baseado em texto pequeno.
- Não passar receita NUNCA — sequer mencionar dose específica.
- Encaminhar pra avaliação:
  > "Pra avaliar isso direito preciso de história, exame físico e exames. Procura uma consulta — comigo se quiser, senão com qualquer nefrologista da sua rede."
- 200-400 chars (curto e direto)

---

## Tipo 4 — Elogio / Agradecimento

**Sinais:** "Obrigado(a) doutor!", "Que aula maravilhosa!", "Conteúdo incrível!", emojis 👏❤️🙏 isolados ou com curta frase

**Política:**
- Resposta curta e autêntica. **NÃO** falar de medicina.
- Gratidão direta, primeira pessoa, 1 emoji opcional no fim.
- 50-200 chars.
- Exemplos:
  > "Muito obrigado pelo feedback! Fico feliz que tenha agregado. 🙏🏻"
  > "Valeu demais! Esse é o objetivo — informação útil e baseada em evidência."

**Anti-padrão:**
- ❌ Não engatar conteúdo médico ("e por falar nisso, sabia que...")
- ❌ Não fechar com CTA de consulta
- ❌ Não mais que 2 emojis

---

## Tipo 5 — Objeção / Crítica

**Sinais:** "Isso não é verdade", "Médicos só querem vender remédio", "Vocês exageram"

**Política:**
- Resposta **sem confronto**. Validar parcialmente o ponto se houver substrato (ex.: sim, conflitos de interesse existem na medicina). Apresentar evidência.
- Estrutura: reconhecer a preocupação → apresentar dado/mecanismo → fechar abrindo diálogo.
- 400-800 chars.
- Exemplo de abertura:
  > "Entendo a preocupação — e tem fundamento histórico nisso. Mas no caso específico de [X], a evidência atual mostra..."
- **NUNCA** ironia, sarcasmo, "respondi de forma educada mas com agressividade".
- Fechamento: "Discussão saudável — abre espaço pra mais perguntas se quiser."

---

## Tipo 6 — Spam / Bot / Não-relevante

**Sinais:** texto promocional ("ganhe X mil reais"), idioma estranho ao contexto, link suspeito, padrão de bot (resposta genérica copia-cola que não engaja com o tema)

**Política:**
- **Não responder.**
- **NÃO usar `INSTAGRAM_DELETE_COMMENT`** sem pedido explícito do usuário — irreversível e pode afetar usuários reais marcados como falso positivo.
- Marcar como "descartado" no relatório final só pra contagem.
- Se for ofensivo (xingamento, ataque pessoal): trazer pro usuário decidir se quer hide/delete manualmente.

---

## Tipo 7 — Agendamento / Comercial

**Sinais:** "Quero marcar consulta", "Como funciona o atendimento?", "Quanto custa?", "Atende plano X?"

**Política:**
- DM: resposta cordial direcionando pro canal certo.
- Comentário em post público: cuidado com info comercial em público. Direciona pra DM ou link de agendamento.
- Resposta tipo:
  > "Olá [nome]! Pra agendar e tirar dúvidas sobre o atendimento, fala com a gente no WhatsApp da Plenya: [WhatsApp link]. Ou no link da bio. Te aguardamos!"
- **NUNCA** dar preço/condições em público — sempre redirecionar pra contato direto.
- Verificar memória `getulio_canonical_handles.md` pra links corretos da Plenya.

---

## Casos especiais — posts pessoais

Quando o post é **pessoal/familiar** (ex.: Dia das Mães, aniversário, anúncio de gestação, perda):

- **Tratamento clínico = NÃO aplicar.** Mesmo se alguém comentar algo médico ali, é fora de tom.
- **Volume de comentários em massa (parabéns/❤️):** sugere ao usuário que **não responda individualmente** — curtir todos é suficiente. Apresenta a lista de quem comentou e pergunta se quer responder alguém específico (família, amigos íntimos).
- **Quando responder individualmente** (esposa, irmãos, primos, padrinhos): tom 100% pessoal, sem qualquer coisa médica. Curto, afetuoso.

Como identificar post pessoal: olhar caption.
- Pessoal: menciona família por nome, marcos pessoais, sem citações científicas
- Educacional: começa com pergunta retórica clínica, tem referências, hashtags médicas

---

## Casos especiais — DMs >24h fora da janela

Se a última mensagem do usuário tem >24h, a janela "free messaging" da Meta fechou. Opções:
1. **Não responder** (a mais segura)
2. **Reply via "human agent"** — Meta permite respostas humanas até 7 dias após última msg (se o app tem tag `HUMAN_AGENT`). Skill **não sabe** se o tag está aprovado — então AVISA o usuário antes de tentar:
   > "Atenção: essa DM tem [N]h, fora da janela de 24h da Meta. Posso tentar responder, mas pode falhar com erro de política. Procede mesmo assim?"

Se ele aprova e a API der erro, registrar no relatório final como "tentativa falhou — janela 24h fechada" e seguir.
