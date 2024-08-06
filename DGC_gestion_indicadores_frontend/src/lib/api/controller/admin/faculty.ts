import { deleteFacultyRoute, postFacultyRoute, putFacultyRoute } from "$lib/api/routes/admin/faculty";

import type { PostFacultyRequest, PutFacultyRequest } from "$lib/api/model/admin/faculty";
import type { Faculty } from "$lib/api/model/api/faculty";
import type { CommonError } from "$lib/api/model/errors";
import { generateCommonErrorFromFetchError } from "$lib/utils";
import type { CommonDeleteResponse } from "$lib/api/model/common";

export async function PostFaculty(faculty: PostFacultyRequest, token: string): Promise<Faculty | CommonError> {
    try {
        const response = await fetch(postFacultyRoute, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': token
            },
            body: JSON.stringify(faculty)
        });
        if (!response.ok) {
            const error: CommonError = await response.json()
            return error
        }
        const facultyCreated: Faculty = await response.json()
        return facultyCreated;
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}

export async function PutFaculty(faculty: PutFacultyRequest, token: string): Promise<Faculty | CommonError> {
    try {
        const response = await fetch(putFacultyRoute + faculty.ID.toString(), {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': token
            },
            body: JSON.stringify(faculty)
        });
        if (!response.ok) {
            const error: CommonError = await response.json()
            return error
        }
        const facultyCreated: Faculty = await response.json()
        return facultyCreated;
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}

export async function DeleteFaculty(id: string, token: string): Promise<CommonDeleteResponse | CommonError> {
    try {
        const response = await fetch(deleteFacultyRoute + id, {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': token
            }
        })
        if (!response.ok) {
            const error: CommonError = await response.json()
            return error
        }
        const commonDeleteResponse: CommonDeleteResponse = await response.json()        
        return commonDeleteResponse
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}