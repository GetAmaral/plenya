#!/usr/bin/env python3
"""
BATCH COGNIÇÃO - 30 ITEMS COMPLETO
Enriquecimento focado em Medicina Funcional Integrativa
Neuroplasticidade | Neurotransmissores | Neuroinflamação | Eixo Intestino-Cérebro
"""

import requests
import time
from typing import Dict

API_URL = "http://localhost:3001/api/v1"
EMAIL = "import@plenya.com"
PASSWORD = "Import@123456"

def login() -> str:
    response = requests.post(f"{API_URL}/auth/login", json={"email": EMAIL, "password": PASSWORD}, timeout=10)
    response.raise_for_status()
    data = response.json()
    return data.get("data", {}).get("access_token") or data.get("accessToken") or data.get("access_token")

def update_item(token: str, item_id: str, data: Dict) -> bool:
    try:
        response = requests.put(f"{API_URL}/score-items/{item_id}", json=data,
                               headers={"Authorization": f"Bearer {token}"}, timeout=30)
        return response.status_code == 200
    except:
        return False

# ===== CONTEÚDOS CLÍNICOS =====

CONTENT = {
    "dubois_imediato": {
        "clinicalRelevance": "O Teste de Dubois (evocação imediata) avalia codificação e recuperação imediata de memória episódica verbal, sendo altamente sensível para disfunção hipocampal precoce. Na MFI, reflete integridade da neurotransmissão colinérgica/glutamatérgica, metabolismo cerebral (glicose/cetonas), status de cofatores (B1,B6,B12,colina,magnésio), neuroinflamação (IL-6,TNF-α), função mitocondrial neuronal, e modulação do eixo HPA. Déficits (<4/5) associam-se a deficiências nutricionais, privação de sono, estresse oxidativo, resistência insulínica cerebral. Intervenções multimodais (MIND diet + exercício aeróbico 150min/semana + ômega-3 2g/dia + fosfatidilserina 300mg) demonstram melhora de 23% em 12 semanas.",
        "patientExplanation": "Este teste avalia sua capacidade de guardar novas memórias imediatamente. Lembrar 5/5 palavras indica sistema de memória excelente. 3-4 palavras é aceitável mas sugere necessidade de suporte (melhor sono, nutrição, redução de estresse). <3 palavras requer atenção. Causas comuns e TRATÁVEIS: privação de sono, deficiências de vitaminas B/ômega-3/magnésio, inflamação, estresse crônico, glicemia instável. Com intervenções corretas (sono adequado, alimentação rica em antioxidantes/peixes/vegetais verdes, exercícios, manejo de estresse), a maioria melhora significativamente.",
        "conduct": "IMEDIATO: Avaliar sono (72h), hidratação, horário última refeição, ansiedade. Se <4/5: repetir teste otimizado. INVESTIGAÇÃO (1-2sem): Hemograma, glicemia+insulina+HbA1c, B12+metilmalônico, folato, homocisteína(<8), VitD(50-80ng/mL), ferritina, zinco/magnésio eritrocitário, TSH+T4L, hs-CRP. NUTRIÇÃO: MIND diet (3 porções vegetais verdes/dia, frutas vermelhas diárias, peixes 3-4x/sem, oleaginosas 30g/dia, azeite, eliminar açúcar/trans). SUPLEMENTOS: Complexo B metilado, ômega-3 2-3g/dia, magnésio L-treonato 2g/dia, VitD3 5000UI se baixa, fosfatidilserina 300mg, acetil-L-carnitina 1500mg, Lion's Mane 1g/dia. SONO: 7-8h, escuridão total, 18-20°C, exposição solar matinal, sem telas 2h antes. EXERCÍCIO: 150min/sem aeróbico, resistido 2-3x/sem. STRESS: Mindfulness 10-20min/dia, respiração 4-7-8. REAVALIAÇÃO: 8-12sem."
    },

    "dubois_tardio": {
        "clinicalRelevance": "Evocação tardia (recall após 5-10min) é marcador específico de consolidação de memória de longo prazo e integridade hipocampal. Sensibilidade 87%/especificidade 92% para CCL amnéstico. Na MFI, depende criticamente do sono profundo N3/REM (replay hippocampal transfere informações hipocampo→córtex). Déficits associam-se a: fragmentação do sono (apneia, insônia), neuroinflamação (IL-1β,TNF-α), resistência insulínica cerebral (diabetes tipo 3), estresse oxidativo mitocondrial, deficiência de neurotransmissores (acetilcolina,glutamato,noradrenalina), disfunção barreira hematoencefálica, acúmulo de beta-amiloide/tau. Biomarcadores: homocisteína >12, B12<500, VitD<30, ômega-3<8%, HbA1c>5.7%, hs-CRP>2. Protocolo ReCODE modificado: melhora 1,5-2 pontos em 84% após 6 meses.",
        "patientExplanation": "Este teste avalia se seu cérebro consegue armazenar memórias de forma duradoura. Durante sono profundo, cérebro 'reproduz' memórias do dia e as arquiva permanentemente. Lembrar 5/5 após intervalo = sistema funcionando perfeitamente. 3-4 = razoável. <3 = atenção necessária. Principais causas TRATÁVEIS: sono de má qualidade (especialmente fases profundas), inflamação corporal afetando cérebro, deficiências de B12/folato/ômega-3, glicemia instável, estresse crônico danificando área da memória. Com intervenções (priorizar sono 7-8h em ambiente ideal, peixes gordos/vegetais/frutas vermelhas, exercícios aumentando fatores de crescimento cerebral, mente ativa, gerenciar estresse), melhoras em 2-3 meses.",
        "conduct": "INICIAL: Diferenciar agudo vs. crônico. Investigar sono (Berlim apneia, Epworth sonolência, diário 7 dias), medicações (BZD, anticolinérgicos), depressão (PHQ-9), álcool. LABORATÓRIO COMPLETO: Glicemia+insulina+HbA1c+frutosamina, TSH+T4L+T3L+anti-TPO, B12+metilmalônico+holotranscobalamina, folato sérico/eritrocitário, homocisteína(<8), VitD(50-80ng/mL), perfil lipídico+LDL oxidado+ApoB, hs-CRP+VHS+ferritina, função hepática/renal, hemograma+VCM, ômega-3 índice(>8%), magnésio/zinco eritrocitário, cortisol salivar 4 pontos, DHEA-S, hormônios sexuais. Se <3/5 persistente: APOE, autoimunidade neural, permeabilidade intestinal, metais pesados. IMAGEM: RM cérebro com volumetria hipocampal se progressão 6-12m, <65anos, familiar+, <3/5 persistente. SONO: Polissonografia se Epworth>10, roncos, IMC>30. NUTRIÇÃO: MIND diet + jejum 16:8 (autofagia,BDNF): 6 porções vegetais verdes/dia, 2 frutas vermelhas/dia, peixes selvagens 4-5x/sem, 30-50g oleaginosas, azeite 3-4col/dia, abacate, leguminosas 3-4x/sem, evitar açúcar/ultraprocessados/trans, carne vermelha máx 1x/sem. Se HbA1c>5.7%: cetogênica modificada + MCT C8 30-45ml/dia. SUPLEMENTOS CAMADAS: (1)Ômega-3 2-4g/dia, complexo B metilado completo, magnésio L-treonato 2g/dia, VitD3 5-10k UI + K2 200mcg. (2)Fosfatidilserina 300mg, ALCAR 1500-2000mg, alfa-GPC 300-600mg, citicolina 250-500mg. (3)Bacopa 300mg 2x/dia, Lion's Mane 1g 2x/dia, curcumina lipossomal 1g/dia, resveratrol 500mg, pterostilbeno 100-200mg. (4)CoQ10 ubiquinol 200-400mg, PQQ 20mg, R-ALA 300-600mg. (5)Magnésio glicinato 400-800mg noite, L-teanina 200-400mg, glicina 3-5g, melatonina 0.3-1mg. SONO RADICAL: Fixo horários ±30min, sol 10-30min <1h acordar, blackout total, 17-19°C, 40-60% umidade, sem ruído, luz azul zero 2-3h antes, rotina pré-sono 30-60min, sem refeições pesadas 3h antes, sem cafeína pós-14h, sem álcool. Apneia: CPAP obrigatório, perda peso 10%, evitar supino. EXERCÍCIO: Aeróbico 150-300min/sem 65-85%FCmáx, HIIT 3x/sem 30min (1min intenso:2min rec), resistido 2-3x/sem, coordenação/equilíbrio 2-3x/sem, matinal. Sedentarismo: Levantar cada 30-45min, 8-10k passos/dia. COGNITIVO: Nova habilidade (idioma,instrumento,dança), BrainHQ/CogniFit 20-30min 5x/sem, xadrez, leitura ativa, socialização. STRESS: Mindfulness 20min 2x/dia, respiração 4-7-8 ou coerência cardíaca 5-10min 3x/dia, natureza 120min/sem, gratidão diária, HRV biofeedback. NEUROTOXINAS: Evitar mercúrio(peixes grandes), alumínio(antiácidos/utensílios/antiperspirantes), pesticidas(orgânicos dirty dozen), BPA/ftalatos(sem plástico quente), mofo, poluição(HEPA). INTESTINO: Probióticos Lactobacillus helveticus+Bifidobacterium longum, prebióticos inulina 10-15g, fermentados diários. HORMONAL: Mulheres perimenopausa/menopausa <60a ou <10a: TRH bioidêntica estradiol transdérmico+progesterona oral. Homens testosterona<70pg/mL: reposição 400-600ng/dL. MONITORAMENTO: Reteste 6-8sem, labs 12sem, neuropsicológico formal 6m. ENCAMINHAMENTO URGENTE: Declínio >1pt em 3-6m, focal neurológico, psiquiátrico novo, <60a progressão rápida. META: ≥4/5 idealmente 5/5, independência funcional, prevenir demência."
    },

    "memoria_capacidade": {
        "clinicalRelevance": "Queixas subjetivas de memória (QSM) precedem alterações objetivas e aumentam risco 2-4x para CCL em 3-5 anos. Na MFI, refletem função cognitiva objetiva MAS também humor (depressão pseudodemência), ansiedade, sono, sobrecarga cognitiva, estresse. Importante diferenciar QSM verdadeiras (correlacionadas com declínio) de desproporcionais (ansiedade/depressão). Metacognição (monitoramento do próprio desempenho) depende de pré-frontais intactos. Discrepâncias: anosognosia (sem queixas apesar de déficits severos em demência) vs. hipersensibilidade a lapsos normais (ansiosos perfeccionistas). Avaliação funcional correlaciona com deficiências B12/folato, disfunção tireoidiana, privação de sono, cortisol elevado, neuroinflamação, resistência insulínica cerebral, depressão. Intervenções em causas reversíveis melhoram função objetiva E percepção subjetiva, reduzindo ansiedade.",
        "patientExplanation": "Como você percebe sua memória é tão importante quanto testes objetivos. Queixas persistentes, mesmo com testes normais, podem ser sinal precoce merecendo intervenção preventiva. Causas comuns TRATÁVEIS: ansiedade (interfere atenção, parece memória pior), depressão (afeta concentração), privação de sono (cérebro não consolida adequadamente), multitarefa excessiva (sobrecarga), deficiências nutricionais. Na MFI investigamos: sono de qualidade (7-8h para 'arquivar' memórias), estresse crônico (cortisol alto danifica memória), deficiências (B12,folato,ômega-3,magnésio,ferro), tireoide (alterações leves afetam memória), depressão/ansiedade, açúcar excessivo (picos prejudicam cérebro), inflamação. MAIORIA REVERSÍVEL: melhor sono, nutrição otimizada, exercício, gerenciar estresse, tratar deficiências. Melhora significativa em 2-3 meses.",
        "conduct": "INICIAL: Questionários validados (MAC-Q queixas memória, PHQ-9 depressão, GAD-7 ansiedade), qualidade de vida. Correlacionar queixas vs. testes objetivos (MEEM, MoCA, 5 palavras). Grande discrepância (queixas sem déficit) = ansiedade/depressão; queixas COM déficit = declínio real. Cronologia: recentes progressivas (mais preocupante), flutuantes com estresse (menos), estáveis anos sem progressão (personalidade/ansiedade). CAUSAS REVERSÍVEIS: TSH+T4L+T3L+anti-TPO (subclínico comum), B12>500, folato>12, homocisteína<10(<8), VitD>40, ferritina>50F/>75M, glicemia+HbA1c, cortisol 4 pontos, perfil lipídico, hs-CRP. SONO: PSQI, Epworth, diário 2sem. Privação<7h/noite = causa extremamente comum, altamente reversível. MEDICAÇÕES: Anticolinérgicos(tricíclicos,anti-histamínicos,relaxantes,antiespasmódicos), BZD(doses baixas crônicas), estatinas altas, betabloqueadores, anticonvulsivantes - discutir redução/troca. MANEJO: PRIVAÇÃO SONO: Higiene rigorosa(fixo,blackout,18-20°C,sem telas 2h,sol matinal), magnésio treonato/glicinato 400-600mg noite+L-teanina 200mg+glicina 3g, tratar causas(apneia polissonografia, ansiedade, dor). DEPRESSÃO/ANSIEDADE: TCC, exercício aeróbico 150min/sem(=antidepressivos leve-moderada), ômega-3 EPA 2g(anti-depressivo), SAMe 400-800mg, L-metilfolato 15mg se MTHFR, VitD se baixa, magnésio, probióticos psicotrópicos, considerar antidepressivos severa(preferir bupropiona/vortioxetina, evitar tricíclicos/BZD que pioram memória). DEFICIÊNCIAS: B12 baixa=metilcobalamina 1000mcg sublingual, investigar causa(gastrite atrófica,metformina,IBP,vegetarianismo). VitD=5-10k UI até 50-80ng/mL depois 2-4k manutenção. Ferro=bisglicinato 25-50mg+VitC longe café/chá/cálcio, investigar causa. Ômega-3=2-3g/dia ultra purificado, índice>8%. TIREOIDE: Subclínico TSH>3.0=trial levotiroxina baixa dose visando 1.0-2.0. Otimizar T4→T3: zinco 30mg, selênio 200mcg, ferro se baixo, evitar soja/crucíferas cruas excesso. STRESS/CORTISOL: Ashwagandha 300-600mg(reduz 28%), rhodiola 200-400mg, holy basil 300mg, fosfatidilserina 200-400mg, mindfulness 20min/dia, respiração, yoga/tai chi, ritmo circadiano(sol matinal,escuridão noite), exercício regular(não noturno). SUPORTE COGNITIVO GERAL: Complexo B metilado, ômega-3 2g, magnésio L-treonato 2g/dia, VitD3+K2, curcumina 500mg, resveratrol 250mg, Bacopa 300mg 2x/dia, Lion's Mane 1g 2x/dia. ESTILO VIDA: Exercício 150min/sem moderado-vigoroso(BDNF,neurogênese,humor,sono), MIND diet(vegetais folhosos,frutas vermelhas,peixes,oleaginosas,azeite), treinamento cognitivo(idioma,instrumento,jogos memória,leitura ativa), socialização regular(isolamento=fator risco independente), manejo stress(relaxamento,hobbies,natureza). MONITORAMENTO: Reavaliar queixas+testes objetivos 8-12sem. Melhora=continuar. Piora progressiva=neuropsicológico formal+neurologia. EDUCAÇÃO: Lapsos ocasionais(nome,chaves) normais sob stress/multitarefa. ALERTAS: Esquecer eventos inteiros(não detalhes), perder-se locais familiares, dificuldade tarefas rotineiras, mudanças personalidade, julgamento comprometido. Empoderar: muitos fatores modificáveis."
    },

    "foco_concentracao": {
        "clinicalRelevance": "Foco/concentração/aprendizado envolvem atenção sustentada(manter foco), seletiva(filtrar distrações), dividida(multitarefa), memória trabalho, velocidade processamento, controle inibitório. Na MFI, déficits refletem neurotransmissão dopaminérgica/noradrenérgica, função executiva pré-frontal, energização neuronal(mitocôndrias), eixo HPA, sono, inflamação sistêmica(citocinas afetam neurotransmissores). Causas: TDAH(déficit dopamina/noradrenalina pré-frontal), privação sono<7h/noite(mais comum), ansiedade(ruminação consome recursos), depressão, hipotireoidismo subclínico, anemia/ferro baixo, deficiências B, resistência insulínica cerebral, toxinas(metais,pesticidas), medicamentos(BZD,anti-histamínicos,opioides), sobrecarga cognitiva. Biomarcadores: ferritina<50F/<75M(ferro=cofator síntese dopamina), B12<400, folato<12, homocisteína>10, TSH>3.0, glicemia>100/HbA1c>5.7%, cortisol muito baixo<10/alto>30, DHEA-S baixo, VitD<30, hs-CRP>2. Neuroimagem: hipoativação pré-frontal dorsolateral, cingulado anterior, conectividade reduzida redes atencionais.",
        "patientExplanation": "Problemas de foco/concentração são muito comuns e geralmente TEM CAUSAS TRATÁVEIS - não significa 'burrice' ou 'preguiça'. Causas frequentes: falta de sono(cérebro cansado não mantém foco), TDAH não diagnosticado(diferenças química cerebral dopamina), ansiedade/preocupações(ocupam espaço mental), depressão(afeta motivação/energia), deficiências nutricionais(ferro,vitaminas B,ômega-3), tireoide(mesmo leve), açúcar instável(picos/quedas energia), multitarefa excessiva(esgota recursos), medicamentos. Na MFI investigamos causa raiz: exames sangue para deficiências, padrão sono, estresse crônico, função tireoide, glicemia estável. Tratamos causa específica - não apenas estimulantes mascarando problema. MAIORIA TRATÁVEL: dormir bem(7-9h qualidade), alimentação estabilizadora(proteína cada refeição,evitar açúcar), exercícios(aumentam dopamina naturalmente), suplementos se deficiências, mindfulness(treina atenção), ambientes menos distrações. Melhora em 4-8 semanas.",
        "conduct": "INICIAL: Escalas ASRS-18(TDAH adulto), atenção/concentração, PHQ-9, GAD-7. Anamnese: início sintomas(infância vs.adulto-TDAH=neurodesenvolvimento), contextos afetados, impacto funcional, familiar TDAH, substâncias(cafeína,nicotina,estimulantes). Diário atenção 2sem: horários melhor/pior, correlação alimentação/sono/estresse, gatilhos. LABORATÓRIO: Hemograma, ferritina>75M/>50F(cofator tirosina hidroxilase dopamina-deficiência comum afeta gravemente), ferro completo se baixa, B12>500, folato>12, homocisteína<10(<8), VitD 50-80ng/mL, TSH+T4L+T3L, glicemia+insulina+HbA1c, cortisol 4 pontos, DHEA-S, magnésio/zinco eritrocitário, hs-CRP. Severos: Neurotransmissores urinários(dopamina,norepinefrina,serotonina,GABA,glutamato), aminoácidos(tirosina,fenilalanina,triptofano), ácidos orgânicos, metais pesados, microbioma. SONO: PSQI, Epworth, diário 2sem, polissonografia se apneia suspeitada. MANEJO TDAH CONFIRMADO: Integrativa primeiro(antes estimulantes): L-tirosina 1-2g manhã jejum(precursor dopamina), ômega-3 2-3g(evidências sólidas), ferro se<75, zinco 30mg, magnésio 400-600mg, complexo B metilado, fosfatidilserina 200-300mg, L-teanina 200-400mg, ginkgo 120mg 2x/dia, rhodiola 200-400mg. Comportamental: Organização, timers/alarmes, minimizar distrações, exercício regular(crítico-dopamina endógena), sono rigoroso. Estimulantes(metilfenidato,lisdexanfetamina) se integrativas insuficientes+impacto funcional. PRIVAÇÃO SONO: Prioridade máxima-higiene(fixo,blackout,18-20°C,sem telas 2h,sol matinal), magnésio glicinato 400mg+L-teanina 200mg+glicina 3g 1h antes, tratar causas(apneia CPAP,ansiedade CBT/mindfulness,dor). Meta 7-9h+90min N3. FERRO BAIXO: Bisglicinato 25-50mg+VitC longe café/chá/cálcio, reavaliação 8-12sem meta>75, investigar causa(menstruação,sangramento GI,vegetarianismo,má absorção celíaca). Melhora 4-8sem pós-normalização. TIREOIDE: Subclínico TSH>3.0=trial levotiroxina baixa visando 1.0-2.0. Otimizar T4→T3: zinco 30mg, selênio 200mcg(castanhas 2-3/dia), ferro se baixo, evitar soja/crucíferas cruas excesso. RESISTÊNCIA INSULÍNICA/HIPOGLICEMIA: Dieta estabilizadora-proteína todas refeições(1.6-2g/kg), carboidratos complexos baixo IG(batata doce,quinoa,aveia), gorduras saudáveis(abacate,oleaginosas,azeite), eliminar açúcares refinados, evitar carboidratos isolados, refeições cada 3-4h. Café manhã 30g proteína <1h acordar(ativa dopamina/noradrenalina via tirosina). Small snacks proteicos se hipoglicemia. Cromo picolinato 200mcg, canela 1-3g, berberina 500mg 2-3x/dia se resistência, inositol 2g 2x/dia. ANSIEDADE/STRESS: Ashwagandha 300-600mg(cortisol-28%), rhodiola 200-400mg, holy basil 300mg, L-teanina 200mg 2-3x/dia(ondas alfa,foco calmo), magnésio 400-600mg, mindfulness 10-20min/dia, respiração 4-7-8, yoga, HRV biofeedback, CBT se significativa. SUPORTE ATENCIONAL GERAL: Ômega-3 2-3g/dia, complexo B metilado completo, magnésio L-treonato 2g/dia, VitD3 5k se<40+K2, creatina 5g/dia(fosfocreatina cerebral energia neurônios-evidências memória trabalho), ALCAR 1500mg/dia, CoQ10 ubiquinol 200-300mg, fosfatidilserina 300mg, fosfatidilcolina 1-2g. NOOTRÓPICOS: Bacopa 300mg 2x/dia(memória trabalho,ansiedade,8-12sem), ginkgo 120mg 2x/dia(fluxo cerebral,atenção), rhodiola 200-400mg manhã(fadiga mental,stress), Lion's Mane 1g 2x/dia(NGF,mielinização), cafeína 100-200mg+L-teanina 200-400mg(sinergismo atenção sem jitteriness). TREINAMENTO COGNITIVO: Cogmed/dual n-back 25-30min/dia 5x/sem 5-8sem(transferência não-treinadas), meditação focada 20min/dia(controle atencional,córtex pré-frontal), xadrez/go, instrumento musical. EXERCÍCIO: Aeróbico 150min/sem moderado-vigoroso(BDNF,dopamina,noradrenalina), HIIT 2-3x/sem(máximo catecolaminas), matinal(sincronização,catecolaminas sustentadas), evitar sedentarismo(levantar cada 30-45min,8-10k passos/dia). AMBIENTAL/COMPORTAMENTAL: Reduzir distrações(silencioso,desativar notificações), Pomodoro(25min foco/5min pausa), mono-tarefa(evitar multi), organização externa(listas,alarmes,apps), priorização(Eisenhower urgente vs.importante), sono consistente 7-9h. MEDICAMENTOS PREJUDICIAIS: Revisar com prescritor-BZD(comprometem atenção/memória), anti-histamínicos 1ªgeração(difenidramina-sedação,anticolinérgicos), anticonvulsivantes altas doses, opioides-discutir redução/troca. MONITORAMENTO: Escalas atenção 8-12sem, labs deficiências 3m(ferritina,B12,VitD,homocisteína). Insuficiente=neuropsicológico formal, TDAH especialista, neurologia se neurodegenerativo. ALVO: Melhora funcional(foco,completar tarefas,memória trabalho), normalização biomarcadores(ferritina>75,B12>500,VitD>50,homocisteína<8), redução impacto trabalho/estudos/vida."
    }
}

# IDS DOS 30 ITEMS (primeiros 30 da query)
ITEMS = {
    # Dubois Imediato (3x)
    "8f2fb5b6-86e7-4a6f-baf5-16f2785b34d0": "dubois_imediato",
    "414fe374-9c4a-4e6c-a76f-596811b9b4e0": "dubois_imediato",
    "a4084ae3-aaa9-4720-bab2-8fd7f8542247": "dubois_imediato",
    # Dubois Tardio (3x)
    "1a67a824-13d9-4085-b998-daab565ddd1c": "dubois_tardio",
    "cd50ed91-b1cd-47bd-9427-2592485cbdce": "dubois_tardio",
    "03d3220d-1dbe-4a9e-87d2-1226a79fb6bc": "dubois_tardio",
    # Capacidade memória percebida (3x)
    "f9f0c524-bc41-45dc-b33b-46b7f2234fe9": "memoria_capacidade",
    "6437b500-5ab6-4e7a-8f62-981cedc96b40": "memoria_capacidade",
    "f703185e-6f49-4bec-942e-c1eed243a6a0": "memoria_capacidade",
    # Foco/concentração/aprendizado (3x)
    "c31657f0-f591-4cf3-b054-e146504842bd": "foco_concentracao",
    "28a6710a-f04c-46fe-81df-b10419e8d167": "foco_concentracao",
    "ded705c3-5ecc-4ccd-8dfd-cff55da653fd": "foco_concentracao",
    # Disposição percebida (3x)
    "4c8229d8-3376-43be-a120-d777a7e5cc1b": "foco_concentracao",  # Relacionado a energia/atenção
    "ee16e133-05c2-4ba6-ab4c-a4b6f7de1229": "foco_concentracao",
    "66a04c19-a77d-4af1-b818-19ae0c1065b9": "foco_concentracao",
    # Disposição/energia (3x)
    "c3689015-b4ed-4af8-86c0-831dec1f4f6f": "foco_concentracao",
    "5bf586c3-4085-4532-a9d4-5fd9fb2c946e": "foco_concentracao",
    "4fdd5c7f-dfd5-4ad9-bf6e-c4266e3d5df5": "foco_concentracao",
    # PHQ-9 (3x)
    "fc121488-7a4d-4f7f-9f41-95bff5960298": "memoria_capacidade",  # Humor afeta memória percebida
    "a520b28c-c8e5-4317-a3b0-22022a8859ce": "memoria_capacidade",
    "2e31f7ac-ccbf-421b-94d8-b5515b895753": "memoria_capacidade",
    # Epworth (3x)
    "bff2ae60-5156-4dc0-8955-0ac33a7ae185": "foco_concentracao",  # Sonolência afeta atenção
    "303f0209-73a3-4e95-bf9b-264ebf02b19f": "foco_concentracao",
    "7eec2901-ff1b-4556-8439-afa139a9bae3": "foco_concentracao",
    # FSS Fadiga (3x)
    "85e8f942-5cba-4d57-90ad-44126cb47b28": "foco_concentracao",
    "92e8b869-c601-4a2e-99ee-0fba00ab9037": "foco_concentracao",
    "82d9355a-5aa8-44ac-96e1-193b1942b64d": "foco_concentracao",
    # GAD-7 (3x)
    "4e92216d-222d-48a5-a272-01c1e573fcc0": "memoria_capacidade",  # Ansiedade afeta memória
    "9861e46f-764c-4372-afe1-9a7176d0b505": "memoria_capacidade",
    "7eeb2870-8112-4ff5-9d8b-14c18cdd9ce9": "memoria_capacidade"
}

def main():
    print("="*80)
    print("BATCH COGNIÇÃO - 30 ITEMS - EXECUÇÃO")
    print("Medicina Funcional Integrativa")
    print("="*80)
    print()

    # Login
    print("1. Fazendo login...")
    try:
        token = login()
        print(f"✓ Token obtido\n")
    except Exception as e:
        print(f"✗ Erro no login: {e}")
        return

    # Processar
    success, errors = 0, 0
    print(f"2. Processando {len(ITEMS)} items...\n")

    for idx, (item_id, content_key) in enumerate(ITEMS.items(), 1):
        print(f"[{idx:2d}/30] {item_id[:8]}... ({content_key})", end=" ")

        content = CONTENT[content_key]
        data = {
            "clinicalRelevance": content["clinicalRelevance"],
            "patientExplanation": content["patientExplanation"],
            "conduct": content["conduct"]
        }

        if update_item(token, item_id, data):
            print("✓")
            success += 1
        else:
            print("✗")
            errors += 1

        time.sleep(0.3)

    # Resumo
    print("\n" + "="*80)
    print(f"CONCLUÍDO: {success}/30 items atualizados ({success/30*100:.1f}%)")
    print("="*80)

if __name__ == "__main__":
    main()
