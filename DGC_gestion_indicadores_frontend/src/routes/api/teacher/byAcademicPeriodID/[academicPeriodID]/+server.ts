import { GetTeachersByAcademicPeriodID } from "$lib/api/controller/api/teacher";
import { generateErrorFromCommonError, type CommonError } from "$lib/api/model/errors";
import { json, type RequestHandler } from "@sveltejs/kit";

export const GET: RequestHandler = async ({ params, cookies }) => {
    const academicPeriodID = params.academicPeriodID;
    const token = cookies.get("AuthorizationToken");
    const response = await GetTeachersByAcademicPeriodID(token!, academicPeriodID!);

    if ((response as CommonError).status_code) {
        return generateErrorFromCommonError(response as CommonError)
    }
    return json(response);
}
