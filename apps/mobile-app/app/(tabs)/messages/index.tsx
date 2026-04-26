import { ScrollView } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { EmptyState } from '@plenya/ui-mobile';

export default function MessagesScreen() {
  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <ScrollView contentContainerClassName="gap-4 p-4">
        <EmptyState
          title="Mensagens chegam na Sprint 4"
          description="Thread paciente↔clínica + push de nova mensagem — em construção."
        />
      </ScrollView>
    </SafeAreaView>
  );
}
