import type { PageLoad } from './$types';
import { superValidate } from 'sveltekit-superforms';
import { signInSchema } from '../schema';
import { zod } from 'sveltekit-superforms/adapters';

export const load: PageLoad = async () => {
	const form = await superValidate(zod(signInSchema));

	return { form };
};
