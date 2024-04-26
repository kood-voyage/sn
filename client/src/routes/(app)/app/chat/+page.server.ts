import { addUserToChat, createChat, getAllChats, getChatUsers, type GetAllChats } from "$lib/server/api/chat-requests";
import { getUsersFromArray, mainGetAllUsers, type UserRowType } from "$lib/server/db/user";
import { fail, redirect } from "@sveltejs/kit";
import type { Actions, PageServerLoad } from "./$types";
import { v4 as uuidv4 } from 'uuid';
import type { ReturnToClientType } from "$lib/types/requests";
import type { UserModel } from "$lib/types/user";
import { getUserIdFromCookie } from "$lib/client/jwt-handle";
import { getGroup, type GetGroup, type GroupJson } from "$lib/server/api/group-requests";

export type ChatsWithUsers = {
  [key: string]: {
    users: UserModel[];
    group: GroupJson | { name: string }
  }
}
type DataType = { usersData: UserRowType[], chatsData: ChatsWithUsers }

type LoadResp = ReturnToClientType<DataType>

export const load: PageServerLoad = async (event): Promise<LoadResp> => {

  // console.log(event.params.name)
  // console.log((await event.parent()))
  const usersResp = (mainGetAllUsers())
  // const usersResp = { ok: false, err: new Error("NO PROBLEM") }
  if (!usersResp.ok) {
    console.error(usersResp.error)
    return { ok: usersResp.ok, message: usersResp.message }

  }

  const chatsResp = await getAllChats(event)
  if (!chatsResp.ok) {
    console.error(chatsResp.error)
    return { ok: chatsResp.ok, message: chatsResp.message }
  }




  const usersData = usersResp.data as UserRowType[]
  const chatsData = chatsResp.data as GetAllChats[]

  const chatsWithUserInfo: ChatsWithUsers = {}
  if (chatsData != undefined && chatsData.length != 0)
    for (const chat of chatsData) {
      const chatUserResp = await getChatUsers(event, chat.id)
      if (!chatUserResp.ok) {
        console.error(chatUserResp.error)
        return { ok: chatUserResp.ok, message: chatUserResp.message }
      }
      const users = getUsersFromArray(chatUserResp.data)


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


      chatsWithUserInfo[chat.id] = { users, group: groupJson }
    }

  const dataOut = { usersData, chatsData: chatsWithUserInfo }

  return { ok: true, data: dataOut }
}

export const actions: Actions = {
  NewChat: async (event) => {
    const formData = await event.request.formData()

    // Create and Get variables like chat_id and user_id, respectively.
    const id = uuidv4()

    // Create Chat 
    const createResp = await createChat(event, id)
    if (!createResp.ok) {
      console.error(createResp.error)
      return fail(400, { ok: false })
    }
    if (createResp.data <= 200, createResp.data >= 299) {
      const err = new Error("Creating chat failed with Status Code >>> " + createResp.data.toString())
      console.error(err)
      return fail(createResp.data, { ok: false })
    }

    // Get the target User to add to chat
    const target_user = formData.get("target") as string
    if (typeof target_user != "string" || target_user == undefined) {
      console.error(new Error("User ID not found on form"))
      return fail(400, { ok: false })
    }

    // Get the user creating the post 
    const userIDResp = getUserIdFromCookie(event)
    if (!userIDResp.ok) {
      console.error(userIDResp.error)
      return fail(400, { ok: false })
    }



    // Add User to the created Chat
    let addUserResp = await addUserToChat(event, target_user, id)
    if (!addUserResp.ok) {
      console.error(addUserResp.error)
      return fail(400, { ok: false })
    }
    if (addUserResp.data <= 200, addUserResp.data >= 299) {
      const err = new Error("Adding User To Chat failed with Status Code >>> " + addUserResp.data.toString())
      console.error(err)
      return fail(addUserResp.data, { ok: false })
    }

    // If the users are the same then just return and don't try to add a second one
    if (userIDResp.user_id == target_user) {
      return redirect(303, "/app/chat")
    }

    // Otherwise add the second user
    addUserResp = await addUserToChat(event, userIDResp.user_id as string, id)
    if (!addUserResp.ok) {
      console.error(addUserResp.error)
      return fail(400, { ok: false })
    }
    if (addUserResp.data <= 200, addUserResp.data >= 299) {
      const err = new Error("Adding User To Chat failed with Status Code >>> " + addUserResp.data.toString())
      console.error(err)
      return fail(addUserResp.data, { ok: false })
    }


    return redirect(303, "/app/chat")
  },

};