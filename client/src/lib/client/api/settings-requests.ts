


// export async function sendTo()


// function = {
//   default: async (event) => {
//     const data = await event.request.formData()
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