import { writable } from "svelte/store";

export type ServerMessage = {
  type: string,
  address: string,
  id: string,
  source_id: string,
  data: string | object,
}

export type WebSocketStore = { websocket: WebSocket | undefined, access_token: string | undefined }
export const webSocketStore = writable<WebSocketStore>({ websocket: undefined, access_token: undefined })

export type MessageStore = Array<ServerMessage>
export const messageStore = writable<MessageStore>([])

export type UserStatusStore = { [key: string]: boolean }
export const userStatusStore = writable<UserStatusStore>({})

export type NotificationStore = Array<ServerMessage>
export const notificationStore = writable<NotificationStore>([])
