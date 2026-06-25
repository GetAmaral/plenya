import { PlenyaSymbol } from "@plenya/brand/logo";
import { cn } from "@/lib/utils";

/**
 * PlenyaMark — selo de marca oficial para uso na UI (sidebars, login, headers).
 *
 * Tile petrol arredondado com o símbolo P∞ em dourado, exatamente a paleta do
 * brandbook. Substitui o antigo /logo_infinity.svg (export CorelDRAW em bege
 * fora da marca) e o ícone genérico do login. Renderiza via @plenya/brand/logo,
 * a mesma fonte vetorial usada no site institucional.
 */
export function PlenyaMark({
  className,
  symbolClassName,
  variant = "tile",
}: {
  className?: string;
  symbolClassName?: string;
  /** "tile": símbolo gold sobre fundo petrol. "bare": só o símbolo (herda cor). */
  variant?: "tile" | "bare";
}) {
  if (variant === "bare") {
    return (
      <PlenyaSymbol
        className={cn("h-auto w-full text-gold", symbolClassName, className)}
      />
    );
  }
  return (
    <div
      className={cn(
        "flex items-center justify-center rounded-xl bg-petrol shadow-sm",
        className,
      )}
    >
      <PlenyaSymbol className={cn("h-auto w-1/2 text-gold", symbolClassName)} />
    </div>
  );
}
