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

//1. server folderis api folder api folderis grouprequests.ts --> fetch request there
//2. page.server.ts peab looma g sisse 
//3. returnin loadis data --> data tuleb valja page.sveltes --> naeb marco name filedes script +page.server.ts ja page.svelte
//4. 