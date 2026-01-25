import { create } from "zustand";

interface LoadingState {
  isNavigating: boolean;
  setNavigating: (isNavigating: boolean) => void;
}

export const useLoadingState = create<LoadingState>((set) => ({
  isNavigating: false,
  setNavigating: (isNavigating) => set({ isNavigating }),
}));
