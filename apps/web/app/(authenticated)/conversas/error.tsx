'use client';

import { useEffect } from 'react';

export default function ConversasError({
  error,
  reset,
}: {
  error: Error & { digest?: string };
  reset: () => void;
}) {
  useEffect(() => {
    // Log no console pra DevTools/Firecrawl ler
    console.error('[/conversas] crash:', error);
  }, [error]);

  return (
    <div className="p-6">
      <h2 className="text-lg font-semibold mb-2 text-rose-700">
        Erro ao carregar /conversas
      </h2>
      <pre className="bg-rose-50 border border-rose-200 p-3 rounded text-xs whitespace-pre-wrap wrap-break-word mb-3">
        <strong>{error.name}:</strong> {error.message}
        {error.digest && (
          <>
            {'\n\n'}<strong>digest:</strong> {error.digest}
          </>
        )}
        {error.stack && (
          <>
            {'\n\n'}<strong>stack:</strong>{'\n'}{error.stack.slice(0, 2000)}
          </>
        )}
      </pre>
      <button
        type="button"
        onClick={reset}
        className="rounded-md bg-rose-600 px-4 py-2 text-sm text-white hover:bg-rose-700"
      >
        Tentar novamente
      </button>
    </div>
  );
}
