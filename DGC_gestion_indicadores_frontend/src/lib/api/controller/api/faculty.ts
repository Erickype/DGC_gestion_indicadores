import { getFacultiesRoute } from "$lib/api/routes/api/faculty";
import type { Message } from "$lib/components/combobox/combobox";
import type { Faculty } from "$lib/api/model/api/faculty";
import { generateErrorFromCommonError, type CommonError } from "$lib/api/model/errors";
import { generateCommonErrorFromFetchError } from "$lib/utils";

export async function GetFaculties(token: string) {
    try {
        const response = await fetch(getFacultiesRoute, {
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
        const faculties: Faculty[] = await response.json()
        return faculties
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}


export async function LoadFacultiesWithComboMessages(token: string) {
    const response = await GetFaculties(token);
    if ((response as CommonError).status) {
        throw generateErrorFromCommonError(response as CommonError)
    }
    const faculties = response as Faculty[];
    let messages: Message[] = []
    messages = messages.concat(
        faculties.map((faculty) => ({
            value: faculty.ID,
            label: faculty.abbreviation,
        }))
    );

    return {
        faculties,
        messages
    }
}