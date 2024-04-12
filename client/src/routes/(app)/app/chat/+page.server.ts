import { createChat } from "$lib/server/api/chat-requests";
import { mainGetAllUsers, type UserRowType } from "$lib/server/db/user";
import { fail } from "@sveltejs/kit";
import type { Actions, PageServerLoad } from "./$types";
import { v4 as uuidv4 } from 'uuid';


type ReturnType = { ok: true, data: UserRowType[] } | {
  ok: false
}

export const load: PageServerLoad = async (): Promise<ReturnType> => {

  // console.log(event.params.name)
  // console.log((await event.parent()))
  const usersResp = (mainGetAllUsers())
  // const usersResp = { ok: false, err: new Error("NO PROBLEM") }
  if (!usersResp.ok) {
    console.error(usersResp.message)
    return { ok: false }

  }

  const usersData = usersResp.data as UserRowType[]
  // console.log(typeof data)
  // console.log(typeof data.data)
  return { ok: true, data: usersData }

}

export const actions: Actions = {
  NewChat: async (event) => {
    const formData = await event.request.formData()
    console.log(formData.get("target"))
    console.log("RAN")

    const id = uuidv4()

    const createResp = await createChat(event, id)
    if (!createResp.ok) {
      console.error(createResp.error)
      return fail(400, { ok: false })
    }
    if (createResp.data[0] != 200) {
      const err = new Error("Creating chat failed with Status Code >>> " + createResp.data[0].toString())
      console.error(err)
      return fail(createResp.data[0], { ok: false })
    }



    console.log(createResp.data)


    return { status: createResp.data[0], success: true }
  },

};