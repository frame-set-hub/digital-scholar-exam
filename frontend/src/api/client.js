/**
 * Base URL สำหรับ API
 * - ว่าง: ใช้ path สัมพัทธ์ `/api` (คู่กับ proxy ใน vite.config — แนะนำตอน dev)
 * - ตั้ง VITE_API_BASE_URL: เช่น http://localhost:8080
 */
export function apiBase() {
  const b = import.meta.env.VITE_API_BASE_URL
  return typeof b === 'string' ? b.replace(/\/$/, '') : ''
}

export function apiUrl(path) {
  const p = path.startsWith('/') ? path : `/${path}`
  const base = apiBase()
  return base ? `${base}${p}` : p
}

export async function fetchJSON(url, options = {}) {
  const headers = { ...options.headers }
  if (options.body != null && headers['Content-Type'] == null) {
    headers['Content-Type'] = 'application/json'
  }
  const res = await fetch(url, {
    ...options,
    headers,
  })
  const text = await res.text()
  let data
  try {
    data = text ? JSON.parse(text) : null
  } catch {
    throw new Error(text || res.statusText)
  }
  if (!res.ok) {
    const msg = data?.error || res.statusText || 'Request failed'
    const err = new Error(msg)
    err.status = res.status
    err.data = data
    throw err
  }
  return data
}
