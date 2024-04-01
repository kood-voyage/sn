import type { RequestEvent } from "@sveltejs/kit";
import { getGroup } from "../api/group-requests";

export async function mainGetGroup(event: RequestEvent, group_name: string) {

  const respGroup = await getGroup(event, group_name)
  if (!respGroup.ok) {
    return { ...respGroup }
  }

  console.log(respGroup.data)
  return { ...respGroup }
}