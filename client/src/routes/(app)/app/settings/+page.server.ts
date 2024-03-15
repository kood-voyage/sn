import { saveUserAvatarToS3 } from "$lib/server/images/upload";
import { getUserIdFromCookie } from "$lib/server/jwt-handle";
import type { Actions, PageServerLoad } from "./$types";

export const load: PageServerLoad = async (event) => {

  // return { img: (await getUserBannerFromS3(event))?.valueOf() }

}


export const actions: Actions = {
  profile: async (event) => {
    const data = await event.request.formData()


    const user_id = getUserIdFromCookie(event)



    const avatar = data.get("fileInputAvatar") as File
    // const banner = data.get("fileInputBanner") as File
    // const description = data.get("description")

    // const str = (await getUserBannerFromS3(event))

    const {ok,resp,error} = await saveUserAvatarToS3(user_id, avatar)

    if (!ok){
      return error
    }

    console.log(resp)

    








    // console.log(avatar)
    // console.log(banner)
    // console.log(description)







  }
}