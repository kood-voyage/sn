// FUNCTION FOR SENDING MESSAGE
// 1. ADD MESSAGE TO DB
// 2. ADD NOTIFICATION TO DB
// 3. send WEBSOCKET MESSAGE TO TARGET USER/GROUP

import { addChatLine } from "$lib/client/api/chat-requests"



export const sendMessageTo = async (message: string, chat_id: string, user_id: string) => {
  console.log("message >>>", message)
  console.log("chat_id >>>", chat_id)
  console.log("user_id >>>", user_id)
  const addLineResp = await addChatLine(chat_id, user_id, message)
  if (!addLineResp.ok) {
    console.error(addLineResp.error)
    return { ok: false }
  }

  console.log(addLineResp)
}