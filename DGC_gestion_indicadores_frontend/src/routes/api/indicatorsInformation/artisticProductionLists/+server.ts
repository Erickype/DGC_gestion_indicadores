import { GetArtisticProductionLists } from "$lib/api/controller/api/indicatorsInformation/artisticProductionLists";

import { generateErrorFromCommonError, type CommonError } from "$lib/api/model/errors";
import { json, type RequestHandler } from "@sveltejs/kit";

export const GET: RequestHandler = async ({ cookies }) => {
    const token = cookies.get("AuthorizationToken")
    const response = await GetArtisticProductionLists(token!)

    if ((response as CommonError).status_code) {
        return generateErrorFromCommonError(response as CommonError)
    }
    return json(response);
};