/**
 * Ilustrações instrucionais inline para itens do Escore Light que exigem execução
 * (não autorrelato): Prancha, Cintura abdominal, Pescoço.
 *
 * Imagens geradas via OpenAI gpt-image-1 (estilo lineart minimalista),
 * salvas em /public/escore-light/instructions/.
 */
import Image from 'next/image';

const INSTRUCTION_NAMES: Record<string, 'plank' | 'waist' | 'neck'> = {
  'Resistência Neuromuscular - Prancha (Homens)': 'plank',
  'Resistência Neuromuscular - Prancha (Mulheres)': 'plank',
  'Abdominal (cintura em cm) - homem': 'waist',
  'Abdominal (cintura em cm) - mulher': 'waist',
  'Pescoço - homem': 'neck',
  'Pescoço - mulher': 'neck',
};

type InstructionType = 'plank' | 'waist' | 'neck';

const INSTRUCTIONS: Record<
  InstructionType,
  { title: string; image: string; alt: string; bullets: React.ReactNode[] }
> = {
  plank: {
    title: 'Como executar',
    image: '/escore-light/instructions/plank.png',
    alt: 'Ilustração de uma pessoa em prancha frontal correta — corpo em linha reta da cabeça aos calcanhares',
    bullets: [
      <>Apoie antebraços e pontas dos pés no chão</>,
      <>Mantenha o corpo em <strong className="text-petrol">linha reta</strong> (sem afundar quadril nem levantar o bumbum)</>,
      <>Olhe para o chão, pescoço neutro</>,
      <>Cronometre desde que entrar até <strong className="text-petrol">perder a forma</strong></>,
    ],
  },
  waist: {
    title: 'Como medir',
    image: '/escore-light/instructions/waist.png',
    alt: 'Ilustração de torso humano com fita métrica posicionada na altura do umbigo',
    bullets: [
      <>Em pé, expire normalmente (não prenda a barriga)</>,
      <>Posicione a fita na <strong className="text-petrol">altura do umbigo</strong></>,
      <>Mantenha a fita paralela ao chão, sem apertar</>,
      <>Anote o valor em <strong className="text-petrol">centímetros</strong></>,
    ],
  },
  neck: {
    title: 'Como medir',
    image: '/escore-light/instructions/neck.png',
    alt: 'Ilustração de pescoço humano com fita métrica posicionada logo abaixo do pomo de adão',
    bullets: [
      <>Em pé, olhar para frente, pescoço neutro</>,
      <>Posicione a fita <strong className="text-petrol">logo abaixo do pomo de adão</strong> (parte mais estreita)</>,
      <>Mantenha a fita paralela ao chão, sem apertar</>,
      <>Anote o valor em <strong className="text-petrol">centímetros</strong></>,
    ],
  },
};

export function getInstructionType(itemName: string): InstructionType | null {
  return INSTRUCTION_NAMES[itemName] ?? null;
}

export function ItemInstruction({ type }: { type: InstructionType }) {
  const data = INSTRUCTIONS[type];
  return (
    <div className="bg-paper border border-petrol/15 rounded-md p-5 my-3">
      <p className="label-upper text-gold mb-3 text-[10px]">{data.title}</p>
      <div className="flex flex-col sm:flex-row gap-5 items-start">
        <div className="shrink-0 w-44 h-44 relative">
          <Image
            src={data.image}
            alt={data.alt}
            fill
            sizes="176px"
            className="object-contain"
          />
        </div>
        <ul className="text-sm text-petrol/75 space-y-1.5 leading-relaxed">
          {data.bullets.map((b, i) => (
            <li key={i}>— {b}</li>
          ))}
        </ul>
      </div>
    </div>
  );
}
