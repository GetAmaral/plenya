'use client';

import { useState } from 'react';
import { Play } from 'lucide-react';
import { cn } from '@/lib/cn';

type Provider = 'bunny' | 'cloudflare' | 'youtube' | 'mp4';

type Props = {
  provider: Provider;
  videoId: string;
  poster?: string;
  title: string;
  className?: string;
  aspect?: 'video' | 'square';
};

function buildEmbedUrl(provider: Provider, videoId: string): string {
  switch (provider) {
    case 'bunny':
      return `https://iframe.mediadelivery.net/embed/${videoId}?autoplay=true`;
    case 'cloudflare':
      return `https://customer-${videoId.split(':')[0]}.cloudflarestream.com/${videoId.split(':')[1]}/iframe?autoplay=true`;
    case 'youtube':
      return `https://www.youtube-nocookie.com/embed/${videoId}?autoplay=1&rel=0&modestbranding=1`;
    case 'mp4':
      return videoId;
  }
}

export function VideoPlayer({ provider, videoId, poster, title, className, aspect = 'video' }: Props) {
  const [playing, setPlaying] = useState(false);

  return (
    <div
      className={cn(
        'relative overflow-hidden rounded-2xl border border-petrol/15 bg-petrol',
        aspect === 'square' ? 'aspect-square' : 'aspect-video',
        className,
      )}
    >
      {!playing ? (
        <button
          type="button"
          onClick={() => setPlaying(true)}
          aria-label={`Reproduzir vídeo: ${title}`}
          className="group relative h-full w-full"
        >
          {poster ? (
            <img
              src={poster}
              alt={title}
              loading="lazy"
              className="absolute inset-0 h-full w-full object-cover transition-transform duration-700 group-hover:scale-[1.02]"
            />
          ) : (
            <div className="absolute inset-0 bg-linear-to-br from-petrol via-petrol-light to-petrol" />
          )}
          <div className="absolute inset-0 bg-petrol/30 group-hover:bg-petrol/20 transition" />
          <div className="absolute inset-0 flex items-center justify-center">
            <span className="inline-flex h-20 w-20 items-center justify-center rounded-full bg-gold text-petrol shadow-lg shadow-petrol/40 group-hover:scale-110 transition">
              <Play size={28} fill="currentColor" className="ml-1" />
            </span>
          </div>
          <div className="absolute bottom-6 left-6 right-6 flex items-end justify-between gap-4">
            <p className="font-mono text-[11px] uppercase tracking-widest text-cream/90">{title}</p>
          </div>
        </button>
      ) : provider === 'mp4' ? (
        <video
          src={videoId}
          controls
          autoPlay
          poster={poster}
          className="absolute inset-0 h-full w-full"
        />
      ) : (
        <iframe
          src={buildEmbedUrl(provider, videoId)}
          title={title}
          loading="lazy"
          allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture"
          allowFullScreen
          className="absolute inset-0 h-full w-full"
        />
      )}
    </div>
  );
}
