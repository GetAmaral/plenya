'use client';

import { useRouter } from 'next/navigation';
import Link from 'next/link';
import { ArrowLeft } from 'lucide-react';

import { useRequireAuth } from '@/lib/use-auth';
import { Button } from '@/components/ui/button';
import { CampaignForm } from '@/components/campaigns/campaign-form';
import { useCreateCampaign, type CreateCampaignInput } from '@/lib/api/campaigns-api';
import { toast } from 'sonner';

export default function NewCampaignPage() {
  useRequireAuth();
  const router = useRouter();
  const create = useCreateCampaign();

  const onSubmit = (input: CreateCampaignInput) => {
    create.mutate(input, {
      onSuccess: (c) => {
        toast.success('Campanha criada');
        router.push(`/campaigns/${c.id}`);
      },
      onError: (err) => {
        toast.error('Erro ao criar', {
          description: err instanceof Error ? err.message : 'Tente novamente',
        });
      },
    });
  };

  return (
    <div className="space-y-6 max-w-3xl">
      <Link href="/campaigns">
        <Button variant="ghost" size="sm">
          <ArrowLeft className="mr-2 h-4 w-4" /> Voltar
        </Button>
      </Link>
      <div>
        <h1 className="text-2xl font-semibold">Nova campanha</h1>
        <p className="text-sm text-muted-foreground mt-1">
          UTMs definem como leads desta ação serão identificados no CRM.
        </p>
      </div>
      <CampaignForm onSubmit={onSubmit} submitting={create.isPending} />
    </div>
  );
}
