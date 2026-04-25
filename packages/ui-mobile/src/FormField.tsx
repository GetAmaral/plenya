import { Text, View } from 'react-native';
import { cn } from './cn';

export interface FormFieldProps {
  label?: string;
  helper?: string;
  error?: string;
  required?: boolean;
  className?: string;
  children: React.ReactNode;
}

export function FormField({
  label,
  helper,
  error,
  required,
  className,
  children,
}: FormFieldProps) {
  return (
    <View className={cn('gap-1.5', className)}>
      {label && (
        <Text className="text-sm font-medium text-foreground">
          {label}
          {required && <Text className="text-destructive"> *</Text>}
        </Text>
      )}
      {children}
      {error ? (
        <Text className="text-xs text-destructive">{error}</Text>
      ) : helper ? (
        <Text className="text-xs text-muted-foreground">{helper}</Text>
      ) : null}
    </View>
  );
}
