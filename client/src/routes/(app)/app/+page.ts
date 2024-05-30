import { getFeed } from '$lib/client/api/post-requests';
import type { PageLoad } from './$types';

export const ssr = false
export const load: PageLoad = async () => {
  const feedResp = await getFeed()
  return { feed: feedResp };
};
