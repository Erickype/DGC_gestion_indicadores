import type { CommonDeleteResponse } from "$lib/api/model/common";
import type { CommonError } from "$lib/api/model/errors";
import type { PostEvaluationPeriodRequest, UpdateEvaluationPeriodRequest } from "$lib/api/model/view/evaluationPeriod";
import { deleteEvaluationPeriod, postEvaluationPeriod, updateEvaluationPeriod } from "$lib/api/routes/admin/evaluationPeriod";
import { generateCommonErrorFromFetchError } from "$lib/utils";

export async function PostEvaluationPeriod(period: PostEvaluationPeriodRequest, token: string) {
    return await fetch(postEvaluationPeriod, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': token
        },
        body: JSON.stringify(period)
    });
}

export async function UpdateEvaluationPeriod(period: UpdateEvaluationPeriodRequest, token: string) {
    return await fetch(updateEvaluationPeriod + period.ID.toString(), {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': token
        },
        body: JSON.stringify(period)
    });
}

export async function DeleteEvaluationPeriod(id: string, token: string): Promise<CommonDeleteResponse | CommonError> {
    try {
        const response = await fetch(deleteEvaluationPeriod + id, {
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
        const author: CommonDeleteResponse = await response.json()
        return author;
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}