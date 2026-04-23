'use client';

import { forwardRef, useState } from 'react';

/**
 * Input de telefone BR (10 ou 11 dígitos), formatação visual `(XX) XXXXX-XXXX`.
 * O `onChange` recebe sempre o valor em E.164 (`+55XXXXXXXXXXX`) ou string vazia se inválido.
 *
 * Sem dependência externa — libphonenumber-js seria 200KB. Aceita apenas BR no MVP.
 */

type PhoneInputProps = {
  value?: string; // E.164 (+55...)
  onChange?: (e164: string) => void;
  placeholder?: string;
  required?: boolean;
  disabled?: boolean;
  className?: string;
  id?: string;
  name?: string;
  /**
   * Quando true (default), aceita apenas celular BR: 11 dígitos com 9° dígito = '9'.
   * Use false em formulários onde landline também é aceito (ex: contato genérico).
   */
  requireMobile?: boolean;
  'aria-describedby'?: string;
};

function digitsOnly(s: string): string {
  return s.replace(/\D/g, '');
}

function formatBR(digits: string): string {
  const d = digits.slice(0, 11);
  if (d.length === 0) return '';
  if (d.length <= 2) return `(${d}`;
  if (d.length <= 6) return `(${d.slice(0, 2)}) ${d.slice(2)}`;
  if (d.length <= 10) return `(${d.slice(0, 2)}) ${d.slice(2, 6)}-${d.slice(6)}`;
  return `(${d.slice(0, 2)}) ${d.slice(2, 7)}-${d.slice(7, 11)}`;
}

function toE164(digits: string, requireMobile: boolean): string {
  const d = digits.replace(/\D/g, '');
  if (requireMobile) {
    // Celular BR: 11 dígitos, 9° dígito (índice 2 do número local após DDD) = '9'
    if (d.length !== 11) return '';
    if (d[2] !== '9') return '';
    return `+55${d}`;
  }
  if (d.length !== 10 && d.length !== 11) return '';
  return `+55${d}`;
}

function fromE164(e164: string): string {
  if (!e164) return '';
  return e164.replace(/^\+55/, '');
}

export const PhoneInput = forwardRef<HTMLInputElement, PhoneInputProps>(function PhoneInput(
  {
    value,
    onChange,
    placeholder = '(11) 99999-9999',
    required,
    disabled,
    className,
    id,
    name,
    requireMobile = true,
    ...rest
  },
  ref
) {
  const [display, setDisplay] = useState(() => formatBR(fromE164(value ?? '')));

  function handleChange(e: React.ChangeEvent<HTMLInputElement>) {
    const raw = digitsOnly(e.target.value);
    const formatted = formatBR(raw);
    setDisplay(formatted);
    onChange?.(toE164(raw, requireMobile));
  }

  return (
    <input
      ref={ref}
      id={id}
      name={name}
      type="tel"
      inputMode="tel"
      autoComplete="tel"
      placeholder={placeholder}
      value={display}
      onChange={handleChange}
      required={required}
      disabled={disabled}
      className={className}
      aria-describedby={rest['aria-describedby']}
    />
  );
});
