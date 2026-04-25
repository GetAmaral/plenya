const preset = require('@plenya/ui-mobile/tailwind-preset').default;

module.exports = {
  content: [
    './app/**/*.{ts,tsx}',
    './features/**/*.{ts,tsx}',
    './components/**/*.{ts,tsx}',
    './lib/**/*.{ts,tsx}',
    '../../packages/ui-mobile/src/**/*.{ts,tsx}',
  ],
  presets: [preset],
  theme: { extend: {} },
  plugins: [],
};
