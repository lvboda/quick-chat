function genKey(key: string, ns = "[ROOT]"): string {
  return `${ns}:${key}`;
}
export function getFromLocal<T>(key: string, ns?: string): T | null {
  const str = window.localStorage.getItem(genKey(key, ns));
  return str ? JSON.parse(str) : null;
}

export function setToLocal(key: string, value: any, ns?: string): void {
  window.localStorage.setItem(genKey(key, ns), JSON.stringify(value));
}

export function removeFromLocal(key: string, ns?: string): void {
  window.localStorage.removeItem(genKey(key, ns));
}
