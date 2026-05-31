"""
Run gpt-image-2 translations for all 38 ebook figures.

Reads figure-substitutions.yaml, builds a strict-prompt for each figure,
calls translate-figure.sh in parallel batches, and reports progress.
"""

import subprocess
import sys
import time
from pathlib import Path
from concurrent.futures import ThreadPoolExecutor, as_completed

EBOOK = Path("/home/user/plenya/ebook/Ebook 1 - Performance e Longevidade")
YAML_PATH = EBOOK / "figuras" / "figure-substitutions.yaml"
SCRIPT = EBOOK / "build" / "translate-figure.sh"

# concurrency limit (OpenAI rate limits / network)
MAX_CONCURRENCY = 4


def parse_yaml(path: Path):
    """Minimal YAML parser tailored for this file's structure.
    Returns list of dicts: [{figure, substitutions: [(pt, en)], notes, keep_unchanged}]
    """
    figures = []
    current = None
    section = None  # 'substitutions' | 'notes' | 'keep_unchanged' | None
    pending_pt = None

    with open(path, "r", encoding="utf-8") as f:
        for line in f:
            raw = line.rstrip("\n")
            stripped = raw.strip()

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
                section = "substitutions"
                continue
            if stripped == "notes:":
                section = "notes"
                continue
            if stripped == "keep_unchanged:":
                section = "keep_unchanged"
                continue

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


def build_prompt(fig: dict) -> str:
    """Build a strict-prompt of PT→EN substitutions for a single figure."""
    lines = []
    for pt, en in fig["substitutions"]:
        lines.append(f"- '{pt}' → use EXACT '{en}'")

    if fig.get("notes"):
        lines.append("")
        lines.append("Additional rules:")
        for n in fig["notes"]:
            lines.append(f"- {n}")

    if fig.get("keep_unchanged"):
        lines.append("")
        lines.append("Keep these unchanged (do not translate):")
        for k in fig["keep_unchanged"]:
            lines.append(f"- {k}")

    return "\n".join(lines)


def translate_one(fig: dict, idx: int, total: int) -> tuple[str, bool, str]:
    figname = fig["figure"]
    prompt = build_prompt(fig)
    print(f"[{idx}/{total}] starting: {figname}", flush=True)
    t0 = time.time()
    try:
        result = subprocess.run(
            [str(SCRIPT), figname, prompt],
            capture_output=True, text=True, timeout=300,
        )
        elapsed = time.time() - t0
        if result.returncode == 0:
            print(f"[{idx}/{total}] ✓ done {figname} ({elapsed:.0f}s)", flush=True)
            return (figname, True, result.stdout.strip())
        else:
            print(f"[{idx}/{total}] ✗ FAILED {figname} ({elapsed:.0f}s)", flush=True)
            return (figname, False, (result.stderr or result.stdout)[:500])
    except subprocess.TimeoutExpired:
        print(f"[{idx}/{total}] ✗ TIMEOUT {figname}", flush=True)
        return (figname, False, "timeout")
    except Exception as e:
        print(f"[{idx}/{total}] ✗ ERROR {figname}: {e}", flush=True)
        return (figname, False, str(e))


def main():
    figures = parse_yaml(YAML_PATH)
    print(f"loaded {len(figures)} figures from YAML")

    # optional filter: pass figure names as CLI args
    if len(sys.argv) > 1:
        wanted = set(sys.argv[1:])
        figures = [f for f in figures if f["figure"] in wanted]
        print(f"filtered to {len(figures)} (matching CLI args)")

    total = len(figures)
    results = []

    with ThreadPoolExecutor(max_workers=MAX_CONCURRENCY) as pool:
        futures = {
            pool.submit(translate_one, fig, idx + 1, total): fig["figure"]
            for idx, fig in enumerate(figures)
        }
        for fut in as_completed(futures):
            results.append(fut.result())

    print()
    print(f"=== Results: {sum(1 for _, ok, _ in results if ok)}/{total} succeeded ===")
    for figname, ok, msg in sorted(results):
        flag = "✓" if ok else "✗"
        print(f"  {flag} {figname}{('  ' + msg[:120]) if not ok else ''}")


if __name__ == "__main__":
    main()
