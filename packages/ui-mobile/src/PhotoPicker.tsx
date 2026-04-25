import { useState } from 'react';
import { Image, Pressable, View } from 'react-native';
import { Button } from './Button';
import { cn } from './cn';
import { Text } from './Text';

export interface SelectedPhoto {
  uri: string;
  width?: number;
  height?: number;
  mimeType?: string;
  fileName?: string;
}

export interface PhotoPickerProps {
  photos: SelectedPhoto[];
  onAdd: () => void | Promise<void>;
  onRemove: (index: number) => void;
  max?: number;
  label?: string;
  className?: string;
  disabled?: boolean;
}

/**
 * Visual primitive — não toca expo-image-picker direto pra manter o package
 * livre de deps nativas. O caller passa onAdd que abre o picker.
 */
export function PhotoPicker({
  photos,
  onAdd,
  onRemove,
  max = 6,
  label = 'Fotos',
  className,
  disabled,
}: PhotoPickerProps) {
  const [busy, setBusy] = useState(false);
  const limitReached = photos.length >= max;

  async function handleAdd() {
    setBusy(true);
    try {
      await onAdd();
    } finally {
      setBusy(false);
    }
  }

  return (
    <View className={cn('gap-2', className)}>
      <View className="flex-row items-center justify-between">
        <Text variant="caption">
          {label} {max > 1 ? `(${photos.length}/${max})` : ''}
        </Text>
      </View>

      <View className="flex-row flex-wrap gap-2">
        {photos.map((p, i) => (
          <View key={`${p.uri}-${i}`} className="relative">
            <Image source={{ uri: p.uri }} style={{ width: 96, height: 96, borderRadius: 8 }} />
            {!disabled && (
              <Pressable
                onPress={() => onRemove(i)}
                className="absolute -right-2 -top-2 h-6 w-6 items-center justify-center rounded-full bg-destructive"
              >
                <Text className="text-xs font-bold text-white">×</Text>
              </Pressable>
            )}
          </View>
        ))}
      </View>

      <Button
        variant="outline"
        size="sm"
        loading={busy}
        disabled={disabled || limitReached}
        onPress={handleAdd}
      >
        {limitReached ? 'Limite atingido' : 'Adicionar foto'}
      </Button>
    </View>
  );
}
