import { mainDashboarRoute } from "$lib/api/util/paths";
import type { PageServerLoad, Actions } from "./$types";
import { redirect } from "@sveltejs/kit"

import { filterAcademicPeriodsAuxSchema, generateFormMessageFromHttpResponse, generateFormMessageFromInvalidForm } from "$lib/utils";
import { superValidate } from "sveltekit-superforms";
import { zod } from "sveltekit-superforms/adapters";

import { LoadAcademicPeriodsWithComboMessages } from "$lib/api/controller/view/academicPeriod";
import { LoadCareersWithComboMessages } from "$lib/api/controller/api/career";
import { addGradeRateListSchema, updateGradeRateListSchema } from "./schema";
import { PostGradeRateList } from "$lib/api/controller/api/indicatorsInformation/gradeRateLists";
import type { GradeRateList } from "$lib/api/model/api/indicatorsInformation/gradeRateLists";

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

export const actions: Actions = {
    postGradeRateList: async (event) => {
        const form = await superValidate(event, zod(addGradeRateListSchema))

        if (!form.valid) {
            return generateFormMessageFromInvalidForm(form)
        }

        const token = event.cookies.get("AuthorizationToken")
        const data = form.data
        const gradeRateList: GradeRateList = {
            academic_period_id: data.academic_period_id,
            career_id: data.career_id,
            count_graduated_students: data.count_graduated_students,
            count_court_students: data.count_court_students,
            count_admitted_matriculated_students: data.count_admitted_matriculated_students,
            count_admitted_students: data.count_admitted_students
        }

        const response = await PostGradeRateList(token!, gradeRateList)

        return generateFormMessageFromHttpResponse(form, response)
    }
};