import { getProfile } from "$lib/server/db"
import type { PageServerLoad } from './$types';








export const load: PageServerLoad = async ({params}) => {




    const data = getProfile(params.username)




  return data

}