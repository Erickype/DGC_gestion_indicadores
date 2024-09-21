import type { Actions, PageServerLoad } from "./$types";
import { mainDashboarRoute } from "$lib/api/util/paths";
import { redirect } from "@sveltejs/kit";

import { addAcademicProductionSchema } from "./schema";
import { superValidate } from "sveltekit-superforms";
import { zod } from "sveltekit-superforms/adapters";

import { LoadGetScienceMagazinesWithComboMessages } from "$lib/api/controller/api/academicProduction/scienceMagazines/scienceMagazine";
import { LoadAcademicPeriodsWithComboMessages } from "$lib/api/controller/view/academicPeriod";
import { LoadGetImpactCoefficientssWithComboMessages } from "$lib/api/controller/api/academicProduction/impactCoefficients/impactCoefficient";

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
        scienceMagazinesData: await LoadGetScienceMagazinesWithComboMessages(token!),
        impactCoefficientsData: await LoadGetImpactCoefficientssWithComboMessages(token!),
    }
};