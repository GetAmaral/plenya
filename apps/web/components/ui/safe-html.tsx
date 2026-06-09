import * as React from "react";
import { sanitizeHtml } from "@/lib/security";

type SafeHtmlProps = {
  html: string | null | undefined;
  /** Elemento wrapper (default: div). */
  as?: React.ElementType;
} & Omit<React.HTMLAttributes<HTMLElement>, "dangerouslySetInnerHTML" | "children">;

/**
 * SafeHtml — ÚNICO ponto autorizado a injetar HTML no DOM. Sempre sanitiza via sanitizeHtml.
 * O lint `react/no-danger` é erro em todo o resto do código pra forçar o uso deste componente
 * (sem buraco de XSS por sink novo). Ver docs/emr/estudo-xss-hardening.md.
 */
export function SafeHtml({ html, as: Tag = "div", ...props }: SafeHtmlProps) {
  return (
    <Tag
      {...props}
      // eslint-disable-next-line react/no-danger -- sanitizado por sanitizeHtml; sink central
      dangerouslySetInnerHTML={{ __html: sanitizeHtml(html) }}
    />
  );
}
