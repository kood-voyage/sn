import type { PageServerLoad, Actions, } from './$types';
import { superValidate } from 'sveltekit-superforms';


import { apiCreateUser } from '$lib/server/api/user-requests';

import { signInSchema } from '../schema';

import { zod } from 'sveltekit-superforms/adapters';
import { checkSessionExists, checkUserExists, deleteSession } from '$lib/server/db';

import { fail } from '@sveltejs/kit';
import { createTokens } from '$lib/server/jwt-handle';


export const load: PageServerLoad = async () => {
	const form = await superValidate(zod(signInSchema));

	return { form };
};


export const actions: Actions = {
	signin: async (event) => {
		const form = await superValidate(event, zod(signInSchema));

		const respUser = checkUserExists(form.data.login, form.data.password)
		if (!respUser.ok) {
			console.error(respUser.error)
			return fail(400, {
				form,
				message: "Username/Email or Password not found"
			})
		} else if (!respUser.authorized) {
			console.error(respUser.error)
			return fail(400, {
				form,
				message: "Username/Email or Password incorrect"
			})
		}

		const user_id = respUser.id as string
		const resp = checkSessionExists(user_id)
		if (resp.ok) deleteSession(user_id)

		const respToken = createTokens(event, user_id)




		if (!respToken.ok) {
			console.error(respToken.error)
			respToken.error = undefined
			return fail(400, {
				form,
				...respToken
			})
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
