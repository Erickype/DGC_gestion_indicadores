import { mainDashboarRoute } from "$lib/api/util/paths";
import type { PageServerLoad, Actions } from "./$types";
import { redirect } from "@sveltejs/kit"

import { generateFormMessageFromHttpResponse, generateFormMessageFromInvalidForm } from "$lib/utils";
import { addArtisticProductionListSchema, updateArtisticProductionListSchema } from "./schema";
import { superValidate } from "sveltekit-superforms";
import { zod } from "sveltekit-superforms/adapters";

import { PostArtisticProductionList, UpdateArtisticProductionList } from "$lib/api/controller/api/indicatorsInformation/artisticProductionLists";
import type { ArtisticProductionList } from "$lib/api/model/api/indicatorsInformation/artisticProductionLists";

export const load: PageServerLoad = async ({ locals }) => {
    if (!locals.user) {
        throw redirect(302, "/login")
    }
    if (locals.user.role === 3) {
        throw redirect(302, mainDashboarRoute)
    }

    return {
        addResearchInnovationProjectListForm: await superValidate(zod(addArtisticProductionListSchema)),
        updateResearchInnovationProjectListForm: await superValidate(zod(updateArtisticProductionListSchema)),
    }
};

export const actions: Actions = {
    postArtisticProductionList: async (event) => {
        const form = await superValidate(event, zod(addArtisticProductionListSchema))

        if (!form.valid) {
            return generateFormMessageFromInvalidForm(form)
        }

        const token = event.cookies.get("AuthorizationToken")
        const data = form.data
        const request: ArtisticProductionList = {
            academic_period_id: data.academic_period_id,
            international_artistic_work: data.international_artistic_work,
            national_artistic_work: data.national_artistic_work,
            intellectual_property: data.intellectual_property
        }

        const response = await PostArtisticProductionList(token!, request)

        return generateFormMessageFromHttpResponse(form, response)
    },

    updateArtisticProductionList: async (event) => {
        const form = await superValidate(event, zod(updateArtisticProductionListSchema))

        if (!form.valid) {
            return generateFormMessageFromInvalidForm(form)
        }

        const token = event.cookies.get("AuthorizationToken")
        const data = form.data
        const request: ArtisticProductionList = {
            academic_period_id: data.academic_period_id,
            international_artistic_work: data.international_artistic_work,
            national_artistic_work: data.national_artistic_work,
            intellectual_property: data.intellectual_property
        }

        const response = await UpdateArtisticProductionList(token!, request)

        return generateFormMessageFromHttpResponse(form, response)
    },
};