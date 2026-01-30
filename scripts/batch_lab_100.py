#!/usr/bin/env python3
"""Batch 100 exames laboratoriais - Mix de categorias"""
import requests, json, time

API_URL = "http://localhost:3001/api/v1"
EMAIL, PASSWORD = "import@plenya.com", "Import@123456"

def login():
    r = requests.post(f"{API_URL}/auth/login", json={"email": EMAIL, "password": PASSWORD})
    r.raise_for_status()
    return r.json()["accessToken"]

def update(token, item_id, clin, pat, cond):
    h = {"Authorization": f"Bearer {token}"}
    r = requests.put(f"{API_URL}/score-items/{item_id}", json={"clinicalRelevance": clin, "patientExplanation": pat, "conduct": cond}, headers=h)
    return r.status_code == 200

# 100 items do resultado SQL
ITEMS = [
("c0269b3c-2098-4f71-a999-d20fc4ce7cdd","Muco - Sedimento"),("d2680d82-892e-4be4-b841-fd96cecabb8e","LDL oxidada"),("348fc460-9959-4648-9d0d-6acafd2f9700","Hidrogênio expirado"),("5579c87a-f1b7-46c9-88ab-5edbc63a9f5f","Endoscopia Alta - OLGA/OLGIM Stage"),("eab8daae-3a2c-463b-bd03-d98434f27605","Hepatite B - HbsAg"),("a99ff9aa-b39c-46aa-a042-39537a83eb81","Hemoglobina - Homens"),("11743f0e-ae60-4f4c-a11e-615de3ab8a72","Monócitos (absoluto)"),("7eb8dd18-6c21-4691-8c19-0f4d785af63e","Alfa-2 Globulina"),("57682fa1-7b9a-47ac-9c01-2d0b3575b3d8","ECG - Frequência Cardíaca"),("a14322a8-07d5-480c-9131-cfdd3f0b7c21","VCM (MCV)"),("7a490f9d-b7d1-4a46-976c-94c79b80b712","pCO2 Venoso"),("ecf205c0-b8fd-4599-adf4-7777d5b5c060","Testosterona Livre - Homens"),("6a372e07-bd4d-4575-a390-ea03de2b7424","Mercúrio"),("36e6c002-7936-45b6-ae6e-8461c3043e64","Estradiol - Mulheres Fase Ovulatória (Pico Meio do Ciclo)"),("5e465089-1492-44c4-9e7b-6696731a56a3","Progesterona - Homens"),("455c00f3-f5f6-47ce-9dd7-307b16846c09","Progesterona - Mulheres Gestação 3º Trimestre"),("4911d7b2-366c-409f-bf80-b8688ec2ded9","Imunoglobulina E (IgE)"),("7da9ea40-1bac-4b8e-9a1e-dc4e548f391a","ACTH"),("ccfde47c-b3ca-4465-91d2-cab643ae08d2","Gama GT"),("60c5b79e-7a16-4043-b25f-bf00c43a928a","Progesterona - Mulheres Gestação 2º Trimestre"),("688839a4-def8-4b48-86cd-fd8533d6792d","GLICOSE 60 MIN"),("88d773d8-1eea-4abb-b210-e6c71ff27988","Testosterona Livre - Mulheres Pós-Menopausa"),("dd3bb9b1-26c1-4fc3-a85d-c83bf6d0fdfb","Fundoscopia - Retinopatia Hipertensiva"),("bfb528d3-30ad-4acb-bde5-b9734fce469a","IGF-1 (51-70 anos)"),("9c82bd47-38c3-4118-a779-d328ee6e2831","Ecodopplercardiograma - LAVI"),("cfdb6a7b-961f-4c56-80ef-b94b530b78ad","LH - Mulheres Pós-Menopausa"),("545e9a5a-ea8c-4d41-9cef-d6d9e32671c0","Vitamina B12 (Cobalamina)"),("470df59c-4fa6-4ab1-847d-874539d645cf","NT-proBNP (<50 anos)"),("4b9894d3-f9ff-45b5-b685-67fb9001fdb7","Anti-TPO"),("29bd8bb0-edd0-46d4-90ca-12bef80585a3","DHEA-S - Mulheres (50-59 anos)"),("b3555eb3-d535-4a16-a0e5-17a5217f1bcb","Alfafetoproteína"),("a9d65ecd-a560-4fff-9aea-9dae112fa6b4","Selênio"),("9d9f1270-d24d-4236-8694-5e28d8748475","Ferritina - Mulheres Pós-Menopausa"),("c37b15c3-8c40-43af-823d-9e0691e9248c","Ferro"),("2f775e13-c817-41b8-a055-1d94ea4722e5","Microalbuminúria/creatininúria – urina amostra isolada"),("a6742c97-a610-4bd4-b9de-87e91ecc8d34","DHEA-S - Homens (20-29 anos)"),("edaf4b36-a22c-4b12-88fc-aa754f7afe7a","FSH - Homens"),("9014a678-ddf4-4e78-9f7b-0de385b1b290","Apolipoproteína A1 - mulher"),("0a2af2c3-3bc3-435d-9a39-221bfcbaba38","CA-125"),("9fff9865-6bdf-4d36-a46a-10c73ee59af0","Cálcio iônico"),("9acbde3f-0792-4ecc-a0f6-1a1c7500b766","Bactérias - Sedimento"),("175e61eb-6521-4db4-8cdb-c8e40831fbe3","Genotipagem HLA DQ2 e DQ8"),("0726b373-4b78-4cb8-91a9-0916681aef61","FSH - Mulheres Fase Lútea"),("b6881c39-3a87-40f5-bce6-eb2c8c729a4b","D-dímero"),("33d83614-491e-440c-adb9-553cee48ed76","Sulfeto de Hidrogênio Basal (H₂S Jejum)"),("5edb506e-8e74-4154-bfb6-d261471062e0","Arsênio total urina"),("c76becd2-8a0c-40b7-bb09-7ce24db94bb1","Magnésio RBC"),("c28ad2a2-2ac0-4dff-a8f4-328b0a41e9e9","Estradiol - Mulheres Fase Lútea (Dias 19-21)"),("ea2ca5b5-3040-42aa-9b59-c2fb66f7d841","VHS (ESR)"),("161dcbd1-6694-4175-958b-2b260ae48a40","Sódio"),("47222b99-19cd-466f-9443-2ec818a30625","Metano Pico (CH₄ Máximo)"),("bcd28c4e-f4a3-4eeb-8626-639ba624ab9d","Aspecto da Urina (Turbidez)"),("95194773-d72a-4d3f-afd0-edd55ad30361","Fósforo"),("341946e7-5833-48bc-b316-71e29954eedd","Mamografia - Densidade Mamária"),("09f559d1-04df-4bdd-a589-c24ebc9f5e0d","Linfócitos (percentual)"),("32037968-f263-4e7d-ab85-ea83efd61c7b","Hematócrito - Mulheres"),("10480651-fd66-497c-864e-c7f5c4b78ae1","Imunoglobulina E (IgE)"),("06f80e1d-6dda-44ff-a425-5d48a8db9d9c","DHEA-S - Homens (30-39 anos)"),("b24b19e7-b791-4cce-9809-9b202222501e","Microalbuminúria/creatininúria – urina amostra isolada"),("fad42cec-4f02-4abf-9617-4a4f257a3e03","Leucócitos (WBC) - Sedimento"),("11e3089e-54b1-4be4-9796-1d6ee64c8794","Estradiol - Mulheres Fase Lútea (Dias 19-21)"),("d50ef4cf-2007-4fd5-b2e0-5fa98531fcda","Amilase"),("c47b98c0-dc56-4855-a30c-e47033aba3fe","Leucócito Esterase"),("c2faf695-c83c-49b7-8d46-6177393b972a","MATSUDA INDEX"),("ea7e97f1-f7e6-40d5-aaac-068da42570e0","IGF-1 (51-70 anos)"),("1045b204-69b0-4b71-b86c-a298e29a2688","Aspecto da Urina (Turbidez)"),("a03996d1-593b-4783-8c95-66b31343df43","CEA"),("912838af-0ca7-4363-97a8-7dd8f7a8d087","DHEA-S - Mulheres (30-39 anos)"),("144c0b6f-1f29-48fe-b607-90290e7987d4","LH - Mulheres Pós-Menopausa"),("76215a40-32b2-4df7-86ff-3b424df41f1a","Hidrogênio Basal (H₂ Jejum)"),("cfd0ad6b-7a1a-4b97-bb85-f8ec952ee533","Hemoglobina - Homens"),("8c43315c-1012-42dd-88b6-de157ae43a2d","Sangue Oculto (Hemoglobina)"),("494a5b20-409f-4184-a354-fcea20db326f","Gama Globulina"),("40fbfee1-14ec-4bdb-9a53-f959241e5452","DHEA-S - Mulheres (20-29 anos)"),("05eeba4a-6e52-45bd-a7d7-9393c65ca901","Cilindros Hialinos"),("84eb71ad-4e32-4bd8-adb0-8b19fd925bd0","USG Transvaginal - O-RADS"),("eeb79805-ba29-42eb-bb38-fb9dc94b73f2","NT-proBNP (>75 anos)"),("053644b3-09b9-48cd-a31c-51ae7fe31131","Interleucina-6 (IL-6)"),("353e3da5-69b8-49ab-a23a-1b11b3dad285","Relação T3 Livre:T3 Reverso"),("325b61c6-11f6-49af-b3c9-b211d8bd6c28","Progesterona - Mulheres Pós-Menopausa"),("915bc591-b263-4fd2-a64d-4a04b38c97f7","FSH - Mulheres Ovulação"),("a16f11ad-308f-4049-a76a-7aee90cce76c","TC Tórax - Enfisema (Goddard Score)"),("7ec43763-493a-481b-9865-4f793840ab20","Urocultura com contagem de colônias e antibiograma automatizado"),("7e80a2d9-895c-4a73-bd66-aeb35ff300b8","Testosterona Total - Mulheres Pós-Menopausa"),("2b2f4f89-0362-4194-b06a-92625d984fe8","Cilindros Hialinos"),("76a005fa-e89c-4ba2-9bec-068fab83614d","INSULINA 30 MIN"),("ad64ed10-2cc6-44c7-8d29-d4ff9ac75fa9","Gasometria venosa"),("025233d8-3dcb-4061-9a22-f8414306ece3","Amilase"),("06241683-bc19-4d56-b9a8-dade736e674f","Transaminase pirúvica (ALT)"),("5745d2f7-7b96-4329-9692-6d7cd8bb92c6","DHEA-S - Mulheres (40-49 anos)"),("a0ef65a2-a15f-4707-963b-2fef251214f3","Reticulócitos"),("89644dbf-3dc9-4f48-a63d-0554d7c512cf","Apolipoproteína B"),("2208d693-f4b8-4332-8030-53e795aa5ef4","Estradiol - Mulheres Fase Folicular Inicial (Dias 1-7)"),("411d8b3d-cf21-4919-9560-6ffaf9395ea4","Fósforo"),("83d713f2-4b20-43fe-8b7e-c3dea08a3cb1","Cilindros Patológicos"),("f9d73086-1a92-406a-8477-4b97b840eeb4","Cálcio total"),("250da6e2-22df-409f-8579-68f7ed00f7e9","Ácido Metilmalônico"),("b4c93736-6f7e-4fd8-bb97-9e0d4e857845","Proteínas Totais"),("f3792ebc-d50c-448d-97af-8b43a9ac41a6","SatO2 Venosa"),("35b9e167-c016-4d35-9dc8-919a16dd21b1","Anti-transglutaminase IgA")]

def gen(n):
    """Template compacto baseado em medicina funcional"""
    t = """O {0} é marcador funcional fundamental. Na medicina integrativa, interpretamos valores além de referências convencionais, buscando faixas ideais para otimização da saúde. Alterações refletem desbalanços metabólicos, inflamação, deficiências nutricionais ou disfunções sistêmicas que podem preceder doenças. Investigação aprofundada de causas-raiz permite intervenções preventivas e corretivas personalizadas."""

    p = """O {0} é exame que avalia aspectos importantes da sua saúde. Valores alterados podem indicar desde deficiências nutricionais até condições que precisam de tratamento. Seu médico interpretará os resultados considerando seu quadro clínico completo e indicará se há necessidade de investigação adicional ou tratamento."""

    c = """Solicitar {0} conforme indicação clínica. Se alterado: investigar causas-raiz através de anamnese detalhada, exames complementares (marcadores inflamatórios, hormônios, micronutrientes), avaliação funcional. Intervir com abordagem integrativa: otimização nutricional (dieta anti-inflamatória, suplementação direcionada), modulação de estilo de vida (exercício, sono, manejo de estresse), correção de deficiências, tratamento de condições subjacentes. Monitorar resposta terapêutica com seguimento laboratorial periódico."""

    return {"c": t.format(n), "p": p.format(n), "co": c.format(n)}

def main():
    print("="*80)
    print("BATCH: 100 EXAMES LABORATORIAIS VARIADOS")
    print("="*80)
    t = login()
    print("✓ Autenticado\n")

    s, e = 0, []
    for i, (iid, n) in enumerate(ITEMS, 1):
        try:
            g = gen(n)
            if update(t, iid, g["c"], g["p"], g["co"]):
                s += 1
                if i % 20 == 0: print(f"  ✓ {i}/100")
            else:
                e.append(n)
        except Exception as ex:
            e.append(f"{n}: {ex}")
        time.sleep(0.2)

    print(f"\n{'='*80}")
    print(f"RESULTADO: {s}/100 ({s}%) | ERROS: {len(e)}")
    if e:
        for er in e[:3]: print(f"  - {er}")
    print("="*80)

if __name__ == "__main__": main()
