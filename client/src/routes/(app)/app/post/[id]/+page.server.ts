import { LOCAL_PATH } from "$env/static/private"
import {getUser } from "$lib/server/db/profile";
import type { PageServerLoad, Actions, } from './$types';
import { commentSchema } from "../../comment-schema";
import { superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';
import { redirect } from "@sveltejs/kit";
import { getUserIdFromCookie } from "$lib/server/jwt-handle";

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

export interface Comment {
  id:string
  user_id:string
  post_id:string
  parent_id:string
  content:string
  image_path?: string[]
  created_at: string
  count?:string

}

export const load: PageServerLoad = async (event) => {
	const post_id = event.params.id
	const postResponse = await fetch(`${LOCAL_PATH}/api/v1/auth/posts/${post_id}`,{
		method: 'GET',
		headers: {
            'Authorization': `Bearer ${event.cookies.get('at')}`}
			
	})

  	const commentsResponse = await fetch(`${LOCAL_PATH}/api/v1/auth/comment/${post_id}`,{
		method: 'GET',
		headers: {
            'Authorization': `Bearer ${event.cookies.get('at')}`}
			
	})

  const form = await superValidate(zod(commentSchema));

	const postJson = await postResponse.json()
	const post: Post = postJson.data


  const postAuthorResponse = await getUser(post.user_id)
  const postAuthor = await postAuthorResponse.data

  const comments = await commentsResponse.json()
  // const comments = commentsJson.data


  console.log("HERE IS")
  console.log(comments)


  return {post,postAuthor,form,comments}

}
  

export const actions: Actions = {
commentSubmit: async (event) => {

		const {user_id}= getUserIdFromCookie(event)
    const formData = await event.request.formData();
    const post_id = formData.get('post_id') as string;
    const parent_id = formData.get('parent_id') as string
    const content = formData.get('content') as string;
    const user_name =formData.get('user_name') as string;
    const user_avatar =formData.get('user_avatar') as string;


		// const imagesURL: string[] = []
    // Extracting uploaded images
    // const images = formData.getAll('images') as File[];
		// for(const [i,image] of images.entries()){
		// 	const resp = await saveToS3(("post"+(i+1)),post_id,image,"post")
		// 	imagesURL.push(S3_BUCKET + resp)
		// }

		const json: Comment ={
			user_id,
      post_id,
      parent_id,
      content,
      user_name,
      user_avatar
		}

try {
    const response = await fetch(`${LOCAL_PATH}/api/v1/auth/comment/create`, {
        method: 'POST',
        headers: {
            'Authorization': `Bearer ${event.cookies.get('at')}`,
            'Content-Type': 'application/json' // Specify JSON content type
        },
        body: JSON.stringify(json) // Convert the JSON object to a string
    });

    if (!response.ok) {
        throw new Error('Failed to create post'); // Throw an error if response is not OK
    }

    // Handle successful response
}  catch (err) {
    if (err instanceof Error) {
      return { ok: false, error: err, message: err.message }
    } else {
      return { ok: false, error: err, message: "Unknown Error" }
    }
  }
		redirect(304, `/app/post/${post_id}`)
  }
};