#!/usr/bin/env python3
"""
Script para enriquecer 20 Score Items de HISTÓRICO DE DOENÇAS
Batch 2: Adrenalectomia, Alergias, Alterações em pelos, Amputação, Medicamentos
Metodologia: Medicina Funcional Integrativa - textos profundos e detalhados
"""

import requests
import json
import time
from typing import Dict, List

# Configuração
API_BASE_URL = "http://localhost:3001/api/v1"
EMAIL = "import@plenya.com"
PASSWORD = "Import@123456"

# IDs dos 20 items
ITEMS = [
    {"id": "d05a5afa-90d2-4622-8f1c-4a16ae4026af", "name": "Adrenalectomia", "group": "cirurgias"},
    {"id": "ed9f91de-a4e6-42ee-8b15-cda98a584c37", "name": "Adrenalectomia", "group": "cirurgias"},
    {"id": "838f53ac-c20d-4fd7-ba3e-7610ccd8e0fc", "name": "Adrenalectomia", "group": "cirurgias"},
    {"id": "1105bcdc-5c3d-4c8c-ab4b-b855a81f5a94", "name": "Alergias", "group": "alergias"},
    {"id": "7bbaabb1-1a43-4f31-8acc-136e5d63aded", "name": "Alergias", "group": "alergias"},
    {"id": "d54a2c69-ef58-4960-8314-6d8441981ca5", "name": "Alergias", "group": "alergias"},
    {"id": "a3db4c94-f6ff-435f-afde-83bc62a0563c", "name": "Alterações em pelos (excesso / falta)", "group": "pelos"},
    {"id": "2d92e515-db23-4862-9472-2f0a22062a61", "name": "Alterações em pelos (excesso / falta)", "group": "pelos"},
    {"id": "08306502-4181-4997-866e-2182f73ed8c1", "name": "Alterações em pelos (excesso / falta)", "group": "pelos"},
    {"id": "c56d1e62-1fb8-46fd-9155-1dbbec16fb79", "name": "Amputação de membro", "group": "amputacao"},
    {"id": "c37b8ae9-1d9b-4f06-b183-aa029ba5137c", "name": "Amputação de membro", "group": "amputacao"},
    {"id": "96080ff6-ed14-4632-a1de-097fd28e1da7", "name": "Amputação de membro", "group": "amputacao"},
    {"id": "4a5e1d32-6f84-4b2a-a2aa-e05183f5fe97", "name": "Analgésicos / Opioides / AINES / Relaxantes musculares", "group": "analgesicos"},
    {"id": "77890546-0a36-485b-a255-311ee0fc7bf4", "name": "Analgésicos / Opioides / AINES / Relaxantes musculares", "group": "analgesicos"},
    {"id": "64988fe3-f35b-4f89-bf54-1c8f149c05a8", "name": "Analgésicos / Opioides / AINES / Relaxantes musculares", "group": "analgesicos"},
    {"id": "01a6f6ca-068d-4de4-bb4d-0e886a128fb2", "name": "Ansiolíticos / Benzodiazepínicos / Sedativos", "group": "ansioliticos"},
    {"id": "cc86de0a-99fd-4d40-a4ea-908531ead927", "name": "Ansiolíticos / Benzodiazepínicos / Sedativos", "group": "ansioliticos"},
    {"id": "747fc220-3813-4595-a7c1-bf522ca6d79d", "name": "Ansiolíticos / Benzodiazepínicos / Sedativos", "group": "ansioliticos"},
    {"id": "fe579635-390e-4ce0-a970-351333e28cae", "name": "Anti-hipertensivos (IECA, BRA, betabloqueadores, BCC, diuréticos)", "group": "anti-hipertensivos"},
    {"id": "f6437fb6-32e6-4655-beb1-187d198da125", "name": "Anti-hipertensivos (IECA, BRA, betabloqueadores, BCC, diuréticos)", "group": "anti-hipertensivos"}
]

class DiseaseHistoryEnricher:
    def __init__(self):
        self.token = None
        self.session = requests.Session()

    def login(self) -> bool:
        """Faz login e obtém token JWT"""
        try:
            response = self.session.post(
                f"{API_BASE_URL}/auth/login",
                json={"email": EMAIL, "password": PASSWORD}
            )
            response.raise_for_status()
            data = response.json()
            self.token = data["accessToken"]
            self.session.headers.update({"Authorization": f"Bearer {self.token}"})
            print("✓ Login realizado com sucesso!")
            return True
        except Exception as e:
            print(f"✗ Erro no login: {e}")
            return False

    def update_score_item(self, item_id: str, clinical_relevance: str,
                         patient_explanation: str, conduct: str) -> bool:
        """Atualiza um score item com os textos clínicos"""
        try:
            payload = {
                "clinicalRelevance": clinical_relevance,
                "patientExplanation": patient_explanation,
                "conduct": conduct
            }

            response = self.session.put(
                f"{API_BASE_URL}/score-items/{item_id}",
                json=payload
            )
            response.raise_for_status()
            return True
        except Exception as e:
            print(f"✗ Erro ao atualizar item {item_id}: {e}")
            if hasattr(e, 'response') and e.response is not None:
                print(f"  Resposta: {e.response.text}")
            return False

    def get_content_for_group(self, group: str) -> Dict:
        """Retorna conteúdo clínico baseado no grupo do item"""

        content_map = {
            "cirurgias": {
                "clinical_relevance": """A adrenalectomia, remoção cirúrgica de uma ou ambas glândulas suprarrenais, representa intervenção de alta complexidade com implicações endócrinas, metabólicas e imunológicas profundas. Na medicina funcional integrativa, o histórico de adrenalectomia sinaliza necessidade crítica de reposição hormonal adequada e manejo especializado de múltiplas cascatas metabólicas interdependentes.

As glândulas suprarrenais são órgãos endócrinos vitais que produzem cortisol, aldosterona, andrógenos adrenais (DHEA, androstenediona) e catecolaminas (adrenalina, noradrenalina). A adrenalectomia unilateral preserva função parcial, mas pode resultar em insuficiência adrenal relativa sob estresse. A adrenalectomia bilateral causa insuficiência adrenal primária absoluta (doença de Addison iatrogênica), condição potencialmente fatal sem reposição hormonal adequada.

As indicações incluem tumores adrenais benignos ou malignos (adenomas funcionantes, feocromocitoma, carcinoma adrenocortical), síndrome de Cushing refratária, hiperaldosteronismo primário (síndrome de Conn) e, raramente, metástases. Na perspectiva funcional, a adrenalectomia impacta criticamente: metabolismo energético (via cortisol), equilíbrio hidroeletrolítico (via aldosterona), resposta ao estresse (eixo HPA), função imunológica, pressão arterial, metabolismo de carboidratos, proteínas e lipídios, densidade óssea e função sexual.

Pacientes pós-adrenalectomia bilateral requerem reposição vitalícia de glicocorticoides (hidrocortisona 15-25mg/dia dividida) e mineralocorticoides (fludrocortisona 0,05-0,2mg/dia). A dosagem deve ser ajustada conforme situações de estresse fisiológico (infecções, cirurgias, traumas) para prevenir crise addisoniana, emergência médica com mortalidade significativa. A medicina funcional enfatiza individualização da reposição hormonal, suporte nutricional específico (sódio, vitamina C, magnésio, ácido pantotênico) e monitoramento de parâmetros metabólicos, eletrólitos e densidade óssea.""",

                "patient_explanation": """A adrenalectomia é a cirurgia para remover uma ou ambas as glândulas suprarrenais, pequenos órgãos localizados acima dos rins que produzem hormônios essenciais para a vida. Se você passou por essa cirurgia, significa que havia um problema sério nessas glândulas (tumor, produção excessiva de hormônios ou outras condições).

As suprarrenais produzem três tipos principais de hormônios: cortisol (que controla o estresse, metabolismo e inflamação), aldosterona (que regula sal e pressão arterial) e hormônios sexuais (DHEA). Quando uma glândula é removida, a outra geralmente compensa parcialmente. Quando ambas são removidas, você precisa repor esses hormônios para o resto da vida através de medicamentos.

O mais importante é entender que o cortisol é essencial para a vida. Se você fez adrenalectomia bilateral (duas glândulas), você tem insuficiência adrenal e NUNCA pode parar os medicamentos. Além disso, em situações de estresse físico (infecções, cirurgias, acidentes), você precisa aumentar a dose do cortisol temporariamente - converse sempre com seu médico sobre isso. É fundamental usar uma pulseira ou cartão de identificação médica informando que você tem insuficiência adrenal, pois em emergências essa informação pode salvar sua vida.""",

                "conduct": """1. Avaliação endocrinológica especializada contínua (trimestral inicialmente, depois semestral)

2. Monitoramento laboratorial regular:
   - Sódio e potássio séricos (mensalmente nos primeiros meses, depois conforme estabilidade)
   - Renina plasmática e aldosterona (ajuste de fludrocortisona)
   - Cortisol sérico matinal (antes da dose, avaliar adequação da reposição)
   - ACTH (se adrenalectomia bilateral)
   - Glicemia e HbA1c (cortisol afeta metabolismo glicídico)
   - Pressão arterial (monitoramento domiciliar regular)
   - Densidade óssea (DEXA) anualmente (risco de osteoporose)

3. Reposição hormonal individualizada:
   - Se adrenalectomia bilateral:
     * Hidrocortisona 15-25mg/dia dividida (maior dose pela manhã, imitando ritmo circadiano)
     * Fludrocortisona 0,05-0,2mg/dia (conforme eletrólitos e pressão arterial)
     * DHEA 25-50mg/dia (mulheres, melhora qualidade de vida)
   - Se adrenalectomia unilateral:
     * Avaliar função da glândula remanescente
     * Reposição apenas se insuficiência documentada

4. Protocolo de estresse - CRÍTICO:
   - Aumentar hidrocortisona 2-3x em infecções, febres, cirurgias menores
   - Hidrocortisona injetável disponível em casa para emergências
   - Instruções claras para doença aguda ("sick day rules")
   - Comunicação com todos os médicos sobre necessidade de ajuste hormonal pré-procedimentos

5. Identificação médica obrigatória:
   - Pulseira ou cartão de identificação informando insuficiência adrenal
   - Prescrição de hidrocortisona injetável para emergências
   - Carta médica explicando a condição (para viagens, emergências)

6. Suporte nutricional direcionado:
   - Sódio: 3-5g/dia (essencial se deficiência de aldosterona)
   - Vitamina C: 500-1000mg/dia (glândulas adrenais têm alta concentração)
   - Magnésio: 400-600mg/dia (cofator em síntese hormonal)
   - Vitamina D: manter níveis >40ng/mL (proteção óssea)
   - Cálcio: 1200-1500mg/dia (prevenção de osteoporose)
   - Ácido pantotênico (B5): 250-500mg/dia (precursor de CoA)
   - Complexo B: suporte metabólico geral

7. Modulação do estilo de vida:
   - Evitar estresse excessivo (baixa reserva hormonal)
   - Exercícios regulares moderados (evitar overtraining)
   - Sono adequado (7-9h/noite)
   - Hidratação generosa (2-3L/dia)
   - Técnicas de manejo de estresse (mindfulness, respiração)
   - Evitar jejum prolongado (risco de hipoglicemia)

8. Prevenção de crise addisoniana:
   - Educação contínua do paciente e familiares
   - Plano de ação escrito para situações de emergência
   - Acesso rápido a atendimento médico
   - Uso de kit de emergência com hidrocortisona IM

9. Monitoramento de complicações:
   - Osteoporose (DEXA anual, considerar bisfosfonatos se necessário)
   - Diabetes secundário ao cortisol
   - Hipertensão (se excesso de fludrocortisona)
   - Hipotensão (se insuficiente reposição de mineralocorticoide)
   - Ganho de peso (ajuste fino da dose de cortisol)
   - Disfunção sexual (considerar reposição de DHEA)

10. Reavaliação multidisciplinar periódica:
    - Endocrinologista (principal)
    - Nutricionista funcional (otimização nutricional)
    - Psicólogo (adaptação à condição crônica)
    - Médico de família (coordenação de cuidados)"""
            },

            "alergias": {
                "clinical_relevance": """O histórico de alergias representa manifestação de hipersensibilidade imunológica a antígenos ambientais, alimentares ou farmacológicos, refletindo desequilíbrio na resposta imune predominantemente mediada por IgE (hipersensibilidade tipo I). Na medicina funcional integrativa, as alergias são compreendidas como expressão de disfunção imunológica sistêmica envolvendo desregulação Th1/Th2, hiperatividade mastocitária, permeabilidade de barreiras mucosas (intestino, respiratória, cutânea) e inflamação de baixo grau.

As manifestações clínicas abrangem amplo espectro: alergia respiratória (rinite alérgica, asma brônquica), cutânea (dermatite atópica, urticária), alimentar (desde sintomas gastrointestinais leves até anafilaxia), medicamentosa e a insetos (himenópteros). A prevalência global de doenças alérgicas tem aumentado significativamente nas últimas décadas, fenômeno associado à "hipótese da higiene", alterações da microbiota, exposição precoce a alérgenos, poluição ambiental e mudanças dietéticas.

A abordagem funcional reconhece que alergias frequentemente coexistem com outras condições atópicas (tríade atópica: rinite, asma, dermatite), sensibilidades alimentares não-IgE mediadas, síndrome de ativação mastocitária e deficiências imunológicas sutis. O histórico de alergias graves (anafilaxia) confere risco significativo de reações futuras potencialmente fatais, exigindo disponibilidade de epinefrina auto-injetável e plano de ação emergencial.

A patogênese alérgica envolve sensibilização inicial (exposição ao alérgeno com produção de IgE específica), seguida de reação imediata em reexposição (degranulação mastocitária com liberação de histamina, leucotrienos, prostaglandinas e citocinas). A fase tardia, 4-12 horas após, envolve infiltração eosinofílica e inflamação persistente. A medicina funcional investiga fatores predisponentes modificáveis: disbiose intestinal, deficiência de vitamina D, estresse oxidativo, exposição a toxinas ambientais, deficiências de nutrientes imunomoduladores (ômega-3, zinco, quercetina) e hiperpermeabilidade intestinal (leaky gut).""",

                "patient_explanation": """Ter histórico de alergias significa que seu sistema imunológico reage de forma exagerada a substâncias normalmente inofensivas (pólen, ácaros, alimentos, medicamentos, pelos de animais). Quando você entra em contato com essas substâncias, seu corpo libera histamina e outras substâncias que causam os sintomas alérgicos: espirros, coceira, coriza, falta de ar, urticária ou, em casos graves, anafilaxia.

Na medicina funcional, entendemos que alergias não são apenas "azar genético" - há fatores que podemos modificar para reduzir a intensidade e frequência das reações. Muitas vezes, alergias estão relacionadas a desequilíbrios no intestino (onde fica 70% do sistema imunológico), deficiência de vitamina D, inflamação crônica ou exposição excessiva a toxinas ambientais.

O objetivo não é apenas controlar sintomas com anti-histamínicos (embora sejam importantes quando necessário), mas também trabalhar para "acalmar" o sistema imunológico através de: alimentação anti-inflamatória, fortalecimento da saúde intestinal, suplementação estratégica (vitamina D, ômega-3, quercetina), redução de carga tóxica ambiental e manejo de estresse. Muitas pessoas conseguem redução significativa dos sintomas alérgicos com essa abordagem integrativa.""",

                "conduct": """1. Avaliação alergológica completa:
   - História detalhada de exposições e reações
   - Identificação de alérgenos específicos
   - Testes cutâneos (prick test) ou IgE específica sérica
   - Avaliação de gravidade e risco de anafilaxia
   - Pesquisa de comorbidades atópicas (asma, dermatite, rinoconjuntivite)

2. Investigação de fatores predisponentes funcionais:
   - Avaliação da microbiota intestinal (disbiose correlaciona com atopia)
   - Vitamina D sérica (níveis <30ng/mL associados a alergias)
   - Marcadores de permeabilidade intestinal (zonulina, se disponível)
   - IgG alimentares (sensibilidades não-IgE podem amplificar inflamação)
   - Marcadores inflamatórios (PCR-us, histamina, triptase)
   - Status de micronutrientes (zinco, selênio, magnésio)

3. Estratégia de evitação de alérgenos:
   - Identificação clara dos gatilhos
   - Controle ambiental (ácaros, fungos, pólens)
   - Leitura rigorosa de rótulos (alergias alimentares)
   - Evitação de medicamentos alergênicos
   - Uso de filtros HEPA (ambiente doméstico)
   - Controle de umidade (<50%)

4. Plano de emergência (se história de anafilaxia ou alto risco):
   - Prescrição de epinefrina auto-injetável (EpiPen ou similar)
   - Treinamento do paciente e familiares no uso
   - Pulseira de identificação médica
   - Plano de ação escrito
   - Anti-histamínicos e corticoides de resgate

5. Tratamento farmacológico convencional:
   - Anti-histamínicos (2ª geração - loratadina, cetirizina, desloratadina)
   - Corticosteroides tópicos (rinite, dermatite)
   - Broncodilatadores (se asma associada)
   - Imunoterapia alérgeno-específica (dessensibilização) quando apropriado
   - Anti-IgE (omalizumabe) em casos graves refratários

6. Modulação imunológica funcional:
   - Probióticos específicos (Lactobacillus rhamnosus GG, Bifidobacterium lactis)
     * 10-50 bilhões UFC/dia
     * Evidências em prevenção e redução de sintomas atópicos
   - Vitamina D (2.000-5.000 UI/dia, alvo >40ng/mL)
     * Imunomodulação, redução Th2, melhora barreira epitelial
   - Ômega-3 (2-3g/dia EPA+DHA)
     * Efeito anti-inflamatório, redução de leucotrienos
   - Quercetina (500-1000mg/dia)
     * Estabilizador natural de mastócitos, anti-histamínico natural
   - N-acetilcisteína (600-1200mg/dia)
     * Antioxidante, mucolítico, suporte de glutationa
   - Vitamina C (1-2g/dia)
     * Anti-histamínico natural, antioxidante

7. Suporte da barreira intestinal:
   - L-glutamina (5-10g/dia)
   - Zinco carnosina (75-150mg/dia)
   - Colostro bovino (quantidade conforme produto)
   - Enzimas digestivas (se má digestão contribui para sensibilidades)
   - Prebióticos (fibras fermentáveis - alimento para microbiota)

8. Intervenção nutricional anti-inflamatória:
   - Dieta de eliminação (se alergias alimentares)
   - Redução de alimentos histaminérgicos (quando relevante)
   - Aumento de alimentos anti-inflamatórios (vegetais coloridos, ômega-3, cúrcuma)
   - Hidratação adequada
   - Evitar aditivos alimentares, corantes, conservantes

9. Redução de carga alergênica ambiental:
   - Purificação do ar (filtros HEPA)
   - Remoção de tapetes e cortinas pesadas
   - Capas antiácaros em colchões e travesseiros
   - Lavagem de roupas de cama 60°C semanalmente
   - Controle de umidade
   - Evitar tabagismo (ativo e passivo)

10. Modulação do estilo de vida:
    - Manejo de estresse (amplifica respostas alérgicas)
    - Sono adequado (7-9h/noite)
    - Exercícios regulares moderados (imunomodulação)
    - Exposição solar controlada (vitamina D)
    - Técnicas de relaxamento (redução de cortisol e inflamação)

11. Monitoramento e reavaliação:
    - Acompanhamento alergológico periódico
    - Reavaliação de sensibilizações (podem mudar ao longo do tempo)
    - Ajuste de estratégias conforme resposta
    - Consideração de imunoterapia (dessensibilização) para casos selecionados"""
            },

            "pelos": {
                "clinical_relevance": """Alterações no padrão de crescimento piloso, manifestando-se como excesso (hirsutismo, hipertricose) ou redução (alopecia, hipotricose), representam manifestações clínicas de desequilíbrios endócrinos, metabólicos, nutricionais ou autoimunes com implicações diagnósticas e terapêuticas significativas. Na medicina funcional integrativa, essas alterações são reconhecidas como sinalizadores de disfunções sistêmicas subjacentes que exigem investigação abrangente.

O HIRSUTISMO, crescimento excessivo de pelos terminais em distribuição masculina em mulheres (face, tórax, abdome, região inguinal), afeta 5-10% das mulheres e reflete predominantemente hiperandrogenismo. As causas principais incluem síndrome dos ovários policísticos (SOP - 70-80% dos casos), hiperplasia adrenal congênita não clássica, tumores secretores de andrógenos (ovarianos ou adrenais), síndrome de Cushing, hiperprolactinemia, resistência à insulina e uso de medicamentos androgênicos. A avaliação funcional reconhece que o hiperandrogenismo frequentemente associa-se a resistência insulínica, inflamação de baixo grau, disbiose intestinal e estresse oxidativo.

A HIPERTRICOSE, aumento difuso da pilosidade não dependente de andrógenos, pode ser congênita ou adquirida, associada a medicamentos (minoxidil, ciclosporina, corticoides), hipotireoidismo, porfiria, anorexia nervosa ou síndromes genéticas raras. Diferencia-se do hirsutismo pela distribuição não androgênica e afeta ambos os sexos.

A ALOPECIA (perda de cabelo) apresenta etiologias diversas: alopecia androgenética (calvície de padrão masculino/feminino, associada a andrógenos e predisposição genética), alopecia areata (autoimune), eflúvio telógeno (queda difusa pós-estresse fisiológico, parto, cirurgias, deficiências nutricionais), deficiência de ferro, hipotireoidismo, deficiência de zinco/biotina, excesso de vitamina A, doenças autoimunes sistêmicas e quimioterapia. A medicina funcional investiga causas raiz: inflamação sistêmica, deficiências nutricionais (ferro, zinco, biotina, vitaminas do complexo B, vitamina D), desequilíbrios hormonais (tireoide, andrógenos, estrogênio), toxicidade por metais pesados e estresse oxidativo mitocondrial.""",

                "patient_explanation": """Alterações nos pelos podem se manifestar de duas formas principais: crescimento excessivo (especialmente em mulheres, em locais como rosto, peito, barriga) ou queda/falta de pelos (principalmente no couro cabeludo, mas também sobrancelhas e corpo).

Quando há crescimento excessivo de pelos em mulheres (hirsutismo), geralmente está relacionado a hormônios masculinos (andrógenos) elevados. As causas mais comuns são síndrome dos ovários policísticos (SOP), resistência à insulina, problemas na tireoide ou, mais raramente, tumores. Já a queda de cabelo pode ter várias causas: deficiências nutricionais (ferro, zinco, biotina, vitaminas), problemas na tireoide, estresse intenso, questões autoimunes ou simplesmente genética (calvície hereditária).

Na medicina funcional, não queremos apenas tratar o sintoma (tirar pelos com depilação ou usar minoxidil para queda), mas entender e corrigir a causa raiz. Isso geralmente envolve: investigação hormonal completa (andrógenos, tireoide, cortisol), avaliação de deficiências nutricionais, controle de resistência à insulina (se presente), redução de inflamação e suplementação direcionada. Muitas mulheres com SOP, por exemplo, conseguem melhora significativa do hirsutismo ao controlar a resistência à insulina através de alimentação adequada, exercícios e suplementos específicos.""",

                "conduct": """1. Avaliação clínica detalhada:
   - Padrão de crescimento/perda pilosa (distribuição, velocidade de início)
   - História menstrual (mulheres) - irregularidades sugerem SOP
   - Medicamentos em uso (muitos afetam pelos)
   - História familiar (genética)
   - Sintomas associados (acne, ganho de peso, fadiga)
   - Escore de Ferriman-Gallwey (hirsutismo)

2. Investigação laboratorial hormonal (hirsutismo/excesso de pelos):
   - Testosterona total e livre
   - DHEA-S (origem adrenal vs ovariana de andrógenos)
   - Androstenediona
   - 17-OH-progesterona (triagem para hiperplasia adrenal congênita)
   - Prolactina (hiperprolactinemia causa amenorreia e hirsutismo)
   - TSH, T4L, T3L (hipotireoidismo associa-se a alterações pilosas)
   - LH/FSH (relação elevada em SOP)
   - Glicemia, insulina, HOMA-IR (resistência insulínica)
   - Cortisol urinário 24h (se suspeita de Cushing)

3. Investigação laboratorial (alopecia/queda de cabelo):
   - Hemograma completo (anemia)
   - Ferritina (ideal >70ng/mL para saúde capilar ótima)
   - Ferro sérico, transferrina, saturação de transferrina
   - Zinco sérico
   - Vitamina D
   - Vitaminas B12 e folato
   - TSH, T4L, T3L (hipo e hipertireoidismo causam queda)
   - DHEA-S e testosterona (alopecia androgenética em mulheres)
   - ANA, anti-TPO (rastreamento autoimunidade)

4. Avaliação de imagem quando indicada:
   - Ultrassom pélvico (SOP - ovários policísticos)
   - Ultrassom ou TC de adrenais (se DHEA-S muito elevado - tumor?)
   - Tricoscopia/biópsia de couro cabeludo (casos selecionados de alopecia)

5. Tratamento específico para hirsutismo/excesso:
   A. Tratamento farmacológico (sob supervisão médica):
      - Anticoncepcionais orais combinados (1ª linha - reduzem andrógenos)
      - Espironolactona 50-200mg/dia (antiandrogênico)
      - Metformina 1500-2000mg/dia (se resistência insulínica/SOP)
      - Finasterida 2,5-5mg/dia (inibidor 5-alfa-redutase)

   B. Métodos de remoção local:
      - Depilação a laser (redução permanente)
      - Eletrólise
      - Creme de eflornitina (reduz crescimento facial)

6. Tratamento específico para alopecia:
   A. Alopecia androgenética:
      - Minoxidil tópico 5% (homens) ou 2% (mulheres)
      - Finasterida 1mg/dia (homens) - inibidor 5-alfa-redutase
      - Espironolactona (mulheres com hiperandrogenismo)
      - Transplante capilar (casos avançados)

   B. Alopecia areata (autoimune):
      - Corticoides tópicos ou intralesionais
      - Minoxidil
      - Imunoterapia tópica (difenciprona)
      - JAK inibidores (casos graves)

   C. Eflúvio telógeno:
      - Identificar e tratar causa (estresse, deficiências, medicamentos)
      - Geralmente reversível com correção da causa

7. Intervenções funcionais nutricionais e suplementação:

   A. Para hirsutismo/SOP:
      - Inositol (mio-inositol 2g + d-quiro-inositol 50mg, 2x/dia)
        * Melhora sensibilidade insulínica e reduz andrógenos em SOP
      - Berberina 500mg 3x/dia (sensibilidade insulínica)
      - NAC (N-acetilcisteína) 600mg 2x/dia (reduz andrógenos, melhora ovulação)
      - Ômega-3 2-3g/dia (anti-inflamatório)
      - Vitamina D (adequação >40ng/mL)
      - Zinco 30mg/dia (reduz conversão testosterona→DHT)
      - Saw palmetto (serenoa repens) 160mg 2x/dia (antiandrogênico natural)

   B. Para alopecia/queda de cabelo:
      - Ferro (se ferritina <70ng/mL - alvo 70-100ng/mL)
      - Zinco 30-50mg/dia (essencial para ciclo capilar)
      - Biotina 5-10mg/dia (fortalece cabelo, reduz quebra)
      - Vitaminas do complexo B (B6, B12, folato)
      - Vitamina D (adequação >40ng/mL)
      - Sílica (colágeno tipo II) 10-20mg/dia
      - L-cisteína e metionina (aminoácidos sulfurados - estrutura capilar)
      - Extrato de saw palmetto (alopecia androgenética)
      - Colágeno hidrolisado tipo I 10g/dia

   C. Antioxidantes e anti-inflamatórios (ambas condições):
      - Curcumina (500mg 2x/dia)
      - Resveratrol (200-500mg/dia)
      - Vitamina E (400 UI/dia)
      - Selênio 200mcg/dia

8. Intervenção nutricional:
   A. Para hirsutismo/SOP:
      - Dieta de baixa carga glicêmica (controle insulina)
      - Restrição de carboidratos refinados e açúcares
      - Proteínas de qualidade (1,2-1,5g/kg)
      - Gorduras saudáveis (ômega-3, abacate, azeite)
      - Fibras (>30g/dia - modulam excreção de andrógenos)
      - Chá de hortelã (spearmint) - propriedades antiandrogênicas

   B. Para alopecia:
      - Proteína adequada (cabelo é 95% queratina - proteína)
      - Alimentos ricos em ferro (carnes vermelhas, fígado, vegetais verde-escuros)
      - Biotina (ovos, nozes, sementes)
      - Zinco (ostras, carnes, sementes de abóbora)
      - Ômega-3 (peixes gordos, linhaça, chia)
      - Vitamina C (absorção de ferro)

9. Modulação do estilo de vida:
   - Controle de peso (redução de 5-10% melhora hirsutismo em SOP)
   - Exercícios regulares (melhora sensibilidade insulínica)
   - Manejo de estresse (reduz cortisol que agrava queda capilar)
   - Sono adequado (7-9h/noite)
   - Evitar penteados muito puxados (tração)
   - Cuidados capilares gentis (evitar calor excessivo, químicas agressivas)

10. Tratamento de condições subjacentes:
    - Otimização de hipotireoidismo
    - Controle de resistência insulínica/diabetes
    - Tratamento de deficiências nutricionais
    - Manejo de condições autoimunes
    - Suspensão ou substituição de medicamentos causadores

11. Monitoramento e reavaliação:
    - Acompanhamento endocrinológico (hirsutismo)
    - Dermatologia especializada (alopecia)
    - Reavaliação hormonal trimestral (primeiros 6 meses)
    - Reavaliação de micronutrientes (3-6 meses)
    - Documentação fotográfica (avaliar progressão/melhora)
    - Ajustes terapêuticos conforme resposta

**IMPORTANTE: Muitas alterações pilosas respondem lentamente ao tratamento (3-6 meses mínimo), pois o ciclo de crescimento capilar é longo. Paciência e aderência são essenciais para resultados.**"""
            },

            "amputacao": {
                "clinical_relevance": """A amputação de membro representa procedimento cirúrgico com profundas implicações físicas, metabólicas, psicológicas e funcionais. Na medicina funcional integrativa, o histórico de amputação sinaliza necessidade de abordagem multidimensional abrangendo reabilitação física, suporte metabólico-nutricional, manejo de dor neuropática (incluindo dor fantasma), prevenção de complicações do membro residual, otimização protética e suporte psicoemocional.

As indicações de amputação incluem trauma severo com destruição tecidual irreparável, doença vascular periférica (principalmente em diabetes mellitus - 85% das amputações não traumáticas), infecções severas (gangrena, osteomielite refratária), tumores ósseos ou de partes moles malignos, malformações congênitas e, raramente, controle de dor refratária. A localização anatômica (digital, transmetatarsiana, transtibial, transfemoral, desarticulação, etc.) determina potencial de reabilitação, gasto energético de deambulação e necessidade protética.

A amputação desencadeia múltiplas consequências sistêmicas: reorganização neuroplástica cortical (gerando dor fantasma em 60-80% dos casos), alterações biomecânicas com sobrecarga do membro contralateral e coluna vertebral, aumento do gasto energético de locomoção (20-60% maior dependendo do nível), risco de síndrome metabólica (redução de atividade física), alterações psicológicas (luto, depressão, ansiedade, transtorno de estresse pós-traumático) e complicações do coto (neuromas, contraturas, úlceras de pressão, infecções, dor residual).

A medicina funcional reconhece que pacientes amputados por causas vasculares (diabetes, aterosclerose) apresentam doença sistêmica com risco substancial de progressão (amputação contralateral, eventos cardiovasculares). A otimização metabólica intensiva (controle glicêmico rigoroso, revascularização quando indicada, cuidados podológicos preventivos) pode reduzir significativamente o risco de amputações subsequentes. O suporte nutricional adequado é fundamental para cicatrização do coto, prevenção de infecções, manutenção de massa muscular e densidade óssea.""",

                "patient_explanation": """A amputação é a remoção cirúrgica de parte ou totalidade de um membro (braço, perna, dedos). Se você passou por uma amputação, isso representa uma mudança profunda que afeta não apenas o aspecto físico, mas também o emocional, funcional e até metabólico do seu corpo.

As causas mais comuns são complicações do diabetes (falta de circulação e feridas que não cicatrizam), acidentes graves, infecções severas ou tumores. Após a amputação, muitas pessoas experimentam "dor fantasma" - uma sensação real de dor no membro que não existe mais, causada por reorganização no cérebro.

A reabilitação é fundamental e envolve várias etapas: cicatrização adequada do coto, fisioterapia, adaptação à prótese (quando indicada), fortalecimento muscular, reaprendizado de movimentos e suporte psicológico. Na medicina funcional, também focamos em: nutrição adequada para cicatrização, controle rigoroso do diabetes (se for a causa), prevenção de problemas no membro que restou, manejo da dor neuropática e suporte para saúde mental. O objetivo é recuperar máxima independência e qualidade de vida.""",

                "conduct": """1. Avaliação multidisciplinar coordenada:
   - Cirurgião vascular ou ortopedista (seguimento do coto)
   - Fisiatra (medicina física e reabilitação)
   - Fisioterapeuta especializado em amputados
   - Terapeuta ocupacional
   - Protesista (avaliação e adaptação de prótese)
   - Psicólogo/psiquiatra (suporte emocional, TEPT, depressão)
   - Nutricionista (suporte metabólico e cicatrização)
   - Especialista em dor (manejo de dor fantasma e neuropática)

2. Cuidados com o coto (membro residual):
   - Inspeção diária (especialmente diabéticos - visão, espelhos)
   - Higiene adequada (limpeza, secagem cuidadosa)
   - Hidratação da pele (evitar ressecamento e fissuras)
   - Enfaixamento compressivo (primeiras semanas - reduz edema, molda coto)
   - Prevenção de contraturas (posicionamento adequado, alongamentos)
   - Tratamento de neuromas dolorosos (infiltração, cirurgia se refratário)
   - Monitoramento de úlceras de pressão
   - Ajuste de prótese conforme mudanças no coto

3. Controle rigoroso da condição de base (se amputação vascular):
   - Diabetes: HbA1c <7% (idealmente <6,5%)
   - Pressão arterial <130/80 mmHg
   - LDL-colesterol <70mg/dL
   - Triglicerídeos <150mg/dL
   - Cessação absoluta de tabagismo
   - Antiagregação plaquetária (AAS, clopidogrel)
   - Estatina em dose adequada

4. Proteção do membro contralateral (amputação de MMII):
   - Avaliação vascular periférica (ITB, ultrassom Doppler)
   - Cuidados podológicos rigorosos
   - Inspeção diária dos pés
   - Calçados adequados (macios, sem costuras internas)
   - Tratamento precoce de calos, unhas encravadas, micoses
   - Hidratação diária dos pés
   - Controle de pressão plantar (palmilhas quando indicado)

5. Manejo de dor fantasma e neuropática:
   A. Farmacológico:
      - Gabapentina 900-3600mg/dia (dividida 3x)
      - Pregabalina 150-600mg/dia (dividida 2x)
      - Antidepressivos tricíclicos (amitriptilina 25-100mg noite)
      - Duloxetina 60-120mg/dia
      - Tramadol (analgésico com ação em dor neuropática)
      - Evitar opioides fortes (risco de dependência)

   B. Não farmacológico:
      - Terapia do espelho (Mirror Box Therapy - evidência em dor fantasma)
      - Dessensibilização gradual do coto
      - TENS (estimulação elétrica transcutânea)
      - Acupuntura
      - Realidade virtual
      - Mindfulness e técnicas de relaxamento
      - Biofeedback

6. Reabilitação física progressiva:
   - Fortalecimento do membro residual
   - Fortalecimento de core e membro contralateral
   - Treino de equilíbrio
   - Treino de transferências
   - Treino de marcha (com auxiliares inicialmente, depois prótese)
   - Condicionamento cardiovascular adaptado
   - Prevenção de contraturas (alongamentos diários)
   - Trabalho de propriocepção

7. Adaptação e treinamento protético:
   - Avaliação do potencial de protetização (nível, condição clínica)
   - Escolha de prótese adequada (nível de atividade K0-K4)
   - Tipos: mecânicas, microprocessadas, mioelét ricas
   - Período de adaptação gradual
   - Ajustes frequentes inicialmente
   - Treinamento de marcha com fisioterapeuta
   - Manutenção regular da prótese

8. Suporte nutricional direcionado:
   A. Para cicatrização e integridade do coto:
      - Proteína adequada (1,2-1,5g/kg - síntese de colágeno)
      - Vitamina C (500-1000mg/dia - síntese de colágeno)
      - Zinco (30-50mg/dia - cicatrização)
      - Vitamina A (10.000 UI/dia curto prazo - reepitelização)
      - Arginina (10-20g/dia - cicatrização, fluxo sanguíneo)

   B. Para saúde óssea e muscular:
      - Cálcio (1000-1200mg/dia)
      - Vitamina D (2.000-5.000 UI/dia, alvo >40ng/mL)
      - Magnésio (400-600mg/dia)
      - Vitamina K2 (100-200mcg/dia)
      - Proteína de qualidade (manutenção de massa muscular)

   C. Para função nervosa e dor neuropática:
      - Complexo B de alta potência (B1, B6, B12)
      - Ácido alfa-lipóico (600mg/dia - neuroproteção)
      - Acetil-L-carnitina (1-2g/dia - regeneração nervosa)
      - Ômega-3 (2-3g/dia EPA+DHA - neuroinflamação)

   D. Antioxidantes e anti-inflamatórios:
      - Curcumina (500mg 2x/dia)
      - Resveratrol (200-500mg/dia)
      - Vitamina E (400 UI/dia)
      - Selênio (200mcg/dia)

9. Intervenção nutricional para prevenção de síndrome metabólica:
   - Controle calórico (gasto energético pode reduzir por menor atividade)
   - Dieta anti-inflamatória mediterrânea
   - Controle de carboidratos (se diabético ou pré-diabético)
   - Proteína adequada (preservação de massa magra)
   - Hidratação (2-2,5L/dia)

10. Suporte psicoemocional:
    - Psicoterapia individual (processo de luto, aceitação)
    - Terapia cognitivo-comportamental (manejo de dor, ansiedade)
    - EMDR (se trauma/TEPT)
    - Grupos de suporte (compartilhar experiências)
    - Tratamento de depressão e ansiedade (farmacológico se indicado)
    - Suporte familiar (educação, envolvimento)

11. Prevenção de complicações do membro contralateral e sistêmicas:
    - Prevenção de sobrecarga (orteses, adequação de marcha)
    - Exercícios de alongamento e fortalecimento bilateral
    - Monitoramento de dor lombar e articular (adaptações biomecânicas)
    - Vigilância cardiovascular (risco aumentado)
    - Prevenção de novas amputações (cuidados vasculares rigorosos)

12. Adaptações ambientais e de estilo de vida:
    - Modificações domiciliares (barras, rampas, banheiro adaptado)
    - Cadeira de rodas (mobilidade quando sem prótese)
    - Adaptações veiculares (dirigir com controles manuais se MMII)
    - Retorno ao trabalho (adaptações ocupacionais)
    - Atividades recreacionais e esportivas adaptadas
    - Vida sexual (adaptações, comunicação com parceiro)

13. Monitoramento e reavaliação periódica:
    - Consultas cirúrgicas (3, 6, 12 meses e depois anuais)
    - Fisioterapia contínua (manutenção funcional)
    - Ajustes protéticos (conforme mudanças no coto)
    - Reavaliação de dor (ajustes de tratamento)
    - Monitoramento de saúde mental
    - Reavaliação laboratorial (diabetes, fatores de risco CV)

**IMPORTANTE: A amputação é um evento de vida transformador que exige paciência, persistência e abordagem multidisciplinar. A maioria dos amputados alcança independência funcional satisfatória com reabilitação adequada e suporte contínuo.**"""
            },

            "analgesicos": {
                "clinical_relevance": """O histórico de uso de analgésicos, opioides, anti-inflamatórios não esteroidais (AINEs) e relaxantes musculares representa informação clínica de alta relevância, sinalizando presença de dor crônica ou recorrente, condições musculoesqueléticas, risco de eventos adversos medicamentosos e, no caso de opioides, potencial para dependência e síndrome de abstinência. Na medicina funcional integrativa, o uso crônico dessas medicações é compreendido como manifestação de processos álgicos que frequentemente possuem causas-raiz modificáveis (inflamação sistêmica, disfunções musculoesqueléticas, sensibilização central, deficiências nutricionais, estresse oxidativo).

Os AINEs (ibuprofeno, naproxeno, diclofenaco, cetoprofeno, celecoxibe) atuam inibindo ciclooxigenases (COX-1/COX-2), reduzindo prostaglandinas inflamatórias. Embora eficazes para dor aguda, o uso crônico associa-se a gastropatia (úlceras, sangramento gastrointestinal - risco 3-5x aumentado), nefropatia (insuficiência renal, nefrite intersticial), eventos cardiovasculares (especialmente COX-2 seletivos e AINEs não seletivos em altas doses), hipertensão arterial e comprometimento de cicatrização óssea e tendínea.

Os OPIOIDES (codeína, tramadol, morfina, oxicodona, metadona, fentanil) atuam em receptores mu, delta e kappa, modulando transmissão nociceptiva. São eficazes para dor moderada a severa aguda (pós-operatória, trauma, câncer), mas o uso crônico em dor não oncológica é controverso devido a: tolerância progressiva (necessidade de doses crescentes), hiperalgesia induzida por opioides (paradoxalmente aumentam sensibilidade à dor), dependência física, síndrome de abstinência, constipação severa, disfunção endócrina (hipogonadismo hipogonadotrófico, hipocortisolismo), imunossupressão, depressão respiratória e risco de overdose fatal. A epidemia de opioides nos EUA demonstra os perigos da prescrição indiscriminada.

Os RELAXANTES MUSCULARES (ciclobenzaprina, carisoprodol, baclofeno, tizanidina) atuam centralmente reduzindo espasticidade e tensão muscular. Associam-se a sedação, tontura, dependência (especialmente carisoprodol) e risco de quedas em idosos. Na perspectiva funcional, espasmos musculares frequentemente refletem desequilíbrios biomecânicos, deficiência de magnésio, desidratação ou disfunções neurológicas subjacentes que merecem investigação.""",

                "patient_explanation": """Se você usa ou usou regularmente analgésicos, anti-inflamatórios, opioides ou relaxantes musculares, isso indica que você experimenta dor recorrente ou crônica. Esses medicamentos são importantes para controlar a dor, mas o uso prolongado pode trazer problemas.

Anti-inflamatórios (ibuprofeno, diclofenaco) podem causar úlceras no estômago, problemas nos rins e aumentar pressão arterial com uso prolongado. Opioides (codeína, tramadol, morfina) são potentes para dor intensa, mas criam tolerância (você precisa de doses maiores para o mesmo efeito), podem viciar e causar constipação severa. Relaxantes musculares causam sonolência e tontura.

Na medicina funcional, queremos entender POR QUE você tem dor crônica e trabalhar nas causas raiz: pode ser inflamação sistêmica, deficiências nutricionais (magnésio, vitamina D), problemas posturais, sensibilização do sistema nervoso ou condições autoimunes. O objetivo é reduzir gradualmente a dependência de medicações através de: alimentação anti-inflamatória, suplementação direcionada, fisioterapia, técnicas de manejo de dor (mindfulness, acupuntura), correção postural e exercícios terapêuticos.""",

                "conduct": """1. Avaliação detalhada do uso de analgésicos:
   - Tipo de medicação (AINE, opioide, relaxante muscular)
   - Dose e frequência (diária, semanal, conforme necessidade)
   - Duração do uso (agudo <3 meses, crônico >3 meses)
   - Indicação original (dor aguda, crônica, pós-operatória)
   - Eficácia percebida (controle adequado da dor?)
   - Escalada de dose (tolerância)?
   - Tentativas prévias de suspensão
   - Sintomas de abstinência (especialmente opioides)

2. Caracterização da síndrome álgica:
   - Localização, intensidade (escala 0-10), qualidade da dor
   - Fatores desencadeantes e atenuantes
   - Padrão temporal (constante, intermitente)
   - Impacto funcional (AVDs, sono, trabalho)
   - Investigação de bandeiras vermelhas (dor neoplásica, infecciosa, fraturas)

3. Avaliação de complicações do uso crônico:

   A. AINEs:
      - Endoscopia digestiva alta (se uso >3 meses, sintomas dispépticos, >60 anos)
      - Função renal (creatinina, TFG, EAS)
      - Pressão arterial (monitoramento)
      - Risco cardiovascular (eventos coronarianos, AVC)

   B. Opioides:
      - Sinais de dependência (critérios DSM-5)
      - Constipação (geralmente presente)
      - Função hormonal: testosterona (homens), estradiol/LH/FSH (mulheres)
        * Opioides suprimem eixo hipotálamo-hipófise-gonadal
      - Função adrenal (cortisol se suspeita)
      - Densidade óssea (hipogonadismo→osteoporose)
      - Depressão respiratória (oximetria, polissonografia se apneia)
      - Hiperalgesia induzida por opioides (piora paradoxal da dor)

   C. Relaxantes musculares:
      - Sedação excessiva, risco de quedas (especialmente idosos)
      - Função hepática (alguns metabolizados hepaticamente)

4. Investigação de causas subjacentes da dor:
   - Marcadores inflamatórios (PCR-us, VHS)
   - Vitamina D (níveis <30ng/mL associados a dor crônica)
   - Magnésio (deficiência causa espasmos, dor muscular)
   - Vitamina B12 (neuropatia periférica)
   - TSH (hipotireoidismo causa mialgias)
   - Fator reumatoide, anti-CCP (artrite reumatoide)
   - ANA (doenças autoimunes)
   - Ácido úrico (gota)
   - Ressonância ou tomografia (conforme localização da dor)

5. Estratégia de redução gradual (tapering):

   **CRÍTICO: NUNCA suspender opioides ou benzodiazepínicos abruptamente - risco de síndrome de abstinência severa**

   A. AINEs:
      - Redução de dose gradual (25% a cada 1-2 semanas)
      - Mudança para uso "conforme necessidade" vs diário
      - Proteção gástrica se uso contínuo necessário (IBP, misoprostol)

   B. Opioides (SEMPRE sob supervisão médica):
      - Redução lenta: 10% da dose a cada 1-2 semanas (ou mais devagar se sintomas)
      - Monitoramento de síndrome de abstinência:
        * Precoce: ansiedade, sudorese, lacrimejamento, rinorreia, insônia
        * Tardia: náuseas, vômitos, diarreia, dores musculares, piloereção
      - Suporte farmacológico da abstinência:
        * Clonidina (sintomas autonômicos)
        * Loperamida (diarreia)
        * Antieméticos
        * AINEs (mialgias)
      - Considerar rotação de opioides ou buprenorfina (casos complexos)

   C. Relaxantes musculares:
      - Redução gradual de dose
      - Substituição por alternativas (magnésio, alongamentos, fisioterapia)

6. Estratégias não farmacológicas de manejo de dor:
   - Fisioterapia (exercícios terapêuticos, fortalecimento, alongamento)
   - Terapia manual (osteopatia, quiropraxia)
   - Acupuntura (evidência em dor lombar, osteoartrite, cefaleia)
   - TENS (estimulação elétrica transcutânea)
   - Termoterapia (calor) ou crioterapia (frio)
   - Mindfulness e meditação (redução de catastrofização)
   - Terapia cognitivo-comportamental (manejo de dor crônica)
   - Biofeedback
   - Yoga terapêutico
   - Hidroterapia/natação
   - Massagem terapêutica

7. Intervenções nutricionais anti-inflamatórias e analgésicas:
   - Dieta anti-inflamatória (mediterrânea, low-carb se obesidade)
   - Eliminação de alimentos pró-inflamatórios (açúcares, óleos vegetais refinados, trans)
   - Aumento de ômega-3 (peixes gordos, linhaça, chia)
   - Especiarias anti-inflamatórias (cúrcuma, gengibre)
   - Alimentos ricos em antioxidantes (frutas vermelhas, vegetais coloridos)
   - Hidratação adequada (2-2,5L/dia)

8. Suplementação funcional com propriedades analgésicas/anti-inflamatórias:
   - Ômega-3 (2-4g/dia EPA+DHA) - potente anti-inflamatório natural
   - Curcumina (500-1000mg 2x/dia, forma biodisponível) - analgésico natural
   - Boswellia serrata (300-500mg 3x/dia) - anti-inflamatório, dor articular
   - Gengibre (1-2g/dia) - COX-2 inibidor natural
   - MSM (metilsulfonilmetano) 2-3g/dia - dor articular
   - Glucosamina + condroitina (1500mg + 1200mg/dia) - osteoartrite
   - Colágeno tipo II (40mg/dia) - dor articular
   - Magnésio (400-600mg/dia) - relaxamento muscular, dor neuropática
   - Vitamina D (adequação >40ng/mL) - dor musculoesquelética
   - Ácido alfa-lipóico (600mg/dia) - dor neuropática
   - Acetil-L-carnitina (1-2g/dia) - dor neuropática
   - S-adenosilmetionina (SAMe) 400-800mg/dia - osteoartrite, fibromialgia
   - CBD (canabidiol) 20-40mg/dia - dor neuropática, inflamatória (onde legal)
   - Complexo B de alta potência (neuropatias)

9. Alternativas farmacológicas não opioides para dor crônica:
   - Paracetamol (acetaminofeno) 3-4g/dia - analgésico puro, menos efeitos GI
   - Gabapentina/pregabalina (dor neuropática)
   - Antidepressivos (amitriptilina, duloxetina) - dor neuropática, fibromialgia
   - Cremes tópicos (diclofenaco, capsaicina, lidocaína)
   - Infiltrações (corticoides, ácido hialurônico - articular)
   - Bloqueios anestésicos (clínica de dor)

10. Correção de deficiências que perpetuam dor:
    - Vitamina D (suplementação agressiva se <20ng/mL)
    - Magnésio (reposição oral, até tolerância intestinal)
    - Vitamina B12 (se deficiente - IM se grave)
    - Ferro (se anemia contribui para fadiga e piora percepção de dor)

11. Abordagens intervencionistas (casos refratários):
    - Bloqueios nervosos
    - Radiofrequência
    - Neuromodulação (estimulação medular)
    - Bombas de infusão intratecal
    - Procedimentos cirúrgicos definitivos (se indicação clara)

12. Suporte psicoemocional:
    - Avaliação de depressão e ansiedade (comórbidas em dor crônica)
    - Terapia cognitivo-comportamental para dor
    - Mindfulness-based stress reduction (MBSR)
    - Grupos de suporte
    - Tratamento de insônia (frequente em dor crônica)

13. Monitoramento e reavaliação:
    - Consultas regulares durante período de redução de medicação
    - Escala de dor (acompanhar tendência)
    - Funcionalidade (capacidade de realizar AVDs)
    - Qualidade de vida (questionários validados)
    - Sinais de abstinência
    - Reavaliação laboratorial (função renal, hepática, hormonal)
    - Ajustes terapêuticos conforme resposta

14. Educação do paciente:
    - Compreensão da natureza da dor crônica (neuroplasticidade)
    - Objetivos realistas (melhora funcional vs cura completa)
    - Papel ativo no manejo (não passivo/dependente de medicação)
    - Importância de abordagem multimodal
    - Prevenção de recaídas

**IMPORTANTE: A redução de opioides em uso crônico deve ser SEMPRE gradual e supervisionada por médico experiente. Síndrome de abstinência pode ser perigosa e desconfortável. Paciência e suporte são essenciais.**"""
            },

            "ansioliticos": {
                "clinical_relevance": """O histórico de uso de ansiolíticos, benzodiazepínicos e sedativos representa informação clínica crítica, sinalizando presença de transtornos de ansiedade, insônia ou outras condições psiquiátricas, além de risco significativo de dependência física e psicológica, tolerância, síndrome de abstinência potencialmente grave e comprometimento cognitivo. Na medicina funcional integrativa, o uso crônico de benzodiazepínicos é reconhecido como problemático e frequentemente perpetua os sintomas que pretende tratar, além de mascarar disfunções subjacentes (desequilíbrios de neurotransmissores, eixo HPA disfuncional, deficiências nutricionais, estresse oxidativo).

Os benzodiazepínicos (BZDs - diazepam, clonazepam, alprazolam, lorazepam, bromazepam) atuam como agonistas do receptor GABA-A, potencializando a neurotransmissão inibitória GABAérgica. São eficazes para ansiedade aguda, crises de pânico, insônia transitória, abstinência alcoólica e controle de convulsões. Entretanto, o uso crônico (>4 semanas) está associado a: desenvolvimento rápido de tolerância (necessidade de doses progressivamente maiores), dependência física (síndrome de abstinência severa com risco de convulsões e delirium), dependência psicológica, comprometimento cognitivo (memória, atenção, função executiva), sedação, risco de quedas e fraturas (especialmente idosos), interações medicamentosas (potencialização com álcool, opioides, antidepressivos), depressão respiratória (doses altas ou combinações), demência (evidências controversas mas preocupantes) e piora paradoxal de ansiedade e insônia (ansiedade de rebote).

Os chamados "Z-drugs" (zolpidem, zopiclona, eszopiclona), embora quimicamente distintos, compartilham mecanismo GABAérgico e riscos similares, incluindo amnésia anterógrada, sonambulismo complexo (dirigir dormindo, comer) e dependência. Outros sedativos incluem barbitúricos (uso atual muito limitado devido a toxicidade) e anti-histamínicos sedantes (difenidramina, doxilamina) com perfil de segurança melhor mas eficácia limitada.

A síndrome de abstinência de benzodiazepínicos pode ser SEVERA e potencialmente fatal, incluindo: ansiedade e pânico rebote intensos, insônia grave, tremores, sudorese, taquicardia, hipertensão, náuseas, hipersensibilidade sensorial (luz, som), despersonalização, desrealização, convulsões (especialmente com BZDs de meia-vida curta como alprazolam) e, raramente, delirium. A medicina funcional enfatiza abordagem integrativa para ansiedade e insônia, investigando causas-raiz (hipoglicemia, deficiência de magnésio e vitaminas do complexo B, disfunção tireoidiana, excesso de cafeína, apneia do sono, estresse crônico, trauma psicológico) e implementando estratégias não farmacológicas (terapia cognitivo-comportamental, mindfulness, suplementação GABA-érgica natural, adaptógenos, higiene do sono).""",

                "patient_explanation": """Se você usa ou usou ansiolíticos (especialmente benzodiazepínicos como diazepam, clonazepam, alprazolam), isso indica que você sofre com ansiedade, pânico ou insônia. Esses medicamentos são muito eficazes no curto prazo para "desligar" a ansiedade ou ajudar a dormir, mas têm um problema sério: criam dependência rapidamente.

Dependência significa duas coisas: 1) Seu corpo se acostuma e você precisa de doses maiores para o mesmo efeito (tolerância), e 2) Se você parar de repente, pode ter sintomas graves de abstinência - ansiedade intensa, insônia pior que antes, tremores, e em casos graves até convulsões. Por isso, NUNCA deve parar esses medicamentos de uma hora para outra.

Na medicina funcional, investigamos POR QUE você tem ansiedade ou insônia: pode ser deficiência de magnésio (mineral calmante), vitaminas do complexo B, problemas na tireoide, excesso de cafeína, açúcar no sangue instável, apneia do sono ou estresse crônico não tratado. O objetivo é reduzir GRADUALMENTE e com segurança o medicamento enquanto implementamos alternativas: terapia cognitivo-comportamental, técnicas de respiração, exercícios, suplementos calmantes naturais (magnésio, L-teanina, ashwagandha), melhora do sono e tratamento das causas raiz.""",

                "conduct": """1. Avaliação detalhada do uso de benzodiazepínicos/ansiolíticos:
   - Medicação específica (tipo, dose, frequência)
   - Duração do uso (<4 semanas = curto prazo; >4 semanas = dependência provável)
   - Indicação original (ansiedade, pânico, insônia, convulsões, abstinência alcoólica)
   - Prescritor original (psiquiatra, clínico geral, neurologista)
   - Padrão de uso (regular vs conforme necessidade)
   - Escalada de dose (tolerância)
   - Tentativas prévias de suspensão e resultados
   - Sintomas de abstinência prévia
   - Uso concomitante de álcool ou outras substâncias

2. Avaliação do transtorno psiquiátrico subjacente:
   - Caracterização do transtorno de ansiedade:
     * Ansiedade generalizada (TAG)
     * Transtorno do pânico
     * Fobia social
     * TEPT (transtorno de estresse pós-traumático)
     * Transtorno obsessivo-compulsivo
   - Insônia (primária vs secundária)
   - Comorbidades psiquiátricas (depressão, bipolaridade)
   - Gravidade dos sintomas (escalas validadas - GAD-7, BAI)
   - Impacto funcional (trabalho, relacionamentos, AVDs)
   - Fatores estressores psicossociais

3. Avaliação de complicações do uso crônico:
   - Função cognitiva (memória, atenção, concentração)
     * Testes neuropsicológicos se comprometimento significativo
   - Quedas e fraturas (especialmente idosos)
   - Acidentes (direção veicular sob efeito)
   - Interações medicamentosas perigosas
   - Depressão respiratória (especialmente se uso com opioides, álcool)
   - Critérios de dependência (DSM-5):
     * Uso em quantidades maiores ou por mais tempo que pretendido
     * Desejo persistente ou esforços mal-sucedidos para reduzir
     * Muito tempo gasto obtendo, usando ou recuperando-se
     * Fissura/desejo intenso
     * Tolerância (necessidade de doses maiores)
     * Abstinência (sintomas quando suspende)

4. Investigação funcional de causas subjacentes de ansiedade/insônia:
   - Função tireoidiana (hiper e hipotireoidismo causam ansiedade)
     * TSH, T4L, T3L, anti-TPO
   - Glicemia e insulina (hipoglicemia reativa causa ansiedade)
   - Magnésio sérico e eritrocitário (deficiência→ansiedade, insônia)
   - Vitaminas do complexo B (B6, B12, folato - cofatores em síntese de neurotransmissores)
   - Vitamina D (níveis baixos associados a ansiedade e depressão)
   - Cortisol (ritmo circadiano - saliva 4 pontos ou urinário 24h)
     * Cortisol elevado à noite impede sono
   - Ferro e ferritina (anemia contribui para ansiedade)
   - Neurotransmissores urinários (dopamina, serotonina, GABA, glutamato - se disponível)
   - Marcadores inflamatórios (inflamação sistêmica afeta humor)
   - Polissonografia (se suspeita de apneia do sono causando insônia)

5. Protocolo de redução gradual (tapering) - **CRÍTICO**:

   **NUNCA suspender benzodiazepínicos abruptamente - risco de convulsões e síndrome de abstinência severa**

   A. Princípios gerais:
      - Redução muito lenta (mais lento que opioides)
      - Velocidade: 10% da dose a cada 1-2 semanas (ou mais devagar se sintomas)
      - Alguns casos exigem 6-12 meses para retirada segura
      - Monitoramento médico/psiquiátrico regular
      - Suporte psicológico durante processo

   B. Estratégia de conversão e redução:
      - Converter para BZD de meia-vida longa (diazepam) se uso de meia-vida curta (alprazolam)
        * Alprazolam 0,5mg = Diazepam 5mg
        * Clonazepam 0,5mg = Diazepam 10mg
        * Lorazepam 1mg = Diazepam 10mg
      - Dividir dose total em 2-3 tomadas/dia (evitar picos e vales)
      - Reduzir 10% a cada 1-2 semanas (da dose ORIGINAL, não da dose atual)
      - Pausar redução se sintomas de abstinência intensos
      - Últimos 25% da dose são os mais difíceis (reduzir mais devagar)

   C. Sintomas de abstinência (iniciar 1-4 dias após redução/suspensão):
      - Precoces: ansiedade rebote, insônia, irritabilidade, inquietação
      - Intermediários: tremores, sudorese, taquicardia, náusea
      - Graves: convulsões (especialmente dias 3-8), delirium, alucinações
      - Prolongados (semanas-meses): ansiedade, insônia, despersonalização, hipersensibilidade sensorial

   D. Manejo de sintomas de abstinência:
      - Betabloqueadores (propranolol) - sintomas autonômicos (tremor, taquicardia)
      - Anti-histamínicos (hidroxizina) - ansiedade leve, insônia
      - Anticonvulsivantes (carbamazepina, valproato) - prevenção de convulsões em casos de alto risco
      - Antidepressivos (ISRS/IRSN) - tratamento do transtorno de base
      - Suporte sintomático (antieméticos, analgésicos)

6. Tratamento do transtorno de ansiedade/insônia subjacente:

   A. Farmacológico não benzodiazepínico:
      - Ansiedade:
        * ISRS (sertralina, escitalopram) - 1ª linha para TAG, pânico, fobia social
        * IRSN (venlafaxina, duloxetina) - TAG
        * Buspirona 15-60mg/dia - ansiolítico não GABAérgico, sem dependência
        * Pregabalina 150-600mg/dia - TAG (sem potencial de abuso vs BZD)
        * Betabloqueadores (propranolol) - sintomas físicos de ansiedade
        * Antidepressivos tricíclicos (clomipramina) - TOC, pânico

      - Insônia:
        * Trazodona 50-200mg - sedação, sem dependência
        * Mirtazapina 7,5-15mg - sedação, ganho de apetite
        * Doxepina 3-6mg - antagonista H1, aprovado para insônia
        * Melatonina 2-10mg (1-2h antes de dormir)
        * Ramelteon - agonista de receptores de melatonina
        * Suvorexant - antagonista de orexina (caro)

   B. Psicoterapia (ESSENCIAL - eficácia igual ou superior a medicamentos):
      - Terapia cognitivo-comportamental para ansiedade (TCC-A)
      - Terapia cognitivo-comportamental para insônia (TCC-I) - padrão-ouro
      - Terapia de exposição (fobias, pânico, TEPT)
      - EMDR (TEPT, traumas)
      - Mindfulness-based stress reduction (MBSR)
      - Terapia de aceitação e compromisso (ACT)

7. Suplementação funcional ansiolítica e promotora de sono:

   A. Suporte GABAérgico (substituição gradual de BZDs):
      - GABA (500-1000mg) - controverso se atravessa barreira hematoencefálica
      - L-teanina (200-400mg) - promove ondas alfa, relaxamento sem sedação
      - Taurina (500-2000mg) - agonista GABA, ansiolítico suave
      - Magnésio L-treonato ou glicinato (400-600mg) - ESSENCIAL - agonista GABA, relaxamento muscular

   B. Precursores de serotonina (ansiedade, insônia, depressão):
      - 5-HTP (50-200mg) - precursor direto de serotonina
        * Tomar à noite (promove sono)
        * CUIDADO: não combinar com ISRS (risco de síndrome serotoninérgica)
      - Triptofano (500-2000mg) - precursor de 5-HTP→serotonina→melatonina

   C. Adaptógenos (modulação da resposta ao estresse):
      - Ashwagandha (KSM-66) 300-600mg 2x/dia - reduz cortisol, ansiedade
      - Rhodiola rosea 200-400mg/dia - adaptógeno, fadiga, ansiedade
      - Holy basil (Tulsi) 300-600mg 2x/dia - reduz cortisol
      - L-teanina + cafeína (para ansiedade sem sedação diurna)

   D. Promotores de sono (substituição gradual de sedativos):
      - Melatonina 0,5-10mg (1-2h antes dormir) - restaura ritmo circadiano
        * Começar com doses baixas (0,5-1mg), titular conforme necessário
      - Glicina 3g antes de dormir - melhora qualidade do sono
      - Magnésio glicinato 400-600mg - relaxamento, sono profundo
      - Valeriana 300-600mg - sedação suave, não usar >4-6 semanas (pode causar dependência leve)
      - Passiflora (maracujá) 300-500mg - ansiolítico e sedativo suave
      - Lavanda (Silexan) 80-160mg - reduz ansiedade

   E. Cofatores para síntese de neurotransmissores:
      - Complexo B ativado (B6, B12, folato) - essencial para síntese de serotonina, GABA, dopamina
      - Vitamina D 2.000-5.000 UI/dia (alvo >40ng/mL)
      - Ômega-3 (2-3g/dia EPA+DHA) - neuroinflamação, função cerebral
      - Probióticos (eixo intestino-cérebro) - cepas específicas produzem GABA
      - N-acetilcisteína (NAC) 600-1200mg 2x/dia - precursor de glutationa, TOC, tricotilomania

   F. Outros:
      - Inositol 12-18g/dia - pânico, TOC (doses altas)
      - CBD (canabidiol) 20-40mg/dia - ansiedade (onde legal)

8. Intervenções comportamentais e de estilo de vida:

   A. Higiene do sono (insônia):
      - Horário regular (dormir e acordar mesma hora, incluindo fins de semana)
      - Evitar telas 1-2h antes de dormir (luz azul suprime melatonina)
      - Quarto escuro, silencioso, fresco (18-20°C ideal)
      - Cama apenas para dormir e sexo (não TV, celular, trabalho)
      - Evitar cafeína após 14h
      - Evitar álcool (fragmenta sono)
      - Exercício regular (mas não 3h antes de dormir)
      - Exposição à luz natural pela manhã (ajusta ritmo circadiano)
      - Técnica 4-7-8 respiração (relaxamento pré-sono)
      - Banho quente 1-2h antes de dormir (queda de temperatura corporal induz sono)

   B. Manejo de ansiedade:
      - Exercício regular (150min/semana - ansiolítico natural)
      - Meditação diária (10-20min - reduz atividade amígdala)
      - Respiração diafragmática (ativa sistema parassimpático)
      - Yoga (evidência em redução de ansiedade)
      - Contato com natureza (reduz cortisol)
      - Redução de cafeína (<200mg/dia ou eliminação)
      - Eliminação de álcool (ansiedade rebote)
      - Rede de suporte social
      - Journaling (processamento emocional)

9. Intervenções nutricionais:
   - Dieta anti-inflamatória mediterrânea (inflamação afeta humor)
   - Controle glicêmico (evitar montanha-russa de glicose→ansiedade)
   - Proteína adequada (aminoácidos precursores de neurotransmissores)
   - Alimentos ricos em triptofano (peru, ovos, sementes, banana)
   - Alimentos ricos em magnésio (vegetais verde-escuros, nozes, sementes)
   - Probióticos naturais (kefir, iogurte, kimchi, chucrute)
   - Hidratação (desidratação causa ansiedade)
   - Evitar alimentos processados, açúcar, cafeína excessiva

10. Monitoramento durante retirada de benzodiazepínicos:
    - Consultas semanais ou quinzenais (fase inicial)
    - Escalas de ansiedade (GAD-7, BAI) - acompanhar tendência
    - Diário de sintomas (identificar padrões)
    - Sinais vitais (PA, FC - abstinência causa hiperatividade simpática)
    - Sinais de alerta para convulsões
    - Suporte psicológico contínuo
    - Ajuste de velocidade de redução conforme tolerância

11. Suporte psicossocial:
    - Educação sobre dependência e processo de retirada
    - Grupos de suporte (Benzodiazepine Withdrawal Support Groups)
    - Envolvimento familiar (compreensão e suporte)
    - Gerenciamento de expectativas (processo pode levar meses)
    - Reforço positivo (celebrar marcos)

12. Prevenção de recaídas:
    - Plano de manejo de crises de ansiedade sem BZDs
    - Técnicas de coping (respiração, grounding)
    - Continuidade de psicoterapia
    - Manutenção de suplementação
    - Estilo de vida saudável (sono, exercício, alimentação)
    - Identificação precoce de gatilhos
    - Suporte contínuo

**AVISOS CRÍTICOS:**
1. NUNCA suspender benzodiazepínicos abruptamente - risco de convulsões potencialmente fatais
2. Redução deve ser SEMPRE supervisionada por médico experiente (idealmente psiquiatra)
3. Processo pode levar 6-12 meses em usuários crônicos de altas doses
4. Síndrome de abstinência prolongada pode durar meses após suspensão completa
5. Terapia cognitivo-comportamental é ESSENCIAL para sucesso a longo prazo
6. Paciência, persistência e suporte são fundamentais"""
            },

            "anti-hipertensivos": {
                "clinical_relevance": """O histórico de uso de anti-hipertensivos representa marcador de doença cardiovascular estabelecida, confere diagnóstico de hipertensão arterial sistêmica e sinaliza risco aumentado para eventos cardiovasculares maiores (IAM, AVC, insuficiência cardíaca, doença renal crônica). Na medicina funcional integrativa, o uso de anti-hipertensivos é reconhecido como importante para controle pressórico e redução de risco cardiovascular, mas a abordagem enfatiza investigação e tratamento de causas subjacentes modificáveis (resistência à insulina, disfunção endotelial, desequilíbrio autonômico, inflamação vascular, deficiências nutricionais) que podem permitir redução ou suspensão gradual da farmacoterapia em casos selecionados.

As principais classes de anti-hipertensivos incluem:

INIBIDORES DA ECA (IECA - enalapril, captopril, lisinopril, ramipril): bloqueiam conversão de angiotensina I→II, reduzindo vasoconstrição e retenção de sódio. Benefícios adicionais: nefroproteção (especialmente diabéticos), proteção cardiovascular pós-IAM, redução de remodelamento ventricular. Efeitos adversos: tosse seca (10-15% - acúmulo de bradicinina), hipercalemia, angioedema (raro mas grave), deterioração de função renal (estenose bilateral de artéria renal), hipotensão.

BLOQUEADORES DE RECEPTORES DE ANGIOTENSINA (BRA - losartana, valsartana, telmisartana, olmesartana): bloqueiam receptor AT1 da angiotensina II. Efeitos similares a IECAs mas SEM tosse (não afetam bradicinina). Preferidos quando IECA não tolerado. Benefícios adicionais: redução de ácido úrico (losartana), proteção renal, cardiovascular.

BETABLOQUEADORES (BB - atenolol, metoprolol, carvedilol, bisoprolol, propranolol): bloqueiam receptores beta-adrenérgicos, reduzindo frequência cardíaca, contratilidade e débito cardíaco. Indicações preferenciais: pós-IAM, insuficiência cardíaca, arritmias, angina. Efeitos adversos: fadiga, bradicardia, broncoespasmo (evitar em asmáticos), disfunção erétil, mascaramento de hipoglicemia (diabéticos), piora de perfil lipídico e glicêmico, extremidades frias. CRÍTICO: nunca suspender abruptamente (risco de rebote hipertensivo, angina, IAM).

BLOQUEADORES DE CANAIS DE CÁLCIO (BCC - anlodipino, nifedipino, diltiazem, verapamil): bloqueiam entrada de cálcio em células musculares lisas vasculares, causando vasodilatação. Divididos em di-hidropiridínicos (anlodipino, nifedipino - mais vasodilatadores) e não di-hidropiridínicos (diltiazem, verapamil - efeito cronotrópico negativo). Efeitos adversos: edema maleolar (especialmente di-hidropiridínicos), cefaleia, rubor, constipação (verapamil), bradicardia (não di-hidropiridínicos).

DIURÉTICOS (tiazídicos - hidroclorotiazida, clortalidona; poupadores de potássio - espironolactona, amilorida; alça - furosemida): reduzem reabsorção de sódio e água, diminuindo volume intravascular. Tiazídicos são 1ª linha em hipertensão essencial. Efeitos adversos: hipocalemia (tiazídicos, alça), hipercalemia (poupadores de K), hiponatremia, hiperuricemia (gota), hiperglicemia, dislipidemia, disfunção erétil, ginecomastia (espironolactona).

A medicina funcional investiga causas raiz de hipertensão: obesidade visceral, resistência à insulina/síndrome metabólica (60% dos hipertensos), consumo excessivo de sódio com deficiência de potássio e magnésio, deficiência de vitamina D, estresse crônico (excesso de cortisol), apneia obstrutiva do sono (hiperatividade simpática), disfunção endotelial (déficit de óxido nítrico), inflamação vascular crônica, desequilíbrio do microbioma intestinal (produção de TMAO), toxicidade por metais pesados (chumbo, cádmio) e causas secundárias (hiperaldosteronismo primário, feocromocitoma, estenose de artéria renal, doença renal, hipertireoidismo, hiperparatireoidismo).""",

                "patient_explanation": """Se você usa medicamentos para pressão alta (anti-hipertensivos), significa que você tem hipertensão arterial - condição em que o sangue circula com força excessiva nas artérias, sobrecarregando coração e vasos sanguíneos. A pressão alta é perigosa porque aumenta muito o risco de derrame (AVC), infarto, insuficiência cardíaca e problemas nos rins.

Os medicamentos para pressão são muito importantes e geralmente devem ser tomados continuamente - NUNCA pare por conta própria, pois a pressão pode subir perigosamente. Existem vários tipos: IECAs e BRAs (protegem coração e rins), betabloqueadores (reduzem batimentos cardíacos), bloqueadores de cálcio (relaxam vasos sanguíneos) e diuréticos (eliminam líquido e sal).

Na medicina funcional, além dos medicamentos essenciais, investigamos as CAUSAS da sua pressão alta: pode estar relacionada a excesso de peso, resistência à insulina, estresse crônico, consumo excessivo de sal, deficiências de magnésio e potássio, apneia do sono ou inflamação crônica. Em muitos casos, com mudanças intensivas no estilo de vida (perda de peso, alimentação adequada, exercícios, suplementação, manejo de estresse), é possível reduzir ou até suspender gradualmente os medicamentos - MAS isso só deve ser feito sob supervisão médica rigorosa.""",

                "conduct": """1. Avaliação detalhada da terapia anti-hipertensiva:
   - Medicações em uso (classes, doses, horários)
   - Duração do tratamento
   - Adesão ao tratamento (tomar regularmente?)
   - Eficácia (pressão controlada nas metas?)
   - Efeitos adversos (tosse, edema, fadiga, tontura, disfunção sexual)
   - Histórico de ajustes de dose
   - Tentativas prévias de suspensão

2. Monitoramento da pressão arterial:
   - Medidas em consultório (técnica adequada - repouso, braço apoiado)
   - Automonitorização domiciliar (AMPA - 2x/dia, 3 dias/semana)
   - Monitorização ambulatorial 24h (MAPA) - padrão-ouro
   - Identificação de hipertensão do avental branco ou mascarada
   - Metas pressóricas:
     * Geral: <130/80 mmHg
     * Diabéticos, DRC, pós-AVC: <130/80 mmHg
     * Idosos frágeis: individualizar (evitar <120 mmHg)

3. Avaliação de lesão de órgãos-alvo (LOA):
   - Cardiovascular:
     * Ecocardiograma (hipertrofia ventricular esquerda, disfunção diastólica)
     * ECG (sinais de HVE - Sokolow-Lyon, Cornell)
     * Peptídeos natriuréticos (BNP/NT-proBNP se suspeita IC)
   - Renal:
     * Creatinina, TFG estimada
     * Microalbuminúria/proteinúria (marcador precoce)
     * Relação albumina/creatinina urinária
   - Vascular:
     * Espessura médio-intimal carotídea (ultrassom)
     * Índice tornozelo-braquial (doença arterial periférica)
     * Velocidade de onda de pulso (rigidez arterial)
   - Cerebral:
     * Avaliação de AVC prévio ou sinais neurológicos
   - Oftalmológico:
     * Fundo de olho (retinopatia hipertensiva)

4. Investigação de causas secundárias de hipertensão (especialmente se refratária, início <30 anos, ou início abrupto):
   - Renal:
     * Doppler de artérias renais (estenose)
     * Função renal, urinálise
   - Endócrina:
     * Hiperaldosteronismo primário: aldosterona, renina, relação aldosterona/renina
     * Feocromocitoma: metanefrinas urinárias ou plasmáticas
     * Síndrome de Cushing: cortisol urinário 24h, teste de supressão com dexametasona
     * Hiper/hipotireoidismo: TSH, T4L
     * Hiperparatireoidismo: cálcio, paratormônio
   - Apneia obstrutiva do sono: polissonografia (se roncos, sonolência, obesidade)
   - Medicamentos/substâncias: AINEs, anticoncepcionais, corticoides, descongestionantes, cocaína, anfetaminas, alcaçuz, álcool excessivo

5. Avaliação funcional de fatores contribuintes:
   - Composição corporal (bioimpedância, DEXA):
     * Obesidade, especialmente visceral/abdominal
     * Relação cintura-quadril
   - Resistência à insulina:
     * Glicemia, insulina de jejum, HOMA-IR
     * HbA1c
     * Teste de tolerância oral à glicose (se borderline)
   - Perfil lipídico avançado (frequente em hipertensos)
   - Marcadores inflamatórios:
     * PCR ultrassensível
     * Homocisteína (risco cardiovascular independente)
   - Micronutrientes:
     * Magnésio (eritrocitário mais acurado que sérico)
     * Potássio sérico
     * Vitamina D (níveis <30ng/mL associados a hipertensão)
     * Zinco, selênio
   - Função endotelial (se disponível - dilatação mediada por fluxo)
   - Metais pesados (chumbo, cádmio - se exposição ocupacional)
   - Sódio urinário 24h (avaliar consumo)

6. Otimização da farmacoterapia anti-hipertensiva:
   - Avaliar adequação das classes em uso (conforme comorbidades)
   - Preferências conforme perfil:
     * Diabéticos: IECA ou BRA (nefroproteção) + outros conforme necessário
     * Pós-IAM: Betabloqueador + IECA/BRA
     * Insuficiência cardíaca: IECA/BRA + betabloqueador + espironolactona + diurético
     * Idosos, negros: BCC ou diuréticos (melhor resposta)
     * Gestantes: metildopa, nifedipino, labetalol (evitar IECA/BRA - teratogênicos)
   - Combinações racionais:
     * IECA/BRA + BCC + diurético tiazídico (combinação sinérgica)
     * Evitar IECA + BRA juntos (sem benefício, mais eventos adversos)
   - Ajuste de horários (alguns preferem anti-hipertensivos à noite - cronoterapia)

7. Intervenções não farmacológicas intensivas:

   A. Redução de peso (se sobrepeso/obesidade):
      - Meta: perda de 5-10% do peso corporal
      - Cada 1kg perdido→redução ~1mmHg
      - Restrição calórica moderada (500-750kcal/dia déficit)
      - Dieta de baixa carga glicêmica (controle insulina)

   B. Padrão alimentar DASH (Dietary Approaches to Stop Hypertension):
      - Frutas (4-5 porções/dia)
      - Vegetais (4-5 porções/dia)
      - Grãos integrais (6-8 porções/dia)
      - Laticínios magros (2-3 porções/dia)
      - Proteínas magras (aves, peixes)
      - Nozes, sementes, leguminosas
      - Redução de carne vermelha, doces, bebidas açucaradas
      - Efeito: redução de 8-14 mmHg

   C. Restrição de sódio:
      - Meta: <2g/dia (ideal <1,5g/dia)
      - Evitar alimentos processados (maior fonte)
      - Não adicionar sal à comida
      - Ler rótulos (sódio oculto)
      - Efeito: redução de 2-8 mmHg

   D. Aumento de potássio (se função renal normal):
      - Meta: 3,5-5g/dia
      - Fontes: frutas (banana, laranja), vegetais (batata, espinafre, abacate)
      - Suplementação se necessário (sob supervisão)
      - Relação ideal Na:K = 1:2
      - Efeito: redução de 2-8 mmHg

   E. Atividade física regular:
      - Aeróbica: 150min/semana intensidade moderada ou 75min vigorosa
      - Resistida: 2-3x/semana (todos grupos musculares)
      - Isométrica (handgrip): evidência emergente
      - Efeito: redução de 4-9 mmHg

   F. Moderação de álcool:
      - Homens: ≤2 doses/dia
      - Mulheres: ≤1 dose/dia
      - Excesso causa hipertensão
      - Efeito da moderação: redução de 2-4 mmHg

   G. Cessação de tabagismo:
      - Cada cigarro→aumento agudo de PA
      - Fumo aumenta risco cardiovascular sinergicamente
      - Suporte para cessação (TRN, vareniclina, bupropiona, aconselhamento)

8. Suplementação funcional anti-hipertensiva baseada em evidências:

   A. Essenciais (deficiências comuns em hipertensos):
      - Magnésio 400-600mg/dia (glicinato, taurato, citrato)
        * Vasodilatação, antagonista natural de cálcio
        * Redução média: 3-4 mmHg
      - Potássio 2-4g/dia (se função renal normal, não em uso de poupadores de K)
        * Reduz reabsorção de sódio
      - Vitamina D 2.000-5.000 UI/dia (alvo >40ng/mL)
        * Regula sistema renina-angiotensina
        * Redução média: 2-6 mmHg (se deficiente)

   B. Promotores de óxido nítrico (vasodilatação):
      - L-arginina 3-6g/dia (precursor de NO)
        * Cuidado se história de herpes (pode reativar)
      - L-citrulina 3-6g/dia (converte em arginina, melhor biodisponibilidade)
      - Beetroot (suco de beterraba) 250-500ml/dia ou pó
        * Rico em nitratos→nitrito→NO
        * Redução média: 4-10 mmHg
      - Alho envelhecido (aged garlic extract) 600-1200mg/dia
        * Produção de NO, vasodilatação
        * Redução média: 3-8 mmHg

   C. Ômega-3 (EPA+DHA) 2-4g/dia:
      - Anti-inflamatório, melhora função endotelial
      - Redução média: 2-4 mmHg

   D. Coenzima Q10 100-300mg/dia:
      - Antioxidante, melhora função endotelial
      - ESSENCIAL se em uso de estatinas (depletam CoQ10)
      - Redução média: 10-15 mmHg (evidência moderada)

   E. Probióticos (cepas específicas):
      - Lactobacillus plantarum, L. acidophilus
      - Modulação do eixo intestino-rim-cérebro
      - Redução de TMAO (proaterogênico)
      - Redução média: 3-5 mmHg

   F. Antioxidantes:
      - Vitamina C 500-1000mg/dia (melhora função endotelial)
      - Vitamina E 400 UI/dia (tocoferóis mistos)
      - Resveratrol 150-300mg/dia (ativa NO sintase)
      - Cacau/chocolate amargo (flavonóis) 30-50g/dia

   G. Outros:
      - Extrato de folha de oliveira 500-1000mg/dia (oleuropeína)
      - Hibisco (chá) 1-2g/dia
      - Taurina 1-3g/dia (modula cálcio, diurético suave)
      - N-acetilcisteína (NAC) 600mg 2x/dia (antioxidante)

9. Manejo de estresse e modulação autonômica:
   - Respiração lenta (6 respirações/min, 15min/dia) - reduz atividade simpática
   - Meditação mindfulness (20min/dia)
   - Yoga (evidência em redução de PA)
   - Biofeedback
   - Sono adequado (7-9h/noite - privação aumenta PA)
   - Técnicas de relaxamento (progressivo muscular)
   - Redução de estresse psicossocial

10. Monitoramento e ajustes:
    - AMPA regular (2x/dia, 3 dias/semana)
    - Consultas mensais inicialmente (ajustes de medicação/estilo de vida)
    - Reavaliação laboratorial (3-6 meses):
      * Eletrólitos (K, Na, Mg) - especialmente se uso de diuréticos
      * Função renal
      * Glicemia, HbA1c
      * Perfil lipídico
      * Vitamina D
    - MAPA 24h (anualmente ou se PA mal controlada)
    - Ecocardiograma (avaliação de regressão de HVE)

11. Tentativa de redução/suspensão de anti-hipertensivos (casos selecionados):

    **APENAS se:**
    - PA bem controlada (<120/80) por >6 meses
    - Perda de peso significativa (>10% se obeso)
    - Mudanças intensivas e sustentadas de estilo de vida
    - Supervisão médica rigorosa

    **Estratégia:**
    - Reduzir uma medicação por vez
    - Começar pela de menor impacto cardiovascular
    - Redução gradual (25-50% da dose a cada 2-4 semanas)
    - EXCEÇÃO: Betabloqueadores NUNCA suspender abruptamente (redução muito gradual)
    - Monitoramento intensivo de PA (diário durante redução)
    - Interromper redução se PA >130/80 ou aumento >10mmHg
    - Manter IECA/BRA se diabético ou DRC (nefroproteção independente de PA)

12. Prevenção de eventos cardiovasculares:
    - Controle rigoroso de todos fatores de risco:
      * PA <130/80
      * LDL <100mg/dL (idealmente <70 se alto risco)
      * HbA1c <7% (se diabético)
      * Cessação tabagismo
      * Peso adequado
    - Antiagregação plaquetária (AAS) se alto risco cardiovascular (escore de Framingham >10%)
    - Estatina se indicação (dislipidemia, diabetes, doença cardiovascular)
    - Controle de estresse
    - Atividade física regular

**IMPORTANTE:**
1. NUNCA suspender anti-hipertensivos sem orientação médica
2. Betabloqueadores devem ser reduzidos MUITO gradualmente (risco de rebote)
3. Mudanças de estilo de vida são essenciais mas não substituem medicação na maioria dos casos
4. Meta é PA <130/80 para maioria dos pacientes
5. Hipertensão é fator de risco cardiovascular major - controle rigoroso salva vidas"""
            }
        }

        return content_map.get(group, {
            "clinical_relevance": f"Conteúdo clínico para {group} em desenvolvimento.",
            "patient_explanation": f"Explicação para {group} em desenvolvimento.",
            "conduct": f"Conduta para {group} em desenvolvimento."
        })

    def enrich_item(self, item: Dict) -> bool:
        """Enriquece um item específico"""
        print(f"\n{'='*80}")
        print(f"[{ITEMS.index(item) + 1}/{len(ITEMS)}] Processando: {item['name']}")
        print(f"ID: {item['id']}")
        print(f"Grupo: {item['group']}")
        print(f"{'='*80}\n")

        # Obter conteúdo baseado no grupo
        content = self.get_content_for_group(item['group'])

        # Atualizar via API
        success = self.update_score_item(
            item['id'],
            content['clinical_relevance'],
            content['patient_explanation'],
            content['conduct']
        )

        if success:
            print(f"✓ Item '{item['name']}' enriquecido com sucesso!\n")
        else:
            print(f"✗ Falha ao enriquecer item '{item['name']}'\n")

        return success

    def process_all_items(self):
        """Processa todos os items"""
        results = {
            "success": [],
            "failed": [],
            "total": len(ITEMS)
        }

        print(f"\n{'='*80}")
        print(f"ENRIQUECIMENTO DE HISTÓRICO DE DOENÇAS - BATCH 2")
        print(f"Total de items: {len(ITEMS)}")
        print(f"{'='*80}\n")

        for item in ITEMS:
            success = self.enrich_item(item)
            if success:
                results["success"].append(item['name'])
            else:
                results["failed"].append(item['name'])

            # Delay entre requisições
            time.sleep(0.3)

        # Sumário final
        print(f"\n{'='*80}")
        print("SUMÁRIO DO PROCESSAMENTO")
        print(f"{'='*80}")
        print(f"Total de items: {results['total']}")
        print(f"✓ Sucesso: {len(results['success'])}")
        print(f"✗ Falhas: {len(results['failed'])}")

        if results['failed']:
            print(f"\nItems com falha:")
            for name in results['failed']:
                print(f"  - {name}")

        print(f"{'='*80}\n")

        return results

def main():
    enricher = DiseaseHistoryEnricher()

    # Login
    if not enricher.login():
        print("✗ Erro: Não foi possível fazer login. Encerrando.")
        return

    # Processar todos items
    results = enricher.process_all_items()

    # Salvar resultados
    output_file = '/home/user/plenya/DISEASE-HISTORY-BATCH2-RESULTS.json'
    with open(output_file, 'w', encoding='utf-8') as f:
        json.dump(results, f, indent=2, ensure_ascii=False)

    print(f"✓ Resultados salvos em: {output_file}")

if __name__ == "__main__":
    main()
