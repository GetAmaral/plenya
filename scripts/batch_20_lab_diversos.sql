-- ============================================================================
-- ENRIQUECIMENTO DE 20 ITEMS DE EXAMES LABORATORIAIS DIVERSOS
-- Data: 2026-01-27
-- Items: Albumina, Alfa-1 Globulina, Alfa-2 Globulina, Alfafetoproteína,
--        Alumínio, Amilase, Anti-LA (SSB), Anti-RO (SSA), Anti-TPO, Anti-Tireoglobulina
-- ============================================================================

-- 1. ALBUMINA
UPDATE score_items
SET
  clinical_explanation = 'A albumina é a principal proteína plasmática sintetizada exclusivamente pelos hepatócitos no fígado, representando aproximadamente 60% das proteínas totais do plasma sanguíneo. Com peso molecular de cerca de 66 kDa, a albumina desempenha múltiplas funções fisiológicas essenciais, incluindo a manutenção da pressão oncótica plasmática (responsável por 70-80% da pressão oncótica), o transporte de substâncias hidrofóbicas (hormônios tireoidianos, ácidos graxos livres, bilirrubina não conjugada, drogas), e ação antioxidante através de seus grupos tiol. A síntese hepática de albumina é influenciada por diversos fatores, incluindo estado nutricional, pressão oncótica plasmática, hormônios (insulina, cortisol, hormônio do crescimento), e processos inflamatórios. Sua meia-vida plasmática é de aproximadamente 20 dias, tornando-a um marcador útil para avaliar o estado nutricional proteico em médio prazo e a função de síntese hepática. A dosagem de albumina sérica é amplamente utilizada na prática clínica para avaliação do estado nutricional, função hepática, perdas proteicas renais ou gastrointestinais, e como componente de escores prognósticos em diversas condições clínicas.',

  low_explanation = 'A hipoalbuminemia (albumina sérica <3,5 g/dL) resulta de três mecanismos fisiopatológicos principais: síntese hepática diminuída, perdas aumentadas ou redistribuição. A redução da síntese ocorre em hepatopatias crônicas (cirrose, hepatite crônica), onde a massa de hepatócitos funcionantes está reduzida, e em estados de desnutrição proteico-calórica, onde há deficiência de substratos. Processos inflamatórios agudos e crônicos também suprimem a síntese de albumina através de citocinas inflamatórias (IL-6, TNF-α), que redirecionam a síntese hepática para proteínas de fase aguda. As perdas aumentadas ocorrem na síndrome nefrótica (proteinúria >3,5 g/dia), enteropatia perdedora de proteínas (doença inflamatória intestinal, linfangiectasia intestinal), e queimaduras extensas. A redistribuição da albumina do compartimento intravascular para o interstício ocorre em estados de aumento da permeabilidade capilar, como sepse e SDRA. A hipoalbuminemia tem importantes implicações clínicas, incluindo desenvolvimento de edema e ascite, alteração na farmacocinética de medicamentos, e aumento do risco de infecções.',

  high_explanation = 'A hiperalbuminemia verdadeira (albumina sérica >5,0 g/dL) é incomum e geralmente representa desidratação ou hemoconcentração, não refletindo aumento absoluto da massa de albumina corporal. Os mecanismos envolvem perda de água livre desproporcional à perda de proteínas, resultando em concentração relativa da albumina plasmática. Causas comuns incluem desidratação por redução da ingesta hídrica (particularmente em idosos), perdas gastrintestinais (vômitos, diarreia sem reposição adequada), perdas renais (diabetes insipidus), perdas insensíveis aumentadas (febre, hiperventilação, sudorese excessiva), e uso excessivo de diuréticos. Em alguns casos, a hemoconcentração pode ocorrer em situações de sequestro de líquidos no terceiro espaço, como em obstrução intestinal ou pancreatite aguda. A interpretação deve sempre considerar o contexto clínico e sinais de desidratação.',

  clinical_significance = 'A albumina sérica é um dos marcadores bioquímicos mais importantes na prática clínica, refletindo múltiplos aspectos da homeostase fisiológica. Valores normais variam entre 3,5-5,0 g/dL. Na avaliação nutricional, a albumina é considerada um marcador de desnutrição proteico-calórica crônica. Valores <3,0 g/dL indicam desnutrição moderada, e <2,5 g/dL desnutrição grave, correlacionando-se com aumento de morbimortalidade, complicações pós-operatórias, e resposta imunológica comprometida. Na avaliação de hepatopatias, a albumina é componente essencial de escores prognósticos como Child-Pugh e MELD. Em doenças renais, a hipoalbuminemia na síndrome nefrótica correlaciona-se com o grau de proteinúria. A albumina também é utilizada no cálculo do gradiente de albumina sérica-ascite (GASA), fundamental para investigação etiológica de ascite. Na medicina intensiva, a hipoalbuminemia é marcador de gravidade em sepse e preditor de mortalidade.',

  interpretation_guide = 'Na interpretação da dosagem de albumina sérica, considerar os valores de referência (3,5-5,0 g/dL), variações fisiológicas, e contexto clínico. Valores entre 3,0-3,5 g/dL sugerem hipoalbuminemia leve, entre 2,5-3,0 g/dL moderada, e <2,5 g/dL grave. Quando encontrada hipoalbuminemia, verificar as proteínas totais: se ambas reduzidas, considerar perdas proteicas ou desnutrição; se apenas a albumina reduzida com globulinas normais ou elevadas, pensar em hepatopatia ou inflamação. A investigação deve incluir: proteinúria de 24 horas, provas de função hepática, marcadores inflamatórios (PCR, VHS), eletroforese de proteínas, e avaliação nutricional. A albumina deve ser solicitada em avaliação pré-operatória, seguimento de hepatopatias crônicas, investigação de edema/ascite, e monitorização de pacientes em terapia nutricional. A decisão de repor albumina exógena deve considerar a etiologia, gravidade, e situações específicas.',

  recommendations = ARRAY[
    'Solicitar albumina sérica como parte da avaliação nutricional em pacientes com doenças crônicas, desnutrição suspeita, ou em preparo pré-operatório de cirurgias de grande porte',
    'Avaliar em conjunto com proteína C reativa (PCR) para diferenciar desnutrição (albumina baixa, PCR normal) de processo inflamatório (ambos alterados)',
    'Em hipoalbuminemia, investigar perdas renais com proteinúria de 24 horas e perdas gastrintestinais com calprotectina fecal e endoscopia se indicado',
    'Monitorar albumina sérica em pacientes cirróticos a cada 3-6 meses como parte da avaliação prognóstica (escore de Child-Pugh)',
    'Considerar albumina como marcador prognóstico em pacientes críticos, mas não como alvo isolado de terapia de reposição',
    'Em pacientes com síndrome nefrótica e albumina <2,0 g/dL, avaliar risco tromboembólico aumentado e considerar anticoagulação profilática',
    'Repetir dosagem após 2-4 semanas em pacientes em terapia nutricional para avaliar resposta (melhora lenta esperada)',
    'Na presença de ascite, calcular o gradiente de albumina sérica-ascite (GASA) para investigação etiológica'
  ],

  related_conditions = ARRAY[
    'Cirrose hepática',
    'Síndrome nefrótica',
    'Desnutrição proteico-calórica',
    'Doença inflamatória intestinal',
    'Linfangiectasia intestinal',
    'Sepse e síndrome de resposta inflamatória sistêmica',
    'Síndrome de má-absorção intestinal',
    'Queimaduras extensas',
    'Neoplasias avançadas',
    'Insuficiência cardíaca congestiva'
  ],

  patient_friendly_explanation = 'A albumina é a principal proteína presente no sangue, produzida pelo fígado. Ela ajuda a manter o líquido dentro dos vasos sanguíneos (evitando inchaço), transporta medicamentos e hormônios pelo corpo, e fornece nutrição aos tecidos. Este exame mede a quantidade de albumina no sangue e é útil para avaliar se o fígado está funcionando bem, se você está recebendo nutrição adequada, e se há perda de proteínas pelos rins ou intestino. Valores baixos podem indicar problemas no fígado, desnutrição, inflamação crônica, ou perda de proteínas pelos rins ou intestino. Quando a albumina está baixa, pode haver inchaço nas pernas e barriga. Valores muito altos são raros e geralmente indicam desidratação. O resultado deve ser interpretado junto com outros exames e sua condição clínica geral.',

  updated_at = CURRENT_TIMESTAMP
WHERE id = '6cb96be1-1095-4641-88cc-a403fb034c8a';

-- 2-3. ALFA-1 GLOBULINA (3 items com mesmo nome)
UPDATE score_items
SET
  clinical_explanation = 'A alfa-1 globulina é uma fração proteica sérica identificada por eletroforese de proteínas, compreendendo principalmente a alfa-1-antitripsina (AAT), que representa aproximadamente 90% desta fração. A alfa-1-antitripsina é uma glicoproteína de 52 kDa sintetizada principalmente pelos hepatócitos e macrófagos, pertencente à família das serpinas (serine protease inhibitors). Sua principal função fisiológica é inibir elastase neutrofílica e outras proteases séricas, protegendo tecidos da degradação por enzimas liberadas durante processos inflamatórios, particularmente no pulmão. A AAT constitui uma proteína de fase aguda positiva, com níveis elevando-se 3-4 vezes durante processos inflamatórios. A deficiência genética de alfa-1-antitripsina é uma condição autossômica codominante causada por mutações no gene SERPINA1. Os fenótipos mais comuns são PiMM (normal), PiMZ (deficiência leve) e PiZZ (deficiência grave). A deficiência predispõe a enfisema pulmonar precoce e doença hepática.',

  low_explanation = 'A redução da fração alfa-1 globulina sugere principalmente deficiência de alfa-1-antitripsina, podendo ser de origem genética ou adquirida. A deficiência genética de AAT resulta de mutações no gene SERPINA1, sendo a deficiência hereditária mais comum em caucasianos (1:2.500-5.000). O fenótipo PiZZ é o mais grave, com níveis de AAT <35% do normal, predispondo a enfisema pulmonar precoce (terceira-quarta década em fumantes) de predomínio em bases pulmonares, hepatopatia crônica (cirrose infantil, cirrose em adultos, hepatocarcinoma), e raramente paniculite necrosante. Heterozigotos PiMZ têm risco discretamente aumentado de doença pulmonar, especialmente se fumantes. A deficiência adquirida pode ocorrer em hepatopatias graves devido à redução da síntese hepática, síndrome nefrótica com perdas proteicas maciças, enteropatia perdedora de proteínas, ou desnutrição grave. A investigação deve incluir dosagem quantitativa de AAT, fenotipagem, provas de função pulmonar, e tomografia de tórax em pacientes sintomáticos.',

  high_explanation = 'O aumento da fração alfa-1 globulina geralmente reflete elevação de proteínas de fase aguda, principalmente alfa-1-antitripsina e alfa-1-glicoproteína ácida, em resposta a processos inflamatórios, infecciosos ou neoplásicos. Durante a resposta de fase aguda, citocinas pró-inflamatórias (IL-6, IL-1, TNF-α) estimulam os hepatócitos a aumentar a síntese de proteínas de fase aguda, incluindo a alfa-1-antitripsina, que pode elevar-se 3-4 vezes acima dos valores basais. Condições associadas incluem infecções bacterianas agudas ou crônicas, processos inflamatórios sistêmicos (artrite reumatoide, doença inflamatória intestinal, vasculites), neoplasias (especialmente linfomas e alguns carcinomas), trauma extenso, queimaduras, cirurgias de grande porte, e infarto agudo do miocárdio. Em gestação, há aumento fisiológico da fração alfa-1. A interpretação deve considerar o padrão completo da eletroforese de proteínas e confirmação com marcadores inflamatórios específicos.',

  clinical_significance = 'A avaliação da fração alfa-1 globulina tem importância clínica principalmente na triagem de deficiência de alfa-1-antitripsina e na avaliação de processos inflamatórios. Valores normais variam entre 0,1-0,4 g/dL ou 2-5% das proteínas totais. A principal aplicação clínica é na suspeita de deficiência de AAT, que deve ser considerada em: pacientes com enfisema precoce (<45 anos), enfisema de predomínio basal, DPOC em não fumantes ou com tabagismo leve, doença hepática crônica de causa indeterminada, hepatite neonatal, história familiar de enfisema ou hepatopatia, paniculite necrosante, ou vasculite c-ANCA positiva. A eletroforese mostrando redução da fração alfa-1 (<0,1 g/dL ou <2%) é sugestiva mas não diagnóstica, requerendo confirmação com dosagem quantitativa de AAT (valores normais 90-200 mg/dL; deficiência grave <57 mg/dL) e fenotipagem. O rastreamento de familiares de pacientes com deficiência é recomendado.',

  interpretation_guide = 'Na interpretação da fração alfa-1 globulina, considerar o contexto clínico e correlacionar com outras frações proteicas. Valores normais são 0,1-0,4 g/dL ou 2-5% das proteínas totais. Alfa-1 reduzida (<0,1 g/dL ou <2%) sugere deficiência de alfa-1-antitripsina e requer dosagem quantitativa específica de AAT. Em pacientes com AAT <140 mg/dL, considerar fenotipagem (PiMM normal, PiMZ heterozigoto, PiZZ homozigoto). Alfa-1 aumentada (>0,4 g/dL ou >5%) geralmente reflete processos inflamatórios. A eletroforese de proteínas deve ser solicitada em: investigação de enfisema precoce ou em não fumantes, hepatopatia crônica de etiologia indeterminada, avaliação de processos inflamatórios crônicos, investigação de gamopatias monoclonais. Quando deficiência de AAT é confirmada (AAT <57 mg/dL, fenótipo PiZZ), avaliar função pulmonar com espirometria e tomografia de tórax, enzimas hepáticas, e considerar terapia de reposição. A cessação do tabagismo é IMPERATIVA em pacientes com deficiência de AAT.',

  recommendations = ARRAY[
    'Solicitar eletroforese de proteínas em pacientes com enfisema precoce (<45 anos), enfisema em não fumantes, ou DPOC com características atípicas',
    'Em alfa-1 globulina reduzida (<0,1 g/dL), solicitar dosagem quantitativa de alfa-1-antitripsina e fenotipagem para confirmar deficiência',
    'Rastrear familiares de primeiro grau de pacientes com deficiência confirmada de alfa-1-antitripsina (padrão PiZZ ou PiSZ)',
    'Em deficiência confirmada de AAT, realizar espirometria anualmente e tomografia de tórax a cada 2-3 anos para monitorar progressão',
    'Orientar rigorosamente pacientes com deficiência de AAT sobre a NECESSIDADE ABSOLUTA de cessação do tabagismo',
    'Considerar terapia de reposição de alfa-1-antitripsina em pacientes com fenótipo PiZZ, AAT <57 mg/dL, e enfisema progressivo',
    'Em alfa-1 globulina elevada, investigar processo inflamatório subjacente e correlacionar com proteína C reativa',
    'Vacinar pacientes com deficiência de AAT contra pneumococo (PCV13 e PPSV23) e influenza anualmente'
  ],

  related_conditions = ARRAY[
    'Deficiência de alfa-1-antitripsina',
    'Enfisema pulmonar precoce',
    'Doença pulmonar obstrutiva crônica (DPOC)',
    'Cirrose hepática por deficiência de AAT',
    'Hepatite neonatal colestática',
    'Paniculite necrosante',
    'Processos inflamatórios agudos e crônicos',
    'Infecções bacterianas',
    'Neoplasias (linfomas, carcinomas)',
    'Síndrome nefrótica'
  ],

  patient_friendly_explanation = 'A alfa-1 globulina é um grupo de proteínas no sangue, sendo a principal delas a alfa-1-antitripsina (AAT), produzida pelo fígado. A função mais importante da AAT é proteger os pulmões da destruição por enzimas liberadas durante inflamações. Este exame mede a quantidade dessas proteínas e é especialmente importante para detectar a deficiência de alfa-1-antitripsina, uma condição hereditária que pode levar ao desenvolvimento precoce de enfisema pulmonar (principalmente em fumantes) e, em alguns casos, doença no fígado. Valores baixos podem indicar deficiência genética de AAT, que é importante diagnosticar porque requer cuidados especiais, especialmente evitar o tabagismo de forma absoluta. Valores altos geralmente indicam inflamação ou infecção no corpo. Se o resultado sugerir deficiência de AAT, exames adicionais serão solicitados para confirmar o diagnóstico, e pode ser recomendado investigar outros membros da família.',

  updated_at = CURRENT_TIMESTAMP
WHERE id IN (
  '0ed3b126-3e60-4189-bc2c-e46b9606975a',
  '88081d50-7089-4f41-b463-c23347afedbc',
  'de7fa5ad-a023-49df-8063-8cfffa07de85'
);

-- Continua no próximo bloco...
