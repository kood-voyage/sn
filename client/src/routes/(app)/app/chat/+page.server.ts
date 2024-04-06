import { mainGetAllUsers } from "$lib/server/db/user";
import type { PageServerLoad } from "./$types";


export const load: PageServerLoad = async () => {


  // console.log(event.params.name)
  // console.log((await event.parent()))
  const usersResp = (mainGetAllUsers())
  if (!usersResp.ok) {
    console.error(usersResp.message)
    return { data: usersResp.data }

  }
  // console.log(typeof data)
  // console.log(typeof data.data)
  return { data: usersResp.data }

}