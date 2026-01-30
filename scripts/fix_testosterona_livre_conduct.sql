-- Fix conduct field length for Testosterona Livre - Mulheres Pré-Menopausa
-- Reduzir de 2759 para ~2450 caracteres

UPDATE score_items
SET
    conduct = 'INVESTIGAÇÃO DIAGNÓSTICA: Solicitar testosterona livre em mulheres com sinais clínicos de hiperandrogenismo (hirsutismo, acne refratária, alopecia androgenética, irregularidade menstrual) especialmente quando testosterona total está normal. Coletar preferencialmente fase folicular precoce (dias 3-5 ciclo) pela manhã (8-10h) por variação circadiana. Preferir testosterona livre calculada pela fórmula de Vermeulen (requer testosterona total por LC-MS/MS + SHBG) se diálise de equilíbrio não disponível. Solicitar painel completo: testosterona total e livre, SHBG, DHEA-S, androstenediona, 17-hidroxiprogesterona (excluir hiperplasia adrenal), LH/FSH (razão LH:FSH elevada sugere SOP), prolactina, TSH, glicemia jejum e HbA1c. Em casos de testosterona livre elevada: realizar ultrassom transvaginal para morfologia ovariana (critérios Rotterdam para SOP: ≥12 folículos 2-9mm ou volume >10mL), avaliar síndrome metabólica (circunferência abdominal, pressão arterial, lipidograma, glicemia), calcular HOMA-IR. Se testosterona livre >2-3x limite superior ou progressão rápida (<6 meses), investigar tumor secretor (ultrassom/RNM pélvica e adrenal). INTERPRETAÇÃO: Valores 0.3-1.9 pg/mL (1.0-6.6 pmol/L) normais; 2.0-3.5 pg/mL leve elevação (SOP provável); >4.0 pg/mL significativo (investigar tumor se >8.0 pg/mL). Considerar influência SHBG: obesidade reduz SHBG (aumenta testosterona livre desproporcionalmente), anticoncepcionais aumentam SHBG (reduzem testosterona livre). CONDUTA TERAPÊUTICA: SOP sem desejo gestacional: anticoncepcionais orais combinados primeira linha (etinilestradiol 20-35mcg + progestágeno antiandrogênico como drospirenona, ciproterona, dienogeste) reduzem testosterona livre por suprimir LH e aumentar SHBG. Se hirsutismo persiste após 6 meses: adicionar espironolactona 50-200mg/dia (monitorar potássio) ou finasterida 2.5-5mg/dia. Se resistência insulínica ou síndrome metabólica: metformina 1500-2000mg/dia melhora sensibilidade insulínica e pode reduzir testosterona livre. SOP com infertilidade: encaminhar reprodução assistida (citrato clomifeno, letrozol, gonadotrofinas). Modificações estilo vida fundamentais: perda ponderal 5-10% melhora hiperandrogenismo, exercício resistido e aeróbico, dieta baixo índice glicêmico. SEGUIMENTO: Reavaliar testosterona livre, glicemia e lipidograma após 3-6 meses tratamento. Monitorar anualmente função metabólica (síndrome metabólica em 40% das SOP), risco cardiovascular e saúde reprodutiva.',

    updated_at = CURRENT_TIMESTAMP

WHERE id = 'bb8e93f4-97f3-45de-8446-e138235953f0';

-- Verificar novo tamanho
SELECT
    'conduct' as campo,
    LENGTH(conduct) as caracteres,
    CASE
        WHEN LENGTH(conduct) BETWEEN 1500 AND 2500 THEN '✓ OK'
        ELSE '✗ Fora do range 1500-2500'
    END as validacao
FROM score_items
WHERE id = 'bb8e93f4-97f3-45de-8446-e138235953f0';
