import { forwardRef } from 'react';
import { TextInput, type TextInputProps } from 'react-native';
import { cn } from './cn';

export interface InputProps extends TextInputProps {
  className?: string;
  error?: boolean;
}

export const Input = forwardRef<TextInput, InputProps>(
  ({ className, error, editable = true, ...rest }, ref) => (
    <TextInput
      ref={ref}
      editable={editable}
      placeholderTextColor="#9CA3AF"
      accessibilityState={{ disabled: !editable }}
      className={cn(
        'h-11 rounded-lg border border-border bg-background px-3 text-base text-foreground',
        error && 'border-destructive',
        !editable && 'opacity-60',
        className,
      )}
      {...rest}
    />
  ),
);
Input.displayName = 'Input';
