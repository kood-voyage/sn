import type { PageServerLoad, Actions, } from './$types';
import { superValidate } from 'sveltekit-superforms';
import { v4 as uuidv4 } from 'uuid';
import { zod } from 'sveltekit-superforms/adapters';
import { groupSchema } from '$lib/types/group-schema';
import { saveToS3 } from '$lib/server/images/upload';
import { LOCAL_PATH, S3_BUCKET } from '$env/static/private';
import { redirect } from '@sveltejs/kit';



// interface post {

// 	user_id: string,
// 	post_id: string,
// 	title: string,
// 	content: string,
// 	privacy: string,
// 	images: string[],

// }

// import { fail } from '@sveltejs/kit';
// import { createTokens } from '$lib/server/jwt-handle';


export const load: PageServerLoad = async ({ request }) => {
	const form = await superValidate(request, zod(groupSchema));
	// Just returns { form } with the message (and status code 200).
	return { form };
}


export const actions: Actions = {
	default: async (event) => {
		const formData = await event.request.formData()

		const file = formData.get("image") as File

		console.log(`Sent File size ${file.size / 1024 / 1024} MB`)
		// console.log([...formData.entries()])

		const topic = "group"
		const id = uuidv4()
		const type = "cover"

		const key = await saveToS3(type, id, file, topic)
		const img_path = S3_BUCKET + key

		// console.log(formData)
		const group = {
			id,
			name: formData.get("title"),
			description: formData.get("content"),
			image_path: [img_path],
			privacy: formData.get("privacy"),
		}
		console.log("GROUP >>>", JSON.stringify(group))
		
		try {
			const response = await fetch(`${LOCAL_PATH}/api/v1/auth/group/create`, {
				method: 'POST',
				headers: {
					'Authorization': `Bearer ${event.cookies.get('at')}`,
					'Content-Type': 'application/json' // Specify JSON content type
				},
				body: JSON.stringify(group) // Convert the JSON object to a string
			});
			
			console.log("didnt throw error by itself", response.status)
			console.log(await response.json())
			// if (!response.ok) {
				// 	throw new Error('Failed to create group'); // Throw an error if response is not OK
				// }
				
				// Handle successful response


			//add creator to the group also
			const resp = await fetch(`${LOCAL_PATH}/api/v1/auth/group/join/${group.name}`, {
				method: 'GET',
				headers: {
					'Authorization': `Bearer ${event.cookies.get('at')}`,
				},
			})

		} catch (err) {
			if (err instanceof Error) {
				console.log("ERROR >>>", err.name)
				//  { ok: false, error: err, message: err.message }
			} else {
				//  { ok: false, error: err, message: "Unknown Error" }
			}
		}
		console.log("hello")
		redirect(303, `g/${group.name}`)

	},

};