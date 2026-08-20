-- +goose Up
-- Correções de nomenclatura e de convenção de dose no catálogo magistral.
--
-- Quatro correções que nasceram como seed em dev e precisam chegar em produção, onde a carga da
-- 00081 gravou os nomes crus do formulário das parceiras:
--
--   1. MINERAL DIZ A MOLÉCULA. O formulário escreve "COBRE QUELATO" e o mapa de nomes colapsou
--      para "Cobre". Receita magistral não se escreve assim: a farmácia precisa saber qual sal ou
--      quelato pesar. Junto vem a marca de que a dose é do ELEMENTO, que é como mineral se
--      prescreve — sem ela, "cobre quelato 1 mg" seria lido como 1 mg do pó.
--   2. COMPONENTE ÓRFÃO. 20 substâncias estavam gravadas com o texto cru ("Gimnema silvestre",
--      "Cromo (GTF ou picolinato)", "Weg lem 70(ganoderma lucidum)") e não casavam com nenhuma
--      entrada do catálogo — perdiam faixa de dose, densidade, teto da IN 28, interferência em
--      exame e a marca de elementar. Todas existiam sob o nome canônico.
--   3. NOTA DE RODAPÉ NO NOME DO NUTRIENTE. O Anexo IV numera as notas em romano e a numeração
--      veio grudada: "Cálciov", "Vitamina Dii", "Niacina ix". O nome é a chave que liga o
--      componente ao teto, então os dois lados mudam juntos.
--   4. NOTA REDUNDANTE NO RECEITUÁRIO. Com o PDF imprimindo "(do elemento)" ao lado da
--      quantidade, a observação "Dose em cobre elementar." repetia a informação no meio da linha.
--
-- Tudo em UPDATE condicional: rodar de novo não faz efeito. Down não desfaz — reverter seria
-- reintroduzir nome errado em receita já emitida.


-- ==============================================================================================
-- fonte: docs/emr/magistral-minerais-forma-e-elemento.sql
-- ==============================================================================================
-- Minerais: o nome diz a MOLÉCULA e a dose diz que é do ELEMENTO.
--
-- ERRO QUE ISTO CORRIGE: o formulário das parceiras escreve "COBRE QUELATO", "MANGANÊS QUELATO",
-- "BORO QUELATO", e o mapa de nomes colapsou tudo para o elemento puro — "Cobre". Prescrição
-- magistral não se escreve assim: a farmácia precisa saber qual sal ou quelato pesar, e a dose
-- precisa dizer se é do elemento ou do insumo.
--
-- O percentual elementar fica NULO onde depende do fornecedor: quelato de cobre de um fabricante
-- não tem a mesma concentração do de outro, e inventar número aqui seria pior que não ter. Fica a
-- observação pedindo a ficha técnica. Onde a estequiometria resolve (iodeto de potássio) ou o
-- insumo é padronizado de mercado (bisglicinato ferroso), o número entra.


-- ---------------------------------------------------------------------------------------------
-- 1. Selênio some: a forma que o prescritor usa é a selenometionina
-- ---------------------------------------------------------------------------------------------
UPDATE magistral_formula_template_components c
   SET substance = 'Selenometionina', as_elemental = true,
       note = CASE WHEN coalesce(note,'') = '' THEN 'Dose em selênio elementar.' ELSE note END
  FROM magistral_formula_templates t
 WHERE t.id = c.template_id AND c.substance = 'Selênio';

-- Soft delete, não DELETE: `prescription_formula_components.magistral_component_id` tem FK para
-- cá, então apagar a linha derruba a migration inteira no dia em que uma receita já emitida
-- apontar para ela — e o deploy junto. Todo o código filtra `deleted_at IS NULL`, então some da
-- busca do mesmo jeito, sem quebrar receita antiga.
UPDATE public.magistral_components SET deleted_at = now()
 WHERE name = 'Selênio' AND deleted_at IS NULL
   AND NOT EXISTS (SELECT 1 FROM public.magistral_formula_template_components c
                    WHERE c.substance = 'Selênio' AND c.deleted_at IS NULL);

-- ---------------------------------------------------------------------------------------------
-- 2. O nome passa a dizer a molécula
-- ---------------------------------------------------------------------------------------------
UPDATE magistral_components SET name = 'Cobre quelato',
       synonyms = 'cobre, COBRE QUELATO, COBRE QUELADO, bisglicinato de cobre',
       correction_note = 'Percentual de cobre no quelato varia com o fornecedor: confirmar na ficha técnica antes de converter para massa de insumo.'
 WHERE name = 'Cobre';

UPDATE magistral_components SET name = 'Cálcio quelato',
       synonyms = 'cálcio, CALCIO QUELADO, CÁLCIO QUELATO, bisglicinato de cálcio',
       correction_note = 'Percentual de cálcio varia com o quelato: confirmar na ficha técnica.'
 WHERE name = 'Cálcio';

UPDATE magistral_components SET name = 'Ferro quelato',
       synonyms = 'ferro, ferro bisglicinato, bisglicinato ferroso, ferrochel',
       elemental_percent = 20,
       correction_note = 'O bisglicinato ferroso de mercado traz cerca de 20% de ferro elementar.'
 WHERE name = 'Ferro';

UPDATE magistral_components SET name = 'Manganês quelato',
       synonyms = 'manganês, MANGANÊS QUELATO, bisglicinato de manganês',
       correction_note = 'Percentual de manganês varia com o quelato: confirmar na ficha técnica.'
 WHERE name = 'Manganês';

UPDATE magistral_components SET name = 'Molibdênio quelato',
       synonyms = 'molibdênio, MOLIBEDÊNIO QUELATO, quelato de molibdênio',
       correction_note = 'Percentual varia com o quelato: confirmar na ficha técnica.'
 WHERE name = 'Molibdênio';

UPDATE magistral_components SET name = 'Potássio quelato',
       synonyms = 'potássio, POTÁSSIO QUELATO, citrato de potássio',
       correction_note = 'Percentual de potássio varia com o sal: confirmar na ficha técnica.'
 WHERE name = 'Potássio';

UPDATE magistral_components SET name = 'Boro quelato',
       synonyms = 'boro, BORO QUELATO, glicinato de boro',
       correction_note = 'Percentual de boro varia com o quelato: confirmar na ficha técnica.'
 WHERE name = 'Boro';

UPDATE magistral_components SET name = 'Vanádio quelato',
       synonyms = 'vanádio, VANADIO QUELADO, sulfato de vanadila',
       correction_note = 'Sulfato de vanadila e quelatos têm percentuais diferentes: confirmar na ficha técnica.'
 WHERE name = 'Vanádio';

UPDATE magistral_components SET name = 'Iodeto de potássio',
       synonyms = 'iodo, IODO, iodeto de potássio, KI',
       elemental_percent = 76.4,
       correction_note = 'Estequiometria do iodeto de potássio: 76,4% de iodo.'
 WHERE name = 'Iodo';

-- ---------------------------------------------------------------------------------------------
-- 3. As fórmulas acompanham o nome novo, e a dose passa a ser declarada do ELEMENTO
--
-- É a convenção: mineral se prescreve em quantidade do elemento, e a conversão para a massa do
-- insumo é da farmácia. Sem a marca, "cobre quelato 1 mg" seria lido como 1 mg do pó.
-- ---------------------------------------------------------------------------------------------
UPDATE magistral_formula_template_components c SET substance = m.novo, as_elemental = true,
       note = CASE WHEN coalesce(c.note,'') = '' THEN 'Dose em ' || m.elemento || ' elementar.' ELSE c.note END
  FROM (VALUES
    ('Cobre',      'Cobre quelato',       'cobre'),
    ('Cálcio',     'Cálcio quelato',      'cálcio'),
    ('Ferro',      'Ferro quelato',       'ferro'),
    ('Manganês',   'Manganês quelato',    'manganês'),
    ('Molibdênio', 'Molibdênio quelato',  'molibdênio'),
    ('Potássio',   'Potássio quelato',    'potássio'),
    ('Boro',       'Boro quelato',        'boro'),
    ('Vanádio',    'Vanádio quelato',     'vanádio'),
    ('Iodo',       'Iodeto de potássio',  'iodo')
  ) AS m(velho, novo, elemento)
 WHERE c.substance = m.velho AND c.deleted_at IS NULL;

-- Zinco e magnésio já tinham o nome certo, mas nem sempre a dose declarada como do elemento.
UPDATE magistral_formula_template_components c SET as_elemental = true,
       note = CASE WHEN coalesce(c.note,'') = '' THEN 'Dose em elemento.' ELSE c.note END
 WHERE c.deleted_at IS NULL AND c.as_elemental = false
   AND (c.substance ILIKE 'zinco%' OR c.substance ILIKE 'magnésio%' OR c.substance ILIKE 'selenometionina%'
        OR c.substance ILIKE 'picolinato de cromo%' OR c.substance ILIKE 'cromo %');

-- ==============================================================================================
-- fonte: docs/emr/magistral-orfaos-nome-canonico.sql
-- ==============================================================================================
-- Componentes de fórmula que não achavam entrada no catálogo.
--
-- Todos existiam sob o nome canônico; o que estava gravado era o texto cru do formulário das
-- parceiras ("Gimnema silvestre", "Cromo (GTF ou picolinato)", "Weg lem 70(ganoderma lucidum)").
-- Sem casar com o catálogo, o componente perde faixa de dose, densidade, teto da IN 28,
-- interferência de exame e marca de elementar — ou seja, some de todo o motor de conferência.
--
-- Onde o texto original carrega informação que o nome canônico perde (marca de insumo,
-- padronização), ela vai para a nota do componente, que a farmácia lê.


UPDATE magistral_formula_template_components c SET substance = m.novo,
       note = CASE WHEN coalesce(c.note,'') = '' THEN m.nota ELSE c.note || ' ' || m.nota END
  FROM (VALUES
    ('Alfa-GPC (L-alfa-glicerofosfocolina)',            'Alfa-GPC',                 ''),
    ('BHB',                                             'Beta-hidroxibutirato',     ''),
    ('Citrimax (ácido hidroxicítrico - HCA)',           'Ácido hidroxicítrico',     'Formulário indica Citrimax, extrato padronizado em HCA.'),
    ('Coenzima Q10 (ubiquinona) ou ubiquinol',          'Coenzima Q10',             'Ubiquinol aceito como alternativa.'),
    ('Cromo (GTF ou picolinato)',                       'Picolinato de cromo',      'Cromo GTF aceito como alternativa.'),
    ('Curcumina padronizada (95% curcuminoides)',       'Curcumina',                'Extrato padronizado em 95% de curcuminoides.'),
    ('Gimnema silvestre',                               'Gymnema silvestre',        ''),
    ('Ginostema pentaphyllum (Gynostemma pentaphyllum)','Gynostemma pentaphyllum',  ''),
    ('Glutamina',                                       'L-glutamina',              ''),
    ('Inositol',                                        'Mio-inositol',             ''),
    ('Magnésio glicina',                                'Magnésio quelato',         'Bisglicinato.'),
    ('Magnésio treonato',                               'Magnésio L-treonato',      ''),
    ('Piridoxal-5-fosfato (P5P)',                       'Piridoxal-5-fosfato',      ''),
    ('PQQ (pirroloquinolina quinona)',                  'PQQ',                      ''),
    ('Taurina',                                         'L-taurina',                ''),
    ('Vitamina B2 (riboflavina)',                       'Riboflavina',              ''),
    ('Vitamina B3 (nicotinamida)',                      'Nicotinamida',             ''),
    ('Vitamina B6 (piridoxal-5-fosfato)',               'Piridoxal-5-fosfato',      ''),
    ('Weg lem 70(ganoderma lucidum)',                   'Ganoderma lucidum',        'Formulário indica WEG LEM 70.')
  ) AS m(velho, novo, nota)
 WHERE c.substance = m.velho AND c.deleted_at IS NULL;

-- "Veiculo oleoso qsp" não é componente: é o veículo, que o parser leu como linha da fórmula.
-- Vitaminas lipossolúveis em solução pedem veículo oleoso; o que estava gravado ("hidroalcoólico")
-- veio de outra fórmula. As 5 gotas são a posologia, que a fórmula já tem.
UPDATE magistral_formula_templates SET vehicle = 'Veículo oleoso q.s.p.'
 WHERE name = 'Osteoporose – MIX vit lipolíticas';
DELETE FROM magistral_formula_template_components WHERE substance = 'Veiculo oleoso qsp';

-- Magnésio quelato: 30% é de quelato tamponado (cortado com óxido), não de bisglicinato puro.
-- O bisglicinato anidro Mg(C2H4NO2)2 pesa 172,4 g/mol para 24,305 de magnésio: 14,1%.
-- Fica o número conservador — subestimar o percentual superestima a massa de insumo, e errar
-- para "cabe numa cápsula maior" é o lado seguro do erro.
UPDATE magistral_components SET elemental_percent = 14.1,
  correction_note = 'Bisglicinato de magnésio anidro: 14,1% de magnésio elementar por estequiometria (24,305/172,4). Quelatos tamponados com óxido chegam a 20-30% — confirmar no laudo do lote. A tabela da parceira prescreve a dose já em elementar: 50 a 500 mg/dia.'
 WHERE lower(name) = 'magnésio quelato';

-- ==============================================================================================
-- fonte: docs/emr/in28-nomes-nota-rodape.sql
-- ==============================================================================================
-- Tetos da IN 28: tirar o marcador de nota de rodapé colado no nome.
--
-- O Anexo IV numera as notas em romano e a numeração veio grudada no nome do nutriente na carga:
-- "Cálciov", "Vitamina Dii", "Extrato de cacauviii", "Niacina ix". Duas consequências: o nome sai
-- torto no alerta, e — pior — o nome é a chave que liga `magistral_components.in28_nutrient` a
-- `in28_limits.nutrient`; limpar um lado só desligaria o teto em silêncio. Os dois mudam na mesma
-- transação.
--
-- Nenhum nutriente legítimo desta tabela termina em i/v/x minúsculo (conferido: 13 linhas, todas
-- com marcador). Duas linhas colidem depois da limpeza e cada uma tem tratamento próprio:
--
--   · "Niacina" e "Niacina ix" trazem o MESMO teto (35 mg) — é duplicata, e a segunda sai.
--   · "Beta-glucana (Euglena gracilis)" aparece com 187,5 mg e com 244 mg. São duas linhas
--     distintas do anexo, sob notas diferentes; fundir escolheria um número que a norma não diz.
--     A nota vira parte legível do nome e as duas continuam existindo.


DELETE FROM in28_limits d USING in28_limits k
 WHERE d.nutrient = 'Niacina ix' AND k.nutrient = 'Niacina'
   AND k.max_adult IS NOT DISTINCT FROM d.max_adult AND k.unit = d.unit;

-- O nome novo: tira o romano do fim. Quando a limpeza colidiria com uma linha que já existe com
-- o nome limpo, a nota entra legível no lugar de sumir — é o caso da beta-glucana de Euglena, que
-- aparece com 187,5 mg e com 244 mg sob notas diferentes. O EXISTS enxerga o estado do início do
-- comando, então a segunda atualização decide pelo mesmo critério da primeira.
--
-- Sem tabela temporária de propósito: `ON COMMIT DROP` só sobrevive porque o goose roda a
-- migration em transação, e uma migration não deve depender disso para não quebrar em silêncio.
UPDATE public.magistral_components c
   SET in28_nutrient = CASE
         WHEN EXISTS (SELECT 1 FROM public.in28_limits o
                       WHERE o.nutrient = regexp_replace(l.nutrient, '\s?[ivx]+$', ''))
         THEN regexp_replace(l.nutrient, '\s?[ivx]+$', '')
              || ' (nota ' || regexp_replace(l.nutrient, '^.*?\s?([ivx]+)$', '\1') || ')'
         ELSE regexp_replace(l.nutrient, '\s?[ivx]+$', '') END
  FROM public.in28_limits l
 WHERE c.in28_nutrient = l.nutrient AND l.nutrient ~ '[ivx]$';

UPDATE public.in28_limits l
   SET nutrient = CASE
         WHEN EXISTS (SELECT 1 FROM public.in28_limits o
                       WHERE o.nutrient = regexp_replace(l.nutrient, '\s?[ivx]+$', ''))
         THEN regexp_replace(l.nutrient, '\s?[ivx]+$', '')
              || ' (nota ' || regexp_replace(l.nutrient, '^.*?\s?([ivx]+)$', '\1') || ')'
         ELSE regexp_replace(l.nutrient, '\s?[ivx]+$', '') END
 WHERE l.nutrient ~ '[ivx]$';

-- ==============================================================================================
-- fonte: docs/emr/magistral-notas-redundantes.sql
-- ==============================================================================================
-- Volta o palmitato para dose do ativo, e tira as notas que viraram redundância no receituário.
--
-- ERRO CORRIGIDO AQUI: desmarquei os 25 componentes de palmitato de ascorbila achando que a dose
-- era do insumo. Não é. A fonte das parceiras escreve "VITAMINA C 100 mg" e a curadoria
-- (magistral-formulario-correcoes.sql) trocou a FORMA para palmitato mantendo a dose em vitamina
-- C, justamente para a troca não cortar a dose a 43%. Desmarcar cortava a dose de todas as 25.
--
-- O sinal do catálogo acompanha: `dose_as_elemental` é "a dose se escreve em ativo/elemento e a
-- farmácia converte", que é exatamente o caso — mineral por convenção, palmitato por decisão de
-- forma preferida.
--
-- E as notas "Dose em cobre elementar." que eu tinha posto em cada mineral saem: o PDF passou a
-- imprimir "(do elemento)" ao lado da quantidade, então a nota repetia a mesma informação no meio
-- da linha do receituário.


UPDATE magistral_components SET dose_as_elemental = true WHERE name = 'Palmitato de ascorbila';

UPDATE magistral_formula_template_components c SET as_elemental = true
 WHERE c.substance = 'Palmitato de ascorbila' AND c.deleted_at IS NULL
   AND c.note LIKE 'Forma trocada para palmitato%';

-- Nota mais curta: o "(do elemento)" já diz que a dose é do ativo.
UPDATE magistral_formula_template_components SET note = 'Dose de vitamina C; o palmitato é a forma preferida do prescritor.'
 WHERE substance = 'Palmitato de ascorbila' AND note LIKE 'Forma trocada para palmitato%';

UPDATE magistral_formula_template_components SET note = ''
 WHERE note ~ '^Dose em ([a-zà-ú]+ )?element(o|ar)\.$';


-- As 5 restantes vieram da mesma troca, por passes de curadoria diferentes ("troca do
-- prescritor"). Uma delas dizia na nota que a dose era do ativo e estava marcada como insumo —
-- contradição dentro da própria linha. Com `dose_as_elemental` no catálogo, palmitato é elementar
-- em todas.
UPDATE magistral_formula_template_components
   SET as_elemental = true,
       note = 'Dose de vitamina C; o palmitato é a forma preferida do prescritor.'
 WHERE substance = 'Palmitato de ascorbila' AND deleted_at IS NULL AND NOT as_elemental;

-- ==============================================================================================
-- Depois das renomeações, remarcar quem se prescreve em elemento: em produção a 00083 rodou sobre
-- os nomes antigos ("Cobre", "Iodo"), e os novos ("Cobre quelato", "Iodeto de potássio") precisam
-- do mesmo sinal. É a mesma regra da 00083, reaplicada.
-- ==============================================================================================
UPDATE public.magistral_components SET dose_as_elemental = true
 WHERE deleted_at IS NULL AND NOT dose_as_elemental
   AND (in28_nutrient IN ('Cálcio','Cobre','Cromo','Ferro','Fósforo','Iodo','Magnésio','Manganês',
                          'Molibdênio','Potássio','Selênio','Zinco','Boro','Vanádio')
        OR name ~* '^(cobre|cálcio|ferro|manganês|molibdênio|potássio|boro|vanádio|zinco|magnésio|selênio|cromo|iodo)( |$)'
        OR name ~* '(quelato|quelado|bisglicinato)$'
        OR name IN ('Iodeto de potássio','Selenometionina','Picolinato de cromo','Zinco carnosina',
                    'Palmitato de ascorbila'));

-- +goose Down
-- Sem volta: desfazer devolveria "Cobre" no lugar de "Cobre quelato" e apagaria a marca de dose
-- elementar de receita já emitida. Corrige-se para a frente.
SELECT 1;
