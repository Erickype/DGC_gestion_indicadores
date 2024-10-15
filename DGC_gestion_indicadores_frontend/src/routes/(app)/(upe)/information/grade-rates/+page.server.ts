import { mainDashboarRoute } from "$lib/api/util/paths";
import type { PageServerLoad } from "./$types";
import { redirect } from "@sveltejs/kit";

import { filterAcademicPeriodsAuxSchema } from "$lib/utils";
import { superValidate } from "sveltekit-superforms";
import { zod } from "sveltekit-superforms/adapters";

import { LoadAcademicPeriodsWithComboMessages } from "$lib/api/controller/view/academicPeriod";
import { LoadCareersWithComboMessages } from "$lib/api/controller/api/career";
import { addGradeRateListSchema, updateGradeRateListSchema } from "./schema";

export const load: PageServerLoad = async ({ cookies, locals }) => {
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
        addGradeRateListForm: await superValidate(zod(addGradeRateListSchema)),
        updateGradeRateListForm: await superValidate(zod(updateGradeRateListSchema))
    }
};