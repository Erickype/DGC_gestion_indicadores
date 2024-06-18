import type { PostPersonRequest } from "$lib/api/model/api/person";
import { postPersonRoute } from "$lib/api/routes/admin/person";

export async function PostPerson(person: PostPersonRequest, token: string) {
    return await fetch(postPersonRoute, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': token
        },
        body: JSON.stringify(person)
    });
}