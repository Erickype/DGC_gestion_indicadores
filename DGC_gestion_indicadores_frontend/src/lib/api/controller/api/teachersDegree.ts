import type { PatchTeachersDegreeByTeachersDegreeIDRequest, TeachersDegree } from "$lib/api/model/api/teachersDegree";
import { patchTeachersDegreeByTeachersDegreeIDRoute } from "$lib/api/routes/api/teachersDegree";
import { generateCommonErrorFromFetchError } from "$lib/utils";
import type { CommonError } from "$lib/api/model/errors";

export async function PatchTeachersDegreeByTeachersDegreeID(token: string, request: PatchTeachersDegreeByTeachersDegreeIDRequest, id: string): Promise<TeachersDegree | CommonError> {
    try {
        const response = await fetch(patchTeachersDegreeByTeachersDegreeIDRoute + id, {
            method: "PATCH",
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
        const teachersDegree: TeachersDegree = await response.json()
        return teachersDegree;
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}