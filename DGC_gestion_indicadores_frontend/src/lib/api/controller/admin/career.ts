import type { PostCareerRequest, PutCareerRequest } from "$lib/api/model/admin/career";
import type { Career } from "$lib/api/model/api/career";
import type { CommonError } from "$lib/api/model/errors";
import { deleteCareerRoute, postCareerRoute, putCareerRoute } from "$lib/api/routes/admin/career";
import { generateCommonErrorFromFetchError } from "$lib/utils";

export async function PostCareer(career: PostCareerRequest, token: string): Promise<Career | CommonError> {
    try {
        const response = await fetch(postCareerRoute, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': token
            },
            body: JSON.stringify(career)
        });
        if (!response.ok) {
            const error: CommonError = await response.json()
            return error
        }
        const careerCreated: Career = await response.json()
        return careerCreated;
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}

export async function PutCareer(career: PutCareerRequest, token: string) {
    try {
        const response = await fetch(putCareerRoute + career.ID.toString(), {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': token
            },
            body: JSON.stringify(career)
        });
        if (!response.ok) {
            const error: CommonError = await response.json()
            return error
        }
        const careerCreated: Career = await response.json()
        return careerCreated;
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}

export async function DeleteCareer(id: string, token: string) {
    return await fetch(deleteCareerRoute + id, {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': token
        }
    })
}