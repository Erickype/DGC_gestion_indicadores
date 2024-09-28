import { error, redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import { mainDashboarRoute } from "$lib/api/util/paths";
import { superValidate } from "sveltekit-superforms";
import { zod } from "sveltekit-superforms/adapters";
import { addAcademicProductionListsAuthorSchema } from "./schema";
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