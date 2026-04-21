'use client';

import { useMemo, useState } from 'react';
import { useRouter } from 'next/navigation';
import type {
  LightConfig,
  LightItemConfig,
  SessionResponse,
} from '@/lib/score-light/types';
import { createSession } from '@/lib/score-light/api';

type Demographics = {
  age: number | '';
  gender: 'male' | 'female' | 'other' | '';
  postMenopause: boolean | null;
  height: number | '';
  weight: number | '';
};

type ResponseMap = Record<string, SessionResponse>;

const STEP_INTRO = 'intro';
const STEP_DEMO = 'demographics';

function appliesTo(item: LightItemConfig, demo: Demographics): boolean {
  if (item.gender && item.gender !== 'not_applicable' && item.gender !== demo.gender) {
    return false;
  }
  if (typeof item.ageRangeMin === 'number' && typeof demo.age === 'number' && demo.age < item.ageRangeMin) {
    return false;
  }
  if (typeof item.ageRangeMax === 'number' && typeof demo.age === 'number' && demo.age > item.ageRangeMax) {
    return false;
  }
  if (typeof item.postMenopause === 'boolean' && demo.gender === 'female') {
    if (demo.postMenopause === null) return false;
    if (item.postMenopause !== demo.postMenopause) return false;
  }
  return true;
}

function levelLabel(level: { name: string; lowerLimit?: string; upperLimit?: string; operator: string }): string {
  if (level.name) return level.name;
  if (level.operator === 'between') return `${level.lowerLimit ?? '—'} a ${level.upperLimit ?? '—'}`;
  return `${level.operator} ${level.lowerLimit ?? level.upperLimit ?? ''}`.trim();
}

export function EscoreLightForm({ config, locale }: { config: LightConfig; locale: string }) {
  const router = useRouter();
  const [step, setStep] = useState<string>(STEP_INTRO);
  const [demo, setDemo] = useState<Demographics>({
    age: '',
    gender: '',
    postMenopause: null,
    height: '',
    weight: '',
  });
  const [responses, setResponses] = useState<ResponseMap>({});
  const [submitting, setSubmitting] = useState(false);
  const [submitError, setSubmitError] = useState<string | null>(null);

  // Lista de grupos com items aplicáveis (filtragem dinâmica por demografia)
  const groupSteps = useMemo(() => {
    return config.groups
      .map((g) => ({
        ...g,
        items: g.subgroups
          .flatMap((sg) => sg.items)
          .filter((it) => appliesTo(it, demo))
          .sort((a, b) => a.lightOrder - b.lightOrder),
      }))
      .filter((g) => g.items.length > 0);
  }, [config, demo]);

  const totalSteps = 1 + 1 + groupSteps.length; // intro + demo + N grupos
  const currentStepIndex =
    step === STEP_INTRO
      ? 0
      : step === STEP_DEMO
      ? 1
      : 2 + groupSteps.findIndex((g) => g.id === step);

  const setResponse = (itemId: string, partial: Partial<SessionResponse>) => {
    setResponses((prev) => ({
      ...prev,
      [itemId]: { ...prev[itemId], scoreItemId: itemId, ...partial },
    }));
  };

  const goNext = () => {
    if (step === STEP_INTRO) {
      setStep(STEP_DEMO);
      return;
    }
    if (step === STEP_DEMO) {
      if (groupSteps[0]) setStep(groupSteps[0].id);
      else void submit();
      return;
    }
    const idx = groupSteps.findIndex((g) => g.id === step);
    const next = groupSteps[idx + 1];
    if (next) setStep(next.id);
    else void submit();
  };

  const goBack = () => {
    if (step === STEP_DEMO) {
      setStep(STEP_INTRO);
      return;
    }
    const idx = groupSteps.findIndex((g) => g.id === step);
    if (idx === 0) {
      setStep(STEP_DEMO);
      return;
    }
    const prev = groupSteps[idx - 1];
    if (prev) setStep(prev.id);
  };

  const demoIsValid = useMemo(() => {
    if (typeof demo.age !== 'number' || demo.age < 18 || demo.age > 120) return false;
    if (!demo.gender) return false;
    if (demo.gender === 'female' && demo.postMenopause === null) return false;
    return true;
  }, [demo]);

  const groupComplete = (groupId: string): boolean => {
    const g = groupSteps.find((x) => x.id === groupId);
    if (!g) return true;
    return g.items.every((it) => responses[it.id] !== undefined);
  };

  async function submit() {
    setSubmitting(true);
    setSubmitError(null);
    try {
      if (typeof demo.age !== 'number' || !demo.gender) {
        throw new Error('Dados demográficos incompletos');
      }
      const payload = {
        age: demo.age,
        gender: demo.gender,
        postMenopause: demo.gender === 'female' ? demo.postMenopause ?? undefined : undefined,
        height: typeof demo.height === 'number' ? demo.height : undefined,
        weight: typeof demo.weight === 'number' ? demo.weight : undefined,
        responses: Object.values(responses),
      };
      const session = await createSession(payload);
      router.push(`/${locale}/escore-plenya/resultado/${session.publicCode}`);
    } catch (err) {
      setSubmitError(err instanceof Error ? err.message : 'Falha ao enviar respostas');
      setSubmitting(false);
    }
  }

  // === Renderers ===

  const renderProgress = () => (
    <div className="flex items-center gap-2 mb-12">
      {Array.from({ length: totalSteps }).map((_, i) => (
        <span
          key={i}
          className={`h-1.5 w-8 rounded-full transition-all ${
            i <= currentStepIndex ? 'bg-gold' : 'bg-petrol/15'
          }`}
        />
      ))}
    </div>
  );

  const renderIntro = () => (
    <div className="space-y-8 max-w-2xl">
      <p className="label-upper text-gold">Escore Plenya Light</p>
      <h1 className="heading-section text-petrol text-3xl md:text-4xl leading-tight">
        Uma fotografia clara da sua saúde — em poucas perguntas.
      </h1>
      <p className="text-petrol/80 text-lg leading-relaxed">
        Esta é a versão pública do Escore Plenya. Você responde {config.itemCount} perguntas
        sobre hábitos, histórico e medidas básicas. Ao final, recebe um radar com sua
        pontuação em cada pilar do Método AGIR.
      </p>
      <ul className="space-y-3 text-petrol/70">
        <li className="flex gap-3">
          <span className="text-gold">—</span>
          <span>10 a 15 minutos. Anônimo. Sem cadastro.</span>
        </li>
        <li className="flex gap-3">
          <span className="text-gold">—</span>
          <span>Você pode salvar o resultado por email no final, se quiser.</span>
        </li>
        <li className="flex gap-3">
          <span className="text-gold">—</span>
          <span>Não substitui consulta médica. É um ponto de partida para a conversa.</span>
        </li>
      </ul>
      <div className="pt-6">
        <button onClick={goNext} className="btn-gold">
          Começar
        </button>
      </div>
    </div>
  );

  const renderDemographics = () => (
    <div className="space-y-8 max-w-xl">
      <p className="label-upper text-gold">Sobre você</p>
      <h2 className="heading-section text-petrol text-2xl md:text-3xl">
        Algumas informações básicas para personalizar a avaliação.
      </h2>

      <div className="space-y-6">
        <label className="block">
          <span className="block text-petrol/80 mb-2">Idade</span>
          <input
            type="number"
            min={18}
            max={120}
            value={demo.age}
            onChange={(e) =>
              setDemo({ ...demo, age: e.target.value === '' ? '' : Number(e.target.value) })
            }
            className="w-full border border-petrol/20 bg-cream px-4 py-3 text-petrol focus:border-gold focus:outline-none"
            placeholder="Ex: 47"
          />
        </label>

        <fieldset>
          <legend className="block text-petrol/80 mb-3">Sexo biológico</legend>
          <div className="flex gap-3">
            {(['male', 'female', 'other'] as const).map((g) => (
              <button
                key={g}
                type="button"
                onClick={() => setDemo({ ...demo, gender: g, postMenopause: null })}
                className={`px-5 py-3 border transition ${
                  demo.gender === g
                    ? 'bg-petrol text-cream border-petrol'
                    : 'bg-cream text-petrol border-petrol/20 hover:border-petrol/50'
                }`}
              >
                {g === 'male' ? 'Masculino' : g === 'female' ? 'Feminino' : 'Outro'}
              </button>
            ))}
          </div>
        </fieldset>

        {demo.gender === 'female' && (
          <fieldset>
            <legend className="block text-petrol/80 mb-3">Status de menopausa</legend>
            <div className="flex gap-3">
              {[
                { v: true, label: 'Pós-menopausa' },
                { v: false, label: 'Pré-menopausa' },
              ].map((opt) => (
                <button
                  key={String(opt.v)}
                  type="button"
                  onClick={() => setDemo({ ...demo, postMenopause: opt.v })}
                  className={`px-5 py-3 border transition ${
                    demo.postMenopause === opt.v
                      ? 'bg-petrol text-cream border-petrol'
                      : 'bg-cream text-petrol border-petrol/20 hover:border-petrol/50'
                  }`}
                >
                  {opt.label}
                </button>
              ))}
            </div>
          </fieldset>
        )}

        <div className="grid grid-cols-2 gap-4">
          <label className="block">
            <span className="block text-petrol/80 mb-2">Altura (cm) — opcional</span>
            <input
              type="number"
              value={demo.height}
              onChange={(e) =>
                setDemo({
                  ...demo,
                  height: e.target.value === '' ? '' : Number(e.target.value),
                })
              }
              className="w-full border border-petrol/20 bg-cream px-4 py-3 text-petrol focus:border-gold focus:outline-none"
              placeholder="Ex: 170"
            />
          </label>
          <label className="block">
            <span className="block text-petrol/80 mb-2">Peso (kg) — opcional</span>
            <input
              type="number"
              value={demo.weight}
              onChange={(e) =>
                setDemo({
                  ...demo,
                  weight: e.target.value === '' ? '' : Number(e.target.value),
                })
              }
              className="w-full border border-petrol/20 bg-cream px-4 py-3 text-petrol focus:border-gold focus:outline-none"
              placeholder="Ex: 72"
            />
          </label>
        </div>
      </div>

      <div className="flex items-center gap-4 pt-6">
        <button onClick={goBack} className="text-petrol/60 hover:text-petrol transition underline-offset-4 hover:underline">
          ← Voltar
        </button>
        <button
          onClick={goNext}
          disabled={!demoIsValid}
          className={`btn-gold ${!demoIsValid ? 'opacity-40 cursor-not-allowed' : ''}`}
        >
          Continuar
        </button>
      </div>
    </div>
  );

  const renderGroupStep = (groupId: string) => {
    const group = groupSteps.find((g) => g.id === groupId);
    if (!group) return null;
    return (
      <div className="space-y-8 max-w-2xl">
        <p className="label-upper text-gold">{group.name}</p>
        <h2 className="heading-section text-petrol text-2xl md:text-3xl">
          {group.items.length} {group.items.length === 1 ? 'pergunta' : 'perguntas'} sobre {group.name.toLowerCase()}.
        </h2>
        <div className="space-y-10">
          {group.items.map((item) => renderItemQuestion(item))}
        </div>
        <div className="flex items-center gap-4 pt-6">
          <button onClick={goBack} className="text-petrol/60 hover:text-petrol transition underline-offset-4 hover:underline">
            ← Voltar
          </button>
          <button
            onClick={goNext}
            disabled={!groupComplete(groupId) || submitting}
            className={`btn-gold ${
              !groupComplete(groupId) || submitting ? 'opacity-40 cursor-not-allowed' : ''
            }`}
          >
            {submitting
              ? 'Enviando...'
              : groupSteps[groupSteps.length - 1].id === groupId
              ? 'Ver meu resultado'
              : 'Continuar'}
          </button>
        </div>
        {submitError && <p className="text-red-700 text-sm">{submitError}</p>}
      </div>
    );
  };

  const renderItemQuestion = (item: LightItemConfig) => {
    const resp = responses[item.id];
    const isLevelChoice = item.levels.length > 0 && !item.labTestCode;

    return (
      <div key={item.id} className="border-t border-petrol/15 pt-6">
        <p className="text-petrol text-lg leading-relaxed mb-4">{item.lightQuestion}</p>
        {isLevelChoice ? (
          <div className="grid sm:grid-cols-2 gap-2">
            {item.levels
              .slice()
              .sort((a, b) => a.level - b.level)
              .map((lv) => {
                const selected = resp?.selectedLevel === lv.level;
                return (
                  <button
                    key={lv.id}
                    type="button"
                    onClick={() => setResponse(item.id, { selectedLevel: lv.level })}
                    className={`text-left px-4 py-3 border transition ${
                      selected
                        ? 'bg-petrol text-cream border-petrol'
                        : 'bg-cream text-petrol border-petrol/20 hover:border-petrol/50'
                    }`}
                  >
                    {levelLabel(lv)}
                  </button>
                );
              })}
          </div>
        ) : (
          <div className="flex items-center gap-3">
            <input
              type="number"
              value={resp?.numericValue ?? ''}
              onChange={(e) =>
                setResponse(item.id, {
                  numericValue: e.target.value === '' ? undefined : Number(e.target.value),
                })
              }
              className="flex-1 border border-petrol/20 bg-cream px-4 py-3 text-petrol focus:border-gold focus:outline-none"
              placeholder={item.unit ?? 'Valor numérico'}
            />
            {item.unit && <span className="text-petrol/60">{item.unit}</span>}
          </div>
        )}
      </div>
    );
  };

  return (
    <section className="bg-cream min-h-screen">
      <div className="site-container pt-32 pb-32 md:pt-40 md:pb-40">
        {step !== STEP_INTRO && renderProgress()}
        {step === STEP_INTRO && renderIntro()}
        {step === STEP_DEMO && renderDemographics()}
        {step !== STEP_INTRO && step !== STEP_DEMO && renderGroupStep(step)}
      </div>
    </section>
  );
}
