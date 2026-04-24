'use client';

import { useRequireAuth } from '@/lib/use-auth';
import { usePatientGuard } from '@/lib/use-patient-guard';
import { useConversations } from '@/lib/api/conversations-api';

export default function ConversasPage() {
  useRequireAuth();
  usePatientGuard();

  const { data, isLoading, isError, error } = useConversations({});

  return (
    <div className="p-4">
      <h1 className="text-2xl font-bold mb-4">Conversas (debug)</h1>
      {isLoading && <p>Carregando...</p>}
      {isError && (
        <pre className="bg-red-50 border border-red-200 p-3 rounded text-xs whitespace-pre-wrap break-words">
          {String(error)}
        </pre>
      )}
      {data && (
        <div>
          <p className="text-sm text-muted-foreground mb-2">
            {data.items.length} conversa(s)
          </p>
          <ul className="space-y-2">
            {data.items.map((it) => (
              <li
                key={`${it.ownerType}-${it.ownerId}`}
                className="border rounded p-3 text-sm"
              >
                <div className="font-semibold">{it.name}</div>
                <div className="text-xs text-muted-foreground">
                  {it.ownerType} · {it.email ?? it.phone ?? '-'} · unread={it.unreadCount}
                </div>
                <div className="mt-1 text-xs italic line-clamp-2">{it.lastSnippet}</div>
              </li>
            ))}
          </ul>
        </div>
      )}
    </div>
  );
}
