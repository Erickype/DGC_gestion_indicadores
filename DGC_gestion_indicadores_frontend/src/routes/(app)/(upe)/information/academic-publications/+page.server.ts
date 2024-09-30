import type { Actions, PageServerLoad } from "./$types";
import { mainDashboarRoute } from "$lib/api/util/paths";
import { redirect } from "@sveltejs/kit";

import { addAcademicProductionSchema, updateAcademicProductionSchema } from "./schema";
import { superValidate } from "sveltekit-superforms";
import { zod } from "sveltekit-superforms/adapters";

import { LoadGetImpactCoefficientssWithComboMessages } from "$lib/api/controller/api/academicProduction/impactCoefficients/impactCoefficient";
import { LoadGetScienceMagazinesWithComboMessages } from "$lib/api/controller/api/academicProduction/scienceMagazines/scienceMagazine";
import { PatchAcademicProductionList, PostAcademicProductionList } from "$lib/api/controller/api/indicatorsInformation/academicProductionLists";
import type { AcademicProductionList } from "$lib/api/model/api/indicatorsInformation/academicProductionLists";
import { generateFormMessageFromHttpResponse, generateFormMessageFromInvalidForm } from "$lib/utils";
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
        academicProductionForm: await superValidate(zod(addAcademicProductionSchema)),
        updateAcademicProductionForm: await superValidate(zod(updateAcademicProductionSchema)),
        scienceMagazinesData: await LoadGetScienceMagazinesWithComboMessages(token!),
        impactCoefficientsData: await LoadGetImpactCoefficientssWithComboMessages(token!),
    }
};

export const actions: Actions = {
    postAcademicProductionList: async (event) => {
        const form = await superValidate(event, zod(addAcademicProductionSchema))

        if (!form.valid) {
            return generateFormMessageFromInvalidForm(form)
        }

        const token = event.cookies.get("AuthorizationToken")
        const data = form.data
        const academicProduction: AcademicProductionList = {
            DOI: data.doi,
            publication_date: new Date(data.publication_date).toISOString(),
            publication_name: data.publication_name,
            academic_period_id: data.academicPeriod,
            detailed_field_id: data.detailed_field_id,
            science_magazine_id: data.science_magazine_id,
            impact_coefficient_id: data.impact_coefficient_id,
            intercultural_component: data.intercultural_component
        }

        const response = await PostAcademicProductionList(token!, academicProduction)

        return generateFormMessageFromHttpResponse(form, response)
    },

    patchAcademicProductionList: async (event) => {
        const form = await superValidate(event, zod(updateAcademicProductionSchema))
        const token = event.cookies.get("AuthorizationToken")

        if (!form.valid) {
            return generateFormMessageFromInvalidForm(form)
        }

        const data = form.data        
        const academicProductionList: AcademicProductionList = {
            ID: data.ID,
            DOI: data.doi,
            publication_date: new Date(data.publication_date).toISOString(),
            publication_name: data.publication_name,
            academic_period_id: data.academicPeriod,
            detailed_field_id: data.detailed_field_id,
            science_magazine_id: data.science_magazine_id,
            impact_coefficient_id: data.impact_coefficient_id,
            intercultural_component: data.intercultural_component
        }
        const response = await PatchAcademicProductionList(token!, academicProductionList)

        return generateFormMessageFromHttpResponse(form, response)
    }
};