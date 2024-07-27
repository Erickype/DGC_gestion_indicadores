import { redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import { superValidate } from "sveltekit-superforms";
import { zod } from "sveltekit-superforms/adapters";
import { addFacultySchema, updateFacultySchema } from "./schema";

export const load: PageServerLoad = async ({locals}) => {
    if(!locals.user){
        throw redirect(302, "/")
    }

    return{
        addFacultyForm: await superValidate(zod(addFacultySchema)),
        updateFacultyForm: await superValidate(zod(updateFacultySchema))
    }
};