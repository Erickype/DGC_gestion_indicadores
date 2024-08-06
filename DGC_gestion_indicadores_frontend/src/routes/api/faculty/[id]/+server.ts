import { DeleteFaculty } from "$lib/api/controller/admin/faculty";
import { generateErrorFromCommonError, type CommonError } from "$lib/api/model/errors";
import { json, type RequestHandler } from "@sveltejs/kit";

export const DELETE: RequestHandler = async ({ cookies, params }) => {
    const id = params.id
    const token = cookies.get("AuthorizationToken")
    const response = await DeleteFaculty(id!, token!)
    if ((response as CommonError).status_code) {
        return generateErrorFromCommonError(response as CommonError)
    }
    return json(response);
};