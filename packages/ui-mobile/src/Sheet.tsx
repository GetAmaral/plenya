import { Modal, Pressable, View, type ModalProps } from 'react-native';
import { cn } from './cn';

export interface SheetProps extends Omit<ModalProps, 'children'> {
  open: boolean;
  onClose: () => void;
  children: React.ReactNode;
  className?: string;
}

/**
 * Thin bottom-sheet built on top of Modal. Replace with @gorhom/bottom-sheet
 * if gestures become a requirement.
 */
export function Sheet({ open, onClose, children, className, ...rest }: SheetProps) {
  return (
    <Modal
      visible={open}
      animationType="slide"
      transparent
      onRequestClose={onClose}
      {...rest}
    >
      <Pressable className="flex-1 bg-black/40" onPress={onClose} />
      <View
        className={cn(
          'absolute bottom-0 left-0 right-0 rounded-t-2xl bg-card p-4 pb-8',
          className,
        )}
      >
        {children}
      </View>
    </Modal>
  );
}
