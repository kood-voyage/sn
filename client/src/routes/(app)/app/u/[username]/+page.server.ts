import { getProfile } from "$lib/server/db/profile"
import type { PageServerLoad } from './$types';








export const load: PageServerLoad = async (event) => {




  const data = (await getProfile(event, event.params.username))
  if (data.error) {
    return {}
  }




  return data.user

}