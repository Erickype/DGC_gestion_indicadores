import { PostFilterTeachersListsByAcademicPeriodID } from "$lib/api/controller/api/indicatorsInformation/teachersList";
import { PostFilterPeople } from "$lib/api/controller/api/person";
import { generateErrorFromCommonError, type CommonError } from "$lib/api/model/errors";
import { json, type RequestHandler } from "@sveltejs/kit";

export const POST: RequestHandler = async ({ cookies, request }) => {
    const token = cookies.get("AuthorizationToken")
    const response = await PostFilterTeachersListsByAcademicPeriodID(token!, await request.json())
    
    if ((response as CommonError).status_code) {
        return generateErrorFromCommonError(response as CommonError)
    }
    return json(response);
};