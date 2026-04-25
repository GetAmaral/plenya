import { forwardRef } from 'react';
import { Text as RNText, type TextProps as RNTextProps } from 'react-native';
import { cn } from './cn';

export type TextVariant =
  | 'display'
  | 'heading'
  | 'title'
  | 'body'
  | 'caption'
  | 'mono';

const variantClass: Record<TextVariant, string> = {
  display: 'text-4xl font-bold text-foreground',
  heading: 'text-2xl font-semibold text-foreground',
  title: 'text-lg font-semibold text-foreground',
  body: 'text-base text-foreground',
  caption: 'text-sm text-muted-foreground',
  mono: 'text-sm font-mono text-foreground',
};

export interface TextProps extends RNTextProps {
  variant?: TextVariant;
  className?: string;
}

export const Text = forwardRef<RNText, TextProps>(
  ({ variant = 'body', className, ...props }, ref) => (
    <RNText ref={ref} className={cn(variantClass[variant], className)} {...props} />
  ),
);
Text.displayName = 'Text';
