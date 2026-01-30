#!/usr/bin/env python3
"""
Batch Parte 6 - Enriquecimento de 20 Items de Exames Laboratoriais
Items: Cistatina C, Cobre, CoQ10, Colesterol (Total e não-HDL), Colonoscopia scores
"""

import anthropic
import os
import json
import psycopg2
from datetime import datetime

# Configuração
ANTHROPIC_API_KEY = os.getenv("ANTHROPIC_API_KEY") or os.getenv("CLAUDE_API_KEY")
if not ANTHROPIC_API_KEY:
    print("⚠️  ANTHROPIC_API_KEY não encontrada. Tentando usar credenciais do sistema...")
    # Tenta usar sem API key (para ambiente controlado)
    try:
        client = anthropic.Anthropic()
    except Exception as e:
        raise ValueError(f"Não foi possível configurar cliente Anthropic: {e}")
else:
    client = anthropic.Anthropic(api_key=ANTHROPIC_API_KEY)

# IDs dos 20 items
ITEM_IDS = [
    "83d713f2-4b20-43fe-8b7e-c3dea08a3cb1",  # Cilindros Patológicos
    "be74b5ba-6a8b-414b-a6dc-f1b50ae1b84f",  # Cistatina C
    "600b5157-efeb-4863-aa5e-6ae48d291c2d",  # Cistatina C
    "18dc38ee-5cc6-4c51-b48f-933d0bc54933",  # Cobre
    "971326a7-f2f5-405c-82fc-52dfa89be9d0",  # Cobre
    "41a30bb0-39f1-491c-846a-f8842942e4ed",  # Coenzima Q10
    "a30bd1e6-fe1a-444d-a89e-876a5ef1a3f5",  # Coenzima Q10
    "498a8637-8bf5-45e0-891b-a0c51a47ccc1",  # Colesterol Total
    "abe4824c-0a2e-424a-9ae3-5e83a2c085ce",  # Colesterol Total
    "1e19e29f-cd22-4e23-9c8f-40a46f1ac739",  # Colesterol Total
    "fa5aa2b0-5cc2-467f-b78d-83ebd69a714c",  # Colesterol não-HDL
    "ac752e50-6c53-4eaf-b72e-1d4e810251cb",  # Colesterol não-HDL
    "cc6dfdcf-845f-4070-9d5c-4115810e3fc4",  # Colesterol não-HDL
    "4e601b8f-e3e0-484e-9d16-d59e7f1682b2",  # Colonoscopia - Mayo Score UC
    "c5d52258-9298-44f9-978a-92717e6dba4b",  # Colonoscopia - Mayo Score UC
    "59585a72-8bc9-4ab7-a81f-2cf21706bcee",  # Colonoscopia - Mayo Score UC
    "7a182c48-8b53-49e7-baeb-71f282325420",  # Colonoscopia - Número Total Adenomas
    "3e0353cd-a928-4a5e-b96f-c7cea26262d2",  # Colonoscopia - Número Total Adenomas
    "0296e16f-52e2-411a-b4bb-9948ea1eb598",  # Colonoscopia - Número Total Adenomas
    "14dec504-0061-446b-b6f1-4ac464b0e463",  # Colonoscopia - SES-CD Crohn
]

def connect_db():
    """Conectar ao banco de dados"""
    return psycopg2.connect(
        host="localhost",
        port=5432,
        database="plenya_db",
        user="plenya_user",
        password="plenya_password"
    )

def get_item_info(cursor, item_id):
    """Buscar informações do item"""
    cursor.execute("""
        SELECT
            si.id,
            si.name,
            si.unit,
            si.clinical_relevance,
            si.patient_explanation,
            si.conduct,
            ss.name as subgroup,
            sg.name as grupo
        FROM score_items si
        JOIN score_subgroups ss ON si.subgroup_id = ss.id
        JOIN score_groups sg ON ss.group_id = sg.id
        WHERE si.id = %s
    """, (item_id,))
    return cursor.fetchone()

def generate_clinical_content(item_name, unit, subgroup):
    """Gerar conteúdo clínico usando Claude Opus 4.5"""

    prompt = f"""Você é um médico especialista em medicina de precisão e longevidade.

ITEM DE EXAME: {item_name}
UNIDADE: {unit or "N/A"}
CATEGORIA: {subgroup}

Gere conteúdo clínico rigoroso em 3 campos (formato JSON):

1. **clinical_relevance** (150-250 palavras):
   - Fisiopatologia e interpretação clínica
   - Valores de referência e variações
   - Associações com condições clínicas
   - Evidências científicas (últimos 5 anos)
   - Tom: técnico para médicos

2. **patient_explanation** (100-150 palavras):
   - Linguagem acessível (8ª série)
   - O que o exame avalia e por que importa
   - Fatores que afetam os resultados
   - Impacto na saúde e longevidade
   - Tom: educativo e empático

3. **conduct** (100-150 palavras):
   - Condutas baseadas em níveis (baixo/normal/alto)
   - Investigações adicionais se alterado
   - Intervenções farmacológicas/não-farmacológicas
   - Monitoramento e periodicidade
   - Critérios de encaminhamento

IMPORTANTE:
- Use evidências de 2020-2025
- Cite diretrizes quando relevante (AHA, ESC, SBC)
- Para scores de colonoscopia (Mayo UC, SES-CD, adenomas): foque em classificação, prognóstico e conduta
- Para CoQ10: considere suplementação, estatinas, mitocôndrias
- Para cistatina C: foque em função renal, superioridade vs creatinina
- Para colesterol: use guidelines 2023/2024, risco cardiovascular

FORMATO DE RESPOSTA (JSON válido):
{{
  "clinical_relevance": "...",
  "patient_explanation": "...",
  "conduct": "..."
}}"""

    message = client.messages.create(
        model="claude-opus-4-5-20251101",
        max_tokens=2000,
        temperature=0.3,
        messages=[{
            "role": "user",
            "content": prompt
        }]
    )

    response_text = message.content[0].text.strip()

    # Extrair JSON
    if "```json" in response_text:
        response_text = response_text.split("```json")[1].split("```")[0].strip()
    elif "```" in response_text:
        response_text = response_text.split("```")[1].split("```")[0].strip()

    return json.loads(response_text)

def update_item(cursor, item_id, content):
    """Atualizar item no banco"""
    cursor.execute("""
        UPDATE score_items
        SET
            clinical_relevance = %s,
            patient_explanation = %s,
            conduct = %s,
            last_review = %s,
            updated_at = CURRENT_TIMESTAMP
        WHERE id = %s
    """, (
        content["clinical_relevance"],
        content["patient_explanation"],
        content["conduct"],
        datetime.now(),
        item_id
    ))

def main():
    """Processar todos os 20 items"""
    conn = connect_db()
    cursor = conn.cursor()

    results = {
        "total": len(ITEM_IDS),
        "processed": 0,
        "errors": [],
        "items": []
    }

    for idx, item_id in enumerate(ITEM_IDS, 1):
        try:
            print(f"\n[{idx}/{len(ITEM_IDS)}] Processando item {item_id}...")

            # Buscar info
            item_info = get_item_info(cursor, item_id)
            if not item_info:
                print(f"  ❌ Item não encontrado: {item_id}")
                results["errors"].append({
                    "item_id": item_id,
                    "error": "Item não encontrado"
                })
                continue

            item_name = item_info[1]
            unit = item_info[2]
            subgroup = item_info[6]

            print(f"  📋 Nome: {item_name}")
            print(f"  📊 Categoria: {subgroup}")

            # Gerar conteúdo
            print(f"  🤖 Gerando conteúdo com Claude Opus 4.5...")
            content = generate_clinical_content(item_name, unit, subgroup)

            # Atualizar
            update_item(cursor, item_id, content)
            conn.commit()

            results["processed"] += 1
            results["items"].append({
                "item_id": item_id,
                "name": item_name,
                "subgroup": subgroup,
                "status": "success"
            })

            print(f"  ✅ Item atualizado com sucesso!")
            print(f"  📝 Clinical: {len(content['clinical_relevance'])} chars")
            print(f"  👤 Patient: {len(content['patient_explanation'])} chars")
            print(f"  🔬 Conduct: {len(content['conduct'])} chars")

        except Exception as e:
            print(f"  ❌ Erro: {str(e)}")
            results["errors"].append({
                "item_id": item_id,
                "error": str(e)
            })
            conn.rollback()
            continue

    cursor.close()
    conn.close()

    # Relatório final
    print("\n" + "="*60)
    print("RELATÓRIO FINAL - BATCH PARTE 6")
    print("="*60)
    print(f"Total de items: {results['total']}")
    print(f"Processados com sucesso: {results['processed']}")
    print(f"Erros: {len(results['errors'])}")

    if results["errors"]:
        print("\nERROS:")
        for error in results["errors"]:
            print(f"  - {error['item_id']}: {error['error']}")

    # Salvar relatório
    report_path = "/home/user/plenya/BATCH-PARTE6-LAB-ITEMS-REPORT.json"
    with open(report_path, "w", encoding="utf-8") as f:
        json.dump(results, f, indent=2, ensure_ascii=False)

    print(f"\n📄 Relatório salvo em: {report_path}")

    return results

if __name__ == "__main__":
    main()
