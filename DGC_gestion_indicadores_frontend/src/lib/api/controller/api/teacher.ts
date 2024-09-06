import { deleteTeacherRoute, getTeachersPersonJoinedByTeacherIDRoute, postFilterTeachersRoute, postTeacherRoute, updateTeacherRoute } from "$lib/api/routes/api/teacher";
import type { AddTeacherRequest, FilterTeachersRequest, FilterTeachersResponse, Teacher, TeacherPerson, UpdateTeacherRequest } from "$lib/api/model/api/teacher";

import { generateCommonErrorFromFetchError } from "$lib/utils";
import type { CommonError } from "$lib/api/model/errors";
import type { Message } from "$lib/components/combobox/combobox";

export async function PostFilterTeachers(token: string, filterTeachersRequest: FilterTeachersRequest): Promise<FilterTeachersResponse | CommonError> {
    try {
        const response = await fetch(postFilterTeachersRoute, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': token
            },
            body: JSON.stringify(filterTeachersRequest)
        })
        if (!response.ok) {
            const error: CommonError = await response.json()
            return error
        }
        const filterPeopleResponse: FilterTeachersResponse = await response.json()
        return filterPeopleResponse

    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}

export async function GetTeacherPersonJoinedByTeacherID(token: string, teacherID: string): Promise<TeacherPerson | CommonError> {
    try {
        const response = await fetch(getTeachersPersonJoinedByTeacherIDRoute + teacherID, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': token
            },
        })
        if (!response.ok) {
            const error: CommonError = await response.json()
            return error
        }
        const teacherPerson: TeacherPerson = await response.json()
        return teacherPerson

    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}

export async function CreateTeacher(token: string, request: AddTeacherRequest) {
    try {
        const response = await fetch(postTeacherRoute, {
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
        const teacherCreated: Teacher = await response.json()
        return teacherCreated;
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}

export async function DeleteTeacher(token: string, teacherID: string) {
    return await fetch(deleteTeacherRoute + teacherID, {
        method: "DELETE",
        headers: {
            'Content-Type': 'application/json',
            'Authorization': token
        },
    })
}


export async function UpdateTeacher(token: string, request: UpdateTeacherRequest, teacherID: string) {
    return await fetch(updateTeacherRoute + teacherID, {
        method: "PUT",
        headers: {
            'Content-Type': 'application/json',
            'Authorization': token
        },
        body: JSON.stringify(request)
    })
}

export function GenerateComboMessagesFromTeachers(teachers: TeacherPerson[]): Message[] {
    let messages: Message[] = []
    messages = messages.concat(
        teachers.map((person) => ({
            value: person.ID,
            label: person.identity + " " + person.name + " " + person.lastname,
        }))
    );
    return messages
}