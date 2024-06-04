import type { PostAcademicPeriodRequest, UpdateAcademicPeriodRequest } from "$lib/api/model/view/academicPeriod";
import { deleteAcademicPeriod, postAcademicPeriod, updateAcademicPeriod } from "$lib/api/routes/admin/academicPeriod";

export async function PostAcademicPeriod(period: PostAcademicPeriodRequest, token: string) {
    return await fetch(postAcademicPeriod, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': token
        },
        body: JSON.stringify(period)
    });
}

export async function UpdateAcademicPeriod(period: UpdateAcademicPeriodRequest, token: string) {
    return await fetch(updateAcademicPeriod + period.ID.toString(), {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': token
        },
        body: JSON.stringify(period)
    });
}

export async function DeleteAcademicPeriod(id: string, token: string) {
    return await fetch(deleteAcademicPeriod + id, {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': token
        }
    })
}