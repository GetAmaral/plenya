#!/usr/bin/env python3
"""
Transform batch SQL files to use correct schema columns.
Maps various column names to: clinical_relevance, patient_explanation, conduct
"""

import re
import sys

def transform_batch_1(input_file, output_file):
    """
    Batch 1 uses: interpretation, low/medium/high_level_description, low/medium/high_level_recommendation, articles
    Map to: clinical_relevance, conduct
    """
    with open(input_file, 'r', encoding='utf-8') as f:
        content = f.read()

    # Pattern to match UPDATE statements in Batch 1
    pattern = r"UPDATE score_items SET\s+(.*?)\s+WHERE id = '([a-f0-9-]+)';"

    def replace_update(match):
        fields_block = match.group(1)
        item_id = match.group(2)

        # Extract fields
        interpretation = re.search(r"interpretation = '(.*?)',", fields_block, re.DOTALL)
        low_desc = re.search(r"low_level_description = '(.*?)',", fields_block, re.DOTALL)
        med_desc = re.search(r"medium_level_description = '(.*?)',", fields_block, re.DOTALL)
        high_desc = re.search(r"high_level_description = '(.*?)',", fields_block, re.DOTALL)
        low_rec = re.search(r"low_level_recommendation = '(.*?)',", fields_block, re.DOTALL)
        med_rec = re.search(r"medium_level_recommendation = '(.*?)',", fields_block, re.DOTALL)
        high_rec = re.search(r"high_level_recommendation = '(.*?)',", fields_block, re.DOTALL)

        # Combine clinical_relevance
        clinical_parts = []
        if interpretation:
            clinical_parts.append(interpretation.group(1))
        if low_desc:
            clinical_parts.append(f"\\n\\n**Nível Ótimo (Low):**\\n{low_desc.group(1)}")
        if med_desc:
            clinical_parts.append(f"\\n\\n**Nível Intermediário (Medium):**\\n{med_desc.group(1)}")
        if high_desc:
            clinical_parts.append(f"\\n\\n**Nível Alto (High):**\\n{high_desc.group(1)}")

        clinical_relevance = ''.join(clinical_parts)

        # Combine conduct
        conduct_parts = []
        if low_rec:
            conduct_parts.append(f"**Conduta - Nível Ótimo:**\\n{low_rec.group(1)}")
        if med_rec:
            conduct_parts.append(f"\\n\\n**Conduta - Nível Intermediário:**\\n{med_rec.group(1)}")
        if high_rec:
            conduct_parts.append(f"\\n\\n**Conduta - Nível Alto:**\\n{high_rec.group(1)}")

        conduct = ''.join(conduct_parts)

        # Create new UPDATE statement
        return f"""UPDATE score_items SET
  clinical_relevance = '{clinical_relevance}',
  conduct = '{conduct}',
  updated_at = NOW()
WHERE id = '{item_id}';"""

    # Replace all UPDATE statements
    transformed = re.sub(pattern, replace_update, content, flags=re.DOTALL)

    with open(output_file, 'w', encoding='utf-8') as f:
        f.write(transformed)

    print(f"✅ Transformed Batch 1: {output_file}")

def main():
    print("🔧 Transforming batch SQL files to match schema...")
    print()

    # Transform Batch 1
    try:
        transform_batch_1('batch_final_1_exames_A.sql', 'batch_final_1_exames_A_FIXED.sql')
    except Exception as e:
        print(f"❌ Error transforming Batch 1: {e}")
        sys.exit(1)

    print()
    print("✅ All transformations complete!")
    print("📋 Next step: Review the _FIXED.sql files and execute them")

if __name__ == '__main__':
    main()
