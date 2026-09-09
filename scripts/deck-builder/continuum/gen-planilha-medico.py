#!/usr/bin/env python3
"""
Gera planilha-medico.html (A4) a partir do MODELO de custo do Continuum Médico.
Os números não são digitados: saem do mesmo cálculo da planilha
docs/continuum/precificacao-continuum-medico.md. Render: node render-a4.js
"""
import datetime, pathlib

# ---- parâmetros herdados da planilha MVP (docs/continuum/precificacao-mvp.md) ----
IMPOSTO, GATEWAY = 0.1433, 0.0299
F = 1 - IMPOSTO - GATEWAY                    # 0,8268
TARIFA_SESSAO = 600                          # médico, tarifa MVP (Trim #1)
WPP_MES = 200                                # retainer enxuto, cortado pela metade
ENC = {"sem": 6, "anu": 12}
PRECO = {"sem": 12000, "anu": 20000}
RISCO = 0.05

# pools anuais já com a fatia Continuum de 80%
POOL_INFRA, POOL_CONTAB, POOL_SEGURO = 2048, 7680, 2400
POOL_EMR, POOL_LGPD, POOL_JUR_SETUP = 18386, 3840, 9600
POOL_CAC = 55680 + 24000 + 4086              # IG+tráfego · mentoria · sites

def blocos(N, plano):
    anual = plano == "anu"
    m = 2 if anual else 1                    # o que escala com a duração
    meses = 12 if anual else 6
    enc = ENC[plano] * TARIFA_SESSAO
    wpp = WPP_MES * meses
    A = enc + wpp
    C = m*round(POOL_INFRA/N) + m*70 + m*round(POOL_EMR/(3*N))
    D = m*round(POOL_CONTAB/N) + m*round(POOL_SEGURO/N)
    E = round(POOL_CAC/N)
    Fj = round(POOL_JUR_SETUP/(3*N)) + m*round(POOL_LGPD/N)
    sub = A + C + D + E + Fj
    I = round(sub*RISCO)
    return dict(enc=enc, wpp=wpp, A=A, B=0, C=C, D=D, E=E, F=Fj, J=0, I=I, cfixo=sub+I)

def margem(plano, cfixo, preco=None):
    P = preco or PRECO[plano]
    return (P*F - cfixo)/P*100

def lucro(plano, cfixo, preco=None):
    P = preco or PRECO[plano]
    return P*F - cfixo

def breakeven(plano):
    N = 8
    while lucro(plano, blocos(N, plano)["cfixo"]) < 0:
        N += 1
    return N

def br(v, dec=0):
    s = f"{v:,.{dec}f}".replace(",", "\x00").replace(".", ",").replace("\x00", ".")
    return s

b24 = {p: blocos(24, p) for p in ("sem", "anu")}
b32 = {p: blocos(32, p) for p in ("sem", "anu")}
b40 = {p: blocos(40, p) for p in ("sem", "anu")}

# ---- comparação com o Continuum completo (planilha MVP travada em 24-25/05) ----
COMPLETO = {"sem": dict(cfixo=28878, preco=37000, A=19350, enc=29),
            "anu": dict(cfixo=50571, preco=65000, A=36450, enc=52)}

def linha_bloco(rot, chave, origem):
    return (f"<tr><td class='rot'>{rot}</td>"
            f"<td class='num'>{br(b24['sem'][chave])}</td>"
            f"<td class='num'>{br(b24['anu'][chave])}</td>"
            f"<td class='org'>{origem}</td></tr>")

hoje = datetime.date(2026, 9, 8).strftime("%d/%m/%Y")

# ---------- sensibilidade de sala ----------
def sala_row(valor_hora):
    ds = ENC["sem"]*valor_hora*(1+RISCO)
    da = ENC["anu"]*valor_hora*(1+RISCO)
    return (f"<tr><td class='rot'>R$ {br(valor_hora)} por hora</td>"
            f"<td class='num {'neg' if margem('sem', b24['sem']['cfixo']+ds)<0 else ''}'>{br(margem('sem', b24['sem']['cfixo']+ds),1)}%</td>"
            f"<td class='num {'neg' if margem('anu', b24['anu']['cfixo']+da)<0 else ''}'>{br(margem('anu', b24['anu']['cfixo']+da),1)}%</td></tr>")

# ---------- sensibilidade do WhatsApp ----------
def wpp_row(rot, valor_mes, destaque=False):
    ds = (valor_mes - WPP_MES)*6*(1+RISCO)
    da = (valor_mes - WPP_MES)*12*(1+RISCO)
    ms, ma = margem('sem', b24['sem']['cfixo']+ds), margem('anu', b24['anu']['cfixo']+da)
    cls = " class='hl'" if destaque else ""
    return (f"<tr{cls}><td class='rot'>{rot}</td>"
            f"<td class='num {'neg' if ms<0 else ''}'>{br(ms,1)}%</td>"
            f"<td class='num {'neg' if ma<0 else ''}'>{br(ma,1)}%</td></tr>")

HTML = f"""<!DOCTYPE html>
<html lang="pt-BR"><head><meta charset="UTF-8">
<title>Continuum Médico · Estrutura de custos</title>
<link href="https://fonts.googleapis.com/css2?family=Cormorant+Garamond:ital,wght@0,400;0,500;0,600;1,400&family=Inter:wght@300;400;500;600&display=swap" rel="stylesheet">
<style>
  :root {{
    --gold:#B38645; --gold-soft:#D4A86B; --petrol:#063B4F; --petrol-deep:#041F2A;
    --ocean:#417E8E; --sage:#92B8B4; --cream:#EAE7DA; --cream-soft:#F7F4EC;
    --ink:#0A1F26; --muted:#5A6B70; --rule:#D8D2C4;
    --serif:'Cormorant Garamond',Georgia,serif; --sans:'Inter',system-ui,sans-serif;
  }}
  @page {{ size: A4; margin: 20mm 18mm 18mm 18mm; }}
  * {{ margin:0; padding:0; box-sizing:border-box; }}
  body {{ font-family: var(--sans); font-size: 9.4pt; line-height:1.55; color: var(--ink); background:#fff; }}

  .capa {{ height: 247mm; display:flex; flex-direction:column; justify-content:center; page-break-after: always; }}
  .capa .mark {{ display:flex; align-items:center; gap:14px; margin-bottom:52px; }}
  .capa .mark img {{ height:34px; }}
  .capa h1 {{ font-family:var(--serif); font-weight:500; font-size:34pt; line-height:1.08; letter-spacing:-.02em; color:var(--petrol); margin-bottom:14px; }}
  .capa .sub {{ font-family:var(--serif); font-style:italic; font-size:14pt; color:var(--gold); margin-bottom:44px; }}
  .capa .meta {{ font-size:8.6pt; color:var(--muted); line-height:1.9; border-top:1px solid var(--rule); padding-top:18px; max-width:105mm; }}
  .capa .meta b {{ color:var(--ink); font-weight:500; }}
  .capa .conf {{ margin-top:34px; font-size:8pt; letter-spacing:.16em; text-transform:uppercase; color:var(--gold); }}

  h2 {{ font-family:var(--serif); font-weight:500; font-size:17pt; color:var(--petrol); letter-spacing:-.01em;
       margin:0 0 4px; padding-bottom:7px; border-bottom:1.5px solid var(--gold); }}
  h2 + .lead {{ font-family:var(--serif); font-style:italic; font-size:11pt; color:var(--gold); margin:8px 0 14px; }}
  h3 {{ font-family:var(--serif); font-weight:600; font-size:11.5pt; color:var(--petrol); margin:18px 0 7px; }}
  section {{ margin-bottom:26px; }}
  p {{ margin-bottom:9px; max-width:none; }}
  ul {{ margin:0 0 10px 15px; }} li {{ margin-bottom:5px; }}
  strong {{ font-weight:600; color:var(--petrol-deep); }}
  .brk {{ page-break-before: always; }}
  .avoid {{ page-break-inside: avoid; }}

  table {{ width:100%; border-collapse:collapse; margin:10px 0 12px; font-size:9pt; page-break-inside:avoid; }}
  thead {{ display:table-header-group; }}
  th {{ text-align:left; font-weight:500; font-size:7.8pt; letter-spacing:.1em; text-transform:uppercase;
        color:var(--muted); padding:0 8px 6px; border-bottom:1px solid var(--petrol); }}
  th.num, td.num {{ text-align:right; font-variant-numeric: tabular-nums; }}
  td {{ padding:6px 8px; border-bottom:1px solid var(--rule); vertical-align:top; }}
  td.rot {{ color:var(--ink); }}
  td.org {{ color:var(--muted); font-size:8.2pt; }}
  tr.tot td {{ border-top:1.5px solid var(--petrol); border-bottom:none; font-weight:600; color:var(--petrol-deep); padding-top:8px; }}
  tr.hl td {{ background:var(--cream-soft); font-weight:600; }}
  td.neg {{ color:#8C2F1D; }}
  .nota {{ background:var(--cream-soft); border-left:2.5px solid var(--gold); padding:11px 14px; margin:12px 0; font-size:8.8pt; }}
  .nota b {{ color:var(--petrol-deep); }}
  .kpi {{ display:flex; gap:0; margin:14px 0 16px; border-top:1px solid var(--rule); border-bottom:1px solid var(--rule); }}
  .kpi div {{ flex:1; padding:12px 14px; border-right:1px solid var(--rule); }}
  .kpi div:last-child {{ border-right:none; }}
  .kpi .k {{ font-size:7.6pt; letter-spacing:.12em; text-transform:uppercase; color:var(--muted); margin-bottom:5px; }}
  .kpi .v {{ font-family:var(--serif); font-size:19pt; font-weight:500; color:var(--petrol); line-height:1; }}
  .kpi .u {{ font-size:8pt; color:var(--muted); margin-top:4px; }}
</style></head><body>

<div class="capa">
  <div class="mark">
    <img src="../../../apps/site/public/brand/symbol/gold.png" alt="">
    <img src="../../../apps/site/public/brand/wordmark/ink.png" alt="Plenya" style="height:22px;">
  </div>
  <h1>Continuum Médico<br>Estrutura de custos<br>e formação de preço</h1>
  <div class="sub">Versão enxuta do Continuum, conduzida por médico</div>
  <div class="meta">
    <b>Data:</b> {hoje}<br>
    <b>Base de custo:</b> planilha MVP travada em 24 e 25/05/2026<br>
    <b>Volume de referência:</b> 24 pacientes por ano<br>
    <b>Regime:</b> Lucro Presumido, presunção de 32%
  </div>
  <div class="conf">Documento interno · uso restrito aos sócios</div>
</div>

<section>
  <h2>O que é o Continuum Médico</h2>
  <p class="lead">O mesmo método, sem a equipe multidisciplinar e sem o box físico.</p>
  <p>O <strong>Continuum Plenya</strong> completo custa {br(COMPLETO['sem']['preco'])} reais no semestral e
  {br(COMPLETO['anu']['preco'])} no anual, e reúne quatro profissionais, encontro semanal e quatro boxes por ciclo.
  Ele continua existindo e não foi reprecificado, mas <strong>sai de venda</strong>. O que vai a mercado é o
  Continuum Médico: o mesmo Método AGIR e o mesmo Escore, conduzidos só por médico.</p>
  <p>A razão econômica da mudança está num bloco só. Os quatro profissionais custavam
  R$ {br(COMPLETO['sem']['A'])} por ciclo semestral, {br(COMPLETO['sem']['A']/COMPLETO['sem']['cfixo']*100,0)}% de todo
  o custo do programa. O médico sozinho, com a mesma tarifa, custa R$ {br(b24['sem']['A'])}.</p>

  <table class="avoid">
    <tr><th>&nbsp;</th><th>Continuum (stand-by)</th><th>Continuum Médico</th></tr>
    <tr><td class="rot">Equipe no ciclo</td><td>4 profissionais</td><td>só médico</td></tr>
    <tr><td class="rot">Encontros, semestral e anual</td><td>{COMPLETO['sem']['enc']} e {COMPLETO['anu']['enc']} toques</td><td>{ENC['sem']} e {ENC['anu']}</td></tr>
    <tr><td class="rot">Box Plenya</td><td>4 boxes por ciclo</td><td>não existe</td></tr>
    <tr><td class="rot">Formato dos encontros</td><td>100% online</td><td>online ou presencial em Londrina</td></tr>
    <tr><td class="rot">WhatsApp com o médico</td><td>R$ 400 por mês</td><td>R$ {br(WPP_MES)} por mês</td></tr>
    <tr class="tot"><td class="rot">Preço, semestral</td><td>R$ {br(COMPLETO['sem']['preco'])}</td><td>R$ {br(PRECO['sem'])}</td></tr>
    <tr class="tot"><td class="rot">Preço, anual</td><td>R$ {br(COMPLETO['anu']['preco'])}</td><td>R$ {br(PRECO['anu'])}</td></tr>
  </table>
  <p>O que ficou de fora não sumiu: o painel genético é <strong>add-on cobrado à parte</strong>, e suplementos
  e manipulados passam a ser prescrição que o paciente compra na farmácia de sua escolha.</p>
</section>

<section>
  <h2>Como o preço foi formado</h2>
  <p class="lead">A âncora veio primeiro; a planilha disse quanto de margem sobrava.</p>
  <p>Imposto e gateway são percentuais da receita, não custos fixos, então o preço é um cálculo circular.
  Ele se resolve isolando os custos que <em>não</em> dependem do preço, o C fixo:</p>
  <div class="nota" style="text-align:center; font-family:var(--serif); font-size:12pt;">
    Preço = C fixo &divide; (1 &minus; {br(IMPOSTO*100,2)}% &minus; {br(GATEWAY*100,2)}% &minus; margem)
    &nbsp;&nbsp;&rarr;&nbsp;&nbsp; fator de escala <b>{br(F,4)}</b>
  </div>
  <p>A carga tributária efetiva de {br(IMPOSTO*100,2)}% soma IRPJ, CSLL, PIS, COFINS e o ISS de Londrina, em
  cenário conservador de Lucro Presumido. O gateway é a taxa do Asaas no cartão parcelado. A margem é
  <strong>lucro líquido sobre o preço</strong>, depois de todos os custos.</p>
  <p>Diferente do produto completo, aqui o preço foi decidido antes: fixou-se a âncora comercial e ajustou-se
  a estrutura até caber. Foi o que motivou o corte do retainer de WhatsApp, detalhado adiante.</p>
</section>

<section class="brk">
  <h2>Estrutura de custo por bloco</h2>
  <p class="lead">Custo por paciente e por ciclo, a 24 pacientes por ano.</p>
  <table>
    <tr><th>Bloco</th><th class="num">Semestral</th><th class="num">Anual</th><th>Origem</th></tr>
    <tr><td class="rot">A &middot; Encontros ({ENC['sem']} e {ENC['anu']} &times; R$ {br(TARIFA_SESSAO)})</td>
        <td class="num">{br(b24['sem']['enc'])}</td><td class="num">{br(b24['anu']['enc'])}</td>
        <td class="org">tarifa MVP do médico</td></tr>
    <tr><td class="rot">A &middot; WhatsApp (R$ {br(WPP_MES)} por mês)</td>
        <td class="num">{br(b24['sem']['wpp'])}</td><td class="num">{br(b24['anu']['wpp'])}</td>
        <td class="org">retainer de disponibilidade</td></tr>
    <tr><td class="rot"><b>A &middot; total de honorários</b></td>
        <td class="num"><b>{br(b24['sem']['A'])}</b></td><td class="num"><b>{br(b24['anu']['A'])}</b></td><td class="org"></td></tr>
    <tr><td class="rot">B &middot; Box Plenya</td><td class="num">0</td><td class="num">0</td>
        <td class="org">removido do produto</td></tr>
    {linha_bloco('C &middot; Infraestrutura e amortização do EMR', 'C', 'VPS, vídeo, storage, dev')}
    {linha_bloco('D &middot; Contabilidade e seguro de responsabilidade', 'D', 'rateado por volume')}
    {linha_bloco('E &middot; Aquisição de cliente', 'E', 'tráfego, mentoria, sites')}
    {linha_bloco('F &middot; Jurídico e LGPD', 'F', 'setup amortizado e sob demanda')}
    <tr><td class="rot">J &middot; Painel genético e coordenação</td><td class="num">0</td><td class="num">0</td>
        <td class="org">add-on e absorvido</td></tr>
    {linha_bloco('I &middot; Risco e contingência (5%)', 'I', 'inadimplência, cancelamento, troca')}
    <tr class="tot"><td class="rot">C fixo por ciclo</td>
        <td class="num">R$ {br(b24['sem']['cfixo'])}</td><td class="num">R$ {br(b24['anu']['cfixo'])}</td><td class="org"></td></tr>
  </table>
  <p>Para comparação, o Continuum completo tem C fixo de R$ {br(COMPLETO['sem']['cfixo'])} no semestral e
  R$ {br(COMPLETO['anu']['cfixo'])} no anual. O Médico custa
  <strong>{br(b24['sem']['cfixo']/COMPLETO['sem']['cfixo']*100,0)}%</strong> e
  <strong>{br(b24['anu']['cfixo']/COMPLETO['anu']['cfixo']*100,0)}%</strong> disso.</p>
  <div class="nota">
    <b>O segundo maior custo é a aquisição.</b> Sem o box, o CAC passa a ser
    {br(b24['sem']['E']/b24['sem']['cfixo']*100,0)}% do C fixo semestral. Ele é alto porque o volume é baixo:
    o gasto de marketing é dividido por 24 pacientes. É a única alavanca relevante que não se resolve
    cortando, e sim vendendo mais.
  </div>
</section>

<section class="brk">
  <h2>Preço, margem e o peso do volume</h2>
  <p class="lead">Nos preços travados, a rentabilidade é quase toda função do volume.</p>
  <div class="kpi avoid">
    <div><div class="k">Semestral</div><div class="v">R$ {br(PRECO['sem'])}</div><div class="u">6 &times; R$ {br(PRECO['sem']/6)}</div></div>
    <div><div class="k">Anual</div><div class="v">R$ {br(PRECO['anu'])}</div><div class="u">12 &times; R$ {br(PRECO['anu']/12,2)}</div></div>
    <div><div class="k">Break-even</div><div class="v">{breakeven('sem')} e {breakeven('anu')}</div><div class="u">pacientes por ano</div></div>
  </div>
  <table>
    <thead><tr><th>Volume</th><th class="num">C fixo sem.</th><th class="num">Lucro sem.</th><th class="num">Margem sem.</th>
        <th class="num">C fixo anual</th><th class="num">Lucro anual</th><th class="num">Margem anual</th></tr></thead>
    {''.join(f"<tr><td class='rot'>{n} pacientes por ano</td>"
             f"<td class='num'>{br(b['sem']['cfixo'])}</td><td class='num'>{br(lucro('sem', b['sem']['cfixo']))}</td>"
             f"<td class='num'>{br(margem('sem', b['sem']['cfixo']),1)}%</td>"
             f"<td class='num'>{br(b['anu']['cfixo'])}</td><td class='num'>{br(lucro('anu', b['anu']['cfixo']))}</td>"
             f"<td class='num'>{br(margem('anu', b['anu']['cfixo']),1)}%</td></tr>"
             for n, b in ((24, b24), (32, b32), (40, b40)))}
  </table>
  <div class="nota">
    <b>Os dois planos nascem perto do break-even.</b> No volume de hoje o semestral rende
    R$ {br(lucro('sem', b24['sem']['cfixo']))} por ciclo e o anual R$ {br(lucro('anu', b24['anu']['cfixo']))}. Ambos são ruído.
    Até a cadência do anual subir para {ENC['anu']} encontros, ele rendia 9,2% e sustentava o produto;
    agora não sustenta mais. <b>A margem do negócio deixou de estar na modalidade e passou a estar no volume:</b>
    aos 32 pacientes por ano os dois voltam a cerca de 10%.
  </div>
  <p>O anual continua valendo mais para quem compra e para nós. Dois semestrais custam
  R$ {br(2*PRECO['sem'])} contra R$ {br(PRECO['anu'])} do anual, uma economia de
  R$ {br(2*PRECO['sem']-PRECO['anu'])} ou {br((2*PRECO['sem']-PRECO['anu'])/(2*PRECO['sem'])*100,1)}%, porque aquisição e
  onboarding acontecem uma vez só. No produto completo esse mesmo desconto é de 12,2%.</p>
</section>

<section class="brk">
  <h2>As duas decisões que definiram a margem</h2>

  <h3>1. O retainer de WhatsApp foi cortado pela metade</h3>
  <p>A R$ 400 por mês o retainer custava R$ {br(400*6)} no ciclo semestral. No Continuum completo isso era 6,5% de
  um preço de R$ {br(COMPLETO['sem']['preco'])}. Num produto de R$ {br(PRECO['sem'])} vira <strong>20% do preço</strong>,
  e sozinho colocava a âncora em prejuízo. Cortado para R$ {br(WPP_MES)}, com prazo de resposta mais estreito em troca.</p>
  <table class="avoid">
    <tr><th>Cenário, a 24 pacientes por ano</th><th class="num">Margem semestral</th><th class="num">Margem anual</th></tr>
    {wpp_row('R$ 400 por mês, como no completo', 400)}
    {wpp_row(f'R$ {br(WPP_MES)} por mês, adotado', WPP_MES, destaque=True)}
    {wpp_row('sem retainer', 0)}
  </table>

  <h3>2. O presencial entrou sem custo de sala, por premissa</h3>
  <p>Desde 08/09 o paciente escolhe entre encontro online e presencial na clínica de Londrina. A planilha
  <strong>não ganhou linha de custo de sala</strong>, pela mesma convenção que já rege a aquisição: a clínica
  existe, é usada pela Consulta Plenya e pelo presencial avulso, e é carregada pelo resto da prática. Um
  paciente do Continuum que escolhe presencial ocupa uma sala já paga, e o honorário do médico é o mesmo nas
  duas modalidades.</p>
  <p><strong>Esta é a premissa mais frágil do modelo.</strong> O semestral absorve
  R$ {br(lucro('sem', b24['sem']['cfixo'])/(1+RISCO)/ENC['sem'])} por encontro de hora-sala antes de zerar; o anual
  absorve R$ {br(lucro('anu', b24['anu']['cfixo'])/(1+RISCO)/ENC['anu'])}.</p>
  <table class="avoid">
    <tr><th>Se a sala for cobrada do programa, 100% presencial</th><th class="num">Margem semestral</th><th class="num">Margem anual</th></tr>
    {sala_row(50)}{sala_row(100)}{sala_row(150)}
  </table>
  <div class="nota">
    <b>Qualquer rateio de sala, por menor que seja, põe os dois planos em prejuízo.</b> Enquanto o anual tinha
    10 encontros ele ainda tinha colchão; com {ENC['anu']}, não tem mais. Se a clínica passar a cobrar hora-sala dos
    programas, ou se o presencial virar a escolha da maioria e a capacidade tiver que crescer, os dois preços
    precisam ser revistos.
  </div>
</section>

<section>
  <h2>Riscos assumidos, explicitamente</h2>
  <ul>
    <li><strong>Margem fina no volume atual.</strong> Os dois planos rendem valores que são ruído a 24
      pacientes por ano. O produto foi desenhado para converter e para escalar, não para dar margem já.</li>
    <li><strong>O CAC é um artefato de volume baixo.</strong> Cada paciente a mais no ano derruba o custo de
      todos os outros. É o item que mais responde a crescimento.</li>
    <li><strong>O rateio de 80% da aquisição é o parâmetro mais sensível.</strong> Os
      R$ {br(b24['sem']['E'])} do bloco E já são a fatia do <em>programa</em> dentro do gasto de marketing; os outros
      20% ficam com o resto da prática, que segue existindo. Se a Consulta Plenya sair do funil, esse número sobe
      para R$ {br(round(POOL_CAC/24/0.8))} e a âncora do semestral passa a dar prejuízo.</li>
    <li><strong>Caixa e payback.</strong> Aquisição e onboarding são pagos adiantados e as parcelas entram ao
      longo de 6 a 12 meses. Com margem desta ordem <em>não há folga para antecipar recebível</em>: a
      antecipação do Asaas custa a partir de 1,25% ao mês e comeria a margem inteira.</li>
    <li><strong>Estimativas ainda não cotadas.</strong> O seguro de responsabilidade civil é estimativa de
      mercado, sem apólice fechada.</li>
  </ul>
</section>

<section>
  <h2>Em aberto</h2>
  <ul>
    <li>Se o anual segue a R$ {br(PRECO['anu'])} tendo passado de 10 para {ENC['anu']} encontros, isto é, 20% mais entrega
      sem mudança de preço. A R$ 22.000, em 12 parcelas de R$ 1.833,33, a margem voltaria a cerca de 10%.</li>
    <li>Se o presencial tem teto de encontros por ciclo. Hoje não há limite definido.</li>
    <li>Se há upgrade do Médico para o Continuum completo no meio do ciclo, e como seria precificado.</li>
    <li>Cotação real da apólice de responsabilidade civil.</li>
  </ul>
  <div class="nota" style="margin-top:16px;">
    Todos os números deste documento são gerados a partir do mesmo modelo de custo, sem transcrição manual.
    Detalhamento em <b>docs/continuum/precificacao-continuum-medico.md</b>; parâmetros herdados em
    <b>docs/continuum/precificacao-mvp.md</b>.
  </div>
</section>

</body></html>"""

out = pathlib.Path(__file__).parent / "planilha-medico.html"
out.write_text(HTML, encoding="utf-8")
print(f"escrito: {out}")
print(f"C_fixo sem {b24['sem']['cfixo']} / anu {b24['anu']['cfixo']}  |  margens {margem('sem', b24['sem']['cfixo']):.1f}% / {margem('anu', b24['anu']['cfixo']):.1f}%")
