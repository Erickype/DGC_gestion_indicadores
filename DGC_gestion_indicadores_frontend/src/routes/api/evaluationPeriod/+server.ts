import { type CommonError, generateErrorFromCommonError } from "$lib/api/model/errors";
import { GetEvaluationPeriods } from "$lib/api/controller/view/evaluationPeriod";
import { json, type RequestHandler } from "@sveltejs/kit";

export const GET: RequestHandler = async () => {
    const response = await GetEvaluationPeriods()
    if ((response as CommonError).status_code) {
        return generateErrorFromCommonError(response as CommonError)
    }
    return json(response);
};