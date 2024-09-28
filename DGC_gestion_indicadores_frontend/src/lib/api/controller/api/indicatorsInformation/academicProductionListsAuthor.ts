import { getAcademicProductionListsAuthorPreviousCareersByAuthorID, getAcademicProductionListsAuthorsJoinedByAcademicProductionListID, postAcademicProductionListsAuthorCareersRoute } from "$lib/api/routes/api/indicatorsInformation/academicProductionListsAuthor";
import type { AcademicProductionListsAuthorsCareersJoined, PostAcademicProductionListsAuthorCareersRequest } from "$lib/api/model/api/indicatorsInformation/academicProductionListsAuthor";
import type { Career } from "$lib/api/model/api/career";

import { generateCommonErrorFromFetchError } from "$lib/utils";
import type { CommonError } from "$lib/api/model/errors";

export async function PostAcademicProductionListsAuthorCareers(token: string, request: PostAcademicProductionListsAuthorCareersRequest) {
    try {
        const response = await fetch(postAcademicProductionListsAuthorCareersRoute, {
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
        const postAcademicProductionListsAuthorCareersRequest: PostAcademicProductionListsAuthorCareersRequest[] = await response.json()
        return postAcademicProductionListsAuthorCareersRequest
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}

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
