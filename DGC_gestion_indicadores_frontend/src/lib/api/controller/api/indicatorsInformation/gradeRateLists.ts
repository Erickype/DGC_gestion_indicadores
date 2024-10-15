import { getGradeRateListsByAcademicPeriodRoute, postGradeRateListRoute } from "$lib/api/routes/api/indicatorsInformation/gradeRateLists";
import type { GradeRateList, GradeRateListJoined } from "$lib/api/model/api/indicatorsInformation/gradeRateLists";

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

export async function PostGradeRateList(token: string, request: GradeRateList): Promise<GradeRateList | CommonError> {
    try {
        const response = await fetch(postGradeRateListRoute, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': token
            },
            body: JSON.stringify(request)
        });

        if (!response.ok) {
            const error: CommonError = await response.json()
            return error
        }
        const gradeRateList: GradeRateList = await response.json()
        return gradeRateList
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}