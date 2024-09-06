import { GetTeacherPersonJoinedByTeacherID } from "$lib/api/controller/api/teacher";
import { generateErrorFromCommonError, type CommonError } from "$lib/api/model/errors";
import { json, type RequestHandler } from "@sveltejs/kit";

export const GET: RequestHandler = async ({ cookies, params }) => {
    const teacherID = params.id
    const token = cookies.get("AuthorizationToken")
    const response = await GetTeacherPersonJoinedByTeacherID(token!, teacherID!)

    if ((response as CommonError).status_code) {
        return generateErrorFromCommonError(response as CommonError)
    }
    return json(response);
};