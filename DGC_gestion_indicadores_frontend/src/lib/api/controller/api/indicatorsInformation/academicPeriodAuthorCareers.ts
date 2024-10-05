import { getAcademicPeriodAuthorPreviousCareers } from "$lib/api/routes/api/indicatorsInformation/academicPeriodAuthorCareers";
import type { Career } from "$lib/api/model/api/career";

import { generateCommonErrorFromFetchError } from "$lib/utils";
import type { CommonError } from "$lib/api/model/errors";

export async function GetAcademicPeriodAuthorPreviousCareers(token: string, authorID: string): Promise<Career[] | CommonError> {
    try {
        const response = await fetch(getAcademicPeriodAuthorPreviousCareers + authorID, {
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
        const careers: Career[] = await response.json()
        return careers
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}