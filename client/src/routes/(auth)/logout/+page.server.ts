
import { redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
export const prerender = false;

export const load: PageServerLoad = async () => {
    redirect(303, "/signin")
};
