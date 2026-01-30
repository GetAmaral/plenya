#!/usr/bin/env python3
"""
Script standalone para enriquecer 20 Score Items de ALIMENTAÇÃO
Instala suas próprias dependências se necessário
"""

import subprocess
import sys
import os

# Tentar importar dependências, instalar se necessário
def ensure_dependencies():
    """Garante que dependências estejam instaladas"""
    required = ["anthropic", "requests"]
    missing = []

    for package in required:
        try:
            __import__(package)
        except ImportError:
            missing.append(package)

    if missing:
        print(f"📦 Instalando dependências: {', '.join(missing)}")
        subprocess.check_call([
            sys.executable, "-m", "pip", "install", "--user", "--quiet"
        ] + missing)
        print("✅ Dependências instaladas\n")

ensure_dependencies()

import json
import requests
import anthropic

# Configurações
API_URL = "http://localhost:3001/api/v1"
EMAIL = "import@plenya.com"
PASSWORD = "Import@123456"
# API Key será passada via argumento ou variável de ambiente
ANTHROPIC_API_KEY = os.getenv("ANTHROPIC_API_KEY") or (sys.argv[1] if len(sys.argv) > 1 else None)

# IDs dos 20 itens a enriquecer
TARGET_ITEMS = [
    "dc29ca47-3fa1-4ada-8cb0-85e83bcd38eb",  # Gluten 1
    "503b91b1-50d5-41f8-bc2d-772822096033",  # Gluten 2
    "a4bd778b-63f1-4b5f-9435-08411b954ffe",  # Gluten 3
    "21128901-d9bb-422f-8a95-5d4f2202fa49",  # Histamina 1
    "e76cf533-e63a-4602-98ff-5dbfa54c6079",  # Histamina 2
    "2436e6d5-945f-4646-885d-8c2df6ed6da7",  # Histamina 3
    "2a12cebb-111e-492a-8388-55222be9c346",  # Histórico familiar alimentação 1
    "64f67026-3778-461f-b5eb-4674bbb46c58",  # Histórico familiar alimentação 2
    "207b3a03-01a4-4165-9b48-a1fd0c42ae35",  # Histórico familiar alimentação 3
    "2a2e420e-1a0c-44b4-bb86-67d46e33c572",  # Hábitos mãe 1
    "73050428-0624-4b9b-8ebe-f920606d2335",  # Hábitos mãe 2
    "f946af49-7962-4371-b56e-794fcfb1d505",  # Hábitos mãe 3
    "b8401d7b-edb1-4c06-8dbd-dfa3e8754e41",  # Hábitos pai 1
    "92df69af-cf17-43fc-8a18-251bc6b8ebdf",  # Hábitos pai 2
    "8f4bbe8c-41ee-45c0-8455-c758fcc22bb0",  # Hábitos pai 3
    "2cf7a504-b4ce-4b50-9471-30fe89b19758",  # Problemas alimentares 1
    "e75bec23-9bc7-4fdc-9911-c9c9a2f48411",  # Problemas alimentares 2
    "f267b2b2-8a63-4801-9029-166aabb83176",  # Problemas alimentares 3
    "1daa1598-27a4-4750-9330-dc9967c345c6",  # Qualidade alimentação parentes 1
    "b781e715-372d-4f05-b01e-4db68c05d8db",  # Qualidade alimentação parentes 2
]


def login():
    """Faz login e retorna o access token"""
    response = requests.post(
        f"{API_URL}/auth/login",
        json={"email": EMAIL, "password": PASSWORD},
        timeout=10
    )
    response.raise_for_status()
    return response.json()["accessToken"]


def get_score_item(token, item_id):
    """Busca um score item pelo ID"""
    headers = {"Authorization": f"Bearer {token}"}
    response = requests.get(
        f"{API_URL}/score-items/{item_id}",
        headers=headers,
        timeout=10
    )
    response.raise_for_status()
    return response.json()


def get_relevant_lectures(token, item_name):
    """Busca lectures MFI relevantes"""
    headers = {"Authorization": f"Bearer {token}"}

    response = requests.get(
        f"{API_URL}/articles?limit=100",
        headers=headers,
        timeout=15
    )
    response.raise_for_status()
    all_articles = response.json()

    # Filtrar relevantes
    relevant = []
    name_lower = item_name.lower()

    for article in all_articles:
        title = article.get("title", "").lower()

        # Priorizar lectures MFI
        if "mfi" in title:
            relevant.append(article)
        elif "programação metabólica" in title:
            relevant.append(article)
        elif "genética" in title or "epigenética" in title:
            relevant.append(article)
        elif "gluten" in title and "gluten" in name_lower:
            relevant.append(article)
        elif "histamin" in title and "histamin" in name_lower:
            relevant.append(article)
        elif "nutrição" in title or "alimentação" in title:
            # Incluir, mas com menor prioridade
            if len(relevant) < 15:
                relevant.append(article)

    return relevant


def enrich_with_claude(item, lectures):
    """Usa Claude Sonnet 4.5 para gerar conteúdo clínico"""

    client = anthropic.Anthropic(api_key=ANTHROPIC_API_KEY)

    # Preparar contexto das lectures
    lectures_context = ""
    for i, lecture in enumerate(lectures[:8], 1):  # Top 8 lectures
        title = lecture.get("title", "Sem título")
        abstract = lecture.get("abstract", "")
        full_content = lecture.get("full_content", "")

        # Preferir full_content, fallback para abstract
        content = full_content if full_content else abstract

        if content:
            # Limitar a 4000 chars por lecture
            content = content[:4000]
            lectures_context += f"\n### LECTURE {i}: {title}\n{content}\n\n"

    if not lectures_context:
        lectures_context = "Sem lectures disponíveis. Use conhecimento médico geral sobre medicina funcional integrativa."

    # Criar prompt contextualizado
    item_name = item.get('name', '')

    # Definir foco baseado no nome do item
    if "gluten" in item_name.lower():
        focus = """Foque em:
- Sensibilidade ao glúten não-celíaca (SGNC)
- Doença celíaca
- Permeabilidade intestinal e zonulina
- Reatividade cruzada e mimetismo molecular
- Impacto neurológico (ataxia, neuropatia)
- Inflamação sistêmica e autoimunidade"""

    elif "histamina" in item_name.lower():
        focus = """Foque em:
- Intolerância à histamina
- Déficit de DAO (diamino oxidase)
- Mastócitos e MCAS (síndrome de ativação de mastócitos)
- Alimentos ricos em histamina e liberadores de histamina
- Sintomas (cefaleia, urticária, sintomas GI, taquicardia)
- Protocolo de dieta de eliminação"""

    elif "familiar" in item_name.lower() or "mãe" in item_name.lower() or "pai" in item_name.lower() or "parentes" in item_name.lower():
        focus = """Foque em:
- Programação metabólica e epigenética
- Janelas críticas de desenvolvimento (gestação, primeira infância)
- Influências transgeracionais
- DOHaD (Developmental Origins of Health and Disease)
- Impacto de hábitos parentais (tabagismo, etilismo, alimentação)
- Modificações epigenéticas herdáveis
- Prevenção e intervenção precoce"""

    else:
        focus = """Foque em aspectos nutricionais e metabólicos relevantes para o item específico."""

    prompt = f"""Você é um médico especialista em Medicina Funcional Integrativa. Baseado nas lectures do MFI fornecidas, gere conteúdo clínico profundo e baseado em evidências para este Score Item.

**ITEM:** {item_name}

**CONTEXTO DAS LECTURES MFI:**
{lectures_context}

**DIRECIONAMENTO:**
{focus}

**TAREFA:**
Gere um objeto JSON com 3 campos:

1. **clinicalRelevance** (200-400 palavras):
   - Fisiopatologia detalhada na perspectiva funcional
   - Mecanismos bioquímicos e moleculares
   - Impacto sistêmico e conexões (ex: eixo intestino-cérebro)
   - Dados epidemiológicos quando relevante
   - CITE CONCEITOS DAS LECTURES MFI quando aplicável

2. **patientExplanation** (100-200 palavras):
   - Linguagem acessível, sem jargões
   - Explicar "o que é" e "por que importa"
   - Tom empático e educativo
   - Ajudar o paciente a entender a relevância para sua saúde

3. **conduct** (150-300 palavras):
   - Protocolo clínico estruturado e prático
   - Investigação: anamnese dirigida, exames específicos
   - Intervenções: dietéticas, suplementação, estilo de vida
   - Monitoramento e critérios de reavaliação
   - Sinais de alerta e quando referenciar

**FORMATO:**
Retorne APENAS um objeto JSON válido com as 3 chaves.
Não inclua markdown code blocks.
Não inclua texto explicativo extra.
Apenas o JSON puro."""

    try:
        message = client.messages.create(
            model="claude-sonnet-4-5-20250929",
            max_tokens=4096,
            temperature=0.7,
            messages=[{"role": "user", "content": prompt}]
        )

        response_text = message.content[0].text.strip()

        # Limpar possíveis markdown code blocks
        if response_text.startswith("```"):
            lines = response_text.split("\n")
            if lines[0].startswith("```json"):
                response_text = "\n".join(lines[1:-1])
            elif lines[0] == "```":
                response_text = "\n".join(lines[1:-1])

        # Parsear JSON
        enriched_data = json.loads(response_text)

        return enriched_data

    except json.JSONDecodeError as e:
        print(f"   ⚠️  Erro ao parsear JSON: {str(e)}")
        print(f"   Resposta recebida: {response_text[:500]}...")
        return None
    except Exception as e:
        print(f"   ⚠️  Erro ao chamar Claude: {str(e)}")
        return None


def update_score_item(token, item_id, data):
    """Atualiza um score item"""
    headers = {
        "Authorization": f"Bearer {token}",
        "Content-Type": "application/json"
    }

    response = requests.put(
        f"{API_URL}/score-items/{item_id}",
        headers=headers,
        json=data,
        timeout=15
    )
    response.raise_for_status()
    return response.json()


def main():
    print("=" * 80)
    print("ENRIQUECIMENTO DE 20 SCORE ITEMS DE ALIMENTAÇÃO")
    print("Usando lectures MFI + Claude Sonnet 4.5")
    print("=" * 80)
    print()

    if not ANTHROPIC_API_KEY:
        print("❌ ERRO: Variável ANTHROPIC_API_KEY não configurada")
        print("   Configure: export ANTHROPIC_API_KEY='sk-...'")
        return 1

    # Login
    print("🔐 Fazendo login...")
    try:
        token = login()
        print("✅ Login realizado\n")
    except Exception as e:
        print(f"❌ Erro no login: {str(e)}")
        return 1

    # Processar cada item
    results = []
    success_count = 0
    error_count = 0

    for idx, item_id in enumerate(TARGET_ITEMS, 1):
        print(f"\n{'=' * 80}")
        print(f"ITEM {idx}/{len(TARGET_ITEMS)}")
        print(f"{'=' * 80}")

        try:
            # 1. Buscar item
            print(f"📖 Buscando item...")
            item = get_score_item(token, item_id)
            item_name = item.get("name", "Sem nome")
            print(f"   Nome: {item_name}")
            print(f"   ID: {item_id}")

            # 2. Buscar lectures
            print(f"\n📚 Buscando lectures MFI...")
            lectures = get_relevant_lectures(token, item_name)
            print(f"   Encontradas: {len(lectures)} lectures")

            if lectures:
                print(f"   Top lectures:")
                for i, lec in enumerate(lectures[:5], 1):
                    print(f"      {i}. {lec.get('title', 'Sem título')}")

            # 3. Enriquecer com Claude
            print(f"\n🤖 Gerando conteúdo com Claude Sonnet 4.5...")
            enriched = enrich_with_claude(item, lectures)

            if not enriched:
                print(f"   ❌ Falha ao gerar conteúdo")
                error_count += 1
                results.append({
                    "item_id": item_id,
                    "name": item_name,
                    "status": "error",
                    "error": "Falha ao gerar conteúdo com Claude"
                })
                continue

            # Verificar tamanhos
            cr_len = len(enriched.get("clinicalRelevance", ""))
            pe_len = len(enriched.get("patientExplanation", ""))
            cd_len = len(enriched.get("conduct", ""))

            print(f"   ✅ Conteúdo gerado:")
            print(f"      - Clinical Relevance: {cr_len} caracteres")
            print(f"      - Patient Explanation: {pe_len} caracteres")
            print(f"      - Conduct: {cd_len} caracteres")

            # 4. Atualizar no banco
            print(f"\n💾 Salvando no banco de dados...")
            updated = update_score_item(token, item_id, enriched)
            print(f"   ✅ Item atualizado com sucesso!")

            success_count += 1
            results.append({
                "item_id": item_id,
                "name": item_name,
                "status": "success",
                "lectures_used": len(lectures),
                "content_sizes": {
                    "clinical_relevance": cr_len,
                    "patient_explanation": pe_len,
                    "conduct": cd_len
                }
            })

        except Exception as e:
            print(f"   ❌ ERRO: {str(e)}")
            error_count += 1
            results.append({
                "item_id": item_id,
                "name": item.get("name", "Desconhecido") if 'item' in locals() else "Desconhecido",
                "status": "error",
                "error": str(e)
            })

    # Relatório final
    print(f"\n\n{'=' * 80}")
    print("RELATÓRIO FINAL")
    print(f"{'=' * 80}\n")

    print(f"✅ Sucesso: {success_count}/{len(TARGET_ITEMS)}")
    print(f"❌ Erros: {error_count}/{len(TARGET_ITEMS)}")
    print(f"📊 Taxa de sucesso: {(success_count/len(TARGET_ITEMS)*100):.1f}%")
    print()

    if error_count > 0:
        print("ERROS DETALHADOS:")
        for r in results:
            if r["status"] == "error":
                print(f"  - {r['name']} ({r['item_id']})")
                print(f"    Erro: {r.get('error', 'Desconhecido')}")
        print()

    # Salvar relatório JSON
    report_path = "/home/user/plenya/DIETARY-BATCH-20-REPORT.json"
    with open(report_path, "w", encoding="utf-8") as f:
        json.dump(results, f, indent=2, ensure_ascii=False)

    print(f"📄 Relatório JSON salvo em: {report_path}")
    print()

    return 0 if error_count == 0 else 1


if __name__ == "__main__":
    sys.exit(main())
