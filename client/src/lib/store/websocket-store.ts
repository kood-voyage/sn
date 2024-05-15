import type { ChatLine } from "$lib/client/api/chat-requests";
import type { UserType } from "$lib/types/user";
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

export type MessageStore = Array<ChatLine>
export const messageStore = writable<MessageStore>([])

export type UserStatusStore = { [key: string]: boolean }
export const userStatusStore = writable<UserStatusStore>({})

export type CustomNotification = {
  created_at: string
  id: string
  message: string
  parent_id: string
  source_id: string
  source_information: UserType
  target_id: string
  target_information: UserType
  type_id: number
}

export type NotificationStore = Array<CustomNotification>
export const notificationStore = writable<NotificationStore>([])
