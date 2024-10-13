import { GetCalculateIndicatorByTypeIDAndAcademicPeriod } from "$lib/api/controller/api/indicators/academicPeriods";

import { generateErrorFromCommonError, type CommonError } from "$lib/api/model/errors";
import { json, type RequestHandler } from "@sveltejs/kit";

export const GET: RequestHandler = async ({ cookies, params }) => {
    const token = cookies.get("AuthorizationToken")
    const indicatorTypeID = params.indicatorTypeID
    const academicPeriodID = params.academicPeriodID

    const response = await GetCalculateIndicatorByTypeIDAndAcademicPeriod(academicPeriodID!, indicatorTypeID!, token!)

    if ((response as CommonError).status_code) {
        return generateErrorFromCommonError(response as CommonError)
    }
    return json(response);
};