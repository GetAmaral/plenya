#!/usr/bin/env python3
# Gerador determinístico de points + escada de níveis para os 361 genes.
# Direção: nível ALTO = bom, nível 0 = ruim. Saída: /tmp/genes_plan.csv (revisão) e, com --sql, o SQL.
import csv, re, sys, uuid

ROWS = list(csv.DictReader(open('/tmp/genes_full.csv')))
NS = uuid.UUID('a1f0c0de-0000-0000-0000-00000000ge00'.replace('ge','9e'))  # namespace fixo p/ UUIDs determinísticos

def lvl_uuid(code, level):
    return str(uuid.uuid5(NS, f"plenya-genelevel|{code}|{level}"))

# ---- calibração ----
TIER_PTS = {'T3':8,'T2':2,'T1':1,'T0':1}
SUBGROUP_TIER = {
 'Aptidão Física & Lesão':'T1','Performance':'T1','Pele, Estética & Capilar':'T1',
 'Circadiano & Sono':'T1','Neuro, Humor & Comportamento':'T1','Detoxificação':'T1','Outros':'T1',
 'Imune & Inflamatório':'T2','Imunidade':'T2','Vitaminas & Micronutrientes':'T2',
 'Lipídico & Cardiovascular':'T2','Cardiovascular':'T2','Cardiovascular & Coagulação':'T2',
 'Metabolismo':'T2','Metabolismo, Apetite & Energia':'T2','Neurodegeneração':'T2',
}
# tier/points por símbolo de gene (override do default do subgrupo)
PTS_OVERRIDE = {
 'ACE':10,'ABCC8':8,'ABCA1':8,'FTO':8,'TCF7L2':8,'PCSK9':8,'LDLR':9,'LPL':7,'APOA5':7,'APOB':7,
 'KCNJ11':7,'PPARG':7,'MC4R':7,'CDKAL1':6,'GCK':8,'HNF1A':8,'HNF4A':8,'HNF1B':8,'SLC30A8':6,
 'F5':8,'F2':8,'F13A1':5,'FGB':4,'SERPINE1':5,'HFE':7,'SERPINA1':7,'LRRK2':6,'SNCA':6,'MAPT':6,
 'GRN':6,'C9orf72':7,'APP':8,'PSEN1':9,'PSEN2':9,'MTHFR':6,'IL6':3,
}
T3_SYMBOLS = {'ACE','FTO','TCF7L2','PCSK9','LDLR','APOE','GCK','HNF1A','HNF4A','HNF1B','F5','F2',
 'HFE','SERPINA1','APP','PSEN1','PSEN2','C9orf72'}
# het = 3 (aditivo/dose-resposta) p/ estes símbolos; default 2
DOSE_RESPONSE = {'MTHFR','MTHFD1','MTRR','MTR','FADS1','FADS2','ELOVL2','VDR','GC','CYP2R1','FUT2',
 'TCN2','BCO1','COMT','MCM6','SOD2','NQO1','CAT','GPX1','SLC23A1','SLC23A2','TMPRSS6','HFE'}

def symbol(item):  # 1º token do nome
    m = re.match(r'^([A-Za-z0-9./-]+)', item.strip()); return m.group(1) if m else item

# ---- overrides explícitos (trade-off / sem-direção / casos especiais) ----
# cada um: (points, [(level, name), ...])  -> ladder explícita
def gen_neutral(name='Genótipo (informativo)'):
    return [(3, name)]   # nível único neutro = 60%, não premia nem penaliza
OVR = {
 # trade-offs puros (neutros, informativos)
 'ACTN3_RS1815739_R577X':(1,[(4,'RR (poder)'),(4,'RX (misto)'),(3,'XX (resistência)')]),
 'CYP1A2_RS762551':(2,[(5,'AA (metabolizador rápido)'),(3,'AC (intermediário)'),(2,'CC (metabolizador lento)')]),
 'ADH1B_RS1229984_ARG48HIS':(1,[(4,'His/His (rápido)'),(4,'Arg/His (intermediário)'),(3,'Arg/Arg (lento)')]),
 'IL6_RS1800795_174G_C':(3,[(5,'CC (↓IL-6)'),(3,'GC (intermediário)'),(0,'GG (↑IL-6)')]),  # pop-variável (europeu)
 # APOE SNPs individuais (direção biológica conhecida; peso baixo p/ não triplicar APOE-ε)
 'APOE_RS429358':(3,[(5,'TT (sem ε4)'),(2,'TC (1× ε4)'),(0,'CC (ε4/ε4)')]),
 'APOE_RS7412':(3,[(5,'TT (ε2/ε2 — proteção)'),(4,'CT (1× ε2)'),(3,'CC (sem ε2 — basal)')]),
 # OTHER (variantes de risco nomeadas)
 'LRRK2':(6,[(5,'Não portador (favorável)'),(0,'Portador G2019S (risco)')]),  # casa por símbolo abaixo
 # detox/funcional com alelo entre parênteses (parser não pega)
 'EPHX1_RS1051740_TYR113HIS':(2,[(5,'Tyr/Tyr GG (normal)'),(3,'Tyr/His (intermediário)'),(0,'His/His AA (↓atividade)')]),
 'GSTP1_RS1695_ILE105VAL':(2,[(5,'Ile/Ile AA (normal)'),(3,'Ile/Val (intermediário)'),(0,'Val/Val GG (↓atividade)')]),
 'SOD2_RS4880_ALA16VAL':(2,[(5,'Ala/Ala TT (basal)'),(3,'Ala/Val (intermediário)'),(2,'Val/Val CC (variável)')]),
 'CYP2A6_RS1801272':(2,[(5,'*1/*1 (normal)'),(4,'*1/*2 ou *1/*4 (intermediário)'),(3,'*2/*4 (↓atividade)')]),
 'MCM6_RS4988235':(2,[(5,'TT (persistência da lactase)'),(4,'CT (persistência)'),(2,'CC (intolerância)')]),
 # HLA celíaco — necessário-não-suficiente; informativo-moderado
 # GeneHair capilar (13) — informativo neutro (cosmético; não move escore de longevidade)
}
GENEHAIR_INFORMATIVE = {  # code -> points 1, neutro
 'SRD5A2_RS523349','SRD5A1_RS39848','AR_RS6152','CRABP2_RS12724719','CYP19A1_RS2470152',
 'HDAC9_RS2249817','HDAC9_RS756853','PTGES2_RS13283456','WNT10A_RS7349332','SULT1A1_RS1042028',
 'LIPH_RS201249971','LIPH_RS201868115',
}
# códigos com prefixo intergênico / nomes irregulares no painel capilar
GENEHAIR_NAME_RE = re.compile(r'capilar', re.I)

# ---- parser de legenda -> (kind, risk_token, prot_token) ----
def parse(leg, item):
    body = re.sub(r'^Genótipo\s*\|\s*','',(leg or '').strip())
    body = re.sub(r'^Genótipo\s*','',body).strip()
    low = body.lower()
    if not body or body.startswith('('):
        return ('NODIR',None,None)
    # rare deleterious
    if re.search(r'(muta|expansão|deficiência)',low) and 'proteç' not in low:
        m=re.search(r'(\w+)=defici',body) or re.search(r'(Mutação|Expansão[^=]*)',body)
        tok=m.group(1) if m else 'Mutação'
        return ('RARE_DEL',tok,None)
    # rare protective
    if 'proteç' in low and 'risco' not in low:
        m=re.search(r'(\S+)=proteç',body); return ('RARE_PROT', m.group(1) if m else 'variante', None)
    # functional (arrow/normal)
    if ('↓' in body or '↑' in body or 'normal' in low or 'basal' in low
            or 'persist' in low or 'conversão' in low) and '=risco' not in low:
        mr=re.search(r'([A-Za-z]{1,3})\s*=\s*[↓↑]',body)
        mp=re.search(r'([A-Za-z]{1,3})\s*=\s*(?:normal|basal|persist\w*)',body)
        if mr: return ('TERN', mr.group(1), mp.group(1) if mp else None)
    # risk + prot
    mr=re.search(r'(\w+)=risco',body); mp=re.search(r'(\w+)=(?:proteç\w+|normal)',body)
    if mr:
        return ('TERN', mr.group(1), mp.group(1) if mp else None)
    # OTHER: "G2019S=risco alto", "DQ2.5/2.2=alto risco"
    if 'risco' in low:
        m=re.search(r'(\S+)=.*risco',body) or re.search(r'(\S+).*risco',body)
        return ('RISK_NAMED', m.group(1) if m else None, None)
    return ('UNKNOWN',None,None)

def ladder_tern(risk, prot, het):
    # genotipos. risk/prot podem ser 1 letra, 2 letras (XX) ou aminoácido (3 letras)
    r=risk; p=prot
    if r and len(r)==2 and r[0]==r[1]:  # "AA=risco" -> alelo de risco = letra
        ra=r[0]
        return [(5,'Sem alelo de risco (favorável)'),(het,f'Heterozigoto (1× {ra})'),(0,f'{ra}{ra} (risco)')]
    if r and len(r)==1:
        if p and len(p)==1:
            g_fav=f'{p}{p}'; g_het=''.join(sorted([p,r])); g_risk=f'{r}{r}'
            return [(5,f'{g_fav} (favorável)'),(het,f'{g_het} (heterozigoto)'),(0,f'{g_risk} (risco)')]
        return [(5,'Sem alelo de risco (favorável)'),(het,f'Heterozigoto (1× {r})'),(0,f'{r}{r} (risco)')]
    # aminoácido (His/Arg/Trp/Gly/...) ou token composto
    if r:
        if p:
            return [(5,f'{p}/{p} (favorável)'),(het,f'{p}/{r} (heterozigoto)'),(0,f'{r}/{r} (risco)')]
        return [(5,'Sem alelo de risco (favorável)'),(het,f'Heterozigoto (1× {r})'),(0,f'{r}/{r} (risco)')]
    return None

def build(r):
    code=r['code']; item=r['item']; sub=r['subgroup']; sym=symbol(item)
    flags=[]
    # APOE-ε consolidado: rebuild fixo (corrige inversão)
    if code=='APOE_GENOTIPO_E':
        pts=9
        lad=[(5,'E2/E2 (proteção)'),(4,'E2/E3 (favorável)'),(3,'E3/E3 (mais comum)'),
             (2,'E2/E4'),(1,'E3/E4 (portador ε4 — risco aumentado)'),(0,'E4/E4 (homozigoto ε4 — risco alto)')]
        return pts,lad,'APOE-eps rebuild'
    # GeneHair capilar informativo
    if code in GENEHAIR_INFORMATIVE or (sub=='Pele, Estética & Capilar' and GENEHAIR_NAME_RE.search(item) and parse(r['legend0'],item)[0] in ('NODIR','UNKNOWN')):
        return 1, gen_neutral(), 'GeneHair informativo'
    # overrides explícitos
    if code in OVR:
        pts,lad=OVR[code]; return pts,lad,'override'
    if sym in OVR and OVR[sym][1]:  # LRRK2 etc by symbol
        pts,lad=OVR[sym]; return pts,lad,'override-sym'
    kind,risk,prot=parse(r['legend0'],item)
    # tier/points
    tier = 'T3' if sym in T3_SYMBOLS else SUBGROUP_TIER.get(sub,'T1')
    pts = PTS_OVERRIDE.get(sym, TIER_PTS[tier])
    het = 3 if (sym in DOSE_RESPONSE or kind=='TERN' and ('↓' in (r['legend0'] or '') or '↑' in (r['legend0'] or '') or 'normal' in (r['legend0'] or '').lower())) else 2
    if kind=='TERN':
        lad=ladder_tern(risk,prot,het)
        if not lad: flags.append('TERN-noladder')
        return pts,lad,(';'.join(flags) or f'tern het={het}')
    if kind=='RARE_DEL':
        return pts,[(5,'Não portador (favorável)'),(0,f'Portador ({risk}) — risco')],'rare-del'
    if kind=='RARE_PROT':
        return max(pts,3),[(5,f'Portador ({risk}) — proteção'),(3,'Não portador (basal)')],'rare-prot'
    if kind=='RISK_NAMED':
        return pts,[(5,'Não portador (favorável)'),(0,f'Portador {risk} (risco)')],'risk-named'
    # NODIR/UNKNOWN sem override -> FLAG, não gerar
    return None,None,'!!NEEDS_REVIEW!!'

out=[]; review=[]
for r in ROWS:
    pts,lad,note=build(r)
    if lad is None:
        review.append((r['code'],r['subgroup'],r['item'],r['legend0'],note))
    out.append((r,pts,lad,note))

# CSV de revisão
with open('/tmp/genes_plan.csv','w',newline='') as f:
    w=csv.writer(f); w.writerow(['code','subgroup','gene','points','note','ladder'])
    for r,pts,lad,note in out:
        ladstr=' | '.join(f'{lv}:{nm}' for lv,nm in (lad or []))
        w.writerow([r['code'],r['subgroup'],r['item'],pts,note,ladstr])

# resumo
from collections import Counter
tot=sum(pts for _,pts,lad,_ in out if pts)
print(f"genes={len(out)}  total_points_genetica={tot}")
print("notas:", dict(Counter(n for *_,n in out)))
print(f"\nNEEDS_REVIEW: {len(review)}")
for c,s,i,l,n in review: print(f"  [{s}] {i}  ::  {l}")
if '--sql' in sys.argv:
    with open('/tmp/genes_points_levels.sql','w') as f:
        f.write("-- Genética: points + escada de níveis (gerado por gen_genes.py). Idempotente.\n")
        f.write("-- Direção: nível alto=bom, 0=ruim. UUIDs de nível determinísticos (dev≡prod).\nBEGIN;\n")
        for r,pts,lad,note in out:
            if lad is None: continue
            f.write(f"UPDATE score_items SET points={pts}, updated_at=now() WHERE id='{r['id']}';\n")
            f.write(f"DELETE FROM score_levels WHERE item_id='{r['id']}';\n")
            for idx,(lv,nm) in enumerate(lad):
                lid=lvl_uuid(r['id'],f"{lv}-{idx}"); nm_esc=nm.replace("'","''")
                f.write("INSERT INTO score_levels (id,level,name,operator,item_id,created_at,updated_at) VALUES "
                        f"('{lid}',{lv},'{nm_esc}','=','{r['id']}',now(),now());\n")
        f.write("COMMIT;\n")
    print("\nSQL -> /tmp/genes_points_levels.sql")
