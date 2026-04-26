import { ScrollView } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { EmptyState } from '@plenya/ui-mobile';

export default function ExamsScreen() {
  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <ScrollView contentContainerClassName="gap-4 p-4">
        <EmptyState
          title="Exames e escores chegam na Sprint 2"
          description="Lab batches, escores e download PDF — em construção."
        />
      </ScrollView>
    </SafeAreaView>
  );
}
