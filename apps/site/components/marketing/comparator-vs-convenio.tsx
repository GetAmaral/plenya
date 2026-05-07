import { Check, X } from 'lucide-react';
import { getLocale } from 'next-intl/server';
import { defaultLocale, isLocale, type Locale } from '@/lib/i18n/config';

type Row = { trait: string; plenya: string | true; convenio: string | false };

const rowsPt: Row[] = [
  { trait: 'Tempo de consulta', plenya: '60–90 min', convenio: '8–15 min' },
  { trait: 'Leitura funcional dos exames', plenya: true, convenio: false },
  { trait: 'Equipe multidisciplinar discutindo o seu caso', plenya: true, convenio: false },
  { trait: 'Plano por escrito, com metas e revisões', plenya: true, convenio: false },
  { trait: 'Escore Plenya — pontuação evolutiva', plenya: true, convenio: false },
  { trait: 'Reavaliação programada (3, 6, 12 meses)', plenya: true, convenio: false },
  { trait: 'Acesso direto à equipe entre consultas', plenya: true, convenio: false },
  { trait: 'Foco em prevenção, não em reação', plenya: true, convenio: false },
  { trait: 'Conduta integrada (medicação + estilo de vida + suplementação)', plenya: true, convenio: '“peça encaminhamento”' },
  { trait: 'Continuidade — mesmo médico ao longo do tempo', plenya: true, convenio: 'depende' },
  { trait: 'Painel ampliado de exames além do convencional', plenya: true, convenio: false },
];

const rowsEn: Row[] = [
  { trait: 'Visit length', plenya: '60–90 min', convenio: '8–15 min' },
  { trait: 'Functional reading of labs', plenya: true, convenio: false },
  { trait: 'Multidisciplinary team discussing your case', plenya: true, convenio: false },
  { trait: 'Written plan, with goals and reviews', plenya: true, convenio: false },
  { trait: 'Plenya Score — an evolving measurement', plenya: true, convenio: false },
  { trait: 'Scheduled re-evaluation (3, 6, 12 months)', plenya: true, convenio: false },
  { trait: 'Direct access to the team between visits', plenya: true, convenio: false },
  { trait: 'Focus on prevention, not reaction', plenya: true, convenio: false },
  { trait: 'Integrated plan (medication + lifestyle + supplementation)', plenya: true, convenio: '“get a referral”' },
  { trait: 'Continuity — the same physician over time', plenya: true, convenio: 'depends' },
  { trait: 'Expanded lab panel beyond conventional checkups', plenya: true, convenio: false },
];

type Strings = {
  title: string;
  label: string;
  plenyaSubLabel: string;
  plenyaSubTitle: string;
  conventionalLabel: string;
  conventionalTitle: string;
  yes: string;
  no: string;
  footnote: string;
};

const stringsPt: Strings = {
  title: 'O que muda entre uma Consulta Plenya e o atendimento convencional.',
  label: 'Comparativo',
  plenyaSubLabel: 'Plenya',
  plenyaSubTitle: 'Consulta & Continuum',
  conventionalLabel: 'Tradicional',
  conventionalTitle: 'Convênio / consulta padrão',
  yes: 'Sim',
  no: 'Não',
  footnote:
    'Comparativo aplicável à média dos atendimentos por convênio e consultas particulares padrão no Brasil. Casos individuais podem variar.',
};

const stringsEn: Strings = {
  title: 'What changes between a Plenya Consultation and conventional care.',
  label: 'Comparison',
  plenyaSubLabel: 'Plenya',
  plenyaSubTitle: 'Consultation & Continuum',
  conventionalLabel: 'Conventional',
  conventionalTitle: 'Insurance / standard private visit',
  yes: 'Yes',
  no: 'No',
  footnote:
    'Comparison reflects the average of insurance-based and standard private-pay visits in Brazil. Individual cases may vary.',
};

export async function ComparatorVsConvenio({
  title,
  label,
  rows,
  bg = 'bg-paper',
}: {
  title?: string;
  label?: string;
  rows?: Row[];
  bg?: string;
}) {
  const rawLocale = await getLocale();
  const locale: Locale = isLocale(rawLocale) ? rawLocale : defaultLocale;
  const s = locale === 'en' ? stringsEn : stringsPt;
  const resolvedRows = rows ?? (locale === 'en' ? rowsEn : rowsPt);
  const resolvedTitle = title ?? s.title;
  const resolvedLabel = label ?? s.label;

  return (
    <section className={bg}>
      <div className="site-container section">
        <div className="max-w-3xl mb-12 space-y-4">
          <p className="label-upper text-gold">{resolvedLabel}</p>
          <h2 className="heading-section text-petrol text-3xl md:text-4xl">{resolvedTitle}</h2>
        </div>

        <div className="overflow-x-auto -mx-6 md:mx-0">
          <table className="w-full min-w-[640px] border-collapse">
            <thead>
              <tr className="border-b-2 border-petrol/20">
                <th className="text-left py-5 px-4 md:px-6 label-upper text-petrol/55 font-normal">
                  &nbsp;
                </th>
                <th className="py-5 px-4 md:px-6 text-left">
                  <span className="label-upper text-gold block">{s.plenyaSubLabel}</span>
                  <span className="heading-section text-petrol text-lg block mt-1">{s.plenyaSubTitle}</span>
                </th>
                <th className="py-5 px-4 md:px-6 text-left">
                  <span className="label-upper text-petrol/45 block">{s.conventionalLabel}</span>
                  <span className="heading-section text-petrol/60 text-lg block mt-1">{s.conventionalTitle}</span>
                </th>
              </tr>
            </thead>
            <tbody>
              {resolvedRows.map((r, i) => (
                <tr
                  key={r.trait}
                  className={i % 2 === 0 ? 'bg-cream/50' : ''}
                >
                  <td className="py-6 px-4 md:px-6 text-petrol/85 text-base md:text-lg align-middle min-h-[64px]">
                    {r.trait}
                  </td>
                  <td className="py-6 px-4 md:px-6 align-middle">
                    {r.plenya === true ? (
                      <span className="inline-flex items-center gap-2 text-gold">
                        <Check size={26} strokeWidth={2.5} />
                        <span className="sr-only">{s.yes}</span>
                      </span>
                    ) : (
                      <span className="text-petrol text-base md:text-lg">{r.plenya}</span>
                    )}
                  </td>
                  <td className="py-6 px-4 md:px-6 align-middle">
                    {r.convenio === false ? (
                      <span className="inline-flex items-center gap-2 text-petrol/30">
                        <X size={26} strokeWidth={2.5} />
                        <span className="sr-only">{s.no}</span>
                      </span>
                    ) : (
                      <span className="text-petrol/55 text-base md:text-lg italic">{r.convenio}</span>
                    )}
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>

        <p className="text-petrol/55 text-sm leading-relaxed mt-6 max-w-2xl">
          {s.footnote}
        </p>
      </div>
    </section>
  );
}
