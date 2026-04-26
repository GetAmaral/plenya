import { Linking, Platform, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { useLocalSearchParams } from 'expo-router';
import { Button, Text } from '@plenya/ui-mobile';

const STORE_URL = Platform.select({
  ios: 'https://apps.apple.com/app/id000000000',
  android: 'https://play.google.com/store/apps/details?id=com.plenya.pro',
}) as string;

export default function MinVersionScreen() {
  const { required, current } = useLocalSearchParams<{ required: string; current: string }>();

  return (
    <SafeAreaView className="flex-1 bg-background">
      <View className="flex-1 items-center justify-center gap-4 px-6">
        <Text variant="heading" className="text-center">
          Atualize o Plenya Pro
        </Text>
        <Text variant="caption" className="text-center">
          Sua versão ({current}) é incompatível. Versão mínima: {required}.
        </Text>
        <Button onPress={() => Linking.openURL(STORE_URL)} fullWidth size="lg">
          Abrir loja
        </Button>
      </View>
    </SafeAreaView>
  );
}
