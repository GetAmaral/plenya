'use client';

/**
 * Formulário público de exercício dos direitos do titular (LGPD art. 18).
 *
 * Strategy: usa mailto: como entrega — abre o cliente de email do usuário
 * com payload pré-preenchido para o DPO. Sem backend dedicado.
 *
 * Vantagens: zero infra, zero risco de coleta indevida de dados, registro
 * fica no email do próprio titular (transparência).
 */
import { useState } from 'react';
import { useTranslations } from 'next-intl';

const RIGHT_IDS = ['access', 'correct', 'delete', 'anonymize', 'portability', 'sharing', 'revoke', 'other'] as const;
type RightId = (typeof RIGHT_IDS)[number];

const RIGHT_KEY: Record<RightId, { labelKey: string; descKey: string }> = {
  access: { labelKey: 'rightAccessLabel', descKey: 'rightAccessDesc' },
  correct: { labelKey: 'rightCorrectLabel', descKey: 'rightCorrectDesc' },
  delete: { labelKey: 'rightDeleteLabel', descKey: 'rightDeleteDesc' },
  anonymize: { labelKey: 'rightAnonymizeLabel', descKey: 'rightAnonymizeDesc' },
  portability: { labelKey: 'rightPortabilityLabel', descKey: 'rightPortabilityDesc' },
  sharing: { labelKey: 'rightSharingLabel', descKey: 'rightSharingDesc' },
  revoke: { labelKey: 'rightRevokeLabel', descKey: 'rightRevokeDesc' },
  other: { labelKey: 'rightOtherLabel', descKey: 'rightOtherDesc' },
};

export function LGPDRightsForm({ dpoEmail }: { dpoEmail: string }) {
  const t = useTranslations('lgpdRights');
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [right, setRight] = useState<string>('');
  const [details, setDetails] = useState('');
  const [sessionCode, setSessionCode] = useState('');

  const isValid = name.trim() !== '' && email.trim() !== '' && right !== '';

  function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    if (!isValid) return;

    const rightLabel = right in RIGHT_KEY ? t(RIGHT_KEY[right as RightId].labelKey) : right;

    const subject = t('emailSubject', { right: rightLabel, name });
    const body = `${t('emailHeader')}

${t('emailRightLabel')}: ${rightLabel}

${t('emailRequester')}:
- ${t('emailRequesterName')}: ${name}
- ${t('emailRequesterEmail')}: ${email}
${sessionCode ? `- ${t('emailSessionCode')}: ${sessionCode}\n` : ''}
${t('emailDetails')}:
${details || t('emailDetailsNone')}

---
${t('emailFooter')}
`;

    window.location.href = `mailto:${dpoEmail}?subject=${encodeURIComponent(subject)}&body=${encodeURIComponent(body)}`;
  }

  return (
    <form onSubmit={handleSubmit} className="space-y-6 border border-petrol/15 bg-paper p-6 rounded-md">
      <div>
        <label className="block text-petrol/80 mb-2 text-sm font-medium">
          {t('formRightQuestion')} <span className="text-red-600">*</span>
        </label>
        <div className="space-y-2">
          {RIGHT_IDS.map((id) => (
            <label key={id} className="flex gap-3 items-start cursor-pointer hover:bg-cream/60 p-2 rounded-sm transition">
              <input
                type="radio"
                name="right"
                value={id}
                checked={right === id}
                onChange={() => setRight(id)}
                className="mt-1 accent-gold"
              />
              <span>
                <span className="text-petrol font-medium block">{t(RIGHT_KEY[id].labelKey)}</span>
                <span className="text-petrol/60 text-sm">{t(RIGHT_KEY[id].descKey)}</span>
              </span>
            </label>
          ))}
        </div>
      </div>

      <div className="grid sm:grid-cols-2 gap-4">
        <div>
          <label className="block text-petrol/80 mb-2 text-sm font-medium">
            {t('formNameLabel')} <span className="text-red-600">*</span>
          </label>
          <input
            type="text"
            required
            value={name}
            onChange={(e) => setName(e.target.value)}
            className="w-full border border-petrol/20 bg-cream px-4 py-3 text-petrol focus:border-gold focus:outline-hidden"
          />
        </div>
        <div>
          <label className="block text-petrol/80 mb-2 text-sm font-medium">
            {t('formEmailLabel')} <span className="text-red-600">*</span>
          </label>
          <input
            type="email"
            required
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            className="w-full border border-petrol/20 bg-cream px-4 py-3 text-petrol focus:border-gold focus:outline-hidden"
            placeholder={t('formEmailPlaceholder')}
          />
        </div>
      </div>

      <div>
        <label className="block text-petrol/80 mb-2 text-sm font-medium">
          {t('formSessionCodeLabel')}
        </label>
        <input
          type="text"
          value={sessionCode}
          onChange={(e) => setSessionCode(e.target.value)}
          placeholder={t('formSessionCodePlaceholder')}
          className="w-full border border-petrol/20 bg-cream px-4 py-3 text-petrol focus:border-gold focus:outline-hidden font-mono text-sm"
        />
        <p className="text-petrol/55 text-xs mt-1">
          {t('formSessionCodeHelpPart1')}
          <code>/resultado/<strong>jG5RAMX2089a</strong></code>
          {t('formSessionCodeHelpPart2')}
        </p>
      </div>

      <div>
        <label className="block text-petrol/80 mb-2 text-sm font-medium">
          {t('formDetailsLabel')}
        </label>
        <textarea
          value={details}
          onChange={(e) => setDetails(e.target.value)}
          rows={4}
          className="w-full border border-petrol/20 bg-cream px-4 py-3 text-petrol focus:border-gold focus:outline-hidden"
          placeholder={t('formDetailsPlaceholder')}
        />
      </div>

      <div className="bg-cream border-l-2 border-gold pl-3 py-2 text-petrol/70 text-xs">
        {t('formNotice')}
      </div>

      <div>
        <button
          type="submit"
          disabled={!isValid}
          className={`btn-gold ${!isValid ? 'opacity-40 cursor-not-allowed' : ''}`}
        >
          {t('formSubmit')}
        </button>
      </div>
    </form>
  );
}
