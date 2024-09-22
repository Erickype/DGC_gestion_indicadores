import type { FilterDetailedFieldRequest, FilterDetailedFieldResponse } from "$lib/api/model/api/knowledgeFields/detailedFields"
import { postFilterDetailedFieldsRoute } from "$lib/api/routes/api/knowledgeFields/detailedFields"

import { generateCommonErrorFromFetchError } from "$lib/utils"
import type { CommonError } from "$lib/api/model/errors"

export async function PostFilterDetailedFields(token: string, filterDetailedFieldRequest: FilterDetailedFieldRequest): Promise<FilterDetailedFieldResponse | CommonError> {
    try {
        const response = await fetch(postFilterDetailedFieldsRoute, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': token
            },
            body: JSON.stringify(filterDetailedFieldRequest)
        })
        if (!response.ok) {
            const error: CommonError = await response.json()
            return error
        }
        const filterDetailedFieldResponse: FilterDetailedFieldResponse = await response.json()
        return filterDetailedFieldResponse

    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}