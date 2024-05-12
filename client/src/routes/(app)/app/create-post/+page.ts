// import type { PageServerLoad, Actions, } from './$types';
import { superValidate } from 'sveltekit-superforms';
import { postSchema } from "../../../../lib/types/post-schema"
import { zod } from 'sveltekit-superforms/adapters';


export const ssr = false

export const load: PageServerLoad = async () => {
    const form = await superValidate(zod(postSchema));
    return { form };
};



