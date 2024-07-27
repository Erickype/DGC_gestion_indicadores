import { redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import { superValidate } from "sveltekit-superforms";
import { zod } from "sveltekit-superforms/adapters";
import { updateFacultySchema } from "./schema";

export const load: PageServerLoad = async ({locals}) => {
    if(!locals.user){
        throw redirect(302, "/")
    }

    return{
        updateFacultyForm: await superValidate(zod(updateFacultySchema))
    }
};