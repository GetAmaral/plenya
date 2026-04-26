import { useState } from 'react';
import { Alert, ScrollView, TextInput, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { router, useLocalSearchParams } from 'expo-router';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { options, queryKeys, leadSourceLabels, leadStatusLabels } from '@plenya/api-client';
import {
  Button,
  Card,
  CardHeader,
  Dialog,
  ErrorState,
  FormField,
  Sheet,
  Spinner,
  Text,
  useToast,
} from '@plenya/ui-mobile';
import { formatRelative } from '@plenya/domain';

type ReplyChannel = 'whatsapp' | 'email';

export default function LeadDetailScreen() {
  const { id } = useLocalSearchParams<{ id: string }>();
  const leadId = id ?? '';
  const queryClient = useQueryClient();
  const toast = useToast();

  const lead = useQuery(options.leadOptions(leadId));

  const [composeOpen, setComposeOpen] = useState<ReplyChannel | null>(null);
  const [composeText, setComposeText] = useState('');
  const [confirmConvert, setConfirmConvert] = useState(false);

  const sendReply = useMutation({
    mutationFn: (vars: { channel: ReplyChannel; content: string }) =>
      options.leadMutations.sendReply(leadId, vars),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.leads.detail(leadId) });
      queryClient.invalidateQueries({ queryKey: queryKeys.leads.all() });
      setComposeOpen(null);
      setComposeText('');
      toast.show('Mensagem enviada', 'success');
    },
    onError: (err) => toast.show(err instanceof Error ? err.message : 'Falha ao enviar', 'error'),
  });

  const convert = useMutation({
    mutationFn: () => options.leadMutations.convert(leadId),
    onSuccess: (data) => {
      queryClient.invalidateQueries({ queryKey: queryKeys.leads.all() });
      toast.show('Lead convertido em paciente', 'success');
      setConfirmConvert(false);
      router.replace(`/(tabs)/patients/${data.patientId}` as never);
    },
    onError: (err) => toast.show(err instanceof Error ? err.message : 'Falha ao converter', 'error'),
  });

  if (lead.isLoading) return <Spinner centered />;
  if (lead.isError) {
    return (
      <SafeAreaView className="flex-1 bg-background">
        <ErrorState onRetry={() => lead.refetch()} />
      </SafeAreaView>
    );
  }
  if (!lead.data) return null;

  const l = lead.data;
  const alreadyConverted = Boolean(l.convertedPatientId);

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <ScrollView contentContainerClassName="gap-3 p-4">
        <Text variant="heading">{l.name ?? l.email ?? l.phone ?? 'Lead'}</Text>
        <Text variant="caption">
          {leadSourceLabels[l.source] ?? l.source} ·{' '}
          {leadStatusLabels[l.status] ?? l.status}
        </Text>

        <Card>
          <CardHeader>
            <Text variant="title">Contato</Text>
          </CardHeader>
          <Text variant="body">{l.phone ?? '—'}</Text>
          <Text variant="caption">{l.email ?? '—'}</Text>
        </Card>

        <Button
          onPress={() => router.push(`/(tabs)/conversations/lead/${leadId}` as never)}
          fullWidth
          variant="outline"
        >
          Abrir conversa (com anexos + IA)
        </Button>

        <View className="flex-row gap-2">
          {l.phone && (
            <Button
              variant="outline"
              size="sm"
              onPress={() => setComposeOpen('whatsapp')}
              className="flex-1"
            >
              Resposta rápida WA
            </Button>
          )}
          {l.email && (
            <Button
              variant="outline"
              size="sm"
              onPress={() => setComposeOpen('email')}
              className="flex-1"
            >
              Resposta rápida email
            </Button>
          )}
        </View>

        {alreadyConverted ? (
          <Button
            variant="ghost"
            onPress={() => router.push(`/(tabs)/patients/${l.convertedPatientId}` as never)}
            fullWidth
          >
            Ver paciente convertido
          </Button>
        ) : (
          <Button onPress={() => setConfirmConvert(true)} fullWidth>
            Converter em paciente
          </Button>
        )}

        <Card>
          <CardHeader>
            <Text variant="title">Timeline</Text>
          </CardHeader>
          <View className="gap-3">
            {l.activities.map((a) => (
              <View key={a.id}>
                <Text variant="body">{a.content}</Text>
                <Text variant="caption">
                  {a.kind} · {formatRelative(a.createdAt)}
                  {a.actorName ? ` · ${a.actorName}` : ''}
                </Text>
              </View>
            ))}
            {l.activities.length === 0 && <Text variant="caption">Sem atividades</Text>}
          </View>
        </Card>
      </ScrollView>

      <Sheet open={composeOpen !== null} onClose={() => setComposeOpen(null)}>
        <Text variant="title" className="mb-3">
          {composeOpen === 'whatsapp' ? 'Mensagem WhatsApp' : 'Mensagem por email'}
        </Text>
        <FormField label={composeOpen === 'email' ? 'Corpo do email' : 'Mensagem'}>
          <View className="rounded-lg border border-border bg-background p-3">
            <TextInput
              value={composeText}
              onChangeText={setComposeText}
              multiline
              textAlignVertical="top"
              placeholder="Digite sua resposta..."
              placeholderTextColor="#9CA3AF"
              editable={!sendReply.isPending}
              style={{ minHeight: 140, fontSize: 15, color: '#0f172a' }}
            />
          </View>
        </FormField>
        <View className="mt-3 flex-row gap-2">
          <Button
            variant="outline"
            onPress={() => setComposeOpen(null)}
            className="flex-1"
            disabled={sendReply.isPending}
          >
            Cancelar
          </Button>
          <Button
            onPress={() => {
              if (!composeText.trim() || !composeOpen) return;
              sendReply.mutate({ channel: composeOpen, content: composeText.trim() });
            }}
            loading={sendReply.isPending}
            className="flex-1"
          >
            Enviar
          </Button>
        </View>
      </Sheet>

      <Dialog open={confirmConvert} onClose={() => setConfirmConvert(false)}>
        <Text variant="title" className="mb-2">
          Converter em paciente?
        </Text>
        <Text variant="body" className="mb-4">
          Esta ação cria um cadastro de paciente vinculado ao lead. Não é reversível pelo app.
        </Text>
        <View className="flex-row gap-2">
          <Button
            variant="outline"
            onPress={() => setConfirmConvert(false)}
            disabled={convert.isPending}
            className="flex-1"
          >
            Cancelar
          </Button>
          <Button
            onPress={() => convert.mutate()}
            loading={convert.isPending}
            className="flex-1"
          >
            Converter
          </Button>
        </View>
      </Dialog>
    </SafeAreaView>
  );
}
