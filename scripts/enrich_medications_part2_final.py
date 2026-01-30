#!/usr/bin/env python3
"""
BATCH FINAL - 7 items restantes
"""
import requests, json, time
from typing import Dict

API_BASE_URL = "http://localhost:3001/api/v1"
EMAIL, PASSWORD = "import@plenya.com", "Import@123456"

CONTENT = {
    "18faeba8-2ccb-4adf-afed-64e725f1f4c7": {  # Antiosteoporóticos
        "name": "Antiosteoporóticos",
        "cr": """Antiosteoporóticos indicam osteoporose (densidade mineral óssea T-score ≤-2,5) ou alto risco de fraturas. Principais classes: BISFOSFONATOS (alendronato, risedronato, ácido zoledrônico) - inibem reabsorção óssea por osteoclastos. Eficazes mas efeitos adversos: esofagite (tomar com água abundante, ficar ereto 30min), osteonecrose de mandíbula (raro, risco maior se procedimentos dentários), fraturas atípicas de fêmur (uso >5 anos). DENOSUMABE (Prolia) - anticorpo anti-RANKL, SC semestral. Eficaz mas risco de fraturas vertebrais múltiplas se descontinuação abrupta. TERIPARATIDA (Forteo) - PTH recombinante, anabólico ósseo, uso máximo 2 anos. RALOXIFENO - modulador seletivo receptor estrogênio, reduz risco fraturas vertebrais. Reposição hormonal (TRH) em mulheres pós-menopausa selecionadas. Deficiências induzidas: bisfosfonatos podem reduzir cálcio, magnésio. Medicina funcional investiga causas: deficiência vitamina D/K2, baixa ingestão cálcio, sedentarismo, tabagismo, álcool, deficiência hormonal, hiperparatireoidismo, hipercortisolismo, acidose metabólica crônica, deficiência proteica, disbiose intestinal (má absorção). Suporte: cálcio 1000-1200mg/dia, vitamina D 2000-4000UI (alvo >40ng/mL), vitamina K2 100-200mcg, magnésio 400-600mg, exercícios resistidos e impacto, proteína adequada 1-1,2g/kg.""",
        "pe": """Antiosteoporóticos tratam osteoporose (ossos frágeis, risco de fraturas). Bisfosfonatos são os mais usados (alendronato semanal, risedronato). IMPORTANTE: tomar em jejum com água abundante, ficar sentado/em pé 30 minutos (evitar refluxo). Denosumabe é injeção semestral. Todos reduzem risco de fraturas. Efeitos colaterais: azia, dor muscular, raramente osteonecrose mandíbula (evitar extrações dentárias). Essencial: cálcio e vitamina D diários, exercícios com peso, não fumar.""",
        "cd": """1. DEXA scan baseline, repetir 1-2 anos\n2. Cálcio 1000-1200mg/dia + vitamina D 2000-4000UI (alvo >40ng/mL)\n3. Vitamina K2 (MK-7) 100-200mcg/dia\n4. Magnésio 400-600mg/dia\n5. Proteína 1-1,2g/kg/dia\n6. Exercícios resistidos 2-3x/semana + impacto\n7. Avaliação dentária antes bisfosfonatos\n8. PTH, 25(OH)D, cálcio, fósforo, marcadores ósseos\n9. Reavaliação necessidade após 3-5 anos (holiday drug se baixo risco)"""
    },

    "8c91bc11-46ed-4f20-b5f0-eb1edf5f13be": {  # Antiparkinsonianos
        "name": "Antiparkinsonianos",
        "cr": """Antiparkinsonianos indicam doença de Parkinson ou parkinsonismo. Principal: LEVODOPA + CARBIDOPA (Prolopa, Sinemet) - L-DOPA é precursor dopamina, carbidopa inibe conversão periférica. Tratamento gold-standard mas complicações motoras após 5-10 anos (flutuações on-off, discinesias). Efeitos adversos: náuseas, hipotensão postural, alucinações, sonolência súbita, compulsões (jogo, sexo, compras). AGONISTAS DOPAMINÉRGICOS (pramipexol, ropinirol, rotigotina) - estimulam diretamente receptores. Usados como monoterapia inicial (<60 anos) ou adjuvantes. Efeitos adversos: náusea, edema, sonolência extrema, compulsões (>10%). INIBIDORES MAO-B (selegilina, rasagilina) - inibem degradação dopamina. Monoterapia inicial ou adjuvante. INIBIDORES COMT (entacapone, tolcapone) - prolongam ação levodopa. Tolcapone risco hepatotoxicidade. ANTICOLINÉRGICOS (biperideno) - tremor refratário. Efeitos anticolinérgicos intensos (confusão, retenção urinária, constipação - evitar idosos). AMANTADINA - reduz discinesias. Deficiências induzidas: Levodopa depleta vitaminas B6, B12, folato, SAMe, aumenta homocisteína. Medicina funcional investiga: toxicidade metais pesados, disfunção mitocondrial, estresse oxidativo, neuroinflamação. Suporte: CoQ10 1200-2400mg/dia, vitamina E, ômega-3, creatina, NAC, vitaminas B.""",
        "pe": """Antiparkinsonianos tratam doença de Parkinson (tremor, rigidez, lentidão de movimentos). Levodopa é o mais eficaz, transforma-se em dopamina no cérebro. Tomar antes das refeições para melhor absorção (proteínas competem). Efeitos colaterais: náuseas (melhoram com tempo), tonturas ao levantar, sonolência, movimentos involuntários (após anos), compulsões (avisar médico). NUNCA parar abruptamente. Exercícios regulares são essenciais (retardam progressão).""",
        "cd": """1. Titulação lenta conforme resposta\n2. Levodopa: tomar 30-60min antes refeições (evitar competição com proteínas)\n3. Suplementação OBRIGATÓRIA:\n   - Vitamina B6, B12, folato (prevenir hiperhomocisteinemia)\n   - CoQ10 1200-2400mg/dia\n   - Vitamina E 400-800UI\n   - Ômega-3 2-3g/dia\n   - NAC 1-2g/dia\n   - Creatina 5-10g/dia\n4. Exercícios: aeróbicos, tai chi, dança\n5. Fisioterapia, fonoaudiologia\n6. Monitorar compulsões, alucinações\n7. Evitar antipsicóticos típicos (bloqueiam dopamina)"""
    },

    "3a5048fe-1ff2-4f39-a8fd-1332fa906fa5": {  # Antipsicóticos
        "name": "Antipsicóticos",
        "cr": """Antipsicóticos (neurolépticos) tratam esquizofrenia, transtorno bipolar (mania), psicose, agitação. TÍPICOS (haloperidol, clorpromazina) - antagonistas D2. Eficazes mas efeitos extrapiramidais intensos: parkinsonismo, acatisia (inquietação), distonia aguda, discinesia tardia (movimentos involuntários - pode ser irreversível). Síndrome neuroléptica maligna (rara, grave): rigidez, hipertermia, rabdomiólise. ATÍPICOS (risperidona, olanzapina, quetiapina, aripiprazol, clozapina) - bloqueio D2 + 5-HT2A. Menos efeitos extrapiramidais mas SÍNDROME METABÓLICA: ganho de peso massivo (olanzapina >10kg), diabetes, dislipidemia, hipertensão. Risperidona: hiperprolactinemia (galactorreia, disfunção sexual, osteoporose). Quetiapina: sedação intensa, usado off-label insônia/ansiedade. CLOZAPINA: mais eficaz em refratários mas agranulocitose (1%) - hemograma semanal (6 meses), depois quinzenal. Todos prolongam QT (risco torsades). Deficiências induzidas: antipsicóticos podem reduzir CoQ10. Medicina funcional investiga psicose: neuroinflamação, autoimunidade (NMDA), toxinas, deficiências nutricionais (B12, folato, niacina), doença celíaca, infecções, disfunção mitocondrial. Suporte: ômega-3, NAC, glicina, vitaminas B.""",
        "pe": """Antipsicóticos tratam doenças mentais graves (esquizofrenia, psicose, mania). Bloqueiam dopamina no cérebro. Divididos em típicos (haloperidol - mais efeitos motores) e atípicos (risperidona, olanzapina, quetiapina - mais ganho peso/diabetes). Efeitos colaterais: ganho peso significativo, sonolência, tremor, rigidez, inquietação, diabetes, colesterol alto. Clozapina é mais eficaz mas requer exames de sangue frequentes (risco baixa glóbulos brancos). NUNCA parar abruptamente. Exercícios e dieta são essenciais para controlar peso.""",
        "cd": """1. Monitoramento metabólico OBRIGATÓRIO:\n   - Peso, IMC, circunferência abdominal mensal (3 meses), depois trimestral\n   - Glicemia jejum, HbA1c baseline, 3 meses, anual\n   - Perfil lipídico baseline, 3 meses, anual\n   - Pressão arterial\n2. Hemograma (clozapina: semanal 6 meses, depois quinzenal)\n3. Prolactina (risperidona, paliperidona)\n4. ECG baseline (QTc)\n5. Função hepática\n6. Prevenção síndrome metabólica:\n   - Metformina 500-1500mg/dia (previne ganho peso/diabetes)\n   - Exercícios regulares\n   - Dieta rigorosa\n7. Suplementação:\n   - Ômega-3 2-3g/dia\n   - NAC 1-2g/dia\n   - Vitaminas B\n   - CoQ10 100-300mg/dia\n   - Glicina 15-30g/dia (adjuvante sintomas negativos)\n8. Manejo efeitos extrapiramidais: biperideno, propranolol (acatisia)\n9. Reavaliação necessidade/dose regularmente"""
    },

    "0a0507a5-ce18-4c92-99e3-48fb31d32b5f": {  # Antivirais
        "name": "Antivirais",
        "cr": """Antivirais tratam infecções virais específicas. HIV: TARV (terapia antirretroviral) com inibidores nucleosídeos transcriptase reversa (lamivudina, tenofovir, abacavir), não-nucleosídeos (efavirenz, nevirapina), inibidores protease (atazanavir, darunavir), inibidores integrase (dolutegravir, bictegravir). Objetivo: carga viral indetectável, CD4 >500. Efeitos adversos: lipodistrofia, dislipidemia, resistência insulínica, nefrotoxicidade (tenofovir), osteoporose, hepatotoxicidade. Hepatite C: DAAs (sofosbuvir, ledipasvir, velpatasvir) - cura >95% em 8-12 semanas, bem tolerados. Hepatite B: entecavir, tenofovir - supressão viral. Herpes: aciclovir, valaciclovir - reduzem duração/gravidade. Influenza: oseltamivir (Tamiflu) - reduz 1 dia sintomas se início precoce. CMV: ganciclovir (transplantados, HIV). Deficiências induzidas: TARV pode reduzir vitamina D, carnitina. Medicina funcional enfatiza suporte imunológico.""",
        "pe": """Antivirais tratam vírus específicos. HIV: medicamentos modernos (1 comprimido/dia) controlam vírus, permitem vida normal. Hepatite C: cura com 8-12 semanas tratamento. Herpes: aciclovir reduz surtos. Importante: tomar exatamente como prescrito (resistência viral se falhar), fazer exames regulares. HIV: uso correto leva a carga viral indetectável = não transmite sexualmente (I=I). Efeitos colaterais variam: náusea, diarreia, alterações metabólicas.""",
        "cd": """HIV:\n1. TARV imediato, objetivo CV indetectável (<50 cópias/mL)\n2. Monitorar CD4 e CV trimestral\n3. Perfil lipídico, glicemia, função renal semestral\n4. Densidade óssea (DEXA) baseline\n5. Vitamina D >40ng/mL\n6. Rastreamento infecções oportunistas conforme CD4\n7. Profilaxia conforme CD4 (<200: Pneumocystis, <100: Toxoplasma)\n8. Suporte: ômega-3, vitamina D, probióticos, NAC\n\nHEPATITE C:\n1. Tratamento DAAs 8-12 semanas\n2. RVS12 (cura): HCV-RNA indetectável 12 semanas pós-tratamento\n3. Se cirrose: continuar rastreamento hepatocarcinoma"""
    },

    "1647a98b-9883-4dad-baf7-f92b6f987100": {  # Agonistas GLP-1
        "name": "Agonistas de GLP-1",
        "cr": """Agonistas GLP-1 (liraglutida, semaglutida, dulaglutida) mimetizam incretina GLP-1: estimulam secreção insulina glicose-dependente, inibem glucagon, retardam esvaziamento gástrico, aumentam saciedade. Uso: DM2, obesidade (semaglutida 2,4mg - Wegovy). BENEFÍCIOS: perda peso 10-15% (dose obesidade), redução HbA1c 1-1,5%, proteção cardiovascular (redução IAM, AVC, morte CV), possível proteção renal, não causam hipoglicemia. DESVANTAGENS: custo altíssimo (R$1000-2000/mês), injetáveis SC semanal, náusea intensa (30-50% - melhora com titulação lenta), vômitos, diarreia, constipação, risco pancreatite (raro), possível risco câncer medular tireoide (contraindicado se história pessoal/familiar CMT ou neoplasia endócrina múltipla tipo 2), gastroparesia, perda massa muscular junto com gordura. Efeito rebote se descontinuação (ganho peso). Medicina funcional reconhece utilidade mas enfatiza mudanças estilo vida sustentáveis.""",
        "pe": """Agonistas GLP-1 (Ozempic, Victoza, Trulicity) são injeções semanais que imitam hormônio intestinal. Usados em diabetes tipo 2 e obesidade. Benefícios: perdem 10-15% peso, melhoram diabetes, protegem coração. Efeitos colaterais: náusea/vômitos (30-50%, melhoram em semanas), prisão de ventre. Importante: injeção sob pele 1x/semana, dieta rica em proteína (evitar perda muscular), exercícios resistidos. Custo alto. Efeito rebote se parar (ganho peso volta). Não usar se histórico câncer tireoide na família.""",
        "cd": """1. Titulação lenta (reduzir náusea)\n2. Orientar injeção SC (abdome, coxa)\n3. Dieta proteica (1,6-2g/kg) + exercícios resistidos (prevenir sarcopenia)\n4. Hidratação adequada\n5. Monitorar náusea/vômitos, pancreatite, gastroparesia\n6. HbA1c trimestral\n7. Peso mensal\n8. NÃO usar se: história pessoal/familiar câncer medular tireoide, NEM2\n9. Estratégia saída: transição para hábitos sustentáveis antes descontinuar"""
    },

    "e94f9287-d972-43ad-b6d1-ee59ed991d31": {  # Arritmia (condição)
        "name": "Arritmia (condição)",
        "cr": """Arritmia representa distúrbio ritmo cardíaco: taquiarritmias (FA, flutter, taquicardia supraventricular, ventricular), bradiarritmias (disfunção sinusal, bloqueios AV), extrassístoles. FA é mais comum, risco AVC 5x. Medicina funcional investiga: distúrbios eletrolíticos (Mg, K), disfunção tireoidiana, apneia sono, estimulantes, estresse, inflamação. Tratamento: antiarrítmicos, betabloqueadores, anticoagulação (FA), ablação, marca-passo. Suporte: magnésio 400-600mg/dia, CoQ10, ômega-3, taurina, redução cafeína/álcool.""",
        "pe": """Arritmia são batimentos cardíacos irregulares. Tipos: muito rápido (taquicardia), muito lento (bradicardia), batidas "extras". Fibrilação atrial é comum, faz coração bater desorganizado, risco coágulos/derrame. Tratamento: medicamentos controlam ritmo/velocidade, anticoagulantes previnem derrame. Importante: evitar cafeína/álcool em excesso, tratar apneia sono, suplementar magnésio.""",
        "cd": """1. ECG, Holter 24-48h, ecocardiograma\n2. Investigar causas: eletrólitos (Mg, K), TSH, polissonografia\n3. Anticoagulação se FA (CHA2DS2-VASc ≥2)\n4. Antiarrítmicos, betabloqueadores conforme tipo\n5. Magnésio 400-600mg/dia\n6. CoQ10 100-200mg/dia\n7. Ômega-3 2g/dia\n8. Taurina 1-3g/dia\n9. Redução cafeína, álcool, estresse\n10. Tratamento apneia sono (CPAP)"""
    },

    "8b232987-3a9a-48ab-8051-04356897d97b": {  # Artrite
        "name": "Artrite",
        "cr": """Artrite refere-se a inflamação articular. Principais: OSTEOARTRITE (degenerativa, mecânica, idade >50) e ARTRITE REUMATOIDE (autoimune, inflamatória, articulações simétricas). OA: dor mecânica, rigidez matinal <30min, nódulos Heberden/Bouchard. Tratamento: analgésicos, AINEs, infiltração corticoide, condroitina/glicosamina, reabilitação, prótese (casos graves). AR: dor inflamatória, rigidez >1h, artrite simétrica mãos/pés, fator reumatoide/anti-CCP positivos, elevação VHS/PCR. Tratamento: DMARDs (metotrexato, sulfassalazina), biológicos (anti-TNF, anti-IL6), corticoides. Medicina funcional investiga: autoimunidade, permeabilidade intestinal, disbiose, intolerâncias alimentares, deficiências (vit D, ômega-3). Suporte: dieta anti-inflamatória, ômega-3 2-4g/dia, cúrcuma, boswellia, condroitina/glicosamina (OA), vitamina D.""",
        "pe": """Artrite é inflamação das articulações causando dor, inchaço, rigidez. Tipos: Osteoartrite (desgaste, comum em idosos, joelhos/quadris/mãos) e Artrite Reumatoide (autoimune, ataca articulações, mãos/pés). Tratamento OA: analgésicos, fisioterapia, perda peso, suplementos articulares. AR: medicamentos imunossupressores (metotrexato, biológicos). Suporte: dieta anti-inflamatória, ômega-3, cúrcuma, exercícios leves, compressas.""",
        "cd": """OSTEOARTRITE:\n1. Analgésicos (paracetamol), AINEs (uso curto)\n2. Fisioterapia, exercícios baixo impacto\n3. Perda peso (se sobrepeso)\n4. Sulfato glicosamina 1500mg/dia + condroitina 1200mg/dia\n5. Colágeno tipo II 40mg/dia\n6. Cúrcuma 500-1000mg/dia\n7. Ômega-3 2-3g/dia\n\nARTRITE REUMATOIDE:\n1. DMARDs (metotrexato 15-25mg/semana + ácido fólico)\n2. Biológicos se refratária\n3. Vitamina D >40ng/mL\n4. Ômega-3 3-4g/dia\n5. Dieta eliminação (glúten, laticínios)\n6. Probióticos\n7. Cúrcuma, boswellia\n8. Monitorar: hemograma, enzimas hepáticas (metotrexato)"""
    },

    "2fd99065-70ad-4d33-86b8-0ac8be744547": {  # Asma
        "name": "Asma",
        "cr": """Asma é doença inflamatória crônica vias aéreas com hiperreatividade brônquica, obstrução reversível. Sintomas: dispneia, sibilos, tosse, aperto torácico (piora noturna/manhã). Classificação: intermitente, persistente leve/moderada/grave. Gatilhos: alérgenos (ácaros, pólen, animais), exercício, ar frio, infecções, DRGE, AINEs. Diagnóstico: espirometria (VEF1/CVF <0,7, reversibilidade >12% pós-broncodilatador). Tratamento: CONTROLADORES (corticoides inalados - CI, CI+LABA, anti-leucotrienos), RESGATE (beta-2 agonistas curta ação - salbutamol). Grave: anti-IgE (omalizumab), anti-IL5 (mepolizumab). Efeitos adversos CI: candidíase oral (enxaguar boca), rouquidão, supressão adrenal (doses altas). Medicina funcional investiga: alergias alimentares, deficiência vit D, disbiose, inflamação, estresse. Suporte: vitamina D >40ng/mL, ômega-3, magnésio, vitamina C, quercetina, controle ambiental.""",
        "pe": """Asma é doença crônica que inflama e estreita vias respiratórias, causando falta de ar, chiado, tosse. Gatilhos: poeira, pólen, animais, exercício, ar frio. Tratamento: bombinhas controladas diárias (corticoide inalado - previne crises) + bombinha resgate (salbutamol - alivia crise). Importante: usar controlador TODOS os dias mesmo sem sintomas, enxaguar boca após uso (evitar sapinho), evitar gatilhos, ter plano ação para crises. Exercícios são benéficos (natação). Suplementar vitamina D.""",
        "cd": """1. Espirometria baseline, anual\n2. Plano ação escrito (zona verde/amarela/vermelha)\n3. Técnica inalatória correta (espaçador)\n4. CI dose baixa/moderada diária\n5. Beta-2 resgate (se uso >2x/semana = controle inadequado)\n6. Controle ambiental: capas antiácaros, evitar carpetes, animais\n7. Tratamento DRGE se presente\n8. Vitamina D 2000-5000UI (alvo >40ng/mL)\n9. Ômega-3 2-3g/dia\n10. Magnésio 400-600mg/dia\n11. Vitamina C 1-2g/dia\n12. Quercetina 500-1000mg/dia\n13. Exercícios respiratórios, natação\n14. Vacinação influenza anual, pneumococo"""
    }
}

class Enricher:
    def __init__(self):
        self.token = None
        self.session = requests.Session()

    def login(self):
        try:
            r = self.session.post(f"{API_BASE_URL}/auth/login", json={"email": EMAIL, "password": PASSWORD})
            r.raise_for_status()
            self.token = r.json()["accessToken"]
            self.session.headers.update({"Authorization": f"Bearer {self.token}"})
            print("✓ Login OK\n")
            return True
        except: return False

    def update(self, item_id, data):
        try:
            payload = {"clinicalRelevance": data["cr"], "patientExplanation": data["pe"], "conduct": data["cd"]}
            r = self.session.put(f"{API_BASE_URL}/score-items/{item_id}", json=payload)
            r.raise_for_status()
            return True
        except: return False

    def run(self):
        results = {"success": [], "failed": [], "total": len(CONTENT)}
        print("="*80)
        print("BATCH FINAL - 7 ITEMS")
        print("="*80 + "\n")

        for idx, (item_id, data) in enumerate(CONTENT.items(), 1):
            print(f"[{idx}/7] {data['name']}")
            if self.update(item_id, data):
                print("    ✓ OK\n")
                results["success"].append(data['name'])
            else:
                print("    ✗ FALHOU\n")
                results["failed"].append(data['name'])
            time.sleep(0.5)

        print("="*80)
        print(f"RESULTADO: {len(results['success'])}/7 SUCESSO")
        if results['failed']:
            print(f"Falhas: {', '.join(results['failed'])}")
        print("="*80)
        return results

def main():
    e = Enricher()
    if not e.login(): return
    e.run()

if __name__ == "__main__":
    main()
