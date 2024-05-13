import { commentSchema } from "../../comment-schema";
import { superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';
import { PUBLIC_LOCAL_PATH } from '$env/static/public';
import { getUserById } from "$lib/client/api/user-requests";

import type { UserType } from "$lib/types/user";


export interface Post {
  id: string;
  user_id: string;
  title: string;
  content: string;
  image_path: string[];
  community_id: string;
  created_at: Date;
  privacy: string;
  user_information: UserType
}

export interface Comment {
  id: string
  user_id: string
  post_id: string
  parent_id: string
  content: string
  image_path?: string[]
  created_at: string
  count?: string

}


export const load: PageLoad = async ({params}) => {
  const post_id = params.id



  const postResponse = await fetch(`${PUBLIC_LOCAL_PATH}/api/v1/auth/posts/${post_id}`, {
			method: 'GET',
			headers: {
				'Content-Type': 'application/json',
				'Access-Control-Request-Method': 'GET'
			},
			credentials: 'include'

  })

  const commentsResponse = await fetch(`${PUBLIC_LOCAL_PATH}/api/v1/auth/comment/${post_id}`, {
      method: 'GET',
			headers: {
				'Content-Type': 'application/json',
				'Access-Control-Request-Method': 'GET'
			},
			credentials: 'include'
  })



  const form = await superValidate(zod(commentSchema));
  const postJson = await postResponse.json()
  const post: Post = postJson.data



 const   comments = await commentsResponse.json()

  return { post, form, comments }

}
