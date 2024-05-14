import { superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';
import { groupSchema } from '$lib/types/group-schema';
import type { PageLoad } from '../$types';





export const load: PageLoad = async ({ request }) => {
	const form = await superValidate(request, zod(groupSchema));
	return { form };
}

