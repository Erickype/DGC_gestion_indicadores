import { GetAcademicProductionListsAuthorsJoinedByAcademicProductionListID } from "$lib/api/controller/api/indicatorsInformation/academicProductionListsAuthor";
import { generateErrorFromCommonError, type CommonError } from "$lib/api/model/errors";
import { json, type RequestHandler } from "@sveltejs/kit";

export const GET: RequestHandler = async ({ cookies, params }) => {
    const token = cookies.get("AuthorizationToken")
    const academicProductionID = params.academicProductionListID
    const response = await GetAcademicProductionListsAuthorsJoinedByAcademicProductionListID(token!, academicProductionID!)

    if ((response as CommonError).status_code) {
        return generateErrorFromCommonError(response as CommonError)
    }
    return json(response);
};