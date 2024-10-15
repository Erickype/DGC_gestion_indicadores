import { getGradeRateListsByAcademicPeriodRoute } from "$lib/api/routes/api/indicatorsInformation/gradeRateLists";
import type { GradeRateListJoined } from "$lib/api/model/api/indicatorsInformation/gradeRateLists";

import { generateCommonErrorFromFetchError } from "$lib/utils";
import type { CommonError } from "$lib/api/model/errors";

export async function GetGradeRateListsByAcademicPeriod(token: string, academicPeriodID: string): Promise<GradeRateListJoined[] | CommonError> {
    try {
        const response = await fetch(getGradeRateListsByAcademicPeriodRoute + academicPeriodID, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': token
            },
        });

        if (!response.ok) {
            const error: CommonError = await response.json()
            return error
        }
        const gradeRateListsJoined: GradeRateListJoined[] = await response.json()
        return gradeRateListsJoined
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}