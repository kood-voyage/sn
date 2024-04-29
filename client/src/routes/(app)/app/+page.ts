import { GetAllUsers } from '$lib/client/api/user-requests';
import type { UserType } from '$lib/types/user';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ fetch }) => {

  // console.log()
  const resp = await GetAllUsers(fetch)
  if (!resp.ok) {
    return { allUsers: [] }
  }

  return { allUsers: resp.allUsers as UserType[] };
};
