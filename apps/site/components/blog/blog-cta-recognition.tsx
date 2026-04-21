import { Link } from '@/lib/i18n/navigation';

export function BlogCTARecognition() {
  return (
    <div className="my-16 border-l-2 border-gold pl-8 py-6 space-y-6">
      <p className="label-upper text-gold">Reconhecimento</p>
      <p className="heading-section text-petrol text-2xl max-w-xl">
        Esse texto descreve o que você sente?
      </p>
      <p className="text-petrol/75 max-w-xl leading-relaxed">
        Se sim, talvez faça sentido entender como o Continuum Plenya conduz esse
        tipo de cuidado. Cinco perguntas curtas devolvem a leitura — ou a
        equipe pode ouvir você diretamente.
      </p>
      <div className="flex flex-wrap gap-3">
        <Link href="/diagnostico" className="btn-gold">
          Fazer o diagnóstico
        </Link>
        <Link
          href="/contato"
          className="inline-flex items-center px-6 py-3 border border-petrol/30 text-petrol hover:bg-petrol hover:text-cream transition label-upper"
        >
          Conversar com a equipe
        </Link>
      </div>
    </div>
  );
}
