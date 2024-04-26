import { PUBLIC_LOCAL_PATH } from "$env/static/public"
import type { CreateUser } from "$lib/types/user"
import type { RequestEvent } from "@sveltejs/kit"



export async function apiCreateUser(user: CreateUser) {
  try {
    const resp = await fetch(`${PUBLIC_LOCAL_PATH}/api/v1/user/create`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Request-Method": "POST",
      },
      credentials: "include",
      body: JSON.stringify(user)
    })
    if (resp.ok) {
      return { ok: resp.ok, status: resp.statusText }
    } else {
      return { ok: resp.ok, status: resp.statusText }
    }

  } catch (err) {
    console.log("ERRRRRR", err)
    if (err instanceof Error) {
      return { ok: false, error: err, message: err.message }
    } else {
      return { ok: false, error: err, message: "Unknown Error" }
    }
  }
}

export async function getUserFollowing(event: RequestEvent, user_id: string) {
  try {
    const fetchResp = await fetch(`${PUBLIC_LOCAL_PATH}/api/v1/auth/user/following/${user_id}`, {
      headers: {
        "Authorization": `Bearer ${event.cookies.get('at')}`
      }

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

export async function getUserFollowers(event: RequestEvent, user_id: string) {

  try {
    const fetchResp = await fetch(`${PUBLIC_LOCAL_PATH}/api/v1/auth/user/followers/${user_id}`, {
      headers: {
        "Authorization": `Bearer ${event.cookies.get('at')?.valueOf()}`
      }

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

