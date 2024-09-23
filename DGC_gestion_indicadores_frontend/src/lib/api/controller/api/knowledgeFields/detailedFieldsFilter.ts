import type { DetailedFieldJoined, FilterDetailedFieldRequest, FilterDetailedFieldResponse } from "$lib/api/model/api/knowledgeFields/detailedFields"
import { getDetailedFieldJoinedByDetailedFieldIDRoute, postFilterDetailedFieldsRoute } from "$lib/api/routes/api/knowledgeFields/detailedFields"

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

export async function GetDetailedFieldJoinedByDetailedFieldID(token: string, detailedFieldID: string): Promise<DetailedFieldJoined | CommonError> {
    try {
        const response = await fetch(getDetailedFieldJoinedByDetailedFieldIDRoute + detailedFieldID, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': token
            },
        })
        if (!response.ok) {
            const error: CommonError = await response.json()
            return error
        }
        const detailedField: DetailedFieldJoined = await response.json()
        return detailedField

    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}

export function GenerateComboMessagesFromDetailedFieldJoined(detailedFields: DetailedFieldJoined[]): Message[] {
    let messages: Message[] = []
    messages = messages.concat(
        detailedFields.map((detailedField) => ({
            value: detailedField.detailed_field_id,
            label: detailedField.detailed_field,
        }))
    );
    return messages
}