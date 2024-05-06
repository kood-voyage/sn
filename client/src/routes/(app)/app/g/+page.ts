import { GetAllGroups, type GroupJson } from "$lib/client/api/group-requests";
import { postSchema } from "$lib/types/post-schema";
import type { PageLoad } from "../$types";
import { superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';

export const ssr = false

export const load: PageLoad = async ({ fetch }) => {

  const data = (await GetAllGroups(fetch))
  if (!data.ok) {
    console.error(data.message)
    return {
      groups: undefined
    }
  }


  return { groups: data.allGroups as GroupJson[] }

}