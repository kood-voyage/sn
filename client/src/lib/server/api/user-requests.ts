import { LOCAL_PATH } from "$env/static/private"
import type { RequestEvent } from "@sveltejs/kit"
import { getUser } from "../db/profile"
import type { Resp } from "./api-types"

export async function apiCreateUser(privacy_state: string, event: RequestEvent) {
  // console.log(event.cookies.get('at')?.valueOf())
  try {
    const resp = await fetch(`${LOCAL_PATH}/api/v1/auth/user/create/${privacy_state}`, {
      headers: {
        "Authorization": `Bearer ${event.cookies.get('at')?.valueOf()}`
      }

    })
    // console.log("Create User >>>", resp.statusText)
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
    const fetchResp = await fetch(`${LOCAL_PATH}/api/v1/auth/user/following/${user_id}`, {
      headers: {
        "Authorization": `Bearer ${event.cookies.get('at')?.valueOf()}`
      }

    })
    // console.log("GETUSERFOLLOWING >>>", fetchResp.statusText)
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
    const fetchResp = await fetch(`${LOCAL_PATH}/api/v1/auth/user/followers/${user_id}`, {
      headers: {
        "Authorization": `Bearer ${event.cookies.get('at')?.valueOf()}`
      }

    })
    // console.log("GETUSERFOLLOWERS >>>", fetchResp.statusText)
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

export async function getUserPosts(event: RequestEvent, username: string): Promise<Resp> {
  try {
    const userInfo = getUser(username)
    if (!userInfo.ok) {
      throw new Error("User with this username not found")
    }

    const user_id = userInfo.data?.id

    const url = `${LOCAL_PATH}/api/v1/auth/user/posts/${user_id}`
    // console.log("url >>", url)

    const fetchResp = await fetch(url, {
      headers: {
        "Authorization": `Bearer ${event.cookies.get('at')?.valueOf()}`
      }

    })
    // console.log("GETUSERPOSTS >>>", fetchResp.statusText)
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