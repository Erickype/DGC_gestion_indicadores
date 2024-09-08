import type { AddDegreeAndTeachersListsDegreeRequest } from "$lib/api/model/api/indicatorsInformation/teachersListsDegree";
import { postAddDegreeAndTeachersListsDegreeRoute } from "$lib/api/routes/api/indicatorsInformation/teachersDegree";
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