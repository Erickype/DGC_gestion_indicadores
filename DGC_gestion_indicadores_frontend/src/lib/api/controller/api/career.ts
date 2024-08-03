import type { Career } from "$lib/api/model/api/career";
import { generateErrorFromCommonError, type CommonError } from "$lib/api/model/errors";
import { getCareersRoute } from "$lib/api/routes/api/career";
import type { Message } from "$lib/components/combobox/combobox";
import { generateCommonErrorFromFetchError } from "$lib/utils";

export async function GetCareers(token: string) {
    try {
        const response = await fetch(getCareersRoute, {
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
        const careers: Career[] = await response.json()
        return careers
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}


export async function LoadCareersWithComboMessages(token: string) {
    const response = await GetCareers(token);
    if ((response as CommonError).status) {
        throw generateErrorFromCommonError(response as CommonError)
    }
    const careers: Career[] = response as Career[];
    let messages: Message[] = []
    messages = messages.concat(
        careers.map((career) => ({
            value: career.ID,
            label: career.abbreviation,
        }))
    );

    return {
        careers,
        messages
    }
}