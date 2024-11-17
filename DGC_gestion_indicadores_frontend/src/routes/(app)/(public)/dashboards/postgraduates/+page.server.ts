import { mainDashboarRoute } from "$lib/api/util/paths";
import type { PageServerLoad } from "./$types";
import { redirect } from "@sveltejs/kit"

import { superValidate } from "sveltekit-superforms";
import { zod } from "sveltekit-superforms/adapters";

import { filterCohortsAuxSchema } from "$lib/utils";

export const load: PageServerLoad = async ({ cookies, locals }) => {
    const token = cookies.get("AuthorizationToken")
    if (!locals.user) {
        throw redirect(302, "/login")
    }
    if (locals.user.role === 3) {
        throw redirect(302, mainDashboarRoute)
    }

    return {
        filterCohortsAuxForm: await superValidate(zod(filterCohortsAuxSchema)),
    }
};