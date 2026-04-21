'use client';

import { useState } from 'react';
import { Link } from '@/lib/i18n/navigation';
import { ScoreRadarChart } from './score-radar-chart';
import { requestClaim } from '@/lib/score-light/api';
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
  const [email, setEmail] = useState('');
  const [claimStatus, setClaimStatus] = useState<'idle' | 'sending' | 'sent' | 'error'>(
    'idle'
  );
  const [claimError, setClaimError] = useState<string | null>(null);

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

  async function handleClaim(e: React.FormEvent) {
    e.preventDefault();
    if (!email) return;
    setClaimStatus('sending');
    setClaimError(null);
    try {
      await requestClaim(session.publicCode, email);
      setClaimStatus('sent');
    } catch (err) {
      setClaimError(err instanceof Error ? err.message : 'Falha ao enviar email');
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
          <div className="flex items-baseline gap-3 mt-2">
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
                  <div className="text-right">
                    <span className="heading-section text-petrol text-2xl tabular-nums">
                      {Math.round(g.scorePercentage)}
                    </span>
                    <span className="text-petrol/50">/ 100</span>
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
              <p className="font-medium">Link enviado para {email}.</p>
              <p className="text-petrol/70 text-sm mt-2">
                Se a sessão estiver ativa, o email chegará em alguns minutos. Verifique
                também o spam.
              </p>
            </div>
          ) : (
            <form onSubmit={handleClaim} className="mt-8 flex flex-col sm:flex-row gap-3">
              <input
                type="email"
                required
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                placeholder="seu@email.com"
                className="flex-1 border border-petrol/20 bg-cream px-4 py-3 text-petrol focus:border-gold focus:outline-none"
              />
              <button
                type="submit"
                disabled={claimStatus === 'sending' || !email}
                className={`btn-gold ${
                  claimStatus === 'sending' || !email ? 'opacity-40 cursor-not-allowed' : ''
                }`}
              >
                {claimStatus === 'sending' ? 'Enviando...' : 'Receber meu link'}
              </button>
            </form>
          )}
          {claimError && <p className="text-red-700 text-sm mt-3">{claimError}</p>}

          <p className="text-petrol/50 text-xs mt-6 leading-relaxed">
            Seu email será usado apenas para enviar o link de acesso. Você pode salvar
            esta URL diretamente nos favoritos como alternativa.
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
    </>
  );
}
