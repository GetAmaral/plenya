import axios from 'axios'

const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:3001/api/v1'

console.log('🔍 API Client initialized with baseURL:', API_URL)

export const apiClient = axios.create({
  baseURL: API_URL,
  headers: {
    'Content-Type': 'application/json',
  },
})

// Request interceptor to add auth token
apiClient.interceptors.request.use(
  (config) => {
    // Get token from localStorage
    if (typeof window !== 'undefined') {
      const authStore = localStorage.getItem('plenya-auth')
      if (authStore) {
        try {
          const parsed = JSON.parse(authStore)
          const token = parsed?.state?.accessToken
          if (token) {
            config.headers.Authorization = `Bearer ${token}`
          }
        } catch (e) {
          console.error('Error parsing auth store:', e)
        }
      }
    }
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

// Response interceptor to handle errors
apiClient.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      // Token expired or invalid
      if (typeof window !== 'undefined') {
        localStorage.removeItem('plenya-auth')
        window.location.href = '/login'
      }
    }
    return Promise.reject(error)
  }
)
