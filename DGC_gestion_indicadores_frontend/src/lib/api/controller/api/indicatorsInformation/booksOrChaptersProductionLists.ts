import { postBookOrChaptersProductionListRoute } from "$lib/api/routes/api/indicatorsInformation/booksOrChaptersProductionLists";

import { generateCommonErrorFromFetchError } from "$lib/utils";
import type { CommonError } from "$lib/api/model/errors";
import type { BooksOrChaptersProductionList } from "$lib/api/model/api/indicatorsInformation/booksOrChaptersProductionLists";

export async function PostBooksOrChaptersProductionLists(token: string, request: BooksOrChaptersProductionList) {
    try {
        const response = await fetch(postBookOrChaptersProductionListRoute, {
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
        const booksOrChaptersProductionList: BooksOrChaptersProductionList = await response.json()
        return booksOrChaptersProductionList;
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}