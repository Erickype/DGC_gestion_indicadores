import type { FilterSocialProjectListsByAcademicPeriodRequest, FilterSocialProjectListsByAcademicPeriodResponse } from "$lib/api/model/api/indicatorsInformation/socialProjectLists";
import { postFilterSocialProjectListsByAcademicPeriodRoute } from "$lib/api/routes/api/indicatorsInformation/socialProjectLists";

import { generateCommonErrorFromFetchError } from "$lib/utils";
import type { CommonError } from "$lib/api/model/errors";

export async function PostFilterSocialProjectListsByAcademicPeriod(token: string, request: FilterSocialProjectListsByAcademicPeriodRequest) {
    try {
        const response = await fetch(postFilterSocialProjectListsByAcademicPeriodRoute, {
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
        const filterSocialProjectListsByAcademicPeriodResponse: FilterSocialProjectListsByAcademicPeriodResponse[] = await response.json()
        return filterSocialProjectListsByAcademicPeriodResponse
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}