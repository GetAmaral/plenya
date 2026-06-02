-- Esqueleto clínico P2b: lista de problemas (com CID-10 opcional), medicações
-- em uso (reconciliação) e catálogo CURADO de CID-10 para autocomplete.
-- O catálogo cobre condições crônicas/preventivas comuns (não é a ~14k oficial);
-- a descrição do problema é texto livre, então o código é apenas conveniência.

-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.cid_codes (
    code character varying(10) NOT NULL,
    description character varying(300) NOT NULL,
    chapter character varying(160),
    version character varying(8) NOT NULL DEFAULT '10',
    CONSTRAINT cid_codes_pkey PRIMARY KEY (code)
);

CREATE TABLE IF NOT EXISTS public.problem_list_items (
    id uuid NOT NULL,
    patient_id uuid NOT NULL,
    description character varying(300) NOT NULL,
    cid_code character varying(10),
    cid_version character varying(8) NOT NULL DEFAULT 'CID-10',
    status character varying(10) NOT NULL DEFAULT 'active',
    onset_date date,
    resolved_date date,
    onset_appointment_id uuid,
    notes text,
    recorded_by_user_id uuid NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone,
    CONSTRAINT problem_list_items_pkey PRIMARY KEY (id),
    CONSTRAINT chk_problem_list_items_status CHECK (status IN ('active','resolved','inactive'))
);
CREATE INDEX IF NOT EXISTS idx_problem_list_items_patient_id ON public.problem_list_items (patient_id);
CREATE INDEX IF NOT EXISTS idx_problem_list_items_active ON public.problem_list_items (patient_id) WHERE status = 'active' AND deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_problem_list_items_cid ON public.problem_list_items (cid_code);
CREATE INDEX IF NOT EXISTS idx_problem_list_items_deleted_at ON public.problem_list_items (deleted_at);

CREATE TABLE IF NOT EXISTS public.medications_in_use (
    id uuid NOT NULL,
    patient_id uuid NOT NULL,
    medication_name character varying(200) NOT NULL,
    active_ingredient character varying(200) NOT NULL DEFAULT '',
    dosage character varying(100) NOT NULL DEFAULT '',
    frequency character varying(100) NOT NULL DEFAULT '',
    route character varying(50) NOT NULL DEFAULT '',
    source character varying(20) NOT NULL DEFAULT 'patient_reported',
    status character varying(12) NOT NULL DEFAULT 'active',
    source_prescription_id uuid,
    reconciled_appointment_id uuid,
    start_date date,
    end_date date,
    notes text,
    recorded_by_user_id uuid NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone,
    CONSTRAINT medications_in_use_pkey PRIMARY KEY (id),
    CONSTRAINT chk_medications_in_use_source CHECK (source IN ('prescribed_here','external','patient_reported')),
    CONSTRAINT chk_medications_in_use_status CHECK (status IN ('active','suspended','stopped'))
);
CREATE INDEX IF NOT EXISTS idx_medications_in_use_patient_id ON public.medications_in_use (patient_id);
CREATE INDEX IF NOT EXISTS idx_medications_in_use_active ON public.medications_in_use (patient_id) WHERE status = 'active' AND deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_medications_in_use_deleted_at ON public.medications_in_use (deleted_at);

-- Seed curado de CID-10 (condições crônicas/preventivas comuns). Idempotente.
INSERT INTO public.cid_codes (code, description, chapter, version) VALUES
('E03.9','Hipotireoidismo não especificado','Endócrinas, nutricionais e metabólicas','10'),
('E04.9','Bócio não-tóxico não especificado','Endócrinas, nutricionais e metabólicas','10'),
('E05.9','Tireotoxicose não especificada (hipertireoidismo)','Endócrinas, nutricionais e metabólicas','10'),
('E06.3','Tireoidite autoimune','Endócrinas, nutricionais e metabólicas','10'),
('E10.9','Diabetes mellitus tipo 1, sem complicações','Endócrinas, nutricionais e metabólicas','10'),
('E11.9','Diabetes mellitus tipo 2, sem complicações','Endócrinas, nutricionais e metabólicas','10'),
('E16.2','Hipoglicemia não especificada','Endócrinas, nutricionais e metabólicas','10'),
('E28.2','Síndrome do ovário policístico','Endócrinas, nutricionais e metabólicas','10'),
('E55.9','Deficiência de vitamina D não especificada','Endócrinas, nutricionais e metabólicas','10'),
('E66.0','Obesidade devida a excesso de calorias','Endócrinas, nutricionais e metabólicas','10'),
('E66.9','Obesidade não especificada','Endócrinas, nutricionais e metabólicas','10'),
('E78.0','Hipercolesterolemia pura','Endócrinas, nutricionais e metabólicas','10'),
('E78.1','Hipertrigliceridemia pura','Endócrinas, nutricionais e metabólicas','10'),
('E78.2','Hiperlipidemia mista','Endócrinas, nutricionais e metabólicas','10'),
('E78.5','Hiperlipidemia não especificada','Endócrinas, nutricionais e metabólicas','10'),
('E79.0','Hiperuricemia sem sinais de artrite inflamatória','Endócrinas, nutricionais e metabólicas','10'),
('E88.9','Distúrbio metabólico não especificado','Endócrinas, nutricionais e metabólicas','10'),
('I10','Hipertensão essencial (primária)','Aparelho circulatório','10'),
('I11.9','Doença cardíaca hipertensiva sem insuficiência cardíaca','Aparelho circulatório','10'),
('I20.9','Angina pectoris não especificada','Aparelho circulatório','10'),
('I25.1','Doença aterosclerótica do coração','Aparelho circulatório','10'),
('I25.9','Doença isquêmica crônica do coração não especificada','Aparelho circulatório','10'),
('I48','Flutter e fibrilação atrial','Aparelho circulatório','10'),
('I49.9','Arritmia cardíaca não especificada','Aparelho circulatório','10'),
('I50.9','Insuficiência cardíaca não especificada','Aparelho circulatório','10'),
('I63.9','Infarto cerebral não especificado','Aparelho circulatório','10'),
('I64','Acidente vascular cerebral não especificado','Aparelho circulatório','10'),
('I70.9','Aterosclerose generalizada e não especificada','Aparelho circulatório','10'),
('I73.9','Doença vascular periférica não especificada','Aparelho circulatório','10'),
('I83.9','Varizes dos membros inferiores sem úlcera nem inflamação','Aparelho circulatório','10'),
('N18.1','Doença renal crônica, estágio 1','Aparelho geniturinário','10'),
('N18.2','Doença renal crônica, estágio 2 (leve)','Aparelho geniturinário','10'),
('N18.3','Doença renal crônica, estágio 3 (moderada)','Aparelho geniturinário','10'),
('N18.4','Doença renal crônica, estágio 4 (grave)','Aparelho geniturinário','10'),
('N18.5','Doença renal crônica, estágio 5','Aparelho geniturinário','10'),
('N18.9','Doença renal crônica não especificada','Aparelho geniturinário','10'),
('N20.0','Calculose do rim (litíase renal)','Aparelho geniturinário','10'),
('N39.0','Infecção do trato urinário de localização não especificada','Aparelho geniturinário','10'),
('N40','Hiperplasia da próstata','Aparelho geniturinário','10'),
('F10.2','Transtornos mentais devidos ao uso de álcool — síndrome de dependência','Transtornos mentais e comportamentais','10'),
('F17.2','Transtornos mentais devidos ao uso de fumo — síndrome de dependência','Transtornos mentais e comportamentais','10'),
('F32.9','Episódio depressivo não especificado','Transtornos mentais e comportamentais','10'),
('F33.9','Transtorno depressivo recorrente não especificado','Transtornos mentais e comportamentais','10'),
('F41.0','Transtorno de pânico','Transtornos mentais e comportamentais','10'),
('F41.1','Ansiedade generalizada','Transtornos mentais e comportamentais','10'),
('F41.9','Transtorno ansioso não especificado','Transtornos mentais e comportamentais','10'),
('F43.2','Transtornos de adaptação','Transtornos mentais e comportamentais','10'),
('F51.0','Insônia não orgânica','Transtornos mentais e comportamentais','10'),
('G40.9','Epilepsia não especificada','Sistema nervoso','10'),
('G43.9','Enxaqueca não especificada','Sistema nervoso','10'),
('G47.0','Distúrbios do início e da manutenção do sono (insônias)','Sistema nervoso','10'),
('G47.3','Apneia do sono','Sistema nervoso','10'),
('G62.9','Polineuropatia não especificada','Sistema nervoso','10'),
('M10.9','Gota não especificada','Osteomuscular e tecido conjuntivo','10'),
('M15.9','Poliartrose não especificada','Osteomuscular e tecido conjuntivo','10'),
('M16.9','Coxartrose não especificada','Osteomuscular e tecido conjuntivo','10'),
('M17.9','Gonartrose não especificada','Osteomuscular e tecido conjuntivo','10'),
('M19.9','Artrose não especificada','Osteomuscular e tecido conjuntivo','10'),
('M54.2','Cervicalgia','Osteomuscular e tecido conjuntivo','10'),
('M54.4','Lumbago com ciática','Osteomuscular e tecido conjuntivo','10'),
('M54.5','Dor lombar baixa','Osteomuscular e tecido conjuntivo','10'),
('M75.9','Lesão do ombro não especificada','Osteomuscular e tecido conjuntivo','10'),
('M79.7','Fibromialgia','Osteomuscular e tecido conjuntivo','10'),
('M81.0','Osteoporose pós-menopáusica','Osteomuscular e tecido conjuntivo','10'),
('M81.9','Osteoporose não especificada','Osteomuscular e tecido conjuntivo','10'),
('K21.0','Doença de refluxo gastroesofágico com esofagite','Aparelho digestivo','10'),
('K21.9','Doença de refluxo gastroesofágico sem esofagite','Aparelho digestivo','10'),
('K29.7','Gastrite não especificada','Aparelho digestivo','10'),
('K58.9','Síndrome do intestino irritável sem diarreia','Aparelho digestivo','10'),
('K59.0','Constipação','Aparelho digestivo','10'),
('K76.0','Degeneração gordurosa do fígado (esteatose hepática)','Aparelho digestivo','10'),
('K80.2','Calculose da vesícula biliar sem colecistite (colelitíase)','Aparelho digestivo','10'),
('J30.4','Rinite alérgica não especificada','Aparelho respiratório','10'),
('J44.9','Doença pulmonar obstrutiva crônica não especificada','Aparelho respiratório','10'),
('J45.9','Asma não especificada','Aparelho respiratório','10'),
('D50.9','Anemia por deficiência de ferro não especificada','Sangue e órgãos hematopoéticos','10'),
('D64.9','Anemia não especificada','Sangue e órgãos hematopoéticos','10'),
('L40.9','Psoríase não especificada','Pele e tecido subcutâneo','10'),
('L70.9','Acne não especificada','Pele e tecido subcutâneo','10'),
('B18.9','Hepatite viral crônica não especificada','Doenças infecciosas e parasitárias','10'),
('R03.0','Pressão arterial elevada, sem diagnóstico de hipertensão','Sintomas, sinais e achados anormais','10'),
('R53','Mal estar e fadiga','Sintomas, sinais e achados anormais','10'),
('R73.0','Anormalidade do exame de tolerância à glicose (pré-diabetes)','Sintomas, sinais e achados anormais','10'),
('R73.9','Hiperglicemia não especificada','Sintomas, sinais e achados anormais','10'),
('Z00.0','Exame médico geral','Fatores que influenciam o estado de saúde','10'),
('Z71.3','Aconselhamento e vigilância dietética','Fatores que influenciam o estado de saúde','10'),
('Z72.0','Uso de tabaco','Fatores que influenciam o estado de saúde','10'),
('Z72.3','Falta de exercício físico','Fatores que influenciam o estado de saúde','10')
ON CONFLICT (code) DO NOTHING;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.medications_in_use;
DROP TABLE IF EXISTS public.problem_list_items;
DROP TABLE IF EXISTS public.cid_codes;
-- +goose StatementEnd
