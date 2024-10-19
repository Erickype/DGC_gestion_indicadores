import type { PageServerLoad } from "./$types";
import { redirect } from "@sveltejs/kit";

import { superValidate } from "sveltekit-superforms";
import { zod } from "sveltekit-superforms/adapters";
import { updateUserSchema } from "./schema";

export const load: PageServerLoad = async ({ locals }) => {
    if(!locals.user){
        throw redirect(302, "/login")
    }

    if(locals.user.role !== 1){
        throw redirect(302, "/")
    }

    return {
        updateUserForm: await superValidate(zod(updateUserSchema)),
    }
};