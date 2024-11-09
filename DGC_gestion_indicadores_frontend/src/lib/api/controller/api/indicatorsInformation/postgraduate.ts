import { getPostgraduateProgramByProgramIDRoute, postFilterPostgraduateProgramsRoute, postPostgraduateProgramRoute, updatePostgraduateProgramRoute } from "$lib/api/routes/api/indicatorsInformation/postgraduate";
import type { FilterPostgraduateProgramsRequest, FilterPostGraduateProgramsResponse, PostgraduateProgram } from "$lib/api/model/api/indicatorsInformation/postgraduate";
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