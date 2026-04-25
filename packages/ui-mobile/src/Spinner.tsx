import { ActivityIndicator, View, type ActivityIndicatorProps } from 'react-native';
import { cn } from './cn';

export interface SpinnerProps extends ActivityIndicatorProps {
  centered?: boolean;
  className?: string;
}

export function Spinner({ centered, className, ...rest }: SpinnerProps) {
  if (centered) {
    return (
      <View className={cn('flex-1 items-center justify-center', className)}>
        <ActivityIndicator {...rest} />
      </View>
    );
  }
  return <ActivityIndicator className={className} {...rest} />;
}
