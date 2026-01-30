-- Fix Clinical Relevance for Leptina - Mulheres to meet character count target
BEGIN;

UPDATE score_items
SET
    clinical_relevance = 'A leptina é um hormônio peptídico de 16 kDa produzido principalmente pelo tecido adiposo branco que desempenha papel crucial na regulação do apetite, metabolismo energético e função reprodutiva feminina. Em mulheres, os níveis de leptina variam significativamente durante o ciclo menstrual, com elevações geralmente observadas na fase lútea média (11.4 ng/mL) comparada à fase folicular (10.0 ng/mL), apresentando pico no meio do ciclo durante o pico de LH (até 21.7 ng/mL). Esta variação cíclica está intimamente relacionada ao eixo hipotálamo-hipófise-gônadas (HPG), conectando homeostase energética à reprodução através de receptores específicos no hipotálamo. Níveis adequados de leptina são essenciais para manutenção de ciclos menstruais regulares e fertilidade, atuando como sinal permissivo para a secreção pulsátil de GnRH. A deficiência de leptina (amenorreia hipotalâmica em atletas ou anorexia nervosa) ou excesso (obesidade, síndrome dos ovários policísticos) podem causar disfunções reprodutivas significativas incluindo anovulação e infertilidade. Mulheres com SOP frequentemente apresentam hiperleptinemia (>11.58 ng/mL) associada à resistência insulínica (HOMA-IR >2.5), hiperandrogenismo (testosterona total >60 ng/dL) e anovulação crônica. Estudos recentes demonstram que a combinação de leptina elevada com hormônio antimülleriano (AMH) apresenta excelente acurácia diagnóstica para SOP (AUC 92.3%, sensibilidade 93.3%). Na transição menopausal e pós-menopausa, a leptina contribui para alterações metabólicas adversas, acúmulo de gordura visceral, resistência insulínica e aumento do risco cardiovascular e de síndrome metabólica. A avaliação dos níveis de leptina, especialmente quando interpretada em conjunto com IMC, fase do ciclo menstrual, perfil hormonal reprodutivo e marcadores metabólicos, auxilia no diagnóstico diferencial de distúrbios reprodutivos e metabólicos, e no planejamento de intervenções terapêuticas personalizadas focadas em modificação de estilo de vida e tratamento farmacológico quando indicado.',
    updated_at = NOW()
WHERE id = '6d18335b-09d8-41b6-a5d1-db49a4fa62fa';

-- Verify the update
DO $$
DECLARE
    v_clinical_length INTEGER;
BEGIN
    SELECT LENGTH(clinical_relevance)
    INTO v_clinical_length
    FROM score_items
    WHERE id = '6d18335b-09d8-41b6-a5d1-db49a4fa62fa';

    RAISE NOTICE 'Updated Clinical Relevance: % chars (target: 1500-2000)', v_clinical_length;

    IF v_clinical_length BETWEEN 1500 AND 2000 THEN
        RAISE NOTICE '✓ Clinical Relevance now meets target criteria';
    END IF;
END $$;

COMMIT;
