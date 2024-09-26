import type { AcademicProductionListsAuthorsCareersJoined } from "$lib/api/model/api/indicatorsInformation/academicProductionListsAuthor";
import type { CommonError } from "$lib/api/model/errors";
import { getAcademicProductionListsAuthorsJoinedByAcademicProductionListID } from "$lib/api/routes/api/indicatorsInformation/academicProductionListsAuthor";
import { generateCommonErrorFromFetchError } from "$lib/utils";

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