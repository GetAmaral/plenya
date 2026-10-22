#!/usr/bin/env python3
"""Gera o áudio do roteiro de fala da aula (gpt-4o-mini-tts) com as pausas marcadas.

Lê docs/transplante/cihdott-2026-10-22/roteiro-de-fala.md a partir de "## Prólogo".
Cada trecho entre marcações vira uma chamada de TTS; `[pausa Ns]` vira N segundos de
silêncio real; `[olhar a plateia]` vira 1,5 s. Cache por hash do texto em audio/cache/.

Uso: python3 scripts/transplante/tts-roteiro.py [--voice ash] [--only 01,02]
Saída: ~/lecture-output/transplante-rede-cihdott/audio/roteiro-completo.mp3

Preço (developers.openai.com, consultado em 2026-10-21): gpt-4o-mini-tts US$ 0,60 por 1M
tokens de texto + US$ 12 por 1M tokens de áudio. Limite de 2.000 tokens por chamada.
"""
import argparse, concurrent.futures as cf, hashlib, json, os, pathlib, re, subprocess, urllib.request

ROOT = pathlib.Path(__file__).resolve().parents[2]
SRC = ROOT / "docs/transplante/cihdott-2026-10-22/roteiro-de-fala.md"
OUT = pathlib.Path.home() / "lecture-output/transplante-rede-cihdott/audio"
CACHE = OUT / "cache"
INSTR = ("Fale em português do Brasil, com sotaque brasileiro natural. Você é um médico nefrologista "
         "experiente dando uma palestra no estilo TED para profissionais de saúde: voz calorosa e firme, "
         "ritmo calmo e pausado, emoção contida, ênfase nas palavras-chave, um respiro entre as frases. "
         "Nas cenas da história, fale baixo e devagar, como quem conta algo importante.")


def key():
    k = os.environ.get("OPENAI_API_KEY")
    if k:
        return k
    for line in (ROOT / "apps/api/.env").read_text().splitlines():
        if line.startswith("OPENAI_API_KEY="):
            return line.split("=", 1)[1].strip()
    raise SystemExit("OPENAI_API_KEY não encontrada")


def parse():
    s = SRC.read_text()
    s = s[s.index("## Prólogo"):]
    slides = re.split(r"^### \[(\d+)\][^\n]*$", s, flags=re.M)
    items = []  # ("t", texto) | ("p", segundos)
    for i in range(1, len(slides), 2):
        num, body = slides[i], slides[i + 1]
        if num == "45":
            break
        body = re.sub(r"^#.*$|^---$", "", body, flags=re.M)
        body = re.sub(r"\([^)]*\)", "", body)
        body = body.replace("[olhar a plateia]", "[pausa 1.5s]")
        for tok in re.split(r"(\[pausa [\d.]+s\])", body):
            m = re.match(r"\[pausa ([\d.]+)s\]", tok)
            if m:
                items.append(("p", float(m.group(1))))
                continue
            paras = [" ".join(p.split()) for p in tok.split("\n\n") if p.strip()]
            for k, para in enumerate(paras):
                if k:
                    items.append(("p", 0.7))
                items.append(("t", para))
        items.append(("p", 1.2))
    return items


def tts(text, voice, k):
    h = hashlib.sha1(f"{voice}|{INSTR}|{text}".encode()).hexdigest()[:16]
    f = CACHE / f"{h}.mp3"
    if f.exists() and f.stat().st_size > 0:
        return f
    body = json.dumps({"model": "gpt-4o-mini-tts", "voice": voice, "input": text,
                       "instructions": INSTR, "response_format": "mp3"}).encode()
    req = urllib.request.Request("https://api.openai.com/v1/audio/speech", data=body,
                                 headers={"Authorization": f"Bearer {k}", "Content-Type": "application/json"})
    for attempt in range(4):
        try:
            with urllib.request.urlopen(req, timeout=180) as r:
                f.write_bytes(r.read())
            return f
        except Exception as e:  # noqa: BLE001
            err = e
    raise SystemExit(f"falhou: {text[:60]}… {err}")


def silence(sec):
    f = CACHE / f"sil-{sec:.2f}.mp3"
    if not f.exists():
        subprocess.run(["ffmpeg", "-loglevel", "error", "-y", "-f", "lavfi", "-i",
                        "anullsrc=r=24000:cl=mono", "-t", f"{sec}", "-b:a", "128k", str(f)], check=True)
    return f


def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("--voice", default="ash")
    a = ap.parse_args()
    CACHE.mkdir(parents=True, exist_ok=True)
    items = parse()
    k = key()
    texts = [t for kind, t in items if kind == "t"]
    print(f"{len(texts)} trechos, {sum(len(t) for t in texts)} caracteres")
    with cf.ThreadPoolExecutor(6) as ex:
        files = dict(zip(texts, ex.map(lambda t: tts(t, a.voice, k), texts)))
    lst = OUT / "concat.txt"
    with lst.open("w") as fh:
        for kind, v in items:
            f = files[v] if kind == "t" else silence(v)
            fh.write(f"file '{f}'\n")
    out = OUT / "roteiro-completo.mp3"
    subprocess.run(["ffmpeg", "-loglevel", "error", "-y", "-f", "concat", "-safe", "0", "-i", str(lst),
                    "-ar", "24000", "-ac", "1", "-b:a", "128k", str(out)], check=True)
    dur = float(subprocess.run(["ffprobe", "-v", "error", "-show_entries", "format=duration", "-of", "csv=p=0",
                                str(out)], capture_output=True, text=True).stdout)
    print(f"{out} · {int(dur // 60)}:{int(dur % 60):02d}")


if __name__ == "__main__":
    main()
