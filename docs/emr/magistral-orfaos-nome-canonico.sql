-- Componentes de fórmula que não achavam entrada no catálogo.
--
-- Todos existiam sob o nome canônico; o que estava gravado era o texto cru do formulário das
-- parceiras ("Gimnema silvestre", "Cromo (GTF ou picolinato)", "Weg lem 70(ganoderma lucidum)").
-- Sem casar com o catálogo, o componente perde faixa de dose, densidade, teto da IN 28,
-- interferência de exame e marca de elementar — ou seja, some de todo o motor de conferência.
--
-- Onde o texto original carrega informação que o nome canônico perde (marca de insumo,
-- padronização), ela vai para a nota do componente, que a farmácia lê.

BEGIN;

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

COMMIT;
