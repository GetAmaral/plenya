import Image from 'next/image';
import { useTranslations } from 'next-intl';
import { PlenyaSymbol } from '@plenya/brand/logo';
import { Link } from '@/lib/i18n/navigation';

export function DrGetulioPreview() {
  const tCta = useTranslations('cta');

  return (
    <section className="bg-paper">
      <div className="site-container section grid gap-16 lg:grid-cols-2 items-center">
        <div className="relative aspect-[3/4] overflow-hidden">
          <Image
            src="/images/dr-getulio.jpg"
            alt="Dr. Getúlio Amaral Filho — Direção Clínica Plenya"
            fill
            className="object-cover"
            sizes="(min-width: 1024px) 540px, 100vw"
          />
        </div>

        <div className="space-y-8">
          <div className="flex items-center gap-4">
            <PlenyaSymbol aria-hidden="true" className="h-7 w-auto text-gold" />
            <p className="label-upper text-gold">Direção Clínica</p>
          </div>

          <h2 className="heading-section text-petrol text-4xl md:text-5xl">
            Dr. Getúlio Amaral Filho
          </h2>

          <p className="label-upper text-petrol/50">
            CRM-PR 21.876 · RQE 16.038 · Nefrologia · Medicina Funcional Integrativa
          </p>

          <p className="text-petrol/80 text-lg leading-relaxed max-w-lg">
            Médico desde 2004. Especialista em nefrologia e clínica médica pela Santa Casa de
            Londrina, onde coordena a residência médica em nefrologia. Concluiu pós-graduação
            em medicina funcional integrativa pela ABMFI em 2026.
          </p>

          <div className="flex flex-wrap gap-5 pt-2">
            <Link href="/dr-getulio" className="btn-gold">
              Conhecer Dr. Getúlio
            </Link>
            <Link href="/contato" className="btn-outline-dark">
              {tCta('scheduleConsultation')}
            </Link>
          </div>
        </div>
      </div>
    </section>
  );
}
