import type { Person } from "$lib/api/model/api/person";
import { generateErrorFromCommonError, type CommonError } from "$lib/api/model/errors";
import { getPeopleRoute } from "$lib/api/routes/api/person";
import type { Message } from "$lib/components/combobox/combobox";
import { generateCommonErrorFromFetchError } from "$lib/utils";

export async function GetPeople(token: string): Promise<Person[] | CommonError> {
    try {
        const response = await fetch(getPeopleRoute, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': token
            }
        })
        if (!response.ok) {
            const error: CommonError = await response.json()
            return error
        }
        const periods: Person[] = await response.json()
        return periods

    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}


export async function LoadPeopleWithComboMessages(token: string) {
    const response = await GetPeople(token);
    if ((response as CommonError).status) {
        throw generateErrorFromCommonError(response as CommonError)
    }
    const people = response as Person[]
    let messages: Message[] = []
    messages = messages.concat(
        people.map((person) => ({
            value: person.ID,
            label: person.identity,
        }))
    );

    return {
        people,
        messages
    }
}