import { DeletePerson } from "$lib/api/controller/api/person";
import { error, json, type RequestHandler } from "@sveltejs/kit";

export const DELETE: RequestHandler = async ({ cookies, params }) => {
    const id = params.id
    const token = cookies.get("AuthorizationToken")
    const res = await DeletePerson(id!, token!)
    if (!res.ok) {
        throw error(res.status, { message: await res.json() })
    }
    return json(await res.json());
};