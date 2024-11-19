import { GetCalculateIndicatorsByPostgraduateCohortYear } from "$lib/api/controller/api/indicators/postgraduateCohortYears";

import { generateErrorFromCommonError, type CommonError } from "$lib/api/model/errors";
import { json, type RequestHandler } from "@sveltejs/kit";

export const GET: RequestHandler = async ({ cookies, params }) => {
    const token = cookies.get("AuthorizationToken")
    const cohortYear = params.cohortYear

    const response = await GetCalculateIndicatorsByPostgraduateCohortYear(cohortYear!, token!)

    if ((response as CommonError).status_code) {
        return generateErrorFromCommonError(response as CommonError)
    }
    return json(response);
};