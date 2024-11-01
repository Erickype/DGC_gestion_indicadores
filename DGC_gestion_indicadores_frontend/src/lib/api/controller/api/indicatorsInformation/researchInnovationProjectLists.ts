import { getResearchInnovationProjectListsByAcademicPeriodRoute, postResearchInnovationProjectListRoute, updateResearchInnovationProjectListRoute } from "$lib/api/routes/api/indicatorsInformation/researchInnovationProjectLists";
import type { ResearchInnovationProjectList } from "$lib/api/model/api/indicatorsInformation/researchInnovationProjectLists";

import { generateCommonErrorFromFetchError } from "$lib/utils";
import type { CommonError } from "$lib/api/model/errors";

export async function GetResearchInnovationProjectListsByAcademicPeriod(token: string, academicPeriodID: string): Promise<ResearchInnovationProjectList[] | CommonError> {
    try {
        const response = await fetch(getResearchInnovationProjectListsByAcademicPeriodRoute + academicPeriodID, {
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
        const researchInnovationProjectListsJoined: ResearchInnovationProjectList[] = await response.json()
        return researchInnovationProjectListsJoined
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}

export async function PostResearchInnovationProjectList(token: string, request: ResearchInnovationProjectList): Promise<ResearchInnovationProjectList | CommonError> {
    try {
        const response = await fetch(postResearchInnovationProjectListRoute, {
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
        const researchInnovationProjectList: ResearchInnovationProjectList = await response.json()
        return researchInnovationProjectList
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}

export async function UpdateResearchInnovationProjectList(token: string, request: ResearchInnovationProjectList): Promise<ResearchInnovationProjectList | CommonError> {
    try {
        const response = await fetch(updateResearchInnovationProjectListRoute + request.academic_period_id, {
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
        const researchInnovationProjectList: ResearchInnovationProjectList = await response.json()
        return researchInnovationProjectList
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}