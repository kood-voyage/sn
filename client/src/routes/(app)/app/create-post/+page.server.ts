import type { PageServerLoad, Actions, } from './$types';
import { superValidate } from 'sveltekit-superforms';
import { postSchema } from "../../../../lib/types/post-schema"

import { zod } from 'sveltekit-superforms/adapters';
import { getUserIdFromCookie } from '$lib/server/jwt-handle';


import { v4 as uuidv4 } from 'uuid';
import { saveToS3 } from '$lib/server/images/upload';
import { LOCAL_PATH, S3_BUCKET, WEBSITE_PATH } from '$env/static/private';
import { redirect } from '@sveltejs/kit';




export const load: PageServerLoad = async () => {
    const form = await superValidate(zod(postSchema));

    return { form };
};


export type Privacy = {
    privacy: 'private' | 'public' | 'selected';
}


export type Post = {
    id: string;
    user_id: string;
    title: string;
    content: string;
    image_path?: string[];
    community_id: string;
    privacy: Privacy
}

export const actions: Actions = {
    postSubmit: async (event) => {
        // Validate the form data
        // const form = await superValidate(event, zod(postSchema));


        const { user_id } = getUserIdFromCookie(event)
        const formData = await event.request.formData();


        // Extracting other form fields
        const post_id = uuidv4()
        const title = formData.get('title') as string;
        const content = formData.get('content') as string;
        const privacy = formData.get('privacy') as Privacy


        const imagesURL: string[] = []

        // Extracting uploaded images
        const images = formData.getAll('images') as File[];


        for (const [i, image] of images.entries()) {
            const resp = await saveToS3(("post" + (i + 1)), post_id, image, "post")
            imagesURL.push(S3_BUCKET + resp)
        }


        const json: Post = {
            id: post_id,
            user_id: user_id,
            title: title,
            content: content,
            privacy: privacy,
            community_id: "",
            image_path: imagesURL

        }
        console.log(json)

        try {
            const response = await fetch(`${LOCAL_PATH}/api/v1/auth/posts/create`, {
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
        } catch (err) {
            if (err instanceof Error) {
                return { ok: false, error: err, message: err.message }
            } else {
                return { ok: false, error: err, message: "Unknown Error" }
            }
        }
        redirect(304, `/app/u`)
    }
};