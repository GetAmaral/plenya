/**
 * Plenya brand colors — extracted from the official brandbook (PLENYA Visual Identity, 2026).
 *
 *   gold   #b38645  — Dourado/Ocre        (calor, confiança, sofisticação)
 *   petrol #063b4f  — Azul Petróleo       (profundidade, ciência, autoridade)
 *   ocean  #417e8e  — Azul Acinzentado    (clareza, equilíbrio — wordmark principal)
 *   sage   #92b8b4  — Verde Suave/Sálvia  (cuidado, calma, vitalidade)
 *   cream  #eae7da  — Bege/Off White      (papel editorial, respiração)
 */
export const colors = {
  gold: {
    DEFAULT: '#b38645',
    50: '#faf4ea',
    100: '#f3e6cf',
    200: '#e6cc9f',
    300: '#d4ad75',
    400: '#c39858',
    500: '#b38645',
    600: '#956d36',
    700: '#75552a',
    800: '#523c1e',
    900: '#332615',
  },
  petrol: {
    DEFAULT: '#063b4f',
    50: '#e6eef2',
    100: '#c2d4dc',
    200: '#85a9b6',
    300: '#467e90',
    400: '#1b566c',
    500: '#063b4f',
    600: '#052f40',
    700: '#042330',
    800: '#021921',
    light: '#417e8e',
  },
  ocean: {
    DEFAULT: '#417e8e',
    50: '#eaf2f4',
    100: '#cfdfe3',
    200: '#a0bfc7',
    300: '#6d9fab',
    400: '#56909d',
    500: '#417e8e',
    600: '#356674',
    700: '#284e58',
    800: '#1c373f',
  },
  sage: {
    DEFAULT: '#92b8b4',
    50: '#f0f6f5',
    100: '#dde9e7',
    200: '#bcd3d0',
    300: '#a4c3bf',
    400: '#92b8b4',
    500: '#7aa39e',
    600: '#608580',
    700: '#496661',
    800: '#324744',
  },
  cream: {
    DEFAULT: '#eae7da',
    50: '#fbfaf6',
    100: '#f3f1e7',
    200: '#eae7da',
    300: '#d8d1ba',
    400: '#bcb094',
  },
  ink: '#0a0a0a',
  paper: '#fbfaf6',
} as const;

export type ColorToken = typeof colors;
