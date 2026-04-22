/**
 * Widget customizado para a Escala PHQ-9 (humor).
 * Renderiza as 9 perguntas com escala 0-3, soma automaticamente,
 * classifica e mapeia para o `selectedLevel` do item Light correspondente.
 *
 * Mapeamento PHQ-9 → níveis no DB (item c77cedd3-2800-721b-94d8-b5515b895753):
 *   N0: ≥ 20 (depressão grave)
 *   N1: 15-19 (moderada-grave)
 *   N2: 10-14 (moderada)
 *   N3: 5-9 (leve)
 *   N4: 2-4 (mínima)
 *   N5: ≤ 1 (ausente)
 */
import { useState, useEffect } from 'react';

// PHQ-9 oficial (versão validada em PT-BR)
const QUESTIONS = [
  'Pouco interesse ou prazer em fazer as coisas',
  'Sentir-se triste, deprimido(a) ou sem perspectiva',
  'Dificuldade para pegar no sono ou continuar dormindo, ou dormir demais',
  'Sentir-se cansado(a) ou com pouca energia',
  'Falta de apetite ou comer demais',
  'Sentir-se mal consigo mesmo(a) — ou achar que é um fracasso ou que decepcionou sua família ou a si mesmo(a)',
  'Dificuldade para se concentrar nas coisas (como ler o jornal ou ver televisão)',
  'Lentidão para se mover ou falar (a ponto de outras pessoas perceberem) — ou o oposto: estar tão agitado(a) que tem se mexido bem mais que de costume',
  'Pensar em se ferir de alguma forma ou que seria melhor estar morto(a)',
];

const ANSWERS = [
  { value: 0, label: 'Nenhuma vez' },
  { value: 1, label: 'Vários dias' },
  { value: 2, label: 'Mais da metade dos dias' },
  { value: 3, label: 'Quase todos os dias' },
];

function classify(total: number): { level: number; label: string; color: string } {
  if (total >= 20) return { level: 0, label: 'Depressão grave', color: 'text-red-700' };
  if (total >= 15) return { level: 1, label: 'Moderadamente grave', color: 'text-orange-600' };
  if (total >= 10) return { level: 2, label: 'Moderada', color: 'text-amber-600' };
  if (total >= 5) return { level: 3, label: 'Leve', color: 'text-yellow-600' };
  if (total >= 2) return { level: 4, label: 'Mínima', color: 'text-emerald-600' };
  return { level: 5, label: 'Ausente', color: 'text-emerald-700' };
}

export function PHQ9Widget({
  itemId,
  onResult,
  onClear,
  hasExistingResponse,
}: {
  itemId: string;
  onResult: (level: number) => void;
  onClear: () => void;
  hasExistingResponse: boolean;
}) {
  const [answers, setAnswers] = useState<Record<number, number>>({});

  // Reset interno quando o usuário limpa pelo botão "limpar" do form
  useEffect(() => {
    if (!hasExistingResponse && Object.keys(answers).length > 0) {
      setAnswers({});
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [hasExistingResponse]);

  const answeredCount = Object.keys(answers).length;
  const allAnswered = answeredCount === QUESTIONS.length;
  const total = Object.values(answers).reduce((s, v) => s + v, 0);
  const result = allAnswered ? classify(total) : null;

  const choose = (qIdx: number, value: number) => {
    setAnswers((prev) => {
      // toggle: clicar de novo na mesma resposta desmarca
      if (prev[qIdx] === value) {
        const next = { ...prev };
        delete next[qIdx];
        // se tirou a última, limpa o item no form
        if (Object.keys(next).length === 0) {
          onClear();
        }
        return next;
      }
      return { ...prev, [qIdx]: value };
    });
  };

  // Quando 9/9 respondidas, registra o level no form pai
  useEffect(() => {
    if (allAnswered && result) {
      onResult(result.level);
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [allAnswered, total]);

  return (
    <div className="space-y-5">
      <p className="text-petrol/70 text-sm leading-relaxed">
        Nas últimas 2 semanas, com que frequência você foi incomodado(a) por algum dos
        problemas abaixo?
      </p>

      <div className="space-y-3">
        {QUESTIONS.map((q, qIdx) => (
          <div key={qIdx} className="border border-petrol/15 rounded-md bg-cream p-4">
            <p className="text-petrol leading-snug mb-3">
              <span className="text-gold/80 mr-2">{qIdx + 1}.</span>
              {q}
            </p>
            <div className="grid grid-cols-2 sm:grid-cols-4 gap-2">
              {ANSWERS.map((a) => {
                const selected = answers[qIdx] === a.value;
                return (
                  <button
                    key={a.value}
                    type="button"
                    onClick={() => choose(qIdx, a.value)}
                    aria-pressed={selected}
                    className={`text-left text-sm px-3 py-2 border transition leading-snug ${
                      selected
                        ? 'bg-petrol text-cream border-petrol'
                        : 'bg-paper text-petrol border-petrol/15 hover:border-petrol/40'
                    }`}
                  >
                    {a.label}
                  </button>
                );
              })}
            </div>
          </div>
        ))}
      </div>

      {/* Resultado */}
      <div className="border-t border-petrol/15 pt-4 flex items-center justify-between gap-4">
        <div>
          <p className="label-upper text-petrol/55 text-[10px]">Pontuação</p>
          <p className="text-petrol text-xl tabular-nums mt-1">
            {answeredCount === 0
              ? '— / 27'
              : allAnswered
              ? `${total} / 27`
              : `${answeredCount} de 9 respondidas`}
          </p>
        </div>
        {result && (
          <div className="text-right">
            <p className="label-upper text-petrol/55 text-[10px]">Classificação</p>
            <p className={`text-xl mt-1 ${result.color}`}>{result.label}</p>
          </div>
        )}
      </div>
    </div>
  );
}

// Lista de IDs de items que devem usar o widget PHQ-9
export const PHQ9_ITEM_IDS = new Set([
  'c77cedd3-2800-721b-94d8-b5515b895753',
]);
