'use client';

import { useEffect, useRef, useState } from 'react';
import { Save } from 'lucide-react';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { Textarea } from '@/components/ui/textarea';
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';
import { Card, CardContent } from '@/components/ui/card';
import type { Campaign, CreateCampaignInput } from '@/lib/api/campaigns-api';

const LANDING_OPTIONS: { value: string; label: string }[] = [
  { value: '/escore-plenya/painel', label: 'Escore Plenya — Painel ampliado (83 itens)' },
  { value: '/escore-plenya/avaliar', label: 'Escore Plenya — Triagem (público padrão)' },
  { value: '/escore-plenya', label: 'Escore Plenya — página explicativa' },
  { value: '/contato', label: 'Página de contato' },
  { value: '/', label: 'Home' },
];

const SOURCE_HINTS = [
  'instagram',
  'whatsapp',
  'qr-cartao',
  'qr-cartaz',
  'email',
  'youtube',
  'podcast',
  'google',
  'meta-ads',
  'google-ads',
  'evento-presencial',
];

const MEDIUM_HINTS = [
  'stories',
  'post',
  'bio',
  'reel',
  'qr',
  'broadcast',
  'newsletter',
  'paid-ads',
  'organic',
  'flyer',
];

function slugify(s: string): string {
  return s
    .toLowerCase()
    .normalize('NFD')
    .replace(/[̀-ͯ]/g, '')
    .replace(/[^a-z0-9]+/g, '-')
    .replace(/^-+|-+$/g, '');
}

export function CampaignForm({
  initial,
  onSubmit,
  submitting,
  submitLabel = 'Salvar',
}: {
  initial?: Campaign;
  onSubmit: (input: CreateCampaignInput) => void;
  submitting?: boolean;
  submitLabel?: string;
}) {
  const formRef = useRef<HTMLFormElement>(null);
  const [name, setName] = useState(initial?.name ?? '');
  const [slug, setSlug] = useState(initial?.slug ?? '');
  const [slugTouched, setSlugTouched] = useState(!!initial?.slug);
  const [description, setDescription] = useState(initial?.description ?? '');
  const [landingPath, setLandingPath] = useState(
    initial?.landingPath ?? '/escore-plenya/painel',
  );
  const [utmSource, setUtmSource] = useState(initial?.utmSource ?? '');
  const [utmMedium, setUtmMedium] = useState(initial?.utmMedium ?? '');
  const [utmCampaign, setUtmCampaign] = useState(initial?.utmCampaign ?? '');
  const [utmCampaignTouched, setUtmCampaignTouched] = useState(!!initial?.utmCampaign);
  const [utmTerm, setUtmTerm] = useState(initial?.utmTerm ?? '');

  // Auto-derivações: slug ← name; utm_campaign ← slug (até o usuário tocar manual)
  useEffect(() => {
    if (!slugTouched) setSlug(slugify(name));
  }, [name, slugTouched]);

  useEffect(() => {
    if (!utmCampaignTouched) setUtmCampaign(slug);
  }, [slug, utmCampaignTouched]);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    onSubmit({
      name: name.trim(),
      slug: slug.trim() || undefined,
      description: description.trim() || undefined,
      landingPath,
      utmSource: utmSource.trim(),
      utmMedium: utmMedium.trim(),
      utmCampaign: utmCampaign.trim() || undefined,
      utmTerm: utmTerm.trim() || undefined,
    });
  };

  const previewURL = (() => {
    if (!utmSource || !utmMedium) return null;
    const params = new URLSearchParams();
    params.set('utm_source', utmSource);
    params.set('utm_medium', utmMedium);
    params.set('utm_campaign', utmCampaign || slug);
    if (utmTerm) params.set('utm_term', utmTerm);
    return `https://plenyasaude.com.br${landingPath}?${params.toString()}`;
  })();

  return (
    <form ref={formRef} onSubmit={handleSubmit} className="space-y-6">
      <Card>
        <CardContent className="pt-6 space-y-4">
          <div>
            <Label htmlFor="name">Nome interno *</Label>
            <Input
              id="name"
              value={name}
              onChange={(e) => setName(e.target.value)}
              placeholder="Stories Insta — Janeiro 2026"
              required
              maxLength={160}
            />
            <p className="text-xs text-muted-foreground mt-1">Visível só no CRM.</p>
          </div>

          <div>
            <Label htmlFor="slug">Slug</Label>
            <Input
              id="slug"
              value={slug}
              onChange={(e) => {
                setSlug(slugify(e.target.value));
                setSlugTouched(true);
              }}
              placeholder="auto-gerado do nome"
              maxLength={120}
            />
          </div>

          <div>
            <Label htmlFor="description">Descrição (opcional)</Label>
            <Textarea
              id="description"
              value={description}
              onChange={(e) => setDescription(e.target.value)}
              placeholder="Briefing, audiência, peças usadas, prazo…"
              rows={3}
            />
          </div>

          <div>
            <Label htmlFor="landingPath">Página de destino *</Label>
            <Select value={landingPath} onValueChange={setLandingPath}>
              <SelectTrigger id="landingPath">
                <SelectValue />
              </SelectTrigger>
              <SelectContent>
                {LANDING_OPTIONS.map((o) => (
                  <SelectItem key={o.value} value={o.value}>
                    {o.label}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
          </div>
        </CardContent>
      </Card>

      <Card>
        <CardContent className="pt-6 space-y-4">
          <h3 className="font-medium">Atribuição UTM</h3>
          <p className="text-sm text-muted-foreground">
            Esses valores entram na URL e ficam gravados na sessão do Escore + Lead. Use
            sempre os mesmos vocabulários pra agrupar relatórios.
          </p>
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div>
              <Label htmlFor="utmSource">utm_source *</Label>
              <Input
                id="utmSource"
                value={utmSource}
                onChange={(e) => setUtmSource(e.target.value.toLowerCase().replace(/\s+/g, '-'))}
                list="utm-source-hints"
                placeholder="instagram"
                required
                maxLength={80}
              />
              <datalist id="utm-source-hints">
                {SOURCE_HINTS.map((h) => (
                  <option key={h} value={h} />
                ))}
              </datalist>
            </div>
            <div>
              <Label htmlFor="utmMedium">utm_medium *</Label>
              <Input
                id="utmMedium"
                value={utmMedium}
                onChange={(e) => setUtmMedium(e.target.value.toLowerCase().replace(/\s+/g, '-'))}
                list="utm-medium-hints"
                placeholder="stories"
                required
                maxLength={80}
              />
              <datalist id="utm-medium-hints">
                {MEDIUM_HINTS.map((h) => (
                  <option key={h} value={h} />
                ))}
              </datalist>
            </div>
            <div>
              <Label htmlFor="utmCampaign">utm_campaign</Label>
              <Input
                id="utmCampaign"
                value={utmCampaign}
                onChange={(e) => {
                  setUtmCampaign(e.target.value);
                  setUtmCampaignTouched(true);
                }}
                placeholder="default = slug"
                maxLength={120}
              />
            </div>
            <div>
              <Label htmlFor="utmTerm">utm_term (opcional)</Label>
              <Input
                id="utmTerm"
                value={utmTerm}
                onChange={(e) => setUtmTerm(e.target.value)}
                placeholder="variação A/B, criativo, etc."
                maxLength={120}
              />
            </div>
          </div>
        </CardContent>
      </Card>

      {previewURL && (
        <Card>
          <CardContent className="pt-6">
            <Label className="text-xs uppercase tracking-wider text-muted-foreground">
              Preview da URL
            </Label>
            <p className="mt-2 font-mono text-xs break-all bg-muted p-3 rounded-md">
              {previewURL}
            </p>
          </CardContent>
        </Card>
      )}

      <div className="flex gap-3">
        <Button type="submit" disabled={submitting || !name || !utmSource || !utmMedium}>
          <Save className="mr-2 h-4 w-4" />
          {submitting ? 'Salvando…' : submitLabel}
        </Button>
      </div>
    </form>
  );
}
