import { GetAcademicPeriods } from "$lib/api/controller/view/academicPeriod";
import { generateErrorFromCommonError, type CommonError } from "$lib/api/model/errors";
import { json, type RequestHandler } from "@sveltejs/kit";

export const GET: RequestHandler = async () => {
    const response = await GetAcademicPeriods()
    
    if ((response as CommonError).StatusCode) {
        return generateErrorFromCommonError(response as CommonError)
    }
    return json(response);
};