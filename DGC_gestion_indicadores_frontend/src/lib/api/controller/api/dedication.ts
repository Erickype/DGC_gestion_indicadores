import type { Dedication } from "$lib/api/model/api/dedication";
import { getDedicationsRoute } from "$lib/api/routes/api/dedications";
import type { Message } from "$lib/components/combobox/combobox";

export async function GetDedications(token: string) {
    return await fetch(getDedicationsRoute, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': token
        }
    });
}


export async function LoadDedicationsWithComboMessages(token: string) {
    const res = await GetDedications(token);
    if (!res.ok) {
        throw new Error('Fallo cargando dedicaciones.');
    }
    const dedications: Dedication[] = await res.json();
    let messages: Message[] = []
    messages = messages.concat(
        dedications.map((Dedication) => ({
            value: Dedication.ID,
            label: Dedication.name,
        }))
    );

    return {
        dedications,
        messages
    }
}