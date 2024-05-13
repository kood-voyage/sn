// import { getUser, setAvatar, setCover, setDescription } from "$lib/server/db/profile";
// import { saveUserAvatarToS3, saveUserCoverToS3 } from "$lib/client/images/upload";
// import { getUserIdFromCookie } from "$lib/client/jwt-handle";
// import type { Actions, PageServerLoad } from "./$types";


export const load: PageServerLoad = async (event) => {
  const idResp = getUserIdFromCookie(event)
  if (!idResp.ok) {
    return
  }

  const user_id = idResp.user_id as string
  const data = getUser(user_id)

  if (data.error) {
    return {}
  }
  return data?.data
}


// export const actions: Actions = {
//   default: async (event) => {
//     const data = await event.request.formData()
//     const user_resp = getUserIdFromCookie(event)
//     const id = user_resp.user_id as string
//     const avatar = data.get("fileInputAvatar") as File
//     const cover = data.get("fileInputCover") as File

//     if (avatar.type.startsWith('image/')) {
//       const avatarResp = await saveUserAvatarToS3(user_resp, avatar)
//       if (!avatarResp.ok) {
//         return avatarResp.error
//       }

//       const avatarUrl = avatarResp.resp
//       if (avatarUrl) setAvatar(avatarUrl, id)
//     } else {
//       console.log("The selected file is not an image.");
//     }

//     if (cover.type.startsWith('image/')) {
//       const coverResp = await saveUserCoverToS3(user_resp, cover)

//       if (!coverResp.ok) {
//         return coverResp.error
//       }

//       const coverUrl = coverResp.resp

//       if (coverUrl) setCover(coverUrl, id)

//     } else {
//       console.log("The selected file is not an image.");
//     }

//     const description = data.get("description") as string
//     if (description) setDescription(description, id)

//   }
// }