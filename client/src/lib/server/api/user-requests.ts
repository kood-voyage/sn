import { LOCAL_PATH, WEBSITE_PATH } from "$env/static/private"
import type { RequestEvent } from "@sveltejs/kit"

export async function apiCreateUser(privacy_state: string, event: RequestEvent) {
  // console.log(event.cookies.get('at')?.valueOf())
  try {
    const resp = await fetch(`${LOCAL_PATH}/api/v1/auth/user/create/${privacy_state}`, {
      headers: {
        "Authorization": `Bearer ${event.cookies.get('at')?.valueOf()}`
      }

    })
    console.log(resp.statusText)
    if (resp.ok) {
      return { ok: resp.ok, status: resp.statusText }
    } else {
      return { ok: resp.ok, status: resp.statusText }
    }

  } catch (err) {
    if (err instanceof Error) {
      return { ok: false, error: err, message: err.message }
    } else {
      return { ok: false, error: err, message: "Unknown Error" }
    }
  }
}

export async function getUserFollowing(event: RequestEvent, user_id: string) {
  // const userResp = getUserId(event)
  // if (!userResp.ok) {
  //   return { ok: userResp.ok, error: userResp.error, message: userResp.message }
  // }

  try {
    const fetchResp = await fetch(`${WEBSITE_PATH}/api/v1/auth/user/following/${user_id}`, {
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

export async function getUserFollowers(event: RequestEvent, user_id: string) {
  // const userResp = getUserId(event)
  // if (!userResp.ok) {
  //   return { ok: userResp.ok, error: userResp.error, message: userResp.message }
  // }

  try {
    const fetchResp = await fetch(`${WEBSITE_PATH}/api/v1/auth/user/followers/${user_id}`, {
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