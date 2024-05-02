import { PUBLIC_LOCAL_PATH } from "$env/static/public"
import type { ReturnEntryType, ReturnType } from "$lib/types/requests"
import type { RequestEvent } from "@sveltejs/kit"


type Fetch = {
  (input: RequestInfo | URL, init?: RequestInit | undefined): Promise<Response>;
  (input: string | Request | URL, init?: RequestInit | undefined): Promise<Response>;
}

export type GetAllChats = {
  id: string,
  group_id: string
}

type ChatsResp = ReturnType<GetAllChats[]>
export async function getAllChats(customFetch: Fetch = fetch): Promise<ChatsResp> {
  try {
    const fetchResp = await customFetch(`${PUBLIC_LOCAL_PATH}/api/v1/auth/chats`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Request-Method": "GET",
      },
      credentials: "include"
    })
    const json = (await fetchResp.json()).data

    return { ok: true, data: json }
  } catch (err) {
    if (err instanceof Error) {
      return { ok: false, error: err, message: err.message }
    } else {
      return { ok: false, error: err, message: "Unknown Error" }
    }
  }

}

type CreateChat = ReturnEntryType<"status", number>
export async function createChat(id: string, group_id: string = ""): Promise<CreateChat> {
  try {
    const fetchResp = await fetch(`${PUBLIC_LOCAL_PATH}/api/v1/auth/chats/create`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Request-Method": "POST",
      },
      credentials: "include",
      body: JSON.stringify({
        id: id,
        group_id: group_id
      })
    })
    console.log(await fetchResp.json())

    if (fetchResp.ok) {
      return { ok: true, status: fetchResp.status }
    } else {
      return { ok: false, error: new Error("Fetch Response went wrong trying to create chat"), message: await fetchResp.json() }
    }


  } catch (err) {
    if (err instanceof Error) {
      return { ok: false, error: err, message: err.message }
    } else {
      return { ok: false, error: err, message: "Unknown Error" }
    }
  }
}

type AddUserToChat = ReturnEntryType<"status", number>
export async function addUserToChat(target_user_id: string, chat_id: string): Promise<AddUserToChat> {
  try {
    const fetchResp = await fetch(`${PUBLIC_LOCAL_PATH}/api/v1/auth/chats/add/user`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Request-Method": "POST",
      },
      credentials: "include",
      body: JSON.stringify({
        user_id: target_user_id,
        chat_id: chat_id
      })
    })

    return { ok: true, status: fetchResp.status }

  } catch (err) {
    if (err instanceof Error) {
      return { ok: false, error: err, message: err.message }
    } else {
      return { ok: false, error: err, message: "Unknown Error" }
    }
  }
}

export type GetChatUsers = {
  id: string,
  member_type: number
}

type GetChatResp = ReturnType<GetChatUsers[]>
export async function getChatUsers(chat_id: string, customFetch: Fetch = fetch): Promise<GetChatResp> {
  try {
    const fetchResp = await customFetch(`${PUBLIC_LOCAL_PATH}/api/v1/auth/chats/get/users/${chat_id}`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Request-Method": "GET",
      },
      credentials: "include"
    })

    const json = (await fetchResp.json()).data

    // console.log(json)

    return { ok: true, data: json }
  } catch (err) {
    if (err instanceof Error) {
      return { ok: false, error: err, message: err.message }
    } else {
      return { ok: false, error: err, message: "Unknown Error" }
    }
  }
}