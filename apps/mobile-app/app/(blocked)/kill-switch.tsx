import { View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { useLocalSearchParams } from 'expo-router';
import { Text } from '@plenya/ui-mobile';

export default function KillSwitchScreen() {
  const { message } = useLocalSearchParams<{ message: string }>();

  return (
    <SafeAreaView className="flex-1 bg-background">
      <View className="flex-1 items-center justify-center gap-4 px-6">
        <Text variant="heading" className="text-center">
          App temporariamente indisponível
        </Text>
        <Text variant="caption" className="text-center">
          {message || 'Tente novamente em instantes.'}
        </Text>
      </View>
    </SafeAreaView>
  );
}
