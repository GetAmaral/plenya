import { useRef, useState } from 'react';
import {
  FlatList,
  KeyboardAvoidingView,
  Platform,
  TextInput,
  View,
} from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { useMutation } from '@tanstack/react-query';
import { options, type TrainingAIChatMessage } from '@plenya/api-client';
import { Button, Card, EmptyState, Text, useToast } from '@plenya/ui-mobile';

export default function TrainingAiAgentScreen() {
  const [messages, setMessages] = useState<TrainingAIChatMessage[]>([]);
  const [draft, setDraft] = useState('');
  const listRef = useRef<FlatList<TrainingAIChatMessage>>(null);
  const toast = useToast();

  const chat = useMutation({
    mutationFn: (next: TrainingAIChatMessage[]) =>
      options.trainingAiMutations.chat({ messages: next }),
    onSuccess: (res) => {
      setMessages((prev) => [...prev, { role: 'assistant', content: res.message }]);
      requestAnimationFrame(() =>
        listRef.current?.scrollToEnd({ animated: true }),
      );
    },
    onError: (err) =>
      toast.show(err instanceof Error ? err.message : 'Falha na IA', 'error'),
  });

  function handleSend() {
    const trimmed = draft.trim();
    if (!trimmed || chat.isPending) return;
    const next: TrainingAIChatMessage[] = [
      ...messages,
      { role: 'user', content: trimmed },
    ];
    setMessages(next);
    setDraft('');
    chat.mutate(next);
  }

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <KeyboardAvoidingView
        behavior={Platform.OS === 'ios' ? 'padding' : undefined}
        className="flex-1"
      >
        {messages.length === 0 ? (
          <EmptyState
            title="Agente IA de fisiologia"
            description="Pergunte sobre programação de treino, fisiologia do exercício, periodização, ACSM, NSCA. Histórico fica no app, não persiste no servidor."
          />
        ) : (
          <FlatList
            ref={listRef}
            data={messages}
            keyExtractor={(_, i) => String(i)}
            contentContainerClassName="gap-2 p-4"
            renderItem={({ item }) => <Bubble msg={item} />}
            onContentSizeChange={() =>
              listRef.current?.scrollToEnd({ animated: false })
            }
          />
        )}

        {chat.isPending && (
          <View className="px-4 pb-2">
            <Text variant="caption" className="italic">
              Pensando…
            </Text>
          </View>
        )}

        <View className="flex-row gap-2 border-t border-border bg-card p-3">
          <View className="flex-1 rounded-lg border border-border bg-background px-3 py-2">
            <TextInput
              value={draft}
              onChangeText={setDraft}
              multiline
              placeholder="Pergunte ao agente..."
              placeholderTextColor="#9CA3AF"
              editable={!chat.isPending}
              style={{ minHeight: 36, maxHeight: 120, fontSize: 15, color: '#0f172a' }}
            />
          </View>
          <Button onPress={handleSend} loading={chat.isPending} size="sm">
            Enviar
          </Button>
        </View>
      </KeyboardAvoidingView>
    </SafeAreaView>
  );
}

function Bubble({ msg }: { msg: TrainingAIChatMessage }) {
  const isUser = msg.role === 'user';
  return (
    <Card
      className={`max-w-[90%] ${isUser ? 'self-end bg-primary' : 'self-start bg-muted'}`}
    >
      <Text
        variant="body"
        className={isUser ? 'text-primary-foreground' : 'text-foreground'}
      >
        {msg.content}
      </Text>
    </Card>
  );
}
