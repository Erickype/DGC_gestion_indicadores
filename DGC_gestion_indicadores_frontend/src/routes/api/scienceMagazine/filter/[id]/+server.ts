import { GetScienceMagazineFilterJoinedByScienceMagazineID } from "$lib/api/controller/api/academicProduction/scienceMagazines/scienceMagazine";
import { generateErrorFromCommonError, type CommonError } from "$lib/api/model/errors";
import { json, type RequestHandler } from "@sveltejs/kit";

export const GET: RequestHandler = async ({ cookies, params }) => {
    const scienceMagazineID = params.id
    const token = cookies.get("AuthorizationToken")
    const response = await GetScienceMagazineFilterJoinedByScienceMagazineID(token!, scienceMagazineID!)

    if ((response as CommonError).status_code) {
        return generateErrorFromCommonError(response as CommonError)
    }
    return json(response);
};