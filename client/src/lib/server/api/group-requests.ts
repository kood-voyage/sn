import { LOCAL_PATH } from "$env/static/private";
import type { RequestEvent } from "@sveltejs/kit";


export type GroupJson = {
  id: string,
  creator_id: string,
  name: string,
  description: string,
  image_path: Array<string>,
  privacy: string,
  members: null | Array<string>
}

export async function getGroup(event: RequestEvent, group_name: string) {

  try {

    const fetchResp = await fetch(`${LOCAL_PATH}/api/v1/auth/group/${group_name}`, {
      headers: {
        "Authorization": `Bearer ${event.cookies.get('at')?.valueOf()}`
      }
    })
    const json = (await fetchResp.json()).data as GroupJson


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