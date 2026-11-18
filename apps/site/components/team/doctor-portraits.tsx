import Image from 'next/image';
import { getAllDoctors } from '@/lib/team';

/**
 * Retratos individuais do núcleo médico, lado a lado. Substitui as fotos de grupo
 * (equipe-formal / equipe-candid), que mostravam profissionais que não fazem mais
 * parte da equipe.
 */
export async function DoctorPortraits({
  columns = 2,
  sizes,
  className = '',
}: {
  columns?: 2 | 4;
  sizes: string;
  className?: string;
}) {
  const doctors = (await getAllDoctors()).filter((d) => d.category === 'medico' && d.photo);
  const grid = columns === 4 ? 'grid-cols-2 md:grid-cols-4' : 'grid-cols-2';

  return (
    <div className={`grid ${grid} gap-3 ${className}`}>
      {doctors.map((d) => (
        <figure key={d.slug} className="space-y-2">
          <div className="relative aspect-3/4 overflow-hidden bg-petrol/5">
            <Image src={d.photo!} alt={d.name} fill className="object-cover object-top" sizes={sizes} />
          </div>
          <figcaption className="label-upper-sm text-current opacity-60">{d.name}</figcaption>
        </figure>
      ))}
    </div>
  );
}
