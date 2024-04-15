import { redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import { mainDashboarRoute } from "$lib/api/util/paths";
import { LoadAcademicPeriodsWithComboMessages } from "$lib/api/controller/view/academicPeriod";
import { LoadPeopleWithComboMessages } from "$lib/api/controller/api/person";

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
        peopleData: await LoadPeopleWithComboMessages(token!),
    }
};