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
 * sanitizeEmailHtml — sanitização para corpo de E-MAIL inbound (HTML arbitrário,
 * fonte não-confiável). Mais permissiva que sanitizeHtml (preserva tabelas, imagens
 * e estilos inline que e-mails reais usam pra layout), mas remove qualquer vetor
 * ativo: <script>, <iframe>, <object>/<embed>, <form> e controles, <base>/<meta>/<link>
 * (evita redirecionar âncoras / injetar CSS externo), além dos handlers on* e URLs
 * javascript: que o DOMPurify já tira por padrão.
 *
 * IMPORTANTE: o resultado é renderizado SEMPRE dentro de um <iframe sandbox> (sem
 * allow-same-origin) pelo EmailHtmlView — dupla barreira: mesmo que algo escape da
 * allowlist, o sandbox isola origem, cookies e DOM do app.
 */
export function sanitizeEmailHtml(html: string | null | undefined): string {
  if (!html) return "";
  return DOMPurify.sanitize(html, {
    FORBID_TAGS: [
      "script", "iframe", "object", "embed", "form", "input", "button",
      "textarea", "select", "base", "meta", "link", "title", "noscript",
    ],
    // target é controlado pelo <base target="_blank"> do wrapper (abre em nova aba).
    FORBID_ATTR: ["target", "ping"],
  });
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
