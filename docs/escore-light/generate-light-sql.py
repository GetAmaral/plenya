#!/usr/bin/env python3
"""Generate SQL UPDATE statements to mark Light items in score_items.

Reads the v5 curation list, queries DB for each item by name pattern,
and outputs:
  - UPDATE statements with is_light_version=true, light_order, light_question
  - A report of items found / not found / multiple matches (to review)
"""
import subprocess
import json
import re
from pathlib import Path

# Item list: (light_order, name_patterns, light_question)
# Patterns are SQL ILIKE matches. Multiple patterns mean multiple records (M/F pairs etc).
ITEMS = [
    # ============ A — ATIVIDADE FÍSICA ============
    (1, ["Estratégia macro atual"],
     "Como você descreveria seu nível atual de atividade física?"),
    (2, ["Lesões relacionadas ao exercício"],
     "Você tem ou já teve lesões relacionadas ao exercício?"),
    (3, ["Resistência Neuromuscular - Prancha (Homens)", "Resistência Neuromuscular - Prancha (Mulheres)"],
     "Por quantos segundos você consegue sustentar uma prancha frontal? (com cronômetro, em forma correta)"),

    # ============ A — ALIMENTAÇÃO ============
    (4, ["Qualidade da alimentação na vida adulta"],
     "Como você descreveria a qualidade da sua alimentação nos últimos anos?"),
    (5, ["Consumo de Verduras e Legumes"],
     "Quantas porções de verduras e legumes você consome por dia? (1 porção ≈ 1 prato pequeno ou 1 xícara)"),
    (6, ["Refrigerantes e energéticos"],
     "Com que frequência você consome refrigerantes ou bebidas energéticas?"),
    (7, ["Onde e como come"],
     "Em quantas das suas refeições diárias você se senta à mesa, com calma e sem distratores (TV, celular)?"),
    (8, ["Divisão e Horários das refeições"],
     "Como são a divisão e os horários das suas refeições?"),
    (9, ["Todos seguem o mesmo padrão alimentar"],
     "As pessoas que moram com você seguem o mesmo padrão alimentar que o seu?"),

    # ============ A — COMPOSIÇÃO CORPORAL ============
    (10, ["Abdominal (cintura em cm) - homem", "Abdominal (cintura em cm) - mulher"],
     "Qual é a sua circunferência abdominal? Meça na altura do umbigo, sem apertar a fita."),
    (11, ["Pescoço - homem", "Pescoço - mulher"],
     "Qual é a circunferência do seu pescoço? Meça logo abaixo do pomo de adão."),
    (12, ["IMC (kg/m²)"],
     "IMC calculado automaticamente a partir do seu peso e altura."),
    (13, ["Razão cintura/altura"],
     "Razão cintura/altura calculada automaticamente."),
    (14, ["ASMI (kg/m²) - homem", "ASMI (kg/m²) - mulher"],
     "Se você fez bioimpedância ou DEXA: qual o seu ASMI (Índice de Massa Apendicular)?"),
    (15, ["Gordura visceral (cm²) - homem", "Gordura visceral (cm²) - mulher"],
     "Se você fez bioimpedância ou DEXA: qual seu valor de gordura visceral em cm²?"),
    (16, ["% Gordura corporal - homem", "% Gordura corporal - mulher"],
     "Se você fez bioimpedância ou DEXA: qual seu percentual de gordura corporal?"),

    # ============ G — HISTÓRICO DE DOENÇAS ============
    (17, ["Doença cardiovascular (IAM, revascularização, AVC, etc"],
     "Você tem ou já teve doença cardiovascular (infarto, angioplastia, AVC)?"),
    (18, ["^Insuficiência cardíaca$"],
     "Você tem diagnóstico de insuficiência cardíaca?"),
    (19, ["^Diabetes mellitus$"],
     "Você tem diagnóstico de diabetes mellitus?"),
    (20, ["^Pré-diabetes / resistência a insulina$"],
     "Você tem diagnóstico de pré-diabetes ou resistência insulínica?"),
    (21, ["^Arritmia$"],
     "Você tem diagnóstico de arritmia cardíaca?"),
    (22, ["^Asma$"],
     "Você tem diagnóstico de asma?"),

    # ============ G — HISTÓRICO FAMILIAR ============
    (23, ["Doença cardiovascular (IAM, revascularização, arritmias, AVC)"],
     "Algum parente próximo (pais, irmãos, avós) tem ou teve doença cardiovascular?"),
    (24, ["^Câncer$"],  # HF Câncer — exact match
     "Algum parente próximo tem ou teve câncer?"),
    (25, ["Diabetes mellitus / pré-diabetes / resistência insulínica"],
     "Algum parente próximo tem diabetes, pré-diabetes ou resistência insulínica?"),
    (26, ["^Dislipidemia$"],
     "Algum parente próximo tem dislipidemia (colesterol/triglicérides alterados)?"),
    (27, ["^Hipertensão arterial$"],
     "Algum parente próximo tem hipertensão arterial?"),
    (28, ["^Obesidade$"],
     "Algum parente próximo tem obesidade?"),

    # ============ G — MEDICAMENTOS ============
    (29, ["Anti-hipertensivos (IECA"],
     "Você usa medicamentos para pressão alta atualmente?"),
    (30, ["Estatinas / Hipolipemiantes"],
     "Você usa estatinas ou outros medicamentos para colesterol?"),
    (31, ["Análogos de GLP-1"],
     "Você usa medicamentos análogos de GLP-1 (Ozempic, Mounjaro, Wegovy, Saxenda)?"),
    (32, ["^Insulinas$"],
     "Você usa insulina?"),

    # ============ G — SINTOMAS SENTINELAS ============
    (33, ["Dor torácica"],
     "Você sente ou já sentiu dor no peito (dor torácica)?"),
    (34, ["Dispnéia"],
     "Você sente falta de ar (dispneia) em situações cotidianas?"),
    (35, ["Palpitação"],
     "Você sente palpitações cardíacas?"),
    (36, ["Claudicação"],
     "Você sente dor nas pernas ao caminhar que melhora ao parar (claudicação)?"),
    (37, ["^Edema$"],
     "Você apresenta inchaço (edema) nos membros?"),
    (38, ["^Azia$"],
     "Você sente azia ou queimação no estômago?"),
    (39, ["^Obstipação$"],
     "Você tem dificuldade para evacuar (intestino preso/obstipação)?"),

    # ============ G — SAÚDE BUCAL ============
    (40, ["Histórico Periodontal"],
     "Como é seu histórico de doença periodontal (gengivite, periodontite)?"),
    (41, ["Higiene e Acompanhamento"],
     "Como é sua rotina de higiene bucal e acompanhamento odontológico?"),

    # ============ G — LAB CARDIOVASCULAR ============
    (42, ["^Apolipoproteína B$"],
     "Se você tem o resultado: qual seu valor de Apolipoproteína B (ApoB) em mg/dL?"),
    (43, ["^Lipoproteína A$"],
     "Se você tem o resultado: qual seu valor de Lipoproteína(a) em nmol/L?"),
    (44, ["PCR ultrassensível"],
     "Se você tem o resultado: qual seu valor de PCR ultrassensível (hs-CRP) em mg/L?"),
    (45, ["^LDL Colesterol$"],
     "Se você tem o resultado: qual seu valor de LDL colesterol em mg/dL?"),
    (46, ["HDL Colesterol"],
     "Se você tem o resultado: qual seu valor de HDL colesterol em mg/dL?"),
    (47, ["^Triglicerídeos$"],
     "Se você tem o resultado: qual seu valor de triglicerídeos em mg/dL?"),
    (48, ["Relação Triglicerídeos/HDL"],
     "Razão TG/HDL calculada automaticamente."),
    (49, ["NT-proBNP"],
     "Se você tem o resultado: qual seu valor de NT-proBNP em pg/mL?"),

    # ============ G — LAB METABÓLICO ============
    (50, ["Hemoglobina glicada"],
     "Se você tem o resultado: qual sua hemoglobina glicada (HbA1c) em %?"),
    (51, ["^HOMA-IR$"],
     "Se você tem o resultado: qual seu HOMA-IR?"),

    # ============ G — LAB NUTRICIONAL ============
    (52, ["25-hidroxivitamina D"],
     "Se você tem o resultado: qual seu valor de Vitamina D (25-OH) em ng/mL?"),
    (53, ["^Homocisteína$"],
     "Se você tem o resultado: qual seu valor de homocisteína em µmol/L?"),
    (54, ["Ferritina - Homens", "Ferritina - Mulheres Pré-Menopausa", "Ferritina - Mulheres Pós-Menopausa"],
     "Se você tem o resultado: qual seu valor de ferritina em ng/mL?"),

    # ============ G — LAB HORMONAL ============
    (55, ["^TSH$"],
     "Se você tem o resultado: qual seu valor de TSH em mIU/L?"),
    (56, ["^T3 Livre$"],
     "Se você tem o resultado: qual seu valor de T3 Livre em pg/mL?"),
    (57, ["Testosterona Total - Homens", "Testosterona Total - Mulheres Pré-Menopausa", "Testosterona Total - Mulheres Pós-Menopausa"],
     "Se você tem o resultado: qual seu valor de testosterona total em ng/dL?"),
    (58, ["Estradiol - Homens",
          "Estradiol - Mulheres Fase Folicular Inicial",
          "Estradiol - Mulheres Fase Ovulatória",
          "Estradiol - Mulheres Fase Lútea",
          "Estradiol - Mulheres Pós-Menopausa (Sem TRH)",
          "Estradiol - Mulheres Pós-Menopausa (Com TRH)"],
     "Se você tem o resultado: qual seu valor de estradiol em pg/mL?"),

    # ============ G — RENAL ============
    (59, ["Microalbuminúria"],
     "Se você tem o resultado: qual seu valor de microalbuminúria/creatininúria em mg/g?"),

    # ============ G — IMAGEM ============
    (60, ["TC coração para escore de cálcio coronariano"],
     "Se você fez: qual seu escore de cálcio coronariano (Agatston Units)?"),
    (61, ["USG Abdome - Esteatose Hepática"],
     "Se você fez ultrassom abdominal: qual o grau de esteatose hepática?"),
    (62, ["Densitometria - T-Score Colo Femoral"],
     "Se você fez densitometria óssea: qual seu T-score do colo femoral?"),

    # ============ I — MENTE-CORPO/SOCIAL/VÍCIOS ============
    (63, ["Escala PHQ-9"],
     "Escala PHQ-9 (humor): some os 9 itens que vão aparecer e informe seu total /27."),
    (64, ["Capacidade da memória percebida"],
     "Como está a sua memória hoje, comparada ao seu auge?"),
    (65, ["Socialização atual"],
     "Como está sua vida social nos últimos 6 meses?"),
    (66, ["Fontes de stress percebidas"],
     "Como você descreveria seu nível atual de estresse e capacidade de manejá-lo?"),
    (67, ["^Tabaco$"],
     "Qual seu padrão de consumo de tabaco?"),
    (68, ["^Álcool$"],
     "Qual seu padrão de consumo de álcool?"),
    (69, ["Situação familiar"],
     "Como está sua situação familiar e rede de suporte?"),

    # ============ R — RITMO CIRCADIANO ============
    (70, ["Tempo de sono"],
     "Quantas horas, em média, você dorme por noite?"),
    (71, ["Hora de dormir"],
     "A que horas você costuma deitar para dormir?"),
    (72, ["Qualidade percebida do sono"],
     "Como você percebe a qualidade do seu sono?"),
    (73, ["Regularidade no acordar"],
     "Em quantos dias do mês você acorda no mesmo horário?"),
    (74, ["Uso de BiPAP/CPAP"],
     "Você usa CPAP ou BiPAP para dormir?"),
    (75, ["^Apneias$"],
     "Você apresenta pausas respiratórias durante o sono (apneias)?"),
    (76, ["^Roncos$"],
     "Você ronca durante o sono?"),
    (77, ["Tempo tela noturna"],
     "Com que frequência você usa telas (celular, TV, tablet) na hora antes de dormir?"),
]


def query_items_by_pattern(pattern: str):
    """Return list of (id, name) matching pattern (case-insensitive)."""
    # Use POSIX regex for ^anchors etc; otherwise ILIKE
    # Escape single quotes by doubling
    pat_escaped = pattern.replace("'", "''")
    # Filter for items with pts > 0 to eliminate header/duplicate items
    pt_filter = "AND points > 0"
    if pattern.startswith("^") or pattern.endswith("$"):
        sql = f"SELECT id::text, name FROM score_items WHERE deleted_at IS NULL {pt_filter} AND name ~* '{pat_escaped}' ORDER BY name;"
    else:
        like_pat = "%" + pat_escaped + "%"
        sql = f"SELECT id::text, name FROM score_items WHERE deleted_at IS NULL {pt_filter} AND name ILIKE '{like_pat}' ORDER BY name;"
    result = subprocess.run(
        ["docker", "compose", "exec", "-T", "db", "psql", "-U", "plenya_user", "-d", "plenya_db", "-A", "-F", "|", "-t", "-c", sql],
        capture_output=True, text=True
    )
    rows = []
    for line in result.stdout.strip().split("\n"):
        if "|" in line:
            iid, name = line.split("|", 1)
            rows.append((iid.strip(), name.strip()))
    return rows


def main():
    sql_lines = []
    report = []
    not_found = []
    multi_matches = []
    found_ids = []

    for order, patterns, question in ITEMS:
        for p_idx, pattern in enumerate(patterns):
            matches = query_items_by_pattern(pattern)
            if not matches:
                not_found.append((order, pattern))
                report.append(f"❌ #{order} pattern '{pattern}' — NOT FOUND")
                continue
            if len(matches) > 1:
                multi_matches.append((order, pattern, matches))
                report.append(f"ℹ️  #{order} pattern '{pattern}' — {len(matches)} matches (all included):")
                for m in matches:
                    report.append(f"     {m[0]}  {m[1]}")
            else:
                report.append(f"✓ #{order} '{matches[0][1]}'  →  {matches[0][0]}")
            # Include ALL matches (pts > 0 filter ensures only valid items)
            for iid, name in matches:
                found_ids.append((iid, order, name, question.replace("'", "''")))

    # Build SQL
    sql_lines.append("-- Escore Plenya Light — SQL gerado a partir da curadoria v5")
    sql_lines.append("-- Marca items com is_light_version=true, light_order, light_question")
    sql_lines.append("BEGIN;")
    sql_lines.append("")
    sql_lines.append("-- 1. Resetar todos os items para garantir estado limpo")
    sql_lines.append("UPDATE score_items SET is_light_version = false, light_order = NULL, light_question = NULL WHERE is_light_version = true;")
    sql_lines.append("")
    sql_lines.append("-- 2. Marcar items selecionados")

    for iid, order, name, question in found_ids:
        sql_lines.append(
            f"UPDATE score_items SET is_light_version = true, light_order = {order}, "
            f"light_question = '{question}' WHERE id = '{iid}'; -- {name}"
        )

    sql_lines.append("")
    sql_lines.append("COMMIT;")
    sql_lines.append("")
    sql_lines.append(f"-- Total: {len(found_ids)} registros marcados")

    out_sql = Path("/home/user/plenya/docs/escore-light/light-curation-v5.sql")
    out_sql.write_text("\n".join(sql_lines), encoding="utf-8")

    out_report = Path("/home/user/plenya/docs/escore-light/light-curation-v5-report.txt")
    out_report.write_text("\n".join(report), encoding="utf-8")

    print(f"SQL: {out_sql} ({len(found_ids)} records)")
    print(f"Report: {out_report}")
    print(f"\nNot found: {len(not_found)}")
    print(f"Multi-match (took first): {len(multi_matches)}")

if __name__ == "__main__":
    main()
