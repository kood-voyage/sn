// import { LOCAL_PATH } from "$env/static/private";
// import { type RequestEvent } from "@sveltejs/kit";
// import { type ReturnType } from "$lib/types/requests";


// export type GroupJson = {
//   id: string,
//   creator_id: string,
//   name: string,
//   description: string,
//   image_path: Array<string>,
//   privacy: string,
//   members: null | Array<{ id: string, member_type: number }>
// }

// export type GroupPostJson = {
//   id: string,
//   title: string,
//   content: string,
//   image_path: string[],
//   user_id: string,
//   community_id: string | null,
//   created_at: Date,
//   privacy: string | null
// }

// export type GetGroup = ReturnType<GroupJson>

// export async function getGroup(event: RequestEvent, group_name: string): Promise<GetGroup> {

//   try {

//     const fetchResp = await fetch(`${LOCAL_PATH}/api/v1/auth/group/${group_name}`, {
//       headers: {
//         "Authorization": `Bearer ${event.cookies.get('at')?.valueOf()}`
//       }
//     })
//     const json = (await fetchResp.json()).data


//     console.log(fetchResp.status)
//     console.log(json)

//     return { ok: true, data: json }

//   } catch (err) {
//     if (err instanceof Error) {
//       return { ok: false, error: err, message: err.message }
//     } else {
//       return { ok: false, error: err, message: "Unknown Error" }
//     }
//   }

// }

// export async function getGroups(event: RequestEvent) {

//   try {
//     const fetchResp = await fetch(`${LOCAL_PATH}/api/v1/auth/group`, {
//       headers: {
//         "Authorization": `Bearer ${event.cookies.get('at')?.valueOf()}`
//       }
//     })
//     const json = (await fetchResp.json()).data as GroupJson[]

//     return { ok: true, data: json }

//   } catch (err) {
//     if (err instanceof Error) {
//       return { ok: false, error: err, message: err.message }
//     } else {
//       return { ok: false, error: err, message: "Unknown Error" }
//     }
//   }
// }

// export async function getGroupPosts(event: RequestEvent, group_name: string): Promise<ReturnType<GroupPostJson[]>> {

//   try {
//     const fetchResp = await fetch(`${LOCAL_PATH}/api/v1/auth/group/posts/${group_name.replace("_", " ")}`, {
//       headers: {
//         "Authorization": `Bearer ${event.cookies.get('at')?.valueOf()}`
//       }
//     })
//     const json = (await fetchResp.json()).data as GroupPostJson[]

//     console.log("posts", json)
//     return { ok: true, data: json }

//   } catch (err) {
//     if (err instanceof Error) {
//       return { ok: false, error: err, message: err.message }
//     } else {
//       return { ok: false, error: err, message: "Unknown Error" }
//     }
//   }
// }

// export async function joinGroup(event: RequestEvent, group_name: string) {
//   try {
//     const fetchResp = await fetch(`${LOCAL_PATH}/api/v1/auth/group/join/${group_name.replace("_", " ")}`, {
//       method: "GET",
//       headers: {
//         "Authorization": `Bearer ${event.cookies.get('at')?.valueOf()}`
//       }
//     });

//     if (!fetchResp.ok) {
//       const errorMessage = await fetchResp.json();
//       return { ok: false, error: errorMessage, message: "error " }
//     }

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