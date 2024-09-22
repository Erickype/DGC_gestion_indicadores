import type { DetailedFilterJoined, FilterDetailedFieldRequest, FilterDetailedFieldResponse } from "$lib/api/model/api/knowledgeFields/detailedFields"
import { postFilterDetailedFieldsRoute } from "$lib/api/routes/api/knowledgeFields/detailedFields"

import { generateCommonErrorFromFetchError } from "$lib/utils"
import type { CommonError } from "$lib/api/model/errors"
import type { Message } from "$lib/components/combobox/combobox"

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

export function GenerateComboMessagesFromDetailedFilterJoined(detailedFields: DetailedFilterJoined[]): Message[] {
    let messages: Message[] = []
    messages = messages.concat(
        detailedFields.map((detailedField) => ({
            value: detailedField.detailed_field_id,
            label: detailedField.detailed_field,
        }))
    );
    return messages
}