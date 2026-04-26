import * as Sentry from '@sentry/react-native';
import { env } from '../env';

const PHI_KEYS = new Set([
  'cpf',
  'rg',
  'name',
  'fullName',
  'email',
  'phone',
  'patient',
  'patientName',
  'birthDate',
  'address',
  'reason',
  'diagnosis',
  'doctorNotes',
  'patientNotes',
  'freeText',
  'token',
  'accessToken',
  'refreshToken',
  'authorization',
]);

function scrub(value: unknown, depth = 0): unknown {
  if (depth > 6) return '[truncated]';
  if (value === null || value === undefined) return value;
  if (Array.isArray(value)) return value.map((v) => scrub(v, depth + 1));
  if (typeof value !== 'object') return value;

  const out: Record<string, unknown> = {};
  for (const [k, v] of Object.entries(value as Record<string, unknown>)) {
    if (PHI_KEYS.has(k)) {
      out[k] = '[stripped]';
    } else {
      out[k] = scrub(v, depth + 1);
    }
  }
  return out;
}

let initialized = false;

/**
 * Inicializa Sentry com tracing nativo + strip de PHI no beforeSend.
 * No-op se SENTRY_DSN não estiver setado (dev local sem Sentry).
 */
export function initSentry(): void {
  if (initialized) return;
  if (!env.sentryDsn) return;

  Sentry.init({
    dsn: env.sentryDsn,
    environment: env.appVariant,
    enableAutoSessionTracking: true,
    sessionTrackingIntervalMillis: 30_000,
    tracesSampleRate: env.appVariant === 'production' ? 0.1 : 1.0,
    beforeSend(event) {
      if (event.user) {
        delete event.user.email;
        delete event.user.username;
        delete event.user.ip_address;
      }
      if (event.contexts) {
        event.contexts = scrub(event.contexts) as typeof event.contexts;
      }
      if (event.extra) {
        event.extra = scrub(event.extra) as typeof event.extra;
      }
      if (event.breadcrumbs) {
        event.breadcrumbs = event.breadcrumbs.map((b) => ({
          ...b,
          data: b.data ? (scrub(b.data) as typeof b.data) : b.data,
        }));
      }
      if (event.request?.url) {
        event.request.url = event.request.url.replace(/\/[0-9a-f-]{36}/g, '/<uuid>');
      }
      return event;
    },
  });

  initialized = true;
}

export { Sentry };
