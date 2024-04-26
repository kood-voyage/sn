
export type ReturnType<T> = { ok: true, data: T } | { ok: false, error: Error | unknown, message: string }

export type ReturnToClientType<T> = { ok: true, data: T } | { ok: false, message: string }