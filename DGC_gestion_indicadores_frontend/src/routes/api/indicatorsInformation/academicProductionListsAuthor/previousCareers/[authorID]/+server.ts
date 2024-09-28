import { GetAcademicProductionListAuthorPreviousCareers } from "$lib/api/controller/api/indicatorsInformation/academicProductionListsAuthor";
import { generateErrorFromCommonError, type CommonError } from "$lib/api/model/errors";
import { json, type RequestHandler } from "@sveltejs/kit";

export const GET: RequestHandler = async ({ cookies, params }) => {
    const token = cookies.get("AuthorizationToken")
    const authorID = params.authorID
    const response = await GetAcademicProductionListAuthorPreviousCareers(token!, authorID!)

    if ((response as CommonError).status_code) {
        return generateErrorFromCommonError(response as CommonError)
    }
    return json(response);
};