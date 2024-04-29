import { PUBLIC_LOCAL_PATH } from "$env/static/public"
import type { ReturnType } from "$lib/types/requests";
import type { UserModel, UserType } from "$lib/types/user"
import type { SignIn } from "../../../routes/(auth)/signin/type"


type Fetch = {
  (input: RequestInfo | URL, init?: RequestInit | undefined): Promise<Response>;
  (input: string | Request | URL, init?: RequestInit | undefined): Promise<Response>;
}

export async function GetAllUsers(customFetch?: Fetch) {
  if (!customFetch) {
    customFetch = fetch
  }

  try {
    const resp = await customFetch(`${PUBLIC_LOCAL_PATH}/api/v1/auth/user/all`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Request-Method": "GET",
      },
      credentials: "include"
    })


    if (resp.ok) {
      return { ok: resp.ok, allUsers: (await resp.json()).data as UserType[] }
    } else {
      return { ok: resp.ok, allUsers: (await resp.json()).data as UserType[] }
    }

  } catch (err) {
    if (err instanceof Error) {

      console.log(err)
      return { ok: false, error: err, message: err.message }
    } else {

      console.log(err)
      return { ok: false, error: err, message: "Unknown Error" }
    }
  }

}

export async function RegisterUser(user: UserModel) {
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

      console.log(err)
      return { ok: false, error: err, message: err.message }
    } else {

      console.log(err)
      return { ok: false, error: err, message: "Unknown Error" }
    }
  }
}

type CurrentUser = ReturnType<UserType>

export async function currentUser(customFetch?: Fetch): Promise<CurrentUser> {
  if (!customFetch) {
    customFetch = fetch
  }

  try {
    const resp = await customFetch(`${PUBLIC_LOCAL_PATH}/api/v1/auth/user/current`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Request-Method": "GET",
      },
      credentials: "include",
    })



    const data = (await resp.json()).data


    if (resp.ok) {
      return { ok: resp.ok, data }
    } else {
      throw new Error(data)
    }


  } catch (err) {
    console.log("ERRRRRR", err)
    if (err instanceof Error) {

      console.log(err)
      return { ok: false, error: err, message: err.message }
    } else {

      console.log(err)
      return { ok: false, error: err, message: "Unknown Error" }
    }
  }
}


export async function LoginUser(credentials: SignIn) {

  try {
    const resp = await fetch(`${PUBLIC_LOCAL_PATH}/api/v1/user/login`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Request-Method": "POST",
      },
      credentials: "include",
      body: JSON.stringify(credentials)
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

export async function getUserFollowing(user_id: string, customFetch?: Fetch) {
  if (!customFetch) {
    customFetch = fetch
  }
  try {
    const fetchResp = await customFetch(`${PUBLIC_LOCAL_PATH}/api/v1/auth/user/following/${user_id}`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Request-Method": "GET",
      },
      credentials: "include",

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

export async function getUserFollowers(user_id: string, customFetch?: Fetch) {
  if (!customFetch) {
    customFetch = fetch
  }

  try {
    const fetchResp = await customFetch(`${PUBLIC_LOCAL_PATH}/api/v1/auth/user/followers/${user_id}`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Request-Method": "GET",
      },
      credentials: "include",
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

