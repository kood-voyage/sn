import { mainGetAllUsers, type UserRowType } from "$lib/server/db/user";
import type { PageServerLoad } from "./$types";


type ReturnType = { ok: true, data: UserRowType[] } | {
  ok: false
}

export const load: PageServerLoad = async (): Promise<ReturnType> => {

  // console.log(event.params.name)
  // console.log((await event.parent()))
  const usersResp = (mainGetAllUsers())
  if (!usersResp.ok) {
    console.error(usersResp.message)
    return { ok: false }

  }

  const usersData = usersResp.data as UserRowType[]
  // console.log(typeof data)
  // console.log(typeof data.data)
  return { ok: true, data: usersData }

}