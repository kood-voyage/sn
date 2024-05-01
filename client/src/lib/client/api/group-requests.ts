
import { type RequestEvent } from "@sveltejs/kit";
import { type ReturnEntryType, type ReturnType } from "$lib/types/requests";
import { PUBLIC_LOCAL_PATH } from "$env/static/public";


export type GroupJson = {
  id: string,
  creator_id: string,
  name: string,
  description: string,
  image_path: Array<string>,
  privacy: string,
  members: null | Array<{ id: string, member_type: number }>
}

export type GroupPostJson = {
  id: string,
  title: string,
  content: string,
  image_path: string[],
  user_id: string,
  community_id: string | null,
  created_at: Date,
  privacy: string | null
}

type Fetch = {
  (input: RequestInfo | URL, init?: RequestInit | undefined): Promise<Response>;
  (input: string | Request | URL, init?: RequestInit | undefined): Promise<Response>;
}

type Group = ReturnEntryType<"group", GroupJson>


export async function GetGroup(group_name: string, customFetch: Fetch = fetch): Promise<Group> {

  try {

    const fetchResp = await fetch(`${PUBLIC_LOCAL_PATH}/api/v1/auth/group/${group_name}`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Request-Method": "GET",
      },
      credentials: "include"
    })
    const json = (await fetchResp.json()).data as GroupJson


    console.log(fetchResp.status)
    console.log(json)

    return { ok: true, group: json }

  } catch (err) {
    if (err instanceof Error) {
      return { ok: false, error: err, message: err.message }
    } else {
      return { ok: false, error: err, message: "Unknown Error" }
    }
  }

}

type AllGroups = ReturnEntryType<"allGroups", GroupJson[]>

export async function GetAllGroups(customFetch: Fetch = fetch): Promise<AllGroups> {
  try {
    const fetchResp = await customFetch(`${PUBLIC_LOCAL_PATH}/api/v1/auth/group`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Request-Method": "GET",
      },
      credentials: "include"
    })
    const json = (await fetchResp.json()).data as GroupJson[]

    return { ok: true, allGroups: json }
  } catch (err) {
    if (err instanceof Error) {
      return { ok: false, error: err, message: err.message }
    } else {
      return { ok: false, error: err, message: "Unknown Error" }
    }
  }
}

type AllGroupPosts = ReturnEntryType<"allGroupPosts", GroupPostJson[]>

export async function GetGroupPosts(group_name: string, customFetch: Fetch = fetch): Promise<AllGroupPosts> {

  try {
    const fetchResp = await customFetch(`${PUBLIC_LOCAL_PATH}/api/v1/auth/group/posts/${group_name.replace("_", " ")}`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Request-Method": "GET",
      },
      credentials: "include"
    })
    const json = (await fetchResp.json()).data as GroupPostJson[]

    console.log("posts", json)
    return { ok: true, allGroupPosts: json }

  } catch (err) {
    if (err instanceof Error) {
      return { ok: false, error: err, message: err.message }
    } else {
      return { ok: false, error: err, message: "Unknown Error" }
    }
  }
}

export async function joinGroup(event: RequestEvent, group_name: string) {
  try {
    const fetchResp = await fetch(`${PUBLIC_LOCAL_PATH}/api/v1/auth/group/join/${group_name.replace("_", " ")}`, {
      method: "GET",
      headers: {
        "Authorization": `Bearer ${event.cookies.get('at')?.valueOf()}`
      }
    });

    if (!fetchResp.ok) {
      const errorMessage = await fetchResp.json();
      return { ok: false, error: errorMessage, message: "error " }
    }

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