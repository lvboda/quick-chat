export async function wait(ms: number): Promise<void> {
  return new Promise((res) => setTimeout(res, ms));
}

export async function waitCall<T>(fn: () => T, ms: number): Promise<T> {
  await new Promise((res) => setTimeout(res, ms));
  return fn();
}
