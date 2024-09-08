import type { AddDegreeAndTeachersListsDegreeRequest, GetTeachersListsDegreesJoinedResponse } from "$lib/api/model/api/indicatorsInformation/teachersListsDegree";
import { getTeachersListsDegreesJoinedRoute, postAddDegreeAndTeachersListsDegreeRoute } from "$lib/api/routes/api/indicatorsInformation/teachersDegree";
import { generateCommonErrorFromFetchError } from "$lib/utils";
import type { CommonError } from "$lib/api/model/errors";

export async function AddDegreeAndTeachersListsDegree(token: string, request: AddDegreeAndTeachersListsDegreeRequest) {
    try {
        const response = await fetch(postAddDegreeAndTeachersListsDegreeRoute, {
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
        const addDegreeAndTeachersListsDegreeRequest: AddDegreeAndTeachersListsDegreeRequest = await response.json()
        return addDegreeAndTeachersListsDegreeRequest;
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}

export async function GetTeachersListsDegreesJoined(academic_period_id: string, teacher_id: string, token: string) {
    try {
        const response = await fetch(getTeachersListsDegreesJoinedRoute + academic_period_id + "/" + teacher_id, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': token
            }
        });        

        if (!response.ok) {
            const error: CommonError = await response.json()
            return error
        }
        const degrees: GetTeachersListsDegreesJoinedResponse[] = await response.json()
        return degrees
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}