#!/usr/bin/env python3
"""
Script COMPLETO para processar Score Items de exames laboratoriais
Enriquece itens com conteúdo clínico baseado em Medicina Funcional Integrativa
"""

import requests
import json
import sys
from typing import Dict

# Configurações da API
API_BASE_URL = "http://localhost:3001/api/v1"
LOGIN_EMAIL = "import@plenya.com"
LOGIN_PASSWORD = "Import@123456"

# Importar geradores
import sys
sys.path.append('/home/user/plenya/scripts')
from lab_content_generators import (
    generate_colesterol_total_texts,
    generate_hdl_texts,
    generate_ldl_texts
)
from lab_generators_extra import (
    generate_triglicerides_texts,
    generate_creatinina_texts
)

# Items a processar (ID, nome, tipo)
TARGET_ITEMS = [
    # Glicose (já processados, mas incluindo para teste)
    ("6d272987-00de-4f8d-a3d8-abd729dc24f7", "GLICOSE 0 MIN (JEJUM)", "glicose_jejum"),
    ("cf39e4a8-e1ec-4e94-a2a1-596a2a3cdf44", "GLICOSE 120 MIN", "glicose_120"),

    # Colesterol
    ("498a8637-8bf5-45e0-891b-a0c51a47ccc1", "Colesterol Total", "colesterol_total"),
    ("55787b95-d165-48d6-8eb4-e496bd62d509", "HDL Colesterol", "hdl"),
    ("d55af9e8-4358-4cfa-a75c-05dfcd1bfb4d", "LDL Colesterol", "ldl"),

    # Triglicérides
    ("dfa58a95-7d7a-491f-b7c5-d2fe8c694daa", "Triglicerídeos", "triglicerides"),

    # Creatinina
    ("8234724e-f5d6-4862-9fde-84e78e136f88", "Creatinina", "creatinina"),
]


class ScoreItemProcessor:
    def __init__(self):
        self.token = None
        self.processed = []

    def authenticate(self) -> bool:
        """Autentica na API e obtém token"""
        try:
            response = requests.post(
                f"{API_BASE_URL}/auth/login",
                json={"email": LOGIN_EMAIL, "password": LOGIN_PASSWORD},
                timeout=10
            )

            if response.status_code == 200:
                data = response.json()
                self.token = data.get("accessToken")
                print(f"✓ Autenticado com sucesso")
                return True
            else:
                print(f"✗ Falha na autenticação: {response.status_code}")
                return False
        except Exception as e:
            print(f"✗ Erro na autenticação: {e}")
            return False

    def update_item(self, item_id: str, name: str, texts: Dict[str, str]) -> bool:
        """Atualiza um Score Item com conteúdo clínico"""
        try:
            headers = {
                "Authorization": f"Bearer {self.token}",
                "Content-Type": "application/json"
            }

            response = requests.put(
                f"{API_BASE_URL}/score-items/{item_id}",
                json=texts,
                headers=headers,
                timeout=30
            )

            if response.status_code == 200:
                print(f"✓ {name}: Atualizado com sucesso")
                self.processed.append({
                    "id": item_id,
                    "name": name,
                    "success": True,
                    "texts_generated": True
                })
                return True
            else:
                print(f"✗ {name}: Falha {response.status_code} - {response.text}")
                self.processed.append({
                    "id": item_id,
                    "name": name,
                    "success": False,
                    "error": response.text
                })
                return False

        except Exception as e:
            print(f"✗ {name}: Erro {e}")
            self.processed.append({
                "id": item_id,
                "name": name,
                "success": False,
                "error": str(e)
            })
            return False

    def generate_report(self) -> Dict:
        """Gera relatório final"""
        success_count = sum(1 for p in self.processed if p.get("success"))
        failed_count = len(self.processed) - success_count

        return {
            "processed": self.processed,
            "summary": {
                "total": len(self.processed),
                "success": success_count,
                "failed": failed_count
            }
        }


def get_texts_for_type(item_type: str) -> Dict[str, str]:
    """Retorna textos clínicos baseado no tipo de exame"""
    generators = {
        "colesterol_total": generate_colesterol_total_texts,
        "hdl": generate_hdl_texts,
        "ldl": generate_ldl_texts,
        "triglicerides": generate_triglicerides_texts,
        "creatinina": generate_creatinina_texts,
    }

    generator = generators.get(item_type)
    if generator:
        return generator()
    else:
        return None


def main():
    """Função principal"""
    processor = ScoreItemProcessor()

    print("=== Processador COMPLETO de Score Items - Exames Laboratoriais ===\n")

    # Autenticar
    if not processor.authenticate():
        sys.exit(1)

    print(f"\n{'='*70}")
    print(f"Processando {len(TARGET_ITEMS)} items...")
    print(f"{'='*70}\n")

    # Processar cada item
    for item_id, item_name, item_type in TARGET_ITEMS:
        print(f"\n--- {item_name} ({item_type}) ---")

        # Gerar textos baseado no tipo
        texts = get_texts_for_type(item_type)

        if not texts:
            print(f"⊘ Pulando (sem gerador específico para '{item_type}')")
            processor.processed.append({
                "id": item_id,
                "name": item_name,
                "success": False,
                "error": f"No generator for type: {item_type}"
            })
            continue

        # Atualizar via API
        processor.update_item(item_id, item_name, texts)

    # Gerar relatório
    print(f"\n{'='*70}")
    print("RELATÓRIO FINAL")
    print(f"{'='*70}\n")

    report = processor.generate_report()
    print(json.dumps(report, indent=2, ensure_ascii=False))

    # Salvar relatório
    report_path = "/home/user/plenya/lab_items_complete_report.json"
    with open(report_path, "w", encoding="utf-8") as f:
        json.dump(report, f, indent=2, ensure_ascii=False)

    print(f"\n✓ Relatório salvo em: {report_path}")

    # Sumário executivo
    print(f"\n{'='*70}")
    print("SUMÁRIO EXECUTIVO")
    print(f"{'='*70}")
    print(f"Total processado: {report['summary']['total']}")
    print(f"Sucesso: {report['summary']['success']}")
    print(f"Falhas: {report['summary']['failed']}")
    print(f"Taxa de sucesso: {report['summary']['success']/report['summary']['total']*100:.1f}%")


if __name__ == "__main__":
    main()
