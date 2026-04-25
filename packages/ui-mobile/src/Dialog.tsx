import { Modal, Pressable, View, type ModalProps } from 'react-native';
import { cn } from './cn';

export interface DialogProps extends Omit<ModalProps, 'children'> {
  open: boolean;
  onClose: () => void;
  children: React.ReactNode;
  className?: string;
}

export function Dialog({ open, onClose, children, className, ...rest }: DialogProps) {
  return (
    <Modal
      visible={open}
      animationType="fade"
      transparent
      onRequestClose={onClose}
      {...rest}
    >
      <Pressable className="flex-1 items-center justify-center bg-black/50 p-6" onPress={onClose}>
        <Pressable onPress={(e) => e.stopPropagation()}>
          <View className={cn('w-full max-w-md rounded-2xl bg-card p-5', className)}>
            {children}
          </View>
        </Pressable>
      </Pressable>
    </Modal>
  );
}
