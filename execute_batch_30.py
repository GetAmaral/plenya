#!/usr/bin/env python3
"""
Batch 30 - Processar 30 Score Items de Exames Laboratoriais
Sistema Plenya EMR - Medicina Funcional Integrativa
"""

import requests
import json

API_URL = "http://localhost:3001/api/v1"
TOKEN = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiJmOTliMzk2OC0zNTI4LTQ2YTUtYWRkNi03ODYyNmQ1NGJmMTgiLCJlbWFpbCI6ImltcG9ydEBwbGVueWEuY29tIiwicm9sZSI6ImFkbWluIiwiaXNzIjoicGxlbnlhLWVtciIsImV4cCI6MTc2OTM5NDY4MiwiaWF0IjoxNzY5MzkzNzgyfQ.lxNqy-W_YTJjPXS12eDDgI0iaJdtMXs19o1ZL_xK0WQ"

headers = {
    "Authorization": f"Bearer {TOKEN}",
    "Content-Type": "application/json"
}

# Items 3-10 (1 e 2 já foram processados)
items_data = [
    {
        "id": "0899a7f5-e9d5-4adf-a825-4374ecc3292f",
        "name": "pO2 Venoso",
        "clinicalRelevance": "A pressão parcial de oxigênio venoso (PvO2) mede a tensão de oxigênio no sangue venoso, refletindo a oferta de oxigênio aos tecidos após extração metabólica. Valores normais situam-se entre 35-45 mmHg. A PvO2 é um indicador crucial de oxigenação tecidual e balanço entre oferta e consumo de oxigênio. Em medicina funcional, PvO2 <35 mmHg indica hipoxia tecidual, podendo resultar de baixo débito cardíaco, anemia severa, hipoxemia arterial ou aumento do consumo de oxigênio (febre, sepse, tireotoxicose). A hipoxia tecidual prolongada leva à produção de lactato (acidose lática), disfunção mitocondrial e dano celular. PvO2 reduzida é particularmente crítica em pacientes com insuficiência cardíaca, hipertensão pulmonar e doenças críticas. O monitoramento da saturação venosa de oxigênio (SvO2), derivada da PvO2, tem valor prognóstico em pacientes com hipertensão pulmonar (PvO2 <35 mmHg associa-se a pior prognóstico). A interpretação deve considerar o contexto clínico, incluindo níveis de hemoglobina, débito cardíaco e demanda metabólica.",
        "patientExplanation": "A PvO2 (pressão de oxigênio no sangue venoso) mede quanto oxigênio sobra no sangue depois que os tecidos do corpo o utilizaram. Valores normais são 35-45 mmHg. Níveis baixos (<35 mmHg) indicam que seus tecidos estão retirando muito oxigênio do sangue, o que pode acontecer por falta de oxigênio chegando aos órgãos (problema cardíaco, pulmonar ou anemia) ou porque o corpo está consumindo oxigênio demais (febre, infecções graves). Este exame ajuda a avaliar se seus órgãos estão recebendo oxigênio suficiente e é especialmente importante em doenças cardíacas e pulmonares graves.",
        "conduct": "**Interpretação:** PvO2 35-45 mmHg (normal); <35 mmHg (hipoxia tecidual) - investigar causa; >45 mmHg (aumento de oferta ou redução de consumo). **Para PvO2 baixa (<35 mmHg):** Avaliar hemoglobina (tratar anemia se Hb <10 g/dL com ferro, B12, folato); otimizar função cardíaca (avaliar fração de ejeção, considerar inotrópicos se insuficiência cardíaca); investigar hipoxemia arterial (gasometria arterial, oxigenoterapia se PaO2 <60 mmHg); tratar causas de hipermetabolismo (febre, sepse, tireotoxicose). **Suporte nutricional:** Coenzima Q10 100-300 mg/dia (função mitocondrial), L-carnitina 1-2 g/dia, antioxidantes (vitamina C 500-1.000 mg/dia, vitamina E 400 UI/dia). **Monitoramento:** Gasometria venosa seriada em pacientes críticos; SvO2 contínua em UTI se disponível; ecocardiograma e teste de caminhada de 6 minutos em hipertensão pulmonar."
    },
    {
        "id": "08afead7-7ce4-4435-93d0-1b0147ed4131",
        "name": "Potássio",
        "clinicalRelevance": "O potássio é o principal cátion intracelular, essencial para função neuromuscular, contratilidade cardíaca, equilíbrio ácido-base e regulação da pressão arterial. A faixa de referência é 3,5-5,0 mEq/L, mas em medicina funcional e cardiologia preventiva, valores ótimos situam-se entre 4,0-5,0 mEq/L, especialmente em pacientes com doenças cardiovasculares. Estudos recentes (POTCAST trial) demonstram que manter potássio em 4,5-5,0 mEq/L reduz arritmias ventriculares e mortalidade cardiovascular em 24% comparado ao manejo padrão. Hipocalemia (<3,5 mEq/L) aumenta risco de arritmias, fraqueza muscular, constipação, rabdomiólise e nefropatia crônica. Em pacientes com insuficiência cardíaca, potássio <4,1 mEq/L é preditor independente de morte súbita. Hipercalemia (>5,5 mEq/L) causa arritmias letais, especialmente em insuficiência renal ou uso de IECA/BRA. A reposição agressiva de potássio (alvos 4,5-5,0 mEq/L) beneficia pacientes cardiopatas, hipertensos e usuários de diuréticos.",
        "patientExplanation": "O potássio é um mineral essencial para o coração, músculos e nervos. Níveis normais estão entre 3,5-5,0 mEq/L, mas valores ideais para saúde cardiovascular são 4,0-5,0 mEq/L. Potássio baixo (<3,5 mEq/L) pode causar fraqueza muscular, cãibras, batimentos cardíacos irregulares, constipação e fadiga. Potássio alto (>5,5 mEq/L) é perigoso, podendo causar arritmias graves. Níveis na faixa alta-normal (4,5-5,0 mEq/L) protegem o coração contra arritmias e morte súbita, especialmente se você tem problemas cardíacos, pressão alta ou usa diuréticos. Alimentação rica em vegetais, frutas e legumes ajuda a manter bons níveis de potássio.",
        "conduct": "**Interpretação:** <3,5 mEq/L (hipocalemia) - repor urgente; 3,5-4,0 mEq/L (baixo-normal) - otimizar em cardiopatas; 4,0-5,0 mEq/L (ótimo para saúde cardiovascular); >5,5 mEq/L (hipercalemia) - investigar causa e tratar. **Para hipocalemia (<4,0 mEq/L em cardiopatas):** Suplementação oral de cloreto de potássio 20-40 mEq/dia (ajustar conforme resposta); aumentar alimentos ricos em potássio (banana, laranja, batata-doce, abacate, espinafre, feijão - alvo: 4.700 mg/dia); considerar antagonistas de aldosterona (espironolactona 25-50 mg/dia) se hipertensão ou insuficiência cardíaca. **Para hipercalemia (>5,5 mEq/L):** Suspender suplementos de potássio e IECA/BRA temporariamente; resina quelante (sorcal) se K >6,0 mEq/L; hemodiálise se K >7,0 mEq/L ou arritmias. **Monitoramento:** Dosar potássio semanalmente durante reposição; ECG se K <3,0 ou >6,0 mEq/L; ajustar meta para 4,5-5,0 mEq/L em alto risco cardiovascular; monitorar função renal (creatinina) concomitantemente."
    },
    {
        "id": "08c5d89b-976b-437b-8771-c6066d46d1db",
        "name": "pH Venoso",
        "clinicalRelevance": "O pH venoso mede a acidez ou alcalinidade do sangue venoso, refletindo o equilíbrio ácido-base sistêmico. A faixa normal é 7,32-7,42 (ligeiramente mais baixo que o pH arterial 7,35-7,45 devido ao CO2 tecidual). O pH venoso é um indicador sensível de distúrbios metabólicos e respiratórios, sendo menos invasivo que gasometria arterial. Em medicina funcional, pH <7,32 indica acidose (metabólica por excesso de ácidos/perda de bicarbonato ou respiratória por retenção de CO2), enquanto pH >7,42 sugere alcalose (metabólica por perda de ácidos/excesso de bicarbonato ou respiratória por hiperventilação). Acidose metabólica crônica de baixo grau está associada a perda de massa muscular, desmineralização óssea, resistência à insulina e progressão de doença renal crônica. A interpretação requer análise conjunta de pCO2, HCO3 e ânion gap. Distúrbios ácido-base afetam função enzimática, transporte de oxigênio, contratilidade cardíaca e função neurológica.",
        "patientExplanation": "O pH venoso mede se seu sangue está ácido ou alcalino (valores normais: 7,32-7,42). O corpo mantém este equilíbrio rigorosamente para que células e órgãos funcionem bem. pH baixo (<7,32) indica acidose, que pode ser causada por diabetes descontrolado, diarreia severa, insuficiência renal, problemas pulmonares ou exercício intenso. pH alto (>7,42) indica alcalose, que pode resultar de vômitos, uso excessivo de antiácidos ou hiperventilação. Desequilíbrios graves afetam respiração, coração, músculos e consciência. Este exame ajuda a diagnosticar e monitorar distúrbios metabólicos e respiratórios.",
        "conduct": "**Interpretação:** pH 7,32-7,42 (normal); <7,32 (acidose) - classificar como metabólica (HCO3 baixo) ou respiratória (pCO2 alto); >7,42 (alcalose) - classificar como metabólica (HCO3 alto) ou respiratória (pCO2 baixo). **Para acidose metabólica (pH <7,32, HCO3 <22 mEq/L):** Calcular ânion gap (Na - Cl - HCO3); investigar cetoacidose diabética (glicemia, cetonas), acidose láctica (lactato), insuficiência renal (ureia, creatinina), intoxicações; tratar causa base; bicarbonato de sódio IV se pH <7,1 e HCO3 <10 mEq/L. **Para alcalose metabólica (pH >7,42, HCO3 >28 mEq/L):** Investigar vômitos, uso de diuréticos, hiperaldosteronismo; repor cloreto de potássio; reduzir diuréticos se possível. **Dieta alcalinizante para acidose leve:** Aumentar vegetais, frutas, reduzir proteínas animais; citrato de potássio 20 mEq/dia. **Monitoramento:** Gasometria venosa seriada até normalização; tratar distúrbio de base; avaliar eletrólitos (K, Cl, Na) concomitantemente."
    }
]

print("Iniciando processamento do Batch 30 - Exames Laboratoriais")
print("=" * 70)

success_count = 0
error_count = 0

for idx, item in enumerate(items_data, start=3):
    try:
        print(f"\n[{idx}/30] Processando: {item['name']} ({item['id'][:8]}...)")

        payload = {
            "clinicalRelevance": item["clinicalRelevance"],
            "patientExplanation": item["patientExplanation"],
            "conduct": item["conduct"]
        }

        response = requests.put(
            f"{API_URL}/score-items/{item['id']}",
            headers=headers,
            json=payload,
            timeout=10
        )

        if response.status_code == 200:
            print(f"  ✅ Sucesso: {item['name']}")
            success_count += 1
        else:
            print(f"  ❌ Erro {response.status_code}: {item['name']}")
            print(f"     Response: {response.text[:100]}")
            error_count += 1

    except Exception as e:
        print(f"  ❌ Exceção: {item['name']} - {str(e)}")
        error_count += 1

print("\n" + "=" * 70)
print(f"Processamento concluído!")
print(f"  ✅ Sucessos: {success_count}")
print(f"  ❌ Erros: {error_count}")
print(f"  📊 Total: {len(items_data)}")
print("=" * 70)
