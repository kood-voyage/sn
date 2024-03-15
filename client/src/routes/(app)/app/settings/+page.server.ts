import { saveUserAvatarToS3 } from "$lib/server/images/upload";
import type { Actions, PageServerLoad } from "./$types";

export const load: PageServerLoad = async (event) => {

  // return { img: (await getUserBannerFromS3(event))?.valueOf() }

}


export const actions: Actions = {
  profile: async (event) => {
    const data = await event.request.formData()

    const avatar = data.get("fileInputAvatar") as File
    const banner = data.get("fileInputBanner") as File
    const description = data.get("description")

    // const str = (await getUserBannerFromS3(event))

    saveUserAvatarToS3(event, avatar)

    // console.log(avatar)
    // console.log(banner)
    // console.log(description)


    




  }
}