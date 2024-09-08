import { error, redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import { superValidate } from "sveltekit-superforms";
import { zod } from "sveltekit-superforms/adapters";
import { addTeachersListsDegreeSchema } from "./schema";
import { LoadDegreeLevelsWithComboMessages } from "$lib/api/controller/api/degreeLevels";
import { mainDashboarRoute } from "$lib/api/util/paths";

export const load: PageServerLoad = async ({ params, locals, cookies }) => {
    const token = cookies.get("AuthorizationToken")

    if (!locals.user) {
        throw redirect(302, "/login")
    }

    if (locals.user.role === 3) {
        throw redirect(302, mainDashboarRoute)
    }

    const academicPeriodID = parseInt(params.teacherID)
    const teacherID = parseInt(params.teacherID)
    if (academicPeriodID > 0 && teacherID > 0) {
        return {
            academicPeriodID: academicPeriodID,
            teacherID: teacherID,
            addTeachersListsDegreeForm: await superValidate(zod(addTeachersListsDegreeSchema)),
            degreeLevelsData: await LoadDegreeLevelsWithComboMessages(token!)
        }
    }

    error(404, 'Not found');
};