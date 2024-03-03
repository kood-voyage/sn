import type { PageServerLoad, Actions } from './$types';
import { fail, redirect } from '@sveltejs/kit';
import { superValidate } from 'sveltekit-superforms';

import { signUpSchema } from '../schema';


import { User,UserStyle} from '$lib/types/user';

import { zod} from 'sveltekit-superforms/adapters';

import { createUser } from '$lib/server/db';



export const load: PageServerLoad = async () => {
	const form = await superValidate(zod(signUpSchema));
	return { form };
};



export const actions: Actions = {
	signup: async (event) => {
		const form = await superValidate(event, zod(signUpSchema));

		

		const userType: UserStyle = {
			username: form.data.username,
			email: form.data.email,
			dateOfBirth: form.data.dateOfBirth,
			firstName: form.data.firstName,
			lastName: form.data.lastName,
			password: form.data.password,
			repeatPassword:form.data.repeatPassword
		}


		const user = new User(userType)

		const result =  createUser(user)


		if(result.ok){
			redirect(300,"/signin")
		}






		if (!form.valid) {
			return fail(400, {
				
			});
		}




		redirect(300, "/signin")


	}
};
