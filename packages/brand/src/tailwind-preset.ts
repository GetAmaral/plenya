import type { Config } from 'tailwindcss';
import { colors, fontFamilies, fontSizes, letterSpacing, radii, containers } from './tokens';

const preset: Partial<Config> = {
  theme: {
    extend: {
      colors: {
        gold: colors.gold,
        petrol: colors.petrol,
        sage: colors.sage,
        cream: colors.cream,
        ink: colors.ink,
        paper: colors.paper,
      },
      fontFamily: {
        heading: fontFamilies.heading,
        mono: fontFamilies.mono,
        sans: fontFamilies.body,
      },
      fontSize: fontSizes,
      letterSpacing,
      borderRadius: radii,
      maxWidth: {
        prose: containers.prose,
        narrow: containers.narrow,
        container: containers.default,
        wide: containers.wide,
      },
    },
  },
};

export default preset;
