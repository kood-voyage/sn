import type { PageServerLoad, Actions,  } from './$types';
import { type Cookies} from '@sveltejs/kit';
import { superValidate } from 'sveltekit-superforms';

import { signInSchema } from '../schema';

import { zod} from 'sveltekit-superforms/adapters';
import { checkUserExists, getUserId } from '$lib/server/db';

import jwt from 'jsonwebtoken'
import { JWT_KEY } from '$env/static/private';

import { v4 as uuidv4 } from 'uuid';


export const load: PageServerLoad = async () => {
	const form = await superValidate(zod(signInSchema));

	return { form };
};


export const actions: Actions = {
	signin: async (event) => {
    // const login = "http://localhost:8080/api/v1/auth/login"



		const form = await superValidate(event,zod(signInSchema));

		console.log("DOES THE USER EXIST >>> ", checkUserExists(form.data.login, form.data.password))

		if(form.data.login && form.data.password){
			const user_id  = getUserId(form.data.login)


		const access_token_id:string = uuidv4()


		const access_token = jwt.sign({
				exp: Math.floor(Date.now() / 1000) + (60 * 15),
				user_id,
				id: access_token_id
		}, JWT_KEY,{ algorithm: 'HS256' })


		const refresh_token = jwt.sign({
				exp: Math.floor(Date.now() / 1000) + (60 * 60 *24 *7),
				access_token_id: access_token_id
		}, JWT_KEY,{ algorithm: 'HS256' })


		event.cookies.set("at"  ,access_token,{path:"/"})
		event.cookies.set("rt", refresh_token,{path:"/"})

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
