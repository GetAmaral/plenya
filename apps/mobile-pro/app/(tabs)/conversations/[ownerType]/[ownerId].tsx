import { useEffect, useState } from 'react';
import { FlatList, KeyboardAvoidingView, Platform, TextInput, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { useLocalSearchParams } from 'expo-router';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import {
  options,
  type ConversationMessage,
  type ConversationOwnerType,
} from '@plenya/api-client';
import {
  Button,
  Card,
  ErrorState,
  Sheet,
  Spinner,
  Text,
  useToast,
} from '@plenya/ui-mobile';
import { formatRelative } from '@plenya/domain';
import { conversationKeysFor } from '@plenya/api-client/options/conversations';
import { useScreenCaptureProtection } from '../../../../lib/security/screenCapture';

export default function ConversationDetailScreen() {
  useScreenCaptureProtection();
  const { ownerType, ownerId } = useLocalSearchParams<{
    ownerType: ConversationOwnerType;
    ownerId: string;
  }>();
  const queryClient = useQueryClient();
  const toast = useToast();

  const messages = useQuery(options.conversationMessagesOptions(ownerType, ownerId ?? ''));

  const [composerOpen, setComposerOpen] = useState<'email' | 'whatsapp' | null>(null);
  const [draft, setDraft] = useState('');
  const [aiOpen, setAiOpen] = useState(false);
  const [summary, setSummary] = useState<string | null>(null);
  const [suggestion, setSuggestion] = useState<string | null>(null);

  // Marca como lida ao abrir
  const markRead = useMutation({
    mutationFn: () => options.conversationMutations.markRead(ownerType, ownerId ?? ''),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: conversationKeysFor.all() });
      queryClient.invalidateQueries({ queryKey: ['plenya', 'notifications'] });
    },
  });

  useEffect(() => {
    if (ownerId) markRead.mutate();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [ownerId]);

  const sendEmail = useMutation({
    mutationFn: () =>
      options.conversationMutations.sendEmail(ownerType, ownerId ?? '', { content: draft }),
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: conversationKeysFor.messages(ownerType, ownerId ?? ''),
      });
      setComposerOpen(null);
      setDraft('');
      toast.show('Email enviado', 'success');
    },
    onError: (err) => toast.show(err instanceof Error ? err.message : 'Falha', 'error'),
  });

  const sendWA = useMutation({
    mutationFn: () =>
      options.conversationMutations.sendWhatsApp(ownerType, ownerId ?? '', { content: draft }),
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: conversationKeysFor.messages(ownerType, ownerId ?? ''),
      });
      setComposerOpen(null);
      setDraft('');
      toast.show('WhatsApp enviado', 'success');
    },
    onError: (err) => toast.show(err instanceof Error ? err.message : 'Falha', 'error'),
  });

  const aiSummary = useMutation({
    mutationFn: () => options.conversationMutations.aiSummary(ownerType, ownerId ?? ''),
    onSuccess: (res) => setSummary(res.summary),
    onError: (err) => toast.show(err instanceof Error ? err.message : 'IA falhou', 'error'),
  });

  const aiSuggest = useMutation({
    mutationFn: () => options.conversationMutations.aiSuggestReply(ownerType, ownerId ?? ''),
    onSuccess: (res) => setSuggestion(res.reply),
    onError: (err) => toast.show(err instanceof Error ? err.message : 'IA falhou', 'error'),
  });

  if (messages.isLoading) return <Spinner centered />;
  if (messages.isError) {
    return (
      <SafeAreaView className="flex-1 bg-background">
        <ErrorState onRetry={() => messages.refetch()} />
      </SafeAreaView>
    );
  }

  const items = messages.data ?? [];

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <KeyboardAvoidingView
        behavior={Platform.OS === 'ios' ? 'padding' : undefined}
        className="flex-1"
      >
        <FlatList
          data={items}
          inverted
          keyExtractor={(item) => item.id}
          contentContainerClassName="gap-2 p-4"
          renderItem={({ item }) => <MessageBubble msg={item} />}
        />

        <View className="flex-row gap-2 border-t border-border bg-card p-3">
          <Button
            variant="outline"
            size="sm"
            onPress={() => setAiOpen(true)}
            className="flex-1"
          >
            ✨ IA
          </Button>
          <Button
            variant="outline"
            size="sm"
            onPress={() => setComposerOpen('whatsapp')}
            className="flex-1"
          >
            💬 WA
          </Button>
          <Button
            size="sm"
            onPress={() => setComposerOpen('email')}
            className="flex-1"
          >
            ✉ Email
          </Button>
        </View>
      </KeyboardAvoidingView>

      <Sheet open={composerOpen !== null} onClose={() => setComposerOpen(null)}>
        <Text variant="title" className="mb-2">
          {composerOpen === 'email' ? 'Enviar email' : 'Enviar WhatsApp'}
        </Text>
        <View className="rounded-lg border border-border bg-background p-3">
          <TextInput
            value={draft}
            onChangeText={setDraft}
            multiline
            textAlignVertical="top"
            placeholder="Digite a mensagem..."
            placeholderTextColor="#9CA3AF"
            editable={!sendEmail.isPending && !sendWA.isPending}
            style={{ minHeight: 140, fontSize: 15, color: '#0f172a' }}
          />
        </View>
        <View className="mt-3 flex-row gap-2">
          <Button
            variant="outline"
            onPress={() => setComposerOpen(null)}
            className="flex-1"
            disabled={sendEmail.isPending || sendWA.isPending}
          >
            Cancelar
          </Button>
          <Button
            onPress={() => {
              if (!draft.trim()) return;
              if (composerOpen === 'email') sendEmail.mutate();
              else sendWA.mutate();
            }}
            loading={sendEmail.isPending || sendWA.isPending}
            className="flex-1"
          >
            Enviar
          </Button>
        </View>
      </Sheet>

      <Sheet
        open={aiOpen}
        onClose={() => {
          setAiOpen(false);
          setSummary(null);
          setSuggestion(null);
        }}
      >
        <Text variant="title" className="mb-2">
          Assistente IA
        </Text>
        <View className="gap-2">
          <Button
            variant="outline"
            onPress={() => aiSummary.mutate()}
            loading={aiSummary.isPending}
          >
            Resumir conversa
          </Button>
          <Button
            variant="outline"
            onPress={() => aiSuggest.mutate()}
            loading={aiSuggest.isPending}
          >
            Sugerir resposta
          </Button>
        </View>

        {summary && (
          <Card className="mt-3">
            <Text variant="caption" className="font-semibold">
              Resumo
            </Text>
            <Text variant="body" className="mt-1">
              {summary}
            </Text>
          </Card>
        )}

        {suggestion && (
          <Card className="mt-3">
            <Text variant="caption" className="font-semibold">
              Sugestão de resposta
            </Text>
            <Text variant="body" className="mt-1">
              {suggestion}
            </Text>
            <View className="mt-2 flex-row gap-2">
              <Button
                size="sm"
                onPress={() => {
                  setDraft(suggestion);
                  setAiOpen(false);
                  setComposerOpen('email');
                }}
                className="flex-1"
              >
                Usar no email
              </Button>
              <Button
                variant="outline"
                size="sm"
                onPress={() => {
                  setDraft(suggestion);
                  setAiOpen(false);
                  setComposerOpen('whatsapp');
                }}
                className="flex-1"
              >
                Usar no WA
              </Button>
            </View>
          </Card>
        )}
      </Sheet>
    </SafeAreaView>
  );
}

function MessageBubble({ msg }: { msg: ConversationMessage }) {
  const isOut = msg.direction === 'out';
  const channelIcon =
    msg.channel === 'email' ? '✉' : msg.channel === 'whatsapp' ? '💬' : '·';

  return (
    <View
      className={`max-w-[85%] rounded-2xl px-3 py-2 ${
        isOut ? 'self-end bg-primary' : 'self-start bg-muted'
      }`}
    >
      <Text
        variant="body"
        className={isOut ? 'text-primary-foreground' : 'text-foreground'}
      >
        {msg.content || `(${msg.type})`}
      </Text>
      {msg.attachments && msg.attachments.length > 0 && (
        <View className="mt-1 gap-0.5">
          {msg.attachments.map((a, i) => (
            <Text
              key={`${a.url}-${i}`}
              variant="caption"
              className={isOut ? 'text-primary-foreground' : 'text-muted-foreground'}
            >
              📎 {a.filename}
            </Text>
          ))}
        </View>
      )}
      <Text
        variant="caption"
        className={`mt-0.5 ${isOut ? 'text-primary-foreground' : 'text-muted-foreground'}`}
      >
        {channelIcon} {formatRelative(msg.createdAt)}
        {msg.actorName ? ` · ${msg.actorName}` : ''}
      </Text>
    </View>
  );
}
