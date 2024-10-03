
import { generateCommonErrorFromFetchError } from "$lib/utils";
import type { CommonError } from "$lib/api/model/errors";
import { postBookOrChaptersProductionListsAuthorRoute } from "$lib/api/routes/api/indicatorsInformation/booksOrChaptersProductionListsAuthor";
import type { PostBooksOrChaptersProductionListsAuthorCareersRequest } from "$lib/api/model/api/indicatorsInformation/booksOrChaptersProductionListsAuthor";

export async function PostBooksOrChaptersProductionListsAuthorCareers(token: string, request: PostBooksOrChaptersProductionListsAuthorCareersRequest) {
    try {
        const response = await fetch(postBookOrChaptersProductionListsAuthorRoute, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': token
            },
            body: JSON.stringify(request)
        });

        if (!response.ok) {
            const error: CommonError = await response.json()
            return error
        }
        const postBooksOrChaptersProductionListsAuthorCareersRequest: PostBooksOrChaptersProductionListsAuthorCareersRequest[] = await response.json()
        return postBooksOrChaptersProductionListsAuthorCareersRequest
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}