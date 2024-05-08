import { getChatLines, getChatUsers, type ChatLine } from "$lib/client/api/chat-requests";
import { currentUser } from "$lib/client/api/user-requests";
import type { ReturnEntryType } from "$lib/types/requests";
import type { PageLoad } from "./$types";


export type DisplayData = {
  chat_id: string;
  user_id: string;
  display_name: string;
  cover: string;
  avatar: string;
};


type linesData = ChatLine[]

type LoadOut = ReturnEntryType<"chatData", { display_data: DisplayData, lines_data: linesData }>

export const ssr = false
// export const csr = false

export const load: PageLoad = async ({ fetch, params }): Promise<LoadOut> => {
  // console.log(params.chatid)

  const currentUserResp = await currentUser(fetch)
  if (!currentUserResp.ok) {
    console.error(currentUserResp.error)
    return { ...currentUserResp }
  }
  const current_user_id = currentUserResp.data.id

  const usersResp = await getChatUsers(params.chatid, fetch)
  if (!usersResp.ok) {
    console.error(usersResp.error)
    return { ...usersResp }
  }

  const chatUsers = usersResp.data.filter((value) => value.id != current_user_id)
  const display_data = {
    chat_id: params.chatid,
    user_id: chatUsers[0].id,
    display_name: chatUsers[0].username,
    cover: chatUsers[0].cover as string,
    avatar: chatUsers[0].avatar as string
  };


  const linesResp = await getChatLines(params.chatid, fetch)
  if (!linesResp.ok) {
    console.error(linesResp.error)
    return { ...linesResp }
  }

  const lines_data = linesResp.chatLines
  return { ok: true, chatData: { lines_data, display_data } }

}