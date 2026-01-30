#!/usr/bin/env python3
"""
Script para enriquecer 15 Score Items de MEDICAMENTOS - Parte 2
Conteúdo clínico estruturado e validado
"""

import requests
import json
import time
from typing import Dict

# Configuração
API_BASE_URL = "http://localhost:3001/api/v1"
EMAIL = "import@plenya.com"
PASSWORD = "Import@123456"

# Base de conhecimento clínico estruturado
MEDICATION_CONTENT = {
    "84d98093-e62f-470f-9e08-6ab525f7d357": {  # Antiarrítmicos
        "name": "Antiarrítmicos",
        "clinical_relevance": """O histórico de uso de antiarrítmicos representa marcador de cardiopatia estrutural ou elétrica subjacente, indicando presença de arritmias cardíacas clinicamente significativas que requereram intervenção farmacológica. Na medicina funcional integrativa, os antiarrítmicos são reconhecidos como terapias essenciais para controle de arritmias potencialmente letais, mas a abordagem enfatiza investigação e correção de fatores desencadeantes modificáveis.

As principais classes de antiarrítmicos (Classificação de Vaughan Williams):

CLASSE IA (quinidina, procainamida, disopiramida): bloqueiam canais de sódio de forma moderada, prolongam QT. Uso limitado atualmente devido a efeitos pró-arrítmicos. Efeitos adversos: prolongamento QT com risco de torsades de pointes, efeitos anticolinérgicos (boca seca, visão turva, retenção urinária), náuseas, diarreia (quinidina), síndrome lúpus-like (procainamida).

CLASSE IB (lidocaína, mexiletina): bloqueiam canais de sódio rapidamente, encurtam QT. Lidocaína é usada IV em emergências (TV/FV). Mexiletina é via oral. Efeitos adversos: tontura, tremor, parestesias, confusão mental, convulsões (doses altas).

CLASSE IC (flecainida, propafenona): bloqueio intenso de canais de sódio. CONTRAINDICADOS em cardiopatia estrutural (estudo CAST mostrou aumento de mortalidade pós-IAM). Uso restrito a FA em pacientes sem cardiopatia. Efeitos adversos: pró-arritmia (especialmente TV), tontura, visão turva, broncoespasmo (propafenona).

CLASSE II (betabloqueadores): já descritos em anti-hipertensivos. Essenciais em arritmias adrenérgicas, prevenção de morte súbita pós-IAM.

CLASSE III (amiodarona, sotalol, dofetilida, ibutilida): bloqueiam canais de potássio, prolongam QT. AMIODARONA é o mais usado, altamente eficaz mas toxicidade significativa. Efeitos adversos da amiodarona: disfunção tireoidiana (hipo ou hipertireoidismo - 15-20%), toxicidade pulmonar (fibrose pulmonar - pode ser fatal), hepatotoxicidade, fotossensibilidade cutânea, depósitos corneanos, neuropatia periférica, coloração azul-acinzentada da pele, prolongamento QT. SOTALOL: combina efeito betabloqueador + prolongamento QT, risco de torsades de pointes, contraindicado se ClCr <40mL/min.

CLASSE IV (diltiazem, verapamil): bloqueadores de canais de cálcio não di-hidropiridínicos. Úteis em taquicardias supraventriculares. Efeitos adversos: bradicardia, bloqueio AV, hipotensão, constipação (verapamil), edema.

OUTROS: DIGOXINA (glicosídeo cardíaco) - inibe bomba Na-K-ATPase, aumenta contratilidade e reduz condução AV. Janela terapêutica estreita (0,5-2,0ng/mL). Toxicidade digitalica: náuseas, vômitos, confusão, visão amarelada, arritmias (extrassístoles ventriculares, bloqueios AV, taquicardia atrial com bloqueio). Fatores de risco: hipocalemia, hipomagnesemia, insuficiência renal, idosos. ADENOSINA (uso agudo IV para reversão de taquicardias supraventriculares).

A medicina funcional investiga causas subjacentes de arritmias: distúrbios eletrolíticos (magnésio, potássio, cálcio), disfunção tireoidiana (hipo e hipertireoidismo), apneia obstrutiva do sono (hipóxia intermitente → hiperatividade simpática), consumo excessivo de estimulantes (cafeína, álcool, nicotina), estresse crônico (excesso catecolaminérgico), inflamação sistêmica, desequilíbrio autonômico, deficiências nutricionais (magnésio, CoQ10, L-carnitina, taurina), toxicidade medicamentosa, cardiopatia estrutural subjacente.""",

        "patient_explanation": """Antiarrítmicos são medicamentos usados para controlar batimentos cardíacos irregulares (arritmias). Quando seu coração bate muito rápido, muito devagar ou de forma desorganizada, esses medicamentos ajudam a restaurar um ritmo mais normal e prevenir complicações graves.

Existem vários tipos de antiarrítmicos, cada um funcionando de forma diferente:
- Alguns tornam o coração menos "excitável"
- Outros controlam a velocidade dos batimentos
- Alguns previnem ritmos perigosos que podem causar morte súbita

O medicamento mais comum é a amiodarona, muito eficaz mas que requer monitoramento cuidadoso porque pode afetar tireoide, pulmões e fígado. Betabloqueadores também são frequentemente usados para arritmias, especialmente após infarto.

É FUNDAMENTAL tomar esses medicamentos exatamente como prescrito - nunca pule doses ou interrompa por conta própria, pois isso pode desencadear arritmias graves. Muitos antiarrítmicos têm "janela terapêutica estreita", ou seja, a dose que funciona e a dose que causa problemas são próximas.

Efeitos colaterais variam conforme o medicamento mas podem incluir cansaço, tontura, problemas na tireoide (amiodarona), batimentos muito lentos. Procure seu médico imediatamente se sentir: desmaios, tontura intensa, falta de ar súbita, batimentos muito rápidos ou irregulares.""",

        "conduct": """1. Avaliação cardiológica especializada:
   - ECG basal e comparativo
   - Holter 24-48h (avaliar carga arrítmica)
   - Ecocardiograma (descartar cardiopatia estrutural)
   - Teste ergométrico (arritmias induzidas por esforço)

2. Monitoramento medicamentoso específico:

   AMIODARONA:
   - Função tireoidiana (TSH, T4L) pré-tratamento e a cada 6 meses
   - Radiografia de tórax e PFTs baseline + anual
   - Transaminases hepáticas baseline + trimestral
   - Exame oftalmológico anual
   - ECG (QTc - risco de torsades se >500ms)
   - Fotoproteção solar

   DIGOXINA:
   - Digoxinemia (alvo 0,5-2,0ng/mL)
   - Eletrólitos (K, Mg) - corrigir hipocalemia/hipomagnesemia
   - Função renal (ajustar dose)
   - ECG (avaliar intoxicação digitalica)

   SOTALOL/DOFETILIDA:
   - QTc (risco torsades se >500ms)
   - Eletrólitos (K, Mg)
   - Função renal (contraindicado se ClCr baixa)

3. Investigação de causas subjacentes:
   - Eletrólitos séricos (magnésio, potássio, cálcio)
   - Função tireoidiana completa (TSH, T4L, T3L)
   - Polissonografia (rastreamento apneia do sono)
   - Ecocardiograma (avaliação estrutural)

4. Correção de distúrbios eletrolíticos:
   - Magnésio sérico >2,0mg/dL (suplementar se <2,0)
   - Potássio 4,0-5,0mEq/L

5. Modulação do estilo de vida:
   - Redução/cessação de cafeína (café, energéticos, chás)
   - Moderação ou cessação de álcool
   - Cessação de tabagismo
   - Manejo de estresse (técnicas de respiração, meditação)
   - Exercícios regulares sob supervisão médica

6. Suplementação funcional baseada em evidências:
   - Magnésio (400-600mg/dia) - ESSENCIAL
   - Ômega-3 (2-3g/dia EPA+DHA) - propriedades antiarrítmicas
   - Coenzima Q10 (100-200mg/dia)
   - Taurina (1-3g/dia) - estabilizador de membrana
   - L-carnitina (1-2g/dia)
   - Potássio (através da dieta - frutas, vegetais)

7. Tratamento de comorbidades:
   - Otimização de hipotireoidismo/hipertireoidismo
   - Controle rigoroso de hipertensão
   - Tratamento de apneia obstrutiva do sono (CPAP)
   - Controle de diabetes

8. Educação do paciente:
   - Sinais de toxicidade medicamentosa
   - Importância da adesão estrita
   - Nunca suspender abruptamente (especialmente betabloqueadores)
   - Quando procurar atendimento de emergência

9. Monitoramento regular com cardiologista (3-6 meses) e ajustes conforme resposta clínica e carga arrítmica"""
    },

    "1c173e0e-3f0d-48a9-805d-b206544d27b9": {  # Antibióticos prolongados
        "name": "Antibióticos de uso prolongado",
        "clinical_relevance": """O histórico de uso prolongado de antibióticos (>7-14 dias ou uso repetitivo) representa marcador de infecções bacterianas recorrentes ou crônicas, condições que exigem antibioticoterapia estendida (tuberculose, osteomielite, endocardite) ou possível uso inadequado/excessivo. Na medicina funcional integrativa, o uso prolongado de antibióticos é reconhecido como intervenção necessária em infecções graves, mas a abordagem enfatiza consequências metabólicas, imunológicas e microbiológicas significativas que requerem manejo proativo.

IMPACTOS DO USO PROLONGADO DE ANTIBIÓTICOS:

DISBIOSE INTESTINAL: Antibióticos de amplo espectro destroem tanto bactérias patogênicas quanto simbióticas benéficas, causando depleção da diversidade microbiana intestinal. Consequências: diarreia associada a antibióticos (10-30% dos casos), infecção por Clostridioides difficile (especialmente fluoroquinolonas, clindamicina, cefalosporinas), síndrome do intestino irritável pós-infecciosa, sobrecrescimento de Candida, aumento de permeabilidade intestinal ("leaky gut"), redução de produção de ácidos graxos de cadeia curta (butirato), comprometimento da síntese de vitaminas (K2, B12).

RESISTÊNCIA BACTERIANA: Uso repetido ou prolongado favorece seleção de cepas resistentes (MRSA, bactérias produtoras de ESBL, Pseudomonas multirresistentes), comprometendo opções terapêuticas futuras tanto individuais quanto comunitárias.

DEFICIÊNCIAS NUTRICIONAIS INDUZIDAS: Tetraciclinas quelam cálcio, magnésio, ferro, zinco (reduzindo absorção). Fluoroquinolonas aumentam risco de tendinopatia/ruptura tendínea (especialmente tendão de Aquiles) e neuropatia periférica. Sulfonamidas e trimetoprima inibem metabolismo de folato. Isoniazida (tuberculose) depleta piridoxina (B6).

IMUNOMODULAÇÃO: Alterações do microbioma intestinal comprometem 70% do sistema imunológico associado ao intestino (GALT), reduzindo resistência a infecções futuras e favorecendo recorrências.

TOXICIDADES ESPECÍFICAS: Aminoglicosídeos (gentamicina, amicacina) - nefrotoxicidade e ototoxicidade (irreversível). Vancomicina - nefrotoxicidade, "síndrome do homem vermelho". Fluoroquinolonas - tendinopatia, neuropatia periférica, prolongamento QT, fotossensibilidade, risco de dissecção de aorta. Macrolídeos - prolongamento QT (azitromicina, claritromicina), hepatotoxicidade. Metronidazol - neuropatia periférica (uso >2 semanas), efeito dissulfiram (evitar álcool).

CONDIÇÕES QUE EXIGEM ANTIBIOTICOTERAPIA PROLONGADA: Tuberculose (6-9 meses RIPE - rifampicina, isoniazida, pirazinamida, etambutol), endocardite bacteriana (4-6 semanas IV), osteomielite (6-12 semanas), infecções protéticas ortopédicas, abscesso hepático/esplênico, actinomicose, brucelose, doença de Lyme disseminada, infecções por micobactérias atípicas (MAC), profilaxia de febre reumática recorrente, profilaxia de ITU recorrente (controverso).

A medicina funcional enfatiza: uso racional de antibióticos (apenas quando bacteriano confirmado, evitar em viroses), escolha de espectro mais estreito possível, duração adequada mas não excessiva, proteção da microbiota durante tratamento, repovoamento pós-tratamento, investigação de causas de infecções recorrentes (imunodeficiências, alterações anatômicas, biofilmes bacterianos), otimização da imunidade através de nutrição, sono, manejo de estresse.""",

        "patient_explanation": """Antibióticos são medicamentos que matam bactérias ou impedem seu crescimento, salvando vidas em infecções graves. No entanto, quando usados por períodos prolongados (semanas ou meses) ou de forma repetitiva, podem causar efeitos indesejados que precisam ser gerenciados.

O principal problema do uso prolongado de antibióticos é a destruição da flora intestinal benéfica (microbiota). Nosso intestino abriga trilhões de bactérias "boas" que ajudam na digestão, produção de vitaminas, proteção contra infecções e fortalecimento do sistema imunológico. Antibióticos não distinguem bactérias boas de ruins, eliminando ambas.

Consequências comuns do uso prolongado:
- Diarreia (10-30% dos casos)
- Infecção por Clostridium difficile (bactéria oportunista que causa diarreia grave)
- Candidíase vaginal ou oral (proliferação de fungos)
- Maior susceptibilidade a infecções futuras
- Problemas digestivos persistentes

Alguns antibióticos têm riscos específicos: fluoroquinolonas (como ciprofloxacino, levofloxacino) podem causar ruptura de tendões, formigamento nas mãos/pés e problemas nos nervos. Aminoglicosídeos podem afetar rins e audição.

SE VOCÊ PRECISA DE ANTIBIÓTICOS PROLONGADOS (tuberculose, infecção óssea, etc.), é ESSENCIAL:
- Tomar exatamente como prescrito, sem pular doses
- Nunca interromper antes do prazo (mesmo se sentir melhor)
- Proteger seu intestino com probióticos
- Manter hidratação adequada
- Reportar efeitos colaterais ao médico
- Fazer exames de monitoramento conforme solicitado

Procure seu médico imediatamente se tiver: diarreia intensa ou com sangue, erupções cutâneas, inchaço na face ou língua, dor intensa em tendões, formigamento persistente.""",

        "conduct": """1. Documentação e justificativa:
   - Confirmar diagnóstico bacteriano (cultura + antibiograma quando possível)
   - Documentar duração necessária conforme protocolo
   - Reavaliar periodicamente necessidade de continuidade

2. Escolha racional do antibiótico:
   - Espectro mais estreito possível (evitar "canhão para matar mosca")
   - Considerar perfil de resistência local
   - Avaliar via oral vs IV (trocar para VO quando possível)

3. Proteção da microbiota durante tratamento:
   - Probióticos de alta potência (≥50 bilhões UFC/dia)
   - Cepas específicas: Saccharomyces boulardii (reduz diarreia associada a ATB)
   - Lactobacillus rhamnosus GG, Lactobacillus casei
   - Tomar probiótico 2-3h separado do antibiótico
   - Continuar por 2-4 semanas após fim do antibiótico

4. Monitoramento de toxicidade específica:

   AMINOGLICOSÍDEOS (gentamicina, amicacina):
   - Função renal (creatinina) 2x/semana
   - Níveis séricos de pico e vale
   - Audiometria baseline e durante tratamento

   FLUOROQUINOLONAS:
   - Orientar sobre risco de tendinopatia (suspender se dor)
   - Evitar exercícios intensos durante tratamento
   - Monitorar sintomas neurológicos

   VANCOMICINA:
   - Função renal
   - Níveis séricos (vale 15-20mcg/mL)

   TUBERCULOSTÁTICOS:
   - Transaminases hepáticas mensais (hepatotoxicidade)
   - Piridoxina (B6) 25-50mg/dia com isoniazida
   - Função renal
   - Acuidade visual (etambutol - neurite óptica)

5. Prevenção de deficiências nutricionais:
   - Tetraciclinas: suplementar cálcio, magnésio (2h separado)
   - Isoniazida: piridoxina (B6) 25-50mg/dia
   - Trimetoprima: ácido fólico (5mg/dia)
   - Geral: multivitamínico durante tratamento prolongado

6. Prevenção de Clostridioides difficile:
   - Evitar inibidores de bomba de prótons (IBP) desnecessários
   - Probióticos profiláticos
   - Higiene rigorosa (lavar mãos com água e sabão, álcool gel não mata esporos)

7. Repovoamento pós-antibiótico (após término):
   - Probióticos de múltiplas cepas (30-100 bilhões UFC/dia) por 4-8 semanas
   - Prebióticos (fibras fermentáveis): inulina, FOS, acácia
   - Alimentos fermentados (kefir, chucrute, kimchi, kombucha)
   - Dieta rica em fibras diversificadas

8. Suporte imunológico:
   - Vitamina D (manter >40ng/mL)
   - Vitamina C (1-2g/dia durante infecção)
   - Zinco (15-30mg/dia)
   - Ômega-3 (anti-inflamatório)
   - Glutamina (5-10g/dia) - suporte da mucosa intestinal

9. Investigação de infecções recorrentes:
   - Avaliação imunológica (dosagem de imunoglobulinas)
   - Investigação de alterações anatômicas
   - Avaliação de comorbidades (diabetes, HIV)
   - Otimização nutricional e metabólica

10. Educação do paciente:
    - Importância da adesão completa
    - Reconhecimento de efeitos adversos
    - Não compartilhar antibióticos
    - Não usar "sobras" de tratamentos anteriores
    - Probióticos durante e após tratamento

11. Monitoramento clínico regular conforme condição específica"""
    },

    "9bc0c04e-dffe-4356-b350-182b7d4fc25a": {  # Anticoagulantes
        "name": "Anticoagulantes",
        "clinical_relevance": """O histórico de uso de anticoagulantes indica presença de condições tromboembólicas (fibrilação atrial, tromboembolismo venoso, próteses valvares) ou alto risco de eventos tromboembólicos, representando situações clínicas de morbimortalidade significativa. Na medicina funcional integrativa, os anticoagulantes são reconhecidos como terapias salva-vidas essenciais, mas a abordagem enfatiza compreensão dos estados pró-trombóticos subjacentes, otimização da terapia e prevenção de complicações hemorrágicas.

PRINCIPAIS ANTICOAGULANTES:

ANTAGONISTAS DA VITAMINA K (warfarina): Inibem reciclagem de vitamina K, reduzindo síntese hepática dos fatores de coagulação II, VII, IX, X. Monitora nto via RNI/INR (Razão Normalizada Internacional). Alvo terapêutico varia conforme indicação: FA 2,0-3,0; prótese valvar metálica 2,5-3,5. Janela terapêutica estreita com alta variabilidade interindividual (genética CYP2C9, VKORC1; interações medicamentosas extensas; interações dietéticas com vitamina K). Efeitos adversos: hemorragia (major 1-3%/ano), necrose cutânea (deficiência de proteína C/S), síndrome do dedo roxo, osteoporose (uso prolongado), calcificação vascular paradoxal. Interações: antibióticos (alteram flora produtora de vit K), antifúngicos, amiodarona, AINES, inibidores de CYP2C9. Alimentos ricos em vitamina K (vegetais verde-escuros, fígado, natto) reduzem efeito; manter ingestão consistente.

ANTICOAGULANTES ORAIS DIRETOS (DOACs/NOACs): Dabigatrana (inibidor direto da trombina), rivaroxabana, apixabana, edoxabana (inibidores do fator Xa). Vantagens: dose fixa, sem monitoramento rotineiro de coagulação, menos interações alimentares, início rápido de ação, meia-vida curta. Desvantagens: custo elevado, contraindicados em insuficiência renal grave (depuração renal significativa), risco hemorrágico sem antídoto universal facilmente disponível (idarucizumab para dabigatrana, andexanet alfa para anti-Xa - caros e pouco disponíveis), adesão crítica (esquecer dose = desanticoagulação rápida). Efeitos adversos: hemorragia (similar ou ligeiramente menor que warfarina), dispepsia (dabigatrana 10-30%), interação com inibidores/indutores de P-gp e CYP3A4.

HEPARINAS (heparina não-fracionada HNF, heparinas de baixo peso molecular HBPM - enoxaparina, dalteparina): Uso parenteral (SC ou IV). HNF para situações agudas (TEP, IAM, AVC isquêmico agudo) com monitoramento via aPTT ou anti-Xa. HBPM para profilaxia de TEV, tratamento ambulatorial de TVP. Efeitos adversos: hemorragia, trombocitopenia induzida por heparina (HIT - grave, requer suspensão imediata e anticoagulante alternativo), osteoporose (uso >3 meses), reações alérgicas. HIT ocorre em 1-3% (HNF) e 0,1-1% (HBPM), pode causar trombose paradoxal grave.

INDICAÇÕES PRINCIPAIS: Fibrilação atrial (FA) - prevenção de AVC embólico (CHA2DS2-VASc ≥2 homens, ≥3 mulheres). Tromboembolismo venoso (TVP, TEP) - tratamento agudo (3-6 meses) e profilaxia secundária em trombofilia. Próteses valvares metálicas (warfarina obrigatória). Síndrome coronariana aguda (heparina fase aguda). Prevenção de TEV em situações de alto risco (cirurgias ortopédicas, imobilização prolongada, câncer ativo).

COMPLICAÇÕES PRINCIPAIS: Hemorragia (major: intracraniana, gastrointestinal, necessidade de transfusão). Fatores de risco para sangramento: idade >75 anos, história de sangramento prévio, insuficiência renal, uso concomitante de antiagregantes plaquetários ou AINES, hipertensão não controlada, queda frequente, alcoolismo, anemia, trombocitopenia. Escore HAS-BLED avalia risco de sangramento.

A medicina funcional investiga estados pró-trombóticos subjacentes: trombofilias hereditárias (Fator V Leiden, mutação da protrombina G20210A, deficiência de antitrombina III/proteína C/S), síndrome antifosfolípide, hiperhomocisteinemia (deficiência de B6/B12/folato, variante MTHFR), níveis elevados de lipoproteína(a), inflamação sistêmica crônica, obesidade, resistência à insulina, estase venosa, desidratação crônica.""",

        "patient_explanation": """Anticoagulantes são medicamentos que "afinam" o sangue, reduzindo sua capacidade de formar coágulos. Embora os coágulos sejam essenciais para estancar sangramentos quando nos machucamos, em certas condições eles podem se formar dentro de vasos sanguíneos e causar problemas graves como derrame (AVC), embolia pulmonar ou trombose nas pernas.

Você pode estar usando anticoagulante porque tem:
- Fibrilação atrial (batimento irregular do coração que facilita formação de coágulos)
- Histórico de trombose nas pernas ou pulmões
- Válvula cardíaca artificial
- Alto risco de formar coágulos por outros motivos

TIPOS DE ANTICOAGULANTES:
- Warfarina (Marevan, Coumadin): mais antigo, requer exames mensais de sangue (INR) para ajustar dose. Interage com muitos alimentos e medicamentos
- Anticoagulantes modernos (Rivaroxabana/Xarelto, Apixabana/Eliquis, Dabigatrana/Pradaxa, Edoxabana/Lixiana): dose fixa, sem necessidade de exames frequentes, mas mais caros

O MAIOR RISCO dos anticoagulantes é sangramento. Por isso é FUNDAMENTAL:
- Tomar exatamente como prescrito, no mesmo horário todos os dias
- NUNCA pular doses ou tomar dose dupla
- Se usar warfarina: manter quantidade de vegetais verde-escuros (couve, espinafre, brócolis) constante na dieta
- Evitar anti-inflamatórios (ibuprofeno, diclofenaco) - use paracetamol se precisar de analgésico
- Informar TODOS os médicos e dentistas que você usa anticoagulante
- Cuidado extra para evitar cortes e quedas
- Usar escova de dentes macia

PROCURE EMERGÊNCIA IMEDIATAMENTE se tiver: sangramento que não para, sangue nas fezes ou urina preta, vômito com sangue, tosse com sangue, dor de cabeça súbita e intensa, tontura extrema, queda ou trauma na cabeça.""",

        "conduct": """1. Avaliação inicial e indicação:
   - Confirmar indicação (FA com CHA2DS2-VASc ≥2, TEV, prótese valvar)
   - Avaliar risco hemorrágico (escore HAS-BLED)
   - Função renal (clearance de creatinina - crucial para DOACs)
   - Função hepática
   - Hemograma completo

2. Escolha do anticoagulante:
   WARFARINA:
   - Prótese valvar metálica (única opção)
   - Insuficiência renal grave (ClCr <15-30mL/min)
   - Custo-efetividade

   DOACs (preferidos em FA não-valvar):
   - Apixabana ou rivaroxabana se ClCr >30mL/min
   - Dose ajustada conforme função renal, idade, peso
   - Contraindicados se ClCr <15-30mL/min

3. Monitoramento - WARFARINA:
   - INR semanal até estabilização, depois mensal
   - Alvo: FA 2,0-3,0; prótese valvar metálica 2,5-3,5
   - Ajustes de dose conforme protocolo
   - Revisar medicamentos e dieta a cada consulta

4. Monitoramento - DOACs:
   - Função renal (creatinina) a cada 6-12 meses (mais frequente se >75 anos ou ClCr limítrofe)
   - Função hepática anual
   - Hemograma anual
   - NÃO requer monitoramento de coagulação rotineiro

5. Prevenção de deficiências nutricionais:
   WARFARINA:
   - Vitamina K2 (MK-7) - controverso, mas pode melhorar estabilidade do INR em doses baixas (45-90mcg/dia)
   - Vitamina D (protege contra calcificação vascular)
   - Magnésio

   HEPARINA prolongada:
   - Vitamina D + cálcio (prevenir osteoporose)
   - Monitorar densidade óssea se uso >3 meses

6. Prevenção de complicações hemorrágicas:
   - Educação intensiva do paciente
   - Evitar AINEs (preferir paracetamol)
   - Evitar quedas (fisioterapia, exercícios de equilíbrio em idosos)
   - Controle rigoroso de hipertensão
   - Tratar anemia e plaquetopenia se presentes
   - Evitar injeções intramusculares

7. Suplementação funcional compatível:
   - Ômega-3: CUIDADO - pode potencializar sangramento em doses altas (>3g/dia). Doses moderadas (1-2g/dia) geralmente seguras mas discutir com médico
   - Vitamina D: seguro e recomendado
   - Magnésio: seguro
   - Coenzima Q10: seguro

   EVITAR/CAUTELA:
   - Vitamina E (>400UI - aumenta sangramento)
   - Alho em suplementos (aumenta sangramento)
   - Gengibre em doses altas
   - Ginkgo biloba (aumenta sangramento)
   - Cúrcuma em doses altas
   - Açafrão
   - Vitamina K (antagoniza warfarina)

8. Investigação de causas pró-trombóticas (especialmente em TEV não provocado):
   - Painel de trombofilia (Fator V Leiden, mutação protrombina)
   - Anticorpos antifosfolípides
   - Homocisteína
   - Lipoproteína (a)
   - Proteínas C, S, antitrombina III

9. Redução de fatores de risco modificáveis:
   - Controle de peso
   - Mobilização precoce pós-cirúrgica
   - Hidratação adequada
   - Cessação de tabagismo
   - Uso de meias de compressão (se TVP prévia)

10. Tratamento de hiperhomocisteinemia (se presente):
    - Vitaminas B6 (50mg), B12 (1000mcg), folato (1-5mg)
    - Meta: homocisteína <10μmol/L

11. Planejamento pré-cirúrgico/procedimentos:
    - WARFARINA: suspender 5 dias antes, bridging com HBPM se alto risco
    - DOACs: suspender 1-4 dias antes conforme função renal e risco cirúrgico
    - Coordenar com cirurgião e cardiologista

12. Reversão de emergência:
    - WARFARINA: vitamina K IV (10mg) + plasma fresco ou concentrado de complexo protrombínico
    - DABIGATRANA: idarucizumab
    - Anti-Xa: andexanet alfa (ou concentrado de complexo protrombínico)

13. Reavaliação periódica (6-12 meses):
    - Necessidade de continuidade
    - Risco-benefício individualizado
    - Ajustes conforme mudanças clínicas"""
    },

    "36575fbe-db94-4640-b590-3019cc7e12c2": {  # Anticoncepcionais
        "name": "Anticoncepcionais hormonais",
        "clinical_relevance": """O histórico de uso de anticoncepcionais hormonais representa intervenção farmacológica significativa no eixo hipotálamo-hipófise-gonadal feminino, com implicações metabólicas, cardiovasculares, tromboembólicas e nutricionais que persistem durante e eventualmente após o uso. Na medicina funcional integrativa, os anticoncepcionais são reconhecidos como método contraceptivo eficaz e tratamento de certas condições ginecológicas, mas a abordagem enfatiza compreensão dos efeitos sistêmicos, individualização da escolha e prevenção de complicações.

TIPOS DE ANTICONCEPCIONAIS HORMONAIS:

COMBINADOS (estrogênio + progestagênio):
- Pílulas orais combinadas (POC): Etinilestradiol (EE) 20-35mcg + progestagênio (levonorgestrel, desogestrel, gestodeno, drospirenona, acetato de ciproterona)
- Adesivo transdérmico (Evra): libera EE + norelgestromina
- Anel vaginal (NuvaRing): libera EE + etonogestrel
- Mecanismo: supressão do eixo HHG → bloqueio da ovulação, espessamento do muco cervical, atrofia endometrial

PROGESTAGÊNIO ISOLADO (mais seguros cardiovascularmente):
- Minipílula (desogestrel 75mcg, noretisterona)
- Injetável trimestral (acetato de medroxiprogesterona - Depo-Provera)
- Implante subdérmico (etonogestrel - Implanon, 3 anos)
- DIU hormonal (levonorgestrel - Mirena, Kyleena, 5 anos)

BENEFÍCIOS ALÉM DA CONTRACEPÇÃO:
- Regulação de ciclos menstruais irregulares
- Redução de dismenorreia (cólicas), menorragia (fluxo intenso)
- Tratamento de síndrome dos ovários policísticos (SOP) - drospirenona, acetato de ciproterona
- Redução de acne e hirsutismo (efeito antiandrogênico)
- Tratamento de endometriose
- Redução de risco de câncer de ovário e endométrio (uso >5 anos)
- Tratamento de síndrome pré-menstrual (SPM)

EFEITOS ADVERSOS E RISCOS:

TROMBOEMBOLISMO VENOSO (TEV): Risco aumentado 2-6x conforme tipo. Maior risco no 1º ano de uso. POC de 3ª/4ª geração (desogestrel, gestodeno, drospirenona) têm risco ligeiramente maior que 2ª geração (levonorgestrel). Fatores de risco adicionais: trombofilia hereditária (Fator V Leiden - contraindicação relativa), obesidade, tabagismo, imobilização, cirurgia, câncer. CONTRAINDICAÇÃO ABSOLUTA: história pessoal de TVP/TEP, mutação homozigótica Fator V Leiden ou protrombina, trombofilias graves.

CARDIOVASCULARES: Pequeno aumento de risco de IAM e AVC isquêmico (principalmente se tabagismo, hipertensão, enxaqueca com aura). CONTRAINDICADO: >35 anos + tabagismo, hipertensão não controlada, diabetes com complicações vasculares, enxaqueca com aura.

METABÓLICOS: Resistência à insulina (especialmente progestagênios androgênicos), alterações lipídicas discretas (redução HDL, aumento triglicerídeos), retenção hídrica (estrogênio), ganho de peso (progestagênios androgênicos, Depo-Provera).

HEPÁTICOS: Aumento de risco de adenoma hepatocelular (raro, uso prolongado), colestase intra-hepática, exacerbação de porfiria. CONTRAINDICADO: doença hepática ativa, tumores hepáticos.

PSIQUIÁTRICOS: Alterações de humor (depressão, ansiedade) em subgrupo de mulheres, redução de libido (supressão de testosterona livre).

DEFICIÊNCIAS NUTRICIONAIS INDUZIDAS: Anticoncepcionais orais combinados DEPLETEM: Vitaminas B2, B6, B12, ácido fólico (crucial se gravidez pós-descontinuação), vitamina C, vitamina E, magnésio, zinco, selênio, triptofano. Consequências: fadiga, alterações de humor (déficit B6 → redução serotonina), hiperhomocisteinemia (déficit B6/B12/folato → risco cardiovascular), comprometimento imunológico.

CÂNCER DE MAMA: Pequeno aumento de risco (RR 1,2-1,3) durante uso e até 10 anos após descontinuação, retornando ao basal após. Deve ser discutido no consentimento informado. CONTRAINDICADO: câncer de mama atual ou recente.

OUTRAS: Náuseas (primeiros meses), cefaleia, mastalgia, sangramento de escape (primeiros 3 meses), ausência de menstruação (amenorreia), cloasma (manchas escuras na face - fotossensibilidade), piora de varizes.

A medicina funcional investiga causas subjacentes de irregularidades menstruais (SOP, hipotireoidismo, hiperprolactinemia, estresse crônico, distúrbios alimentares) e oferece alternativas quando apropriado: regulação do ciclo via nutrição, suplementação, manejo de estresse, melhora da sensibilidade insulínica (SOP), acupuntura. Também enfatiza suporte nutricional proativo durante uso de anticoncepcionais.""",

        "patient_explanation": """Anticoncepcionais hormonais são medicamentos que contêm hormônios femininos sintéticos (estrogênio e/ou progesterona) usados principalmente para prevenir gravidez, mas também para tratar várias condições como cólicas intensas, fluxo menstrual excessivo, acne, síndrome dos ovários policísticos e endometriose.

TIPOS:
- Pílula combinada: toma 1 por dia, contém estrogênio + progesterona
- Minipílula: só progesterona, mais segura para quem tem risco cardiovascular
- Injeção trimestral: aplicação a cada 3 meses
- Implante: inserido no braço, dura 3 anos
- DIU hormonal: colocado no útero, dura 5 anos
- Adesivo e anel vaginal: outras formas de entrega

COMO FUNCIONAM: Bloqueiam a ovulação (liberação do óvulo), tornam o muco cervical mais espesso (dificulta passagem do esperma) e deixam o útero menos receptivo.

BENEFÍCIOS além da contracepção:
- Redução de cólicas e fluxo menstrual
- Melhora da acne e pelos em excesso
- Ciclos mais regulares e previsíveis
- Proteção contra câncer de ovário e útero

EFEITOS COLATERAIS COMUNS (geralmente melhoram após 3 meses):
- Náusea, sensibilidade nas mamas, dor de cabeça
- Sangramento irregular nos primeiros meses
- Alterações de humor em algumas mulheres
- Retenção de líquidos
- Redução da libido em algumas mulheres

RISCOS IMPORTANTES:
- Pequeno aumento de risco de trombose (coágulos) - maior no 1º ano
- Contraindicado se: histórico de trombose, fumante >35 anos, enxaqueca com aura, pressão alta não controlada

DEFICIÊNCIAS NUTRICIONAIS: Anticoncepcionais reduzem vitaminas B6, B12, ácido fólico, magnésio e zinco. É recomendado suplementar para prevenir fadiga, alterações de humor e outros problemas.

É FUNDAMENTAL:
- Tomar no mesmo horário todos os dias (pílulas)
- Não fumar (aumenta muito o risco de trombose e AVC)
- Informar seu médico se tiver dor intensa na perna, falta de ar súbita, dor no peito, dor de cabeça intensa, visão turva
- Fazer consultas de acompanhamento regularmente""",

        "conduct": """1. Avaliação pré-prescrição:
   - Anamnese detalhada (fatores de risco cardiovascular, tromboembólico)
   - História familiar de trombose, câncer de mama
   - Pressão arterial
   - IMC
   - Tabagismo (CONTRAINDICAÇÃO se >35 anos + tabagismo)
   - Enxaqueca com aura (CONTRAINDICAÇÃO para combinados)

   CONTRAINDICAÇÕES ABSOLUTAS (OMS categoria 4):
   - Trombose venosa ou arterial atual/prévia
   - Trombofilias conhecidas (Fator V Leiden homozigoto)
   - Tabagismo >15 cigarros/dia + >35 anos
   - Hipertensão ≥160/100mmHg ou com complicações
   - Diabetes com complicações vasculares
   - Câncer de mama atual/recente
   - Doença hepática ativa, tumores hepáticos
   - Enxaqueca com aura (contraindicação para combinados)
   - <6 semanas pós-parto se amamentando

2. Escolha individualizada:

   BAIXO RISCO (jovem, sem fatores de risco):
   - POC com levonorgestrel (menor risco TEV)
   - Drospirenona se SOP, acne, retenção hídrica

   ALTO RISCO CARDIOVASCULAR/TROMBÓTICO:
   - Progestagênio isolado (minipílula, implante, DIU hormonal)
   - Evitar estrogênio

   HISTÓRICO DE DEPRESSÃO:
   - Monitorar cuidadosamente humor
   - Considerar DIU cobre (não hormonal)

3. Monitoramento durante uso:
   - Pressão arterial: 3 meses após início, depois anual
   - Triagem de câncer de mama conforme guidelines
   - Papanicolau conforme rotina
   - Reavaliação anual de contraindicações

4. Suplementação para prevenir deficiências (ESSENCIAL):
   - Complexo B de alta potência ou:
     - Vitamina B6 (25-50mg/dia) - crucial para humor
     - Vitamina B12 (500-1000mcg/dia)
     - Ácido fólico (400-800mcg/dia) - CRÍTICO se possível gravidez futura
   - Magnésio (300-400mg/dia)
   - Zinco (15-30mg/dia)
   - Vitamina C (500mg/dia)
   - Vitamina E (200-400UI/dia)
   - Selênio (100-200mcg/dia)

5. Otimização metabólica:
   - Controle de peso (exercícios, alimentação balanceada)
   - Melhora de sensibilidade insulínica (se SOP ou resistência)
   - Suporte hepático (detoxificação de estrogênios)

6. Redução de risco tromboembólico:
   - Cessação de tabagismo (CRÍTICO)
   - Manter peso saudável
   - Atividade física regular
   - Hidratação adequada
   - Suspender 4 semanas antes de cirurgias eletivas de grande porte
   - Profilaxia com HBPM em voos longos se alto risco

7. Suporte hepático (metabolização de estrogênios):
   - Vegetais crucíferos (brócolis, couve, repolho) - induzem fase II detox
   - DIM (di-indolilmetano) 100-200mg/dia ou I3C (indol-3-carbinol)
   - Cálcio-D-glucarato (500mg/dia) - facilita excreção de estrogênios
   - Cardo mariano (silimarina)

8. Manejo de efeitos colaterais:

   NÁUSEA: tomar à noite com alimento, vitamina B6

   SANGRAMENTO DE ESCAPE: aguardar 3 meses, se persistir trocar formulação

   RETENÇÃO HÍDRICA: drospirenona (efeito antimineralocorticoide), magnésio, reduzir sódio

   REDUÇÃO DE LIBIDO: considerar trocar para progestagênio menos androgênico ou DIU

   ALTERAÇÕES DE HUMOR: B6, magnésio, ômega-3; se persistir considerar suspensão

9. Avaliação de trombofilia (considerar se):
   - História familiar forte de trombose
   - Trombose em familiar jovem (<50 anos)
   - Pode incluir: Fator V Leiden, mutação protrombina, proteína C/S, antitrombina III

10. Transição pós-anticoncepcionais:
    - Regulação natural do ciclo pode levar 3-12 meses
    - Suporte nutricional intensivo (B6, B12, folato, zinco, magnésio)
    - Otimização da fase lútea (progesterona natural se necessário)
    - Tratamento de SOP subjacente se presente

11. Planejamento de gravidez futura:
    - Descontinuar 3-6 meses antes de tentar engravidar (ideal)
    - Suplementar ácido fólico 400-800mcg/dia ANTES de engravidar
    - Otimizar nutrientes (B12, ferro, vitamina D, ômega-3)

12. Educação da paciente:
    - Sinais de TEV (dor na perna, falta de ar, dor no peito) - EMERGÊNCIA
    - Sinais de AVC (dor de cabeça súbita intensa, visão turva, fraqueza unilateral)
    - Importância da adesão (tomar no mesmo horário)
    - Interações medicamentosas (antibióticos, anticonvulsivantes reduzem eficácia)
    - Backup contraceptivo se vômito/diarreia ou antibióticos"""
    }
}

# Continuar com os outros 11 medicamentos...
# (continuação no próximo arquivo devido ao tamanho)

class MedicationEnricher:
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
            print("✓ Login realizado com sucesso!\n")
            return True
        except Exception as e:
            print(f"✗ Erro no login: {e}")
            return False

    def update_score_item(self, item_id: str, content: Dict) -> bool:
        """Atualiza um score item com os textos clínicos"""
        try:
            payload = {
                "clinicalRelevance": content["clinical_relevance"],
                "patientExplanation": content["patient_explanation"],
                "conduct": content["conduct"]
            }

            response = self.session.put(
                f"{API_BASE_URL}/score-items/{item_id}",
                json=payload
            )
            response.raise_for_status()
            return True
        except Exception as e:
            print(f"  ✗ Erro ao atualizar item: {e}")
            return False

    def process_all(self):
        """Processa todos os items disponíveis"""
        results = {"success": [], "failed": [], "total": 0}

        print("="*80)
        print("ENRIQUECIMENTO DE MEDICAMENTOS - PARTE 2")
        print("="*80 + "\n")

        for item_id, content in MEDICATION_CONTENT.items():
            results["total"] += 1
            print(f"[{results['total']}/4] Processando: {content['name']}")
            print(f"    ID: {item_id}")

            success = self.update_score_item(item_id, content)

            if success:
                print(f"    ✓ Sucesso!")
                results["success"].append(content['name'])
            else:
                print(f"    ✗ Falhou")
                results["failed"].append(content['name'])

            print()
            time.sleep(0.5)

        # Sumário
        print("="*80)
        print("SUMÁRIO")
        print("="*80)
        print(f"Total: {results['total']}")
        print(f"Sucesso: {len(results['success'])}")
        print(f"Falhas: {len(results['failed'])}")
        print("="*80 + "\n")

        return results

def main():
    enricher = MedicationEnricher()

    if not enricher.login():
        return

    results = enricher.process_all()

    # Salvar resultados
    with open('/home/user/plenya/medications_part2_batch1_results.json', 'w') as f:
        json.dump(results, f, indent=2, ensure_ascii=False)

if __name__ == "__main__":
    main()
