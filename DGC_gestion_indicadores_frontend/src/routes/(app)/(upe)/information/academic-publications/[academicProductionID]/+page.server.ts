import { addAcademicProductionListsAuthorSchema } from "./schema";
import { superValidate } from "sveltekit-superforms";
import { zod } from "sveltekit-superforms/adapters";

import { error, redirect, type Actions } from "@sveltejs/kit";
import { mainDashboarRoute } from "$lib/api/util/paths";
import type { PageServerLoad } from "./$types";

import type { PostAcademicProductionListsAuthorCareersRequest } from "$lib/api/model/api/indicatorsInformation/academicProductionListsAuthor";
import { PostAcademicProductionListsAuthorCareers } from "$lib/api/controller/api/indicatorsInformation/academicProductionListsAuthor";
import { generateFormMessageFromHttpResponse, generateFormMessageFromInvalidForm } from "$lib/utils";
import { LoadCareersWithComboMessages } from "$lib/api/controller/api/career";

export const load: PageServerLoad = async ({ locals, cookies, params }) => {
    const token = cookies.get("AuthorizationToken")

    if (!locals.user) {
        throw redirect(302, "/login")
    }

    if (locals.user.role === 3) {
        throw redirect(302, mainDashboarRoute)
    }

    const academicProductionID = parseInt(params.academicProductionID)
    if (academicProductionID > 0) {
        return {
            academicProductionID: academicProductionID,
            addAcademicProductionListsAuthorForm: await superValidate(zod(addAcademicProductionListsAuthorSchema)),
            careersData: await LoadCareersWithComboMessages(token!),
        }
    }

    error(404, 'Not found');
};

export const actions: Actions = {
    postAcademicProductionListsAuthorCareers: async (event) => {
        const form = await superValidate(event, zod(addAcademicProductionListsAuthorSchema))

        if (!form.valid) {
            return generateFormMessageFromInvalidForm(form)
        }

        const token = event.cookies.get("AuthorizationToken")
        const data = form.data
        const academicProductionListsAuthorCareersRequest: PostAcademicProductionListsAuthorCareersRequest = {
            academic_production_list_id: data.academicProductionID,
            author_id: data.authorID,
            careers: data.careers
        }

        const response = await PostAcademicProductionListsAuthorCareers(token!, academicProductionListsAuthorCareersRequest)

        return generateFormMessageFromHttpResponse(form, response)
    },
};