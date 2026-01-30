#!/usr/bin/env python3
"""
Batch enrichment for ALIMENTAÇÃO items - Part 2
20 items: Qualidade alimentação parentes, Infância, Intolerâncias, etc.
"""

import anthropic
import os
import json
import time
from datetime import datetime

# API Configuration
# Try environment variable first, then check for key file
API_KEY = os.getenv("ANTHROPIC_API_KEY")

if not API_KEY:
    # Try to read from config file
    key_file = os.path.expanduser("~/.anthropic_key")
    if os.path.exists(key_file):
        with open(key_file, 'r') as f:
            API_KEY = f.read().strip()

if not API_KEY:
    print("\nERROR: ANTHROPIC_API_KEY not found.")
    print("Please set it as environment variable or create ~/.anthropic_key file")
    print("\nUsage:")
    print("  export ANTHROPIC_API_KEY='your-key-here'")
    print("  python3 batch_alimentacao_parte2.py")
    print("\nOr:")
    print("  echo 'your-key-here' > ~/.anthropic_key")
    print("  python3 batch_alimentacao_parte2.py")
    exit(1)

client = anthropic.Anthropic(api_key=API_KEY)

# Items to process
ITEMS = [
    {
        "id": "b781e715-372d-4f05-b01e-4db68c05d8db",
        "title": "Qualidade alimentação parentes",
        "current_question": "Como você descreveria a alimentação dos seus pais e familiares próximos?",
        "context": "Dietary quality of parents and close relatives - family nutritional patterns"
    },
    {
        "id": "2aa102b7-9d48-4870-a1f1-53f53b29a3e6",
        "title": "Infância",
        "current_question": "Como era sua alimentação durante a infância?",
        "context": "Childhood nutrition patterns and dietary habits formation"
    },
    {
        "id": "fb918f6e-32c6-494f-b7a9-cc6ee738ce32",
        "title": "Intolerâncias",
        "current_question": "Você tem alguma intolerância alimentar conhecida?",
        "context": "Food intolerances (lactose, fructose, histamine, etc.)"
    },
    {
        "id": "39d2c6c6-7b8b-41a8-a90e-05b396ad61dc",
        "title": "Introdução alimentar",
        "current_question": "Como foi sua introdução alimentar quando bebê?",
        "context": "Weaning and complementary feeding introduction during infancy"
    },
    {
        "id": "7ba24074-1ffa-444d-82d7-0decb8db039d",
        "title": "Lactose",
        "current_question": "Você tem intolerância à lactose?",
        "context": "Lactose intolerance - primary or secondary"
    },
    {
        "id": "ce6181cc-827b-4ad8-89ce-a555884c71c6",
        "title": "Alimentação parental pré/durante gestação",
        "current_question": "Se você conhece, como era a alimentação dos seus pais antes e durante sua gestação?",
        "context": "Parental nutrition before and during pregnancy - epigenetic programming"
    },
    {
        "id": "d8696390-b619-4752-b366-5c63f4730811",
        "title": "Recordatório",
        "current_question": "O que você comeu nas últimas 24 horas?",
        "context": "24-hour dietary recall - quantitative food assessment"
    },
    {
        "id": "71cc4a7b-5ebd-437a-bbee-6f85105ac8dc",
        "title": "Onde e como come",
        "current_question": "Onde e como você faz suas refeições habitualmente?",
        "context": "Eating environment and circumstances - mindful eating assessment"
    },
    {
        "id": "5dc99862-e9c2-4ddb-b0b3-0631a32e3dc4",
        "title": "Ordem alimentos",
        "current_question": "Em que ordem você consome os alimentos durante as refeições?",
        "context": "Food sequencing and meal composition order - glycemic impact"
    },
    {
        "id": "3dcaa794-d1f1-4551-ad62-8bae902abaf8",
        "title": "Outros",
        "current_question": "Há algum outro aspecto importante sobre sua alimentação que não foi mencionado?",
        "context": "Additional dietary information and observations"
    },
    {
        "id": "f7a0aa7a-cfea-4c61-8391-52d82b236429",
        "title": "Perfil metabólico parental",
        "current_question": "Seus pais ou familiares próximos têm diabetes, obesidade ou síndrome metabólica?",
        "context": "Parental metabolic profile - genetic and familial risk factors"
    },
    {
        "id": "5cac2737-c0c5-47d9-aaa8-2ca4fad39f47",
        "title": "Piores períodos",
        "current_question": "Quais foram os piores períodos da sua alimentação e por quê?",
        "context": "Worst dietary periods - crisis points and nutritional history"
    },
    {
        "id": "c3b9b8b0-2e94-40f8-b3d6-143a7d8224f1",
        "title": "Preferências",
        "current_question": "Quais são suas preferências alimentares e alimentos favoritos?",
        "context": "Food preferences and favorite foods - palatability and adherence"
    },
    {
        "id": "a7fa71ab-72c6-430d-b6f6-ca6c95611197",
        "title": "Proteína do leite",
        "current_question": "Você tem alergia ou sensibilidade à proteína do leite?",
        "context": "Milk protein allergy/sensitivity (casein, whey) - IgE and non-IgE mediated"
    },
    {
        "id": "d8de21dc-24bd-489c-93ed-aa094428f292",
        "title": "Quanto come",
        "current_question": "Como você descreveria as porções que você come?",
        "context": "Portion sizes and serving amounts - quantitative intake"
    },
    {
        "id": "a35a3012-60ea-4659-a742-fbbcb741c6db",
        "title": "Quem cozinha",
        "current_question": "Quem prepara suas refeições habitualmente?",
        "context": "Meal preparation responsibility and cooking skills"
    },
    {
        "id": "27582ef6-b5e1-4d6f-9a7a-368c27fdf42a",
        "title": "Regras alimentares",
        "current_question": "Você segue alguma regra ou padrão alimentar específico?",
        "context": "Dietary rules, patterns and eating behaviors"
    },
    {
        "id": "6998408b-341d-487c-9c99-7e8bc72de3fb",
        "title": "Relação com comida",
        "current_question": "Como é sua relação emocional com a comida?",
        "context": "Emotional relationship with food - eating psychology"
    },
    {
        "id": "8f24a93b-8cd7-4157-adf4-3c9a7985e368",
        "title": "Restrições pessoais",
        "current_question": "Você evita algum alimento por escolha pessoal (não por intolerância)?",
        "context": "Personal food restrictions and dietary choices"
    },
    {
        "id": "01498f3d-551a-4285-bc84-86d015569d31",
        "title": "Satisfação pós-refeição",
        "current_question": "Como você se sente após as refeições?",
        "context": "Post-prandial satisfaction and digestive comfort"
    }
]

def generate_enrichment(item):
    """Generate clinical enrichment using Claude API"""

    prompt = f"""You are a clinical nutritionist creating evidence-based content for an EMR system.

ITEM: {item['title']}
CURRENT QUESTION: {item['current_question']}
CONTEXT: {item['context']}

Generate a JSON object with these fields (ALL IN PORTUGUESE):

1. **question** (string): Improved clinical question - clear, specific, patient-friendly
   - Should be comprehensive but easy to understand
   - Include timeframes when relevant (últimas semanas, atualmente, etc.)

2. **clinical_relevance** (string): Why this matters clinically (2-3 sentences)
   - Specific pathophysiological mechanisms
   - Impact on metabolic health, longevity, disease prevention

3. **interpretation_guide** (string): How to interpret responses (3-4 sentences)
   - What constitutes optimal, suboptimal, and concerning patterns
   - Key red flags or positive indicators
   - Contextual factors to consider

4. **health_implications** (array of strings): 4-6 specific health implications
   - Each should be concrete and evidence-based
   - Cover immediate and long-term effects
   - Include metabolic, inflammatory, and systemic impacts

5. **followup_questions** (array of strings): 3-5 contextual follow-up questions
   - Should deepen understanding based on initial response
   - Help quantify or qualify the information
   - Explore related factors

6. **red_flags** (array of strings): 3-4 concerning patterns requiring attention
   - Clear clinical warning signs
   - Patterns suggesting intervention needed

7. **recommendations** (array of strings): 4-6 evidence-based recommendations
   - Actionable, specific guidance
   - Applicable to different scenarios/patient types
   - Based on current nutrition science

8. **scientific_background** (string): Brief scientific context (2-3 sentences)
   - Key research findings
   - Relevant mechanisms
   - Current scientific consensus

Output ONLY valid JSON. No markdown, no extra text.
"""

    try:
        message = client.messages.create(
            model="claude-opus-4-5-20250929",
            max_tokens=16000,
            temperature=1,
            thinking={
                "type": "enabled",
                "budget_tokens": 10000
            },
            messages=[{
                "role": "user",
                "content": prompt
            }]
        )

        # Extract JSON from response
        response_text = ""
        for block in message.content:
            if block.type == "text":
                response_text = block.text
                break

        # Parse JSON
        enrichment = json.loads(response_text)

        return {
            "success": True,
            "data": enrichment,
            "thinking_tokens": sum(b.thinking.get("usage", {}).get("tokens", 0)
                                 for b in message.content if hasattr(b, "thinking")),
            "output_tokens": message.usage.output_tokens
        }

    except Exception as e:
        return {
            "success": False,
            "error": str(e)
        }

def main():
    """Process all items"""

    print(f"\n{'='*80}")
    print(f"BATCH ALIMENTAÇÃO PARTE 2 - 20 ITEMS")
    print(f"{'='*80}\n")
    print(f"Started at: {datetime.now().strftime('%Y-%m-%d %H:%M:%S')}\n")

    results = []
    total_thinking = 0
    total_output = 0

    for idx, item in enumerate(ITEMS, 1):
        print(f"\n[{idx}/20] Processing: {item['title']}")
        print(f"ID: {item['id']}")

        result = generate_enrichment(item)

        if result["success"]:
            print(f"✓ Success")
            print(f"  Thinking tokens: {result['thinking_tokens']:,}")
            print(f"  Output tokens: {result['output_tokens']:,}")

            total_thinking += result['thinking_tokens']
            total_output += result['output_tokens']

            results.append({
                "id": item["id"],
                "title": item["title"],
                "enrichment": result["data"],
                "tokens": {
                    "thinking": result['thinking_tokens'],
                    "output": result['output_tokens']
                }
            })
        else:
            print(f"✗ Error: {result['error']}")
            results.append({
                "id": item["id"],
                "title": item["title"],
                "error": result["error"]
            })

        # Rate limiting
        if idx < len(ITEMS):
            time.sleep(2)

    # Save results
    output_file = "/home/user/plenya/batch_alimentacao_parte2_results.json"
    with open(output_file, 'w', encoding='utf-8') as f:
        json.dump(results, f, ensure_ascii=False, indent=2)

    # Generate SQL
    sql_statements = []
    sql_statements.append("-- BATCH ALIMENTAÇÃO PARTE 2 - 20 ITEMS")
    sql_statements.append(f"-- Generated: {datetime.now().strftime('%Y-%m-%d %H:%M:%S')}\n")

    for r in results:
        if "enrichment" in r:
            e = r["enrichment"]
            sql_statements.append(f"-- {r['title']}")
            sql_statements.append(f"UPDATE score_items SET")
            sql_statements.append(f"  question = '{e['question'].replace(chr(39), chr(39)+chr(39))}',")
            sql_statements.append(f"  clinical_relevance = '{e['clinical_relevance'].replace(chr(39), chr(39)+chr(39))}',")
            sql_statements.append(f"  interpretation_guide = '{e['interpretation_guide'].replace(chr(39), chr(39)+chr(39))}',")
            sql_statements.append(f"  health_implications = ARRAY[{', '.join([chr(39) + hi.replace(chr(39), chr(39)+chr(39)) + chr(39) for hi in e['health_implications']])}]::text[],")
            sql_statements.append(f"  followup_questions = ARRAY[{', '.join([chr(39) + fq.replace(chr(39), chr(39)+chr(39)) + chr(39) for fq in e['followup_questions']])}]::text[],")
            sql_statements.append(f"  red_flags = ARRAY[{', '.join([chr(39) + rf.replace(chr(39), chr(39)+chr(39)) + chr(39) for rf in e['red_flags']])}]::text[],")
            sql_statements.append(f"  recommendations = ARRAY[{', '.join([chr(39) + rec.replace(chr(39), chr(39)+chr(39)) + chr(39) for rec in e['recommendations']])}]::text[],")
            sql_statements.append(f"  scientific_background = '{e['scientific_background'].replace(chr(39), chr(39)+chr(39))}'")
            sql_statements.append(f"WHERE id = '{r['id']}';\n")

    sql_file = "/home/user/plenya/batch_alimentacao_parte2.sql"
    with open(sql_file, 'w', encoding='utf-8') as f:
        f.write('\n'.join(sql_statements))

    # Summary
    print(f"\n{'='*80}")
    print(f"EXECUTION SUMMARY")
    print(f"{'='*80}\n")
    print(f"Total items processed: {len(results)}")
    print(f"Successful: {sum(1 for r in results if 'enrichment' in r)}")
    print(f"Failed: {sum(1 for r in results if 'error' in r)}")
    print(f"\nToken Usage:")
    print(f"  Thinking tokens: {total_thinking:,}")
    print(f"  Output tokens: {total_output:,}")
    print(f"  Total tokens: {(total_thinking + total_output):,}")
    print(f"\nFiles generated:")
    print(f"  JSON: {output_file}")
    print(f"  SQL: {sql_file}")
    print(f"\nCompleted at: {datetime.now().strftime('%Y-%m-%d %H:%M:%S')}")

if __name__ == "__main__":
    main()
