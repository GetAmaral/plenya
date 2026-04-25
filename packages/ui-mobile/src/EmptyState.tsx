import { View } from 'react-native';
import { cn } from './cn';
import { Text } from './Text';

export interface EmptyStateProps {
  title: string;
  description?: string;
  action?: React.ReactNode;
  className?: string;
}

export function EmptyState({ title, description, action, className }: EmptyStateProps) {
  return (
    <View className={cn('flex-1 items-center justify-center gap-3 px-6 py-12', className)}>
      <Text variant="title" className="text-center">
        {title}
      </Text>
      {description && (
        <Text variant="caption" className="text-center">
          {description}
        </Text>
      )}
      {action}
    </View>
  );
}
