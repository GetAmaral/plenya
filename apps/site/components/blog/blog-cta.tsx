import { Link } from '@/lib/i18n/navigation';

export function BlogCTA() {
  return (
    <div className="my-16 border-l-2 border-gold pl-8 py-6 space-y-4">
      <p className="label-upper text-gold">Plenya</p>
      <p className="heading-section text-petrol text-2xl max-w-xl">
        Quer aplicar o Método AGIR na sua vida? Agende uma consulta com Dr. Getúlio.
      </p>
      <div>
        <Link href="/contato" className="btn-gold">
          Agendar consulta
        </Link>
      </div>
    </div>
  );
}
