import type { FilterPeopleRequest, FilterPeopleResponse, Person, PersonWithRoles, PostPersonRequest, PostPersonWithRolesRequest, PutPersonRequest, UpdatePersonWithRolesRequest } from "$lib/api/model/api/person";
import { deletePersonRoute, getPeopleRoute, postFilterPeopleRoute, postPersonRoute, postPersonWithRolesRoute, putPersonRoute, updatePersonWithRolesRoute } from "$lib/api/routes/api/person";

import { generateErrorFromCommonError, type CommonError } from "$lib/api/model/errors";
import type { Message } from "$lib/components/combobox/combobox";
import { generateCommonErrorFromFetchError } from "$lib/utils";

export async function GetPeople(token: string): Promise<PersonWithRoles[] | CommonError> {
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
        const periods: PersonWithRoles[] = await response.json()
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

export async function PostPersonWithRoles(postPersonWithRolesRequest: PostPersonWithRolesRequest, token: string) {
    try {
        const response = await fetch(postPersonWithRolesRoute, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': token
            },
            body: JSON.stringify(postPersonWithRolesRequest)
        });
        if (!response.ok) {
            const error: CommonError = await response.json()
            return error
        }
        const postPersonWithRolesResponse: PostPersonWithRolesRequest = await response.json()

        return postPersonWithRolesResponse
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}

export async function UpdatePersonWithRoles(updatePersonWithRolesRequest: UpdatePersonWithRolesRequest, token: string) {
    try {
        const response = await fetch(updatePersonWithRolesRoute, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': token
            },
            body: JSON.stringify(updatePersonWithRolesRequest)
        });
        if (!response.ok) {
            const error: CommonError = await response.json()
            return error
        }
        const postPersonWithRolesResponse: UpdatePersonWithRolesRequest = await response.json()

        return postPersonWithRolesResponse
    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
}

export async function PutPerson(person: PutPersonRequest, token: string) {
    try {
        const response = await fetch(putPersonRoute + person.ID.toString(), {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': token
            },
            body: JSON.stringify(person)
        });
        if (!response.ok) {
            const error: CommonError = await response.json()
            return error
        }
        const putPersonResponse: Person = await response.json()

        return putPersonResponse

    } catch (error) {
        return generateCommonErrorFromFetchError(error)
    }
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
    const people = response as PersonWithRoles[]
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

export function GenerateComboMessagesFromPeople(people: PersonWithRoles[]): Message[] {
    let messages: Message[] = []
    messages = messages.concat(
        people.map((person) => ({
            value: person.ID,
            label: person.identity + " " + person.name + " " + person.lastname,
        }))
    );
    return messages
}