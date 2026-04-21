'use client';

import { useState } from 'react';
import { Link } from '@/lib/i18n/navigation';
import { track } from '@/lib/plausible';

type RetratoId =
  | 'tudo-certo'
  | 'susto'
  | 'carrega'
  | 'transicao'
  | 'isolados'
  | 'mede'
  | 'sinais'
  | 'inteiro';

type Retrato = {
  id: RetratoId;
  titulo: string;
  descricao: string;
};

const RETRATOS: Retrato[] = [
  {
    id: 'tudo-certo',
    titulo: 'Faz tudo certo e não vê resposta.',
    descricao:
      'Treina, come com cuidado, dorme as horas — e mesmo assim a disposição não volta, o peso não cede, os marcadores não melhoram. Falta o pilar que ninguém integrou.',
  },
  {
    id: 'susto',
    titulo: 'Recebeu o susto.',
    descricao:
      'Um exame que assustou, um diagnóstico de família, um aniversário que doeu. A janela parecia distante e ficou perto demais.',
  },
  {
    id: 'carrega',
    titulo: 'Carrega mais do que si mesmo.',
    descricao:
      'Sócios, equipe, filhos, pais — gente depende de você inteiro. Saúde virou variável crítica do que você sustenta no mundo.',
  },
  {
    id: 'transicao',
    titulo: 'Está numa transição que ninguém nomeia.',
    descricao:
      'O sono mudou, o humor mudou, o corpo mudou — e a resposta clínica que você recebe é "ainda não é nada" ou "é da idade".',
  },
  {
    id: 'isolados',
    titulo: 'Tem profissionais isolados.',
    descricao:
      'Personal, nutricionista, terapeuta, médico — cada um cuidando do seu pedaço, ninguém olhando o conjunto.',
  },
  {
    id: 'mede',
    titulo: 'Mede tudo e não decide nada.',
    descricao:
      'Wearables, painéis de exames, relatórios de longevidade — dado não falta. Falta alguém que leia o conjunto e acompanhe a decisão no tempo.',
  },
  {
    id: 'sinais',
    titulo: 'Tem sinais pequenos que ninguém soma.',
    descricao:
      'Adoeceu três vezes este ano. A barriga endureceu, embora a balança não acuse. O intestino bagunçou, a pele perdeu viço, o sono não recupera. Cada queixa parece pequena demais — junto, descreve um terreno em desgaste.',
  },
  {
    id: 'inteiro',
    titulo: 'Quer estar inteiro para quem ainda vai chegar.',
    descricao:
      'Filho pequeno aos 45, projetos que não cabem em dez anos, gente nova chegando. A pergunta deixou de ser quanto vou viver — virou como vou estar quando precisarem de mim inteiro.',
  },
];

type Option = {
  label: string;
  scores: Partial<Record<RetratoId, number>>;
};

type Question = {
  id: string;
  prompt: string;
  options: Option[];
};

const QUESTIONS: Question[] = [
  {
    id: 'idade',
    prompt: 'Em qual faixa você está hoje?',
    options: [
      { label: '35 a 44 anos', scores: { 'tudo-certo': 1, carrega: 1, inteiro: 2 } },
      { label: '45 a 54 anos', scores: { transicao: 1, susto: 1, mede: 1, inteiro: 1 } },
      { label: '55 anos ou mais', scores: { susto: 2, transicao: 1 } },
      { label: 'Menos de 35', scores: { 'tudo-certo': 1, mede: 1, inteiro: 1 } },
    ],
  },
  {
    id: 'motivo',
    prompt: 'O que mais te trouxe a esta página hoje?',
    options: [
      {
        label: 'Faço o que dizem ser certo e não vejo o resultado que esperava.',
        scores: { 'tudo-certo': 3 },
      },
      {
        label: 'Tive um susto recente — meu ou de família — e quero antecipar.',
        scores: { susto: 3 },
      },
      {
        label: 'Tenho gente que depende de mim inteiro: equipe, sócios, família.',
        scores: { carrega: 3 },
      },
      {
        label: 'Sinto que algo está mudando e ninguém consegue nomear o quê.',
        scores: { transicao: 3 },
      },
      {
        label: 'Já tenho profissionais separados — falta alguém juntando.',
        scores: { isolados: 3 },
      },
      {
        label: 'Tenho dado demais (wearables, exames) e decisão de menos.',
        scores: { mede: 3 },
      },
      {
        label: 'Acumulo sintomas pequenos — gripe, intestino, pele, sono — que ninguém soma.',
        scores: { sinais: 3 },
      },
      {
        label: 'Quero garantir presença e energia para quem ainda vai chegar — filhos, projetos, anos pela frente.',
        scores: { inteiro: 3 },
      },
    ],
  },
  {
    id: 'sono',
    prompt: 'Como está seu sono na maior parte das semanas?',
    options: [
      {
        label: 'Acordo descansado, energia estável o dia todo.',
        scores: {},
      },
      {
        label: 'Durmo as horas mas acordo sem energia.',
        scores: { 'tudo-certo': 2, transicao: 2, sinais: 2 },
      },
      {
        label: 'Não durmo o suficiente — deito tarde, acordo cedo.',
        scores: { carrega: 2, mede: 1 },
      },
      {
        label: 'Sono fragmentado — acordo várias vezes ou ronco muito.',
        scores: { susto: 1, transicao: 2, sinais: 2 },
      },
    ],
  },
  {
    id: 'exames',
    prompt: 'Qual sua relação com os exames que você já fez?',
    options: [
      {
        label: 'Os médicos dizem que está tudo normal — mas eu sinto que não está.',
        scores: { 'tudo-certo': 2, transicao: 1, sinais: 2 },
      },
      {
        label: 'Algo apareceu recentemente que me fez parar para pensar.',
        scores: { susto: 3 },
      },
      {
        label: 'Já fiz exames sofisticados e não sei o que fazer com o resultado.',
        scores: { mede: 3, isolados: 1 },
      },
      {
        label: 'Faço só o básico anual e gostaria de uma leitura mais profunda.',
        scores: { carrega: 1, susto: 1, inteiro: 2 },
      },
    ],
  },
  {
    id: 'cadencia',
    prompt: 'Encontros semanais online, em rotação entre quatro profissionais — cabe na sua vida?',
    options: [
      {
        label: 'Cabe — semanal é exatamente o ritmo que faz diferença.',
        scores: { 'tudo-certo': 1, carrega: 1, isolados: 1, mede: 1, sinais: 1, inteiro: 1 },
      },
      {
        label: 'Cabe com alguma flexibilidade de horário.',
        scores: { carrega: 2, inteiro: 1 },
      },
      {
        label: 'Preciso entender melhor antes de me comprometer.',
        scores: { susto: 1, transicao: 1 },
      },
    ],
  },
];

type Answers = Record<string, number>;

function computeTopRetratos(answers: Answers, allOptions: Question[]): Retrato[] {
  const totals: Record<RetratoId, number> = {
    'tudo-certo': 0,
    susto: 0,
    carrega: 0,
    transicao: 0,
    isolados: 0,
    mede: 0,
    sinais: 0,
    inteiro: 0,
  };

  for (const q of allOptions) {
    const optIdx = answers[q.id];
    if (optIdx === undefined) continue;
    const opt = q.options[optIdx];
    if (!opt) continue;
    for (const [k, v] of Object.entries(opt.scores)) {
      totals[k as RetratoId] += v ?? 0;
    }
  }

  const ranked = [...RETRATOS].sort((a, b) => totals[b.id] - totals[a.id]);
  const topScore = totals[ranked[0].id];
  if (topScore === 0) return ranked.slice(0, 3);
  return ranked.slice(0, 3);
}

export function DiagnosticoQuiz() {
  const [phase, setPhase] = useState<'intro' | 'quiz' | 'result'>('intro');
  const [step, setStep] = useState(0);
  const [answers, setAnswers] = useState<Answers>({});

  function start() {
    track('diagnostico_started');
    setPhase('quiz');
    setStep(0);
  }

  function answer(optionIdx: number) {
    const q = QUESTIONS[step];
    const next = { ...answers, [q.id]: optionIdx };
    setAnswers(next);
    if (step + 1 < QUESTIONS.length) {
      setStep(step + 1);
    } else {
      track('diagnostico_completed');
      setPhase('result');
    }
  }

  function back() {
    if (step > 0) setStep(step - 1);
    else setPhase('intro');
  }

  function restart() {
    setAnswers({});
    setStep(0);
    setPhase('intro');
  }

  if (phase === 'intro') {
    return (
      <div className="max-w-2xl">
        <p className="label-upper text-gold mb-4">Como funciona</p>
        <p className="text-petrol/85 text-lg leading-relaxed mb-4">
          Cinco perguntas. Sem diagnóstico. Sem cobrança de e-mail.
        </p>
        <p className="text-petrol/75 leading-relaxed mb-10">
          O resultado mostra quais retratos do Continuum Plenya mais se aproximam do
          seu momento — e abre o caminho para conversar com a equipe se fizer
          sentido.
        </p>
        <button
          type="button"
          onClick={start}
          className="btn-gold"
        >
          Começar
        </button>
      </div>
    );
  }

  if (phase === 'quiz') {
    const q = QUESTIONS[step];
    return (
      <div className="max-w-3xl">
        <div className="flex gap-2 mb-12">
          {QUESTIONS.map((_, i) => (
            <span
              key={i}
              className={`h-1 w-8 rounded-full transition ${
                i <= step ? 'bg-gold' : 'bg-petrol/15'
              }`}
            />
          ))}
        </div>

        <p className="label-upper text-gold mb-4">
          Pergunta {step + 1} de {QUESTIONS.length}
        </p>
        <h2 className="heading-section text-petrol text-2xl md:text-4xl mb-10">
          {q.prompt}
        </h2>

        <div className="space-y-3">
          {q.options.map((opt, i) => (
            <button
              key={i}
              type="button"
              onClick={() => answer(i)}
              className="w-full text-left p-6 bg-paper hover:bg-cream-100 border border-petrol/10 hover:border-gold transition group"
            >
              <span className="text-petrol/85 group-hover:text-petrol leading-relaxed">
                {opt.label}
              </span>
            </button>
          ))}
        </div>

        <button
          type="button"
          onClick={back}
          className="mt-10 text-petrol/60 hover:text-petrol transition text-sm label-upper"
        >
          ← Voltar
        </button>
      </div>
    );
  }

  const tops = computeTopRetratos(answers, QUESTIONS);

  return (
    <div className="max-w-4xl">
      <p className="label-upper text-gold mb-4">Sua leitura</p>
      <h2 className="heading-section text-petrol text-3xl md:text-5xl mb-6">
        O que você descreve se aproxima de três retratos do Continuum Plenya.
      </h2>
      <p className="text-petrol/75 text-lg leading-relaxed max-w-2xl mb-16">
        Não é diagnóstico. É a forma como costumamos reconhecer o momento em
        que faz sentido entrar num programa contínuo. A equipe pode aprofundar
        em conversa direta.
      </p>

      <div className="grid md:grid-cols-3 gap-px bg-petrol/10 border-y border-petrol/10 mb-16">
        {tops.map((r, i) => (
          <div key={r.id} className="bg-paper p-8 space-y-3">
            <p className="label-upper text-gold">Retrato {i + 1}</p>
            <h3 className="heading-section text-petrol text-xl">{r.titulo}</h3>
            <p className="text-petrol/75 leading-relaxed text-sm">
              {r.descricao}
            </p>
          </div>
        ))}
      </div>

      <div className="space-y-4 mb-12 max-w-2xl">
        <p className="text-petrol/85 leading-relaxed">
          O Continuum Plenya é um programa <strong className="text-petrol">semestral ou anual</strong>,
          100% online, conduzido por quatro profissionais que se reúnem,
          discutem o seu caso e desenham um plano único.
        </p>
        <p className="text-petrol/85 leading-relaxed">
          Tudo começa pelo <strong className="text-petrol">Escore Plenya</strong> (mais
          de 800 itens) e segue organizado pelo <strong className="text-petrol">Método AGIR</strong> —
          atividade física e alimentação, gestão clínica e metabólica,
          integração mente-corpo, ritmo circadiano e repouso — em encontros
          semanais que cobrem cada pilar a cada quatro semanas.
        </p>
      </div>

      <div className="flex flex-wrap gap-4">
        <Link href="/contato" className="btn-gold">
          Conversar com a equipe
        </Link>
        <Link
          href="/continuum"
          className="inline-flex items-center px-6 py-3 border border-petrol/30 text-petrol hover:bg-petrol hover:text-cream transition label-upper"
        >
          Ler sobre o Continuum Plenya
        </Link>
      </div>

      <button
        type="button"
        onClick={restart}
        className="mt-12 text-petrol/60 hover:text-petrol transition text-sm label-upper"
      >
        ← Refazer o diagnóstico
      </button>
    </div>
  );
}
