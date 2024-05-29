import { DeleteEvaluationPeriod } from "$lib/api/controller/admin/evaluationPeriod";
import { error, json, type RequestHandler } from "@sveltejs/kit";

export const DELETE: RequestHandler = async ({ cookies, params }) => {
    const id = params.id
    const token = cookies.get("AuthorizationToken")
    const res = await DeleteEvaluationPeriod(id!, token!)
    if (!res.ok) {
        throw error(res.status, { message: await res.json() })
    }
    return json(await res.json());
};