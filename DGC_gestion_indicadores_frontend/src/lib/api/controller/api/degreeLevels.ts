import type { DegreeLevel } from "$lib/api/model/api/degreeLevels";
import { generateErrorFromCommonError, type CommonError } from "$lib/api/model/errors";
import { getDegreeLevelsRoute } from "$lib/api/routes/api/degreeLevel";
import type { Message } from "$lib/components/combobox/combobox";
import { generateCommonErrorFromFetchError } from "$lib/utils";

export async function GetDegreeLevels(token: string) {
    try {
        const response = await fetch(getDegreeLevelsRoute, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': token
            }
        });

        if (!response.ok) {
            const error: CommonError = await response.json()
            return error
        }
        const degreeLevels: DegreeLevel[] = await response.json()
        return degreeLevels
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}


export async function LoadDegreeLevelsWithComboMessages(token: string) {
    const response = await GetDegreeLevels(token);
    if ((response as CommonError).status) {
        throw generateErrorFromCommonError(response as CommonError)
    }
    const degreeLevels = response as DegreeLevel[];
    let messages: Message[] = []
    messages = messages.concat(
        degreeLevels.map((faculty) => ({
            value: faculty.ID,
            label: faculty.abbreviation,
        }))
    );

    return {
        degreeLevels,
        messages
    }
}