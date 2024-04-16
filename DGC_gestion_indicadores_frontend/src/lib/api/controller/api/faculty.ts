import type { Faculty } from "$lib/api/model/api/faculty";
import { getFacultiesRoute } from "$lib/api/routes/api/faculty";
import type { Message } from "$lib/components/combobox/combobox";

export async function GetFaculties(token: string) {
    return await fetch(getFacultiesRoute, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': token
        }
    });
}


export async function LoadFacultiesWithComboMessages(token: string) {
    const res = await GetFaculties(token);
    if (!res.ok) {
        throw new Error('Fallo cargando facultades.');
    }
    const faculties: Faculty[] = await res.json();
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