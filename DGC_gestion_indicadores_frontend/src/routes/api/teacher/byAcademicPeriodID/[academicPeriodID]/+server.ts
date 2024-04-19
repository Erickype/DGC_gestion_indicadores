import { GetTeachersByAcademicPeriodID } from "$lib/api/controller/api/teacher";
import { error, json, type RequestHandler } from "@sveltejs/kit";

export const GET: RequestHandler = async ({ params, cookies }) => {
    const academicPeriodID = params.academicPeriodID;
    const token = cookies.get("AuthorizationToken");
    const res = await GetTeachersByAcademicPeriodID(token!, academicPeriodID!);

    if (!res.ok) {
        throw error(res.status);
    }
    return json(await res.json());
}
