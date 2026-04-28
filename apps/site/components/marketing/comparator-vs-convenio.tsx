import { Check, X } from 'lucide-react';

type Row = { trait: string; plenya: string | true; convenio: string | false };

const defaultRows: Row[] = [
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
  { trait: 'Pagamento', plenya: 'Particular', convenio: 'Convênio / SUS' },
];

export function ComparatorVsConvenio({
  title = 'O que muda entre uma Consulta Plenya e o atendimento convencional.',
  label = 'Comparativo',
  rows = defaultRows,
  bg = 'bg-paper',
}: {
  title?: string;
  label?: string;
  rows?: Row[];
  bg?: string;
}) {
  return (
    <section className={bg}>
      <div className="site-container section">
        <div className="max-w-3xl mb-12 space-y-4">
          <p className="label-upper text-gold">{label}</p>
          <h2 className="heading-section text-petrol text-3xl md:text-4xl">{title}</h2>
        </div>

        <div className="overflow-x-auto -mx-6 md:mx-0">
          <table className="w-full min-w-[640px] border-collapse">
            <thead>
              <tr className="border-b-2 border-petrol/20">
                <th className="text-left py-5 px-4 md:px-6 label-upper text-petrol/55 font-normal">
                  &nbsp;
                </th>
                <th className="py-5 px-4 md:px-6 text-left">
                  <span className="label-upper text-gold block">Plenya</span>
                  <span className="heading-section text-petrol text-lg block mt-1">Consulta &amp; Continuum</span>
                </th>
                <th className="py-5 px-4 md:px-6 text-left">
                  <span className="label-upper text-petrol/45 block">Tradicional</span>
                  <span className="heading-section text-petrol/60 text-lg block mt-1">Convênio / consulta padrão</span>
                </th>
              </tr>
            </thead>
            <tbody>
              {rows.map((r, i) => (
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
                        <span className="sr-only">Sim</span>
                      </span>
                    ) : (
                      <span className="text-petrol text-base md:text-lg">{r.plenya}</span>
                    )}
                  </td>
                  <td className="py-6 px-4 md:px-6 align-middle">
                    {r.convenio === false ? (
                      <span className="inline-flex items-center gap-2 text-petrol/30">
                        <X size={26} strokeWidth={2.5} />
                        <span className="sr-only">Não</span>
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
          Comparativo aplicável à média dos atendimentos por convênio e consultas particulares
          padrão no Brasil. Casos individuais podem variar.
        </p>
      </div>
    </section>
  );
}
