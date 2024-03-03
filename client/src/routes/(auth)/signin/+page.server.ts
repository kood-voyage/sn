import type { PageServerLoad, Actions } from './$types';
import { fail, redirect } from '@sveltejs/kit';
import { superValidate } from 'sveltekit-superforms';

import { signInSchema } from '../schema';

import { zod} from 'sveltekit-superforms/adapters';
import { checkUserExists } from '$lib/server/db';

export const load: PageServerLoad = async () => {
	const form = await superValidate(zod(signInSchema));

	return { form };
};

export const actions: Actions = {
	signin: async (event) => {
    const login = "http://localhost:8080/api/v1/auth/login"
		const form = await superValidate(event, zod(signInSchema));

		console.log("DOES THE USER EXIST >>> ", checkUserExists(form.data.login, form.data.password))

		if (!event.locals.user) redirect(302, "/signin")
		console.log('HERE');
		console.log(form.data);
		console.log(login);

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
