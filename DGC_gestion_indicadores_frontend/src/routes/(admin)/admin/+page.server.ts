import { redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ locals }) => {
    if(!locals.user){
        throw redirect(302, "/login")
    }

    if(locals.user.role !== 1){
        throw redirect(302, "/")
    }
};