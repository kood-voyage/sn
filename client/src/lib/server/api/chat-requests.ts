import { LOCAL_PATH } from "$env/static/private"
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