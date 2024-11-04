import { getArtisticProductionListsRoute, postArtisticProductionListRoute, updateArtisticProductionListRoute } from "$lib/api/routes/api/indicatorsInformation/artisticProductionLists";
import type { ArtisticProductionList, ArtisticProductionListJoined } from "$lib/api/model/api/indicatorsInformation/artisticProductionLists";

import { generateCommonErrorFromFetchError } from "$lib/utils";
import type { CommonError } from "$lib/api/model/errors";

export async function GetArtisticProductionLists(token: string): Promise<ArtisticProductionListJoined[] | CommonError> {
    try {
        const response = await fetch(getArtisticProductionListsRoute, {
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
        const artisticProductionListJoined: ArtisticProductionListJoined[] = await response.json()
        return artisticProductionListJoined
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}

export async function PostArtisticProductionList(token: string, request: ArtisticProductionList): Promise<ArtisticProductionList | CommonError> {
    try {
        const response = await fetch(postArtisticProductionListRoute, {
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
        const artisticProductionList: ArtisticProductionList = await response.json()
        return artisticProductionList
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}

export async function UpdateArtisticProductionList(token: string, request: ArtisticProductionList): Promise<ArtisticProductionList | CommonError> {
    try {
        const response = await fetch(updateArtisticProductionListRoute + request.academic_period_id, {
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
        const artisticProductionList: ArtisticProductionList = await response.json()
        return artisticProductionList
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}