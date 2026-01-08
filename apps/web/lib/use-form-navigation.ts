import { useEffect, RefObject } from "react";

interface UseFormNavigationOptions {
  /**
   * Ref do formulário onde aplicar a navegação
   */
  formRef: RefObject<HTMLFormElement>;

  /**
   * Se true, desabilita a navegação automática
   * @default false
   */
  disabled?: boolean;

  /**
   * Callback chamado quando Enter é pressionado no último campo
   * Se não fornecido, faz submit automático do formulário
   */
  onLastFieldEnter?: () => void;
}

/**
 * Hook para navegação de formulário com Enter
 *
 * Comportamento:
 * - Enter em qualquer campo (exceto textarea) move para o próximo campo
 * - Enter no último campo faz submit do formulário
 * - Tab continua funcionando normalmente
 *
 * @example
 * ```tsx
 * const formRef = useRef<HTMLFormElement>(null);
 * useFormNavigation({ formRef });
 *
 * return (
 *   <form ref={formRef} onSubmit={handleSubmit}>
 *     <input name="field1" />
 *     <input name="field2" />
 *     <button type="submit">Salvar</button>
 *   </form>
 * );
 * ```
 */
export function useFormNavigation({
  formRef,
  disabled = false,
  onLastFieldEnter,
}: UseFormNavigationOptions) {
  useEffect(() => {
    if (disabled || !formRef.current) return;

    const form = formRef.current;

    const handleKeyDown = (event: KeyboardEvent) => {
      // Só processa Enter
      if (event.key !== "Enter") return;

      const target = event.target as HTMLElement;

      // Não processa se for textarea (Enter deve quebrar linha)
      // Não processa se for botão (deixa o comportamento padrão)
      if (
        target.tagName === "TEXTAREA" ||
        target.tagName === "BUTTON" ||
        target.getAttribute("type") === "submit"
      ) {
        return;
      }

      // Só processa se for um elemento de input
      const isInputElement =
        target.tagName === "INPUT" ||
        target.tagName === "SELECT" ||
        target.getAttribute("role") === "combobox" || // Para Select do shadcn
        target.getAttribute("contenteditable") === "true";

      if (!isInputElement) return;

      // Previne o submit padrão
      event.preventDefault();

      // Pega todos os elementos focáveis do formulário
      const focusableElements = Array.from(
        form.querySelectorAll<HTMLElement>(
          'input:not([disabled]):not([type="hidden"]), select:not([disabled]), textarea:not([disabled]), button[type="submit"]:not([disabled]), [role="combobox"]:not([disabled])'
        )
      );

      // Encontra o índice do elemento atual
      const currentIndex = focusableElements.findIndex(
        (el) => el === target || el.contains(target)
      );

      if (currentIndex === -1) return;

      // Próximo elemento
      const nextIndex = currentIndex + 1;

      // Se existe próximo elemento e não é botão de submit, foca nele
      if (nextIndex < focusableElements.length) {
        const nextElement = focusableElements[nextIndex];

        // Se o próximo é botão de submit, significa que é o último campo
        if (nextElement.getAttribute("type") === "submit") {
          // Chegou no último campo, faz submit
          if (onLastFieldEnter) {
            onLastFieldEnter();
          } else {
            // Dispara o submit do formulário
            form.requestSubmit();
          }
        } else {
          // Foca no próximo campo
          nextElement.focus();

          // Se for input de texto, seleciona o conteúdo
          if (
            nextElement instanceof HTMLInputElement &&
            (nextElement.type === "text" ||
              nextElement.type === "email" ||
              nextElement.type === "tel" ||
              nextElement.type === "number")
          ) {
            nextElement.select();
          }
        }
      } else {
        // Não há próximo elemento, chegou no último campo
        if (onLastFieldEnter) {
          onLastFieldEnter();
        } else {
          // Dispara o submit do formulário
          form.requestSubmit();
        }
      }
    };

    // Adiciona listener ao formulário
    form.addEventListener("keydown", handleKeyDown);

    // Cleanup
    return () => {
      form.removeEventListener("keydown", handleKeyDown);
    };
  }, [formRef, disabled, onLastFieldEnter]);
}
