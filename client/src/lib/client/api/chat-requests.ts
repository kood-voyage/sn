import { PUBLIC_LOCAL_PATH } from "$env/static/public"
import type { ReturnEntryType, ReturnType } from "$lib/types/requests"
import type { UserType } from "$lib/types/user";

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

    if (!fetchResp.ok) throw new Error(await fetchResp.json())

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

    if (!fetchResp.ok) throw new Error(await fetchResp.json())

    return { ok: true, status: fetchResp.status }
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

    console.log(await fetchResp.json())

    if (!fetchResp.ok) throw new Error(await fetchResp.json())

    return { ok: true, status: fetchResp.status }

  } catch (err) {
    if (err instanceof Error) {
      return { ok: false, error: err, message: err.message }
    } else {
      return { ok: false, error: err, message: "Unknown Error" }
    }
  }
}

// export type GetChatUsers = {
//   id: string,
//   member_type: number
// }

type GetChatResp = ReturnType<UserType[]>
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

    if (!fetchResp.ok) throw new Error(await fetchResp.json())

    return { ok: true, data: json }
  } catch (err) {
    if (err instanceof Error) {
      return { ok: false, error: err, message: err.message }
    } else {
      return { ok: false, error: err, message: "Unknown Error" }
    }
  }
}

type AddLineResp = ReturnEntryType<"chatLine", ChatLine>
export async function addChatLine(chat_id: string, user_id: string, message: string, customFetch: Fetch = fetch): Promise<AddLineResp> {
  try {
    const fetchResp = await customFetch(`${PUBLIC_LOCAL_PATH}/api/v1/auth/chats/add/line`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Request-Method": "POST",
      },
      credentials: "include",
      body: JSON.stringify({
        chat_id,
        user_id,
        message
      })
    })
    const data = await fetchResp.json()

    if (!fetchResp.ok) throw new Error(data)
    return { ok: true, chatLine: data }
  } catch (err) {
    if (err instanceof Error) {
      return { ok: false, error: err, message: err.message }
    } else {
      return { ok: false, error: err, message: "Unknown Error" }
    }
  }
}

export type ChatLine = {
  chat_id: string
  created_at: string
  id: string
  message: string
  user_id: string
}

type GetLineResp = ReturnEntryType<"chatLines", ChatLine[]>
export async function getChatLines(chat_id: string, customFetch: Fetch = fetch): Promise<GetLineResp> {
  try {
    const fetchResp = await customFetch(`${PUBLIC_LOCAL_PATH}/api/v1/auth/chats/${chat_id}`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Request-Method": "GET",
      },
      credentials: "include",
    })

    const json = (await fetchResp.json()).data as ChatLine[]

    return { ok: true, chatLines: json }
  } catch (err) {
    if (err instanceof Error) {
      return { ok: false, error: err, message: err.message }
    } else {
      return { ok: false, error: err, message: "Unknown Error" }
    }
  }

}