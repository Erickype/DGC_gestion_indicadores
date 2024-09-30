import type { AcademicProductionList, FilterAcademicProductionListsByAcademicPeriodRequest, FilterAcademicProductionListsByAcademicPeriodResponse } from "$lib/api/model/api/indicatorsInformation/academicProductionLists";
import { patchAcademicProductionListRoute, postAcademicProductionListRoute, postFilterAcademicProductionListByAcademicPeriodIDRoute } from "$lib/api/routes/api/indicatorsInformation/academicProductionLists";
import { generateCommonErrorFromFetchError } from "$lib/utils";
import type { CommonError } from "$lib/api/model/errors";

export async function PostFilterAcademicProductionListsByAcademicPeriodID(token: string, request: FilterAcademicProductionListsByAcademicPeriodRequest) {
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

export async function PostAcademicProductionList(token: string, request: AcademicProductionList) {
    try {
        const response = await fetch(postAcademicProductionListRoute, {
            method: "POST",
            headers: {
                'Content-Type': 'application/json',
                'Authorization': token
            },
            body: JSON.stringify(request)
        })
        if (!response.ok) {
            const error: CommonError = await response.json()
            return error
        }
        const academicProduction: AcademicProductionList = await response.json()
        return academicProduction;
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}

export async function PatchAcademicProductionList(token: string, request: AcademicProductionList) {
    try {
        const response = await fetch(patchAcademicProductionListRoute + request.ID, {
            method: "PATCH",
            headers: {
                'Content-Type': 'application/json',
                'Authorization': token
            },
            body: JSON.stringify(request)
        })
        if (!response.ok) {
            const error: CommonError = await response.json()
            return error
        }
        const academicProductionList: AcademicProductionList = await response.json()
        return academicProductionList;
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}