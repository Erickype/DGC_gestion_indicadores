import type { AddDegreeAndTeachersListsDegreeRequest, GetDegreesNotAssignedResponse, GetTeachersListsDegreesJoinedResponse, TeachersListsDegree } from "$lib/api/model/api/indicatorsInformation/teachersListsDegree";
import { getDegreesNotAssignedRoute, getTeachersListsDegreesJoinedRoute, postAddDegreeAndTeachersListsDegreeRoute, postTeachersListsDegreeRoute } from "$lib/api/routes/api/indicatorsInformation/teachersDegree";
import { generateErrorFromCommonError, type CommonError } from "$lib/api/model/errors";
import type { Message } from "$lib/components/combobox/combobox";
import { generateCommonErrorFromFetchError } from "$lib/utils";

export async function PostTeachersListsDegree(token: string, request: TeachersListsDegree) {
    try {
        const response = await fetch(postTeachersListsDegreeRoute, {
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
        const teachersListDegree: TeachersListsDegree = await response.json()
        return teachersListDegree;
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}

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

export async function GetDegreesNotAssigned(academic_period_id: string, teacher_id: string, token: string) {
    try {
        const response = await fetch(getDegreesNotAssignedRoute + academic_period_id + "/" + teacher_id, {
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
        const degrees: GetDegreesNotAssignedResponse[] = await response.json()
        return degrees
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

export async function LoadDegreesNotAssignedWithComboMessages(academicPeriodID: string, teacherID: string, token: string) {
    const response = await GetDegreesNotAssigned(academicPeriodID, teacherID, token);
    if ((response as CommonError).status) {
        throw generateErrorFromCommonError(response as CommonError)
    }
    const degress = response as GetDegreesNotAssignedResponse[];
    let messages: Message[] = []
    messages = messages.concat(
        degress.map((degree) => ({
            value: degree.teachers_degree_id,
            label: `${degree.abbreviation} ${degree.name}`,
        }))
    );

    return {
        degress,
        messages
    }
}

export function GenerateComboMessagesFromDegreesNotAssigned(degrees: GetDegreesNotAssignedResponse[]): Message[] {
    let messages: Message[] = []
    messages = messages.concat(
        degrees.map((degree) => ({
            value: degree.teachers_degree_id,
            label: `${degree.abbreviation} ${degree.name}`,
        }))
    );
    return messages
}