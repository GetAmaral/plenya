import { forwardRef } from 'react';
import {
  ActivityIndicator,
  Pressable,
  Text,
  type PressableProps,
} from 'react-native';
import { cn } from './cn';

export type ButtonVariant = 'primary' | 'secondary' | 'outline' | 'ghost' | 'destructive';
export type ButtonSize = 'sm' | 'md' | 'lg';

const variantClass: Record<ButtonVariant, string> = {
  primary: 'bg-primary active:opacity-90',
  secondary: 'bg-secondary active:opacity-90',
  outline: 'border border-border bg-transparent active:bg-muted',
  ghost: 'bg-transparent active:bg-muted',
  destructive: 'bg-destructive active:opacity-90',
};

const textClass: Record<ButtonVariant, string> = {
  primary: 'text-primary-foreground font-semibold',
  secondary: 'text-secondary-foreground font-semibold',
  outline: 'text-foreground font-semibold',
  ghost: 'text-foreground font-semibold',
  destructive: 'text-destructive-foreground font-semibold',
};

const sizeClass: Record<ButtonSize, string> = {
  sm: 'h-9 px-3 rounded-md',
  md: 'h-11 px-4 rounded-lg',
  lg: 'h-14 px-6 rounded-xl',
};

const textSize: Record<ButtonSize, string> = {
  sm: 'text-sm',
  md: 'text-base',
  lg: 'text-lg',
};

export interface ButtonProps extends Omit<PressableProps, 'children'> {
  children: React.ReactNode;
  variant?: ButtonVariant;
  size?: ButtonSize;
  loading?: boolean;
  fullWidth?: boolean;
  className?: string;
}

export const Button = forwardRef<React.ComponentRef<typeof Pressable>, ButtonProps>(
  (
    { variant = 'primary', size = 'md', loading, fullWidth, disabled, className, children, ...rest },
    ref,
  ) => {
    const isDisabled = disabled || loading;
    return (
      <Pressable
        ref={ref}
        accessibilityRole="button"
        disabled={isDisabled}
        className={cn(
          'flex-row items-center justify-center',
          variantClass[variant],
          sizeClass[size],
          fullWidth && 'w-full',
          isDisabled && 'opacity-50',
          className,
        )}
        {...rest}
      >
        {loading ? (
          <ActivityIndicator size="small" />
        ) : typeof children === 'string' ? (
          <Text className={cn(textClass[variant], textSize[size])}>{children}</Text>
        ) : (
          children
        )}
      </Pressable>
    );
  },
);
Button.displayName = 'Button';
