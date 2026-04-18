import type { Metadata } from 'next';
import Image from 'next/image';
import { setRequestLocale } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';

export const metadata: Metadata = {
  title: 'Escore Plenya — instrumento de medida do Método AGIR',
  description:
    'Mais de 800 itens em uma pontuação clara, evolutiva e personalizada. Histórico, bioquímica e genética traduzidos em uma medida única do seu estado de saúde.',
};

// ── Radar AGIR — 22 pilares distribuídos em 4 setores de 90° ──────────
// Replica o display do EMR: vértices por pilar + anel externo colorido por letra.
const RADAR_CX = 200;
const RADAR_CY = 200;
const RADAR_MAX = 150;

const letterColors: Record<'A' | 'G' | 'I' | 'R', string> = {
  A: '#92b8b4', // sage
  G: '#b38645', // gold
  I: '#caa56b', // gold suave
  R: '#417e8e', // ocean
};

const radarPillars: { letter: 'A' | 'G' | 'I' | 'R'; angle: number; score: number }[] = [
  // A — top quadrant (-45° → +45°), 4 pilares
  { letter: 'A', angle: -33.75, score: 82 },
  { letter: 'A', angle: -11.25, score: 78 },
  { letter: 'A', angle:  11.25, score: 86 },
  { letter: 'A', angle:  33.75, score: 74 },
  // G — right quadrant (45° → 135°), 10 pilares
  { letter: 'G', angle:  49.5, score: 76 },
  { letter: 'G', angle:  58.5, score: 70 },
  { letter: 'G', angle:  67.5, score: 80 },
  { letter: 'G', angle:  76.5, score: 72 },
  { letter: 'G', angle:  85.5, score: 84 },
  { letter: 'G', angle:  94.5, score: 78 },
  { letter: 'G', angle: 103.5, score: 71 },
  { letter: 'G', angle: 112.5, score: 82 },
  { letter: 'G', angle: 121.5, score: 76 },
  { letter: 'G', angle: 130.5, score: 73 },
  // I — bottom quadrant (135° → 225°), 5 pilares
  { letter: 'I', angle: 144, score: 68 },
  { letter: 'I', angle: 162, score: 72 },
  { letter: 'I', angle: 180, score: 75 },
  { letter: 'I', angle: 198, score: 70 },
  { letter: 'I', angle: 216, score: 74 },
  // R — left quadrant (225° → 315°), 3 pilares
  { letter: 'R', angle: 240, score: 84 },
  { letter: 'R', angle: 270, score: 88 },
  { letter: 'R', angle: 300, score: 80 },
];

// Helper: angulo em graus (0 = norte, sentido horário) → coordenada cartesiana
function polar(angleDeg: number, radius: number) {
  const rad = ((angleDeg - 90) * Math.PI) / 180;
  return {
    x: RADAR_CX + radius * Math.cos(rad),
    y: RADAR_CY + radius * Math.sin(rad),
  };
}

const radarPoints = radarPillars.map((p) => {
  const r = (p.score / 100) * RADAR_MAX;
  return { ...p, ...polar(p.angle, r) };
});

const polygonStr = radarPoints.map((p) => `${p.x.toFixed(1)},${p.y.toFixed(1)}`).join(' ');

// Arco externo colorido (uma path por letra)
const ARC_RADIUS = RADAR_MAX + 18;
function arcPath(startAngle: number, endAngle: number, radius: number) {
  const s = polar(startAngle, radius);
  const e = polar(endAngle, radius);
  return `M ${s.x.toFixed(1)} ${s.y.toFixed(1)} A ${radius} ${radius} 0 0 1 ${e.x.toFixed(1)} ${e.y.toFixed(1)}`;
}

const letterArcs: { letter: 'A' | 'G' | 'I' | 'R'; start: number; end: number; midAngle: number }[] = [
  { letter: 'A', start: -45, end:  45, midAngle:   0 },
  { letter: 'G', start:  45, end: 135, midAngle:  90 },
  { letter: 'I', start: 135, end: 225, midAngle: 180 },
  { letter: 'R', start: 225, end: 315, midAngle: 270 },
];

const LETTER_LABEL_RADIUS = ARC_RADIUS + 18;

const steps = [
  {
    n: '01',
    title: 'Avaliação ampla',
    body:
      'Mais de 800 itens são levantados — do histórico pré-natal e familiar aos exames laboratoriais e marcadores genéticos. Não há pergunta solta: cada item alimenta uma decisão clínica.',
  },
  {
    n: '02',
    title: 'Organização AGIR',
    body:
      'Cada item entra em um pilar clínico do Método AGIR — agrupado em uma das quatro letras: Alimentação e Atividade Física, Gestão Metabólica, Integração Mente-Corpo, Ritmo Circadiano.',
  },
  {
    n: '03',
    title: 'Pontuação que evolui',
    body:
      'Você recebe uma pontuação global e uma por letra. A cada reavaliação, a curva mostra o que melhorou, o que estagnou e onde a próxima intervenção precisa entrar.',
  },
];

const lensColumns = [
  {
    title: 'Histórico',
    body:
      'Pré-natal, infância, adolescência, vida adulta. Doenças crônicas, cirurgias, medicamentos, vícios. Histórico familiar em três gerações.',
  },
  {
    title: 'Bioquímica',
    body:
      'Hormônios completos (testosterona livre, estradiol, tireoide, cortisol), metabolismo (insulina, HOMA-IR, ApoB), inflamação (PCR-us, IL-6), vitaminas e minerais, e exames de imagem como CAC e densitometria.',
  },
  {
    title: 'Genética',
    body:
      'Cerca de 80 variantes integradas ao pilar clínico correspondente — APOE no risco cardiovascular, MTHFR no controle glicêmico, ALDH2 na função hepática, ACTN3 na prescrição de exercícios.',
  },
];

const tiers = [
  {
    title: 'Versão Completa',
    desc: 'Aplicada pela equipe Plenya durante o acompanhamento. Avaliação clínica, laboratorial, comportamental e funcional integrais. Gera relatório detalhado com metas e plano personalizado.',
  },
  {
    title: 'Versão Intermediária',
    desc: 'Disponível após a primeira consulta. Recorte focado nas metas prioritárias identificadas pela equipe.',
  },
  {
    title: 'Versão Light',
    desc: 'Disponível online em breve. Permite começar a entender sua saúde de forma acessível, sem necessidade de consulta prévia.',
  },
];

export default async function ScorePage({ params }: { params: Promise<{ locale: string }> }) {
  const { locale } = await params;
  setRequestLocale(locale);

  return (
    <>
      {/* HERO */}
      <section className="bg-petrol text-cream">
        <div className="site-container pt-32 pb-24 md:pt-40 md:pb-32">
          <p className="label-upper text-gold mb-6">Instrumento</p>
          <h1 className="heading-hero text-[clamp(2.5rem,6vw,5rem)] text-cream max-w-2xl">
            Escore Plenya
          </h1>
          <p className="text-cream/75 text-xl md:text-2xl mt-8 max-w-2xl leading-relaxed font-light">
            O instrumento de medida do Método AGIR. Mais de 800 itens em uma
            pontuação clara, evolutiva e personalizada do seu estado de saúde.
          </p>
        </div>
      </section>

      {/* O QUE É — texto âncora curto */}
      <section className="bg-cream">
        <div className="site-container section">
          <div className="max-w-3xl space-y-6">
            <p className="label-upper text-gold">O que é</p>
            <p className="heading-section text-petrol text-2xl md:text-4xl leading-tight">
              O Escore Plenya não substitui um exame. Substitui o vácuo entre eles.
            </p>
            <p className="text-petrol/80 text-lg leading-relaxed">
              É a medida que o cuidado precisava ter para deixar de ser uma sequência
              desconexa de consultas e virar um processo com direção e progresso visível.
              Onde antes havia laudos isolados, agora há uma pontuação que sintetiza —
              e que evolui a cada reavaliação.
            </p>
          </div>
        </div>
      </section>

      {/* COMO FUNCIONA — 3 etapas */}
      <section className="bg-paper">
        <div className="site-container section">
          <div className="max-w-3xl mb-14 space-y-4">
            <p className="label-upper text-gold">Como funciona</p>
            <h2 className="heading-section text-petrol text-3xl md:text-4xl">
              Três etapas. Uma pontuação que faz sentido.
            </h2>
          </div>

          <ol className="grid md:grid-cols-3 gap-10 max-w-5xl">
            {steps.map((s) => (
              <li key={s.n} className="space-y-4">
                <p className="heading-section text-gold text-5xl leading-none">{s.n}</p>
                <h3 className="heading-section text-petrol text-xl">{s.title}</h3>
                <p className="text-petrol/75 leading-relaxed">{s.body}</p>
              </li>
            ))}
          </ol>
        </div>
      </section>

      {/* PONTUAÇÃO — radar AGIR como peça visual central */}
      <section className="bg-cream">
        <div className="site-container section grid lg:grid-cols-[1fr_auto] gap-16 items-center">
          <div className="space-y-6 max-w-xl">
            <p className="label-upper text-gold">Pontuação</p>
            <h2 className="heading-section text-petrol text-3xl md:text-5xl leading-tight">
              Não é uma nota. <em className="not-italic text-gold">É um mapa.</em>
            </h2>
            <p className="text-petrol/80 text-lg leading-relaxed">
              Você recebe uma pontuação global e quatro pontuações por letra do AGIR.
              Juntas, elas formam uma constelação — onde está concentrada a força,
              onde mora a lacuna.
            </p>
            <p className="text-petrol/70 leading-relaxed">
              Quando o radar é simétrico, o cuidado está distribuído. Quando uma ponta
              afunda, é ali que a próxima intervenção entra.
            </p>
          </div>

          {/* Radar AGIR — 22 pilares com anéis externos coloridos por letra */}
          <figure className="flex flex-col items-center gap-5">
            <svg viewBox="0 0 400 400" className="w-80 h-80 md:w-[26rem] md:h-[26rem]" aria-label="Exemplo de Escore Plenya: 22 pilares organizados em A 80 · G 76 · I 72 · R 84, total 78">
              {/* Anéis concêntricos de fundo — 25/50/75/100 */}
              {[37.5, 75, 112.5, 150].map((r) => (
                <circle key={r} cx={RADAR_CX} cy={RADAR_CY} r={r} fill="none" stroke="#063b4f" strokeOpacity="0.07" strokeWidth="1" />
              ))}

              {/* Eixos cardinais sutis */}
              <line x1={RADAR_CX} y1={RADAR_CY - RADAR_MAX} x2={RADAR_CX} y2={RADAR_CY + RADAR_MAX} stroke="#063b4f" strokeOpacity="0.06" strokeWidth="1" />
              <line x1={RADAR_CX - RADAR_MAX} y1={RADAR_CY} x2={RADAR_CX + RADAR_MAX} y2={RADAR_CY} stroke="#063b4f" strokeOpacity="0.06" strokeWidth="1" />

              {/* Polígono ligando os 22 vértices (fill suave gold) */}
              <polygon
                points={polygonStr}
                fill="#b38645"
                fillOpacity="0.16"
                stroke="#b38645"
                strokeOpacity="0.85"
                strokeWidth="1.5"
                strokeLinejoin="round"
              />

              {/* Pontos por pilar — cor do setor (letra) */}
              {radarPoints.map((p, i) => (
                <circle key={i} cx={p.x} cy={p.y} r="3.5" fill={letterColors[p.letter]} stroke="#fbfaf6" strokeWidth="1.2" />
              ))}

              {/* Anel externo colorido — uma path por letra */}
              {letterArcs.map((arc) => (
                <path
                  key={arc.letter}
                  d={arcPath(arc.start + 2, arc.end - 2, ARC_RADIUS)}
                  fill="none"
                  stroke={letterColors[arc.letter]}
                  strokeWidth="7"
                  strokeLinecap="round"
                  opacity="0.9"
                />
              ))}

              {/* Labels A G I R nos pontos cardinais externos */}
              {letterArcs.map((arc) => {
                const p = polar(arc.midAngle, LETTER_LABEL_RADIUS);
                return (
                  <text
                    key={`label-${arc.letter}`}
                    x={p.x}
                    y={p.y + 8}
                    textAnchor="middle"
                    fontFamily="'Cormorant Garamond', serif"
                    fontSize="26"
                    fontWeight="500"
                    fill="#063b4f"
                  >
                    {arc.letter}
                  </text>
                );
              })}

              {/* Score global no centro */}
              <circle cx={RADAR_CX} cy={RADAR_CY} r="34" fill="#fbfaf6" stroke="#063b4f" strokeOpacity="0.18" strokeWidth="1.5" />
              <text x={RADAR_CX} y={RADAR_CY + 11} textAnchor="middle" fontFamily="'Cormorant Garamond', serif" fontSize="34" fill="#063b4f" letterSpacing="-1">78</text>
            </svg>

            {/* Legenda de scores por letra */}
            <div className="flex gap-4 md:gap-6 font-mono text-[11px] uppercase tracking-[0.2em]">
              {[
                { l: 'A', score: 80 },
                { l: 'G', score: 76 },
                { l: 'I', score: 72 },
                { l: 'R', score: 84 },
              ].map((s) => (
                <span key={s.l} className="flex items-center gap-2 text-petrol/70">
                  <span className="w-2 h-2 rounded-full" style={{ background: letterColors[s.l as 'A' | 'G' | 'I' | 'R'] }} />
                  <span>{s.l} {s.score}</span>
                </span>
              ))}
            </div>
            <p className="label-upper text-petrol/40 text-center text-[10px]">
              Exemplo · 22 pilares · escala 0–100
            </p>
          </figure>
        </div>
      </section>

      {/* FAIXAS — legenda semântica */}
      <section className="bg-paper border-y border-petrol/10">
        <div className="site-container py-16 md:py-20">
          <div className="max-w-3xl mb-10 space-y-3">
            <p className="label-upper text-gold">Como ler</p>
            <h2 className="heading-section text-petrol text-2xl md:text-3xl">
              Cinco faixas. Direção, não punição.
            </h2>
          </div>

          <div className="grid sm:grid-cols-2 lg:grid-cols-5 gap-3 md:gap-4 max-w-5xl">
            {[
              { label: 'Crítico',  range: '0–30',   color: 'bg-[#c1542d]' },
              { label: 'Atenção',  range: '31–50',  color: 'bg-[#d89345]' },
              { label: 'Regular',  range: '51–70',  color: 'bg-[#caa54a]' },
              { label: 'Bom',      range: '71–85',  color: 'bg-[#7ba07a]' },
              { label: 'Ótimo',    range: '86–100', color: 'bg-[#3f7d6e]' },
            ].map((f) => (
              <div key={f.label} className="flex items-center gap-3 p-4 border border-petrol/10 rounded-sm">
                <span className={`w-3 h-10 rounded-sm ${f.color} flex-shrink-0`} aria-hidden />
                <div>
                  <p className="heading-section text-petrol text-base leading-tight">{f.label}</p>
                  <p className="font-mono text-xs text-petrol/50 mt-0.5">{f.range}</p>
                </div>
              </div>
            ))}
          </div>

          <p className="text-petrol/60 leading-relaxed mt-8 max-w-3xl">
            A faixa não é veredito — é o ponto de partida da próxima conversa clínica.
            Um <strong className="text-petrol">71% em queda</strong> de 85% conta uma história
            diferente de um <strong className="text-petrol">71% subindo</strong> de 60%.
          </p>
        </div>
      </section>

      {/* EVOLUÇÃO — gráfico simbólico */}
      <section className="bg-cream">
        <div className="site-container section grid lg:grid-cols-[1fr_auto] gap-16 items-center">
          <div className="space-y-6 max-w-xl">
            <p className="label-upper text-gold">Evolução</p>
            <h2 className="heading-section text-petrol text-3xl md:text-4xl">
              A curva que mostra se o cuidado está funcionando.
            </h2>
            <p className="text-petrol/80 text-lg leading-relaxed">
              A cada reavaliação, um novo ponto entra na curva. O médico, a equipe e
              você veem o mesmo gráfico: progresso real, estagnação ou queda — sem
              espaço para opinião.
            </p>
          </div>

          <figure className="flex flex-col items-center gap-3">
            <svg viewBox="0 0 420 220" className="w-full max-w-md md:max-w-lg" aria-label="Exemplo de curva de evolução do Escore Plenya">
              {/* Grid horizontal */}
              {[40, 90, 140, 190].map((y) => (
                <line key={y} x1="40" y1={y} x2="400" y2={y} stroke="#063b4f" strokeOpacity="0.06" strokeWidth="1" />
              ))}
              {/* Eixo Y labels */}
              {[
                { y: 40, v: '100' },
                { y: 90, v: '75' },
                { y: 140, v: '50' },
                { y: 190, v: '25' },
              ].map((t) => (
                <text key={t.v} x="32" y={t.y + 4} textAnchor="end" fontFamily="monospace" fontSize="10" fill="#063b4f" fillOpacity="0.4">{t.v}</text>
              ))}
              {/* Linha de tendência exemplo */}
              <polyline
                points="60,160 140,140 220,110 300,85 380,68"
                fill="none"
                stroke="#b38645"
                strokeWidth="2.5"
                strokeLinecap="round"
                strokeLinejoin="round"
              />
              {/* Pontos */}
              {[
                { x: 60, y: 160 },
                { x: 140, y: 140 },
                { x: 220, y: 110 },
                { x: 300, y: 85 },
                { x: 380, y: 68 },
              ].map((p, i) => (
                <circle key={i} cx={p.x} cy={p.y} r="4" fill="#063b4f" stroke="#fbfaf6" strokeWidth="2" />
              ))}
              {/* Eixo X labels */}
              {['T0', '3m', '6m', '9m', '12m'].map((label, i) => (
                <text key={label} x={60 + i * 80} y="210" textAnchor="middle" fontFamily="monospace" fontSize="10" fill="#063b4f" fillOpacity="0.4">{label}</text>
              ))}
            </svg>
            <p className="label-upper text-petrol/40 text-center text-[10px]">
              Exemplo · 12 meses de acompanhamento
            </p>
          </figure>
        </div>
      </section>

      {/* O QUE O ESCORE ENXERGA — 3 lentes */}
      <section className="bg-paper">
        <div className="site-container section">
          <div className="max-w-3xl mb-12 space-y-4">
            <p className="label-upper text-gold">O que o Escore enxerga</p>
            <h2 className="heading-section text-petrol text-3xl md:text-4xl">
              Três lentes sobre uma vida inteira.
            </h2>
            <p className="text-petrol/70 leading-relaxed">
              Para cada pilar do Método AGIR, o Escore combina três dimensões de evidência —
              não apenas o exame do dia.
            </p>
          </div>

          <div className="grid md:grid-cols-3 gap-10 max-w-5xl">
            {lensColumns.map((col) => (
              <div key={col.title} className="space-y-4 border-t-2 border-gold pt-6">
                <h3 className="heading-section text-petrol text-xl">{col.title}</h3>
                <p className="text-petrol/75 leading-relaxed">{col.body}</p>
              </div>
            ))}
          </div>

          <p className="label-upper text-petrol/50 mt-12 text-xs max-w-3xl">
            Lista parcial — a versão completa cobre 800+ marcadores específicos por caso.
          </p>
        </div>
      </section>

      {/* IMAGEM ÂNCORA */}
      <section className="bg-cream">
        <div className="site-container">
          <div className="relative aspect-[16/7] overflow-hidden">
            <Image
              src="/images/hero-score.jpg"
              alt="Pausa, reflexão e autoavaliação"
              fill
              className="object-cover"
              sizes="100vw"
            />
          </div>
        </div>
      </section>

      {/* VERSÕES */}
      <section className="bg-cream">
        <div className="site-container section">
          <div className="max-w-3xl mb-12 space-y-4">
            <p className="label-upper text-gold">Versões</p>
            <h2 className="heading-section text-petrol text-3xl md:text-4xl">
              Três níveis de profundidade.
            </h2>
          </div>
          <div className="grid md:grid-cols-3 gap-8">
            {tiers.map((tier, i) => (
              <div key={tier.title} className="border-t border-petrol/15 pt-8 space-y-4">
                <span className="label-upper text-petrol/40">0{i + 1}</span>
                <h3 className="heading-section text-petrol text-2xl">{tier.title}</h3>
                <p className="text-petrol/70 leading-relaxed">{tier.desc}</p>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* SAÍDA — link único pro método */}
      <section className="bg-petrol text-cream">
        <div className="site-container section text-center space-y-8">
          <div className="max-w-2xl mx-auto space-y-4">
            <p className="label-upper text-gold">Próximo passo</p>
            <h2 className="heading-section text-cream text-3xl md:text-4xl">
              O Escore mostra onde você está. <br className="hidden md:block" />
              O Método AGIR é como atuamos a partir daí.
            </h2>
          </div>
          <Link href="/metodo-agir" className="btn-gold">
            Conhecer o Método AGIR
          </Link>
        </div>
      </section>
    </>
  );
}
