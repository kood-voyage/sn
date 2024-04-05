import { mainGetAllUsers } from "$lib/server/db/user";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async () => {



  const data = mainGetAllUsers()
  if (!data.ok) {
    console.error(data.message)
    return {
      allUsers: {
        ...data

      }
    }
    // console.log(typeof data)
    // console.log(typeof data.data)

  }
  return { allUsers: { ...data } }
}