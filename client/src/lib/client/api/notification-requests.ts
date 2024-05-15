import { PUBLIC_LOCAL_PATH } from "$env/static/public";
import type { CustomNotification, NotificationStore } from "$lib/store/websocket-store";
import type { ReturnEntryType } from "$lib/types/requests";

type Fetch = {
  (input: RequestInfo | URL, init?: RequestInit | undefined): Promise<Response>;
  (input: string | Request | URL, init?: RequestInit | undefined): Promise<Response>;
}



type GetUserNotifications = ReturnEntryType<"notifications", NotificationStore>
export async function getUserNotifications(customFetch: Fetch = fetch): Promise<GetUserNotifications> {
  try {
    const fetchResp = await customFetch(`${PUBLIC_LOCAL_PATH}/api/v1/auth/user/notifications`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Request-Method": "GET",
      },
      credentials: "include"
    })

    const json = (await fetchResp.json()).data

    if (!fetchResp.ok) throw new Error(await fetchResp.json())

    return { ok: true, notifications: json }
  } catch (err) {
    if (err instanceof Error) {
      return { ok: false, error: err, message: err.message }
    } else {
      return { ok: false, error: err, message: "Unknown Error" }
    }
  }
}

type CreateNotification = ReturnEntryType<"createdNotif", CustomNotification>
export async function createNotification(source_id: string, target_id: string, message: string, group_id: string = "", customFetch: Fetch = fetch): Promise<CreateNotification> {
  try {
    const fetchResp = await customFetch(`${PUBLIC_LOCAL_PATH}/api/v1/auth/notification/create`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Request-Method": "POST",
      },
      credentials: "include",
      body: JSON.stringify({ source_id, target_id, message, parent_id: group_id })
    })

    const json = (await fetchResp.json()).data

    if (!fetchResp.ok) throw new Error(await fetchResp.json())

    return { ok: true, createdNotif: json }
  } catch (err) {
    if (err instanceof Error) {
      return { ok: false, error: err, message: err.message }
    } else {
      return { ok: false, error: err, message: "Unknown Error" }
    }
  }
}

export async function deleteNotification(notif_id: string, customFetch: Fetch = fetch): Promise<CreateNotification> {
  try {
    const fetchResp = await customFetch(`${PUBLIC_LOCAL_PATH}/api/v1/auth/notification/delete/${notif_id}`, {
      method: "DELETE",
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Request-Method": "DELETE",
      },
      credentials: "include",
    })

    const json = (await fetchResp.json()).data

    if (!fetchResp.ok) throw new Error(await fetchResp.json())

    return { ok: true, createdNotif: json }
  } catch (err) {
    if (err instanceof Error) {
      return { ok: false, error: err, message: err.message }
    } else {
      return { ok: false, error: err, message: "Unknown Error" }
    }
  }
}
