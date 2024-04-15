import { getGroups } from "$lib/server/api/group-requests";
import { mainGetGroup } from "$lib/server/db/group"
import { postSchema } from "$lib/types/post-schema";
import type { PageServerLoad } from "./$types"
import { superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';


export const load: PageServerLoad = async (event) => {
  const data = (await getGroups(event))
  if (!data.ok) {
    console.error(data.message)
    return {
      groups: { ...data }
    }
  }

  return { groups: { ...data }}

}