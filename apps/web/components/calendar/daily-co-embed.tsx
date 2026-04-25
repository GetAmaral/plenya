'use client';

/**
 * DailyCoEmbed — embed inline da sala Daily.co pra teleconsulta.
 *
 * Daily.co prebuilt iframe é responsive nativo (não precisa SDK extra).
 * Permissões `camera; microphone; fullscreen` são essenciais — sem elas o
 * usuário não consegue ativar mídia.
 *
 * Uso:
 *   <DailyCoEmbed roomURL={appointment.dailyRoomUrl} />
 */
import { useState } from 'react';
import { Video, X } from 'lucide-react';
import { Button } from '@/components/ui/button';

interface DailyCoEmbedProps {
  roomURL: string;
  /** Opcional: callback quando user clica "Sair" */
  onLeave?: () => void;
}

export function DailyCoEmbed({ roomURL, onLeave }: DailyCoEmbedProps) {
  const [hidden, setHidden] = useState(false);

  if (hidden) {
    return (
      <div className="rounded-lg border border-dashed p-6 text-center">
        <Video className="mx-auto mb-3 h-8 w-8 text-muted-foreground" />
        <p className="mb-3 text-sm text-muted-foreground">
          Sala de teleconsulta encerrada
        </p>
        <Button variant="outline" size="sm" onClick={() => setHidden(false)}>
          Reabrir sala
        </Button>
      </div>
    );
  }

  return (
    <div className="relative w-full overflow-hidden rounded-lg border bg-stone-950">
      <iframe
        src={roomURL}
        title="Teleconsulta Plenya"
        allow="camera; microphone; fullscreen; speaker; display-capture; autoplay"
        className="h-[600px] w-full md:h-[640px]"
      />
      <Button
        type="button"
        size="sm"
        variant="secondary"
        className="absolute right-3 top-3 z-10 gap-1 shadow-lg"
        onClick={() => {
          setHidden(true);
          onLeave?.();
        }}
      >
        <X className="h-3.5 w-3.5" />
        Sair
      </Button>
    </div>
  );
}
