import { LOCAL_PATH } from '$env/static/private';
import type { Actions } from './$types';

import { getProfile } from "$lib/server/db/profile"
import type { PageServerLoad } from './$types';


export const load: PageServerLoad = async (event) => {

  const data = (await getProfile(event, event.params.username))
  if (data.error) {
    console.error(data)
    return {
      username: "Undefined"
    }
  }
  
  return data

}


export const actions: Actions = {
	follow: async (event) => {
  const data = await event.request.formData()
  const target_id = data.get("target_id")

    try {
    await fetch(`${LOCAL_PATH}/api/v1/auth/follow/${target_id}`, {
      headers: {
        "Authorization": `Bearer ${event.cookies.get('at')}`
      }
    })

  } catch (err) {
    if (err instanceof Error) {
      return { ok: false, error: err, message: err.message }
    } else {
      return { ok: false, error: err, message: "Unknown Error" }
    }
  }

  },



  unfollow: async (event) => {
  const data = await event.request.formData()
  const target_id = data.get("target_id")

    try {
    await fetch(`${LOCAL_PATH}/api/v1/auth/unfollow/${target_id}`, {
      headers: {
        "Authorization": `Bearer ${event.cookies.get('at')}`
      }
    })
  } catch (err) {
    if (err instanceof Error) {
      return { ok: false, error: err, message: err.message }
    } else {
      return { ok: false, error: err, message: "Unknown Error" }
    }
  }

  }





}

