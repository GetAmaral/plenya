'use client';

import { useEffect, useRef, useState } from 'react';
import { sanitizeEmailHtml } from '@/lib/security';

/**
 * EmailHtmlView — renderiza o corpo HTML de um e-mail inbound com fidelidade, mas
 * isolado do app. Dupla barreira de segurança:
 *   1. sanitizeEmailHtml remove <script>/<iframe>/<form>/handlers (allowlist).
 *   2. <iframe sandbox> SEM allow-same-origin → o documento roda em origem opaca,
 *      sem acesso a cookies, localStorage ou DOM do app; o CSS do e-mail também
 *      não vaza pra fora do iframe.
 *
 * Auto-altura: um script TRUSTED (nosso, não do e-mail) injetado no wrapper mede
 * scrollHeight e faz postMessage; o pai ajusta a altura (cap p/ não explodir).
 * allow-scripts é seguro aqui porque o único script presente é o nosso — os do
 * e-mail foram removidos na sanitização.
 */
export function EmailHtmlView({
  html,
  className,
  maxHeight = 900,
}: {
  html: string;
  className?: string;
  maxHeight?: number;
}) {
  const ref = useRef<HTMLIFrameElement>(null);
  const [height, setHeight] = useState(140);

  useEffect(() => {
    function onMessage(e: MessageEvent) {
      const win = ref.current?.contentWindow;
      if (!win || e.source !== win) return;
      if (e.data && e.data.t === 'pl-email-h') {
        const h = Number(e.data.h);
        if (h > 0) setHeight(Math.min(Math.ceil(h) + 8, maxHeight));
      }
    }
    window.addEventListener('message', onMessage);
    return () => window.removeEventListener('message', onMessage);
  }, [maxHeight]);

  const srcDoc = buildDocument(sanitizeEmailHtml(html));

  return (
    <iframe
      ref={ref}
      title="Corpo do e-mail"
      // sem allow-same-origin de propósito: origem opaca isola cookies/DOM/CSS.
      sandbox="allow-scripts allow-popups allow-popups-to-escape-sandbox"
      srcDoc={srcDoc}
      className={className}
      style={{ width: '100%', height, border: 0, background: '#fff', borderRadius: 6, display: 'block' }}
    />
  );
}

function buildDocument(bodyHtml: string): string {
  // Script de medição: nosso, confiável. Reporta a altura no load, em resize e via
  // ResizeObserver (imagens que carregam depois reajustam).
  const resizeScript = `(function(){
    function send(){try{parent.postMessage({t:'pl-email-h',h:document.documentElement.scrollHeight},'*');}catch(e){}}
    window.addEventListener('load',send);
    window.addEventListener('resize',send);
    if(window.ResizeObserver){try{new ResizeObserver(send).observe(document.documentElement);}catch(e){}}
    document.addEventListener('load',send,true);
    setTimeout(send,250);setTimeout(send,800);setTimeout(send,2000);
  })();`;

  return `<!doctype html><html><head><meta charset="utf-8">
<base target="_blank">
<style>
  html,body{margin:0;padding:12px;background:#fff;color:#0a0a0a;
    font:14px/1.55 -apple-system,BlinkMacSystemFont,"Segoe UI",Roboto,Helvetica,Arial,sans-serif;
    word-break:break-word;overflow-wrap:anywhere;}
  *{max-width:100%;box-sizing:border-box;}
  img{max-width:100%!important;height:auto;}
  table{max-width:100%!important;border-collapse:collapse;}
  a{color:#417e8e;}
  blockquote{margin:0 0 0 .5rem;padding-left:.75rem;border-left:3px solid #e5e5e5;color:#555;}
  pre{white-space:pre-wrap;word-break:break-word;}
</style></head>
<body>${bodyHtml}<script>${resizeScript}</script></body></html>`;
}
