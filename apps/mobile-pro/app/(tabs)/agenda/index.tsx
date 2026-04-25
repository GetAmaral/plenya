import { View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { Text } from '@plenya/ui-mobile';

export default function AgendaScreen() {
  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <View className="flex-1 items-center justify-center p-6">
        <Text variant="heading">Agenda</Text>
        <Text variant="caption" className="mt-2 text-center">
          Integração com /api/v1/appointments — Sprint 5.
        </Text>
      </View>
    </SafeAreaView>
  );
}
