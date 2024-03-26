import type { PageServerLoad, Actions, } from './$types';
import { superValidate } from 'sveltekit-superforms';
import {postSchema} from "../post-schema"

import { zod } from 'sveltekit-superforms/adapters';
import { getUserIdFromCookie } from '$lib/server/jwt-handle';


import { v4 as uuidv4 } from 'uuid';



interface post {

	user_id: string,
	post_id: string,
	title: string,
	content: string,
	privacy: string,
	images: string[],

}

// import { fail } from '@sveltejs/kit';
// import { createTokens } from '$lib/server/jwt-handle';


export const load: PageServerLoad = async () => {
	const form = await superValidate(zod(postSchema));

	return { form };
};


export const actions: Actions = {
    postSubmit: async (event) => {
        // Validate the form data
        // const form = await superValidate(event, zod(postSchema));

		




		const user_id = getUserIdFromCookie(event)        
        const formData = await event.request.formData();


		

        // Extracting other form fields
		const post_id = uuidv4()
        const title = formData.get('title') as string;
        const content = formData.get('content') as string;

		const privacy = formData.get('privacy') as string

        // Extracting uploaded images
        const images = formData.getAll('images') as File[];


		// for(const image of images){


			

		// }





		console.log(user_id.user_id)
		console.log(post_id)
		console.log(title)
		console.log(content)
		console.log(privacy)
		console.log(images)


        // You can now use the extracted form data and uploaded images to perform further actions (e.g., database operations)

        // Return a response as needed
        return {
            status: 200,
            body: { message: 'Form submitted successfully' }
        };
    }
};