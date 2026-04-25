import { View } from 'react-native';
import { Link } from 'expo-router';
import { Card, Text } from '@plenya/ui-mobile';

export interface SubsectionLinkProps {
  href: string;
  title: string;
  description?: string;
  count?: number;
}

/**
 * Card clicável usado no overview do paciente para navegar para anamnese,
 * exames, prescrições, escores. Mostra contagem opcional.
 */
export function SubsectionLink({ href, title, description, count }: SubsectionLinkProps) {
  return (
    <Link href={href as never} asChild>
      <Card>
        <View className="flex-row items-center justify-between">
          <View className="flex-1">
            <Text variant="title">{title}</Text>
            {description && (
              <Text variant="caption" className="mt-0.5">
                {description}
              </Text>
            )}
          </View>
          {typeof count === 'number' && (
            <View className="ml-3 h-7 min-w-7 items-center justify-center rounded-full bg-primary px-2">
              <Text className="text-xs font-semibold text-primary-foreground">{count}</Text>
            </View>
          )}
        </View>
      </Card>
    </Link>
  );
}
