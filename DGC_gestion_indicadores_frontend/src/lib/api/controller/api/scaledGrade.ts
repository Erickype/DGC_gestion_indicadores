import type { ScaledGrade } from "$lib/api/model/api/scaledGrade";
import { getScaledGradesRoute } from "$lib/api/routes/api/scaledGrade";
import type { Message } from "$lib/components/combobox/combobox";

export async function GetScaledGrades(token: string) {
    return await fetch(getScaledGradesRoute, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': token
        }
    });
}


export async function LoadScaledGradesWithComboMessages(token: string) {
    const res = await GetScaledGrades(token);
    if (!res.ok) {
        throw new Error('Fallo cargando grados escalafonados.');
    }
    const scaledGrades: ScaledGrade[] = await res.json();
    let messages: Message[] = []
    messages = messages.concat(
        scaledGrades.map((scaledGrade) => ({
            value: scaledGrade.ID,
            label: scaledGrade.abbreviation,
        }))
    );

    return {
        scaledGrades,
        messages
    }
}