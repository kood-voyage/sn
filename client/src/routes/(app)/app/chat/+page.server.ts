import { addUserToChat, createChat, getAllChats, getChatUsers, type GetAllChats } from "$lib/server/api/chat-requests";
import { getUsersFromArray, mainGetAllUsers, type UserRowType } from "$lib/server/db/user";
import { fail, redirect } from "@sveltejs/kit";
import type { Actions, PageServerLoad } from "./$types";
import { v4 as uuidv4 } from 'uuid';
import type { ReturnType } from "$lib/types/requests";
import type { UserModel } from "$lib/types/user";
import { getUserIdFromCookie } from "$lib/server/jwt-handle";

export type ChatsWithUsers = {
  [key: string]: {
    users: UserModel[];
    group_id: string
  }
}
type DataType = { usersData: UserRowType[], chatsData: ChatsWithUsers }

type LoadResp = ReturnType<DataType>

export const load: PageServerLoad = async (event): Promise<LoadResp> => {

  // console.log(event.params.name)
  // console.log((await event.parent()))
  const usersResp = (mainGetAllUsers())
  // const usersResp = { ok: false, err: new Error("NO PROBLEM") }
  if (!usersResp.ok) {
    console.error(usersResp.message)
    return { ...usersResp }

  }

  const chatsResp = await getAllChats(event)
  if (!chatsResp.ok) {
    console.error(chatsResp.message)
    return { ...chatsResp }
  }



  const usersData = usersResp.data as UserRowType[]
  const chatsData = chatsResp.data as GetAllChats[]

  const chatsWithUserInfo: ChatsWithUsers = {}
  if (chatsData != undefined && chatsData.length != 0)
    for (const chat of chatsData) {
      const chatUserResp = await getChatUsers(event, chat.id)
      if (!chatUserResp.ok) {
        console.error(chatUserResp.message)
        return { ...chatUserResp }
      }
      const users = await getUsersFromArray(chatUserResp.data)

      chatsWithUserInfo[chat.id] = users
    }

  const dataOut = { usersData, chatsData: chatsWithUserInfo }

  return { ok: true, data: dataOut }
}

export const actions: Actions = {
  NewChat: async (event) => {
    const formData = await event.request.formData()

    // Create and Get variables like chat_id and user_id, respectively.
    const id = uuidv4()
    const target_user = formData.get("target") as string
    if (typeof target_user != "string" || target_user == undefined) {
      console.error(new Error("User ID not found on form"))
      return fail(400, { ok: false })
    }

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

    const userIDResp = getUserIdFromCookie(event)
    if (!userIDResp.ok) {
      console.error(userIDResp.error)
      return fail(400, { ok: false })
    }

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