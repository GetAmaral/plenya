#!/usr/bin/env python3
"""Monta lecture.html da aula "Transplante renal: a volta inteira" (CIHDOTT, 22/10/2026).

Roteiro e fontes: /home/user/plenya/docs/transplante/cihdott-2026-10-22/
(arco-proposta.md e verificacao.md). Rodar: python3 build.py
"""
import math, pathlib

HERE = pathlib.Path(__file__).parent
TOTAL = 24
BRAND = "Santa Casa de Londrina"

STATIONS = ["UBS", "Nefrologia e diálise", "Avaliação", "Fila", "CIHDOTT e OPO",
            "Captação", "Cirurgia", "UTI", "Enfermaria e alta", "Ambulatório de transplante"]
SHORT = ["UBS", "Nefrologia", "Avaliação", "Fila", "Doação", "Captação",
         "Cirurgia", "UTI", "Enfermaria", "Ambulatório"]
WHO = ["agente comunitário, enfermagem, médico de família",
       "nefrologia, enfermagem da diálise",
       "equipe multiprofissional",
       "laboratório HLA, Central Estadual",
       "comissões intra-hospitalares, OPO",
       "cirurgia, logística, transporte",
       "cirurgia, anestesia, enfermagem",
       "enfermagem, intensivistas",
       "enfermagem e equipe multi",
       "nefrologia do transplante, enfermagem"]


def ring(cx, cy, r, active=None, who=False, fs=40, wfs=30):
    out = []
    out.append(f'<circle cx="{cx}" cy="{cy}" r="{r}" fill="none" stroke="#E2E8E4" stroke-width="6"/>')
    # seta do sentido da volta
    n = len(STATIONS)
    for i, name in enumerate(STATIONS):
        a = math.radians(-90 + i * 360 / n)
        x, y = cx + r * math.cos(a), cy + r * math.sin(a)
        on = (i == active)
        col = "#B38645" if (on or i == 0) else "#8FB09E"
        rr = 26 if on or i == 0 else 17
        out.append(f'<circle cx="{x:.0f}" cy="{y:.0f}" r="{rr}" fill="{col}" stroke="#FFFFFF" stroke-width="5"/>')
        lx, ly = cx + (r + 52) * math.cos(a), cy + (r + 52) * math.sin(a)
        c = math.cos(a)
        anchor = "middle" if abs(c) < 0.2 else ("start" if c > 0 else "end")
        if abs(c) < 0.2:
            ly += -22 if math.sin(a) < 0 else 40
        else:
            ly += 12
        weight = "500" if (on or i == 0) else "400"
        fill = "#3C5257"
        out.append(f'<text x="{lx:.0f}" y="{ly:.0f}" font-size="{fs}" font-weight="{weight}" fill="{fill}" text-anchor="{anchor}">{name}</text>')
        if who:
            out.append(f'<text x="{lx:.0f}" y="{ly + wfs + 12:.0f}" font-size="{wfs}" fill="#6E7F83" text-anchor="{anchor}">{WHO[i]}</text>')
    return "\n".join(out)


def map_svg(vb_w, vb_h, cx, cy, r, center_lines, who=False, fs=40, wfs=30):
    c = []
    y0 = cy - (len(center_lines) - 1) * 32
    for k, (t, size, color, w) in enumerate(center_lines):
        c.append(f'<text x="{cx}" y="{y0 + k * 64}" font-size="{size}" font-weight="{w}" fill="{color}" text-anchor="middle" font-family="Jost, sans-serif">{t}</text>')
    return (f'<svg viewBox="0 0 {vb_w} {vb_h}" role="img" aria-label="Mapa circular da rede do transplante renal, da UBS à UBS">'
            f'<g font-family="Inter, sans-serif">{ring(cx, cy, r, who=who, fs=fs, wfs=wfs)}{"".join(c)}</g></svg>')


def trail(active):
    if active is None:
        return ""
    dots = "".join(f'<i class="{"on" if i == active else ("done" if i < active else "")}"></i>' for i in range(len(SHORT)))
    return f'<div class="trail"><span class="tl">{SHORT[active]}</span>{dots}</div>'


def foot(n, refs="", wv=False):
    r = f'<span class="ref">{refs}</span>' if refs else ""
    return f'<div class="foot{" wv" if wv else ""}"><span class="pg">{n:02d} / {TOTAL}</span>{r}</div>'


def vivo(*ps, who=None):
    w = f'<p class="who">{who}</p>' if who else ""
    return '<aside class="vivo"><div class="vl">No doador vivo</div>' + w + "".join(f"<p>{p}</p>" for p in ps) + "</aside>"


def brand():
    return f'<div class="brand">{BRAND}</div>'


ANAT = (HERE / "anatomia.svg.html").read_text()

S = []

# 01 capa
S.append(f'''<section class="slide s" id="s01">
  <div class="ph ph--cover" style="width:54%"><img src="images/santacasa-iscal.jpg" alt="" style="object-position:50% 50%"></div>
  <p class="credit">Foto: Irmandade Santa Casa de Londrina</p>
  <div class="body cover" style="right:calc(54% + 80px)">
    <h1 class="cover-t">Transplante renal: a volta inteira</h1>
    <p class="cover-s">Da UBS à UBS, pelas mãos de cada um nesta sala</p>
    <div class="rule"></div>
    <p class="cover-m">
      <b>Dr. Getúlio Amaral Filho</b><br>
      Serviço de Transplante Renal<br>
      Santa Casa de Londrina<br>
      22 de outubro de 2026
    </p>
  </div>
</section>''')

# 02 duas cenas
S.append(f'''<section class="slide s" id="s02">
  <div class="duo"><img src="images/familia.png" alt=""><img src="images/caso.png" alt=""></div>
  <p class="duo-line">Duas cenas ligadas por um rim.</p>
  <div class="duo-cap">
    <p>Num hospital da macrorregião, uma família diz sim.</p>
    <p>Em Londrina, o telefone de um homem em hemodiálise toca.</p>
  </div>
</section>''')

# 03 mapa
S.append(f'''<section class="slide s" id="s03">
  {brand()}
  <div class="body top" style="top:120px; bottom:110px">
    <div class="map">
      {map_svg(1720, 740, 860, 372, 242, [("A volta inteira", 50, "#5A7075", 300), ("de um rim", 50, "#5A7075", 300)])}
    </div>
    <div class="note note--warm" style="margin-top:4px"><b>6.697</b> transplantes renais no Brasil em 2025, recorde, metade da necessidade. <b style="color:var(--alert)">38.925</b> na fila do rim, 2.093 no Paraná.</div>
  </div>
  {foot(3, "<b>[1]</b>Registro Brasileiro de Transplantes 2025. ABTO, 2026")}
</section>''')

# 04 catorze anos antes
S.append(f'''<section class="slide s has-vivo" id="s04">
  {brand()}{trail(0)}
  <div class="body wv">
    <p class="kicker">Ato I · Catorze anos antes</p>
    <h2>Aos 38 anos, diabetes na UBS</h2>
    <ol class="n">
      <li>Taxa de filtração glomerular (TFG) abaixo de 30: encaminhar ao nefrologista</li>
      <li>TFG abaixo de 20: conversar sobre transplante e avaliação</li>
      <li>Inscrição em lista pela norma do Sistema Nacional de Transplantes</li>
    </ol>
    <div class="note note--warm"><span class="lbl">O que aconteceu</span>Ele chegou ao nefrologista com TFG 12 e começou a hemodiálise de urgência, por cateter.</div>
  </div>
  {vivo("Ana, 34 anos, glomerulonefrite.", "Encaminhada cedo. Começou a diálise por fístula, programada, com o irmão já em avaliação.")}
  {foot(4, "<b>[2]</b>MS 2014 e 2015 · <b>[3]</b>KDIGO 2020 · <b>[4]</b>Portaria 8.041/2025", wv=True)}
</section>''')

# 05 dois caminhos
S.append(f'''<section class="slide s" id="s05">
  {brand()}{trail(1)}
  <div class="body top">
    <p class="kicker">Ato I · Dois caminhos</p>
    <h2>Doador falecido e doador vivo</h2>
    <table>
      <thead><tr><th></th><th>Falecido · Sr. J.</th><th>Vivo · Ana e Paulo</th></tr></thead>
      <tbody>
        <tr><td class="k">Transplantes renais em 2025</td><td>87%</td><td>13%</td></tr>
        <tr><td class="k">Isquemia fria</td><td>horas</td><td>minutos</td></tr>
        <tr><td class="k">Função tardia do enxerto</td><td class="gold">54% no Brasil</td><td>3,6%</td></tr>
        <tr><td class="k">Enxerto em 5 anos</td><td>70%</td><td>85%</td></tr>
      </tbody>
    </table>
    <div class="note note--vivo"><span class="lbl">Quem pode doar em vida</span>Cônjuge, companheiro ou parente até o 4º grau. Não aparentado, só com autorização judicial. Pode desistir até a cirurgia.</div>
  </div>
  {foot(5, "<b>[1]</b>RBT 2025 · <b>[5]</b>Sandes-Freitas 2021 · <b>[6]</b>Redfield 2016 · <b>[7]</b>Lei 9.434/1997, Decreto 9.175/2017")}
</section>''')

# 06 avaliação
S.append(f'''<section class="slide s has-vivo" id="s06">
  {brand()}{trail(2)}
  <div class="body wv top">
    <p class="kicker">Ato I · Avaliação</p>
    <h2>A avaliação é de toda a equipe</h2>
    <div class="chips">
      <span>Nefrologia</span><span>Cirurgia</span><span>Cardiologia</span>
      <span>Odontologia</span><span>Enfermagem</span><span>Psicologia</span>
      <span>Serviço social</span><span>Nutrição</span><span>Farmácia</span>
    </div>
    <div class="note note--warm"><span class="lbl">O que a UBS pode adiantar</span>Vacinas de vírus vivo (tríplice viral, varicela, febre amarela) pelo menos 4 semanas antes, porque depois do transplante ficam contraindicadas. Hepatite B em 4 doses dobradas. Dentista. Rastreios de câncer em dia.</div>
  </div>
  {vivo("Paulo tem avaliação própria: clínica, renal e psicológica.", "A psicologia confirma que a decisão é livre, sem pressão da família.")}
  {foot(6, "<b>[8]</b>Manual dos CRIE, 6ª ed. MS, 2023", wv=True)}
</section>''')

# 07 fila
S.append(f'''<section class="slide s has-vivo" id="s07">
  {brand()}{trail(3)}
  <div class="body wv">
    <p class="kicker">Ato I · Fila</p>
    <h2>Estar na fila é ter tarefa</h2>
    <ul class="b tight">
      <li>Grupo sanguíneo idêntico e pontos por compatibilidade HLA, em que o DR pesa mais</li>
      <li>Somam pontos o tempo de espera, a sensibilização, o diabetes, a idade abaixo de 18 e ter sido doador vivo</li>
      <li><strong>Soro a cada 90 dias</strong> no laboratório de histocompatibilidade, e de novo após transfusão</li>
      <li>Telefone atualizado. Transporte e acompanhante combinados com o serviço social</li>
    </ul>
  </div>
  {vivo("Ana não espera chamada.", "A cirurgia tem data marcada.")}
  {foot(7, "<b>[4]</b>Portaria 2.600/2009 (PRC 4), vigente. Portaria 8.041/2025 publicada, vigência adiada", wv=True)}
</section>''')

# 08 duas famílias
S.append(f'''<section class="slide s" id="s08">
  <div class="ph ph--full"><img src="images/familia.png" alt=""></div>
  {trail(4)}
  <div class="body wide">
    <p class="kicker">Ato II · O sim</p>
    <p class="closing" style="font-size:64px">A parte mais difícil da rede acontece com vocês. E quase nunca vocês ficam sabendo o que veio depois.</p>
    <div class="note note--vivo" style="max-width:24ch"><span class="lbl">No doador vivo</span>Paulo disse sim meses antes, numa consulta, com tempo para pensar e para desistir.</div>
  </div>
  {foot(8)}
</section>''')

# 09 cuidado com o doador
S.append(f'''<section class="slide s has-vivo" id="s09">
  {brand()}{trail(5)}
  <div class="body wv">
    <p class="kicker">Ato II · CIHDOTT e OPO</p>
    <h2>O cuidado com o doador é o primeiro cuidado com o receptor</h2>
    <div class="nums n2">
      <div><div class="num">54%</div><p class="num-cap">dos rins de doador falecido têm função tardia no Brasil, de 30% a 88% por centro.</p></div>
      <div><div class="num">+4%</div><p class="num-cap">de chance de função tardia por hora de isquemia fria.</p></div>
    </div>
    <div class="note note--warm"><span class="lbl">Por quê</span>Encaminhamento tardio e manutenção inadequada do doador. Metas cumpridas: 17% contra 30%.</div>
  </div>
  {vivo("As duas salas operam lado a lado.", "Isquemia fria de minutos.")}
  {foot(9, "<b>[5]</b>Sandes-Freitas, Transpl Int 2021 · <b>[9]</b>Malinoski, Am J Transplant 2013", wv=True)}
</section>''')

# 10 2h
S.append(f'''<section class="slide s" id="s10">
  {brand()}
  <div class="ph"><img src="images/caso.png" alt=""></div>
  <div class="body narrow">
    <p class="kicker">Ato II · Hora zero</p>
    <h1>2h da manhã</h1>
    <p class="lead">O telefone do Sr. J. toca.</p>
    <ul class="b airy">
      <li>Doador de 41 anos, morte encefálica</li>
      <li>Até a reperfusão: <strong>18 horas</strong> de isquemia fria</li>
    </ul>
  </div>
  {foot(10)}
</section>''')

# 11 pré
S.append(f'''<section class="slide s has-vivo" id="s11">
  {brand()}{trail(6)}
  <div class="body wv">
    <p class="kicker">Ato III · Pré-operatório</p>
    <h2>O preparo</h2>
    <ul class="b tight">
      <li>Última diálise, potássio e peso seco</li>
      <li>Jejum de 6 h para sólidos leves e 2 h para líquidos claros</li>
      <li>Pele: tricotomizador elétrico só se necessário, imediatamente antes. Nunca lâmina</li>
      <li>Acesso central pela nefrologia. Sonda vesical no centro cirúrgico</li>
    </ul>
    <div class="note note--warm"><span class="lbl">Psicologia</span>Quem é chamado às 2h depois de quatro anos de fila está com medo, não com alegria.</div>
  </div>
  {vivo("Internação programada, diálise na véspera, preparo sem pressa.", "O doador interna junto.")}
  {foot(11, "<b>[10]</b>Cuidados PO Tx Renal v1.1, ISCAL/Nefroclínica", wv=True)}
</section>''')

# 12 intra
S.append(f'''<section class="slide s" id="s12">
  {brand()}{trail(6)}
  <div class="body top">
    <p class="kicker">Ato III · Intraoperatório</p>
    <h2>O rim é acrescentado, não trocado</h2>
    <div class="fig fig--anat">
{ANAT}
    </div>
    <div class="note note--vivo" style="margin-top:6px"><span class="lbl">No doador vivo</span>Na sala ao lado, nefrectomia por videolaparoscopia, em geral do rim esquerdo.</div>
  </div>
  {foot(12, "<b>[11]</b>KDIGO doador vivo, 2017")}
</section>''')

# 13 UTI
S.append(f'''<section class="slide s has-vivo" id="s13">
  {brand()}{trail(7)}
  <div class="body wv">
    <p class="kicker">Ato III · UTI, dias 0 a 3</p>
    <h2>A enfermagem à beira do leito</h2>
    <ul class="b tight">
      <li>Sinais vitais a cada 15 min nas 3 primeiras horas, depois de hora em hora</li>
      <li>Diurese de hora em hora e balanço rigoroso</li>
      <li>Pressão arterial média (PAM) entre 80 e 110 mmHg</li>
      <li>Diurese acima de 100 ml/h: repor 70% dela na hora seguinte</li>
      <li>Enxerto palpável, sonda sem presilha, débito do dreno anotado</li>
    </ul>
  </div>
  {vivo("Ana urina 100 a 200 ml por hora desde a sala. A reposição volêmica é o trabalho.", "Paulo é paciente cirúrgico, no leito ao lado.")}
  {foot(13, "<b>[10]</b>Cuidados PO Tx Renal v1.1, ISCAL/Nefroclínica", wv=True)}
</section>''')

# 14 imunossupressão
S.append(f'''<section class="slide s" id="s14">
  {brand()}{trail(7)}
  <div class="body top">
    <p class="kicker">Ato III · Imunossupressão</p>
    <h2>As regras do serviço e um horário</h2>
    <div class="g2">
      <div>
        <div class="col-lbl">No nosso serviço</div>
        <ul class="b tight">
          <li>Indução com timoglobulina só com 5 ou mais incompatibilidades HLA</li>
          <li>Tríplice: tacrolimus, micofenolato, prednisona</li>
          <li>Sulfametoxazol-trimetoprima por 3 meses, sempre</li>
          <li>Profilaxia de citomegalovírus só com indução, ou doador positivo e receptor negativo</li>
        </ul>
      </div>
      <div>
        <div class="col-lbl">O nível do tacrolimus</div>
        <div class="hour"><div>7h30<small>colher</small></div><div>8h<small>dar a dose</small></div></div>
        <p class="t">Colhido fora dessa janela, o número é falso e a dose sai errada.</p>
      </div>
    </div>
  </div>
  {foot(14, "<b>[10]</b>Cuidados PO Tx Renal v1.1 · <b>[12]</b>KDIGO receptor, 2009")}
</section>''')

# 15 dia 2
S.append(f'''<section class="slide s has-vivo" id="s15">
  {brand()}{trail(7)}
  <div class="body wv">
    <p class="kicker">Ato III · Dia 2</p>
    <h2>O rim não urina. A família pergunta: rejeitou?</h2>
    <ul class="b tight">
      <li><strong>Função tardia do enxerto:</strong> precisar de diálise na primeira semana</li>
      <li>Esperada com isquemia longa e doador instável. Aumenta em cerca de 41% o risco de perder o enxerto em 3 anos</li>
      <li>A enfermagem muda o desfecho: PAM no alvo, volemia, diurese anotada com honestidade</li>
    </ul>
    <div class="note note--alert"><span class="lbl">Esperado não é abrupto</span>Queda súbita da diurese em quem urinava bem é trombose até prova em contrário. Doppler agora.</div>
  </div>
  {vivo("Função tardia em 3,6% dos transplantes com doador vivo.", "O Sr. J. dialisa no dia 3 e começa a urinar no dia 7.")}
  {foot(15, "<b>[13]</b>Braz J Nephrol 2026 · <b>[6]</b>Redfield 2016", wv=True)}
</section>''')

# 16 equipe
S.append(f'''<section class="slide s" id="s16">
  {brand()}{trail(8)}
  <div class="body top">
    <p class="kicker">Ato III · A equipe</p>
    <h2>A equipe que não aparece na foto</h2>
    <div class="g4">
      <div><div class="col-lbl">Farmácia</div><ul class="b">
        <li>Reconciliação na admissão e na alta</li>
        <li>Interações com o tacrolimus</li>
        <li>Pedir o remédio de alto custo antes da alta</li></ul></div>
      <div><div class="col-lbl">Nutrição</div><ul class="b">
        <li>Da dieta da diálise para a do transplantado</li>
        <li>Diabetes pós-transplante em 10% a 30% no 1º ano</li>
        <li>Nada cru ou não pasteurizado. Toranja, não</li></ul></div>
      <div><div class="col-lbl">Psicologia</div><ul class="b">
        <li>O medo do chamado</li>
        <li>A frustração do rim que não urina</li>
        <li>O rim de outra pessoa: gratidão e culpa</li>
        <li class="vv">Doador vivo: Paulo e a relação entre os irmãos</li></ul></div>
      <div><div class="col-lbl">Serviço social</div><ul class="b">
        <li>Transporte para as consultas</li>
        <li>Acesso ao remédio</li>
        <li>A volta ao trabalho</li></ul></div>
    </div>
  </div>
  {foot(16, "<b>[14]</b>PCDT Imunossupressão Tx Renal 2021 · <b>[12]</b>KDIGO 2009 · <b>[15]</b>Avery 2019")}
</section>''')

# 17 enfermaria e alta
S.append(f'''<section class="slide s has-vivo" id="s17">
  {brand()}{trail(8)}
  <div class="body wv top">
    <p class="kicker">Ato III · Enfermaria, dias 4 a 8</p>
    <h2>A alta é uma prova oral</h2>
    <p class="t">Peso diário na mesma balança, creatinina diária, sonda, dreno e acesso retirados cedo.</p>
    <div class="col-lbl" style="margin-top:34px">Ele responde sem consultar papel</div>
    <ul class="b tight">
      <li>Cada remédio, a dose e o horário</li>
      <li>O que fazer se esquecer ou vomitar a dose</li>
      <li>Por que nunca parar, nem quando estiver bem</li>
      <li>Quando ligar, e que a coleta é antes da dose</li>
    </ul>
  </div>
  {vivo("Paulo tem alta em poucos dias, antes da irmã.")}
  {foot(17, "<b>[10]</b>Cuidados PO Tx Renal v1.1", wv=True)}
</section>''')

# 18 quem cuida do quê
S.append(f'''<section class="slide s has-vivo" id="s18">
  {brand()}{trail(9)}
  <div class="body wv top">
    <p class="kicker">Ato IV · De volta para casa</p>
    <h2>Quem cuida do quê</h2>
    <div class="g2" style="gap:50px">
      <div><div class="col-lbl">Ambulatório de transplante</div><ul class="b tight">
        <li>Enxerto e imunossupressão</li>
        <li>Níveis, complicações, profilaxias</li>
        <li>Pele e gestação</li>
        <li>O doador vivo</li></ul></div>
      <div><div class="col-lbl">UBS</div><ul class="b tight">
        <li>Hipertensão, diabetes, dislipidemia, muitas nascidas do próprio remédio</li>
        <li>A porta de entrada para todo o resto</li></ul></div>
    </div>
    <div class="note note--warm"><span class="lbl">Contrarreferência</span>O telefone do ambulatório tem de estar no prontuário da UBS.</div>
  </div>
  {vivo("Paulo volta ao ambulatório pelo menos uma vez por ano, para sempre: pressão, creatinina, albuminúria.", "Na UBS, alguém com um rim só, escrito no prontuário.")}
  {foot(18, "<b>[11]</b>KDIGO doador vivo, 2017 · <b>[12]</b>KDIGO receptor, 2009", wv=True)}
</section>''')

# 19 UBS e crônicas
S.append(f'''<section class="slide s" id="s19">
  {brand()}{trail(0)}
  <div class="body top">
    <p class="kicker">Ato IV · Na UBS</p>
    <h2>Antes de prescrever: isso mexe no tacrolimus?</h2>
    <table>
      <thead><tr><th>Sobem o nível</th><th>Derrubam o nível</th><th>Agridem o enxerto</th></tr></thead>
      <tbody><tr><td>Diltiazem, verapamil, claritromicina, fluconazol</td><td>Rifampicina, fenitoína, carbamazepina</td><td>Anti-inflamatórios</td></tr></tbody>
    </table>
    <ul class="b tight" style="margin-top:22px">
      <li>Imunossupressor não é suspenso nem ajustado na UBS. Se faltar, é urgência: avisar o ambulatório no mesmo dia</li>
      <li>Vacina de vírus vivo, não: tríplice viral, varicela, febre amarela, BCG, dengue. Influenza todo ano, sim</li>
    </ul>
  </div>
  {foot(19, "<b>[12]</b>KDIGO 2009 · <b>[8]</b>Manual dos CRIE 2023")}
</section>''')

# 20 quando acionar
S.append(f'''<section class="slide s" id="s20">
  {brand()}
  <div class="ph"><img src="images/ubs.png" alt=""></div>
  <div class="body narrow">
    <p class="kicker">Ato IV · Sinais</p>
    <h2>Quando a UBS aciona o ambulatório</h2>
    <ul class="b tight">
      <li>Febre acima de 38 °C</li>
      <li>Queda da diurese ou dor no enxerto</li>
      <li>Vômito que faz perder a dose</li>
      <li>Diarreia persistente: <strong>sobe</strong> o tacrolimus</li>
      <li>Creatinina subindo, mesmo sem sintoma</li>
      <li>Qualquer remédio novo de uso contínuo</li>
    </ul>
  </div>
  {foot(20, "<b>[16]</b>Lemahieu, Am J Transplant 2005")}
</section>''')

# 21 ACS
S.append(f'''<section class="slide s" id="s21">
  {brand()}
  <div class="ph"><img src="images/acs.png" alt=""></div>
  <div class="body narrow">
    <p class="kicker">Ato IV · Na comunidade</p>
    <h1>A conversa começa em casa</h1>
    <p class="lead">45% das famílias recusaram a doação em 2025.</p>
    <ul class="b airy">
      <li>Quase sempre porque nunca ouviram a pessoa dizer que queria doar</li>
      <li>A visita domiciliar é onde essa conversa pode começar, antes do hospital</li>
    </ul>
  </div>
  {foot(21, "<b>[1]</b>RBT 2025")}
</section>''')

# 22 um ano depois
S.append(f'''<section class="slide s" id="s22">
  <div class="ph ph--full"><img src="images/domingo.png" alt=""></div>
  <div class="body wide">
    <p class="kicker">Um ano depois</p>
    <p class="closing">O Sr. J. não faz mais diálise. Ana e Paulo almoçam juntos no domingo.</p>
    <div class="rule"></div>
    <p class="t">O Sr. J. nunca vai saber o nome da família que disse sim. A lei guarda a identidade dos dois lados.</p>
  </div>
  {foot(22, "<b>[7]</b>Decreto 9.175/2017, art. 52")}
</section>''')

# 23 volta inteira
S.append(f'''<section class="slide s s--soft" id="s23">
  <div class="body top" style="top:70px; bottom:90px">
    <div class="map">
      {map_svg(1672, 800, 836, 378, 252, [("O rim do Sr. J.", 40, "#5A7075", 400), ("passou por mais de vinte mãos", 34, "#6E7F83", 300), ("antes de voltar para a UBS.", 34, "#6E7F83", 300)], who=True, fs=38, wfs=30)}
    </div>
    <p class="closing" style="max-width:none; text-align:center; font-size:60px; margin-top:6px">A maior parte delas está nesta sala.</p>
  </div>
</section>''')

# 24 referências
REFS = [
    "ABTO. Registro Brasileiro de Transplantes 2025, jan/dez. 2026.",
    "MS. Diretrizes clínicas para o cuidado ao paciente com DRC no SUS, 2014. Protocolos de encaminhamento, Nefrologia, 2015.",
    "KDIGO 2020, candidatos a transplante renal. Transplantation 2020;104(4):708-14.",
    "Portaria GM/MS 2.600/2009 (PRC 4/2017, Anexo I). Portaria GM/MS 8.041/2025.",
    "Sandes-Freitas TV et al. Transpl Int 2021. PMID <span class=\"pmid\">33742470</span>",
    "Redfield RR et al. Transpl Int 2016. PMID <span class=\"pmid\">26432507</span>",
    "Lei 9.434/1997. Decreto 9.175/2017.",
    "MS. Manual dos CRIE, 6ª ed., 2023.",
    "Malinoski DJ et al. Am J Transplant 2013. PMID <span class=\"pmid\">23406284</span>",
    "Amaral Filho G, Moro M. Cuidados PO Tx Renal v1.1. ISCAL/Nefroclínica, 2023.",
    "Lentine KL et al. KDIGO doador vivo. Transplantation 2017. PMID <span class=\"pmid\">28742762</span>",
    "KDIGO, cuidado do receptor. Am J Transplant 2009. PMID <span class=\"pmid\">19845597</span>",
    "A nova era dos tratamentos para receptores de transplante renal. Braz J Nephrol 2026;48(4).",
    "MS. PCDT Imunossupressão no Transplante Renal. Portaria Conjunta nº 1/2021.",
    "Avery RK, Michaels MG. Clin Transplant 2019. PMID <span class=\"pmid\">30844096</span>",
    "Lemahieu W et al. Am J Transplant 2005. PMID <span class=\"pmid\">15888045</span>",
]
S.append(f'''<section class="slide s" id="s24">
  {brand()}
  <div class="body top" style="top:130px">
    <h2>Referências</h2>
    <ol class="refs refs--dense">{"".join(f"<li>{r}</li>" for r in REFS)}</ol>
  </div>
  {foot(24)}
</section>''')

assert len(S) == TOTAL, len(S)
html = f'''<!DOCTYPE html>
<html lang="pt-BR">
<head>
<meta charset="UTF-8">
<title>Transplante renal: a volta inteira</title>
<link rel="stylesheet" href="lecture.css">
<style>
ol.refs.refs--dense{{ gap:16px 60px; margin-top:28px }}
.fig--anat svg{{ max-height:440px; width:auto; max-width:100% }}
li.vv{{ color:var(--sage-deep) }}
.credit{{ position:absolute; right:28px; bottom:22px; z-index:4; font-family:var(--body); font-size:22px; color:rgba(255,255,255,.85); text-shadow:0 1px 3px rgba(0,0,0,.4) }}
ol.refs.refs--dense li{{ font-size:30px; line-height:1.3; padding-left:48px }}
ol.refs.refs--dense li::before{{ font-size:30px }}
</style>
</head>
<body>
{chr(10).join(S)}
</body>
</html>
'''
(HERE / "lecture.html").write_text(html)
print("ok", len(S), "slides")
