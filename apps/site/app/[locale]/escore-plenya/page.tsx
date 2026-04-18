import type { Metadata } from 'next';
import Image from 'next/image';
import { setRequestLocale } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';

export const metadata: Metadata = {
  title: 'Escore Plenya',
  description: 'O diagnóstico estruturado que mostra onde sua saúde está — 800+ itens organizados em 13 pilares clínicos dentro do Método AGIR.',
};

const tiers = [
  {
    title: 'Versão Completa',
    desc: 'Aplicada pela equipe Plenya durante o acompanhamento. Inclui avaliação clínica, laboratorial, comportamental e funcional. Gera relatório detalhado com metas e plano de cuidado personalizado.',
  },
  {
    title: 'Versão Intermediária',
    desc: 'Disponível após a primeira consulta. Recorte focado nas metas prioritárias identificadas pela equipe médica.',
  },
  {
    title: 'Versão Light',
    desc: 'Disponível online em breve. Permite começar a entender sua saúde de forma acessível, sem necessidade de consulta prévia.',
  },
];

const agirStructure = [
  {
    code: 'A',
    name: 'Alimentação e Atividade Física',
    color: 'bg-gold/10 border-gold/30',
    pillars: ['Avaliação Nutricional', 'Prescrição de Exercícios', 'Composição Corporal'],
  },
  {
    code: 'G',
    name: 'Gestão Metabólica',
    color: 'bg-ocean/10 border-ocean/30',
    pillars: ['Controle Glicêmico', 'Perfil Lipídico', 'Função Hepática', 'Função Renal'],
  },
  {
    code: 'I',
    name: 'Integração Mente-Corpo',
    color: 'bg-sage/10 border-sage/30',
    pillars: ['Avaliação Psicológica', 'Técnicas de Relaxamento', 'Função Cognitiva'],
  },
  {
    code: 'R',
    name: 'Ritmo Circadiano',
    color: 'bg-petrol/5 border-petrol/20',
    pillars: ['Qualidade do Sono', 'Cronobiologia', 'Exposição à Luz'],
  },
];

const labCategories = [
  { title: 'Hematologia', items: 'Hemograma completo, reticulócitos, ferritina, B12, folato, imunoglobulinas (IgA, IgG, IgM, IgE).' },
  { title: 'Metabolismo', items: 'Glicemia, insulina, HOMA-IR, HbA1c, peptídeo C, leptina, IGF-1 ajustado por faixa etária.' },
  { title: 'Lipídios e cardiovascular', items: 'Perfil lipídico completo, ApoB, lipoproteína(a), LDL oxidada, NT-proBNP, PCR ultrassensível, IL-6.' },
  { title: 'Hormônios', items: 'Painel completo: testosterona livre, SHBG, estradiol, progesterona, DHEA-S, cortisol, LH, FSH, prolactina, tireoide, PTH, PSA.' },
  { title: 'Vitaminas e minerais', items: 'Vitamina D, B12, ferro, magnésio sérico e RBC, selênio, manganês, zinco, ômega-3.' },
  { title: 'Função renal e urinária', items: 'Microalbuminúria, ureia, creatinina, sódio, potássio, urinálise completa com sedimento.' },
  { title: 'Toxicidade', items: 'Mercúrio e outros metais pesados quando indicado.' },
  { title: 'Imagem', items: 'CAC (escore de cálcio coronariano), bioimpedância e outros exames conforme indicação clínica.' },
];

const geneticPanels = [
  { title: 'Metabolismo', body: 'Risco de diabetes (TCF7L2, CDKAL1, KCNJ11), obesidade (FTO, MC4R, LEPR), resposta à vitamina D (VDR), conversão de ômega-3 (FADS1/2), tolerância à lactose (MCM6), homocisteína (MTHFR).' },
  { title: 'Cardiovascular', body: 'Hipertensão (ACE, AGT, GNB3, NOS3), perfil lipídico (PCSK9, APOA5, LDLR), risco de Alzheimer e dislipidemia (APOE).' },
  { title: 'Neurodegeneração', body: 'Alzheimer familial (PSEN1/2, APP), Parkinson (LRRK2, SNCA, PARK2/7, PINK1), demências frontotemporais (C9orf72, GRN, MAPT).' },
  { title: 'Detoxificação', body: 'Metabolismo de álcool (ADH1B, ALDH2), cafeína (CYP1A2), nicotina (CYP2A6), glutationa (GPX1, GST), antioxidantes (SOD2, CAT).' },
  { title: 'Imunidade', body: 'Resposta inflamatória (TNF, IL1B, IL6), proteína C reativa (CRP), risco de doença celíaca (HLA-DQ2/DQ8).' },
  { title: 'Performance', body: 'Tipo de fibra muscular (ACTN3), risco de lesão tendínea (COL5A1), densidade óssea (COL1A1, ESR1).' },
  { title: 'Outros', body: 'Marcadores adicionais conforme indicação clínica e suspeitas específicas.' },
];

export default async function ScorePage({ params }: { params: Promise<{ locale: string }> }) {
  const { locale } = await params;
  setRequestLocale(locale);

  return (
    <>
      <section className="bg-petrol text-cream">
        <div className="site-container pt-32 pb-24 md:pt-40 md:pb-32">
          <p className="label-upper text-gold mb-6">Diagnóstico</p>
          <h1 className="heading-hero text-[clamp(2.5rem,6vw,5rem)] text-cream max-w-2xl">
            Escore Plenya
          </h1>
          <p className="text-cream/70 text-lg mt-6 max-w-2xl">
            <strong className="text-cream">Mais de 800 itens avaliados</strong> — organizados nos
            13 pilares clínicos do Método AGIR. O diagnóstico estruturado que mostra onde sua
            saúde está hoje e define para onde ir.
          </p>
        </div>
      </section>

      {/* Escore x AGIR — diferenciação conceitual */}
      <section className="bg-cream">
        <div className="site-container section">
          <div className="grid md:grid-cols-2 gap-10 max-w-4xl">
            <div className="space-y-3 border-l-4 border-gold pl-6">
              <p className="label-upper text-gold">Escore Plenya</p>
              <h2 className="heading-section text-petrol text-2xl md:text-3xl">Onde você está.</h2>
              <p className="text-petrol/75 leading-relaxed">
                A ferramenta de diagnóstico. Mapeia 800+ itens ao longo de toda sua trajetória —
                do histórico familiar ao DNA, do sono à bioquímica — e traduz em uma pontuação
                que mostra o estado atual e o quanto você pode evoluir.
              </p>
            </div>
            <div className="space-y-3 border-l-4 border-petrol/30 pl-6">
              <p className="label-upper text-petrol/60">Método AGIR</p>
              <h2 className="heading-section text-petrol text-2xl md:text-3xl">Para onde vamos.</h2>
              <p className="text-petrol/75 leading-relaxed">
                O framework de atuação. Toma o diagnóstico do Escore e organiza o cuidado em
                quatro frentes interdependentes — Alimentação e Atividade Física, Gestão
                Metabólica, Integração Mente-Corpo e Ritmo Circadiano.
              </p>
            </div>
          </div>
        </div>
      </section>

      {/* Como funciona — texto + foto */}
      <section className="bg-paper">
        <div className="site-container section grid lg:grid-cols-2 gap-16 items-center">
          <div className="space-y-8">
            <p className="label-upper text-gold">Como funciona</p>
            <p className="text-petrol/80 text-lg leading-relaxed">
              Cada um dos 800+ itens pertence a um dos 13 pilares clínicos, e cada pilar a uma
              das quatro letras do AGIR. O Escore produz uma pontuação global <em>e</em> uma
              pontuação por letra — A, G, I, R — revelando em qual frente você está mais forte e
              qual precisa de atenção.
            </p>
            <p className="text-petrol/80 text-lg leading-relaxed">
              A cada reavaliação, o escore reflete as mudanças reais no seu corpo e nos seus
              hábitos. Não é uma nota de aprovação: é um mapa de onde você está e para onde pode
              ir — atualizado em tempo real ao longo do acompanhamento.
            </p>
          </div>

          <div className="relative aspect-[3/4] overflow-hidden">
            <Image
              src="/images/hero-score.jpg"
              alt="Pausa, reflexão e autoavaliação"
              fill
              className="object-cover"
              sizes="(min-width: 1024px) 540px, 100vw"
            />
          </div>
        </div>
      </section>

      {/* 13 pilares organizados em AGIR */}
      <section className="bg-cream">
        <div className="site-container section">
          <div className="max-w-3xl mb-12 space-y-4">
            <p className="label-upper text-gold">13 pilares em 4 letras</p>
            <h2 className="heading-section text-petrol text-3xl md:text-4xl">
              Os pilares que o Escore mede.
            </h2>
            <p className="text-petrol/70 leading-relaxed">
              Cada item avaliado se encaixa em um pilar específico, e cada pilar em uma das letras
              do AGIR. Você recebe um score por pilar, por letra e no total — vendo com precisão
              onde estão as forças e as lacunas.
            </p>
          </div>

          <div className="grid md:grid-cols-2 lg:grid-cols-4 gap-6">
            {agirStructure.map((letter) => (
              <div key={letter.code} className={`p-6 border rounded-sm ${letter.color} space-y-4`}>
                <div className="space-y-1">
                  <span className="heading-section text-petrol text-5xl block leading-none">{letter.code}</span>
                  <h3 className="heading-section text-petrol text-base pt-3">{letter.name}</h3>
                </div>
                <ul className="space-y-2 pt-2 border-t border-petrol/10">
                  {letter.pillars.map((p) => (
                    <li key={p} className="text-petrol/75 text-sm leading-snug flex gap-2">
                      <span className="text-gold">—</span>
                      <span>{p}</span>
                    </li>
                  ))}
                </ul>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* Painel laboratorial */}
      <section className="bg-paper">
        <div className="site-container section">
          <div className="max-w-3xl mb-12 space-y-4">
            <p className="label-upper text-gold">Painel laboratorial</p>
            <h2 className="heading-section text-petrol text-3xl md:text-4xl">
              Mais de 150 marcadores alimentam o Escore.
            </h2>
            <p className="text-petrol/70 leading-relaxed">
              Não é o painel básico do check-up de plano. São biomarcadores que revelam o que
              está acontecendo dez anos antes do diagnóstico — distribuídos pelos pilares do AGIR.
            </p>
          </div>

          <dl className="grid md:grid-cols-2 gap-x-12 gap-y-8 max-w-5xl">
            {labCategories.map((cat) => (
              <div key={cat.title} className="space-y-2 border-l-2 border-gold pl-5">
                <dt className="heading-section text-petrol text-lg">{cat.title}</dt>
                <dd className="text-petrol/70 leading-relaxed">{cat.items}</dd>
              </div>
            ))}
          </dl>

          <p className="label-upper text-petrol/50 mt-12 max-w-3xl">
            Lista parcial · marcadores adicionados conforme indicação clínica.
          </p>
        </div>
      </section>

      {/* Painel genético */}
      <section className="bg-petrol text-cream">
        <div className="site-container section">
          <div className="max-w-3xl mb-12 space-y-4">
            <p className="label-upper text-gold">Painel genético</p>
            <h2 className="heading-section text-cream text-3xl md:text-4xl">
              Sete painéis. Cerca de 80 variantes mapeadas.
            </h2>
            <p className="text-cream/70 leading-relaxed">
              A genética não é destino — é direção. Saber suas predisposições muda como
              monitoramos, prevenimos e intervimos em cada pilar do AGIR.
            </p>
          </div>

          <dl className="grid md:grid-cols-2 gap-x-12 gap-y-10 max-w-5xl">
            {geneticPanels.map((panel) => (
              <div key={panel.title} className="space-y-2">
                <dt className="heading-section text-gold text-lg">{panel.title}</dt>
                <dd className="text-cream/70 leading-relaxed text-sm">{panel.body}</dd>
              </div>
            ))}
          </dl>
        </div>
      </section>

      {/* Pontuação exemplo */}
      <section className="bg-cream-100">
        <div className="site-container section flex flex-col items-center gap-4">
          <svg viewBox="0 0 300 300" className="w-64 h-64 md:w-80 md:h-80" aria-label="Exemplo de pontuação do Escore Plenya: 78">
            <circle cx="150" cy="150" r="140" stroke="#063b4f" strokeOpacity="0.08" strokeWidth="1" fill="none" />
            <circle cx="150" cy="150" r="110" stroke="#063b4f" strokeOpacity="0.12" strokeWidth="1" fill="none" />
            <circle cx="150" cy="150" r="80" stroke="#063b4f" strokeOpacity="0.16" strokeWidth="1" fill="none" />
            <path d="M 150 10 A 140 140 0 1 1 40 260" stroke="#b38645" strokeWidth="2" fill="none" strokeLinecap="round" />
            <text x="150" y="160" textAnchor="middle" fontFamily="'Cormorant Garamond', serif" fontSize="56" fill="#063b4f" letterSpacing="-2">78</text>
          </svg>
          <p className="label-upper text-petrol/50 text-center">
            Exemplo de pontuação global · escala 0–100
          </p>
        </div>
      </section>

      {/* Versões do escore */}
      <section className="bg-paper">
        <div className="site-container section">
          <p className="label-upper text-gold mb-10">Versões do Escore</p>
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

      {/* Cross-links */}
      <section className="bg-cream">
        <div className="site-container section grid md:grid-cols-2 gap-8">
          <Link href="/metodo-agir" className="border-t border-petrol/15 pt-8 space-y-3 group">
            <p className="label-upper text-gold">Método</p>
            <p className="heading-section text-petrol text-xl group-hover:text-gold transition">Como agimos a partir do Escore →</p>
          </Link>
          <Link href="/planos" className="border-t border-petrol/15 pt-8 space-y-3 group">
            <p className="label-upper text-gold">Planos</p>
            <p className="heading-section text-petrol text-xl group-hover:text-gold transition">Comece seu acompanhamento →</p>
          </Link>
        </div>
      </section>
    </>
  );
}
