import { View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { router } from 'expo-router';
import { Button, Text } from '@plenya/ui-mobile';

export default function UnreachableScreen() {
  return (
    <SafeAreaView className="flex-1 bg-background">
      <View className="flex-1 items-center justify-center gap-4 px-6">
        <Text variant="heading" className="text-center">
          Sem conexão
        </Text>
        <Text variant="caption" className="text-center">
          Não conseguimos contatar o servidor. Verifique sua conexão e tente de novo.
        </Text>
        <Button onPress={() => router.replace('/')} fullWidth size="lg">
          Tentar novamente
        </Button>
      </View>
    </SafeAreaView>
  );
}
