import { GetPostgraduateProgramMissingCohortYearsByProgramID } from "$lib/api/controller/api/indicatorsInformation/postgraduate";
import { generateErrorFromCommonError, type CommonError } from "$lib/api/model/errors";
import { json, type RequestHandler } from "@sveltejs/kit";

export const GET: RequestHandler = async ({ cookies, params }) => {
    const programID = params.programID
    const token = cookies.get("AuthorizationToken")
    const response = await GetPostgraduateProgramMissingCohortYearsByProgramID(token!, programID!)

    if ((response as CommonError).status_code) {
        return generateErrorFromCommonError(response as CommonError)
    }
    return json(response);
};