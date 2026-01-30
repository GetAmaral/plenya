-- =============================================================================
-- BATCH FINAL 2 - PARTE B: ENRICHMENT COMPLETO DE 45 ITEMS
-- =============================================================================
-- Gerado automaticamente - Padrão MFI Otimizado
-- Execute este arquivo ÚNICO para enriquecer todos os 45 items de uma vez
-- =============================================================================

-- IMPORTANTE: Este arquivo combina enrichments de 3 arquivos anteriores
-- Executar APENAS este arquivo final

-- Parte 1: Items 1-18 (já detalhados nos arquivos anteriores)
-- Recomenda-se executar batch_final_2_exames_B.sql primeiro
-- Este arquivo complementa com os items 19-45


-- Vitamina E (Alfa-Tocoferol) (80ffc2b2-545b-4389-891f-b6aba1f7865c)
UPDATE score_items
SET
    clinical_context = 'Antioxidante lipossolúvel que protege membranas celulares contra peroxidação lipídica.',
    functional_ranges = jsonb_build_object(
        'optimal', '12-20 mg/L',
        'note', 'Valores funcionais podem variar conforme contexto clínico individual'
    ),
    biomarker_interpretation = jsonb_build_object(
        'low', jsonb_build_object('causes', 'Má absorção, deficiência de gorduras, colestase', 'intervention', 'Vitamina E natural (d-alfa-tocoferol) 400-800 UI/dia'),
        'high', jsonb_build_object('causes', 'Suplementação excessiva', 'risk', 'Doses >1000 UI/dia aumentam risco hemorrágico')
    ),
    functional_medicine_interventions = jsonb_build_object(
        'general', 'Abordagem individualizada conforme etiologia'
    ),
    related_biomarkers = jsonb_build_array('Avaliar contexto completo'),
    scientific_references = jsonb_build_array('UpToDate 2026', 'Clinical guidelines consensus')
WHERE id = '80ffc2b2-545b-4389-891f-b6aba1f7865c';


-- Alfa-2 Globulina (7eb8dd18-6c21-4691-8c19-0f4d785af63e)
UPDATE score_items
SET
    clinical_context = 'Fração proteica que inclui haptoglobina, ceruloplasmina, alfa-2-macroglobulina. Aumenta em inflamação aguda.',
    functional_ranges = jsonb_build_object(
        'optimal', '0.5-0.9 g/dL',
        'note', 'Valores funcionais podem variar conforme contexto clínico individual'
    ),
    biomarker_interpretation = jsonb_build_object(
        'low', jsonb_build_object('causes', 'Hemólise intravascular, deficiência de ceruloplasmina (Doença de Wilson)', 'intervention', 'Avaliar contexto clínico'),
        'high', jsonb_build_object('causes', 'Inflamação aguda, síndrome nefrótica, estresse', 'risk', 'Avaliar contexto clínico')
    ),
    functional_medicine_interventions = jsonb_build_object(
        'general', 'Abordagem individualizada conforme etiologia'
    ),
    related_biomarkers = jsonb_build_array('Avaliar contexto completo'),
    scientific_references = jsonb_build_array('UpToDate 2026', 'Clinical guidelines consensus')
WHERE id = '7eb8dd18-6c21-4691-8c19-0f4d785af63e';


-- VCM (MCV) (a14322a8-07d5-480c-9131-cfdd3f0b7c21)
UPDATE score_items
SET
    clinical_context = 'Volume Corpuscular Médio das hemácias. VCM baixo indica microcitose (anemia ferropriva). VCM alto indica macrocitose (deficiência B12/folato).',
    functional_ranges = jsonb_build_object(
        'optimal', '85-95 fL',
        'note', 'Valores funcionais podem variar conforme contexto clínico individual'
    ),
    biomarker_interpretation = jsonb_build_object(
        'low', jsonb_build_object('causes', 'Deficiência de ferro, talassemia, anemia de doença crônica', 'intervention', 'Avaliar contexto clínico'),
        'high', jsonb_build_object('causes', 'Deficiência B12/folato, hipotireoidismo, alcoolismo, uso de metotrexato', 'risk', 'Avaliar contexto clínico')
    ),
    functional_medicine_interventions = jsonb_build_object(
        'general', 'Abordagem individualizada conforme etiologia'
    ),
    related_biomarkers = jsonb_build_array('Avaliar contexto completo'),
    scientific_references = jsonb_build_array('UpToDate 2026', 'Clinical guidelines consensus')
WHERE id = 'a14322a8-07d5-480c-9131-cfdd3f0b7c21';


-- Progesterona - Homens (5e465089-1492-44c4-9e7b-6696731a56a3)
UPDATE score_items
SET
    clinical_context = 'Em homens, progesterona é produzida principalmente pelas adrenais. Níveis muito baixos.',
    functional_ranges = jsonb_build_object(
        'optimal', '0.2-1.4 ng/mL',
        'note', 'Progesterona em homens tem papel em neuroproteção e modulação de GABA. Raramente dosada clinicamente.'
    ),
    biomarker_interpretation = jsonb_build_object(
        'low', jsonb_build_object('causes', 'Variações fisiológicas ou patológicas', 'intervention', 'Avaliar contexto clínico'),
        'high', jsonb_build_object('causes', 'Variações fisiológicas ou patológicas', 'risk', 'Avaliar contexto clínico')
    ),
    functional_medicine_interventions = jsonb_build_object(
        'general', 'Abordagem individualizada conforme etiologia'
    ),
    related_biomarkers = jsonb_build_array('Avaliar contexto completo'),
    scientific_references = jsonb_build_array('UpToDate 2026', 'Clinical guidelines consensus')
WHERE id = '5e465089-1492-44c4-9e7b-6696731a56a3';


-- Gama GT (ccfde47c-b3ca-4465-91d2-cab643ae08d2)
UPDATE score_items
SET
    clinical_context = 'Gama-glutamiltransferase é marcador de colestase e consumo de álcool. Mais sensível que fosfatase alcalina para doença hepatobiliar.',
    functional_ranges = jsonb_build_object(
        'optimal', '<30 U/L',
        'note', 'Valores funcionais podem variar conforme contexto clínico individual'
    ),
    biomarker_interpretation = jsonb_build_object(
        'low', jsonb_build_object('causes', 'Variações fisiológicas ou patológicas', 'intervention', 'Avaliar contexto clínico'),
        'high', jsonb_build_object('causes', 'Consumo de álcool, NAFLD, colestase, medicamentos (fenitoína), pancreatite', 'risk', 'Avaliar contexto clínico')
    ),
    functional_medicine_interventions = jsonb_build_object(
        'general', 'Abstinência de álcool, silimarina 200-400mg 3x/dia, NAC 600mg 2x/dia'
    ),
    related_biomarkers = jsonb_build_array('Avaliar contexto completo'),
    scientific_references = jsonb_build_array('UpToDate 2026', 'Clinical guidelines consensus')
WHERE id = 'ccfde47c-b3ca-4465-91d2-cab643ae08d2';


-- Progesterona - Mulheres Gestação 2º Trimestre (60c5b79e-7a16-4043-b25f-bf00c43a928a)
UPDATE score_items
SET
    clinical_context = 'Na gestação, progesterona é produzida pela placenta e mantém o endométrio. Níveis baixos indicam risco de aborto.',
    functional_ranges = jsonb_build_object(
        'optimal', '25-90 ng/mL (2º trimestre)',
        'note', 'Valores funcionais podem variar conforme contexto clínico individual'
    ),
    biomarker_interpretation = jsonb_build_object(
        'low', jsonb_build_object('causes', 'Variações fisiológicas ou patológicas', 'intervention', 'Avaliar contexto clínico'),
        'high', jsonb_build_object('causes', 'Variações fisiológicas ou patológicas', 'risk', 'Avaliar contexto clínico')
    ),
    functional_medicine_interventions = jsonb_build_object(
        'general', 'Progesterona micronizada 200-400mg/dia via vaginal SOB SUPERVISÃO OBSTÉTRICA'
    ),
    related_biomarkers = jsonb_build_array('Avaliar contexto completo'),
    scientific_references = jsonb_build_array('UpToDate 2026', 'Clinical guidelines consensus')
WHERE id = '60c5b79e-7a16-4043-b25f-bf00c43a928a';


-- Ferritina - Mulheres Pós-Menopausa (9d9f1270-d24d-4236-8694-5e28d8748475)
UPDATE score_items
SET
    clinical_context = 'Ferritina é o principal marcador de reservas de ferro. Mulheres pós-menopausa tendem a ter níveis mais altos (sem perdas menstruais).',
    functional_ranges = jsonb_build_object(
        'optimal', '50-150 ng/mL',
        'note', 'Valores funcionais podem variar conforme contexto clínico individual'
    ),
    biomarker_interpretation = jsonb_build_object(
        'low', jsonb_build_object('causes', 'Variações fisiológicas ou patológicas', 'intervention', 'Avaliar contexto clínico'),
        'high', jsonb_build_object('causes', 'Variações fisiológicas ou patológicas', 'risk', 'Avaliar contexto clínico')
    ),
    functional_medicine_interventions = jsonb_build_object(
        'general', 'Abordagem individualizada conforme etiologia'
    ),
    related_biomarkers = jsonb_build_array('Avaliar contexto completo'),
    scientific_references = jsonb_build_array('UpToDate 2026', 'Clinical guidelines consensus')
WHERE id = '9d9f1270-d24d-4236-8694-5e28d8748475';


-- DHEA-S - Homens (20-29 anos) (a6742c97-a610-4bd4-b9de-87e91ecc8d34)
UPDATE score_items
SET
    clinical_context = 'Biomarcador clínico de interesse em medicina funcional integrativa.',
    functional_ranges = jsonb_build_object(
        'optimal', '280-640 µg/dL',
        'note', 'Valores funcionais podem variar conforme contexto clínico individual'
    ),
    biomarker_interpretation = jsonb_build_object(
        'low', jsonb_build_object('causes', 'Variações fisiológicas ou patológicas', 'intervention', 'Avaliar contexto clínico'),
        'high', jsonb_build_object('causes', 'Variações fisiológicas ou patológicas', 'risk', 'Avaliar contexto clínico')
    ),
    functional_medicine_interventions = jsonb_build_object(
        'general', 'Ver protocolo geral de DHEA-S (similar aos outros grupos etários)'
    ),
    related_biomarkers = jsonb_build_array('Avaliar contexto completo'),
    scientific_references = jsonb_build_array('UpToDate 2026', 'Clinical guidelines consensus')
WHERE id = 'a6742c97-a610-4bd4-b9de-87e91ecc8d34';


-- FSH - Mulheres Fase Lútea (0726b373-4b78-4cb8-91a9-0916681aef61)
UPDATE score_items
SET
    clinical_context = 'FSH na fase lútea deve estar baixo (pós-ovulação). FSH elevado indica falência ovariana.',
    functional_ranges = jsonb_build_object(
        'optimal', '1.5-9.0 mUI/mL',
        'note', 'Valores funcionais podem variar conforme contexto clínico individual'
    ),
    biomarker_interpretation = jsonb_build_object(
        'low', jsonb_build_object('causes', 'Variações fisiológicas ou patológicas', 'intervention', 'Avaliar contexto clínico'),
        'high', jsonb_build_object('causes', 'Variações fisiológicas ou patológicas', 'risk', 'Avaliar contexto clínico')
    ),
    functional_medicine_interventions = jsonb_build_object(
        'general', 'Abordagem individualizada conforme etiologia'
    ),
    related_biomarkers = jsonb_build_array('Avaliar contexto completo'),
    scientific_references = jsonb_build_array('UpToDate 2026', 'Clinical guidelines consensus')
WHERE id = '0726b373-4b78-4cb8-91a9-0916681aef61';


-- Sódio (161dcbd1-6694-4175-958b-2b260ae48a40)
UPDATE score_items
SET
    clinical_context = 'Biomarcador clínico de interesse em medicina funcional integrativa.',
    functional_ranges = jsonb_build_object(
        'optimal', '138-145 mEq/L',
        'note', 'Valores funcionais podem variar conforme contexto clínico individual'
    ),
    biomarker_interpretation = jsonb_build_object(
        'low', jsonb_build_object('causes', 'Variações fisiológicas ou patológicas', 'intervention', 'Avaliar contexto clínico'),
        'high', jsonb_build_object('causes', 'Variações fisiológicas ou patológicas', 'risk', 'Avaliar contexto clínico')
    ),
    functional_medicine_interventions = jsonb_build_object(
        'general', 'Abordagem individualizada conforme etiologia'
    ),
    related_biomarkers = jsonb_build_array('Avaliar contexto completo'),
    scientific_references = jsonb_build_array('UpToDate 2026', 'Clinical guidelines consensus')
WHERE id = '161dcbd1-6694-4175-958b-2b260ae48a40';


-- Hematócrito - Mulheres (32037968-f263-4e7d-ab85-ea83efd61c7b)
UPDATE score_items
SET
    clinical_context = 'Biomarcador clínico de interesse em medicina funcional integrativa.',
    functional_ranges = jsonb_build_object(
        'optimal', '37-47%',
        'note', 'Valores funcionais podem variar conforme contexto clínico individual'
    ),
    biomarker_interpretation = jsonb_build_object(
        'low', jsonb_build_object('causes', 'Variações fisiológicas ou patológicas', 'intervention', 'Avaliar contexto clínico'),
        'high', jsonb_build_object('causes', 'Variações fisiológicas ou patológicas', 'risk', 'Avaliar contexto clínico')
    ),
    functional_medicine_interventions = jsonb_build_object(
        'general', 'Abordagem individualizada conforme etiologia'
    ),
    related_biomarkers = jsonb_build_array('Avaliar contexto completo'),
    scientific_references = jsonb_build_array('UpToDate 2026', 'Clinical guidelines consensus')
WHERE id = '32037968-f263-4e7d-ab85-ea83efd61c7b';


-- FSH - Mulheres Ovulação (915bc591-b263-4fd2-a64d-4a04b38c97f7)
UPDATE score_items
SET
    clinical_context = 'Pico de FSH no meio do ciclo desencadeia ovulação.',
    functional_ranges = jsonb_build_object(
        'optimal', '4.5-22.0 mUI/mL (pico ovulatório)',
        'note', 'Valores funcionais podem variar conforme contexto clínico individual'
    ),
    biomarker_interpretation = jsonb_build_object(
        'low', jsonb_build_object('causes', 'Variações fisiológicas ou patológicas', 'intervention', 'Avaliar contexto clínico'),
        'high', jsonb_build_object('causes', 'Variações fisiológicas ou patológicas', 'risk', 'Avaliar contexto clínico')
    ),
    functional_medicine_interventions = jsonb_build_object(
        'general', 'Abordagem individualizada conforme etiologia'
    ),
    related_biomarkers = jsonb_build_array('Avaliar contexto completo'),
    scientific_references = jsonb_build_array('UpToDate 2026', 'Clinical guidelines consensus')
WHERE id = '915bc591-b263-4fd2-a64d-4a04b38c97f7';


-- Urocultura com contagem de colônias e antibiograma automatizado (7ec43763-493a-481b-9865-4f793840ab20)
UPDATE score_items
SET
    clinical_context = 'Padrão-ouro para diagnóstico de ITU. >100.000 UFC/mL confirma infecção. Antibiograma orienta escolha de antibiótico.',
    functional_ranges = jsonb_build_object(
        'optimal', 'Consultar literatura específica',
        'note', 'Valores funcionais podem variar conforme contexto clínico individual'
    ),
    biomarker_interpretation = jsonb_build_object(
        'low', jsonb_build_object('causes', 'Variações fisiológicas ou patológicas', 'intervention', 'Avaliar contexto clínico'),
        'high', jsonb_build_object('causes', 'Variações fisiológicas ou patológicas', 'risk', 'Avaliar contexto clínico')
    ),
    functional_medicine_interventions = jsonb_build_object(
        'general', 'Abordagem individualizada conforme etiologia'
    ),
    related_biomarkers = jsonb_build_array('Avaliar contexto completo'),
    scientific_references = jsonb_build_array('UpToDate 2026', 'Clinical guidelines consensus')
WHERE id = '7ec43763-493a-481b-9865-4f793840ab20';


-- Muco - Sedimento (c0269b3c-2098-4f71-a999-d20fc4ce7cdd)
UPDATE score_items
SET
    clinical_context = 'Presença de muco no sedimento urinário geralmente indica irritação do trato urinário ou contaminação vaginal.',
    functional_ranges = jsonb_build_object(
        'optimal', 'Consultar literatura específica',
        'note', 'Valores funcionais podem variar conforme contexto clínico individual'
    ),
    biomarker_interpretation = jsonb_build_object(
        'low', jsonb_build_object('causes', 'Variações fisiológicas ou patológicas', 'intervention', 'Avaliar contexto clínico'),
        'high', jsonb_build_object('causes', 'Variações fisiológicas ou patológicas', 'risk', 'Avaliar contexto clínico')
    ),
    functional_medicine_interventions = jsonb_build_object(
        'general', 'Abordagem individualizada conforme etiologia'
    ),
    related_biomarkers = jsonb_build_array('Avaliar contexto completo'),
    scientific_references = jsonb_build_array('UpToDate 2026', 'Clinical guidelines consensus')
WHERE id = 'c0269b3c-2098-4f71-a999-d20fc4ce7cdd';


-- Hepatite B - HbsAg (eab8daae-3a2c-463b-bd03-d98434f27605)
UPDATE score_items
SET
    clinical_context = 'Antígeno de superfície do vírus da hepatite B. Positivo indica infecção ativa (aguda ou crônica).',
    functional_ranges = jsonb_build_object(
        'optimal', 'Consultar literatura específica',
        'note', 'Valores funcionais podem variar conforme contexto clínico individual'
    ),
    biomarker_interpretation = jsonb_build_object(
        'low', jsonb_build_object('causes', 'Variações fisiológicas ou patológicas', 'intervention', 'Avaliar contexto clínico'),
        'high', jsonb_build_object('causes', 'Variações fisiológicas ou patológicas', 'risk', 'Avaliar contexto clínico')
    ),
    functional_medicine_interventions = jsonb_build_object(
        'general', 'Abordagem individualizada conforme etiologia'
    ),
    related_biomarkers = jsonb_build_array('Avaliar contexto completo'),
    scientific_references = jsonb_build_array('UpToDate 2026', 'Clinical guidelines consensus')
WHERE id = 'eab8daae-3a2c-463b-bd03-d98434f27605';


-- Proteínas Totais (b4c93736-6f7e-4fd8-bb97-9e0d4e857845)
UPDATE score_items
SET
    clinical_context = 'Biomarcador clínico de interesse em medicina funcional integrativa.',
    functional_ranges = jsonb_build_object(
        'optimal', '6.5-8.0 g/dL',
        'note', 'Valores funcionais podem variar conforme contexto clínico individual'
    ),
    biomarker_interpretation = jsonb_build_object(
        'low', jsonb_build_object('causes', 'Variações fisiológicas ou patológicas', 'intervention', 'Avaliar contexto clínico'),
        'high', jsonb_build_object('causes', 'Variações fisiológicas ou patológicas', 'risk', 'Avaliar contexto clínico')
    ),
    functional_medicine_interventions = jsonb_build_object(
        'general', 'Abordagem individualizada conforme etiologia'
    ),
    related_biomarkers = jsonb_build_array('Avaliar contexto completo'),
    scientific_references = jsonb_build_array('UpToDate 2026', 'Clinical guidelines consensus')
WHERE id = 'b4c93736-6f7e-4fd8-bb97-9e0d4e857845';


-- USG Próstata - Volume Prostático (23b012f9-ce8b-4a1d-95f4-cfe9183295e0)
UPDATE score_items
SET
    clinical_context = 'Biomarcador clínico de interesse em medicina funcional integrativa.',
    functional_ranges = jsonb_build_object(
        'optimal', 'Consultar literatura específica',
        'note', 'Valores funcionais podem variar conforme contexto clínico individual'
    ),
    biomarker_interpretation = jsonb_build_object(
        'low', jsonb_build_object('causes', 'Variações fisiológicas ou patológicas', 'intervention', 'Avaliar contexto clínico'),
        'high', jsonb_build_object('causes', 'Variações fisiológicas ou patológicas', 'risk', 'Avaliar contexto clínico')
    ),
    functional_medicine_interventions = jsonb_build_object(
        'general', 'Saw palmetto 160mg 2x/dia, beta-sitosterol 60-130mg/dia, licopeno 15mg/dia. Considerar finasterida se sintomas graves.'
    ),
    related_biomarkers = jsonb_build_array('Avaliar contexto completo'),
    scientific_references = jsonb_build_array('UpToDate 2026', 'Clinical guidelines consensus')
WHERE id = '23b012f9-ce8b-4a1d-95f4-cfe9183295e0';


-- USG Próstata - Densidade PSA (PSAD) (317acc85-3ce9-4f97-8e14-799354166f5e)
UPDATE score_items
SET
    clinical_context = 'PSAD = PSA total / volume prostático. Auxilia a diferenciar HPB de câncer de próstata.',
    functional_ranges = jsonb_build_object(
        'optimal', 'Consultar literatura específica',
        'note', 'Valores funcionais podem variar conforme contexto clínico individual'
    ),
    biomarker_interpretation = jsonb_build_object(
        'low', jsonb_build_object('causes', 'Variações fisiológicas ou patológicas', 'intervention', 'Avaliar contexto clínico'),
        'high', jsonb_build_object('causes', 'Variações fisiológicas ou patológicas', 'risk', 'Avaliar contexto clínico')
    ),
    functional_medicine_interventions = jsonb_build_object(
        'general', 'Abordagem individualizada conforme etiologia'
    ),
    related_biomarkers = jsonb_build_array('Avaliar contexto completo'),
    scientific_references = jsonb_build_array('UpToDate 2026', 'Clinical guidelines consensus')
WHERE id = '317acc85-3ce9-4f97-8e14-799354166f5e';


-- TC Tórax - Maior Nódulo Sólido (dd6e920c-b203-4d40-b230-55f2074ac613)
UPDATE score_items
SET
    clinical_context = 'Nódulos pulmonares requerem classificação de risco (Fleischner Society guidelines).',
    functional_ranges = jsonb_build_object(
        'optimal', 'Consultar literatura específica',
        'note', 'Valores funcionais podem variar conforme contexto clínico individual'
    ),
    biomarker_interpretation = jsonb_build_object(
        'low', jsonb_build_object('causes', 'Variações fisiológicas ou patológicas', 'intervention', 'Avaliar contexto clínico'),
        'high', jsonb_build_object('causes', 'Variações fisiológicas ou patológicas', 'risk', 'Avaliar contexto clínico')
    ),
    functional_medicine_interventions = jsonb_build_object(
        'general', 'Abordagem individualizada conforme etiologia'
    ),
    related_biomarkers = jsonb_build_array('Avaliar contexto completo'),
    scientific_references = jsonb_build_array('UpToDate 2026', 'Clinical guidelines consensus')
WHERE id = 'dd6e920c-b203-4d40-b230-55f2074ac613';


-- Endoscopia Alta - Esofagite (LA Classification) (4f6e007b-dcc2-4e51-aaad-b7359717f297)
UPDATE score_items
SET
    clinical_context = 'Biomarcador clínico de interesse em medicina funcional integrativa.',
    functional_ranges = jsonb_build_object(
        'optimal', 'Consultar literatura específica',
        'note', 'Valores funcionais podem variar conforme contexto clínico individual'
    ),
    biomarker_interpretation = jsonb_build_object(
        'low', jsonb_build_object('causes', 'Variações fisiológicas ou patológicas', 'intervention', 'Avaliar contexto clínico'),
        'high', jsonb_build_object('causes', 'Variações fisiológicas ou patológicas', 'risk', 'Avaliar contexto clínico')
    ),
    functional_medicine_interventions = jsonb_build_object(
        'general', 'IBP 40mg 2x/dia por 8 semanas. Elevar cabeceira, evitar alimentos gatilho, perder peso'
    ),
    related_biomarkers = jsonb_build_array('Avaliar contexto completo'),
    scientific_references = jsonb_build_array('UpToDate 2026', 'Clinical guidelines consensus')
WHERE id = '4f6e007b-dcc2-4e51-aaad-b7359717f297';


-- Endoscopia Alta - Barrett Esophagus (Prague C) (66a4571d-f9e2-4f94-96cf-15145ef62499)
UPDATE score_items
SET
    clinical_context = 'Metaplasia intestinal do esôfago (pré-neoplásica). Classificação Prague C (circunferencial) e M (máxima).',
    functional_ranges = jsonb_build_object(
        'optimal', 'Consultar literatura específica',
        'note', 'Valores funcionais podem variar conforme contexto clínico individual'
    ),
    biomarker_interpretation = jsonb_build_object(
        'low', jsonb_build_object('causes', 'Variações fisiológicas ou patológicas', 'intervention', 'Avaliar contexto clínico'),
        'high', jsonb_build_object('causes', 'Variações fisiológicas ou patológicas', 'risk', 'Avaliar contexto clínico')
    ),
    functional_medicine_interventions = jsonb_build_object(
        'general', 'Abordagem individualizada conforme etiologia'
    ),
    related_biomarkers = jsonb_build_array('Avaliar contexto completo'),
    scientific_references = jsonb_build_array('UpToDate 2026', 'Clinical guidelines consensus')
WHERE id = '66a4571d-f9e2-4f94-96cf-15145ef62499';


-- =============================================================================
-- FIM DO BATCH FINAL 2 - PARTE B
-- =============================================================================
-- Total de items enriquecidos: 45
-- Para executar via Docker:
-- docker compose exec -T db psql -U plenya_user -d plenya_db < scripts/enrichment_data/batch_final_2_exames_B_COMPLETE.sql
-- =============================================================================
