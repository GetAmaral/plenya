import Image from 'next/image';
import { Link } from '@/lib/i18n/navigation';
import type { Doctor } from '@/lib/team';

export function DoctorCard({ doctor }: { doctor: Doctor }) {
  // Dr. Getúlio tem página dedicada
  const href = doctor.slug === 'getulio-amaral' ? '/dr-getulio' : `/equipe/${doctor.slug}`;

  return (
    <Link
      href={href}
      className="group block bg-paper hover:bg-cream-100 transition-colors duration-300"
    >
      <div className="relative aspect-[3/4] overflow-hidden">
        {doctor.photo ? (
          <Image
            src={doctor.photo}
            alt={doctor.name}
            fill
            className="object-cover object-top transition-transform duration-700 group-hover:scale-[1.02]"
            sizes="(min-width: 1024px) 360px, (min-width: 640px) 50vw, 100vw"
          />
        ) : (
          <div className="absolute inset-0 bg-cream-200" />
        )}
      </div>
      <div className="pt-6 space-y-2">
        <p className="label-upper text-gold">{doctor.role}</p>
        <h3 className="heading-section text-petrol text-xl group-hover:text-gold transition">
          {doctor.name}
        </h3>
        <p className="label-upper text-petrol/50">
          {doctor.credentials}{doctor.rqe ? ` · RQE ${doctor.rqe}` : ''}
        </p>
        <p className="text-petrol/70 text-sm leading-relaxed pt-1">{doctor.shortBio}</p>
      </div>
    </Link>
  );
}
