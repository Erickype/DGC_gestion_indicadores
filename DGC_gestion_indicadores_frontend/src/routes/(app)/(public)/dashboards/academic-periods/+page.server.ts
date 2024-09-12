import { LoadAcademicPeriodsWithComboMessages } from "$lib/api/controller/view/academicPeriod";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async () => {
    return {
        academicPeriodsData: await LoadAcademicPeriodsWithComboMessages(),
    }
};