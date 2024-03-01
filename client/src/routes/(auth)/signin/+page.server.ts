import type { PageServerLoad, Actions } from './$types';
import { fail, redirect } from '@sveltejs/kit';
import { superValidate } from 'sveltekit-superforms';

import { signInSchema } from '../schema';

import { zod} from 'sveltekit-superforms/adapters';

export const load: PageServerLoad = async () => {
	const form = await superValidate(zod(signInSchema));

	return { form };
};

export const actions: Actions = {
	signin: async (event) => {
		const form = await superValidate(event, zod(signInSchema));
		if (!event.locals.user) redirect(302, "/signin")
		console.log('HERE');
		console.log(form.data);

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
