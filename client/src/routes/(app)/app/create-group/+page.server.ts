import type { PageServerLoad, Actions, } from './$types';
import { superValidate } from 'sveltekit-superforms';


import { zod } from 'sveltekit-superforms/adapters';
import { groupSchema } from '../group-schema';



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
		// console.log(file)
		console.log(`Sent File size ${file.size / 1024 / 1024} MB`)
		// const formData = await event.request.formData();
		// Process formData as necessary...

		// Assuming 'form' is a representation of your form data or state you wish to return
		console.log([...formData.entries()])

		// console.log(formData)
		const group = {
			name: formData.get("title"),
			description: formData.get("content"),
			privacy: formData.get("privacy"),
		}
		console.log(group)


		// const file = formData.get("image") as File
		// console.log(`Sent File size ${file.size / 1024 / 1024} MB`)





		// redirect(303, "/app/create-group")

	},

};