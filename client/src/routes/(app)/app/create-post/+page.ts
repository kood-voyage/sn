// import type { PageServerLoad, Actions, } from './$types';
import { superValidate } from 'sveltekit-superforms';
import { postSchema } from "../../../../lib/types/post-schema"
import { zod } from 'sveltekit-superforms/adapters';
import type { PageLoad } from '../$types';


export const ssr = false

export const load: PageLoad = async () => {
    const form = await superValidate(zod(postSchema));
    return { form };
};



