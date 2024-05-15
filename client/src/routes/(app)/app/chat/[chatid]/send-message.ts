

import { addChatLine, type ChatLine } from "$lib/client/api/chat-requests"
import { sendMessage } from "$lib/client/websocket"
import type { ReturnEntryType } from "$lib/types/requests"

type SendMessage = ReturnEntryType<"chatLine", ChatLine>

export const sendMessageTo = async (message: string, chat_id: string, user_id: string): Promise<SendMessage> => {
  console.log("message >>>", message)
  console.log("chat_id >>>", chat_id)
  console.log("user_id >>>", user_id)
  const addLineResp = await addChatLine(chat_id, user_id, message)
  if (!addLineResp.ok) {
    console.error(addLineResp.error)
    return { ...addLineResp }
  }

  return { ...addLineResp }
}

export const sendMessageByWebsocket = (target_id: string, source_id: string, message: ChatLine) => {
  sendMessage('message', 'direct', target_id, source_id, message);
}