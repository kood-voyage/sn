// import { PUBLIC_LOCAL_PATH } from "$env/static/public"
// import type { ReturnType } from "$lib/types/requests"
// import type { RequestEvent } from "@sveltejs/kit"

// export type GetAllChats = {
//   id: string,
//   group_id: string
// }

// type ChatsResp = ReturnType<GetAllChats[]>
// export async function getAllChats(event: RequestEvent): Promise<ChatsResp> {
//   try {
//     const fetchResp = await fetch(`${PUBLIC_LOCAL_PATH}/api/v1/auth/chats`, {
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

// type CreateChat = ReturnType<number>
// export async function createChat(event: RequestEvent, id: string, group_id: string = ""): Promise<CreateChat> {
//   try {
//     const fetchResp = await fetch(`${PUBLIC_LOCAL_PATH}/api/v1/auth/chats/create`, {
//       method: "post",
//       headers: {
//         "Authorization": `Bearer ${event.cookies.get('at')?.valueOf()}`,
//         'Content-Type': 'application/json' // Specify JSON content type
//       },
//       body: JSON.stringify({
//         id: id,
//         group_id: group_id
//       })
//     })

//     return { ok: true, data: fetchResp.status }

//   } catch (err) {
//     if (err instanceof Error) {
//       return { ok: false, error: err, message: err.message }
//     } else {
//       return { ok: false, error: err, message: "Unknown Error" }
//     }
//   }
// }

// type AddUserToChat = ReturnType<number>
// export async function addUserToChat(event: RequestEvent, target_user_id: string, chat_id: string): Promise<AddUserToChat> {
//   try {
//     const fetchResp = await fetch(`${PUBLIC_LOCAL_PATH}/api/v1/auth/chats/add/user`, {
//       method: "post",
//       headers: {
//         "Authorization": `Bearer ${event.cookies.get('at')?.valueOf()}`,
//         'Content-Type': 'application/json' // Specify JSON content type
//       },
//       body: JSON.stringify({
//         user_id: target_user_id,
//         chat_id: chat_id
//       })
//     })

//     return { ok: true, data: fetchResp.status }

//   } catch (err) {
//     if (err instanceof Error) {
//       return { ok: false, error: err, message: err.message }
//     } else {
//       return { ok: false, error: err, message: "Unknown Error" }
//     }
//   }
// }

// export type GetChatUsers = {
//   id: string,
//   member_type: number
// }

// type GetChatResp = ReturnType<GetChatUsers[]>
// export async function getChatUsers(event: RequestEvent, chat_id: string): Promise<GetChatResp> {
//   try {
//     const fetchResp = await fetch(`${PUBLIC_LOCAL_PATH}/api/v1/auth/chats/get/users/${chat_id}`, {
//       method: "get",
//       headers: {
//         "Authorization": `Bearer ${event.cookies.get('at')?.valueOf()}`,
//         'Content-Type': 'application/json' // Specify JSON content type
//       },
//     })
//     const json = (await fetchResp.json()).data

//     // console.log(json)

//     return { ok: true, data: json }
//   } catch (err) {
//     if (err instanceof Error) {
//       return { ok: false, error: err, message: err.message }
//     } else {
//       return { ok: false, error: err, message: "Unknown Error" }
//     }
//   }
// }