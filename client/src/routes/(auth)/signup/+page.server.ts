import type { PageServerLoad, Actions } from './$types';
import { fail, redirect } from '@sveltejs/kit';
import { superValidate } from 'sveltekit-superforms';

import { signUpSchema } from '../schema';

// import { v4 as uuidv4 } from 'uuid';


import jwt from 'jsonwebtoken'


import { JWT_KEY, LOCAL_PATH} from '$env/static/private';

const access_token_id: string = uuidv4()


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


		const access_token = jwt.sign({
			exp: Math.floor(Date.now() / 1000) + (60 * 15), 
			user_id : user.id,
			access_token_id
		}, JWT_KEY, { algorithm: 'HS256' })


		await fetch(`${LOCAL_PATH}/api/v1/auth/user/create/public`, {
		headers: {
			"Authorization": `Bearer ${access_token}`
		}
		
		})
			redirect(300,"/signin")
		}
		if (!form.valid) {
			return fail(400, {
				
			});
		}
	}
};
function uuidv4(): string {
	throw new Error('Function not implemented.');
}

