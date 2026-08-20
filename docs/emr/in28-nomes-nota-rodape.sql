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

BEGIN;

DELETE FROM in28_limits d USING in28_limits k
 WHERE d.nutrient = 'Niacina ix' AND k.nutrient = 'Niacina'
   AND k.max_adult IS NOT DISTINCT FROM d.max_adult AND k.unit = d.unit;

CREATE TEMP TABLE limpeza_in28 ON COMMIT DROP AS
WITH alvo AS (
  SELECT nutrient AS velho, regexp_replace(nutrient, '\s?([ivx]+)$', '') AS base,
         regexp_replace(nutrient, '^.*?\s?([ivx]+)$', '\1') AS nota
    FROM in28_limits WHERE nutrient ~ '[ivx]$'
)
SELECT a.velho,
       CASE WHEN EXISTS (SELECT 1 FROM in28_limits o WHERE o.nutrient <> a.velho AND o.nutrient = a.base)
            THEN a.base || ' (nota ' || a.nota || ')'
            ELSE a.base END AS novo
  FROM alvo a;

UPDATE magistral_components c SET in28_nutrient = l.novo
  FROM limpeza_in28 l WHERE c.in28_nutrient = l.velho;

UPDATE in28_limits t SET nutrient = l.novo
  FROM limpeza_in28 l WHERE t.nutrient = l.velho;

COMMIT;
