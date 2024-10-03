import { postBookOrChaptersProductionListRoute, postFilterBooksOrChaptersProductionListByAcademicPeriodIDRoute, updateBookOrChaptersProductionListRoute } from "$lib/api/routes/api/indicatorsInformation/booksOrChaptersProductionLists";

import { generateCommonErrorFromFetchError } from "$lib/utils";
import type { CommonError } from "$lib/api/model/errors";
import type { BooksOrChaptersProductionList, FilterBooksOrChaptersProductionListsByAcademicPeriodRequest, FilterBooksOrChaptersProductionListsByAcademicPeriodResponse } from "$lib/api/model/api/indicatorsInformation/booksOrChaptersProductionLists";

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

export async function UpdateBooksOrChaptersProductionList(token: string, request: BooksOrChaptersProductionList) {
    try {
        const response = await fetch(updateBookOrChaptersProductionListRoute + request.ID, {
            method: "PUT",
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

export async function PostFilterBooksOrChaptersProductionListsByAcademicPeriodID(token: string, request: FilterBooksOrChaptersProductionListsByAcademicPeriodRequest) {
    try {
        const response = await fetch(postFilterBooksOrChaptersProductionListByAcademicPeriodIDRoute, {
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
        const filterBooksOrChaptersProductionListsByAcademicPeriodResponse: FilterBooksOrChaptersProductionListsByAcademicPeriodResponse[] = await response.json()
        return filterBooksOrChaptersProductionListsByAcademicPeriodResponse
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}