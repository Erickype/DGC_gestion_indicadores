import { getPostgraduateCohortListsByProgramIDRoute, getPostgraduateProgramByProgramIDRoute, getPostgraduateProgramMissingCohortYearsByProgramIDRoute, postFilterCohortsRoute, postFilterPostgraduateProgramsRoute, postPostgraduateCohortListRoute, postPostgraduateProgramRoute, updatePostgraduateCohortListRoute, updatePostgraduateProgramRoute } from "$lib/api/routes/api/indicatorsInformation/postgraduate";
import type { FilterCohortListsRequest, FilterCohortListsResponse, FilterPostgraduateProgramsRequest, FilterPostGraduateProgramsResponse, GetPostgraduateProgramMissingCohortYearsByProgramIDResponse, PostgraduateCohortList, PostgraduateProgram } from "$lib/api/model/api/indicatorsInformation/postgraduate";
import type { Message } from "$lib/components/combobox/combobox";

import { generateCommonErrorFromFetchError } from "$lib/utils";
import type { CommonError } from "$lib/api/model/errors";

export async function PostFilterPostGraduatePrograms(token: string, request: FilterPostgraduateProgramsRequest) {
    try {
        const response = await fetch(postFilterPostgraduateProgramsRoute, {
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
        const filterPostGraduateProgramsResponse: FilterPostGraduateProgramsResponse[] = await response.json()
        return filterPostGraduateProgramsResponse
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}

export async function GetPostgraduateProgramByProgramID(token: string, programID: string) {
    try {
        const response = await fetch(getPostgraduateProgramByProgramIDRoute + programID, {
            method: "GET",
            headers: {
                'Content-Type': 'application/json',
                'Authorization': token
            },
        })
        if (!response.ok) {
            const error: CommonError = await response.json()
            return error
        }
        const postgraduateProgram: PostgraduateProgram = await response.json()
        return postgraduateProgram;
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}

export async function PostPostgraduateProgram(token: string, request: PostgraduateProgram) {
    try {
        const response = await fetch(postPostgraduateProgramRoute, {
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
        const postgraduateProgram: PostgraduateProgram = await response.json()
        return postgraduateProgram;
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}

export async function UpdatePostgraduateProgram(token: string, request: PostgraduateProgram) {
    try {
        const response = await fetch(updatePostgraduateProgramRoute + request.ID, {
            method: "PUT",
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
        const postgraduateProgram: PostgraduateProgram = await response.json()
        return postgraduateProgram;
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}

export function GenerateComboMessagesFromAuthors(programs: PostgraduateProgram[]): Message[] {
    let messages: Message[] = []
    messages = messages.concat(
        programs.map((program) => ({
            value: program.ID!,
            label: program.name,
        }))
    );
    return messages
}

export async function GetPostgraduateProgramMissingCohortYearsByProgramID(token: string, programID: string) {
    try {
        const response = await fetch(getPostgraduateProgramMissingCohortYearsByProgramIDRoute + programID, {
            method: "GET",
            headers: {
                'Content-Type': 'application/json',
                'Authorization': token
            },
        })
        if (!response.ok) {
            const error: CommonError = await response.json()
            return error
        }
        const missingYears: GetPostgraduateProgramMissingCohortYearsByProgramIDResponse = await response.json()
        return missingYears;
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}

export function GenerateComboMessagesFromMissingYears(programs: number[]): Message[] {
    let messages: Message[] = []
    messages = messages.concat(
        programs.map((year) => ({
            value: year,
            label: year.toString(),
        }))
    );
    return messages
}

export async function GetPostgraduateCohortListsByProgramID(token: string, programID: string) {
    try {
        const response = await fetch(getPostgraduateCohortListsByProgramIDRoute + programID, {
            method: "GET",
            headers: {
                'Content-Type': 'application/json',
                'Authorization': token
            },
        })
        if (!response.ok) {
            const error: CommonError = await response.json()
            return error
        }
        const postgraduateCohortLists: PostgraduateCohortList[] = await response.json()
        return postgraduateCohortLists;
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}

export async function PostPostgraduateCohortList(token: string, request: PostgraduateCohortList) {
    try {
        const response = await fetch(postPostgraduateCohortListRoute, {
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
        const postgraduateCohortList: PostgraduateCohortList = await response.json()
        return postgraduateCohortList;
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}

export async function UpdatePostgraduateCohortList(token: string, request: PostgraduateCohortList) {
    try {
        const response = await fetch(updatePostgraduateCohortListRoute + request.postgraduate_program_id + "/" + request.year, {
            method: "PUT",
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
        const postgraduateCohortList: PostgraduateCohortList = await response.json()
        return postgraduateCohortList;
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}

export async function PostFilterCohortLists(token: string, request: FilterCohortListsRequest) {
    try {
        const response = await fetch(postFilterCohortsRoute, {
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
        const filterCohortListsResponse: FilterCohortListsResponse = await response.json()
        return filterCohortListsResponse
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}
