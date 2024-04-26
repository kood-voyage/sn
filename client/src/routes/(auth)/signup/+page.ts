import type { PageLoad } from './$types';
import { superValidate } from 'sveltekit-superforms';
import { signUpSchema } from '../schema';
import { zod } from 'sveltekit-superforms/adapters';


export const load: PageLoad = async () => {
	const form = await superValidate(zod(signUpSchema));
	return { form };
};



