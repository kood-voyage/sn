import { getAllChats, getChatUsers, type GetAllChats } from "$lib/client/api/chat-requests";
import type { ReturnEntryType } from "$lib/types/requests";
import type { UserType } from "$lib/types/user";
import { type GroupJson } from "$lib/client/api/group-requests";
import { GetAllUsers, currentUser } from "$lib/client/api/user-requests";
import type { LayoutLoad } from "./$types";

export type ChatsWithUsers = {
  [key: string]: {
    users: UserType[];
    group: GroupJson | { name: string }
  }
}
type DataType = { usersData: UserType[], chatsData: ChatsWithUsers }

type LoadResp = ReturnEntryType<"chatLoadData", DataType>

export const ssr = false

export const load: LayoutLoad = async ({ fetch }): Promise<LoadResp> => {

  // console.log(event.params.name)
  // console.log((await event.parent()))

  const currentUserResp = await currentUser(fetch)
  if (!currentUserResp.ok) {
    console.error(currentUserResp.error)
    return { ...currentUserResp }
  }

  const current_user_id = currentUserResp.data.id

  const usersResp = await GetAllUsers(fetch)
  // const usersResp = { ok: false, err: new Error("NO PROBLEM") }
  if (!usersResp.ok) {
    console.error(usersResp.error)
    return { ...usersResp }

  }
  // console.log(usersResp)

  const chatsResp = await getAllChats(fetch)
  if (!chatsResp.ok) {
    console.error(chatsResp.error)
    return { ...chatsResp }
  }

  // console.log("cahtsDATA >>>", chatsResp)



  const usersData = usersResp.allUsers as UserType[]
  const chatsData = chatsResp.data as GetAllChats[]

  const chatsWithUserInfo: ChatsWithUsers = {}
  if (chatsData != undefined && chatsData.length != 0)
    for (const chat of chatsData) {
      const chatUserResp = await getChatUsers(chat.id, fetch)
      if (!chatUserResp.ok) {
        console.error(chatUserResp.error)
        return { ...chatUserResp }
      }
      // console.log(current_user_id)
      const chatUsers = chatUserResp.data.filter((value) => value.id != current_user_id)
      // console.log(chatUsers)
      // const users = getUsersFromArray(chatUserResp.data)

      // ---------------- IF GROUPS CHATS IMPLEMENTATION -------------------------------------
      // let groupResp: GetGroup = { ok: false, error: new Error("GroupResp Not declared yet"), message: "No group Resp Yet" }
      // let groupJson: GroupJson | { name: string } = { name: chat.group_id }
      // if (chat.group_id != "" && chat.group_id != undefined) {
      //   groupResp = await getGroup(chat.group_id)
      //   if (!groupResp.ok) {
      //     console.log("WELP")
      //     console.error(groupResp.error)
      //     return { ok: groupResp.ok, error: groupResp.error, message: groupResp.message }
      //   }
      //   groupJson = groupResp.data
      //   console.log(groupResp)
      // }
      // -------------------------------------------------------------------------------------


      chatsWithUserInfo[chat.id] = { users: chatUsers, group: { name: "nop" } }
      // console.log("YO", chatsWithUserInfo)
    }

  const dataOut = { usersData, chatsData: chatsWithUserInfo }

  return { ok: true, chatLoadData: dataOut }
}



// export const actions: Actions = {
//   NewChat: async (event) => {
//     

// };

