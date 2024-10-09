import { redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import { LoadEvaluationPeriodsWithComboMessages } from "$lib/api/controller/view/evaluationPeriod";

export const load: PageServerLoad = async ({ locals }) => {
    if (!locals.user) {
        throw redirect(302, "/login")
    }

    return {
        evaluationPeriodsData: await LoadEvaluationPeriodsWithComboMessages()
    }
};