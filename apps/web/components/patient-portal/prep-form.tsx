"use client";

/**
 * PrepForm — renderizador do formulário de preparação pré-consulta.
 *
 * Reusa o mesmo estilo da Triagem pública (grupos -> itens, botões de nível /
 * input numérico / texto), porém sem next-intl e dirigido pela config curada do
 * backend (ScoreItems com PrepOrder != nil). Os itens serão curados depois; até
 * lá a config vem vazia e o formulário mostra um estado neutro.
 */
import type { PrepConfig, PrepItemConfig, PrepResponse } from "@/lib/api/patient-portal-api";
import { cn } from "@/lib/utils";

export type PrepAnswers = Record<string, PrepResponse>;

function levelLabel(lv: PrepItemConfig["levels"][number]): string {
  if (lv.name && lv.name.trim() !== "") return lv.name;
  if (lv.operator === "between" && lv.lowerLimit && lv.upperLimit) {
    return `${lv.lowerLimit} a ${lv.upperLimit}`;
  }
  if (lv.lowerLimit && [">", ">=", "<", "<="].includes(lv.operator)) {
    return `${lv.operator} ${lv.lowerLimit}`;
  }
  return lv.lowerLimit ?? String(lv.level);
}

export function PrepForm({
  config,
  answers,
  onChange,
}: {
  config: PrepConfig;
  answers: PrepAnswers;
  onChange: (next: PrepAnswers) => void;
}) {
  const setAnswer = (scoreItemId: string, patch: Partial<PrepResponse> | null) => {
    const next = { ...answers };
    if (patch === null) {
      delete next[scoreItemId];
    } else {
      next[scoreItemId] = { scoreItemId, ...patch };
    }
    onChange(next);
  };

  const groups = [...config.groups].sort((a, b) => a.order - b.order);

  return (
    <div className="space-y-8">
      {groups.map((g) => {
        const items = g.subgroups
          .flatMap((sg) => sg.items)
          .sort((a, b) => a.lightOrder - b.lightOrder);
        if (items.length === 0) return null;
        return (
          <section key={g.id} className="space-y-4">
            <h3 className="text-base font-semibold text-foreground">{g.name}</h3>
            <div className="space-y-5">
              {items.map((item) => (
                <PrepItem
                  key={item.id}
                  item={item}
                  value={answers[item.id]}
                  onChange={(patch) => setAnswer(item.id, patch)}
                />
              ))}
            </div>
          </section>
        );
      })}
    </div>
  );
}

function PrepItem({
  item,
  value,
  onChange,
}: {
  item: PrepItemConfig;
  value: PrepResponse | undefined;
  onChange: (patch: Partial<PrepResponse> | null) => void;
}) {
  const hasLevels = item.levels && item.levels.length > 0;
  const isNumeric = !hasLevels && !!item.unit;

  return (
    <div className="rounded-lg border border-border p-4">
      <p className="text-sm font-medium text-foreground">{item.lightQuestion || item.name}</p>
      {item.patientExplanation && (
        <p className="mt-1 text-xs text-muted-foreground">{item.patientExplanation}</p>
      )}

      <div className="mt-3">
        {hasLevels && (
          <div className="grid grid-cols-1 gap-2 sm:grid-cols-2">
            {[...item.levels]
              .sort((a, b) => a.level - b.level)
              .map((lv) => {
                const selected = value?.selectedLevel === lv.level;
                return (
                  <button
                    key={lv.id}
                    type="button"
                    onClick={() =>
                      onChange(selected ? null : { selectedLevel: lv.level })
                    }
                    className={cn(
                      "rounded-lg border px-3 py-2 text-left text-sm transition-colors",
                      selected
                        ? "border-primary bg-primary/10 text-primary"
                        : "border-border hover:bg-accent",
                    )}
                  >
                    {levelLabel(lv)}
                  </button>
                );
              })}
          </div>
        )}

        {isNumeric && (
          <div className="flex items-center gap-2">
            <input
              type="number"
              inputMode="decimal"
              value={value?.numericValue ?? ""}
              onChange={(e) => {
                const v = e.target.value;
                onChange(v === "" ? null : { numericValue: Number(v) });
              }}
              placeholder="Valor"
              className="w-40 rounded-lg border border-border bg-background px-3 py-2 text-sm"
            />
            {item.unit && <span className="text-sm text-muted-foreground">{item.unit}</span>}
          </div>
        )}

        {!hasLevels && !isNumeric && (
          <textarea
            value={value?.textValue ?? ""}
            onChange={(e) => {
              const v = e.target.value;
              onChange(v === "" ? null : { textValue: v });
            }}
            rows={2}
            placeholder="Sua resposta"
            className="w-full rounded-lg border border-border bg-background px-3 py-2 text-sm"
          />
        )}
      </div>
    </div>
  );
}
