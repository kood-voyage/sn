

import { getUser } from "$lib/server/db/profile"
import { getUserIdFromCookie } from "$lib/server/jwt-handle"
import type { LayoutServerLoad } from "./$types"









export const load: LayoutServerLoad = async (event) => {
  const idResp = getUserIdFromCookie(event)
  if (!idResp.ok) {
    return
  }

  const user_id = idResp.user_id as string
  const data = getUser(user_id)

  if (data.error) {
    return {}
  }
  console.log("LAYOUT DATA >>>", data)






}