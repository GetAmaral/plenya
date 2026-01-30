-- BATCH 41: SONO PARTE 3 - FINAL
-- Items finais: Terceira seção de Equipamento, Ambiente e Higiene do Sono

-- ============================================
-- SEÇÃO: EQUIPAMENTO DO SONO (TERCEIRA SEÇÃO - ATUAL)
-- ============================================

-- Equipamento do sono (categoria - terceira seção)
UPDATE score_items SET
  clinical_relevance = 'Equipamento de sono atual (cama, colchão, travesseiro, roupa de cama) é determinante modificável de qualidade do sono e dor musculoesquelética. Investimento em equipamento adequado tem ROI (retorno sobre investimento) alto: melhora qualidade de sono, reduz dor crônica (lombar, cervical), aumenta produtividade, reduz necessidade de analgésicos. Custo-benefício superior a medicamentos. Priorização: colchão >10 anos = troca urgente, travesseiro >2 anos = trocar, cama pequena para casal = upgrade para Queen/King.',
  values_low = '{"min":0,"max":0,"interpretation":"Não aplicável - categoria de agrupamento."}',
  values_high = '{"min":1,"max":999,"interpretation":"Não aplicável - categoria de agrupamento."}'
WHERE id = '9ca5a90d-ed46-49f7-9e53-013004f6e3db';

-- Cama (terceira seção - atual)
UPDATE score_items SET
  clinical_relevance = 'Tamanho atual da cama é fator frequentemente negligenciado mas com impacto significativo na qualidade do sono de casais. Cama de casal padrão (138cm de largura) fornece apenas 69cm por pessoa - menos que um berço infantil (70cm). Espaço insuficiente aumenta despertares noturnos por movimentos do parceiro (média 6-8 movimentos/noite), eleva temperatura corporal por proximidade excessiva (dificultando termorregulação), reduz tempo em sono profundo. Estudos mostram que upgrade de cama de casal para Queen/King reduz despertares em 30-40%.',
  values_low = '{"min":0,"max":0,"interpretation":"Não aplicável - item descritivo."}',
  values_high = '{"min":1,"max":999,"interpretation":"Avaliar tamanho atual: solteiro 88cm (adequado para 1 pessoa), casal 138cm (inadequado para 2 adultos, apenas 69cm/pessoa), queen 158cm (79cm/pessoa, mínimo aceitável), king 193cm (96,5cm/pessoa, ideal), super king 203cm. Cama de casal para dois adultos = upgrade PRIORITÁRIO. Avaliar também: estabilidade (rangidos ao movimentar), adequação ao peso combinado (reforço central se >160kg), altura da cama (facilidade de entrar/sair, especialmente idosos)."}'
WHERE id = 'a59269ac-50fb-4aac-9cf8-ea99accfeb68';

-- Colchão (terceira seção - atual)
UPDATE score_items SET
  clinical_relevance = 'Colchão atual adequado é investimento de maior ROI em qualidade de sono e saúde musculoesquelética. Colchão desgastado (>8-10 anos), com afundamento >5cm ou perda de suporte causa: desalinhamento espinhal (lordose/cifose excessiva em decúbito), dor lombar matinal (melhora após 30-60min de atividade), cervicalgia, sono fragmentado. Firmeza adequada depende de IMC: IMC <25 = médio-macio, IMC 25-30 = médio-firme, IMC >30 = firme. Teste de adequação: mão não deve passar facilmente sob lombar em decúbito dorsal (muito firme) nem afundar >5cm (muito mole).',
  values_low = '{"min":0,"max":0,"interpretation":"Não aplicável - item descritivo."}',
  values_high = '{"min":1,"max":999,"interpretation":"Avaliar detalhadamente: idade do colchão (trocar obrigatoriamente a cada 8-10 anos, ou antes se desgaste visível), tipo (molas ensacadas: durabilidade e suporte, espuma de alta densidade 28-33kg/m³, látex natural: hipoalergênico, viscoelástico: alívio de pressão), firmeza adequada ao IMC atual, presença de afundamento (teste: deitar e verificar se >5cm), deformidade visível, dor lombar/cervical/ombros matinal. Dor matinal que melhora após atividade = colchão inadequado (troca prioritária). Investimento típico: R$2.000-5.000, durabilidade 8-10 anos."}'
WHERE id = '7b24a7f0-7c7a-415c-a182-eba24bbcfd43';

-- Travesseiros (terceira seção - atual)
UPDATE score_items SET
  clinical_relevance = 'Travesseiro atual adequado é essencial para alinhamento cervical e prevenção de cervicalgia/cefaleia tensional. Travesseiro inadequado (altura, firmeza, idade >2 anos) é causa tratável de dor cervical matinal e sono fragmentado. Altura ideal depende de posição de dormir predominante: (1) Lado (70% população): travesseiro alto (10-15cm) para preencher distância ombro-orelha, (2) Costas (20%): travesseiro médio (8-10cm), (3) Bruços (10%): sem travesseiro ou muito baixo (não recomendado por rotação cervical excessiva). Teste de adequação: orelha deve alinhar com ombro em decúbito lateral.',
  values_low = '{"min":0,"max":0,"interpretation":"Não aplicável - item descritivo."}',
  values_high = '{"min":1,"max":999,"interpretation":"Avaliar detalhadamente: número de travesseiros (empilhar 2+ travesseiros = inadequado, causa hiperflexão cervical), altura atual, firmeza (muito mole perde suporte, muito firme causa pressão excessiva), idade (trocar a cada 1-2 anos: ácaros acumulados, perda de suporte), tipo (espuma: econômico, látex: durável e hipoalergênico, viscoelástico: contorna cervical, penas: não recomendado), posição de dormir (lado/costas/bruços), cervicalgia matinal (melhora ao longo do dia?), cefaleia tensional. Travesseiro inadequado + cervicalgia = travesseiro cervical ergonômico (investimento R$200-500)."}'
WHERE id = 'c6fe49d7-6d6c-4127-9cd6-4259d3ccb69f';

-- Pijamas (terceira seção - atual)
UPDATE score_items SET
  clinical_relevance = 'Vestimenta noturna atual impacta diretamente termorregulação, essencial para iniciação e manutenção do sono. Fisiologia: temperatura corporal central deve cair 0,5-1°C para iniciar sono (vasodilatação periférica dissipa calor). Roupas muito quentes (flanela, lã, sintéticos) impedem dissipação de calor → superaquecimento → sudorese → despertares noturnos. Tecidos sintéticos (poliéster, acrílico, nylon) retêm calor e umidade (má respirabilidade). Tecidos naturais (algodão, linho, bambú, modal, tencel) permitem respirabilidade e termorregulação adequada. Dormir sem roupa ou roupa mínima é ideal em climas quentes (>25°C).',
  values_low = '{"min":0,"max":0,"interpretation":"Não aplicável - item descritivo."}',
  values_high = '{"min":1,"max":999,"interpretation":"Avaliar: tipo de roupa noturna atual (pijama completo pesado, leve, camiseta+short, roupa íntima, sem roupa), tecido (sintético: poliéster/acrílico vs natural: algodão/linho/bambú), adequação à temperatura ambiente do quarto (18-21°C ideal), presença de desconforto térmico/sudorese noturna, estação do ano. Roupas pesadas + ambiente quente + sudorese = trocar por tecidos naturais leves ou sem roupa. Recomendar: algodão leve/bambú (verão), camadas ajustáveis (inverno), sem roupa se preferência."}'
WHERE id = 'db478afd-c211-4667-9d80-cc142b56e5cd';

-- Lençóis e cobertas (terceira seção - atual)
UPDATE score_items SET
  clinical_relevance = 'Roupa de cama atual (lençóis, cobertas) impacta termorregulação, higiene e qualidade do sono. Lençóis: tecidos sintéticos (microfibra poliéster) retêm calor e umidade → sudorese noturna e despertares. Tecidos naturais (algodão percal 200-400 fios, algodão cetim 300-600 fios, linho, bambú) melhoram respirabilidade e conforto térmico. Cobertas: peso inadequado (muito pesada: superaquecimento, muito leve: frio) causa despertares. Temperatura ideal de sono: 18-21°C. Higiene: lavar lençóis semanalmente (ácaros, células mortas, suor acumulado → alergia respiratória, asma, rinite).',
  values_low = '{"min":0,"max":0,"interpretation":"Não aplicável - item descritivo."}',
  values_high = '{"min":1,"max":999,"interpretation":"Avaliar: tipo de tecido dos lençóis (natural: algodão/linho/bambú vs sintético: microfibra/poliéster), contagem de fios (algodão percal 200-400: fresco, cetim 300-600: mais quente), peso das cobertas (leve 200-300g/m², média 300-500g/m², pesada >500g/m²), adequação à temperatura ambiente, frequência de troca (semanal mínimo), presença de sudorese noturna, sintomas alérgicos (rinite, asma). Lençóis sintéticos + sudorese = trocar por algodão. Recomendar: algodão percal (verão), cetim (inverno), lavar semanalmente 60°C (elimina ácaros)."}'
WHERE id = 'f6db908c-880d-4a6b-a091-4499560783cd';

-- ============================================
-- SEÇÃO: AMBIENTE DO SONO (TERCEIRA SEÇÃO - ATUAL)
-- ============================================

-- Ambiente do sono (categoria - terceira seção)
UPDATE score_items SET
  clinical_relevance = 'Ambiente atual do sono é conjunto de fatores modificáveis de ALTA prioridade terapêutica, com evidência robusta de impacto na qualidade do sono. Componentes críticos e metas: (1) Escuridão total <1 lux (cortinas blackout + remover LEDs), (2) Temperatura 18-21°C (AC, ventilador, janela), (3) Silêncio <30 dB ou ruído branco 40-50 dB, (4) Zero telas 1-2h pré-sono (luz azul <10 lux), (5) Aromas relaxantes (lavanda: +20% sono profundo). Otimização do ambiente é intervenção de primeira linha na TCC-I, eficácia 70-80%, superior a hipnóticos, sem efeitos adversos, custo-benefício excelente.',
  values_low = '{"min":0,"max":0,"interpretation":"Não aplicável - categoria de agrupamento."}',
  values_high = '{"min":1,"max":999,"interpretation":"Não aplicável - categoria de agrupamento."}'
WHERE id = '6526cc60-99d8-4147-972b-e88d4a9b2acc';

-- Quarto (terceira seção - atual)
UPDATE score_items SET
  clinical_relevance = 'Características atuais do quarto são determinantes modificáveis de qualidade do sono. Fatores críticos: (1) Tamanho adequado: mínimo 9m² para casal, 6m² para solteiro (claustrofobia aumenta ansiedade e cortisol noturno), (2) Ventilação adequada: CO2 >1.000 ppm causa cefaleia matinal, déficit cognitivo, sono fragmentado (abrir janela ou ventilador), (3) Organização: desordem visual aumenta cortisol basal em 20%, dificulta relaxamento (técnica: arrumar quarto 30min pré-sono como ritual), (4) Uso exclusivo para sono/sexo: cama usada para trabalho/estudo cria associação mental incompatível com sono (princípio da TCC-I).',
  values_low = '{"min":0,"max":0,"interpretation":"Não aplicável - item descritivo."}',
  values_high = '{"min":1,"max":999,"interpretation":"Avaliar detalhadamente: tamanho do quarto (m², sensação de claustrofobia/excesso de espaço), ventilação (janelas abertas à noite? ventilador? AC?), CO2 estimado (cefaleia matinal sugere >1.000 ppm), organização (desordem visual: roupas espalhadas, objetos acumulados), uso múltiplo (home office no quarto: mesa de trabalho, computador), sensação subjetiva de conforto/segurança. Quarto desorganizado ou usado para trabalho = intervenção prioritária: remover equipamentos de trabalho, organizar 30min pré-sono, uso exclusivo para sono/sexo."}'
WHERE id = 'e4265a0d-e800-420b-b653-049c48cd975e';

-- Iluminação (terceira seção - atual)
UPDATE score_items SET
  clinical_relevance = 'Iluminação noturna atual é fator modificável de MÁXIMA prioridade, com impacto direto e dose-dependente na supressão de melatonina. Fisiologia: luz >10 lux detectada por células ganglionares retinianas intrinsecamente fotossensíveis (ipRGCs, sensíveis a azul 450-480nm) → supressão de melatonina pelo núcleo supraquiasmático (NSQ). Dose-resposta: 10-30 lux suprime 20%, 50-100 lux suprime 50%, >200 lux suprime >80%. Escuridão total <1 lux é meta. Fontes comuns: luz de rua (100-1.000 lux), LEDs de eletrônicos (5-50 lux), luz de banheiro. Intervenção: cortinas blackout (bloqueio >99%) + máscara de dormir + remover/cobrir LEDs.',
  values_low = '{"min":0,"max":0,"interpretation":"Não aplicável - item descritivo."}',
  values_high = '{"min":1,"max":999,"interpretation":"Avaliar PRIORITARIAMENTE: escuridão total (<1 lux: não consegue ver a própria mão) ou presença de luz (janela sem blackout: luz de rua/postes 100-1.000 lux, luz de quarto ao lado vazando pela porta, eletrônicos: despertador digital 10-30 lux, TV em standby 5-10 lux, roteador WiFi, carregadores com LED), uso de cortinas blackout ou máscara de dormir. Luz visível no quarto = supressão significativa de melatonina. INTERVENÇÃO PRIORITÁRIA: cortinas blackout (investimento R$200-500, ROI alto), máscara de dormir (R$30-100), remover eletrônicos com LEDs, cobrir LEDs inevitáveis com fita isolante preta, despertador analógico ou com luz vermelha <5 lux."}'
WHERE id = 'ab010c3d-1001-412c-b62f-652209abc623';

-- Janelas / claridade (terceira seção - atual)
UPDATE score_items SET
  clinical_relevance = 'Janelas no quarto têm papel DUAL essencial no ritmo circadiano: (1) Manhã: entrada de luz natural >1.000 lux (sol direto ou próximo à janela aberta) sincroniza relógio circadiano, suprime melatonina residual, aumenta cortisol e alerta, adianta fase circadiana (facilita acordar mais cedo), melhora humor (previne depressão sazonal). (2) Noite: entrada de luz de rua >10 lux suprime melatonina, atrasa sono. Solução: cortinas blackout (bloqueio >99% à noite) + abertura completa pela manhã. Posição ideal: janela face LESTE (sol matinal 6-8h, intensidade 10.000-100.000 lux).',
  values_low = '{"min":0,"max":0,"interpretation":"Não aplicável - item descritivo."}',
  values_high = '{"min":1,"max":999,"interpretation":"Avaliar: presença de janelas no quarto (se ausente: exposição matinal insuficiente, risco de atraso de fase), exposição à luz natural matinal (abre cortinas ao acordar? tempo de exposição >30min? face leste ideal), cortinas blackout (presentes? eficácia de bloqueio), luz de rua/postes à noite (intensidade estimada), posição geográfica (leste: sol matinal ideal, oeste: sol poente pode atrasar sono, norte/sul: luz lateral). INTERVENÇÃO: cortinas blackout obrigatórias se luz de rua + abertura TOTAL das cortinas ao acordar (exposição >30min, >1.000 lux) + considerar mudança para quarto face leste se cronótipo vespertino problemático."}'
WHERE id = '8533025d-d0bb-4fc1-b2ca-df71df8a148e';

-- Televisão / telas (terceira seção - atual)
UPDATE score_items SET
  clinical_relevance = 'TV e dispositivos eletrônicos no quarto atualmente são obstáculos de MÁXIMA prioridade à higiene do sono. Mecanismos triplos: (1) Luz azul 450-480nm (pico de sensibilidade das ipRGCs): smartphone a 30cm emite 30-60 lux, suprime melatonina >50% por 2h, atrasa fase circadiana 1-2h. (2) Estimulação cognitiva: conteúdo ativador (trabalho, redes sociais, notícias, jogos) aumenta cortisol, frequência cardíaca, atividade simpática, inibe transição vigília-sono. (3) Condicionamento comportamental: associação mental de quarto/cama com vigília (não sono). Regra de ouro TCC-I: quarto livre de telas, celular em modo avião FORA da cama (outro cômodo ideal).',
  values_low = '{"min":0,"max":0,"interpretation":"Não aplicável - item descritivo."}',
  values_high = '{"min":1,"max":999,"interpretation":"Avaliar CRITICAMENTE: TV no quarto (presente? ligada até adormecer? adormece com TV ligada?), celular (localização: sob travesseiro/criado-mudo/cama, modo avião?), uso de telas <1h pré-sono (minutos, dispositivos), brilho (máximo/médio/baixo), tipo de conteúdo (trabalho: máxima ativação, redes sociais: dopamina e comparação social, vídeos/séries: ativação emocional, jogos: ativação máxima). TV/celular na cama até adormecer = PIOR cenário (supressão máxima de melatonina + condicionamento inadequado). INTERVENÇÃO MÁXIMA PRIORIDADE: remover TV do quarto, celular em modo avião em outro cômodo (ou mínimo >1m da cama, modo avião), substituir por leitura (livro físico, luz âmbar <30 lux), blue-blockers (Uvex SCT-Orange) se uso inevitável."}'
WHERE id = 'af4763f5-a7bc-4594-a058-797c371bd7b1';

-- Barulho (terceira seção - atual)
UPDATE score_items SET
  clinical_relevance = 'Ruído noturno atual é fator ambiental crítico. Fisiologia: ruído >35 dB (conversação baixa) causa microdespertares (fragmentação do sono, redução de N3 e REM), ativação simpática (aumento de frequência cardíaca, pressão arterial), aumento de cortisol noturno. Ruído intermitente (trânsito, latidos, vizinhos) é mais disruptivo que ruído constante (ventilador, AC). Ruído crônico noturno >55 dB aumenta risco cardiovascular: hipertensão +10%, IAM +20% (mecanismo: ativação simpática crônica, inflamação). Soluções: (1) Protetores auriculares (redução 20-30 dB, NRR 30), (2) Ruído branco/rosa 40-50 dB (mascara ruídos), (3) Isolamento acústico.',
  values_low = '{"min":0,"max":0,"interpretation":"Não aplicável - item descritivo."}',
  values_high = '{"min":1,"max":999,"interpretation":"Avaliar: fonte primária de ruído (trânsito: rodovia/avenida próxima, vizinhos: música/TV/discussões, animais: latidos/miados, parceiro: roncos/AOS, outros), intensidade estimada (dB: 30 sussurro, 35 biblioteca, 40 geladeira, 50 conversação, 60 TV, 70 trânsito), frequência (intermitente: mais disruptivo vs constante), impacto no sono (quantos despertares por ruído?), medidas atuais (proteção auricular? ruído branco?). Ruído >35 dB fragmenta sono significativamente. INTERVENÇÃO: protetores auriculares de silicone moldáveis (NRR 30, R$20-50), máquina de ruído branco/app (Relax Melodies, White Noise), tratar AOS do parceiro (CPAP elimina roncos), janelas acústicas se trânsito intenso (investimento alto mas ROI excelente)."}'
WHERE id = '02f8f992-b47a-407f-808e-bd95f8d4d390';

-- Odores (terceira seção - atual)
UPDATE score_items SET
  clinical_relevance = 'Aromas atuais no ambiente de sono impactam qualidade via sistema límbico (processamento emocional) e amígdala (ansiedade). Lavanda (Lavandula angustifolia): evidência RCT robusta de melhora de qualidade do sono (PSQI -2 pontos), aumento de tempo em sono profundo N3 (+20% por EEG), redução de ansiedade (STAI -15%), redução de pressão arterial (-5 mmHg), aumento de atividade parassimpática (HRV). Dose: 5-10 gotas óleo essencial em difusor ou 2-3 gotas em sachê. Outros aromas com evidência: camomila romana, bergamota (reduz cortisol), valeriana. Odores irritantes (mofo, fumaça, VOCs de limpeza) causam inflamação de vias aéreas, congestão nasal, sono fragmentado.',
  values_low = '{"min":0,"max":0,"interpretation":"Não aplicável - item descritivo."}',
  values_high = '{"min":1,"max":999,"interpretation":"Avaliar: aromas relaxantes presentes (lavanda: óleo essencial/sachê/spray de travesseiro, camomila, outros), odores irritantes (mofo/umidade: causa de rinite e asma, fumaça: tabagismo passivo, produtos de limpeza: VOCs, tintas/reformas recentes, animais de estimação no quarto), ventilação (renovação de ar), uso de difusores/óleos essenciais. Odores irritantes = eliminar PRIORITARIAMENTE (mofo: desumidificador, fumaça: proibir fumo em casa, VOCs: ventilar, animais: fora do quarto se alergia). Recomendar aromaterapia: óleo essencial de lavanda 100% puro (5-10 gotas em difusor ultrassônico 30min pré-sono, ou sachê com 2-3 gotas no travesseiro). Investimento baixo (R$30-80), sem efeitos adversos, evidência robusta."}'
WHERE id = 'fc12c639-66a4-4196-95e3-ab690caab087';

-- ============================================
-- SEÇÃO: HIGIENE DO SONO (TERCEIRA SEÇÃO - ATUAL)
-- ============================================

-- Higiene do sono (categoria - terceira seção)
UPDATE score_items SET
  clinical_relevance = 'Higiene do sono atual é conjunto de comportamentos modificáveis de MÁXIMA prioridade, com evidência robusta de eficácia. Componentes críticos: (1) Cafeína: zero após 14h (meia-vida 5-6h), (2) Refeição: leve 2-3h pré-sono, (3) Telas: ZERO 1-2h pré-sono, (4) Exercício: regular mas não >3h pré-sono, (5) Banho quente: 1-2h pré-sono (vasodilatação facilita sono), (6) Rotina pré-sono: relaxante e consistente (leitura, meditação, alongamento), (7) Horários: regulares ±30min (incluindo fins de semana). TCC-I (terapia cognitivo-comportamental para insônia) é tratamento de primeira linha: eficácia 70-80%, superior a hipnóticos (eficácia 50-60%), sem efeitos adversos, benefícios duradouros.',
  values_low = '{"min":0,"max":0,"interpretation":"Não aplicável - categoria de agrupamento."}',
  values_high = '{"min":1,"max":999,"interpretation":"Não aplicável - categoria de agrupamento."}'
WHERE id = 'b256e241-ac54-4043-82a3-0ec5087702bc';

-- Estimulantes noturnos (ex: cafeína) (terceira seção - atual)
UPDATE score_items SET
  clinical_relevance = 'Uso atual de cafeína após 14h é obstáculo modificável de ALTA prioridade. Farmacologia: cafeína antagonista competitivo de receptores de adenosina A1 (promotora de sono), meia-vida 5-6h (variável por polimorfismo CYP1A2: metabolizadores rápidos *1A/*1A 3-4h, lentos *1F/*1F 8-10h, ~50% população). Efeitos: aumenta latência do sono +10-30min, reduz tempo total de sono -30-60min, suprime sono profundo N3 -20-30%, aumenta despertares +2-4/noite. Doses típicas: café expresso 60-80mg, café coado 100-150mg, Red Bull 80mg, pré-treino 200-400mg, termogênico 300-500mg. Sensibilidade individual varia 10x (genética + tolerância).',
  values_low = '{"min":0,"max":0,"interpretation":"Não aplicável - item descritivo."}',
  values_high = '{"min":1,"max":999,"interpretation":"Avaliar DETALHADAMENTE: horário EXATO da última dose de cafeína hoje (se após 14h: interferência provável), quantidade estimada (mg: café coado 150ml = 100mg, expresso 30ml = 80mg, lata Red Bull = 80mg), fontes (café, chá preto/verde: 30-50mg, chá mate: 30mg, refrigerante cola 350ml: 35mg, chocolate >70%: 20-30mg/30g, pré-treino: 200-400mg, termogênicos: 300-500mg, medicamentos: Dorflex 65mg), sensibilidade individual (insônia após café da tarde? palpitações?), polimorfismo CYP1A2 (se conhecido). Cafeína >6h pré-sono (após 14h se sono 22h) = interferência CERTA. INTERVENÇÃO OBRIGATÓRIA: corte TOTAL após 14h (metabolizadores lentos) ou 16h (rápidos), substituir por descafeinado (resíduo 2-5mg), chás sem cafeína (camomila, erva-doce, hortelã), considerar farmacogenética CYP1A2 se insônia persistente."}'
WHERE id = '67185b74-1b99-4397-90e1-9971e7d24353';

-- Dieta noturna (terceira seção - atual)
UPDATE score_items SET
  clinical_relevance = 'Dieta noturna atual impacta qualidade do sono via múltiplos mecanismos. (1) Termogênese pós-prandial: refeição pesada <3h pré-sono aumenta temperatura corporal +0,5-1°C por 2-3h (incompatível com sono, que requer queda de 0,5-1°C). (2) RGE (refluxo): decúbito + estômago cheio → refluxo ácido → despertar, tosse noturna. (3) Desconforto: distensão abdominal, flatulência. (4) Hipoglicemia: jejum >12h → hipoglicemia noturna → despertar (mecanismo: liberação de cortisol e adrenalina). Alimentos pró-sono (triptofano → 5-HTP → serotonina → melatonina): peru, frango, ovos, nozes, banana, aveia, leite. Alimentos anti-sono: cafeína, açúcar simples (pico glicêmico → hiperinsulinemia → hipoglicemia rebound), álcool (fragmenta REM), picantes (RGE), gorduras (digestão lenta).',
  values_low = '{"min":0,"max":0,"interpretation":"Não aplicável - item descritivo."}',
  values_high = '{"min":1,"max":999,"interpretation":"Avaliar DETALHADAMENTE: horário EXATO da última refeição hoje (<3h pré-sono?), volume (refeição completa: prato cheio vs lanche leve), composição (proteínas: digestão 3-4h, carboidratos complexos: 2-3h, gorduras: 4-6h, açúcar simples: pico glicêmico), presença de sintomas noturnos (refluxo/azia, desconforto abdominal, despertar por fome), tipo de alimentos (pró-sono: triptofano vs anti-sono: cafeína/açúcar/álcool/picantes). Refeição pesada <3h pré-sono = termogênese incompatível com sono. INTERVENÇÃO: jantar 2-3h pré-sono (ex: 19h se sono 22h), lanche leve 1h pré-sono se necessário (banana + pasta de amendoim, aveia com leite, iogurte natural + nozes), priorizar fontes de triptofano (peru 350mg/100g, frango 290mg/100g, ovos 170mg/100g), evitar TOTALMENTE: cafeína, açúcar, álcool, picantes, frituras."}'
WHERE id = '98a7ddf2-4c79-44c9-8724-bd5baf7c709f';

-- Tempo tela noturna (1h antes de dormir) (terceira seção - atual)
UPDATE score_items SET
  clinical_relevance = 'Exposição atual a telas <1h pré-sono é obstáculo de MÁXIMA prioridade, mecanismo triplo: (1) Supressão de melatonina dose-dependente por luz azul 450-480nm via ipRGCs: smartphone 30cm/brilho máximo emite 40-60 lux → supressão >50% por 2h, tablet 50cm 30-40 lux → supressão 30-40%, TV 3m 10-20 lux → supressão 15-25%. (2) Estimulação cognitiva por conteúdo: trabalho (cortisol +30%), redes sociais (dopamina, comparação social, ansiedade), notícias (estresse, hipervigilância), jogos (ativação máxima). (3) Atraso de fase circadiana: cada 30min de tela <1h pré-sono atrasa sono 20-30min. Blue-blockers (óculos Uvex SCT-Orange, bloqueio >98% de 450-500nm) mitigam ~60% da supressão de melatonina mas não eliminam estimulação cognitiva.',
  values_low = '{"min":0,"max":0,"interpretation":"Não aplicável - item descritivo."}',
  values_high = '{"min":1,"max":999,"interpretation":"Avaliar CRITICAMENTE: minutos EXATOS de exposição a telas <1h pré-sono HOJE (0-15min: impacto moderado, 15-30min: significativo, >30min: severo), dispositivos específicos (smartphone: pior por proximidade 30cm, tablet: intermediário 50cm, TV: menor 3m, laptop: variável), distância dos olhos, brilho (máximo: pior, automático: intermediário, mínimo: melhor mas insuficiente), uso de blue-blockers (Uvex SCT-Orange) ou modo noturno (filtro âmbar: redução parcial ~30%), tipo de conteúdo (trabalho: cortisol máximo, redes sociais: dopamina + ansiedade, vídeos passivos: menor mas não zero, jogos: ativação máxima). >30min tela <1h pré-sono = supressão SEVERA de melatonina + atraso ~30min. INTERVENÇÃO MÁXIMA PRIORIDADE: ZERO telas 1-2h pré-sono (regra de ouro TCC-I), substituir por livro físico (luz âmbar <30 lux), alongamento, meditação, conversa. Blue-blockers são segunda linha (usar se telas inevitáveis)."}'
WHERE id = 'd492ae28-656e-4e88-8567-505644266875';

-- Campos eletromagnéticos (terceira seção - atual)
UPDATE score_items SET
  clinical_relevance = 'Exposição atual a campos eletromagnéticos (CEM) de radiofrequência (RF) 900-2.400 MHz durante o sono é tema controverso mas com evidências crescentes. Estudos experimentais (exposição controlada): RF 900 MHz (GSM) reduz sono profundo N3 -10-15%, aumenta atividade beta no EEG (incompatível com sono), aumenta latência do sono +8-15min. Mecanismo hipotético: alteração de canais de cálcio voltagem-dependentes, estresse oxidativo, interferência com melatonina. OMS classifica RF como "possivelmente carcinogênica" (Grupo 2B). Princípio da precaução: minimizar exposição desnecessária. Fontes: celular em standby 0,01-1 W/m² (sob travesseiro: máxima exposição), WiFi 0,001-0,1 W/m², Bluetooth 0,0001-0,01 W/m². Solução simples: modo avião (RF zero).',
  values_low = '{"min":0,"max":0,"interpretation":"Não aplicável - item descritivo."}',
  values_high = '{"min":1,"max":999,"interpretation":"Avaliar: celular durante sono (localização EXATA: sob travesseiro = máxima exposição 1 W/m², criado-mudo <50cm 0,1-0,5 W/m², >1m <0,01 W/m², modo avião? se não: standby com RF ativa), WiFi roteador (ligado à noite? distância do quarto? desligar com timer?), dispositivos Bluetooth (fones, smartwatch, caixas de som: Bluetooth ativo à noite?), estação base de telefone sem fio (DECT: RF contínua 1.900 MHz, desligar), baby monitor (RF ativa, minimizar distância). Celular sob travesseiro/criado-mudo SEM modo avião = exposição máxima desnecessária. INTERVENÇÃO (princípio da precaução): modo avião noturno OBRIGATÓRIO (RF zero), desligar WiFi/roteador à noite (timer 23h-6h, economia energia + redução RF), manter celular >1m da cabeça, despertador analógico ou digital com bateria (sem RF), remover smartwatch para dormir, desligar Bluetooth de todos dispositivos."}'
WHERE id = '8ec69bbf-b7c1-4eb0-bd38-31f3b1156ef6';

-- Uso atual de medicamentos/suplementos para dormir (terceira seção - atual)
UPDATE score_items SET
  clinical_relevance = 'Uso atual de hipnóticos requer avaliação CRÍTICA e plano de descontinuação gradual. Benzodiazepínicos (BZD: diazepam 5-10mg, clonazepam 0,5-2mg, alprazolam 0,25-1mg, lorazepam 1-2mg) e Z-drugs (zolpidem 5-10mg, zopiclona 7,5mg, eszopiclona 2-3mg): mecanismo GABA-A agonista. Riscos: (1) Dependência física/psicológica (3-4 semanas), (2) Tolerância (perda de eficácia), (3) Supressão de sono REM -20-30%, (4) Quedas/fraturas (+50% risco, especialmente idosos), (5) Déficit cognitivo (memória, atenção), (6) Demência: uso >3 meses aumenta risco +50% (Billioti de Gage, BMJ 2014), (7) Mortalidade: uso regular aumenta +3-5x (Kripke, BMJ Open 2012). Alternativas seguras: melatonina 0,5-5mg, magnésio 400mg, glicina 3g, L-teanina 200mg.',
  values_low = '{"min":0,"max":0,"interpretation":"Sem uso atual de hipnóticos. Menor risco de dependência, efeitos adversos e sequelas de longo prazo."}',
  values_high = '{"min":1,"max":999,"interpretation":"Uso atual de hipnóticos. Avaliar URGENTEMENTE: medicamento/suplemento ESPECÍFICO (BZD, Z-drug, antidepressivo sedativo: trazodona 50-100mg/mirtazapina 7,5-30mg/amitriptilina 25-50mg, antihistamínico: difenidramina, antipsicótico: quetiapina, suplemento: melatonina/magnésio), dose, horário de tomada (quanto tempo pré-sono?), duração de uso (<3 meses: curto prazo com risco moderado, >3 meses: uso crônico com risco ALTO de dependência e demência), eficácia REAL (melhora latência? quantos minutos? manutenção? quantas horas de sono TOTAL?), efeitos adversos (sonolência residual matinal: hangover, quedas, amnésia anterógrada, tolerância: necessidade de aumento de dose), tentativas prévias de retirada (insônia de rebote? ansiedade?). PLANO OBRIGATÓRIO: desmame GRADUAL (BZD/Z-drugs: redução 10-25% a cada 1-2 semanas, total 2-6 meses) + TCC-I simultânea (eficácia 70-80%) + substituição por suplementos seguros (melatonina 0,5-3mg: dose menor mais fisiológica, magnésio glicinato 400mg, glicina 3g, L-teanina 200mg)."}'
WHERE id = '43dae3c4-04d0-49a5-9a0f-b8d0ebb6e1a1';

-- ============================================
-- SEÇÃO: SINTOMAS NOTURNOS (TERCEIRA SEÇÃO - FINAL)
-- ============================================

-- Sintomas noturnos (categoria - terceira seção)
UPDATE score_items SET
  clinical_relevance = 'Sintomas noturnos atuais (bruxismo, pesadelos, roncos, apneias, sudorese) são marcadores de distúrbios específicos do sono que requerem avaliação diagnóstica direcionada e tratamento específico. Cada sintoma tem causa, fisiopatologia, impacto e tratamento distintos: Bruxismo → placa oclusal + TCC + polissonografia se AOS. Pesadelos → IRT + prazosina + avaliação psiquiátrica TEPT. Roncos/apneias → polissonografia URGENTE + CPAP. Sudorese → investigação hormonal/infecciosa/neoplásica. Sintomas noturnos NÃO tratados agravam qualidade de sono, saúde metabólica, cardiovascular e mental.',
  values_low = '{"min":0,"max":0,"interpretation":"Não aplicável - categoria de agrupamento."}',
  values_high = '{"min":1,"max":999,"interpretation":"Não aplicável - categoria de agrupamento."}'
WHERE id = '128659b2-fbe3-4c96-81e9-15189f22f250';

-- Bruxismo (terceira seção - final)
UPDATE score_items SET
  clinical_relevance = 'Bruxismo atual (ranger/apertar dentes durante sono) é distúrbio de movimento relacionado ao sono. Etiologia multifatorial: (1) Estresse psicológico (hiperatividade simpática noturna, correlação com cortisol elevado), (2) Ansiedade/TAG (40% dos bruxistas), (3) AOS (40-50% associação: mecanismo = despertar cortical por hipoxemia → ativação mandibular reflexa), (4) Má oclusão dentária, (5) Medicamentos (ISRS 20%, anfetaminas, cocaína). Consequências progressivas: desgaste dentário severo (exposição de dentina, sensibilidade, fraturas), DTM (dor articular, limitação de abertura bucal), cefaleia tensional matinal (tensão muscular persistente), hipertrofia de masseter (face quadrada).',
  values_low = '{"min":0,"max":0,"interpretation":"Sem bruxismo atual. Menor risco de complicações odontológicas, DTM e cefaleia tensional."}',
  values_high = '{"min":1,"max":999,"interpretation":"Bruxismo presente. Avaliar URGENTEMENTE: intensidade (audível? intensidade suficiente para despertar parceiro?), desgaste dentário (avaliação odontológica: facetas de desgaste visíveis, esmalte perdido, dentina exposta, sensibilidade dentária), dor mandibular matinal (DTM: dor ao mastigar, abertura limitada <35mm, estalos articulares), cefaleia tensional (localização: têmporas, região occipital), hipertrofia de masseter (palpável), uso atual de placa oclusal (presente? eficácia?), nível de estresse (escala 0-10), ansiedade (GAD-7), AOS (Epworth >10, roncos, apneias, obesidade). TRATAMENTO: encaminhar URGENTE para dentista (placa oclusal miorrelaxante: uso noturno obrigatório, eficácia 70-80%), polissonografia se AOS suspeitada (Epworth >10 ou roncos), TCC se estresse/ansiedade alta, toxina botulínica no masseter (25-50U bilateral) se refratário."}'
WHERE id = 'e7f8abf5-0988-4add-af05-8375aa5585a5';

-- FIM COMPLETO BATCH 41 SONO PARTE 3
