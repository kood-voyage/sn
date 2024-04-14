import type { RequestEvent } from "@sveltejs/kit";
import { getGroup, type GroupJson } from "$lib/server/api/group-requests";
import { type ReturnType } from "$lib/types/requests";

type GetGroupType = ReturnType<GroupJson>

export async function mainGetGroup(event: RequestEvent, group_name: string): Promise<GetGroupType> {
  const originalGroupName = group_name.replace(/_/g, ' ');

  const respGroup = await getGroup(event, originalGroupName);
  if (!respGroup.ok) {
    return { ...respGroup };
  }

  return { ...respGroup };
}