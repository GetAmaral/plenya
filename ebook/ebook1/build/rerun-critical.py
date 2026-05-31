"""
Re-run the 11 critical figures with surgical anti-paraphrase prompts
that name the SPECIFIC errors detected in the QA pass and demand exact
strings.
"""

import subprocess
import time
from pathlib import Path
from concurrent.futures import ThreadPoolExecutor, as_completed

EBOOK = Path("/home/user/plenya/ebook/Ebook 1 - Performance e Longevidade")
SCRIPT = EBOOK / "build" / "translate-figure.sh"
YAML_PATH = EBOOK / "figuras" / "figure-substitutions.yaml"

# Per-figure surgical critique sections appended to the standard YAML prompt.
# Each entry highlights ERRORS detected in the previous attempt and demands
# the exact replacement.
CRITICAL_FIXES = {
    "Cap01 Fig02.PNG": """
=== CRITICAL FIXES (previous attempt got these wrong, this attempt MUST get them right) ===
1. The word 'optimal' is a key technical term in this book — DO NOT replace it with 'great', 'best', or 'ideal'. Use 'optimal' verbatim wherever YAML says 'optimal'.
2. The phrase 'Lab upper limit of normal' must appear EXACTLY — do NOT paraphrase to 'Laboratory normal range' or 'Laboratory normal limit'.
3. ALL decimal numbers must use period (.) not comma (,). Specifically: 2.4 (not 2,4), 13.8 (not 13,8), 3.0 (not 3,0), 1.0 (not 1,0). Apply this to EVERY numeric value visible in the figure.
""",
    "Cap04 Fig01.PNG": """
=== CRITICAL FIXES ===
1. ALL decimals must use period: 1.5, 1.9, 12.4, 3.9, 1.0, 2.0, 3.5. Do NOT keep PT comma decimals.
2. Use possessive: 'Fernanda's value' (not 'Fernanda value').
3. Use 'Reference scale' (not 'Reference range').
4. Use 'no defined clinical reference' EXACTLY (not 'no laboratory limit defined').
5. Do NOT add the phrase 'where risk raises death risk' — that phrase is invented.
6. The word 'optimal' is a key term — keep it verbatim. Do NOT use 'great' or 'ideal'.
""",
    "Cap04 Fig02.PNG": """
=== CRITICAL FIXES ===
1. ALL decimals MUST use period: 4.8 (not 4,8), 5.2 (not 5,2), 5.4 (not 5,4), 6.5 (not 6,5), 3.5 (not 3,5), 1.0 (not 1,0). Apply globally.
2. Use 'Triglycerides/HDL ratio' (with the 's' — plural, not 'Triglyceride').
3. Use 'Subclinical myocardial injury' (not 'damage').
4. Use the EXACT phrase 'Genetic atherosclerotic risk — measure once in a lifetime' (NOT 'measured one once in life' — that is broken English).
5. Use 'True kidney function — independent of muscle mass' (not 'does not depend on muscle mass').
""",
    "Cap06 Fig04.PNG": """
=== CRITICAL FIXES ===
1. The TITLE must be EXACTLY: "André's OGTT: when fasting lies."
   Do NOT translate it as 'when the meal deceives the mind' — that is COMPLETELY WRONG. The PT 'jejum' means 'fasting'. The PT 'mente' is the verb 'to lie' (third person), NOT the noun 'mind'. So 'quando o jejum mente' = 'when fasting lies'. Translate this title verbatim: "André's OGTT: when fasting lies."
2. Subtitle must be EXACTLY: "Apparently normal glucose with disproportionate insulin response." (no 'a' before 'disproportionate').
3. Use 'µU/mL' as written, do not change to 'µIU/mL'.
4. Use 'Moderate glycemia.' (not 'Moderate glucose.').
""",
    "Cap08 Fig01.PNG": """
=== CRITICAL FIXES ===
1. The badge at top-left must read 'FIGURE 1' in English — do NOT keep it as 'FIGURA 1'.
2. ALL decimals must use period: '+1.8 MET' (not '+1,8 MET'). Apply to every numeric value.
3. The footer sentence is CRITICAL — the previous attempt INVERTED its meaning. The footer must read EXACTLY:
   "+1.8 MET = 25–30% lower all-cause death risk."
   Do NOT write 'raises death risk' — that is the OPPOSITE of the correct meaning. The PT 'menor risco' = 'lower risk', NOT 'raises risk'. Use the word 'lower' specifically.
4. Use 'VO₂ MAX · STRESS TEST' (preferred) or 'VO₂ MAX · ERGOMETRY' — either is fine, but match the YAML.
""",
    "Cap08 Fig02.PNG": """
=== CRITICAL FIXES ===
1. The dutasteride dose must use period decimal: '0.5 mg/day' (not '0,5 mg/day').
2. The NOT-prescribe block must read EXACTLY: 'We do NOT prescribe finasteride.' Do NOT add 'or dutasteride' — that is invented and wrong. PT only mentions finasteride here.
3. Do NOT add 'iron optimization' to the alternatives list — that is invented. The list is exactly: 'scalp PRP; microneedling; correction of ferritin, TSH, vitamin D, B12, zinc; thyroid nutrition.'
4. Use 'scalp PRP' (keep 'scalp' — do not omit it).
""",
    "Cap09 Fig01.PNG": """
=== CRITICAL FIXES ===
1. The badge must read 'FIGURE 1' in English — do NOT keep 'FIGURA 1'.
2. The subtitle is CRITICAL — the previous attempt INVERTED its meaning. The subtitle must read EXACTLY:
   "Why three drug classes stopped being 'diabetes drugs' in the last five years."
   Do NOT write 'have become the remedy for diabetes' — that is the OPPOSITE meaning. The PT 'deixaram de ser remédio de diabético' = 'stopped being diabetes drugs'. Use 'stopped being' specifically.
""",
    "Cap11 Fig01.PNG": """
=== CRITICAL FIXES ===
1. The title is CRITICAL. PT 'reposição' = 'replacement' (hormone replacement therapy). It does NOT mean 'rest'. The title must read EXACTLY:
   "Paulo: 6 months without replacement — and the trajectory changed"
   Do NOT use 'rest', 'break', or any other word for 'reposição'. Use 'replacement' specifically.
2. ALL decimals use period: 1.0, 0.9, 1.7, 4.8, 11.2 (not 1,0 / 0,9 / 1,7 / 4,8 / 11,2).
""",
    "Cap12 Fig01.PNG": """
=== CRITICAL FIXES ===
1. The first stressor in Box 1 is CRITICAL. PT 'após o expediente' means 'after work hours' / 'after office hours'. It does NOT mean 'after the shipment'. The text must read EXACTLY:
   "Emails after work hours"
   Do NOT translate it as 'Emails after the shipment' — 'shipment' is completely wrong. 'Expediente' refers to the work day, not a delivery.
""",
    "Cap12 Fig02.PNG": """
=== CRITICAL FIXES ===
1. ALL decimals use period: 1.8 (not 1,8), 0.7 (not 0,7). Apply to all CRP values.
2. The subtitle must clearly state that biological optimization moved 0 markers (i.e., none of the biological markers — there are 2 biological ones: CRP and morning cortisol). Use the YAML wording verbatim.
""",
    "Cap13 Fig02.PNG": """
=== CRITICAL FIXES ===
1. Use 'labs' — NOT 'exams'. The title must read EXACTLY: '"The labs are good. I am not."' (with 'I am', not 'I'm'.) PT 'exames' translates to 'labs' in this clinical context, not 'exams' (which would mean tests/quizzes in EN).
2. The PT case mentions 'amigos chegando 3h-5h' which means 'friends arriving for 3-5 hours' — translate as such, NOT 'screens 3h-5h' (screens is a completely different word).
3. PT 'ritual laico' = 'secular ritual' (laico means non-religious/secular). Do NOT translate as 'laid-back ritual'.
""",
}


def run_one(figname: str, critique: str, idx: int, total: int):
    print(f"[{idx}/{total}] starting: {figname}", flush=True)
    t0 = time.time()
    try:
        result = subprocess.run(
            [str(SCRIPT), figname, critique],
            capture_output=True, text=True, timeout=300,
        )
        elapsed = time.time() - t0
        if result.returncode == 0:
            print(f"[{idx}/{total}] ✓ done {figname} ({elapsed:.0f}s)", flush=True)
            return (figname, True, "")
        else:
            err = (result.stderr or result.stdout)[:300]
            print(f"[{idx}/{total}] ✗ FAILED {figname}: {err}", flush=True)
            return (figname, False, err)
    except subprocess.TimeoutExpired:
        return (figname, False, "timeout")


def _parse_yaml(path):
    figures = []
    current = None
    section = None
    pending_pt = None
    with open(path, "r", encoding="utf-8") as f:
        for line in f:
            stripped = line.strip()
            if not stripped or stripped.startswith("#"):
                continue
            if stripped.startswith("- figure:"):
                if current:
                    figures.append(current)
                fig = stripped.split(":", 1)[1].strip().strip('"')
                current = {"figure": fig, "substitutions": [], "notes": [], "keep_unchanged": []}
                section = None
                pending_pt = None
                continue
            if stripped == "substitutions:":
                section = "substitutions"; continue
            if stripped == "notes:":
                section = "notes"; continue
            if stripped == "keep_unchanged:":
                section = "keep_unchanged"; continue
            if section == "substitutions":
                if stripped.startswith("- pt:"):
                    pt_val = stripped.split(":", 1)[1].strip()
                    if pt_val.startswith('"') and pt_val.endswith('"'):
                        pt_val = pt_val[1:-1]
                    pending_pt = pt_val
                elif stripped.startswith("en:") and pending_pt is not None:
                    en_val = stripped.split(":", 1)[1].strip()
                    if en_val.startswith('"') and en_val.endswith('"'):
                        en_val = en_val[1:-1]
                    current["substitutions"].append((pending_pt, en_val))
                    pending_pt = None
            elif section in ("notes", "keep_unchanged"):
                if stripped.startswith("- "):
                    val = stripped[2:].strip()
                    if val.startswith('"') and val.endswith('"'):
                        val = val[1:-1]
                    current[section].append(val)
    if current:
        figures.append(current)
    return figures


def parse_yaml_for_figure(target_figure: str) -> str:
    figs = _parse_yaml(YAML_PATH)
    for f in figs:
        if f["figure"] == target_figure:
            lines = []
            for pt, en in f["substitutions"]:
                lines.append(f"- '{pt}' → use EXACT '{en}'")
            if f.get("notes"):
                lines.append("")
                lines.append("Additional rules:")
                for n in f["notes"]:
                    lines.append(f"- {n}")
            return "\n".join(lines)
    return ""


def main():
    figures = sorted(CRITICAL_FIXES.keys())
    total = len(figures)

    # Build the full prompt = standard YAML substitutions + surgical critique
    jobs = []
    for fig in figures:
        base_prompt = parse_yaml_for_figure(fig)
        full = base_prompt + "\n\n" + CRITICAL_FIXES[fig]
        jobs.append((fig, full))

    results = []
    with ThreadPoolExecutor(max_workers=4) as pool:
        futures = {pool.submit(run_one, f, p, i + 1, total): f for i, (f, p) in enumerate(jobs)}
        for fut in as_completed(futures):
            results.append(fut.result())

    print()
    print(f"=== Critical re-runs: {sum(1 for _, ok, _ in results if ok)}/{total} succeeded ===")
    for fn, ok, msg in sorted(results):
        flag = "✓" if ok else "✗"
        print(f"  {flag} {fn}{('  ' + msg) if not ok else ''}")


if __name__ == "__main__":
    main()
