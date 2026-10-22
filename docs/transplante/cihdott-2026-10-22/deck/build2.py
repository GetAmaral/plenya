#!/usr/bin/env python3
"""Aula "Transplante renal: a volta inteira" v2 (CIHDOTT, 22/10/2026).

Padrão: um slide, uma ideia. Foto full-bleed para emoção, escuro para número e frase,
claro só para o essencial. As frases são do deck de 2025 do Dr. Getúlio.
Fontes dos números: docs/transplante/cihdott-2026-10-22/verificacao.md
Rodar: python3 build2.py  → lecture.html
"""
import math, pathlib

HERE = pathlib.Path(__file__).parent
S = []


def pg():
    return f'<div class="pg">{len(S) + 1:02d}</div>'


def src(t):
    return f'<div class="src">{t}</div>' if t else ""


def quote(t):
    return f'<p class="quote">{t}</p>' if t else ""


def photo(img, h1, kicker="", sub="", q="", s="", extra="", soft=False, pos="center", night=False):
    k = f'<p class="kicker">{kicker}</p>' if kicker else ""
    sb = f'<p class="sub">{sub}</p>' if sub else ""
    S.append(f'''<section class="slide photo{' soft' if soft else ''}{' night' if night else ''}">
  <img class="bg" src="images/{img}" style="object-position:{pos}" alt="">
  <div class="txt">{k}<h1>{h1}</h1>{sb}{extra}</div>
  {quote(q)}{src(s)}{pg()}
</section>''')


def chapter(img, n, h1, sub, pos="center"):
    S.append(f'''<section class="slide chapter">
  <img class="bg" src="images/{img}" style="object-position:{pos}" alt="">
  <div class="txt"><div class="num">{n}</div><h1>{h1}</h1><p class="sub">{sub}</p></div>
  {pg()}
</section>''')


def bg(img):
    return f'<img class="bgimg" src="images/{img[:-4]}-blur.jpg" alt=""><div class="veil"></div>' if img else ""


def statement(text, sub="", small="", s="", q="", img=""):
    sb = f'<p class="sub">{sub}</p>' if sub else ""
    sm = f'<p class="small">{small}</p>' if small else ""
    qq = f'<p class="quote" style="margin-top:56px">{q}</p>' if q else ""
    S.append(f'''<section class="slide dark">{bg(img)}
  <div class="center"><p class="statement">{text}</p>{sb}{sm}{qq}</div>
  {src(s)}{pg()}
</section>''')


def number(big, cap, small="", s="", alert=False, q="", img=""):
    sm = f'<p class="small">{small}</p>' if small else ""
    qq = f'<p class="quote" style="margin-top:50px">{q}</p>' if q else ""
    S.append(f'''<section class="slide dark">{bg(img)}
  <div class="center"><div class="big{' alert' if alert else ''}">{big}</div><p class="big-cap">{cap}</p>{sm}{qq}</div>
  {src(s)}{pg()}
</section>''')


def light(h2, body, q="", s="", img="", pos="center"):
    ph = f'<div class="ph"><img src="images/{img}" style="object-position:{pos}" alt=""></div>' if img else ""
    S.append(f'''<section class="slide light{' split' if img else ''}">{ph}
  <div class="wrap"><h2>{h2}</h2>{body}</div>
  {quote(q)}{src(s)}{pg()}
</section>''')


def rows(items):
    return '<div class="rows">' + "".join(f"<div><b>{a}</b><span>{b}</span></div>" for a, b in items) + "</div>"


def ul(items):
    return '<ul class="list">' + "".join(f"<li>{i}</li>" for i in items) + "</ul>"


# ------------------------------------------------------------------ PRÓLOGO
S.append('''<section class="slide cover">
  <div class="ph"><img src="images/santacasa-iscal.jpg" alt=""></div>
  <p class="credit">Foto: Irmandade Santa Casa de Londrina</p>
  <div class="txt">
    <p class="ev">CIHDOTT · Macronorte · PR</p>
    <h1>Transplante renal:<br><em>a volta inteira</em></h1>
    <p class="sub">Da UBS à UBS, pelas mãos de cada um nesta sala</p>
    <p class="meta"><b>Dr. Getúlio Amaral Filho · Dr. Michel Moro<br>Dra. Taynara Fratoni · Dra. Karoline Araujo</b><br>Nefrologistas · Equipe Nefroclínica<br>Santa Casa de Londrina<br>22 de outubro de 2026</p>
  </div>
</section>''')

photo("cihdott-f.jpg", "Uma família diz sim.", kicker="23h · Interior do Paraná", pos="60% center", night=True)
photo("telefone-f.jpg", "Um telefone toca.", kicker="2h · Londrina", pos="60% center", night=True)
statement("Entre essas duas cenas existe uma rede inteira.<br><em>E ela está nesta sala.</em>",
          sub="Hoje vamos percorrê-la de ponta a ponta, da UBS à UBS.", img="maos.jpg")
number("38.925", "brasileiros esperavam um rim em dezembro de 2025. <b style='font-weight:400;color:#E2C48F'>2.093</b> no Paraná.",
       s="Registro Brasileiro de Transplantes 2025, ABTO", alert=True)
S.append(f'''<section class="slide dark teamslide">{bg("equipe-nefroclinica.jpg")}
  <div class="center">
    <div>
      <p class="kicker">Quem somos</p>
      <h2>Equipe Nefroclínica</h2>
      <p class="quote" style="margin-top:44px">Cada transplante que acompanhamos é uma história de superação que nos motiva todos os dias.</p>
    </div>
    <div class="photo-frame"><img src="images/equipe-nefroclinica.jpg" alt="Equipe Nefroclínica"></div>
  </div>{pg()}
</section>''')

# mapa da volta
STN = ["UBS", "Nefrologia e diálise", "Avaliação", "Fila", "CIHDOTT e OPO", "Captação",
       "Cirurgia", "UTI", "Enfermaria e alta", "Ambulatório"]
cx, cy, r = 960, 548, 372
parts = [f'<circle cx="{cx}" cy="{cy}" r="{r}" fill="none" stroke="rgba(201,160,99,.45)" stroke-width="3" stroke-dasharray="2 14" stroke-linecap="round"/>']
for i, n in enumerate(STN):
    a = math.radians(-90 + i * 36)
    x, y = cx + r * math.cos(a), cy + r * math.sin(a)
    first = i == 0
    parts.append(f'<circle cx="{x:.0f}" cy="{y:.0f}" r="{30 if first else 16}" fill="{"#C9A063" if first else "#8FB09E"}"/>')
    c = math.cos(a)
    lx, ly = cx + (r + 50) * c, cy + (r + 50) * math.sin(a) + 14
    anchor = "middle" if abs(c) < .2 else ("start" if c > 0 else "end")
    if abs(c) < .2:
        ly += -26 if math.sin(a) < 0 else 30
    parts.append(f'<text x="{lx:.0f}" y="{ly:.0f}" font-size="48" font-weight="300" fill="{"#E2C48F" if first else "#FFFFFF"}" text-anchor="{anchor}" font-family="Jost, sans-serif">{n}</text>')
parts.append(f'<text x="{cx}" y="{cy - 6}" font-size="64" font-weight="300" fill="#FFFFFF" text-anchor="middle" font-family="Fraunces, serif" font-style="italic">A volta inteira</text>')
parts.append(f'<text x="{cx}" y="{cy + 58}" font-size="34" fill="rgba(255,255,255,.65)" text-anchor="middle" font-family="Jost, sans-serif" letter-spacing="6">DE UM RIM</text>')
S.append(f'''<section class="slide dark ring">
  <svg viewBox="0 0 1920 1080">{"".join(parts)}</svg>{pg()}
</section>''')

# ------------------------------------------------------------------ 1 · ANTES DO TELEFONE
chapter("ubs-rua.jpg", "1", "Antes do telefone", "a rede básica · a nefrologia · a fila")
photo("ubs-f.jpg", "A história de um transplante começa na UBS.", kicker="Catorze anos antes",
      sub="Sr. J., 38 anos. Diabetes descoberto numa consulta de rotina.", pos="65% center")
light("Três marcos quando o rim começa a falhar",
      '<div class="marks"><div><b>&lt; 60</b><span>TFG: encaminhar ao nefrologista</span></div>'
      '<div><b>&lt; 20</b><span>TFG: conversar sobre transplante</span></div>'
      '<div><b>90 dias</b><span>do início da diálise: oferecer a inscrição na lista</span></div></div>',
      s="TFG: taxa de filtração glomerular, mL/min/1,73 m². Prática do serviço · KDIGO 2020 · Portaria GM/MS 8.041/2025 · MS 2014")
photo("dialise.jpg", "Ele chegou tarde.", kicker="O que aconteceu",
      sub="TFG 12. Diálise de urgência, por cateter.<br>Quatro anos de máquina, três vezes por semana.", pos="65% center")
photo("irmaos-f.jpg", "Ana tem 34 anos.<br>O doador é o irmão.", kicker="Ao mesmo tempo, outra história",
      sub="Encaminhada cedo. Diálise programada, por fístula. Cirurgia com data marcada.", pos="60% center")
S.append(f'''<section class="slide dark">
  <div class="center">
    <p class="statement" style="font-size:80px">Dois caminhos.</p>
    <div class="vs">
      <div><p class="who">Doador falecido · Sr. J.</p>
        <div class="r"><span>transplantes renais no Brasil, 2025</span><b>87%</b></div>
        <div class="r"><span>isquemia fria</span><b>20 a 24 h</b></div>
        <div class="r"><span>função tardia do enxerto (Brasil)</span><b>54%</b></div>
        <div class="r"><span>rim funcionando em 5 anos</span><b>70%</b></div></div>
      <div class="alt"><p class="who">Doador vivo · Ana e Paulo</p>
        <div class="r"><span>transplantes renais no Brasil, 2025</span><b>13%</b></div>
        <div class="r"><span>isquemia fria</span><b>~1 h</b></div>
        <div class="r"><span>função tardia do enxerto (EUA)</span><b>3,6%</b></div>
        <div class="r"><span>rim funcionando em 5 anos</span><b>85%</b></div></div>
    </div>
  </div>
  {src("RBT 2025, ABTO · Sandes-Freitas, Transpl Int 2021 · Redfield, Transpl Int 2016 · isquemia fria: prática do serviço")}{pg()}
</section>''')
photo("equipe.jpg", "Ninguém entra na fila sozinho.", kicker="Avaliação pré-transplante",
      sub="Nefrologia, cirurgia, urologia, cardiologia, odontologia, enfermagem, psicologia, serviço social, nutrição e farmácia.",
      q="Cada cuidado nosso é uma ponte para uma nova vida.", pos="70% center")
light("O que a UBS pode adiantar",
      ul(["Vacina de vírus vivo 4 semanas antes", "Dentes tratados", "Rastreios de câncer em dia"]) +
      '<p class="note">Depois do transplante, <b>tríplice viral, varicela e febre amarela ficam contraindicadas.</b></p>',
      s="Manual dos CRIE, 6ª ed., Ministério da Saúde, 2023", img="vacina.jpg", pos="35% center")
light("Estar na fila é ter tarefa",
      ul(["Soro no laboratório a cada 90 dias", "Avisar toda transfusão", "Telefone sempre atualizado"]) +
      '<p class="note"><span class="vivo-tag">Doador vivo</span>Ana não espera chamada. A cirurgia tem data.</p>',
      s="Portaria GM/MS 2.600/2009 (PRC 4/2017) vigente · 8.041/2025 com vigência adiada", img="laboratorio.jpg", pos="70% center")

# ------------------------------------------------------------------ 2 · O SIM
chapter("captacao.jpg", "2", "O sim", "comissões · OPO · central · captação", pos="70% center")
statement("O cuidado com o doador<br>é o <em>primeiro cuidado</em> com o receptor.", img="uti.jpg")
number("1 em 2", "rins de doador falecido no Brasil demora a funcionar.",
       small="O receptor ainda precisa de diálise na primeira semana. As causas, segundo os autores: encaminhamento tardio e manutenção inadequada do doador.",
       s="Sandes-Freitas TV et al. Transpl Int 2021 · 3.992 transplantes, 18 centros")
statement("Doador mantido nas metas:<br>esse risco cai <em>quase pela metade.</em>",
          sub="Pressão, diurese e dose de vasopressor dentro das metas no momento do consentimento da família.",
          s="Malinoski DJ et al. Am J Transplant 2013 · 722 rins, EUA: função tardia de 30% para 17% com as metas cumpridas",
          img="captacao.jpg")
photo("caixa.jpg", "Cada hora conta.", kicker="Isquemia fria",
      sub="Cada hora a mais aumenta em cerca de 4% a chance de o rim demorar a funcionar.<br>No doador vivo, as salas são vizinhas: cerca de uma hora.",
      q="Somos guardiões de milagres que acontecem todos os dias.", s="Sandes-Freitas TV et al. Transpl Int 2021", pos="60% center")
number("22 horas", "de isquemia fria até o rim do Sr. J. voltar a receber sangue.", img="caixa.jpg")

# ------------------------------------------------------------------ 3 · DENTRO DO HOSPITAL
chapter("cirurgia.jpg", "3", "Dentro do hospital", "pré · intra · pós", pos="60% center")
light("Pré-operatório",
      ul(["Jejum: 6 h sólidos, 2 h líquidos claros", "Tricotomia só se preciso, nunca lâmina",
          "Acesso central antes, sonda no centro cirúrgico", "Medo, não alegria: acolher"]),
      q="Cada detalhe no preparo é um tijolo na construção do sucesso.",
      s="Cuidados PO Tx Renal v1.1, ISCAL / Nefroclínica", img="preop.jpg", pos="48% center")
photo("enxerto.jpg", "Três a quatro horas.", kicker="Intraoperatório",
      sub="Cirurgia, anestesia e enfermagem do centro cirúrgico.<br>Na sala ao lado, Paulo doa o rim esquerdo, por vídeo.", pos="65% center")
S.append(f'''<section class="slide anat">
  <img src="images/anatomia.jpg" alt="">
  <div class="txt">
    <h2>O rim é acrescentado, não trocado.</h2>
    <p class="cap">O rim novo fica embaixo, na fossa ilíaca direita.</p>
    <ol><li>Os rins nativos ficam no lugar</li><li>Artéria e veia na ilíaca externa</li><li>Ureter na bexiga, sobre duplo J</li></ol>
    <p class="note" style="color:#5C6F74;font-size:32px;margin-top:26px"><span class="vivo-tag">Doador vivo</span>Na sala ao lado, Paulo doa o rim esquerdo, por vídeo.</p>
  </div>
  {quote("Conhecer a cirurgia nos torna parceiros do cirurgião no cuidado.")}{pg()}
</section>''')
photo("uti2.jpg", "A enfermagem mede o rim de hora em hora.", kicker="UTI · dias 0 a 3",
      extra='<div class="chips"><span>Sinais vitais 15/15 min</span><span>Diurese horária</span><span>PAM 80 a 110 mmHg</span><span>Repor 70% acima de 100 ml/h</span></div>',
      q="Na UTI, nossa vigilância é o escudo protetor do novo rim.", s="Cuidados PO Tx Renal v1.1, ISCAL / Nefroclínica", pos="60% center", night=True)
photo("coleta.jpg", "<span style='color:#E2C48F'>7h30</span> colhe.<br><span style='color:#E2C48F'>8h</span> dá a dose.", kicker="O nível do tacrolimus",
      sub="Colhido fora da janela, o número mente, e a dose sai errada.",
      q="Somos os guardiões da precisão que protege o novo rim.", pos="60% center")
light("As regras do nosso serviço",
      rows([("Indução", "timoglobulina, só com 5 ou mais incompatibilidades HLA"),
            ("Manutenção", "tacrolimus, micofenolato e prednisona"),
            ("Pneumocistose", "sulfametoxazol-trimetoprima por 3 meses, sempre"),
            ("Citomegalovírus", "profilaxia se houve indução, ou doador positivo e receptor negativo")]),
      s="Serviço de Transplante Renal, Santa Casa de Londrina", img="imunossupressores.jpg", pos="60% center")
photo("mao.jpg", "Dia 2.<br>O rim não urina.", kicker="A virada", sub="A família pergunta: rejeitou?", pos="60% center")
statement("Não. O rim carrega as <em>últimas horas</em> de outra pessoa.",
          sub="Função tardia do enxerto: esperada no doador falecido, e ainda assim 41% mais risco de perder o rim em 3 anos. PAM no alvo, volemia e diurese anotada com honestidade mudam esse desfecho.",
          small="No doador vivo, a função tardia é de 3,6%.",
          s="Braz J Nephrol 2026;48(4) · Redfield, Transpl Int 2016", img="mao.jpg")
photo("dia7.jpg", "Dia 7.<br>Ele urina.", kicker="A espera", sub="Dialisou no terceiro dia. No sétimo, o rim acordou.", pos="65% center")
statement("Esperado é estável.<br><em style='color:#E08A70'>Súbito é urgência.</em>",
       sub="Queda súbita da diurese em quem urinava bem é trombose até prova em contrário. Doppler agora.",
       q="Nossa observação atenta pode ser a diferença entre sucesso e complicação.")
S.append(f'''<section class="slide dark">{bg("dia7.jpg")}
  <svg viewBox="0 0 1920 1080" style="position:absolute;inset:0;width:100%;height:100%;z-index:2"><path d="M1306.5,110.3 A430,430 0 0 1 1693.6,391.5 L1491.8,465.8 A215,215 0 0 0 1298.3,325.2 Z" fill="#C9A063"/><text x="1480" y="271" text-anchor="middle" font-family="Jost, sans-serif" font-size="40" font-weight="500" fill="#0A2730">Enfermagem</text><text x="1480" y="319" text-anchor="middle" font-family="Fraunces, serif" font-style="italic" font-size="36" fill="#0A2730">a vigilância</text><path d="M1703.8,422.9 A430,430 0 0 1 1555.9,877.9 L1423.0,709.0 A215,215 0 0 0 1496.9,481.5 Z" fill="#8FB09E"/><text x="1597" y="610" text-anchor="middle" font-family="Jost, sans-serif" font-size="38" font-weight="500" fill="#0A2730">Farmácia</text><text x="1597" y="650" text-anchor="middle" font-family="Jost, sans-serif" font-size="38" font-weight="500" fill="#0A2730">clínica</text><text x="1597" y="696" text-anchor="middle" font-family="Fraunces, serif" font-style="italic" font-size="36" fill="#0A2730">a prescrição</text><path d="M1529.2,897.3 A430,430 0 0 1 1050.8,897.3 L1170.4,718.7 A215,215 0 0 0 1409.6,718.7 Z" fill="#5E9AA8"/><text x="1290" y="854" text-anchor="middle" font-family="Jost, sans-serif" font-size="40" font-weight="500" fill="#0A2730">Nutrição</text><text x="1290" y="902" text-anchor="middle" font-family="Fraunces, serif" font-style="italic" font-size="36" fill="#0A2730">o prato</text><path d="M1024.1,877.9 A430,430 0 0 1 876.2,422.9 L1083.1,481.5 A215,215 0 0 0 1157.0,709.0 Z" fill="#D49A7A"/><text x="983" y="632" text-anchor="middle" font-family="Jost, sans-serif" font-size="40" font-weight="500" fill="#0A2730">Psicologia</text><text x="983" y="680" text-anchor="middle" font-family="Fraunces, serif" font-style="italic" font-size="36" fill="#0A2730">a adesão</text><path d="M886.4,391.5 A430,430 0 0 1 1273.5,110.3 L1281.7,325.2 A215,215 0 0 0 1088.2,465.8 Z" fill="#B7C7A0"/><text x="1100" y="249" text-anchor="middle" font-family="Jost, sans-serif" font-size="38" font-weight="500" fill="#0A2730">Serviço</text><text x="1100" y="289" text-anchor="middle" font-family="Jost, sans-serif" font-size="38" font-weight="500" fill="#0A2730">social</text><text x="1100" y="335" text-anchor="middle" font-family="Fraunces, serif" font-style="italic" font-size="36" fill="#0A2730">o caminho</text><defs><clipPath id="cc"><circle cx="1290" cy="540" r="193"/></clipPath></defs><image href="images/sr-j-rosto.jpg" x="1097" y="347" width="386" height="386" clip-path="url(#cc)" preserveAspectRatio="xMidYMid slice"/><circle cx="1290" cy="540" r="193" fill="none" stroke="#C9A063" stroke-width="4"/></svg>
  <div style="position:absolute;z-index:3;left:130px;top:0;bottom:0;width:640px;display:flex;flex-direction:column;justify-content:center">
    <p class="kicker">A equipe multiprofissional</p>
    <p class="statement" style="font-size:84px">Cinco guardiões em volta de um rim.</p>
    <p class="sub" style="font-size:38px;margin-top:34px">E a enfermagem em todos os turnos.</p>
    <p class="quote" style="margin-top:56px">Técnica com humanização transforma procedimentos em cuidado integral.</p>
  </div>{pg()}
</section>''')
photo("alta-f.jpg", "A alta é uma prova oral.", kicker="Enfermaria e alta",
      sub="Ele sabe cada remédio e horário, o que fazer se vomitar a dose, por que nunca parar e quando ligar.",
      q="Educar é empoderar: um paciente bem orientado é nosso melhor aliado.", pos="60% center")

# ------------------------------------------------------------------ 4 · DE VOLTA PARA CASA
chapter("chegada.jpg", "4", "De volta para casa", "ambulatório · UBS · comunidade", pos="70% center")
light("Quem cuida do quê",
      '<div class="cols"><div><h3>Ambulatório de transplante</h3><p>O enxerto, a imunossupressão, as complicações.</p><p>E o doador vivo, todo ano, para sempre.</p></div>'
      '<div><h3>UBS</h3><p>Hipertensão, diabetes, colesterol, muitas vezes causados pelo próprio remédio.</p><p>A porta de entrada para todo o resto.</p></div></div>',
      s="KDIGO, doador vivo, 2017 · KDIGO, receptor, 2009")
S.append(f'''<section class="slide dark dsplit">
  <div class="ph"><img src="images/prescricao.jpg" alt=""></div>
  <div class="center">
    <p class="statement" style="font-size:72px">Antes de prescrever: <em>isso mexe no tacrolimus?</em></p>
    <div class="drows">
      <div><b>Sobem o nível</b><span>diltiazem, verapamil, claritromicina, fluconazol</span></div>
      <div><b>Derrubam o nível</b><span>rifampicina, fenitoína, carbamazepina</span></div>
      <div><b>Agridem o rim</b><span>anti-inflamatórios</span></div>
    </div>
    <p class="small" style="margin-top:56px">E o imunossupressor nunca é suspenso na UBS. Se faltar, é urgência.</p>
  </div>
  <p class="quote" style="position:absolute;left:130px;bottom:96px">Conhecer os riscos nos torna detectores precoces de problemas.</p>
  {src("KDIGO 2009 · bula do tacrolimo")}{pg()}
</section>''')
light("Quando ligar para o ambulatório",
      '<div class="cols" style="margin-top:30px">' + ul(["Febre acima de 38 °C", "Diurese caindo", "Dor no enxerto"]) +
      ul(["Vômito que perde a dose", "Diarreia: o tacrolimus sobe", "Creatinina subindo, sem sintoma"]) + "</div>",
      s="Lemahieu W et al. Am J Transplant 2005")
photo("acs-f.jpg", "A conversa sobre doação começa em casa.", kicker="Na comunidade",
      sub="45% das famílias disseram não em 2025. Quase sempre porque nunca ouviram a pessoa dizer que queria doar.",
      s="Registro Brasileiro de Transplantes 2025, ABTO", pos="65% center")

# ------------------------------------------------------------------ EPÍLOGO
photo("um-ano-f.jpg", "Sem diálise.", kicker="Um ano depois",
      sub="A família que disse sim nunca vai ver esta cena. Ele nunca vai saber o nome dela. A lei guarda os dois lados.",
      extra='<p class="sub" style="font-size:34px;color:#E2C48F;margin-top:26px">E Ana e Paulo almoçam juntos todo domingo.</p>',
      s="Decreto 9.175/2017, art. 52", pos="65% center")
photo("maos.jpg", "O rim do Sr. J. passou por mais de vinte mãos.", sub="A maior parte delas está nesta sala.", soft=True)
S.append(f'''<section class="slide dark">{bg("um-ano.jpg")}
  <div class="center" style="align-items:center;text-align:center">
    <p class="statement" style="font-style:italic;max-width:1400px">O conhecimento nos capacita, mas é o <em>amor pelo que fazemos</em> que transforma vidas.</p>
    <p class="sub" style="margin-top:70px">Converse com a sua família sobre doação de órgãos.</p>
    <p class="small" style="margin-top:60px;max-width:none">Dr. Getúlio Amaral Filho · Santa Casa de Londrina</p>
  </div>{pg()}
</section>''')

REFS = ["ABTO. Registro Brasileiro de Transplantes 2025.",
        "Sandes-Freitas TV et al. Transpl Int 2021. PMID 33742470.",
        "Malinoski DJ et al. Am J Transplant 2013. PMID 23406284.",
        "Redfield RR et al. Transpl Int 2016. PMID 26432507.",
        "KDIGO 2020, candidatos a transplante renal. Transplantation 2020;104(4):708-14.",
        "KDIGO, cuidado do receptor de transplante renal. Am J Transplant 2009. PMID 19845597.",
        "Lentine KL et al. KDIGO, doador vivo. Transplantation 2017. PMID 28742762.",
        "A nova era dos tratamentos para receptores de transplante renal. Braz J Nephrol 2026;48(4).",
        "Lemahieu W et al. Am J Transplant 2005. PMID 15888045.",
        "Ministério da Saúde. Diretrizes clínicas para o cuidado ao paciente com DRC no SUS, 2014.",
        "Ministério da Saúde. Manual dos CRIE, 6ª ed., 2023.",
        "Portaria GM/MS 2.600/2009 (PRC 4/2017). Portaria GM/MS 8.041/2025.",
        "Lei 9.434/1997. Decreto 9.175/2017.",
        "Amaral Filho G, Moro M. Cuidados PO Tx Renal v1.1. ISCAL / Nefroclínica, 2023."]
S.append(f'''<section class="slide light">
  <div class="wrap" style="top:110px;bottom:110px;justify-content:flex-start"><h2>Referências</h2>
  <ol class="refs">{"".join(f"<li>{r}</li>" for r in REFS)}</ol></div>{pg()}
</section>''')

S = [x.replace('<section class="slide', f'<section id="s{i + 1:02d}" class="slide', 1) for i, x in enumerate(S)]
html = f'''<!DOCTYPE html>
<html lang="pt-BR"><head><meta charset="UTF-8"><title>Transplante renal: a volta inteira</title>
<link rel="stylesheet" href="v2.css"></head><body>
{chr(10).join(S)}
</body></html>'''
(HERE / "lecture.html").write_text(html)
print("ok", len(S), "slides")
