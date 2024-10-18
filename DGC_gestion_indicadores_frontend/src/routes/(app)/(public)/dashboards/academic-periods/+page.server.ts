import { LoadAcademicPeriodsWithComboMessages } from "$lib/api/controller/view/academicPeriod";
import type { PageServerLoad } from "./$types";
import { redirect } from "@sveltejs/kit";

export const load: PageServerLoad = async ({locals}) => {
    if (!locals.user) {
        throw redirect(302, "/login")
    }

    return {
        academicPeriodsData: await LoadAcademicPeriodsWithComboMessages(),
    }
};