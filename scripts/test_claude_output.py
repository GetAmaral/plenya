#!/usr/bin/env python3
import subprocess
import tempfile
import os

prompt = "Retorne apenas um JSON com três campos: clinical_relevance, patient_explanation e conduct. Cada campo deve ter 'teste' como valor."

with tempfile.NamedTemporaryFile(mode='w', suffix='.txt', delete=False) as f:
    f.write(prompt)
    prompt_file = f.name

result = subprocess.run(
    ['claude', '-p', '--model', 'claude-sonnet-4-5-20250929', prompt_file],
    capture_output=True,
    text=True,
    timeout=60
)

os.unlink(prompt_file)

print("STDOUT:")
print(result.stdout)
print("\nSTDERR:")
print(result.stderr)
print(f"\nReturn code: {result.returncode}")
