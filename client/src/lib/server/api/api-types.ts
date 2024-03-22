
export type Resp = {
  ok: true;
  data: string;
} | {
  ok: false;
  error: unknown;
  message: string;
}