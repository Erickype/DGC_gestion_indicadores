import type { Actions, PageServerLoad } from "./$types";
import { mainDashboarRoute } from "$lib/api/util/paths";
import { redirect } from "@sveltejs/kit";

import { filterAcademicPeriodsAuxSchema, generateFormMessageFromHttpResponse, generateFormMessageFromInvalidForm } from "$lib/utils";
import { addSocialProjectListSchema, updateSocialProjectListSchema } from "./schema";
import { superValidate } from "sveltekit-superforms";
import { zod } from "sveltekit-superforms/adapters";

import { PostSocialProjectList, PutSocialProjectList } from "$lib/api/controller/api/indicatorsInformation/socialProjectLists";
import type { SocialProjectList } from "$lib/api/model/api/indicatorsInformation/socialProjectLists";
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
        addSocialProjectListForm: await superValidate(zod(addSocialProjectListSchema)),
        updateSocialProjectListForm: await superValidate(zod(updateSocialProjectListSchema))
    }
};

export const actions: Actions = {
    postSocialProjectList: async (event) => {
        const form = await superValidate(event, zod(addSocialProjectListSchema))

        if (!form.valid) {
            return generateFormMessageFromInvalidForm(form)
        }

        const token = event.cookies.get("AuthorizationToken")
        const data = form.data
        const socialProjectList: SocialProjectList = {
            academic_period_id: data.academic_period_id,
            career_id: data.career_id,
            name: data.name
        }

        const response = await PostSocialProjectList(token!, socialProjectList)

        return generateFormMessageFromHttpResponse(form, response)
    },

    putSocialProjectList: async (event) => {
        const form = await superValidate(event, zod(updateSocialProjectListSchema))

        if (!form.valid) {
            return generateFormMessageFromInvalidForm(form)
        }

        const token = event.cookies.get("AuthorizationToken")
        const data = form.data
        const socialProjectList: SocialProjectList = {
            ID: data.ID,
            academic_period_id: data.academic_period_id,
            career_id: data.career_id,
            name: data.name
        }

        const response = await PutSocialProjectList(token!, socialProjectList)

        return generateFormMessageFromHttpResponse(form, response)
    }
};