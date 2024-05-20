import { getFeed } from '$lib/client/api/post-requests';
import { GetAllUsers } from '$lib/client/api/user-requests';
import type { UserType } from '$lib/types/user';
import type { PageLoad } from './$types';

export const ssr = false

export const load: PageLoad = async ({ fetch }) => {

  // const resp = await GetAllUsers(fetch)
  // if (!resp.ok) {
  //   return { allUsers: [] }
  // }




  const feedResp = await getFeed()


  console.group(feedResp)

  



  return { feed: feedResp };



  
};
