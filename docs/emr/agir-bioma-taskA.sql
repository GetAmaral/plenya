-- Task A (alinhar 81 genes à Bioma): 4 trocas + 6 adoções + SNPs extras. DEV-only.
BEGIN;
UPDATE score_items SET name='APOA1 rs1799837 (HDL)', anamnese_item_code='APOA1_RS1799837', updated_at=now() WHERE id='019c1a2b-a36f-7754-a61f-537e73e7def7';
UPDATE score_levels SET name='Genótipo | CC=risco', updated_at=now() WHERE item_id='019c1a2b-a36f-7754-a61f-537e73e7def7';
UPDATE score_items SET name='FADS2 rs174616 (Ômega-3/DHA)', anamnese_item_code='FADS2_RS174616', updated_at=now() WHERE id='019c1a2b-a36f-7b94-a267-7244733470b9';
UPDATE score_levels SET name='Genótipo | A=risco', updated_at=now() WHERE item_id='019c1a2b-a36f-7b94-a267-7244733470b9';
UPDATE score_items SET name='IRS1 rs2943641 (Resistência Insulina)', anamnese_item_code='IRS1_RS2943641', updated_at=now() WHERE id='019c1a2b-a36f-79d4-9f02-3ea19df6b4c3';
UPDATE score_levels SET name='Genótipo | C=risco', updated_at=now() WHERE item_id='019c1a2b-a36f-79d4-9f02-3ea19df6b4c3';
UPDATE score_items SET name='LPL rs13702 (Triglicerídeos)', anamnese_item_code='LPL_RS13702', updated_at=now() WHERE id='019c1a2b-a36f-7421-bc36-6ec4e2ff9595';
UPDATE score_levels SET name='Genótipo | TT=risco', updated_at=now() WHERE item_id='019c1a2b-a36f-7421-bc36-6ec4e2ff9595';
UPDATE score_items SET name='APOE rs429358 (Alzheimer e Lipídios)', anamnese_item_code='APOE_RS429358', updated_at=now() WHERE id='019c1a2b-a36f-7d3f-84e7-eca470ed8908';
UPDATE score_levels SET name='Genótipo (Bioma)', updated_at=now() WHERE item_id='019c1a2b-a36f-7d3f-84e7-eca470ed8908';
UPDATE score_items SET name='GSTM1 rs2071487 (Detoxificação)', anamnese_item_code='GSTM1_RS2071487', updated_at=now() WHERE id='019c1a2b-a36f-7b15-8286-10badc9d2e6f';
UPDATE score_levels SET name='Genótipo | CC=risco', updated_at=now() WHERE item_id='019c1a2b-a36f-7b15-8286-10badc9d2e6f';
UPDATE score_items SET name='GSTT1 rs2266633 (Detoxificação)', anamnese_item_code='GSTT1_RS2266633', updated_at=now() WHERE id='019c1a2b-a36f-7581-b3a6-19d4e16efd0f';
UPDATE score_levels SET name='Genótipo | TT=risco', updated_at=now() WHERE item_id='019c1a2b-a36f-7581-b3a6-19d4e16efd0f';
UPDATE score_items SET name='NAT2 rs1495741 (Acetilador)', anamnese_item_code='NAT2_RS1495741', updated_at=now() WHERE id='019c1a2b-a36f-79ff-8265-bc69a1329f40';
UPDATE score_levels SET name='Genótipo | A=risco', updated_at=now() WHERE item_id='019c1a2b-a36f-79ff-8265-bc69a1329f40';
UPDATE score_items SET name='INS rs3741208 VNTR (Diabetes Tipo 1)', anamnese_item_code='INS_RS3741208', updated_at=now() WHERE id='019c1a2b-a36f-7fb0-bf56-5376d91a2d40';
UPDATE score_levels SET name='Genótipo | AA=risco', updated_at=now() WHERE item_id='019c1a2b-a36f-7fb0-bf56-5376d91a2d40';
UPDATE score_items SET name='ALPL rs1697421 (Hipofosfatasia)', anamnese_item_code='ALPL_RS1697421', updated_at=now() WHERE id='019c1a2b-a36f-789e-8b05-59cd5a745d09';
UPDATE score_levels SET name='Genótipo | TT=risco', updated_at=now() WHERE item_id='019c1a2b-a36f-789e-8b05-59cd5a745d09';
INSERT INTO score_items (id,name,subgroup_id,"order",is_light_version,anamnese_item_code,created_at,updated_at) VALUES
  ('b1030000-0000-7000-8000-000000000000','APOE rs7412','019c1a2b-a36f-7909-b035-cadd0684c8af',0,false,'APOE_RS7412',now(),now()),
  ('b1030000-0000-7000-8000-000000000001','GSTM1 rs74837985 (Destoxificação de xenobióticos)','b1050000-0000-7000-8000-000000000005',0,false,'GSTM1_RS74837985',now(),now()),
  ('b1030000-0000-7000-8000-000000000002','GSTT1 rs2266637 (Destoxificação de xenobióticos)','b1050000-0000-7000-8000-000000000005',0,false,'GSTT1_RS2266637',now(),now()),
  ('b1030000-0000-7000-8000-000000000003','EPHX1 rs2234922 (Destoxificação de epóxidos)','b1050000-0000-7000-8000-000000000005',0,false,'EPHX1_RS2234922',now(),now()),
  ('b1030000-0000-7000-8000-000000000004','FTO rs8050136 (Supressão de Grelina)','019c1a2b-a36f-7df0-a68b-ebec86a2588b',0,false,'FTO_RS8050136',now(),now()),
  ('b1030000-0000-7000-8000-000000000005','GSTP1 rs1138272 (Destoxificação de xenobióticos)','b1050000-0000-7000-8000-000000000005',0,false,'GSTP1_RS1138272',now(),now()),
  ('b1030000-0000-7000-8000-000000000006','LDLR rs6511720 (Receptor hepático de LDL, envolvido com a endocito)','019c1a2b-a36f-7909-b035-cadd0684c8af',0,false,'LDLR_RS6511720',now(),now()),
  ('b1030000-0000-7000-8000-000000000007','LIPC rs2070895 (Regula o transporte reverso de colesterol)','019c1a2b-a36f-7909-b035-cadd0684c8af',0,false,'LIPC_RS2070895',now(),now()),
  ('b1030000-0000-7000-8000-000000000008','MC4R rs12970134 (Regulação do apetite e gasto energético)','019c1a2b-a36f-7df0-a68b-ebec86a2588b',0,false,'MC4R_RS12970134',now(),now()),
  ('b1030000-0000-7000-8000-000000000009','MTHFR rs1801131 (Conversão de ácido fólico em metilfolato)','019c1a2b-a36f-7df0-a68b-ebec86a2588b',0,false,'MTHFR_RS1801131',now(),now()),
  ('b1030000-0000-7000-8000-00000000000a','PCSK9 rs11206510 (Degradação de receptores hepáticos de LDL)','019c1a2b-a36f-7909-b035-cadd0684c8af',0,false,'PCSK9_RS11206510',now(),now()),
  ('b1030000-0000-7000-8000-00000000000b','PCSK9 rs505151 (Degradação de receptores hepáticos de LDL)','019c1a2b-a36f-7909-b035-cadd0684c8af',0,false,'PCSK9_RS505151',now(),now()),
  ('b1030000-0000-7000-8000-00000000000c','TCF7L2 rs12255372 (Homeostase de glicose)','019c1a2b-a36f-7df0-a68b-ebec86a2588b',0,false,'TCF7L2_RS12255372',now(),now())
ON CONFLICT (id) DO NOTHING;
INSERT INTO score_levels (id,level,name,item_id,operator,created_at,updated_at) VALUES
  ('b1040000-0000-7000-8000-000000000000',0,'Genótipo (Bioma)','b1030000-0000-7000-8000-000000000000','=',now(),now()),
  ('b1040000-0000-7000-8000-000000000001',0,'Genótipo | CC=risco','b1030000-0000-7000-8000-000000000001','=',now(),now()),
  ('b1040000-0000-7000-8000-000000000002',0,'Genótipo | TT=risco','b1030000-0000-7000-8000-000000000002','=',now(),now()),
  ('b1040000-0000-7000-8000-000000000003',0,'Genótipo | GG=risco','b1030000-0000-7000-8000-000000000003','=',now(),now()),
  ('b1040000-0000-7000-8000-000000000004',0,'Genótipo | A=risco','b1030000-0000-7000-8000-000000000004','=',now(),now()),
  ('b1040000-0000-7000-8000-000000000005',0,'Genótipo | T=risco','b1030000-0000-7000-8000-000000000005','=',now(),now()),
  ('b1040000-0000-7000-8000-000000000006',0,'Genótipo | GG=risco','b1030000-0000-7000-8000-000000000006','=',now(),now()),
  ('b1040000-0000-7000-8000-000000000007',0,'Genótipo | A=risco','b1030000-0000-7000-8000-000000000007','=',now(),now()),
  ('b1040000-0000-7000-8000-000000000008',0,'Genótipo | A=risco','b1030000-0000-7000-8000-000000000008','=',now(),now()),
  ('b1040000-0000-7000-8000-000000000009',0,'Genótipo | GG=risco','b1030000-0000-7000-8000-000000000009','=',now(),now()),
  ('b1040000-0000-7000-8000-00000000000a',0,'Genótipo | T=risco','b1030000-0000-7000-8000-00000000000a','=',now(),now()),
  ('b1040000-0000-7000-8000-00000000000b',0,'Genótipo | G=risco','b1030000-0000-7000-8000-00000000000b','=',now(),now()),
  ('b1040000-0000-7000-8000-00000000000c',0,'Genótipo | TT=risco','b1030000-0000-7000-8000-00000000000c','=',now(),now())
ON CONFLICT (id) DO NOTHING;
COMMIT;
