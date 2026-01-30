-- BATCH 41: SONO PARTE 3 - CONTINUAÇÃO
-- Items restantes: Equipamento, Ambiente e Higiene do Sono (Atual)

-- ============================================
-- SEÇÃO: EQUIPAMENTO DO SONO (ATUAL)
-- ============================================

-- Equipamento do sono (categoria - atual)
UPDATE score_items SET
  clinical_relevance = 'Qualidade atual do equipamento de sono é determinante modificável de qualidade do sono e dor musculoesquelética. Colchão inadequado (desgastado, firmeza inadequada) causa dor lombar/cervical e sono fragmentado. Travesseiro inadequado causa cervicalgia e cefaleia tensional. Investimento em equipamento de sono adequado tem alto ROI: melhora qualidade de sono, reduz dor crônica, aumenta produtividade. Custo-benefício superior a medicamentos.',
  values_low = '{"min":0,"max":0,"interpretation":"Não aplicável - categoria de agrupamento."}',
  values_high = '{"min":1,"max":999,"interpretation":"Não aplicável - categoria de agrupamento."}'
WHERE id = 'b991b0a9-7b65-4fd8-87b3-ed45772c794f';

-- Cama (atual)
UPDATE score_items SET
  clinical_relevance = 'Tamanho atual da cama impacta qualidade do sono, especialmente se compartilhada. Cama de casal (138cm) fornece apenas 69cm por pessoa (menos que berço infantil - 70cm). Espaço insuficiente aumenta despertares noturnos por movimentos do parceiro, eleva temperatura corporal por proximidade excessiva e reduz tempo em sono profundo. Cama Queen (158cm) ou King (193cm) melhora significativamente qualidade do sono de casais.',
  values_low = '{"min":0,"max":0,"interpretation":"Não aplicável - item descritivo."}',
  values_high = '{"min":1,"max":999,"interpretation":"Avaliar: tamanho atual (solteiro 88cm, casal 138cm, queen 158cm, king 193cm), compartilhada ou individual, adequação ao número de ocupantes, despertares por movimentos do parceiro. Cama de casal para dois adultos é inadequada. Recomendar upgrade para Queen/King se orçamento permitir. ROI alto em qualidade de sono."}'
WHERE id = 'a92458e9-f6db-4a09-bdfc-cdbac22ac227';

-- Colchão (atual)
UPDATE score_items SET
  clinical_relevance = 'Colchão atual inadequado é causa tratável de dor musculoesquelética e sono fragmentado. Colchão desgastado (>8-10 anos), com afundamento ou perda de suporte causa desalinhamento espinhal, dor lombar/cervical matinal e sono de má qualidade. Firmeza inadequada: excesso (pressão em pontos de contato) ou insuficiente (falta de suporte lombar). Colchão adequado deve manter coluna neutra em decúbito lateral. Teste: mão não deve passar sob lombar (muito firme) ou afundar >5cm (muito mole).',
  values_low = '{"min":0,"max":0,"interpretation":"Não aplicável - item descritivo."}',
  values_high = '{"min":1,"max":999,"interpretation":"Avaliar: idade do colchão (trocar a cada 8-10 anos), tipo (molas ensacadas, espuma de alta densidade, látex, viscoelástico), firmeza adequada ao IMC (maior peso = maior firmeza), presença de afundamento/deformidade, dor matinal (lombar, cervical, ombros). Colchão >10 anos requer troca. Dor lombar matinal que melhora após levantar sugere colchão inadequado. Investimento prioritário."}'
WHERE id = 'c9261b30-8c20-4a31-af68-1a702327b2e2';

-- Travesseiros (atual)
UPDATE score_items SET
  clinical_relevance = 'Travesseiro atual inadequado causa cervicalgia, cefaleia tensional matinal e sono fragmentado. Altura deve manter alinhamento cervical neutro (teste: orelha alinhada com ombro em decúbito lateral). Travesseiro muito alto: flexão cervical, tensão muscular posterior. Travesseiro muito baixo: extensão cervical, compressão facetária. Travesseiro desgastado (>1-2 anos) perde suporte e acumula ácaros (alergia). Posição de dormir determina altura ideal: lado (alto), costas (médio), bruços (sem travesseiro).',
  values_low = '{"min":0,"max":0,"interpretation":"Não aplicável - item descritivo."}',
  values_high = '{"min":1,"max":999,"interpretation":"Avaliar: número de travesseiros (empilhar = inadequado), altura, firmeza, idade (trocar a cada 1-2 anos), tipo (espuma, látex, penas, viscoelástico), posição de dormir predominante (lado 70%, costas 20%, bruços 10%), cervicalgia/cefaleia matinal. Cervicalgia matinal + melhora ao longo do dia = travesseiro inadequado. Recomendar travesseiro cervical ergonômico se dor persistente."}'
WHERE id = '3bb71150-7f78-4684-8829-6cce50554892';

-- Pijamas (atual)
UPDATE score_items SET
  clinical_relevance = 'Vestimenta noturna atual impacta termorregulação e qualidade do sono. Temperatura corporal central deve cair 0,5-1°C para iniciar sono (vasodilatação periférica). Roupas muito quentes impedem dissipação de calor, causam superaquecimento, sudorese e despertares. Tecidos sintéticos (poliéster, acrílico) retêm calor e umidade. Tecidos naturais (algodão, linho, bambú, modal) favorecem respirabilidade. Dormir sem roupa (ou roupa mínima) é ideal em climas quentes: facilita termorregulação e melhora qualidade do sono.',
  values_low = '{"min":0,"max":0,"interpretation":"Não aplicável - item descritivo."}',
  values_high = '{"min":1,"max":999,"interpretation":"Avaliar: tipo de roupa noturna (pesada, leve, sem roupa), tecido (sintético vs natural), adequação à temperatura ambiente, presença de desconforto/sudorese noturna, preferência pessoal. Roupas pesadas + ambiente quente = fragmentação do sono. Recomendar: tecidos naturais (algodão, linho), roupas leves ou sem roupa em climas quentes, camadas ajustáveis em climas frios."}'
WHERE id = '4d3e40dc-9a1b-4b1c-9332-20e800cf46ac';

-- Lençóis e cobertas (atual)
UPDATE score_items SET
  clinical_relevance = 'Roupa de cama atual impacta termorregulação, higiene e qualidade do sono. Lençóis sintéticos retêm calor e umidade, aumentando sudorese noturna e despertares. Cobertas inadequadas: excesso (superaquecimento) ou insuficiência (frio, despertares). Temperatura ideal de sono: 18-21°C (mais frio que vigília). Tecidos naturais (algodão percal/cetim, linho, bambú) melhoram respirabilidade. Higiene: lavar semanalmente (ácaros, células mortas, suor acumulado podem causar alergia/asma).',
  values_low = '{"min":0,"max":0,"interpretation":"Não aplicável - item descritivo."}',
  values_high = '{"min":1,"max":999,"interpretation":"Avaliar: tipo de tecido (natural vs sintético), peso das cobertas, adequação à temperatura ambiente (18-21°C ideal), frequência de troca (mínimo semanal), presença de sudorese noturna, sintomas alérgicos. Lençóis sintéticos + sudorese = trocar por algodão. Recomendar: lençóis de algodão alta contagem de fios (300+), coberta leve no verão, camadas ajustáveis no inverno, lavagem semanal."}'
WHERE id = '3ee86f2a-38a0-491d-abf7-6ed993c65d59';

-- ============================================
-- SEÇÃO: AMBIENTE DO SONO (ATUAL)
-- ============================================

-- Ambiente do sono (categoria - atual)
UPDATE score_items SET
  clinical_relevance = 'Ambiente atual do sono inadequado é causa modificável e de alta prioridade para intervenção. Fatores críticos: escuridão total <1 lux (luz suprime melatonina dose-dependente), temperatura 18-21°C (frio facilita sono), silêncio <30 dB ou ruído branco, ausência de telas 1-2h pré-sono (luz azul suprime melatonina >50%), aromas relaxantes (lavanda). Otimização do ambiente é intervenção de primeira linha, custo-benefício superior a medicamentos. Componente obrigatório da TCC-I.',
  values_low = '{"min":0,"max":0,"interpretation":"Não aplicável - categoria de agrupamento."}',
  values_high = '{"min":1,"max":999,"interpretation":"Não aplicável - categoria de agrupamento."}'
WHERE id = '74dad4e7-7184-4bea-886c-ed632a8e7d23';

-- Quarto (atual)
UPDATE score_items SET
  clinical_relevance = 'Características atuais do quarto determinam qualidade do sono. Tamanho adequado (mínimo 9m² para casal), ventilação suficiente (CO2 >1.000 ppm causa cefaleia matinal, déficit cognitivo), organização (desordem visual aumenta cortisol e ansiedade), uso exclusivo para sono/sexo (cama usada para trabalho cria associação incompatível com sono - componente da TCC-I). Quarto deve ser santuário do sono: escuro, fresco, silencioso, organizado, livre de trabalho.',
  values_low = '{"min":0,"max":0,"interpretation":"Não aplicável - item descritivo."}',
  values_high = '{"min":1,"max":999,"interpretation":"Avaliar: tamanho, ventilação (janelas, ventilador, AC), organização (desordem visual?), uso múltiplo (home office no quarto?), sensação de conforto/segurança. Quarto desorganizado ou usado para trabalho prejudica sono (associação mental inadequada). Recomendar: uso exclusivo para sono/sexo, ventilação adequada (janela aberta ou ventilador), organização, remover equipamentos de trabalho."}'
WHERE id = '22a2a975-6b17-463a-ad12-bf7cfe936984';

-- Iluminação (atual)
UPDATE score_items SET
  clinical_relevance = 'Iluminação noturna atual é determinante crítico de secreção de melatonina e qualidade do sono. Luz >10 lux à noite suprime melatonina de forma dose-dependente: 10-30 lux suprime 20%, >100 lux suprime >50%. Luz azul (450-480nm) de LEDs e telas tem máximo impacto. Escuridão total <1 lux é ideal. Fontes comuns: luz de rua (janelas sem blackout), LEDs de eletrônicos (despertador digital, TV, roteador), luz de corredores. Solução: cortinas blackout + máscara de dormir + remover eletrônicos.',
  values_low = '{"min":0,"max":0,"interpretation":"Não aplicável - item descritivo."}',
  values_high = '{"min":1,"max":999,"interpretation":"Avaliar: escuridão total (<1 lux) ou presença de luz (janela, luz de rua, eletrônicos), uso de cortinas blackout ou máscara de dormir, luz de dispositivos (despertador digital, TV, roteador, carregadores com LED). Luz visível no quarto = supressão de melatonina. Recomendar ALTA PRIORIDADE: cortinas blackout, máscara de dormir, remover/cobrir LEDs, despertador analógico ou com luz vermelha."}'
WHERE id = '0f87c64b-dbb0-4eec-9a64-9c5e0bb9b123';

-- Janelas / claridade (atual)
UPDATE score_items SET
  clinical_relevance = 'Janelas no quarto têm papel dual: permitem entrada de luz natural matinal (sincroniza ritmo circadiano, >1.000 lux) mas também entrada de luz noturna (suprime melatonina). Janela face leste: sol matinal direto (ideal para acordar naturalmente, avanço de fase). Janela face oeste: sol poente (pode atrasar sono). Luz de rua à noite requer cortinas blackout (bloqueio >99%). Exposição matinal à luz natural >30min reduz latência do sono em 1h à noite. Janela + blackout = melhor combinação.',
  values_low = '{"min":0,"max":0,"interpretation":"Não aplicável - item descritivo."}',
  values_high = '{"min":1,"max":999,"interpretation":"Avaliar: presença de janelas, exposição à luz natural matinal (face leste ideal), cortinas blackout (bloqueio de luz noturna), luz de rua/postes à noite, posição geográfica do quarto. Janela SEM blackout + luz de rua = supressão de melatonina. Recomendar: cortinas blackout (bloqueio >99%), exposição matinal à luz natural >30min (abrir cortinas ao acordar), face leste preferível."}'
WHERE id = '075e0074-16bc-4db3-8f0b-93594686032e';

-- Televisão / telas (atual)
UPDATE score_items SET
  clinical_relevance = 'TV e dispositivos eletrônicos no quarto atualmente são obstáculos principais à higiene do sono. Mecanismos: (1) luz azul 450-480nm suprime melatonina >50% mesmo em baixa intensidade, (2) conteúdo estimulante aumenta cortisol e alerta cognitivo, (3) associação mental de quarto com vigília (não sono). Uso de telas <1h pré-sono atrasa início do sono ~1h. TV ligada até adormecer = pior cenário. Regra de ouro TCC-I: quarto livre de telas, celular em modo avião FORA da cama.',
  values_low = '{"min":0,"max":0,"interpretation":"Não aplicável - item descritivo."}',
  values_high = '{"min":1,"max":999,"interpretation":"Avaliar: TV no quarto (ligada até adormecer?), celular na cama/sob travesseiro, uso de telas <1h pré-sono, brilho das telas, tipo de conteúdo (trabalho/redes sociais/vídeos). TV/celular até adormecer = MÁXIMA supressão de melatonina. Recomendar ALTA PRIORIDADE: remover TV do quarto, celular em modo avião fora da cama, substituir por leitura (livro físico), blue-blockers se necessário usar telas (segunda linha)."}'
WHERE id = 'f5d21770-c844-4b1a-9a25-42492ab7e3b5';

-- Barulho (atual)
UPDATE score_items SET
  clinical_relevance = 'Ruído noturno atual >35 dB (conversação baixa) fragmenta arquitetura do sono, reduz tempo em sono profundo (N3), aumenta despertares e eleva cortisol noturno. Ruído intermitente (trânsito, vizinhos, latidos) é mais disruptivo que ruído constante (ventilador). Ruído crônico noturno >55 dB aumenta risco cardiovascular (hipertensão, IAM) por ativação simpática persistente. Soluções: protetores auriculares (redução 20-30 dB), ruído branco/rosa (40-50 dB, mascara ruídos externos), isolamento acústico.',
  values_low = '{"min":0,"max":0,"interpretation":"Não aplicável - item descritivo."}',
  values_high = '{"min":1,"max":999,"interpretation":"Avaliar: fonte de ruído atual (trânsito, vizinhos, animais, parceiro roncando), intensidade estimada (dB), frequência (intermitente/constante), impacto no sono (despertares), uso de proteção (tampões auriculares, fones noise-canceling, ruído branco). Ruído >35 dB fragmenta sono. Recomendar: protetores auriculares de silicone (NRR 30), máquina de ruído branco, tratamento de AOS do parceiro (CPAP), mudança de quarto se possível."}'
WHERE id = '2b7e6b36-5578-4c46-901c-16e38005e8d9';

-- Odores (atual)
UPDATE score_items SET
  clinical_relevance = 'Aromas atuais no ambiente do sono impactam qualidade via sistema límbico e amígdala. Lavanda (Lavandula angustifolia): evidência RCT de melhora de qualidade do sono (+20% tempo em sono profundo), redução de ansiedade, redução de pressão arterial. Outros aromas com evidência: camomila romana, bergamota, valeriana. Odores irritantes (mofo, produtos de limpeza, fumaça): inflamação de vias aéreas, congestão nasal, sono fragmentado. Aromaterapia é adjuvante seguro, custo-benefício alto, sem efeitos adversos.',
  values_low = '{"min":0,"max":0,"interpretation":"Não aplicável - item descritivo."}',
  values_high = '{"min":1,"max":999,"interpretation":"Avaliar: presença de aromas relaxantes (lavanda, camomila), odores irritantes (mofo, fumaça, produtos químicos fortes), ventilação, uso de difusores/óleos essenciais/sachês aromáticos. Odores irritantes pioram qualidade do sono. Recomendar: óleo essencial de lavanda (5-10 gotas em difusor ou 2-3 gotas em sachê no travesseiro), eliminar fontes de odores irritantes, ventilação adequada (renovação de ar)."}'
WHERE id = 'b01110cb-1fb2-4a5e-a9b0-c52b197a46d1';

-- ============================================
-- SEÇÃO: HIGIENE DO SONO (ATUAL)
-- ============================================

-- Higiene do sono (categoria - atual)
UPDATE score_items SET
  clinical_relevance = 'Higiene do sono atual inadequada é causa primária e modificável de insônia e sono de má qualidade. Componentes críticos de intervenção: (1) evitar cafeína >6h pré-sono, (2) refeição leve >3h pré-sono, (3) zero telas 1-2h pré-sono, (4) exercício regular (não >3h pré-sono), (5) banho quente 1-2h pré-sono (vasodilatação pós-banho facilita sono), (6) rotina pré-sono relaxante, (7) horários regulares. TCC-I (terapia cognitivo-comportamental para insônia) é tratamento de primeira linha, superior a hipnóticos.',
  values_low = '{"min":0,"max":0,"interpretation":"Não aplicável - categoria de agrupamento."}',
  values_high = '{"min":1,"max":999,"interpretation":"Não aplicável - categoria de agrupamento."}'
WHERE id = 'df8df017-2f55-443a-8f14-683006b14adb';

-- Estimulantes noturnos (ex: cafeína) (atual)
UPDATE score_items SET
  clinical_relevance = 'Uso atual de cafeína após 14h interfere no sono das 22h (meia-vida 5-6h, variável por CYP1A2). Cafeína bloqueia receptores de adenosina (promotora de sono), aumenta latência do sono, reduz tempo total de sono, suprime sono profundo (N3) e aumenta despertares noturnos. Metabolizadores lentos (CYP1A2*1F, ~50% da população) têm meia-vida >8h. Dose típica: café expresso 80mg, café coado 150mg, Red Bull 80mg, pré-treino 200-400mg. Sensibilidade individual varia 10x.',
  values_low = '{"min":0,"max":0,"interpretation":"Não aplicável - item descritivo."}',
  values_high = '{"min":1,"max":999,"interpretation":"Avaliar: horário última dose de cafeína hoje, quantidade estimada (mg), fontes (café, chá preto/verde, refrigerante cola, chocolate >70%, pré-treinos, termogênicos), sensibilidade individual (insônia após café da tarde?). Cafeína >6h pré-sono interfere. Recomendar: corte TOTAL após 14h (ou 12h se metabolizador lento), substituir por descafeinado/chás sem cafeína (camomila, erva-doce), considerar teste farmacogenético CYP1A2 se insônia persistente."}'
WHERE id = '25aa750f-d458-4e39-b37c-b5a66502340a';

-- Dieta noturna (atual)
UPDATE score_items SET
  clinical_relevance = 'Dieta noturna atual impacta diretamente qualidade do sono. Refeição pesada <3h pré-sono causa: (1) termogênese pós-prandial (elevação de temperatura corporal incompatível com sono), (2) refluxo gastroesofágico (RGE) em decúbito, (3) desconforto abdominal. Jejum excessivo (>12h) pode causar hipoglicemia noturna e despertar. Alimentos pró-sono (ricos em triptofano → serotonina → melatonina): peru, frango, ovos, nozes, banana, aveia. Alimentos anti-sono: cafeína, açúcar simples (pico glicêmico), álcool (fragmenta REM), alimentos picantes (RGE).',
  values_low = '{"min":0,"max":0,"interpretation":"Não aplicável - item descritivo."}',
  values_high = '{"min":1,"max":999,"interpretation":"Avaliar: horário da última refeição hoje, volume (refeição completa vs lanche), composição (proteínas, carboidratos complexos, gorduras), presença de refluxo/azia/desconforto noturno, despertar por fome. Refeição pesada <3h pré-sono fragmenta sono. Recomendar: jantar 2-3h pré-sono, lanche leve 1h pré-sono se necessário (banana + nozes, aveia), fontes de triptofano, evitar: cafeína, açúcar, picantes, gorduras."}'
WHERE id = '2ce2bbb8-31d7-44e1-88b3-d7282aeb79a7';

-- Tempo tela noturna (1h antes de dormir) (atual)
UPDATE score_items SET
  clinical_relevance = 'Exposição atual a telas <1h pré-sono é fator modificável de alta prioridade. Mecanismos: (1) luz azul 450-480nm suprime melatonina >50% por 2h (dose-dependente: smartphone 30cm > tablet 50cm > TV 3m), (2) estimulação cognitiva por conteúdo ativador (trabalho, redes sociais, notícias), (3) hipervigilância e aumento de cortisol. Cada 30min de tela <1h pré-sono atrasa sono ~30min. Blue-blockers (óculos bloqueadores de luz azul) mitigam ~50% do impacto. Regra de ouro TCC-I: ZERO telas 1-2h pré-sono.',
  values_low = '{"min":0,"max":0,"interpretation":"Não aplicável - item descritivo."}',
  values_high = '{"min":1,"max":999,"interpretation":"Avaliar: minutos de exposição a telas <1h pré-sono HOJE, dispositivos (smartphone, tablet, TV, laptop), distância, brilho (máximo/médio/baixo), uso de blue-blockers ou modo noturno (filtro âmbar), tipo de conteúdo (trabalho/redes sociais/vídeos/jogos). >30min de tela <1h pré-sono = supressão significativa de melatonina. Recomendar ALTA PRIORIDADE: ZERO telas 1-2h pré-sono, substituir por livro físico, blue-blockers (Uvex SCT-Orange) se necessário, modo noturno (segunda linha)."}'
WHERE id = '498aada9-e14e-431c-a13a-4bc1498abb85';

-- Campos eletromagnéticos (atual)
UPDATE score_items SET
  clinical_relevance = 'Exposição atual a campos eletromagnéticos (CEM) de radiofrequência durante o sono é tema controverso mas com evidências crescentes. Estudos mostram: redução de sono profundo (N3), aumento de atividade beta no EEG (incompatível com sono), aumento de latência do sono com exposição a RF 900MHz (celular). OMS classifica RF como "possivelmente carcinogênica" (2B). Princípio da precaução: minimizar exposição. Soluções simples: modo avião, desligar WiFi/Bluetooth, manter dispositivos >1m da cabeça. CEM de baixa frequência (60Hz, fiação) tem menor evidência.',
  values_low = '{"min":0,"max":0,"interpretation":"Não aplicável - item descritivo."}',
  values_high = '{"min":1,"max":999,"interpretation":"Avaliar: celular próximo à cabeça durante sono (sob travesseiro/criado-mudo, em modo avião?), WiFi ligado à noite, dispositivos Bluetooth (fones, smartwatch), distância de dispositivos eletrônicos da cama. Celular sob travesseiro = máxima exposição. Recomendar: modo avião noturno OBRIGATÓRIO, desligar WiFi/roteador (timer), manter celular >1m da cama, despertador analógico ou com bateria, remover smartwatch para dormir."}'
WHERE id = '073e9b8b-f3b8-411d-996c-e4b6ae6b71e1';

-- Uso atual de medicamentos/suplementos para dormir (atual)
UPDATE score_items SET
  clinical_relevance = 'Uso atual de hipnóticos requer avaliação crítica de indicação, eficácia, segurança e plano de descontinuação. Benzodiazepínicos (diazepam 5-10mg, clonazepam 0,5-2mg, alprazolam 0,25-1mg) e Z-drugs (zolpidem 5-10mg, zopiclona 7,5mg): risco de dependência, tolerância, quedas, fraturas, déficit cognitivo, supressão de sono REM, aumento de mortalidade. Uso >3 meses aumenta risco de demência em 50%. Alternativas: melatonina 0,5-5mg (dose menor mais fisiológica), magnésio 400mg (glicinato), glicina 3g, L-teanina 200mg, valeriana 400-600mg.',
  values_low = '{"min":0,"max":0,"interpretation":"Sem uso atual de hipnóticos. Menor risco de dependência e efeitos adversos."}',
  values_high = '{"min":1,"max":999,"interpretation":"Uso atual de hipnóticos. Avaliar: medicamento/suplemento específico, dose, horário de tomada, duração de uso (<3 meses: curto prazo; >3 meses: uso crônico), eficácia (melhora latência/manutenção? quantas horas de sono?), efeitos adversos (sonolência residual, quedas, amnésia anterógrada, tolerância), tentativas prévias de retirada (insônia de rebote?). Planejar DESMAME GRADUAL (redução 10-25%/mês) + TCC-I. Substituir por: melatonina 0,5-3mg, magnésio 400mg, glicina 3g."}'
WHERE id = '0a087bbe-fb67-4969-bcad-feda46b447e4';

-- ============================================
-- SEÇÃO: SINTOMAS NOTURNOS (ATUAL - FINAL)
-- ============================================

-- Sintomas noturnos (categoria - atual - duplicata)
UPDATE score_items SET
  clinical_relevance = 'Sintomas noturnos atuais indicam distúrbios específicos do sono que requerem avaliação diagnóstica e tratamento direcionado. Bruxismo: placa oclusal + TCC + avaliação de AOS. Pesadelos: IRT + prazosina + avaliação psiquiátrica. Roncos/apneias: polissonografia URGENTE + CPAP se IAH >15. Sudorese: investigar causas hormonais/infecciosas/neoplásicas. Parceiro com AOS: tratamento melhora sono de ambos. Sintomas noturnos impactam qualidade de vida, saúde metabólica e cardiovascular.',
  values_low = '{"min":0,"max":0,"interpretation":"Não aplicável - categoria de agrupamento."}',
  values_high = '{"min":1,"max":999,"interpretation":"Não aplicável - categoria de agrupamento."}'
WHERE id = '725810f4-ddff-4fac-aede-00e0d1a12bc6';

-- Bruxismo (atual - duplicata)
UPDATE score_items SET
  clinical_relevance = 'Bruxismo atual indica hiperatividade mastigatória noturna persistente. Etiologia multifatorial: estresse psicológico (hiperatividade simpática), ansiedade, AOS (40% dos bruxistas têm AOS, mecanismo: despertar cortical + ativação mandibular), má oclusão, medicamentos (ISRS, anfetaminas). Consequências: desgaste dentário severo (exposição de dentina, fraturas), disfunção temporomandibular (DTM), cefaleia tensional matinal, hipertrofia do masseter. Tratamento: placa oclusal (primeira linha), TCC, polissonografia se AOS, toxina botulínica (casos refratários).',
  values_low = '{"min":0,"max":0,"interpretation":"Sem bruxismo atual. Menor risco de complicações odontológicas e DTM."}',
  values_high = '{"min":1,"max":999,"interpretation":"Bruxismo presente. Avaliar urgentemente: intensidade (audível fora do quarto? testemunhas?), desgaste dentário visível (facetas de desgaste, esmalte perdido), dor mandibular matinal (DTM), cefaleia tensional, limitação de abertura bucal, uso de placa oclusal, nível de estresse/ansiedade, AOS (Epworth >10, roncos, apneias). Encaminhar para dentista (placa oclusal urgente), polissonografia se AOS, TCC se estresse alto. Toxina botulínica no masseter se refratário."}'
WHERE id = 'efe2b4bb-8980-420d-aa75-48508ea4d25d';

-- Pesadelos (atual - duplicata)
UPDATE score_items SET
  clinical_relevance = 'Pesadelos frequentes atuais (>1/semana) caracterizam transtorno de pesadelos (DSM-5), associado a TEPT (50-70% dos casos), transtornos ansiosos, depressão maior, ou iatrogenia medicamentosa (betabloqueadores, antidepressivos ISRS/IRSN, pramipexol, donepezil, vareniclina). Pesadelos fragmentam sono REM, causam insônia de manutenção (medo de dormir), agravam sintomas de TEPT (retraumatização noturna). Tratamento: IRT (Imagery Rehearsal Therapy, eficácia 70%), prazosina 1-6mg (alfa-1-bloqueador, reduz pesadelos em 50%), TCC para TEPT.',
  values_low = '{"min":0,"max":0,"interpretation":"Sem pesadelos frequentes. Sono REM preservado."}',
  values_high = '{"min":1,"max":999,"interpretation":"Pesadelos frequentes. Avaliar: frequência (quantos/semana), conteúdo temático (traumas recorrentes específicos? TEPT provável), impacto no sono (insônia de manutenção, medo de dormir, fadiga diurna), diagnósticos psiquiátricos (TEPT, ansiedade, depressão), medicamentos causais (betabloqueadores, ISRS, dopaminérgicos). Solicitar avaliação psiquiátrica. Iniciar IRT (Imagery Rehearsal Therapy), considerar prazosina 1mg (escalar até 6mg), suspender/trocar medicamentos causais se possível."}'
WHERE id = 'ba3a8842-9d98-433c-b0a0-3d04c6c51bc9';

-- Roncos (atual - duplicata)
UPDATE score_items SET
  clinical_relevance = 'Roncos habituais atuais (>3 noites/semana) são marcador de resistência aumentada das vias aéreas superiores e preditor independente de AOS. Roncos primários (sem apneias) ainda estão associados a hipertensão, espessamento de carótidas, disfunção endotelial e risco cardiovascular aumentado. Roncos progressivos (piora com envelhecimento/ganho de peso) ou associados a sonolência diurna (Epworth >10) indicam AOS provável. Polissonografia está indicada se: roncos + Epworth >10, ou roncos + comorbidades (HAS, DM2, FA), ou roncos + apneias testemunhadas.',
  values_low = '{"min":0,"max":0,"interpretation":"Sem roncos habituais. Menor risco de AOS."}',
  values_high = '{"min":1,"max":999,"interpretation":"Roncos habituais presentes. Avaliar: frequência (noites/semana), intensidade (audível fora do quarto?), pausas respiratórias/engasgos testemunhados, sonolência diurna (Epworth), obesidade (IMC, circunferência cervical >40cm H />37cm M), HAS, DM2. Solicitar polissonografia se: Epworth >10, ou apneias testemunhadas, ou comorbidades cardiovasculares/metabólicas. Tratamento: perda de peso (reduz IAH 30-50%), CPAP se IAH >15, aparelho intraoral se IAH 5-15, evitar álcool/sedativos noturnos."}'
WHERE id = 'f8c9dfea-cfed-41c5-a99e-965389158850';

-- Apneias (atual - duplicata)
UPDATE score_items SET
  clinical_relevance = 'Apneias testemunhadas atuais (pausas respiratórias >10s durante sono) têm VPP >90% para AOS. AOS não tratada causa: (1) hipoxemia intermitente (dessaturação <90%), (2) fragmentação do sono (microdespertares), (3) ativação simpática noturna persistente, (4) inflamação sistêmica (PCR, IL-6), (5) estresse oxidativo. Consequências: hipertensão resistente (50% dos hipertensos refratários têm AOS), resistência insulínica e DM2, síndrome metabólica, fibrilação atrial, AVC, mortalidade cardiovascular 3x aumentada. Tratamento com CPAP reduz eventos CV em 30-50%.',
  values_low = '{"min":0,"max":0,"interpretation":"Sem apneias testemunhadas. Menor probabilidade de AOS."}',
  values_high = '{"min":1,"max":999,"interpretation":"Apneias presentes. AOS ALTAMENTE PROVÁVEL. Solicitar polissonografia URGENTE (prioridade se comorbidades CV). Avaliar: duração das pausas (>10s?), frequência estimada (quantas/noite?), engasgos, sonolência diurna (Epworth), cefaleia matinal, HAS (especialmente resistente), obesidade (IMC, circunferência cervical >40cm H />37cm M), DM2, FA. Tratamento: CPAP (primeira linha se IAH >15 ou >5 com sintomas), perda de peso (reduz IAH 30-50% se perda >10kg), aparelho intraoral se IAH leve-moderado, cirurgia (UPPP, avanço maxilomandibular) em casos selecionados."}'
WHERE id = '6fc90b17-bd93-4d30-b0ec-13817d2e1756';

-- Sudorese noturna (atual - duplicata final)
UPDATE score_items SET
  clinical_relevance = 'Sudorese noturna atual que exige troca de roupa/lençóis é sintoma B (sintoma constitucional) que exige investigação urgente de causas potencialmente graves. Causas principais: (1) Hormonais: menopausa/perimenopausa (fogachos, deficiência estrogênica), hipogonadismo masculino (testosterona <300 ng/dL), (2) Metabólicas: hipoglicemia noturna (DM, insulinoma, sulfoniluréias, jejum), hipertireoidismo, (3) Infecciosas: tuberculose (sudorese + febre + perda de peso), endocardite, HIV, (4) Neoplásicas: linfoma (Hodgkin/não-Hodgkin), leucemia, (5) Outros: AOS grave, DRGE, medicamentos (antidepressivos, tamoxifeno).',
  values_low = '{"min":0,"max":0,"interpretation":"Sem sudorese noturna. Termorregulação autonômica preservada."}',
  values_high = '{"min":1,"max":999,"interpretation":"Sudorese noturna presente. Investigar URGENTEMENTE: intensidade (troca de roupa/lençóis?), frequência, sintomas B (febre >38°C, perda de peso >10% em 6 meses, fadiga intensa), sintomas hormonais (fogachos diurnos, libido reduzida), hipoglicemia noturna (DM). Solicitar URGENTE: hemograma completo + diferencial, VHS, PCR, glicemia de madrugada (3-4h), TSH + T4L, testosterona total/estradiol, RX tórax (descartar TB, linfoma). TC tórax/abdome/pelve se sinais de alarme (febre + perda de peso + sudorese). PPD/Quantiferon se risco TB."}'
WHERE id = 'b35b777e-929a-497e-bcc9-652bb329d315';

-- Hábitos do cônjuge/parceiro de quarto (atual - duplicata)
UPDATE score_items SET
  clinical_relevance = 'Hábitos atuais do parceiro de quarto impactam significativamente qualidade do sono do paciente. Principais interferências: (1) Roncos/apneias do parceiro (ruído intermitente fragmenta sono, média 4-6 despertares/noite), (2) Movimentos periódicos de pernas (síndrome das pernas inquietas, PLMD), (3) Cronótipos incompatíveis (matutino vs vespertino, >2h diferença), (4) Uso de telas na cama (luz azul suprime melatonina do parceiro). Parceiro com AOS não tratada reduz qualidade do sono em 40-60% e aumenta risco de insônia, ansiedade, divórcio. Tratamento de AOS do parceiro melhora sono de ambos.',
  values_low = '{"min":0,"max":0,"interpretation":"Não aplicável - item descritivo."}',
  values_high = '{"min":1,"max":999,"interpretation":"Avaliar detalhadamente: parceiro ronca (frequência, intensidade, AOS provável?), movimentos noturnos que despertam paciente, horários de sono incompatíveis (diferença >2h?), uso de telas na cama pelo parceiro, preferências diferentes de temperatura/cobertas. Roncos do parceiro = encaminhar para polissonografia (tratamento de AOS melhora sono de ambos). Considerar quartos separados temporariamente se interferência severa e AOS não tratável imediatamente. Avaliar síndrome das pernas inquietas do parceiro (ferritina <75)."}'
WHERE id = 'b15a28a0-707d-4684-aed1-d3c1fcd61726';

-- ============================================
-- SEÇÃO: DESPERTARES NOTURNOS (ATUAL)
-- ============================================

-- Despertares noturnos (categoria - atual - duplicata)
UPDATE score_items SET
  clinical_relevance = 'Despertares noturnos frequentes atuais (>2/noite) fragmentam arquitetura do sono, reduzem tempo em sono profundo (N3, essencial para restauração metabólica, secreção de GH) e sono REM (consolidação de memória emocional, regulação de humor), prejudicam consolidação de memória e aumentam fadiga diurna, sonolência, déficit cognitivo. Causas principais: noctúria (HBP, DM descompensado, AOS, diuréticos noturnos, poliúria noturna), AOS (microdespertares por hipoxemia), dor crônica (artrite, fibromialgia), ansiedade (insônia de manutenção, ruminação), síndrome das pernas inquietas, medicamentos.',
  values_low = '{"min":0,"max":0,"interpretation":"Não aplicável - categoria de agrupamento."}',
  values_high = '{"min":1,"max":999,"interpretation":"Não aplicável - categoria de agrupamento."}'
WHERE id = '428052bb-0cbd-4ab2-977c-a6a39cedc3fa';

-- Número (de despertares noturnos - atual - duplicata)
UPDATE score_items SET
  clinical_relevance = 'Número atual de despertares noturnos quantifica objetivamente fragmentação do sono e prediz impacto diurno. >2 despertares/noite está associado a maior sonolência diurna (Epworth >10), déficit cognitivo (memória de trabalho, atenção), pior qualidade de vida, maior risco de acidentes, resistência insulínica e obesidade. Despertares frequentes (>3/noite) reduzem tempo em sono profundo (N3) em 30-50% e sono REM em 20-40%. Causas mais comuns: noctúria (60%), AOS (40%), dor crônica (30%), ansiedade (25%).',
  values_low = '{"min":0,"max":1,"interpretation":"Sono consolidado com mínimos despertares (0-1/noite). Arquitetura do sono provavelmente preservada, sono profundo e REM adequados."}',
  values_high = '{"min":3,"max":20,"interpretation":"Despertares frequentes (≥3/noite). Sono severamente fragmentado. Investigar URGENTEMENTE causas: noctúria (volume urinado, frequência, avaliar HBP, DM, uso de diuréticos, AOS), apneias/roncos (Epworth, polissonografia), dor crônica (localização, intensidade, analgesia inadequada), ansiedade (ruminação noturna, preocupações, GAD-7), síndrome das pernas inquietas (ferritina <75), medicamentos. Quantificar impacto diurno: Epworth, fadiga, déficit cognitivo. Tratamento direcionado à causa."}'
WHERE id = 'b718b0b6-283d-4051-9e06-12b19e8ed296';

-- Motivos (dos despertares noturnos - atual - duplicata)
UPDATE score_items SET
  clinical_relevance = 'Identificação específica dos motivos atuais de despertares noturnos permite intervenção direcionada e eficaz. Principais motivos e intervenções: (1) Noctúria: avaliar HBP (finasterida, tansulosina), DM (controle glicêmico), AOS (CPAP aumenta peptídeo natriurético → poliúria), diuréticos (mudar para manhã), (2) Apneias: polissonografia + CPAP, (3) Dor: otimizar analgesia noturna (analgésico 1h pré-sono, dose extra se despertar), (4) Ansiedade: TCC-I, higiene do sono, meditação mindfulness, (5) Movimentos de pernas: ferritina >75 (suplementar ferro se <75), pramipexol 0,125-0,5mg.',
  values_low = '{"min":0,"max":0,"interpretation":"Não aplicável - item descritivo."}',
  values_high = '{"min":1,"max":999,"interpretation":"Investigar detalhadamente CADA motivo de despertar: (1) Noctúria: frequência (quantas vezes/noite), volume urinado (>500mL/noite = poliúria), sintomas de HBP (jato fraco, gotejamento), DM descompensado (glicemia de madrugada), uso de diuréticos (horário), (2) Apneias/engasgos: Epworth, roncos, obesidade, (3) Dor: localização precisa, EVA, analgesia atual inadequada, (4) Ansiedade: ruminação (conteúdo das preocupações), GAD-7, (5) Movimentos de pernas: sensação de formigamento/queimação, necessidade de mover pernas, (6) Sede/boca seca: DM, Sjögren, respiração oral (AOS), (7) Outros. Intervenção direcionada à causa específica."}'
WHERE id = '53066b08-838c-4edb-9cfa-776be8461071';

-- FIM BATCH 41 SONO PARTE 3
