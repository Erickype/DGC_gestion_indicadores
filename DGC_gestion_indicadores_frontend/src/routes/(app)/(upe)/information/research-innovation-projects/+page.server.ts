import { mainDashboarRoute } from "$lib/api/util/paths";
import type { PageServerLoad, Actions } from "./$types";
import { redirect } from "@sveltejs/kit"

import { addResearchInnovationProjectListSchema, updateResearchInnovationProjectListSchema } from "./schema";
import { generateFormMessageFromHttpResponse, generateFormMessageFromInvalidForm } from "$lib/utils";
import { superValidate } from "sveltekit-superforms";
import { zod } from "sveltekit-superforms/adapters";

import { PostResearchInnovationProjectList, UpdateResearchInnovationProjectList } from "$lib/api/controller/api/indicatorsInformation/researchInnovationProjectLists";
import type { ResearchInnovationProjectList } from "$lib/api/model/api/indicatorsInformation/researchInnovationProjectLists";

export const load: PageServerLoad = async ({ cookies, locals }) => {
    const token = cookies.get("AuthorizationToken")
    if (!locals.user) {
        throw redirect(302, "/login")
    }
    if (locals.user.role === 3) {
        throw redirect(302, mainDashboarRoute)
    }

    return {
        addResearchInnovationProjectListForm: await superValidate(zod(addResearchInnovationProjectListSchema)),
        updateResearchInnovationProjectListForm: await superValidate(zod(updateResearchInnovationProjectListSchema)),
    }
};

export const actions: Actions = {
    postResearchInnovationProjectList: async (event) => {
        const form = await superValidate(event, zod(addResearchInnovationProjectListSchema))

        if (!form.valid) {
            return generateFormMessageFromInvalidForm(form)
        }

        const token = event.cookies.get("AuthorizationToken")
        const data = form.data
        const request: ResearchInnovationProjectList = {
            academic_period_id: data.academic_period_id,
            total_projects_uep: data.total_projects_uep,
            projects_external_funding: data.projects_external_funding,
            international_cooperation_projects: data.international_cooperation_projects,
            national_cooperation_projects: data.national_cooperation_projects
        }

        const response = await PostResearchInnovationProjectList(token!, request)

        return generateFormMessageFromHttpResponse(form, response)
    },

    updateResearchInnovationProjectList: async (event) => {
        const form = await superValidate(event, zod(updateResearchInnovationProjectListSchema))

        if (!form.valid) {
            return generateFormMessageFromInvalidForm(form)
        }

        const token = event.cookies.get("AuthorizationToken")
        const data = form.data
        const request: ResearchInnovationProjectList = {
            academic_period_id: data.academic_period_id,
            total_projects_uep: data.total_projects_uep,
            projects_external_funding: data.projects_external_funding,
            international_cooperation_projects: data.international_cooperation_projects,
            national_cooperation_projects: data.national_cooperation_projects
        }

        const response = await UpdateResearchInnovationProjectList(token!, request)

        return generateFormMessageFromHttpResponse(form, response)
    },
};