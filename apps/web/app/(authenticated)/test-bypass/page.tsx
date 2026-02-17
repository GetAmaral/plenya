"use client";

export default function TestBypassPage() {
  const devBypass = process.env.NEXT_PUBLIC_DEV_BYPASS_AUTH === 'true';

  return (
    <div className="p-8">
      <h1 className="text-2xl font-bold mb-4">Teste Dev Bypass Auth</h1>

      <div className="space-y-4">
        <div className="p-4 border rounded">
          <h2 className="font-semibold">Variável de Ambiente:</h2>
          <pre className="mt-2 bg-gray-100 p-2 rounded">
            NEXT_PUBLIC_DEV_BYPASS_AUTH = "{process.env.NEXT_PUBLIC_DEV_BYPASS_AUTH || 'undefined'}"
          </pre>
        </div>

        <div className="p-4 border rounded">
          <h2 className="font-semibold">Resultado da Verificação:</h2>
          <p className="mt-2 text-lg">
            {devBypass ? (
              <span className="text-green-600 font-bold">✅ BYPASS ATIVO</span>
            ) : (
              <span className="text-red-600 font-bold">❌ BYPASS DESATIVADO</span>
            )}
          </p>
        </div>

        <div className="p-4 border rounded">
          <h2 className="font-semibold">Todas as Variáveis NEXT_PUBLIC_*:</h2>
          <pre className="mt-2 bg-gray-100 p-2 rounded text-sm">
            {JSON.stringify(
              Object.keys(process.env)
                .filter(key => key.startsWith('NEXT_PUBLIC_'))
                .reduce((acc, key) => ({ ...acc, [key]: process.env[key] }), {}),
              null,
              2
            )}
          </pre>
        </div>
      </div>
    </div>
  );
}
