import { Link } from 'expo-router';
import { View } from 'react-native';
import { Text } from '@plenya/ui-mobile';

export default function NotFoundScreen() {
  return (
    <View className="flex-1 items-center justify-center bg-background p-6">
      <Text variant="heading">Página não encontrada</Text>
      <Link href="/" className="mt-4 text-primary">
        Voltar ao início
      </Link>
    </View>
  );
}
