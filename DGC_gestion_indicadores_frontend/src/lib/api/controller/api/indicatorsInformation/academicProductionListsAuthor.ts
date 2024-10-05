import { getAcademicProductionListsAuthorsJoinedByAcademicProductionListID, postAcademicProductionListsAuthorCareersRoute, updateAcademicProductionListsAuthorCareersRoute } from "$lib/api/routes/api/indicatorsInformation/academicProductionListsAuthor";
import type { AcademicProductionListsAuthorsCareersJoined, PostAcademicProductionListsAuthorCareersRequest, UpdateAcademicProductionListsAuthorCareersRequest } from "$lib/api/model/api/indicatorsInformation/academicProductionListsAuthor";

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

export async function UpdateAcademicProductionListsAuthorCareers(token: string, request: UpdateAcademicProductionListsAuthorCareersRequest) {
    try {
        const response = await fetch(updateAcademicProductionListsAuthorCareersRoute, {
            method: 'PUT',
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
        const postAcademicProductionListsAuthorCareersRequest: UpdateAcademicProductionListsAuthorCareersRequest[] = await response.json()
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
