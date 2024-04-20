import type { CreateTeacherRequest } from "$lib/api/model/api/teacher";
import { deleteTeacherRoute, getTeachersByAcademicPeriodIDRoute, postTeacherRoute } from "$lib/api/routes/api/teacher";

export async function GetTeachersByAcademicPeriodID(token: string, academicPeriod: string) {
    return await fetch(getTeachersByAcademicPeriodIDRoute + academicPeriod, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': token
        }
    });
}

export async function CreateTeacher(token: string, request: CreateTeacherRequest) {
    return await fetch(postTeacherRoute, {
        method: "POST",
        headers: {
            'Content-Type': 'application/json',
            'Authorization': token
        },
        body: JSON.stringify(request)
    })
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
