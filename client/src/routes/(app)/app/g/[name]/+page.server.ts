import { mainGetGroup } from "$lib/server/db/group"
import type { PageServerLoad } from "./$types"

export const load: PageServerLoad = async (event) => {

  console.log(event.params.name)
  const data = (await mainGetGroup(event, event.params.name))
  if (!data.ok) {
    console.error(data)
    return {
      username: "Undefined"
    }
  }

  return data

}