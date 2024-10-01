import type { Actions, PageServerLoad } from "./$types";
import { mainDashboarRoute } from "$lib/api/util/paths";
import { redirect } from "@sveltejs/kit";

import { superValidate } from "sveltekit-superforms";
import { zod } from "sveltekit-superforms/adapters";

import { LoadAcademicPeriodsWithComboMessages } from "$lib/api/controller/view/academicPeriod";

export const load: PageServerLoad = async ({ locals, cookies }) => {
    const token = cookies.get("AuthorizationToken")

    if (!locals.user) {
        throw redirect(302, "/login")
    }

    if (locals.user.role === 3) {
        throw redirect(302, mainDashboarRoute)
    }
    return {
        academicPeriodsData: await LoadAcademicPeriodsWithComboMessages(),
    }
};

export const actions: Actions = {
};