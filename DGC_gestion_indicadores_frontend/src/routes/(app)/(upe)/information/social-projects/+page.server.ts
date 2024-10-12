import type { Actions, PageServerLoad } from "./$types";
import { mainDashboarRoute } from "$lib/api/util/paths";
import { redirect } from "@sveltejs/kit";

import { filterAcademicPeriodsAuxSchema } from "$lib/utils";
import { addSocialProjectListSchema } from "./schema";
import { superValidate } from "sveltekit-superforms";
import { zod } from "sveltekit-superforms/adapters";

import { LoadAcademicPeriodsWithComboMessages } from "$lib/api/controller/view/academicPeriod";
import { LoadCareersWithComboMessages } from "$lib/api/controller/api/career";

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
        careersData: await LoadCareersWithComboMessages(token!),
        filterAcademicPeriodsAuxForm: await superValidate(zod(filterAcademicPeriodsAuxSchema)),
        addSocialProjectListForm: await superValidate(zod(addSocialProjectListSchema))
    }
};