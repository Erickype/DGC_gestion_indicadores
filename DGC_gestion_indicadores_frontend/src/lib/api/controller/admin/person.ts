import type { PostPersonRequest, PutPersonRequest } from "$lib/api/model/api/person";
import { deletePersonRoute, postPersonRoute, putPersonRoute } from "$lib/api/routes/admin/person";

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

export async function PutPerson(person: PutPersonRequest, token: string) {
    return await fetch(putPersonRoute + person.ID.toString(), {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': token
        },
        body: JSON.stringify(person)
    });
}

export async function DeletePerson(id: string, token: string) {
    return await fetch(deletePersonRoute + id, {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': token
        }
    })
}