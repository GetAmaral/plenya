import { dirname } from "path";
import { fileURLToPath } from "url";
import { FlatCompat } from "@eslint/eslintrc";

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

const compat = new FlatCompat({
  baseDirectory: __dirname,
});

const eslintConfig = [
  ...compat.extends("next/core-web-vitals", "next/typescript"),
  {
    // Anti-XSS: HTML cru só via <SafeHtml> (que sanitiza). Bloqueia sink novo sem sanitização.
    // Ver docs/emr/estudo-xss-hardening.md.
    rules: {
      "react/no-danger": "error",
    },
  },
];

export default eslintConfig;
