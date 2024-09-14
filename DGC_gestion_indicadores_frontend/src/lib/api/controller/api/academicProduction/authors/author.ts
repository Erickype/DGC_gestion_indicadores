import { postAuthorFromPersontRoute, postAuthorRoute } from "$lib/api/routes/api/academicProduction/authors/author";
import type { Author, PostAuthorFromPersonRequest } from "$lib/api/model/api/academicProduction/authors/author";
import { generateCommonErrorFromFetchError } from "$lib/utils";
import type { CommonError } from "$lib/api/model/errors";

export async function PostAuthor(token: string, request: Author) {
    try {
        const response = await fetch(postAuthorRoute, {
            method: "POST",
            headers: {
                'Content-Type': 'application/json',
                'Authorization': token
            },
            body: JSON.stringify(request)
        })
        if (!response.ok) {
            const error: CommonError = await response.json()
            return error
        }
        const author: Author = await response.json()
        return author;
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}

export async function PostAuthorFromPerson(token: string, request: PostAuthorFromPersonRequest) {
    try {
        const response = await fetch(postAuthorFromPersontRoute, {
            method: "POST",
            headers: {
                'Content-Type': 'application/json',
                'Authorization': token
            },
            body: JSON.stringify(request)
        })
        if (!response.ok) {
            const error: CommonError = await response.json()
            return error
        }
        const author: Author = await response.json()
        return author;
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}