import type { Faculty } from "$lib/api/model/api/faculty";
import type { CommonError } from "$lib/api/model/errors";
import type { Message } from "../combobox/combobox";
import { goto } from "$app/navigation";

export async function fetchFaculties(): Promise<Faculty[]> {
    const url = `/api/faculty`;
    const response = await fetch(url, {
        method: 'GET',
    });
    if (!response.ok) {
        const errorData = (await response.json()) as CommonError;
        if (response.status === 401) {
            throw goto('/');
        }
        throw errorData;
    }
    return response.json();
}

export function GenerateComboMessagesFromFaculties(faculties: Faculty[]): Message[] {
    let messages: Message[] = []
    messages = messages.concat(
        faculties.map((faculty) => ({
            value: faculty.ID,
            label: faculty.abbreviation,
        }))
    );
    return messages
}