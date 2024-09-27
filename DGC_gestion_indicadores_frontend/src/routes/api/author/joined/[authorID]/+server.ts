import { GetAuthorPersonByAuthorID } from "$lib/api/controller/api/academicProduction/authors/author";
import { generateErrorFromCommonError, type CommonError } from "$lib/api/model/errors";
import { json, type RequestHandler } from "@sveltejs/kit";

export const GET: RequestHandler = async ({ cookies, params }) => {
    const authorID = params.authorID
    const token = cookies.get("AuthorizationToken")
    const response = await GetAuthorPersonByAuthorID(token!, parseInt(authorID!))

    if ((response as CommonError).status_code) {
        return generateErrorFromCommonError(response as CommonError)
    }
    return json(response);
};