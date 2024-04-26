// import { LOCAL_PATH } from "$env/static/private"
// import type { RequestEvent } from "@sveltejs/kit"





// export async function getUserPosts(event: RequestEvent, user_id: string) {
//   try {
//     const fetchResp = await fetch(`${LOCAL_PATH}/api/v1/auth/user/posts/${user_id}`, {
//       headers: {
//         "Authorization": `Bearer ${event.cookies.get('at')?.valueOf()}`
//       }

//     })
//     const json = (await fetchResp.json()).data



//     return { ok: true, data: json }

//   } catch (err) {
//     if (err instanceof Error) {
//       return { ok: false, error: err, message: err.message }
//     } else {
//       return { ok: false, error: err, message: "Unknown Error" }
//     }
//   }

// }