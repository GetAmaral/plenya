'use client';

import { useState } from 'react';
import { useRouter } from 'next/navigation';
import { Link } from '@/lib/i18n/navigation';
import { ScoreRadarChart } from './score-radar-chart';
import { PhoneInput } from '@/components/forms/phone-input';
import { requestClaim, deleteSession, exportSessionURL } from '@/lib/score-light/api';
import { PRIVACY_POLICY_VERSION } from '@/lib/legal';
import type { PublicSession } from '@/lib/score-light/types';

function scoreLabel(pct: number): string {
  if (pct >= 80) return 'Muito bom';
  if (pct >= 60) return 'Bom — com pontos a fortalecer';
  if (pct >= 40) return 'Médio — vale revisar';
  return 'Atenção — vale uma conversa';
}

export function EscoreLightResultado({
  session,
  locale,
}: {
  session: PublicSession;
  locale: string;
}) {
  const snapshot = session.snapshot;
  const router = useRouter();
  const [wantEmail, setWantEmail] = useState(true);
  const [wantWhatsApp, setWantWhatsApp] = useState(false);
  const [email, setEmail] = useState('');
  const [phoneE164, setPhoneE164] = useState('');
  const [newsletterOptIn, setNewsletterOptIn] = useState(false);
  const [claimStatus, setClaimStatus] = useState<'idle' | 'sending' | 'sent' | 'error'>(
    'idle'
  );
  const [claimError, setClaimError] = useState<string | null>(null);
  const [deleteStatus, setDeleteStatus] = useState<'idle' | 'confirming' | 'deleting' | 'done' | 'error'>('idle');
  const [deleteError, setDeleteError] = useState<string | null>(null);

  async function handleDelete() {
    setDeleteStatus('deleting');
    setDeleteError(null);
    try {
      await deleteSession(session.publicCode);
      setDeleteStatus('done');
      // Redireciona para home após 2s
      setTimeout(() => router.push(`/${locale}`), 2000);
    } catch (err) {
      setDeleteError(err instanceof Error ? err.message : 'Falha ao excluir');
      setDeleteStatus('error');
    }
  }

  if (!snapshot) {
    return (
      <section className="bg-cream">
        <div className="site-container pt-32 pb-24 max-w-xl">
          <p className="label-upper text-gold">Resultado indisponível</p>
          <h1 className="heading-section text-petrol text-3xl mt-6">
            Não foi possível calcular seu radar.
          </h1>
          <p className="text-petrol/80 mt-6 leading-relaxed">
            Pode ter havido respostas insuficientes para gerar a avaliação. Tente
            refazer a autoavaliação.
          </p>
          <div className="pt-8">
            <Link href="/escore-plenya/avaliar" className="btn-gold">
              Refazer
            </Link>
          </div>
        </div>
      </section>
    );
  }

  const totalPct = snapshot.totalScorePercentage;
  const radarData = snapshot.groupResults
    .filter((g) => g.itemsEvaluatedCount > 0)
    .map((g) => ({ label: g.groupName, value: g.scorePercentage }));

  const emailValid = wantEmail ? /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email) : true;
  // WhatsApp exige celular (11 dígitos com 9° dígito = 9). PhoneInput já filtra isso
  // — se phoneE164 estiver não-vazio, é móvel válido.
  const phoneValid = wantWhatsApp ? /^\+55\d{2}9\d{8}$/.test(phoneE164) : true;
  const canSubmit = (wantEmail || wantWhatsApp) && emailValid && phoneValid;

  async function handleClaim(e: React.FormEvent) {
    e.preventDefault();
    if (!canSubmit) return;
    setClaimStatus('sending');
    setClaimError(null);
    try {
      await requestClaim(session.publicCode, {
        email: wantEmail ? email : undefined,
        phone: wantWhatsApp ? phoneE164 : undefined,
        newsletterOptIn,
        consentVersion: PRIVACY_POLICY_VERSION,
      });
      setClaimStatus('sent');
    } catch (err) {
      setClaimError(err instanceof Error ? err.message : 'Falha ao enviar');
      setClaimStatus('error');
    }
  }

  return (
    <>
      {/* Hero — pontuação total */}
      <section className="bg-petrol text-cream">
        <div className="site-container pt-32 pb-20 md:pt-40 md:pb-24 max-w-3xl">
          <p className="label-upper text-gold">Seu Escore Plenya Light</p>
          <p className="text-cream/60 text-sm mt-6">Pontuação geral</p>
          <div className="flex items-baseline gap-3 mt-2 whitespace-nowrap">
            <span className="text-7xl md:text-8xl font-light text-cream tabular-nums">
              {Math.round(totalPct)}
            </span>
            <span className="text-2xl text-cream/60">/ 100</span>
          </div>
          <p className="text-cream/80 text-xl mt-4">{scoreLabel(totalPct)}</p>
          <p className="text-cream/60 text-sm mt-8 max-w-xl leading-relaxed">
            {snapshot.itemsEvaluatedCount} de {snapshot.itemsEvaluatedCount + snapshot.itemsNotEvaluatedCount}{' '}
            itens avaliados. Os pilares com pontuação mais baixa são onde a equipe
            Plenya costuma começar.
          </p>
        </div>
      </section>

      {/* Radar */}
      <section className="bg-cream">
        <div className="site-container section grid lg:grid-cols-[1fr_1fr] gap-16 items-center">
          <div>
            <p className="label-upper text-gold">Radar por pilar</p>
            <h2 className="heading-section text-petrol text-2xl md:text-3xl mt-4">
              Cada eixo é um grupo clínico do Escore Plenya.
            </h2>
            <p className="text-petrol/70 mt-6 leading-relaxed">
              Quanto mais próximo da borda, mais pontos você somou naquele pilar. As
              áreas mais retraídas indicam onde há margem de evolução.
            </p>
          </div>
          <div className="flex justify-center">
            {radarData.length >= 3 ? (
              <ScoreRadarChart data={radarData} size={400} />
            ) : (
              <p className="text-petrol/60">
                Dados insuficientes para radar. Necessário ao menos 3 pilares avaliados.
              </p>
            )}
          </div>
        </div>
      </section>

      {/* Detalhamento por grupo */}
      <section className="bg-paper">
        <div className="site-container section">
          <p className="label-upper text-gold mb-10">Detalhamento</p>
          <div className="border-t border-petrol/15">
            {snapshot.groupResults
              .filter((g) => g.itemsEvaluatedCount > 0)
              .map((g) => (
                <div
                  key={g.groupId}
                  className="grid grid-cols-[1fr_auto] gap-6 py-6 border-b border-petrol/10 items-center"
                >
                  <div>
                    <p className="text-petrol text-lg">{g.groupName}</p>
                    <p className="text-petrol/60 text-sm mt-1">
                      {g.itemsEvaluatedCount} {g.itemsEvaluatedCount === 1 ? 'item' : 'itens'}
                    </p>
                  </div>
                  <div className="text-right whitespace-nowrap">
                    <span className="text-petrol text-3xl font-light tabular-nums">
                      {Math.round(g.scorePercentage)}
                    </span>
                    <span className="text-petrol/50 text-base ml-1">/ 100</span>
                  </div>
                </div>
              ))}
          </div>
        </div>
      </section>

      {/* Salvar resultado (claim opcional) */}
      <section className="bg-cream">
        <div className="site-container section max-w-2xl">
          <p className="label-upper text-gold">Salvar para acompanhar</p>
          <h2 className="heading-section text-petrol text-2xl md:text-3xl mt-4">
            Quer guardar este resultado e refazer em três meses?
          </h2>
          <p className="text-petrol/70 mt-6 leading-relaxed">
            Informe seu email — enviamos um link único para você acessar seu resultado
            sempre que quiser e comparar a evolução. Sem cadastro, sem senha.
          </p>

          {claimStatus === 'sent' ? (
            <div className="mt-8 border border-gold/40 bg-gold/5 p-6 text-petrol">
              <p className="font-medium">Link enviado.</p>
              <p className="text-petrol/70 text-sm mt-2">
                {wantEmail && wantWhatsApp
                  ? `Você receberá o link por email (${email}) e WhatsApp em alguns minutos.`
                  : wantEmail
                    ? `Verifique seu email (${email}) — pode chegar na pasta de spam.`
                    : 'Você receberá o link por WhatsApp em alguns minutos.'}
              </p>
            </div>
          ) : (
            <form onSubmit={handleClaim} className="mt-8 space-y-5">
              <fieldset className="space-y-4">
                <legend className="text-petrol/70 text-sm mb-2">Como prefere receber?</legend>

                <div className="space-y-2">
                  <label className="flex items-start gap-3 cursor-pointer">
                    <input
                      type="checkbox"
                      checked={wantEmail}
                      onChange={(e) => setWantEmail(e.target.checked)}
                      className="mt-1 h-4 w-4 accent-gold cursor-pointer"
                    />
                    <span className="text-petrol">Email</span>
                  </label>
                  {wantEmail && (
                    <input
                      type="email"
                      value={email}
                      onChange={(e) => setEmail(e.target.value)}
                      placeholder="seu@email.com"
                      className="ml-7 w-[calc(100%-1.75rem)] border border-petrol/20 bg-cream px-4 py-3 text-petrol focus:border-gold focus:outline-none"
                    />
                  )}
                </div>

                <div className="space-y-2">
                  <label className="flex items-start gap-3 cursor-pointer">
                    <input
                      type="checkbox"
                      checked={wantWhatsApp}
                      onChange={(e) => setWantWhatsApp(e.target.checked)}
                      className="mt-1 h-4 w-4 accent-gold cursor-pointer"
                    />
                    <span className="text-petrol">WhatsApp</span>
                  </label>
                  {wantWhatsApp && (
                    <>
                      <PhoneInput
                        value={phoneE164}
                        onChange={setPhoneE164}
                        className="ml-7 w-[calc(100%-1.75rem)] border border-petrol/20 bg-cream px-4 py-3 text-petrol focus:border-gold focus:outline-none"
                      />
                      <p className="ml-7 text-xs text-petrol/55 leading-relaxed">
                        Ao informar WhatsApp, seus dados são processados pela Meta (EUA) sob cláusulas-padrão de
                        proteção. Detalhes na <Link href="/privacidade" className="underline underline-offset-4">Política de Privacidade</Link>.
                      </p>
                    </>
                  )}
                </div>
              </fieldset>

              <label className="flex items-start gap-3 cursor-pointer pt-2 border-t border-petrol/10">
                <input
                  type="checkbox"
                  checked={newsletterOptIn}
                  onChange={(e) => setNewsletterOptIn(e.target.checked)}
                  className="mt-1 h-4 w-4 accent-gold cursor-pointer"
                />
                <span className="text-petrol/80 text-sm leading-relaxed">
                  Quero receber a newsletter Plenya — conteúdos sobre longevidade e medicina personalizada (opcional, pode cancelar quando quiser).
                </span>
              </label>

              <button
                type="submit"
                disabled={claimStatus === 'sending' || !canSubmit}
                className={`btn-gold w-full ${
                  claimStatus === 'sending' || !canSubmit ? 'opacity-40 cursor-not-allowed' : ''
                }`}
              >
                {claimStatus === 'sending' ? 'Enviando…' : 'Receber meu link'}
              </button>
            </form>
          )}
          {claimError && <p className="text-red-700 text-sm mt-3">{claimError}</p>}

          <p className="text-petrol/50 text-xs mt-6 leading-relaxed">
            Seus dados de contato serão usados apenas para enviar o link de acesso (e a newsletter, se você marcou). Você
            pode salvar esta URL diretamente nos favoritos como alternativa.
          </p>
        </div>
      </section>

      {/* CTAs finais */}
      <section className="bg-petrol text-cream">
        <div className="site-container section max-w-3xl">
          <p className="label-upper text-gold">Próximos passos</p>
          <h2 className="heading-section text-cream text-2xl md:text-3xl mt-4">
            O Light é um ponto de partida — não um diagnóstico.
          </h2>
          <p className="text-cream/80 mt-6 leading-relaxed">
            Se quiser conversar sobre o que esse radar mostra, a equipe Plenya está
            disponível. A versão completa do Escore — com exames, anamnese aprofundada
            e análise da equipe multidisciplinar — vive dentro do Continuum.
          </p>
          <div className="flex flex-wrap gap-4 mt-10">
            <Link href="/contato" className="btn-gold">
              Conversar com a equipe
            </Link>
            <Link
              href="/continuum"
              className="px-6 py-3 border border-cream/40 text-cream hover:bg-cream/10 transition"
            >
              Conhecer o Continuum
            </Link>
          </div>
        </div>
      </section>

      {/* LGPD — direito de eliminação (art. 18 VI) */}
      <section className="bg-cream border-t border-petrol/10">
        <div className="site-container section max-w-2xl">
          <p className="label-upper text-petrol/45">LGPD · Seus direitos</p>
          <h2 className="text-petrol text-xl mt-3">Quer apagar este resultado?</h2>
          <p className="text-petrol/70 text-sm mt-3 leading-relaxed">
            Você pode excluir esta sessão e todos os dados associados a ela a
            qualquer momento. A exclusão é definitiva — não há como recuperar
            depois. Para outros direitos (acesso, correção, portabilidade), use a{' '}
            <Link href="/lgpd/direitos" className="text-gold underline underline-offset-4">
              página de direitos LGPD
            </Link>.
          </p>

          {deleteStatus === 'idle' && (
            <div className="mt-5 flex flex-wrap gap-6 items-center">
              <a
                href={exportSessionURL(session.publicCode)}
                download
                className="text-petrol/70 hover:text-petrol text-sm underline underline-offset-4 transition"
              >
                Baixar meus dados (JSON)
              </a>
              <button
                type="button"
                onClick={() => setDeleteStatus('confirming')}
                className="text-petrol/60 hover:text-red-700 text-sm underline underline-offset-4 transition"
              >
                Excluir esta sessão
              </button>
            </div>
          )}

          {deleteStatus === 'confirming' && (
            <div className="mt-5 border border-red-200 bg-red-50 p-4 rounded-md space-y-3">
              <p className="text-red-900 text-sm">
                Tem certeza? Esta ação é definitiva e remove imediatamente seu
                radar, suas respostas e qualquer dado associado a esta sessão.
              </p>
              <div className="flex gap-3">
                <button
                  type="button"
                  onClick={handleDelete}
                  className="px-4 py-2 bg-red-700 text-cream rounded text-sm hover:bg-red-800 transition"
                >
                  Sim, excluir definitivamente
                </button>
                <button
                  type="button"
                  onClick={() => setDeleteStatus('idle')}
                  className="px-4 py-2 border border-petrol/30 text-petrol rounded text-sm hover:bg-petrol/5 transition"
                >
                  Cancelar
                </button>
              </div>
            </div>
          )}

          {deleteStatus === 'deleting' && (
            <p className="mt-5 text-petrol/70 text-sm">Excluindo…</p>
          )}

          {deleteStatus === 'done' && (
            <div className="mt-5 border border-emerald-200 bg-emerald-50 p-4 rounded-md text-emerald-900 text-sm">
              ✓ Sessão excluída. Você será redirecionado em alguns segundos.
            </div>
          )}

          {deleteStatus === 'error' && deleteError && (
            <div className="mt-5 border border-red-200 bg-red-50 p-4 rounded-md text-red-900 text-sm">
              {deleteError}
              <button
                type="button"
                onClick={() => setDeleteStatus('idle')}
                className="ml-3 underline underline-offset-4"
              >
                tentar de novo
              </button>
            </div>
          )}
        </div>
      </section>
    </>
  );
}
