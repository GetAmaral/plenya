'use client';

import { MessageCircle } from 'lucide-react';
import { track } from '@/lib/plausible';

const PHONE = process.env.NEXT_PUBLIC_WHATSAPP_NUMBER || '5543999999999';

export function WhatsAppBubble() {
  return (
    <a
      href={`https://wa.me/${PHONE}`}
      target="_blank"
      rel="noreferrer"
      onClick={() => track('whatsapp_clicou', { source: 'bubble' })}
      aria-label="WhatsApp Plenya"
      className="fixed bottom-6 right-6 z-50 inline-flex h-14 w-14 items-center justify-center rounded-full bg-petrol text-gold shadow-lg shadow-petrol/30 ring-1 ring-gold/40 hover:bg-petrol-light transition"
    >
      <MessageCircle size={22} />
    </a>
  );
}
