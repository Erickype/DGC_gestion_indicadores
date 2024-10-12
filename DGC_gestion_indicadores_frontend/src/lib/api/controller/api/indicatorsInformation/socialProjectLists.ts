import type { FilterSocialProjectListsByAcademicPeriodRequest, FilterSocialProjectListsByAcademicPeriodResponse, SocialProjectList } from "$lib/api/model/api/indicatorsInformation/socialProjectLists";
import { postFilterSocialProjectListsByAcademicPeriodRoute, postSocialProjectListRoute, putSocialProjectListRoute } from "$lib/api/routes/api/indicatorsInformation/socialProjectLists";

import { generateCommonErrorFromFetchError } from "$lib/utils";
import type { CommonError } from "$lib/api/model/errors";

export async function PostSocialProjectList(token: string, request: SocialProjectList): Promise<SocialProjectList | CommonError> {
    try {
        const response = await fetch(postSocialProjectListRoute, {
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
        const socialProjectList: SocialProjectList = await response.json()
        return socialProjectList
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}

export async function PutSocialProjectList(token: string, request: SocialProjectList): Promise<SocialProjectList | CommonError> {
    try {
        const response = await fetch(putSocialProjectListRoute + request.ID, {
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
        const socialProjectList: SocialProjectList = await response.json()
        return socialProjectList
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}

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