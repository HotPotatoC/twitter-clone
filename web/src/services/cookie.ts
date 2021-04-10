export function get(name: string): string {
  let cookie: Record<string, string> = {}
  document.cookie.split(';').forEach(function (el) {
    let [k, v] = el.split('=')
    cookie[k.trim()] = v
  })
  return cookie[name]
}
