import type { RequestEvent } from "@sveltejs/kit";
import { getGroup } from "../api/group-requests";

export async function mainGetGroup(event: RequestEvent, group_name: string) {
  const originalGroupName = group_name.replace(/_/g, ' ');

  const respGroup = await getGroup(event, originalGroupName);
  if (!respGroup.ok) {
    return { ...respGroup };
  }

  return { ...respGroup };
}