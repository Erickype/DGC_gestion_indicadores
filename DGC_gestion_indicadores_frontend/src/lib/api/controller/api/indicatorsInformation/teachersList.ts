import type { CreateTeachersListsRequest, TeachersList } from "$lib/api/model/api/indicatorsInformation/teachersLists";
import type { CommonError } from "$lib/api/model/errors";
import { postTeachersListRoute } from "$lib/api/routes/api/indicatorsInformation/teachersLists";
import { generateCommonErrorFromFetchError } from "$lib/utils";

export async function CreateTeachersList(token: string, request: CreateTeachersListsRequest) {
    try {
        const response = await fetch(postTeachersListRoute, {
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
        const teacherCreated: TeachersList = await response.json()
        return teacherCreated;
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}