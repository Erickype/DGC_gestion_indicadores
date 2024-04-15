import type { Person } from "$lib/api/model/api/person";
import { getPeopleRoute } from "$lib/api/routes/api/person";
import type { Message } from "$lib/components/combobox/combobox";

export async function GetPeople(token: string) {
    return await fetch(getPeopleRoute, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': token
        }
    });
}


export async function LoadPeopleWithComboMessages(token: string) {
    const res = await GetPeople(token);
    if (!res.ok) {
        throw new Error('Fallo cargando personas.');
    }
    const people: Person[] = await res.json();
    let messages: Message[] = []
    messages = messages.concat(
        people.map((person) => ({
            value: person.identity,
            label: person.identity,
        }))
    );

    return {
        people,
        messages
    }
}