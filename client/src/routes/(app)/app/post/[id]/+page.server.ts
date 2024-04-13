import { LOCAL_PATH } from "$env/static/private"
import type { PageServerLoad } from "./$types"


export interface Post {
    id:           string;
    user_id:      string;
    title:        string;
    content:      string;
    image_path:   string[];
    community_id: string;
    created_at:   Date;
    privacy:      string;
}



export const load: PageServerLoad = async (event) => {
	const post_id = event.params.id
	const response = await fetch(`${LOCAL_PATH}/api/v1/auth/posts/${post_id}`,{
		method: 'GET',
		headers: {
            'Authorization': `Bearer ${event.cookies.get('at')}`}
			
	})

	const postJson = await response.json()


	const post: Post = postJson.data

  return {post}

}


