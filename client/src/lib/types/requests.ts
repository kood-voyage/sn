
export type ReturnType<T> = { ok: true, data: T } | { ok: false, error: Error | unknown, message: string }