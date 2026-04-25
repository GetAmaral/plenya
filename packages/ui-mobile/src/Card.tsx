import { View, type ViewProps } from 'react-native';
import { cn } from './cn';

export interface CardProps extends ViewProps {
  className?: string;
  padded?: boolean;
}

export function Card({ className, padded = true, ...rest }: CardProps) {
  return (
    <View
      className={cn(
        'rounded-xl border border-border bg-card',
        padded && 'p-4',
        className,
      )}
      {...rest}
    />
  );
}

export function CardHeader({ className, ...rest }: CardProps) {
  return <View className={cn('mb-2 gap-1', className)} {...rest} />;
}

export function CardContent({ className, ...rest }: CardProps) {
  return <View className={cn('gap-2', className)} {...rest} />;
}

export function CardFooter({ className, ...rest }: CardProps) {
  return <View className={cn('mt-3 flex-row items-center gap-2', className)} {...rest} />;
}
