'use client';

import { MessageCircle } from 'lucide-react';
import { track } from '@/lib/plausible';

export function WhatsAppShare({ title, url }: { title: string; url: string }) {
  const text = `${title} — ${url}`;
  const href = `https://wa.me/?text=${encodeURIComponent(text)}`;
  return (
    <a
      href={href}
      target="_blank"
      rel="noreferrer"
      onClick={() => track('whatsapp_clicou', { source: 'blog-share' })}
      className="btn-outline-dark inline-flex items-center gap-2"
    >
      <MessageCircle size={16} />
      Compartilhar no WhatsApp
    </a>
  );
}
