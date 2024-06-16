import { superValidate } from "sveltekit-superforms";
import type { PageServerLoad } from "./$types";
import { redirect, type Actions } from "@sveltejs/kit";
import { zod } from "sveltekit-superforms/adapters";
import { addPersonSchema } from "./schema";


export const load: PageServerLoad = async ({ locals }) => {
    if (!locals.user) {
        throw redirect(302, "/")
    }

    return {
        addPersonForm: await superValidate(zod(addPersonSchema))
    }
};
