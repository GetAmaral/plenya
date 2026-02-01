/**
 * Suppress unnecessary development logs in browser console
 * This improves developer experience by hiding verbose Next.js/React logs
 */

if (process.env.NODE_ENV === 'development' && typeof window !== 'undefined') {
  // Suppress React DevTools suggestion (annoying in every session)
  const originalConsoleLog = console.log
  console.log = (...args) => {
    if (
      typeof args[0] === 'string' &&
      args[0].includes('Download the React DevTools')
    ) {
      return
    }
    originalConsoleLog.apply(console, args)
  }

  // Suppress HMR/Fast Refresh logs (too verbose)
  const originalConsoleInfo = console.info
  console.info = (...args) => {
    if (
      typeof args[0] === 'string' &&
      (args[0].includes('[HMR]') ||
        args[0].includes('[Fast Refresh]'))
    ) {
      return
    }
    originalConsoleInfo.apply(console, args)
  }

  // Keep errors and warnings visible (important for debugging)
  // console.error and console.warn are NOT suppressed
}

export {}
