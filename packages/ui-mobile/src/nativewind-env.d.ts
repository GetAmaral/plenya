/// <reference types="nativewind/types" />

import 'react-native';

declare module 'react-native' {
  interface PressableProps {
    className?: string;
  }
  interface ActivityIndicatorProps {
    className?: string;
  }
  interface ModalProps {
    className?: string;
  }
}
