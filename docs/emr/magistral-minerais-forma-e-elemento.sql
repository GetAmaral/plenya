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

BEGIN;

-- ---------------------------------------------------------------------------------------------
-- 1. Selênio some: a forma que o prescritor usa é a selenometionina
-- ---------------------------------------------------------------------------------------------
UPDATE magistral_formula_template_components c
   SET substance = 'Selenometionina', as_elemental = true,
       note = CASE WHEN coalesce(note,'') = '' THEN 'Dose em selênio elementar.' ELSE note END
  FROM magistral_formula_templates t
 WHERE t.id = c.template_id AND c.substance = 'Selênio';

DELETE FROM magistral_components
 WHERE name = 'Selênio'
   AND NOT EXISTS (SELECT 1 FROM magistral_formula_template_components c WHERE c.substance = 'Selênio');

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

COMMIT;
