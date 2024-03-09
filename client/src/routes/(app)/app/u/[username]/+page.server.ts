import { mainGetProfile } from "$lib/server/db/profile";
import type { PageServerLoad } from "./$types";


export const load: PageServerLoad = async (event) => {
  const username = event.url.pathname.split("/")[event.url.pathname.split("/").length - 1]
  console.log(username)
  mainGetProfile(username)



};