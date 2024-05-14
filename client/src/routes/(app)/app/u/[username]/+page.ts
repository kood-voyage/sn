export const ssr = false 

import { getUserPosts } from '$lib/client/api/post-requests';
import { getUserById, getUserFollowers, getUserFollowing } from '$lib/client/api/user-requests';
import type { PageLoad } from './$types';

export const load: PageLoad = async({ params }) => {

	const user = await getUserById(params.username)
	const posts = await getUserPosts(user.data.id)
	const followers = await getUserFollowers(user.data.id)
	const following = await getUserFollowing(user.data.id)



	return {params,posts,user,followers,following};
};