import { create } from 'zustand';
import type { ConversationOwnerType } from './api/conversations-api';

export type DockSelection = { type: ConversationOwnerType; id: string } | null;

/**
 * Estado do dock global de WhatsApp. Vive FORA da árvore de rotas (store em memória),
 * então sobrevive à navegação entre telas do EMR — abrir o dock, responder e continuar
 * trabalhando sem perder o contexto. `openWith` permite outras telas (lead, paciente)
 * abrirem direto numa conversa.
 */
interface WhatsAppDockState {
  isOpen: boolean;
  selected: DockSelection;
  open: () => void;
  close: () => void;
  toggle: () => void;
  openWith: (type: ConversationOwnerType, id: string) => void;
  setSelected: (sel: DockSelection) => void;
}

export const useWhatsAppDock = create<WhatsAppDockState>((set) => ({
  isOpen: false,
  selected: null,
  open: () => set({ isOpen: true }),
  close: () => set({ isOpen: false }),
  toggle: () => set((s) => ({ isOpen: !s.isOpen })),
  openWith: (type, id) => set({ isOpen: true, selected: { type, id } }),
  setSelected: (selected) => set({ selected }),
}));
