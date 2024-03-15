import {deleteTokens} from "$lib/server/jwt-handle"
import { redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";




export const load: PageServerLoad = async (event) => {
    
    console.log("GERE")
    
    deleteTokens(event)

    redirect(303,"/signin")


};
