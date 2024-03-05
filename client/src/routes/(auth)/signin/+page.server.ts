import type { PageServerLoad, Actions,  } from './$types';
import { superValidate } from 'sveltekit-superforms';

import { signInSchema } from '../schema';

import { zod} from 'sveltekit-superforms/adapters';
import { checkUserExists } from '$lib/server/db';

import { fail } from '@sveltejs/kit';
import { createTokens } from '$lib/server/jwt-handle';


export const load: PageServerLoad = async () => {
	const form = await superValidate(zod(signInSchema));

	return { form };
};


export const actions: Actions = {
	signin: async (event) => {
    // const login = "http://localhost:8080/api/v1/auth/login"
		// console.log(event.url)
		//  await fetch("http://localhost:8080/api/v1/auth/user/create/123", 
		//  {
		// 	credentials: "include"
		//  }
		// )
		// console.log(resps)

		const form = await superValidate(event,zod(signInSchema));

		const resp = checkUserExists(form.data.login, form.data.password)
		if(resp.ok && resp.authorized){
			const user_id = resp.id as string
			createTokens(event, user_id)
		}


		if (!form.valid) {
			return fail(400, {
				form 
			});
		}
		return {
			form
		};
	}
};
