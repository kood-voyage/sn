import { getAllChats, getChatUsers, type GetAllChats } from "$lib/client/api/chat-requests";
// import { getUsersFromArray, mainGetAllUsers, type UserRowType } from "$lib/client/db/user";
// import { fail, redirect } from "@sveltejs/kit";
import type { PageLoad } from "./$types";
import type { ReturnToClientType } from "$lib/types/requests";
import type { UserModel, UserType } from "$lib/types/user";
import { getGroup, type GetGroup, type GroupJson } from "$lib/client/api/group-requests";
import { GetAllUsers } from "$lib/client/api/user-requests";

export type ChatsWithUsers = {
  [key: string]: {
    users: UserModel[];
    group: GroupJson | { name: string }
  }
}
type DataType = { usersData: UserType[], chatsData: ChatsWithUsers }

type LoadResp = ReturnToClientType<DataType>

export const ssr = false

export const load: PageLoad = async ({ fetch }): Promise<LoadResp> => {

  // console.log(event.params.name)
  // console.log((await event.parent()))

  const usersResp = await GetAllUsers(fetch)
  // const usersResp = { ok: false, err: new Error("NO PROBLEM") }
  if (!usersResp.ok) {
    console.error(usersResp.error)
    return { ok: usersResp.ok, message: usersResp.message }

  }
  console.log(usersResp)

  const chatsResp = await getAllChats(fetch)
  if (!chatsResp.ok) {
    console.error(chatsResp.error)
    return { ok: chatsResp.ok, message: chatsResp.message }
  }

  console.log(chatsResp)



  const usersData = usersResp.allUsers as UserType[]
  const chatsData = chatsResp.data as GetAllChats[]

  const chatsWithUserInfo: ChatsWithUsers = {}
  if (chatsData != undefined && chatsData.length != 0)
    for (const chat of chatsData) {
      const chatUserResp = await getChatUsers(chat.id, fetch)
      if (!chatUserResp.ok) {
        console.error(chatUserResp.error)
        return { ok: chatUserResp.ok, message: chatUserResp.message }
      }
      console.log(chat)
      // const users = getUsersFromArray(chatUserResp.data)


      let groupResp: GetGroup = { ok: false, error: new Error("GroupResp Not declared yet"), message: "No group Resp Yet" }
      let groupJson: GroupJson | { name: string } = { name: chat.group_id }
      if (chat.group_id != "" && chat.group_id != undefined) {
        groupResp = await getGroup(event, chat.group_id)
        if (!groupResp.ok) {
          console.log("WELP")
          console.error(groupResp.error)
          return { ok: groupResp.ok, message: groupResp.message }
        }
        groupJson = groupResp.data
        console.log(groupResp)
      }


      chatsWithUserInfo[chat.id] = { users: [], group: groupJson }
    }

  const dataOut = { usersData, chatsData: chatsWithUserInfo }

  return { ok: true, data: dataOut }
}

// export const actions: Actions = {
//   NewChat: async (event) => {
//     

// };