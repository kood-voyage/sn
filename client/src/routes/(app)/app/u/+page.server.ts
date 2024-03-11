import { mainGetAllUsers } from "$lib/server/db/user"







export const load = async () => {

  const data = mainGetAllUsers()

  if (data.ok) {
    return data
  } else {
    //
  }



}