import { deleteTeacherRoute, getTeachersByAcademicPeriodIDRoute, postTeacherRoute, updateTeacherRoute } from "$lib/api/routes/api/teacher";
import type { AddTeacherRequest, CreateTeacherRequest, Teacher, UpdateTeacherRequest } from "$lib/api/model/api/teacher";

import { generateCommonErrorFromFetchError } from "$lib/utils";
import type { CommonError } from "$lib/api/model/errors";

export async function GetTeachersByAcademicPeriodID(token: string, academicPeriod: string) {
    try {
        const response = await fetch(getTeachersByAcademicPeriodIDRoute + academicPeriod, {
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
        const teachers: Teacher[] = await response.json()
        return teachers
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