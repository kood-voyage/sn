import { mainGetAllUsers, type UserRowType } from "$lib/server/db/user";
import type { Actions, PageServerLoad } from "./$types";


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

    return { success: true }
  },

};