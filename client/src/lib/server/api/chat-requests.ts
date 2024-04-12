import { LOCAL_PATH } from "$env/static/private"
import type { ReturnType } from "$lib/types/requests"
import type { RequestEvent } from "@sveltejs/kit"

// type chatType = {}

// type ChatsResp = ReturnType<>


export async function getUserChats(event: RequestEvent) {
  try {
    const fetchResp = await fetch(`${LOCAL_PATH}/api/v1/auth/chats`, {
      headers: {
        "Authorization": `Bearer ${event.cookies.get('at')?.valueOf()}`
      }

    })
    const json = (await fetchResp.json()).data

    console.log(json)


    return { ok: true, data: json }

  } catch (err) {
    if (err instanceof Error) {
      return { ok: false, error: err, message: err.message }
    } else {
      return { ok: false, error: err, message: "Unknown Error" }
    }
  }

}

type CreateChat = ReturnType<number>

export async function createChat(event: RequestEvent, id: string, group_id: string = ""): Promise<CreateChat> {
  try {
    const fetchResp = await fetch(`${LOCAL_PATH}/api/v1/auth/chats/create`, {
      method: "post",
      headers: {
        "Authorization": `Bearer ${event.cookies.get('at')?.valueOf()}`,
        'Content-Type': 'application/json' // Specify JSON content type
      },
      body: JSON.stringify({
        id: id,
        group_id: group_id
      })
    })
    const json = (await fetchResp.json()).data as string[]

    // console.log(fetchResp.status)
    console.log("json", json)
    console.log(JSON.stringify({
      id: id,
      group_id: group_id
    }))

    return { ok: true, data: [fetchResp.status] }

  } catch (err) {
    if (err instanceof Error) {
      console.log("YEP")
      return { ok: false, error: err, message: err.message }
    } else {
      return { ok: false, error: err, message: "Unknown Error" }
    }
  }
}