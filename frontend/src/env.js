// In Docker/nginx proxy mode, use same origin (empty string)
// For local dev without proxy, set to backend URL directly
export const HOST_URL = typeof window !== 'undefined' && window.__HOST_URL__
  ? window.__HOST_URL__
  : '';
export const APP_ENV = 'development';
