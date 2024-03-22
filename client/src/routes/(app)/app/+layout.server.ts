import { getUser } from "$lib/server/db/profile"
import { getUserIdFromCookie } from "$lib/server/jwt-handle"
import type { RequestEvent } from "@sveltejs/kit"
import type { RouteParams } from "../../(auth)/signin/$types"
import type { LayoutServerLoad } from "./$types"










export const load: LayoutServerLoad = async (event: RequestEvent<RouteParams, "/(auth)/signin"> | RequestEvent<Partial<Record<string, string>>, string | null>) => {
  const idResp = getUserIdFromCookie(event)
  if (!idResp.ok) {
    return
  }


  const user_id = idResp.user_id as string
  const data = getUser(user_id)

  if (data.error) {
    return {}
  }
  return data?.data


}

