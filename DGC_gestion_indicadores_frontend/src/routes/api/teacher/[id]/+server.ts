import { DeleteTeacher } from "$lib/api/controller/api/teacher";
import { error, json, type RequestHandler } from "@sveltejs/kit";

export const DELETE: RequestHandler = async ({ params, cookies }) => {
    const teacherID = params.id
    const token = cookies.get("AuthorizationToken");
    const res = await DeleteTeacher(token!, teacherID!)
    if (!res.ok) {
        throw error(res.status);
    }
    return json(await res.json());
}