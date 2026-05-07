'use client';

import { useState } from 'react';
import { useTranslations } from 'next-intl';
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

const RETRATO_IDS: RetratoId[] = [
  'tudo-certo',
  'susto',
  'carrega',
  'transicao',
  'isolados',
  'mede',
  'sinais',
  'inteiro',
];

type Option = {
  scores: Partial<Record<RetratoId, number>>;
};

type Question = {
  id: 'idade' | 'motivo' | 'sono' | 'exames' | 'cadencia';
  options: Option[];
};

const QUESTIONS: Question[] = [
  {
    id: 'idade',
    options: [
      { scores: { 'tudo-certo': 1, carrega: 1, inteiro: 2 } },
      { scores: { transicao: 1, susto: 1, mede: 1, inteiro: 1 } },
      { scores: { susto: 2, transicao: 1 } },
      { scores: { 'tudo-certo': 1, mede: 1, inteiro: 1 } },
    ],
  },
  {
    id: 'motivo',
    options: [
      { scores: { 'tudo-certo': 3 } },
      { scores: { susto: 3 } },
      { scores: { carrega: 3 } },
      { scores: { transicao: 3 } },
      { scores: { isolados: 3 } },
      { scores: { mede: 3 } },
      { scores: { sinais: 3 } },
      { scores: { inteiro: 3 } },
    ],
  },
  {
    id: 'sono',
    options: [
      { scores: {} },
      { scores: { 'tudo-certo': 2, transicao: 2, sinais: 2 } },
      { scores: { carrega: 2, mede: 1 } },
      { scores: { susto: 1, transicao: 2, sinais: 2 } },
    ],
  },
  {
    id: 'exames',
    options: [
      { scores: { 'tudo-certo': 2, transicao: 1, sinais: 2 } },
      { scores: { susto: 3 } },
      { scores: { mede: 3, isolados: 1 } },
      { scores: { carrega: 1, susto: 1, inteiro: 2 } },
    ],
  },
  {
    id: 'cadencia',
    options: [
      { scores: { 'tudo-certo': 1, carrega: 1, isolados: 1, mede: 1, sinais: 1, inteiro: 1 } },
      { scores: { carrega: 2, inteiro: 1 } },
      { scores: { susto: 1, transicao: 1 } },
    ],
  },
];

type Answers = Record<string, number>;

function computeTopRetratos(answers: Answers, allOptions: Question[]): RetratoId[] {
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

  const ranked = [...RETRATO_IDS].sort((a, b) => totals[b] - totals[a]);
  return ranked.slice(0, 3);
}

export function DiagnosticoQuiz() {
  const t = useTranslations('diagnostic');
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
        <p className="label-upper text-gold mb-4">{t('quizIntroLabel')}</p>
        <p className="text-petrol/85 text-lg leading-relaxed mb-4">
          {t('quizIntroLine1')}
        </p>
        <p className="text-petrol/75 leading-relaxed mb-10">
          {t('quizIntroLine2')}
        </p>
        <button
          type="button"
          onClick={start}
          className="btn-gold"
        >
          {t('quizStart')}
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
          {t('quizProgress', { step: step + 1, total: QUESTIONS.length })}
        </p>
        <h2 className="heading-section text-petrol text-2xl md:text-4xl mb-10">
          {t(`questions.${q.id}.prompt`)}
        </h2>

        <div className="space-y-3">
          {q.options.map((_, i) => (
            <button
              key={i}
              type="button"
              onClick={() => answer(i)}
              className="w-full text-left p-6 bg-paper hover:bg-cream-100 border border-petrol/10 hover:border-gold transition group"
            >
              <span className="text-petrol/85 group-hover:text-petrol leading-relaxed">
                {t(`questions.${q.id}.opt${i}` as 'questions.idade.opt0')}
              </span>
            </button>
          ))}
        </div>

        <button
          type="button"
          onClick={back}
          className="mt-10 text-petrol/60 hover:text-petrol transition text-sm label-upper"
        >
          {t('quizBack')}
        </button>
      </div>
    );
  }

  const tops = computeTopRetratos(answers, QUESTIONS);

  return (
    <div className="max-w-4xl">
      <p className="label-upper text-gold mb-4">{t('resultLabel')}</p>
      <h2 className="heading-section text-petrol text-3xl md:text-5xl mb-6">
        {t('resultTitle')}
      </h2>
      <p className="text-petrol/75 text-lg leading-relaxed max-w-2xl mb-16">
        {t('resultIntro')}
      </p>

      <div className="grid md:grid-cols-3 gap-px bg-petrol/10 border-y border-petrol/10 mb-16">
        {tops.map((id, i) => (
          <div key={id} className="bg-paper p-8 space-y-3">
            <p className="label-upper text-gold">{t('resultRetratoLabel', { n: i + 1 })}</p>
            <h3 className="heading-section text-petrol text-xl">{t(`retratos.${id}.titulo`)}</h3>
            <p className="text-petrol/75 leading-relaxed text-sm">
              {t(`retratos.${id}.descricao`)}
            </p>
          </div>
        ))}
      </div>

      <div className="space-y-4 mb-12 max-w-2xl">
        <p className="text-petrol/85 leading-relaxed">
          {t('resultDescP1Part1')}
          <strong className="text-petrol">{t('resultDescP1Strong')}</strong>
          {t('resultDescP1Part2')}
        </p>
        <p className="text-petrol/85 leading-relaxed">
          {t('resultDescP2Part1')}
          <strong className="text-petrol">{t('resultDescP2Strong1')}</strong>
          {t('resultDescP2Part2')}
          <strong className="text-petrol">{t('resultDescP2Strong2')}</strong>
          {t('resultDescP2Part3')}
        </p>
      </div>

      <div className="flex flex-wrap gap-4">
        <Link href="/contato" className="btn-gold">
          {t('resultCtaContact')}
        </Link>
        <Link
          href="/continuum"
          className="inline-flex items-center px-6 py-3 border border-petrol/30 text-petrol hover:bg-petrol hover:text-cream transition label-upper"
        >
          {t('resultCtaContinuum')}
        </Link>
      </div>

      <button
        type="button"
        onClick={restart}
        className="mt-12 text-petrol/60 hover:text-petrol transition text-sm label-upper"
      >
        {t('resultRestart')}
      </button>
    </div>
  );
}
