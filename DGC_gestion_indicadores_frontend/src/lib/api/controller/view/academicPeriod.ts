import { getAcademicPeriodsRoute } from "$lib/api/routes/view/academicPeriod";

export async function GetAcademicPeriods() {
    return await fetch(getAcademicPeriodsRoute, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        }
    });
}