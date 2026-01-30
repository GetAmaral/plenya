#!/usr/bin/env python3
"""Enriquecer 14 items restantes"""
import requests, json, time

API_BASE_URL = "http://localhost:3001/api/v1"
EMAIL, PASSWORD = "import@plenya.com", "Import@123456"

# Content estruturado para items faltantes
MEDICATIONS = {
    "2df78cbd-debf-4032-9901-797275b5b9a9": ("Anti-hipertensivos", "fe579635-390e-4ce0-a970-351333e28cae"),
    "e362aec1-31f5-4a13-a523-beb1be9ebae7": ("Antiarrítmicos", "84d98093-e62f-470f-9e08-6ab525f7d357"),
    "ea412847-5928-4dec-9ee4-e8ed1335b026": ("Antiarrítmicos", "84d98093-e62f-470f-9e08-6ab525f7d357"),
    "e94f9287-d972-43af-b6d1-ee59ed991d31": ("Anticonvulsivantes", "2799ebd7-4c10-4362-9119-d0238a5c4890"),
    "51f52c0e-0d7d-4e2e-a21d-894312525e74": ("Antidiabéticos", "ae0c5229-76f9-4c8a-8b48-761739ac2268"),
    "439634e4-d503-4ff7-a5aa-144f08cae74c": ("Antidiabéticos", "ae0c5229-76f9-4c8a-8b48-761739ac2268"),
    "67050bc8-7933-426d-8b88-ec707e11d1f8": ("Artrite", "8b232987-3a9a-48ab-8051-04356897d97b"),
    "2a70a9bb-e432-46bb-bbc2-81329e69c008": ("Artrite", "8b232987-3a9a-48ab-8051-04356897d97b"),
    "7fb3a872-d7a5-4a36-bfa3-24df546f8617": ("Artrite", "8b232987-3a9a-48ab-8051-04356897d97b"),
    "8addcebb-7bc7-4a97-95fc-d3a4f5010355": ("Asma", "2fd99065-70ad-4d33-86b8-0ac8be744547"),
    "78643ead-ce00-4ef5-b90a-7b4b6cc19d5b": ("Asma", "2fd99065-70ad-4d33-86b8-0ac8be744547"),
    "eaf309c1-6cfc-48cb-8e19-d1268a66431f": ("Asma", "2fd99065-70ad-4d33-86b8-0ac8be744547"),
}

class Enricher:
    def __init__(self):
        self.token = None
        self.session = requests.Session()
        self.content_cache = {}

    def login(self):
        try:
            r = self.session.post(f"{API_BASE_URL}/auth/login", json={"email": EMAIL, "password": PASSWORD})
            r.raise_for_status()
            self.token = r.json()["accessToken"]
            self.session.headers.update({"Authorization": f"Bearer {self.token}"})
            print("✓ Login OK\n")
            return True
        except Exception as e:
            print(f"✗ Login falhou: {e}")
            return False

    def get_content(self, reference_id):
        """Busca conteúdo de um item de referência já enriquecido"""
        if reference_id in self.content_cache:
            return self.content_cache[reference_id]

        try:
            r = self.session.get(f"{API_BASE_URL}/score-items/{reference_id}")
            r.raise_for_status()
            data = r.json()
            content = {
                "cr": data.get("clinicalRelevance", ""),
                "pe": data.get("patientExplanation", ""),
                "cd": data.get("conduct", "")
            }
            self.content_cache[reference_id] = content
            return content
        except Exception as e:
            print(f"    ✗ Erro ao buscar referência: {e}")
            return None

    def update(self, item_id, content):
        try:
            payload = {
                "clinicalRelevance": content["cr"],
                "patientExplanation": content["pe"],
                "conduct": content["cd"]
            }
            r = self.session.put(f"{API_BASE_URL}/score-items/{item_id}", json=payload)
            r.raise_for_status()
            return True
        except Exception as e:
            print(f"    ✗ Erro update: {e}")
            return False

    def run(self):
        results = {"success": 0, "failed": 0}
        print("="*80)
        print("ENRIQUECIMENTO DE 12 ITEMS FALTANTES (CÓPIA DE REFERÊNCIAS)")
        print("="*80 + "\n")

        for idx, (item_id, (name, ref_id)) in enumerate(MEDICATIONS.items(), 1):
            print(f"[{idx}/12] {name} (ID: {item_id[:8]}...)")

            # Buscar conteúdo do item de referência
            content = self.get_content(ref_id)
            if not content or not content["cr"]:
                print(f"    ✗ Referência vazia\n")
                results["failed"] += 1
                continue

            # Aplicar no item alvo
            if self.update(item_id, content):
                print(f"    ✓ Copiado de referência {ref_id[:8]}...\n")
                results["success"] += 1
            else:
                results["failed"] += 1

            time.sleep(0.3)

        print("="*80)
        print(f"RESULTADO FINAL: {results['success']} SUCESSO, {results['failed']} FALHAS")
        print("="*80)
        return results

def main():
    e = Enricher()
    if not e.login():
        return
    e.run()

if __name__ == "__main__":
    main()
