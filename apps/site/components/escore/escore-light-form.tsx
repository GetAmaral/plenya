'use client';

import { useEffect, useMemo, useState } from 'react';
import { useRouter } from 'next/navigation';
import { useTranslations } from 'next-intl';
import type {
  LightConfig,
  LightItemConfig,
  SessionResponse,
} from '@/lib/score-light/types';
import { createSession } from '@/lib/score-light/api';
import { getGroupIcon } from './group-icons';
import { ItemInstruction, getInstructionType } from './item-instructions';
import { PHQ9Widget, PHQ9_ITEM_IDS } from './phq9-widget';
import { NumericClassifierInput, isNumericClassifiable } from './numeric-classifier-input';
import { LabPDFUpload } from './lab-pdf-upload';
import type { LightLabMatch } from '@/lib/score-light/api';
import { PRIVACY_POLICY_VERSION } from '@/lib/legal';
import { captureUTMFromURL, utmToPayload } from '@/lib/utm';

type Demographics = {
  age: number | '';
  gender: 'male' | 'female' | 'other' | '';
  height: number | '';
  weight: number | '';
};

type ResponseMap = Record<string, SessionResponse>;

const STEP_INTRO = 'intro';
const STEP_DEMO = 'demographics';

// IDs de itens derivados de outros (calculados automaticamente).
// Quando os inputs estão disponíveis, o valor é injetado em responses[id].
const DERIVED_ITEM_IDS = {
  IMC: '019bf31d-2ef0-7bb8-8305-e0e0285aeb80',
  WAIST_HEIGHT_RATIO: '019bf31d-2ef0-7171-a061-3bcc500957f7',
} as const;

const CINTURA_ITEM_IDS = {
  male: 'c77cedd3-2800-7aac-985c-c075f020e9e0',
  female: 'c77cedd3-2800-7a74-99ad-56ca4b6dddc1',
} as const;

// DERIVED_NOTES agora vem de t() — chaves: formDerivedBMI / formDerivedWaistHeight

function isDerivedItem(itemId: string): boolean {
  return (Object.values(DERIVED_ITEM_IDS) as string[]).includes(itemId);
}

function computeDerivedValue(
  itemId: string,
  demo: Demographics,
  responses: ResponseMap,
): number | undefined {
  if (itemId === DERIVED_ITEM_IDS.IMC) {
    if (typeof demo.height !== 'number' || typeof demo.weight !== 'number') return undefined;
    if (demo.height <= 0 || demo.weight <= 0) return undefined;
    const heightM = demo.height / 100;
    return Math.round((demo.weight / (heightM * heightM)) * 10) / 10;
  }
  if (itemId === DERIVED_ITEM_IDS.WAIST_HEIGHT_RATIO) {
    if (typeof demo.height !== 'number' || demo.height <= 0) return undefined;
    const cinturaId =
      demo.gender === 'female' ? CINTURA_ITEM_IDS.female : CINTURA_ITEM_IDS.male;
    const cintura = responses[cinturaId]?.numericValue;
    if (typeof cintura !== 'number') return undefined;
    return Math.round((cintura / demo.height) * 100) / 100;
  }
  return undefined;
}

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
  // Items que dependem de status de menopausa foram desabilitados do Light
  // (ver curadoria v5). Qualquer item que ainda venha do backend com
  // postMenopause definido é ignorado defensivamente.
  if (typeof item.postMenopause === 'boolean') {
    return false;
  }
  return true;
}

function levelLabel(level: { name: string; lowerLimit?: string; upperLimit?: string; operator: string }): string {
  if (level.name) return level.name;
  if (level.operator === 'between') return `${level.lowerLimit ?? '—'} a ${level.upperLimit ?? '—'}`;
  return `${level.operator} ${level.lowerLimit ?? level.upperLimit ?? ''}`.trim();
}

export function EscoreLightForm({
  config,
  locale,
  tierLabel = 'Escore Plenya',
}: {
  config: LightConfig;
  locale: string;
  tierLabel?: string;
}) {
  const t = useTranslations('escoreLight');
  const router = useRouter();
  const derivedNote = (id: string): string | undefined => {
    if (id === DERIVED_ITEM_IDS.IMC) return t('formDerivedBMI');
    if (id === DERIVED_ITEM_IDS.WAIST_HEIGHT_RATIO) return t('formDerivedWaistHeight');
    return undefined;
  };
  const [step, setStep] = useState<string>(STEP_INTRO);
  const [demo, setDemo] = useState<Demographics>({
    age: '',
    gender: '',
    height: '',
    weight: '',
  });
  const [responses, setResponses] = useState<ResponseMap>({});
  const [submitting, setSubmitting] = useState(false);
  const [submitError, setSubmitError] = useState<string | null>(null);
  const [consentAccepted, setConsentAccepted] = useState(false);

  // Captura UTMs da URL no mount → sessionStorage. Garante atribuição correta
  // mesmo se o usuário entrar via /painel?utm=... e completar o form vários
  // minutos depois (ou navegar pra demais páginas e voltar).
  useEffect(() => {
    captureUTMFromURL();
  }, []);

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

  const clearResponse = (itemId: string) => {
    setResponses((prev) => {
      if (!(itemId in prev)) return prev;
      const next = { ...prev };
      delete next[itemId];
      return next;
    });
  };

  // Scroll ao topo sempre que o step mudar (mais robusto que chamar manualmente
  // em cada handler — garante que o scroll roda APÓS o render do novo conteúdo).
  useEffect(() => {
    if (typeof window !== 'undefined') {
      window.scrollTo({ top: 0, behavior: 'smooth' });
    }
  }, [step]);

  // Sincroniza valores derivados (IMC, Razão cintura/altura) sempre que
  // os inputs base mudam (peso, altura, gênero, cintura).
  const cinturaMaleVal = responses[CINTURA_ITEM_IDS.male]?.numericValue;
  const cinturaFemaleVal = responses[CINTURA_ITEM_IDS.female]?.numericValue;
  useEffect(() => {
    setResponses((prev) => {
      let changed = false;
      const next = { ...prev };
      for (const id of Object.values(DERIVED_ITEM_IDS)) {
        const computed = computeDerivedValue(id, demo, prev);
        const current = prev[id]?.numericValue;
        if (computed !== undefined && computed !== current) {
          next[id] = { scoreItemId: id, numericValue: computed };
          changed = true;
        }
        if (computed === undefined && current !== undefined) {
          delete next[id];
          changed = true;
        }
      }
      return changed ? next : prev;
    });
  }, [demo.height, demo.weight, demo.gender, cinturaMaleVal, cinturaFemaleVal]);

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

  /** Navegação direta para um grupo (via clique na barra de progresso). */
  const goToGroup = (groupId: string) => {
    setStep(groupId);
  };

  const demoIsValid = useMemo(() => {
    if (typeof demo.age !== 'number' || demo.age < 18 || demo.age > 120) return false;
    if (!demo.gender) return false;
    return true;
  }, [demo]);

  const responseCount = useMemo(() => {
    return Object.values(responses).filter(
      (r) =>
        r.selectedLevel !== undefined ||
        r.numericValue !== undefined ||
        (r.textValue !== undefined && r.textValue !== ''),
    ).length;
  }, [responses]);

  async function submit() {
    setSubmitting(true);
    setSubmitError(null);
    try {
      if (typeof demo.age !== 'number' || !demo.gender) {
        throw new Error(t('formErrorIncompleteDemo'));
      }
      // Só envia respostas de fato preenchidas — items em branco ficam fora do cálculo
      const validResponses = Object.values(responses).filter(
        (r) =>
          r.selectedLevel !== undefined ||
          r.numericValue !== undefined ||
          (r.textValue !== undefined && r.textValue !== ''),
      );
      if (validResponses.length === 0) {
        throw new Error(t('formErrorAnswerOne'));
      }
      const utm = utmToPayload(captureUTMFromURL());
      const payload = {
        age: demo.age,
        gender: demo.gender,
        height: typeof demo.height === 'number' ? demo.height : undefined,
        weight: typeof demo.weight === 'number' ? demo.weight : undefined,
        responses: validResponses,
        consentVersion: PRIVACY_POLICY_VERSION,
        ...utm,
      };
      const session = await createSession(payload);
      router.push(`/${locale}/escore-plenya/resultado/${session.publicCode}`);
    } catch (err) {
      setSubmitError(err instanceof Error ? err.message : t('formErrorSubmit'));
      setSubmitting(false);
    }
  }

  // === Renderers ===

  type IconType = ReturnType<typeof getGroupIcon>;
  const renderProgress = () => {
    type ProgressDot = {
      key: string;
      label: string;
      reached: boolean;
      onClick?: () => void;
      Icon?: IconType;
    };

    const dots: ProgressDot[] = [
      { key: 'intro', label: t('formProgressIntro'), reached: currentStepIndex >= 0 },
      { key: 'demo', label: t('formProgressDemo'), reached: currentStepIndex >= 1, onClick: () => setStep(STEP_DEMO) },
      ...groupSteps.map((g, i) => ({
        key: g.id,
        label: g.name,
        reached: currentStepIndex >= 2 + i,
        onClick: () => goToGroup(g.id),
        Icon: getGroupIcon(g.name),
      })),
    ];

    return (
      <div className="flex items-center gap-1.5 mb-12 flex-wrap">
        {dots.map((d) => {
          const isCurrent =
            (d.key === 'intro' && step === STEP_INTRO) ||
            (d.key === 'demo' && step === STEP_DEMO) ||
            d.key === step;
          const Icon = d.Icon;
          return (
            <button
              key={d.key}
              type="button"
              onClick={d.onClick}
              disabled={!d.onClick}
              title={d.label}
              aria-label={d.label}
              aria-current={isCurrent ? 'step' : undefined}
              className={`group relative inline-flex items-center justify-center transition ${
                d.onClick ? 'cursor-pointer' : 'cursor-default'
              }`}
            >
              {Icon ? (
                <span
                  className={`inline-flex items-center justify-center h-7 w-7 rounded-full border transition ${
                    isCurrent
                      ? 'bg-gold text-cream border-gold'
                      : d.reached
                      ? 'bg-gold/20 text-gold border-gold/40 hover:bg-gold/30'
                      : 'bg-paper text-petrol/30 border-petrol/15'
                  }`}
                >
                  <Icon className="w-3.5 h-3.5" strokeWidth={2} />
                </span>
              ) : (
                <span
                  className={`block h-1.5 w-7 rounded-full transition ${
                    isCurrent ? 'bg-gold' : d.reached ? 'bg-gold/60' : 'bg-petrol/15'
                  }`}
                />
              )}
              {/* Tooltip ao passar o mouse */}
              <span
                role="tooltip"
                className="pointer-events-none absolute top-full left-1/2 -translate-x-1/2 mt-2 px-2 py-1 bg-petrol text-cream text-[11px] rounded whitespace-nowrap opacity-0 group-hover:opacity-100 transition-opacity z-10 shadow-lg"
              >
                {d.label}
              </span>
            </button>
          );
        })}
      </div>
    );
  };

  const renderIntro = () => (
    <div className="space-y-8 max-w-2xl">
      <p className="label-upper text-gold">{tierLabel}</p>
      <h1 className="heading-section text-petrol text-3xl md:text-4xl leading-tight">
        {t('formIntroTitle')}
      </h1>
      <p className="text-petrol/80 text-lg leading-relaxed">
        {t('formIntroDescPart1')}{config.itemCount}{t('formIntroDescPart2')}
      </p>
      <ul className="space-y-3 text-petrol/70">
        <li className="flex gap-3">
          <span className="text-gold">—</span>
          <span><strong className="text-petrol">{t('formIntroBullet1Strong')}</strong>{t('formIntroBullet1Rest')}</span>
        </li>
        <li className="flex gap-3">
          <span className="text-gold">—</span>
          <span>{t('formIntroBullet2')}</span>
        </li>
        <li className="flex gap-3">
          <span className="text-gold">—</span>
          <span>{t('formIntroBullet3')}</span>
        </li>
        <li className="flex gap-3">
          <span className="text-gold">—</span>
          <span>{t('formIntroBullet4')}</span>
        </li>
      </ul>
      {/* Consentimento LGPD obrigatório (art. 11 §1º — dados sensíveis exigem consentimento específico) */}
      <div className="border border-petrol/15 bg-paper rounded-md p-5 mt-2">
        <label className="flex gap-3 items-start cursor-pointer group">
          <input
            type="checkbox"
            checked={consentAccepted}
            onChange={(e) => setConsentAccepted(e.target.checked)}
            className="mt-1 w-4 h-4 accent-gold cursor-pointer"
          />
          <span className="text-petrol/80 text-sm leading-relaxed">
            {t('formConsentPart1')}
            <a href={`/${locale}/privacidade`} target="_blank" rel="noopener noreferrer" className="text-gold underline underline-offset-4">
              {t('formConsentPrivacyLink')}
            </a>
            {t('formConsentPart2')}
            <a href={`/${locale}/termos`} target="_blank" rel="noopener noreferrer" className="text-gold underline underline-offset-4">
              {t('formConsentTermsLink')}
            </a>
            {t('formConsentPart3')}
            <strong>{t('formConsentNotDiagnosis')}</strong>
            {t('formConsentPart4')}
          </span>
        </label>
      </div>

      <div className="pt-2">
        <button
          onClick={goNext}
          disabled={!consentAccepted}
          className={`btn-gold ${!consentAccepted ? 'opacity-40 cursor-not-allowed' : ''}`}
          title={!consentAccepted ? t('formIntroConsentRequired') : undefined}
        >
          {t('formIntroStart')}
        </button>
      </div>
    </div>
  );

  const renderDemographics = () => (
    <div className="space-y-8 max-w-xl">
      <p className="label-upper text-gold">{t('formDemoLabel')}</p>
      <h2 className="heading-section text-petrol text-2xl md:text-3xl">
        {t('formDemoTitle')}
      </h2>

      <div className="space-y-6">
        <label className="block">
          <span className="block text-petrol/80 mb-2">{t('formDemoAge')}</span>
          <input
            type="number"
            min={18}
            max={120}
            value={demo.age}
            onChange={(e) =>
              setDemo({ ...demo, age: e.target.value === '' ? '' : Number(e.target.value) })
            }
            className="w-full border border-petrol/20 bg-cream px-4 py-3 text-petrol focus:border-gold focus:outline-none"
            placeholder={t('formDemoAgePlaceholder')}
          />
        </label>

        <fieldset>
          <legend className="block text-petrol/80 mb-3">{t('formDemoGender')}</legend>
          <div className="flex gap-3">
            {(['male', 'female', 'other'] as const).map((g) => (
              <button
                key={g}
                type="button"
                onClick={() => setDemo({ ...demo, gender: g })}
                className={`px-5 py-3 border transition ${
                  demo.gender === g
                    ? 'bg-petrol text-cream border-petrol'
                    : 'bg-cream text-petrol border-petrol/20 hover:border-petrol/50'
                }`}
              >
                {g === 'male'
                  ? t('formDemoGenderMale')
                  : g === 'female'
                  ? t('formDemoGenderFemale')
                  : t('formDemoGenderOther')}
              </button>
            ))}
          </div>
        </fieldset>

        <div className="grid grid-cols-2 gap-4">
          <label className="block">
            <span className="block text-petrol/80 mb-2">{t('formDemoHeight')}</span>
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
              placeholder={t('formDemoHeightPlaceholder')}
            />
          </label>
          <label className="block">
            <span className="block text-petrol/80 mb-2">{t('formDemoWeight')}</span>
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
              placeholder={t('formDemoWeightPlaceholder')}
            />
          </label>
        </div>
      </div>

      <div className="flex items-center gap-4 pt-6">
        <button onClick={goBack} className="text-petrol/60 hover:text-petrol transition underline-offset-4 hover:underline">
          {t('formBack')}
        </button>
        <button
          onClick={goNext}
          disabled={!demoIsValid}
          className={`btn-gold ${!demoIsValid ? 'opacity-40 cursor-not-allowed' : ''}`}
        >
          {t('formContinue')}
        </button>
      </div>
    </div>
  );

  const renderGroupStep = (groupId: string) => {
    const group = groupSteps.find((g) => g.id === groupId);
    if (!group) return null;
    const isLastGroup = groupSteps[groupSteps.length - 1].id === groupId;
    const GroupIcon = getGroupIcon(group.name);
    const isLabsGroup = group.name === 'Exames';
    return (
      <div className="space-y-8 max-w-2xl">
        <div className="flex items-center gap-4">
          <span className="inline-flex items-center justify-center w-14 h-14 rounded-full bg-gold/10 text-gold shrink-0">
            <GroupIcon className="w-7 h-7" strokeWidth={1.5} />
          </span>
          <h2 className="heading-section text-petrol text-2xl md:text-3xl">{group.name}</h2>
        </div>
        <p className="text-petrol/60 text-sm">
          {t('formGroupOptional')}
        </p>
        {isLabsGroup && (
          <LabPDFUpload
            onExtracted={(matched: LightLabMatch[]) => {
              // Injeta valores extraídos diretamente em responses
              setResponses((prev) => {
                const next = { ...prev };
                for (const m of matched) {
                  next[m.scoreItemId] = {
                    scoreItemId: m.scoreItemId,
                    numericValue: m.numericValue,
                  };
                }
                return next;
              });
            }}
          />
        )}
        <div className="space-y-10">
          {group.items.map((item) => renderItemQuestion(item))}
        </div>
        <div className="flex items-center gap-4 pt-6">
          <button onClick={goBack} className="text-petrol/60 hover:text-petrol transition underline-offset-4 hover:underline">
            {t('formBack')}
          </button>
          <button
            onClick={goNext}
            disabled={submitting || (isLastGroup && responseCount === 0)}
            className={`btn-gold ${
              submitting || (isLastGroup && responseCount === 0)
                ? 'opacity-40 cursor-not-allowed'
                : ''
            }`}
            title={isLastGroup && responseCount === 0 ? t('formGroupOneRequired') : undefined}
          >
            {submitting
              ? t('formSubmitting')
              : isLastGroup
              ? `${t('formSeeResult')}${responseCount > 0 ? ` (${responseCount})` : ''}`
              : t('formContinue')}
          </button>
        </div>
        {submitError && <p className="text-red-700 text-sm">{submitError}</p>}
      </div>
    );
  };

  const renderItemQuestion = (item: LightItemConfig) => {
    const resp = responses[item.id];
    const isPHQ9 = PHQ9_ITEM_IDS.has(item.id);
    const numericMode = !isPHQ9 && isNumericClassifiable(item);
    const isLevelChoice = !numericMode && item.levels.length > 0 && !item.labTestCode;
    const hasResponse = resp?.selectedLevel !== undefined || resp?.numericValue !== undefined;
    const instructionType = getInstructionType(item.name);

    // PHQ-9 usa widget customizado de 9 perguntas
    if (isPHQ9) {
      return (
        <div key={item.id} className="border-t border-petrol/15 pt-6">
          <div className="flex items-baseline justify-between gap-3 mb-4">
            <p className="text-petrol text-lg leading-relaxed">
              {t('formPhq9Question')} <span className="text-petrol/50 text-base">{t('formPhq9Scale')}</span>
            </p>
            <button
              type="button"
              onClick={() => clearResponse(item.id)}
              disabled={!hasResponse}
              aria-hidden={!hasResponse}
              className={`shrink-0 text-xs underline underline-offset-4 transition ${
                hasResponse
                  ? 'text-petrol/40 hover:text-petrol/80 cursor-pointer'
                  : 'text-transparent pointer-events-none select-none'
              }`}
              title={hasResponse ? t('formClearResponse') : ''}
            >
              {t('formClear')}
            </button>
          </div>
          <PHQ9Widget
            itemId={item.id}
            hasExistingResponse={hasResponse}
            onResult={(level) => setResponse(item.id, { selectedLevel: level })}
            onClear={() => clearResponse(item.id)}
          />
        </div>
      );
    }

    return (
      <div key={item.id} className="border-t border-petrol/15 pt-6">
        <div className="flex items-baseline justify-between gap-3 mb-4">
          <p className="text-petrol text-lg leading-relaxed">{item.lightQuestion}</p>
          <button
            type="button"
            onClick={() => clearResponse(item.id)}
            disabled={!hasResponse}
            aria-hidden={!hasResponse}
            className={`shrink-0 text-xs underline underline-offset-4 transition ${
              hasResponse
                ? 'text-petrol/40 hover:text-petrol/80 cursor-pointer'
                : 'text-transparent pointer-events-none select-none'
            }`}
            title={hasResponse ? t('formClearResponse') : ''}
          >
            {t('formClear')}
          </button>
        </div>
        {instructionType && <ItemInstruction type={instructionType} />}
        {numericMode ? (
          <NumericClassifierInput
            item={item}
            value={resp?.numericValue}
            readOnly={isDerivedItem(item.id)}
            derivedNote={isDerivedItem(item.id) ? derivedNote(item.id) : undefined}
            onChange={(v) => {
              if (v === undefined) {
                clearResponse(item.id);
              } else {
                setResponse(item.id, { numericValue: v });
              }
            }}
          />
        ) : isLevelChoice ? (
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
                    onClick={() => {
                      if (selected) {
                        clearResponse(item.id);
                      } else {
                        setResponse(item.id, { selectedLevel: lv.level });
                      }
                    }}
                    aria-pressed={selected}
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
              onChange={(e) => {
                if (e.target.value === '') {
                  clearResponse(item.id);
                } else {
                  setResponse(item.id, { numericValue: Number(e.target.value) });
                }
              }}
              className="flex-1 border border-petrol/20 bg-cream px-4 py-3 text-petrol focus:border-gold focus:outline-none"
              placeholder={item.unit ?? t('formNumericPlaceholder')}
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
