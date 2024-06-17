import { GetPeople } from "$lib/api/controller/api/person";
import { generateErrorFromCommonError, type CommonError } from "$lib/api/model/errors";
import { json, type RequestHandler } from "@sveltejs/kit";

export const GET: RequestHandler = async ({cookies}) => {
    const token = cookies.get("AuthorizationToken")
    const response = await GetPeople(token!)
    
    if ((response as CommonError).status_code) {
        return generateErrorFromCommonError(response as CommonError)
    }
    return json(response);
};