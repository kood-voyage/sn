import { getProfile } from "$lib/server/db/profile"
import type { PageServerLoad } from './$types';








export const load: PageServerLoad = async (event) => {

  // saveToS3("user1", { test: "worked" })
  // mainUpload()

  const data = (await getProfile(event, event.params.username))
  if (data.error) {
    console.error(data)
    return {
      username: "Undefined"
    }
  }

  // console.log(data)


  return data.user

}