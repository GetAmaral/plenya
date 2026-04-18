import { useTranslations } from 'next-intl';
import { PlenyaSymbol } from '@plenya/brand/logo';
import { Link } from '@/lib/i18n/navigation';

export function ScoreSection() {
  const t = useTranslations('home');
  return (
    <section className="bg-cream">
      <div className="site-container section grid gap-16 lg:grid-cols-2 items-center">
        <div className="space-y-8">
          <p className="label-upper text-gold">Diagnóstico</p>
          <h2 className="heading-section text-petrol text-3xl md:text-5xl">{t('scoreTitle')}</h2>
          <p className="text-petrol/80 text-lg leading-relaxed max-w-lg">{t('scoreSubtitle')}</p>

          <ul className="space-y-4 text-petrol/70">
            <li className="flex gap-4 items-start">
              <span className="text-gold text-lg leading-none mt-0.5">—</span>
              <span>Versão Completa — aplicada pela equipe Plenya</span>
            </li>
            <li className="flex gap-4 items-start">
              <span className="text-gold text-lg leading-none mt-0.5">—</span>
              <span>Versão Intermediária — após primeira consulta</span>
            </li>
            <li className="flex gap-4 items-start">
              <span className="text-gold text-lg leading-none mt-0.5">—</span>
              <span>Versão Light — disponível online em breve</span>
            </li>
          </ul>

          <Link href="/escore-plenya" className="btn-outline-dark">
            Entender o Escore
          </Link>
        </div>

        <div className="relative flex items-center justify-center">
          <svg viewBox="0 0 300 300" className="w-64 h-64 md:w-80 md:h-80" aria-hidden>
            <circle cx="150" cy="150" r="140" stroke="#063b4f" strokeOpacity="0.08" strokeWidth="1" fill="none" />
            <circle cx="150" cy="150" r="110" stroke="#063b4f" strokeOpacity="0.12" strokeWidth="1" fill="none" />
            <circle cx="150" cy="150" r="80" stroke="#063b4f" strokeOpacity="0.16" strokeWidth="1" fill="none" />
            <path d="M 150 10 A 140 140 0 1 1 40 260" stroke="#b38645" strokeWidth="2" fill="none" strokeLinecap="round" />
          </svg>
          <PlenyaSymbol
            aria-hidden="true"
            focusable="false"
            className="absolute h-20 w-20 md:h-24 md:w-24 text-petrol"
          />
        </div>
      </div>
    </section>
  );
}
