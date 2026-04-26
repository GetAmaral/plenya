import { useEffect, useRef, useState } from 'react';
import {
  FlatList,
  KeyboardAvoidingView,
  Platform,
  TextInput,
  View,
} from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { useFocusEffect } from 'expo-router';
import { useCallback } from 'react';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { options, queryKeys, type PatientMessage } from '@plenya/api-client';
import {
  Button,
  Card,
  EmptyState,
  Spinner,
  Text,
  useToast,
} from '@plenya/ui-mobile';
import { formatDateTime } from '@plenya/domain';

/**
 * Thread paciente↔clínica. Backend retorna ordem DESC (mais nova primeiro);
 * invertemos pra mostrar como chat (mais antiga em cima, scroll fica no fim).
 *
 * Mark-read dispara em focus + após enviar.
 */

function Bubble({ msg }: { msg: PatientMessage }) {
  const isMe = msg.from === 'patient';
  return (
    <Card
      className={`max-w-[85%] ${isMe ? 'self-end bg-primary' : 'self-start bg-card'}`}
    >
      {!isMe && msg.authorName && (
        <Text variant="caption" className="font-semibold">
          {msg.authorName}
        </Text>
      )}
      <Text
        variant="body"
        className={isMe ? 'text-primary-foreground' : 'text-foreground'}
      >
        {msg.content}
      </Text>
      <Text
        variant="caption"
        className={isMe ? 'text-primary-foreground/70 mt-1' : 'mt-1'}
      >
        {formatDateTime(msg.createdAt)}
      </Text>
    </Card>
  );
}

export default function MessagesScreen() {
  const qc = useQueryClient();
  const toast = useToast();
  const messages = useQuery(options.patientMeMessagesOptions());
  const [draft, setDraft] = useState('');
  const listRef = useRef<FlatList<PatientMessage>>(null);

  const send = useMutation({
    mutationFn: (content: string) =>
      options.patientMeMutations.sendMessage({ content }),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: queryKeys.patientMe.messages() });
      setDraft('');
    },
    onError: (err) =>
      toast.show(err instanceof Error ? err.message : 'Erro ao enviar', 'error'),
  });

  const markRead = useMutation({
    mutationFn: () => options.patientMeMutations.markMessagesRead(),
    onSuccess: () => {
      qc.invalidateQueries({
        queryKey: [...queryKeys.patientMe.messages(), 'unread'],
      });
    },
  });

  useFocusEffect(
    useCallback(() => {
      markRead.mutate();
      // intencionalmente sem dep de markRead: estável referencialmente o suficiente
      // eslint-disable-next-line react-hooks/exhaustive-deps
    }, []),
  );

  const inverted = (messages.data ?? []).slice().reverse();

  useEffect(() => {
    requestAnimationFrame(() =>
      listRef.current?.scrollToEnd({ animated: false }),
    );
  }, [inverted.length]);

  function handleSend() {
    const trimmed = draft.trim();
    if (!trimmed || send.isPending) return;
    send.mutate(trimmed);
  }

  if (messages.isLoading) return <Spinner centered />;

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <KeyboardAvoidingView
        behavior={Platform.OS === 'ios' ? 'padding' : undefined}
        className="flex-1"
      >
        {inverted.length === 0 ? (
          <EmptyState
            title="Sem mensagens ainda"
            description="Envie a primeira mensagem pra sua clínica."
          />
        ) : (
          <FlatList
            ref={listRef}
            data={inverted}
            keyExtractor={(m) => m.id}
            contentContainerClassName="gap-2 p-4"
            renderItem={({ item }) => <Bubble msg={item} />}
            onContentSizeChange={() =>
              listRef.current?.scrollToEnd({ animated: false })
            }
          />
        )}

        <View className="flex-row gap-2 border-t border-border bg-card p-3">
          <View className="flex-1 rounded-lg border border-border bg-background px-3 py-2">
            <TextInput
              value={draft}
              onChangeText={setDraft}
              multiline
              placeholder="Mensagem pra clínica…"
              placeholderTextColor="#9CA3AF"
              editable={!send.isPending}
              style={{ minHeight: 36, maxHeight: 120, fontSize: 15, color: '#0f172a' }}
            />
          </View>
          <Button onPress={handleSend} loading={send.isPending} size="sm">
            Enviar
          </Button>
        </View>
      </KeyboardAvoidingView>
    </SafeAreaView>
  );
}
