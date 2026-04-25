import type { Metadata } from 'next';
import { setRequestLocale } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';

export const metadata: Metadata = {
  title: 'Agendar consulta — Plenya',
  description:
    'Agende sua Consulta Plenya — presencial em Londrina ou online por telemedicina. Resposta em até 2 horas em dias úteis.',
};

export default async function AgendarPage({ params }: { params: Promise<{ locale: string }> }) {
  const { locale } = await params;
  setRequestLocale(locale);

  return (
    <>
      <section className="bg-petrol text-cream">
        <div className="site-container pt-32 pb-20 md:pt-40 md:pb-24 max-w-3xl">
          <p className="label-upper text-gold mb-6">Agendamento</p>
          <h1 className="heading-hero text-[clamp(2.5rem,6vw,4.5rem)] text-cream">
            Marcar a sua Consulta Plenya.
          </h1>
          <p className="text-cream/75 text-lg leading-relaxed mt-8 max-w-2xl">
            Presencial em Londrina-PR ou online por telemedicina, com a mesma profundidade clínica.
            Atendimento particular, sem convênios.
          </p>
        </div>
      </section>

      <section className="bg-cream">
        <div className="site-container section grid lg:grid-cols-[1fr_1fr] gap-16">
          <div className="space-y-8">
            <div className="space-y-3">
              <p className="label-upper text-gold">Por enquanto, agendamento conduzido pela equipe</p>
              <p className="text-petrol/80 text-lg leading-relaxed max-w-md">
                Estamos finalizando a integração de calendário online. Até lá, o agendamento é
                feito diretamente com a equipe da clínica — pelo formulário ao lado, por WhatsApp
                ou por e-mail. Resposta em até 2 horas em dias úteis (8h–18h).
              </p>
            </div>

            <div className="space-y-3 border-t border-petrol/15 pt-6">
              <p className="label-upper text-petrol/55">O que você recebe ao agendar</p>
              <ul className="space-y-3 text-petrol/75">
                <li className="flex gap-3"><span className="text-gold leading-none mt-1.5">—</span><span>Confirmação por WhatsApp em até 2 horas</span></li>
                <li className="flex gap-3"><span className="text-gold leading-none mt-1.5">—</span><span>Pré-orientação para a consulta (exames, jejum se houver)</span></li>
                <li className="flex gap-3"><span className="text-gold leading-none mt-1.5">—</span><span>Lembretes 48h e 24h antes</span></li>
              </ul>
            </div>
          </div>

          <div className="bg-paper p-8 md:p-10 space-y-6 border-t-2 border-gold">
            <p className="label-upper text-gold">Marcar agora</p>
            <h2 className="heading-section text-petrol text-2xl md:text-3xl">
              Três caminhos diretos.
            </h2>
            <div className="space-y-5 pt-2">
              <Link href="/contato" className="btn-gold w-full text-center block">
                Preencher formulário
              </Link>
              <a
                href="https://wa.me/554399999999"
                target="_blank"
                rel="noreferrer"
                className="btn-outline-dark w-full text-center block"
              >
                Conversar no WhatsApp
              </a>
              <a
                href="mailto:contato@plenyasaude.com.br"
                className="block text-center label-upper text-petrol/60 hover:text-gold transition pt-2"
              >
                contato@plenyasaude.com.br
              </a>
            </div>
          </div>
        </div>
      </section>
    </>
  );
}
