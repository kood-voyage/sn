import type { Actions } from "./$types";


export const actions: Actions = {
  profile: async (event) => {
    const data = await event.request.formData()

    const avatar = data.get("fileInputAvatar") as File
    const banner = data.get("fileInputBanner") as File
    const description = data.get("description")





    console.log(avatar)
    console.log(banner)
    console.log(description)





  }
}