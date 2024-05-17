import type { PostEvaluationPeriodRequest } from "$lib/api/model/view/evaluationPeriod";
import { postEvaluationPeriod } from "$lib/api/routes/admin/evaluationPeriod";

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
