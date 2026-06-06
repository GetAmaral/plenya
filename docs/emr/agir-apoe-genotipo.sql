-- APOE genótipo ε (E2/E3/E4) consolidado — readout derivado da Bioma que a extração por-SNP perdeu. DEV-only.
BEGIN;
INSERT INTO score_items (id,name,subgroup_id,"order",is_light_version,anamnese_item_code,created_at,updated_at)
VALUES ('b1090000-0000-7000-8000-000000000001','APOE — Genótipo ε (E2/E3/E4)',
  (SELECT sg.id FROM score_subgroups sg JOIN score_groups g ON g.id=sg.group_id WHERE g."order"=13 AND sg.name='Cardiovascular' LIMIT 1),
  99,false,'APOE_GENOTIPO_E',now(),now())
ON CONFLICT (id) DO NOTHING;
INSERT INTO score_levels (id,level,name,item_id,operator,created_at,updated_at) VALUES
  ('b10a0000-0000-7000-8000-000000000000',0,'E2/E2 (proteção)','b1090000-0000-7000-8000-000000000001','=',now(),now()),
  ('b10a0000-0000-7000-8000-000000000001',1,'E2/E3','b1090000-0000-7000-8000-000000000001','=',now(),now()),
  ('b10a0000-0000-7000-8000-000000000002',2,'E3/E3 (mais comum)','b1090000-0000-7000-8000-000000000001','=',now(),now()),
  ('b10a0000-0000-7000-8000-000000000003',3,'E2/E4','b1090000-0000-7000-8000-000000000001','=',now(),now()),
  ('b10a0000-0000-7000-8000-000000000004',4,'E3/E4 (portador ε4 — risco aumentado)','b1090000-0000-7000-8000-000000000001','=',now(),now()),
  ('b10a0000-0000-7000-8000-000000000005',5,'E4/E4 (homozigoto ε4 — risco alto)','b1090000-0000-7000-8000-000000000001','=',now(),now())
ON CONFLICT (id) DO NOTHING;
INSERT INTO score_item_method_pillars (score_item_id,method_pillar_id) VALUES
  ('b1090000-0000-7000-8000-000000000001','a91d9014-0000-7000-8000-000000000000'),
  ('b1090000-0000-7000-8000-000000000001','a91d9010-0000-7000-8000-000000000000'),
  ('b1090000-0000-7000-8000-000000000001','a91d9001-0000-7000-8000-000000000000')
ON CONFLICT DO NOTHING;
COMMIT;
