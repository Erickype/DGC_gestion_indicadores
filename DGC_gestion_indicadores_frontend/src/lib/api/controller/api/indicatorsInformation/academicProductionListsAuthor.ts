import { getAcademicProductionListsAuthorPreviousCareersByAuthorID, getAcademicProductionListsAuthorsJoinedByAcademicProductionListID } from "$lib/api/routes/api/indicatorsInformation/academicProductionListsAuthor";
import type { AcademicProductionListsAuthorsCareersJoined } from "$lib/api/model/api/indicatorsInformation/academicProductionListsAuthor";
import type { Career } from "$lib/api/model/api/career";

import { generateCommonErrorFromFetchError } from "$lib/utils";
import type { CommonError } from "$lib/api/model/errors";

export async function GetAcademicProductionListsAuthorsJoinedByAcademicProductionListID(token: string, academicProductionListID: string): Promise<AcademicProductionListsAuthorsCareersJoined[] | CommonError> {
    try {
        const response = await fetch(getAcademicProductionListsAuthorsJoinedByAcademicProductionListID + academicProductionListID, {
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
        const academicProductionListsAuthorsCareersJoined: AcademicProductionListsAuthorsCareersJoined[] = await response.json()
        return academicProductionListsAuthorsCareersJoined
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}

export async function GetAcademicProductionListAuthorPreviousCareers(token: string, authorID: string): Promise<Career[] | CommonError> {
    try {
        const response = await fetch(getAcademicProductionListsAuthorPreviousCareersByAuthorID + authorID, {
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
