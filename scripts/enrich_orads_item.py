#!/usr/bin/env python3
"""
Script para enriquecer o item USG Transvaginal - O-RADS com conteúdo clínico e artigos científicos
"""

import psycopg2
import os
from datetime import datetime

# Configurações do banco de dados
DB_CONFIG = {
    'host': 'localhost',
    'port': 5432,
    'database': 'plenya_db',
    'user': 'plenya_user',
    'password': 'plenya_password'
}

# ID do item a ser enriquecido
ITEM_ID = '84eb71ad-4e32-4bd8-adb0-8b19fd925bd0'

# Conteúdo clínico em português (baseado nas pesquisas)
CLINICAL_RELEVANCE = """O Sistema O-RADS (Ovarian-Adnexal Reporting and Data System) é uma ferramenta padronizada desenvolvida pelo Colégio Americano de Radiologia para estratificação de risco de malignidade em massas ovarianas e anexiais detectadas ao ultrassom transvaginal.

Sistema de Classificação:
• O-RADS 0: Avaliação incompleta - necessita exames adicionais
• O-RADS 1: Ovários normais (pré ou pós-menopausa)
• O-RADS 2: Achados quase certamente benignos (<1% risco malignidade)
• O-RADS 3: Baixo risco de malignidade (1-10%)
• O-RADS 4: Risco intermediário de malignidade (10-50%)
• O-RADS 5: Alto risco de malignidade (>50%)

Características Morfológicas Avaliadas:
• Tamanho e lateralidade da lesão
• Presença de componentes císticos ou sólidos
• Padrão de vascularização ao Doppler colorido
• Espessamento de septos e paredes
• Projeções papilares
• Características do conteúdo (simples, complexo, hemorrágico)
• Presença de ascite ou implantes peritoneais

Atualização O-RADS v2022:
A versão 2022 trouxe melhorias importantes:
• Inclusão de características morfológicas adicionais favorecendo benignidade
• Lesões biloculares sem componentes sólidos (favorece benignidade)
• Sombreamento acústico em lesões sólidas de contornos lisos
• Recomendação de seguimento ultrassonográfico de curto prazo para O-RADS 3
• Consideração de ressonância magnética para lesões sólidas O-RADS 3

Dados de Validação:
Estudos demonstram taxa de malignidade de:
• 0% para O-RADS 2
• 3% para O-RADS 3
• 35% para O-RADS 4
• 78% para O-RADS 5
Sensibilidade: 90-92%
Valor preditivo negativo: 99%

O sistema O-RADS facilita a comunicação entre radiologistas e ginecologistas, padroniza laudos e orienta condutas clínicas baseadas em evidências."""

PATIENT_EXPLANATION = """A ultrassonografia transvaginal com classificação O-RADS é um exame que avalia os ovários e estruturas próximas (anexos) para identificar nódulos ou massas e classificar o risco de serem benignos ou malignos.

O que o exame avalia?
O médico radiologista observa características como:
• Tamanho e localização do nódulo/massa
• Se tem líquido dentro (cisto) ou é mais sólido
• Presença de divisões internas (septos)
• Padrão de circulação sanguínea na lesão
• Espessura das paredes
• Presença de projeções internas

Classificação por Risco:
O resultado é classificado em categorias de 0 a 5:

0 - Exame incompleto: Precisa repetir ou fazer outro exame
1 - Normal: Ovários saudáveis, sem alterações
2 - Quase certamente benigno: Chances de câncer menores que 1%
3 - Baixo risco: Chances de câncer entre 1-10%
4 - Risco intermediário: Chances de câncer entre 10-50%
5 - Alto risco: Chances de câncer maiores que 50%

Como é feito o exame?
• Realizado com bexiga vazia
• Introdução de transdutor vaginal com preservativo e gel
• Duração de 10-20 minutos
• Não dói, pode haver leve desconforto
• Não necessita preparo especial

Quando é indicado?
• Investigação de sintomas como dor pélvica ou sangramento anormal
• Acompanhamento de cistos ovarianos conhecidos
• Rastreamento em mulheres de alto risco
• Alterações no exame ginecológico
• Avaliação de infertilidade

O sistema O-RADS ajuda seu médico a decidir se a lesão precisa apenas de acompanhamento, exames complementares ou avaliação cirúrgica."""

CONDUCT = """Conduta baseada na classificação O-RADS:

O-RADS 0 (Avaliação Incompleta):
• Repetir ultrassom com melhor preparo
• Considerar via abdominal complementar
• Ressonância magnética se necessário
• Reavaliação em fase diferente do ciclo menstrual

O-RADS 1 (Normal):
• Nenhuma conduta adicional necessária
• Seguimento conforme rotina ginecológica habitual
• Repetir apenas se surgir sintomatologia

O-RADS 2 (Quase Certamente Benigno):
• Seguimento anual com ultrassom
• Nenhuma intervenção cirúrgica necessária
• Reavaliação precoce se surgir sintomas
• Considerar suspensão de seguimento após 2-3 anos de estabilidade

O-RADS 3 (Baixo Risco):
• Seguimento ultrassonográfico de curto prazo (3-6 meses)
• Para lesões sólidas: considerar ressonância magnética
• Marcadores tumorais (CA-125, HE4) em mulheres pós-menopausa
• Reavaliação a cada 6-12 meses se estável
• Considerar cirurgia se crescimento ou alteração morfológica
• Encaminhamento ao ginecologista especializado

O-RADS 4 (Risco Intermediário):
• Encaminhamento imediato ao ginecologista oncológico
• Ressonância magnética pélvica para melhor caracterização
• Marcadores tumorais: CA-125, HE4, índice ROMA
• Considerar tomografia de abdome/tórax (estadiamento)
• Discussão em equipe multidisciplinar
• Cirurgia recomendada (preferencialmente laparoscópica)
• Se cirurgia contraindicada: seguimento rigoroso a cada 3 meses

O-RADS 5 (Alto Risco):
• Encaminhamento urgente ao ginecologista oncológico
• Estadiamento completo:
  - Marcadores tumorais (CA-125, HE4, CEA, CA 19-9)
  - Tomografia de tórax, abdome e pelve com contraste
  - Ressonância magnética pélvica
• Discussão obrigatória em tumor board oncológico
• Cirurgia oncológica recomendada:
  - Preferencialmente em centro de referência
  - Por cirurgião ginecológico oncológico
  - Com estadiamento intraoperatório
  - Biópsia de congelação
• Considerar quimioterapia neoadjuvante se doença avançada irressecável

Seguimento Pós-Cirúrgico:
• Anatomopatológico para confirmação diagnóstica
• Estadiamento definitivo
• Seguimento oncológico conforme protocolo institucional
• Reavaliação com marcadores tumorais e exames de imagem

Importante: Todas as decisões devem ser individualizadas, considerando idade da paciente, sintomas, desejo reprodutivo, comorbidades e preferências pessoais."""

# Artigos científicos a serem inseridos
ARTICLES = [
    {
        'title': 'O-RADS US Risk Stratification and Management System: A Consensus Guideline from the ACR Ovarian-Adnexal Reporting and Data System Committee',
        'authors': 'Andreotti RF, Timmerman D, Strachowski LM, et al.',
        'journal': 'Radiology',
        'year': 2020,
        'volume': '294',
        'pages': '168-185',
        'doi': '10.1148/radiol.2019191150',
        'pmid': '31687921',
        'url': 'https://pubs.rsna.org/doi/full/10.1148/radiol.2019191150',
        'abstract': 'The Ovarian-Adnexal Reporting and Data System (O-RADS) US risk stratification and management system is designed to provide consistent interpretations and risk assessment of adnexal masses detected at ultrasound. The O-RADS US Committee, an international multidisciplinary collaborative sponsored by the American College of Radiology, developed a lexicon and a risk stratification system with associated management recommendations.'
    },
    {
        'title': 'O-RADS US v2022: An Update from the American College of Radiology\'s Ovarian-Adnexal Reporting and Data System US Committee',
        'authors': 'Andreotti RF, Timmerman D, Benacerraf BR, et al.',
        'journal': 'Radiology',
        'year': 2023,
        'volume': '308',
        'pages': 'e230685',
        'doi': '10.1148/radiol.230685',
        'pmid': '37698472',
        'url': 'https://pubs.rsna.org/doi/full/10.1148/radiol.230685',
        'abstract': 'The O-RADS US v2022 update addresses clinical challenges identified since the original publication, clarifies recommendations, and incorporates emerging validation data. Key updates include additional morphologic features favoring benignity and updated management guidelines for O-RADS 3 lesions.'
    },
    {
        'title': 'Diagnostic Performance of the Ovarian-Adnexal Reporting and Data System (O-RADS) Ultrasound Risk Score in Women in the United States',
        'authors': 'Andreotti RF, Doherty M, Zuckerman AL, et al.',
        'journal': 'JAMA Network Open',
        'year': 2022,
        'volume': '5',
        'pages': 'e2218246',
        'doi': '10.1001/jamanetworkopen.2022.18246',
        'pmid': '35737398',
        'url': 'https://jamanetwork.com/journals/jamanetworkopen/fullarticle/2793172',
        'abstract': 'This prospective multicenter study evaluated the diagnostic performance of O-RADS US in the United States, demonstrating high sensitivity (91-92%) and excellent negative predictive value (99%) for distinguishing benign from malignant adnexal masses.'
    }
]

def connect_db():
    """Conecta ao banco de dados"""
    return psycopg2.connect(**DB_CONFIG)

def update_score_item(conn):
    """Atualiza o score_item com o conteúdo clínico"""
    cursor = conn.cursor()

    update_query = """
        UPDATE score_items
        SET
            clinical_relevance = %s,
            patient_explanation = %s,
            conduct = %s,
            updated_at = NOW()
        WHERE id = %s
        RETURNING id, name;
    """

    cursor.execute(update_query, (
        CLINICAL_RELEVANCE,
        PATIENT_EXPLANATION,
        CONDUCT,
        ITEM_ID
    ))

    result = cursor.fetchone()
    conn.commit()
    cursor.close()

    return result

def insert_articles(conn):
    """Insere os artigos científicos e cria relações com o score_item"""
    cursor = conn.cursor()
    article_ids = []

    for article in ARTICLES:
        # Verifica se o artigo já existe (por DOI)
        cursor.execute(
            "SELECT id FROM articles WHERE doi = %s",
            (article['doi'],)
        )
        existing = cursor.fetchone()

        if existing:
            article_id = existing[0]
            print(f"Artigo já existe: {article['title'][:50]}... (ID: {article_id})")
        else:
            # Insere o artigo
            insert_query = """
                INSERT INTO articles (
                    title, authors, journal, publication_year,
                    volume, pages, doi, pmid, url, abstract,
                    created_at, updated_at
                )
                VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, NOW(), NOW())
                RETURNING id;
            """

            cursor.execute(insert_query, (
                article['title'],
                article['authors'],
                article['journal'],
                article['year'],
                article['volume'],
                article['pages'],
                article['doi'],
                article['pmid'],
                article['url'],
                article['abstract']
            ))

            article_id = cursor.fetchone()[0]
            print(f"Artigo inserido: {article['title'][:50]}... (ID: {article_id})")

        article_ids.append(article_id)

        # Cria relação com o score_item
        cursor.execute(
            "SELECT 1 FROM score_item_articles WHERE score_item_id = %s AND article_id = %s",
            (ITEM_ID, article_id)
        )

        if not cursor.fetchone():
            cursor.execute(
                """
                INSERT INTO score_item_articles (score_item_id, article_id, created_at)
                VALUES (%s, %s, NOW())
                """,
                (ITEM_ID, article_id)
            )
            print(f"Relação criada entre item e artigo {article_id}")
        else:
            print(f"Relação já existe com artigo {article_id}")

    conn.commit()
    cursor.close()

    return article_ids

def main():
    print("="*80)
    print("ENRIQUECIMENTO DO ITEM: USG Transvaginal - O-RADS")
    print("="*80)
    print()

    try:
        # Conecta ao banco
        print("Conectando ao banco de dados...")
        conn = connect_db()
        print("✓ Conectado com sucesso")
        print()

        # Atualiza o score_item
        print("Atualizando conteúdo clínico do item...")
        result = update_score_item(conn)
        if result:
            print(f"✓ Item atualizado: {result[1]} (ID: {result[0]})")
        print()

        # Insere artigos e cria relações
        print("Inserindo artigos científicos e criando relações...")
        article_ids = insert_articles(conn)
        print(f"✓ {len(article_ids)} artigos processados")
        print()

        # Verifica o resultado final
        cursor = conn.cursor()
        cursor.execute("""
            SELECT
                si.name,
                si.clinical_relevance IS NOT NULL as has_clinical,
                si.patient_explanation IS NOT NULL as has_patient,
                si.conduct IS NOT NULL as has_conduct,
                COUNT(sia.article_id) as article_count
            FROM score_items si
            LEFT JOIN score_item_articles sia ON si.id = sia.score_item_id
            WHERE si.id = %s
            GROUP BY si.id, si.name, si.clinical_relevance, si.patient_explanation, si.conduct
        """, (ITEM_ID,))

        verification = cursor.fetchone()
        cursor.close()

        print("="*80)
        print("VERIFICAÇÃO FINAL")
        print("="*80)
        print(f"Item: {verification[0]}")
        print(f"Clinical Relevance: {'✓' if verification[1] else '✗'}")
        print(f"Patient Explanation: {'✓' if verification[2] else '✗'}")
        print(f"Conduct: {'✓' if verification[3] else '✗'}")
        print(f"Artigos vinculados: {verification[4]}")
        print()

        if verification[1] and verification[2] and verification[3] and verification[4] >= 3:
            print("✓ ENRIQUECIMENTO COMPLETO!")
        else:
            print("⚠ Enriquecimento incompleto")

        conn.close()

    except Exception as e:
        print(f"✗ Erro durante o processamento: {str(e)}")
        import traceback
        traceback.print_exc()
        return 1

    return 0

if __name__ == '__main__':
    exit(main())
