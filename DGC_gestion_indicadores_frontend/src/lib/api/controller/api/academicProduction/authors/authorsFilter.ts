import type { FilterAuthorsRequest, FilterAuthorsResponse } from "$lib/api/model/api/academicProduction/authors/authorsFilter"
import { postFilterAuthorsRoute } from "$lib/api/routes/api/academicProduction/authors/author"
import { generateCommonErrorFromFetchError } from "$lib/utils"
import type { CommonError } from "$lib/api/model/errors"

export async function PostFilterAuthors(token: string, filterAuthorsRequest: FilterAuthorsRequest): Promise<FilterAuthorsResponse | CommonError> {
    try {
        const response = await fetch(postFilterAuthorsRoute, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': token
            },
            body: JSON.stringify(filterAuthorsRequest)
        })
        if (!response.ok) {
            const error: CommonError = await response.json()
            return error
        }
        const filterAuthorsResponse: FilterAuthorsResponse = await response.json()
        return filterAuthorsResponse

    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}