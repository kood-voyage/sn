// import { getUserPosts } from "$lib/server/api/user-requests";
import { getProfile } from "$lib/server/db/profile"
// import { getFromS3, mainUpload, saveToS3 } from "$lib/server/images/upload";
import type { PageServerLoad } from './$types';







export const load: PageServerLoad = async (event) => {

  // getFromS3("user1")
  // saveToS3("user1", { test: "worked" })
  // mainUpload()
  // const dat = await getUserPosts(event, event.params.username)
  // console.log(dat)

  const data = (await getProfile(event, event.params.username))
  if (data.error) {
    console.error(data)
    return {
      username: "Undefined"
    }
  }
  console.log(data)

  // console.log(data)


  return data.user

}