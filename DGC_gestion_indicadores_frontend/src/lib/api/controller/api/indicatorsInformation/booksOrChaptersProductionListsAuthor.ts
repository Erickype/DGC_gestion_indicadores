
import { getBooksOrChaptersProductionListsAuthorsJoinedByBooksOrChaptersProductionListIDRoute, postBookOrChaptersProductionListsAuthorRoute, updateBookOrChaptersProductionListsAuthorRoute } from "$lib/api/routes/api/indicatorsInformation/booksOrChaptersProductionListsAuthor";
import type { BooksOrChaptersProductionListsAuthorsCareersJoined, PostBooksOrChaptersProductionListsAuthorCareersRequest, UpdateBooksOrChaptersProductionListsAuthorCareersRequest } from "$lib/api/model/api/indicatorsInformation/booksOrChaptersProductionListsAuthor";

import { generateCommonErrorFromFetchError } from "$lib/utils";
import type { CommonError } from "$lib/api/model/errors";

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

export async function UpdateBooksOrChaptersProductionListsAuthorCareers(token: string, request: UpdateBooksOrChaptersProductionListsAuthorCareersRequest) {
    try {
        const response = await fetch(updateBookOrChaptersProductionListsAuthorRoute, {
            method: 'PUT',
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
        const updateBooksOrChaptersProductionListsAuthorCareersRequest: UpdateBooksOrChaptersProductionListsAuthorCareersRequest[] = await response.json()
        return updateBooksOrChaptersProductionListsAuthorCareersRequest
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}

export async function GetBooksOrChaptersProductionListsAuthorsJoinedByBooksOrChaptersProductionListID(token: string, booksOrChaptersProductionListID: string): Promise<BooksOrChaptersProductionListsAuthorsCareersJoined[] | CommonError> {
    try {
        const response = await fetch(getBooksOrChaptersProductionListsAuthorsJoinedByBooksOrChaptersProductionListIDRoute + booksOrChaptersProductionListID, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': token
            },
        });

        if (!response.ok) {
            const error: CommonError = await response.json()
            return error
        }
        const booksOrChaptersProductionListsByAcademicPeriodJoined: BooksOrChaptersProductionListsAuthorsCareersJoined[] = await response.json()
        return booksOrChaptersProductionListsByAcademicPeriodJoined
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}