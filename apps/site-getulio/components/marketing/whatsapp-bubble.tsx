'use client';

import { MessageCircle } from 'lucide-react';

// WhatsApp da consulta (Plenya · longevidade). Número fixo da clínica.
const PHONE = '5543999748899';

export function WhatsAppBubble({ label = 'WhatsApp' }: { label?: string }) {
  return (
    <a
      href={`https://wa.me/${PHONE}`}
      target="_blank"
      rel="noreferrer"
      aria-label={label}
      className="fixed bottom-6 right-6 z-50 inline-flex h-14 w-14 items-center justify-center rounded-full bg-gold text-paper shadow-lg shadow-gold/30 ring-1 ring-gold/40 transition hover:brightness-110"
    >
      <MessageCircle size={22} />
    </a>
  );
}
