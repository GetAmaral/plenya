import DOMPurify from "isomorphic-dompurify";

/**
 * sanitizeHtml — ponto ÚNICO de sanitização de HTML antes de injetar via SafeHtml.
 * Usa a config padrão do DOMPurify (remove <script>, handlers on*, javascript: URLs),
 * que é segura para os relatórios HTML gerados (treino/anamnese/consulta) sem quebrar
 * estilos inline legítimos. Ver docs/emr/estudo-xss-hardening.md.
 */
export function sanitizeHtml(html: string | null | undefined): string {
  if (!html) return "";
  return DOMPurify.sanitize(html);
}

/**
 * safeUrl — valida URLs vindas de dado não-confiável (links de mensagens, campanhas) antes de
 * usar em href/src. Bloqueia esquemas perigosos (javascript:, data:, vbscript:, file:) que
 * permitiriam XSS ao clicar. Retorna undefined quando a URL não é segura.
 */
export function safeUrl(url?: string | null): string | undefined {
  if (!url) return undefined;
  const t = url.trim();
  if (!t) return undefined;
  // Esquemas perigosos — barra explicitamente.
  if (/^(javascript|data|vbscript|file):/i.test(t)) return undefined;
  // Relativos, âncoras, query, protocol-relative → seguros (mesmo site).
  if (/^(\/|#|\?|\.)/.test(t) || t.startsWith("//")) return t;
  // Absolutos: http(s), mailto, tel; blob: (object URLs gerados pelo próprio app — mídia/download).
  if (/^(https?:\/\/|mailto:|tel:|blob:)/i.test(t)) return t;
  // Sem esquema reconhecido → descarta (evita interpretação ambígua).
  return undefined;
}
