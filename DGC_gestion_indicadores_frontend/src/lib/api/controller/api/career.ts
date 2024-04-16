import type { Career } from "$lib/api/model/api/career";
import { getCareersRoute } from "$lib/api/routes/api/career";
import type { Message } from "$lib/components/combobox/combobox";

export async function GetCareers(token: string) {
    return await fetch(getCareersRoute, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': token
        }
    });
}


export async function LoadCareersWithComboMessages(token: string) {
    const res = await GetCareers(token);
    if (!res.ok) {
        throw new Error('Fallo cargando carreras.');
    }
    const careers: Career[] = await res.json();
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