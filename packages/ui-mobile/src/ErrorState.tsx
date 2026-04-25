import { View } from 'react-native';
import { Button } from './Button';
import { cn } from './cn';
import { Text } from './Text';

export interface ErrorStateProps {
  title?: string;
  message?: string;
  onRetry?: () => void;
  className?: string;
}

export function ErrorState({
  title = 'Algo deu errado',
  message = 'Não foi possível carregar. Tente de novo.',
  onRetry,
  className,
}: ErrorStateProps) {
  return (
    <View className={cn('flex-1 items-center justify-center gap-3 px-6 py-12', className)}>
      <Text variant="title" className="text-center">
        {title}
      </Text>
      <Text variant="caption" className="text-center">
        {message}
      </Text>
      {onRetry && (
        <Button onPress={onRetry} variant="outline">
          Tentar novamente
        </Button>
      )}
    </View>
  );
}
