import { GetBooksOrChaptersProductionListsAuthorsJoinedByBooksOrChaptersProductionListID } from "$lib/api/controller/api/indicatorsInformation/booksOrChaptersProductionListsAuthor";
import { generateErrorFromCommonError, type CommonError } from "$lib/api/model/errors";
import { json, type RequestHandler } from "@sveltejs/kit";

export const GET: RequestHandler = async ({ cookies, params }) => {
    const token = cookies.get("AuthorizationToken")
    const booksOrChaptersProductionListID = params.booksOrChaptersProductionListID

    const response = await GetBooksOrChaptersProductionListsAuthorsJoinedByBooksOrChaptersProductionListID(token!, booksOrChaptersProductionListID!)

    if ((response as CommonError).status_code) {
        return generateErrorFromCommonError(response as CommonError)
    }
    return json(response);
};