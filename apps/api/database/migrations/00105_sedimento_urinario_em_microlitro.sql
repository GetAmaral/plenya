-- +goose Up
-- SEDIMENTO URINÁRIO EM /µL: o item existia só em células/campo, e o laboratório reporta em /µL.
--
-- O catálogo guarda hemácias, leucócitos e células epiteliais do sedimento com unidade `/µL`, e o
-- item de escore dos três está em `células/campo`. A bactéria tem o mesmo problema numa forma
-- diferente: o item só tem níveis com nome de cruz (Ausentes, Raras 1+, Moderadas 2+) e o laudo de
-- citometria entrega um número em /µL, que não casa com nome nenhum.
--
-- Nos três de contagem celular são grandezas diferentes — contagem por campo do microscópio contra
-- concentração — então a guarda de unidade recusa classificar e o resultado fica sem nível. A guarda está certa: sem ela, 0,6/µL cairia na faixa "≤5 células/campo" e o
-- paciente leria ÓTIMO sobre um número que ninguém comparou.
--
-- O efeito colateral é que o sedimento inteiro ficava fora do escore em todo laudo de citometria de
-- fluxo, que é o padrão dos laboratórios grandes hoje. Num paciente investigando doença glomerular,
-- o dado que mais pesa é justamente a contagem de hemácias — e era o que o escore não vigiava.
--
-- POR QUE UM ITEM NOVO E NÃO UMA CONVERSÃO
--
-- Não existe fator de conversão válido entre /µL e células/campo: a relação depende do volume
-- centrifugado, do volume de ressuspensão e da área do campo do microscópio, e muda de protocolo
-- para protocolo. Um fator fixo seria um número inventado com cara de rigor. É a mesma situação da
-- lipoproteína(a), que já tem um item por unidade (nmol/L e mg/dL) pelo mesmo motivo, e o motor já
-- sabe escolher a escala pela unidade do laudo (`filtraPelaUnidade`). O gêmeo da unidade errada
-- entra no snapshot como `not_applicable`, que NÃO conta no denominador — conferido: o denominador
-- do snapshot é a soma exata dos itens `evaluated`.
--
-- AS FAIXAS, E DE ONDE VÊM
--
-- O limiar que importa é o da hematúria: a AUA define hematúria microscópica como 3 ou mais
-- hemácias por campo, e o equivalente aceito em citometria de fluxo urinária é 25/µL. Bate com a
-- referência impressa pelo próprio laboratório nos laudos que originaram esta migration
-- (hemácias < 23/µL, leucócitos < 25/µL, células epiteliais < 31/µL).
--
-- Acima do limiar a literatura não sustenta uma tabela de faixas finas em /µL, e inventar fronteira
-- é pior do que ter menos níveis — mesma decisão tomada na 00102 para a razão apoB/apoA-1. Daí
-- quatro níveis por item (5, 3, 1 e 0) em vez de seis: o normal, a alteração confirmada, a
-- alteração importante e a intensa.
--
-- Célula epitelial não é doença: em quantidade alta ela diz que a amostra veio contaminada e que a
-- coleta precisa ser repetida, e é assim que o nível 0 deve ser lido.
--
-- A BACTÉRIA TEM DUAS ÂNCORAS QUE DISCORDAM, e as duas ficam registradas aqui em vez de eu escolher
-- em silêncio. A validação do Sysmex UF-5000 para rastreio de infecção urinária usa 30/µL (95,2% de
-- sensibilidade e 91,2% de valor preditivo negativo contra urocultura); o laboratório que emitiu os
-- laudos desta migration imprime como referência < 1.200/µL. São quarenta vezes de diferença, e ela
-- é real: o corte de rastreio quer não perder infecção, e o do laboratório quer não alarmar. Então
-- o nível 5 fica no corte de rastreio, o nível 0 no do laboratório, e entre os dois há um nível 3
-- que diz o que deve ser dito — que sozinho não decide, e pede leucocitúria e cultura. Se o Getúlio
-- preferir outro corte, é aqui que se mexe.
-- +goose StatementBegin
DO $sed$
DECLARE
  v_alvo   record;
  v_item   uuid;
  v_gemeo  record;
BEGIN
  FOR v_alvo IN
    SELECT * FROM (VALUES
      ('PLN6B5C27A4', 'Hemácias (RBC) - Sedimento (/µL)',
       'Contagem de hemácias na urina por citometria de fluxo. O limiar de hematúria microscópica é '
       '25/µL, equivalente aceito das 3 hemácias por campo da definição da AUA. Hematúria persistente '
       'com sedimento no restante limpo é o achado que abre investigação de doença glomerular; '
       'confirmada, pede microscopia com contraste de fase para procurar hemácia dismórfica, que é o '
       'que separa origem glomerular de urológica.'),
      ('PLN0CE7409E', 'Leucócitos (WBC) - Sedimento (/µL)',
       'Contagem de leucócitos na urina por citometria de fluxo. Acima de 25/µL configura piúria. '
       'Piúria com bacteriúria aponta infecção; piúria sem bactéria (piúria estéril) obriga a pensar '
       'em nefrite intersticial, litíase, tuberculose urinária e uso de anti-inflamatório.'),
      ('PLN0868A698', 'Bactérias - Sedimento (/µL)',
       'Contagem de bactérias na urina por citometria de fluxo. Sozinha não fecha infecção urinária: '
       'bacteriúria sem leucocitúria e sem sintoma costuma ser contaminação de coleta ou bacteriúria '
       'assintomática, que na maioria dos adultos não se trata. O que dá peso ao número é a companhia '
       'de piúria, de sintoma e de urocultura.'),
      ('PLN444CDDE8', 'Células Epiteliais - Sedimento (/µL)',
       'Contagem de células epiteliais escamosas na urina. Não é marcador de doença: é marcador de '
       'QUALIDADE DA AMOSTRA. Contagem alta indica contaminação por células da uretra distal ou da '
       'genitália externa, e o que ela pede é repetir a coleta de jato médio com higiene, não '
       'investigar o rim. Referência do laboratório: abaixo de 31/µL.')
    ) AS t(codigo, nome, relevancia)
  LOOP
    -- `IS DISTINCT FROM` e não `<>`: com unit NULL o `<>` devolve NULL, o gêmeo não é achado e a
    -- migration vira um no-op silencioso para aquele exame — goose não trata RAISE NOTICE como
    -- falha, então ninguém saberia.
    -- O gêmeo é o item que já existe para este exame em OUTRA unidade: 'células/campo' nos três de
    -- contagem celular e 'Descritivo' na bactéria, cujos níveis têm nome de cruz. Dele saem
    -- subgrupo, ordem, pontos e pilares: os dois são o MESMO exame e têm de pesar igual.
    SELECT si.subgroup_id, si."order", si.points INTO v_gemeo
      FROM public.score_items si
     WHERE si.lab_test_code = v_alvo.codigo AND si.unit IS DISTINCT FROM '/µL' AND si.deleted_at IS NULL
     ORDER BY si."order"
     LIMIT 1;

    IF v_gemeo.subgroup_id IS NULL THEN
      RAISE NOTICE 'sem item gêmeo para %: 00105 pulou este exame', v_alvo.codigo;
      CONTINUE;
    END IF;

    SELECT id INTO v_item FROM public.score_items
     WHERE lab_test_code = v_alvo.codigo AND unit = '/µL' AND deleted_at IS NULL
     LIMIT 1;

    IF v_item IS NULL THEN
      INSERT INTO public.score_items
        (id, name, unit, points, "order", subgroup_id, lab_test_code, clinical_relevance,
         created_at, updated_at)
      VALUES (uuid_generate_v7(), v_alvo.nome, '/µL', v_gemeo.points, v_gemeo."order",
              v_gemeo.subgroup_id, v_alvo.codigo, v_alvo.relevancia, now(), now())
      RETURNING id INTO v_item;
    END IF;

    INSERT INTO public.score_levels (id, item_id, level, name, operator, lower_limit, upper_limit, created_at, updated_at)
    SELECT uuid_generate_v7(), v_item, v.nivel, v.rotulo, v.op, v.inf, v.sup, now(), now()
    FROM (VALUES
      ('PLN6B5C27A4', 5, '≤25',      '<=',      NULL,  '25'),
      ('PLN6B5C27A4', 3, '25-100',   'between', '25',  '100'),
      ('PLN6B5C27A4', 1, '100-500',  'between', '100', '500'),
      ('PLN6B5C27A4', 0, '>500',     '>',       '500', NULL),
      ('PLN0CE7409E', 5, '≤25',      '<=',      NULL,  '25'),
      ('PLN0CE7409E', 3, '25-100',   'between', '25',  '100'),
      ('PLN0CE7409E', 1, '100-500',  'between', '100', '500'),
      ('PLN0CE7409E', 0, '>500',     '>',       '500', NULL),
      ('PLN0868A698', 5, '≤30',      '<=',      NULL,  '30'),
      ('PLN0868A698', 3, '30-1200',  'between', '30',  '1200'),
      ('PLN0868A698', 0, '>1200',    '>',       '1200', NULL),
      ('PLN444CDDE8', 5, '≤31',      '<=',      NULL,  '31'),
      ('PLN444CDDE8', 3, '31-100',   'between', '31',  '100'),
      ('PLN444CDDE8', 1, '100-200',  'between', '100', '200'),
      ('PLN444CDDE8', 0, '>200 (contaminação)', '>', '200', NULL)
    ) AS v(codigo, nivel, rotulo, op, inf, sup)
    WHERE v.codigo = v_alvo.codigo
      AND NOT EXISTS (SELECT 1 FROM public.score_levels sl
                       WHERE sl.item_id = v_item AND sl.deleted_at IS NULL);

    -- Mesmos pilares AGIR do gêmeo.
    INSERT INTO public.score_item_method_pillars (score_item_id, method_pillar_id)
    SELECT v_item, mp.method_pillar_id
      FROM public.score_item_method_pillars mp
      JOIN public.score_items g ON g.id = mp.score_item_id
     WHERE g.lab_test_code = v_alvo.codigo AND g.unit IS DISTINCT FROM '/µL' AND g.deleted_at IS NULL
       AND NOT EXISTS (SELECT 1 FROM public.score_item_method_pillars x
                        WHERE x.score_item_id = v_item AND x.method_pillar_id = mp.method_pillar_id);
  END LOOP;
END
$sed$;

-- UROBILINOGÊNIO: o laudo entrega a PALAVRA "Normal" e o item só tem faixas numéricas, então nada
-- casa e o resultado sai do escore. Um nível chamado "Normal" resolve, e a única pergunta é qual
-- nível ele recebe. Na fita, "normal" é 0,2 a 1,0 mg/dL, que atravessa o nível 4 (0,1-0,6) e o
-- nível 5 (0,6-1,1) deste item. Escolhido o 4, que é o lado que NÃO infla o escore: dizer "bom"
-- quando o laudo disse apenas "normal" é o erro caro. Se o Getúlio preferir o 5, é um UPDATE de uma
-- linha aqui.
INSERT INTO public.score_levels (id, item_id, level, name, operator, created_at, updated_at)
SELECT uuid_generate_v7(), si.id, 4, 'Normal', '=', now(), now()
  FROM public.score_items si
 WHERE si.lab_test_code = 'PLN2515CA09' AND si.deleted_at IS NULL
   AND NOT EXISTS (SELECT 1 FROM public.score_levels sl
                    WHERE sl.item_id = si.id AND sl.name = 'Normal' AND sl.deleted_at IS NULL);
-- +goose StatementEnd

-- +goose Down
UPDATE public.score_levels sl SET deleted_at = now()
  FROM public.score_items si
 WHERE sl.item_id = si.id AND si.lab_test_code = 'PLN2515CA09'
   AND sl.name = 'Normal' AND sl.deleted_at IS NULL;
-- SOFT delete, pelo mesmo motivo da 00103: `patient_score_item_results` referencia o item com
-- RESTRICT assim que alguém é pontuado, e apagar as linhas dependentes destruiria o histórico.
DELETE FROM public.score_item_method_pillars p USING public.score_items si
 WHERE p.score_item_id = si.id
   AND si.lab_test_code IN ('PLN6B5C27A4', 'PLN0CE7409E', 'PLN444CDDE8', 'PLN0868A698') AND si.unit = '/µL';
UPDATE public.score_levels sl SET deleted_at = now()
  FROM public.score_items si
 WHERE sl.item_id = si.id AND sl.deleted_at IS NULL
   AND si.lab_test_code IN ('PLN6B5C27A4', 'PLN0CE7409E', 'PLN444CDDE8', 'PLN0868A698') AND si.unit = '/µL';
UPDATE public.score_items SET deleted_at = now()
 WHERE lab_test_code IN ('PLN6B5C27A4', 'PLN0CE7409E', 'PLN444CDDE8', 'PLN0868A698')
   AND unit = '/µL' AND deleted_at IS NULL;
