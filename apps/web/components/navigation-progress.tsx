"use client";

import { useEffect } from "react";
import { usePathname, useSearchParams } from "next/navigation";
import NProgress from "nprogress";
import "nprogress/nprogress.css";

// Configuração do NProgress
NProgress.configure({
  showSpinner: false,
  trickleSpeed: 200,
  minimum: 0.08,
});

export function NavigationProgress() {
  const pathname = usePathname();
  const searchParams = useSearchParams();

  useEffect(() => {
    // Adiciona cursor wait ao body durante navegação
    const handleStart = () => {
      NProgress.start();
      document.body.style.cursor = "wait";
    };

    const handleComplete = () => {
      NProgress.done();
      document.body.style.cursor = "default";
    };

    // Simula início de navegação quando pathname muda
    handleStart();

    // Completa após render
    const timer = setTimeout(handleComplete, 100);

    return () => {
      clearTimeout(timer);
      handleComplete();
    };
  }, [pathname, searchParams]);

  return null;
}
