'use client';

/**
 * O plano de devolutiva na TELA do paciente.
 *
 * Não é o PDF embutido nem o slide de 1920×1080 com zoom. O slide impresso é uma moldura fixa
 * pensada para projetor e papel; no celular ele viraria uma imagem minúscula que o paciente
 * beliscaria para ler. Aqui o mesmo conteúdo vira uma pilha de seções que se adapta à largura:
 * tipografia em `clamp`, grades que caem para uma coluna, e a régua em SVG com `viewBox` — que
 * encolhe de 1714px para a largura do telefone sem perder nada.
 *
 * As cores e a gramática são as mesmas do deck impresso, de propósito: o paciente vê na tela o
 * documento que recebe em PDF.
 */
import { cn } from '@/lib/utils';
import type { DeckSlide } from '@/lib/api/patient-plans';

/** Rampa contínua do nível 0 (pior) ao 5 (ótimo). Espelha pdfdoc.rulerRamp. */
const RAMPA = ['#B3503C', '#CD8674', '#E1C6B9', '#AFC9D5', '#5D93AC', '#0E4C6B'];

type RulerSegment = { level: number; a: number; b: number };
type RulerPoint = { value: number; text?: string; date?: string };
type Ruler = {
  code?: string;
  display?: string;
  sub?: string;
  unit?: string;
  axis?: [number, number] | number[];
  segments?: RulerSegment[];
  history?: RulerPoint[];
  note?: string;
};

const rampColor = (level: number) => RAMPA[Math.min(Math.max(level, 0), 5)];

/**
 * A régua, redesenhada para tela vertical.
 *
 * No impresso o nome fica numa coluna à esquerda da barra; num telefone isso deixaria a barra com
 * uns dois centímetros. Aqui o nome vai ACIMA e a barra ocupa a largura toda. O ponto e a seta do
 * valor anterior continuam, porque a direção da mudança é o que mais importa e é onde o leitor
 * mais erra.
 */
function PortalRuler({ ruler }: { ruler: Ruler }) {
  const axis = ruler.axis ?? [0, 1];
  const [lo, hi] = [Number(axis[0] ?? 0), Number(axis[1] ?? 1)];
  const span = hi - lo || 1;
  const pct = (v: number) => ((Math.min(Math.max(v, lo), hi) - lo) / span) * 100;

  const segs = [...(ruler.segments ?? [])].sort((x, y) => x.a - y.a);
  const hist = ruler.history ?? [];
  const atual = hist.length > 0 ? hist[hist.length - 1] : null;
  const anterior = hist.length > 1 ? hist[hist.length - 2] : null;

  return (
    <div className="py-3">
      <div className="mb-2 flex items-baseline justify-between gap-3">
        <div className="min-w-0">
          <div className="truncate text-[15px] font-semibold text-[#0A1F26]">{ruler.display}</div>
          {ruler.sub && <div className="truncate text-[13px] text-[#5A6B70]">{ruler.sub}</div>}
        </div>
        {atual && (
          <div className="shrink-0 text-right">
            <span className="text-[17px] font-semibold tabular-nums text-[#0A1F26]">
              {atual.text ?? atual.value}
            </span>
            {ruler.unit && <span className="ml-1 text-[13px] text-[#5A6B70]">{ruler.unit}</span>}
          </div>
        )}
      </div>

      <div className="relative h-6">
        {/* Cada faixa é posicionada pelo próprio valor, não empilhada em fluxo. Empilhar assume que
            as faixas cobrem o eixo sem buraco; um nível com limite ilegível é descartado na
            derivação, e aí TODAS as faixas seguintes escorregariam para a esquerda e a bolinha do
            paciente cairia sobre a cor errada. O fundo é a pior faixa, como no PDF. */}
        <div className="absolute inset-0 overflow-hidden rounded-full" style={{ background: RAMPA[0] }}>
          {segs.map((s, i) => (
            <div
              key={i}
              className="absolute inset-y-0"
              style={{
                left: `${pct(s.a)}%`,
                width: `${Math.max(pct(s.b) - pct(s.a), 0.5)}%`,
                background: rampColor(s.level),
              }}
            />
          ))}
        </div>
        {anterior && (
          <div
            className="absolute top-1/2 h-3 w-3 -translate-x-1/2 -translate-y-1/2 rounded-full border-2 border-[#0A1F26] bg-[#F5F1E8]"
            style={{ left: `${pct(anterior.value)}%` }}
            aria-hidden
          />
        )}
        {atual && (
          <div
            className="absolute top-1/2 h-4 w-4 -translate-x-1/2 -translate-y-1/2 rounded-full bg-[#0A1F26] ring-2 ring-[#F5F1E8]"
            style={{ left: `${pct(atual.value)}%` }}
            aria-hidden
          />
        )}
      </div>

      {anterior && atual && (
        <div className="mt-1 text-[12px] text-[#5A6B70]">
          antes {anterior.text ?? anterior.value} · agora {atual.text ?? atual.value}
        </div>
      )}
      {ruler.note && <div className="mt-1 text-[13px] italic text-[#5A6B70]">{ruler.note}</div>}
    </div>
  );
}

/** A legenda: a rampa inteira é a escala do exame, do pior ao ótimo. */
function RampaLegenda() {
  return (
    <div className="mt-3 flex items-center gap-2 text-[12px] uppercase tracking-wide text-[#5A6B70]">
      <span>pior</span>
      <span className="flex h-3 flex-1 overflow-hidden rounded-full">
        {RAMPA.map((c) => (
          <i key={c} className="flex-1" style={{ background: c }} />
        ))}
      </span>
      <span>ótimo</span>
    </div>
  );
}

/**
 * Texto com a ênfase que o autor escreveu, reconstruída como elementos React — não há
 * `dangerouslySetInnerHTML` aqui: o conteúdo vem do banco e não pode chegar ao DOM como HTML.
 *
 * A lista de tags tem que ser a MESMA do servidor (`pdfdoc.inlineAllowed`). Quando divergiam, um
 * `<br>` escrito num punch quebrava a linha nos dois PDFs e aparecia como o texto literal "<br>"
 * na tela do paciente.
 */
const TAGS_INLINE = /(<\/?(?:em|strong|b|i|small)>|<br\s*\/?>)/gi;

function RichText({ text, className }: { text?: string; className?: string }) {
  if (!text) return null;
  const partes = text.split(TAGS_INLINE);
  let enfase = false;
  let miudo = false;
  const nos: React.ReactNode[] = [];
  partes.forEach((p, i) => {
    if (/^<br\s*\/?>$/i.test(p)) {
      nos.push(<br key={i} />);
      return;
    }
    if (/^<\/?small>$/i.test(p)) {
      miudo = !p.startsWith('</');
      return;
    }
    if (/^<\/?(em|strong|b|i)>$/i.test(p)) {
      enfase = !p.startsWith('</');
      return;
    }
    if (!p) return;
    const cls = cn(enfase && 'font-medium not-italic text-[#8A6534]', miudo && 'text-[0.85em] opacity-80');
    nos.push(
      <span key={i} className={cls || undefined}>
        {p}
      </span>,
    );
  });
  return <p className={className}>{nos}</p>;
}

function Cartao({
  children,
  tone,
}: {
  children: React.ReactNode;
  tone?: 'bom' | 'ruim' | 'escuro';
}) {
  return (
    <div
      className={cn(
        'rounded-xl border p-4',
        tone === 'escuro'
          ? 'border-transparent bg-[#063B4F] text-[#EAE7DA]'
          : 'border-[#0A1F26]/10 bg-white/70',
        tone === 'bom' && 'border-l-4 border-l-[#0E4C6B]',
        tone === 'ruim' && 'border-l-4 border-l-[#B3503C]',
      )}
    >
      {children}
    </div>
  );
}

function Slide({ slide }: { slide: DeckSlide }) {
  const escuro = slide.variant === 'dark' || slide.variant === 'deep';
  const rulers = (slide.rulers ?? []) as Ruler[];
  const summary = slide.summary as
    | {
        cards?: { title: string; tone?: string; lines?: { name: string; sub?: string; value: string; unit?: string; ruler?: Ruler }[] }[];
        stepsTitle?: string;
        steps?: string[];
      }
    | undefined;
  const cards = (slide.cards ?? []) as { kicker?: string; body?: string; dim?: boolean; focus?: boolean }[];
  const steps = (slide.steps ?? []) as { when: string; what: string; detail?: string }[];
  const take = slide.takeaway as
    | {
        highlight?: { when?: string; name: string; obs?: string; dose?: string; unit?: string };
        groups?: { title: string; items?: { name: string; sub?: string; dose?: string }[] }[];
        note?: string;
      }
    | undefined;

  return (
    <section
      className={cn(
        'scroll-mt-16 rounded-2xl px-5 py-7 sm:px-8 sm:py-10',
        escuro ? 'bg-[#041F2A] text-[#EAE7DA]' : 'bg-[#F5F1E8] text-[#0A1F26]',
      )}
    >
      {slide.eyebrow && (
        <div
          className={cn(
            'mb-2 text-[12px] font-semibold uppercase tracking-[0.14em]',
            escuro ? 'text-[#D4A86B]' : 'text-[#B38645]',
          )}
        >
          {slide.eyebrow}
        </div>
      )}
      {slide.title && (
        // clamp: o mesmo título serve o telefone e o desktop sem virar duas versões.
        // Passa pelo RichText porque o servidor também aplica ênfase no título.
        <h2 className="font-serif text-[clamp(1.5rem,5vw,2.25rem)] leading-tight">
          <RichText text={slide.title} />
        </h2>
      )}
      {slide.lede && (
        <RichText
          text={slide.lede}
          className={cn('mt-3 text-[clamp(1rem,3.4vw,1.15rem)] leading-relaxed', escuro ? 'text-[#92B8B4]' : 'text-[#5A6B70]')}
        />
      )}

      {rulers.length > 0 && (
        <div className="mt-5 divide-y divide-[#0A1F26]/10">
          {rulers.map((r, i) => (
            <PortalRuler key={r.code ?? i} ruler={r} />
          ))}
          {slide.legend && <RampaLegenda />}
        </div>
      )}

      {summary && (
        <div className="mt-5 grid gap-4 md:grid-cols-2">
          {summary.cards?.map((c, i) => (
            <Cartao key={i} tone={c.tone === 'bom' || c.tone === 'ruim' ? c.tone : undefined}>
              <div
                className={cn(
                  'mb-3 text-[12px] font-bold uppercase tracking-[0.14em]',
                  c.tone === 'ruim' ? 'text-[#B3503C]' : 'text-[#0E4C6B]',
                )}
              >
                {c.title}
              </div>
              <div className="divide-y divide-[#0A1F26]/8">
                {c.lines?.map((ln, j) => (
                  <div key={j} className="py-2">
                    <div className="flex items-baseline justify-between gap-3">
                      <span className="text-[15px] font-semibold">{ln.name}</span>
                      <span className="shrink-0 text-[15px] font-bold tabular-nums">
                        {ln.value}
                        {ln.unit && <span className="ml-1 font-normal text-[#5A6B70]">{ln.unit}</span>}
                      </span>
                    </div>
                    {ln.sub && <div className="text-[13px] text-[#5A6B70]">{ln.sub}</div>}
                  </div>
                ))}
              </div>
            </Cartao>
          ))}
          {summary.steps && summary.steps.length > 0 && (
            <div className="md:col-span-2">
              <Cartao tone="escuro">
                <div className="mb-3 text-[12px] font-bold uppercase tracking-[0.14em] text-[#D4A86B]">
                  {summary.stepsTitle || 'O que vamos fazer'}
                </div>
                <ol className="grid gap-3 sm:grid-cols-2 lg:grid-cols-4">
                  {summary.steps.map((st, i) => (
                    <li key={i} className="flex gap-3">
                      <span className="font-serif text-2xl leading-none text-[#D4A86B]">{i + 1}</span>
                      <span className="text-[15px] leading-snug">{st}</span>
                    </li>
                  ))}
                </ol>
              </Cartao>
            </div>
          )}
        </div>
      )}

      {cards.length > 0 && (
        <div className="mt-5 grid gap-4 md:grid-cols-2">
          {cards.map((c, i) => (
            <div key={i} className={cn(c.dim && 'opacity-60')}>
              <Cartao>
                {c.kicker && (
                  <div className="mb-2 text-[12px] font-bold uppercase tracking-[0.1em] text-[#5A6B70]">
                    {c.kicker}
                  </div>
                )}
                <RichText text={c.body} className="text-[15px] leading-relaxed" />
              </Cartao>
            </div>
          ))}
        </div>
      )}

      {steps.length > 0 && (
        <ol className="mt-5 divide-y divide-[#0A1F26]/10">
          {steps.map((st, i) => (
            <li key={i} className="grid gap-1 py-3 sm:grid-cols-[180px_1fr] sm:gap-6">
              <div className="text-[12px] font-bold uppercase tracking-[0.1em] text-[#8A6534]">
                {st.when}
              </div>
              <div>
                <RichText text={st.what} className="text-[16px] leading-snug" />
                {st.detail && <div className="mt-1 text-[14px] text-[#5A6B70]">{st.detail}</div>}
              </div>
            </li>
          ))}
        </ol>
      )}

      {take && (
        <div className="mt-5 space-y-4">
          {take.highlight && (
            <Cartao tone="escuro">
              <div className="flex flex-wrap items-end justify-between gap-4">
                <div>
                  {take.highlight.when && (
                    <div className="text-[12px] font-bold uppercase tracking-[0.14em] text-[#D4A86B]">
                      {take.highlight.when}
                    </div>
                  )}
                  <div className="font-serif text-[clamp(1.35rem,4.5vw,2rem)]">{take.highlight.name}</div>
                  {take.highlight.obs && (
                    <div className="mt-1 text-[14px] text-[#92B8B4]">{take.highlight.obs}</div>
                  )}
                </div>
                <div className="text-right">
                  <div className="font-serif text-[clamp(1.75rem,7vw,3rem)] leading-none">
                    {take.highlight.dose}
                  </div>
                  {take.highlight.unit && (
                    <div className="mt-1 text-[13px] text-[#92B8B4]">{take.highlight.unit}</div>
                  )}
                </div>
              </div>
            </Cartao>
          )}
          <div className="grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
            {take.groups?.map((g, i) => (
              <Cartao key={i}>
                <div className="mb-2 border-b border-[#B38645]/35 pb-2 text-[12px] font-bold uppercase tracking-[0.14em] text-[#8A6534]">
                  {g.title}
                </div>
                <div className="divide-y divide-[#0A1F26]/8">
                  {g.items?.map((it, j) => (
                    <div key={j} className="flex items-baseline justify-between gap-3 py-2">
                      <span>
                        <span className="text-[15px] font-semibold">{it.name}</span>
                        {it.sub && <span className="block text-[13px] text-[#5A6B70]">{it.sub}</span>}
                      </span>
                      <span className="shrink-0 text-[15px] font-bold tabular-nums text-[#063B4F]">
                        {it.dose}
                      </span>
                    </div>
                  ))}
                </div>
              </Cartao>
            ))}
          </div>
          {take.note && <p className="text-[14px] leading-relaxed text-[#5A6B70]">{take.note}</p>}
        </div>
      )}

      {slide.kicker && (
        <RichText text={slide.kicker} className={cn('mt-5 text-[15px] leading-relaxed', escuro ? 'text-[#92B8B4]' : 'text-[#5A6B70]')} />
      )}
      {slide.source && (
        <RichText text={slide.source} className="mt-3 text-[13px] leading-relaxed text-[#5A6B70]" />
      )}
      {slide.punch && (
        <RichText
          text={slide.punch}
          className={cn(
            'mt-6 border-t-2 pt-4 font-serif text-[clamp(1.05rem,3.8vw,1.4rem)] leading-snug',
            escuro ? 'border-[#D4A86B]/40 text-[#EAE7DA]' : 'border-[#B38645]/35 text-[#063B4F]',
          )}
        />
      )}
    </section>
  );
}

/** O plano inteiro: uma pilha de seções, na ordem em que o médico montou. */
export function PlanDeck({ slides }: { slides: DeckSlide[] }) {
  return (
    <div className="space-y-4">
      {slides.map((s, i) => (
        <Slide key={i} slide={s} />
      ))}
    </div>
  );
}
