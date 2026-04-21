import type { Metadata } from 'next';
import Image from 'next/image';
import { setRequestLocale } from 'next-intl/server';
import { isLocale, defaultLocale } from '@/lib/i18n/config';
import { Link } from '@/lib/i18n/navigation';
import { getAllDoctors } from '@/lib/team';
import { DoctorCard } from '@/components/team/doctor-card';

export const metadata: Metadata = {
  title: 'Equipe Plenya',
  description: 'Médicos e equipe multidisciplinar da Plenya — guardiões do Método AGIR.',
};

export default async function TeamPage({ params }: { params: Promise<{ locale: string }> }) {
  const { locale: rawLocale } = await params;
  const locale = isLocale(rawLocale) ? rawLocale : defaultLocale;
  setRequestLocale(locale);

  const all = await getAllDoctors();
  const direcao = all.find((d) => d.slug === 'getulio-amaral');
  const nucleoMedico = all.filter(
    (d) => d.category === 'medico' && d.slug !== 'getulio-amaral',
  );
  const multidisciplinar = all.filter((d) => d.category !== 'medico');

  return (
    <>
      {/* Hero */}
      <section className="bg-petrol text-cream">
        <div className="site-container pt-32 pb-16 md:pt-40 md:pb-20 grid lg:grid-cols-[1fr_auto] gap-12 lg:gap-20 items-center">
          <div className="max-w-xl">
            <p className="label-upper text-gold mb-6">Guardiões do Método</p>
            <h1 className="heading-hero text-[clamp(2.5rem,6vw,5rem)] text-cream">
              Equipe Plenya
            </h1>
            <p className="text-cream/80 text-lg mt-6 leading-relaxed">
              Médicos, nutricionista, psicóloga e educador físico falando a
              mesma língua — sobre o mesmo painel, sobre a mesma pessoa. A
              equipe inteira se reúne, discute o seu caso e desenha uma
              conduta única, personalizada e de precisão.
            </p>
            <p className="text-cream/60 mt-4 leading-relaxed">
              Não são especialidades operando em paralelo. É uma leitura
              clínica integrada — traduzida em prescrição compartilhada e
              revisada a cada ciclo, porque o corpo não funciona em partes e o
              cuidado também não.
            </p>
          </div>
          <div className="relative w-full max-w-md lg:w-[360px] lg:max-w-none aspect-[1064/1891] overflow-hidden bg-petrol/40">
            <Image
              src="/images/team/equipe-formal.jpg"
              alt="Equipe Plenya"
              fill
              priority
              className="object-cover object-top"
              sizes="(min-width: 1024px) 360px, 100vw"
            />
          </div>
        </div>
      </section>

      {/* Direção clínica — Dr. Getúlio em destaque */}
      {direcao && (
        <section className="bg-cream">
          <div className="site-container section">
            <div className="flex items-baseline gap-4 mb-12">
              <h2 className="heading-section text-petrol text-2xl md:text-3xl">
                Direção clínica
              </h2>
              <span className="label-upper text-petrol/40">quem conduz o cuidado</span>
            </div>
            <Link
              href="/dr-getulio"
              className="group grid md:grid-cols-[280px_1fr] gap-10 items-start bg-paper hover:bg-cream-100 transition-colors duration-300 p-6 md:p-10"
            >
              <div className="relative aspect-[3/4] overflow-hidden">
                {direcao.photo && (
                  <Image
                    src={direcao.photo}
                    alt={direcao.name}
                    fill
                    className="object-cover object-top transition-transform duration-700 group-hover:scale-[1.02]"
                    sizes="(min-width: 768px) 280px, 100vw"
                  />
                )}
              </div>
              <div className="space-y-4 pt-2">
                <p className="label-upper text-gold">{direcao.role}</p>
                <h3 className="heading-section text-petrol text-3xl md:text-4xl group-hover:text-gold transition">
                  {direcao.name}
                </h3>
                <p className="label-upper text-petrol/50">
                  {direcao.credentials}
                  {direcao.rqe ? ` · RQE ${direcao.rqe}` : ''}
                </p>
                <p className="text-petrol/80 leading-relaxed text-base md:text-lg max-w-2xl">
                  {direcao.shortBio}
                </p>
                <p className="label-upper text-gold pt-2 group-hover:underline underline-offset-4">
                  Conhecer a história →
                </p>
              </div>
            </Link>
          </div>
        </section>
      )}

      {/* Núcleo médico — 3 médicos */}
      {nucleoMedico.length > 0 && (
        <section className="bg-paper">
          <div className="site-container section">
            <div className="flex items-baseline gap-4 mb-12">
              <h2 className="heading-section text-petrol text-2xl md:text-3xl">
                Núcleo médico
              </h2>
              <span className="label-upper text-petrol/40">
                a mesma escola, a mesma língua
              </span>
            </div>
            <div className="grid gap-8 sm:grid-cols-2 lg:grid-cols-3">
              {nucleoMedico.map((d) => (
                <DoctorCard key={d.slug} doctor={d} />
              ))}
            </div>
          </div>
        </section>
      )}

      {/* Equipe multidisciplinar — nutri + psico (+ educador físico futuro) */}
      {multidisciplinar.length > 0 && (
        <section className="bg-cream">
          <div className="site-container section">
            <div className="flex items-baseline gap-4 mb-12">
              <h2 className="heading-section text-petrol text-2xl md:text-3xl">
                Cuidado integrado
              </h2>
              <span className="label-upper text-petrol/40">
                porque o corpo não funciona em partes
              </span>
            </div>
            <div className="grid gap-8 sm:grid-cols-2 lg:grid-cols-3">
              {multidisciplinar.map((d) => (
                <DoctorCard key={d.slug} doctor={d} />
              ))}
            </div>
          </div>
        </section>
      )}

      <section className="bg-petrol text-cream">
        <div className="site-container section text-center space-y-6">
          <p className="heading-section text-cream text-2xl md:text-3xl max-w-2xl mx-auto">
            Quer conhecer a equipe em conversa direta?
          </p>
          <Link href="/contato" className="btn-gold">
            Falar com a Plenya
          </Link>
        </div>
      </section>
    </>
  );
}
