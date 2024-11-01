import { mainDashboarRoute } from "$lib/api/util/paths";
import type { PageServerLoad, Actions } from "./$types";
import { redirect } from "@sveltejs/kit"

import { generateFormMessageFromHttpResponse, generateFormMessageFromInvalidForm } from "$lib/utils";
/* import { addGradeRateListSchema, updateGradeRateListSchema } from "./schema"; */
import { superValidate } from "sveltekit-superforms";
import { zod } from "sveltekit-superforms/adapters";

export const load: PageServerLoad = async ({ cookies, locals }) => {
    const token = cookies.get("AuthorizationToken")
    if (!locals.user) {
        throw redirect(302, "/login")
    }
    if (locals.user.role === 3) {
        throw redirect(302, mainDashboarRoute)
    }

    return {
    }
};

export const actions: Actions = {
};