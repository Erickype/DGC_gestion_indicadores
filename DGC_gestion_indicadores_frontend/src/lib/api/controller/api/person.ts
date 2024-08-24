import type { FilterPeopleRequest, FilterPeopleResponse, Person, PostPersonRequest, PutPersonRequest } from "$lib/api/model/api/person";
import { generateErrorFromCommonError, type CommonError } from "$lib/api/model/errors";
import { deletePersonRoute, getPeopleRoute, postFilterPeopleRoute, postPersonRoute, putPersonRoute } from "$lib/api/routes/api/person";
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

export async function PostFilterPeople(token: string, filterPeopleRequest: FilterPeopleRequest): Promise<FilterPeopleResponse[] | CommonError> {
    try {
        const response = await fetch(postFilterPeopleRoute, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': token
            },
            body: JSON.stringify(filterPeopleRequest)
        })
        if (!response.ok) {
            const error: CommonError = await response.json()
            return error
        }
        const filterPeopleResponse: FilterPeopleResponse[] = await response.json()
        return filterPeopleResponse

    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}

export async function PostPerson(person: PostPersonRequest, token: string) {
    return await fetch(postPersonRoute, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': token
        },
        body: JSON.stringify(person)
    });
}

export async function PutPerson(person: PutPersonRequest, token: string) {
    return await fetch(putPersonRoute + person.ID.toString(), {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': token
        },
        body: JSON.stringify(person)
    });
}

export async function DeletePerson(id: string, token: string) {
    return await fetch(deletePersonRoute + id, {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': token
        }
    })
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