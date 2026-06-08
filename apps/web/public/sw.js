// Service worker do EMR Plenya — só Web Push (avisos sem app nativo).
// Plano: docs/emr/plano-webpush-notificacoes.md.

self.addEventListener('install', () => {
  self.skipWaiting()
})

self.addEventListener('activate', (event) => {
  event.waitUntil(self.clients.claim())
})

self.addEventListener('push', (event) => {
  let payload = {}
  try {
    payload = event.data ? event.data.json() : {}
  } catch (_e) {
    payload = { title: 'Plenya', body: event.data ? event.data.text() : '' }
  }

  const title = payload.title || 'Plenya'
  const url = payload.url || '/'
  const options = {
    body: payload.body || '',
    icon: '/icon-192.png',
    badge: '/icon-192.png',
    tag: payload.data && payload.data.tag ? payload.data.tag : undefined,
    renotify: true,
    data: { url, ...(payload.data || {}) },
  }

  event.waitUntil(self.registration.showNotification(title, options))
})

self.addEventListener('notificationclick', (event) => {
  event.notification.close()
  const targetUrl = (event.notification.data && event.notification.data.url) || '/'

  event.waitUntil(
    self.clients.matchAll({ type: 'window', includeUncontrolled: true }).then((clients) => {
      // Se já há uma aba do EMR aberta, foca e navega.
      for (const client of clients) {
        if ('focus' in client) {
          client.focus()
          if ('navigate' in client && targetUrl) {
            try {
              client.navigate(targetUrl)
            } catch (_e) {
              /* navigate pode falhar cross-origin; ignora */
            }
          }
          return
        }
      }
      // Senão, abre uma nova.
      if (self.clients.openWindow) {
        return self.clients.openWindow(targetUrl)
      }
    })
  )
})
