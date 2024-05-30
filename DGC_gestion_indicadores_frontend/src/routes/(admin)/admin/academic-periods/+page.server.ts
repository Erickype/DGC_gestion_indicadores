import type { PageServerLoad } from "./$types";
import { redirect } from "@sveltejs/kit";
import { addAcademicPeriodSchema, updateAcademicPeriodSchema } from "./schema";
import { superValidate } from "sveltekit-superforms";
import { zod } from "sveltekit-superforms/adapters";

export const load: PageServerLoad = async ({ locals }) => {
    if (!locals.user) {
        throw redirect(302, "/")
    }
    return {
        addAcademiPeriodForm: await superValidate(zod(addAcademicPeriodSchema)),
        updateAcademicPeriodForm: await superValidate(zod(updateAcademicPeriodSchema))
    }
};