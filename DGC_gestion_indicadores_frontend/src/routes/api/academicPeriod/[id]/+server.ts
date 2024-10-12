import { DeleteAcademicPeriod } from "$lib/api/controller/admin/academicPeriod";
import { GetAcademicPeriodByID } from "$lib/api/controller/view/academicPeriod";
import { generateErrorFromCommonError, type CommonError } from "$lib/api/model/errors";
import { error, json, type RequestHandler } from "@sveltejs/kit";

export const GET: RequestHandler = async ({ params }) => {
    const id = params.id
    const response = await GetAcademicPeriodByID(id!)
    if ((response as CommonError).status_code) {
        return generateErrorFromCommonError(response as CommonError)
    }
    return json(response);
};

export const DELETE: RequestHandler = async ({ cookies, params }) => {
    const id = params.id
    const token = cookies.get("AuthorizationToken")
    const res = await DeleteAcademicPeriod(id!, token!)
    if (!res.ok) {
        throw error(res.status, { message: await res.json() })
    }
    return json(await res.json());
};