import type { PageServerLoad, Actions } from './$types';
import { fail, redirect } from '@sveltejs/kit';
import { superValidate } from 'sveltekit-superforms';

import { signUpSchema } from '../schema';


import { User} from '$lib/types/user';

import { zod} from 'sveltekit-superforms/adapters';

import { createUser } from '$lib/server/db';



export const load: PageServerLoad = async () => {
	const form = await superValidate(zod(signUpSchema));
	return { form };
};



export const actions: Actions = {
	signup: async (event) => {
		const form = await superValidate(event, zod(signUpSchema));
		const user = new User(form.data)
		const result =  createUser(user)
		if(result.ok){
			redirect(300,"/signin")
		}
		if (!form.valid) {
			return fail(400, {
				
			});
		}
	}
};
