import { DeleteAcademicPeriod } from "$lib/api/controller/admin/academicPeriod";
import { error, json, type RequestHandler } from "@sveltejs/kit";

export const DELETE: RequestHandler = async ({ cookies, params }) => {
    const id = params.id
    const token = cookies.get("AuthorizationToken")
    const res = await DeleteAcademicPeriod(id!, token!)
    if (!res.ok) {
        throw error(res.status, { message: await res.json() })
    }
    return json(await res.json());
};