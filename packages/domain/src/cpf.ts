export function stripCpf(value: string): string {
  return value.replace(/\D/g, '');
}

export function formatCpf(value: string): string {
  const digits = stripCpf(value);
  if (digits.length !== 11) return value;
  return `${digits.slice(0, 3)}.${digits.slice(3, 6)}.${digits.slice(6, 9)}-${digits.slice(9)}`;
}

export function maskCpf(value: string): string {
  const digits = stripCpf(value);
  if (digits.length !== 11) return value;
  return `***.${digits.slice(3, 6)}.${digits.slice(6, 9)}-**`;
}

export function isValidCpf(value: string): boolean {
  const cpf = stripCpf(value);
  if (cpf.length !== 11) return false;
  if (/^(\d)\1{10}$/.test(cpf)) return false;

  const digits = cpf.split('').map(Number);

  const check = (slice: number[], factor: number) => {
    const sum = slice.reduce((acc, d, i) => acc + d * (factor - i), 0);
    const rest = (sum * 10) % 11;
    return rest === 10 ? 0 : rest;
  };

  return (
    check(digits.slice(0, 9), 10) === digits[9] &&
    check(digits.slice(0, 10), 11) === digits[10]
  );
}
