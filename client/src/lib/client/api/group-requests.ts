
import { type RequestEvent } from "@sveltejs/kit";
import { type ReturnEntryType, type ReturnType } from "$lib/types/requests";
import { PUBLIC_LOCAL_PATH } from "$env/static/public";
import type { User, UserType } from "$lib/types/user";


export type GroupJson = {
  id: string,
  creator_id: string,
  name: string,
  description: string,
  image_path: Array<string>,
  privacy: string,
  members: null | Array<UserType>
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

export type GroupEventJson = {
  id: string,
  user_id: string,
  group_id: string,
  name: string,
  description: string,
  created_at: Date,
  date: Date,
  user_information: User,
  participants: User[],
  is_participant: boolean,
  event_status: string,
}

type Fetch = {
  (input: RequestInfo | URL, init?: RequestInit | undefined): Promise<Response>;
  (input: string | Request | URL, init?: RequestInit | undefined): Promise<Response>;
}

export type Group = ReturnEntryType<"group", GroupJson>


export async function GetGroup(group_name: string, customFetch: Fetch = fetch): Promise<Group> {

  try {

    const fetchResp = await customFetch(`${PUBLIC_LOCAL_PATH}/api/v1/auth/group/${group_name}`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Request-Method": "GET",
      },
      credentials: "include"
    })
    const json = (await fetchResp.json()).data as GroupJson

    return { ok: true, group: json }

  } catch (err) {
    if (err instanceof Error) {
      return { ok: false, error: err, message: err.message }
    } else {
      return { ok: false, error: err, message: "Unknown Error" }
    }
  }

}

export type AllGroups = ReturnEntryType<"allGroups", GroupJson[]>

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

export type AllGroupPosts = ReturnEntryType<"allGroupPosts", GroupPostJson[]>

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


export async function JoinGroup(group_name: string, customFetch: Fetch = fetch) {
  try {
    const fetchResp = await customFetch(`${PUBLIC_LOCAL_PATH}/api/v1/auth/group/join/${group_name.replace("_", " ")}`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Request-Method": "GET",
      },
      credentials: "include"
    });

    if (!fetchResp.ok) {
      const errorMessage = await fetchResp.json();
      return { ok: false, error: errorMessage, message: "error " }
    }

    const json = (await fetchResp.json()).data
    console.log("this is join group resp", json)
    return { ok: true, data: json }

  } catch (err) {
    if (err instanceof Error) {
      return { ok: false, error: err, message: err.message }
    } else {
      return { ok: false, error: err, message: "Unknown Error" }
    }
  }
}


export type GroupEvent = ReturnEntryType<"groupEvent", GroupEventJson>

export async function CreateGroupEvent(event:GroupEventJson, customFetch: Fetch = fetch) {
  try {
    const fetchResp = await customFetch(`${PUBLIC_LOCAL_PATH}/api/v1/auth/group/event/create`, {
      method: "POST",
      headers: {
          "Content-Type": "application/json",
          "Access-Control-Request-Method": "POST",
      },
      credentials: "include",
      body: JSON.stringify(event),
    })
  
      if (!fetchResp.ok) {
        const errorMessage = await fetchResp.json();
        return { ok: false, error: errorMessage, message: "error " }
      }
  
      const json = (await fetchResp.json()).data as GroupEventJson
      json.is_participant = false
      event.event_status = ""
      console.log("this is group event", json)
      return { ok: true, groupEvent: json }

  } catch (err) {
    if (err instanceof Error) {
      return { ok: false, error: err, message: err.message }
    } else {
      return { ok: false, error: err, message: "Unknown Error" }
    }
  }
}

export type AllGroupEvents = ReturnEntryType<"allGroupEvents", GroupEventJson[]>


export async function GetGroupEvents(group_id:string, customFetch: Fetch = fetch) {
  try {
    const fetchResp = await customFetch(`${PUBLIC_LOCAL_PATH}/api/v1/auth/group/${group_id}/event/all`, {
      method: "GET",
      headers: {
          "Content-Type": "application/json",
          "Access-Control-Request-Method": "GET",
      },
      credentials: "include",
    })
  
      if (!fetchResp.ok) {
        const errorMessage = await fetchResp.json();
        return { ok: false, error: errorMessage, message: "error " }
      }
  
      const json = (await fetchResp.json()).data as GroupEventJson[]
      console.log("These are all group events", json)
      if (json) {
        json.forEach((event) => {
          event.is_participant = false
          event.event_status = ""
        })
      }
      return { ok: true, allGroupEvents: json }

  } catch (err) {
    if (err instanceof Error) {
      return { ok: false, error: err, message: err.message }
    } else {
      return { ok: false, error: err, message: "Unknown Error" }
    }
  }
}