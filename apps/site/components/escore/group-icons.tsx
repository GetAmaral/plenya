/**
 * Mapeia nome de grupo do Escore Light → ícone Lucide.
 * Usado no cabeçalho de cada step do formulário.
 */
import {
  Apple,
  Dumbbell,
  Moon,
  Brain,
  Waves,
  Ruler,
  HeartPulse,
  Users,
  Home,
  TestTube,
  HelpCircle,
  type LucideIcon,
} from 'lucide-react';

const ICON_MAP: Record<string, LucideIcon> = {
  'Alimentação': Apple,
  'Movimento e atividade física': Dumbbell,
  'Sono': Moon,
  'Cognição': Brain,
  'Stress': Waves,
  'Composição corporal': Ruler,
  'Histórico de doenças': HeartPulse,
  'Histórico Familiar de Doenças': Users,
  'Social': Home,
  'Exames': TestTube,
};

export function getGroupIcon(groupName: string): LucideIcon {
  return ICON_MAP[groupName] ?? HelpCircle;
}
