import type { Config } from 'tailwindcss';
import plenyaPreset from '@plenya/brand/tailwind-preset';

const config: Config = {
  presets: [plenyaPreset as Config],
  content: [
    './app/**/*.{ts,tsx,mdx}',
    './components/**/*.{ts,tsx}',
    './content/**/*.{md,mdx}',
    '../../packages/brand/src/**/*.{ts,tsx}',
  ],
  theme: {
    extend: {},
  },
  plugins: [],
};

export default config;
