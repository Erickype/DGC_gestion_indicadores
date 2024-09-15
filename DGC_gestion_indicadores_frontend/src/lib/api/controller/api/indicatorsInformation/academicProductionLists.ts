import type { FilterAcademicProductionListsByAcademicPeriodRequest, FilterAcademicProductionListsByAcademicPeriodResponse } from "$lib/api/model/api/indicatorsInformation/academicProductionLists";
import { postFilterAcademicProductionListByAcademicPeriodIDRoute } from "$lib/api/routes/api/indicatorsInformation/academicProductionLists";
import { generateCommonErrorFromFetchError } from "$lib/utils";
import type { CommonError } from "$lib/api/model/errors";

export async function PostAcademicProductionTeachersListsByAcademicPeriodID(token: string, request: FilterAcademicProductionListsByAcademicPeriodRequest) {
    try {
        const response = await fetch(postFilterAcademicProductionListByAcademicPeriodIDRoute, {
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
        const filterAcademicProductionListsByAcademicPeriodResponse: FilterAcademicProductionListsByAcademicPeriodResponse[] = await response.json()
        return filterAcademicProductionListsByAcademicPeriodResponse
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}