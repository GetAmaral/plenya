#!/usr/bin/env python3
"""
Script para processar Score Items de exames laboratoriais
Enriquece itens com conteúdo clínico baseado em Medicina Funcional Integrativa
"""

import requests
import json
import sys
from typing import Dict, List, Optional

# Configurações da API
API_BASE_URL = "http://localhost:3001/api/v1"
LOGIN_EMAIL = "import@plenya.com"
LOGIN_PASSWORD = "Import@123456"

# Items a processar (ID, nome)
TARGET_ITEMS = [
    # Hemoglobina
    ("a99ff9aa-b39c-46aa-a042-39537a83eb81", "Hemoglobina - Homens"),
    ("bdd1777f-2da3-4c20-879e-21446040170e", "Hemoglobina - Mulheres"),
    ("660be9ac-31cd-4255-a15a-f049c099b68b", "Hemoglobina glicada"),

    # Glicose
    ("6d272987-00de-4f8d-a3d8-abd729dc24f7", "GLICOSE 0 MIN (JEJUM)"),
    ("cf39e4a8-e1ec-4e94-a2a1-596a2a3cdf44", "GLICOSE 120 MIN"),

    # Colesterol
    ("498a8637-8bf5-45e0-891b-a0c51a47ccc1", "Colesterol Total"),
    ("55787b95-d165-48d6-8eb4-e496bd62d509", "HDL Colesterol"),
    ("d55af9e8-4358-4cfa-a75c-05dfcd1bfb4d", "LDL Colesterol"),

    # Triglicérides
    ("dfa58a95-7d7a-491f-b7c5-d2fe8c694daa", "Triglicerídeos"),

    # Creatinina
    ("8234724e-f5d6-4862-9fde-84e78e136f88", "Creatinina"),
]


class ScoreItemProcessor:
    def __init__(self):
        self.token = None
        self.processed = []

    def authenticate(self) -> bool:
        """Autentica na API e obtém token"""
        try:
            response = requests.post(
                f"{API_BASE_URL}/auth/login",
                json={"email": LOGIN_EMAIL, "password": LOGIN_PASSWORD},
                timeout=10
            )

            if response.status_code == 200:
                data = response.json()
                self.token = data.get("accessToken")
                print(f"✓ Autenticado com sucesso")
                return True
            else:
                print(f"✗ Falha na autenticação: {response.status_code}")
                return False
        except Exception as e:
            print(f"✗ Erro na autenticação: {e}")
            return False

    def update_item(self, item_id: str, name: str, texts: Dict[str, str]) -> bool:
        """Atualiza um Score Item com conteúdo clínico"""
        try:
            headers = {
                "Authorization": f"Bearer {self.token}",
                "Content-Type": "application/json"
            }

            response = requests.put(
                f"{API_BASE_URL}/score-items/{item_id}",
                json=texts,
                headers=headers,
                timeout=30
            )

            if response.status_code == 200:
                print(f"✓ {name}: Atualizado com sucesso")
                self.processed.append({
                    "id": item_id,
                    "name": name,
                    "success": True,
                    "texts_generated": True
                })
                return True
            else:
                print(f"✗ {name}: Falha {response.status_code} - {response.text}")
                self.processed.append({
                    "id": item_id,
                    "name": name,
                    "success": False,
                    "error": response.text
                })
                return False

        except Exception as e:
            print(f"✗ {name}: Erro {e}")
            self.processed.append({
                "id": item_id,
                "name": name,
                "success": False,
                "error": str(e)
            })
            return False

    def generate_report(self) -> Dict:
        """Gera relatório final"""
        success_count = sum(1 for p in self.processed if p.get("success"))
        failed_count = len(self.processed) - success_count

        return {
            "processed": self.processed,
            "summary": {
                "total": len(self.processed),
                "success": success_count,
                "failed": failed_count
            }
        }


def generate_hemoglobina_texts(name: str) -> Dict[str, str]:
    """Gera textos clínicos para Hemoglobina"""

    clinical_relevance = """A hemoglobina é uma proteína presente nos eritrócitos (glóbulos vermelhos) responsável pelo transporte de oxigênio dos pulmões para todos os tecidos do corpo e pelo retorno de parte do dióxido de carbono aos pulmões para eliminação. Cada molécula de hemoglobina contém ferro em seu núcleo, elemento essencial para a ligação do oxigênio.

Na medicina funcional integrativa, avaliamos não apenas se os valores estão dentro da faixa laboratorial de referência, mas se estão em níveis funcionais ótimos. Valores de hemoglobina no limite inferior da normalidade, embora não configurem anemia franca, podem comprometer a oxigenação tecidual ideal, levando a fadiga, redução da capacidade física, prejuízo cognitivo e recuperação inadequada do exercício.

A produção adequada de hemoglobina depende não apenas de ferro, mas de um ecossistema nutricional completo: vitamina B12, ácido fólico, vitamina B6, cobre, vitamina A e proteínas de qualidade. A medicina funcional reconhece que a deficiência de ferro pode ocorrer mesmo com ferritina aparentemente normal, quando avaliamos marcadores mais sensíveis como saturação de transferrina, capacidade total de ligação do ferro (TIBC) e receptor solúvel de transferrina.

Valores elevados de hemoglobina podem sinalizar desidratação crônica, tabagismo, doenças pulmonares obstrutivas, apneia do sono grave ou policitemia. Na abordagem funcional, investigamos o contexto: atletas de altitude, exposição a hipóxia, uso de testosterona ou outros fatores que estimulam a eritropoiese.

A hemoglobina também se relaciona com outros marcadores: volume corpuscular médio (VCM) ajuda a classificar o tipo de anemia; relação com ferritina, ferro sérico e saturação de transferrina define o status de ferro corporal; e a correlação com contagem de reticulócitos indica a capacidade de resposta medular."""

    patient_explanation = """A hemoglobina é a proteína que carrega oxigênio no sangue. Imagine ela como um trem que pega oxigênio nos pulmões e leva para todo o corpo - músculos, cérebro, coração e todos os outros órgãos.

Quando a hemoglobina está baixa, seu corpo não recebe oxigênio suficiente. Você pode sentir cansaço constante, falta de ar ao subir escadas, dificuldade de concentração, pele pálida e até queda de cabelo. É como se o corpo estivesse sempre "sem gasolina".

Quando está alta demais, o sangue fica "grosso", o que pode ser causado por desidratação, cigarro ou problemas respiratórios. Em algumas pessoas, pode aumentar o risco de coágulos.

O exame é importante porque detecta problemas antes de você sentir sintomas graves. Níveis ideais garantem energia, disposição mental e boa recuperação após exercícios."""

    conduct = """**Avaliação Funcional:**

Valores ótimos funcionais diferem das referências laboratoriais. Para homens, busca-se hemoglobina entre 14,5-16,5 g/dL; para mulheres, 13,5-15,5 g/dL. Valores no terço inferior da faixa merecem investigação mesmo sem anemia.

**Quando Hemoglobina Baixa (ou no limite inferior):**

Solicitar painel completo: hemograma com índices (VCM, HCM, RDW), ferritina, ferro sérico, saturação de transferrina, TIBC, vitamina B12, ácido fólico, homocisteína e proteína C-reativa. Em casos refratários, considerar receptor solúvel de transferrina e zinco protoporfirina.

Investigar causas: sangramento oculto (gastrointestinal, menstrual intenso), má absorção (doença celíaca, gastrite atrófica, SIBO), inflamação crônica (anemia de doença crônica), deficiências nutricionais ou doenças hematológicas.

**Abordagem Terapêutica:**

1. **Nutrição:** Priorizar fontes de ferro heme (carnes vermelhas, fígado, aves), combinadas com vitamina C para melhor absorção. Incluir folhas verde-escuras, leguminosas e oleaginosas. Evitar fitatos e taninos (chá, café) próximos às refeições principais.

2. **Suplementação:** Ferro bisglicinato 25-50 mg/dia (melhor tolerância gastrointestinal), vitamina B12 sublingual 1.000-2.000 mcg, ácido fólico metilado 400-800 mcg, vitamina C 500 mg junto ao ferro. Ajustar doses conforme resposta.

3. **Correção de disbiose e permeabilidade intestinal:** Fundamentais para absorção adequada.

4. **Monitoramento:** Reavaliar hemograma em 8-12 semanas; ferritina funcional ideal acima de 50-70 ng/mL em mulheres e 70-100 ng/mL em homens.

**Quando Hemoglobina Alta:**

Investigar desidratação, tabagismo, apneia do sono (polissonografia), doenças pulmonares (espirometria, gasometria). Em atletas ou usuários de testosterona, monitorar viscosidade sanguínea e risco trombótico. Considerar flebotomia terapêutica se hematócrito >54% em homens ou >48% em mulheres, com sintomas."""

    return {
        "clinicalRelevance": clinical_relevance,
        "patientExplanation": patient_explanation,
        "conduct": conduct
    }


def generate_hemoglobina_glicada_texts() -> Dict[str, str]:
    """Gera textos clínicos para Hemoglobina Glicada (HbA1c)"""

    clinical_relevance = """A hemoglobina glicada (HbA1c) reflete a média glicêmica dos últimos 2-3 meses, pois mede a porcentagem de hemoglobina que sofreu glicação irreversível pela exposição contínua à glicose. É o marcador-padrão para diagnóstico e monitoramento de diabetes, mas na medicina funcional integrativa sua interpretação vai além.

Valores de HbA1c entre 5,7-6,4% definem pré-diabetes; ≥6,5% confirmam diabetes. No entanto, valores "normais" entre 5,3-5,6% já sinalizam resistência insulínica incipiente e risco metabólico aumentado. A medicina funcional busca HbA1c funcional ideal abaixo de 5,3%, refletindo controle glicêmico ótimo e menor inflamação sistêmica.

A glicação é um processo de envelhecimento acelerado: proteínas glicadas (AGEs - advanced glycation end products) geram estresse oxidativo, inflamação e disfunção endotelial, acelerando aterosclerose, doença renal, retinopatia e neuropatia. Portanto, HbA1c não é apenas um número diagnóstico, mas um marcador de "velocidade de envelhecimento" tecidual.

Limitações do teste: anemias, hemoglobinopatias e alta renovação eritrocitária (hemólise, sangramentos) podem falsear resultados. Nesses casos, frutosamina ou glicose média estimada (AGE) via monitoramento contínuo são alternativas.

A HbA1c deve ser interpretada em conjunto com: glicemia de jejum, curva glicêmica/insulinêmica pós-prandial (TOTG), insulina de jejum, HOMA-IR, peptídeo C e marcadores inflamatórios. A combinação permite identificar resistência insulínica precoce, hiperinsulinemia compensatória e disfunção de células beta pancreáticas antes do diabetes franco."""

    patient_explanation = """A hemoglobina glicada mostra a "média" do seu açúcar no sangue nos últimos 2 a 3 meses. É como se fosse um "boletim escolar" do controle da glicose: não adianta fazer o exame de jejum e "estudar" na véspera - a hemoglobina glicada revela a verdade.

Quando você come açúcar ou carboidratos, a glicose gruda na hemoglobina do sangue e fica lá por meses. Quanto mais açúcar no sangue, mais hemoglobina "grudada" (glicada) você terá.

Valores acima de 5,7% indicam que o corpo está com dificuldade de controlar o açúcar - é o pré-diabetes. Acima de 6,5%, já é diabetes. Mas mesmo valores entre 5,3% e 5,6% podem significar que seu corpo está "cansando" de processar tanto açúcar, mesmo que ainda não seja doença.

Quando a hemoglobina fica glicada demais por muito tempo, ela "estraga" as proteínas do corpo, danificando nervos, rins, olhos e vasos sanguíneos. Por isso é importante manter esse valor baixo, idealmente abaixo de 5,3%."""

    conduct = """**Avaliação Funcional Integrativa:**

Meta funcional ideal: HbA1c <5,3%. Valores 5,3-5,6% merecem intervenção preventiva; 5,7-6,4% exigem protocolo intensivo de reversão de pré-diabetes; ≥6,5% requerem abordagem multifatorial imediata.

**Investigação Complementar:**

Solicitar painel metabólico completo: glicemia de jejum, TOTG com curva de glicose e insulina (0, 30, 60, 90, 120 min), insulina de jejum, HOMA-IR, peptídeo C, perfil lipídico avançado (triglicerídeos/HDL, ApoB, partículas de LDL), marcadores inflamatórios (PCR-us, homocisteína), função hepática e renal.

Avaliar composição corporal (bioimpedância ou DEXA) para quantificar gordura visceral, principal driver da resistência insulínica.

**Estratégias de Manejo:**

**1. Nutrição (pilar central):**
- Dieta low carb ou cetogênica (20-50g carboidratos/dia) para reversão rápida de resistência insulínica
- Jejum intermitente 16/8 ou protocolo 5:2 para ativar mitofagia e reduzir carga glicêmica
- Priorizar proteínas de qualidade, gorduras saudáveis (ômega-3, monoinsaturadas), vegetais não-amiláceos
- Eliminar ultraprocessados, açúcares adicionados, farináceos refinados

**2. Exercício:**
- Musculação 3-4x/semana: contração muscular capta glicose independente de insulina
- HIIT 2-3x/semana: melhora sensibilidade insulínica aguda e crônica
- Caminhar pós-refeição (10-15 min): reduz pico glicêmico pós-prandial

**3. Correção de Disbiose e Permeabilidade Intestinal:**
- Endotoxemia metabólica (LPS bacteriano) perpetua inflamação e resistência insulínica
- Prebióticos, probióticos específicos, glutamina, zinco-carnosina

**4. Suplementação:**
- Berberina 500 mg 2-3x/dia: potência comparável à metformina
- Cromo picolinato 200-400 mcg/dia
- Ácido alfa-lipoico 600 mg/dia
- Magnésio bisglicinato 400-600 mg/dia
- Canela de Ceilão 1-3g/dia
- Ômega-3 (EPA+DHA) 2-3g/dia
- Coenzima Q10 100-200 mg/dia (suporte mitocondrial)

**5. Sono e Manejo de Estresse:**
- Sono <6h ou fragmentado piora resistência insulínica agudamente
- Técnicas de redução de estresse: meditação, respiração diafragmática, yoga
- Cortisol cronicamente elevado antagoniza insulina

**6. Monitoramento:**
- HbA1c a cada 3 meses até normalização, depois semestral
- Glicemia capilar pós-prandial 1-2h (meta <120 mg/dL)
- Reavaliar marcadores em 8-12 semanas

**Em Diabéticos:**
Meta inicial <7%, depois progressivamente <6,5% se sem hipoglicemias. Combinar abordagem funcional com farmacoterapia (metformina, iSGLT2, iGLP1) conforme necessidade. Rastreio anual de complicações: fundo de olho, microalbuminúria, monofilamento (neuropatia)."""

    return {
        "clinicalRelevance": clinical_relevance,
        "patientExplanation": patient_explanation,
        "conduct": conduct
    }


def generate_glicose_jejum_texts() -> Dict[str, str]:
    """Gera textos clínicos para Glicose de Jejum"""

    clinical_relevance = """A glicemia de jejum representa a concentração de glicose sanguínea após 8-12 horas sem ingestão calórica, refletindo primariamente a produção hepática de glicose (gliconeogênese e glicogenólise) e a sensibilidade hepática à insulina.

Valores laboratoriais de referência consideram normal até 99 mg/dL; pré-diabetes 100-125 mg/dL; diabetes ≥126 mg/dL (confirmado em duas ocasiões). Contudo, a medicina funcional integrativa adota faixas mais restritivas: glicemia funcional ótima situa-se entre 70-85 mg/dL. Valores entre 86-99 mg/dL, embora "normais", já associam-se a risco aumentado de progressão para diabetes e eventos cardiovasculares.

A glicemia de jejum é marcador tardio de disfunção metabólica. Antes dela se alterar, ocorrem anos de hiperinsulinemia compensatória (o pâncreas produz mais insulina para manter a glicose controlada). Portanto, glicemia de jejum normal NÃO descarta resistência insulínica. É essencial dosar insulina de jejum, calcular HOMA-IR e avaliar glicemia/insulina pós-prandial via TOTG.

A glicemia é influenciada por múltiplos fatores: padrão alimentar (excesso de carboidratos refinados, frutose industrial), obesidade visceral, sedentarismo, sono inadequado, estresse crônico (cortisol elevado estimula gliconeogênese), disbiose intestinal (endotoxemia metabólica), deficiências nutricionais (magnésio, cromo, vitamina D) e medicamentos (corticoides, betabloqueadores).

Fisiopatologicamente, a elevação da glicemia de jejum indica falência da capacidade hepática de responder à insulina noturna e resistência insulínica hepática estabelecida. Essa condição perpetua dislipidemia aterogênica (aumento de triglicerídeos, redução de HDL, partículas de LDL pequenas e densas), inflamação sistêmica e estresse oxidativo mitocondrial."""

    patient_explanation = """A glicose de jejum mede o açúcar no seu sangue após você ficar sem comer por pelo menos 8 horas, geralmente durante a noite. É um dos exames mais comuns para avaliar como seu corpo está controlando o açúcar.

O ideal é que esse valor fique entre 70 e 85 mg/dL. Valores acima de 100 mg/dL já indicam pré-diabetes, e acima de 126 mg/dL significa diabetes. Mas mesmo valores entre 86 e 99, apesar de serem considerados "normais" no laboratório, já podem indicar que seu corpo está tendo dificuldade para controlar o açúcar.

Quando a glicose de jejum está alta, significa que seu fígado está produzindo açúcar demais durante a noite, mesmo sem você comer. Isso acontece porque as células do corpo estão "resistentes" à insulina - o hormônio que controla o açúcar.

Ter glicose alta é perigoso porque danifica nervos, rins, olhos e vasos sanguíneos ao longo do tempo. Mas a boa notícia é que mudanças no estilo de vida - alimentação, exercícios e sono - podem reverter esse quadro antes que vire diabetes."""

    conduct = """**Interpretação Funcional:**

- **Ótimo:** 70-85 mg/dL
- **Limítrofe (investigar resistência insulínica):** 86-99 mg/dL
- **Pré-diabetes:** 100-125 mg/dL
- **Diabetes:** ≥126 mg/dL (confirmar em segunda dosagem)

**Investigação Complementar Obrigatória:**

Nunca avaliar glicemia de jejum isoladamente. Solicitar:

- **Painel glicêmico completo:** Insulina de jejum, HOMA-IR, peptídeo C, hemoglobina glicada (HbA1c)
- **TOTG (curva glicêmica com insulina):** 0, 30, 60, 90, 120 min - detecta resistência insulínica precoce e picos pós-prandiais ocultos
- **Perfil lipídico avançado:** Triglicerídeos, HDL, relação TG/HDL (>3 sugere resistência insulínica), ApoB, partículas de LDL
- **Marcadores inflamatórios:** PCR-us, homocisteína
- **Micronutrientes:** Magnésio sérico/eritrocitário, vitamina D, zinco
- **Avaliação hepática:** TGO, TGP, GGT (esteatose hepática frequente em resistência insulínica)

**Abordagem Terapêutica:**

**Glicemia 86-99 mg/dL (normal-alta) ou pré-diabetes:**

**1. Intervenção Nutricional Intensiva:**
- Redução drástica de carboidratos refinados e açúcares: eliminar pães brancos, massas, doces, refrigerantes, sucos
- Dieta low carb (50-100g/dia) ou cetogênica (<50g/dia) conforme tolerância
- Jejum intermitente 16/8: concentra refeições em janela de 8h (ex: 12h-20h), reduz carga glicêmica diária
- Priorizar proteínas de qualidade (carnes, ovos, peixes), gorduras saudáveis (abacate, azeite, oleaginosas), vegetais fibrosos
- Evitar combinação de gordura + carboidrato na mesma refeição (sobrecarga mitocondrial)

**2. Exercício Físico (não negociável):**
- Musculação 3-4x/semana: aumenta massa muscular (principal captador de glicose)
- HIIT 2-3x/semana: melhora sensibilidade insulínica
- Caminhada pós-prandial 10-15 min: reduz pico glicêmico

**3. Correção de Disbiose:**
- Prebióticos (inulina, FOS), probióticos (Lactobacillus, Bifidobacterium), restaurar barreira intestinal

**4. Suplementação:**
- **Berberina** 500 mg 2-3x/dia (evidência similar à metformina)
- **Magnésio bisglicinato** 400-600 mg/dia (cofator de 300+ enzimas metabólicas)
- **Cromo picolinato** 200-400 mcg/dia
- **Ômega-3** EPA+DHA 2-3g/dia (anti-inflamatório)
- **Canela de Ceilão** 1-3g/dia
- **Ácido alfa-lipoico** 600 mg/dia

**5. Otimização de Sono e Estresse:**
- Sono 7-9h/noite, higiene do sono
- Redução de cortisol: meditação, adaptógenos (ashwagandha, rhodiola)

**Diabetes estabelecido (≥126 mg/dL):**

Abordagem funcional combinada com farmacoterapia. Metformina 1.500-2.000 mg/dia como primeira linha, associada a todas as estratégias acima. Considerar iSGLT2 ou iGLP1 conforme perfil. Meta glicêmica: <100 mg/dL jejum, <140 mg/dL pós-prandial.

**Monitoramento:**

- Glicemia de jejum mensal até estabilização
- HbA1c trimestral
- Reavaliação de insulina, HOMA-IR e perfil lipídico em 12 semanas
- Glicemia capilar domiciliar pós-prandial útil para ajuste fino"""

    return {
        "clinicalRelevance": clinical_relevance,
        "patientExplanation": patient_explanation,
        "conduct": conduct
    }


def main():
    """Função principal"""
    processor = ScoreItemProcessor()

    print("=== Processador de Score Items - Exames Laboratoriais ===\n")

    # Autenticar
    if not processor.authenticate():
        sys.exit(1)

    print(f"\n{'='*60}")
    print("Processando items...")
    print(f"{'='*60}\n")

    # Processar cada item
    for item_id, item_name in TARGET_ITEMS[:3]:  # Processar primeiro lote
        print(f"\n--- {item_name} ---")

        # Gerar textos baseado no tipo
        if "Hemoglobina glicada" in item_name:
            texts = generate_hemoglobina_glicada_texts()
        elif "Hemoglobina" in item_name:
            texts = generate_hemoglobina_texts(item_name)
        elif "GLICOSE" in item_name and "JEJUM" in item_name:
            texts = generate_glicose_jejum_texts()
        else:
            print(f"⊘ Pulando (sem gerador específico)")
            continue

        # Atualizar via API
        processor.update_item(item_id, item_name, texts)

    # Gerar relatório
    print(f"\n{'='*60}")
    print("RELATÓRIO FINAL")
    print(f"{'='*60}\n")

    report = processor.generate_report()
    print(json.dumps(report, indent=2, ensure_ascii=False))

    # Salvar relatório
    with open("/home/user/plenya/lab_items_report.json", "w", encoding="utf-8") as f:
        json.dump(report, f, indent=2, ensure_ascii=False)

    print(f"\n✓ Relatório salvo em: /home/user/plenya/lab_items_report.json")


if __name__ == "__main__":
    main()
