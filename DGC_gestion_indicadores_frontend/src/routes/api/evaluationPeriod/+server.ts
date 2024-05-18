import { GetEvaluationPeriods } from "$lib/api/controller/view/evaluationPeriod";
import { error, json, type RequestHandler } from "@sveltejs/kit";

export const GET: RequestHandler = async () => {
    const res = await GetEvaluationPeriods()
    if (!res.ok) {
        throw error(res.status, { message: await res.json() })
    }
    return json(await res.json());
};