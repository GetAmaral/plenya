#!/usr/bin/env python3
"""
BATCH LAB PART 5 - 20 Items de Exames Laboratoriais (Versão Manual)
Gera conteúdo clínico estruturado sem API externa
"""

import requests
import json
import time
from typing import Dict, Any, List

# Config
API_BASE = "http://localhost:3001/api/v1"
LOGIN_EMAIL = "import@plenya.com"
LOGIN_PASSWORD = "Import@123456"

# IDs dos 20 items
ITEM_IDS = [
    "62d59b40-6dab-46c8-a6f5-53bafe3b7cae",  # CA-125
    "f36dd23c-3aa9-434f-bc86-18a3b5ee7020",  # CEA
    "eb19354a-d925-4301-bec5-2bf488501cea",  # CHCM (MCHC)
    "5c8cdefa-2c8d-4fbf-a5ca-f0ad586e1cda",  # CIMT Carótidas
    "9de603e9-3027-4fd4-b375-5d1b9298506d",  # Calprotectina fecal
    "b6b25826-63d5-4698-87c6-b23d1390e90f",  # IST
    "2a8611ab-0d4a-4a35-b3f8-569eb041ab9d",  # Chumbo
    "e4a51d97-6e52-4745-8c61-bae621bf0d3e",  # Cilindros Hialinos
    "0498349f-58d5-4c77-a25b-76c7c15c9496",  # Cilindros Patológicos
    "ae49e9b7-e1b2-4177-b8b1-d0e30f85e7fc",  # Cistatina C
    "45b91bea-accb-41a5-87aa-ac932eda0ce1",  # Cobre
    "4755b545-6065-4e41-8b7b-cb3fe0b81388",  # Coenzima Q10
    "995f731f-f891-4928-89e4-b53f27e1bb5c",  # Colesterol Total
    "aeee7b2d-0aa9-490b-bc26-e5b4ea5bc37c",  # Colesterol não-HDL
    "dd16df2d-deb3-46bc-bb4a-fdd7139863c9",  # Colonoscopia - Mayo Score
    "c1ea91b5-132e-419a-ae6e-5105d4df205c",  # Colonoscopia - Adenomas
    "6255db6b-9dfb-4941-a1a6-e79c8acbab17",  # Colonoscopia - SES-CD
    "fed861ee-d777-4d0a-b06d-ee1733c903e9",  # Colonoscopia geral
    "593303eb-371c-4b81-946c-4529d81c7e09",  # Relação Colesterol/HDL
    "8bf5780d-7c03-4a36-bd5b-2562ee525ab6",  # Cristais Patológicos
]

# Conteúdos clínicos estruturados
CLINICAL_CONTENT = {
    "CA-125": {
        "clinical_significance": """O CA-125 (Cancer Antigen 125) é uma glicoproteína de alto peso molecular expressa em células epiteliais derivadas do epitélio celômico, incluindo endométrio, pleura e peritônio. Elevações séricas ocorrem primariamente em carcinoma ovariano epitelial (85% dos casos avançados), sendo o biomarcador mais estabelecido para monitoramento desta neoplasia. Mecanicamente, células tumorais ovarianas secretam CA-125 que é liberado na circulação, refletindo carga tumoral e atividade metastática.

Valores elevados também ocorrem em condições benignas como endometriose (40-50% dos casos), miomas uterinos, doença inflamatória pélvica e insuficiência hepática com ascite, limitando especificidade diagnóstica. A interpretação requer contexto clínico rigoroso, correlacionando com exames de imagem (ultrassom transvaginal, tomografia) e sintomatologia. Em pacientes com câncer ovariano tratado, elevações progressivas precedem recidiva clínica em 3-6 meses, permitindo intervenção precoce.

Concentrações normais são inferiores a 35 U/mL em mulheres na pré-menopausa e 20 U/mL na pós-menopausa. Valores limítrofes requerem seguimento seriado, pois cinética de elevação é mais informativa que valor absoluto isolado.""",

        "longevity_context": """No contexto de longevidade feminina, CA-125 serve como marcador de saúde reprodutiva e inflamação peritoneal crônica. Elevações persistentes mesmo em faixas limítrofes (20-35 U/mL) associam-se a processos inflamatórios subclínicos que aceleram envelhecimento ovariano e aumentam risco cardiovascular por mecanismos de estresse oxidativo sistêmico.

Estudos de coorte demonstram que mulheres com CA-125 consistentemente baixo (<15 U/mL) apresentam menor incidência de síndrome metabólica e melhor perfil de senescência celular. O monitoramento preventivo permite identificação precoce de endometriose não-diagnosticada, condição que impacta qualidade de vida e fertilidade futura.""",

        "clinical_recommendations": """Solicitar CA-125 em mulheres com massa anexial identificada em ultrassom, especialmente na pós-menopausa ou com características suspeitas (componentes sólidos, septações, ascite). Em pacientes com diagnóstico de câncer ovariano, monitorar a cada 3 meses durante tratamento e anualmente após remissão.

Para endometriose, valores elevados correlacionam-se com doença avançada (estágios III-IV) e podem guiar decisão cirúrgica. Fatores que elevam CA-125 incluem menstruação (medir após 5º dia do ciclo), gravidez primeiro trimestre e tabagismo. Redução >50% após 3 ciclos de quimioterapia é fator prognóstico positivo.

Em rastreamento populacional, CA-125 isolado tem baixo valor preditivo positivo (<10%), não sendo recomendado fora de populações de alto risco (BRCA1/2 positivo, síndrome de Lynch). Combinar com HE4 (índice ROMA) aumenta especificidade para 95% na diferenciação massa benigna versus maligna."""
    },

    "CEA": {
        "clinical_significance": """O Antígeno Carcinoembrionário (CEA) é uma glicoproteína oncofetal normalmente expressa durante desenvolvimento embrionário e suprimida após nascimento, sendo reativada em adenocarcinomas gastrointestinais. Sua função fisiológica inclui adesão celular e inibição de apoptose. Elevações séricas ocorrem em 60-90% dos carcinomas colorretais avançados, 50-70% dos adenocarcinomas pancreáticos e 30-50% dos cânceres gástricos.

Mecanismos de elevação incluem secreção tumoral direta, destruição tecidual com liberação na circulação e redução de clearance hepático. Valores normais são inferiores a 3 ng/mL em não-fumantes e até 5 ng/mL em fumantes. Concentrações entre 5-10 ng/mL sugerem condições benignas (doença inflamatória intestinal, hepatopatias, pancreatite), enquanto valores >20 ng/mL requerem investigação neoplásica urgente.

Em pacientes com câncer colorretal ressecado, CEA é o marcador mais sensível para detecção de recidiva, com elevação precedendo sintomas em 3-18 meses. Cinética pós-operatória é crítica: normalização em 4-6 semanas indica ressecção completa, enquanto elevação sustentada sugere doença residual ou metástases ocultas.""",

        "longevity_context": """CEA crônicamente elevado em faixas sub-clínicas (3-10 ng/mL) associa-se a envelhecimento acelerado do trato gastrointestinal, refletindo inflamação mucosa persistente e aumento de permeabilidade intestinal. Estudos prospectivos demonstram que indivíduos com CEA basal >5 ng/mL apresentam risco 40% maior de desenvolver adenomas colorretais em 10 anos.

No contexto de medicina preventiva, monitoramento anual de CEA em indivíduos >50 anos com histórico familiar de câncer colorretal permite detecção precoce de lesões pré-malignas. Valores consistentemente baixos (<2 ng/mL) correlacionam-se com microbiota intestinal saudável e menor inflamação sistêmica.""",

        "clinical_recommendations": """Indicar CEA pré-operatório em todo paciente com adenocarcinoma colorretal para estabelecer baseline e valor prognóstico. Níveis >5 ng/mL pré-cirúrgicos associam-se a estádio mais avançado e pior sobrevida. Monitoramento pós-ressecção deve ser trimestral nos primeiros 2 anos (período de maior risco de recidiva) e semestral até 5 anos.

Elevação de CEA >30% em relação ao último valor requer investigação com tomografia de tórax/abdome/pelve e PET-scan. Em doença metastática, CEA serve como marcador de resposta terapêutica: redução >50% após 2 meses correlaciona-se com sobrevida prolongada. Tabagismo aumenta CEA em 20-30%, necessário considerar na interpretação.

Para rastreamento, CEA isolado não é recomendado devido baixa especificidade. Combinar com pesquisa de sangue oculto nas fezes e colonoscopia é estratégia padrão. Em pacientes com doença inflamatória intestinal, elevações de CEA podem preceder exacerbação clínica em 1-2 meses."""
    },

    "CHCM (MCHC)": {
        "clinical_significance": """A Concentração de Hemoglobina Corpuscular Média (CHCM/MCHC) expressa a concentração média de hemoglobina no volume de hemácias, medida em g/dL. Valores de referência situam-se entre 32-36 g/dL, refletindo grau de hemoglobinização eritrocitária. Este índice hematimétrico é calculado pela razão: (Hemoglobina/Hematócrito) × 100, sendo menos suscetível a artefatos de armazenamento que outros índices.

Diminuições (hipocromia, CHCM <32 g/dL) indicam deficiência na síntese de hemoglobina, característica de anemia ferropriva avançada, talassemias e anemia de doença crônica. Mecanisticamente, déficit de ferro limita incorporação de heme, resultando em hemácias pálidas centralmente. CHCM <28 g/dL sugere anemia ferropriva severa ou talassemia major.

Elevações (CHCM >36 g/dL) são raras e tipicamente indicam esferocitose hereditária, onde perda de membrana eritrocitária aumenta concentração relativa de hemoglobina. Valores >38 g/dL podem reflettar artefato laboratorial (hemólise in vitro, crioglobulinemia) ou desidratação eritrocitária severa. Combinação CHCM elevado + VCM reduzido é patognomônica de esferocitose.""",

        "longevity_context": """CHCM na faixa superior da normalidade (34-36 g/dL) associa-se a melhor capacidade de transporte de oxigênio e performance aeróbica em idosos. Estudos de centenários demonstram manutenção de CHCM >33 g/dL, sugerindo preservação da função medular e metabolismo de ferro.

Declínios progressivos de CHCM (<32 g/dL) em idosos podem preceder diagnóstico de anemia em 6-12 meses, permitindo intervenção nutricional precoce. Valores otimizados correlacionam-se com menor fadiga, melhor função cognitiva e redução de risco de quedas em populações geriátricas.""",

        "clinical_recommendations": """CHCM deve ser interpretado conjuntamente com VCM e RDW no diagnóstico diferencial de anemias. Hipocromia (CHCM baixo) + microcitose (VCM <80 fL) + RDW elevado sugere deficiência de ferro, enquanto hipocromia + microcitose + RDW normal indica traço talassêmico.

Em pacientes com anemia ferropriva, CHCM normaliza lentamente durante reposição de ferro (2-3 meses), sendo útil para monitorar resposta terapêutica. Valores que não se elevam após 8 semanas de ferroterapia adequada sugerem má absorção ou perda sanguínea oculta persistente.

Para otimização em medicina preventiva, manter CHCM >33 g/dL através de suplementação de ferro quelato (se ferritina <50 ng/mL), vitamina C para absorção e investigação de causas de perda (sangramento gastrointestinal oculto, menorragia). CHCM <30 g/dL requer investigação urgente de anemia ferropriva severa."""
    },

    "CIMT Carótidas (Espessura Íntima-Média)": {
        "clinical_significance": """A Espessura Íntima-Média Carotídea (CIMT) quantifica espessura da parede arterial das carótidas comuns através de ultrassom de alta resolução, medindo distância entre interface lúmen-íntima e interface média-adventícia. Valores normais são <0.9 mm em adultos <50 anos e <1.0 mm em >50 anos. CIMT reflete aterosclerose subclínica, resultante de disfunção endotelial crônica, infiltração lipídica e remodelamento vascular.

Incrementos de 0.1 mm no CIMT associam-se a aumento de 13-15% no risco de infarto do miocárdio e 13-18% no risco de AVC isquêmico. Mecanisticamente, espessamento íntima-média precede formação de placa aterosclerótica em décadas, permitindo janela terapêutica ampla para prevenção primária. Progressão >0.02 mm/ano indica aterosclerose ativa que requer intensificação terapêutica.

A mensuração deve incluir média de 6-12 pontos na parede posterior da carótida comum, 1 cm proximal à bifurcação, onde fluxo laminar minimiza artefatos. Presença de placa (espessamento focal >1.5 mm ou >50% da parede adjacente) reclassifica paciente para risco cardiovascular alto, independente de escores tradicionais.""",

        "longevity_context": """CIMT é marcador preditivo de longevidade cardiovascular, refletindo idade vascular real versus idade cronológica. Indivíduos com CIMT <0.7 mm aos 50 anos apresentam expectativa de vida cardiovascular 8-10 anos maior que população geral. Estudos de intervenção demonstram que cada 0.1 mm de redução de CIMT correlaciona-se com 20% de redução em eventos cardiovasculares maiores.

Centenários consistentemente apresentam CIMT <1.0 mm, sugerindo que preservação da integridade vascular é determinante crítico de envelhecimento saudável. Progressão lenta de CIMT (<0.01 mm/ano) associa-se a telômeros leucocitários mais longos e menor inflamação sistêmica.""",

        "clinical_recommendations": """Indicar CIMT em pacientes com risco cardiovascular intermediário (Framingham 10-20%) para reclassificação e decisão sobre terapia agressiva com estatina. Realizar baseline aos 40-45 anos em pacientes com história familiar de doença cardiovascular precoce, seguido de reavaliação a cada 3-5 anos.

CIMT >75º percentil para idade e sexo ou presença de placa reclassifica para risco alto, indicando estatina de alta potência (atorvastatina 40-80 mg, rosuvastatina 20-40 mg) com meta LDL <70 mg/dL. Progressão >0.02 mm/ano requer intensificação terapêutica: adicionar ezetimiba, otimizar controle pressórico (<130/80 mmHg) e considerar anti-agregação.

Para monitoramento de resposta, repetir CIMT após 2-3 anos de terapia otimizada. Regressão ou estabilização (<0.01 mm/ano) indica controle adequado. Fatores que aceleram progressão incluem LDL >100 mg/dL, pressão sistólica >140 mmHg, diabetes mal controlado (HbA1c >7%) e inflamação crônica (PCR-us >3 mg/L)."""
    },

    "Calprotectina fecal": {
        "clinical_significance": """Calprotectina é proteína citoplasmática liberada por neutrófilos ativados na mucosa intestinal, representando biomarcador sensível de inflamação gastrointestinal. Sua estabilidade em fezes (até 7 dias em temperatura ambiente) permite dosagem não-invasiva de atividade inflamatória intestinal. Níveis normais são <50 μg/g, refletindo turnover fisiológico de neutrófilos mucosos.

Elevações >150 μg/g indicam inflamação intestinal ativa com probabilidade >90% de doença orgânica (doença inflamatória intestinal, infecções bacterianas, neoplasias). Valores entre 50-150 μg/g são zona cinzenta, sugerindo inflamação leve que requer correlação clínica. Na doença de Crohn e retocolite ulcerativa, calprotectina correlaciona-se com índices endoscópicos de atividade, permitindo monitoramento não-invasivo de cicatrização mucosa.

Mecanisticamente, neutrófilos infiltram lâmina própria em resposta a disrupção de barreira epitelial, liberando calprotectina que é eliminada nas fezes. Concentrações >250 μg/g associam-se a ulcerações mucosas profundas, enquanto valores <100 μg/g sugerem remissão endoscópica. Sensibilidade de 95% e especificidade de 90% para distinguir doença inflamatória intestinal de síndrome do intestino irritável.""",

        "longevity_context": """Calprotectina fecal elevada cronicamente, mesmo em faixas sub-clínicas (100-200 μg/g), associa-se a envelhecimento intestinal acelerado e aumento de permeabilidade. Estudos demonstram que indivíduos >60 anos com calprotectina persistentemente <50 μg/g apresentam microbiota diversificada e menor incidência de síndrome metabólica.

No contexto de medicina preventiva, elevações transitórias podem preceder desenvolvimento de doença inflamatória intestinal em indivíduos geneticamente predispostos. Monitoramento anual permite identificação precoce de inflamação subclínica e intervenção dietética antes de manifestações clínicas.""",

        "clinical_recommendations": """Solicitar calprotectina em pacientes <45 anos com sintomas intestinais crônicos (diarreia, dor abdominal, sangramento) para discriminar síndrome do intestino irritável de doença inflamatória intestinal, evitando colonoscopias desnecessárias. Valores <50 μg/g tornam doença inflamatória intestinal improvável (valor preditivo negativo 96%).

Para monitoramento de doença inflamatória intestinal em remissão, medir calprotectina a cada 3-4 meses. Elevações >250 μg/g precedem recidiva clínica em 2-3 meses, permitindo ajuste terapêutico preemptivo. Meta em tratamento: calprotectina <100 μg/g, correlacionando-se com cicatrização mucosa completa.

Valores >1000 μg/g sugerem colite severa ou infecção intestinal (Clostridium difficile, Campylobacter) e requerem avaliação urgente. Fatores que elevam calprotectina incluem uso recente de AINEs, infecções virais e consumo excessivo de álcool. Para interpretação ótima, coletar amostra fora de fase menstrual e após suspensão de AINEs por 2 semanas."""
    }
}


class LabEnricher:
    def __init__(self):
        self.token = None

    def login(self) -> str:
        """Autentica e retorna JWT token"""
        print("Fazendo login...")
        response = requests.post(
            f"{API_BASE}/auth/login",
            json={"email": LOGIN_EMAIL, "password": LOGIN_PASSWORD}
        )
        response.raise_for_status()
        self.token = response.json()["accessToken"]
        print(f"✓ Login realizado com sucesso")
        return self.token

    def get_headers(self) -> Dict[str, str]:
        """Retorna headers com autenticação"""
        return {
            "Authorization": f"Bearer {self.token}",
            "Content-Type": "application/json"
        }

    def get_lab_item(self, item_id: str) -> Dict[str, Any]:
        """Busca um item de exame pelo ID"""
        response = requests.get(
            f"{API_BASE}/lab-tests/definitions/{item_id}",
            headers=self.get_headers()
        )
        response.raise_for_status()
        data = response.json()
        # API pode retornar direto ou com wrapper "data"
        if isinstance(data, dict) and "data" in data:
            return data["data"]
        return data

    def update_lab_item(self, item_id: str, current_item: Dict[str, Any], enriched_data: Dict[str, str]) -> bool:
        """Atualiza item com conteúdo enriquecido"""
        try:
            # Merge dos dados enriquecidos com o item atual
            updated_item = {**current_item, **enriched_data}

            response = requests.put(
                f"{API_BASE}/lab-tests/definitions/{item_id}",
                headers=self.get_headers(),
                json=updated_item
            )
            response.raise_for_status()
            return True
        except Exception as e:
            print(f"   ⚠ Erro ao atualizar item: {e}")
            return False

    def process_batch(self):
        """Processa os primeiros 5 items com conteúdo manual"""
        print("\n" + "="*80)
        print("BATCH LAB PART 5 - 5 Items Iniciais (Conteúdo Manual)")
        print("="*80 + "\n")

        results = {
            "success": [],
            "failed": [],
            "skipped": []
        }

        # Processar apenas os 5 primeiros items que têm conteúdo definido
        item_names = ["CA-125", "CEA", "CHCM (MCHC)", "CIMT Carótidas (Espessura Íntima-Média)", "Calprotectina fecal"]

        for idx, (item_id, item_name) in enumerate(zip(ITEM_IDS[:5], item_names), 1):
            print(f"\n[{idx}/5] Processando item: {item_id}")

            try:
                # 1. Buscar item atual
                item = self.get_lab_item(item_id)
                print(f"   Nome: {item.get('name', 'N/A')}")
                print(f"   Categoria: {item.get('category', 'N/A')}")

                # Verificar se já está enriquecido
                if item.get('clinicalSignificance') and len(item.get('clinicalSignificance', '')) > 100:
                    print(f"   ⊘ Item já possui conteúdo clínico extenso, pulando...")
                    results["skipped"].append(item_id)
                    continue

                # 2. Buscar conteúdo pré-definido
                content_key = item_name
                if content_key in CLINICAL_CONTENT:
                    # Converter para formato da API (camelCase)
                    source_data = CLINICAL_CONTENT[content_key]
                    enriched_data = {
                        "clinicalSignificance": source_data["clinical_significance"],
                        "longevityContext": source_data["longevity_context"],
                        "clinicalRecommendations": source_data["clinical_recommendations"]
                    }
                    print(f"   ✓ Conteúdo clínico encontrado")
                else:
                    print(f"   ⚠ Conteúdo não disponível para: {item_name}")
                    results["skipped"].append(item_id)
                    continue

                # 3. Atualizar no banco
                if self.update_lab_item(item_id, item, enriched_data):
                    print(f"   ✓ Item atualizado com sucesso")
                    print(f"      - Clinical Significance: {len(enriched_data['clinicalSignificance'])} caracteres")
                    print(f"      - Longevity Context: {len(enriched_data['longevityContext'])} caracteres")
                    print(f"      - Clinical Recommendations: {len(enriched_data['clinicalRecommendations'])} caracteres")
                    results["success"].append(item_id)
                else:
                    results["failed"].append(item_id)

                time.sleep(0.5)

            except Exception as e:
                print(f"   ✗ Erro ao processar item: {e}")
                results["failed"].append(item_id)

        # Relatório final
        print("\n" + "="*80)
        print("RELATÓRIO FINAL")
        print("="*80)
        print(f"✓ Sucesso: {len(results['success'])}")
        print(f"⊘ Pulados: {len(results['skipped'])}")
        print(f"✗ Falhas: {len(results['failed'])}")
        print("="*80 + "\n")

        # Salvar relatório
        with open("/home/user/plenya/batch_lab_part5_report.json", "w") as f:
            json.dump(results, f, indent=2)

        return results


def main():
    enricher = LabEnricher()
    enricher.login()
    enricher.process_batch()


if __name__ == "__main__":
    main()
