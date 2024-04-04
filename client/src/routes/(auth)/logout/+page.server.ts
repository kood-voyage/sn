import { deleteTokens } from "$lib/server/jwt-handle"
import { redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";

export const prerender = false;


export const load: PageServerLoad = async (event) => {


    deleteTokens(event)

    redirect(303, "/signin")


};
