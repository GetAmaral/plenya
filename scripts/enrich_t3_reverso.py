#!/usr/bin/env python3
"""
Enriquecimento do item T3 Reverso (rT3) com conteúdo científico.
Item ID: 4159c2e3-97e2-4ffc-922d-4513fdbc82aa
"""

import psycopg2
from datetime import datetime
import sys

# Dados dos artigos científicos
ARTICLES = [
    {
        "title": "Reverse T3 in patients with hypothyroidism on different thyroid hormone replacement",
        "authors": "Wilson JB, Hoang TD, Lee ML, Epstein M, Friedman TC",
        "journal": "PLoS One",
        "year": 2025,
        "doi": "10.1371/journal.pone.0325046",
        "url": "https://pmc.ncbi.nlm.nih.gov/articles/PMC12148341/",
        "summary": """Estudo de 2025 com 976 pacientes que investigou prevalência de rT3 elevado em diferentes
        regimes de reposição tireoidiana. Encontrou 11% com rT3 elevado, sendo 20.9% no grupo usando apenas L-T4
        e apenas 3.5% no grupo usando extrato tireoidiano dessecado. O rT3 correlacionou fortemente com T4 livre
        e T3 livre, inversamente com TSH. Desafia o descarte convencional da medição de rT3 e sugere relevância
        clínica para compreender sintomas persistentes em pacientes hipotireoideos."""
    },
    {
        "title": "Clinical and laboratory aspects of 3,3',5'-triiodothyronine (reverse T3)",
        "authors": "Halsall DJ, Oddy S",
        "journal": "Annals of Clinical Biochemistry",
        "year": 2021,
        "doi": "10.1177/0004563220969150",
        "url": "https://pubmed.ncbi.nlm.nih.gov/33040575/",
        "summary": """Revisão sobre aspectos clínicos e laboratoriais do rT3. É produzido a partir da tiroxina
        via desiodinação do anel interno e representa um ponto final metabólico inativo. Aumenta durante síndrome
        do doente eutireoideo e com medicações como amiodarona. Métodos espectrométricos de massa substituíram
        radioimunoensaio, oferecendo redução de interferência. Níveis séricos afetados por condições genéticas
        envolvendo desiodases, transportadores tireoidianos e proteínas de transporte."""
    },
    {
        "title": "Euthyroid Sick Syndrome: Practice Essentials, Pathophysiology, Epidemiology",
        "authors": "Medscape Reference",
        "journal": "Medscape",
        "year": 2024,
        "doi": "",
        "url": "https://emedicine.medscape.com/article/118651-overview",
        "summary": """Revisão clínica sobre síndrome do doente eutireoideo (nonthyroidal illness syndrome).
        Caracterizada por testes de função tireoidiana anormais durante doença não tireoidiana, sem disfunção
        prévia da tireoide ou hipófise, completamente reversível após recuperação. Redução de T3 ocorre em
        40-100% dos casos de doença não tireoidiana. rT3 elevado é achado característico. Mortalidade aumenta
        significativamente quando T4 <4 mcg/dL (50% mortalidade) e <2 mcg/dL (80% mortalidade). Em COVID-19,
        presença de síndrome do doente eutireoideo piora prognóstico (34.1% vs 11.3% mortalidade)."""
    }
]

# Conteúdo clínico em PT-BR
CLINICAL_CONTENT = {
    "clinical_relevance": """T3 Reverso (rT3) é uma forma biologicamente inativa do hormônio tireoidiano T3,
produzida pela desiodinação do anel interno da tiroxina (T4) pelas enzimas desiodases tipo 1 e tipo 3 (D1 e D3).

**Significado Clínico:**

1. **Síndrome do Doente Eutireoideo:** Níveis elevados de rT3 são marcadores característicos desta condição,
observada em 40-100% dos pacientes com doenças graves não tireoidianas. O rT3 elevado ocorre simultaneamente
com T3 baixo ("síndrome do T3 baixo").

2. **Mecanismo Adaptativo:** O aumento do rT3 é considerado um mecanismo protetor do organismo para reduzir
a taxa metabólica e conservar energia durante períodos de estresse severo, doença crítica, jejum prolongado
ou inflamação crônica.

3. **Razão T3/rT3:** A razão T3/rT3 deve ser maior que 10 (pelo menos 10 vezes mais T3 do que rT3 no sangue).
Razão <10 indica conversão periférica ineficiente de T4 em T3 ativo, com produção excessiva de rT3 inativo.

**Condições Associadas a rT3 Elevado:**
- Doenças críticas (sepse, infarto do miocárdio, trauma)
- Estresse crônico e inflamação sistêmica
- Restrição calórica extrema ou jejum prolongado
- Insuficiência cardíaca crônica
- Uso de medicações (amiodarona, propranolol, glicocorticoides)
- Hipotireoidismo tratado apenas com L-T4 (20.9% dos pacientes)

**Importância Prognóstica:**
- Níveis de T4 <4 mcg/dL associados a ~50% mortalidade; <2 mcg/dL com ~80% mortalidade
- Em pacientes COVID-19 com síndrome do doente eutireoideo: 34.1% mortalidade vs 11.3% sem a síndrome
- Normalização completa esperada após recuperação da doença de base

**Métodos Laboratoriais:**
Métodos espectrométricos de massa modernos substituíram radioimunoensaios, oferecendo maior precisão
e redução de interferências de outros iodotironinas.""",

    "patient_explanation": """O T3 Reverso (rT3) é uma forma "inativa" de hormônio da tireoide que seu corpo
produz a partir do hormônio T4.

**Por que é importante?**

Imagine que os hormônios da tireoide são como o acelerador do seu metabolismo. O T3 é a forma "ativa" que
acelera suas células. Já o T3 Reverso (rT3) é como um freio - ele bloqueia a ação do T3 ativo.

**Quando o rT3 aumenta?**

Seu corpo aumenta o rT3 como mecanismo de proteção em situações de:
- **Doenças graves:** Infecções sérias, internação hospitalar, cirurgias grandes
- **Estresse prolongado:** Estresse crônico físico ou emocional
- **Jejum ou dietas muito restritivas:** Quando o corpo precisa economizar energia
- **Inflamação crônica:** Doenças inflamatórias sistêmicas
- **Problemas cardíacos:** Insuficiência cardíaca

**O que significa resultado alterado?**

**rT3 elevado ou Razão T3/rT3 <10:**
- Pode indicar que seu corpo está em "modo de economia de energia"
- Comum em pessoas com fadiga persistente mesmo com TSH normal
- Pode explicar sintomas de hipotireoidismo apesar de exames "normais"
- Em doenças graves, valores muito baixos de T4 indicam pior prognóstico

**Importante saber:**
- O rT3 geralmente normaliza completamente após recuperação da doença de base
- Não é um exame de rotina - é solicitado em situações específicas
- Pacientes usando apenas T4 (Puran, Levoid) tendem a ter rT3 mais alto
- Preparações com T3 (como T3 isolado ou extrato tireoidiano) resultam em rT3 mais baixo

**Quando avisar o médico:**
- Fadiga persistente mesmo com tratamento tireoidiano adequado
- Sintomas de hipotireoidismo com TSH controlado
- Após doença grave ou estresse prolongado
- Dificuldade de perder peso apesar de dieta e exercício""",

    "conduct": """**Interpretação de Resultados:**

1. **Valores de Referência:**
   - rT3: 9-27 ng/dL (varia conforme laboratório e método)
   - Razão T3/rT3: >10 (ideal)
   - Razão <10: sugere conversão periférica prejudicada

2. **rT3 Elevado Isolado:**
   - Avaliar história clínica detalhada: doença recente, estresse, dieta, medicações
   - Solicitar painel tireoidiano completo: TSH, T4 livre, T3 livre, rT3
   - Investigar causas secundárias: inflamação, estresse crônico, restrição calórica

3. **Síndrome do Doente Eutireoideo (baixo T3 + alto rT3):**
   - Focar tratamento da doença de base
   - **NÃO** iniciar reposição tireoidiana em pacientes agudamente doentes
   - Reavaliar função tireoidiana 6-8 semanas após recuperação
   - Normalização espontânea esperada com resolução da doença

4. **Hipotireoidismo com rT3 Elevado:**
   - Considerar otimização da dose de levotiroxina
   - Avaliar necessidade de adicionar T3 ao regime (controverso)
   - Monitorar resposta clínica (energia, cognição, peso) além dos laboratoriais
   - Descartar causas de má absorção de levotiroxina

**Abordagem Terapêutica (Evidências Limitadas):**

**Medidas Gerais:**
- Gerenciamento de estresse: técnicas de relaxamento, sono adequado
- Nutrição adequada: evitar restrição calórica extrema (<1200 kcal/dia)
- Suporte nutricional: selênio (200 mcg/dia), zinco, ferro se deficientes
- Exercício moderado: evitar overtraining

**Ajuste de Medicação Tireoidiana:**
- Pacientes em L-T4 isolado com sintomas persistentes:
  - Verificar TSH, T4 livre, T3 livre, rT3
  - Se T3 livre baixo/normal-baixo + rT3 elevado: considerar adicionar T3
  - Preparações combinadas T4/T3 (ex: 95 mcg T4 + 5 mcg T3) podem reduzir rT3
  - Extrato tireoidiano dessecado associado a menor prevalência de rT3 elevado (3.5%)

**Causas Reversíveis a Corrigir:**
- Medicações: avaliar descontinuação/substituição de beta-bloqueadores, amiodarona, glicocorticoides
- Inflamação crônica: tratar condições inflamatórias subjacentes
- Resistência insulínica: otimizar controle glicêmico
- Deficiências nutricionais: repor selênio, zinco, ferro, vitamina D

**Monitoramento:**
- Reavaliar TSH, T4L, T3L, rT3 após 6-8 semanas de qualquer mudança
- Focar em melhora clínica (energia, peso, cognição) além dos laboratoriais
- Em doença crítica: rT3 não guia tratamento, foco na doença de base

**Controvérsias:**
- Medição rotineira de rT3 NÃO é recomendada por sociedades endocrinológicas tradicionais
- Medicina funcional valoriza mais o rT3 para explicar sintomas persistentes
- Evidências sobre tratamento baseado em rT3 são limitadas e controversas
- Estudos prospectivos necessários para estabelecer causalidade entre rT3 e fadiga

**Indicações para Encaminhamento ao Endocrinologista:**
- Síndrome do doente eutireoideo em paciente crítico
- Sintomas persistentes de hipotireoidismo apesar de TSH adequado
- Dúvida sobre necessidade de adicionar T3 ao regime
- rT3 persistentemente elevado sem causa identificada
- Paciente com múltiplas comorbidades e alterações tireoidianas complexas

**Situações de Urgência (T4 muito baixo em doença grave):**
- T4 total <4 mcg/dL: risco de mortalidade ~50%
- T4 total <2 mcg/dL: risco de mortalidade ~80%
- Considerar suporte intensivo e manejo agressivo da doença de base
- Evitar reposição tireoidiana empírica em fase aguda"""
}

def connect_db():
    """Conecta ao banco de dados PostgreSQL."""
    return psycopg2.connect(
        host="localhost",
        port=5432,
        database="plenya_db",
        user="plenya_user",
        password="senha_segura"
    )

def article_exists(cursor, doi=None, title=None):
    """Verifica se artigo já existe no banco."""
    if doi and doi.strip():
        cursor.execute("SELECT id FROM articles WHERE doi = %s", (doi,))
        result = cursor.fetchone()
        if result:
            return result[0]

    if title:
        cursor.execute("SELECT id FROM articles WHERE title = %s", (title,))
        result = cursor.fetchone()
        if result:
            return result[0]

    return None

def insert_article(cursor, article):
    """Insere artigo no banco de dados."""
    existing_id = article_exists(cursor, article.get('doi'), article['title'])
    if existing_id:
        print(f"  ℹ️  Artigo já existe: {article['title'][:60]}... (ID: {existing_id})")
        return existing_id

    # Converter ano para data (primeiro dia do ano)
    publish_date = f"{article['year']}-01-01"

    cursor.execute("""
        INSERT INTO articles (
            title, authors, journal, publish_date, doi,
            original_link, abstract, created_at, updated_at
        )
        VALUES (%s, %s, %s, %s, %s, %s, %s, NOW(), NOW())
        RETURNING id
    """, (
        article['title'],
        article['authors'],
        article['journal'],
        publish_date,
        article['doi'] if article['doi'] else None,
        article['url'],
        article['summary']
    ))

    article_id = cursor.fetchone()[0]
    print(f"  ✅ Artigo inserido: {article['title'][:60]}... (ID: {article_id})")
    return article_id

def create_article_relation(cursor, article_id, score_item_id):
    """Cria relação many-to-many entre article e score_item."""
    # Verificar se relação já existe
    cursor.execute("""
        SELECT 1 FROM article_score_items
        WHERE article_id = %s AND score_item_id = %s
    """, (article_id, score_item_id))

    if cursor.fetchone():
        print(f"    ℹ️  Relação já existe (article {article_id} <-> score_item)")
        return

    cursor.execute("""
        INSERT INTO article_score_items (article_id, score_item_id)
        VALUES (%s, %s)
    """, (article_id, score_item_id))
    print(f"    ✅ Relação criada (article {article_id} <-> score_item)")

def update_score_item(cursor, score_item_id, content):
    """Atualiza campos clínicos do score_item."""
    cursor.execute("""
        UPDATE score_items
        SET
            clinical_relevance = %s,
            patient_explanation = %s,
            conduct = %s,
            last_review = %s,
            updated_at = NOW()
        WHERE id = %s
    """, (
        content['clinical_relevance'],
        content['patient_explanation'],
        content['conduct'],
        datetime.now(),
        score_item_id
    ))
    print(f"  ✅ Score item atualizado com conteúdo clínico (ID: {score_item_id})")

def main():
    """Função principal."""
    score_item_id = "4159c2e3-97e2-4ffc-922d-4513fdbc82aa"

    print("=" * 80)
    print("ENRIQUECIMENTO: T3 Reverso (rT3)")
    print("=" * 80)

    try:
        conn = connect_db()
        cursor = conn.cursor()

        # Verificar se o score_item existe
        cursor.execute("SELECT name FROM score_items WHERE id = %s", (score_item_id,))
        result = cursor.fetchone()
        if not result:
            print(f"❌ Erro: Score item {score_item_id} não encontrado no banco!")
            return 1

        print(f"\n📋 Item encontrado: {result[0]}")

        # 1. Inserir artigos
        print(f"\n{'─' * 80}")
        print("1️⃣  INSERINDO ARTIGOS CIENTÍFICOS")
        print(f"{'─' * 80}")

        article_ids = []
        for i, article in enumerate(ARTICLES, 1):
            print(f"\nArtigo {i}/{len(ARTICLES)}:")
            article_id = insert_article(cursor, article)
            article_ids.append(article_id)

        # 2. Criar relações article-score_item
        print(f"\n{'─' * 80}")
        print("2️⃣  CRIANDO RELAÇÕES ARTICLE ↔ SCORE_ITEM")
        print(f"{'─' * 80}\n")

        for article_id in article_ids:
            create_article_relation(cursor, article_id, score_item_id)

        # 3. Atualizar score_item com conteúdo clínico
        print(f"\n{'─' * 80}")
        print("3️⃣  ATUALIZANDO SCORE_ITEM COM CONTEÚDO CLÍNICO")
        print(f"{'─' * 80}\n")

        update_score_item(cursor, score_item_id, CLINICAL_CONTENT)

        # Commit das mudanças
        conn.commit()

        print(f"\n{'═' * 80}")
        print("✅ ENRIQUECIMENTO CONCLUÍDO COM SUCESSO!")
        print(f"{'═' * 80}")
        print(f"""
📊 RESUMO:
   • Score Item: T3 Reverso (rT3)
   • ID: {score_item_id}
   • Artigos inseridos: {len(article_ids)}
   • Relações criadas: {len(article_ids)}
   • Campos atualizados: clinical_relevance, patient_explanation, conduct, last_review
   • Data de revisão: {datetime.now().strftime('%Y-%m-%d %H:%M:%S')}
        """)

        cursor.close()
        conn.close()
        return 0

    except Exception as e:
        print(f"\n❌ Erro: {e}")
        import traceback
        traceback.print_exc()
        return 1

if __name__ == "__main__":
    sys.exit(main())
