import { mainGetGroup } from "$lib/server/db/group"
import { postSchema } from "$lib/types/post-schema";
import type { PageServerLoad } from "./$types"
import { superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';


export const load: PageServerLoad = async (event) => {
  const form = await superValidate(zod(postSchema));


  console.log(event.params.name)
  console.log((await event.parent()))
  const data = (await mainGetGroup(event, event.params.name))
  if (!data.ok) {
    console.error(data.message)
    return {
      group: { ...data }, form
    }
  }
  // console.log(typeof data)
  // console.log(typeof data.data)
  return { group: { ...data }, form }

}