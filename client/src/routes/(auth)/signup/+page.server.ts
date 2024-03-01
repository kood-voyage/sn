import type { PageServerLoad, Actions } from './$types';
import { fail, redirect } from '@sveltejs/kit';
import { superValidate } from 'sveltekit-superforms';

import { signUpSchema } from '../schema';

import { zod} from 'sveltekit-superforms/adapters';


// interface RegisterForm {
// 	username:string,
// 	email:string,
// 	dateOfBirth: string,
// 	password: string,
// 	repeatPassword: string,
// 	firstName:string,
// 	lastName:string
// }


export const load: PageServerLoad = async () => {
	const form = await superValidate(zod(signUpSchema));

	return { form };
};



export const actions: Actions = {
	signup: async (event) => {
		const form = await superValidate(event, zod(signUpSchema));

		const {username,email,dateOfBirth,password,repeatPassword,firstName,lastName} = form.data






		if (!form.valid) {
			return fail(400, {
				
			});
		}




		redirect(300, "/signin")


	}
};
